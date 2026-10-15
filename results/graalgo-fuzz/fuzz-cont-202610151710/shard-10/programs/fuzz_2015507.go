package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int32; F1 bool; F2 string }
type GenS3 struct { F0 int32; F1 uint; F2 string }
func genFn0(p0 uint8) uint {
    if true {
        v4 := []int{849, 42}
        _ = v4
    }
    v5 := GenS1{F0: ("hello" + ""), F1: 42}
    if (-77 != 1000) {
        v6 := []int{127}
        v7 := GenS1{F0: "hello", F1: -67}
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v5
    return 110
}
func genFn1(p0 string) uint8 {
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 string = p0
        _ = fi8
        _ = v9
    }
    if true {
        v10 := []int{424}
        v11 := GenS2{F0: 433, F1: (!((p0 < (p0 + (p0 + p0))) && false)), F2: "hello"}
        var v12 int32 = -705
        _ = v10
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 20; fi13++ {
            for fi14 := 0; fi14 < 19; fi14++ {
                switch fi14 {
                case -53:
                    p0 = ""
                }
                _ = fi14
            }
            for fi15 := 0; fi15 < 4; fi15++ {
                if true {
                    switch p0 {
                    case "\n\t":
                        fmt.Println(fi15, p0)
                    }
                }
                v16 := GenS1{F0: p0, F1: -128}
                var v17 uint = genFn0(128)
                fmt.Println(fi15, fi13)
                if (1 <= 32767) {
                    v18 := GenS3{F0: 1, F1: v17, F2: p0}
                    v19 := GenS2{F0: 316, F1: true, F2: ("foo bar" + "\n\t")}
                    _ = v18
                    _ = v19
                } else {
                    v20 := GenS3{F0: 42, F1: 0, F2: p0}
                    if false {
                        if (1 >= 984) {
                            p0 = ("" + "a")
                        }
                        var v21 uint64 = 1
                        _ = v21
                    } else {
                        p0 = p0
                    }
                    v22 := GenS3{F0: -2147483648, F1: (v17 ^ (v17 - (v17 ^ v17))), F2: (p0 + p0)}
                    v17 += genFn0(56)
                    switch fi13 {
                    case 2147483647:
                        switch 4294967295 {
                        case -9223372036854775808:
                            v17 -= v17
                        case -336:
                            var v23 uint32 = 2147483647
                            _ = v23
                        case 397:
                            var v24 float64 = 1e10
                            _ = v24
                        }
                    case 0:
                        p0 = p0
                    case 42:
                        p0 = (p0 + "")
                    }
                    _ = v20
                    _ = v22
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            p0 = (p0 + ("hello" + p0))
            _ = fi13
        }
        switch -1 {
        case 4294967295:
            var v25 uint32 = 2
            _ = v25
        case 128:
            p0 = "hello"
        case -642:
            switch -1 {
            case 356:
                var v26 bool = false
                _ = v26
            }
        }
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        p0 = (p0 + "")
        _ = fi27
    }
    _ = p0
    return 255
}

func main() {
    var v1 int = 0
    v2 := []int{v1, (v1 ^ -61), (v1 ^ v1)}
    v3 := make(map[string]bool)
    if false {
        v28 := GenS3{F0: 1, F1: 1, F2: "\n\t"}
        _ = v28
    } else {
        for fi29 := 0; fi29 < 6; fi29++ {
            v1 = (v1 % ((-31) | 1))
            v1 = (fi29 % ((fi29) | 1))
            _ = fi29
        }
    }
    v1 -= (v1 - 174)
    for fi30 := 0; fi30 < 3; fi30++ {
        v1 += -291
        _ = fi30
    }
    var v31 int = v1
    fmt.Println(v31)
    fmt.Println(v1, v1)
    _ = len(v2)
    v1 += v1
    v3["foo bar"] = false
    switch (v1 >> ((0) & 63)) {
    case 127:
        v1 += (^(v1))
    case -930:
        var v32 uint64 = 2
        _ = v32
    case -608:
        v31 = 42
    default:
        fmt.Println(v1, v31)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
}
