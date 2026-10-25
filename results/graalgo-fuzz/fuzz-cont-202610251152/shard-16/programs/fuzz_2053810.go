package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int8; F1 uint64; F2 int }
type GenS3 struct { F0 uint8; F1 int8; F2 uint64 }
func genFn0() bool {
    var v5 bool = ((1.5 == 3.14159) && true)
    var v6 uint = 517
    if v5 {
        fmt.Println(v6, v6)
        v7 := []int{int(v6), 0}
        var v8 string = "hello"
        _ = v7
        _ = v8
    }
    _ = v5
    _ = v6
    return (v5 && ((v6 | v6) == (v6 + v6)))
}
func genFn1() uint16 {
    switch "foo bar" {
    case "":
        if true {
            var v9 int16 = 42
            _ = v9
        } else {
            var v10 uint16 = 65535
            v11 := []int{1, 2147483647}
            v10 -= v10
            _ = v10
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 float64 = 0.1
            v14 := GenS1{F0: -125, F1: (-(fi12))}
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    var v15 float32 = float32(-2.25)
    _ = v15
    return 1
}

func main() {
    var v1 int8 = -120
    v2 := []int{362}
    var v3 bool = false
    var v4 uint8 = 2
    if ((^(v1)) > (^(v1))) {
        if (2147483647 >= -396) {
            v3 = (v3 && (0 != 2))
            switch (("foo bar" + "a") + ("foo bar" + ("a" + ("\n\t" + "\n\t")))) {
            case "\n\t":
                var v16 int64 = -9223372036854775808
                _ = v16
            }
            var v17 uint16 = 0
            _ = v17
        }
        if false {
            fmt.Println(v4, v1, v1)
        } else {
            v18 := GenS1{F0: 366, F1: -9223372036854775808}
            v19 := GenS2{F0: (v1 ^ (v1 >> ((232) & 63))), F1: 4294967295, F2: -436}
            _ = v18
            _ = v19
        }
    } else {
        v3 = (-236 < 128)
    }
    fmt.Println(v1, v3, v4)
    switch "a" {
    case "foo bar":
        var v20 uint32 = 638
        _ = v20
    case "a":
        v1 -= v1
    }
    v3 = (!false)
    var v21 float32 = float32(float64(-2.25))
    fmt.Println(v1, v4, v3)
    v22 := GenS3{F0: v4, F1: v1, F2: 4294967295}
    if v3 {
        if v3 {
            v23 := GenS3{F0: (^(v4)), F1: v1, F2: 0}
            _ = v23
        } else {
            fmt.Println(v4, v4)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
