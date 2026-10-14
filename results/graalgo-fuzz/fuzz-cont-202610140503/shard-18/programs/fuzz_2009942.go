package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int32 }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]uint32)
    var v3 int64 = -1
    v4 := []int{1}
    _ = len(v2)
    v1 = v1
    v1 -= v1
    for fi5 := 0; fi5 < 15; fi5++ {
        switch fi5 {
        case -2147483648:
            v6 := GenS1{F0: 42, F1: true, F2: false}
            _ = v6
        }
        _ = fi5
    }
    v7 := GenS2{F0: 42}
    switch "hello" {
    case "a":
        v8 := GenS3{F0: -453, F1: 235}
        _ = v8
    case "foo bar":
        var v9 uint32 = 4294967295
        _ = v9
    }
    fmt.Println(v3, v3)
    var v10 string = ("foo bar" + ("hello" + "foo bar"))
    v3 = (-(v3))
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 int32 = -2147483648
            _ = fi12
            _ = v13
        }
        fmt.Println(v10, fi11, fi11)
        var v14 int8 = 1
        _ = fi11
        _ = v14
    }
    switch v10 {
    case "a":
        var v15 float32 = float32(3.14159)
        _ = v15
    case "":
        fmt.Println(v3, v10, v10)
    case "\n\t":
        v16 := GenS2{F0: 1}
        _ = v16
    default:
        var v17 uint = 65535
        _ = v17
    }
    if (65535 < 4294967295) {
        v10 = (v10 + ("\n\t" + v10))
        if false {
            fmt.Println(v1, v1, v1)
        } else {
            switch v10 {
            case "hello":
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 int16 = 1
                    _ = len(v4)
                    v20 := v2["hello"]
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            default:
                _ = len(v2)
            }
            v1 += v1
            for fi21 := 0; fi21 < 5; fi21++ {
                if ((v3 ^ -1) < (v3 ^ 0)) {
                    v22 := GenS2{F0: 197}
                    _ = v22
                }
                _ = fi21
            }
        }
        var v23 int32 = -2147483648
        _ = v23
    }
    v1 = v1
    v24 := GenS3{F0: -1, F1: -347}
    v10 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v7
}
