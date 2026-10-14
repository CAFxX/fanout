package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int64; F1 int64; F2 uint32 }
type GenS3 struct { F0 int64; F1 int32; F2 uint32 }

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]bool)
    var v3 float64 = 0.1
    v1 -= (v1 / ((v1) | 1))
    fmt.Println(v3, v1)
    switch -70 {
    case -784:
        v4 := []int{-1}
        _ = v4
    case -265:
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := GenS1{F0: 0, F1: 91, F2: 0}
            v7 := GenS1{F0: 2147483648, F1: 4294967295, F2: 1}
            _ = fi5
            _ = v6
            _ = v7
        }
    case 2147483647:
        v8 := GenS2{F0: -879, F1: 9223372036854775807, F2: 2147483648}
        _ = v8
    }
    fmt.Println(v3)
    var v9 float32 = (1.5 / 0.1)
    if true {
        fmt.Println(v3)
        var v10 bool = (551 > 4294967295)
        v11 := []int{127, 9223372036854775807, 0}
        _ = v10
        _ = v11
    }
    switch "hello" {
    case "":
        if ((!(!false)) && (true && false)) {
            v12 := GenS2{F0: 1, F1: 261, F2: 4294967295}
            if ((-815 < 32768) || false) {
                v1 = v1
                v13 := []int{128}
                v14 := v2[("" + ("" + (("" + "") + "a")))]
                _ = v13
                _ = v14
            }
            v15 := []int{2147483647, 4294967295}
            _ = v12
            _ = v15
        } else {
            v16 := []int{42}
            v17 := GenS2{F0: -9223372036854775808, F1: 537, F2: 42}
            _ = v16
            _ = v17
        }
    case "foo bar":
        v1 -= v1
    case "\n\t":
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS3{F0: -1, F1: 2147483647, F2: 65535}
            var v20 int = fi18
            _ = fi18
            _ = v19
            _ = v20
        }
    default:
        v21 := v2[""]
        _ = v21
    }
    v22 := GenS2{F0: 9223372036854775807, F1: -9223372036854775808, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v9
}
