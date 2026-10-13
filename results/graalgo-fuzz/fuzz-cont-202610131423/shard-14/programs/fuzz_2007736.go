package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int64; F1 uint32; F2 string }
type GenS3 struct { F0 uint32; F1 int; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }

func main() {
    var v1 string = ("a" + "hello")
    v2 := make(map[string]bool)
    v3 := []int{}
    var v4 float64 = (float64((0.0 / 0.1)) / 0.1)
    var v5 uint = 628
    var v6 string = (v1 + (v1 + "foo bar"))
    v7 := GenS1{F0: -1, F1: 0}
    var v8 int64 = -409
    var v9 int = -1
    if (!(((v4 * (1e10 * v4)) * -2.25) == float64((0.1 + float32(v4))))) {
        if false {
            v10 := GenS1{F0: 42, F1: (v9 >> ((v5) & 63))}
            switch v9 {
            case 650:
                v5 = 639
            default:
                v11 := GenS2{F0: (v8 << ((v5) & 63)), F1: 0, F2: (("foo bar" + v6) + (("foo bar" + ("hello" + (v6 + (v1 + v6)))) + "hello"))}
                _ = v11
            }
            var v12 string = (v6 + v6)
            _ = v10
            _ = v12
        } else {
            if len(v3) > 0 {
                v3[0] = v9
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                fmt.Println(fi13, v4, fi13)
                _ = fi13
            }
            v14 := GenS1{F0: (-(v8)), F1: -832}
            _ = v14
        }
        var v15 uint = (^(v5))
        v1 = ("hello" + v6)
        _ = v15
    }
    var v16 uint = (v5 - v5)
    if len(v3) > 0 {
        v3[0] = int(v5)
    }
    v17 := GenS3{F0: 1, F1: (^(v9)), F2: v8}
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 string = ("\n\t" + v6)
        switch v1 {
        case "foo bar":
            var v20 int32 = int32(v16)
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        switch ("foo bar" + "foo bar") {
        case "hello":
            v22 := GenS4{F0: (-(fi21))}
            _ = v22
        case "a":
            fmt.Println(fi21, v9)
        default:
            var v23 uint32 = 2147483647
            _ = v23
        }
        v16 = 255
        _ = fi21
    }
    switch v9 {
    case 42:
        v24 := GenS4{F0: v9}
        _ = v24
    case 128:
        switch ((v1 + "\n\t") + ((v1 + "foo bar") + (v1 + (v1 + (v6 + v6))))) {
        case "a":
            var v25 int8 = 40
            _ = v25
        case "":
            if false {
                var v26 float64 = ((v4 * v4) * float64(0.0))
                v9 = v9
                _ = v26
            }
        default:
            var v27 bool = (v1 < v1)
            _ = v27
        }
    case 32767:
        for fi28 := 0; fi28 < 2; fi28++ {
            v2["a"] = true
            v29 := GenS2{F0: v8, F1: 1, F2: (v1 + v1)}
            v30 := GenS5{F0: "\n\t"}
            _ = fi28
            _ = v29
            _ = v30
        }
    default:
        var v31 uint64 = 18446744073709551615
        _ = v31
    }
    v8 -= (v8 + (^(v8)))
    var v32 int8 = 127
    v6 = v1
    v2[v6] = true
    if true {
        v16 += (v5 & v5)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v1)
    }
    var v33 uint = 619
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
