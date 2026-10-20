package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint32 }
type GenS2 struct { F0 int; F1 int64; F2 int8 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint; F1 string; F2 uint64 }
func genFn0(p0 int64, p1 int8) int8 {
    for fi7 := 0; fi7 < 11; fi7++ {
        if (255 != 42) {
            switch ("hello" + "\n\t") {
            case "\n\t":
                p0 = (p0 * (p0 >> ((42) & 63)))
            case "foo bar":
                p0 -= 42
            case "":
                var v8 uint = 18446744073709551615
                _ = v8
            }
            v9 := GenS1{F0: -1, F1: fi7, F2: 42}
            _ = v9
        } else {
            p1 += (p1 / (((-(p1))) | 1))
        }
        _ = fi7
    }
    switch 128 {
    case -9223372036854775808:
        switch -2147483648 {
        case -1:
            p0 = 9223372036854775807
        default:
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 float64 = 3.14159
                if (32767 < 32767) {
                    var v12 uint64 = 2
                    var v13 uint64 = 2
                    _ = v12
                    _ = v13
                }
                v11 -= (v11 / v11)
                _ = fi10
                _ = v11
            }
        }
    case -1:
        switch 494 {
        case 128:
            v14 := []int{-9223372036854775808}
            _ = v14
        case -2147483648:
            fmt.Println(p1, p0)
        default:
            fmt.Println(p0, p0)
        }
    case 843:
        var v15 string = "\n\t"
        _ = v15
    }
    v16 := GenS1{F0: p0, F1: -1, F2: 829}
    _ = p0
    _ = p1
    _ = v16
    return -37
}

func main() {
    v1 := []uint32{2147483647, 1, 65535, 1}
    v2 := make(map[string]int)
    var v3 uint32 = 2147483647
    v4 := make(map[int]int)
    v5 := make(map[uint32]string)
    var v6 int = -187
    if false {
        v6 = v6
    } else {
        switch ("" + ("foo bar" + ("a" + "foo bar"))) {
        case "a":
            var v17 uint8 = 73
            _ = v17
        case "\n\t":
            fmt.Println(v3)
        case "":
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS1{F0: 318, F1: 0, F2: v3}
                _ = fi18
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v6, v6, fi20)
            var v21 uint = 225
            _ = fi20
            _ = v21
        }
        var v22 int16 = 0
        _ = v22
    }
    var v23 bool = false
    if (42 >= 1) {
        fmt.Println(v23, v23)
        if v23 {
            v24 := GenS2{F0: (v6 + (v6 / ((-354) | 1))), F1: 990, F2: 1}
            _ = v24
        }
    }
    v25 := GenS3{F0: (v3 * (v3 / (((v3 << ((667) & 63))) | 1))), F1: 1}
    v26 := GenS1{F0: 9223372036854775807, F1: -450, F2: v3}
    v23 = v23
    var v27 int64 = 0
    v23 = (true && (v27 > v27))
    var v28 uint64 = 380
    v29 := GenS4{F0: 18446744073709551615, F1: "foo bar", F2: v28}
    _ = len(v2)
    for fi30 := 0; fi30 < 16; fi30++ {
        if (!(v23 || v23)) {
            v3 = 0
        }
        for fi31 := 0; fi31 < 8; fi31++ {
            v3 -= (v3 * (v3 & 559))
            _ = fi31
        }
        _ = fi30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
