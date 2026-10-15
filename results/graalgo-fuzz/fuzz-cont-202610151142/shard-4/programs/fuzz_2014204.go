package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 uint16 }
func genFn0() int64 {
    if true {
        switch "\n\t" {
        case "foo bar":
            if true {
                for fi7 := 0; fi7 < 19; fi7++ {
                    v8 := GenS1{F0: 32767, F1: fi7}
                    var v9 uint32 = 874
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            } else {
                for fi10 := 0; fi10 < 9; fi10++ {
                    v11 := GenS2{F0: 1, F1: -1}
                    var v12 int = fi10
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v13 uint8 = 162
                if true {
                    if false {
                        var v14 float32 = float32(0.0)
                        _ = v14
                    } else {
                        if (1e10 > (float32(0.0) / 3.14159)) {
                            var v15 int16 = 827
                            _ = v15
                        }
                        v13 = (v13 / ((v13) | 1))
                    }
                }
                _ = v13
            }
        case "a":
            if false {
                var v16 string = ""
                _ = v16
            } else {
                var v17 int32 = -151
                _ = v17
            }
        default:
            var v18 bool = (2147483647 > 2147483648)
            _ = v18
        }
        var v19 uint32 = 42
        _ = v19
    } else {
        if ((-32768 == 1) && (!true)) {
            if true {
                v20 := GenS3{F0: 10}
                _ = v20
            }
        } else {
            v21 := GenS2{F0: 20, F1: 405}
            var v22 uint32 = 65535
            if true {
                var v23 uint8 = 97
                v22 += 614
                v23 = 200
                _ = v23
            } else {
                v22 = v22
                var v24 int32 = 1
                var v25 uint64 = 18446744073709551615
                _ = v24
                _ = v25
            }
            _ = v21
            _ = v22
        }
        if true {
            if false {
                if (-189 != 2147483647) {
                    v26 := []int{}
                    var v27 int16 = -910
                    _ = v26
                    _ = v27
                } else {
                    var v28 int = 127
                    _ = v28
                }
                var v29 uint32 = 1
                _ = v29
            } else {
                var v30 int16 = 1
                v30 -= 0
                for fi31 := 0; fi31 < 7; fi31++ {
                    v32 := []int{(^(fi31)), (fi31 * -504), (fi31 & (fi31 + (fi31 >> ((42) & 63)))), (fi31 * 41)}
                    _ = fi31
                    _ = v32
                }
                _ = v30
            }
            for fi33 := 0; fi33 < 19; fi33++ {
                v34 := []int{fi33}
                v35 := GenS4{F0: fi33, F1: 65535}
                _ = fi33
                _ = v34
                _ = v35
            }
            var v36 int32 = -327
            _ = v36
        } else {
            var v37 int16 = -32768
            if true {
                var v38 uint32 = 985
                if true {
                    for fi39 := 0; fi39 < 1; fi39++ {
                        fmt.Println(v38, v38)
                        _ = fi39
                    }
                    var v40 uint = uint(v38)
                    _ = v40
                } else {
                    var v41 uint = 1
                    v42 := []int{42, 128}
                    _ = v41
                    _ = v42
                }
                _ = v38
            } else {
                v37 -= (v37 / ((v37) | 1))
                if true {
                    fmt.Println(v37)
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        v37 += (v37 | (v37 << ((4294967295) & 63)))
                    case "a":
                        v43 := GenS2{F0: 115, F1: 513}
                        _ = v43
                    case "\n\t":
                        v37 = -32768
                    default:
                        if true {
                            v37 += (^(v37))
                            var v44 int64 = -9223372036854775808
                            _ = v44
                        } else {
                            v45 := []int{-9223372036854775808, 2147483647, 1}
                            var v46 uint16 = 255
                            for fi47 := 0; fi47 < 8; fi47++ {
                                var v48 uint8 = 128
                                _ = len(v45)
                                v37 -= (v37 >> ((48) & 63))
                                _ = fi47
                                _ = v48
                            }
                            _ = v45
                            _ = v46
                        }
                    }
                    v37 = 42
                }
            }
            v49 := []int{871, -54, 32767, 728}
            _ = v37
            _ = v49
        }
    }
    var v50 uint64 = 219
    _ = v50
    return 9223372036854775807
}
func genFn1(p0 string) int32 {
    fmt.Println(p0)
    switch "a" {
    case "\n\t":
        p0 = "foo bar"
    }
    var v51 int32 = -2147483648
    v52 := []int{}
    if true {
        if true {
            for fi53 := 0; fi53 < 3; fi53++ {
                fmt.Println(fi53, v51, v51)
                if len(v52) >= 2 {
                    v52 = v52[1:len(v52)]
                }
                switch p0 {
                case "a":
                    var v54 int16 = -438
                    _ = v54
                case "":
                    _ = len(v52)
                case "\n\t":
                    v55 := GenS3{F0: -128}
                    _ = v55
                }
                _ = fi53
            }
            var v56 int8 = -128
            _ = v56
        } else {
            _ = len(v52)
            p0 = p0
        }
    }
    _ = p0
    _ = v51
    _ = v52
    return v51
}

func main() {
    v1 := []string{"\n\t", "", ("" + "hello"), ("a" + "a")}
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 32767
    var v4 int16 = 751
    var v5 uint16 = 32767
    var v6 int16 = v4
    v57 := GenS3{F0: -19}
    if (!false) {
        v3 = 904
        for fi58 := 0; fi58 < 1; fi58++ {
            var v59 uint16 = 490
            fmt.Println(v59)
            _ = fi58
            _ = v59
        }
        for fi60 := 0; fi60 < 8; fi60++ {
            var v61 int = fi60
            v5 = (^(v5))
            for fi62 := 0; fi62 < 12; fi62++ {
                v4 = (v4 | v6)
                v3 = (v5 / ((32767) | 1))
                v61 -= fi62
                _ = fi62
            }
            _ = fi60
            _ = v61
        }
    }
    v4 = (v6 - (^(v4)))
    for fi63 := 0; fi63 < 4; fi63++ {
        var v64 uint = 4294967295
        var v65 uint16 = v3
        _ = fi63
        _ = v64
        _ = v65
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    var v66 uint8 = 1
    fmt.Println(v3)
    if (1 <= 255) {
        v5 -= (v5 ^ (v3 * (^(v5))))
        switch 32768 {
        case 32768:
            var v67 float64 = float64(-2.25)
            _ = v67
        case -1:
            v3 = v3
        default:
            for fi68 := 0; fi68 < 13; fi68++ {
                switch (^(fi68)) {
                case -910:
                    var v69 float64 = (-2.25 - 0.0)
                    _ = v69
                }
                switch ("a" + ("" + ("" + "foo bar"))) {
                case "foo bar":
                    fmt.Println(fi68, v5, v3)
                }
                v1 = append(v1, "a")
                _ = fi68
            }
        }
        v70 := GenS2{F0: 18446744073709551615, F1: -968}
        _ = v70
    }
    for fi71 := 0; fi71 < 13; fi71++ {
        switch 546 {
        case 127:
            if (true && true) {
                var v72 int8 = -128
                var v73 float64 = 3.14159
                var v74 uint16 = (v5 & (^(v5)))
                _ = v72
                _ = v73
                _ = v74
            }
        default:
            var v75 int64 = v2
            _ = v75
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch fi71 {
        case -1:
            var v76 float32 = 0.0
            _ = v76
        case 4294967295:
            v4 = 32767
        }
        _ = fi71
    }
    if true {
        var v77 int8 = -128
        _ = len(v1)
        var v78 uint32 = 2147483648
        _ = v77
        _ = v78
    }
    v3 = (^(v3))
    _ = len(v1)
    var v79 uint8 = (v66 / ((v66) | 1))
    var v80 uint16 = 557
    var v81 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v66
    _ = v79
    _ = v80
    _ = v81
}
