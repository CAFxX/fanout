package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int64 }
func genFn0(p0 int, p1 uint16, p2 int64) uint32 {
    p2 = 9223372036854775807
    if false {
        if true {
            fmt.Println(p0)
            var v6 float64 = 3.14159
            var v7 uint32 = 1
            _ = v6
            _ = v7
        } else {
            v8 := GenS1{F0: 42, F1: p1, F2: (p2 ^ (p2 / (((p2 ^ 42)) | 1)))}
            _ = v8
        }
        var v9 uint32 = 639
        v10 := []int{p0}
        _ = v10
        _ = v9
    }
    var v11 uint64 = 798
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 65535
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    var v4 uint64 = 339
    var v5 int8 = 127
    if false {
        v12 := GenS1{F0: 747, F1: 255, F2: v1}
        var v13 uint = 0
        _ = v12
        _ = v13
    }
    if false {
        for fi14 := 0; fi14 < 10; fi14++ {
            fmt.Println(v1, fi14, v5)
            _ = fi14
        }
        var v15 string = "a"
        _ = v15
    }
    var v16 uint8 = 6
    var v17 uint8 = (v16 | v16)
    v17 = (v16 + (^(v17)))
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            if false {
                var v20 int64 = v1
                _ = v20
            }
            var v21 int64 = int64(v5)
            _ = fi19
            _ = v21
        }
        v4 = v4
        _ = fi18
    }
    var v22 uint64 = 938
    v23 := GenS1{F0: 0, F1: 2, F2: (^(v1))}
    switch int(v5) {
    case -42:
        delete(v3, 0)
    }
    var v24 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
