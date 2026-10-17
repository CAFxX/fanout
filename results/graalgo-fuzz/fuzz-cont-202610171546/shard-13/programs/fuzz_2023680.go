package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 bool; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int16; F2 string }
type GenS5 struct { F0 bool }
func genFn0(p0 int32, p1 uint8) uint64 {
    if false {
        switch 599 {
        case 32768:
            var v5 int = 1
            _ = v5
        case 0:
            if (2 < 0) {
                fmt.Println(p1)
            } else {
                v6 := GenS1{F0: "\n\t"}
                _ = v6
            }
        case 1:
            v7 := []int{101, 127, 2147483647, -1}
            _ = v7
        }
        v8 := GenS1{F0: "a"}
        _ = v8
    }
    var v9 float64 = (1e10 * 1.5)
    _ = p0
    _ = p1
    _ = v9
    return 42
}
func genFn1(p0 uint32) int32 {
    var v10 int8 = -1
    v11 := GenS1{F0: "a"}
    _ = p0
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 int16 = -1
    var v2 uint = 916
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    v12 := GenS2{F0: -686, F1: (((1e10 == ((-2.25 + (0.1 + 0.1)) + 0.1)) || (-2.25 > 0.0)) || ((0.0 * 3.14159) == (1.5 - (3.14159 - (3.14159 / (((((3.14159 + ((1e10 * ((1.5 + 0.0) - 1e10)) / 0.1)) * 0.1) / 0.1) / ((1.5 - (0.1 / 0.1)) * (-2.25 + -2.25))) / 0.1)))))), F2: "\n\t"}
    if false {
        if (162 >= 2147483647) {
            v13 := v3[("\n\t" + "\n\t")]
            switch 968 {
            case 373:
                var v14 uint8 = 0
                _ = v14
            }
            var v15 uint64 = 42
            _ = v13
            _ = v15
        }
        delete(v3, "\n\t")
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        switch "hello" {
        case "a":
            v17 := GenS1{F0: ""}
            _ = v17
        default:
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 uint16 = 170
                fmt.Println(v1, v2)
                _ = fi18
                _ = v19
            }
        }
        _ = fi16
    }
    v20 := GenS1{F0: "hello"}
    if true {
        var v21 string = "\n\t"
        v22 := GenS3{F0: (v1 + (v1 % (((v1 & (^(v1)))) | 1)))}
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 19; fi23++ {
            fmt.Println(v2, v1)
            if true {
                delete(v4, ("" + ("\n\t" + "")))
                v2 -= v2
            } else {
                if (!true) {
                    v24 := GenS3{F0: (^(v1))}
                    v25 := GenS4{F0: (-(fi23)), F1: (-(v1)), F2: "foo bar"}
                    if true {
                        fmt.Println(fi23, fi23, v1)
                        switch (fi23 * fi23) {
                        case 128:
                            fmt.Println(fi23, v2)
                        }
                    }
                    _ = v24
                    _ = v25
                } else {
                    v26 := GenS5{F0: (((124 >= 1) || true) && false)}
                    var v27 string = "foo bar"
                    _ = v26
                    _ = v27
                }
            }
            var v28 float64 = 0.0
            _ = fi23
            _ = v28
        }
        v2 += (v2 * (v2 - 2))
        var v29 int64 = 9223372036854775807
        _ = v29
    }
    fmt.Println(v1)
    v1 = (-(v1))
    _ = len(v4)
    delete(v4, ("\n\t" + "\n\t"))
    fmt.Println(v1, v1)
    var v30 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
}
