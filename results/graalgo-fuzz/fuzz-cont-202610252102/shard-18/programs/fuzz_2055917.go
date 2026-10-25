package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 bool }
type GenS2 struct { F0 uint8; F1 uint64; F2 int }
type GenS3 struct { F0 int32; F1 int; F2 int16 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 uint; F1 uint32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint32; F1 uint; F2 int8 }
type GenS10 struct { F0 uint16 }
func genFn0(p0 uint32, p1 int8) int8 {
    fmt.Println(p1, p1)
    var v7 int64 = -9223372036854775808
    switch "hello" {
    case "":
        for fi8 := 0; fi8 < 2; fi8++ {
            switch (fi8 - fi8) {
            case -855:
                var v9 int = (fi8 - (fi8 % ((fi8) | 1)))
                _ = v9
            case 0:
                p1 -= int8(p0)
            }
            p1 += p1
            if true {
                v10 := []int{(fi8 % (((fi8 >> ((57) & 63))) | 1)), -2147483648}
                _ = v10
            }
            _ = fi8
        }
    case "\n\t":
        var v11 int16 = 0
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    return (p1 & (p1 ^ p1))
}
func genFn1() int {
    for fi12 := 0; fi12 < 10; fi12++ {
        if false {
            fmt.Println(fi12)
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 int32 = 2147483647
                _ = fi13
                _ = v14
            }
        }
        var v15 uint64 = 994
        for fi16 := 0; fi16 < 1; fi16++ {
            v15 += v15
            _ = fi16
        }
        _ = fi12
        _ = v15
    }
    var v17 int = 4294967295
    _ = v17
    return v17
}

func main() {
    var v1 uint32 = 392
    v2 := []string{(("foo bar" + "hello") + "foo bar"), "foo bar", ""}
    var v3 int32 = 0
    var v4 int32 = 949
    v5 := make(map[int]uint32)
    var v6 int = 127
    _ = len(v5)
    if ((v1 >> ((542) & 63)) < 65535) {
        switch ("hello" + (("foo bar" + "\n\t") + "foo bar")) {
        case "a":
            for fi18 := 0; fi18 < 14; fi18++ {
                for fi19 := 0; fi19 < 10; fi19++ {
                    if ((1.5 - (float32(1.5) + 3.14159)) <= (float32((3.14159 * (-2.25 * 1e10))) + 0.0)) {
                        v20 := GenS1{F0: v4, F1: 255, F2: false}
                        v21 := GenS1{F0: v3, F1: 0, F2: (4294967295 < 888)}
                        v3 += v4
                        _ = v20
                        _ = v21
                    } else {
                        v1 = (v1 % ((v1) | 1))
                        if false {
                            fmt.Println(v6, fi19)
                            if (float32(1.5) != 0.0) {
                                var v22 int8 = -1
                                fmt.Println(v3, v3)
                                _ = v22
                            } else {
                                v1 -= (v1 << ((531) & 63))
                                var v23 uint16 = 42
                                fmt.Println(v23)
                                _ = v23
                            }
                            v6 -= -398
                            switch 1 {
                            case 9223372036854775807:
                                delete(v5, -628)
                            case 32767:
                                var v24 int16 = 0
                                var v25 uint16 = 1
                                switch fi19 {
                                case -9223372036854775808:
                                    if (v25 < (v25 * (v25 % ((v25) | 1)))) {
                                        fmt.Println(v6)
                                    }
                                case -2147483648:
                                    if (2 <= 57) {
                                        v1 = (^(v1))
                                    }
                                default:
                                    v26 := GenS2{F0: uint8(v25), F1: 18446744073709551615, F2: 42}
                                    _ = v26
                                }
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                                if false {
                                    fmt.Println(v1, v3)
                                }
                                _ = v24
                                _ = v25
                            case -737:
                                var v27 int8 = 42
                                _ = v27
                            default:
                                v28 := GenS2{F0: 141, F1: 18446744073709551615, F2: (-(fi19))}
                                _ = v28
                            }
                            if false {
                                var v29 string = "foo bar"
                                v6 -= v6
                                var v30 uint16 = 255
                                _ = v29
                                _ = v30
                            } else {
                                if true {
                                    var v31 float32 = -2.25
                                    var v32 int32 = v4
                                    if false {
                                        var v33 uint64 = 2
                                        v5[(v6 + v6)] = v1
                                        _ = v33
                                    }
                                    if false {
                                        v34 := GenS3{F0: (v32 | (v3 + v3)), F1: -2147483648, F2: -868}
                                        _ = v34
                                    } else {
                                        v32 = int32(v1)
                                    }
                                    _ = v31
                                    _ = v32
                                } else {
                                    v4 = v4
                                    delete(v5, fi18)
                                    if (!(true || (!false))) {
                                        v3 += v4
                                    } else {
                                        fmt.Println(fi19, v1)
                                    }
                                }
                                v1 = 0
                                v35 := GenS4{F0: ((("hello" + ("a" + "hello")) + "foo bar") + ("hello" + "\n\t")), F1: "\n\t"}
                                if (1e10 <= float32(0.1)) {
                                    v36 := GenS4{F0: "hello", F1: (((("\n\t" + "foo bar") + ("foo bar" + "")) + ("\n\t" + "")) + "")}
                                    var v37 uint8 = 157
                                    _ = v36
                                    _ = v37
                                }
                                var v38 uint8 = 42
                                v1 = (v1 + 2147483648)
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                                var v39 int = fi19
                                _ = v35
                                _ = v38
                                _ = v39
                            }
                        } else {
                            switch ("foo bar" + ("a" + "\n\t")) {
                            case "\n\t":
                                switch ("" + "foo bar") {
                                case "a":
                                    switch v6 {
                                    case 901:
                                        v6 -= v6
                                    case 987:
                                        switch (fi18 >> ((1) & 63)) {
                                        case 4294967295:
                                            var v40 uint8 = 2
                                            _ = v40
                                        case -979:
                                            v3 = (v4 / ((v4) | 1))
                                        default:
                                            v41 := GenS5{F0: 127}
                                            if true {
                                                fmt.Println(v4, v4, fi19)
                                                var v42 uint = 2
                                                _ = v42
                                            }
                                            _ = v41
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(v1, fi19, v3)
                                    }
                                case "\n\t":
                                    fmt.Println(v4, fi18)
                                }
                            case "a":
                                if false {
                                    if (1 < 0) {
                                        var v43 int8 = -128
                                        _ = v43
                                    } else {
                                        v44 := GenS2{F0: 128, F1: 752, F2: fi19}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v44
                                    }
                                    v6 -= fi19
                                    switch 128 {
                                    case 128:
                                        var v45 uint16 = 65535
                                        _ = v45
                                    default:
                                        v5[32767] = 42
                                    }
                                } else {
                                    if (!(4294967295 != 0)) {
                                        v5[fi19] = (v1 >> ((42) & 63))
                                        v1 = 920
                                    } else {
                                        var v46 int16 = 0
                                        v3 -= (^(v4))
                                        _ = v46
                                    }
                                    v47 := GenS6{F0: 32768, F1: (fi18 / (((fi18 % (((fi19 + genFn1())) | 1))) | 1))}
                                    _ = v47
                                }
                            case "":
                                v48 := GenS6{F0: genFn1(), F1: -9223372036854775808}
                                _ = v48
                            default:
                                v49 := GenS7{F0: 1, F1: (v1 / ((v1) | 1))}
                                _ = v49
                            }
                            v4 += (v4 + (v4 * v4))
                        }
                        if false {
                            v3 -= v3
                        }
                    }
                    fmt.Println(fi19, fi18)
                    _ = fi19
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if false {
                    _ = len(v5)
                    _ = len(v2)
                    if false {
                        if true {
                            switch v6 {
                            case 0:
                                var v50 int32 = v4
                                _ = v50
                            case 9223372036854775807:
                                fmt.Println(v1)
                            case 128:
                                v6 = fi18
                            }
                            v51 := GenS7{F0: 1, F1: v1}
                            _ = v51
                        }
                        _ = len(v2)
                    }
                } else {
                    var v52 uint = 0
                    v53 := GenS8{F0: 844}
                    if (65535 >= 0) {
                        var v54 int8 = -1
                        var v55 string = "hello"
                        _ = v54
                        _ = v55
                    }
                    _ = v52
                    _ = v53
                }
                _ = fi18
            }
        case "":
            delete(v5, v6)
        default:
            if true {
                var v56 uint8 = 75
                fmt.Println(v56)
                _ = v56
            } else {
                var v57 string = (((("hello" + "") + "") + "") + (("\n\t" + ("" + "\n\t")) + (("foo bar" + "foo bar") + "\n\t")))
                v4 -= (v3 << ((42) & 63))
                _ = v57
            }
        }
        v1 = (v1 << ((575) & 63))
    }
    delete(v5, (v6 << ((0) & 63)))
    switch "foo bar" {
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    default:
        v5[(v6 ^ -2147483648)] = v1
    }
    v1 += v1
    switch ("\n\t" + "a") {
    case "hello":
        v5[v6] = 2147483648
    case "\n\t":
        var v58 int8 = -128
        _ = v58
    case "foo bar":
        delete(v5, (v6 >> ((0) & 63)))
    }
    var v59 uint = 934
    v60 := GenS9{F0: v1, F1: (v59 & (v59 + (v59 % ((42) | 1)))), F2: 74}
    switch -9223372036854775808 {
    case 4294967295:
        var v61 int32 = v4
        _ = v61
    case -764:
        v1 = v1
    case 32768:
        switch -316 {
        case 0:
            v3 += v3
        }
    default:
        _ = len(v5)
    }
    fmt.Println(v4, v59, v4)
    for fi62 := 0; fi62 < 4; fi62++ {
        v63 := GenS7{F0: 303, F1: 336}
        _ = fi62
        _ = v63
    }
    var v64 uint16 = 0
    var v65 float32 = (1e10 / 1.0)
    if false {
        v66 := GenS7{F0: 2, F1: v1}
        v67 := GenS4{F0: "foo bar", F1: "\n\t"}
        if (v6 >= -9223372036854775808) {
            _ = len(v5)
        } else {
            v59 = v59
        }
        _ = v66
        _ = v67
    } else {
        for fi68 := 0; fi68 < 1; fi68++ {
            v69 := GenS10{F0: (v64 % ((v64) | 1))}
            _ = fi68
            _ = v69
        }
    }
    switch (v6 ^ (v6 + (v6 / (((v6 - (^(v6)))) | 1)))) {
    case 127:
        if true {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            fmt.Println(v1, v4)
        }
    case 4294967295:
        v64 -= v64
    }
    var v70 uint = v59
    v71 := GenS1{F0: -2147483648, F1: uint16(v1), F2: (true && (!true))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v64
    _ = v65
    _ = v70
    _ = v71
}
