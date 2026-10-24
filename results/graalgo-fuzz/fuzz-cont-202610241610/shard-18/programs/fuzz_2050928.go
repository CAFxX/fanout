package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 string; F1 uint64 }

func main() {
    var v1 uint8 = 0
    var v2 int64 = 1
    var v3 int64 = (v2 / (((v2 / ((-238) | 1))) | 1))
    var v4 uint = 1
    v5 := []int{9223372036854775807, 128, 9223372036854775807}
    switch "hello" {
    case "a":
        var v6 float64 = 3.14159
        _ = v6
    case "hello":
        var v7 int32 = 397
        _ = v7
    case "":
        v8 := GenS1{F0: 2147483647, F1: v3}
        _ = v8
    }
    var v9 uint32 = 2
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    var v10 int64 = (v3 << (((v4 - (v4 & v4))) & 63))
    var v11 int64 = (v2 % ((v10) | 1))
    _ = len(v5)
    v4 = v4
    var v12 string = (("hello" + "\n\t") + "\n\t")
    switch 128 {
    case 127:
        v2 = (-(v11))
    case -9223372036854775808:
        for fi13 := 0; fi13 < 16; fi13++ {
            for fi14 := 0; fi14 < 7; fi14++ {
                v3 -= (v11 >> ((v4) & 63))
                switch fi13 {
                case -543:
                    v15 := GenS1{F0: 2147483647, F1: v3}
                    _ = v15
                case 32767:
                    v2 += (v11 * 0)
                case 2147483647:
                    var v16 float64 = 1.5
                    _ = v16
                }
                _ = fi14
            }
            var v17 uint64 = 0
            switch 42 {
            case -358:
                for fi18 := 0; fi18 < 1; fi18++ {
                    v19 := GenS1{F0: 42, F1: 504}
                    var v20 uint = (v4 % ((0) | 1))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case 32767:
                v4 = (v4 / (((v4 << ((uint(v9)) & 63))) | 1))
            case -627:
                var v21 uint32 = (v9 % (((^(v9))) | 1))
                _ = v21
            }
            _ = fi13
            _ = v17
        }
    default:
        v10 = 42
    }
    v11 = 271
    var v22 float64 = 3.14159
    var v23 int16 = 1
    switch 32767 {
    case -1:
        switch 0 {
        case 127:
            if true {
                v24 := GenS1{F0: 42, F1: int64(v1)}
                var v25 uint16 = 1
                _ = v24
                _ = v25
            }
        case 4294967295:
            fmt.Println(v1)
        case 191:
            var v26 uint8 = (v1 & (v1 << ((466) & 63)))
            _ = v26
        default:
            for fi27 := 0; fi27 < 5; fi27++ {
                v1 += 0
                _ = len(v5)
                _ = fi27
            }
        }
    case 32767:
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := GenS2{F0: -32768, F1: (^(fi28))}
            _ = fi28
            _ = v29
        }
    case -403:
        switch -9223372036854775808 {
        case 9223372036854775807:
            v12 = v12
        case 42:
            for fi30 := 0; fi30 < 7; fi30++ {
                v31 := GenS2{F0: v23, F1: (fi30 ^ fi30)}
                v3 = (-(v2))
                if (((v12 + v12) + v12) >= (v12 + v12)) {
                    v32 := GenS3{F0: "foo bar", F1: 0}
                    var v33 float64 = v22
                    _ = v32
                    _ = v33
                } else {
                    v22 = ((float64(((1.5 * (((((1e10 + 1.5) / -2.25) / (1.5 * 3.14159)) / -2.25) + 1e10)) - 3.14159)) - v22) + v22)
                    var v34 bool = ((0.0 == v22) || ((-(v2)) >= (v10 % (((-(v11))) | 1))))
                    _ = v34
                }
                _ = fi30
                _ = v31
            }
        default:
            var v35 int64 = v11
            _ = v35
        }
    default:
        var v36 int = -9223372036854775808
        _ = v36
    }
    v37 := GenS1{F0: 43, F1: v2}
    v38 := GenS3{F0: v12, F1: uint64(v4)}
    v23 -= 138
    v39 := GenS3{F0: v12, F1: 1}
    v12 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v9
}
