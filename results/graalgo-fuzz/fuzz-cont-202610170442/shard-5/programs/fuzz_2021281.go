package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 string }
type GenS2 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint }
func genFn0(p0 uint) uint16 {
    v5 := []int{int(p0), 516}
    var v6 uint8 = 255
    var v7 float64 = 1.5
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 int16) uint64 {
    switch 42 {
    case -1:
        p0 = (p0 - (p0 + 1000))
    case 254:
        var v8 uint64 = 18446744073709551615
        _ = v8
    case 42:
        p0 += p0
    }
    v9 := GenS1{F0: 42, F1: (p0 * p0), F2: "a"}
    var v10 int64 = 699
    _ = p0
    _ = v10
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 0.0
    v3 := make(map[string]string)
    var v4 float32 = v2
    var v11 string = "\n\t"
    var v12 float64 = 1e10
    v13 := []int{}
    var v14 uint32 = 2
    v2 = v4
    v15 := v1[42]
    if (4294967295 <= -782) {
        if true {
            if (1 != -2147483648) {
                switch 128 {
                case -1:
                    v4 = v2
                case 32767:
                    var v16 int64 = 0
                    _ = v16
                case 2147483647:
                    if true {
                        fmt.Println(v14, v2)
                        v17 := GenS1{F0: 2147483647, F1: -32768, F2: ("hello" + "foo bar")}
                        if false {
                            v4 = v4
                            v11 = "\n\t"
                            v3[((v11 + "a") + v11)] = "a"
                        }
                        _ = v17
                    } else {
                        switch ("foo bar" + ((v11 + "foo bar") + (((v11 + ((v11 + ("foo bar" + (v11 + (v11 + v11)))) + (v11 + "\n\t"))) + (v11 + v11)) + v11))) {
                        case "a":
                            if (!false) {
                                var v18 float64 = 3.14159
                                for fi19 := 0; fi19 < 1; fi19++ {
                                    v20 := GenS2{F0: 1, F1: -936, F2: 2}
                                    _ = fi19
                                    _ = v20
                                }
                                var v21 string = v11
                                _ = v18
                                _ = v21
                            } else {
                                v14 = 1
                            }
                        }
                    }
                default:
                    v4 = v2
                }
            }
            v2 -= v2
        }
        _ = len(v3)
        v12 = (-2.25 * -2.25)
    }
    if true {
        var v22 int = -9223372036854775808
        fmt.Println(v12, v14)
        _ = v22
    }
    var v23 int8 = -128
    if (-354 >= -189) {
        var v24 float64 = v12
        fmt.Println(v14, v14, v15)
        _ = v24
    }
    var v25 int8 = 100
    v26 := GenS3{F0: -1, F1: v15, F2: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
