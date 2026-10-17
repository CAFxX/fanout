package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint8 }

func main() {
    var v1 float64 = 0.0
    v2 := []int{-351}
    var v3 uint32 = 2
    v4 := []int{-852, 9223372036854775807, 2147483647}
    var v5 uint64 = 42
    var v6 uint8 = 1
    v5 = (v5 - v5)
    if (!(!true)) {
        var v7 string = "foo bar"
        _ = v7
    } else {
        for fi8 := 0; fi8 < 13; fi8++ {
            fmt.Println(v1)
            for fi9 := 0; fi9 < 2; fi9++ {
                if (2 == 682) {
                    v2 = append(v2, (fi9 << ((1) & 63)))
                    v3 = 656
                }
                _ = fi9
            }
            v6 -= v6
            _ = fi8
        }
        v6 += v6
        v2 = append(v2, 128)
    }
    var v10 int64 = 677
    fmt.Println(v6)
    for fi11 := 0; fi11 < 14; fi11++ {
        switch "foo bar" {
        case "hello":
            v12 := GenS1{F0: 692, F1: (v5 >> ((1) & 63)), F2: 0}
            _ = v12
        case "foo bar":
            var v13 uint32 = (v3 - 4294967295)
            _ = v13
        case "a":
            v1 -= (float64(0.1) + v1)
        }
        _ = fi11
    }
    if false {
        if len(v2) > 0 {
            v2[0] = 220
        }
        var v14 uint32 = v3
        _ = v14
    }
    v15 := GenS1{F0: 860, F1: (^(v5)), F2: v6}
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 uint8 = (v6 * (v6 - v6))
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS1{F0: 0, F1: (v5 / (((v5 ^ v5)) | 1)), F2: (v17 - uint8(v3))}
            _ = fi18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    var v20 int16 = -239
    switch -1 {
    case 42:
        var v21 int32 = 1
        _ = v21
    case 32768:
        fmt.Println(v3, v6, v3)
    default:
        switch "foo bar" {
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "\n\t":
            for fi22 := 0; fi22 < 16; fi22++ {
                var v23 uint = 42
                fmt.Println(v23, v20, v1)
                _ = fi22
                _ = v23
            }
        case "":
            if false {
                switch ((("hello" + (("hello" + "a") + "a")) + "foo bar") + "hello") {
                case "a":
                    var v24 int = 4294967295
                    _ = v24
                }
                var v25 uint = 2
                _ = v25
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
