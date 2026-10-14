package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 int32; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int; F2 uint64 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16; F1 int16; F2 string }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int32 = 1
    for fi4 := 0; fi4 < 12; fi4++ {
        switch fi4 {
        case 0:
            var v5 int32 = v3
            _ = v5
        default:
            fmt.Println(v3, fi4)
        }
        var v6 uint = 1
        _ = fi4
        _ = v6
    }
    var v7 uint32 = 2147483648
    v8 := GenS1{F0: "a", F1: v3, F2: v7}
    v9 := GenS2{F0: (-(v3)), F1: 100, F2: 18446744073709551615}
    if false {
        v1["foo bar"] = v7
    } else {
        v10 := GenS3{F0: 1}
        if true {
            if false {
                v11 := GenS1{F0: ((((("foo bar" + "") + "") + "hello") + "\n\t") + ""), F1: v3, F2: (v7 + (v7 / ((v7) | 1)))}
                var v12 bool = ((895 == 366) || (42 == 110))
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 1; fi13++ {
                fmt.Println(v7)
                _ = fi13
            }
        } else {
            v3 -= v3
        }
        var v14 int = 42
        _ = v10
        _ = v14
    }
    v7 += v7
    if true {
        delete(v1, "")
        if true {
            switch "a" {
            case "foo bar":
                delete(v2, (v7 ^ 2))
            case "a":
                if true {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        v16 := GenS4{F0: 0, F1: (fi15 << ((648) & 63)), F2: 940}
                        _ = fi15
                        _ = v16
                    }
                    v2[(v7 ^ 366)] = v7
                    v3 += v3
                } else {
                    var v17 int8 = 100
                    _ = v17
                }
            }
            switch ("hello" + "a") {
            case "a":
                delete(v2, (v7 << ((42) & 63)))
            }
            if (false && (1.5 == (-2.25 * (float32(0.1) * (0.1 + 0.0))))) {
                _ = len(v2)
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    var v18 int32 = v3
                    _ = v18
                case "hello":
                    for fi19 := 0; fi19 < 6; fi19++ {
                        if (1000 == 1) {
                            var v20 int8 = 56
                            for fi21 := 0; fi21 < 17; fi21++ {
                                v20 = 42
                                if (65535 < 0) {
                                    if ((0 > 1) && (699 > 133)) {
                                        var v22 int8 = (-(v20))
                                        _ = v22
                                    } else {
                                        v23 := GenS3{F0: 2}
                                        v3 = int32(fi19)
                                        _ = v23
                                    }
                                }
                                _ = fi21
                            }
                            for fi24 := 0; fi24 < 15; fi24++ {
                                fmt.Println(v7, v3, v7)
                                if true {
                                    v20 += v20
                                    v25 := GenS5{F0: fi24, F1: 1}
                                    v26 := GenS6{F0: 1, F1: ""}
                                    _ = v25
                                    _ = v26
                                }
                                _ = fi24
                            }
                            _ = v20
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                if false {
                                    v7 += (v7 & v7)
                                    var v27 int16 = -796
                                    _ = v27
                                } else {
                                    for fi28 := 0; fi28 < 16; fi28++ {
                                        v3 -= (v3 - -1)
                                        _ = fi28
                                    }
                                    for fi29 := 0; fi29 < 8; fi29++ {
                                        v7 += (^(v7))
                                        var v30 string = ("" + (("foo bar" + (("a" + (("" + "\n\t") + "hello")) + ("foo bar" + "foo bar"))) + (("a" + "") + "hello")))
                                        v3 = -2147483648
                                        v30 = v30
                                        v2[v7] = v7
                                        if false {
                                            var v31 int8 = 100
                                            var v32 int8 = v31
                                            v32 += v31
                                            _ = v31
                                            _ = v32
                                        }
                                        v33 := []int{127, fi19}
                                        var v34 int32 = (^(v3))
                                        _ = fi29
                                        _ = v30
                                        _ = v33
                                        _ = v34
                                    }
                                }
                            case "":
                                var v35 bool = true
                                _ = v35
                            default:
                                var v36 string = (("" + "a") + (("foo bar" + "a") + ("hello" + "a")))
                                _ = v36
                            }
                        }
                        v7 = v7
                        v3 = (v3 & v3)
                        _ = fi19
                    }
                default:
                    v7 -= 2147483648
                }
                v37 := GenS6{F0: 1, F1: ""}
                _ = v37
            }
        } else {
            v38 := GenS7{F0: 506}
            _ = v38
        }
    }
    v39 := GenS8{F0: 32767, F1: 1, F2: "a"}
    if false {
        v40 := GenS9{F0: 0}
        for fi41 := 0; fi41 < 15; fi41++ {
            for fi42 := 0; fi42 < 2; fi42++ {
                v3 -= -342
                var v43 uint32 = (^(v7))
                _ = fi42
                _ = v43
            }
            _ = fi41
        }
        _ = v40
    } else {
        for fi44 := 0; fi44 < 14; fi44++ {
            for fi45 := 0; fi45 < 7; fi45++ {
                var v46 int8 = 0
                _ = fi45
                _ = v46
            }
            v47 := GenS2{F0: v3, F1: 1, F2: 0}
            _ = fi44
            _ = v47
        }
        var v48 float32 = 0.0
        _ = v48
    }
    v49 := v1["\n\t"]
    if true {
        fmt.Println(v49)
        fmt.Println(v49, v7)
        if true {
            if false {
                v50 := GenS8{F0: 467, F1: -32768, F2: "foo bar"}
                _ = v50
            } else {
                for fi51 := 0; fi51 < 2; fi51++ {
                    _ = len(v2)
                    v49 = (v49 & (v49 ^ v49))
                    _ = fi51
                }
            }
            if false {
                var v52 int16 = -32768
                v53 := v1[""]
                var v54 int8 = 112
                _ = v52
                _ = v53
                _ = v54
            }
        } else {
            var v55 uint64 = 4294967295
            var v56 int32 = (-(v3))
            _ = v55
            _ = v56
        }
    } else {
        _ = len(v2)
    }
    v57 := GenS10{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v49
    _ = v57
    _ = v7
    _ = v8
    _ = v9
}
