package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint16, p1 uint64) int32 {
    for fi7 := 0; fi7 < 19; fi7++ {
        if ((false && ((fi7 | fi7) <= (fi7 + 32767))) || ((p0 % (((p0 - (p0 / ((730) | 1)))) | 1)) > (p0 - (p0 | p0)))) {
            fmt.Println(fi7)
        } else {
            if true {
                var v8 int16 = -477
                _ = v8
            }
        }
        v9 := GenS1{F0: 778}
        if (!false) {
            v10 := []int{}
            _ = len(v10)
            _ = v10
        }
        _ = fi7
        _ = v9
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS1{F0: 4294967295}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -123
    var v3 int64 = v2
    var v4 uint16 = 0
    v5 := make(map[string]int)
    var v6 uint32 = 2
    v3 -= -9223372036854775808
    if true {
        var v13 int32 = 2147483647
        v14 := v1["hello"]
        var v15 int64 = v2
        _ = v13
        _ = v14
        _ = v15
    } else {
        switch 1 {
        case -2147483648:
            var v16 uint8 = 2
            _ = v16
        case -1:
            var v17 int32 = 0
            _ = v17
        }
        var v18 int16 = 1000
        _ = v18
    }
    _ = len(v1)
    fmt.Println(v2, v4, v6)
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 uint = 255
        v6 = (v6 / (((v6 / ((v6) | 1))) | 1))
        v21 := []int{285, fi19, int(v4)}
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 int64 = v2
    for fi23 := 0; fi23 < 10; fi23++ {
        v24 := GenS2{F0: (^(fi23)), F1: (v6 % ((v6) | 1))}
        _ = fi23
        _ = v24
    }
    v25 := GenS3{F0: 18446744073709551615, F1: 1}
    v6 += 831
    for fi26 := 0; fi26 < 9; fi26++ {
        for fi27 := 0; fi27 < 4; fi27++ {
            fmt.Println(fi27, v6, v2)
            v3 += (v22 << ((531) & 63))
            _ = fi27
        }
        var v28 int16 = 42
        var v29 int32 = genFn0(v4, 2)
        _ = fi26
        _ = v28
        _ = v29
    }
    var v30 float32 = -2.25
    v4 += v4
    for fi31 := 0; fi31 < 8; fi31++ {
        v32 := GenS1{F0: 136}
        v3 += (v22 << ((676) & 63))
        _ = fi31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
