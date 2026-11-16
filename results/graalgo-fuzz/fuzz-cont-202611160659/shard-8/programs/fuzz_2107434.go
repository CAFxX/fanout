package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int16; F1 int64; F2 string }
type GenS5 struct { F0 int32; F1 int32; F2 int64 }
func genFn0(p0 int16, p1 int) uint32 {
    var v6 string = "a"
    fmt.Println(p1, p0, v6)
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 += (^(p0))
        var v8 int16 = 0
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            v6 = (v6 + v6)
            switch "" {
            case "foo bar":
                if (!(!true)) {
                    var v11 int = (fi10 - -383)
                    var v12 uint = 18446744073709551615
                    v13 := GenS1{F0: (p0 % ((p0) | 1)), F1: int16(v11), F2: 179}
                    fmt.Println(fi9)
                    v14 := []int{v11, (fi9 >> (((v12 / ((v12) | 1))) & 63))}
                    _ = v11
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    switch (p1 * (-(fi10))) {
                    case -2147483648:
                        p1 = (^(fi10))
                    default:
                        p1 -= -2147483648
                    }
                    var v15 uint = 0
                    _ = v15
                }
                switch (p1 & -1) {
                case -1:
                    var v16 int32 = 0
                    fmt.Println(v6, fi10, v16)
                    if false {
                        p1 += (fi9 + -1)
                        var v17 uint = 65535
                        _ = v17
                    }
                    var v18 uint = 947
                    fmt.Println(fi10, p1)
                    v19 := GenS1{F0: p0, F1: p0, F2: 127}
                    _ = v16
                    _ = v18
                    _ = v19
                case 32767:
                    if (true || true) {
                        v20 := GenS1{F0: p0, F1: (p0 + (-(p0))), F2: 48}
                        _ = v20
                    }
                case 127:
                    var v21 float64 = 1.5
                    _ = v21
                }
            }
            if false {
                v22 := []int{42}
                v23 := GenS1{F0: p0, F1: (^(p0)), F2: 1}
                var v24 int32 = 882
                _ = v22
                _ = v23
                _ = v24
            } else {
                var v25 bool = true
                _ = v25
            }
            var v26 uint16 = 2
            v6 = "foo bar"
            var v27 int64 = 9223372036854775807
            var v28 bool = ((true || (!false)) || true)
            _ = fi10
            _ = v26
            _ = v27
            _ = v28
        }
        v29 := []int{(fi9 | int(p0)), -1, (p1 | (fi9 & (p1 >> ((18446744073709551615) & 63)))), fi9}
        _ = fi9
        _ = v29
    }
    _ = p0
    _ = p1
    _ = v6
    return 4294967295
}
func genFn1() string {
    v30 := GenS1{F0: 290, F1: 42, F2: 2}
    var v31 uint32 = 355
    if (858 < 927) {
        v32 := GenS1{F0: -1, F1: -32768, F2: 1}
        _ = v32
    }
    v31 += genFn0(-1000, -435)
    v31 += 0
    _ = v30
    _ = v31
    return "a"
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -4
    var v3 int32 = 42
    var v4 uint8 = 0
    var v5 float32 = (-2.25 + -2.25)
    v4 = 42
    for fi33 := 0; fi33 < 20; fi33++ {
        var v34 int64 = 872
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        var v36 uint16 = 0
        _ = fi35
        _ = v36
    }
    if (291 <= -27) {
        v37 := GenS2{F0: -739, F1: 362}
        _ = v37
    } else {
        v38 := []int{}
        switch "" {
        case "foo bar":
            var v39 string = ""
            _ = v39
        case "\n\t":
            var v40 uint8 = 18
            _ = v40
        }
        var v41 int32 = (v3 % (((^(v3))) | 1))
        _ = v38
        _ = v41
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        v43 := v1[(fi42 << ((681) & 63))]
        _ = fi42
        _ = v43
    }
    fmt.Println(v4, v4)
    v44 := GenS2{F0: 9223372036854775807, F1: -227}
    for fi45 := 0; fi45 < 13; fi45++ {
        for fi46 := 0; fi46 < 7; fi46++ {
            v47 := v1[0]
            v3 = (-(v3))
            _ = fi46
            _ = v47
        }
        fmt.Println(v5)
        if false {
            v2 = v2
            fmt.Println(v2, v4)
            v2 = 117
        } else {
            for fi48 := 0; fi48 < 16; fi48++ {
                v1[(fi45 / (((fi45 | (fi48 ^ fi48))) | 1))] = 439
                if false {
                    fmt.Println(v5, v2)
                }
                v3 = (^(v3))
                _ = fi48
            }
        }
        _ = fi45
    }
    switch ("foo bar" + "a") {
    case "hello":
        var v49 int = 127
        _ = v49
    case "a":
        v4 -= v4
    case "":
        var v50 int = -689
        _ = v50
    }
    for fi51 := 0; fi51 < 8; fi51++ {
        v5 += v5
        var v52 int8 = (v2 ^ v2)
        var v53 int64 = -1
        _ = fi51
        _ = v52
        _ = v53
    }
    if (316 > 2) {
        for fi54 := 0; fi54 < 12; fi54++ {
            for fi55 := 0; fi55 < 16; fi55++ {
                v2 = 42
                _ = fi55
            }
            v56 := GenS3{F0: 0, F1: 228}
            var v57 int = 32768
            _ = fi54
            _ = v56
            _ = v57
        }
        if (!true) {
            if (false || (true && false)) {
                v58 := GenS4{F0: 32767, F1: 0, F2: ""}
                v3 += v3
                var v59 float64 = float64(v5)
                _ = v58
                _ = v59
            }
            v60 := GenS4{F0: 1, F1: 9223372036854775807, F2: "foo bar"}
            _ = v60
        }
    }
    v4 -= 2
    v5 -= v5
    _ = len(v1)
    if false {
        v4 = (v4 - (v4 % (((v4 - v4)) | 1)))
        v4 -= 0
        if false {
            var v61 uint8 = 255
            _ = v61
        }
    }
    v62 := GenS5{F0: v3, F1: 1, F2: 42}
    for fi63 := 0; fi63 < 8; fi63++ {
        if false {
            for fi64 := 0; fi64 < 18; fi64++ {
                var v65 int32 = v3
                if true {
                    if false {
                        _ = len(v1)
                    }
                    v66 := GenS3{F0: (fi64 % ((int(v2)) | 1)), F1: 2}
                    var v67 bool = true
                    var v68 float32 = v5
                    v68 += float32((((-2.25 * 0.1) - 1e10) * -2.25))
                    _ = v66
                    _ = v67
                    _ = v68
                } else {
                    if (!true) {
                        fmt.Println(fi64, v65, v65)
                        v3 += (v65 * 138)
                    }
                    v1[(fi63 / ((fi63) | 1))] = 655
                    v5 = v5
                }
                _ = fi64
                _ = v65
            }
            v69 := v1[fi63]
            v2 = int8(v3)
            _ = v69
        } else {
            var v70 uint8 = v4
            _ = v70
        }
        _ = fi63
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v62
}
