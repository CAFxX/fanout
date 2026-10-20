package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int16 }
func genFn0() string {
    var v7 int64 = 9223372036854775807
    fmt.Println(v7)
    v8 := []int{4294967295}
    if true {
        fmt.Println(v7)
        if false {
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(v7, v7)
                if len(v8) > 0 {
                    v8[0] = (fi9 ^ 2147483647)
                }
                var v10 float32 = 1.5
                _ = fi9
                _ = v10
            }
            if len(v8) > 0 {
                v8[0] = 127
            }
        }
    } else {
        v7 = v7
    }
    _ = v7
    _ = v8
    return "a"
}
func genFn1(p0 bool) int8 {
    fmt.Println(p0)
    var v11 uint64 = 367
    _ = p0
    _ = v11
    return 127
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 680
    var v3 int16 = 1000
    v4 := make(map[uint32]string)
    var v5 float64 = 3.14159
    var v6 string = "foo bar"
    var v12 float64 = v5
    v13 := GenS1{F0: 0, F1: -32768, F2: (v2 - 42)}
    for fi14 := 0; fi14 < 1; fi14++ {
        if true {
            v12 = (v5 * (v12 + float64(float32(((v12 + float64(((1.5 * 0.0) * (1.5 - 1e10)))) / float64(-2.25))))))
            fmt.Println(v5)
            var v15 int16 = 671
            _ = v15
        }
        _ = fi14
    }
    v16 := []int{-592, 9223372036854775807, 32767, 0}
    switch -9223372036854775808 {
    case 0:
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 bool = (!true)
            v16 = append(v16, (fi17 << ((2) & 63)))
            v19 := GenS2{F0: ("hello" + ("foo bar" + v6)), F1: 0, F2: "\n\t"}
            _ = fi17
            _ = v18
            _ = v19
        }
    case -456:
        if false {
            if ("a" < "hello") {
                fmt.Println(v12, v2)
            } else {
                v3 += (v3 * v3)
                var v20 int8 = 0
                v3 = v3
                _ = v20
            }
        } else {
            v6 = "hello"
        }
    case -9223372036854775808:
        v5 -= (v12 + (v5 / 1.0))
    }
    v5 -= v12
    switch v6 {
    case "\n\t":
        if true {
            if len(v16) > 0 {
                v16[0] = int(v2)
            }
        }
    case "":
        if (!(!((("hello" + v6) < v6) || (((418 > 4294967295) && (!(!(false && true)))) && true)))) {
            switch v6 {
            case "\n\t":
                switch "foo bar" {
                case "a":
                    v3 = v3
                case "hello":
                    fmt.Println(v3, v12)
                case "":
                    var v21 float32 = -2.25
                    _ = v21
                default:
                    for fi22 := 0; fi22 < 8; fi22++ {
                        var v23 uint64 = 0
                        v5 = 1e10
                        delete(v4, 602)
                        _ = fi22
                        _ = v23
                    }
                }
            }
        }
    default:
        v24 := GenS3{F0: 2147483647}
        _ = v24
    }
    var v25 float32 = (float32((v5 * -2.25)) / 1.0)
    v12 -= (1e10 - (v5 * v5))
    switch v6 {
    case "":
        v26 := GenS4{F0: v3, F1: 0}
        _ = v26
    default:
        v25 = 0.0
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        v28 := GenS5{F0: 276}
        var v29 uint = 65535
        v6 = v6
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int16 = (v3 << ((2) & 63))
    var v31 float64 = v12
    if false {
        v5 -= float64((v25 / 0.1))
        fmt.Println(v6, v3)
    }
    switch 127 {
    case 371:
        v6 = ("foo bar" + v6)
    default:
        delete(v1, (v6 + v6))
    }
    var v32 int = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
