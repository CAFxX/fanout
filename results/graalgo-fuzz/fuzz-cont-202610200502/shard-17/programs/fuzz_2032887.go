package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 string; F2 int8 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int32; F1 uint16; F2 uint16 }
func genFn0(p0 uint64) uint16 {
    if true {
        v7 := GenS1{F0: 128, F1: false, F2: 100}
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 int32 = 2147483647
            v10 := GenS1{F0: 26, F1: true, F2: 42}
            _ = fi8
            _ = v10
            _ = v9
        }
        switch -614 {
        case 2147483647:
            switch "hello" {
            case "foo bar":
                for fi11 := 0; fi11 < 6; fi11++ {
                    var v12 int8 = -21
                    _ = fi11
                    _ = v12
                }
            case "a":
                v13 := GenS1{F0: 1, F1: (!(0.1 > (-2.25 + 1e10))), F2: -63}
                _ = v13
            }
        case 32768:
            fmt.Println(p0)
        }
        _ = v7
    }
    if ("" != ("hello" + "a")) {
        v14 := GenS2{F0: -2147483648}
        _ = v14
    }
    _ = p0
    return 32767
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []int{-861}
    v3 := []uint32{}
    var v4 float32 = 1.5
    var v5 string = "a"
    var v6 bool = false
    if true {
        v15 := GenS3{F0: v1, F1: (v5 + "foo bar"), F2: 0}
        v6 = ((v1 + 42) == 1)
        v16 := GenS4{F0: (1 < 18446744073709551615), F1: 334}
        _ = v15
        _ = v16
    }
    fmt.Println(v6, v1, v4)
    fmt.Println(v5, v4)
    fmt.Println(v5, v1, v5)
    var v17 int64 = 9223372036854775807
    switch -850 {
    case 2147483647:
        v18 := GenS5{F0: -1, F1: 32767, F2: genFn0(0)}
        _ = v18
    case 9223372036854775807:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v19 bool = true
    switch 42 {
    case -19:
        var v20 int = 1
        _ = v20
    case -9223372036854775808:
        var v21 int64 = (v17 & v17)
        _ = v21
    default:
        if true {
            v19 = true
        } else {
            var v22 float64 = (0.1 / float64(float32(1e10)))
            if len(v3) > 0 {
                v3[0] = v1
            }
            _ = v22
        }
    }
    var v23 bool = (310 > -2147483648)
    var v24 float32 = v4
    v24 += -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
