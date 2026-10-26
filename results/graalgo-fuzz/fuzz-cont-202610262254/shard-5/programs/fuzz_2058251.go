package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint; F2 int16 }
type GenS3 struct { F0 int16; F1 int }

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]string)
    var v3 int16 = -32768
    var v4 uint32 = 1
    var v5 uint = 143
    var v6 string = "hello"
    fmt.Println(v1, v5)
    v7 := []int{}
    v6 = ("\n\t" + ("foo bar" + "foo bar"))
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 bool = (18446744073709551615 > 42)
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 uint16 = 255
            var v12 uint64 = 42
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := GenS1{F0: ("a" < ((v6 + (v6 + (v6 + v6))) + (((v6 + "a") + v6) + "")))}
            var v15 uint64 = 916
            if (!(float64(1.5) != 1e10)) {
                switch fi8 {
                case -1:
                    var v16 int64 = -9223372036854775808
                    _ = v16
                }
                switch 32767 {
                case 2147483647:
                    var v17 int32 = 1
                    _ = v17
                case 42:
                    var v18 bool = false
                    v3 -= v3
                    if (-165 == 0) {
                        switch (fi8 + (fi8 - 0)) {
                        case 597:
                            switch v6 {
                            case "foo bar":
                                switch v6 {
                                case "\n\t":
                                    fmt.Println(v3, v3)
                                case "hello":
                                    v19 := GenS1{F0: true}
                                    _ = v19
                                case "a":
                                    v1 = v1
                                }
                            }
                        default:
                            fmt.Println(fi8)
                        }
                        v3 = 42
                        v3 = 0
                    }
                    _ = v18
                default:
                    if len(v7) > 0 {
                        v7[0] = fi13
                    }
                }
                var v20 float64 = 1e10
                fmt.Println(v9, v15, fi13)
                _ = v20
            } else {
                v1 -= (v1 % ((v1) | 1))
                switch (-(fi13)) {
                case -336:
                    switch v6 {
                    case "a":
                        fmt.Println(fi8, fi13)
                    }
                case 604:
                    if true {
                        v3 = (-(v3))
                    }
                case -9223372036854775808:
                    if ((^(v15)) <= v15) {
                        v4 = v4
                        _ = len(v2)
                        v21 := GenS1{F0: (!("\n\t" == v6))}
                        _ = v21
                    }
                }
                var v22 int32 = 0
                fmt.Println(fi13, v3)
                switch int(v5) {
                case 1:
                    v23 := GenS2{F0: v4, F1: 65535, F2: v3}
                    _ = v23
                case -1:
                    if false {
                        v3 = 707
                    } else {
                        var v24 float32 = float32(float64(1e10))
                        if (v9 || v9) {
                            v24 += v24
                            v25 := GenS1{F0: (fi13 == 616)}
                            _ = v25
                        } else {
                            var v26 float64 = (((1e10 - ((0.0 * (1.5 * float64(3.14159))) + 1e10)) - 3.14159) / 1.5)
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            switch v6 {
                            case "foo bar":
                                switch ((v6 + v6) + v6) {
                                case "\n\t":
                                    v22 += -2147483648
                                case "foo bar":
                                    var v27 float64 = v26
                                    _ = v27
                                default:
                                    var v28 int = fi8
                                    if v9 {
                                        var v29 float64 = (v26 / v26)
                                        fmt.Println(fi13, v1)
                                        v30 := GenS1{F0: v9}
                                        _ = v29
                                        _ = v30
                                    }
                                    v5 += v5
                                    _ = v28
                                }
                            }
                            _ = v26
                        }
                        _ = v24
                    }
                }
                _ = v22
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi8
        _ = v9
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        v5 -= 18446744073709551615
        if (2 != 128) {
            var v32 uint = 403
            var v33 bool = (216 <= 447)
            _ = v32
            _ = v33
        }
        _ = fi31
    }
    switch "foo bar" {
    case "":
        var v34 string = "hello"
        _ = v34
    }
    fmt.Println(v4)
    v5 = v5
    v35 := GenS3{F0: (-(v3)), F1: 128}
    for fi36 := 0; fi36 < 4; fi36++ {
        v5 -= (v5 ^ v5)
        if true {
            for fi37 := 0; fi37 < 6; fi37++ {
                var v38 float32 = 0.0
                if true {
                    if false {
                        _ = len(v7)
                        var v39 int32 = -2147483648
                        _ = v39
                    }
                    var v40 uint = v5
                    _ = v40
                } else {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
                v3 = v3
                _ = fi37
                _ = v38
            }
        } else {
            var v41 uint32 = v4
            _ = v41
        }
        fmt.Println(fi36)
        _ = fi36
    }
    var v42 uint64 = 0
    var v43 uint64 = v42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
