package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int64; F2 int16 }
type GenS3 struct { F0 int16; F1 string; F2 int32 }
type GenS4 struct { F0 string; F1 uint32 }
func genFn0(p0 bool, p1 uint) bool {
    fmt.Println(p1)
    switch "a" {
    case "":
        if p0 {
            var v4 uint64 = 18446744073709551615
            _ = v4
        } else {
            fmt.Println(p1, p0)
        }
    case "\n\t":
        var v5 uint64 = 1
        _ = v5
    }
    if false {
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS1{F0: 42}
            v8 := []int{}
            var v9 bool = (0.0 != (0.1 * (((float64(float32(1e10)) - 0.0) - 3.14159) / 1.5)))
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        fmt.Println(p0)
    } else {
        var v10 uint64 = 4294967295
        v11 := GenS2{F0: p1, F1: 1, F2: 1000}
        v12 := GenS1{F0: 65535}
        _ = v10
        _ = v11
        _ = v12
    }
    if p0 {
        p1 = 125
        if ((p1 >> ((p1) & 63)) > 0) {
            var v13 int8 = 100
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(fi14, fi14, p0)
        _ = fi14
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 42
    v3 := make(map[string]string)
    if false {
        var v15 string = ("hello" + "a")
        v16 := []int{42, 127, 32768}
        if len(v16) > 0 {
            v16[0] = int(v2)
        }
        _ = v15
        _ = v16
    }
    if (!("" != "foo bar")) {
        var v17 uint = 2
        var v18 float32 = 3.14159
        _ = v17
        _ = v18
    }
    var v19 int32 = (v2 & v2)
    v20 := GenS3{F0: -855, F1: "foo bar", F2: v19}
    if (true && (!false)) {
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 int = (-(fi21))
            v22 = (v22 << ((255) & 63))
            _ = fi21
            _ = v22
        }
        if (4294967295 < 4294967295) {
            var v23 uint64 = 68
            _ = v23
        } else {
            var v24 uint32 = 20
            v25 := GenS3{F0: -820, F1: "a", F2: v19}
            switch "foo bar" {
            case "hello":
                v2 = (v2 % ((v2) | 1))
            default:
                fmt.Println(v24, v2, v2)
            }
            _ = v24
            _ = v25
        }
    }
    v2 = (v2 ^ v2)
    if true {
        v2 = v2
        if false {
            v2 += -442
            switch 2147483647 {
            case 0:
                var v26 uint64 = 1
                _ = v26
            }
            switch 9223372036854775807 {
            case 759:
                v27 := v1[""]
                _ = v27
            case -79:
                if true {
                    var v28 float32 = 1.5
                    if (!false) {
                        v29 := GenS4{F0: "\n\t", F1: 4294967295}
                        _ = v29
                    }
                    _ = v28
                }
            case 949:
                switch 9223372036854775807 {
                case 32768:
                    fmt.Println(v19)
                case 833:
                    fmt.Println(v2, v2)
                default:
                    v2 = (-(v19))
                }
            default:
                v19 += (^(v2))
            }
        }
        var v30 uint16 = 185
        _ = v30
    } else {
        var v31 uint8 = 0
        v31 = (v31 ^ 128)
        fmt.Println(v31, v31)
        _ = v31
    }
    fmt.Println(v2)
    v19 = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
