package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int8; F1 bool; F2 int64 }
func genFn0() int {
    var v5 int8 = 100
    v6 := []int{323, 127}
    v5 += v5
    _ = v5
    _ = v6
    return 128
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 32767
    var v3 uint16 = 65535
    var v4 uint = 255
    if true {
        v7 := GenS1{F0: 255}
        _ = v7
    } else {
        v8 := GenS1{F0: (v3 << ((v4) & 63))}
        _ = v8
    }
    v9 := GenS2{F0: ("" + ""), F1: 127}
    v2 = 1000
    var v10 uint16 = v3
    var v11 uint16 = v10
    v12 := GenS2{F0: "hello", F1: -126}
    switch ("\n\t" + ("\n\t" + "foo bar")) {
    case "\n\t":
        v13 := GenS3{F0: -73, F1: ((float64(0.0) * 3.14159) <= -2.25), F2: -1}
        _ = v13
    default:
        var v14 int16 = (v2 | (v2 - (v2 & v2)))
        _ = v14
    }
    _ = len(v1)
    v2 -= v2
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 uint32 = 4294967295
        for fi17 := 0; fi17 < 1; fi17++ {
            if (!true) {
                v18 := []int{}
                _ = v18
            }
            var v19 uint16 = 32767
            _ = fi17
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
