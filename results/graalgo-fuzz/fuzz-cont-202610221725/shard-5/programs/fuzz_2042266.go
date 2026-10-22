package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint64 }
type GenS2 struct { F0 int64; F1 string }

func main() {
    var v1 int64 = 1
    var v2 float64 = (1.5 * (0.1 + ((3.14159 * 0.0) + ((-2.25 - -2.25) * (0.0 + 1.5)))))
    v3 := []string{"foo bar", ((("\n\t" + "") + "hello") + ((("hello" + "\n\t") + ("foo bar" + "a")) + "\n\t")), "", "\n\t"}
    var v4 int = 128
    var v5 int32 = 926
    v6 := []int{v4, 379}
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 0, F1: v5, F2: 53}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: (v5 + (v5 << ((65535) & 63))), F2: 4294967295}
    for fi10 := 0; fi10 < 4; fi10++ {
        v6 = append(v6, v4)
        var v11 uint32 = 717
        switch "\n\t" {
        case "foo bar":
            for fi12 := 0; fi12 < 15; fi12++ {
                v5 -= v5
                v2 -= v2
                _ = fi12
            }
        default:
            v3 = append(v3, ("foo bar" + "a"))
        }
        _ = fi10
        _ = v11
    }
    if true {
        var v13 uint64 = 694
        if (1 >= 1) {
            if true {
                _ = len(v3)
                var v14 int64 = 1
                _ = v14
            }
        }
        _ = v13
    }
    if ((v1 * (v1 >> ((2) & 63))) < -618) {
        v15 := GenS2{F0: (v1 & (-(v1))), F1: ("" + "a")}
        _ = v15
    } else {
        v16 := GenS2{F0: v1, F1: "foo bar"}
        _ = v16
    }
    if (int16(v1) < 43) {
        var v17 uint = 1
        _ = v17
    }
    var v18 uint16 = 65535
    var v19 bool = false
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := GenS2{F0: -464, F1: "hello"}
        switch "hello" {
        case "":
            v4 -= -319
        }
        _ = fi20
        _ = v21
    }
    v22 := GenS2{F0: 601, F1: "hello"}
    var v23 int16 = int16(v18)
    if v19 {
        var v24 uint = 65535
        v5 = -1
        var v25 uint8 = 89
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 19; fi26++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
