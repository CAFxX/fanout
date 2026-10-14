package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint64; F2 int }
type GenS3 struct { F0 string }
func genFn0(p0 int32, p1 uint64) int64 {
    if false {
        p1 = (^(p1))
    } else {
        if true {
            var v6 int32 = (p0 / ((-468) | 1))
            _ = v6
        }
    }
    var v7 float32 = 3.14159
    var v8 uint = 42
    v8 = (v8 << ((865) & 63))
    if true {
        var v9 float64 = 0.1
        var v10 string = ("" + ("hello" + "foo bar"))
        var v11 uint = 577
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int) int64 {
    if true {
        p0 += -1
    } else {
        if true {
            p0 = (^(p0))
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := GenS1{F0: 603}
                p0 = (^(fi12))
                _ = fi12
                _ = v13
            }
        }
        var v14 bool = (!(!false))
        v14 = false
        _ = v14
    }
    var v15 int8 = 127
    var v16 int8 = (v15 >> ((318) & 63))
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 string = "a"
        _ = fi17
        _ = v18
    }
    fmt.Println(p0, p0, v16)
    _ = p0
    _ = v15
    _ = v16
    return 42
}

func main() {
    v1 := []string{}
    var v2 float64 = 0.1
    var v3 uint = 607
    var v4 int64 = 9223372036854775807
    v5 := make(map[int]int)
    fmt.Println(v2, v3)
    v2 = float64((0.1 - (float32((float64((1e10 - (0.0 + 1.5))) + (v2 / (v2 / (3.14159 + v2))))) - 0.1)))
    switch 0 {
    case 478:
        v19 := GenS1{F0: v3}
        _ = v19
    }
    if len(v1) > 0 {
        v1[0] = ("a" + "hello")
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := GenS2{F0: -1, F1: 107, F2: (fi20 * (^(fi20)))}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi20
        _ = v21
    }
    var v22 int64 = genFn0(938, 1)
    for fi23 := 0; fi23 < 8; fi23++ {
        if (!true) {
            var v24 bool = ((("a" + "hello") < ("a" + "foo bar")) || true)
            _ = v24
        } else {
            if true {
                _ = len(v1)
                _ = len(v1)
            } else {
                if (32767 <= 42) {
                    v3 = (v3 ^ v3)
                    var v25 int16 = 32767
                    v22 += (v4 * (v22 >> ((v3) & 63)))
                    _ = v25
                } else {
                    var v26 uint = 2
                    _ = v26
                }
            }
            v27 := GenS2{F0: 100, F1: 2, F2: fi23}
            var v28 bool = (!((!((v22 > (v4 ^ v4)) && (false || true))) || true))
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            v2 = 1.5
            v30 := GenS3{F0: "\n\t"}
            _ = fi29
            _ = v30
        }
        _ = fi23
    }
    var v31 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
