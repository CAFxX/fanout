package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int8; F2 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 int16 }

func main() {
    v1 := []int{57}
    v2 := make(map[uint32]bool)
    var v3 uint8 = 127
    v3 = v3
    var v4 int = 1
    var v5 string = ("a" + ("\n\t" + "foo bar"))
    switch v4 {
    case 603:
        if false {
            v6 := GenS1{F0: 42, F1: 42, F2: 42}
            var v7 int32 = 42
            _ = v6
            _ = v7
        } else {
            v8 := GenS2{F0: ""}
            _ = len(v1)
            _ = v8
        }
    case 128:
        v9 := GenS3{F0: true, F1: 0, F2: 9223372036854775807}
        _ = v9
    case -193:
        var v10 float64 = 0.0
        _ = v10
    default:
        for fi11 := 0; fi11 < 4; fi11++ {
            fmt.Println(v3, v3)
            delete(v2, 708)
            if true {
                var v12 string = v5
                _ = v12
            } else {
                if (v4 >= (fi11 + 1)) {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        v14 := GenS1{F0: 2147483648, F1: 32767, F2: 1}
                        _ = fi13
                        _ = v14
                    }
                    v2[734] = (float64((float32(0.1) + (0.0 * 1.5))) > float64(3.14159))
                }
                switch (fi11 ^ (fi11 << ((1) & 63))) {
                case 0:
                    v3 = (v3 * (v3 + 127))
                case 1:
                    if false {
                        for fi15 := 0; fi15 < 18; fi15++ {
                            var v16 int8 = 1
                            v5 = v5
                            if true {
                                v5 = ""
                                if false {
                                    _ = len(v1)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    var v17 uint32 = 42
                                    switch fi11 {
                                    case 0:
                                        var v18 uint16 = 32767
                                        _ = v18
                                    case -9223372036854775808:
                                        switch v4 {
                                        case -623:
                                            var v19 uint32 = uint32(v3)
                                            _ = v19
                                        case 498:
                                            var v20 float64 = ((3.14159 - 3.14159) - 0.1)
                                            _ = v20
                                        case 42:
                                            var v21 bool = (0 <= -1)
                                            _ = v21
                                        }
                                    case 315:
                                        switch fi11 {
                                        case 593:
                                            v5 = v5
                                        }
                                    default:
                                        v3 = v3
                                    }
                                    _ = v17
                                }
                                if true {
                                    switch (v4 / (((fi11 & (fi15 << ((18446744073709551615) & 63)))) | 1)) {
                                    case 9:
                                        switch v5 {
                                        case "hello":
                                            v2[42] = (-1 == 9223372036854775807)
                                        case "":
                                            switch fi15 {
                                            case -663:
                                                v2[4294967295] = (false && true)
                                            default:
                                                if false {
                                                    switch v4 {
                                                    case 1:
                                                        fmt.Println(v16)
                                                    }
                                                    v3 += v3
                                                    _ = len(v1)
                                                    switch (fi15 ^ (fi11 / ((-966) | 1))) {
                                                    case 858:
                                                        v22 := GenS4{F0: 1}
                                                        _ = v22
                                                    }
                                                    v5 = v5
                                                } else {
                                                    v23 := GenS2{F0: (v5 + v5)}
                                                    if false {
                                                        var v24 uint32 = 0
                                                        if false {
                                                            if true {
                                                                v16 = 100
                                                                v16 = 16
                                                            } else {
                                                                var v25 int32 = 1
                                                                var v26 float64 = 1.5
                                                                _ = v25
                                                                _ = v26
                                                            }
                                                            var v27 float32 = 1e10
                                                            _ = v27
                                                        }
                                                        _ = v24
                                                    }
                                                    var v28 int8 = v16
                                                    _ = v23
                                                    _ = v28
                                                }
                                            }
                                        case "\n\t":
                                            var v29 bool = false
                                            _ = v29
                                        }
                                    case 0:
                                        var v30 int32 = 2147483647
                                        _ = v30
                                    }
                                    v16 = (^(v16))
                                    v1 = append(v1, fi11)
                                }
                                delete(v2, 1)
                                v31 := GenS5{F0: 42, F1: 65535}
                                _ = v31
                            }
                            _ = fi15
                            _ = v16
                        }
                    }
                }
                v5 = v5
            }
            _ = fi11
        }
    }
    v32 := GenS1{F0: 2147483647, F1: 65535, F2: 217}
    if (65535 > 65535) {
        var v33 uint16 = 381
        _ = v33
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v34 bool = (-308 != 254)
    if (-32768 != -32768) {
        for fi35 := 0; fi35 < 8; fi35++ {
            for fi36 := 0; fi36 < 16; fi36++ {
                v4 = fi35
                _ = fi36
            }
            if v34 {
                v2[530] = v34
                if len(v1) > 0 {
                    v1[0] = -1
                }
                fmt.Println(fi35, v3)
            } else {
                v34 = ((v5 + v5) != v5)
                var v37 uint = 328
                var v38 string = "a"
                _ = v37
                _ = v38
            }
            v39 := GenS6{F0: 0}
            _ = fi35
            _ = v39
        }
    } else {
        v4 = v4
    }
    v4 = v4
    _ = len(v2)
    if ((^(v3)) != (v3 | 2)) {
        v2[0] = ((v34 || v34) || (-1 < -60))
        fmt.Println(v34, v34)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v4
    _ = v5
}
