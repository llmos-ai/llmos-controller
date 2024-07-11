{{/*
Expand the name of the chart.
*/}}
{{- define "llmos-controller.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "llmos-controller.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "llmos-controller.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "llmos-controller.labels" -}}
helm.sh/chart: {{ include "llmos-controller.chart" . }}
{{ include "llmos-controller.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Webhook labels
*/}}
{{- define "llmos-controller.webhookLabels" -}}
{{ include "llmos-controller.labels" . }}
app.llmos.ai/webhook: "true"
{{- end }}

{{/*
Selector labels
*/}}
{{- define "llmos-controller.selectorLabels" -}}
app.kubernetes.io/name: {{ include "llmos-controller.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Webhook selector labels
*/}}
{{- define "llmos-controller.webhookSelectorLabels" -}}
{{ include "llmos-controller.selectorLabels" . }}
app.llmos.ai/webhook: "true"
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "llmos-controller.serviceAccountName" -}}
{{- if .Values.controller.serviceAccount.create }}
{{- default (include "llmos-controller.fullname" .) .Values.controller.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.controller.serviceAccount.name }}
{{- end }}
{{- end }}
