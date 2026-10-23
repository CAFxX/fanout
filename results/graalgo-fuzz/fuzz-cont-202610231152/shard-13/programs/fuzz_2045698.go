package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint32 }
func genFn0() int8 {
    v7 := GenS1{F0: (!(!("a" > "hello"))), F1: 329}
    var v8 bool = (0.0 == (0.0 * (0.1 / (1e10 - 0.0))))
    v9 := []int{127, 42, 127, -705}
    fmt.Println(v8)
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 string = ("\n\t" + ("" + (("\n\t" + "") + "\n\t")))
    var v2 int8 = 100
    var v3 float64 = 0.1
    v4 := make(map[string]string)
    var v5 uint8 = 1
    var v6 uint32 = 874
    if false {
        switch (v1 + (v1 + v1)) {
        case "\n\t":
            v10 := GenS1{F0: (v5 <= 1), F1: 42}
            _ = v10
        case "foo bar":
            delete(v4, "")
        case "a":
            v11 := GenS1{F0: true, F1: -870}
            _ = v11
        }
    }
    v12 := GenS1{F0: true, F1: 593}
    fmt.Println(v5, v1, v5)
    v3 -= float64((1.5 * (0.0 * 3.14159)))
    if false {
        switch 9223372036854775807 {
        case 42:
            var v13 int16 = 1000
            _ = v13
        case 1:
            delete(v4, (v1 + v1))
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            v5 = v5
            _ = fi14
        }
        var v15 int32 = 2147483647
        _ = v15
    } else {
        v16 := GenS1{F0: (true || true), F1: 42}
        v17 := GenS1{F0: (461 <= 65535), F1: 9223372036854775807}
        for fi18 := 0; fi18 < 3; fi18++ {
            fmt.Println(v1, fi18)
            v19 := GenS2{F0: -689, F1: -2147483648, F2: 1}
            _ = fi18
            _ = v19
        }
        _ = v16
        _ = v17
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := GenS2{F0: fi20, F1: -2147483648, F2: 841}
        delete(v4, v1)
        fmt.Println(v3)
        _ = fi20
        _ = v21
    }
    v22 := []int{4294967295, 916, 32768}
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    var v23 float32 = 1e10
    v24 := GenS3{F0: (v6 << ((2) & 63))}
    switch v1 {
    case "":
        v25 := v4[v1]
        _ = v25
    case "foo bar":
        var v26 int32 = 0
        _ = v26
    case "hello":
        _ = len(v22)
    default:
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
    }
    v27 := GenS2{F0: 0, F1: 0, F2: 281}
    var v28 float32 = 1.5
    var v29 uint8 = v5
    var v30 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
