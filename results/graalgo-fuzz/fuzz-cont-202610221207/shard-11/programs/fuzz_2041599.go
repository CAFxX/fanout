package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 int32, p1 string) int {
    v5 := []int{-280, 0}
    if (false || (3.14159 != (0.1 - (1e10 - (3.14159 - float32((1e10 / (1.5 + (1e10 + (((1.5 + 0.0) / (1e10 * float64(0.1))) + 1.5)))))))))) {
        p1 = "foo bar"
        fmt.Println(p1)
        fmt.Println(p1, p0)
    } else {
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := GenS1{F0: 255, F1: 919, F2: 0}
            _ = fi6
            _ = v7
        }
        if ((-1 <= -1) || false) {
            var v8 uint8 = 0
            _ = v8
        }
        fmt.Println(p1, p1)
    }
    if len(v5) > 0 {
        v5[0] = 32767
    }
    if ((3.14159 / 1.0) < ((0.0 / 1.5) - (1e10 / 1.0))) {
        switch p1 {
        case "hello":
            v9 := GenS2{F0: 1}
            _ = v9
        }
    } else {
        var v10 float32 = 3.14159
        v11 := GenS1{F0: 18446744073709551615, F1: 0, F2: 217}
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(p1, p1)
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v5
    return 128
}

func main() {
    var v1 uint32 = 804
    var v2 int8 = 100
    v3 := make(map[uint32]string)
    var v4 float64 = 3.14159
    _ = len(v3)
    v13 := []int{32767}
    v14 := GenS2{F0: 322}
    v15 := GenS1{F0: 396, F1: 178, F2: v1}
    if ((1 > 392) || (("foo bar" + ("a" + "")) < (("" + "a") + ("\n\t" + ("\n\t" + "\n\t"))))) {
        for fi16 := 0; fi16 < 11; fi16++ {
            v4 += 3.14159
            v4 = (v4 * 0.0)
            var v17 bool = ((42 >= (v2 - int8(v1))) || true)
            _ = fi16
            _ = v17
        }
        v18 := GenS3{F0: 235}
        var v19 float64 = v4
        _ = v18
        _ = v19
    }
    switch ((("foo bar" + ("foo bar" + "hello")) + "foo bar") + ("hello" + ("foo bar" + ("hello" + "hello")))) {
    case "":
        switch "hello" {
        case "hello":
            v20 := GenS2{F0: 245}
            _ = v20
        case "a":
            var v21 uint8 = 108
            _ = v21
        case "\n\t":
            var v22 uint = 4294967295
            _ = v22
        default:
            v23 := v3[(^(v1))]
            _ = v23
        }
    case "\n\t":
        fmt.Println(v1, v2)
    default:
        v24 := GenS4{F0: "\n\t", F1: 65535}
        _ = v24
    }
    v2 = v2
    if true {
        var v25 uint64 = 4294967295
        var v26 uint32 = (v1 ^ 0)
        _ = v25
        _ = v26
    } else {
        v2 -= (v2 % (((v2 & (v2 | v2))) | 1))
        v2 = 68
        var v27 int8 = (v2 / ((100) | 1))
        _ = v27
    }
    switch "foo bar" {
    case "":
        for fi28 := 0; fi28 < 16; fi28++ {
            v29 := GenS2{F0: 703}
            delete(v3, (v1 << ((42) & 63)))
            _ = fi28
            _ = v29
        }
    default:
        if (false && true) {
            if (true && false) {
                var v30 uint64 = 2
                v2 = (^(v2))
                fmt.Println(v30, v30)
                _ = v30
            } else {
                v1 = v1
            }
            fmt.Println(v4, v1, v2)
        }
    }
    switch ("foo bar" + "") {
    case "foo bar":
        v3[4294967295] = (("hello" + "a") + ("foo bar" + "a"))
    case "":
        delete(v3, v1)
    case "hello":
        _ = len(v13)
    default:
        v31 := GenS3{F0: 255}
        _ = v31
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        v33 := GenS4{F0: "a", F1: 2}
        _ = fi32
        _ = v33
    }
    var v34 bool = false
    fmt.Println(v34, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
