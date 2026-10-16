package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 string; F1 int64; F2 int16 }
type GenS5 struct { F0 int32; F1 int64; F2 bool }
type GenS6 struct { F0 bool; F1 int32 }
func genFn0(p0 string, p1 int8, p2 uint32) int32 {
    v6 := GenS1{F0: 255}
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 float32 = 1e10
        _ = fi7
        _ = v8
    }
    fmt.Println(p2, p1, p1)
    p2 -= 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -2147483648
}

func main() {
    v1 := []string{}
    var v2 string = "hello"
    var v3 bool = true
    var v4 uint64 = 18446744073709551615
    var v5 uint32 = 1
    var v9 uint64 = uint64(v5)
    var v10 uint64 = (^(v4))
    v11 := GenS2{F0: "hello"}
    v4 = v10
    for fi12 := 0; fi12 < 3; fi12++ {
        switch (fi12 - 128) {
        case -9223372036854775808:
            fmt.Println(v10, v2, fi12)
        case -1:
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := GenS3{F0: int8(v5), F1: -815}
                _ = fi13
                _ = v14
            }
        case 1:
            var v15 uint16 = 2
            _ = v15
        }
        var v16 uint8 = 42
        _ = fi12
        _ = v16
    }
    v3 = false
    switch v2 {
    case "a":
        v17 := GenS4{F0: v2, F1: 305, F2: 342}
        _ = v17
    case "foo bar":
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 uint = 0
            _ = fi18
            _ = v19
        }
    }
    var v20 float64 = 1.5
    v4 = 4294967295
    v2 = v2
    var v21 uint64 = uint64(v5)
    fmt.Println(v2, v21, v5)
    fmt.Println(v20, v21, v2)
    for fi22 := 0; fi22 < 2; fi22++ {
        v21 = (v10 + v21)
        for fi23 := 0; fi23 < 13; fi23++ {
            if (167 == 0) {
                v10 = uint64(v5)
                if true {
                    v9 = 42
                    if (!false) {
                        var v24 int8 = 42
                        var v25 uint32 = 2147483647
                        if (!((v24 >> ((716) & 63)) > v24)) {
                            var v26 uint16 = 1
                            _ = v26
                        } else {
                            v24 = v24
                            v27 := GenS5{F0: 1, F1: 9223372036854775807, F2: (0 <= 182)}
                            v20 -= ((float64((float32(v20) - (-2.25 + 3.14159))) / (((v20 / 1.0) - v20) * (v20 + v20))) * v20)
                            _ = v27
                        }
                        _ = v24
                        _ = v25
                    }
                }
                fmt.Println(v5, v2, v5)
                fmt.Println(v3)
            }
            _ = fi23
        }
        _ = fi22
    }
    v28 := GenS5{F0: 2147483647, F1: 1, F2: false}
    if (true && (577 < uint(v5))) {
        var v29 int16 = 1
        v29 = (v29 | (v29 - 0))
        v30 := GenS6{F0: v3, F1: 748}
        _ = v29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
