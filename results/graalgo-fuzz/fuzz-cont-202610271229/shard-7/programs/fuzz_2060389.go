package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int16; F2 int64 }

func main() {
    v1 := []uint32{57, 4294967295, 2147483647, 124}
    v2 := []int{32767, 42, 1}
    var v3 float64 = 0.0
    v4 := []uint32{2147483647, 2}
    var v5 int32 = 42
    var v6 int16 = 32767
    var v7 int32 = (v5 - (v5 % (((v5 | (^(v5)))) | 1)))
    _ = len(v4)
    var v8 uint32 = 2147483647
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: 42}
        _ = fi9
        _ = v10
    }
    var v11 uint = 1
    var v12 uint32 = v8
    var v13 float64 = v3
    switch -199 {
    case 2147483647:
        fmt.Println(v3, v8)
    case 0:
        if ((int16(v11) >= v6) && ((32767 != 0) || (false && true))) {
            if true {
                var v14 uint = 949
                if true {
                    v14 += (v11 & 42)
                    v15 := GenS1{F0: -25}
                    for fi16 := 0; fi16 < 10; fi16++ {
                        fmt.Println(v8, v7, v3)
                        v11 -= 65535
                        if false {
                            v2 = append(v2, (fi16 * fi16))
                        } else {
                            if true {
                                v14 = (v14 - 898)
                                v7 += 517
                            }
                            v17 := GenS1{F0: 127}
                            _ = v17
                        }
                        _ = fi16
                    }
                    _ = v15
                }
                _ = v14
            }
        } else {
            if false {
                for fi18 := 0; fi18 < 2; fi18++ {
                    v19 := GenS2{F0: 100}
                    if true {
                        v6 -= v6
                        switch (fi18 % (((fi18 << ((v11) & 63))) | 1)) {
                        case 9223372036854775807:
                            var v20 int = int(v11)
                            _ = v20
                        case 2147483647:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 603:
                            for fi21 := 0; fi21 < 20; fi21++ {
                                fmt.Println(v6, fi21, v8)
                                _ = fi21
                            }
                        default:
                            v5 += v5
                        }
                        if true {
                            _ = len(v2)
                            fmt.Println(v5, v8, v6)
                        } else {
                            var v22 int = (fi18 / ((fi18) | 1))
                            if (4294967295 != 1) {
                                if false {
                                    for fi23 := 0; fi23 < 5; fi23++ {
                                        if ((v8 >> ((v11) & 63)) != v8) {
                                            v3 = ((((v3 - 1e10) * v13) - v13) + v3)
                                            v6 += -32768
                                        } else {
                                            v6 += (v6 + v6)
                                            v24 := GenS1{F0: 56}
                                            var v25 bool = ((v11 >> ((18446744073709551615) & 63)) != 356)
                                            _ = v24
                                            _ = v25
                                        }
                                        switch (^(fi23)) {
                                        case 767:
                                            v8 -= (v8 - 42)
                                        case 4294967295:
                                            var v26 int64 = -1
                                            v1 = append(v1, (^(v8)))
                                            var v27 int8 = -1
                                            _ = v26
                                            _ = v27
                                        case -2147483648:
                                            v13 = (1e10 - float64(1e10))
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = v8
                                            }
                                        }
                                        v8 = 818
                                        _ = fi23
                                    }
                                    switch "hello" {
                                    case "a":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case "\n\t":
                                        for fi28 := 0; fi28 < 20; fi28++ {
                                            v3 += v3
                                            _ = fi28
                                        }
                                    default:
                                        v22 = fi18
                                    }
                                    var v29 bool = true
                                    _ = v29
                                } else {
                                    switch fi18 {
                                    case 1:
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v30 int64 = -9223372036854775808
                                            _ = v30
                                        default:
                                            var v31 uint = (v11 * (v11 & (v11 / (((v11 >> (((v11 ^ (v11 ^ v11))) & 63))) | 1))))
                                            _ = v31
                                        }
                                    default:
                                        var v32 uint32 = (^(v12))
                                        _ = v32
                                    }
                                    v8 = uint32(v11)
                                    if (!((v5 < (^(v5))) && false)) {
                                        for fi33 := 0; fi33 < 14; fi33++ {
                                            if (false || (!((v11 % ((v11) | 1)) > (v11 % ((255) | 1))))) {
                                                var v34 int16 = v6
                                                var v35 string = "\n\t"
                                                var v36 uint8 = 128
                                                v37 := GenS3{F0: true}
                                                _ = v34
                                                _ = v35
                                                _ = v36
                                                _ = v37
                                            }
                                            _ = fi33
                                        }
                                        v5 -= 1
                                        for fi38 := 0; fi38 < 20; fi38++ {
                                            var v39 uint8 = 20
                                            v8 = v12
                                            _ = fi38
                                            _ = v39
                                        }
                                    } else {
                                        v40 := GenS2{F0: -1}
                                        _ = v40
                                    }
                                }
                                _ = len(v1)
                            }
                            _ = v22
                        }
                    }
                    _ = fi18
                    _ = v19
                }
            } else {
                fmt.Println(v13, v3, v3)
            }
            v13 += (-2.25 - -2.25)
            var v41 uint16 = 0
            _ = v41
        }
    case 1:
        switch "foo bar" {
        case "":
            var v42 bool = true
            _ = v42
        case "hello":
            var v43 uint32 = v12
            _ = v43
        default:
            switch -775 {
            case -937:
                v5 = (v7 % (((v5 ^ 1)) | 1))
            }
        }
    }
    var v44 float32 = (-2.25 - 3.14159)
    if true {
        v45 := GenS4{F0: 42, F1: 0, F2: 9223372036854775807}
        _ = len(v1)
        v7 = (^(v7))
        _ = v45
    }
    for fi46 := 0; fi46 < 5; fi46++ {
        v44 = 0.1
        fmt.Println(v11, v11)
        _ = fi46
    }
    var v47 float32 = ((0.1 + v44) * (((1.5 / (((((v44 * v44) + 0.0) / 0.1) / (v44 / 1.5)) / ((v44 - v44) / v44))) / (1.5 + 0.0)) + (v44 / 3.14159)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
