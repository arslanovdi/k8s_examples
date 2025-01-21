{{/*
This file contains helper templates and functions that can be used across multiple templates in the current chart.
*/}}

{{/* Define a helper function to generate the full name of the application */}}
{{- define "stateful-todo.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name -}}
{{- end -}}

{{/* Define a helper function to get the chart name */}}
{{- define "stateful-todo.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version -}}
{{- end -}}