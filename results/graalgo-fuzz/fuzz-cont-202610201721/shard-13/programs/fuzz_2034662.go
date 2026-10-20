package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
func genFn0() int32 {
    var v6 int32 = -363
    var v7 uint64 = 303
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int32 = v6
        var v10 uint64 = 42
        v11 := GenS1{F0: 583, F1: fi8, F2: (v7 + v7)}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v6 += v6
    _ = v6
    _ = v7
    return 821
}

func main() {
    v1 := []int{0, -422, -140, -9223372036854775808}
    var v2 uint32 = 2147483648
    var v3 uint8 = 1
    var v4 int64 = 42
    var v5 int32 = 1
    v12 := GenS1{F0: (v2 ^ (v2 >> ((4294967295) & 63))), F1: 599, F2: 18446744073709551615}
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 string = ("foo bar" + "")
        _ = fi13
        _ = v14
    }
    switch "hello" {
    case "foo bar":
        switch ("\n\t" + "foo bar") {
        case "a":
            var v15 uint = 944
            _ = v15
        default:
            v1 = append(v1, 9223372036854775807)
        }
    case "a":
        v3 = (v3 >> ((4294967295) & 63))
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        _ = len(v1)
        v17 := GenS2{F0: 18446744073709551615, F1: (v5 << ((84) & 63))}
        _ = fi16
        _ = v17
    }
    switch 32767 {
    case 920:
        fmt.Println(v2)
    case -2147483648:
        v18 := GenS3{F0: -38}
        _ = v18
    case 127:
        v5 -= v5
    default:
        switch "foo bar" {
        case "foo bar":
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int = fi19
                switch 128 {
                case 4294967295:
                    for fi21 := 0; fi21 < 19; fi21++ {
                        if (false && true) {
                            fmt.Println(v2, fi21)
                        } else {
                            switch fi19 {
                            case 32768:
                                var v22 int = (^(fi19))
                                _ = v22
                            case 127:
                                switch "\n\t" {
                                case "":
                                    v23 := GenS4{F0: (v20 | 42)}
                                    _ = v23
                                case "a":
                                    if len(v1) > 0 {
                                        v1[0] = 188
                                    }
                                }
                            case -1:
                                if false {
                                    var v24 uint16 = 698
                                    if true {
                                        v5 = v5
                                        _ = len(v1)
                                    } else {
                                        v2 -= v2
                                        var v25 int16 = int16(v5)
                                        if (!true) {
                                            v20 -= 127
                                            v24 += 32767
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        } else {
                                            var v26 int32 = int32(fi21)
                                            var v27 float64 = 0.0
                                            v26 -= -2147483648
                                            v5 = (v5 - genFn0())
                                            v2 = 790
                                            _ = v26
                                            _ = v27
                                        }
                                        _ = v25
                                    }
                                    _ = v24
                                }
                            }
                            var v28 uint32 = v2
                            var v29 string = ("" + "a")
                            _ = v28
                            _ = v29
                        }
                        if true {
                            if false {
                                var v30 uint16 = 272
                                var v31 float64 = (0.0 * -2.25)
                                _ = v30
                                _ = v31
                            }
                            v20 -= (fi21 >> ((0) & 63))
                            fmt.Println(v4)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v20 = -9223372036854775808
                        v5 = v5
                        _ = fi21
                    }
                case -288:
                    if false {
                        var v32 uint = 4294967295
                        _ = v32
                    } else {
                        switch "\n\t" {
                        case "hello":
                            var v33 uint32 = 65535
                            _ = v33
                        default:
                            var v34 uint = 65535
                            _ = v34
                        }
                    }
                default:
                    v35 := GenS2{F0: 18446744073709551615, F1: (-(v5))}
                    _ = v35
                }
                v1 = append(v1, v20)
                _ = fi19
                _ = v20
            }
        default:
            fmt.Println(v5)
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v36 int16 = 1
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
