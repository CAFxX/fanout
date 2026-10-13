package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 string }

func main() {
    v1 := []string{}
    v2 := make(map[int]uint32)
    var v3 int32 = 42
    var v4 uint = 2
    v5 := GenS1{F0: (false && (!true)), F1: -128}
    switch -9223372036854775808 {
    case -910:
        fmt.Println(v3, v4)
    default:
        v6 := GenS1{F0: true, F1: 127}
        _ = v6
    }
    var v7 int16 = -32768
    if false {
        for fi8 := 0; fi8 < 7; fi8++ {
            v4 -= v4
            for fi9 := 0; fi9 < 6; fi9++ {
                v7 = v7
                v10 := GenS1{F0: (v4 < (v4 - v4)), F1: 35}
                fmt.Println(v7, v7, v7)
                _ = fi9
                _ = v10
            }
            if len(v1) > 0 {
                v1[0] = "a"
            }
            _ = fi8
        }
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        v4 += 255
        v4 -= v4
        _ = fi11
    }
    if true {
        v12 := GenS1{F0: ((!(!true)) && true), F1: 100}
        _ = v12
    }
    v4 = (v4 - v4)
    var v13 uint64 = 331
    if (-383 < -9223372036854775808) {
        v14 := GenS2{F0: "hello"}
        var v15 string = (("a" + ("hello" + "a")) + ("foo bar" + "a"))
        _ = v14
        _ = v15
    }
    var v16 float64 = -2.25
    fmt.Println(v4, v7, v16)
    switch 133 {
    case -1:
        if (32767 >= 598) {
            for fi17 := 0; fi17 < 2; fi17++ {
                v4 -= v4
                fmt.Println(v4, fi17, fi17)
                v7 += 32767
                _ = fi17
            }
            fmt.Println(v7, v3, v3)
            for fi18 := 0; fi18 < 1; fi18++ {
                fmt.Println(v13, v13)
                var v19 int16 = (^(v7))
                var v20 string = ("hello" + "foo bar")
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    }
    fmt.Println(v3, v3)
    switch "" {
    case "a":
        for fi21 := 0; fi21 < 2; fi21++ {
            v3 += v3
            var v22 int = 701
            _ = fi21
            _ = v22
        }
    case "":
        var v23 string = "hello"
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
