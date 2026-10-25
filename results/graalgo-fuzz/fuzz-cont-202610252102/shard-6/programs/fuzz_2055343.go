package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 65535
    v3 := make(map[int]uint32)
    var v4 uint32 = 245
    v4 += 2
    var v5 string = ("foo bar" + "\n\t")
    v6 := GenS1{F0: 127, F1: 312}
    delete(v1, 331)
    if (v4 != 42) {
        v5 = "foo bar"
    }
    v4 = (^(v4))
    var v7 int = 32768
    var v8 int16 = -573
    v9 := []int{2147483647}
    var v10 float64 = (float64(3.14159) - float64(3.14159))
    for fi11 := 0; fi11 < 13; fi11++ {
        v4 -= v4
        var v12 uint64 = 0
        if (false && (!true)) {
            var v13 string = "hello"
            v14 := v3[(fi11 - (fi11 * v7))]
            _ = v13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        v5 = v5
        var v16 int64 = 42
        _ = fi15
        _ = v16
    }
    if (0 <= 69) {
        if true {
            if (false || (-64 == -86)) {
                for fi17 := 0; fi17 < 18; fi17++ {
                    if (uint8(v2) <= 75) {
                        if true {
                            v5 = (v5 + "foo bar")
                            var v18 int16 = -32768
                            _ = v18
                        }
                        v19 := GenS1{F0: 1, F1: 4294967295}
                        _ = v19
                    } else {
                        v20 := GenS2{F0: 127}
                        _ = v20
                    }
                    var v21 float32 = 1e10
                    _ = fi17
                    _ = v21
                }
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        } else {
            if (1 <= -127) {
                if true {
                    fmt.Println(v5, v5)
                }
                v22 := v3[(v7 | v7)]
                _ = v22
            }
            var v23 float64 = (float64(-2.25) * 0.0)
            var v24 uint = 4294967295
            _ = v23
            _ = v24
        }
        if false {
            if (v8 != (-(v8))) {
                for fi25 := 0; fi25 < 3; fi25++ {
                    v5 = v5
                    v8 = (-(v8))
                    v2 = (^(v2))
                    _ = fi25
                }
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        switch v7 {
        case -2147483648:
            switch v5 {
            case "a":
                var v27 uint16 = 1
                _ = v27
            case "":
                var v28 int64 = 0
                _ = v28
            case "hello":
                v8 = (v8 - v8)
            }
        case 42:
            v29 := GenS2{F0: 0}
            _ = v29
        case 127:
            var v30 string = v5
            _ = v30
        default:
            fmt.Println(v10)
        }
        var v31 float32 = float32(3.14159)
        v10 -= v10
        _ = fi26
        _ = v31
    }
    var v32 uint32 = (^(v4))
    if true {
        if (true || false) {
            var v33 float64 = v10
            v4 += (v32 >> ((65535) & 63))
            v9 = append(v9, 9223372036854775807)
            _ = v33
        }
        v8 += 32767
    } else {
        var v34 int32 = 42
        var v35 uint16 = (v2 - (v2 >> ((0) & 63)))
        _ = v34
        _ = v35
    }
    _ = len(v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
