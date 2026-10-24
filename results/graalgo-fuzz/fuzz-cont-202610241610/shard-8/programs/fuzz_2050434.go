package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 float32 = float32(-2.25)
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    var v4 int8 = 127
    v5 := []string{"hello", ("\n\t" + ("foo bar" + "a"))}
    var v6 uint16 = 134
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS1{F0: (-(fi8))}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    switch 4294967295 {
    case -254:
        v10 := GenS1{F0: 4294967295}
        _ = v10
    default:
        v11 := GenS1{F0: -353}
        _ = v11
    }
    var v12 int32 = -2147483648
    v4 = (v4 << ((255) & 63))
    var v13 uint = 65535
    v14 := GenS1{F0: 1}
    for fi15 := 0; fi15 < 20; fi15++ {
        v4 -= (v4 ^ (v4 - v4))
        v16 := GenS2{F0: v6, F1: v6, F2: v13}
        for fi17 := 0; fi17 < 6; fi17++ {
            switch ("a" + "foo bar") {
            case "a":
                var v18 uint32 = 65535
                _ = v18
            case "\n\t":
                if (!false) {
                    v19 := GenS3{F0: uint64(v6)}
                    var v20 bool = true
                    _ = v19
                    _ = v20
                }
            case "foo bar":
                var v21 uint32 = 62
                _ = v21
            default:
                v1 = v1
            }
            var v22 int16 = -817
            var v23 bool = true
            _ = fi17
            _ = v22
            _ = v23
        }
        _ = fi15
        _ = v16
    }
    v24 := GenS1{F0: 134}
    _ = len(v5)
    v1 -= v1
    var v25 bool = ((("foo bar" + "a") + "hello") >= "a")
    for fi26 := 0; fi26 < 10; fi26++ {
        fmt.Println(v13, v13)
        _ = fi26
    }
    switch "a" {
    case "a":
        var v27 float32 = v1
        _ = v27
    }
    switch 127 {
    case 4294967295:
        v12 += (v12 * -2147483648)
    default:
        var v28 float64 = 1e10
        _ = v28
    }
    var v29 uint = (^(v13))
    var v30 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
