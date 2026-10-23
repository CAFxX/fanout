package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{"hello"}
    v3 := []uint32{}
    var v4 int8 = 100
    v4 += v4
    var v5 uint32 = 42
    delete(v1, 0)
    v5 = 65535
    var v6 uint16 = 65535
    var v7 bool = true
    var v8 uint = 255
    _ = len(v3)
    if (-9223372036854775808 > int(v8)) {
        for fi9 := 0; fi9 < 1; fi9++ {
            v8 += v8
            _ = fi9
        }
    } else {
        v10 := GenS1{F0: 684}
        _ = v10
    }
    if (!(42 <= -1)) {
        v8 = 440
    }
    var v11 float32 = (3.14159 * (float32(float64(-2.25)) * 0.1))
    for fi12 := 0; fi12 < 5; fi12++ {
        switch "hello" {
        case "hello":
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v4)
            }
        default:
            var v13 uint32 = 618
            _ = v13
        }
        if false {
            v14 := GenS2{F0: "foo bar", F1: (2 > 0)}
            _ = v14
        }
        var v15 float32 = -2.25
        _ = fi12
        _ = v15
    }
    switch -177 {
    case 620:
        v8 = v8
    case -1:
        var v16 int8 = v4
        _ = v16
    case -2147483648:
        if len(v2) > 0 {
            v2[0] = (("\n\t" + "\n\t") + "\n\t")
        }
    default:
        if (-1 <= 0) {
            if v7 {
                v17 := GenS3{F0: (v5 * (^(v5)))}
                var v18 uint8 = 255
                _ = v17
                _ = v18
            }
            fmt.Println(v6, v4)
            v19 := GenS2{F0: "\n\t", F1: (!(false || (!false)))}
            _ = v19
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
