package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int, p1 string) uint {
    v7 := []int{p0, (p0 ^ p0), p0}
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return 255
}
func genFn1() uint32 {
    v8 := GenS1{F0: 518, F1: 0}
    var v9 string = "hello"
    _ = v8
    _ = v9
    return 116
}

func main() {
    var v1 int = 2147483647
    var v2 int8 = 0
    v3 := make(map[uint32]bool)
    var v4 float32 = ((0.0 * -2.25) + 1e10)
    v5 := make(map[uint32]uint32)
    var v6 int8 = v2
    v1 += (v1 + v1)
    v2 = -39
    if (false && (v1 > 0)) {
        switch v1 {
        case 233:
            v2 = (^(v2))
        case 168:
            var v10 uint16 = 0
            _ = v10
        case 32768:
            fmt.Println(v1)
        default:
            v11 := GenS2{F0: 128, F1: 42}
            _ = v11
        }
        v1 = 306
    } else {
        v12 := []int{1, v1}
        if false {
            switch -823 {
            case 0:
                fmt.Println(v2, v2, v4)
            case 1:
                fmt.Println(v6, v2)
            default:
                switch -1 {
                case 2147483647:
                    var v13 uint16 = 21
                    _ = v13
                }
            }
            fmt.Println(v6, v4)
        } else {
            switch ((("" + "hello") + "a") + "a") {
            case "":
                switch ("foo bar" + "foo bar") {
                case "":
                    var v14 string = "foo bar"
                    _ = v14
                case "hello":
                    v6 = -1
                default:
                    var v15 int16 = 42
                    _ = v15
                }
            case "\n\t":
                var v16 float64 = ((3.14159 / 1e10) - (1.5 * 1e10))
                _ = v16
            case "foo bar":
                switch ("a" + "hello") {
                case "":
                    v17 := v3[2147483647]
                    _ = v17
                case "\n\t":
                    v1 += (v1 ^ (v1 * -9223372036854775808))
                case "hello":
                    var v18 float64 = 1e10
                    _ = v18
                default:
                    fmt.Println(v4, v4, v1)
                }
            default:
                v1 -= v1
            }
        }
        if false {
            var v19 int32 = -819
            fmt.Println(v1, v1)
            _ = v19
        }
        _ = v12
    }
    if true {
        var v20 int64 = 42
        var v21 uint = genFn0(-351, (("a" + "a") + "\n\t"))
        _ = v20
        _ = v21
    } else {
        v22 := GenS1{F0: 635, F1: 42}
        switch (v1 & 42) {
        case -2147483648:
            v23 := GenS1{F0: 4294967295, F1: -1}
            _ = v23
        }
        _ = v22
    }
    v1 += v1
    fmt.Println(v6)
    fmt.Println(v4, v4, v1)
    v24 := GenS3{F0: -2147483648}
    for fi25 := 0; fi25 < 15; fi25++ {
        fmt.Println(v2, v6)
        v26 := []int{(fi25 << ((4294967295) & 63)), 1, v1}
        fmt.Println(v6)
        _ = fi25
        _ = v26
    }
    var v27 int = (v1 - (-(v1)))
    var v28 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
