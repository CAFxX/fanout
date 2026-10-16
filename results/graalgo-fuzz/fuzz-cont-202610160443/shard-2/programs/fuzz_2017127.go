package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 bool; F1 bool; F2 uint64 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int; F1 string; F2 uint64 }
func genFn0(p0 uint) uint {
    fmt.Println(p0)
    switch "" {
    case "a":
        var v6 int16 = 0
        _ = v6
    case "\n\t":
        if true {
            var v7 int32 = -2147483648
            if true {
                p0 = (p0 / ((p0) | 1))
            }
            switch 799 {
            case 42:
                var v8 int8 = 81
                _ = v8
            default:
                fmt.Println(v7, v7)
            }
            _ = v7
        }
    case "foo bar":
        var v9 uint32 = 2147483648
        _ = v9
    }
    p0 += p0
    switch 1 {
    case -331:
        var v10 uint16 = 0
        _ = v10
    }
    if false {
        if false {
            for fi11 := 0; fi11 < 5; fi11++ {
                if (25 == 224) {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        v13 := []int{-1, fi11}
                        p0 = p0
                        fmt.Println(p0)
                        _ = fi12
                        _ = v13
                    }
                    p0 = p0
                    p0 = 42
                }
                _ = fi11
            }
            var v14 int16 = 1000
            _ = v14
        } else {
            p0 += (p0 * (^(p0)))
            var v15 int64 = 332
            for fi16 := 0; fi16 < 14; fi16++ {
                v15 = (v15 << ((p0) & 63))
                if (42 > 842) {
                    v17 := []int{(fi16 / ((-2147483648) | 1)), (-(fi16)), fi16, -2147483648}
                    _ = v17
                } else {
                    var v18 int16 = -1
                    v19 := []int{fi16, (fi16 + (fi16 % ((fi16) | 1)))}
                    v18 = v18
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
            _ = v15
        }
    } else {
        v20 := GenS1{F0: 74, F1: 2147483647}
        if (1000 > 42) {
            v21 := GenS2{F0: (0 < 65535), F1: true, F2: 950}
            if (-2.25 <= float64(0.0)) {
                p0 += (p0 ^ 0)
            }
            switch -549 {
            case 128:
                p0 = (p0 ^ (p0 + (^(p0))))
            default:
                if false {
                    var v22 string = ""
                    v23 := GenS3{F0: -2147483648, F1: true}
                    _ = v22
                    _ = v23
                } else {
                    p0 -= p0
                    var v24 string = "\n\t"
                    _ = v24
                }
            }
            _ = v21
        }
        _ = v20
    }
    _ = p0
    return (^(p0))
}
func genFn1(p0 bool, p1 int32, p2 int16) uint8 {
    v25 := []int{42}
    v26 := GenS3{F0: -213, F1: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    return 0
}

func main() {
    var v1 int16 = 0
    v2 := make(map[int]bool)
    var v3 uint8 = 255
    var v4 int8 = -13
    var v5 string = "hello"
    var v27 uint16 = 920
    fmt.Println(v3)
    var v28 uint8 = (^(v3))
    v29 := []int{}
    v3 += v3
    for fi30 := 0; fi30 < 16; fi30++ {
        v31 := GenS4{F0: fi30, F1: v5, F2: 42}
        if len(v29) > 0 {
            v29[0] = -359
        }
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        var v33 bool = (!(!((v3 >> ((255) & 63)) < 0)))
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 3; fi34++ {
        v35 := GenS2{F0: false, F1: (true && (!(v28 != (v28 - 2)))), F2: 0}
        _ = fi34
        _ = v35
    }
    v4 -= v4
    for fi36 := 0; fi36 < 8; fi36++ {
        delete(v2, int(v1))
        var v37 string = ("foo bar" + (v5 + "hello"))
        _ = fi36
        _ = v37
    }
    v1 = v1
    switch v5 {
    case "foo bar":
        for fi38 := 0; fi38 < 9; fi38++ {
            var v39 uint64 = 1
            _ = fi38
            _ = v39
        }
    }
    _ = len(v29)
    v28 = v28
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
