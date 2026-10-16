package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int32; F1 int32 }
func genFn0(p0 bool, p1 int, p2 uint16) uint64 {
    var v6 bool = ((0.0 / (float32(0.1) * float32(-2.25))) == 1e10)
    v7 := []int{p1}
    for fi8 := 0; fi8 < 10; fi8++ {
        if ((3.14159 + 1e10) >= 3.14159) {
            if (255 == 127) {
                var v9 int = 1
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v9
            } else {
                var v10 int64 = 736
                _ = v10
            }
            p0 = (18446744073709551615 > 4294967295)
        }
        _ = fi8
    }
    p1 += p1
    var v11 uint = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 float64 = (3.14159 * -2.25)
    var v2 int8 = 1
    var v3 int32 = 42
    var v4 string = ""
    var v5 int16 = 42
    var v12 uint16 = 0
    if true {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS1{F0: fi13, F1: 1, F2: (v2 / ((5) | 1))}
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 int16 = v5
            v17 := []int{fi15, (fi15 + fi15), (fi15 << ((4294967295) & 63)), 727}
            var v18 bool = true
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
    }
    switch "\n\t" {
    case "a":
        switch ((((v4 + ("\n\t" + v4)) + v4) + v4) + (v4 + "a")) {
        case "foo bar":
            if false {
                v19 := []int{}
                fmt.Println(v12, v1)
                v4 = (v4 + (v4 + v4))
                _ = v19
            }
        case "a":
            if (("\n\t" != v4) && true) {
                v20 := []int{-1, 32768}
                for fi21 := 0; fi21 < 11; fi21++ {
                    v22 := GenS1{F0: 32768, F1: 636, F2: 16}
                    v5 -= v5
                    _ = fi21
                    _ = v22
                }
                _ = v20
            }
        case "hello":
            v4 = "hello"
        default:
            v12 = 942
        }
    default:
        var v23 float32 = (((0.1 / (1.5 / (1.5 - float32(v1)))) - -2.25) / 1e10)
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 1; fi25++ {
            v26 := []int{(-(fi24)), fi25, fi25}
            var v27 uint = 42
            var v28 string = ""
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 10; fi29++ {
            fmt.Println(v5)
            _ = fi29
        }
        v30 := []int{}
        _ = fi24
        _ = v30
    }
    v3 = (v3 + (^(v3)))
    var v31 bool = (!false)
    for fi32 := 0; fi32 < 12; fi32++ {
        v33 := GenS2{F0: 184}
        var v34 int = fi32
        v35 := GenS3{F0: 1000, F1: (^(v5))}
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    v36 := GenS2{F0: 882}
    for fi37 := 0; fi37 < 9; fi37++ {
        fmt.Println(v31, v5)
        v38 := GenS4{F0: -2147483648, F1: 2147483647}
        if (v1 < v1) {
            v39 := GenS4{F0: (v3 + v3), F1: -2147483648}
            fmt.Println(v5)
            _ = v39
        }
        _ = fi37
        _ = v38
    }
    v40 := []int{-9223372036854775808, 9223372036854775807, 128, 32768}
    v41 := GenS1{F0: 32767, F1: 2, F2: (v2 << ((846) & 63))}
    var v42 bool = v31
    fmt.Println(v5, v5)
    var v43 int8 = 100
    var v44 int = 128
    v45 := GenS2{F0: 65535}
    v43 -= v2
    if len(v40) > 0 {
        v40[0] = (v44 - v44)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
