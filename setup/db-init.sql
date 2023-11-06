DROP TABLE IF EXISTS public.co2_inference;

CREATE TABLE public.spatial_forecast
(
  geojson               text,
  pollutant_val         double precision,
  observationdatetime   timestamp without time zone
);

CREATE TABLE public.forecast
(
    id                      text,
    device_id               text,
    observationdatetime     timestamp with time zone,
    air_temperature         float,
    atmospheric_pressure    float,
    relative_humidity       float,
    pm10                    float,
    pm2p5                   float,
    so2                     float,
    no2                     float,
    co                      float,
    co2                     float,
    air_quality_index       float,
    location_coordinates    point
);

CREATE TABLE public.aqm_actual (
  geojson_id            int,
  observationdatetime   timestamp without time zone,
  pollutant_val         double precision
);

CREATE TABLE public.aqm_forecast (
  geojson_id            int,
  observationdatetime   timestamp without time zone,
  pollutant_val         double precision
);

CREATE TABLE public.aqm_geojson (
  geojson               text,
  geojson_id            int
);