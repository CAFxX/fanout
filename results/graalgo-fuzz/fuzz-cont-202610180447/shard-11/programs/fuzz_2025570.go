package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 int8 }
func genFn0() int32 {
    v4 := GenS1{F0: 953, F1: -2147483648, F2: 127}
    var v5 string = "foo bar"
    switch 9223372036854775807 {
    case 2147483647:
        v5 = v5
    case -666:
        v5 = v5
    default:
        if (2147483647 <= 1) {
            v6 := GenS2{F0: 2147483648, F1: 0, F2: 1}
            v7 := []int{-452, -1, -1, 32767}
            _ = v6
            _ = v7
        } else {
            v8 := GenS3{F0: 1, F1: -1}
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(v5, fi9)
                if (2147483647 == 0) {
                    var v10 float32 = -2.25
                    v10 = v10
                    v11 := []int{(fi9 << ((722) & 63)), fi9}
                    _ = v10
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        v13 := GenS1{F0: 0, F1: 350, F2: 1}
                        v14 := []int{(fi12 | -342), 2147483647, fi12, (fi9 % ((fi12) | 1))}
                        v5 = (v5 + ("\n\t" + v5))
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    v5 = (((v5 + "") + "\n\t") + v5)
                    v15 := []int{fi9, (fi9 % ((-9223372036854775808) | 1)), fi9, fi9}
                    _ = v15
                }
                _ = fi9
            }
            if (75 >= 1) {
                var v16 uint64 = 2
                fmt.Println(v16, v5)
                _ = v16
            } else {
                var v17 uint = 2
                _ = v17
            }
            _ = v8
        }
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        if (false || false) {
            var v19 int64 = 1
            var v20 bool = (209 == 0)
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 4; fi21++ {
                if (v5 > v5) {
                    v5 = (v5 + "\n\t")
                    v22 := GenS2{F0: 2147483647, F1: 127, F2: 65535}
                    _ = v22
                } else {
                    v23 := []int{fi18}
                    _ = v23
                }
                var v24 uint64 = 1
                var v25 int8 = 100
                var v26 uint64 = v24
                var v27 uint16 = 1
                var v28 float32 = 1e10
                _ = fi21
                _ = v24
                _ = v25
                _ = v26
                _ = v27
                _ = v28
            }
            v5 = v5
        }
        _ = fi18
    }
    switch -9223372036854775808 {
    case 127:
        var v29 int = 32767
        _ = v29
    case 760:
        v5 = v5
    case 2147483647:
        v30 := []int{-509, 2147483647, -110}
        _ = v30
    default:
        v31 := GenS3{F0: -365, F1: -46}
        _ = v31
    }
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint16, p1 uint32) uint16 {
    var v32 string = ""
    v33 := GenS1{F0: 0, F1: 1, F2: 123}
    p1 += p1
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v32
    _ = v33
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -1
    var v3 uint8 = 42
    if true {
        v3 = v3
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        v2 = int64(fi34)
        var v35 float64 = 3.14159
        for fi36 := 0; fi36 < 18; fi36++ {
            switch (fi36 ^ fi36) {
            case -9223372036854775808:
                if (fi34 < fi36) {
                    if false {
                        v37 := GenS3{F0: v2, F1: 0}
                        if false {
                            v2 = -611
                            v3 = v3
                            if false {
                                var v38 int64 = (^(v2))
                                v3 += v3
                                _ = v38
                            }
                            v39 := GenS1{F0: 867, F1: -2147483648, F2: 0}
                            _ = v39
                        } else {
                            v40 := GenS4{F0: 2, F1: (v2 * -9223372036854775808)}
                            fmt.Println(fi36, fi34)
                            fmt.Println(v35)
                            var v41 bool = false
                            _ = v40
                            _ = v41
                        }
                        _ = v37
                    } else {
                        if (315 < 2) {
                            if true {
                                v42 := GenS5{F0: -1, F1: 1}
                                _ = v42
                            }
                            v43 := GenS6{F0: (0.0 >= -2.25), F1: (!false)}
                            _ = v43
                        }
                        v44 := v1[""]
                        _ = v44
                    }
                    if false {
                        v3 += v3
                        var v45 float64 = (v35 / 1.5)
                        v35 -= (-2.25 - 0.1)
                        v46 := GenS5{F0: (v2 | (v2 * -214)), F1: 65535}
                        v3 -= v3
                        _ = v45
                        _ = v46
                    }
                    if (false || (85 <= -993)) {
                        switch "hello" {
                        case "hello":
                            v35 += ((v35 / (v35 * 1e10)) / -2.25)
                        case "a":
                            v47 := []int{2147483647, 4294967295, (fi36 * fi36)}
                            _ = v47
                        case "foo bar":
                            fmt.Println(v35, v2, fi34)
                        default:
                            if (364 == 2147483647) {
                                v3 = v3
                                if true {
                                    v2 = v2
                                    v35 = v35
                                    v48 := GenS7{F0: fi36, F1: ("foo bar" + ("hello" + ("" + "")))}
                                    _ = v48
                                }
                            } else {
                                v35 -= 1e10
                            }
                        }
                        v49 := []int{fi34, (-(fi34))}
                        var v50 uint64 = 4294967295
                        _ = v49
                        _ = v50
                    } else {
                        v35 += v35
                        switch -1 {
                        case 32768:
                            v51 := []int{(-(fi34)), (fi34 >> ((36) & 63)), (^(fi36)), (fi36 | (-(fi36)))}
                            _ = v51
                        case 42:
                            v52 := []int{1}
                            var v53 int = 127
                            switch 32768 {
                            case -9223372036854775808:
                                v2 -= v2
                            case 0:
                                var v54 int64 = (v2 - (v2 / (((v2 % ((v2) | 1))) | 1)))
                                _ = v54
                            case 32768:
                                fmt.Println(v2)
                            }
                            _ = v52
                            _ = v53
                        default:
                            if (2 < 1) {
                                v55 := GenS6{F0: (-992 <= int16(v2)), F1: false}
                                var v56 string = "a"
                                _ = v55
                                _ = v56
                            }
                        }
                        if false {
                            v57 := GenS7{F0: (fi34 | (fi34 >> ((269) & 63))), F1: ("\n\t" + "\n\t")}
                            v58 := GenS6{F0: (true && false), F1: true}
                            _ = v57
                            _ = v58
                        }
                    }
                }
            case -577:
                v2 += int64(v3)
            case -1:
                if false {
                    var v59 int32 = 716
                    v59 -= int32(fi34)
                    _ = v59
                }
            default:
                var v60 uint16 = 717
                _ = v60
            }
            var v61 string = (("\n\t" + ("" + "\n\t")) + "hello")
            fmt.Println(v2, v2)
            v62 := GenS2{F0: 65535, F1: (v3 << ((2) & 63)), F2: 377}
            v63 := GenS1{F0: 2, F1: 42, F2: v3}
            var v64 uint64 = 4294967295
            _ = fi36
            _ = v61
            _ = v62
            _ = v63
            _ = v64
        }
        _ = fi34
        _ = v35
    }
    for fi65 := 0; fi65 < 4; fi65++ {
        var v66 float32 = ((0.1 * -2.25) * (0.0 + 0.1))
        fmt.Println(fi65, v3, v66)
        _ = fi65
        _ = v66
    }
    var v67 uint16 = 157
    v3 -= 0
    v68 := GenS8{F0: 0}
    for fi69 := 0; fi69 < 16; fi69++ {
        v2 = (v2 >> ((472) & 63))
        var v70 uint64 = 962
        _ = fi69
        _ = v70
    }
    if (-2.25 > -2.25) {
        var v71 int64 = -9223372036854775808
        v72 := v1[""]
        var v73 int8 = 0
        _ = v71
        _ = v72
        _ = v73
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v67
    _ = v68
}
