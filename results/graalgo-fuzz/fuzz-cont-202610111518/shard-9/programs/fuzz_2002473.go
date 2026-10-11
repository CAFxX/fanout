package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint; F2 uint32 }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    var v4 string = ""
    fmt.Println(v1, v1)
    v5 := GenS1{F0: 127}
    switch v4 {
    case "a":
        var v6 int32 = -2147483648
        _ = v6
    case "hello":
        for fi7 := 0; fi7 < 5; fi7++ {
            if true {
                var v8 int16 = 42
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 uint32 = 0
                    _ = fi9
                    _ = v10
                }
                fmt.Println(v1)
                _ = v8
            }
            v11 := []int{}
            _ = fi7
            _ = v11
        }
    case "foo bar":
        for fi12 := 0; fi12 < 16; fi12++ {
            if (1 > 4294967295) {
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 uint16 = 0
                    if false {
                        v15 := GenS1{F0: -1}
                        _ = v15
                    }
                    _ = fi13
                    _ = v14
                }
                var v16 int32 = 1
                v16 = v16
                _ = v16
            } else {
                v17 := GenS1{F0: 0}
                v18 := GenS2{F0: -128, F1: 18446744073709551615, F2: 4294967295}
                _ = v17
                _ = v18
            }
            v19 := GenS2{F0: 100, F1: 0, F2: 2}
            _ = fi12
            _ = v19
        }
    }
    v20 := GenS1{F0: 0}
    var v21 float32 = (v1 - v1)
    if true {
        for fi22 := 0; fi22 < 2; fi22++ {
            switch -2147483648 {
            case -2147483648:
                v23 := GenS3{F0: 945, F1: 565, F2: 4294967295}
                _ = v23
            case 447:
                fmt.Println(v4, v1, v1)
            default:
                v24 := []int{9223372036854775807}
                _ = v24
            }
            v1 = 0.0
            for fi25 := 0; fi25 < 6; fi25++ {
                fmt.Println(v4)
                var v26 int8 = -98
                _ = fi25
                _ = v26
            }
            _ = fi22
        }
        switch -1 {
        case 4294967295:
            var v27 float64 = 0.0
            _ = v27
        case -772:
            var v28 int64 = 291
            _ = v28
        default:
            v4 = "\n\t"
        }
    }
    fmt.Println(v1)
    v21 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
