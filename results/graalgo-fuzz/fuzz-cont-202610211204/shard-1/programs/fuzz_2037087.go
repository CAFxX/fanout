package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS2 struct { F0 bool; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool }

func main() {
    var v1 int = 128
    var v2 int32 = 42
    var v3 uint64 = 18446744073709551615
    v4 := make(map[uint32]int)
    fmt.Println(v1, v1)
    v5 := []int{}
    v6 := v4[4294967295]
    for fi7 := 0; fi7 < 1; fi7++ {
        v3 += 2
        var v8 int64 = -9223372036854775808
        _ = fi7
        _ = v8
    }
    if (false || (false || (!true))) {
        v9 := GenS1{F0: 210, F1: int16(v1), F2: 1}
        var v10 uint8 = 127
        v2 = (v2 - -2147483648)
        _ = v10
        _ = v9
    }
    var v11 int32 = 2147483647
    v12 := GenS1{F0: 255, F1: 862, F2: 65535}
    v6 -= 0
    var v13 int8 = -1
    v11 = (^(v2))
    fmt.Println(v11)
    v1 = (v6 & v1)
    switch (v1 + (v6 ^ (v1 % ((v1) | 1)))) {
    case 32768:
        for fi14 := 0; fi14 < 19; fi14++ {
            v15 := GenS1{F0: 255, F1: -32768, F2: 59}
            v16 := GenS2{F0: ((0 >= int64(v3)) || (true && false)), F1: v2, F2: 1}
            _ = fi14
            _ = v15
            _ = v16
        }
    case 42:
        for fi17 := 0; fi17 < 11; fi17++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v2 += -1
            for fi18 := 0; fi18 < 17; fi18++ {
                switch "\n\t" {
                case "hello":
                    v2 = (v11 << ((4294967295) & 63))
                case "":
                    if len(v5) > 0 {
                        v5[0] = (fi18 ^ fi18)
                    }
                }
                if false {
                    v19 := GenS1{F0: 0, F1: 1, F2: 2}
                    _ = v19
                } else {
                    v4[0] = -1
                    if false {
                        v4[0] = fi17
                    } else {
                        v3 -= v3
                        if (false || (false && (!false))) {
                            v6 = v6
                            if (!(-32768 > 0)) {
                                delete(v4, 2147483648)
                                switch (^(fi17)) {
                                case 127:
                                    v3 = (v3 + (v3 * v3))
                                case 2147483647:
                                    if (-1 != 1000) {
                                        fmt.Println(fi17, v6, v11)
                                        var v20 int32 = -2147483648
                                        var v21 uint64 = (v3 / ((42) | 1))
                                        _ = v20
                                        _ = v21
                                    } else {
                                        v22 := GenS3{F0: 2}
                                        v6 = v1
                                        switch "\n\t" {
                                        case "a":
                                            v2 += v2
                                            v13 = v13
                                        default:
                                            switch "a" {
                                            case "":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v23 string = "\n\t"
                                                    _ = v23
                                                case "\n\t":
                                                    var v24 uint32 = 2147483648
                                                    _ = v24
                                                case "hello":
                                                    if (255 > 0) {
                                                        if (-2.25 < 1e10) {
                                                            fmt.Println(fi17, v6)
                                                            v25 := GenS2{F0: true, F1: (v11 + 0), F2: 786}
                                                            v13 = v13
                                                            _ = v25
                                                        }
                                                        fmt.Println(v13, fi17)
                                                    } else {
                                                        v26 := GenS1{F0: 128, F1: 32767, F2: 999}
                                                        delete(v4, 4294967295)
                                                        _ = v26
                                                    }
                                                default:
                                                    _ = len(v5)
                                                }
                                            default:
                                                v4[52] = (fi18 % (((^(fi18))) | 1))
                                            }
                                            v27 := GenS4{F0: 2}
                                            var v28 string = "hello"
                                            _ = v27
                                            _ = v28
                                        }
                                        if (1000 != 32767) {
                                            var v29 uint16 = 2
                                            _ = v29
                                        } else {
                                            v6 = (v6 / ((v6) | 1))
                                            if true {
                                                var v30 int8 = (v13 << ((293) & 63))
                                                _ = v30
                                            }
                                            v1 -= v6
                                        }
                                        _ = v22
                                    }
                                case 1:
                                    switch ("\n\t" + ((("foo bar" + ("foo bar" + "foo bar")) + "\n\t") + "a")) {
                                    case "":
                                        if false {
                                            v31 := GenS3{F0: 734}
                                            fmt.Println(fi17, fi17)
                                            v3 -= v3
                                            _ = v31
                                        }
                                    case "a":
                                        v4[0] = (fi17 + fi17)
                                    case "foo bar":
                                        switch v1 {
                                        case 0:
                                            v2 -= 333
                                        case -9223372036854775808:
                                            if true {
                                                var v32 int = v6
                                                switch "hello" {
                                                case "a":
                                                    v6 = 1
                                                case "hello":
                                                    if (!false) {
                                                        if true {
                                                            if (0.1 <= (3.14159 / -2.25)) {
                                                                if false {
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        v2 = 0
                                                                    case "\n\t":
                                                                        var v33 string = "a"
                                                                        _ = v33
                                                                    case "hello":
                                                                        switch (v6 % ((4294967295) | 1)) {
                                                                        case 9223372036854775807:
                                                                            v3 -= 4294967295
                                                                        case -9223372036854775808:
                                                                            var v34 float64 = 1.5
                                                                            _ = v34
                                                                        case 4294967295:
                                                                            _ = len(v5)
                                                                        }
                                                                    }
                                                                    var v35 float32 = (0.0 - 1e10)
                                                                    var v36 int8 = 1
                                                                    _ = v35
                                                                    _ = v36
                                                                } else {
                                                                    fmt.Println(fi18, fi17)
                                                                }
                                                            }
                                                        } else {
                                                            v11 = v11
                                                        }
                                                        if (v32 >= fi17) {
                                                            if (75 != 70) {
                                                                v2 -= (v11 / (((v11 << ((18446744073709551615) & 63))) | 1))
                                                            }
                                                            if len(v5) > 0 {
                                                                v5[0] = v6
                                                            }
                                                        }
                                                    }
                                                case "":
                                                    v6 = (v6 / ((v1) | 1))
                                                default:
                                                    v6 += (fi17 / (((v32 << ((18446744073709551615) & 63))) | 1))
                                                }
                                                v6 = int(v13)
                                                var v37 float64 = 1.5
                                                fmt.Println(v32, v1)
                                                _ = v32
                                                _ = v37
                                            } else {
                                                switch "a" {
                                                case "foo bar":
                                                    v13 = v13
                                                case "a":
                                                    switch -9223372036854775808 {
                                                    case 32767:
                                                        v1 = (fi18 | v1)
                                                    case 32768:
                                                        var v38 int64 = 1
                                                        _ = v38
                                                    default:
                                                        v6 = (v1 ^ (v1 + v1))
                                                    }
                                                }
                                                v2 += -2147483648
                                                var v39 bool = true
                                                var v40 string = "hello"
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                    default:
                                        switch "\n\t" {
                                        case "hello":
                                            var v41 uint16 = uint16(v3)
                                            _ = v41
                                        }
                                    }
                                }
                                var v42 uint32 = 905
                                _ = v42
                            } else {
                                fmt.Println(v2)
                                switch ("hello" + (("hello" + ("hello" + ("" + ""))) + "")) {
                                case "hello":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case "foo bar":
                                    if len(v5) > 0 {
                                        v5[0] = v6
                                    }
                                default:
                                    v2 += v2
                                    v6 -= 1
                                }
                            }
                        }
                        switch "a" {
                        case "\n\t":
                            v3 = v3
                        default:
                            if len(v5) > 0 {
                                v5[0] = fi18
                            }
                        }
                    }
                }
                _ = fi18
            }
            _ = fi17
        }
    case 2147483647:
        if false {
            v43 := GenS2{F0: false, F1: (v11 * 431), F2: 42}
            if true {
                v6 -= v1
                v6 -= 395
            } else {
                if false {
                    fmt.Println(v1, v2, v3)
                    v1 = v6
                } else {
                    _ = len(v4)
                    v44 := v4[296]
                    _ = v44
                }
                v4[651] = v1
                v3 -= v3
            }
            var v45 uint16 = uint16(v3)
            _ = v43
            _ = v45
        }
    default:
        if (v13 >= v13) {
            switch (("hello" + "foo bar") + ("hello" + ("foo bar" + "hello"))) {
            case "a":
                var v46 uint = 65535
                _ = v46
            }
            switch "\n\t" {
            case "hello":
                var v47 bool = false
                _ = v47
            }
        }
    }
    for fi48 := 0; fi48 < 17; fi48++ {
        v49 := GenS5{F0: true}
        for fi50 := 0; fi50 < 12; fi50++ {
            var v51 bool = ((true && ((991 > 2147483647) || (false && true))) && false)
            var v52 int32 = v11
            _ = fi50
            _ = v51
            _ = v52
        }
        _ = fi48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
