package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
func genFn0() uint16 {
    if true {
        v5 := GenS1{F0: 201}
        var v6 int = 4294967295
        v6 += v6
        _ = v5
        _ = v6
    } else {
        v7 := []int{-1, -1, -1}
        if len(v7) > 0 {
            v7[0] = -9223372036854775808
        }
        var v8 int32 = 1
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int32 = -1
        switch (fi9 >> ((681) & 63)) {
        case 253:
            v10 = v10
        }
        v11 := GenS1{F0: 0}
        _ = fi9
        _ = v10
        _ = v11
    }
    return 458
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := []string{"", "\n\t", ("foo bar" + ((("\n\t" + "a") + "") + "hello"))}
    v4 := []int{601, 2147483647, 604, 128}
    var v12 string = "a"
    var v13 int32 = -759
    var v14 int = 2147483647
    switch v12 {
    case "\n\t":
        v15 := GenS2{F0: -272}
        _ = v15
    case "a":
        var v16 int32 = (v13 / ((-492) | 1))
        _ = v16
    }
    v12 = "a"
    if true {
        switch v14 {
        case -1:
            var v17 uint64 = 366
            _ = v17
        }
        v14 -= v14
    }
    switch (v14 << ((65535) & 63)) {
    case 42:
        v12 = (("hello" + "") + v12)
    default:
        v13 -= (-(v13))
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        switch (v12 + "a") {
        case "foo bar":
            v19 := v2[1]
            _ = v19
        case "":
            fmt.Println(v14, v13, v14)
        }
        _ = len(v3)
        _ = fi18
    }
    v20 := GenS3{F0: 2}
    var v21 uint16 = 415
    var v22 uint64 = 4294967295
    var v23 uint16 = v21
    if true {
        for fi24 := 0; fi24 < 14; fi24++ {
            v12 = ((("a" + "\n\t") + (("" + ((v12 + (v12 + v12)) + ("" + ("" + v12)))) + ("\n\t" + "a"))) + "a")
            var v25 float64 = 0.1
            _ = fi24
            _ = v25
        }
    } else {
        v21 += v21
        for fi26 := 0; fi26 < 15; fi26++ {
            v1[1] = false
            _ = fi26
        }
        fmt.Println(v12)
    }
    var v27 string = v12
    if len(v3) > 0 {
        v3[0] = v12
    }
    if true {
        var v28 int64 = 42
        var v29 int32 = (v13 - (-(v13)))
        _ = v28
        _ = v29
    }
    v23 = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
}
