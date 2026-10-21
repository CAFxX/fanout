package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32 }

func main() {
    v1 := []int{-397, 1, -136, 128}
    var v2 int8 = 127
    var v3 uint = 2
    v4 := make(map[uint32]int)
    var v5 int64 = 0
    var v6 uint = v3
    v5 = v5
    fmt.Println(v2, v2)
    v5 = (v5 + 900)
    fmt.Println(v2, v2)
    _ = len(v4)
    fmt.Println(v5, v5, v2)
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    v3 = v6
    var v7 string = ("a" + "a")
    if (!true) {
        _ = len(v1)
        var v8 uint64 = 1
        fmt.Println(v8, v6)
        _ = v8
    } else {
        fmt.Println(v7, v5, v3)
        _ = len(v4)
        switch 32767 {
        case 405:
            var v9 float64 = (1e10 + 1e10)
            _ = v9
        case 128:
            switch (((v7 + "hello") + ((((v7 + v7) + v7) + v7) + ("\n\t" + "\n\t"))) + "foo bar") {
            case "":
                if false {
                    for fi10 := 0; fi10 < 19; fi10++ {
                        v11 := GenS1{F0: 65535}
                        fmt.Println(v5, v2, v7)
                        var v12 uint64 = 56
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                } else {
                    v2 = v2
                    v3 += 42
                }
            case "hello":
                for fi13 := 0; fi13 < 7; fi13++ {
                    v3 = 97
                    _ = fi13
                }
            case "\n\t":
                v2 = (v2 << ((v3) & 63))
            default:
                var v14 int64 = -1
                _ = v14
            }
        case 127:
            for fi15 := 0; fi15 < 5; fi15++ {
                if len(v1) > 0 {
                    v1[0] = fi15
                }
                _ = fi15
            }
        }
    }
    if (!(!(8 <= 2147483647))) {
        var v16 uint8 = 0
        _ = v16
    }
    if true {
        _ = len(v4)
        if true {
            if (false && false) {
                v2 -= (v2 >> (((v6 * v3)) & 63))
                for fi17 := 0; fi17 < 1; fi17++ {
                    v5 = 887
                    var v18 int64 = (v5 >> (((^(v3))) & 63))
                    v19 := GenS1{F0: uint16(v6)}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            } else {
                if (v3 != 0) {
                    var v20 int32 = 0
                    var v21 int = 9223372036854775807
                    _ = v20
                    _ = v21
                } else {
                    v3 = 0
                }
            }
            if (((0.0 - 1e10) / 1.0) <= ((0.1 * (-2.25 + float32(1.5))) * (3.14159 * -2.25))) {
                var v22 int16 = 32767
                fmt.Println(v5, v6)
                v5 = int64(v2)
                _ = v22
            }
        } else {
            var v23 float64 = (0.0 * -2.25)
            _ = v23
        }
    } else {
        v24 := GenS2{F0: v5, F1: 4294967295, F2: 4294967295}
        var v25 uint64 = 1
        v25 = v25
        _ = v24
        _ = v25
    }
    if len(v1) > 0 {
        v1[0] = -888
    }
    if ((((1.5 + ((-2.25 * 1.5) / 1.0)) / 1e10) / 3.14159) != (0.0 - 3.14159)) {
        for fi26 := 0; fi26 < 13; fi26++ {
            switch fi26 {
            case 32767:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 1:
                v4[65535] = fi26
            case 4294967295:
                if true {
                    for fi27 := 0; fi27 < 13; fi27++ {
                        _ = len(v1)
                        _ = fi27
                    }
                }
            default:
                v28 := GenS3{F0: 138}
                _ = v28
            }
            for fi29 := 0; fi29 < 13; fi29++ {
                if (2 == 2147483647) {
                    var v30 bool = (false || false)
                    if v30 {
                        var v31 uint = (v6 - (v6 % (((v3 & 4294967295)) | 1)))
                        v31 -= (v3 % ((v31) | 1))
                        _ = v31
                    }
                    v32 := GenS3{F0: 138}
                    _ = v30
                    _ = v32
                } else {
                    var v33 int = (-(fi26))
                    var v34 int32 = -334
                    switch 42 {
                    case 128:
                        v35 := GenS4{F0: (v2 * v2), F1: 866}
                        _ = v35
                    default:
                        switch "" {
                        case "a":
                            if len(v1) > 0 {
                                v1[0] = v33
                            }
                            v3 += (v6 - (v3 * v6))
                            v2 += v2
                            if (75 > 1) {
                                v7 = ""
                                v6 -= v3
                                _ = len(v4)
                            } else {
                                var v36 uint16 = 193
                                var v37 uint64 = 1
                                _ = v36
                                _ = v37
                            }
                        case "hello":
                            v33 -= v33
                        default:
                            v5 += int64(fi26)
                        }
                    }
                    v38 := GenS1{F0: 255}
                    v34 = v34
                    _ = v33
                    _ = v34
                    _ = v38
                }
                var v39 float64 = float64(((0.0 * 0.1) * (0.1 - (0.0 * 0.1))))
                _ = fi29
                _ = v39
            }
            v3 = 18446744073709551615
            _ = fi26
        }
    }
    var v40 uint16 = 34
    for fi41 := 0; fi41 < 1; fi41++ {
        v42 := GenS5{F0: -1}
        if false {
            for fi43 := 0; fi43 < 12; fi43++ {
                fmt.Println(fi43, v5)
                if true {
                    v44 := GenS3{F0: 0}
                    v45 := GenS6{F0: 4294967295}
                    v3 += v6
                    var v46 string = v7
                    _ = v44
                    _ = v45
                    _ = v46
                } else {
                    switch (fi41 % ((fi41) | 1)) {
                    case 42:
                        v5 += int64(v40)
                    default:
                        v3 -= (v6 << (((v6 - v6)) & 63))
                    }
                }
                _ = fi43
            }
            switch "foo bar" {
            case "foo bar":
                switch 127 {
                case -551:
                    switch fi41 {
                    case 9223372036854775807:
                        switch (("" + v7) + (v7 + "hello")) {
                        case "hello":
                            for fi47 := 0; fi47 < 6; fi47++ {
                                fmt.Println(fi47)
                                _ = fi47
                            }
                        case "a":
                            _ = len(v4)
                        case "":
                            _ = len(v4)
                        }
                    }
                }
            case "a":
                fmt.Println(v2, v7, v2)
            case "":
                if len(v1) > 0 {
                    v1[0] = fi41
                }
            default:
                _ = len(v1)
            }
        } else {
            for fi48 := 0; fi48 < 17; fi48++ {
                var v49 uint32 = 2147483648
                var v50 uint = uint(v49)
                _ = fi48
                _ = v49
                _ = v50
            }
            fmt.Println(v7)
            var v51 uint16 = v40
            _ = v51
        }
        _ = fi41
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
