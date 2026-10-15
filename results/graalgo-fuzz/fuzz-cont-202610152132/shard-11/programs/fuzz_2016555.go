package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 uint64; F2 int64 }
type GenS6 struct { F0 int8; F1 int16; F2 uint16 }
type GenS7 struct { F0 uint; F1 uint64; F2 bool }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS10 struct { F0 uint32; F1 int8 }
type GenS11 struct { F0 int; F1 uint64 }
func genFn0(p0 bool) uint32 {
    v7 := GenS1{F0: 1, F1: (((("a" + "a") + "") > ("foo bar" + ("hello" + "a"))) && p0)}
    switch 653 {
    case -932:
        if (4294967295 < -757) {
            v8 := []int{2147483647, 2147483647, 42, 708}
            v9 := GenS2{F0: 42}
            v10 := GenS3{F0: 2147483647, F1: false}
            _ = v10
            _ = v8
            _ = v9
        } else {
            var v11 uint32 = 42
            _ = v11
        }
    default:
        v12 := GenS4{F0: "hello"}
        _ = v12
    }
    p0 = true
    _ = p0
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 0
    var v3 bool = ((65535 != 925) && (!false))
    v4 := []string{"", ("hello" + "a"), ("\n\t" + ("foo bar" + "\n\t"))}
    var v5 int8 = 82
    v6 := []uint32{822, 2147483648, 4294967295}
    v2 = v2
    fmt.Println(v3, v3, v3)
    if v3 {
        switch ("\n\t" + "a") {
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            if true {
                v5 = 1
            } else {
                if (0.0 != 1.5) {
                    v3 = (v3 && false)
                    switch 4294967295 {
                    case 61:
                        var v13 uint = 2
                        _ = v13
                    case 128:
                        if (576 < 0) {
                            var v14 int64 = -1
                            var v15 bool = (739 <= 350)
                            _ = v14
                            _ = v15
                        }
                    case 42:
                        v5 = v5
                    }
                    if v3 {
                        v3 = (-32768 >= 594)
                        v16 := GenS5{F0: 1, F1: (v2 | 4294967295), F2: 0}
                        _ = v16
                    }
                }
                fmt.Println(v3)
                switch 42 {
                case -9223372036854775808:
                    switch 638 {
                    case -2147483648:
                        var v17 string = ("\n\t" + ("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))))
                        _ = v17
                    }
                case 127:
                    if (!v3) {
                        if (0 <= 810) {
                            for fi18 := 0; fi18 < 5; fi18++ {
                                v5 -= -1
                                _ = len(v1)
                                v5 = v5
                                _ = fi18
                            }
                            v1[2147483647] = 1
                        }
                    } else {
                        if (-2.25 > (0.0 / 0.1)) {
                            v19 := GenS5{F0: 0, F1: 4294967295, F2: -9223372036854775808}
                            var v20 bool = (((1.5 - 1.5) * ((3.14159 + 1.5) - -2.25)) >= 0.1)
                            switch "a" {
                            case "":
                                var v21 uint16 = 507
                                _ = v21
                            case "foo bar":
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    for fi23 := 0; fi23 < 17; fi23++ {
                                        switch "foo bar" {
                                        case "":
                                            if true {
                                                switch -1 {
                                                case 32768:
                                                    var v24 string = "foo bar"
                                                    _ = v24
                                                case 32767:
                                                    v25 := GenS3{F0: 1, F1: v3}
                                                    _ = v25
                                                }
                                                if (0.0 > (1e10 * 1e10)) {
                                                    var v26 int64 = 0
                                                    _ = v26
                                                } else {
                                                    var v27 int16 = -32768
                                                    switch (fi23 % ((42) | 1)) {
                                                    case -614:
                                                        v20 = (!v3)
                                                    }
                                                    v28 := GenS6{F0: (v5 - -1), F1: 122, F2: 2}
                                                    _ = v27
                                                    _ = v28
                                                }
                                            }
                                        case "\n\t":
                                            v2 = 42
                                            var v29 int32 = 1
                                            _ = v29
                                        }
                                        if true {
                                            fmt.Println(fi22, v20)
                                        } else {
                                            v30 := GenS7{F0: 2, F1: (^(v2)), F2: v3}
                                            v2 = 4294967295
                                            _ = v30
                                        }
                                        _ = fi23
                                    }
                                    for fi31 := 0; fi31 < 16; fi31++ {
                                        fmt.Println(v5, v3)
                                        _ = fi31
                                    }
                                    var v32 int = (fi22 + 474)
                                    _ = fi22
                                    _ = v32
                                }
                            default:
                                var v33 uint8 = 19
                                _ = v33
                            }
                            _ = v19
                            _ = v20
                        } else {
                            fmt.Println(v2)
                        }
                        for fi34 := 0; fi34 < 9; fi34++ {
                            switch "foo bar" {
                            case "":
                                fmt.Println(v5)
                            }
                            _ = fi34
                        }
                    }
                case 128:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    if ((((v3 || (1 == 0)) && (42 != 0)) || v3) && (!v3)) {
                        var v35 bool = v3
                        v36 := GenS5{F0: -2147483648, F1: v2, F2: -9223372036854775808}
                        v37 := GenS8{F0: 738, F1: (1.5 > 0.0)}
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                }
            }
        }
        var v38 uint32 = 2
        _ = v38
    } else {
        if ((522 != 42) || false) {
            fmt.Println(v5, v2)
        }
        for fi39 := 0; fi39 < 5; fi39++ {
            fmt.Println(v2)
            v2 -= (v2 >> ((0) & 63))
            _ = fi39
        }
        if v3 {
            v40 := v1[212]
            var v41 uint8 = 163
            v42 := GenS6{F0: (v5 / ((v5) | 1)), F1: -1, F2: 65535}
            _ = v40
            _ = v41
            _ = v42
        } else {
            _ = len(v1)
            v2 = v2
            v5 = v5
        }
    }
    if true {
        v3 = (((((((((("a" + "foo bar") + "") + "hello") + "foo bar") + "foo bar") + "foo bar") + "a") + "\n\t") + ((("foo bar" + "foo bar") + ("hello" + "hello")) + "foo bar")) == (("foo bar" + ("foo bar" + "hello")) + "\n\t"))
    }
    v5 += v5
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v43 int = -414
    if v3 {
        v43 = v43
    }
    if (v3 && (0 != 1)) {
        v43 -= (v43 | (v43 ^ v43))
        switch (-(v43)) {
        case -9223372036854775808:
            var v44 bool = (!(-1 < (v5 >> ((65535) & 63))))
            _ = v44
        default:
            if (v3 && ((-2.25 - (1.5 / (1e10 / 3.14159))) >= 1.5)) {
                v45 := GenS5{F0: 0, F1: 0, F2: -226}
                _ = v45
            } else {
                var v46 uint = 42
                if false {
                    for fi47 := 0; fi47 < 19; fi47++ {
                        v43 += fi47
                        fmt.Println(v46)
                        v6 = append(v6, 614)
                        _ = fi47
                    }
                    var v48 string = ("" + "a")
                    v49 := GenS9{F0: 65535, F1: -9223372036854775808, F2: 42}
                    _ = v48
                    _ = v49
                }
                _ = v46
            }
        }
    }
    if v3 {
        switch "foo bar" {
        case "foo bar":
            if false {
                v50 := v1[(v43 % ((32767) | 1))]
                _ = v50
            }
        }
        v51 := GenS7{F0: 255, F1: (^(v2)), F2: (!((!(702 < 891)) && false))}
        v52 := GenS10{F0: 2147483648, F1: v5}
        _ = v51
        _ = v52
    } else {
        if true {
            if len(v6) > 0 {
                v6[0] = 2147483647
            }
            for fi53 := 0; fi53 < 14; fi53++ {
                v54 := GenS10{F0: 4294967295, F1: 0}
                if (true && (!v3)) {
                    switch (fi53 >> ((255) & 63)) {
                    case 4294967295:
                        v5 = (v5 ^ int8(v2))
                    case 0:
                        for fi55 := 0; fi55 < 19; fi55++ {
                            v56 := GenS11{F0: fi53, F1: v2}
                            v2 = v2
                            fmt.Println(fi55, fi53)
                            _ = fi55
                            _ = v56
                        }
                    case 2147483647:
                        var v57 uint = 1
                        _ = v57
                    }
                }
                _ = fi53
                _ = v54
            }
            fmt.Println(v5)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
