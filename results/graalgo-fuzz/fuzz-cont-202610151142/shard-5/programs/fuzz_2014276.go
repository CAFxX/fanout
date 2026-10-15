package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int16; F2 int16 }
type GenS3 struct { F0 int; F1 bool; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 bool; F2 uint32 }
func genFn0(p0 int64, p1 int) int8 {
    var v4 uint = 332
    switch (((("\n\t" + "hello") + "") + "hello") + ("\n\t" + "")) {
    case "\n\t":
        v5 := []int{4294967295, (p1 >> ((v4) & 63)), -1}
        _ = v5
    default:
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 float32 = 3.14159
            p0 += -1
            v8 := []int{2147483647, p1}
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1() uint32 {
    if true {
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := []int{(^(fi9)), 1}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 int64 = 1
            _ = fi11
            _ = v12
        }
    } else {
        if (false && true) {
            var v13 int = 1
            v14 := []int{32768, v13, (-(v13))}
            v13 = 302
            _ = v13
            _ = v14
        }
    }
    v15 := GenS1{F0: 399}
    for fi16 := 0; fi16 < 6; fi16++ {
        if (128 >= 42) {
            var v17 int8 = 0
            for fi18 := 0; fi18 < 19; fi18++ {
                switch (-(fi18)) {
                case -304:
                    var v19 uint8 = 1
                    _ = v19
                case -347:
                    var v20 uint32 = 425
                    _ = v20
                case -2147483648:
                    v17 = (v17 % ((1) | 1))
                }
                if (!true) {
                    switch (fi16 >> ((2) & 63)) {
                    case 308:
                        v17 -= -50
                        v17 += (v17 % (((v17 / (((v17 * v17)) | 1))) | 1))
                    case -9223372036854775808:
                        var v21 int16 = 684
                        _ = v21
                    case -2147483648:
                        v22 := []int{0, fi16, fi18, (^(fi18))}
                        _ = v22
                    default:
                        v23 := GenS2{F0: 1, F1: -1, F2: 484}
                        _ = v23
                    }
                } else {
                    v24 := []int{fi16, (fi18 * (-(fi16))), 50}
                    v25 := GenS3{F0: (fi16 & (fi16 % (((fi16 % ((fi18) | 1))) | 1))), F1: false, F2: 42}
                    v26 := GenS2{F0: (v17 << ((928) & 63)), F1: 1, F2: -1}
                    _ = v24
                    _ = v25
                    _ = v26
                }
                v27 := []int{}
                _ = fi18
                _ = v27
            }
            var v28 bool = false
            _ = v17
            _ = v28
        }
        _ = fi16
    }
    _ = v15
    return 2147483648
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 127
    v3 := []int{2147483647, 128}
    fmt.Println(v2)
    v29 := GenS2{F0: -109, F1: 32767, F2: 941}
    for fi30 := 0; fi30 < 13; fi30++ {
        switch fi30 {
        case 719:
            v31 := GenS4{F0: 571, F1: 4294967295, F2: 2}
            _ = v31
        case -1:
            if ((v2 - (v2 - (^(v2)))) <= 128) {
                if (587 <= -17) {
                    var v32 bool = true
                    _ = v32
                } else {
                    var v33 uint32 = 1
                    if false {
                        if false {
                            var v34 int32 = -394
                            fmt.Println(v34, v33)
                            v2 = (v2 ^ (v2 >> ((42) & 63)))
                            _ = v34
                        } else {
                            var v35 uint = 2
                            _ = v35
                        }
                        _ = len(v3)
                    } else {
                        v1[(fi30 * fi30)] = -656
                        v33 += 42
                    }
                    _ = v33
                }
            }
        case -498:
            var v36 int64 = 663
            _ = v36
        }
        if (42 > 2) {
            var v37 int16 = -32768
            _ = v37
        } else {
            var v38 uint = 0
            v2 = (v2 >> ((v38) & 63))
            _ = v38
        }
        _ = fi30
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        var v40 uint16 = uint16(v2)
        if (v2 >= v2) {
            v40 = 32767
        } else {
            fmt.Println(fi39)
            if (0 != -2147483648) {
                v40 = v40
                fmt.Println(v40, v40, v40)
            } else {
                for fi41 := 0; fi41 < 12; fi41++ {
                    delete(v1, fi41)
                    if len(v3) > 0 {
                        v3[0] = fi39
                    }
                    _ = fi41
                }
            }
            v42 := GenS5{F0: -105}
            _ = v42
        }
        _ = fi39
        _ = v40
    }
    if (!(468 > 1)) {
        if false {
            fmt.Println(v2)
        }
        var v43 int = 2147483647
        var v44 int64 = 9223372036854775807
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 17; fi45++ {
        var v46 float64 = 3.14159
        _ = fi45
        _ = v46
    }
    v2 = v2
    v47 := GenS2{F0: genFn0(0, -1), F1: 0, F2: -32768}
    v2 = 201
    delete(v1, 32767)
    var v48 int64 = 0
    var v49 int8 = 1
    switch 1 {
    case -2147483648:
        switch "a" {
        case "":
            var v50 float32 = 0.0
            _ = v50
        case "hello":
            v51 := v1[-67]
            _ = v51
        default:
            var v52 int = 0
            _ = v52
        }
    case 32767:
        var v53 int16 = -32768
        _ = v53
    case 42:
        _ = len(v1)
    default:
        v49 -= v49
    }
    fmt.Println(v49)
    var v54 uint64 = 1
    if false {
        fmt.Println(v48, v54)
        v2 -= (v2 | (v2 % ((v2) | 1)))
    } else {
        v55 := GenS2{F0: 0, F1: 32767, F2: 0}
        v56 := GenS6{F0: 0, F1: false, F2: 0}
        v2 -= v2
        _ = v55
        _ = v56
    }
    for fi57 := 0; fi57 < 4; fi57++ {
        v2 -= v2
        _ = fi57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v47
    _ = v48
    _ = v49
    _ = v54
}
