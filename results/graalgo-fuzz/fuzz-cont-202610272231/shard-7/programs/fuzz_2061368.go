package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int16; F1 int64 }
func genFn0(p0 int, p1 uint32) int8 {
    var v5 int32 = 991
    switch (p0 >> ((0) & 63)) {
    case -1:
        if (0 == 500) {
            fmt.Println(p1)
            if (!(p0 < 4294967295)) {
                v5 -= (v5 + (v5 - (v5 & v5)))
            } else {
                fmt.Println(p0, v5, p1)
                v6 := GenS1{F0: "hello", F1: 42}
                if (42 >= -128) {
                    for fi7 := 0; fi7 < 5; fi7++ {
                        switch "\n\t" {
                        case "a":
                            if ("hello" < "foo bar") {
                                for fi8 := 0; fi8 < 20; fi8++ {
                                    var v9 int64 = -1
                                    _ = fi8
                                    _ = v9
                                }
                            } else {
                                v10 := GenS1{F0: "hello", F1: 855}
                                _ = v10
                            }
                        case "hello":
                            var v11 uint16 = 65535
                            _ = v11
                        case "\n\t":
                            for fi12 := 0; fi12 < 15; fi12++ {
                                var v13 uint64 = 0
                                _ = fi12
                                _ = v13
                            }
                        }
                        v5 = 2147483647
                        _ = fi7
                    }
                    v14 := []int{}
                    var v15 int = p0
                    _ = v14
                    _ = v15
                }
                _ = v6
            }
            for fi16 := 0; fi16 < 8; fi16++ {
                v5 += (v5 * (^(v5)))
                _ = fi16
            }
        }
    case 2147483647:
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 uint8 = 255
            _ = fi17
            _ = v18
        }
    default:
        var v19 uint64 = 851
        _ = v19
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        p0 += -9223372036854775808
        if (int16(fi20) >= 1) {
            var v21 int8 = 0
            v22 := []int{fi20, -9223372036854775808, fi20}
            fmt.Println(p1)
            _ = v21
            _ = v22
        } else {
            v23 := []int{(fi20 >> ((4294967295) & 63)), p0}
            _ = v23
        }
        _ = fi20
    }
    var v24 int32 = 42
    v25 := []int{p0, (p0 % (((p0 >> ((4294967295) & 63))) | 1)), -9223372036854775808}
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    _ = v5
    return 71
}
func genFn1(p0 int64, p1 uint32) bool {
    var v26 float64 = (1.5 * (-2.25 * ((0.0 + 1.5) - 0.0)))
    var v27 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return true
}

func main() {
    var v1 float32 = 1e10
    var v2 int8 = -23
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    fmt.Println(v2, v2)
    var v28 int16 = 32767
    v28 = v28
    v29 := []int{32767, -2147483648}
    v30 := GenS1{F0: "hello", F1: int64(v2)}
    v1 += v1
    var v31 uint = 65535
    var v32 uint32 = 880
    for fi33 := 0; fi33 < 7; fi33++ {
        v1 = 1.5
        _ = fi33
    }
    fmt.Println(v2, v32, v28)
    for fi34 := 0; fi34 < 3; fi34++ {
        v32 = 2147483647
        _ = fi34
    }
    v35 := GenS1{F0: ("" + ("a" + "\n\t")), F1: 882}
    if (61 < 2) {
        fmt.Println(v2)
        for fi36 := 0; fi36 < 2; fi36++ {
            v37 := GenS2{F0: (-(v28)), F1: 1}
            _ = fi36
            _ = v37
        }
    }
    var v38 int64 = 1
    if (false && (v1 >= v1)) {
        var v39 int8 = v2
        _ = v39
    }
    v1 = float32(((1.5 - float64(v1)) / 1.0))
    var v40 uint8 = 179
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v38
    _ = v4
    _ = v40
}
