package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 int32) uint {
    v4 := GenS1{F0: -9223372036854775808}
    var v5 float64 = 1e10
    for fi6 := 0; fi6 < 19; fi6++ {
        switch "foo bar" {
        case "":
            v5 = v5
        case "a":
            var v7 uint16 = 106
            _ = v7
        }
        _ = fi6
    }
    _ = p0
    _ = v4
    _ = v5
    return 2
}

func main() {
    v1 := []string{"foo bar", ("hello" + ((("hello" + (("foo bar" + "\n\t") + "a")) + (("a" + "foo bar") + "a")) + "")), "a", (("hello" + "hello") + "hello")}
    v2 := make(map[int]int)
    var v3 uint8 = 187
    v8 := v2[511]
    if len(v1) > 0 {
        v1[0] = (("" + ("foo bar" + ("foo bar" + ""))) + "hello")
    }
    fmt.Println(v3, v8)
    if (!(("\n\t" == "a") || (!true))) {
        fmt.Println(v8, v3)
        var v9 uint8 = v3
        _ = v9
    } else {
        if true {
            v8 = v8
            v2[v8] = (v8 / ((v8) | 1))
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS1{F0: 0}
                _ = fi10
                _ = v11
            }
        } else {
            var v12 uint64 = 42
            _ = v12
        }
    }
    fmt.Println(v3)
    if true {
        v3 = (^(v3))
    } else {
        v13 := GenS1{F0: -1}
        v2[v8] = (v8 % ((v8) | 1))
        _ = v13
    }
    v14 := GenS1{F0: -212}
    fmt.Println(v8)
    fmt.Println(v3, v8)
    if (1 < -1) {
        if (0 != 825) {
            fmt.Println(v8, v3)
            v15 := GenS1{F0: -122}
            _ = v15
        }
        var v16 int16 = -400
        v17 := GenS1{F0: -9223372036854775808}
        _ = v16
        _ = v17
    } else {
        switch (-(v8)) {
        case 4294967295:
            var v18 uint8 = (v3 * (v3 ^ (v3 % (((v3 / ((128) | 1))) | 1))))
            _ = v18
        case 210:
            fmt.Println(v8)
        }
    }
    var v19 int32 = 758
    v20 := GenS2{F0: ("\n\t" + "hello"), F1: 0}
    fmt.Println(v19)
    for fi21 := 0; fi21 < 9; fi21++ {
        for fi22 := 0; fi22 < 11; fi22++ {
            fmt.Println(fi21, v3)
            switch (v8 + (fi22 ^ -572)) {
            case 2147483647:
                v3 -= 77
                fmt.Println(fi21)
                v3 -= v3
            case 9223372036854775807:
                v19 = v19
            default:
                v23 := GenS3{F0: v8, F1: (v3 - v3)}
                _ = v23
            }
            v24 := GenS2{F0: "hello", F1: -1}
            _ = fi22
            _ = v24
        }
        _ = fi21
    }
    var v25 uint8 = (v3 ^ (v3 >> ((2) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v8
}
