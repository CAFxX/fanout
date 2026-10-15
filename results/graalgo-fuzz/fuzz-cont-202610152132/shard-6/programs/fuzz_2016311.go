package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint8 = 2
    v2 := []uint32{198}
    v3 := []int{4294967295}
    v4 := make(map[int]string)
    var v5 uint8 = v1
    fmt.Println(v5, v1)
    for fi6 := 0; fi6 < 2; fi6++ {
        if ((!false) || (958 < 0)) {
            v5 = 1
            fmt.Println(v1, fi6, v5)
        }
        fmt.Println(v1)
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := GenS1{F0: -32768}
            v4[fi6] = "a"
            switch ("foo bar" + ("hello" + "\n\t")) {
            case "\n\t":
                v9 := GenS2{F0: -2147483648, F1: 901, F2: 0}
                _ = v9
            case "a":
                v5 += v1
            case "hello":
                v10 := GenS2{F0: 0, F1: 1, F2: 249}
                _ = v10
            default:
                v11 := GenS1{F0: -56}
                _ = v11
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    if ((-2147483648 < 0) || (2 >= 883)) {
        switch "foo bar" {
        case "":
            switch 128 {
            case 127:
                v1 -= v1
            case 32767:
                v5 = (v5 / (((^(v1))) | 1))
            case 9223372036854775807:
                v1 -= v5
            }
        }
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        fmt.Println(v5, v5)
        _ = fi12
    }
    _ = len(v2)
    v5 = v5
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS2{F0: 0, F1: 281, F2: (v5 % ((v5) | 1))}
        fmt.Println(v5, v5, v5)
        _ = fi13
        _ = v14
    }
    var v15 int = 360
    var v16 bool = ((!(true && (true && true))) && false)
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS3{F0: "foo bar"}
        v4[(-(v15))] = "a"
        v19 := GenS3{F0: "a"}
        _ = fi17
        _ = v18
        _ = v19
    }
    v4[v15] = "\n\t"
    var v20 int = v15
    if true {
        v21 := GenS3{F0: "a"}
        var v22 int32 = 686
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS4{F0: 42, F1: (("hello" + ("hello" + "hello")) + (("a" + "foo bar") + (("foo bar" + "\n\t") + "hello")))}
            for fi25 := 0; fi25 < 3; fi25++ {
                fmt.Println(v5, v20, v16)
                _ = fi25
            }
            v26 := GenS5{F0: (v5 | 2), F1: (v16 && v16)}
            _ = fi23
            _ = v24
            _ = v26
        }
        _ = v21
        _ = v22
    }
    v27 := GenS2{F0: 2147483647, F1: 42, F2: 42}
    var v28 bool = (((v1 * (v5 >> ((2) & 63))) != v1) && v16)
    v29 := GenS3{F0: "a"}
    var v30 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
