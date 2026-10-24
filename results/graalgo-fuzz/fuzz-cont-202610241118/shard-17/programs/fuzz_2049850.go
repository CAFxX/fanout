package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int64; F1 string; F2 uint8 }

func main() {
    v1 := []uint32{}
    v2 := []int{2147483647}
    var v3 int = -252
    var v4 float64 = float64((1e10 / 3.14159))
    var v5 uint16 = 529
    var v6 uint64 = 2
    var v7 uint32 = 1
    switch (v3 | 127) {
    case 42:
        var v8 uint = 91
        _ = v8
    case 32768:
        v3 += v3
    case -1:
        switch (v3 + (-(v3))) {
        case 42:
            v9 := GenS1{F0: false, F1: 1}
            _ = v9
        case -851:
            v10 := GenS2{F0: -730}
            _ = v10
        case -1:
            var v11 uint = uint(v6)
            _ = v11
        default:
            var v12 uint64 = v6
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 19; fi13++ {
            v1 = append(v1, v7)
            _ = fi13
        }
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        switch (("" + "hello") + (("hello" + "\n\t") + "foo bar")) {
        case "hello":
            var v15 int = (^(v3))
            _ = v15
        case "":
            switch "hello" {
            case "":
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 uint64 = (v6 % (((v6 << ((65535) & 63))) | 1))
                    switch (v3 % ((int(v17)) | 1)) {
                    case 441:
                        v7 = (^(v7))
                    case 2147483647:
                        var v18 float64 = float64(1.5)
                        _ = v18
                    case 42:
                        v3 += (^(fi14))
                    default:
                        var v19 bool = true
                        _ = v19
                    }
                    v20 := GenS1{F0: (uint(v7) != 343), F1: -2147483648}
                    v3 -= 32767
                    if true {
                        _ = len(v2)
                    }
                    v21 := GenS2{F0: 581}
                    _ = fi16
                    _ = v17
                    _ = v20
                    _ = v21
                }
            }
        case "foo bar":
            var v22 bool = false
            _ = v22
        }
        _ = fi14
    }
    v5 += (v5 / ((v5) | 1))
    var v23 uint = 18446744073709551615
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v23, v23)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 += v3
    v24 := GenS3{F0: v23, F1: v3}
    var v25 uint16 = (v5 >> ((v23) & 63))
    for fi26 := 0; fi26 < 19; fi26++ {
        v27 := GenS3{F0: 236, F1: (v3 * -2147483648)}
        switch "" {
        case "a":
            switch "foo bar" {
            case "a":
                if (-203 != 42) {
                    switch "foo bar" {
                    case "hello":
                        v28 := GenS4{F0: 42, F1: (("" + "hello") + "a"), F2: 0}
                        _ = v28
                    default:
                        if true {
                            for fi29 := 0; fi29 < 7; fi29++ {
                                switch (fi26 % (((fi29 / ((fi26) | 1))) | 1)) {
                                case 32767:
                                    v7 += (v7 * uint32(v6))
                                }
                                _ = fi29
                            }
                            fmt.Println(fi26, v6, v5)
                            for fi30 := 0; fi30 < 19; fi30++ {
                                if true {
                                    v5 -= uint16(v23)
                                    fmt.Println(v6, v6)
                                    fmt.Println(v23, v7)
                                    _ = len(v1)
                                    var v31 float32 = 1.5
                                    _ = v31
                                }
                                _ = fi30
                            }
                        } else {
                            for fi32 := 0; fi32 < 11; fi32++ {
                                fmt.Println(v7, fi26)
                                var v33 bool = false
                                _ = fi32
                                _ = v33
                            }
                            var v34 float64 = (((v4 + (v4 + float64(1.5))) + float64(0.1)) / (float64(float32(0.1)) / v4))
                            _ = v34
                        }
                    }
                    v35 := GenS1{F0: true, F1: 2147483647}
                    for fi36 := 0; fi36 < 2; fi36++ {
                        fmt.Println(v7)
                        fmt.Println(fi26, v3, fi26)
                        _ = fi36
                    }
                    _ = v35
                }
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        case "hello":
            for fi37 := 0; fi37 < 8; fi37++ {
                v7 = (^(v7))
                _ = fi37
            }
        case "foo bar":
            var v38 uint32 = (v7 - v7)
            _ = v38
        default:
            v5 = uint16(v7)
        }
        _ = fi26
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
