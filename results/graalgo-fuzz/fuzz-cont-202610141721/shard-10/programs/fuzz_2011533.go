package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int64; F1 string; F2 string }

func main() {
    var v1 float64 = 1.5
    var v2 float32 = 1e10
    v3 := []string{"\n\t"}
    var v4 float32 = 0.0
    var v5 int16 = 883
    if false {
        v6 := GenS1{F0: 0, F1: 624, F2: 493}
        v7 := GenS2{F0: "foo bar", F1: 611}
        _ = v6
        _ = v7
    } else {
        if false {
            var v8 bool = (true || (20 <= 0))
            _ = len(v3)
            _ = v8
        }
    }
    fmt.Println(v2, v4)
    for fi9 := 0; fi9 < 11; fi9++ {
        v4 = v2
        v4 -= 1e10
        _ = fi9
    }
    v5 = v5
    v1 = (0.0 * 1e10)
    var v10 float32 = (float32(3.14159) / v4)
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int8 = 55
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v5 = 0
        _ = fi13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v16 := GenS3{F0: -761, F1: "\n\t", F2: (("hello" + "hello") + "foo bar")}
                _ = len(v3)
                _ = v16
            }
            var v17 int8 = -93
            _ = fi15
            _ = v17
        }
        fmt.Println(v10, v1)
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
