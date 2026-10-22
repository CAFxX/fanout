package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int; F1 int64; F2 int16 }

func main() {
    var v1 uint16 = 802
    var v2 int32 = -983
    var v3 bool = true
    v4 := make(map[int]bool)
    var v5 uint8 = 128
    v6 := []int{}
    for fi7 := 0; fi7 < 1; fi7++ {
        v6 = append(v6, (fi7 & 339))
        _ = fi7
    }
    fmt.Println(v1, v2)
    switch 128 {
    case 4294967295:
        var v8 string = ("foo bar" + "\n\t")
        _ = v8
    default:
        var v9 uint = 1
        _ = v9
    }
    v10 := GenS1{F0: 65535}
    v3 = (1 > 18446744073709551615)
    switch ((("\n\t" + "a") + "") + ((("\n\t" + ("\n\t" + "a")) + "") + ("foo bar" + ""))) {
    case "foo bar":
        v5 -= v5
    case "":
        v2 += -1
    }
    if (v3 || (((1e10 * 0.1) != -2.25) && v3)) {
        v11 := GenS2{F0: 956, F1: v3}
        if len(v6) > 0 {
            v6[0] = 1
        }
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint = 2
        if true {
            if v3 {
                fmt.Println(v5, v1)
                v14 := GenS1{F0: (v13 ^ 0)}
                var v15 int = fi12
                _ = v14
                _ = v15
            } else {
                fmt.Println(v5, v13)
                for fi16 := 0; fi16 < 3; fi16++ {
                    if (2 != uint64(v1)) {
                        var v17 int = (fi16 / ((fi16) | 1))
                        var v18 int16 = -32768
                        var v19 uint32 = 797
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        v20 := GenS3{F0: (^(fi12)), F1: 0, F2: -32768}
                        v21 := GenS3{F0: (fi12 | (fi16 / ((fi16) | 1))), F1: 832, F2: 981}
                        _ = v20
                        _ = v21
                    }
                    var v22 int32 = v2
                    v22 += (v22 ^ (v22 ^ v22))
                    _ = fi16
                    _ = v22
                }
            }
        } else {
            var v23 uint = uint(v5)
            var v24 uint32 = 2147483648
            _ = v23
            _ = v24
        }
        var v25 uint64 = 18446744073709551615
        _ = fi12
        _ = v13
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
