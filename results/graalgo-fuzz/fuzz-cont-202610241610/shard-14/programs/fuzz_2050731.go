package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint16, p2 int16) bool {
    if true {
        p1 -= (p1 * uint16(p0))
    } else {
        fmt.Println(p1, p2)
    }
    p1 = (p1 | p1)
    fmt.Println(p1, p2)
    var v7 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}
func genFn1(p0 int64, p1 int, p2 uint16) uint8 {
    switch "" {
    case "a":
        fmt.Println(p1)
    case "":
        p2 -= (p2 | 87)
    }
    var v8 uint = 2
    var v9 int8 = 9
    p2 += (p2 & 1)
    v9 = (v9 - int8(v8))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 237
}

func main() {
    var v1 float32 = 1e10
    var v2 string = ("" + "foo bar")
    v3 := make(map[string]bool)
    var v4 int64 = -1
    var v5 uint64 = 0
    var v6 bool = (!((true && true) || (!(!(127 < 100)))))
    var v10 int32 = 2147483647
    switch "hello" {
    case "":
        if ((!true) && false) {
            v11 := []int{0, -122, -212}
            fmt.Println(v2)
            _ = v11
        } else {
            v12 := GenS1{F0: 2, F1: 127, F2: 255}
            for fi13 := 0; fi13 < 15; fi13++ {
                for fi14 := 0; fi14 < 6; fi14++ {
                    v5 -= 62
                    _ = fi14
                }
                if true {
                    var v15 int32 = (v10 * v10)
                    var v16 int = (fi13 ^ (fi13 | -326))
                    for fi17 := 0; fi17 < 11; fi17++ {
                        v18 := v3["a"]
                        v10 -= (v10 | v15)
                        fmt.Println(v5)
                        _ = fi17
                        _ = v18
                    }
                    _ = v15
                    _ = v16
                } else {
                    v2 = (v2 + "")
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v4 += (v4 % (((v4 % (((v4 << ((255) & 63))) | 1))) | 1))
                        var v20 float64 = (3.14159 * ((0.1 - 1.5) - -2.25))
                        var v21 uint32 = 2147483647
                        v22 := v3[""]
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    var v23 uint = 1
                    _ = v23
                }
                _ = fi13
            }
            _ = v12
        }
    case "hello":
        switch -839 {
        case 4294967295:
            switch v2 {
            case "\n\t":
                if true {
                    var v24 int16 = 1
                    v25 := v3[v2]
                    v1 = float32((1e10 * 0.0))
                    _ = v24
                    _ = v25
                } else {
                    v26 := GenS1{F0: genFn1(821, -748, 32767), F1: 127, F2: 2}
                    var v27 uint16 = 42
                    switch "hello" {
                    case "hello":
                        v5 += v5
                    }
                    _ = v26
                    _ = v27
                }
            case "":
                fmt.Println(v6)
            default:
                switch 1 {
                case -2147483648:
                    var v28 int64 = 69
                    _ = v28
                case 32767:
                    var v29 uint8 = 42
                    _ = v29
                case -9223372036854775808:
                    for fi30 := 0; fi30 < 1; fi30++ {
                        v4 += v4
                        v31 := []int{fi30, 32767, -2147483648}
                        v5 -= 18446744073709551615
                        _ = fi30
                        _ = v31
                    }
                }
            }
        case 0:
            var v32 uint64 = (v5 * v5)
            _ = v32
        default:
            var v33 int64 = 633
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        fmt.Println(fi34, v2, v10)
        if (!(42 < 0)) {
            var v35 int64 = -891
            fmt.Println(v10, v5, v10)
            _ = v35
        }
        for fi36 := 0; fi36 < 1; fi36++ {
            _ = len(v3)
            _ = len(v3)
            _ = fi36
        }
        _ = fi34
    }
    v37 := []int{32767, -1, -2147483648, 32767}
    v37 = append(v37, -1)
    v38 := GenS2{F0: 209}
    v39 := GenS1{F0: 0, F1: genFn1((v4 / (((^(v4))) | 1)), -662, 42), F2: 127}
    v40 := GenS3{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
