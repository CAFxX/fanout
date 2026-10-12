package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint16 }
func genFn0() bool {
    var v5 int16 = 35
    if false {
        v5 = -32768
    } else {
        v5 = (v5 - (^(v5)))
        var v6 float32 = 0.0
        v7 := []int{-9223372036854775808, 32768, -556, 4294967295}
        _ = v6
        _ = v7
    }
    var v8 int64 = 746
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: v5, F1: -9223372036854775808, F2: 554}
        _ = fi9
        _ = v10
    }
    var v11 int = 32768
    _ = v11
    _ = v5
    _ = v8
    return false
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = -2.25
    var v3 uint32 = 42
    v4 := make(map[string]uint32)
    var v12 int = 758
    var v13 float32 = v2
    var v14 uint = 4294967295
    delete(v1, "foo bar")
    fmt.Println(v14, v3)
    if ((v2 * v2) != -2.25) {
        v3 -= (v3 - (v3 ^ 251))
        fmt.Println(v13, v3)
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := []int{(v12 - fi15)}
            v18 := GenS1{F0: 1, F1: 0, F2: 255}
            var v19 int64 = -836
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            switch ("foo bar" + (("foo bar" + "hello") + "")) {
            case "\n\t":
                var v21 float32 = v13
                v13 -= float32(1.5)
                _ = v21
            case "foo bar":
                var v22 int32 = 0
                _ = v22
            }
            v4[""] = 65
            var v23 bool = (!(!(100 <= -1)))
            _ = fi20
            _ = v23
        }
        _ = fi15
    }
    var v24 uint64 = 1
    v25 := []int{v12, (v12 ^ v12), v12}
    v2 = 1.5
    v3 = (v3 - v3)
    v13 += 3.14159
    for fi26 := 0; fi26 < 17; fi26++ {
        fmt.Println(v2, fi26, fi26)
        for fi27 := 0; fi27 < 5; fi27++ {
            switch (v12 * fi27) {
            case 173:
                var v28 uint16 = 2
                _ = v28
            case 32767:
                v29 := GenS1{F0: 0, F1: 927, F2: 32767}
                _ = v29
            case 0:
                var v30 uint64 = uint64(v3)
                _ = v30
            default:
                v31 := GenS1{F0: 1000, F1: 9223372036854775807, F2: 145}
                _ = v31
            }
            if (!(!false)) {
                var v32 uint16 = 222
                fmt.Println(v14, v2)
                v33 := GenS1{F0: 32767, F1: -477, F2: (v32 << ((0) & 63))}
                switch fi26 {
                case 9223372036854775807:
                    _ = len(v25)
                case -9223372036854775808:
                    var v34 uint64 = 75
                    _ = v34
                case 128:
                    var v35 int32 = 1
                    _ = v35
                default:
                    var v36 int8 = -128
                    _ = v36
                }
                v2 += v13
                _ = v32
                _ = v33
            }
            _ = fi27
        }
        _ = fi26
    }
    v37 := GenS1{F0: 48, F1: -1, F2: uint16(v3)}
    fmt.Println(v3, v2)
    v14 += 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v4
}
