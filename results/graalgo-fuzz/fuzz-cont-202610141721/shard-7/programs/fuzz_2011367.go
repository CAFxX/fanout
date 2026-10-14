package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 bool; F1 int64; F2 int64 }
func genFn0(p0 int, p1 uint64) uint {
    v4 := []int{p0, 4294967295}
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 uint16 = 32767
        if (1 < 42) {
            fmt.Println(p1)
            var v7 string = "a"
            var v8 uint = 4294967295
            _ = v7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v9 := GenS1{F0: 1}
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            p1 -= p1
            _ = fi11
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := []string{"hello"}
    var v2 int16 = 32767
    v3 := []int{375}
    var v12 int8 = -1
    var v13 uint64 = 0
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS1{F0: 2147483647}
        v1 = append(v1, "a")
        v16 := GenS2{F0: 0}
        _ = fi14
        _ = v15
        _ = v16
    }
    v2 = v2
    var v17 float32 = 0.0
    v17 = (-2.25 + v17)
    v13 = (v13 - v13)
    for fi18 := 0; fi18 < 7; fi18++ {
        v17 = ((v17 * float32(3.14159)) - (0.1 * (v17 + 1.5)))
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 int32 = 0
            v1 = append(v1, (("\n\t" + ((("" + ("a" + "\n\t")) + "") + "hello")) + ("a" + "\n\t")))
            v20 = v20
            _ = len(v1)
            switch int(v13) {
            case 1:
                if ((!true) && (483 <= 4294967295)) {
                    if false {
                        var v21 int16 = v2
                        var v22 int16 = v2
                        var v23 int32 = v20
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        switch (-(fi18)) {
                        case 1:
                            v13 = (v13 ^ v13)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case 127:
                            v24 := GenS1{F0: 65535}
                            _ = v24
                        default:
                            if (2147483647 != 2147483648) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if false {
                                    v17 -= -2.25
                                }
                            }
                        }
                        v20 -= 1
                        var v25 uint32 = 1
                        _ = v25
                    }
                    var v26 int8 = v12
                    v27 := GenS3{F0: v20, F1: 4294967295}
                    if false {
                        if (2 >= 398) {
                            fmt.Println(v2, v13, fi19)
                            v26 += -128
                            v2 = 1000
                        } else {
                            v28 := GenS1{F0: 1}
                            if false {
                                v29 := GenS3{F0: 155, F1: v13}
                                fmt.Println(v17, v26)
                                v17 -= (0.0 - 3.14159)
                                _ = v29
                            } else {
                                if true {
                                    if (((v13 >> ((genFn0(fi18, (v13 / ((42) | 1)))) & 63)) != 4294967295) && false) {
                                        if true {
                                            var v30 float64 = 0.1
                                            v12 = v26
                                            v26 = (v12 * v26)
                                            _ = v30
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            switch 601 {
                                            case -188:
                                                var v31 int32 = (-(v20))
                                                _ = v31
                                            case 9223372036854775807:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v32 int64 = 237
                                                    _ = v32
                                                case "foo bar":
                                                    v3 = append(v3, 1)
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v33 bool = (-1 <= -9223372036854775808)
                                                if len(v1) > 0 {
                                                    v1[0] = "a"
                                                }
                                                _ = v33
                                            default:
                                                var v34 int = (fi18 << ((444) & 63))
                                                if false {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = len(v3)
                                                }
                                                if false {
                                                    if ((2 != 127) || ((fi19 >> ((65535) & 63)) > (v34 / (((-(fi19))) | 1)))) {
                                                        var v35 uint8 = 255
                                                        v26 += -128
                                                        v36 := GenS4{F0: 127, F1: v20}
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                    switch (fi19 % ((fi18) | 1)) {
                                                    case -2147483648:
                                                        if (true && (true && ((-2.25 * 0.0) >= 1e10))) {
                                                            v37 := GenS5{F0: fi18, F1: 204}
                                                            fmt.Println(v17, v20, v13)
                                                            fmt.Println(v17, v20, v2)
                                                            _ = v37
                                                        } else {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            v2 += (v2 / ((v2) | 1))
                                                        }
                                                    case -209:
                                                        if (false || false) {
                                                            fmt.Println(v13, v12, v13)
                                                        } else {
                                                            fmt.Println(v34, v12)
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v12, v17, fi18)
                                                }
                                                _ = v34
                                            }
                                        }
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        v17 += ((float32(((1.5 + ((3.14159 + float64(3.14159)) * (float64(v17) + (1e10 / 1.0)))) * (3.14159 + ((1.5 + 1e10) * 3.14159)))) + v17) * v17)
                                    case "a":
                                        var v38 int8 = (v26 ^ -128)
                                        _ = v38
                                    case "hello":
                                        var v39 int64 = -1
                                        _ = v39
                                    default:
                                        v26 = v26
                                    }
                                    v40 := GenS1{F0: 2147483648}
                                    _ = v40
                                } else {
                                    if (18446744073709551615 >= v13) {
                                        if ((269 > 1) || true) {
                                            if len(v1) > 0 {
                                                v1[0] = "hello"
                                            }
                                        }
                                        _ = len(v1)
                                    } else {
                                        v41 := GenS6{F0: -128, F1: ("a" + (("" + ("" + "")) + "a"))}
                                        fmt.Println(v20)
                                        _ = v41
                                    }
                                }
                            }
                            _ = v28
                        }
                        v20 = -2147483648
                    }
                    _ = v26
                    _ = v27
                } else {
                    var v42 int64 = 42
                    v2 = v2
                    _ = v42
                }
            case 32768:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 32767:
                if (!(!true)) {
                    var v43 uint16 = 65535
                    var v44 uint = 612
                    switch (fi18 & (-(fi19))) {
                    case 103:
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                    }
                    _ = v43
                    _ = v44
                } else {
                    fmt.Println(v20)
                    fmt.Println(v12, v17, fi19)
                }
            default:
                switch "" {
                case "\n\t":
                    switch 2147483647 {
                    case -725:
                        _ = len(v3)
                    }
                    var v45 int8 = 127
                    _ = v45
                case "foo bar":
                    switch ((("a" + "") + "a") + "\n\t") {
                    case "a":
                        v12 = v12
                    case "hello":
                        if false {
                            var v46 float64 = ((3.14159 * 0.1) - (float64(v17) + (3.14159 - (0.1 + 3.14159))))
                            var v47 uint8 = 128
                            _ = v46
                            _ = v47
                        }
                    default:
                        var v48 uint64 = (v13 << ((42) & 63))
                        _ = v48
                    }
                default:
                    _ = len(v3)
                }
                v12 = (v12 * -1)
            }
            _ = fi19
            _ = v20
        }
        switch -1 {
        case -743:
            v49 := GenS6{F0: (v12 >> ((745) & 63)), F1: ""}
            _ = v49
        case 32767:
            v50 := GenS1{F0: 2147483647}
            _ = v50
        default:
            for fi51 := 0; fi51 < 7; fi51++ {
                if (("\n\t" != "\n\t") && (true || (181 >= v13))) {
                    var v52 int = (fi51 << ((18446744073709551615) & 63))
                    switch 42 {
                    case 32767:
                        fmt.Println(v13)
                    case 42:
                        v1 = append(v1, "")
                    default:
                        var v53 string = ("a" + "\n\t")
                        _ = v53
                    }
                    _ = v52
                }
                _ = fi51
            }
        }
        _ = fi18
    }
    if (32767 == 9223372036854775807) {
        if true {
            var v54 bool = (0.1 > v17)
            if true {
                var v55 int16 = 32767
                fmt.Println(v54, v13)
                _ = v55
            }
            v54 = (382 > 0)
            _ = v54
        }
        var v56 string = "hello"
        v57 := GenS4{F0: (v12 % (((v12 & v12)) | 1)), F1: 2147483647}
        _ = v56
        _ = v57
    } else {
        for fi58 := 0; fi58 < 1; fi58++ {
            _ = len(v1)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v17 = -2.25
            _ = fi58
        }
        for fi59 := 0; fi59 < 3; fi59++ {
            switch "\n\t" {
            case "foo bar":
                var v60 uint64 = (v13 + v13)
                _ = v60
            case "\n\t":
                var v61 int16 = v2
                _ = v61
            default:
                _ = len(v1)
            }
            for fi62 := 0; fi62 < 13; fi62++ {
                v63 := GenS5{F0: (fi59 ^ (fi59 ^ 42)), F1: 681}
                switch ("a" + "\n\t") {
                case "a":
                    v64 := GenS7{F0: true, F1: 198, F2: 0}
                    _ = v64
                case "":
                    fmt.Println(v17, fi62, v13)
                }
                var v65 int64 = -717
                _ = fi62
                _ = v63
                _ = v65
            }
            _ = len(v1)
            _ = fi59
        }
    }
    var v66 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v66
}
