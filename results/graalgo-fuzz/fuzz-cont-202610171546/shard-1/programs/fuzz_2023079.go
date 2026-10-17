package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int64; F1 int16 }

func main() {
    var v1 uint8 = 196
    var v2 int = -471
    var v3 string = "\n\t"
    v4 := make(map[string]uint32)
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 8; fi6++ {
            fmt.Println(v2)
            _ = fi6
        }
        if (82 > 0) {
            fmt.Println(v3, v2, v2)
        }
        _ = fi5
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            v1 += v1
            fmt.Println(v1, v2)
            _ = fi8
        }
        _ = fi7
    }
    v1 = v1
    for fi9 := 0; fi9 < 4; fi9++ {
        if (!(!true)) {
            v10 := GenS1{F0: 221}
            _ = v10
        } else {
            if true {
                var v11 uint8 = 255
                _ = v11
            } else {
                v3 = ((v3 + v3) + "\n\t")
                v1 -= v1
            }
            v2 = fi9
        }
        _ = fi9
    }
    v2 = v2
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            if true {
                var v14 int16 = 32767
                _ = v14
            } else {
                fmt.Println(v1)
                switch ((v3 + ((((v3 + v3) + "a") + (v3 + "a")) + v3)) + (v3 + v3)) {
                case "hello":
                    v15 := []int{(fi12 | (v2 / ((fi12) | 1)))}
                    _ = v15
                case "":
                    var v16 int32 = -794
                    _ = v16
                case "a":
                    v17 := GenS2{F0: 1, F1: 42}
                    _ = v17
                }
                var v18 uint8 = v1
                _ = v18
            }
            v19 := GenS3{F0: 42, F1: 1000}
            v2 += -1
            _ = fi13
            _ = v19
        }
        v3 = "a"
        for fi20 := 0; fi20 < 1; fi20++ {
            fmt.Println(fi20, fi12, fi20)
            _ = fi20
        }
        _ = fi12
    }
    v2 = v2
    v21 := []int{(-(v2)), v2, 2147483647, (v2 >> ((18446744073709551615) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
