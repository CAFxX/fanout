package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int) bool {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int = (-(fi5))
        v7 := []int{}
        fmt.Println(p0, v6, v6)
        _ = fi5
        _ = v6
        _ = v7
    }
    switch p0 {
    case 32768:
        p0 -= -9223372036854775808
    case 9223372036854775807:
        for fi8 := 0; fi8 < 11; fi8++ {
            if (749 > 255) {
                v9 := GenS1{F0: 1000}
                _ = v9
            } else {
                var v10 uint32 = 627
                _ = v10
            }
            _ = fi8
        }
    case 193:
        p0 += (p0 | p0)
    }
    var v11 int16 = 1000
    for fi12 := 0; fi12 < 3; fi12++ {
        if true {
            fmt.Println(fi12)
            var v13 uint8 = 128
            v13 = v13
            _ = v13
        } else {
            var v14 int = -814
            fmt.Println(v14, v11)
            _ = v14
        }
        _ = fi12
    }
    _ = p0
    _ = v11
    return (true || true)
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    v4 := []int{37, -458, 32768, -436}
    for fi15 := 0; fi15 < 8; fi15++ {
        fmt.Println(fi15)
        var v16 int8 = -2
        _ = fi15
        _ = v16
    }
    var v17 int64 = -888
    v18 := GenS1{F0: 1}
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 int8 = 42
        _ = fi19
        _ = v20
    }
    var v21 bool = (4294967295 == 4294967295)
    v17 -= 0
    switch 127 {
    case -2147483648:
        v22 := GenS1{F0: -702}
        _ = v22
    case 128:
        var v23 int16 = -32768
        _ = v23
    }
    switch 0 {
    case 41:
        v17 += (v17 - (v17 << ((255) & 63)))
    case -2147483648:
        var v24 int16 = -1
        _ = v24
    case 0:
        var v25 int8 = -128
        _ = v25
    }
    if (990 <= 309) {
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 uint32 = 0
            fmt.Println(fi26)
            v27 = (v27 ^ (v27 * (v27 & (^(v27)))))
            _ = fi26
            _ = v27
        }
    } else {
        if (v21 || (232 > 0)) {
            for fi28 := 0; fi28 < 16; fi28++ {
                v21 = ((!(v17 < v17)) || v21)
                _ = fi28
            }
            _ = len(v4)
        } else {
            v29 := GenS1{F0: 1}
            _ = v29
        }
        var v30 uint64 = 1
        fmt.Println(v30)
        _ = v30
    }
    v31 := GenS1{F0: 1000}
    var v32 uint16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
