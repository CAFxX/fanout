package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 uint32 = 2
    v2 := []string{("\n\t" + "\n\t")}
    var v3 int8 = -110
    var v4 float64 = -2.25
    v5 := make(map[int]string)
    v6 := []int{642}
    fmt.Println(v1, v3, v1)
    var v7 uint8 = 2
    if (false && true) {
        var v8 int = 32767
        var v9 string = ("a" + "hello")
        _ = v8
        _ = v9
    } else {
        var v10 int32 = 2147483647
        v11 := v5[32767]
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        v7 = 128
        var v13 int = (fi12 | fi12)
        _ = fi12
        _ = v13
    }
    var v14 int32 = 346
    if ((^(v7)) <= v7) {
        var v15 float32 = float32(((v4 * 3.14159) / ((v4 - ((v4 / (1e10 - (v4 - v4))) - ((v4 / v4) + (((float64(3.14159) + 1e10) * 0.0) / v4)))) * (-2.25 - 1.5))))
        var v16 int64 = 0
        if (!(false && false)) {
            v5[4294967295] = "foo bar"
        } else {
            var v17 uint = 0
            var v18 int64 = v16
            _ = v17
            _ = v18
        }
        _ = v15
        _ = v16
    }
    if false {
        v1 -= 0
    } else {
        v14 = 2147483647
    }
    if (!false) {
        delete(v5, -915)
        v2 = append(v2, "foo bar")
        for fi19 := 0; fi19 < 13; fi19++ {
            v3 -= (^(v3))
            _ = fi19
        }
    }
    v20 := GenS1{F0: v14}
    v21 := GenS1{F0: (v14 + v14)}
    if true {
        _ = len(v5)
        v7 = 49
    } else {
        for fi22 := 0; fi22 < 11; fi22++ {
            fmt.Println(v7, v1)
            _ = fi22
        }
    }
    v23 := GenS2{F0: 9223372036854775807, F1: v14}
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 int64 = 1
        for fi26 := 0; fi26 < 11; fi26++ {
            v14 = 24
            v27 := GenS1{F0: v14}
            _ = fi26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v4)
    var v28 uint32 = (v1 >> ((432) & 63))
    if ((v3 << ((1) & 63)) != v3) {
        v29 := GenS3{F0: (^(v28)), F1: uint64(v7)}
        _ = v29
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
