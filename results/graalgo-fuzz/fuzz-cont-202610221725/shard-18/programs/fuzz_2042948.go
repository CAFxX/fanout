package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0() string {
    var v4 uint8 = 42
    var v5 int16 = 1
    if (false || true) {
        v6 := GenS1{F0: "foo bar", F1: 9223372036854775807, F2: 32767}
        v7 := []int{980, 9223372036854775807, -1}
        _ = v6
        _ = v7
    } else {
        v4 = 41
        if (65535 == 255) {
            switch "\n\t" {
            case "foo bar":
                v5 = (v5 * -877)
            }
        } else {
            v8 := GenS1{F0: "hello", F1: 42, F2: (^(v5))}
            _ = v8
        }
    }
    _ = v4
    _ = v5
    return ""
}
func genFn1() int32 {
    if true {
        var v9 bool = (788 != 828)
        var v10 string = "hello"
        for fi11 := 0; fi11 < 10; fi11++ {
            fmt.Println(fi11)
            v9 = false
            _ = fi11
        }
        _ = v10
        _ = v9
    }
    var v12 float64 = 0.0
    var v13 string = ("" + "foo bar")
    v14 := []int{-43}
    v15 := GenS2{F0: 2147483648, F1: 0}
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return -766
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = ("a" <= (("" + ("foo bar" + ("" + ("" + "hello")))) + ""))
    var v3 int16 = 1
    if v2 {
        v2 = v2
        switch 32768 {
        case -9223372036854775808:
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 int64 = 0
                var v18 int32 = -1
                fmt.Println(v18)
                _ = fi16
                _ = v17
                _ = v18
            }
        }
    }
    v2 = (v2 || true)
    var v19 uint = 535
    v20 := GenS2{F0: 1, F1: 1}
    fmt.Println(v19, v2, v2)
    fmt.Println(v19, v3, v2)
    v21 := GenS1{F0: "\n\t", F1: 9223372036854775807, F2: (v3 + v3)}
    v19 = v19
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := GenS2{F0: 2147483648, F1: 4294967295}
        v2 = (696 == 2)
        _ = fi22
        _ = v23
    }
    v2 = (false || v2)
    fmt.Println(v3)
    v1[615] = -1
    var v24 uint = (^(v19))
    fmt.Println(v2, v19, v24)
    var v25 uint64 = 614
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
}
