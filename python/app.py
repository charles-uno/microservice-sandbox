#!/usr/bin/env python

import connexion

app = connexion.App(__name__)

app.add_api('swagger.yaml')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080, debug=True)
