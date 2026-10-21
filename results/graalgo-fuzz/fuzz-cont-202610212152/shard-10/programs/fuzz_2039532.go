package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int64, p1 int16, p2 string) int16 {
    p0 = int64(p1)
    v5 := GenS1{F0: -128, F1: 2147483647}
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int = (fi6 - (fi6 / (((fi6 >> ((466) & 63))) | 1)))
        _ = fi6
        _ = v7
    }
    v8 := []int{}
    v9 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    _ = v9
    return int16(p0)
}
func genFn1(p0 bool) int {
    p0 = p0
    switch "a" {
    case "a":
        switch 42 {
        case -998:
            p0 = p0
        default:
            v10 := GenS2{F0: 86}
            _ = v10
        }
    default:
        p0 = (255 != 134)
    }
    _ = p0
    return -517
}

func main() {
    var v1 int32 = 2147483647
    v2 := []string{"foo bar", "a"}
    var v3 bool = false
    var v4 uint16 = 32767
    _ = len(v2)
    for fi11 := 0; fi11 < 10; fi11++ {
        v1 -= -2147483648
        _ = fi11
    }
    fmt.Println(v3)
    v1 = v1
    v3 = ("" >= "")
    v3 = false
    v4 += v4
    fmt.Println(v3)
    for fi12 := 0; fi12 < 5; fi12++ {
        if (!v3) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                switch "" {
                case "hello":
                    var v14 int8 = 42
                    _ = v14
                case "foo bar":
                    v4 -= (v4 - (v4 % ((32767) | 1)))
                }
                v15 := GenS1{F0: -3, F1: genFn1((0 >= 42))}
                v1 = (v1 / ((42) | 1))
                _ = fi13
                _ = v15
            }
        } else {
            var v16 uint64 = 4294967295
            switch 127 {
            case 42:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 1:
                if false {
                    if (553 >= 42) {
                        if v3 {
                            var v17 uint = uint(v4)
                            var v18 int16 = int16(v17)
                            fmt.Println(v17, v3)
                            _ = v17
                            _ = v18
                        }
                    }
                    _ = len(v2)
                    v19 := GenS2{F0: 42}
                    _ = v19
                }
            case 916:
                var v20 float64 = -2.25
                _ = v20
            }
            _ = v16
        }
        _ = len(v2)
        _ = fi12
    }
    var v21 int64 = -951
    fmt.Println(v3, v3, v1)
    v21 -= 0
    var v22 uint8 = 38
    v21 += (v21 << ((1) & 63))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        for fi23 := 0; fi23 < 16; fi23++ {
            v24 := GenS3{F0: v21}
            _ = fi23
            _ = v24
        }
        v21 = (v21 % ((v21) | 1))
    }
    var v25 int64 = (v21 % ((v21) | 1))
    var v26 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
