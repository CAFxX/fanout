package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint16; F1 int; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int8 }
func genFn0(p0 uint, p1 int8) int32 {
    switch "foo bar" {
    case "a":
        switch (("foo bar" + "a") + ("" + ("" + ("" + "a")))) {
        case "":
            fmt.Println(p0, p0)
        }
    case "hello":
        var v6 int64 = 1
        _ = v6
    }
    p1 = p1
    if (!true) {
        p1 -= p1
        switch "hello" {
        case "":
            var v7 float32 = 1.5
            _ = v7
        case "\n\t":
            p0 += 588
        default:
            var v8 int8 = p1
            _ = v8
        }
    } else {
        if (!false) {
            if (2 == 32767) {
                v9 := GenS1{F0: 1, F1: 2}
                _ = v9
            }
        }
        v10 := []int{444}
        if true {
            if len(v10) > 0 {
                v10[0] = -2147483648
            }
        } else {
            v11 := GenS2{F0: 42, F1: 419, F2: 690}
            p1 += p1
            _ = v11
        }
        _ = v10
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint32, p1 uint16) int64 {
    var v12 uint16 = p1
    switch "" {
    case "a":
        switch "hello" {
        case "\n\t":
            fmt.Println(p0, v12)
        case "a":
            fmt.Println(p1)
        default:
            v13 := []int{-790, 384, -961, 751}
            _ = v13
        }
    case "foo bar":
        v14 := GenS1{F0: int16(v12), F1: 255}
        _ = v14
    }
    var v15 int = 127
    _ = p0
    _ = p1
    _ = v12
    _ = v15
    return 0
}

func main() {
    var v1 float32 = (-2.25 - (1e10 + 0.0))
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    var v4 int64 = -1
    var v5 uint8 = 128
    var v16 uint64 = 42
    var v17 uint32 = 480
    for fi18 := 0; fi18 < 10; fi18++ {
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 bool = (815 == 564)
            _ = fi19
            _ = v20
        }
        var v21 int8 = -128
        _ = fi18
        _ = v21
    }
    switch 32767 {
    case 32767:
        if true {
            var v22 int8 = 127
            _ = v22
        }
    case -332:
        for fi23 := 0; fi23 < 17; fi23++ {
            if false {
                fmt.Println(fi23, v16, fi23)
            } else {
                for fi24 := 0; fi24 < 14; fi24++ {
                    if ("\n\t" == "\n\t") {
                        fmt.Println(v17)
                    } else {
                        v25 := []int{(fi24 / (((fi24 >> ((42) & 63))) | 1)), (fi23 * fi23), fi23}
                        v4 -= 9223372036854775807
                        _ = v25
                    }
                    _ = fi24
                }
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                v27 := GenS3{F0: 2, F1: v17, F2: 42}
                fmt.Println(fi23, v17, fi26)
                _ = fi26
                _ = v27
            }
            _ = fi23
        }
    case 127:
        v4 -= genFn1(42, 32767)
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        var v29 uint = 2
        if (!false) {
            var v30 uint64 = (v16 | v16)
            v31 := GenS2{F0: 0, F1: fi28, F2: 836}
            _ = v30
            _ = v31
        }
        if false {
            v32 := GenS1{F0: -32768, F1: v29}
            _ = v32
        } else {
            fmt.Println(v1, v4)
            var v33 int16 = 1000
            _ = v33
        }
        _ = fi28
        _ = v29
    }
    v34 := []int{2147483647, 2147483647, 32767, -907}
    if (v17 <= 2) {
        fmt.Println(v17)
    }
    v17 = (v17 % (((v17 | v17)) | 1))
    v4 = v4
    v35 := GenS2{F0: 1, F1: -2147483648, F2: (v16 % ((v16) | 1))}
    fmt.Println(v4, v1, v1)
    v16 = (^(v16))
    for fi36 := 0; fi36 < 2; fi36++ {
        if ((v5 * (v5 & v5)) <= 0) {
            v37 := GenS1{F0: 42, F1: 2}
            var v38 float64 = (0.0 / 0.1)
            fmt.Println(v16, v38, v4)
            _ = v37
            _ = v38
        } else {
            var v39 uint8 = v5
            v40 := GenS1{F0: 1000, F1: 65535}
            _ = v39
            _ = v40
        }
        _ = fi36
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        if ((0.0 * (0.1 / -2.25)) != 1e10) {
            var v42 int8 = 42
            _ = v42
        }
        fmt.Println(v4, v16)
        _ = fi41
    }
    var v43 int8 = 42
    v44 := GenS2{F0: 42, F1: 32767, F2: v16}
    v45 := GenS1{F0: 1, F1: 786}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
