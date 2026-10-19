package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint8; F2 string }
func genFn0() uint64 {
    v5 := []int{-416, -789, -2147483648, 553}
    var v6 uint = 926
    v7 := GenS1{F0: -1, F1: 32767, F2: 0}
    _ = v5
    _ = v6
    _ = v7
    return 656
}

func main() {
    var v1 int = -1
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    switch "" {
    case "a":
        switch "" {
        case "foo bar":
            v8 := GenS2{F0: 18446744073709551615}
            _ = v8
        default:
            v1 += (v1 - (v1 | (v1 >> ((2) & 63))))
        }
    default:
        if (false || false) {
            v9 := v3["\n\t"]
            var v10 float64 = 1e10
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 14; fi11++ {
                if (100 == -83) {
                    v12 := GenS1{F0: 1, F1: 1, F2: -2147483648}
                    _ = v12
                }
                _ = fi11
            }
            v1 += 936
        }
    }
    v1 = v1
    v1 = v1
    if (false && false) {
        v13 := []int{(v1 * v1)}
        if false {
            switch 0 {
            case 32768:
                v14 := GenS3{F0: 2147483647, F1: 0, F2: ("hello" + "")}
                _ = v14
            case 2147483647:
                v15 := GenS1{F0: 100, F1: 65535, F2: -1}
                _ = v15
            default:
                var v16 bool = (710 > 1)
                _ = v16
            }
            v1 = (v1 + v1)
            var v17 int32 = -1
            _ = v17
        } else {
            v1 = (-(v1))
            if true {
                for fi18 := 0; fi18 < 3; fi18++ {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        v1 -= 32768
                        v20 := GenS3{F0: 642, F1: 127, F2: "a"}
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
            }
            v21 := GenS1{F0: 127, F1: 1, F2: -2147483648}
            _ = v21
        }
        var v22 float32 = (float32(float64(3.14159)) * 3.14159)
        _ = v13
        _ = v22
    } else {
        v1 = (^(v1))
        delete(v3, "hello")
        switch "a" {
        case "\n\t":
            v1 = v1
        }
    }
    switch -204 {
    case 32767:
        _ = len(v3)
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        v1 = fi23
        var v24 uint64 = 4294967295
        if false {
            var v25 uint64 = v24
            v26 := v2["foo bar"]
            _ = v25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    if (113 != 2) {
        v27 := GenS3{F0: 2147483647, F1: 42, F2: ("hello" + "")}
        v28 := v2["a"]
        _ = v27
        _ = v28
    }
    fmt.Println(v1)
    var v29 uint = 2
    var v30 uint16 = 545
    v29 += v29
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
