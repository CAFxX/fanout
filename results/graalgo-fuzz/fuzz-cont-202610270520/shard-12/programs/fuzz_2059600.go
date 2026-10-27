package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint8, p2 uint8) int16 {
    switch "foo bar" {
    case "":
        if (true || ((!((!(4294967295 > 2147483647)) && true)) && true)) {
            v4 := []int{4294967295, -613, -395}
            fmt.Println(p2)
            if len(v4) > 0 {
                v4[0] = -506
            }
            _ = v4
        } else {
            p2 += 232
            p1 = p1
        }
    case "foo bar":
        p2 += 0
    }
    v5 := []int{-503, 127, 32767, -957}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -750
}
func genFn1() uint32 {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 string = "a"
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 69
    _ = v8
    return 150
}

func main() {
    var v1 uint16 = 193
    var v2 string = ("foo bar" + (("hello" + "a") + ((("hello" + "hello") + "foo bar") + "")))
    var v3 float64 = -2.25
    var v9 int8 = 50
    v3 -= v3
    v9 -= v9
    var v10 uint32 = 65535
    var v11 bool = true
    for fi12 := 0; fi12 < 18; fi12++ {
        switch "a" {
        case "a":
            v10 += v10
        case "\n\t":
            v13 := GenS1{F0: ((("\n\t" + "a") + ("\n\t" + "a")) + v2)}
            _ = v13
        default:
            switch 719 {
            case 32768:
                v14 := GenS1{F0: ((v2 + (v2 + v2)) + (v2 + (v2 + v2)))}
                _ = v14
            case 876:
                v15 := GenS1{F0: v2}
                _ = v15
            case 127:
                var v16 uint32 = (v10 / ((v10) | 1))
                _ = v16
            default:
                switch v2 {
                case "":
                    var v17 uint32 = 2147483647
                    _ = v17
                default:
                    v10 = v10
                }
            }
        }
        switch ("hello" + (v2 + (v2 + v2))) {
        case "foo bar":
            var v18 int16 = 0
            _ = v18
        case "a":
            var v19 float32 = (-2.25 * (1.5 - 3.14159))
            _ = v19
        case "hello":
            fmt.Println(v11, v2)
        default:
            if v11 {
                v20 := GenS2{F0: v2}
                v21 := []int{-890, (^(fi12)), 9223372036854775807}
                _ = v20
                _ = v21
            } else {
                v22 := GenS1{F0: (v2 + v2)}
                if v11 {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        v10 += (v10 << ((19) & 63))
                        var v24 float64 = v3
                        _ = fi23
                        _ = v24
                    }
                    var v25 float64 = (v3 / 1.0)
                    v10 -= (v10 & (v10 % ((genFn1()) | 1)))
                    _ = v25
                } else {
                    var v26 int32 = 585
                    var v27 uint8 = 2
                    v28 := GenS3{F0: v1}
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = v22
            }
        }
        _ = fi12
    }
    var v29 uint8 = 245
    v30 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v9
}
