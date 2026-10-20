package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 uint; F2 string }
func genFn0() string {
    v5 := []int{0, 42, -9223372036854775808}
    v6 := GenS1{F0: 830, F1: 42}
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            if (false || (-9223372036854775808 == 628)) {
                if true {
                    var v8 bool = false
                    _ = v8
                }
            }
            _ = fi7
        }
        var v9 int32 = -204
        v9 = v9
        _ = v9
    } else {
        v10 := GenS1{F0: 682, F1: 65535}
        if true {
            var v11 string = "a"
            _ = v11
        }
        _ = v10
    }
    _ = v5
    _ = v6
    return "hello"
}
func genFn1(p0 string, p1 int16, p2 uint16) uint16 {
    p1 += (-(p1))
    switch -463 {
    case 42:
        fmt.Println(p1, p2, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = (0.1 * 1e10)
    var v3 float64 = v2
    v4 := []string{"hello", "\n\t"}
    v12 := GenS2{F0: 73, F1: 513, F2: 1}
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS3{F0: -19}
        v3 -= float64(1e10)
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS3{F0: 42}
            switch (fi13 / ((fi13) | 1)) {
            case -2147483648:
                var v17 bool = (362 > 633)
                _ = v17
            case 32767:
                v4 = append(v4, "hello")
            default:
                if true {
                    if true {
                        v3 = ((3.14159 / v3) / -2.25)
                    }
                } else {
                    v18 := GenS4{F0: 42, F1: 1, F2: ""}
                    var v19 float32 = 0.0
                    if true {
                        v19 = float32(1e10)
                        var v20 string = ""
                        _ = v20
                    } else {
                        v2 = v3
                    }
                    _ = v18
                    _ = v19
                }
            }
            switch "\n\t" {
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v3 -= v3
            case "hello":
                v21 := GenS3{F0: 41}
                _ = v21
            default:
                v3 += v3
            }
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    v3 = float64(1.5)
    fmt.Println(v3, v3)
    var v22 uint32 = 65535
    _ = len(v1)
    var v23 float64 = v3
    for fi24 := 0; fi24 < 8; fi24++ {
        if (float32(v3) == -2.25) {
            v3 = ((v2 - ((1e10 + v2) - v3)) - v2)
        }
        _ = fi24
    }
    v22 = (v22 / ((v22) | 1))
    v25 := GenS1{F0: 42, F1: 1}
    v26 := GenS2{F0: 42, F1: 18446744073709551615, F2: 196}
    for fi27 := 0; fi27 < 6; fi27++ {
        v3 += v3
        _ = fi27
    }
    v23 -= float64((1e10 + 0.1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
