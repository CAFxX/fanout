package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string }
func genFn0(p0 int8, p1 uint64) uint {
    fmt.Println(p1, p0)
    switch 9223372036854775807 {
    case 42:
        for fi4 := 0; fi4 < 6; fi4++ {
            if true {
                if (!true) {
                    fmt.Println(p1, p0)
                    var v5 int16 = -982
                    _ = v5
                }
                var v6 int16 = -32768
                var v7 uint = 1
                _ = v6
                _ = v7
            }
            fmt.Println(p0)
            _ = fi4
        }
    case 2147483647:
        p0 = 119
    default:
        switch "" {
        case "foo bar":
            v8 := GenS1{F0: (false && false), F1: false, F2: false}
            _ = v8
        case "a":
            if false {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 int64 = -1
                    v11 := []int{(fi9 - 42), fi9}
                    v12 := GenS2{F0: 1}
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            }
        }
    }
    v13 := GenS3{F0: 578, F1: p1}
    _ = p0
    _ = p1
    _ = v13
    return 4294967295
}
func genFn1(p0 int16) uint8 {
    v14 := []int{2147483647, int(p0), 0, 128}
    fmt.Println(p0)
    var v15 float32 = (float32(-2.25) - (1e10 - 1e10))
    if false {
        var v16 uint16 = 1
        _ = v16
    }
    _ = p0
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 bool = (!true)
    v2 := make(map[string]uint32)
    var v3 int64 = 42
    v1 = ((32767 <= -274) || ("\n\t" < "\n\t"))
    v17 := GenS2{F0: 255}
    var v18 uint16 = 2
    fmt.Println(v3)
    var v19 int32 = -2147483648
    v2[""] = 42
    switch 32767 {
    case 1:
        v20 := GenS4{F0: (((1.5 + 0.0) * (float64(0.0) / ((1.5 + -2.25) + 0.0))) > (0.0 / (-2.25 * 1.5)))}
        _ = v20
    default:
        var v21 bool = v1
        _ = v21
    }
    if (v19 >= (v19 * (v19 >> ((1) & 63)))) {
        v19 = int32(v18)
        var v22 int32 = -932
        _ = v22
    } else {
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := GenS5{F0: (("foo bar" + ("foo bar" + ("\n\t" + "hello"))) + "hello")}
            var v25 uint32 = 0
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
