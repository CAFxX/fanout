package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint8 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int16; F1 bool; F2 bool }
func genFn0(p0 bool, p1 int32, p2 uint32) string {
    v6 := GenS1{F0: 18446744073709551615, F1: 48, F2: 42}
    var v7 uint = 0
    var v8 int64 = 42
    var v9 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return "foo bar"
}
func genFn1(p0 int64, p1 uint) int64 {
    p0 = p0
    switch ("a" + (("hello" + "") + (("\n\t" + "hello") + ("a" + "a")))) {
    case "hello":
        p0 = (p0 + (p0 % (((p0 << ((p1) & 63))) | 1)))
    case "foo bar":
        p1 = 266
    case "\n\t":
        var v10 int8 = 127
        _ = v10
    default:
        switch (("foo bar" + ("" + "")) + ("\n\t" + "\n\t")) {
        case "\n\t":
            var v11 uint8 = 15
            _ = v11
        case "":
            if (!true) {
                for fi12 := 0; fi12 < 4; fi12++ {
                    v13 := []int{}
                    v13 = append(v13, (fi12 - fi12))
                    _ = fi12
                    _ = v13
                }
                p1 -= (p1 & p1)
            } else {
                p0 = (p0 * -9223372036854775808)
            }
        case "hello":
            if false {
                p0 += (^(p0))
                var v14 int = 9223372036854775807
                _ = v14
            }
        default:
            p1 = (p1 + (p1 << ((p1) & 63)))
        }
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        fmt.Println(fi15, p0)
        var v16 float32 = 1.5
        var v17 int64 = (-(p0))
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 string = "foo bar"
    var v2 int8 = -128
    v3 := []int{}
    var v4 uint32 = 65535
    v5 := make(map[string]bool)
    var v18 float64 = float64((0.1 * 1e10))
    fmt.Println(v2, v2)
    var v19 int16 = -620
    switch ("a" + "a") {
    case "\n\t":
        for fi20 := 0; fi20 < 2; fi20++ {
            fmt.Println(v19)
            _ = fi20
        }
    case "hello":
        delete(v5, (("" + "foo bar") + ""))
    case "a":
        for fi21 := 0; fi21 < 11; fi21++ {
            v22 := GenS2{F0: v1, F1: -2147483648}
            _ = fi21
            _ = v22
        }
    default:
        if false {
            v2 += (v2 | v2)
            v23 := GenS3{F0: 0, F1: (^(v4))}
            var v24 string = "hello"
            _ = v23
            _ = v24
        }
    }
    var v25 int16 = 42
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 12; fi27++ {
            fmt.Println(v1, fi27, v2)
            v5[v1] = true
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi27
        }
        var v28 uint32 = (^(v4))
        _ = fi26
        _ = v28
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        fmt.Println(v4, v25)
        v2 -= 100
        _ = fi29
    }
    v30 := v5[v1]
    var v31 uint = 42
    v18 = (v18 / (v18 - v18))
    if (v18 >= v18) {
        var v32 string = v1
        _ = v32
    }
    v33 := GenS4{F0: (v19 / (((v19 % (((v25 & v25)) | 1))) | 1)), F1: true, F2: (v30 || (!false))}
    var v34 uint64 = 489
    fmt.Println(v31, v2)
    v4 -= v4
    var v35 float32 = (1e10 / ((0.0 / 1.0) * (1.5 - ((float32(v18) * ((((1e10 + (1e10 * (0.1 - (float32(v18) * 0.1)))) + 0.1) + 3.14159) - (1e10 / float32(v18)))) - (float32(((float64(1.5) * -2.25) / (-2.25 * (1e10 / 3.14159)))) + (-2.25 / float32(-2.25)))))))
    v19 = v25
    var v36 uint64 = v34
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
