package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }

func main() {
    var v1 int64 = -204
    var v2 int = 503
    var v3 uint64 = 42
    v1 = (v1 + (v1 | (v1 * (v1 >> ((1) & 63)))))
    fmt.Println(v2, v1)
    if false {
        var v4 uint32 = 1
        _ = v4
    }
    v5 := GenS1{F0: "a", F1: 0}
    var v6 uint16 = 2
    if true {
        var v7 int32 = 2147483647
        var v8 int32 = (v7 % (((v7 ^ -1)) | 1))
        if (false && false) {
            v3 = 2
            v1 -= 553
        } else {
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 int32 = (v7 + (^(v7)))
                for fi11 := 0; fi11 < 14; fi11++ {
                    v12 := []int{(fi9 / (((fi11 * v2)) | 1)), (-(v2)), fi9, (v2 | fi9)}
                    fmt.Println(v1, fi9, fi9)
                    _ = fi11
                    _ = v12
                }
                _ = fi9
                _ = v10
            }
            v6 = v6
        }
        _ = v7
        _ = v8
    }
    if (true || true) {
        var v13 uint16 = v6
        _ = v13
    }
    v14 := []int{v2}
    v3 = v3
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 uint32 = 698
        for fi17 := 0; fi17 < 14; fi17++ {
            fmt.Println(fi17, fi15)
            var v18 int64 = (-(v1))
            _ = fi17
            _ = v18
        }
        v16 = v16
        _ = fi15
        _ = v16
    }
    v2 -= (v2 + v2)
    var v19 bool = (("a" + "\n\t") != "hello")
    for fi20 := 0; fi20 < 12; fi20++ {
        v1 += v1
        _ = fi20
    }
    switch (v2 / (((v2 ^ int(v6))) | 1)) {
    case -318:
        _ = len(v14)
    case 127:
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS1{F0: "\n\t", F1: (^(v2))}
            _ = fi21
            _ = v22
        }
    default:
        v23 := GenS1{F0: "foo bar", F1: (v2 << ((uint(v6)) & 63))}
        _ = v23
    }
    var v24 float32 = (1.5 * ((0.0 * 1.5) / float32(1.5)))
    fmt.Println(v1)
    for fi25 := 0; fi25 < 7; fi25++ {
        fmt.Println(v3)
        v3 -= v3
        v2 -= (v2 ^ (v2 + (-(v2))))
        _ = fi25
    }
    var v26 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v5
    _ = v6
}
