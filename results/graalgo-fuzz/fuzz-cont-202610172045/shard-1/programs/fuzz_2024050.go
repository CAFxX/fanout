package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int) uint {
    v6 := []int{}
    v6 = append(v6, (p0 + (^(p0))))
    p0 -= (-(p0))
    v7 := GenS1{F0: 2, F1: 1, F2: 35}
    _ = p0
    _ = v6
    _ = v7
    return 0
}
func genFn1() uint {
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(fi8)
        v9 := GenS2{F0: 1}
        v10 := []int{2147483647}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 string = "\n\t"
    _ = v11
    return 193
}

func main() {
    var v1 uint32 = 129
    var v2 uint8 = 2
    var v3 uint32 = 1
    var v4 uint = uint(v2)
    var v5 float64 = 1e10
    v12 := GenS3{F0: v2}
    var v13 string = "foo bar"
    v2 = 0
    var v14 uint64 = 4294967295
    if ((!false) || true) {
        for fi15 := 0; fi15 < 13; fi15++ {
            v16 := GenS3{F0: v2}
            v17 := GenS2{F0: -780}
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := GenS1{F0: v14, F1: v14, F2: (v2 >> (((v4 << ((251) & 63))) & 63))}
        var v19 uint32 = (v1 / ((v3) | 1))
        _ = v18
        _ = v19
    } else {
        fmt.Println(v5, v14, v1)
    }
    if (-1 <= 1) {
        fmt.Println(v13, v14, v13)
        switch "a" {
        case "\n\t":
            switch v13 {
            case "hello":
                v14 = 407
            }
        case "foo bar":
            switch v13 {
            case "foo bar":
                fmt.Println(v13, v1, v1)
            default:
                if ((756 > 1) || false) {
                    v20 := GenS3{F0: v2}
                    var v21 uint16 = 25
                    if (v5 > 0.0) {
                        var v22 string = v13
                        fmt.Println(v1, v2)
                        v5 = v5
                        _ = v22
                    } else {
                        if false {
                            var v23 uint8 = v2
                            _ = v23
                        }
                    }
                    _ = v20
                    _ = v21
                }
            }
        case "":
            if (false || true) {
                v24 := GenS3{F0: (v2 << (((v4 % (((^(v4))) | 1))) & 63))}
                _ = v24
            }
        }
        if true {
            var v25 int16 = -1
            _ = v25
        }
    } else {
        var v26 uint32 = (v3 + 4294967295)
        for fi27 := 0; fi27 < 9; fi27++ {
            v1 = v1
            _ = fi27
        }
        _ = v26
    }
    var v28 uint16 = 42
    switch v13 {
    case "hello":
        if true {
            fmt.Println(v4)
            var v29 uint64 = (^(v14))
            _ = v29
        } else {
            if false {
                var v30 bool = (v5 >= v5)
                _ = v30
            }
        }
    case "\n\t":
        var v31 bool = ((32767 <= 128) || ((1 <= 0) || false))
        _ = v31
    default:
        v1 = (v1 & v3)
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        if ("a" > v13) {
            var v33 int64 = 1
            var v34 uint32 = v1
            if ((v3 / (((v1 % ((v1) | 1))) | 1)) > v34) {
                v2 = 128
                v35 := []int{(fi32 >> ((v4) & 63)), int(v2)}
                _ = v35
            }
            _ = v33
            _ = v34
        }
        v36 := []int{(-(fi32)), fi32, 32768, fi32}
        _ = fi32
        _ = v36
    }
    var v37 bool = false
    var v38 int8 = 52
    fmt.Println(v3, v2)
    var v39 int32 = 2147483647
    if ((v2 ^ 128) == (v2 / (((^(v2))) | 1))) {
        if (0 >= -341) {
            v40 := GenS2{F0: int32(v3)}
            for fi41 := 0; fi41 < 17; fi41++ {
                v3 = (v1 / ((uint32(v28)) | 1))
                _ = fi41
            }
            _ = v40
        }
    }
    var v42 uint = v4
    var v43 uint64 = v14
    var v44 int8 = 42
    v4 = v42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
