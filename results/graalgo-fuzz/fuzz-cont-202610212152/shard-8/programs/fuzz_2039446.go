package main

import (
    "fmt"
)

func genFn0() int64 {
    var v4 uint = 760
    v5 := []int{}
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 42
    var v3 uint16 = 42
    v6 := []int{42}
    v3 = v3
    v2 -= uint8(v3)
    if false {
        var v7 uint8 = v2
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint8 = (^(v2))
        _ = fi8
        _ = v9
    }
    v1[4294967295] = 65535
    for fi10 := 0; fi10 < 10; fi10++ {
        switch -1 {
        case -525:
            var v11 uint64 = 4294967295
            _ = v11
        }
        var v12 bool = false
        _ = fi10
        _ = v12
    }
    switch "a" {
    case "":
        v2 = 0
    case "a":
        v13 := v1[556]
        _ = v13
    case "\n\t":
        v3 += 255
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            if false {
                if false {
                    var v16 int16 = 1000
                    var v17 uint16 = v3
                    _ = v16
                    _ = v17
                } else {
                    var v18 int = fi15
                    v2 = (v2 << ((42) & 63))
                    _ = v18
                }
            }
            var v19 float64 = 1.5
            _ = fi15
            _ = v19
        }
        v3 = v3
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
}
