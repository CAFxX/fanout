package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
type GenS2 struct { F0 int64; F1 int16; F2 uint64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint, p1 uint) uint {
    var v4 int64 = 1
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := []int{4294967295, (fi5 << (((p1 | (^(p0)))) & 63))}
        _ = fi5
        _ = v6
    }
    if true {
        switch 128 {
        case 32767:
            var v7 bool = true
            _ = v7
        case 204:
            v8 := []int{32768, 9223372036854775807}
            _ = v8
        }
        if false {
            v9 := GenS1{F0: 127, F1: "", F2: 0}
            var v10 int8 = -37
            v11 := GenS2{F0: (^(v4)), F1: 1000, F2: 4294967295}
            _ = v10
            _ = v11
            _ = v9
        } else {
            for fi12 := 0; fi12 < 3; fi12++ {
                v4 -= (v4 ^ (-(v4)))
                p0 = 65535
                _ = fi12
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 float64 = -2.25
                v4 += 0
                _ = fi13
                _ = v14
            }
        }
        var v15 uint64 = 42
        _ = v15
    }
    if true {
        fmt.Println(p0, p0, p0)
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(p0)
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v4
    return p1
}
func genFn1() uint8 {
    var v17 uint = 401
    var v18 int16 = 1
    v17 -= genFn0((v17 - v17), (^(v17)))
    var v19 int8 = 100
    _ = v17
    _ = v18
    _ = v19
    return 127
}

func main() {
    var v1 bool = (2 < 255)
    var v2 bool = (-21 != 498)
    v3 := []string{"", ((("" + ("hello" + "foo bar")) + "a") + "hello"), "foo bar", "hello"}
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v1 = (2 <= 65535)
    v20 := GenS3{F0: 42, F1: 721}
    var v21 uint64 = 2
    v3 = append(v3, ("a" + ("" + "foo bar")))
    if (948 >= 2) {
        v1 = (255 > 255)
        var v22 int = 2147483647
        _ = v22
    }
    var v23 int32 = 1
    if (88 <= -98) {
        switch 42 {
        case -2147483648:
            v23 = (v23 + 2147483647)
        case 1:
            var v24 uint16 = 942
            _ = v24
        case 127:
            var v25 int16 = int16(v23)
            _ = v25
        default:
            switch (("\n\t" + ("foo bar" + ("" + "\n\t"))) + "\n\t") {
            case "a":
                var v26 int64 = 9223372036854775807
                _ = v26
            case "":
                v1 = (2147483648 < 0)
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        v2 = v2
    } else {
        switch 4294967295 {
        case 32768:
            var v27 int8 = 127
            _ = v27
        case -2147483648:
            var v28 uint64 = (v21 & 0)
            _ = v28
        default:
            fmt.Println(v23, v2)
        }
    }
    if true {
        if (v2 || (v23 >= v23)) {
            _ = len(v3)
        }
        v29 := GenS3{F0: 1, F1: 0}
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = v29
    } else {
        _ = len(v3)
        switch "hello" {
        case "hello":
            for fi30 := 0; fi30 < 1; fi30++ {
                if (389 <= 255) {
                    for fi31 := 0; fi31 < 17; fi31++ {
                        v2 = ((fi30 / ((fi31) | 1)) <= 2147483647)
                        _ = fi31
                    }
                    fmt.Println(v21)
                    for fi32 := 0; fi32 < 2; fi32++ {
                        v33 := GenS1{F0: 127, F1: "\n\t", F2: 18446744073709551615}
                        if true {
                            if true {
                                if true {
                                    v1 = (960 >= 32767)
                                    if len(v3) > 0 {
                                        v3[0] = ("hello" + ("hello" + ""))
                                    }
                                    var v34 float32 = -2.25
                                    switch ("a" + "\n\t") {
                                    case "":
                                        var v35 uint32 = 4294967295
                                        var v36 uint = 65535
                                        if false {
                                            v21 += v21
                                            var v37 int64 = 42
                                            v36 = v36
                                            _ = v37
                                        }
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = v34
                                } else {
                                    v21 += 42
                                    var v38 uint64 = 4294967295
                                    _ = v38
                                }
                            }
                        } else {
                            if ((!(!((v21 >> ((42) & 63)) > 18446744073709551615))) || (float64((1.5 - ((3.14159 * 1e10) / -2.25))) > ((1e10 + (1.5 + (float64(0.1) * (3.14159 + (0.0 / 1.0))))) - 0.0))) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v39 string = ("hello" + "\n\t")
                                var v40 uint16 = 42
                                fmt.Println(v21, v23)
                                _ = v39
                                _ = v40
                            }
                            if v1 {
                                if v1 {
                                    if (255 >= 2) {
                                        var v41 uint32 = 2147483647
                                        if (2147483647 >= v41) {
                                            fmt.Println(v23, v21, v41)
                                            var v42 float32 = 1.5
                                            switch "\n\t" {
                                            case "foo bar":
                                                if (v1 || ((v2 || (!(-748 == -1))) || ((415 >= 1) && (42 > 1)))) {
                                                    var v43 int16 = 0
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v43
                                                }
                                            case "\n\t":
                                                if v2 {
                                                    v2 = v2
                                                    var v44 string = ("\n\t" + ("hello" + ""))
                                                    _ = v44
                                                } else {
                                                    v1 = (882 == 4294967295)
                                                    var v45 uint64 = (v21 | v21)
                                                    _ = v45
                                                }
                                            case "hello":
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v46 uint32 = 2
                                                v1 = (!(true || (!v2)))
                                                fmt.Println(v2)
                                                var v47 int32 = int32(v41)
                                                _ = v46
                                                _ = v47
                                            }
                                            _ = v42
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                        _ = v41
                                    } else {
                                        v2 = ((!(!v1)) && (false || (0.0 != (1e10 - (-2.25 - 0.1)))))
                                    }
                                    v48 := GenS1{F0: 233, F1: "hello", F2: 42}
                                    _ = v48
                                }
                                if (v2 && (149 <= 225)) {
                                    v23 = (v23 / (((v23 & (v23 / ((42) | 1)))) | 1))
                                    v49 := GenS1{F0: 1, F1: "a", F2: 0}
                                    _ = v49
                                } else {
                                    var v50 int = (fi32 & 127)
                                    if (!true) {
                                        _ = len(v3)
                                        v1 = v1
                                        _ = len(v3)
                                    }
                                    _ = v50
                                }
                            }
                            var v51 uint16 = 65535
                            _ = v51
                        }
                        v2 = (1 > 42)
                        _ = fi32
                        _ = v33
                    }
                }
                var v52 int = -2147483648
                v1 = (32767 < 1000)
                _ = fi30
                _ = v52
            }
        case "a":
            var v53 int64 = 305
            _ = v53
        case "\n\t":
            if (4294967295 != 9223372036854775807) {
                v3 = append(v3, (("\n\t" + "\n\t") + "\n\t"))
                v23 -= 694
            } else {
                fmt.Println(v2)
                var v54 bool = ((!(-128 == -11)) || true)
                v21 -= 156
                _ = v54
            }
        default:
            if v2 {
                v55 := GenS3{F0: 278, F1: -174}
                v2 = v2
                _ = v55
            }
        }
    }
    v56 := GenS1{F0: 128, F1: "\n\t", F2: genFn0(4294967295, 0)}
    v57 := GenS4{F0: v1}
    if true {
        v23 = (v23 >> ((genFn0(65535, 4294967295)) & 63))
        switch 32767 {
        case 32768:
            var v58 uint = 65535
            _ = v58
        }
    } else {
        var v59 uint = 867
        v1 = ((v21 * v21) > v21)
        _ = v59
    }
    v1 = v1
    var v60 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v56
    _ = v57
    _ = v60
}
