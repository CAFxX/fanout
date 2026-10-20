package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int16; F2 uint8 }
func genFn0() uint {
    v5 := GenS1{F0: true, F1: ("a" + "hello"), F2: 1}
    var v6 string = "\n\t"
    v7 := GenS1{F0: (1 < 2), F1: "foo bar", F2: 1}
    fmt.Println(v6)
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := []uint32{}
    var v2 bool = false
    v3 := []uint32{308, 986}
    var v4 int = 42
    switch v4 {
    case -9223372036854775808:
        var v8 bool = (v4 == (v4 << ((481) & 63)))
        _ = v8
    case 340:
        v2 = ((-32768 == 1000) && v2)
    }
    var v9 int32 = -1
    for fi10 := 0; fi10 < 8; fi10++ {
        v9 -= (^(v9))
        _ = fi10
    }
    if len(v3) > 0 {
        v3[0] = 996
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        if v2 {
            fmt.Println(v4)
        }
        v12 := GenS1{F0: v2, F1: "hello", F2: (v9 & v9)}
        switch (fi11 * v4) {
        case 32767:
            var v13 uint8 = 255
            _ = v13
        case 1:
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 int16 = -435
                v4 = (fi11 % ((-9223372036854775808) | 1))
                _ = fi14
                _ = v15
            }
        default:
            if (!v2) {
                v16 := GenS1{F0: ((167 <= 22) || v2), F1: "foo bar", F2: 0}
                if v2 {
                    var v17 int32 = (^(v9))
                    _ = v17
                }
                _ = v16
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v18 int64 = -395
                v18 = (^(v18))
                _ = v18
            }
        }
        _ = fi11
        _ = v12
    }
    v19 := GenS1{F0: (!v2), F1: ((("" + (("hello" + "a") + ("hello" + "hello"))) + "hello") + "hello"), F2: (v9 + v9)}
    var v20 int8 = 42
    var v21 uint32 = 2
    if true {
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS2{F0: "hello"}
            _ = fi22
            _ = v23
        }
    }
    v1 = append(v1, v21)
    var v24 bool = (1.5 != float64((float32(float64(1e10)) + 0.0)))
    var v25 float64 = 0.0
    if (1 != 255) {
        v20 -= v20
    } else {
        var v26 uint32 = (^(v21))
        v27 := GenS3{F0: v26, F1: -1, F2: 127}
        _ = v26
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
