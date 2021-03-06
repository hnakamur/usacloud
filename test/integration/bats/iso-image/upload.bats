#!/usr/bin/env bats

load ${BASE_TEST_DIR}/helpers.bash

DUMMY_ISO_FILE="integration-test-dummy.iso"
DUMMY_ISO_DOWNLOAD_FILE="integration-test-dummy.iso"
DUMMY_ISO_DATA_DIR="integration-test-dummy-iso"
ISO_IMAGE_NAME=${TEST_TARGET_NAME}01

function setup() {
  cleanup_dummy_iso_files
  mkdir "${DUMMY_ISO_DATA_DIR}"
  echo "dummy" > "${DUMMY_ISO_DATA_DIR}/dummy.txt"
  $MK_ISO_CMD -o "${DUMMY_ISO_FILE}" "${DUMMY_ISO_DATA_DIR}"
}

function teardown() {
  cleanup_dummy_iso_files
}

function cleanup_dummy_iso_files {
  rm -rf "${DUMMY_ISO_DATA_DIR}"
  rm -f "${DUMMY_ISO_FILE}"
  rm -f "${DUMMY_ISO_DOWNLOAD_FILE}"
}
@test "Usacloud: should can create and upload iso-image" {
  run usacloud_run iso-image create -y --iso-file "${DUMMY_ISO_FILE}" --name "${ISO_IMAGE_NAME}"

  [ -n "${output}" ]
  [ "${status}" -eq 0 ]

  run usacloud_run iso-image delete -y "${ISO_IMAGE_NAME}"
  [ -n "${output}" ]
  [ "${status}" -eq 0 ]
}
@test "Usacloud: should can create and upload iso-image using STDIN" {
  run usacloud_run iso-image create -y --name "${ISO_IMAGE_NAME}" < ${DUMMY_ISO_FILE}

  [ -n "${output}" ]
  [ "${status}" -eq 0 ]
}
@test "Usacloud: should can read iso-image" {
  run usacloud_run iso-image read --out json "${ISO_IMAGE_NAME}"

  [ -n "${output}" ]
  [ "${status}" -eq 0 ]

  # parse JSON
  res=$(echo ${output} | jq ".[]")
  [ "$(echo ${res} | jq ".Availability")" == '"available"'  ]

  [ "$(echo ${res} | jq ".SizeMB")" -eq 5120 ]
}

@test "Usacloud: should can download iso-image" {
  run usacloud_run iso-image -y --file-destination "${DUMMY_ISO_DOWNLOAD_FILE}" "${ISO_IMAGE_NAME}"

  [ -n "${output}" ]
  [ "${status}" -eq 0 ]

  # compare .iso files
  [ -z "$(cmp "${DUMMY_ISO_FILE}" "${DUMMY_ISO_DOWNLOAD_FILE}")" ]
}

@test "Usacloud: should can delete iso-image" {
  run usacloud_run iso-image delete -y "${ISO_IMAGE_NAME}"

  [ -n "${output}" ]
  [ "${status}" -eq 0 ]
}