## Release 1.3.0

BREAKING CHANGES:

#### provider

* `url` was deleted. Instead, added `env` which defines the Coralogix environment. Also, can be set by environment
  variable `CORALOGIX_ENV` instead. Can be one of the following - `[APAC1, APAC2, EUROPE1, EUROPE2, USA1]`.
* `timeout` was deleted. Will be defined by a different timeout for each resource (internally).

#### resource/rule

* The resource `rule` was deleted. Use `rule_group` with single inner `rule` instead.

#### resource/coralogix_rules_group

* `enabled` changed to `active`.
* `rule_matcher` was deleted and `severity`, `applicationName` and `subsystemName` were moved out to previous level as
  separated lists of `severities`, `applications` and `subsystems`.
* `rules` was deleted and replaced by `rule_subgroups` (every `rule-subgroup` is list of `rule`s with 'or' (||)
  operation between).
* `rules.*.group` was deleted and replaced by `rule_subgroups.*.rules`.
* `rules.*.group.*.type` was deleted. Instead, every `rule` inside `rules` (`rule_subgroups.*.rules.*`) can be one of
    - `[parse, block, json_extract, replace, extract_timestamp, remove_fields, json_stringify, extract]`.
* All the other parameters inside `rules.*.group.*` were moved to the specific rule type schemas
  inside `rule_subgroups.*.rules.*`. Any specific rule type schemas contain only its relevant fields.

#### resource/alert

* `severity` changed to `alert_severity` and can be one of the following - `[Info, Warning, Critical, Error]`.
* `type` was removed. Instead, every alert must contain exactly one of
    - `[standard, ratio, new_value, unique_count, time_relative, metric, tracing, flow]`.
* `schedule` changed to `scheduling`.
* `schedule.*.days` changed to `scheduling.*.days_enabled` and can be one of the following
    - `[Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]`.
* `schedule.*.start` changed to `scheduling.*.start_time`.
* `schedule.*.end` changed to `scheduling.*.end_time`.
* All the other parameters inside `alert` were moved to the specific alert type schemas inside `alert`. Any specific
  alert type schemas contain only their relevant fields.

FEATURES:

* **New Resource:** `logs2metric`

IMPROVEMENTS:

#### provider

* `api_key` can be declared as environment variable `CORALOGIX_API_KEY` instead of to terraform configs.
* Add Acceptance Tests.
* Added retrying mechanism.

#### resource/coralogix_rules_group

* Add Acceptance Tests.
* Moved to Coralogix grpc endpoint.

#### resource/alert

* Add Acceptance Tests.
* Moved to Coralogix grpc endpoint.

## Release 1.3.27

BREAKING CHANGES:

#### resource/alert

* `webhook_ids` was changed to `webhooks`.

## Release 1.3.29

BREAKING CHANGES:

#### resource/alert

* `alert_severity` was changed to `severity`.
* `manage_undetected_values.disable_triggering_on_undetected_values` was omitted.
  Instead, it's possible to set `manage_undetected_values.enable_triggering_on_undetected_values = false`
  (`manage_undetected_values.auto_retire_ratio` is not allowed in that case).

## Release 1.3.31

BREAKING CHANGES:

#### resource/alert

* `categories` ,`classes`, `computers`, `ip_addresses`, `methods` and `search_query` are not supported
filters for tracing alert, Therefore they were deleted from the tracing-alert scheme.
* `applications`,`severities` and `subsystems` filters have currently different format in tracing-alert (`field_filters`),
Therefore they were deleted from the tracing-alert scheme.

## Release 1.4.0

BREAKING CHANGES:

#### resource/alert

* `meta_labels` was changed to key-value map. (e.g. - `meta_labels {key = "alert_type" value = "security"} meta_labels {key = "security_severity" value = "high"}` => `meta_labels = {alert_type = "security" security_severity = "high" }`).
([example-usage](docs/resources/alert.md#standard-alert)).
* `scheduling.time_frames` was changed to `time_frame`.
* `standard.occurrences_threshold` and `tracing.occurrences_threshold` were changed to `threshold`.
* `ratio.queries_ratio` was changed to `ratio_threshold`.
* `notification` was changed to list of `notifications_group` and have entire different schema ([nested-schema-for-notifications](docs/resources/alert.md#nested-schema-for-notifications_group)).
* `notification.ignore_infinity` was moved to `ratio.condition.ignore_infinity` and `time_relative.condition.ignore_infinity`.
* `notification.notify_every_min` was changed to `notifications_group.notification.retriggering_period_minutes`.
* `notification.on_trigger_and_resolved` (boolean) was changed to `notifications_group.notification.notify_on` (string).
* `notification.recipients.webhook_id` replaced with `notifications_group.notification.integration_id` and should contain the integration's (webhook's) id instead of the integration's name.
* flow-alert's (`flow`) schema was fixed. Any earlier version contained wrong schema of flow-alert. ([nested-schema-for-flow](docs/resources/alert.md#nested-schema-for-flow)).
* `tracing.field_filters` was removed, and `tracing.applications`, `tracing.applications` and `tracing.services` were added instead.
* `tracing.tag_filters` was changed to `tracing.tag_filter` and contains only `field` and `values`.
* `tracing.tag_filter.values`, `tracing.applications`, `tracing.applications` and `tracing.services` have the same format as the other alerts' filters. ([example-usage](docs/resources/alert.md#tracing-alert)).
* `tracing.latency_threshold_ms` was changed to `latency_threshold_milliseconds`.

## Release 1.4.4

BREAKING CHANGES:

#### resource/alert

* `notifications_group.group_by_fields` was changed from _TypeSet_ (doesn't keep order of declaration) to _TypeList_ (keeps order of declaration). This change can cause to diffs in state.

## Release 1.5.0

BREAKING CHANGES:

#### resource/events2metric (~~logs2metric~~)

* resource and data-source name _logs2metric_ was changed to _events2metric_ and contains `logs_query` and `span_query` option.

## Release 1.5.2

FEATURES:

#### resource/events2metric
* Adding [aggregations](docs/resources/events2metric.md#nested-schema-for-metric_fieldsaggregations) option to `metric_fields`.

## Release 1.5.3

BREAKING CHANGES:

#### resource/tco_policy

* `severities` is now required.
* `order` is now required.

BUG FIXING:

#### resource/tco_policy
* Fixing [TF doesn't detect manually deleted resources](https://coralogix-dev.slack.com/archives/C04CV0JG36H/p1683820712917899)
* Fixing [Order of policies can't be manged by TF](https://coralogix-dev.slack.com/archives/C04CV0JG36H/p1681995853325159)

FEATURES:

* Adding [Custom Domain option](docs/index.md#private-domains)

DEVELOPERS:

* go version was update to 1.20.x

## Release 1.5.4

FEATURES:

* Adding [tco_policy_override](docs/resources/tco_policy_override.md) resource and data source.

## Release 1.5.5

BREAKING CHANGES:

#### resource/recording_rules_group
* Deprecated, and replaced with [recording_rules_groups_set](./docs/resources/recording_rules_groups_set.md) .

## Release 1.5.6

#### resource/tco_policy
* Fixing [TF doesn't detect manually deleted resources](https://coralogix-dev.slack.com/archives/C04CV0JG36H/p1683820712917899)

## Release 1.5.7

BUG FIXING:

#### resource/alert
* Fixing - received an error when updating promql alert condition from less_than to more_than

## Release 1.5.8

BREAKING CHANGES:
#### resource/coralogix_dashboard
* for all the new line chart widgets - `query`, `series_name_template`, `series_count_limit`, `unit` and `scale_type` fields were deprecated. They will be part of `query_definition` now.
* all lists of objects names' 's' suffix was removed (e.g. - `widgets` => `widget`).

## Release 1.5.9

BUG FIXING:
#### resource/coralogix_alert 
* Fixing - tracing alert with `tracing.tag_filter` and `tracing.applications`/`tracing.services`/`tracing.services` filters doesn't work for 'equal' operator.

FEATURES:
#### resource/coralogix_alert
* Adding support for 'notEquals' operator for `tracing.tag_filter` and `tracing.applications`/`tracing.services`/`tracing.services` filters.

## Release 1.6.0

BREAKING CHANGES:
#### resource/coralogix_events2metric
**Please note** - this version contains a [State Upgrader](https://developer.hashicorp.com/terraform/plugin/framework/migrating/resources/state-upgrade#framework). It will upgrade the state to the new schema. Please make sure to back up your state before upgrading.
* `logs_query` type was changed from `Block List, Max: 1` to `Attributes`.
* `spans_query` type was changed from `Block List, Max: 1` to `Attributes`.
* `metric_fields` type was changed from `Block Set` to `Attributes Map`, and `metric_fields`'s `target_base_metric_name` was changed to the map's key. e.g. - `metric_fields {target_base_metric_name = "metric_name" ...}` => `metric_fields = {"metric_name" = {...}}`.
* `metric_fields`'s `aggregations` type was changed from `Block List, Max: 1` to `Attributes`.
* All `aggregations`'s fields (`avg`, `count`, `histogram`, `max`, `min`, `samples`, `sum`) types where changed from `Block List, Max: 1` `Attributes`.
* `metric_labels` type was changed from `Block Set` to `Attributes Map`, and `metric_labels`'s `target_label_name` was changed to the map's key. e.g. - `metric_labels {target_label_name = "label_name" ...}` => `metric_labels = {"label_name" = {...}}`.

BUG FIXING:
#### resource/coralogix_events2metric
* Fixing - [`aggregations` cannot be updated after creation](https://github.com/coralogix/terraform-provider-coralogix/issues/115).

## Release 1.6.1

FEATURES:
#### resource/coralogix_alert
* Adding support for `more_than_usual` condition for `metric.promql` alert.

## Release 1.6.2

FEATURES:
#### resource/coralogix_tco_policy
* Adding support for `archive_retention_id`.

## Release 1.6.3
DEVELOPERS:
#### resource/coralogix_actions
* Resource and Data Source were moved to plugin-framework.

## Release 1.6.4
Breaking Changes:
#### resource/coralogix_alert
* `ratio` and `time_relative`'s `condition`'s `ignore_infinity` conflicts with `condition`'s `less_than`.

BUG FIXING:
#### resource/coralogix_tco_policy
* Fixing - getting panic on creation errors.

## Release 1.6.5
INTERNAL CHANGES:
* `tco_policy` and `tco_policy_override` endpoints were changed.