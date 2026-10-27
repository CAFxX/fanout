package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int; F1 uint32 }

func main() {
    var v1 int8 = -99
    v2 := make(map[int]string)
    var v3 string = ("foo bar" + "a")
    v4 := []uint32{286}
    var v5 uint64 = 0
    fmt.Println(v3, v5, v5)
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: "a", F1: true, F2: 1}
        _ = fi6
        _ = v7
    }
    v1 = (v1 | v1)
    _ = len(v2)
    v3 = ("foo bar" + v3)
    v5 = v5
    var v8 float64 = float64(0.1)
    switch "\n\t" {
    case "foo bar":
        v2[1] = v3
    case "a":
        v9 := GenS2{F0: 882, F1: -9223372036854775808}
        _ = v9
    }
    if true {
        v1 += 0
        v10 := GenS1{F0: v3, F1: (false && false), F2: -1}
        _ = v10
    }
    v11 := GenS2{F0: 42, F1: 2}
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 int8 = (^(v1))
            v4 = append(v4, 2147483648)
            v14 -= v1
            _ = fi13
            _ = v14
        }
        v8 = v8
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 uint = 538
            var v17 int8 = (v1 >> (((v16 ^ (^(v16)))) & 63))
            var v18 int8 = -1
            var v19 uint = (v16 ^ (^(v16)))
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi12
    }
    v20 := v2[4294967295]
    v21 := v2[127]
    if true {
        v2[2147483647] = ""
        v21 = v20
    }
    if false {
        v8 = 0.1
        if (580 != 2) {
            var v22 int8 = (v1 >> ((0) & 63))
            _ = v22
        }
        var v23 float64 = v8
        _ = v23
    } else {
        v1 = v1
        for fi24 := 0; fi24 < 18; fi24++ {
            if (!(fi24 > (fi24 + (-(fi24))))) {
                if (v8 <= float64(3.14159)) {
                    v5 -= v5
                    v25 := GenS1{F0: v20, F1: true, F2: -365}
                    _ = v25
                } else {
                    v26 := GenS2{F0: 42, F1: -272}
                    _ = v26
                }
                v5 -= v5
                if ((v1 / (((-(v1))) | 1)) != 42) {
                    var v27 uint = 634
                    v27 = v27
                    _ = v27
                }
            } else {
                v28 := GenS3{F0: (fi24 * (fi24 ^ fi24)), F1: 4294967295}
                v5 = 4294967295
                _ = v28
            }
            _ = fi24
        }
    }
    v20 = v20
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
