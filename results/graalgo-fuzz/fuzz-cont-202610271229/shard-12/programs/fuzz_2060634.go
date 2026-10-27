package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint) int {
    p0 = p0
    var v4 uint32 = 1
    var v5 uint64 = 85
    var v6 int64 = 0
    if (!false) {
        v7 := GenS1{F0: -1, F1: "\n\t"}
        _ = v7
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 uint32) uint8 {
    p0 = p0
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 bool = true
        v10 := GenS2{F0: ""}
        v11 := []int{127}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 uint16 = 987
    _ = p0
    _ = v12
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    var v13 int64 = 1
    v13 = v13
    switch -678 {
    case 1:
        var v14 int8 = -1
        _ = v14
    default:
        var v15 int64 = v13
        _ = v15
    }
    switch "foo bar" {
    case "\n\t":
        v13 = v13
    case "a":
        if true {
            if false {
                var v16 int8 = 82
                if false {
                    v16 = -128
                } else {
                    if (-257 > 32767) {
                        for fi17 := 0; fi17 < 20; fi17++ {
                            var v18 uint64 = 4294967295
                            v16 -= 100
                            _ = fi17
                            _ = v18
                        }
                        fmt.Println(v13, v16)
                        var v19 uint8 = genFn1(297)
                        _ = v19
                    }
                }
                _ = v16
            } else {
                v13 -= v13
            }
            if true {
                fmt.Println(v13)
                var v20 uint16 = 0
                var v21 float64 = 1e10
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 11; fi22++ {
                v13 = (v13 << ((65535) & 63))
                _ = fi22
            }
        }
    case "foo bar":
        _ = len(v3)
    }
    v23 := GenS3{F0: 1, F1: 260}
    v24 := GenS2{F0: ""}
    for fi25 := 0; fi25 < 10; fi25++ {
        switch genFn0(42) {
        case -1:
            switch fi25 {
            case -328:
                v26 := []int{-2147483648, 670, fi25, (-(fi25))}
                _ = v26
            case -9223372036854775808:
                var v27 uint64 = 935
                _ = v27
            }
        case 1:
            var v28 int16 = 110
            _ = v28
        }
        _ = fi25
    }
    if (0 != 127) {
        _ = len(v1)
    } else {
        fmt.Println(v13)
        _ = len(v1)
    }
    v29 := GenS1{F0: -1, F1: "hello"}
    var v30 int16 = 32767
    for fi31 := 0; fi31 < 6; fi31++ {
        for fi32 := 0; fi32 < 19; fi32++ {
            var v33 uint8 = 42
            v13 = (v13 - 42)
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    v1[""] = 4294967295
    var v34 int64 = (v13 ^ (^(v13)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v34
}
