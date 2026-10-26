package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    var v3 float32 = 0.0
    v4 := make(map[uint32]bool)
    v5 := make(map[int]uint32)
    v6 := []int{32767, -9223372036854775808, 1}
    if (((0 == 45) && false) || false) {
        if false {
            var v7 uint32 = 0
            v8 := GenS1{F0: 1, F1: (("foo bar" + "\n\t") + ("a" + "a"))}
            _ = v7
            _ = v8
        }
        if true {
            var v9 string = ("" + "foo bar")
            for fi10 := 0; fi10 < 10; fi10++ {
                v3 -= -2.25
                var v11 uint64 = 717
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 12; fi12++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v9 = "hello"
                if true {
                    v9 = v9
                }
                _ = fi12
            }
            _ = v9
        } else {
            var v13 string = (((("hello" + "\n\t") + ("" + "hello")) + ("\n\t" + ("foo bar" + "a"))) + "a")
            _ = v13
        }
        var v14 int = 393
        _ = v14
    }
    if true {
        var v15 int16 = 0
        var v16 int32 = 861
        _ = v15
        _ = v16
    }
    v3 = (0.0 * ((0.0 + 0.0) * v3))
    if (42 >= 42) {
        _ = len(v1)
        if false {
            switch 1 {
            case 813:
                var v17 uint32 = 0
                _ = v17
            default:
                var v18 float64 = 1.5
                _ = v18
            }
            _ = len(v4)
        } else {
            var v19 int8 = -89
            _ = v19
        }
        var v20 int = 548
        _ = v20
    } else {
        v21 := GenS2{F0: 128, F1: 177}
        if (!(255 > 2)) {
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
        } else {
            v3 -= v3
            var v22 int64 = 1
            _ = v22
        }
        if len(v6) > 0 {
            v6[0] = 2147483647
        }
        _ = v21
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint64 = 18446744073709551615
        fmt.Println(fi23, fi23)
        fmt.Println(v24, v24)
        _ = fi23
        _ = v24
    }
    v25 := GenS2{F0: -2147483648, F1: 0}
    switch ("a" + "a") {
    case "":
        v26 := GenS3{F0: (false && (!false))}
        _ = v26
    default:
        for fi27 := 0; fi27 < 6; fi27++ {
            fmt.Println(fi27, fi27)
            v28 := GenS3{F0: false}
            _ = fi27
            _ = v28
        }
    }
    v3 += v3
    v3 = 1e10
    fmt.Println(v3)
    fmt.Println(v3)
    var v29 float32 = v3
    var v30 uint = 255
    delete(v5, 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
