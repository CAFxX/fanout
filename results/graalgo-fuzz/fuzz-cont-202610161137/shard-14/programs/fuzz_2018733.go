package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 716
    v3 := []string{"hello", "\n\t", "\n\t", (("foo bar" + "foo bar") + (("" + ("hello" + "")) + ""))}
    var v4 uint8 = 255
    var v5 int16 = 723
    v1[927] = -2147483648
    if ((1e10 * 0.0) != 1e10) {
        var v6 uint64 = 18446744073709551615
        var v7 string = "hello"
        switch 42 {
        case 32768:
            fmt.Println(v7, v2)
        }
        _ = v6
        _ = v7
    } else {
        var v8 int8 = 127
        _ = len(v1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v8
    }
    v4 += 0
    v4 = (v4 - v4)
    if false {
        _ = len(v1)
        var v9 uint32 = 2147483648
        v10 := GenS1{F0: "foo bar"}
        _ = v10
        _ = v9
    }
    v2 -= (v2 / ((int64(v4)) | 1))
    if true {
        v2 -= (v2 | int64(v4))
        _ = len(v3)
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 int8 = 42
                v14 := GenS2{F0: (v13 >> ((3) & 63))}
                var v15 uint16 = 391
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            v16 := GenS2{F0: -1}
            _ = fi11
            _ = v16
        }
        v4 += (v4 | (v4 | (v4 / ((255) | 1))))
    }
    switch 128 {
    case -948:
        for fi17 := 0; fi17 < 2; fi17++ {
            v2 = (v2 - v2)
            if ((v5 | v5) >= v5) {
                fmt.Println(v5, fi17)
                v4 = v4
            }
            v3 = append(v3, "\n\t")
            _ = fi17
        }
    }
    fmt.Println(v4)
    v18 := GenS3{F0: 609, F1: 111}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
