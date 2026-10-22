package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int8 }

func main() {
    var v1 int8 = 127
    var v2 float32 = (1e10 - 1e10)
    var v3 float32 = v2
    var v4 int32 = 42
    var v5 float32 = float32(1e10)
    var v6 float32 = v5
    fmt.Println(v3, v2, v1)
    v7 := []int{}
    v7 = append(v7, 127)
    for fi8 := 0; fi8 < 2; fi8++ {
        if len(v7) > 0 {
            v7[0] = (fi8 + (fi8 << ((18446744073709551615) & 63)))
        }
        _ = fi8
    }
    switch "a" {
    case "a":
        var v9 int = 42
        _ = v9
    case "hello":
        var v10 float64 = 3.14159
        _ = v10
    case "foo bar":
        v5 = v2
    default:
        switch "foo bar" {
        case "":
            switch "" {
            case "foo bar":
                var v11 int8 = v1
                _ = v11
            default:
                v2 = v6
            }
        default:
            v12 := GenS1{F0: v1, F1: ((0.0 + 1e10) <= 0.0)}
            _ = v12
        }
    }
    v3 -= 0.0
    fmt.Println(v6)
    if false {
        fmt.Println(v4)
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS1{F0: v1, F1: (1 < 1)}
            switch (fi13 % ((fi13) | 1)) {
            case 2147483647:
                switch "hello" {
                case "a":
                    var v15 int16 = -32768
                    _ = v15
                }
            case 32767:
                v16 := GenS2{F0: 65535}
                _ = v16
            }
            _ = fi13
            _ = v14
        }
    } else {
        if (((v2 + (v5 + v2)) >= v6) || false) {
            v1 += 127
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 int32 = -480
            v7 = append(v7, fi17)
            _ = fi17
            _ = v18
        }
        var v19 string = (("\n\t" + ("" + "\n\t")) + "\n\t")
        _ = v19
    }
    _ = len(v7)
    for fi20 := 0; fi20 < 10; fi20++ {
        v3 += 3.14159
        for fi21 := 0; fi21 < 12; fi21++ {
            v22 := GenS3{F0: v1}
            if (!true) {
                if len(v7) > 0 {
                    v7[0] = fi20
                }
            } else {
                v2 -= v3
            }
            _ = fi21
            _ = v22
        }
        switch (fi20 & (fi20 - fi20)) {
        case 529:
            if true {
                v23 := GenS4{F0: false, F1: (-(v1))}
                for fi24 := 0; fi24 < 11; fi24++ {
                    v2 = v6
                    _ = len(v7)
                    var v25 int32 = v4
                    _ = fi24
                    _ = v25
                }
                switch -9223372036854775808 {
                case 0:
                    v7 = append(v7, (fi20 + (fi20 << ((4294967295) & 63))))
                default:
                    var v26 uint64 = 873
                    _ = v26
                }
                _ = v23
            }
        case 9223372036854775807:
            switch (fi20 >> ((525) & 63)) {
            case 809:
                fmt.Println(v5, v3)
            case 874:
                if true {
                    v6 += v3
                } else {
                    for fi27 := 0; fi27 < 6; fi27++ {
                        var v28 uint32 = 65535
                        _ = fi27
                        _ = v28
                    }
                    if (!false) {
                        switch (-(fi20)) {
                        case 0:
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        case 127:
                            var v29 int64 = int64(fi20)
                            _ = v29
                        case -299:
                            for fi30 := 0; fi30 < 6; fi30++ {
                                switch fi30 {
                                case 4294967295:
                                    var v31 int64 = 0
                                    _ = v31
                                case 1:
                                    if len(v7) > 0 {
                                        v7[0] = (^(fi30))
                                    }
                                case 2147483647:
                                    fmt.Println(v5, v6, v6)
                                }
                                if len(v7) > 0 {
                                    v7[0] = (fi30 ^ (-(fi20)))
                                }
                                _ = len(v7)
                                _ = fi30
                            }
                        }
                        var v32 uint64 = 4294967295
                        _ = v32
                    }
                }
            case 9223372036854775807:
                var v33 uint8 = 33
                _ = v33
            default:
                if true {
                    v34 := GenS4{F0: (true || (true && (!true))), F1: (v1 ^ -128)}
                    switch (fi20 + (-(fi20))) {
                    case 152:
                        fmt.Println(v6, v5, v5)
                    case 2147483647:
                        var v35 uint16 = 218
                        _ = v35
                    }
                    v3 += (((((v2 + (1.5 / (((3.14159 + 1.5) * (1e10 + v3)) / -2.25))) / (((0.0 + v3) * 0.1) / ((((float32((float64(v5) + float64(v5))) - ((float32(-2.25) - float32(3.14159)) + v3)) - v5) + v3) - float32(-2.25)))) + 0.1) + float32(0.0)) * float32(0.0))
                    _ = v34
                }
            }
        case -2147483648:
            v3 = (v5 + v5)
        }
        _ = fi20
    }
    v3 += (v5 * (((v3 + (v5 / v2)) * v6) - ((v6 / ((((0.1 / 3.14159) + 1.5) + ((((float32(-2.25) / (((0.1 - 0.1) / v6) - v6)) / v3) * v5) / v3)) / (v2 - v3))) + ((-2.25 - float32(-2.25)) / v5))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
