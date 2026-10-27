package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 string; F1 uint64; F2 uint }
func genFn0(p0 string) bool {
    p0 = (((((((p0 + ((p0 + p0) + p0)) + (p0 + p0)) + p0) + p0) + "hello") + p0) + p0)
    var v5 int32 = -2147483648
    _ = p0
    _ = v5
    return true
}
func genFn1() uint16 {
    var v6 bool = true
    fmt.Println(v6)
    v7 := GenS1{F0: 42}
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int8 = 100
    var v2 int16 = -1
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    var v8 float32 = -2.25
    v9 := []int{32767, 38, -9223372036854775808, -364}
    if len(v9) > 0 {
        v9[0] = 4294967295
    }
    if (127 > 162) {
        v10 := GenS2{F0: 0}
        if false {
            var v11 int16 = 32767
            for fi12 := 0; fi12 < 6; fi12++ {
                for fi13 := 0; fi13 < 15; fi13++ {
                    fmt.Println(v1, v1)
                    var v14 int = (-(fi13))
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v11, v8, v8)
                _ = fi12
            }
            v8 -= 0.1
            _ = v11
        } else {
            delete(v3, "hello")
            var v15 float64 = 3.14159
            for fi16 := 0; fi16 < 7; fi16++ {
                v2 = (v2 >> ((2) & 63))
                v2 = v2
                v17 := GenS3{F0: (!(2147483648 <= 42))}
                _ = fi16
                _ = v17
            }
            _ = v15
        }
        v18 := GenS3{F0: false}
        _ = v10
        _ = v18
    } else {
        v19 := GenS2{F0: 32767}
        v2 += (^(v2))
        _ = v19
    }
    var v20 uint32 = 42
    v1 = v1
    switch 127 {
    case 127:
        switch 0 {
        case 127:
            switch (("\n\t" + (("foo bar" + "a") + ("foo bar" + "\n\t"))) + ("a" + ("foo bar" + "hello"))) {
            case "":
                v21 := GenS4{F0: ("" + "foo bar"), F1: -2147483648, F2: 116}
                _ = v21
            case "a":
                var v22 int = -114
                _ = v22
            }
        case 128:
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 int64 = -9223372036854775808
                if false {
                    v2 -= (v2 + 32767)
                } else {
                    var v25 int64 = (v24 & 9223372036854775807)
                    v24 = v24
                    if ((true && (42 > 127)) && true) {
                        v26 := GenS3{F0: false}
                        switch "foo bar" {
                        case "foo bar":
                            var v27 int64 = v24
                            _ = v27
                        case "":
                            for fi28 := 0; fi28 < 13; fi28++ {
                                _ = len(v9)
                                fmt.Println(v25, v2)
                                fmt.Println(v20, v1)
                                _ = fi28
                            }
                        case "\n\t":
                            var v29 uint = 255
                            _ = v29
                        default:
                            var v30 int32 = 205
                            _ = v30
                        }
                        v31 := GenS4{F0: "foo bar", F1: int32(v2), F2: 128}
                        _ = v26
                        _ = v31
                    } else {
                        if (4294967295 < 639) {
                            for fi32 := 0; fi32 < 3; fi32++ {
                                if false {
                                    var v33 uint64 = 46
                                    _ = v33
                                }
                                _ = fi32
                            }
                            v20 = v20
                        }
                        v34 := GenS5{F0: (v20 ^ (v20 ^ v20)), F1: (v2 >> ((1) & 63))}
                        v35 := v3[("foo bar" + "a")]
                        _ = v34
                        _ = v35
                    }
                    _ = v25
                }
                if ((fi23 ^ (fi23 & fi23)) != (fi23 + (fi23 >> ((0) & 63)))) {
                    v36 := GenS6{F0: (v1 / ((v1) | 1)), F1: (v20 << ((760) & 63))}
                    _ = v36
                }
                _ = fi23
                _ = v24
            }
        default:
            var v37 int8 = (v1 << ((607) & 63))
            _ = v37
        }
    case -903:
        if true {
            var v38 uint8 = 2
            _ = v38
        } else {
            v2 += -32768
            v20 -= (v20 ^ v20)
            v2 -= v2
        }
    case 4294967295:
        fmt.Println(v20, v1, v1)
    }
    v8 = float32(0.1)
    var v39 uint = 1
    v39 = 0
    if (803 > 712) {
        switch "hello" {
        case "hello":
            fmt.Println(v1, v2)
        case "\n\t":
            v1 -= 127
        case "a":
            for fi40 := 0; fi40 < 7; fi40++ {
                fmt.Println(v20)
                for fi41 := 0; fi41 < 12; fi41++ {
                    v20 += 428
                    switch "hello" {
                    case "a":
                        fmt.Println(fi41, fi41, v20)
                    case "hello":
                        if (((!true) || (-1 >= -1)) && (776 <= 407)) {
                            if (2147483647 <= -2147483648) {
                                if (((!(!(true || true))) && ((v20 <= (v20 & (^(v20)))) && (fi40 >= (fi41 % ((int(v39)) | 1))))) && true) {
                                    var v42 uint64 = 218
                                    v43 := GenS4{F0: "foo bar", F1: 0, F2: 2}
                                    fmt.Println(v1, v39, v8)
                                    _ = v42
                                    _ = v43
                                }
                            } else {
                                _ = len(v4)
                                v2 = (v2 - (v2 & -1))
                                if false {
                                    v1 += 12
                                    if true {
                                        switch ("foo bar" + "hello") {
                                        case "a":
                                            v4[1] = (^(fi41))
                                        default:
                                            switch "" {
                                            case "hello":
                                                v44 := GenS4{F0: "a", F1: 571, F2: 1}
                                                _ = v44
                                            default:
                                                v39 += uint(v20)
                                            }
                                        }
                                        var v45 float64 = -2.25
                                        var v46 uint64 = 1
                                        var v47 int8 = (^(v1))
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    } else {
                                        fmt.Println(v8, v8, v39)
                                        if (!(!((!(!true)) || (0.0 >= 3.14159)))) {
                                            switch -907 {
                                            case 9223372036854775807:
                                                fmt.Println(v39)
                                            }
                                            fmt.Println(v2)
                                        } else {
                                            v1 += 31
                                            delete(v3, "foo bar")
                                            _ = len(v9)
                                        }
                                        v48 := GenS7{F0: 128}
                                        _ = v48
                                    }
                                    v9 = append(v9, fi40)
                                } else {
                                    if true {
                                        if len(v9) > 0 {
                                            v9[0] = 1
                                        }
                                        v49 := GenS4{F0: ("a" + ("\n\t" + "\n\t")), F1: -2147483648, F2: 10}
                                        _ = v49
                                    }
                                    v50 := GenS7{F0: (^(fi40))}
                                    _ = v50
                                }
                            }
                        } else {
                            v51 := GenS1{F0: (v20 / ((v20) | 1))}
                            var v52 bool = false
                            v4[61] = (fi40 + fi40)
                            _ = v51
                            _ = v52
                        }
                    case "":
                        v53 := GenS8{F0: -366}
                        _ = v53
                    }
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = fi41
                }
                if ((v2 >> ((42) & 63)) > (v2 ^ (v2 - v2))) {
                    v54 := GenS9{F0: -9223372036854775808}
                    var v55 int = (fi40 >> ((v39) & 63))
                    v56 := GenS2{F0: 0}
                    _ = v54
                    _ = v55
                    _ = v56
                }
                _ = fi40
            }
        }
        var v57 uint64 = 18446744073709551615
        var v58 uint = v39
        _ = v57
        _ = v58
    } else {
        var v59 int8 = (v1 / (((v1 << ((42) & 63))) | 1))
        v39 = 65535
        _ = v59
    }
    switch 4294967295 {
    case -557:
        var v60 int16 = v2
        _ = v60
    case -1:
        if false {
            for fi61 := 0; fi61 < 2; fi61++ {
                for fi62 := 0; fi62 < 1; fi62++ {
                    v63 := GenS3{F0: (42 != 0)}
                    _ = fi62
                    _ = v63
                }
                var v64 bool = (int16(v1) >= (v2 >> ((65535) & 63)))
                v8 = v8
                _ = fi61
                _ = v64
            }
            v65 := GenS10{F0: ("hello" + ("" + ("foo bar" + "foo bar"))), F1: 2, F2: v39}
            for fi66 := 0; fi66 < 1; fi66++ {
                var v67 uint8 = 149
                _ = fi66
                _ = v67
            }
            _ = v65
        }
    case 2147483647:
        if false {
            v39 = 2
            v20 += (v20 + (^(v20)))
        } else {
            var v68 float64 = 1.5
            var v69 string = "hello"
            _ = v68
            _ = v69
        }
    }
    var v70 uint64 = 2
    v71 := v4[v20]
    if (uint32(v70) <= (v20 - v20)) {
        v1 = v1
        var v72 float32 = (v8 + ((v8 / 1.0) * v8))
        _ = v72
    }
    var v73 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v70
    _ = v71
    _ = v73
    _ = v8
    _ = v9
}
