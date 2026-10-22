package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 bool; F1 string }
func genFn0(p0 int16) string {
    var v7 bool = (true || (!(false && (!(true && false)))))
    v8 := []int{2147483647}
    if ((0 == p0) || ("hello" != "\n\t")) {
        if v7 {
            _ = len(v8)
            switch 128 {
            case 2147483647:
                v7 = (-453 == -655)
            case 32767:
                switch 2147483647 {
                case 32768:
                    switch 32768 {
                    case 22:
                        if (!("\n\t" < "a")) {
                            for fi9 := 0; fi9 < 8; fi9++ {
                                switch "\n\t" {
                                case "a":
                                    switch "hello" {
                                    case "":
                                        var v10 uint32 = 0
                                        _ = v10
                                    case "foo bar":
                                        p0 -= (p0 + 32767)
                                    }
                                case "hello":
                                    _ = len(v8)
                                default:
                                    var v11 float64 = -2.25
                                    _ = v11
                                }
                                v7 = v7
                                _ = fi9
                            }
                            fmt.Println(v7)
                        }
                    case 1:
                        v12 := GenS1{F0: 0}
                        _ = v12
                    case 0:
                        v8 = append(v8, 9223372036854775807)
                    }
                }
            }
        } else {
            var v13 int32 = -961
            _ = v13
        }
    }
    fmt.Println(p0, v7)
    _ = p0
    _ = v7
    _ = v8
    return ""
}

func main() {
    v1 := []int{9223372036854775807}
    v2 := []uint32{}
    var v3 bool = ("\n\t" > "")
    var v4 int64 = 9223372036854775807
    var v5 float32 = (float32((3.14159 * 0.0)) * (0.1 * (-2.25 * 3.14159)))
    v6 := make(map[int]int)
    v5 = 0.0
    v5 = v5
    v5 += 0.0
    v14 := GenS1{F0: -274}
    v4 = 0
    if true {
        v15 := GenS1{F0: 0}
        v16 := GenS2{F0: 0, F1: 1}
        _ = v15
        _ = v16
    }
    var v17 float32 = 0.0
    v4 += v4
    switch "hello" {
    case "a":
        switch "a" {
        case "\n\t":
            if v3 {
                v18 := GenS2{F0: 829, F1: 0}
                v3 = false
                _ = v18
            }
        case "":
            fmt.Println(v3, v4)
        case "foo bar":
            fmt.Println(v17)
        default:
            v17 = -2.25
        }
    case "foo bar":
        v6[4294967295] = 992
    case "hello":
        v19 := GenS3{F0: ((42 > 709) || v3), F1: ("foo bar" + "hello")}
        _ = v19
    }
    if v3 {
        switch "hello" {
        case "hello":
            var v20 uint32 = 945
            _ = v20
        case "":
            v17 = v5
        default:
            v3 = (718 >= -1)
        }
    } else {
        v5 = v17
        v21 := GenS2{F0: 65535, F1: 42}
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
