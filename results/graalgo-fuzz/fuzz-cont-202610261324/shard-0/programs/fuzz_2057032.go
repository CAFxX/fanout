package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 string; F2 int16 }
func genFn0(p0 bool, p1 bool) string {
    v7 := GenS1{F0: 1}
    var v8 bool = p0
    p1 = false
    var v9 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return ""
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]uint32)
    v3 := []string{"\n\t", ("" + ("a" + (("" + ("\n\t" + ("hello" + "foo bar"))) + "a"))), "foo bar"}
    var v4 float32 = 3.14159
    var v5 uint8 = 51
    v6 := []uint32{988, 1}
    v4 = (v4 + (3.14159 - (v4 * (-2.25 * 1.5))))
    for fi10 := 0; fi10 < 9; fi10++ {
        v5 = 22
        var v11 int32 = -830
        _ = fi10
        _ = v11
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v12 uint16 = 32767
    var v13 int32 = 42
    v14 := GenS1{F0: (^(v1))}
    var v15 string = ("" + "hello")
    if false {
        fmt.Println(v13, v12, v12)
        if true {
            v13 = (v13 >> ((255) & 63))
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 int16 = -32768
                var v18 float64 = 1.5
                if (false && false) {
                    var v19 int64 = 42
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 int64 = (-(v19))
                        v22 := GenS1{F0: (v1 - v1)}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    _ = v19
                } else {
                    if (v5 == v5) {
                        v12 += v12
                    }
                    var v23 float64 = v18
                    fmt.Println(v13)
                    _ = v23
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            switch 42 {
            case -304:
                var v24 uint32 = 652
                _ = v24
            case 89:
                var v25 string = (v15 + (v15 + (v15 + (v15 + v15))))
                _ = v25
            default:
                var v26 int16 = 0
                _ = v26
            }
        }
        switch 9223372036854775807 {
        case -346:
            var v27 int8 = v1
            _ = v27
        }
    } else {
        v3 = append(v3, v15)
    }
    fmt.Println(v4, v4)
    switch (v15 + "\n\t") {
    case "foo bar":
        v13 += -2147483648
    case "hello":
        var v28 uint = 42
        _ = v28
    case "a":
        v29 := GenS2{F0: 556, F1: ("foo bar" + v15), F2: -171}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
