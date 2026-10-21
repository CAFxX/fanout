package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 string }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 8; fi5++ {
        if true {
            var v6 int64 = 9223372036854775807
            v6 = v6
            _ = v6
        } else {
            var v7 int = (fi5 + fi5)
            v8 := GenS1{F0: 0, F1: ((2 >= 42) || (true && true)), F2: "\n\t"}
            v9 := GenS2{F0: -1, F1: -1}
            _ = v7
            _ = v8
            _ = v9
        }
        if (-9223372036854775808 != 9223372036854775807) {
            var v10 uint = 634
            var v11 int64 = -480
            if (((0.1 * (0.1 / float32(1.5))) * 1e10) <= 1e10) {
                fmt.Println(fi5, fi5)
                var v12 string = (("foo bar" + "\n\t") + ("" + "\n\t"))
                _ = v12
            }
            _ = v10
            _ = v11
        }
        fmt.Println(fi5)
        _ = fi5
    }
    var v13 uint64 = 1
    switch 127 {
    case -1:
        for fi14 := 0; fi14 < 19; fi14++ {
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 int16 = 1
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    case -536:
        v13 = 546
    case -2147483648:
        if true {
            if false {
                for fi17 := 0; fi17 < 2; fi17++ {
                    v18 := GenS2{F0: 524, F1: 1000}
                    v13 = (v13 | (v13 % (((v13 << ((177) & 63))) | 1)))
                    _ = fi17
                    _ = v18
                }
                v13 -= (v13 / (((^(v13))) | 1))
            }
            v13 -= (^(v13))
            switch "foo bar" {
            case "foo bar":
                v13 += (v13 / ((4294967295) | 1))
            }
        }
    default:
        v19 := []int{9223372036854775807, 32768}
        _ = v19
    }
    _ = v13
    return 1000
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    var v3 int16 = 42
    v4 := []string{("hello" + "a")}
    if (true || false) {
        v3 = (v3 << ((42) & 63))
        switch 42 {
        case 42:
            v20 := GenS2{F0: (v3 << ((42) & 63)), F1: 1000}
            _ = v20
        }
    }
    v3 -= (v3 - genFn0())
    v21 := GenS2{F0: v3, F1: v3}
    var v22 int = -9223372036854775808
    v23 := GenS1{F0: 1, F1: false, F2: ("foo bar" + (("a" + ("foo bar" + "foo bar")) + "\n\t"))}
    v24 := GenS3{F0: ((v3 / ((genFn0()) | 1)) != 400)}
    v25 := GenS3{F0: (2 >= 41)}
    switch v22 {
    case 128:
        fmt.Println(v3, v3)
    case -448:
        switch (v22 / ((9223372036854775807) | 1)) {
        case 672:
            fmt.Println(v22, v22)
        default:
            var v26 float64 = 1.5
            _ = v26
        }
    }
    v22 = (-(v22))
    v22 = v22
    switch v22 {
    case 2147483647:
        v22 += v22
    }
    v27 := GenS4{F0: 0}
    fmt.Println(v22)
    fmt.Println(v22, v22)
    for fi28 := 0; fi28 < 10; fi28++ {
        v22 = fi28
        fmt.Println(v22, fi28)
        _ = fi28
    }
    fmt.Println(v3, v22)
    v3 = v3
    var v29 float32 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v29
    _ = v3
    _ = v4
}
