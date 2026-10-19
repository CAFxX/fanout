package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int8 }
type GenS2 struct { F0 uint16; F1 string; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int8; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint16, p1 uint) int16 {
    p0 = 42
    var v6 int64 = 42
    if false {
        p1 = p1
        p1 = p1
    }
    p1 -= 4294967295
    if true {
        fmt.Println(p0, v6, p0)
    }
    _ = p0
    _ = p1
    _ = v6
    return 751
}
func genFn1(p0 int8) bool {
    var v7 bool = (!(!false))
    fmt.Println(v7, v7)
    _ = p0
    _ = v7
    return (((!v7) && ((v7 || false) && false)) && false)
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 32767
    v4 := make(map[string]uint32)
    var v5 int32 = -246
    if false {
        v3 += (v3 % (((v3 | (v3 + (v3 >> ((255) & 63))))) | 1))
        if (128 != 1) {
            v3 = v3
            v1 = (v1 & 42)
            v8 := GenS1{F0: -947, F1: 1, F2: -43}
            _ = v8
        } else {
            fmt.Println(v1)
            var v9 uint32 = (v1 + uint32(v3))
            _ = v9
        }
        if false {
            if (true || false) {
                v5 = (v5 & v5)
            } else {
                v10 := GenS2{F0: v3, F1: "foo bar", F2: 925}
                switch int(v3) {
                case -243:
                    for fi11 := 0; fi11 < 17; fi11++ {
                        if (0 <= 80) {
                            var v12 float32 = (1e10 + 3.14159)
                            _ = v12
                        }
                        fmt.Println(fi11, v1)
                        _ = fi11
                    }
                case -292:
                    fmt.Println(v5, v3)
                case 1:
                    v5 = (v5 / (((v5 | (v5 | (v5 << ((4294967295) & 63))))) | 1))
                default:
                    fmt.Println(v5, v3)
                }
                _ = v10
            }
            v1 += (v1 / (((^(v1))) | 1))
        }
    } else {
        var v13 uint = 56
        _ = v13
    }
    fmt.Println(v1, v3)
    var v14 uint8 = 2
    v15 := GenS3{F0: v3}
    for fi16 := 0; fi16 < 5; fi16++ {
        if false {
            if (fi16 == (fi16 ^ (fi16 ^ fi16))) {
                var v17 int32 = (-(v5))
                _ = v17
            }
            fmt.Println(v14, v14, fi16)
        } else {
            fmt.Println(fi16, v14, v1)
            if false {
                for fi18 := 0; fi18 < 6; fi18++ {
                    switch (fi18 | 592) {
                    case 128:
                        var v19 int = fi18
                        _ = v19
                    default:
                        v20 := GenS4{F0: 179, F1: int8(fi18), F2: ""}
                        _ = v20
                    }
                    v21 := GenS4{F0: uint(v3), F1: 127, F2: "\n\t"}
                    _ = fi18
                    _ = v21
                }
                if (!true) {
                    switch -324 {
                    case 583:
                        for fi22 := 0; fi22 < 10; fi22++ {
                            fmt.Println(v3)
                            _ = fi22
                        }
                    }
                    var v23 int32 = (v5 + (^(v5)))
                    v24 := GenS5{F0: (v3 & 304)}
                    _ = v23
                    _ = v24
                } else {
                    v25 := []int{}
                    v26 := v2[(v1 | 210)]
                    var v27 uint8 = v14
                    _ = v25
                    _ = v26
                    _ = v27
                }
                if (!true) {
                    switch (fi16 ^ fi16) {
                    case 199:
                        v14 = v14
                    case -1:
                        v3 -= 228
                    case -465:
                        fmt.Println(v5, v3)
                    }
                    var v28 int = fi16
                    v29 := []int{(fi16 >> ((65535) & 63)), (v28 - (v28 / ((v28) | 1))), fi16}
                    _ = v28
                    _ = v29
                }
            } else {
                v1 = 96
                v14 = (v14 ^ (v14 << ((18446744073709551615) & 63)))
            }
        }
        var v30 float64 = (((((0.0 * 0.0) / 1e10) * 1e10) * 1e10) / 1.0)
        _ = fi16
        _ = v30
    }
    fmt.Println(v14, v5)
    switch 52 {
    case 32767:
        v5 -= (v5 + -736)
    case 287:
        v31 := v4[("foo bar" + "a")]
        _ = v31
    case -2147483648:
        v14 = (^(v14))
    default:
        v14 = 127
    }
    switch ("\n\t" + "foo bar") {
    case "hello":
        switch -664 {
        case 32767:
            if (1 > 127) {
                var v32 float32 = (((1e10 / 0.1) - 3.14159) * -2.25)
                v5 = v5
                v33 := GenS4{F0: 167, F1: 42, F2: ("" + ("" + "\n\t"))}
                _ = v32
                _ = v33
            }
        }
    case "":
        for fi34 := 0; fi34 < 18; fi34++ {
            for fi35 := 0; fi35 < 5; fi35++ {
                var v36 float32 = (((-2.25 - 3.14159) + ((1.5 + 1.5) + 1e10)) * -2.25)
                v37 := GenS2{F0: (v3 + 42), F1: ("hello" + ""), F2: 1}
                var v38 int32 = v5
                _ = fi35
                _ = v36
                _ = v37
                _ = v38
            }
            _ = fi34
        }
    case "\n\t":
        var v39 int32 = -2147483648
        _ = v39
    }
    if (false || false) {
        if ((v14 << ((4294967295) & 63)) < (v14 >> ((389) & 63))) {
            var v40 uint32 = 327
            var v41 string = "a"
            _ = v40
            _ = v41
        } else {
            v42 := GenS6{F0: 18446744073709551615}
            _ = v42
        }
        for fi43 := 0; fi43 < 14; fi43++ {
            fmt.Println(fi43)
            switch "foo bar" {
            case "\n\t":
                v44 := []int{127, (-(fi43))}
                _ = v44
            }
            _ = fi43
        }
    }
    v5 += (-(v5))
    for fi45 := 0; fi45 < 3; fi45++ {
        v46 := GenS6{F0: 2}
        _ = fi45
        _ = v46
    }
    v47 := []int{int(v5), int(v14), int(v1), 210}
    v14 -= (v14 << ((255) & 63))
    if len(v47) > 0 {
        v47[0] = -961
    }
    v48 := GenS5{F0: uint16(v14)}
    if (float32(1e10) == 1e10) {
        var v49 int = 9223372036854775807
        v50 := GenS4{F0: 564, F1: -128, F2: "hello"}
        v47 = append(v47, (-(v49)))
        _ = v49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
