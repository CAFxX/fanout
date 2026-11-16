package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 int }
func genFn0(p0 int32) int {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 string = "\n\t"
        p0 -= (p0 << ((18446744073709551615) & 63))
        switch (v8 + "a") {
        case "":
            if false {
                var v9 uint64 = 23
                v10 := GenS1{F0: 786, F1: (v8 + v8), F2: v9}
                _ = v10
                _ = v9
            } else {
                for fi11 := 0; fi11 < 1; fi11++ {
                    fmt.Println(fi11)
                    _ = fi11
                }
                p0 -= -1
                v12 := GenS1{F0: 42, F1: (("foo bar" + v8) + v8), F2: 2}
                _ = v12
            }
        case "foo bar":
            v13 := GenS1{F0: 552, F1: v8, F2: 2}
            _ = v13
        default:
            for fi14 := 0; fi14 < 7; fi14++ {
                p0 -= (p0 ^ (p0 << ((42) & 63)))
                _ = fi14
            }
        }
        _ = fi7
        _ = v8
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        switch (fi15 | fi15) {
        case -9223372036854775808:
            fmt.Println(p0, fi15)
        case 32768:
            switch int(p0) {
            case -318:
                var v16 uint = 2
                _ = v16
            case 128:
                v17 := []int{}
                _ = v17
            case 1:
                for fi18 := 0; fi18 < 18; fi18++ {
                    switch "foo bar" {
                    case "a":
                        var v19 uint8 = 1
                        _ = v19
                    case "":
                        fmt.Println(p0)
                    case "\n\t":
                        v20 := []int{(^(fi15))}
                        _ = v20
                    }
                    _ = fi18
                }
            }
        case 4294967295:
            for fi21 := 0; fi21 < 13; fi21++ {
                fmt.Println(p0, fi15, fi21)
                _ = fi21
            }
        default:
            for fi22 := 0; fi22 < 16; fi22++ {
                var v23 float64 = (((1.5 + 0.0) / (((((1.5 - 0.0) * (float64(float32(float64((0.0 - 3.14159)))) - 3.14159)) / float64((0.1 - 1.5))) - (1e10 + float64(1.5))) - ((1.5 - 3.14159) * (0.0 * 1.5)))) * 1.5)
                _ = fi22
                _ = v23
            }
        }
        v24 := []int{(fi15 & (fi15 >> ((375) & 63))), (fi15 * fi15)}
        _ = fi15
        _ = v24
    }
    _ = p0
    return 0
}
func genFn1(p0 uint, p1 bool) uint16 {
    if (!p1) {
        fmt.Println(p1)
        if ((!(!(0.0 <= 0.0))) && p1) {
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 uint16 = 2
                var v27 int8 = -1
                if p1 {
                    v28 := GenS1{F0: v26, F1: "\n\t", F2: 2}
                    _ = v28
                }
                _ = fi25
                _ = v26
                _ = v27
            }
        }
        v29 := []int{333, -9223372036854775808, 32767, 32767}
        _ = v29
    } else {
        v30 := []int{4294967295}
        for fi31 := 0; fi31 < 4; fi31++ {
            fmt.Println(p1, p0)
            if len(v30) > 0 {
                v30[0] = genFn0(-2147483648)
            }
            switch "hello" {
            case "a":
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
            case "":
                var v32 int = (fi31 >> (((p0 - p0)) & 63))
                _ = v32
            case "hello":
                for fi33 := 0; fi33 < 18; fi33++ {
                    v34 := GenS2{F0: 1, F1: 4294967295}
                    var v35 uint64 = 0
                    v35 = v35
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            }
            _ = fi31
        }
        for fi36 := 0; fi36 < 12; fi36++ {
            if len(v30) > 0 {
                v30[0] = (fi36 - 4294967295)
            }
            _ = fi36
        }
        _ = v30
    }
    v37 := GenS3{F0: 302, F1: -2147483648}
    v38 := GenS3{F0: 42, F1: 473}
    if (((1e10 * 0.0) / 1.0) != 0.0) {
        var v39 string = "hello"
        switch 127 {
        case 32767:
            v40 := GenS4{F0: 255}
            _ = v40
        case 0:
            var v41 uint = 1
            _ = v41
        case 1:
            p1 = (!false)
        default:
            for fi42 := 0; fi42 < 7; fi42++ {
                v43 := GenS2{F0: 1, F1: 820}
                _ = fi42
                _ = v43
            }
        }
        var v44 uint8 = 0
        _ = v39
        _ = v44
    } else {
        var v45 uint32 = 0
        switch -562 {
        case 2147483647:
            fmt.Println(p0, p1, v45)
        case 42:
            switch 934 {
            case 42:
                v46 := []int{}
                _ = v46
            }
        }
        _ = v45
    }
    v47 := GenS5{F0: 0, F1: 0}
    _ = p0
    _ = p1
    _ = v37
    _ = v38
    _ = v47
    return 36
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{4294967295, 461, 2147483648, 1}
    var v3 int16 = 42
    v4 := []int{604}
    v5 := make(map[int]bool)
    v6 := make(map[string]bool)
    if false {
        v48 := GenS4{F0: -9223372036854775808}
        _ = v48
    }
    v3 = (v3 * (v3 + v3))
    fmt.Println(v3)
    var v49 uint = 42
    fmt.Println(v3, v49)
    switch "hello" {
    case "\n\t":
        v50 := GenS2{F0: 1, F1: 4294967295}
        _ = v50
    case "":
        var v51 uint8 = 127
        _ = v51
    default:
        var v52 float32 = 1e10
        _ = v52
    }
    var v53 int32 = -797
    if (v53 <= v53) {
        v3 = (v3 % ((v3) | 1))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        for fi54 := 0; fi54 < 6; fi54++ {
            var v55 int16 = (v3 | 32767)
            if ((1 < 100) && (244 <= 42)) {
                if true {
                    for fi56 := 0; fi56 < 5; fi56++ {
                        v57 := GenS4{F0: -9223372036854775808}
                        v58 := GenS6{F0: (fi56 << ((18446744073709551615) & 63))}
                        if false {
                            var v59 int8 = int8(fi54)
                            _ = v59
                        }
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                    if false {
                        if true {
                            v55 = 1
                            v60 := GenS6{F0: fi54}
                            _ = v60
                        }
                        var v61 uint64 = 42
                        v49 += (^(v49))
                        _ = v61
                    }
                }
                v62 := GenS2{F0: 65535, F1: 42}
                if len(v4) > 0 {
                    v4[0] = (fi54 / ((fi54) | 1))
                }
                _ = v62
            }
            v63 := GenS1{F0: 219, F1: "foo bar", F2: 42}
            _ = fi54
            _ = v55
            _ = v63
        }
        var v64 int = 32767
        fmt.Println(v3)
        _ = v64
    }
    for fi65 := 0; fi65 < 11; fi65++ {
        var v66 int8 = 82
        _ = fi65
        _ = v66
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        v3 = int16(v53)
    }
    var v67 float32 = 3.14159
    v68 := GenS2{F0: 4294967295, F1: 185}
    switch ("\n\t" + "hello") {
    case "a":
        fmt.Println(v67, v67, v3)
    default:
        var v69 int16 = -1
        _ = v69
    }
    v70 := v1["\n\t"]
    var v71 uint16 = 65535
    var v72 uint16 = v71
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v67:", v67)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v53
    _ = v6
    _ = v67
    _ = v68
    _ = v70
    _ = v71
    _ = v72
}
