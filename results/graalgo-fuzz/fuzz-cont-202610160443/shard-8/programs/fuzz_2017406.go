package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint16; F1 int }
func genFn0(p0 uint32) uint64 {
    v5 := GenS1{F0: 42}
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 string = ""
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return 734
}

func main() {
    var v1 int64 = 1
    var v2 uint16 = 2
    v3 := make(map[uint32]string)
    var v4 uint64 = 575
    var v8 uint8 = 167
    var v9 uint16 = v2
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int16 = -1
        v12 := GenS1{F0: (^(v8))}
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v4)
    if (false || true) {
        var v13 uint8 = (^(v8))
        var v14 uint = 1
        for fi15 := 0; fi15 < 14; fi15++ {
            v16 := GenS1{F0: (^(v8))}
            _ = fi15
            _ = v16
        }
        _ = v13
        _ = v14
    }
    var v17 uint32 = 65535
    if false {
        v18 := GenS2{F0: false, F1: 127}
        v2 -= v9
        fmt.Println(v17)
        _ = v18
    } else {
        var v19 int32 = 1
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 int = (fi20 & fi20)
        v2 = (v2 & 255)
        _ = fi20
        _ = v21
    }
    switch 432 {
    case 894:
        for fi22 := 0; fi22 < 1; fi22++ {
            fmt.Println(v2)
            _ = fi22
        }
    case 1:
        switch 127 {
        case 4294967295:
            v17 = 2
        case -9223372036854775808:
            if (!(-128 == 127)) {
                v23 := GenS2{F0: true, F1: 192}
                fmt.Println(v4)
                for fi24 := 0; fi24 < 6; fi24++ {
                    fmt.Println(fi24)
                    _ = fi24
                }
                _ = v23
            } else {
                fmt.Println(v17)
                var v25 int32 = 2147483647
                delete(v3, v17)
                _ = v25
            }
        default:
            fmt.Println(v17, v1)
        }
    case 9223372036854775807:
        for fi26 := 0; fi26 < 8; fi26++ {
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 bool = ((-2.25 + 1.5) == 1.5)
                _ = fi27
                _ = v28
            }
            var v29 int8 = 42
            _ = fi26
            _ = v29
        }
    }
    v30 := GenS1{F0: (^(v8))}
    var v31 uint16 = (v9 * (v9 % (((^(v9))) | 1)))
    for fi32 := 0; fi32 < 17; fi32++ {
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 uint32 = (v17 & v17)
            v35 := GenS3{F0: 320, F1: 9223372036854775807}
            fmt.Println(v34)
            _ = fi33
            _ = v34
            _ = v35
        }
        v8 -= 255
        _ = fi32
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        fmt.Println(v17, v9, v31)
        _ = fi36
    }
    v37 := []int{502, -1}
    v31 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
    _ = v8
    _ = v9
}
