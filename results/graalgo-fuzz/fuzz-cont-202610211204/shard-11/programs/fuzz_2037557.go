package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint8 }
func genFn0(p0 uint32, p1 int64) int64 {
    switch "a" {
    case "hello":
        fmt.Println(p1)
    case "\n\t":
        var v5 uint16 = 993
        _ = v5
    default:
        v6 := GenS1{F0: 127, F1: 127, F2: 19}
        _ = v6
    }
    v7 := GenS1{F0: -700, F1: 42, F2: 127}
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := []int{(fi8 % (((fi8 - fi8)) | 1))}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    if (true || (false && (797 < 2))) {
        var v10 uint16 = 32767
        v10 += (^(v10))
        v11 := GenS1{F0: 0, F1: 255, F2: 2}
        _ = v10
        _ = v11
    }
    var v12 uint16 = 42
    v13 := GenS2{F0: uint(v12), F1: 100}
    var v14 bool = true
    var v15 int32 = 2147483647
    fmt.Println(v14, v14)
    if v14 {
        var v16 float32 = 0.0
        _ = v16
    }
    switch int(v15) {
    case -273:
        var v17 int16 = 1
        _ = v17
    case 32767:
        switch "a" {
        case "\n\t":
            if (0 >= 42) {
                v18 := GenS2{F0: 2, F1: 14}
                v12 -= v12
                switch "hello" {
                case "":
                    if (v14 && ((0 != 0) && (-128 < 46))) {
                        v15 += (-(v15))
                    }
                case "a":
                    if (0 != 107) {
                        v19 := GenS3{F0: 383, F1: 255, F2: 2}
                        for fi20 := 0; fi20 < 16; fi20++ {
                            if true {
                                for fi21 := 0; fi21 < 20; fi21++ {
                                    var v22 int64 = -226
                                    _ = fi21
                                    _ = v22
                                }
                                var v23 string = "a"
                                _ = v23
                            }
                            v14 = (764 == 0)
                            _ = fi20
                        }
                        if (823 <= 0) {
                            v12 -= v12
                            _ = len(v3)
                            v24 := []int{798}
                            _ = v24
                        }
                        _ = v19
                    }
                }
                _ = v18
            }
        case "foo bar":
            var v25 int64 = -1
            _ = v25
        }
    case 128:
        var v26 float64 = 1e10
        _ = v26
    }
    if (492 == -180) {
        fmt.Println(v12, v15, v12)
        if false {
            _ = len(v4)
            if (2 != 127) {
                v27 := GenS2{F0: 42, F1: 106}
                v12 = (v12 >> ((2) & 63))
                v12 += v12
                _ = v27
            }
        }
    }
    if (v15 != (v15 & 0)) {
        var v28 int16 = -1
        switch (("hello" + "a") + "hello") {
        case "hello":
            var v29 uint8 = 42
            _ = v29
        case "\n\t":
            for fi30 := 0; fi30 < 8; fi30++ {
                for fi31 := 0; fi31 < 10; fi31++ {
                    if v14 {
                        var v32 int64 = 957
                        fmt.Println(fi30)
                        fmt.Println(v14, v12, v14)
                        _ = v32
                    } else {
                        v33 := v4[fi30]
                        fmt.Println(v12)
                        v15 -= (v15 << ((42) & 63))
                        switch ("" + "") {
                        case "a":
                            v34 := []int{(fi31 >> ((65535) & 63)), (fi30 + (^(fi30))), (fi31 ^ fi31), fi30}
                            _ = v34
                        case "":
                            v35 := GenS1{F0: fi30, F1: 42, F2: 42}
                            _ = v35
                        case "foo bar":
                            if v33 {
                                var v36 bool = (1 > 472)
                                var v37 int = 4294967295
                                _ = v36
                                _ = v37
                            }
                        }
                        v12 = 0
                        _ = v33
                    }
                    v38 := v4[(fi30 | 9223372036854775807)]
                    _ = fi31
                    _ = v38
                }
                _ = fi30
            }
        default:
            if (!(!(!v14))) {
                var v39 int = -9223372036854775808
                _ = v39
            }
        }
        _ = v28
    }
    var v40 uint8 = 0
    var v41 uint16 = (^(v12))
    v42 := GenS2{F0: 59, F1: -128}
    v43 := GenS3{F0: uint32(v12), F1: (^(v41)), F2: 2}
    v14 = (false && v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
