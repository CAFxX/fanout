package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint8) uint16 {
    v6 := GenS1{F0: 42, F1: 0, F2: -1}
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int64 = 530
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        var v10 int32 = 876
        _ = v10
    } else {
        v11 := []int{0, -210, 127}
        p0 = (p0 + 267)
        _ = v11
    }
    var v12 int32 = -705
    v13 := GenS1{F0: 1, F1: -83, F2: (v12 | v12)}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v6
    return 470
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 1
    var v3 string = ""
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 255
    if (((1e10 - (float64(-2.25) - (0.1 + 0.1))) < 3.14159) || false) {
        switch (("a" + ("foo bar" + (v3 + "foo bar"))) + "foo bar") {
        case "a":
            for fi14 := 0; fi14 < 2; fi14++ {
                fmt.Println(fi14)
                switch "" {
                case "hello":
                    var v15 int64 = 620
                    _ = v15
                case "foo bar":
                    v2 = v2
                case "":
                    var v16 uint16 = v5
                    _ = v16
                default:
                    var v17 bool = (v3 == v3)
                    _ = v17
                }
                fmt.Println(v5, v3)
                _ = fi14
            }
        case "foo bar":
            var v18 uint8 = 134
            _ = v18
        case "":
            var v19 uint32 = v2
            _ = v19
        default:
            for fi20 := 0; fi20 < 4; fi20++ {
                if (!true) {
                    var v21 uint32 = 0
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        var v23 int16 = 42
                        v2 += v2
                        v24 := GenS1{F0: (v23 | (^(v23))), F1: -59, F2: -1}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    v25 := []int{(fi20 ^ fi20), 128}
                    _ = v25
                }
                var v26 int64 = -1
                _ = fi20
                _ = v26
            }
        }
        for fi27 := 0; fi27 < 12; fi27++ {
            v28 := GenS1{F0: 509, F1: -128, F2: 42}
            _ = fi27
            _ = v28
        }
        switch 9223372036854775807 {
        case 2147483647:
            var v29 int = 32767
            _ = v29
        default:
            var v30 int = 489
            _ = v30
        }
    }
    v31 := v4[v2]
    if (!false) {
        v32 := GenS2{F0: 42, F1: v3}
        _ = v32
    } else {
        v2 -= 649
    }
    var v33 uint = 1
    var v34 string = (v3 + (("hello" + ((("hello" + v3) + ("hello" + (v3 + v3))) + v3)) + v3))
    v35 := GenS3{F0: 244}
    switch -251 {
    case 4294967295:
        v3 = "\n\t"
    case -929:
        var v36 bool = true
        _ = v36
    }
    v5 = (v5 << ((v33) & 63))
    for fi37 := 0; fi37 < 19; fi37++ {
        for fi38 := 0; fi38 < 16; fi38++ {
            v31 -= (v31 % (((v31 ^ v31)) | 1))
            if (-128 != -126) {
                v39 := GenS1{F0: -698, F1: -99, F2: 985}
                _ = v39
            } else {
                var v40 uint32 = v31
                v41 := GenS3{F0: uint8(v33)}
                _ = v40
                _ = v41
            }
            _ = fi38
        }
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
