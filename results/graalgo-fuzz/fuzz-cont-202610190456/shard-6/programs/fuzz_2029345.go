package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int16) uint64 {
    var v4 uint = 18446744073709551615
    for fi5 := 0; fi5 < 7; fi5++ {
        v4 -= 4294967295
        _ = fi5
    }
    _ = p0
    _ = v4
    return 18446744073709551615
}
func genFn1(p0 uint8) int8 {
    p0 = (p0 >> ((4294967295) & 63))
    var v6 bool = (!false)
    v6 = (1 >= 9223372036854775807)
    switch 477 {
    case 9223372036854775807:
        fmt.Println(p0, v6)
    case 573:
        if v6 {
            v7 := []int{-9223372036854775808, 148}
            _ = v7
        } else {
            var v8 int64 = -9223372036854775808
            for fi9 := 0; fi9 < 3; fi9++ {
                switch fi9 {
                case -645:
                    p0 = p0
                case 1:
                    for fi10 := 0; fi10 < 5; fi10++ {
                        var v11 uint64 = 774
                        switch "foo bar" {
                        case "hello":
                            var v12 uint8 = (p0 - p0)
                            _ = v12
                        }
                        _ = fi10
                        _ = v11
                    }
                case -483:
                    if false {
                        p0 = 2
                    } else {
                        var v13 string = (((("a" + "\n\t") + "foo bar") + "hello") + "\n\t")
                        var v14 float32 = (-2.25 / (-2.25 + 3.14159))
                        _ = v13
                        _ = v14
                    }
                default:
                    v15 := []int{}
                    _ = v15
                }
                for fi16 := 0; fi16 < 20; fi16++ {
                    var v17 uint = 255
                    _ = fi16
                    _ = v17
                }
                _ = fi9
            }
            _ = v8
        }
    default:
        v18 := GenS1{F0: 42, F1: -1}
        _ = v18
    }
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 int32 = 356
    v2 := make(map[string]uint32)
    var v3 string = ("" + "hello")
    v3 = ((v3 + v3) + ("\n\t" + v3))
    var v19 int64 = -1
    var v20 uint = 255
    var v21 int8 = -44
    v3 = ("\n\t" + v3)
    var v22 uint16 = 856
    fmt.Println(v20)
    var v23 int64 = v19
    for fi24 := 0; fi24 < 12; fi24++ {
        fmt.Println(v20)
        if (true && (42 == v1)) {
            for fi25 := 0; fi25 < 9; fi25++ {
                fmt.Println(v22, v1)
                v20 = v20
                var v26 int = (^(fi25))
                _ = fi25
                _ = v26
            }
        }
        fmt.Println(fi24)
        _ = fi24
    }
    v22 += (v22 % (((v22 & (v22 % ((525) | 1)))) | 1))
    fmt.Println(v19)
    v19 += (v23 + int64(v21))
    v27 := GenS2{F0: -1}
    switch "" {
    case "":
        if (!(234 != 2147483647)) {
            v3 = v3
        } else {
            fmt.Println(v21)
            if true {
                v1 = int32(v23)
                for fi28 := 0; fi28 < 2; fi28++ {
                    var v29 string = v3
                    _ = fi28
                    _ = v29
                }
                v3 = (v3 + ("" + v3))
            }
            v30 := []int{0, -391}
            _ = v30
        }
    case "foo bar":
        for fi31 := 0; fi31 < 16; fi31++ {
            v32 := GenS3{F0: v1}
            var v33 bool = (true || false)
            var v34 int64 = 9223372036854775807
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
}
