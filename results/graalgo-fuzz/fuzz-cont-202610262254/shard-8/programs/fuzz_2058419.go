package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int8; F1 int16; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint32; F1 int32; F2 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int; F1 uint16 }

func main() {
    v1 := []uint32{504, 2, 32, 777}
    v2 := []uint32{34}
    var v3 float64 = 3.14159
    if false {
        var v4 uint16 = 32767
        if len(v2) > 0 {
            v2[0] = 42
        }
        v3 += 0.1
        _ = v4
    }
    switch "" {
    case "a":
        v5 := GenS1{F0: -113, F1: 0}
        _ = v5
    case "":
        for fi6 := 0; fi6 < 7; fi6++ {
            v3 -= (v3 - v3)
            switch "\n\t" {
            case "hello":
                switch (fi6 ^ fi6) {
                case 79:
                    var v7 uint8 = 165
                    _ = v7
                }
            default:
                fmt.Println(v3, fi6)
            }
            _ = fi6
        }
    case "foo bar":
        v3 = v3
    }
    fmt.Println(v3)
    if false {
        var v8 uint16 = 42
        _ = len(v2)
        v3 = 1e10
        _ = v8
    }
    v9 := GenS2{F0: 2147483647, F1: 127}
    v10 := GenS3{F0: 1, F1: 619, F2: "foo bar"}
    var v11 string = ("hello" + ("hello" + "\n\t"))
    v12 := GenS4{F0: 1}
    if true {
        if false {
            v13 := GenS5{F0: v11, F1: 122}
            var v14 bool = (true || true)
            _ = v13
            _ = v14
        } else {
            var v15 int16 = 32767
            v16 := GenS4{F0: 65}
            var v17 uint16 = 0
            _ = v15
            _ = v16
            _ = v17
        }
        if true {
            var v18 int8 = -8
            _ = v18
        }
    }
    fmt.Println(v11, v11)
    v19 := GenS6{F0: 347, F1: 42, F2: ((v11 + (("" + v11) + "hello")) + v11)}
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS7{F0: -1}
        if false {
            v22 := GenS8{F0: 732, F1: 32767}
            _ = v22
        } else {
            for fi23 := 0; fi23 < 20; fi23++ {
                fmt.Println(v3, v11)
                _ = fi23
            }
            v24 := GenS6{F0: 1, F1: 2147483647, F2: "foo bar"}
            _ = v24
        }
        var v25 float32 = (1.5 + 3.14159)
        _ = fi20
        _ = v21
        _ = v25
    }
    var v26 int8 = -26
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v9
}
