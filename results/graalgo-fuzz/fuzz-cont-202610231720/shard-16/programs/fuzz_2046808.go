package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int; F1 int32 }
func genFn0(p0 int16) int64 {
    v6 := GenS1{F0: (0 < 3), F1: 0}
    p0 = (p0 - p0)
    var v7 int64 = 820
    fmt.Println(p0, v7)
    var v8 int16 = p0
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return v7
}
func genFn1() uint {
    v9 := []int{}
    var v10 bool = (1.5 > 0.0)
    _ = v10
    _ = v9
    return 2
}

func main() {
    var v1 bool = (("hello" + "") >= "\n\t")
    var v2 int8 = 0
    v3 := make(map[string]int)
    var v4 int64 = -1
    var v5 uint32 = 0
    fmt.Println(v1, v2, v4)
    var v11 float32 = 0.0
    v12 := GenS2{F0: ("hello" + ("foo bar" + "hello")), F1: (v4 - (v4 | v4))}
    v2 = 1
    fmt.Println(v2, v4)
    if v1 {
        switch "\n\t" {
        case "":
            var v13 float64 = 0.0
            _ = v13
        case "a":
            v14 := GenS3{F0: (^(v2))}
            _ = v14
        case "\n\t":
            fmt.Println(v4, v11, v5)
        }
        v15 := GenS4{F0: "foo bar", F1: "hello"}
        _ = v15
    } else {
        var v16 int64 = v4
        v5 = (v5 - 37)
        _ = v16
    }
    delete(v3, ("foo bar" + "foo bar"))
    for fi17 := 0; fi17 < 16; fi17++ {
        fmt.Println(v2, v4, v1)
        v1 = v1
        _ = fi17
    }
    switch "" {
    case "\n\t":
        v11 -= (v11 + 0.0)
    default:
        if ((127 <= 1) && false) {
            for fi18 := 0; fi18 < 18; fi18++ {
                v5 = 4294967295
                var v19 uint32 = (^(v5))
                _ = fi18
                _ = v19
            }
        } else {
            v20 := []int{}
            switch "a" {
            case "\n\t":
                var v21 bool = true
                _ = v21
            default:
                for fi22 := 0; fi22 < 1; fi22++ {
                    for fi23 := 0; fi23 < 2; fi23++ {
                        _ = len(v3)
                        var v24 uint32 = (v5 + v5)
                        v1 = (v5 >= v5)
                        _ = fi23
                        _ = v24
                    }
                    if v1 {
                        var v25 int32 = -349
                        _ = v25
                    } else {
                        v26 := GenS3{F0: v2}
                        var v27 string = "foo bar"
                        _ = v26
                        _ = v27
                    }
                    v28 := GenS5{F0: (^(fi22)), F1: 42}
                    _ = fi22
                    _ = v28
                }
            }
            _ = v20
        }
    }
    v4 -= -9223372036854775808
    v2 -= v2
    if (32768 > 0) {
        switch "foo bar" {
        case "foo bar":
            if false {
                var v29 uint8 = 221
                v4 = (-(v4))
                _ = v29
            }
        case "":
            delete(v3, "a")
        case "\n\t":
            for fi30 := 0; fi30 < 4; fi30++ {
                v31 := []int{32768, int(v5), (^(fi30))}
                var v32 int = (fi30 << ((1) & 63))
                _ = fi30
                _ = v31
                _ = v32
            }
        default:
            fmt.Println(v2)
        }
        fmt.Println(v1, v1, v11)
    }
    var v33 uint16 = 923
    v11 = ((v11 - 3.14159) * float32(((0.1 - -2.25) * 1e10)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
