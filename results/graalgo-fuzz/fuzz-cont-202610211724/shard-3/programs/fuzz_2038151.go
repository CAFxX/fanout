package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint; F1 int16; F2 uint64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -827
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 int32 = 717
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: (468 >= 42), F1: (!false), F2: 2147483648}
        _ = fi6
        _ = v7
    }
    v2 += v2
    v8 := []int{32768, 0, 32768}
    fmt.Println(v2)
    for fi9 := 0; fi9 < 13; fi9++ {
        switch "\n\t" {
        case "a":
            delete(v1, (fi9 - 914))
        case "":
            fmt.Println(v2, fi9)
        }
        v10 := GenS2{F0: 201, F1: 620, F2: 255}
        v8 = append(v8, (fi9 * 193))
        _ = fi9
        _ = v10
    }
    if ((!(579 != 1)) && false) {
        v2 = (^(v2))
        var v11 uint32 = 670
        _ = v11
    } else {
        var v12 int32 = 863
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            fmt.Println(fi13)
            var v15 float64 = 0.0
            if false {
                v15 -= (v15 + v15)
            } else {
                var v16 int64 = 9223372036854775807
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    switch -602 {
    case -9223372036854775808:
        if (458 < -9223372036854775808) {
            v2 -= (v2 & v2)
            for fi17 := 0; fi17 < 19; fi17++ {
                v2 = v2
                for fi18 := 0; fi18 < 15; fi18++ {
                    var v19 int32 = 0
                    _ = fi18
                    _ = v19
                }
                var v20 int64 = 9223372036854775807
                _ = fi17
                _ = v20
            }
        } else {
            if true {
                var v21 int16 = 884
                v21 = (v21 | (v21 << ((2) & 63)))
                v22 := GenS3{F0: "foo bar"}
                _ = v21
                _ = v22
            } else {
                for fi23 := 0; fi23 < 4; fi23++ {
                    _ = len(v8)
                    if (1.5 >= 1e10) {
                        if (!false) {
                            _ = len(v8)
                        } else {
                            if true {
                                switch -468 {
                                case -2147483648:
                                    _ = len(v8)
                                case 33:
                                    v24 := GenS3{F0: "foo bar"}
                                    _ = v24
                                default:
                                    v25 := GenS4{F0: -1}
                                    _ = v25
                                }
                            } else {
                                fmt.Println(v2, fi23)
                                var v26 string = "hello"
                                fmt.Println(v26)
                                _ = v26
                            }
                            _ = len(v8)
                        }
                        switch "foo bar" {
                        case "foo bar":
                            switch (fi23 ^ int(v2)) {
                            case 2147483647:
                                fmt.Println(fi23, fi23)
                            case 591:
                                fmt.Println(fi23, fi23)
                            }
                        case "hello":
                            for fi27 := 0; fi27 < 17; fi27++ {
                                v28 := v3[0]
                                if len(v8) > 0 {
                                    v8[0] = (fi27 ^ (fi23 >> ((65535) & 63)))
                                }
                                _ = fi27
                                _ = v28
                            }
                        default:
                            switch (^(fi23)) {
                            case 508:
                                for fi29 := 0; fi29 < 20; fi29++ {
                                    fmt.Println(fi23, fi23)
                                    _ = fi29
                                }
                            case -9:
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            default:
                                v30 := GenS1{F0: true, F1: true, F2: 0}
                                _ = v30
                            }
                        }
                        v2 += v2
                    } else {
                        if (1 == 2) {
                            v31 := GenS3{F0: "hello"}
                            _ = v31
                        }
                    }
                    _ = fi23
                }
            }
            _ = len(v8)
        }
    case 2147483647:
        switch int(v2) {
        case 128:
            v32 := GenS5{F0: 4294967295, F1: 1}
            _ = v32
        case 32767:
            if false {
                var v33 uint16 = 32767
                v34 := GenS6{F0: 171, F1: 1, F2: 2}
                var v35 string = "a"
                _ = v33
                _ = v34
                _ = v35
            } else {
                if true {
                    var v36 uint = 18446744073709551615
                    v36 -= v36
                    _ = v36
                }
                switch 2147483647 {
                case 2147483647:
                    v2 = v2
                case 9223372036854775807:
                    var v37 int8 = 5
                    _ = v37
                case 0:
                    var v38 float64 = (((1.5 + 1.5) * 3.14159) * 1.5)
                    _ = v38
                default:
                    delete(v3, 2)
                }
            }
        default:
            v2 += -1
        }
    }
    delete(v3, 189)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
