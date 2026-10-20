package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int32; F1 string; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 0
    var v3 int32 = 307
    var v4 float64 = 0.1
    var v5 float32 = (float32(1e10) + 1.5)
    v6 := GenS1{F0: v2, F1: 18446744073709551615, F2: 18446744073709551615}
    v7 := GenS2{F0: v3, F1: ((("" + ("a" + "foo bar")) + ("\n\t" + "a")) + "\n\t"), F2: 1}
    v2 += v2
    switch 30 {
    case 2147483647:
        if (1 > 4294967295) {
            v8 := GenS3{F0: (true || (127 >= (v2 % ((v2) | 1))))}
            if (0.0 == float64((v5 + 0.0))) {
                v9 := GenS4{F0: (v2 ^ 204), F1: true}
                if (-9223372036854775808 != 42) {
                    var v10 int16 = -32768
                    if false {
                        v3 -= 42
                        v2 += (v2 % ((v2) | 1))
                    }
                    _ = v10
                } else {
                    fmt.Println(v3)
                    v11 := GenS4{F0: 255, F1: (!true)}
                    switch 9223372036854775807 {
                    case -6:
                        v4 = (-2.25 * (-2.25 - v4))
                    case 32768:
                        if (100 > -1) {
                            v2 = (v2 + v2)
                            delete(v1, uint32(v2))
                            if (-843 < 169) {
                                _ = len(v1)
                            }
                        }
                    case 48:
                        switch "hello" {
                        case "foo bar":
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v2)
                            case "hello":
                                v12 := GenS1{F0: (v2 >> ((0) & 63)), F1: 161, F2: 18446744073709551615}
                                _ = v12
                            case "foo bar":
                                for fi13 := 0; fi13 < 13; fi13++ {
                                    var v14 uint16 = 1
                                    _ = fi13
                                    _ = v14
                                }
                            }
                        }
                    }
                    _ = v11
                }
                switch 42 {
                case 85:
                    if (true && (false && (0 > 893))) {
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            var v15 uint64 = 390
                            _ = v15
                        case "hello":
                            if true {
                                switch "a" {
                                case "":
                                    var v16 float32 = ((0.0 + -2.25) - 0.1)
                                    _ = v16
                                case "a":
                                    v17 := GenS3{F0: ((v2 & v2) >= v2)}
                                    _ = v17
                                default:
                                    if false {
                                        v3 = (v3 ^ 398)
                                        for fi18 := 0; fi18 < 19; fi18++ {
                                            v5 += v5
                                            v19 := []int{fi18, (fi18 % (((fi18 ^ fi18)) | 1)), (fi18 / (((-(fi18))) | 1)), 32768}
                                            var v20 uint16 = 65535
                                            _ = fi18
                                            _ = v19
                                            _ = v20
                                        }
                                        fmt.Println(v4, v5)
                                    }
                                }
                                v21 := []int{4294967295}
                                _ = v21
                            } else {
                                switch -775 {
                                case -2147483648:
                                    v5 = -2.25
                                case 32768:
                                    v22 := GenS5{F0: 42}
                                    _ = v22
                                }
                                v23 := GenS4{F0: v2, F1: ((v4 * (v4 / -2.25)) < (3.14159 + (float64(1.5) + v4)))}
                                for fi24 := 0; fi24 < 16; fi24++ {
                                    var v25 uint16 = 940
                                    _ = fi24
                                    _ = v25
                                }
                                _ = v23
                            }
                        case "a":
                            for fi26 := 0; fi26 < 2; fi26++ {
                                var v27 float64 = 1.5
                                _ = fi26
                                _ = v27
                            }
                        default:
                            var v28 bool = (!(v2 <= v2))
                            _ = v28
                        }
                    } else {
                        v2 = (v2 << ((18446744073709551615) & 63))
                    }
                }
                _ = v9
            }
            v4 = v4
            _ = v8
        } else {
            v29 := []int{-100, 41, int(v3)}
            v2 = (v2 / ((11) | 1))
            _ = v29
        }
    case 0:
        v2 -= 130
    default:
        v2 = v2
    }
    v2 = (v2 + 161)
    for fi30 := 0; fi30 < 15; fi30++ {
        _ = len(v1)
        _ = fi30
    }
    if true {
        var v31 string = ("" + "a")
        var v32 int8 = 100
        _ = v31
        _ = v32
    } else {
        var v33 int8 = -71
        switch 32768 {
        case 329:
            var v34 float64 = float64(0.1)
            _ = v34
        case 127:
            for fi35 := 0; fi35 < 12; fi35++ {
                v36 := []int{(fi35 % ((fi35) | 1))}
                _ = fi35
                _ = v36
            }
        }
        _ = v33
    }
    if false {
        for fi37 := 0; fi37 < 2; fi37++ {
            v3 = v3
            _ = fi37
        }
        var v38 uint16 = 65535
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
