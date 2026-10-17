package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []int{}
    var v2 uint8 = 0
    v3 := []int{-236}
    v4 := make(map[string]int)
    var v5 int16 = 239
    v6 := GenS1{F0: (v5 << ((259) & 63)), F1: uint32(v2)}
    for fi7 := 0; fi7 < 20; fi7++ {
        if false {
            v4["\n\t"] = fi7
        }
        var v8 uint = 255
        _ = fi7
        _ = v8
    }
    var v9 int8 = 1
    if true {
        switch 560 {
        case -1:
            var v10 float64 = -2.25
            _ = v10
        case 1:
            fmt.Println(v2, v5, v5)
        }
    }
    fmt.Println(v5)
    if false {
        v11 := GenS1{F0: 32767, F1: 65535}
        _ = v11
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    _ = len(v4)
    switch -65 {
    case 42:
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 float32 = 0.0
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 uint = 114
                var v16 uint16 = 42
                var v17 uint32 = 42
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi12
            _ = v13
        }
    case 213:
        for fi18 := 0; fi18 < 11; fi18++ {
            v2 += 2
            v5 = v5
            _ = fi18
        }
    case -58:
        v2 = (v2 * v2)
    }
    var v19 bool = (((false && ((42 <= 9223372036854775807) && true)) && true) && (((false && (!(!(v5 < v5)))) && false) || (!((v9 & (^(v9))) > v9))))
    if true {
        _ = len(v4)
    }
    v9 += (v9 - v9)
    if v19 {
        v9 += -1
        v4[(("foo bar" + "foo bar") + ("\n\t" + "a"))] = 128
        v20 := GenS2{F0: 18446744073709551615, F1: 271}
        _ = v20
    } else {
        if (((176 >= 300) || true) && v19) {
            switch "foo bar" {
            case "\n\t":
                v21 := GenS3{F0: -388}
                _ = v21
            case "hello":
                var v22 uint64 = 603
                _ = v22
            case "a":
                v23 := GenS2{F0: 18446744073709551615, F1: -1}
                _ = v23
            default:
                v5 = -150
            }
        }
    }
    var v24 uint = uint(v2)
    v2 = 134
    v9 = v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
