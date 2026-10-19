package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0(p0 int, p1 int8, p2 bool) uint8 {
    var v6 bool = (p2 || true)
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: ("\n\t" + "\n\t"), F1: 221}
        fmt.Println(v6, p2, p2)
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS2{F0: 2, F1: 127}
            switch p0 {
            case -2147483648:
                var v11 uint32 = 2
                _ = v11
            case -1:
                var v12 uint32 = 274
                _ = v12
            case 127:
                v13 := GenS2{F0: 42, F1: p0}
                _ = v13
            }
            var v14 int = (fi9 | (fi9 ^ int(p1)))
            _ = fi9
            _ = v10
            _ = v14
        }
        _ = fi7
        _ = v8
    }
    var v15 int64 = 9223372036854775807
    v16 := []int{(-(p0)), (p0 / ((32768) | 1)), 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v6
    return 128
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = (1e10 + 1e10)
    var v3 int32 = 1
    var v4 uint64 = 112
    var v5 uint8 = 184
    v5 = 146
    var v17 uint64 = (v4 ^ 4294967295)
    if true {
        v17 -= (v17 / ((393) | 1))
    }
    if false {
        for fi18 := 0; fi18 < 7; fi18++ {
            v5 += (v5 % (((v5 * (v5 >> ((42) & 63)))) | 1))
            fmt.Println(v4)
            fmt.Println(fi18, v17)
            _ = fi18
        }
        var v19 int64 = 9223372036854775807
        v17 = (v17 >> ((651) & 63))
        _ = v19
    }
    switch 1 {
    case -1:
        var v20 uint = 1
        _ = v20
    default:
        fmt.Println(v2)
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        _ = len(v1)
        _ = fi21
    }
    v22 := GenS2{F0: (v17 & v4), F1: -951}
    var v23 int64 = 1
    v24 := GenS3{F0: 424, F1: 2147483647}
    for fi25 := 0; fi25 < 12; fi25++ {
        v5 -= genFn0((^(fi25)), -34, ("\n\t" == ("hello" + "\n\t")))
        _ = fi25
    }
    delete(v1, 0)
    fmt.Println(v2)
    switch "" {
    case "":
        v1[42] = "a"
    case "foo bar":
        v26 := []int{32768, -9223372036854775808, 127, -1}
        _ = v26
    case "\n\t":
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 int64 = 42
            v29 := GenS1{F0: ((("hello" + "\n\t") + "hello") + "a"), F1: (v5 << ((4294967295) & 63))}
            _ = fi27
            _ = v28
            _ = v29
        }
    default:
        v4 = v4
    }
    v30 := GenS3{F0: (v3 << ((65535) & 63)), F1: (v3 % (((v3 * (v3 ^ v3))) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
