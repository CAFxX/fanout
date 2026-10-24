package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 uint = 0
    v2 := make(map[string]int)
    var v3 float64 = (1e10 * 1.5)
    var v4 float64 = (float64(0.0) * v3)
    var v5 float32 = (0.1 - (0.1 / 1e10))
    var v6 uint8 = uint8(v1)
    v7 := []int{-836}
    v8 := GenS1{F0: 410, F1: -1, F2: 2}
    v9 := GenS2{F0: -374}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v10 int64 = 0
    var v11 int64 = v10
    var v12 string = "a"
    switch v12 {
    case "foo bar":
        v6 = 127
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS3{F0: 0, F1: 0}
        switch fi13 {
        case -148:
            switch (v12 + "hello") {
            case "a":
                if (1 != 272) {
                    if false {
                        _ = len(v2)
                    } else {
                        for fi15 := 0; fi15 < 9; fi15++ {
                            v16 := GenS4{F0: 4294967295, F1: 18446744073709551615}
                            fmt.Println(v5, v5)
                            _ = fi15
                            _ = v16
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    v6 = (^(v6))
                }
            }
        default:
            var v17 uint16 = 255
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
