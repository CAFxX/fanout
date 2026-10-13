package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int32) uint {
    p0 += 1
    v5 := []int{}
    _ = p0
    _ = v5
    return 18446744073709551615
}
func genFn1() uint64 {
    if false {
        v6 := GenS1{F0: -1}
        v7 := []int{755}
        switch "" {
        case "foo bar":
            for fi8 := 0; fi8 < 7; fi8++ {
                switch ("\n\t" + "hello") {
                case "a":
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                default:
                    var v9 int64 = -748
                    _ = v9
                }
                fmt.Println(fi8)
                v10 := GenS2{F0: 1}
                _ = fi8
                _ = v10
            }
        case "hello":
            var v11 float32 = 1e10
            _ = v11
        default:
            if len(v7) > 0 {
                v7[0] = 128
            }
        }
        _ = v6
        _ = v7
    } else {
        v12 := []int{-777, 0, -2147483648, -9223372036854775808}
        _ = v12
    }
    v13 := []int{1, 32768}
    var v14 uint8 = 127
    _ = v13
    _ = v14
    return 347
}

func main() {
    var v1 int16 = 127
    v2 := []int{128, -9223372036854775808, 1, 32768}
    var v3 int64 = 42
    v4 := make(map[int]string)
    var v15 int32 = -1
    if ("\n\t" >= ((((("a" + "foo bar") + ("a" + ("" + "foo bar"))) + ("foo bar" + "foo bar")) + "\n\t") + ("\n\t" + ("a" + ("foo bar" + "a"))))) {
        var v16 string = "hello"
        _ = v16
    }
    var v17 int32 = (v15 * -323)
    var v18 int16 = (v1 * 0)
    if false {
        v19 := GenS2{F0: (v1 >> ((4294967295) & 63))}
        _ = v19
    }
    var v20 bool = (0 >= 18446744073709551615)
    if v20 {
        if len(v2) > 0 {
            v2[0] = 42
        }
    }
    switch -2147483648 {
    case 42:
        v1 = 789
    case 9223372036854775807:
        switch (("" + ("foo bar" + (("foo bar" + "\n\t") + ("hello" + "foo bar")))) + "hello") {
        case "a":
            v15 = v17
        case "":
            v4[2147483647] = "foo bar"
        }
    case 2147483647:
        switch 42 {
        case 0:
            v21 := GenS3{F0: 0, F1: 0}
            _ = v21
        case 42:
            var v22 int16 = v18
            _ = v22
        case 4294967295:
            if v20 {
                v18 = (v18 / (((v18 >> ((685) & 63))) | 1))
                var v23 uint32 = 60
                _ = v23
            } else {
                v18 += (v1 | v1)
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    default:
        fmt.Println(v17, v20, v18)
    }
    v3 = v3
    fmt.Println(v20, v20, v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v18, v20)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
