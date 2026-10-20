package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 uint8 }
type GenS4 struct { F0 int32; F1 int16; F2 int32 }
func genFn0() uint8 {
    if (1 > -2147483648) {
        v7 := GenS1{F0: 127, F1: 2}
        var v8 uint16 = 71
        _ = v7
        _ = v8
    } else {
        v9 := []int{-126, 42, 0, -44}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        if (!(!(65535 != 2))) {
            var v10 string = ("" + (("foo bar" + "foo bar") + ""))
            _ = v10
        }
        _ = v9
    }
    var v11 int8 = 42
    switch "" {
    case "foo bar":
        v12 := []int{-9223372036854775808, 0}
        _ = v12
    }
    v13 := []int{128}
    _ = v11
    _ = v13
    return 1
}

func main() {
    var v1 uint16 = 421
    var v2 bool = (!(false || ((2 != 60) && true)))
    v3 := []uint32{2}
    var v4 uint64 = 508
    var v5 float32 = (1.5 / 1.5)
    v6 := []string{("" + ""), "hello", "hello"}
    switch 32768 {
    case -896:
        v1 = v1
    case -764:
        var v14 uint16 = v1
        _ = v14
    default:
        fmt.Println(v4, v1)
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        if (v2 && (743 > 255)) {
            var v16 uint = 2
            v3 = append(v3, 42)
            _ = v16
        }
        _ = fi15
    }
    v17 := GenS2{F0: 2}
    v1 = v1
    var v18 bool = v2
    switch "hello" {
    case "a":
        if len(v3) > 0 {
            v3[0] = 746
        }
    default:
        v5 += (v5 * v5)
    }
    switch 32768 {
    case -325:
        switch 127 {
        case -2147483648:
            v19 := GenS3{F0: 1, F1: -666, F2: 0}
            _ = v19
        case 128:
            if (!((0 == 100) && false)) {
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 uint16 = (v1 >> ((2) & 63))
                    _ = fi20
                    _ = v21
                }
            } else {
                var v22 float64 = 0.1
                v2 = v18
                _ = v22
            }
        case -836:
            var v23 float32 = float32(3.14159)
            _ = v23
        default:
            v24 := GenS1{F0: 241, F1: uint64(v1)}
            _ = v24
        }
    case 408:
        v5 = v5
    }
    if v2 {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        var v25 int16 = 0
        _ = v25
    }
    v2 = (true && (-128 == -102))
    v2 = v18
    var v26 float64 = float64(1.5)
    if (470 >= 1) {
        if v18 {
            for fi27 := 0; fi27 < 1; fi27++ {
                if v2 {
                    fmt.Println(v2, v4)
                }
                fmt.Println(v2, v18, v1)
                if (2 > 2147483647) {
                    if ((42 > 2147483647) || (42 < -9223372036854775808)) {
                        v3 = append(v3, 65535)
                        var v28 uint = 4294967295
                        if (uint(v1) != 42) {
                            var v29 int = 128
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v29 -= fi27
                            _ = v29
                        } else {
                            var v30 int8 = 0
                            if (v26 > (((((v26 + 1.5) + (-2.25 / ((1e10 + float64(1.5)) * (v26 / v26)))) * ((0.1 / v26) * v26)) / v26) - v26)) {
                                fmt.Println(fi27, v2, v5)
                                var v31 int = 32767
                                _ = v31
                            }
                            _ = v30
                        }
                        _ = v28
                    } else {
                        if len(v6) > 0 {
                            v6[0] = ""
                        }
                        if (v2 && v18) {
                            var v32 int = (fi27 - (fi27 ^ (fi27 >> ((841) & 63))))
                            if true {
                                var v33 string = ("\n\t" + "hello")
                                _ = v33
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            var v34 int8 = -128
                            _ = v32
                            _ = v34
                        } else {
                            var v35 int8 = 1
                            _ = v35
                        }
                        switch "a" {
                        case "foo bar":
                            fmt.Println(v2, fi27)
                        case "":
                            v36 := GenS2{F0: 42}
                            _ = v36
                        case "\n\t":
                            var v37 int64 = 0
                            _ = v37
                        default:
                            fmt.Println(v5, fi27)
                        }
                    }
                    var v38 float64 = v26
                    fmt.Println(fi27, v18, v1)
                    _ = v38
                } else {
                    fmt.Println(v18)
                }
                _ = fi27
            }
            for fi39 := 0; fi39 < 3; fi39++ {
                if v2 {
                    if (-32768 >= 42) {
                        v40 := GenS1{F0: 2, F1: v4}
                        _ = v40
                    } else {
                        if ((v4 <= (v4 << ((0) & 63))) && v18) {
                            for fi41 := 0; fi41 < 15; fi41++ {
                                var v42 string = "\n\t"
                                _ = fi41
                                _ = v42
                            }
                        }
                        var v43 bool = v2
                        _ = v43
                    }
                    var v44 int32 = -401
                    _ = v44
                }
                v5 = v5
                if v2 {
                    var v45 uint32 = 1
                    _ = v45
                }
                _ = fi39
            }
        }
        v2 = v2
    }
    v2 = (0 == v1)
    var v46 float32 = v5
    v46 = v46
    v47 := GenS2{F0: genFn0()}
    if ((1e10 + (float64((v46 * v5)) * 3.14159)) < v26) {
        v4 = v4
        v5 += ((v46 - (0.1 + (((((v46 - (-2.25 * ((-2.25 - v46) + float32((v26 - (v26 / v26)))))) - float32((float64((((0.0 - (1.5 - 1e10)) - v46) / (0.1 + (v5 + 1e10)))) / (((0.1 / ((((v26 - (0.0 - v26)) + ((v26 * 0.1) / 1.5)) - (v26 - v26)) * (v26 * (((v26 / 3.14159) / (((float64(v46) * v26) + ((-2.25 / float64(0.1)) * v26)) * ((((float64((v5 - v5)) - v26) / (v26 + float64(v5))) + (v26 / float64(v5))) / v26))) + v26)))) / ((float64((float32((1.5 - 3.14159)) - (((v46 + 1.5) + (v46 / v5)) - v5))) * v26) / (v26 * (v26 / v26)))) - v26)))) / ((0.0 + ((v46 + -2.25) + (v5 / v46))) / v5)) * float32(v26)) * (v5 + 0.1)))) * (v5 * -2.25))
        v48 := GenS3{F0: 65535, F1: 32767, F2: 1}
        _ = v48
    } else {
        for fi49 := 0; fi49 < 20; fi49++ {
            var v50 int = 42
            for fi51 := 0; fi51 < 7; fi51++ {
                if true {
                    v2 = v18
                    if v2 {
                        var v52 int16 = -476
                        _ = v52
                    } else {
                        v18 = ((v1 == v1) || false)
                        v5 -= (3.14159 * 0.1)
                        if v2 {
                            fmt.Println(fi51, v18)
                            var v53 string = ("hello" + "")
                            v50 = v50
                            _ = v53
                        } else {
                            var v54 string = "\n\t"
                            _ = v54
                        }
                    }
                } else {
                    switch (fi51 << ((65535) & 63)) {
                    case -501:
                        v55 := GenS2{F0: 128}
                        _ = v55
                    default:
                        if v2 {
                            var v56 uint = 18446744073709551615
                            _ = v56
                        }
                    }
                    var v57 bool = true
                    switch (fi49 / (((fi51 - 924)) | 1)) {
                    case 183:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case 32768:
                        v18 = v18
                    default:
                        switch v50 {
                        case -2147483648:
                            if true {
                                v57 = (2 < 112)
                                v58 := GenS4{F0: 1, F1: 42, F2: 1}
                                var v59 bool = false
                                _ = v58
                                _ = v59
                            }
                        }
                    }
                    _ = v57
                }
                switch "\n\t" {
                case "foo bar":
                    var v60 bool = (2147483647 >= 938)
                    _ = v60
                case "hello":
                    if v18 {
                        v5 += 1.5
                    } else {
                        var v61 uint64 = v4
                        _ = v61
                    }
                }
                _ = fi51
            }
            _ = fi49
            _ = v50
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
