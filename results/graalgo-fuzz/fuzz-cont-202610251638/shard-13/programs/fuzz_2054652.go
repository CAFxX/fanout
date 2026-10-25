package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int32; F2 bool }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 uint32 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 18446744073709551615, F1: 0}
        fmt.Println(fi6)
        _ = fi6
        _ = v7
    }
    var v8 int64 = -9223372036854775808
    var v9 int32 = 42
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int32 = 42
        v12 := GenS1{F0: 4294967295, F1: 825}
        v9 = (v9 * (-(v11)))
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{0, 0, 32768, 0}
    _ = v13
    _ = v8
    _ = v9
    return 83
}
func genFn1(p0 int16, p1 uint64) uint64 {
    v14 := []int{}
    if len(v14) > 0 {
        v14[0] = 127
    }
    _ = len(v14)
    if len(v14) > 0 {
        v14[0] = -9223372036854775808
    }
    _ = p0
    _ = p1
    _ = v14
    return (p1 % ((1) | 1))
}

func main() {
    var v1 int16 = -391
    var v2 uint16 = 1
    v3 := make(map[uint32]uint32)
    var v4 int32 = 594
    v5 := []uint32{4294967295, 470, 21, 479}
    var v15 string = "hello"
    var v16 uint16 = (v2 + v2)
    _ = len(v5)
    v4 -= 809
    v15 = (("foo bar" + "foo bar") + "a")
    v2 = (v2 | v2)
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 bool = true
            fmt.Println(v19)
            var v20 int16 = -576
            var v21 int32 = 1
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        if true {
            switch fi17 {
            case -9223372036854775808:
                var v22 uint = 1
                _ = v22
            default:
                v3[2] = 867
            }
        }
        _ = fi17
    }
    v23 := GenS2{F0: 255}
    switch v15 {
    case "a":
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 int32 = v4
            _ = fi24
            _ = v25
        }
    case "hello":
        if true {
            if (2147483647 >= -2147483648) {
                if true {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        for fi27 := 0; fi27 < 1; fi27++ {
                            v28 := GenS2{F0: 40}
                            v2 = (v16 >> ((786) & 63))
                            _ = fi27
                            _ = v28
                        }
                        fmt.Println(v4, v1)
                        v2 += v2
                        _ = fi26
                    }
                    v1 -= (-(v1))
                }
                v1 = v1
            }
            fmt.Println(v16, v2, v16)
            if len(v5) > 0 {
                v5[0] = 0
            }
        }
    default:
        if false {
            v5 = append(v5, genFn0())
            for fi29 := 0; fi29 < 15; fi29++ {
                for fi30 := 0; fi30 < 18; fi30++ {
                    v2 = v16
                    switch ((v15 + ((v15 + "") + (((v15 + "a") + (v15 + "")) + v15))) + v15) {
                    case "\n\t":
                        var v31 uint8 = 2
                        v4 = int32(v16)
                        _ = v31
                    case "foo bar":
                        v32 := GenS1{F0: 42, F1: 0}
                        _ = v32
                    case "a":
                        v33 := GenS2{F0: 0}
                        _ = v33
                    }
                    if false {
                        _ = len(v3)
                        switch "hello" {
                        case "foo bar":
                            v2 += 559
                        }
                    } else {
                        v34 := v3[1]
                        fmt.Println(v2, v4, v15)
                        v16 = (^(v16))
                        switch (fi29 & (fi30 >> ((255) & 63))) {
                        case 179:
                            var v35 bool = true
                            v35 = v35
                            if v35 {
                                if true {
                                    var v36 uint16 = v2
                                    var v37 bool = (!(!true))
                                    v2 = v36
                                    v3[v34] = v34
                                    if (uint8(v16) != 2) {
                                        var v38 uint = 2
                                        var v39 float32 = float32(1e10)
                                        _ = v38
                                        _ = v39
                                    }
                                    _ = v36
                                    _ = v37
                                }
                                if (v35 || (0.0 > 1e10)) {
                                    if ((v1 << ((18446744073709551615) & 63)) >= (v1 % ((v1) | 1))) {
                                        v40 := GenS2{F0: 255}
                                        v5 = append(v5, v34)
                                        _ = v40
                                    }
                                    switch (fi29 ^ -2147483648) {
                                    case 127:
                                        var v41 int32 = (v4 - v4)
                                        _ = v41
                                    }
                                }
                                fmt.Println(v2, fi30, fi30)
                            }
                            _ = v35
                        case 42:
                            if (-43 == -96) {
                                v42 := GenS2{F0: 42}
                                v16 = (v16 + (v16 | (^(v16))))
                                if false {
                                    var v43 string = v15
                                    fmt.Println(v1, v43)
                                    _ = v43
                                } else {
                                    if (true || true) {
                                        if (1e10 > 0.1) {
                                            switch "" {
                                            case "":
                                                if ("foo bar" >= ((((v15 + v15) + ("\n\t" + ((v15 + v15) + v15))) + "") + (v15 + "foo bar"))) {
                                                    var v44 bool = (true && false)
                                                    if len(v5) > 0 {
                                                        v5[0] = (v34 >> ((4294967295) & 63))
                                                    }
                                                    _ = v44
                                                }
                                            case "foo bar":
                                                var v45 int8 = -1
                                                _ = v45
                                            default:
                                                switch (v15 + "") {
                                                case "":
                                                    v46 := GenS1{F0: 42, F1: 831}
                                                    _ = v46
                                                case "foo bar":
                                                    v47 := v3[genFn0()]
                                                    _ = v47
                                                default:
                                                    var v48 int32 = v4
                                                    _ = v48
                                                }
                                            }
                                            var v49 int = (fi30 << ((477) & 63))
                                            if false {
                                                v50 := GenS3{F0: 127, F1: (^(v2))}
                                                _ = v50
                                            } else {
                                                v34 = (v34 | 0)
                                                v51 := v3[545]
                                                v51 = (v51 | v34)
                                                fmt.Println(v15, v1, v49)
                                                _ = v51
                                            }
                                            _ = v49
                                        }
                                        delete(v3, (v34 + v34))
                                    }
                                }
                                _ = v42
                            }
                        case -1:
                            switch ((v15 + v15) + "foo bar") {
                            case "\n\t":
                                if (305 < 2) {
                                    v15 = ((((v15 + v15) + v15) + ("\n\t" + v15)) + "")
                                    v52 := GenS1{F0: 4294967295, F1: 4294967295}
                                    v16 = v16
                                    _ = v52
                                } else {
                                    v53 := GenS4{F0: (v15 + "hello")}
                                    _ = v53
                                }
                            default:
                                v15 = (v15 + v15)
                            }
                        }
                        _ = v34
                    }
                    _ = fi30
                }
                for fi54 := 0; fi54 < 12; fi54++ {
                    fmt.Println(v15)
                    v55 := GenS5{F0: (v15 + v15), F1: (v4 & (v4 * -2147483648)), F2: false}
                    var v56 uint16 = (^(v2))
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                _ = fi29
            }
            if ((-895 >= -1) && ("" <= ((("foo bar" + v15) + v15) + (v15 + "a")))) {
                switch ("\n\t" + "hello") {
                case "\n\t":
                    switch ((("a" + v15) + v15) + v15) {
                    case "\n\t":
                        v57 := GenS2{F0: 1}
                        _ = v57
                    case "a":
                        v58 := GenS6{F0: -2147483648, F1: 42}
                        _ = v58
                    case "":
                        v59 := GenS1{F0: 65535, F1: 1}
                        _ = v59
                    default:
                        var v60 uint64 = 2
                        _ = v60
                    }
                }
                var v61 bool = ((1.5 + 0.0) <= 0.1)
                v62 := GenS5{F0: v15, F1: v4, F2: (255 != 720)}
                _ = v61
                _ = v62
            } else {
                var v63 int = -9223372036854775808
                v64 := GenS7{F0: v1, F1: false}
                _ = v63
                _ = v64
            }
        }
    }
    v65 := GenS5{F0: v15, F1: 0, F2: false}
    switch 128 {
    case 32767:
        var v66 uint16 = (v16 | v2)
        _ = v66
    case 0:
        switch -850 {
        case 128:
            delete(v3, 4294967295)
        case 711:
            v67 := GenS8{F0: 0}
            _ = v67
        }
    case -9223372036854775808:
        var v68 bool = (1.5 < -2.25)
        _ = v68
    default:
        v3[65535] = 237
    }
    fmt.Println(v1, v16)
    delete(v3, 2147483647)
    if len(v5) > 0 {
        v5[0] = 1
    }
    for fi69 := 0; fi69 < 6; fi69++ {
        var v70 uint = 18446744073709551615
        v71 := GenS3{F0: 78, F1: 65535}
        v15 = ("\n\t" + v15)
        _ = fi69
        _ = v70
        _ = v71
    }
    var v72 bool = true
    fmt.Println(v1, v16, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v72
}
