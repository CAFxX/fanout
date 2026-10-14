package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 string; F1 uint64; F2 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8; F1 bool }
type GenS9 struct { F0 uint64; F1 int8 }
func genFn0(p0 int16) uint8 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 13; fi4++ {
        switch (-(fi4)) {
        case 0:
            switch (fi4 * (fi4 * fi4)) {
            case -208:
                var v5 uint64 = 1
                _ = v5
            case 128:
                var v6 int16 = (p0 & int16(fi4))
                _ = v6
            }
        case 997:
            var v7 bool = ("a" > "\n\t")
            _ = v7
        case 128:
            v8 := GenS1{F0: 42, F1: 2, F2: 565}
            _ = v8
        }
        if (0 < 809) {
            var v9 int64 = 570
            if (0 != 513) {
                if false {
                    v10 := GenS1{F0: 146, F1: 127, F2: (v9 % (((-(v9))) | 1))}
                    v9 = 0
                    _ = v10
                } else {
                    v11 := []int{}
                    _ = v11
                }
            } else {
                if (255 >= 255) {
                    v12 := []int{(-(fi4)), 0}
                    if len(v12) > 0 {
                        v12[0] = 32768
                    }
                    fmt.Println(v9, v9)
                    _ = v12
                } else {
                    var v13 uint = 18446744073709551615
                    var v14 int64 = v9
                    v15 := GenS2{F0: -2147483648}
                    _ = v13
                    _ = v14
                    _ = v15
                }
                p0 -= -274
            }
            v16 := GenS3{F0: 856}
            _ = v16
            _ = v9
        } else {
            fmt.Println(p0, p0)
            p0 = 32767
            var v17 bool = false
            _ = v17
        }
        _ = fi4
    }
    _ = p0
    return 42
}
func genFn1() uint32 {
    v18 := GenS1{F0: 42, F1: genFn0(-123), F2: -930}
    switch "a" {
    case "foo bar":
        var v19 int8 = -1
        _ = v19
    default:
        var v20 float32 = 0.1
        _ = v20
    }
    var v21 uint64 = 18446744073709551615
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 uint64 = v21
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: -864, F1: 0, F2: 9223372036854775807}
    _ = v18
    _ = v21
    _ = v24
    return 965
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    v3 := []int{32768, 32768, -276, 9223372036854775807}
    for fi25 := 0; fi25 < 11; fi25++ {
        v26 := GenS1{F0: 1, F1: 42, F2: 42}
        _ = fi25
        _ = v26
    }
    var v27 uint8 = 42
    fmt.Println(v27)
    var v28 uint8 = (v27 ^ 204)
    v29 := GenS4{F0: 2147483647, F1: 931, F2: 2147483647}
    v30 := GenS5{F0: -86, F1: 42}
    for fi31 := 0; fi31 < 6; fi31++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        _ = len(v1)
        _ = fi32
    }
    if (!true) {
        for fi33 := 0; fi33 < 13; fi33++ {
            if len(v3) > 0 {
                v3[0] = (fi33 | -657)
            }
            switch 128 {
            case 548:
                _ = len(v1)
            case 0:
                v34 := GenS6{F0: "hello", F1: 286, F2: 322}
                _ = v34
            case 4294967295:
                v28 = v27
            default:
                v28 = v28
            }
            v35 := v2[528]
            _ = fi33
            _ = v35
        }
        var v36 string = (("" + "foo bar") + ("\n\t" + ""))
        _ = v36
    }
    if (float32(-2.25) <= -2.25) {
        v28 = (v27 - v27)
    } else {
        _ = len(v3)
    }
    fmt.Println(v28, v27)
    var v37 int8 = 100
    if true {
        for fi38 := 0; fi38 < 17; fi38++ {
            for fi39 := 0; fi39 < 15; fi39++ {
                switch -1 {
                case 32767:
                    v40 := GenS4{F0: 42, F1: 491, F2: 40}
                    v28 = v27
                    _ = v40
                case -9223372036854775808:
                    v41 := v2[42]
                    _ = v41
                case -1:
                    if len(v3) > 0 {
                        v3[0] = (fi39 & fi38)
                    }
                }
                var v42 int8 = 0
                _ = fi39
                _ = v42
            }
            if (1 <= 32767) {
                switch "\n\t" {
                case "":
                    v43 := GenS7{F0: 796}
                    _ = v43
                }
                if (232 == 0) {
                    fmt.Println(fi38, fi38)
                    if false {
                        var v44 string = (("" + ("\n\t" + "\n\t")) + "\n\t")
                        _ = v44
                    }
                } else {
                    v28 += genFn0(1)
                    v28 = 13
                    v28 += (v28 ^ 0)
                }
                v28 = v27
            } else {
                fmt.Println(v27, v28)
                var v45 uint64 = 121
                _ = v45
            }
            v27 = genFn0(775)
            _ = fi38
        }
        v2[154] = 32768
    }
    if ((499 < 1) && (genFn1() < genFn1())) {
        var v46 uint = 65535
        for fi47 := 0; fi47 < 8; fi47++ {
            _ = len(v1)
            v48 := GenS4{F0: -1, F1: 42, F2: 2147483647}
            if ("\n\t" >= "foo bar") {
                _ = len(v3)
                for fi49 := 0; fi49 < 13; fi49++ {
                    if true {
                        v3 = append(v3, -951)
                        v37 -= v37
                    }
                    v50 := GenS8{F0: v27, F1: false}
                    var v51 uint64 = 0
                    v52 := GenS5{F0: 32767, F1: 0}
                    _ = fi49
                    _ = v50
                    _ = v51
                    _ = v52
                }
            } else {
                v46 = (v46 - 297)
                var v53 float64 = float64(0.0)
                _ = v53
            }
            _ = fi47
            _ = v48
        }
        v54 := GenS9{F0: 42, F1: 42}
        _ = v46
        _ = v54
    }
    var v55 bool = ((!(false || true)) || (2147483647 == genFn1()))
    var v56 uint = 702
    v56 -= 171
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v55
    _ = v56
}
