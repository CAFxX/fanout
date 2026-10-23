package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int8; F1 bool }
func genFn0(p0 uint8, p1 int32) int64 {
    var v5 uint32 = 2147483647
    switch -2147483648 {
    case -787:
        fmt.Println(v5)
    }
    p0 = 0
    fmt.Println(p0)
    fmt.Println(v5, v5)
    _ = p0
    _ = p1
    _ = v5
    return -221
}

func main() {
    v1 := make(map[string]int)
    var v2 int = 30
    var v3 uint8 = 42
    v4 := make(map[int]uint32)
    v2 -= (v2 * v2)
    for fi6 := 0; fi6 < 12; fi6++ {
        v2 = fi6
        var v7 bool = false
        var v8 int = v2
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int16 = 991
    fmt.Println(v2, v2, v2)
    var v10 float64 = float64((-2.25 + (0.0 - -2.25)))
    v2 = int(v3)
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        var v11 int64 = 420
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 5; fi12++ {
            if true {
                fmt.Println(fi12)
                if ((!((^(v9)) < 253)) || true) {
                    var v13 uint8 = 2
                    var v14 uint64 = 2
                    _ = v13
                    _ = v14
                }
            }
            v15 := GenS1{F0: "hello", F1: (42 <= 18446744073709551615)}
            _ = fi12
            _ = v15
        }
    case "a":
        fmt.Println(v10)
    default:
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := GenS1{F0: "\n\t", F1: ((true || true) && (0 != 13))}
            fmt.Println(v3)
            _ = fi16
            _ = v17
        }
    }
    if (!((-(v2)) > v2)) {
        v1["\n\t"] = int(v9)
        v1[(("foo bar" + ("\n\t" + "foo bar")) + "\n\t")] = v2
    } else {
        fmt.Println(v10, v9)
        v18 := GenS2{F0: 0, F1: 0}
        var v19 int = v2
        _ = v18
        _ = v19
    }
    v3 -= (v3 >> ((2) & 63))
    v20 := GenS3{F0: -109, F1: true}
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := GenS1{F0: ("hello" + ("hello" + (("" + ("hello" + "\n\t")) + ("" + "")))), F1: (0 > 100)}
        fmt.Println(fi21)
        fmt.Println(fi21, fi21, fi21)
        _ = fi21
        _ = v22
    }
    v23 := []int{v2}
    var v24 uint32 = 505
    v24 = v24
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 uint8 = (v3 >> ((570) & 63))
        _ = fi25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
