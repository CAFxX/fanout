package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int, p1 int16, p2 uint) int32 {
    switch (p0 % (((p0 << ((42) & 63))) | 1)) {
    case -654:
        var v4 int32 = 239
        _ = v4
    case 32767:
        var v5 uint = (^(p2))
        _ = v5
    }
    p1 = (p1 >> (((p2 ^ (p2 >> ((p2) & 63)))) & 63))
    var v6 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1(p0 uint32) uint32 {
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 = (^(p0))
        p0 += (^(p0))
        _ = fi7
    }
    if true {
        v8 := []int{-941, 32768}
        _ = v8
    } else {
        var v9 uint32 = p0
        if (!false) {
            v10 := []int{-2147483648, -1}
            var v11 float64 = (0.0 - ((1.5 - ((1.5 * 0.0) - float64(0.0))) - (1e10 + (((1e10 / float64(0.1)) - 1.5) * (((-2.25 + (1.5 + 1.5)) / 1e10) - 3.14159)))))
            fmt.Println(p0, p0, v11)
            _ = v10
            _ = v11
        } else {
            switch -342 {
            case 4294967295:
                if true {
                    v12 := []int{32767, 4294967295, -819, 833}
                    fmt.Println(p0)
                    v9 -= p0
                    _ = v12
                }
            case 2147483647:
                v13 := GenS1{F0: "\n\t"}
                _ = v13
            case 0:
                v14 := []int{-114, 0, 4294967295, 236}
                _ = v14
            default:
                fmt.Println(p0)
            }
            v15 := []int{872}
            var v16 uint64 = 401
            _ = v15
            _ = v16
        }
        switch (("\n\t" + "") + "\n\t") {
        case "\n\t":
            var v17 uint32 = p0
            _ = v17
        case "foo bar":
            v18 := GenS1{F0: "hello"}
            _ = v18
        }
        _ = v9
    }
    var v19 int = 1
    switch (v19 % ((v19) | 1)) {
    case 2147483647:
        switch "\n\t" {
        case "":
            v20 := GenS2{F0: false, F1: -2147483648, F2: 4294967295}
            _ = v20
        case "\n\t":
            var v21 uint32 = (^(p0))
            _ = v21
        default:
            v22 := GenS1{F0: "hello"}
            _ = v22
        }
    case 4294967295:
        switch 127 {
        case 788:
            var v23 bool = true
            _ = v23
        case -2147483648:
            v19 = v19
        case 9223372036854775807:
            for fi24 := 0; fi24 < 13; fi24++ {
                v25 := GenS3{F0: p0}
                _ = fi24
                _ = v25
            }
        default:
            fmt.Println(v19, p0)
        }
    }
    var v26 int8 = 1
    _ = p0
    _ = v19
    _ = v26
    return (^(p0))
}

func main() {
    v1 := []int{}
    var v2 uint8 = 128
    var v3 uint32 = 4294967295
    v1 = append(v1, 834)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        switch 2147483647 {
        case 545:
            for fi27 := 0; fi27 < 20; fi27++ {
                for fi28 := 0; fi28 < 13; fi28++ {
                    if true {
                        var v29 uint8 = 255
                        _ = v29
                    } else {
                        v1 = append(v1, 1)
                        if len(v1) > 0 {
                            v1[0] = fi28
                        }
                    }
                    v3 -= v3
                    if true {
                        var v30 int64 = 1
                        var v31 string = "hello"
                        _ = v30
                        _ = v31
                    }
                    _ = fi28
                }
                v32 := GenS4{F0: (v3 >> ((1) & 63)), F1: v3, F2: 2}
                _ = fi27
                _ = v32
            }
        case 0:
            switch -532 {
            case -308:
                v3 = (^(v3))
            case 2147483647:
                _ = len(v1)
            default:
                if true {
                    var v33 bool = ((v2 + (^(v2))) == (v2 * (v2 ^ v2)))
                    _ = v33
                } else {
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                    v2 += 186
                }
            }
        }
    }
    fmt.Println(v3, v3)
    var v34 uint32 = 0
    var v35 float32 = 1e10
    v36 := GenS5{F0: (v34 - (v3 << ((65535) & 63))), F1: 237}
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
