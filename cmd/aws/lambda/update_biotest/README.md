# Update Biotest

## Request

```JSON
{
    "id": 12,
    "higher_muscle_density": {
        "neck": null,
        "shoulders": null,
        "back": null,
        "chest": null,
        "back_chest": null,
        "right_relaxed_bicep": null,
        "right_contracted_bicep": null,
        "left_relaxed_bicep": null,
        "left_contracted_bicep": null,
        "right_forearm": null,
        "left_forearm": null,
        "wrists": null,
        "high_abdomen": null,
        "lower_abdomen": null
    },
    "higher_muscle_density_id": 0,
    "lower_muscle_density": {
        "hips": null,
        "right_leg": null,
        "left_leg": null,
        "right_calf": null,
        "left_calf": null
    },
    "lower_muscle_density_id": 0,
    "skin_folds": {
        "subscapular": null,
        "suprailiac": null,
        "bicipital": null,
        "tricipital": null
    },
    "skin_folds_id": 0,
    "customer": {
        "biotype_id": null,
        "bone_density_id": null,
        "gender_id": null,
        "user_uuid": "",
        "avatar_url": "",
        "birthday": "0001-01-01T00:00:00Z",
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z"
    },
    "biotest_uuid": "uuid",
    "corporal_age": 0,
    "chronological_age": 0,
    "glucose": null,
    "resting_heart_rate": null,
    "maximum_heart_rate": null,
    "observations": null,
    "recommendations": null,
    "front_picture": null,
    "back_picture": null,
    "right_side_picture": null,
    "left_side_picture": null,
    "next_evaluation": null,
    "created_at": "0001-01-01T00:00:00Z"
}
```

## Updated

```JSON
{
    "code": 200,
    "status": "OK"
}
```

## Without id

```JSON
{
    "code": 400,
    "status": "Bad Request",
    "body": [
        {
        "tag": "id",
        "validation": "required"
        }
    ]
}
```

## Validation Errors

```JSON
{
    "code": 400,
    "status": "Bad Request",
    "body": {
        "error": [
            {
                "tag": "weight",
                "validation": "required"
            },
            {
                "tag": "height",
                "validation": "required"
            }
        ]
    }
}
```
