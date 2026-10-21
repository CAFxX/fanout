package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int8; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 uint8; F2 int8 }
func genFn0(p0 uint) int16 {
    for fi5 := 0; fi5 < 17; fi5++ {
        p0 = (p0 << ((18446744073709551615) & 63))
        p0 = (p0 * (p0 - (p0 >> ((p0) & 63))))
        v6 := GenS1{F0: 42, F1: 0}
        _ = fi5
        _ = v6
    }
    var v7 int16 = -32768
    var v8 int8 = -128
    _ = p0
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 32767
    var v3 uint = 0
    var v4 uint64 = 1
    v9 := GenS1{F0: 1, F1: 0}
    v10 := []int{v2}
    var v11 uint16 = 65535
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 int8 = int8(fi12)
        for fi14 := 0; fi14 < 6; fi14++ {
            fmt.Println(fi12, v11, v4)
            var v15 uint32 = 42
            v1 = ("\n\t" + v1)
            switch (v1 + (v1 + "\n\t")) {
            case "\n\t":
                v15 -= (^(v15))
            case "":
                var v16 uint64 = 18446744073709551615
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v3 += (v3 - (v3 - v3))
    switch 32768 {
    case -2147483648:
        switch v2 {
        case 4294967295:
            v17 := GenS2{F0: 111, F1: 45, F2: 0}
            _ = v17
        case 127:
            for fi18 := 0; fi18 < 18; fi18++ {
                switch fi18 {
                case -554:
                    v4 += 74
                default:
                    v19 := GenS1{F0: 2, F1: v11}
                    _ = v19
                }
                fmt.Println(v3)
                _ = fi18
            }
        case 1:
            _ = len(v10)
        }
    case 4294967295:
        if true {
            switch (v1 + v1) {
            case "foo bar":
                if false {
                    switch 9223372036854775807 {
                    case 128:
                        var v20 uint16 = (v11 + 528)
                        _ = v20
                    }
                    var v21 float64 = 1e10
                    switch int(v4) {
                    case 4294967295:
                        var v22 int64 = 773
                        _ = v22
                    case 42:
                        for fi23 := 0; fi23 < 8; fi23++ {
                            var v24 string = "a"
                            if (!(!(!true))) {
                                v3 = 255
                                var v25 uint64 = v4
                                for fi26 := 0; fi26 < 6; fi26++ {
                                    v27 := GenS1{F0: 128, F1: 2}
                                    _ = fi26
                                    _ = v27
                                }
                                _ = v25
                            }
                            v28 := GenS2{F0: 0, F1: 0, F2: 182}
                            _ = fi23
                            _ = v24
                            _ = v28
                        }
                    case 9223372036854775807:
                        v21 -= (float64((0.0 * float32(((v21 / float64(1.5)) - float64((-2.25 - float32(v21))))))) - v21)
                    default:
                        v29 := GenS1{F0: 128, F1: v11}
                        _ = v29
                    }
                    _ = v21
                }
            case "":
                fmt.Println(v1)
            }
            var v30 uint = 2
            v11 = v11
            _ = v30
        } else {
            var v31 uint32 = 2147483647
            for fi32 := 0; fi32 < 16; fi32++ {
                var v33 int64 = 1
                for fi34 := 0; fi34 < 6; fi34++ {
                    switch (fi32 - v2) {
                    case 28:
                        var v35 int64 = v33
                        _ = v35
                    case -9223372036854775808:
                        var v36 int32 = -1
                        _ = v36
                    }
                    v37 := GenS1{F0: 1, F1: (v11 + (v11 >> ((v3) & 63)))}
                    _ = fi34
                    _ = v37
                }
                v38 := GenS2{F0: 1, F1: 1, F2: 42}
                _ = fi32
                _ = v33
                _ = v38
            }
            _ = v31
        }
    }
    var v39 int8 = 127
    v40 := GenS2{F0: (^(v39)), F1: 9223372036854775807, F2: 1}
    var v41 int32 = -634
    if (false || true) {
        v1 = ""
        var v42 float64 = float64(1e10)
        var v43 float32 = ((3.14159 * 1.5) + 0.1)
        _ = v42
        _ = v43
    } else {
        if true {
            if ((((-9223372036854775808 > 1) && true) || true) || true) {
                if false {
                    fmt.Println(v3)
                    v4 -= (v4 + (v4 - (v4 - (v4 & v4))))
                }
                v44 := GenS1{F0: 1, F1: 42}
                _ = v44
            } else {
                var v45 float64 = (3.14159 - 0.1)
                var v46 bool = (-118 == v2)
                _ = v45
                _ = v46
            }
            v47 := GenS2{F0: v39, F1: -1, F2: 1}
            _ = v47
        }
        for fi48 := 0; fi48 < 16; fi48++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi48
        }
        for fi49 := 0; fi49 < 16; fi49++ {
            v3 -= (v3 & v3)
            v41 -= v41
            _ = fi49
        }
    }
    fmt.Println(v39, v2, v3)
    v50 := GenS3{F0: 4294967295, F1: (-(v2))}
    v51 := GenS1{F0: 42, F1: v11}
    fmt.Println(v39, v4)
    var v52 uint32 = 2147483647
    for fi53 := 0; fi53 < 15; fi53++ {
        for fi54 := 0; fi54 < 3; fi54++ {
            var v55 int8 = (-(v39))
            if ((-2.25 <= ((1.5 - 1.5) / ((1.5 - 0.1) / 1.0))) || true) {
                v56 := GenS4{F0: 127, F1: (v11 - (v11 - (v11 >> (((v3 * (^(v3)))) & 63))))}
                _ = v56
            } else {
                if (v52 > 0) {
                    v57 := GenS4{F0: 255, F1: v11}
                    var v58 int32 = -906
                    v59 := GenS5{F0: 127}
                    v60 := GenS6{F0: (v52 & (v52 << ((v3) & 63))), F1: 1, F2: v39}
                    if false {
                        v61 := GenS3{F0: (v52 - (v52 / ((42) | 1))), F1: (^(fi54))}
                        var v62 int8 = v39
                        fmt.Println(v52, v58)
                        _ = v61
                        _ = v62
                    }
                    _ = v57
                    _ = v58
                    _ = v59
                    _ = v60
                } else {
                    v63 := GenS3{F0: (v52 + (v52 << ((221) & 63))), F1: (v2 / ((v2) | 1))}
                    var v64 uint32 = (^(v52))
                    v4 -= 879
                    _ = v63
                    _ = v64
                }
            }
            var v65 int16 = genFn0((v3 ^ v3))
            _ = fi54
            _ = v55
            _ = v65
        }
        _ = fi53
    }
    v2 -= int(v11)
    if true {
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v50
    _ = v51
    _ = v52
    _ = v9
}
