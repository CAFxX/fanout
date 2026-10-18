package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 int) uint {
    fmt.Println(p0)
    var v5 int64 = 9223372036854775807
    var v6 uint32 = 2147483648
    _ = p0
    _ = v5
    _ = v6
    return 662
}

func main() {
    v1 := []string{"foo bar", "\n\t"}
    var v2 int8 = -1
    var v3 bool = true
    var v4 bool = (510 > 255)
    var v7 int64 = -31
    v8 := GenS1{F0: ((("\n\t" + "hello") + "") < "hello"), F1: 65535, F2: 455}
    var v9 bool = v3
    var v10 int32 = -207
    if v3 {
        var v11 uint16 = 42
        _ = v11
    } else {
        var v12 int32 = (v10 * (v10 * int32(v2)))
        switch "" {
        case "hello":
            fmt.Println(v10)
        case "\n\t":
            var v13 int8 = int8(v7)
            _ = v13
        case "a":
            fmt.Println(v7, v3)
        }
        var v14 float64 = -2.25
        _ = v12
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v10 = (v10 - -884)
        v4 = v9
        var v16 uint16 = 122
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: -386}
    for fi18 := 0; fi18 < 6; fi18++ {
        v1 = append(v1, "hello")
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi18
    }
    v10 = (v10 % ((v10) | 1))
    v3 = v4
    v10 = v10
    var v19 float64 = 1e10
    for fi20 := 0; fi20 < 20; fi20++ {
        v21 := GenS3{F0: fi20, F1: 42}
        var v22 int64 = (v7 & v7)
        _ = fi20
        _ = v21
        _ = v22
    }
    v1 = append(v1, "hello")
    fmt.Println(v19)
    if (65535 >= 57) {
        var v23 uint16 = 857
        var v24 uint32 = 1
        _ = v23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
