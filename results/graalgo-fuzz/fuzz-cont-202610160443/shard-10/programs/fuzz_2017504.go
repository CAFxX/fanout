package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int; F1 uint32; F2 int8 }
func genFn0(p0 uint8) int32 {
    fmt.Println(p0)
    var v4 int = -2147483648
    _ = p0
    _ = v4
    return -636
}
func genFn1(p0 uint32, p1 uint16, p2 uint) uint32 {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 uint64 = 42
        var v7 int8 = -117
        switch "a" {
        case "hello":
            fmt.Println(p2)
        case "":
            if (false || true) {
                var v8 int16 = 32767
                for fi9 := 0; fi9 < 2; fi9++ {
                    var v10 int64 = 0
                    v11 := GenS1{F0: (v8 | int16(v7)), F1: uint(p0)}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                if ((v7 >> ((p2) & 63)) <= (v7 ^ v7)) {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v8 += v8
                        fmt.Println(v8, p1, v7)
                        v13 := GenS1{F0: (v8 + (-(v8))), F1: 18446744073709551615}
                        _ = fi12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 16; fi14++ {
                        var v15 uint16 = p1
                        _ = fi14
                        _ = v15
                    }
                    switch 127 {
                    case 42:
                        if false {
                            var v16 int8 = v7
                            _ = v16
                        }
                    }
                } else {
                    if true {
                        switch fi5 {
                        case -702:
                            v7 -= (-(v7))
                        default:
                            var v17 bool = (p1 > 255)
                            _ = v17
                        }
                        var v18 float32 = 1e10
                        var v19 int8 = (v7 & (v7 << (((p2 >> ((255) & 63))) & 63)))
                        _ = v18
                        _ = v19
                    }
                    v20 := []int{(fi5 >> (((^(p2))) & 63)), 9223372036854775807}
                    _ = v20
                }
                _ = v8
            } else {
                var v21 int16 = -797
                switch 127 {
                case -477:
                    v6 = v6
                case -1:
                    v22 := GenS1{F0: (^(v21)), F1: 4294967295}
                    _ = v22
                case 9223372036854775807:
                    var v23 uint = 0
                    _ = v23
                }
                _ = v21
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v24 := []int{-131, 1, 947}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return p0
}

func main() {
    var v1 int16 = 1
    v2 := make(map[int]uint32)
    v3 := []int{-184, -468}
    fmt.Println(v1)
    v25 := GenS2{F0: true, F1: "foo bar"}
    v26 := GenS1{F0: (v1 >> ((18446744073709551615) & 63)), F1: 85}
    var v27 bool = false
    if (516 > 17) {
        v28 := GenS1{F0: v1, F1: 843}
        v29 := GenS1{F0: v1, F1: 65535}
        _ = v28
        _ = v29
    } else {
        for fi30 := 0; fi30 < 16; fi30++ {
            v27 = v27
            _ = fi30
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            v3 = append(v3, fi31)
            _ = fi31
        }
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        v1 = (v1 | (v1 / (((v1 & v1)) | 1)))
        v1 = v1
        if v27 {
            if v27 {
                v33 := v2[fi32]
                v33 = (v33 >> ((2) & 63))
                _ = v33
            } else {
                for fi34 := 0; fi34 < 14; fi34++ {
                    var v35 int = fi34
                    var v36 uint8 = 127
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                if v27 {
                    switch "a" {
                    case "a":
                        v27 = v27
                    case "hello":
                        var v37 string = "a"
                        _ = v37
                    case "foo bar":
                        var v38 int = fi32
                        _ = v38
                    default:
                        v39 := GenS1{F0: v1, F1: 532}
                        _ = v39
                    }
                    var v40 int = -629
                    switch fi32 {
                    case -1:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = v40
                }
            }
            if (0 == 42) {
                if v27 {
                    fmt.Println(v1, v27, v1)
                }
                var v41 uint16 = 32767
                _ = v41
            } else {
                v1 -= v1
                for fi42 := 0; fi42 < 8; fi42++ {
                    if (("hello" + ("foo bar" + "")) == "a") {
                        var v43 int64 = 344
                        fmt.Println(v43)
                        fmt.Println(fi32, fi32)
                        v3 = append(v3, (-(fi42)))
                        _ = v43
                    } else {
                        v27 = (0 > 2)
                    }
                    delete(v2, (^(fi42)))
                    if (!v27) {
                        switch (-(fi42)) {
                        case -9223372036854775808:
                            v44 := GenS2{F0: v27, F1: "hello"}
                            _ = v44
                        default:
                            var v45 uint32 = 1
                            _ = v45
                        }
                        v1 = 0
                    }
                    _ = fi42
                }
            }
        } else {
            v46 := GenS2{F0: v27, F1: (("" + "hello") + ("\n\t" + "\n\t"))}
            _ = v46
        }
        _ = fi32
    }
    var v47 string = "a"
    var v48 uint32 = 42
    v49 := GenS2{F0: false, F1: v47}
    _ = len(v2)
    if (1 < 127) {
        v50 := GenS3{F0: 1, F1: v48, F2: -54}
        _ = v50
    } else {
        if len(v3) > 0 {
            v3[0] = -263
        }
        v51 := GenS3{F0: 2147483647, F1: (v48 + v48), F2: 68}
        v52 := GenS1{F0: 42, F1: 71}
        _ = v51
        _ = v52
    }
    for fi53 := 0; fi53 < 5; fi53++ {
        v1 -= v1
        _ = fi53
    }
    v2[127] = (v48 << ((42) & 63))
    var v54 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v47
    _ = v48
    _ = v49
    _ = v54
}
