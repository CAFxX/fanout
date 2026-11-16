package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int64; F2 int8 }
func genFn0(p0 int32) bool {
    v7 := []int{42, 4294967295, -770, 2147483647}
    _ = len(v7)
    _ = p0
    _ = v7
    return (637 >= 1)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 503
    var v3 int32 = -1
    var v4 bool = false
    v5 := make(map[int]int)
    v6 := make(map[int]int)
    if (1 >= 255) {
        v8 := GenS1{F0: (v2 % (((v2 - (v2 | 378))) | 1)), F1: 545}
        _ = v8
    } else {
        if (false || ((!v4) && false)) {
            v9 := []int{4294967295, -470}
            for fi10 := 0; fi10 < 12; fi10++ {
                v2 += v2
                var v11 float32 = 0.1
                _ = fi10
                _ = v11
            }
            _ = v9
        } else {
            fmt.Println(v4)
        }
        fmt.Println(v4, v3, v4)
        switch ("a" + ("hello" + "foo bar")) {
        case "hello":
            switch -257 {
            case -654:
                v4 = false
            case 824:
                fmt.Println(v4, v2)
            default:
                var v12 int32 = -311
                _ = v12
            }
        case "\n\t":
            v13 := []int{}
            _ = v13
        case "":
            v14 := GenS2{F0: 42}
            _ = v14
        }
    }
    if v4 {
        v3 += (v3 | v3)
        var v15 uint16 = 96
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 int64 = 42
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    v18 := []int{-2147483648, 32767}
    v2 = 65535
    v19 := GenS2{F0: 687}
    var v20 uint32 = 726
    var v21 uint16 = (^(v2))
    v22 := GenS3{F0: v20, F1: 9223372036854775807, F2: 1}
    fmt.Println(v2, v21, v2)
    var v23 bool = ((!v4) || (!(127 >= 32767)))
    v24 := GenS1{F0: 814, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
