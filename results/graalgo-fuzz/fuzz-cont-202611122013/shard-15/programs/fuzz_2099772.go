package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int32 }
func genFn0(p0 string) int8 {
    p0 = p0
    if (!true) {
        fmt.Println(p0)
        v4 := GenS1{F0: 569, F1: 0, F2: 909}
        var v5 uint = 65535
        _ = v4
        _ = v5
    } else {
        var v6 int = -390
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int64 = 1
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    _ = p0
    return 42
}
func genFn1() string {
    v9 := []int{}
    var v10 uint64 = 1
    var v11 int16 = 1
    if (0 < 824) {
        v12 := GenS1{F0: 42, F1: 1, F2: 4294967295}
        _ = v12
    }
    _ = v10
    _ = v11
    _ = v9
    return ("foo bar" + "\n\t")
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    var v3 uint = 42
    delete(v2, 564)
    var v13 bool = (128 >= 128)
    v14 := GenS1{F0: (v3 & 1), F1: uint32(v3), F2: 0}
    if v13 {
        var v15 string = ("a" + ("foo bar" + ""))
        v16 := []int{-1, 790}
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS2{F0: (fi17 >> ((18446744073709551615) & 63)), F1: (v15 + "hello")}
            if v13 {
                v19 := GenS2{F0: fi17, F1: v15}
                for fi20 := 0; fi20 < 5; fi20++ {
                    var v21 string = "hello"
                    fmt.Println(v15, fi17)
                    _ = fi20
                    _ = v21
                }
                _ = v19
            } else {
                var v22 uint = 18446744073709551615
                _ = v22
            }
            _ = fi17
            _ = v18
        }
        _ = v15
        _ = v16
    }
    var v23 int = 4294967295
    _ = len(v2)
    var v24 int8 = 0
    if (v3 != v3) {
        var v25 int = 878
        var v26 uint = 4294967295
        for fi27 := 0; fi27 < 6; fi27++ {
            if ((9223372036854775807 < 946) && false) {
                var v28 int16 = 32767
                if (0.1 == float64(0.0)) {
                    for fi29 := 0; fi29 < 10; fi29++ {
                        v26 = 18446744073709551615
                        _ = fi29
                    }
                }
                _ = v28
            } else {
                var v30 float32 = (0.0 - 3.14159)
                fmt.Println(v26)
                _ = v30
            }
            var v31 uint32 = 579
            for fi32 := 0; fi32 < 7; fi32++ {
                v33 := GenS3{F0: 1}
                var v34 string = "\n\t"
                v23 = v23
                v24 -= (^(v24))
                fmt.Println(v23, v23, v34)
                switch -69 {
                case 32768:
                    v35 := GenS4{F0: 128, F1: 444}
                    _ = v35
                default:
                    var v36 int8 = (v24 * (v24 / ((v24) | 1)))
                    _ = v36
                }
                var v37 string = "hello"
                v13 = (9223372036854775807 != 42)
                var v38 int32 = -1
                _ = fi32
                _ = v33
                _ = v34
                _ = v37
                _ = v38
            }
            _ = fi27
            _ = v31
        }
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
