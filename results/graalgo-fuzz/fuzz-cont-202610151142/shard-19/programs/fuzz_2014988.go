package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string; F1 uint32; F2 string }

func main() {
    var v1 uint16 = 255
    v2 := []int{238, -1, 42, 128}
    var v3 uint = 0
    var v4 uint = (v3 << (((v3 ^ (v3 - v3))) & 63))
    switch -1 {
    case 127:
        v5 := GenS1{F0: "", F1: 2}
        _ = v5
    }
    var v6 string = "hello"
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(fi7, v4, v1)
        var v8 uint32 = 65535
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 float32 = (0.0 * (1e10 - (1.5 / 1.0)))
        var v11 int16 = 942
        _ = fi9
        _ = v10
        _ = v11
    }
    if ((!true) || true) {
        var v12 float64 = (-2.25 / 3.14159)
        _ = v12
    } else {
        fmt.Println(v3)
        var v13 string = (v6 + ((v6 + v6) + v6))
        _ = v13
    }
    var v14 int32 = -1
    if true {
        for fi15 := 0; fi15 < 17; fi15++ {
            v2 = append(v2, fi15)
            v14 = v14
            _ = fi15
        }
        switch v6 {
        case "foo bar":
            for fi16 := 0; fi16 < 5; fi16++ {
                if len(v2) > 0 {
                    v2[0] = (fi16 * (fi16 ^ fi16))
                }
                _ = fi16
            }
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = 165
            }
        }
    }
    v17 := GenS2{F0: "foo bar", F1: 65535, F2: "hello"}
    v1 += (v1 >> ((80) & 63))
    var v18 int64 = 75
    v2 = append(v2, 999)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
