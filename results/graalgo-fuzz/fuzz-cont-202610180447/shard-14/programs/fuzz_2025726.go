package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 string; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int; F1 bool }
func genFn0(p0 int, p1 int8, p2 int64) int64 {
    v4 := GenS1{F0: 42, F1: 874, F2: 255}
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p2
}

func main() {
    var v1 uint8 = 2
    var v2 string = ""
    var v3 uint64 = 1
    var v5 float32 = (0.0 - ((1.5 + ((1.5 - 1.5) + (0.0 / 1.0))) + 0.1))
    var v6 uint8 = v1
    switch v2 {
    case "hello":
        v7 := []int{32768}
        _ = v7
    case "foo bar":
        if (v5 != 1.5) {
            var v8 uint8 = 0
            _ = v8
        }
    case "\n\t":
        switch "" {
        case "a":
            v1 = v1
        default:
            var v9 uint = 255
            _ = v9
        }
    default:
        if true {
            var v10 float32 = v5
            _ = v10
        } else {
            v11 := GenS2{F0: (v2 + (v2 + v2)), F1: (v3 * (v3 | (v3 - (v3 / (((v3 + (v3 + v3))) | 1))))), F2: 2}
            _ = v11
        }
    }
    v12 := []int{-1, -9223372036854775808}
    if true {
        var v13 bool = true
        _ = v13
    } else {
        v12 = append(v12, 127)
        switch v2 {
        case "a":
            if (0.1 < 1.5) {
                _ = len(v12)
                if false {
                    _ = len(v12)
                    fmt.Println(v6, v1, v1)
                    if false {
                        v12 = append(v12, 32767)
                        var v14 float32 = v5
                        var v15 int8 = -70
                        _ = v14
                        _ = v15
                    }
                }
                v5 = v5
            } else {
                fmt.Println(v3)
            }
        case "":
            for fi16 := 0; fi16 < 16; fi16++ {
                for fi17 := 0; fi17 < 18; fi17++ {
                    switch (v2 + (v2 + v2)) {
                    case "hello":
                        var v18 int16 = 32767
                        _ = v18
                    case "":
                        if true {
                            var v19 int16 = 32767
                            if true {
                                switch (fi17 >> ((65535) & 63)) {
                                case 834:
                                    v5 -= (v5 * 0.1)
                                }
                                var v20 uint64 = (^(v3))
                                _ = v20
                            }
                            _ = v19
                        }
                    default:
                        var v21 int = -2147483648
                        fmt.Println(fi16, fi16, v2)
                        _ = v21
                    }
                    var v22 int16 = 999
                    var v23 int64 = -949
                    _ = fi17
                    _ = v22
                    _ = v23
                }
                _ = fi16
            }
        case "hello":
            v2 = ("\n\t" + "a")
        default:
            var v24 uint64 = 374
            _ = v24
        }
    }
    switch (v2 + "a") {
    case "a":
        v5 -= 1e10
    }
    switch 128 {
    case 32767:
        v1 = v1
    case 796:
        v25 := GenS1{F0: 969, F1: v3, F2: uint16(v1)}
        _ = v25
    case 0:
        var v26 float64 = 0.0
        _ = v26
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        v28 := GenS1{F0: 9223372036854775807, F1: v3, F2: 255}
        v1 -= 0
        _ = fi27
        _ = v28
    }
    var v29 uint64 = v3
    for fi30 := 0; fi30 < 1; fi30++ {
        if true {
            v31 := GenS3{F0: (fi30 ^ fi30), F1: (4294967295 <= 4294967295)}
            for fi32 := 0; fi32 < 12; fi32++ {
                v2 = (v2 + ((v2 + v2) + "hello"))
                var v33 int32 = 2147483647
                fmt.Println(v5, v5)
                if true {
                    var v34 uint8 = (v1 % ((188) | 1))
                    v35 := GenS1{F0: -1, F1: 0, F2: 433}
                    _ = v34
                    _ = v35
                } else {
                    v36 := GenS1{F0: 0, F1: (v29 + 293), F2: 562}
                    switch fi30 {
                    case 32768:
                        var v37 int16 = -154
                        _ = v37
                    default:
                        var v38 float32 = 3.14159
                        _ = v38
                    }
                    _ = v36
                }
                _ = fi32
                _ = v33
            }
            v39 := GenS2{F0: ("foo bar" + ("a" + ("a" + (("foo bar" + v2) + v2)))), F1: (v3 * v3), F2: 65535}
            _ = v31
            _ = v39
        } else {
            v40 := GenS1{F0: genFn0(-9223372036854775808, 100, -9223372036854775808), F1: (v29 & v29), F2: 42}
            v29 += v29
            _ = v40
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v5
    _ = v6
}
