package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int }

func main() {
    var v1 uint = 1
    v2 := []uint32{897, 814, 2147483647, 0}
    v3 := make(map[int]uint32)
    var v4 int = -363
    var v5 int8 = int8(v4)
    var v6 uint64 = 42
    fmt.Println(v6)
    if (161 != -2147483648) {
        var v7 float64 = -2.25
        if (!(false || (!true))) {
            v7 = v7
            switch ("hello" + "") {
            case "\n\t":
                v8 := GenS1{F0: ""}
                _ = v8
            }
            switch "hello" {
            case "":
                v7 -= 1e10
            default:
                var v9 uint = v1
                _ = v9
            }
        }
        _ = v7
    }
    v10 := GenS1{F0: "a"}
    var v11 string = "foo bar"
    v4 += v4
    fmt.Println(v6, v6)
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            fmt.Println(v5)
            switch v4 {
            case 316:
                var v14 uint32 = uint32(v1)
                _ = v14
            default:
                v1 = 1
            }
            v15 := GenS1{F0: v11}
            _ = fi13
            _ = v15
        }
        _ = fi12
    }
    v2 = append(v2, 554)
    switch 9223372036854775807 {
    case 0:
        v6 = (v6 % ((v6) | 1))
    case 9223372036854775807:
        var v16 int16 = -278
        _ = v16
    case -33:
        if (32767 >= 42) {
            v3[v4] = 478
            for fi17 := 0; fi17 < 13; fi17++ {
                v1 = (v1 / ((v1) | 1))
                var v18 float32 = 1.5
                delete(v3, (fi17 & 1))
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 14; fi19++ {
                fmt.Println(v6, fi19)
                _ = fi19
            }
        } else {
            if false {
                if (42 == (v6 - v6)) {
                    delete(v3, (v4 | v4))
                    var v20 int = (^(v4))
                    _ = v20
                } else {
                    v21 := GenS1{F0: v11}
                    _ = v21
                }
                v3[v4] = 1
            } else {
                for fi22 := 0; fi22 < 7; fi22++ {
                    var v23 int64 = 640
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v25 := GenS2{F0: 1000, F1: fi24}
                        _ = len(v3)
                        _ = fi24
                        _ = v25
                    }
                    _ = len(v3)
                    _ = fi22
                    _ = v23
                }
            }
            v6 += 0
        }
    }
    var v26 int8 = (v5 % ((v5) | 1))
    var v27 uint32 = uint32(v1)
    fmt.Println(v5, v1, v6)
    v6 += (v6 ^ (v6 + v6))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
