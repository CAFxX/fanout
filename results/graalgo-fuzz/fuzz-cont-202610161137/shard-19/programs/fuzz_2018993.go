package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint64) int {
    v6 := GenS1{F0: 2147483647, F1: 1, F2: 754}
    switch "" {
    case "foo bar":
        p0 += p0
    default:
        p0 = (p0 + p0)
    }
    p0 -= (^(p0))
    var v7 int32 = 2147483647
    var v8 int8 = -1
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1() string {
    var v9 uint16 = 534
    v10 := []int{978, 127, 32767}
    var v11 int16 = -491
    if false {
        for fi12 := 0; fi12 < 2; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                _ = len(v10)
                switch "\n\t" {
                case "foo bar":
                    var v14 int = (fi12 << ((155) & 63))
                    v15 := GenS1{F0: (fi12 % ((genFn0(42)) | 1)), F1: -84, F2: (v11 ^ v11)}
                    _ = v14
                    _ = v15
                case "":
                    if ((^(v9)) <= v9) {
                        var v16 uint8 = 255
                        switch "foo bar" {
                        case "a":
                            v17 := GenS1{F0: (fi12 & fi12), F1: 77, F2: (^(v11))}
                            _ = v17
                        case "":
                            v16 = (v16 - v16)
                        default:
                            v11 = (v11 + -676)
                        }
                        _ = v16
                    } else {
                        v18 := GenS2{F0: 42, F1: (!(v11 >= (v11 ^ 42)))}
                        fmt.Println(fi12, fi13, fi13)
                        _ = v18
                    }
                }
                fmt.Println(v11, v11, v9)
                _ = fi13
            }
            _ = fi12
        }
    }
    _ = v10
    _ = v11
    _ = v9
    return ""
}

func main() {
    var v1 uint16 = 882
    var v2 int8 = -60
    var v3 uint8 = 148
    var v4 uint32 = 793
    var v5 int8 = 42
    for fi19 := 0; fi19 < 17; fi19++ {
        switch 0 {
        case 1:
            v20 := GenS3{F0: 1}
            _ = v20
        default:
            fmt.Println(v3)
        }
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 int8 = v2
            _ = fi21
            _ = v22
        }
        if (true || true) {
            switch (fi19 ^ genFn0(0)) {
            case -594:
                var v23 int16 = -1
                _ = v23
            default:
                v24 := []int{-200, genFn0(2)}
                _ = v24
            }
        } else {
            switch "\n\t" {
            case "":
                var v25 int16 = -771
                _ = v25
            case "foo bar":
                v5 = (v5 | v5)
            case "hello":
                for fi26 := 0; fi26 < 17; fi26++ {
                    switch genFn0(770) {
                    case -2147483648:
                        v5 = (v2 + v5)
                    default:
                        v3 = (v3 % ((v3) | 1))
                    }
                    if (42 > 0) {
                        fmt.Println(v1, v1)
                        if true {
                            v27 := []int{int(v2), 0, (fi19 & -9223372036854775808)}
                            _ = v27
                        }
                        var v28 uint = 255
                        _ = v28
                    }
                    _ = fi26
                }
            default:
                if (!(v2 > 100)) {
                    var v29 int64 = int64(v1)
                    v30 := GenS3{F0: (v4 & (^(v4)))}
                    _ = v29
                    _ = v30
                }
            }
            fmt.Println(v3, fi19)
        }
        _ = fi19
    }
    if true {
        fmt.Println(v5, v1)
        fmt.Println(v5, v3, v2)
    }
    v31 := GenS2{F0: (v4 / (((v4 / (((^(v4))) | 1))) | 1)), F1: (3.14159 != ((float64(1e10) / 0.1) / 1.0))}
    if ((-2.25 * 0.1) != 1e10) {
        var v32 float64 = 1e10
        v33 := GenS2{F0: (v4 - (v4 << ((4294967295) & 63))), F1: false}
        _ = v32
        _ = v33
    } else {
        v34 := GenS3{F0: (v4 << ((4294967295) & 63))}
        v4 = 564
        _ = v34
    }
    fmt.Println(v2, v1)
    fmt.Println(v4, v2, v2)
    var v35 int16 = -32768
    v4 = v4
    v36 := GenS2{F0: v4, F1: false}
    var v37 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
