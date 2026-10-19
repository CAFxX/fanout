package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }

func main() {
    v1 := []uint32{0, 2147483648, 0, 0}
    v2 := make(map[int]uint32)
    var v3 int64 = -346
    v4 := []string{"\n\t", "foo bar", ("foo bar" + ("a" + "a")), ""}
    v5 := GenS1{F0: (v3 / (((v3 | -1)) | 1)), F1: (-12 > -1)}
    v3 += (v3 + -35)
    switch -2147483648 {
    case -718:
        if (false || (("foo bar" + "foo bar") > ("a" + "a"))) {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v3)
            case "foo bar":
                fmt.Println(v3)
            default:
                v3 -= (v3 / (((v3 >> ((65535) & 63))) | 1))
            }
            var v6 uint = 0
            _ = v6
        } else {
            switch ("foo bar" + "") {
            case "\n\t":
                v3 += (^(v3))
            }
            var v7 uint16 = 32767
            v8 := GenS2{F0: 861, F1: (v7 + (^(v7)))}
            _ = v7
            _ = v8
        }
    }
    fmt.Println(v3)
    fmt.Println(v3)
    v9 := GenS1{F0: (v3 >> ((1) & 63)), F1: (37 > 255)}
    v3 += v3
    v3 = (^(v3))
    var v10 bool = false
    v3 += (v3 >> ((24) & 63))
    switch 128 {
    case -864:
        v3 -= v3
    case 663:
        for fi11 := 0; fi11 < 7; fi11++ {
            if v10 {
                if v10 {
                    switch (fi11 ^ fi11) {
                    case -899:
                        delete(v2, -2147483648)
                    case 220:
                        v10 = (0 >= 42)
                    case 127:
                        v12 := GenS3{F0: 223}
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 1; fi13++ {
                        v14 := GenS4{F0: fi11}
                        v2[fi11] = 2
                        switch "hello" {
                        case "\n\t":
                            if (12 >= 180) {
                                var v15 uint64 = 1
                                _ = v15
                            } else {
                                var v16 int32 = 2147483647
                                var v17 float32 = float32(1e10)
                                _ = v16
                                _ = v17
                            }
                        default:
                            _ = len(v4)
                        }
                        _ = fi13
                        _ = v14
                    }
                }
                v3 = 0
                delete(v2, (fi11 >> ((255) & 63)))
            }
            v10 = (42 >= fi11)
            v10 = (-1 != 127)
            _ = fi11
        }
    case 904:
        v10 = (200 > 42)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
