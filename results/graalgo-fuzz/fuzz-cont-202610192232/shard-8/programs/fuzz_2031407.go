package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int32) uint8 {
    v5 := []int{-154, 120}
    p0 = p0
    _ = p0
    _ = v5
    return 0
}
func genFn1() int8 {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint = 656
        fmt.Println(fi6, fi6)
        var v8 uint16 = 900
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int16 = -1
    if ((("\n\t" + ("hello" + ("a" + "foo bar"))) + ("hello" + (("\n\t" + "\n\t") + "\n\t"))) < "hello") {
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 uint32 = 2147483647
            _ = fi10
            _ = v11
        }
        if false {
            var v12 float32 = float32((3.14159 / (1e10 + 0.0)))
            v9 = (v9 >> ((1) & 63))
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 bool = false
                _ = fi13
                _ = v14
            }
            _ = v12
        } else {
            var v15 float32 = (0.0 - (1e10 / ((float32(1.5) - 0.1) * 3.14159)))
            v16 := []int{1, -401}
            _ = v15
            _ = v16
        }
        var v17 bool = ((194 >= 255) && true)
        _ = v17
    }
    v9 -= (v9 % (((v9 - (v9 - v9))) | 1))
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-9223372036854775808, 4294967295, -303, 731}
    var v3 int16 = -1
    v4 := make(map[string]bool)
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 uint64 = 2
        var v20 bool = (2 > 951)
        var v21 uint32 = 0
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    switch "" {
    case "\n\t":
        if (-9223372036854775808 == 9223372036854775807) {
            var v22 uint16 = 0
            v3 += (^(v3))
            _ = v22
        }
    default:
        v3 -= v3
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v23 int64 = 459
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 string = ""
        _ = fi24
        _ = v25
    }
    v26 := GenS1{F0: 18446744073709551615}
    switch "foo bar" {
    case "\n\t":
        var v27 int64 = v23
        _ = v27
    case "":
        _ = len(v2)
    case "a":
        delete(v1, 2)
    }
    v3 -= v3
    v28 := GenS1{F0: 464}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v28
    _ = v3
    _ = v4
}
