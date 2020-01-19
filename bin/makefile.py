#!/usr/bin/env python3
import os
import re
import subprocess
import sys
import shutil
import json

made = set()


def consume_makefile(method, *arg, **kwarg):
    if callable(method):
        i = str(method)
        if i not in made:
            made.add(i)
            method(*arg, **kwarg)


# return function in object scope
def sq(obj, attr):
    # raise if not exists
    def __get_object_attr(*_): return getattr(obj, attr)

    return __get_object_attr


# object requirements
def sqs(obj, *attrs): return map(lambda attr: sqs(obj, attr), attrs)


# return function in object scope
def oq(attr):
    # raise if not exists
    def __get_object_attr(obj, *_): return getattr(obj, attr)

    return __get_object_attr


# object requirements
def oqs(*attrs): return map(oq, attrs)


def require(*targets):
    def c(f):
        def wrapper(*args, **kwargs):
            for target in targets:
                target = (getattr(target, '__name__', None) == '__get_object_attr' and target(*args,
                                                                                              **kwargs)) or target
                consume_makefile(target, *args, **kwargs)
            return f(*args, **kwargs)

        return wrapper

    return c


def require_cls(*target): return require(*oqs(*target))


def pipe(cmd, cwd=None):
    p, line, line2 = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE,
                                      cwd=cwd), ' ', ' '
    while len(line) != 0 or len(line2) != 0:
        line = p.stdout.readline()
        line2 = p.stderr.readline()
        if len(line) != 0:
            print(line.decode('utf-8').strip())
        if len(line2) != 0:
            print(line2.decode('utf-8').strip())
    code = p.wait()
    if code != 0:
        print('exit with %d' % code)


def entry(self, args):
    getattr(self, args[0])(*args[1:]) if len(args) > 0 else self.all()


class Makefile:
    current_path = os.path.dirname(os.path.realpath(__file__))
    config_file_name = 'config.toml'
    context = dict()

    compose_template_file = 'docker-compose.template.yml'
    compose_file = 'docker-compose.yml'
    silent = False

    @classmethod
    def hello(cls, *_):
        print('minimum builder')

    @classmethod
    def pipe(cls, cmd, *_):
        if not cls.silent: print(cmd)
        pipe(cmd)

    @classmethod
    def generate(cls, path='./', match=None, *_):
        match = cls._gen_match(match)
        for file in os.listdir(path):
            file = os.path.join(path, file)
            if os.path.isdir(file):
                cls.generate(file, match)
            if os.path.isfile(file) and match.match(file):
                cls.pipe('go generate %s' % file)


    @classmethod
    @require_cls('read_context')
    def image(cls, *_):
        cls.pipe('docker build --tag %s .' % cls.context['node-name'])


    @staticmethod
    def _gen_match(match):
        if match is None:
            match = re.compile(r'^.*\.go$')
        if isinstance(match, str):
            match = re.compile(match)
        return match

    @classmethod
    @require_cls('template')
    def up(cls, *_):
        pipe('docker-compose -f %s up' % (cls.compose_file))

    @classmethod
    @require_cls('template')
    def down(cls, *_):
        pipe('docker-compose -f %s down' % (cls.compose_file))

    @classmethod
    @require_cls('template')
    def start(cls, *_):
        pipe('docker-compose -f %s start' % (cls.compose_file))

    @classmethod
    @require_cls('template')
    def stop(cls, *_):
        pipe('docker-compose -f %s stop' % (cls.compose_file))

    @classmethod
    @require_cls('read_context')
    def template(cls, *_):
        with open(cls.compose_template_file) as f:
            s = f.read().replace('{{redis-root-password}}', cls.context['redis-root-password'])
            s = s.replace('{{mysql-root-password}}', cls.context['mysql-root-password'])
            s = s.replace('{{conf-path}}', cls.context['conf-path'])
            s = s.replace('{{logs-path}}', cls.context['logs-path'])
            s = s.replace('{{data-path}}', cls.context['data-path'])
            s = s.replace('{{goods-picture-path}}', cls.context['goods-picture-path'])
            s = s.replace('{{needs-picture-path}}', cls.context['needs-picture-path'])
            s = s.replace('{{goods-picture-path-target}}', cls.context['goods-picture-path-target'])
            s = s.replace('{{needs-picture-path-target}}', cls.context['needs-picture-path-target'])
            s = s.replace('{{node-name}}', cls.context['node-name'])
            s = s.replace('{{instance-name}}', cls.context['instance-name'])
            s = s.replace('{{target-port}}', str(cls.context['target-port']))
            s = s.replace('{{config-file}}', cls.config_file)
            s = s.replace('{{config-file-target}}', cls.config_file_target)
            s = s.replace('{{mysql-norm-password}}', cls.context['mysql-norm-password'])
            with open(cls.compose_file, 'w') as o:
                o.write(s)

    @classmethod
    def read_context(cls, *_):
        with open('.minimum-lib-env.json', 'r', encoding='utf-8') as f:
            c = f.read().encode('utf-8')
            cls.context = json.loads('{}' if len(c) == 0 else c)
            cls.context['node-name'] = cls.context.get('node-name', 'myriaddreamin/minimum-blog-backend:alpine')
            cls.context['instance-name'] = cls.context.get('instance-name', 'backend')

            cls.context['redis-root-password'] = cls.context.get('redis-root-password', '12345678')
            cls.context['mysql-root-password'] = cls.context.get('mysql-root-password', '12345678')
            cls.context['mysql-norm-password'] = cls.context.get('mysql-norm-password', '12345678')
            
            cls.context['conf-path'] = os.path.join(cls.current_path, cls.context.get('conf-path', 'testdb/conf'))
            print('conf-path', cls.context['conf-path'])
            
            cls.context['logs-path'] = os.path.join(cls.current_path, cls.context.get('logs-path', 'testdb/logs'))
            print('logs-path', cls.context['logs-path'])
            cls.context['data-path'] = os.path.join(cls.current_path, cls.context.get('data-path', 'testdb/data'))
            print('data-path', cls.context['data-path'])
            cls.context['goods-picture-path'] =\
                os.path.join(cls.current_path, cls.context.get('goods-picture-path', 'testdb/goods-picture'))
            cls.context['needs-picture-path'] =\
                os.path.join(cls.current_path, cls.context.get('needs-picture-path', 'testdb/needs-picture'))
            cls.context['goods-picture-path-target'] =\
                os.path.join(cls.current_path, cls.context.get('goods-picture-path-target', '/goods-picture'))
            cls.context['needs-picture-path-target'] =\
                os.path.join(cls.current_path, cls.context.get('needs-picture-path-target', '/needs-picture'))

            cls.context['target-port'] = cls.context.get('target-port', 23335)

            cls.config_file_name = cls.context.get('config-file-name', 'config.toml')
            cls.config_file = os.path.join(cls.current_path, cls.config_file_name)
            cls.config_file_target = os.path.join('/', cls.config_file_name)

    @classmethod
    @require_cls('read_context')
    def clean(cls, *_):
        shutil.rmtree('testdb')

    @classmethod
    @require_cls('up')
    def all(cls, *_):
        pass



if __name__ == '__main__':
    entry(Makefile, sys.argv[1:])
