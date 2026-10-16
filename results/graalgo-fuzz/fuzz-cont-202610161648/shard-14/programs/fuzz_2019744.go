package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int32 }
type GenS2 struct { F0 int32; F1 uint64; F2 bool }

func main() {
    var v1 uint64 = 42
    var v2 int16 = -1
    v3 := make(map[string]uint32)
    var v4 float32 = (float32((-2.25 / 0.1)) * float32((3.14159 - ((1e10 - 1e10) - (1e10 / 3.14159)))))
    v5 := []string{("hello" + (("hello" + ("a" + (("" + ("\n\t" + "")) + "foo bar"))) + "hello"))}
    var v6 string = ("foo bar" + "a")
    var v7 uint64 = (v1 % ((325) | 1))
    switch (v6 + "hello") {
    case "":
        var v8 float64 = -2.25
        _ = v8
    case "\n\t":
        fmt.Println(v6, v1, v1)
    }
    v7 = v1
    v9 := v3[v6]
    switch "\n\t" {
    case "\n\t":
        if (false && true) {
            v7 -= v1
            var v10 uint32 = 65535
            v2 -= (v2 | (v2 ^ v2))
            _ = v10
        }
    case "":
        v7 -= v1
    case "hello":
        if (100 != 0) {
            v7 -= v1
        }
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        _ = len(v3)
        _ = fi11
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 string = v6
        _ = fi12
        _ = v13
    }
    v9 -= v9
    fmt.Println(v4, v9)
    var v14 uint64 = v1
    var v15 string = ""
    v6 = v15
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS1{F0: 1, F1: v6, F2: -2147483648}
        var v18 float64 = 1e10
        switch "foo bar" {
        case "":
            fmt.Println(v14, v2)
        case "foo bar":
            var v19 float32 = ((((v4 + 1.5) + (((v4 * v4) / float32((3.14159 * v18))) / v4)) * (float32((v18 - float64(v4))) + 1.5)) + ((v4 / v4) - float32(float64(v4))))
            _ = v19
        case "a":
            v18 -= (v18 * (v18 * v18))
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    v20 := GenS1{F0: -646, F1: v15, F2: 0}
    for fi21 := 0; fi21 < 17; fi21++ {
        if (42 <= v2) {
            delete(v3, ("a" + v15))
            v3[""] = (v9 - (v9 / ((v9) | 1)))
            var v22 int16 = v2
            _ = v22
        }
        _ = fi21
    }
    v23 := GenS2{F0: 0, F1: 1, F2: (-9223372036854775808 < -2147483648)}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
