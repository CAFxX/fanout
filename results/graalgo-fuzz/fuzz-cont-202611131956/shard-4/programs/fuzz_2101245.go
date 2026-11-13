package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 uint }
func genFn0(p0 uint32, p1 int8) uint8 {
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := []int{fi4}
        if len(v5) > 0 {
            v5[0] = -2147483648
        }
        _ = fi4
        _ = v5
    }
    p0 = p0
    var v6 uint64 = 1
    for fi7 := 0; fi7 < 13; fi7++ {
        if true {
            fmt.Println(p1, v6)
            switch (^(fi7)) {
            case 0:
                p0 = 42
            case 1:
                p0 -= (p0 + 2)
            }
        }
        switch (fi7 | (fi7 ^ fi7)) {
        case 4294967295:
            if (-9223372036854775808 < 285) {
                for fi8 := 0; fi8 < 20; fi8++ {
                    v6 += v6
                    _ = fi8
                }
                for fi9 := 0; fi9 < 10; fi9++ {
                    var v10 uint32 = p0
                    _ = fi9
                    _ = v10
                }
                fmt.Println(p0, p0, fi7)
            } else {
                for fi11 := 0; fi11 < 19; fi11++ {
                    var v12 int64 = 9223372036854775807
                    _ = fi11
                    _ = v12
                }
            }
        case 527:
            if true {
                for fi13 := 0; fi13 < 9; fi13++ {
                    p0 = (^(p0))
                    v14 := []int{0}
                    fmt.Println(fi13)
                    _ = fi13
                    _ = v14
                }
                p0 = (p0 * (p0 / (((p0 % (((p0 / (((^(p0))) | 1))) | 1))) | 1)))
                var v15 bool = true
                _ = v15
            }
        case 32768:
            var v16 uint8 = 0
            _ = v16
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return 2
}

func main() {
    var v1 uint16 = 0
    v2 := []int{386, 42}
    var v3 int8 = -93
    var v17 uint64 = uint64(v1)
    var v18 int16 = 1000
    var v19 uint8 = uint8(v17)
    switch -2147483648 {
    case 42:
        v20 := GenS1{F0: 2, F1: v17}
        _ = v20
    case 32768:
        if ((true || false) && true) {
            for fi21 := 0; fi21 < 18; fi21++ {
                fmt.Println(v3, v3)
                for fi22 := 0; fi22 < 10; fi22++ {
                    switch (fi21 | fi22) {
                    case 4294967295:
                        v17 -= v17
                    case -690:
                        if len(v2) > 0 {
                            v2[0] = 190
                        }
                    default:
                        v23 := GenS2{F0: (v19 ^ uint8(v17))}
                        _ = v23
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v24 uint = 255
                    _ = fi22
                    _ = v24
                }
                fmt.Println(v19)
                _ = fi21
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if len(v2) > 0 {
                v2[0] = 32767
            }
        }
    default:
        var v25 uint = 4294967295
        _ = v25
    }
    fmt.Println(v19, v18, v17)
    v17 = v17
    switch ("" + "foo bar") {
    case "hello":
        v26 := GenS3{F0: (v19 | 2)}
        _ = v26
    case "\n\t":
        var v27 float32 = 3.14159
        _ = v27
    default:
        for fi28 := 0; fi28 < 17; fi28++ {
            fmt.Println(v18)
            for fi29 := 0; fi29 < 9; fi29++ {
                var v30 int64 = 0
                var v31 int = fi28
                v32 := GenS4{F0: v3, F1: v31}
                if (!(true && false)) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v2) > 0 {
                        v2[0] = fi28
                    }
                    v2 = append(v2, v31)
                    if len(v2) > 0 {
                        v2[0] = -746
                    }
                } else {
                    var v33 string = (("foo bar" + "hello") + "hello")
                    v1 = (v1 >> ((18446744073709551615) & 63))
                    if (!(!true)) {
                        v3 += -1
                        if (!(2147483648 != 926)) {
                            v2 = append(v2, v31)
                            switch v33 {
                            case "hello":
                                if (v30 == (v30 | (-(v30)))) {
                                    if true {
                                        switch (v33 + v33) {
                                        case "hello":
                                            fmt.Println(v1, v1, fi29)
                                        case "a":
                                            v2 = append(v2, -9223372036854775808)
                                            var v34 uint8 = (v19 * 127)
                                            if (v34 <= v34) {
                                                fmt.Println(v34)
                                                switch 644 {
                                                case 42:
                                                    var v35 int32 = 2147483647
                                                    _ = v35
                                                case 32768:
                                                    if false {
                                                        fmt.Println(v33)
                                                        v36 := GenS2{F0: genFn0(777, 79)}
                                                        v18 = -524
                                                        _ = v36
                                                    }
                                                case -9223372036854775808:
                                                    v37 := GenS1{F0: 0, F1: v17}
                                                    _ = v37
                                                }
                                                v38 := GenS4{F0: v3, F1: (-(fi28))}
                                                _ = v38
                                            }
                                            _ = v34
                                        case "\n\t":
                                            var v39 uint16 = 255
                                            _ = v39
                                        default:
                                            v40 := GenS5{F0: (v3 | v3), F1: int(v1)}
                                            fmt.Println(v3)
                                            _ = v40
                                        }
                                        var v41 uint16 = 65535
                                        _ = v41
                                    }
                                    var v42 uint16 = 65535
                                    _ = v42
                                }
                            default:
                                _ = len(v2)
                                v18 = (v18 / (((v18 >> ((1) & 63))) | 1))
                            }
                        }
                        v43 := GenS6{F0: v30, F1: 0}
                        _ = v43
                    }
                    _ = v33
                }
                v17 += (v17 + v17)
                var v44 uint64 = 18446744073709551615
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
                _ = v44
            }
            if (float64((0.0 * (3.14159 * 3.14159))) > 0.0) {
                v45 := GenS4{F0: (v3 - (v3 + (v3 - int8(v18)))), F1: 128}
                var v46 uint64 = 42
                _ = v45
                _ = v46
            } else {
                v47 := GenS7{F0: 2}
                var v48 string = ("\n\t" + "")
                _ = v47
                _ = v48
            }
            _ = fi28
        }
    }
    var v49 uint32 = 307
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v49
}
