package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int64 }
type GenS2 struct { F0 string; F1 int; F2 int64 }
func genFn0(p0 uint64, p1 int16, p2 uint) uint16 {
    v7 := GenS1{F0: 816, F1: p0, F2: 1}
    if ((302 >= 7) || false) {
        p0 = (p0 - 18446744073709551615)
        if false {
            p1 = 107
        }
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := []int{(fi8 - (fi9 << (((^(p2))) & 63))), (fi8 / ((32767) | 1))}
            v11 := GenS1{F0: -1, F1: p0, F2: 451}
            p0 = uint64(p2)
            if (!true) {
                var v12 int64 = 1
                var v13 bool = (fi8 <= (fi9 % ((fi9) | 1)))
                _ = v12
                _ = v13
            } else {
                p1 -= 1
                if len(v10) > 0 {
                    v10[0] = (fi8 << ((p2) & 63))
                }
                v14 := GenS2{F0: ("hello" + (((("hello" + "a") + ("" + "foo bar")) + "\n\t") + "")), F1: fi9, F2: -306}
                _ = v14
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        v15 := []int{fi8, fi8, (fi8 ^ (fi8 & 1))}
        _ = fi8
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 81
}
func genFn1(p0 int32, p1 int64, p2 uint32) int8 {
    var v16 uint64 = 14
    for fi17 := 0; fi17 < 2; fi17++ {
        switch "" {
        case "a":
            p2 = 752
        case "\n\t":
            for fi18 := 0; fi18 < 11; fi18++ {
                v16 -= (v16 ^ 18446744073709551615)
                switch fi18 {
                case -38:
                    p0 = (-(p0))
                default:
                    var v19 uint8 = 1
                    _ = v19
                }
                v20 := GenS2{F0: "\n\t", F1: fi17, F2: p1}
                _ = fi18
                _ = v20
            }
        case "hello":
            fmt.Println(fi17, fi17, fi17)
        }
        _ = fi17
    }
    v21 := GenS2{F0: "hello", F1: -1, F2: (p1 / (((p1 % (((p1 ^ p1)) | 1))) | 1))}
    var v22 uint64 = (^(v16))
    v23 := []int{4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v21
    _ = v22
    _ = v23
    return -89
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 1
    var v3 uint16 = v2
    v4 := []int{-631}
    v5 := []string{"hello", "a", ("" + ((("hello" + "") + ("hello" + "\n\t")) + (("" + ("\n\t" + "\n\t")) + "a"))), ("\n\t" + "hello")}
    var v6 string = "\n\t"
    var v24 int = 2147483647
    v3 = v2
    if (!(v2 < (v2 + genFn0(uint64(v2), 1000, 0)))) {
        var v25 int8 = -1
        var v26 int8 = 79
        v3 = v3
        _ = v25
        _ = v26
    }
    var v27 float64 = 0.1
    switch v24 {
    case 32767:
        var v28 int8 = 0
        _ = v28
    default:
        _ = len(v1)
    }
    var v29 uint16 = 65535
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v30 uint = 2
        _ = v30
    }
    if true {
        v6 = v6
        switch v24 {
        case 0:
            var v31 int16 = 1000
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        if ("foo bar" != "foo bar") {
            _ = len(v1)
        }
        fmt.Println(fi32)
        fmt.Println(v6)
        _ = fi32
    }
    v1[448] = 1
    if true {
        v29 -= v29
        switch (v6 + (("" + "foo bar") + "foo bar")) {
        case "hello":
            v33 := v1[v24]
            _ = v33
        }
    }
    _ = len(v1)
    var v34 int16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
