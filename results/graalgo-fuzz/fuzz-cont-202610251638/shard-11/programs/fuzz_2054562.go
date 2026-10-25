package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint }

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 109
    var v3 uint = 156
    var v4 int = 32767
    v2 = 127
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := v1["\n\t"]
        var v7 int8 = 42
        _ = fi5
        _ = v6
        _ = v7
    }
    v4 -= 0
    switch "" {
    case "":
        if false {
            var v8 int64 = -1
            _ = v8
        } else {
            fmt.Println(v3)
        }
    }
    var v9 float64 = -2.25
    switch ("" + "hello") {
    case "a":
        fmt.Println(v2, v4, v9)
    case "hello":
        v1[((("foo bar" + "") + "") + ("a" + ("foo bar" + "\n\t")))] = (("foo bar" + "a") + "")
    default:
        v9 = 0.1
    }
    v10 := v1[("" + ("\n\t" + "hello"))]
    v11 := v1[v10]
    v12 := GenS1{F0: (v4 - (v4 / (((v4 - -2147483648)) | 1))), F1: 0, F2: 42}
    v13 := GenS1{F0: -2147483648, F1: (v3 | 696), F2: 969}
    v10 = "foo bar"
    if (true && (true && false)) {
        if (true || false) {
            var v14 uint64 = 549
            _ = v14
        }
    }
    var v15 bool = true
    if (-355 > 477) {
        var v16 uint16 = 255
        var v17 int64 = 0
        _ = v16
        _ = v17
    } else {
        var v18 uint = v3
        if v15 {
            v19 := GenS2{F0: 471, F1: 560}
            var v20 uint16 = 723
            _ = v19
            _ = v20
        }
        if false {
            switch v4 {
            case 128:
                var v21 int = (v4 % (((v4 - (v4 - (v4 & v4)))) | 1))
                _ = v21
            }
        } else {
            var v22 int64 = -520
            _ = v22
        }
        _ = v18
    }
    switch (v10 + (v11 + v10)) {
    case "\n\t":
        v1[v10] = (v11 + "\n\t")
    default:
        var v23 uint32 = 2
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
