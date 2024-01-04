#!  /usr/local/bin/python3

#!  ../../../../flask/bin/python3

################################################################################
#   main.py  -  Jan-3-2024 by aldebap
#
#   Rest API example
################################################################################

from flask import Flask

def main():
    servicePort = 8080

    print("Listening port " + servicePort + "\n")

    #   instantiate and run the API Web Server
    app = Flask(__name__)

    @app.route('/status')
    def index():
        return '{"status": "Server up and running"}'

    app.run()

#   Application's entry point
if "__main__" == __name__:
    main()
