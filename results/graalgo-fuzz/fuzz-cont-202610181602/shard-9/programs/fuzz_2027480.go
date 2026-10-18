package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 int; F2 int64 }
type GenS4 struct { F0 int32; F1 uint16; F2 bool }
type GenS5 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS6 struct { F0 int }
func genFn0(p0 bool) uint64 {
    var v7 int32 = 685
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := []int{fi8}
        fmt.Println(p0, p0, p0)
        _ = fi8
        _ = v9
    }
    p0 = (791 < 259)
    v10 := GenS1{F0: ((251 > 2147483647) && p0), F1: 32767}
    if true {
        v11 := GenS1{F0: (!false), F1: 1000}
        if p0 {
            var v12 bool = true
            _ = v12
        }
        _ = v11
    }
    _ = p0
    _ = v10
    _ = v7
    return 1
}

func main() {
    var v1 int = -641
    v2 := make(map[int]int)
    v3 := []uint32{65535, 2147483648, 1}
    v4 := []int{-9223372036854775808, (v1 >> ((482) & 63)), v1}
    var v5 int8 = 1
    var v6 string = ("\n\t" + "foo bar")
    if false {
        v1 = v1
        v13 := GenS1{F0: false, F1: 790}
        _ = v13
    }
    if (1000 >= 0) {
        var v14 bool = (18 < -352)
        for fi15 := 0; fi15 < 19; fi15++ {
            switch "hello" {
            case "a":
                v6 = v6
            case "foo bar":
                if len(v4) > 0 {
                    v4[0] = v1
                }
            case "":
                v16 := GenS2{F0: 328}
                _ = v16
            default:
                v6 = v6
            }
            _ = fi15
        }
        var v17 string = ("\n\t" + ("a" + "\n\t"))
        _ = v14
        _ = v17
    }
    v1 = v1
    var v18 int = v1
    if true {
        switch 1 {
        case 128:
            _ = len(v2)
        case 2147483647:
            v19 := GenS2{F0: 42}
            _ = v19
        case 1:
            for fi20 := 0; fi20 < 17; fi20++ {
                v6 = "foo bar"
                _ = fi20
            }
        default:
            v5 = (v5 + 0)
        }
    }
    var v21 int64 = 9223372036854775807
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 uint = 905
        v5 = (v5 + -128)
        for fi24 := 0; fi24 < 20; fi24++ {
            v1 = 1
            v23 += (v23 >> (((^(v23))) & 63))
            v25 := GenS3{F0: (v18 >> ((v23) & 63)), F1: v18, F2: v21}
            _ = fi24
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    v26 := GenS4{F0: -2147483648, F1: 65535, F2: (((1.5 == 1.5) || true) && ("a" >= v6))}
    v27 := GenS3{F0: (v18 / (((-(v1))) | 1)), F1: v18, F2: 382}
    if false {
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 int = fi28
            if (true && false) {
                for fi30 := 0; fi30 < 16; fi30++ {
                    fmt.Println(v21)
                    v3 = append(v3, 1)
                    _ = fi30
                }
                fmt.Println(v5, v21, v6)
                var v31 uint64 = 435
                _ = v31
            } else {
                for fi32 := 0; fi32 < 16; fi32++ {
                    v33 := GenS5{F0: 890, F1: 4294967295, F2: 128}
                    v29 = 610
                    _ = fi32
                    _ = v33
                }
            }
            _ = fi28
            _ = v29
        }
        v34 := GenS3{F0: (v1 & (v1 * v18)), F1: (v1 | (v1 | (v1 * v18))), F2: (^(v21))}
        v35 := GenS6{F0: -1}
        _ = v34
        _ = v35
    } else {
        v6 = "\n\t"
        var v36 uint32 = 1
        _ = v36
    }
    var v37 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
