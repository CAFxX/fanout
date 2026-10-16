package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint32, p1 int) bool {
    fmt.Println(p0, p1)
    var v5 string = (("\n\t" + "\n\t") + "hello")
    if (v5 > ((v5 + v5) + ("" + ""))) {
        var v6 int32 = 2147483647
        fmt.Println(p1, p1)
        var v7 string = (v5 + v5)
        _ = v6
        _ = v7
    } else {
        fmt.Println(v5, p0, p1)
        var v8 bool = false
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return (-1 > 100)
}

func main() {
    var v1 uint = 2
    var v2 int16 = 42
    var v3 int64 = 9223372036854775807
    v4 := []string{("hello" + (("\n\t" + "foo bar") + ("\n\t" + (("foo bar" + "foo bar") + "\n\t")))), "hello", "hello"}
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int = (-(fi9))
        var v11 uint16 = 255
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 int = (-(fi9))
            _ = fi12
            _ = v13
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "" {
    case "a":
        v3 = v3
    case "\n\t":
        v2 += (-(v2))
    }
    v1 = (v1 / ((v1) | 1))
    var v14 int = int(v3)
    var v15 bool = false
    v2 -= (v2 >> (((v1 + (^(v1)))) & 63))
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(v15)
        v14 = -67
        _ = fi16
    }
    switch v14 {
    case -137:
        v4 = append(v4, ("\n\t" + ("a" + "")))
    case 0:
        v14 -= (v14 << (((^(v1))) & 63))
    }
    if ((v15 && (-2.25 >= 0.0)) && true) {
        v17 := GenS1{F0: 1, F1: 32767}
        _ = v17
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
