package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
func genFn0(p0 int, p1 int64) int32 {
    fmt.Println(p1, p0)
    if false {
        var v5 int8 = 127
        switch (("hello" + "foo bar") + "") {
        case "\n\t":
            fmt.Println(p0, p1, v5)
        case "a":
            v6 := GenS1{F0: 1, F1: p0}
            _ = v6
        default:
            for fi7 := 0; fi7 < 5; fi7++ {
                v8 := GenS2{F0: (^(v5))}
                v9 := GenS1{F0: 42, F1: (-(p0))}
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        v10 := []int{(p0 / ((p0) | 1)), p0, p0}
        _ = v10
        _ = v5
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{687, 453, 4294967295}
    v3 := make(map[int]int)
    var v4 float64 = float64((-2.25 + float32(0.1)))
    v11 := GenS2{F0: 42}
    v4 = float64(0.0)
    switch "a" {
    case "":
        fmt.Println(v4)
    }
    if false {
        for fi12 := 0; fi12 < 9; fi12++ {
            for fi13 := 0; fi13 < 20; fi13++ {
                fmt.Println(fi12)
                v4 = v4
                _ = fi13
            }
            if false {
                var v14 uint64 = 18446744073709551615
                v15 := GenS1{F0: 1, F1: 42}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v14
                _ = v15
            }
            _ = fi12
        }
    }
    delete(v1, -1)
    v4 = ((v4 + v4) + v4)
    v4 = v4
    var v16 int8 = 127
    fmt.Println(v16, v16)
    for fi17 := 0; fi17 < 8; fi17++ {
        switch 863 {
        case 32767:
            var v18 float64 = -2.25
            _ = v18
        case 128:
            v16 = (v16 - (^(v16)))
        case 42:
            var v19 int64 = -1
            _ = v19
        default:
            v20 := GenS1{F0: 233, F1: fi17}
            _ = v20
        }
        if (32767 > 755) {
            if false {
                var v21 int = 4294967295
                var v22 int16 = 0
                _ = v21
                _ = v22
            }
        }
        _ = fi17
    }
    if true {
        var v23 int = 128
        switch "\n\t" {
        case "":
            delete(v1, v23)
        }
        _ = v23
    } else {
        v16 += (v16 % (((v16 + (v16 & (v16 ^ (-(v16)))))) | 1))
        for fi24 := 0; fi24 < 3; fi24++ {
            fmt.Println(v16)
            switch ("\n\t" + "") {
            case "\n\t":
                var v25 int64 = 42
                _ = v25
            case "":
                var v26 int = 1
                _ = v26
            default:
                v27 := GenS3{F0: ""}
                _ = v27
            }
            _ = fi24
        }
    }
    var v28 bool = ("" < "hello")
    var v29 int8 = v16
    v30 := GenS3{F0: ""}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
