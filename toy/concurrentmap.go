package toy

import "sync"

var (
    locationMap  sync.Map
)

func JudgeItem(item string) bool {
    _, ok := locationMap.Load(item)
    if !ok {
        locationMap.Store(item, struct{}{})
    }

    return ok
}

func LoadItem(item string) (interface{}, bool) {
    return locationMap.Load(item)
}
