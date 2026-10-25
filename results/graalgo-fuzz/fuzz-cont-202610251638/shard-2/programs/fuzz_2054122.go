package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int16; F2 string }
type GenS4 struct { F0 int8; F1 int; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 int32; F2 int64 }

func main() {
    var v1 int8 = -128
    var v2 int64 = 9223372036854775807
    var v3 bool = false
    var v4 uint32 = 65535
    var v5 uint = 678
    var v6 int64 = v2
    v5 = v5
    v7 := GenS1{F0: (v1 ^ int8(v5)), F1: v3, F2: v1}
    v8 := []int{}
    var v9 uint64 = 18446744073709551615
    if true {
        fmt.Println(v3, v9, v4)
    } else {
        if ((1 > -392) && false) {
            v6 = (v2 + (^(v6)))
            fmt.Println(v2)
            if v3 {
                switch "hello" {
                case "foo bar":
                    var v10 uint64 = v9
                    _ = v10
                case "":
                    fmt.Println(v6)
                case "a":
                    if (int32(v9) == -470) {
                        v11 := GenS2{F0: 1}
                        _ = v11
                    } else {
                        if len(v8) > 0 {
                            v8[0] = -54
                        }
                    }
                default:
                    var v12 int8 = v1
                    _ = v12
                }
            } else {
                v13 := GenS1{F0: v1, F1: (204 != 42), F2: (v1 - v1)}
                _ = v13
            }
        } else {
            var v14 int64 = 0
            if true {
                var v15 int16 = 42
                v9 += v9
                _ = v15
            }
            _ = v14
        }
        var v16 uint8 = 130
        _ = v16
    }
    if (2 > (v9 ^ v9)) {
        switch "hello" {
        case "a":
            if true {
                var v17 float64 = ((float64(float32(0.1)) / 1.0) - 0.1)
                var v18 float32 = (0.1 / ((1e10 / 1.0) / 0.1))
                v3 = ((v1 >> (((v5 - (^(v5)))) & 63)) != (v1 ^ (v1 + (v1 << ((255) & 63)))))
                _ = v17
                _ = v18
            } else {
                if (!v3) {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        fmt.Println(v3, v9)
                        _ = fi19
                    }
                } else {
                    v20 := GenS1{F0: (v1 >> (((v5 * v5)) & 63)), F1: (-533 >= 0), F2: (v1 | 42)}
                    v21 := GenS2{F0: -304}
                    if false {
                        if len(v8) > 0 {
                            v8[0] = 1
                        }
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        for fi22 := 0; fi22 < 20; fi22++ {
                            var v23 uint64 = 4294967295
                            if true {
                                v23 = (v23 | v23)
                                v24 := GenS3{F0: (fi22 << ((v5) & 63)), F1: 1, F2: "a"}
                                _ = v24
                            } else {
                                fmt.Println(v2)
                                if len(v8) > 0 {
                                    v8[0] = (fi22 / ((fi22) | 1))
                                }
                                v3 = (65535 == 65535)
                            }
                            _ = fi22
                            _ = v23
                        }
                    } else {
                        v6 = (v2 / (((v6 & (v2 << (((^(v5))) & 63)))) | 1))
                        if false {
                            if ((65535 != (v5 / ((v5) | 1))) && ("" < "")) {
                                for fi25 := 0; fi25 < 13; fi25++ {
                                    v26 := GenS2{F0: -32768}
                                    var v27 uint16 = 0
                                    var v28 uint16 = (v27 >> (((v5 * (v5 + (v5 ^ (v5 + v5))))) & 63))
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                                for fi29 := 0; fi29 < 10; fi29++ {
                                    if false {
                                        switch ("a" + "foo bar") {
                                        case "\n\t":
                                            for fi30 := 0; fi30 < 3; fi30++ {
                                                if v3 {
                                                    v8 = append(v8, 157)
                                                }
                                                var v31 int64 = -1
                                                if (-672 > -791) {
                                                    if true {
                                                        v32 := GenS4{F0: (v1 - (v1 - (v1 + v1))), F1: (fi29 - 4294967295), F2: int16(v4)}
                                                        switch (("a" + ("foo bar" + ((("\n\t" + "") + "a") + "hello"))) + ("\n\t" + "hello")) {
                                                        case "\n\t":
                                                            fmt.Println(v31, v1)
                                                        case "hello":
                                                            switch -9223372036854775808 {
                                                            case -1:
                                                                v33 := GenS5{F0: 0}
                                                                _ = v33
                                                            case 631:
                                                                v4 = (v4 + (^(v4)))
                                                                switch ("hello" + ("hello" + "\n\t")) {
                                                                case "\n\t":
                                                                    v34 := GenS6{F0: (v6 + -9223372036854775808), F1: int(v31)}
                                                                    _ = v34
                                                                }
                                                                v8 = append(v8, fi30)
                                                            }
                                                        case "":
                                                            var v35 bool = (!v3)
                                                            _ = v35
                                                        }
                                                        if v3 {
                                                            var v36 int16 = 32767
                                                            _ = v36
                                                        }
                                                        _ = v32
                                                    } else {
                                                        var v37 int16 = 32767
                                                        _ = v37
                                                    }
                                                }
                                                _ = fi30
                                                _ = v31
                                            }
                                        case "a":
                                            var v38 float32 = -2.25
                                            _ = v38
                                        case "hello":
                                            for fi39 := 0; fi39 < 10; fi39++ {
                                                var v40 int = fi29
                                                _ = fi39
                                                _ = v40
                                            }
                                        }
                                        if len(v8) > 0 {
                                            v8[0] = -276
                                        }
                                        if (false || ((v9 * v9) > (v9 + (v9 & 2)))) {
                                            if len(v8) > 0 {
                                                v8[0] = fi29
                                            }
                                            var v41 uint8 = 127
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            _ = v41
                                        }
                                    } else {
                                        if false {
                                            var v42 int8 = 0
                                            v2 += v2
                                            _ = v42
                                        }
                                        var v43 int16 = 742
                                        fmt.Println(v3)
                                        _ = v43
                                    }
                                    _ = fi29
                                }
                                var v44 uint32 = v4
                                _ = v44
                            } else {
                                if len(v8) > 0 {
                                    v8[0] = 32767
                                }
                                v45 := GenS7{F0: -999}
                                _ = v45
                            }
                            v46 := GenS3{F0: int(v4), F1: 395, F2: ("foo bar" + "hello")}
                            _ = v46
                        }
                        v3 = v3
                    }
                    _ = v20
                    _ = v21
                }
                for fi47 := 0; fi47 < 6; fi47++ {
                    if ((v5 - v5) < v5) {
                        fmt.Println(v4, v9)
                        var v48 int32 = int32(v4)
                        v2 -= (v2 | int64(v5))
                        _ = v48
                    }
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v49 uint8 = uint8(v9)
                    _ = fi47
                    _ = v49
                }
                _ = len(v8)
            }
        case "hello":
            for fi50 := 0; fi50 < 2; fi50++ {
                for fi51 := 0; fi51 < 19; fi51++ {
                    _ = len(v8)
                    v2 -= v2
                    _ = fi51
                }
                _ = fi50
            }
        }
        var v52 int = 127
        var v53 uint8 = 1
        _ = v52
        _ = v53
    } else {
        var v54 uint8 = 255
        var v55 float64 = 1.5
        _ = v54
        _ = v55
    }
    if v3 {
        var v56 int32 = int32(v5)
        fmt.Println(v6, v3, v2)
        _ = v56
    } else {
        v57 := GenS3{F0: 4294967295, F1: 404, F2: ""}
        v2 += (v2 % (((v6 - v6)) | 1))
        switch 61 {
        case -597:
            switch 42 {
            case 1:
                var v58 bool = v3
                _ = v58
            case -2147483648:
                v59 := GenS8{F0: 589, F1: -1, F2: (v6 ^ v6)}
                _ = v59
            }
        }
        _ = v57
    }
    for fi60 := 0; fi60 < 19; fi60++ {
        v9 += (^(v9))
        var v61 int = -562
        for fi62 := 0; fi62 < 8; fi62++ {
            v6 += 9223372036854775807
            switch ("foo bar" + ("hello" + "foo bar")) {
            case "\n\t":
                var v63 uint32 = (v4 >> (((v5 / ((42) | 1))) & 63))
                _ = v63
            }
            _ = fi62
        }
        _ = fi60
        _ = v61
    }
    var v64 int16 = 403
    v2 += int64(v1)
    var v65 int = 32768
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
    _ = v7
    _ = v8
    _ = v9
}
