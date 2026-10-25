package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint8; F2 int8 }
type GenS4 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint16) int16 {
    var v5 float64 = (0.0 * 0.0)
    switch "" {
    case "\n\t":
        var v6 bool = (8 > 0)
        _ = v6
    case "hello":
        var v7 uint = 18446744073709551615
        _ = v7
    case "a":
        p0 -= 2
    default:
        fmt.Println(v5, v5)
    }
    _ = p0
    _ = v5
    return 1000
}

func main() {
    var v1 int8 = -128
    var v2 float64 = ((0.0 - (1.5 * (1.5 * 0.1))) - 3.14159)
    v3 := make(map[string]int)
    var v4 int = 414
    v8 := GenS1{F0: 0}
    if (-1 == -2147483648) {
        var v9 uint = 4294967295
        _ = v9
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := v3[""]
            _ = fi10
            _ = v11
        }
        v2 = ((v2 + v2) / v2)
        if (0 != -2147483648) {
            delete(v3, "foo bar")
        } else {
            v12 := GenS1{F0: 431}
            v13 := GenS1{F0: 42}
            v14 := GenS2{F0: -1}
            _ = v12
            _ = v13
            _ = v14
        }
    }
    v3[("a" + "a")] = v4
    v3[("hello" + ("" + ((("" + "hello") + ("\n\t" + "hello")) + ("a" + "\n\t"))))] = -452
    fmt.Println(v2, v4, v4)
    if ((v1 - v1) <= v1) {
        switch "" {
        case "":
            var v15 float32 = 3.14159
            _ = v15
        case "a":
            if false {
                for fi16 := 0; fi16 < 8; fi16++ {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        switch (^(fi17)) {
                        case 856:
                            if (326 >= 37) {
                                var v18 int = -2147483648
                                v19 := GenS1{F0: 0}
                                var v20 bool = true
                                _ = v18
                                _ = v19
                                _ = v20
                            } else {
                                fmt.Println(fi16, v2, v1)
                                var v21 int32 = 2147483647
                                var v22 int = (fi17 ^ 32768)
                                _ = v21
                                _ = v22
                            }
                            var v23 uint8 = 127
                            var v24 int8 = v1
                            _ = v23
                            _ = v24
                        default:
                            v25 := GenS2{F0: -322}
                            _ = v25
                        }
                        var v26 float64 = v2
                        v1 = (v1 - v1)
                        _ = fi17
                        _ = v26
                    }
                    _ = fi16
                }
                for fi27 := 0; fi27 < 18; fi27++ {
                    if false {
                        v4 = 127
                        v2 += v2
                        v2 = ((0.1 - 1.5) / v2)
                    }
                    _ = fi27
                }
            } else {
                v28 := GenS3{F0: (v4 ^ v4), F1: 128, F2: v1}
                _ = v28
            }
        }
        var v29 uint32 = 65535
        v1 += v1
        _ = v29
    } else {
        var v30 bool = true
        var v31 string = ""
        _ = len(v3)
        _ = v30
        _ = v31
    }
    _ = len(v3)
    v2 += float64((0.1 * 0.1))
    switch (v4 & (-(v4))) {
    case 4294967295:
        v32 := GenS4{F0: 0, F1: 2}
        _ = v32
    }
    var v33 bool = (!(2147483647 != 2))
    switch (v4 | -678) {
    case -1:
        v34 := GenS3{F0: 2147483647, F1: 42, F2: (v1 / ((0) | 1))}
        _ = v34
    case 0:
        var v35 int8 = 1
        _ = v35
    }
    fmt.Println(v33)
    v36 := GenS3{F0: (v4 - v4), F1: 175, F2: (v1 >> ((255) & 63))}
    for fi37 := 0; fi37 < 2; fi37++ {
        v38 := []int{fi37, fi37}
        var v39 string = ((("\n\t" + ("hello" + ((("" + "a") + "foo bar") + "\n\t"))) + "a") + "")
        v4 = 2147483647
        _ = fi37
        _ = v38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v8
}
