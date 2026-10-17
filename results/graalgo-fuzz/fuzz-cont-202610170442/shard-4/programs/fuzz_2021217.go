package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int8; F1 bool; F2 int64 }
func genFn0() int {
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint32 = 0
            var v10 int = fi7
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := []int{fi7, (fi7 & 446)}
        var v12 uint32 = 65535
        _ = fi7
        _ = v11
        _ = v12
    }
    v13 := []int{}
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int = fi14
        v13 = append(v13, (^(v15)))
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := GenS1{F0: true}
            var v18 uint8 = 79
            fmt.Println(fi16, fi16, v18)
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    var v19 uint16 = 42
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := GenS2{F0: 1, F1: -1}
        if true {
            v22 := GenS3{F0: 805, F1: 42}
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    _ = v13
    _ = v19
    return 0
}

func main() {
    var v1 bool = ((0.0 - -2.25) <= -2.25)
    v2 := make(map[string]uint32)
    var v3 uint = 65535
    v4 := make(map[string]uint32)
    var v5 int = 1
    var v6 uint = (v3 / (((v3 << (((v3 ^ 825)) & 63))) | 1))
    var v23 string = ((("\n\t" + "\n\t") + ("hello" + "\n\t")) + "")
    for fi24 := 0; fi24 < 18; fi24++ {
        v5 -= -415
        _ = fi24
    }
    v5 += (v5 | v5)
    delete(v2, (((v23 + (v23 + v23)) + "") + "foo bar"))
    v5 = genFn0()
    for fi25 := 0; fi25 < 8; fi25++ {
        v5 = (-(v5))
        _ = fi25
    }
    if (v1 || v1) {
        v26 := GenS2{F0: 771, F1: 2147483647}
        _ = v26
    } else {
        switch 42 {
        case 42:
            var v27 float64 = (-2.25 * 3.14159)
            _ = v27
        case 127:
            if false {
                v28 := v4[v23]
                v29 := GenS1{F0: ((v6 << ((v3) & 63)) <= (^(v3)))}
                switch (^(v5)) {
                case -32:
                    if v1 {
                        var v30 int16 = 643
                        _ = v30
                    } else {
                        var v31 uint16 = 0
                        v32 := GenS2{F0: 389, F1: -2147483648}
                        _ = v31
                        _ = v32
                    }
                case -563:
                    fmt.Println(v5, v28, v6)
                }
                _ = v28
                _ = v29
            }
        default:
            delete(v2, ((("a" + (v23 + v23)) + v23) + v23))
        }
    }
    fmt.Println(v23)
    v3 += v6
    if (0 == 2) {
        fmt.Println(v23, v5)
        v33 := GenS4{F0: -128, F1: v1, F2: 854}
        _ = v33
    }
    v23 = (v23 + (v23 + v23))
    for fi34 := 0; fi34 < 7; fi34++ {
        v5 = (v5 % ((v5) | 1))
        _ = fi34
    }
    switch (^(v5)) {
    case -1:
        var v35 int64 = 9223372036854775807
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
