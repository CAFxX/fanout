package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int64; F1 int; F2 int8 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int8; F1 uint64 }

func main() {
    var v1 int64 = -22
    v2 := make(map[uint32]uint32)
    var v3 bool = ((true || (("foo bar" == "a") || (!(!true)))) && false)
    v4 := []string{"foo bar", "a", "\n\t"}
    fmt.Println(v1, v3)
    v5 := GenS1{F0: 43, F1: 42}
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    var v6 int16 = 167
    var v7 string = ""
    if (v3 && (!(32767 <= 748))) {
        if true {
            if v3 {
                var v8 uint = 882
                delete(v2, 65535)
                _ = v8
            }
            fmt.Println(v3)
            v4 = append(v4, "a")
        } else {
            var v9 uint = 0
            _ = v9
        }
    } else {
        v10 := GenS1{F0: v1, F1: 2}
        var v11 uint32 = 4294967295
        v7 = "\n\t"
        _ = v10
        _ = v11
    }
    switch (("hello" + ("hello" + (v7 + (("a" + v7) + ("foo bar" + (v7 + (("foo bar" + v7) + (v7 + v7)))))))) + "\n\t") {
    case "a":
        if (1e10 < 1e10) {
            _ = len(v2)
            v1 -= v1
            var v12 float64 = 1.5
            _ = v12
        }
    case "hello":
        var v13 uint = 827
        _ = v13
    default:
        for fi14 := 0; fi14 < 12; fi14++ {
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 float64 = -2.25
                _ = fi15
                _ = v16
            }
            v1 = (v1 % ((v1) | 1))
            _ = fi14
        }
    }
    fmt.Println(v3)
    v4 = append(v4, "a")
    for fi17 := 0; fi17 < 18; fi17++ {
        switch fi17 {
        case 306:
            if (318 >= 42) {
                if (("hello" + v7) >= "\n\t") {
                    v6 = 1000
                }
                var v18 float64 = 0.0
                for fi19 := 0; fi19 < 2; fi19++ {
                    var v20 uint16 = 2
                    fmt.Println(v6)
                    if (fi19 <= (fi17 / ((fi17) | 1))) {
                        v18 += v18
                        v6 += v6
                    } else {
                        if ((-1 < 526) || v3) {
                            if (true && v3) {
                                v3 = (((0.1 * float32((float64(3.14159) + v18))) * (1.5 * ((((((0.1 * float32(float64((1e10 * float32((0.1 / v18)))))) + (0.0 * float32(1e10))) / -2.25) * 1.5) / 0.1) + (1.5 * 1e10)))) != -2.25)
                                v4 = append(v4, "foo bar")
                            }
                            var v21 bool = (-9223372036854775808 < -1)
                            _ = v21
                        } else {
                            v22 := GenS2{F0: 2, F1: 0}
                            v23 := GenS3{F0: 1, F1: (fi17 - -9223372036854775808), F2: int8(v20)}
                            _ = v22
                            _ = v23
                        }
                    }
                    _ = fi19
                    _ = v20
                }
                _ = v18
            } else {
                switch fi17 {
                case 883:
                    v1 -= v1
                case 0:
                    delete(v2, 1)
                case 32767:
                    if false {
                        if false {
                            var v24 int = (fi17 * 0)
                            var v25 string = "hello"
                            _ = v24
                            _ = v25
                        }
                        var v26 int = 1
                        v27 := GenS4{F0: (-(v6)), F1: 9223372036854775807}
                        _ = v26
                        _ = v27
                    } else {
                        if ((v6 / ((v6) | 1)) > (v6 << ((418) & 63))) {
                            for fi28 := 0; fi28 < 2; fi28++ {
                                var v29 uint32 = 2147483648
                                v30 := GenS5{F0: -32, F1: 445}
                                var v31 int = fi17
                                _ = fi28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        }
                    }
                }
            }
        case -2147483648:
            if len(v4) > 0 {
                v4[0] = "hello"
            }
        case 4294967295:
            for fi32 := 0; fi32 < 16; fi32++ {
                if (129 < 190) {
                    switch (fi17 % ((int(v6)) | 1)) {
                    case -53:
                        v1 += v1
                    }
                    if v3 {
                        v33 := GenS3{F0: -363, F1: fi32, F2: 42}
                        v7 = "a"
                        if false {
                            if (32767 > 32767) {
                                v1 -= -112
                                var v34 string = v7
                                _ = v34
                            }
                        } else {
                            var v35 float64 = float64((1e10 * 0.1))
                            v1 -= (v1 & (-(v1)))
                            _ = v35
                        }
                        switch fi32 {
                        case 127:
                            switch ("a" + v7) {
                            case "foo bar":
                                var v36 uint8 = 127
                                _ = v36
                            case "hello":
                                var v37 uint = 0
                                _ = v37
                            case "\n\t":
                                var v38 uint16 = 0
                                _ = v38
                            default:
                                var v39 uint64 = 2
                                _ = v39
                            }
                        case -1:
                            if true {
                                fmt.Println(v6)
                                fmt.Println(v6, v7)
                                v40 := GenS2{F0: 18446744073709551615, F1: 4294967295}
                                _ = v40
                            } else {
                                var v41 int64 = (v1 ^ v1)
                                var v42 int16 = v6
                                switch v7 {
                                case "\n\t":
                                    v1 = v41
                                case "foo bar":
                                    var v43 int32 = 669
                                    _ = v43
                                default:
                                    if ((139 > 4294967295) || v3) {
                                        v44 := GenS3{F0: v41, F1: 32767, F2: 1}
                                        delete(v2, 813)
                                        var v45 int64 = v41
                                        _ = v44
                                        _ = v45
                                    } else {
                                        v46 := GenS4{F0: v6, F1: (fi17 & fi32)}
                                        v4 = append(v4, (v7 + ("hello" + (("foo bar" + (v7 + (v7 + v7))) + v7))))
                                        _ = v46
                                    }
                                }
                                _ = v41
                                _ = v42
                            }
                        case 813:
                            if (1e10 <= (3.14159 * ((0.0 + 0.1) * 1e10))) {
                                fmt.Println(v1)
                            }
                        }
                        switch ((v7 + "") + (v7 + v7)) {
                        case "a":
                            v1 = (v1 | 50)
                        }
                        switch (fi17 / (((fi17 & -2147483648)) | 1)) {
                        case 1:
                            fmt.Println(fi17)
                        case -9223372036854775808:
                            var v47 int32 = 731
                            _ = v47
                        case 9223372036854775807:
                            v48 := GenS4{F0: int16(fi17), F1: fi32}
                            _ = v48
                        default:
                            v6 -= v6
                        }
                        var v49 bool = v3
                        var v50 int8 = 100
                        if (!(!((-(v1)) <= (v1 << ((65535) & 63))))) {
                            if (!(0.0 < float64(((-2.25 * ((-2.25 + 0.0) / 3.14159)) * 1.5)))) {
                                var v51 int = fi17
                                _ = v51
                            }
                            switch "a" {
                            case "":
                                var v52 int64 = (-(v1))
                                _ = len(v4)
                                v53 := GenS4{F0: (^(v6)), F1: (fi17 / ((fi32) | 1))}
                                _ = v52
                                _ = v53
                            case "a":
                                switch "a" {
                                case "":
                                    fmt.Println(v1, v3, fi32)
                                default:
                                    fmt.Println(fi17)
                                }
                            default:
                                if len(v4) > 0 {
                                    v4[0] = v7
                                }
                            }
                            var v54 string = (((v7 + ("" + v7)) + (("foo bar" + "\n\t") + v7)) + v7)
                            _ = v54
                        } else {
                            var v55 uint64 = 18446744073709551615
                            v49 = false
                            _ = v55
                        }
                        if len(v4) > 0 {
                            v4[0] = v7
                        }
                        v56 := GenS1{F0: 0, F1: 519}
                        v49 = (255 < 145)
                        _ = v33
                        _ = v49
                        _ = v50
                        _ = v56
                    } else {
                        var v57 uint32 = 65535
                        if true {
                            var v58 int32 = 0
                            _ = v58
                        } else {
                            var v59 uint8 = 255
                            v60 := v2[(v57 ^ (^(v57)))]
                            if v3 {
                                fmt.Println(v7)
                            }
                            _ = v59
                            _ = v60
                        }
                        v4 = append(v4, v7)
                        _ = v57
                    }
                } else {
                    v61 := v2[2]
                    var v62 float32 = 1.5
                    var v63 float32 = (v62 - (float32(float64(v62)) + v62))
                    if true {
                        var v64 uint32 = v61
                        _ = v64
                    } else {
                        v6 = v6
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v3 = v3
                    _ = len(v2)
                    if v3 {
                        v1 = (v1 | (v1 + (-(v1))))
                        if v3 {
                            var v65 uint = 599
                            delete(v2, v61)
                            fmt.Println(v65)
                            _ = v65
                        }
                        v6 = (v6 & v6)
                    }
                    var v66 uint32 = v61
                    _ = v61
                    _ = v62
                    _ = v63
                    _ = v66
                }
                _ = fi32
            }
        default:
            var v67 int16 = v6
            _ = v67
        }
        _ = fi17
    }
    v1 = v1
    v68 := GenS2{F0: 1, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v7
}
