package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint16 }
func genFn0(p0 uint64) uint {
    var v4 int16 = 1000
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(v4, v4, p0)
        switch "\n\t" {
        case "foo bar":
            switch "" {
            case "a":
                switch (fi5 / ((fi5) | 1)) {
                case -1:
                    v6 := []int{fi5, 128}
                    _ = v6
                case 0:
                    v7 := GenS1{F0: fi5, F1: (v4 | -1)}
                    _ = v7
                case -236:
                    if true {
                        v4 = (v4 / ((42) | 1))
                        v8 := GenS2{F0: 1, F1: v4, F2: 74}
                        _ = v8
                    }
                default:
                    v4 = 32767
                }
            case "\n\t":
                v4 = (v4 & v4)
            }
        case "\n\t":
            var v9 uint16 = 42
            _ = v9
        case "":
            v10 := []int{fi5, fi5}
            _ = v10
        default:
            v11 := GenS1{F0: 32767, F1: (^(v4))}
            _ = v11
        }
        _ = fi5
    }
    p0 = (p0 ^ (p0 * (p0 << ((65535) & 63))))
    _ = p0
    _ = v4
    return 18446744073709551615
}
func genFn1() int64 {
    var v12 uint32 = 1
    var v13 uint32 = 2147483647
    _ = v12
    _ = v13
    return -1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    v3 := []int{42}
    var v14 int64 = genFn1()
    v14 = (v14 | 871)
    fmt.Println(v14)
    v14 += (-(v14))
    switch (("hello" + ((((("foo bar" + "hello") + "\n\t") + ("hello" + "\n\t")) + "\n\t") + (("hello" + "foo bar") + "a"))) + "") {
    case "\n\t":
        if (1 >= 42) {
            for fi15 := 0; fi15 < 19; fi15++ {
                switch ("" + ("\n\t" + "hello")) {
                case "foo bar":
                    var v16 float64 = 0.0
                    _ = v16
                case "a":
                    fmt.Println(fi15, v14)
                case "hello":
                    var v17 uint32 = 0
                    _ = v17
                }
                _ = fi15
            }
        }
    default:
        fmt.Println(v14)
    }
    v3 = append(v3, 9223372036854775807)
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := GenS2{F0: 2147483648, F1: 1000, F2: 962}
        _ = fi18
        _ = v19
    }
    if len(v3) > 0 {
        v3[0] = -1
    }
    if false {
        v14 = genFn1()
    } else {
        for fi20 := 0; fi20 < 9; fi20++ {
            if false {
                if len(v3) > 0 {
                    v3[0] = fi20
                }
                var v21 uint = 255
                if true {
                    var v22 uint16 = 0
                    _ = v22
                } else {
                    v21 = v21
                }
                _ = v21
            }
            _ = len(v3)
            _ = fi20
        }
        v14 -= (-(v14))
        fmt.Println(v14)
    }
    var v23 int8 = -73
    var v24 float32 = 1.5
    v24 += 1.5
    var v25 int32 = 303
    v26 := GenS2{F0: 42, F1: 1, F2: 42}
    v14 = (v14 * -9223372036854775808)
    v3 = append(v3, 0)
    var v27 float32 = v24
    v27 -= v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
