package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 bool }
func genFn0() int {
    var v4 int32 = 1
    v5 := []int{42, -9223372036854775808, -9223372036854775808}
    if true {
        var v6 uint16 = 117
        v6 = (^(v6))
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: 443, F1: -9223372036854775808, F2: v4}
            switch "\n\t" {
            case "\n\t":
                v6 -= v6
            }
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    switch "hello" {
    case "hello":
        var v9 int64 = 615
        _ = v9
    }
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = float64(1e10)
    var v3 int8 = 100
    v2 += v2
    v10 := GenS1{F0: 32768, F1: int64(v3), F2: -1}
    var v11 int = -536
    if ((!(("a" + (("a" + ((("a" + "hello") + ("a" + "")) + (((("hello" + "\n\t") + "a") + "") + ("\n\t" + "")))) + "foo bar")) != "\n\t")) && (true && false)) {
        var v12 float32 = -2.25
        v13 := []int{(v11 / ((v11) | 1))}
        _ = v12
        _ = v13
    } else {
        if true {
            var v14 uint64 = 4294967295
            v15 := GenS2{F0: 2, F1: 194, F2: 0}
            v11 = (v11 & 32767)
            _ = v14
            _ = v15
        }
    }
    v2 += (((v2 + v2) + (v2 + v2)) * v2)
    switch "a" {
    case "\n\t":
        for fi16 := 0; fi16 < 19; fi16++ {
            v11 -= (fi16 - (v11 << ((993) & 63)))
            var v17 uint = 18446744073709551615
            _ = fi16
            _ = v17
        }
    case "":
        v18 := GenS3{F0: 127, F1: 4294967295}
        _ = v18
    }
    v19 := GenS4{F0: (42 >= 255)}
    fmt.Println(v3)
    v20 := []int{(-(v11))}
    v2 = v2
    if (false && false) {
        if true {
            v3 = (v3 & v3)
            var v21 float32 = (-2.25 / (1.5 / 1e10))
            _ = v21
        }
        switch 9223372036854775807 {
        case 128:
            var v22 int = 32768
            _ = v22
        case 624:
            if false {
                var v23 int = (v11 * (v11 << ((151) & 63)))
                if true {
                    v24 := GenS3{F0: 177, F1: 913}
                    _ = v24
                }
                _ = v23
            } else {
                var v25 int64 = -340
                _ = v25
            }
        default:
            var v26 uint64 = 0
            _ = v26
        }
    }
    v11 += v11
    var v27 string = "hello"
    fmt.Println(v11)
    switch (v11 / ((v11) | 1)) {
    case 32768:
        v2 = (float64(0.1) / (3.14159 / (float64((0.0 / (((0.0 * float32(v2)) * (1.5 / 1.5)) - (1e10 + 0.1)))) - float64(((-2.25 * -2.25) - 0.1)))))
    default:
        var v28 int = (^(v11))
        _ = v28
    }
    if true {
        for fi29 := 0; fi29 < 10; fi29++ {
            v27 = "\n\t"
            _ = fi29
        }
        var v30 int64 = -1
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
}
