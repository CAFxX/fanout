package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint64 }
func genFn0(p0 uint64, p1 int32) string {
    if false {
        p1 = p1
    }
    p1 = p1
    var v4 string = ("" + (("a" + "hello") + ""))
    v5 := GenS1{F0: 42, F1: v4}
    p1 += (p1 * p1)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return "\n\t"
}
func genFn1() int64 {
    var v6 float64 = 3.14159
    v7 := []int{0}
    fmt.Println(v6)
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int32 = 42
    var v2 bool = (false && (42 == -1))
    var v3 float32 = (1.5 + (float32(1e10) + -2.25))
    fmt.Println(v3, v2, v1)
    if (v2 && (v2 && ("a" > "foo bar"))) {
        v8 := GenS2{F0: -128}
        v1 += (v1 | (^(v1)))
        var v9 int64 = 164
        _ = v8
        _ = v9
    }
    v10 := GenS3{F0: ("a" + "a"), F1: 2}
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(v2, v2, fi11)
            v2 = false
            v3 -= v3
            _ = fi12
        }
        _ = fi11
    }
    var v13 uint = 2
    var v14 int16 = -735
    fmt.Println(v3)
    v13 = 255
    var v15 int16 = v14
    var v16 uint8 = 37
    switch 32767 {
    case 9223372036854775807:
        v1 -= (^(v1))
    default:
        v17 := []int{32}
        _ = v17
    }
    if (v1 < (v1 ^ (v1 | 477))) {
        fmt.Println(v2, v13, v3)
        if (!true) {
            fmt.Println(v1)
        } else {
            v18 := GenS1{F0: v13, F1: ("hello" + "foo bar")}
            switch 486 {
            case 307:
                var v19 int = 32768
                _ = v19
            case 0:
                v20 := GenS1{F0: 42, F1: "hello"}
                _ = v20
            case 32768:
                v21 := GenS1{F0: v13, F1: "foo bar"}
                _ = v21
            }
            _ = v18
        }
        v14 = v15
    } else {
        var v22 uint = 65535
        var v23 int32 = v1
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 float32 = 0.1
            _ = fi25
            _ = v26
        }
        v3 -= 0.0
        var v27 int8 = int8(v1)
        _ = fi24
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
