package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 string }
type GenS2 struct { F0 int; F1 uint64; F2 string }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2147483648, 348, 48}
    v3 := []string{"\n\t", "foo bar"}
    var v4 uint = 18446744073709551615
    v5 := make(map[string]uint32)
    var v6 uint16 = 2
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: -32768, F1: "hello", F2: ("" + ("" + "foo bar"))}
        v9 := GenS2{F0: (fi7 ^ (^(fi7))), F1: 4294967295, F2: "a"}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint16 = (v6 | v6)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        switch fi12 {
        case -385:
            v6 += 2
        case -1:
            for fi13 := 0; fi13 < 20; fi13++ {
                v6 -= 65535
                _ = fi13
            }
        }
        v6 = 42
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi12
    }
    if len(v3) > 0 {
        v3[0] = ""
    }
    var v14 int32 = -808
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS3{F0: 2147483647, F1: 42}
        v17 := GenS4{F0: (^(fi15))}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 bool = ((1.5 - (0.1 * (0.0 - (3.14159 / ((float64(-2.25) - 1.5) * ((((3.14159 / 1.0) + 1.5) - float64(3.14159)) + 0.1)))))) <= 0.1)
    var v19 string = "hello"
    delete(v5, v19)
    var v20 bool = (!(-2147483648 > 128))
    v18 = (!(float32(float64(3.14159)) < (1.5 - 1e10)))
    var v21 int64 = -349
    for fi22 := 0; fi22 < 19; fi22++ {
        fmt.Println(v21, v21, fi22)
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
