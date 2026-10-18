package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int16 }
type GenS2 struct { F0 uint; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int16; F2 int8 }

func main() {
    var v1 int32 = 342
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    if true {
        switch "" {
        case "":
            v1 = (v1 * v1)
        case "hello":
            var v4 int16 = 931
            _ = v4
        case "\n\t":
            if false {
                v5 := []int{4294967295, 351, 128}
                fmt.Println(v1)
                _ = v5
            }
        default:
            var v6 int8 = 127
            _ = v6
        }
    } else {
        switch 128 {
        case 32767:
            v7 := []int{}
            _ = v7
        }
        v8 := GenS1{F0: 2147483647, F1: (!false), F2: 42}
        _ = v8
    }
    var v9 int64 = -1
    switch -661 {
    case -1:
        if true {
            v1 += -1
            if (471 > 1) {
                v10 := GenS1{F0: 2, F1: (879 < 2), F2: -225}
                for fi11 := 0; fi11 < 19; fi11++ {
                    v12 := GenS2{F0: 1, F1: v9, F2: 42}
                    for fi13 := 0; fi13 < 6; fi13++ {
                        switch "hello" {
                        case "foo bar":
                            v14 := GenS1{F0: 740, F1: ((1 < 65535) && (false && true)), F2: -1}
                            _ = v14
                        case "\n\t":
                            v15 := GenS3{F0: 38}
                            _ = v15
                        case "a":
                            var v16 int = 128
                            _ = v16
                        default:
                            v17 := []int{(fi11 << ((4294967295) & 63)), fi13, (fi11 + (fi13 >> ((1) & 63))), 32768}
                            _ = v17
                        }
                        v1 -= 1
                        _ = fi13
                    }
                    _ = fi11
                    _ = v12
                }
                _ = v10
            }
        }
    case 999:
        var v18 float32 = -2.25
        _ = v18
    case 552:
        var v19 uint = 42
        _ = v19
    default:
        if (!true) {
            delete(v2, 4294967295)
            var v20 int32 = (-(v1))
            _ = v20
        }
    }
    v1 = v1
    v21 := GenS4{F0: (-9223372036854775808 > v9), F1: 0, F2: 127}
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := []int{}
        _ = fi22
        _ = v23
    }
    fmt.Println(v9, v9)
    if true {
        v24 := []int{-512, 1, 42, 32767}
        _ = v24
    } else {
        for fi25 := 0; fi25 < 4; fi25++ {
            fmt.Println(v1)
            _ = fi25
        }
        v9 = v9
        v1 -= v1
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v9
}
