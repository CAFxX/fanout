package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint64 }
func genFn0(p0 int16) bool {
    p0 = 749
    if false {
        fmt.Println(p0)
        var v7 string = "a"
        fmt.Println(v7, p0)
        _ = v7
    }
    _ = p0
    return (0 > -9223372036854775808)
}
func genFn1(p0 string, p1 int64) uint16 {
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := []int{fi8, (fi8 | (fi8 & fi8)), (^(fi8)), fi8}
        var v10 int32 = 1
        fmt.Println(p1)
        _ = fi8
        _ = v10
        _ = v9
    }
    p0 = p0
    if true {
        p0 = ""
    }
    _ = p0
    _ = p1
    return 822
}

func main() {
    var v1 int32 = -718
    v2 := []string{"foo bar", "hello", "foo bar"}
    v3 := make(map[int]uint32)
    var v4 float32 = 0.1
    var v5 int16 = 0
    v6 := []string{(((("a" + "") + ("hello" + ("a" + ("hello" + "hello")))) + "a") + "a")}
    v1 += v1
    for fi11 := 0; fi11 < 8; fi11++ {
        v1 = v1
        for fi12 := 0; fi12 < 16; fi12++ {
            v13 := GenS1{F0: 42}
            v1 += (v1 + v1)
            if (105 <= (v5 + v5)) {
                switch "foo bar" {
                case "":
                    fmt.Println(v1)
                case "a":
                    v1 -= v1
                }
                v4 -= v4
            } else {
                var v14 uint32 = 1
                if ((^(v1)) > (v1 - v1)) {
                    var v15 int64 = 1
                    _ = v15
                }
                var v16 uint8 = 19
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v4 = v4
                fmt.Println(fi11)
                switch fi12 {
                case 32768:
                    v14 = 42
                case 127:
                    v1 -= 2147483647
                    v17 := GenS2{F0: 627, F1: v16, F2: (("foo bar" + "a") + "hello")}
                    _ = v17
                }
                v3[-351] = (^(v14))
                if true {
                    fmt.Println(fi12)
                }
                fmt.Println(v5, v14, v5)
                var v18 int32 = v1
                _ = v14
                _ = v16
                _ = v18
            }
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v19 bool = (!false)
    if true {
        v20 := GenS3{F0: int64(v1)}
        _ = v20
    } else {
        v5 = (v5 * (v5 ^ (v5 >> ((1) & 63))))
    }
    v21 := GenS4{F0: (-1 <= 75), F1: 42, F2: 584}
    var v22 bool = (128 > 1)
    _ = len(v3)
    if (3.14159 > v4) {
        var v23 int8 = -128
        _ = v23
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch 2147483647 {
    case 127:
        if (v5 == v5) {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v5, v4)
            }
            fmt.Println(v1)
            var v24 int = 0
            _ = v24
        }
    case -2147483648:
        v4 += float32(1e10)
    case 251:
        switch ("" + "\n\t") {
        case "":
            var v25 int32 = (v1 - (v1 >> ((1) & 63)))
            _ = v25
        }
    default:
        if (!((!v22) || ((3.14159 / 1.5) == 3.14159))) {
            fmt.Println(v22, v5)
        }
    }
    switch "hello" {
    case "foo bar":
        _ = len(v6)
    }
    _ = len(v3)
    var v26 string = (("hello" + ("foo bar" + "a")) + (((("" + "") + "\n\t") + ("a" + "")) + ("a" + ("a" + "foo bar"))))
    delete(v3, -9223372036854775808)
    v26 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
