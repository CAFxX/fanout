package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int16; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int64 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]int)
    var v3 int64 = 632
    v4 := make(map[uint32]string)
    if (v3 <= (v3 + 1)) {
        v1 -= 4294967295
    }
    var v5 int64 = v3
    var v6 uint64 = 4294967295
    for fi7 := 0; fi7 < 15; fi7++ {
        v6 -= (v6 ^ (v6 + v6))
        var v8 uint = uint(v6)
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    if false {
        v6 = (v6 >> ((459) & 63))
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: ((true && (1.5 > (0.1 * 1.5))) || (42 <= -937)), F1: v3}
            var v12 float32 = -2.25
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 bool = true
        _ = v13
    } else {
        if false {
            v6 -= 1
            var v14 uint16 = 2
            v15 := GenS2{F0: -1, F1: 188, F2: 1}
            _ = v14
            _ = v15
        }
        fmt.Println(v3, v3)
        for fi16 := 0; fi16 < 16; fi16++ {
            v2[fi16] = 1
            v3 += (-(v3))
            _ = fi16
        }
    }
    var v17 uint64 = 964
    var v18 float64 = 1.5
    var v19 int32 = -2147483648
    var v20 bool = (2 < 241)
    fmt.Println(v17)
    v9 = append(v9, -2147483648)
    v1 = v1
    if (true && (v1 == (v1 & 4294967295))) {
        var v21 int64 = v3
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := GenS2{F0: 32767, F1: 42, F2: 0}
            for fi24 := 0; fi24 < 15; fi24++ {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    var v25 uint = 4294967295
                    _ = v25
                case "a":
                    v21 = (v21 >> ((18446744073709551615) & 63))
                }
                switch int(v19) {
                case -1:
                    var v26 uint32 = v1
                    _ = v26
                case 127:
                    fmt.Println(v3)
                default:
                    switch fi24 {
                    case 0:
                        _ = len(v2)
                    case 128:
                        var v27 int = (^(fi22))
                        _ = v27
                    case 42:
                        _ = len(v9)
                    }
                }
                v28 := GenS3{F0: 226}
                v21 = 118
                var v29 int8 = -128
                v30 := GenS4{F0: fi24, F1: 1}
                delete(v2, (fi24 + (fi22 % ((9223372036854775807) | 1))))
                var v31 uint32 = (v1 << ((42) & 63))
                _ = fi24
                _ = v28
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        var v32 int8 = -1
        _ = v32
    }
    switch -2147483648 {
    case 1:
        v19 = 2147483647
    case 128:
        _ = len(v2)
    }
    fmt.Println(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
