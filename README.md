# background-commons

## Getting started

This repo contains shared logic for both background-jobs and background-processing repositories. The main logic is mostly related to database. That is why this repo also contains the MariaDB helm chart.

Install it as it follows:

```bash
kubectl apply -f mariadb-helm-chart.yaml
```

It will create a StateFulSet into the `mariadb` namespace, plus a service listening in the default port `3306` though a Service (mariadb.mariadb.svc.cluster.local).

## Dependencies

In order to manage the access to the database, I decided to delegate it to a module called [GORP](https://github.com/go-gorp/gorp). It reduces the boilerplating.
