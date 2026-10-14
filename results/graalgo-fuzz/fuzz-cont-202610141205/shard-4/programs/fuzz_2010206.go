package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 uint32; F2 bool }
type GenS6 struct { F0 int64; F1 string; F2 int8 }
type GenS7 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS8 struct { F0 uint32; F1 uint; F2 int64 }
type GenS9 struct { F0 uint16; F1 int; F2 int8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int64; F1 int }
type GenS12 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS13 struct { F0 int16; F1 uint8; F2 string }

func main() {
    var v1 uint32 = 134
    var v2 uint8 = 89
    var v3 uint16 = 0
    v4 := make(map[string]bool)
    fmt.Println(v3, v1, v3)
    v2 = v2
    v5 := []int{0, 127, 1, -400}
    for fi6 := 0; fi6 < 20; fi6++ {
        switch (fi6 % (((fi6 & fi6)) | 1)) {
        case 128:
            var v7 uint = 1
            _ = v7
        default:
            if (132 >= 18446744073709551615) {
                v8 := GenS1{F0: 0, F1: -998}
                v9 := GenS2{F0: v3}
                _ = v8
                _ = v9
            } else {
                v10 := GenS3{F0: 127, F1: 1}
                switch fi6 {
                case 32768:
                    var v11 int8 = -13
                    _ = v11
                case 127:
                    var v12 int = (fi6 * (fi6 % (((fi6 | fi6)) | 1)))
                    _ = v12
                case -903:
                    for fi13 := 0; fi13 < 11; fi13++ {
                        var v14 bool = false
                        v5 = append(v5, (fi6 + (fi13 % ((fi6) | 1))))
                        v14 = true
                        switch fi13 {
                        case -1:
                            if v14 {
                                if (65535 < 350) {
                                    var v15 string = "foo bar"
                                    _ = v15
                                }
                                v16 := GenS1{F0: 0, F1: -410}
                                _ = v16
                            }
                        case -2147483648:
                            switch ("a" + "hello") {
                            case "foo bar":
                                _ = len(v5)
                            case "a":
                                switch "" {
                                case "\n\t":
                                    v17 := GenS4{F0: 2147483647}
                                    _ = v17
                                default:
                                    var v18 int16 = -911
                                    _ = v18
                                }
                            case "hello":
                                if v14 {
                                    _ = len(v5)
                                    v19 := GenS4{F0: 944}
                                    _ = v19
                                }
                            default:
                                if len(v5) > 0 {
                                    v5[0] = (fi13 >> ((42) & 63))
                                }
                            }
                        }
                        _ = fi13
                        _ = v14
                    }
                default:
                    if false {
                        var v20 uint = 742
                        var v21 int8 = 0
                        v5 = append(v5, fi6)
                        _ = v20
                        _ = v21
                    } else {
                        if (493 == 440) {
                            fmt.Println(v2)
                            switch (("hello" + ("foo bar" + "")) + "hello") {
                            case "hello":
                                fmt.Println(v1, v2)
                            case "":
                                fmt.Println(v1, v1, v1)
                            default:
                                v4["\n\t"] = (!(0.0 != 1e10))
                            }
                            v1 = (v1 | v1)
                        }
                        for fi22 := 0; fi22 < 14; fi22++ {
                            v1 = (v1 | (v1 << ((65535) & 63)))
                            switch fi22 {
                            case 32768:
                                v1 += 783
                            case 110:
                                fmt.Println(v1, fi6, fi6)
                            case 0:
                                if len(v5) > 0 {
                                    v5[0] = fi6
                                }
                            }
                            _ = len(v4)
                            _ = fi22
                        }
                    }
                }
                var v23 int8 = -128
                _ = v10
                _ = v23
            }
        }
        switch fi6 {
        case -1:
            v24 := GenS5{F0: int(v1), F1: v1, F2: true}
            _ = v24
        case -9223372036854775808:
            switch ("" + (("foo bar" + "\n\t") + "")) {
            case "hello":
                if true {
                    v2 += v2
                } else {
                    v25 := GenS1{F0: int16(v2), F1: 42}
                    if (fi6 != fi6) {
                        v26 := GenS6{F0: -9223372036854775808, F1: "hello", F2: -128}
                        _ = v26
                    }
                    _ = v25
                }
            case "foo bar":
                if (-1 > -249) {
                    v27 := GenS2{F0: v3}
                    var v28 uint = 65535
                    _ = v27
                    _ = v28
                } else {
                    delete(v4, ((("hello" + "a") + "") + "foo bar"))
                }
            case "":
                v29 := GenS7{F0: 4294967295, F1: 42, F2: (v1 / ((v1) | 1))}
                _ = v29
            default:
                fmt.Println(v2, v3, v3)
            }
        case 2147483647:
            switch (fi6 / ((fi6) | 1)) {
            case 43:
                v30 := GenS6{F0: -26, F1: "\n\t", F2: 127}
                _ = v30
            case 32768:
                v2 = 1
            default:
                for fi31 := 0; fi31 < 4; fi31++ {
                    var v32 int32 = 2147483647
                    v33 := GenS8{F0: 400, F1: 4294967295, F2: 9223372036854775807}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
        default:
            v2 += 108
        }
        _ = fi6
    }
    delete(v4, "")
    switch ("\n\t" + "a") {
    case "hello":
        v2 -= 33
    default:
        var v34 int8 = 100
        _ = v34
    }
    v3 += (v3 >> ((42) & 63))
    switch 42 {
    case 646:
        switch 471 {
        case 732:
            v35 := GenS9{F0: (v3 ^ (v3 << ((428) & 63))), F1: 32768, F2: int8(v2)}
            _ = v35
        default:
            v36 := GenS4{F0: 466}
            _ = v36
        }
    case 127:
        switch "a" {
        case "":
            var v37 string = ("a" + "hello")
            _ = v37
        case "foo bar":
            switch int(v1) {
            case 127:
                var v38 uint16 = v3
                _ = v38
            case 1:
                v2 = 127
            case 2147483647:
                var v39 uint32 = v1
                _ = v39
            }
        case "\n\t":
            var v40 int32 = 2147483647
            _ = v40
        }
    default:
        if (0 >= 4294967295) {
            v2 = v2
            v41 := GenS2{F0: (^(v3))}
            for fi42 := 0; fi42 < 12; fi42++ {
                if true {
                    v3 = (v3 & v3)
                } else {
                    for fi43 := 0; fi43 < 1; fi43++ {
                        v44 := GenS8{F0: v1, F1: 4294967295, F2: 42}
                        _ = fi43
                        _ = v44
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi42
            }
            _ = v41
        } else {
            v45 := v4[""]
            v46 := v4[""]
            _ = v45
            _ = v46
        }
    }
    var v47 bool = (!false)
    v2 = (^(v2))
    var v48 string = "\n\t"
    v4[((v48 + v48) + v48)] = (65535 < 272)
    if ((v1 >> ((212) & 63)) != 2147483648) {
        for fi49 := 0; fi49 < 15; fi49++ {
            var v50 float64 = (3.14159 * 0.0)
            for fi51 := 0; fi51 < 20; fi51++ {
                v47 = (0 <= 9223372036854775807)
                var v52 uint16 = 42
                var v53 uint = 255
                if v47 {
                    var v54 int64 = 0
                    v55 := GenS3{F0: -1, F1: 100}
                    if v47 {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v56 := GenS10{F0: v3}
                        v47 = v47
                        _ = v56
                    }
                    _ = v54
                    _ = v55
                } else {
                    switch v48 {
                    case "foo bar":
                        if ((!(!(int16(v1) > 1))) || (!(int16(fi49) >= -792))) {
                            v2 -= 0
                            var v57 float32 = 1.5
                            _ = v57
                        } else {
                            v2 += (^(v2))
                        }
                    case "\n\t":
                        if v47 {
                            var v58 float64 = (v50 - v50)
                            if (!v47) {
                                v50 += v50
                                var v59 float64 = (float64(-2.25) / (v58 - v58))
                                v58 = v50
                                _ = v59
                            } else {
                                v53 -= (v53 ^ v53)
                                v60 := GenS3{F0: 0, F1: 112}
                                var v61 int8 = 0
                                _ = v60
                                _ = v61
                            }
                            v47 = v47
                            _ = v58
                        }
                        var v62 int32 = -918
                        _ = v62
                    case "a":
                        v63 := GenS11{F0: -9223372036854775808, F1: (fi49 * fi51)}
                        _ = v63
                    default:
                        v47 = v47
                    }
                    if (v47 && ((!v47) || (!true))) {
                        v52 = v52
                        v50 = v50
                        v64 := GenS12{F0: 1, F1: (v2 * 0), F2: -1}
                        _ = v64
                    }
                }
                v65 := GenS13{F0: 32767, F1: (^(v2)), F2: ("\n\t" + (("a" + (v48 + v48)) + v48))}
                _ = fi51
                _ = v52
                _ = v53
                _ = v65
            }
            v4[("hello" + v48)] = false
            _ = fi49
            _ = v50
        }
    }
    _ = len(v5)
    v66 := GenS4{F0: -135}
    v3 = (v3 / ((255) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v66
}
