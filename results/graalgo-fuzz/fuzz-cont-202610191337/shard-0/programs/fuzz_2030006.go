package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint64; F1 int }
func genFn0(p0 string) uint8 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := []int{127, fi4, 0, fi4}
        if true {
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if (!false) {
                    p0 = p0
                    p0 = (p0 + "\n\t")
                }
                p0 = p0
            }
        } else {
            var v6 int = fi4
            switch fi4 {
            case 1:
                var v7 bool = (2 > 472)
                _ = v7
            }
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    v8 := GenS1{F0: -2147483648}
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 137
}

func main() {
    v1 := []string{"hello", ("hello" + (("hello" + "\n\t") + "hello")), ("" + "foo bar")}
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    if (-1 == 1) {
        var v9 int = 32768
        v9 -= v9
        _ = v9
    } else {
        switch "foo bar" {
        case "":
            switch 71 {
            case -2147483648:
                var v10 int32 = 0
                _ = v10
            default:
                var v11 uint64 = 4294967295
                _ = v11
            }
        case "\n\t":
            var v12 int16 = -32768
            _ = v12
        case "foo bar":
            v13 := GenS1{F0: 168}
            _ = v13
        default:
            for fi14 := 0; fi14 < 14; fi14++ {
                for fi15 := 0; fi15 < 17; fi15++ {
                    fmt.Println(fi15)
                    if len(v1) > 0 {
                        v1[0] = ("\n\t" + ("hello" + "a"))
                    }
                    var v16 int64 = 1
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
        }
        v17 := GenS2{F0: 4294967295, F1: 0}
        var v18 int32 = 0
        _ = v17
        _ = v18
    }
    _ = len(v3)
    if len(v1) > 0 {
        v1[0] = ""
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 uint32 = 4294967295
        _ = fi19
        _ = v20
    }
    if (-9223372036854775808 <= 872) {
        var v21 uint8 = 128
        _ = v21
    }
    switch -748 {
    case -368:
        var v22 int64 = 361
        _ = v22
    case -1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 157:
        var v23 uint = 18446744073709551615
        _ = v23
    default:
        var v24 uint = 265
        _ = v24
    }
    var v25 float32 = 1e10
    var v26 float32 = v25
    v26 = 1.5
    fmt.Println(v26)
    v26 = (3.14159 * (0.1 - v25))
    v27 := GenS3{F0: 663, F1: 128}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
