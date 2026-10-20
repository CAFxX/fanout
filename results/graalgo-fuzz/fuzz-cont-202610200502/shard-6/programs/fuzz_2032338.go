package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int32; F2 int32 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 uint; F1 int }
func genFn0() int {
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := []int{32768, (-(fi5)), fi5}
        v7 := GenS1{F0: -911, F1: -32768}
        var v8 bool = (false && false)
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int = (fi9 << ((65535) & 63))
        _ = fi9
        _ = v10
    }
    if false {
        v11 := []int{-2147483648, -271, 981, 930}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 uint8 = 43
            if true {
                v13 -= 127
                var v14 int16 = 42
                if (!false) {
                    switch (fi12 ^ (^(fi12))) {
                    case 0:
                        if false {
                            v15 := GenS1{F0: -2147483648, F1: (v14 & -830)}
                            _ = v15
                        } else {
                            fmt.Println(v13, v14, v14)
                        }
                    case 127:
                        v14 = (v14 - (v14 | (v14 ^ (v14 / ((v14) | 1)))))
                    default:
                        fmt.Println(fi12, v13, fi12)
                    }
                    v16 := []int{fi12, fi12, fi12}
                    _ = v16
                } else {
                    var v17 uint32 = 892
                    v18 := []int{}
                    _ = v17
                    _ = v18
                }
                _ = v14
            }
            var v19 uint8 = v13
            _ = fi12
            _ = v13
            _ = v19
        }
        var v20 bool = false
        v20 = false
        _ = v20
    }
    var v21 uint32 = 71
    var v22 int16 = -886
    _ = v21
    _ = v22
    return 694
}
func genFn1(p0 string, p1 int) uint16 {
    p0 = ("a" + ("hello" + "a"))
    p0 = "a"
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 uint16 = 255
    v2 := []int{32768, 861, -9223372036854775808}
    v3 := make(map[uint32]int)
    var v4 uint64 = 492
    if len(v2) > 0 {
        v2[0] = -101
    }
    v23 := GenS1{F0: -2147483648, F1: 0}
    v4 = 0
    switch 42 {
    case -9223372036854775808:
        switch 128 {
        case 4294967295:
            v1 = v1
        }
    default:
        var v24 int8 = -60
        _ = v24
    }
    var v25 int16 = -1
    var v26 int64 = 432
    var v27 string = "foo bar"
    if (32768 <= 671) {
        for fi28 := 0; fi28 < 19; fi28++ {
            for fi29 := 0; fi29 < 13; fi29++ {
                v25 = (v25 / (((v25 & v25)) | 1))
                _ = fi29
            }
            v4 -= (v4 << ((1) & 63))
            v30 := GenS2{F0: 2}
            _ = fi28
            _ = v30
        }
    } else {
        var v31 int8 = -128
        _ = v31
    }
    if (!(-1 > 759)) {
        v32 := GenS3{F0: 5}
        _ = v32
    }
    v26 -= v26
    for fi33 := 0; fi33 < 19; fi33++ {
        v34 := v3[1]
        if ((!((!false) || false)) || true) {
            if false {
                for fi35 := 0; fi35 < 19; fi35++ {
                    v27 = "\n\t"
                    if (3.14159 < 3.14159) {
                        var v36 string = ("foo bar" + v27)
                        v26 -= int64(v4)
                        fmt.Println(fi33, fi35)
                        v26 = (v26 % ((v26) | 1))
                        var v37 float64 = (0.0 + 0.0)
                        switch (v36 + v36) {
                        case "a":
                            var v38 bool = true
                            _ = v38
                        case "hello":
                            v36 = (("\n\t" + v27) + (v36 + (v27 + ((v36 + v36) + (v27 + v36)))))
                        case "foo bar":
                            fmt.Println(fi33)
                        default:
                            v1 += (^(v1))
                        }
                        _ = v36
                        _ = v37
                    }
                    if true {
                        var v39 uint32 = 4294967295
                        v40 := GenS2{F0: 65535}
                        fmt.Println(v34)
                        v41 := GenS3{F0: -103}
                        _ = v39
                        _ = v40
                        _ = v41
                    } else {
                        delete(v3, 65535)
                        v3[1] = genFn0()
                    }
                    _ = fi35
                }
            }
            v34 = (^(v34))
            if true {
                var v42 int = (v34 + v34)
                _ = v42
            }
        }
        if (!(float32(0.0) == 1e10)) {
            v43 := GenS4{F0: 18446744073709551615, F1: -901, F2: 42}
            for fi44 := 0; fi44 < 18; fi44++ {
                fmt.Println(v34, fi33, v27)
                switch fi33 {
                case -503:
                    var v45 int = (fi44 << ((762) & 63))
                    _ = v45
                }
                _ = fi44
            }
            var v46 bool = (!false)
            _ = v43
            _ = v46
        }
        _ = fi33
        _ = v34
    }
    for fi47 := 0; fi47 < 17; fi47++ {
        v48 := GenS2{F0: 1}
        for fi49 := 0; fi49 < 20; fi49++ {
            v50 := GenS5{F0: v25, F1: 0}
            v51 := GenS5{F0: v25, F1: -2147483648}
            if (1 != v25) {
                v4 = 354
            }
            _ = fi49
            _ = v50
            _ = v51
        }
        v52 := GenS2{F0: 82}
        _ = fi47
        _ = v48
        _ = v52
    }
    if (4294967295 >= 18446744073709551615) {
        var v53 int = genFn0()
        switch 127 {
        case 127:
            var v54 string = v27
            _ = v54
        case 9223372036854775807:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            if true {
                if true {
                    var v55 int = int(v1)
                    fmt.Println(v26, v4, v26)
                    _ = v55
                } else {
                    v1 = (v1 << ((4294967295) & 63))
                }
            } else {
                for fi56 := 0; fi56 < 9; fi56++ {
                    if true {
                        v25 += 1000
                    } else {
                        var v57 uint32 = 915
                        var v58 float64 = float64(1e10)
                        _ = v57
                        _ = v58
                    }
                    _ = len(v2)
                    switch v27 {
                    case "a":
                        var v59 uint32 = 89
                        _ = v59
                    case "\n\t":
                        for fi60 := 0; fi60 < 17; fi60++ {
                            fmt.Println(v25, v1)
                            var v61 uint = 86
                            _ = fi60
                            _ = v61
                        }
                    }
                    _ = fi56
                }
            }
        }
        for fi62 := 0; fi62 < 11; fi62++ {
            if (true || false) {
                if false {
                    v27 = ((((("\n\t" + v27) + (("\n\t" + ("" + v27)) + (v27 + v27))) + "\n\t") + "hello") + ((v27 + v27) + ""))
                    v2 = append(v2, v53)
                }
                var v63 int64 = v26
                _ = v63
            } else {
                v64 := GenS6{F0: 0, F1: -690}
                _ = v64
            }
            v65 := GenS7{F0: 225, F1: (^(fi62))}
            _ = fi62
            _ = v65
        }
        _ = v53
    }
    v2 = append(v2, -9223372036854775808)
    _ = len(v3)
    v1 = genFn1("", 127)
    v25 = v25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
