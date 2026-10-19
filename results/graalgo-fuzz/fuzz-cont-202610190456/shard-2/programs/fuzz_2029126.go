package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0(p0 string, p1 int8, p2 int16) uint32 {
    var v5 uint32 = 42
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := []int{2147483647, 42, 884}
        switch fi6 {
        case -1:
            if true {
                p0 = ((p0 + "foo bar") + p0)
                switch p0 {
                case "a":
                    p2 = p2
                case "foo bar":
                    var v8 int64 = 0
                    _ = v8
                case "\n\t":
                    if len(v7) > 0 {
                        v7[0] = fi6
                    }
                }
                var v9 float64 = 3.14159
                _ = v9
            }
        case -553:
            v10 := GenS1{F0: 2}
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    if (255 != 65535) {
        v11 := []int{32767, 9223372036854775807, 2147483647}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return v5
}

func main() {
    var v1 float32 = 1.5
    var v2 uint8 = 0
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    v1 = 1.5
    var v12 uint16 = 0
    fmt.Println(v1, v12, v2)
    v13 := GenS2{F0: (!false)}
    v2 += (v2 + v2)
    for fi14 := 0; fi14 < 1; fi14++ {
        if ((42 == 546) && true) {
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 int = 42
                _ = fi15
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 uint8 = v2
                var v19 int16 = -300
                v2 += v18
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS3{F0: 0, F1: -117}
            v12 -= (^(v12))
            _ = v20
        }
        var v21 float32 = 1.5
        if (true && true) {
            v22 := []int{fi14}
            _ = v22
        }
        _ = fi14
        _ = v21
    }
    fmt.Println(v1, v1)
    v23 := []int{-29, 32768}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
