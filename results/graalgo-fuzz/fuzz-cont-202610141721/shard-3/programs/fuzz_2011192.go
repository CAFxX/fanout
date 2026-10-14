package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint8; F2 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int32; F1 bool; F2 int8 }
type GenS7 struct { F0 int16; F1 bool }
func genFn0() uint16 {
    switch 32768 {
    case -2147483648:
        for fi6 := 0; fi6 < 18; fi6++ {
            switch (fi6 - 32767) {
            case 42:
                var v7 uint8 = 208
                _ = v7
            case 128:
                if ((-32768 >= 1) || (4294967295 == 2)) {
                    var v8 int16 = 1
                    v9 := GenS1{F0: (!(-605 < 0))}
                    for fi10 := 0; fi10 < 13; fi10++ {
                        var v11 uint16 = 32767
                        v11 += v11
                        var v12 float64 = (float64(0.0) + (-2.25 - 1.5))
                        v12 = (1.5 - v12)
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = v8
                    _ = v9
                }
            default:
                fmt.Println(fi6)
            }
            _ = fi6
        }
    default:
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 int = fi13
            var v15 float32 = (0.1 - 0.0)
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    var v16 int64 = 0
    var v17 float32 = float32(0.0)
    _ = v16
    _ = v17
    return 910
}

func main() {
    var v1 uint = 65535
    var v2 float64 = (((float64(1e10) + (1e10 / 1.5)) / -2.25) * ((0.1 - 1e10) / 1e10))
    var v3 float32 = 0.0
    var v4 int16 = -1
    var v5 uint16 = 1
    v1 = v1
    v18 := []int{-868}
    v19 := GenS2{F0: (v1 / ((773) | 1))}
    v1 = (^(v1))
    switch 127 {
    case -2147483648:
        var v20 uint = 42
        _ = v20
    case 32767:
        v3 += (v3 + v3)
    case 128:
        if (100 <= -61) {
            if true {
                for fi21 := 0; fi21 < 10; fi21++ {
                    v5 = genFn0()
                    _ = fi21
                }
                if true {
                    var v22 int16 = v4
                    if (278 < 42) {
                        switch 42 {
                        case -948:
                            switch "hello" {
                            case "foo bar":
                                for fi23 := 0; fi23 < 9; fi23++ {
                                    for fi24 := 0; fi24 < 8; fi24++ {
                                        v4 += (v22 * (v4 + v4))
                                        _ = fi24
                                    }
                                    _ = fi23
                                }
                            }
                        }
                        v25 := GenS3{F0: int64(v5), F1: 1, F2: (-1 == 0)}
                        _ = v25
                    } else {
                        switch 779 {
                        case 2147483647:
                            var v26 bool = (!false)
                            _ = v26
                        default:
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                        }
                        var v27 float64 = 3.14159
                        v28 := GenS1{F0: false}
                        _ = v27
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 12; fi29++ {
                        for fi30 := 0; fi30 < 11; fi30++ {
                            var v31 uint64 = 18446744073709551615
                            if true {
                                v4 = int16(v5)
                                v32 := GenS1{F0: (0 >= v1)}
                                _ = v32
                            }
                            v18 = append(v18, (-(fi30)))
                            _ = fi30
                            _ = v31
                        }
                        for fi33 := 0; fi33 < 11; fi33++ {
                            switch "a" {
                            case "foo bar":
                                var v34 uint = (v1 | (v1 % (((v1 | 1)) | 1)))
                                _ = v34
                            case "hello":
                                var v35 int = (fi33 * (fi29 + (fi29 >> (((v1 % ((v1) | 1))) & 63))))
                                _ = v35
                            default:
                                switch ("hello" + "foo bar") {
                                case "\n\t":
                                    fmt.Println(fi29, fi29, fi33)
                                case "foo bar":
                                    if true {
                                        switch "a" {
                                        case "":
                                            var v36 uint64 = 2
                                            _ = v36
                                        case "foo bar":
                                            switch ("hello" + "foo bar") {
                                            case "\n\t":
                                                if len(v18) > 0 {
                                                    v18[0] = (fi33 << ((v1) & 63))
                                                }
                                            case "foo bar":
                                                fmt.Println(v5)
                                            case "hello":
                                                var v37 uint = v1
                                                _ = v37
                                            }
                                        default:
                                            if ((true || (((v5 ^ v5) <= 83) && false)) || false) {
                                                var v38 string = "a"
                                                _ = v38
                                            }
                                        }
                                    }
                                case "a":
                                    v2 -= v2
                                }
                                v4 -= 899
                                if false {
                                    if true {
                                        if len(v18) > 0 {
                                            v18[0] = (fi29 >> (((v1 | (v1 * (v1 - v1)))) & 63))
                                        }
                                        var v39 int64 = -9223372036854775808
                                        _ = v39
                                    } else {
                                        v18 = append(v18, (fi29 << (((v1 * uint(v5))) & 63)))
                                        v40 := GenS2{F0: (^(v1))}
                                        _ = v40
                                    }
                                }
                                fmt.Println(v3)
                            }
                            v41 := GenS1{F0: false}
                            fmt.Println(v5)
                            _ = fi33
                            _ = v41
                        }
                        v42 := GenS1{F0: (!(!(v5 < 42)))}
                        _ = fi29
                        _ = v42
                    }
                    _ = v22
                } else {
                    v4 = (v4 >> (((v1 ^ v1)) & 63))
                    v1 += v1
                    fmt.Println(v1, v2, v4)
                }
                v3 = 3.14159
            } else {
                for fi43 := 0; fi43 < 20; fi43++ {
                    v2 += v2
                    v1 = (^(v1))
                    var v44 float32 = v3
                    _ = fi43
                    _ = v44
                }
                fmt.Println(v1, v4, v1)
            }
            v1 = v1
        } else {
            if (false || (v5 >= v5)) {
                _ = len(v18)
                var v45 int8 = -128
                _ = v45
            }
            var v46 float32 = (v3 - (v3 - (0.0 * (float32(3.14159) / v3))))
            _ = v46
        }
    default:
        if false {
            switch "" {
            case "a":
                v5 = v5
            case "foo bar":
                for fi47 := 0; fi47 < 7; fi47++ {
                    var v48 int8 = -1
                    _ = fi47
                    _ = v48
                }
            default:
                for fi49 := 0; fi49 < 2; fi49++ {
                    if false {
                        _ = len(v18)
                        if len(v18) > 0 {
                            v18[0] = (^(fi49))
                        }
                        v1 = v1
                    } else {
                        v2 += (v2 / (v2 * v2))
                    }
                    v2 = v2
                    _ = fi49
                }
            }
            for fi50 := 0; fi50 < 3; fi50++ {
                var v51 int8 = -1
                fmt.Println(v4)
                _ = fi50
                _ = v51
            }
            fmt.Println(v4, v2, v2)
        }
    }
    for fi52 := 0; fi52 < 9; fi52++ {
        var v53 bool = true
        var v54 int16 = (v4 | v4)
        _ = fi52
        _ = v53
        _ = v54
    }
    v55 := GenS4{F0: 2147483647}
    fmt.Println(v1, v5)
    v56 := GenS3{F0: -635, F1: 0, F2: (!(true && (!(!(float64(((-2.25 / float32(((v2 * v2) - 0.0))) * v3)) < v2)))))}
    if ((v4 - v4) < (v4 - v4)) {
        var v57 string = "a"
        v58 := GenS5{F0: 350, F1: (18446744073709551615 > 0)}
        v59 := GenS6{F0: -1, F1: true, F2: int8(v1)}
        _ = v57
        _ = v58
        _ = v59
    } else {
        v5 += genFn0()
    }
    v60 := GenS7{F0: (v4 >> ((v1) & 63)), F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v60
}
