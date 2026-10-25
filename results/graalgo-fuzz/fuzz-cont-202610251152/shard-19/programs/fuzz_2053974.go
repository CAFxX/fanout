package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 uint16 = 42
    var v4 int16 = 0
    v5 := []int{353}
    var v6 float32 = 0.0
    _ = len(v5)
    for fi7 := 0; fi7 < 16; fi7++ {
        v4 -= int16(v3)
        _ = fi7
    }
    switch 32767 {
    case 128:
        for fi8 := 0; fi8 < 4; fi8++ {
            switch fi8 {
            case -50:
                var v9 uint8 = 255
                _ = v9
            case 32767:
                v4 -= v4
            case 42:
                for fi10 := 0; fi10 < 17; fi10++ {
                    if len(v5) > 0 {
                        v5[0] = fi8
                    }
                    if (false || (42 < -1)) {
                        var v11 uint = 2
                        var v12 int16 = 1000
                        _ = v11
                        _ = v12
                    } else {
                        v3 += v3
                        if true {
                            switch 128 {
                            case 127:
                                v6 = 1.5
                            default:
                                if ((true && (!true)) || (!(2147483647 < fi10))) {
                                    v13 := GenS1{F0: ("" + "hello"), F1: fi8, F2: 65535}
                                    var v14 uint32 = 2
                                    var v15 int = fi10
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                }
                            }
                            v16 := GenS1{F0: ("" + "foo bar"), F1: fi8, F2: 276}
                            if (false && (false || true)) {
                                if false {
                                    if (0 < (^(fi8))) {
                                        var v17 uint = 65535
                                        var v18 string = "a"
                                        _ = v17
                                        _ = v18
                                    }
                                    v19 := v1[fi8]
                                    _ = v19
                                }
                            } else {
                                fmt.Println(v3)
                                delete(v1, 4294967295)
                                v20 := GenS1{F0: "", F1: (fi8 & 42), F2: uint(v3)}
                                _ = v20
                            }
                            if true {
                                switch ("hello" + "") {
                                case "hello":
                                    v21 := v1[(-(fi10))]
                                    _ = v21
                                }
                            } else {
                                v22 := GenS1{F0: ("hello" + "a"), F1: -9223372036854775808, F2: 1}
                                _ = v22
                            }
                            v23 := v2[(fi8 / (((fi10 << ((623) & 63))) | 1))]
                            _ = v16
                            _ = v23
                        } else {
                            if false {
                                v4 += 1000
                                fmt.Println(v4)
                            }
                            var v24 uint64 = 557
                            if false {
                                fmt.Println(v6, v24)
                                var v25 int8 = 127
                                _ = v25
                            }
                            _ = v24
                        }
                        v26 := GenS2{F0: 1, F1: 42}
                        switch 128 {
                        case 0:
                            v4 = v4
                        case 599:
                            var v27 uint64 = 419
                            _ = len(v2)
                            _ = v27
                        }
                        _ = v26
                    }
                    switch int(v4) {
                    case 972:
                        v4 = v4
                    case 42:
                        var v28 int8 = 69
                        _ = v28
                    }
                    _ = fi10
                }
            default:
                v4 -= (v4 & (v4 | (v4 * v4)))
            }
            v4 = (v4 & v4)
            v2[(fi8 | (^(fi8)))] = true
            _ = fi8
        }
    case 1:
        switch "foo bar" {
        case "":
            var v29 string = (("" + "\n\t") + "\n\t")
            _ = v29
        case "hello":
            for fi30 := 0; fi30 < 15; fi30++ {
                v6 += (float32(0.1) / 1.0)
                _ = fi30
            }
        case "\n\t":
            v6 += v6
        default:
            v3 += v3
        }
    case -575:
        v6 -= v6
    default:
        fmt.Println(v6, v6, v6)
    }
    if true {
        for fi31 := 0; fi31 < 13; fi31++ {
            fmt.Println(v3, v3, v4)
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 uint = 255
                v33 += v33
                switch (^(fi31)) {
                case -9223372036854775808:
                    v6 = (v6 + v6)
                }
                _ = fi32
                _ = v33
            }
            var v34 int64 = 9223372036854775807
            _ = fi31
            _ = v34
        }
        v6 = 0.1
        v35 := GenS3{F0: 128, F1: ((2 > 42) || true)}
        _ = v35
    } else {
        if false {
            for fi36 := 0; fi36 < 20; fi36++ {
                v37 := GenS2{F0: 18446744073709551615, F1: 2147483647}
                _ = fi36
                _ = v37
            }
            var v38 int32 = 42
            v3 += v3
            _ = v38
        } else {
            _ = len(v5)
            v39 := GenS2{F0: 18446744073709551615, F1: -577}
            _ = v39
        }
    }
    v40 := GenS4{F0: ("\n\t" + (("\n\t" + ("a" + "\n\t")) + "\n\t"))}
    fmt.Println(v4, v3, v3)
    for fi41 := 0; fi41 < 11; fi41++ {
        var v42 int8 = int8(v4)
        var v43 int8 = v42
        _ = fi41
        _ = v42
        _ = v43
    }
    var v44 float32 = (v6 - 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v44
    _ = v5
    _ = v6
}
