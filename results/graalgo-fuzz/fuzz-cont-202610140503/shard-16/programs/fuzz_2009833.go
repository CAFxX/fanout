package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 bool; F1 string }
func genFn0() int {
    for fi4 := 0; fi4 < 4; fi4++ {
        fmt.Println(fi4)
        fmt.Println(fi4)
        var v5 uint8 = 127
        _ = fi4
        _ = v5
    }
    var v6 uint = 0
    if (-128 > -56) {
        v7 := GenS1{F0: 65535, F1: 42, F2: 255}
        fmt.Println(v6)
        _ = v7
    }
    var v8 string = "a"
    _ = v6
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{4294967295, 42, 42}
    var v2 string = "\n\t"
    var v3 float32 = 0.0
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    v9 := GenS1{F0: 18446744073709551615, F1: 52, F2: 56}
    v10 := GenS1{F0: 18446744073709551615, F1: 0, F2: 577}
    if (65535 < 449) {
        for fi11 := 0; fi11 < 14; fi11++ {
            for fi12 := 0; fi12 < 20; fi12++ {
                if (1 != 58) {
                    v2 = v2
                } else {
                    if false {
                        var v13 int8 = -13
                        _ = v13
                    } else {
                        var v14 uint32 = 0
                        v2 = v2
                        if len(v1) > 0 {
                            v1[0] = v14
                        }
                        _ = v14
                    }
                }
                v15 := GenS1{F0: 1, F1: -128, F2: 175}
                switch (-(fi11)) {
                case -147:
                    var v16 float64 = 3.14159
                    _ = v16
                default:
                    v17 := GenS1{F0: 4294967295, F1: 100, F2: 52}
                    _ = v17
                }
                fmt.Println(v3, v2)
                _ = fi12
                _ = v15
            }
            switch 4294967295 {
            case 636:
                var v18 int8 = 1
                _ = v18
            }
            _ = fi11
        }
        var v19 uint32 = 1
        var v20 string = "hello"
        _ = v19
        _ = v20
    }
    v3 = (v3 - (0.0 / 0.1))
    var v21 uint32 = 2147483648
    var v22 uint64 = 1
    if true {
        v22 -= 18446744073709551615
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        switch fi23 {
        case 2147483647:
            for fi24 := 0; fi24 < 14; fi24++ {
                v25 := GenS2{F0: (false && ((v22 & v22) != v22)), F1: "hello"}
                _ = fi24
                _ = v25
            }
        default:
            _ = len(v1)
        }
        fmt.Println(v3)
        _ = fi23
    }
    var v26 int64 = 407
    if false {
        v27 := GenS1{F0: 257, F1: -102, F2: 0}
        _ = v27
    }
    var v28 int64 = (v26 >> ((1) & 63))
    v26 -= (^(v26))
    v21 -= (^(v21))
    v28 = v28
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v28
    _ = v3
    _ = v9
}
