package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint; F2 int }
type GenS3 struct { F0 uint; F1 int; F2 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint; F1 int }
func genFn0() int {
    if true {
        var v5 int = 127
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: 1, F1: 100, F2: -697}
            v8 := GenS1{F0: 1, F1: 127, F2: 542}
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    } else {
        switch -938 {
        case 679:
            var v9 uint8 = 39
            _ = v9
        }
        v10 := GenS1{F0: 459, F1: 100, F2: -1}
        switch ("hello" + ("" + "a")) {
        case "a":
            var v11 uint8 = 127
            _ = v11
        }
        _ = v10
    }
    var v12 int64 = -1
    if ((2147483648 <= 2) || (!(true && true))) {
        v12 += (v12 / ((-9223372036854775808) | 1))
        v12 = v12
    }
    _ = v12
    return -9223372036854775808
}
func genFn1() uint {
    if false {
        if false {
            if (710 >= 1) {
                var v13 float32 = 0.1
                _ = v13
            } else {
                for fi14 := 0; fi14 < 4; fi14++ {
                    if false {
                        for fi15 := 0; fi15 < 6; fi15++ {
                            v16 := GenS1{F0: 0, F1: 62, F2: 9223372036854775807}
                            _ = fi15
                            _ = v16
                        }
                        v17 := GenS2{F0: 1, F1: 2, F2: (^(fi14))}
                        v18 := []int{}
                        _ = v17
                        _ = v18
                    } else {
                        v19 := []int{genFn0(), (fi14 << ((4294967295) & 63)), (fi14 / ((fi14) | 1)), (fi14 | (fi14 - genFn0()))}
                        if ((true && true) || (255 == 1)) {
                            if true {
                                var v20 int64 = -1
                                v20 += (-(v20))
                                fmt.Println(fi14)
                                _ = v20
                            } else {
                                var v21 uint64 = 18446744073709551615
                                var v22 float64 = 3.14159
                                var v23 bool = false
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            var v24 string = "hello"
                            if false {
                                v24 = v24
                                v25 := GenS3{F0: 1, F1: fi14, F2: true}
                                _ = v25
                            } else {
                                for fi26 := 0; fi26 < 10; fi26++ {
                                    var v27 int32 = -528
                                    v27 = v27
                                    v28 := GenS3{F0: 18446744073709551615, F1: (fi26 & genFn0()), F2: false}
                                    var v29 uint16 = 0
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                                var v30 int32 = -368
                                _ = v30
                            }
                            _ = v24
                        } else {
                            for fi31 := 0; fi31 < 16; fi31++ {
                                var v32 uint16 = 896
                                _ = fi31
                                _ = v32
                            }
                        }
                        _ = v19
                    }
                    _ = fi14
                }
            }
        } else {
            if true {
                var v33 float32 = 1.5
                v33 = v33
                v34 := []int{2147483647, -1}
                _ = v33
                _ = v34
            } else {
                if true {
                    var v35 int16 = 1
                    _ = v35
                }
                v36 := GenS1{F0: 42, F1: 42, F2: -1}
                _ = v36
            }
            var v37 uint16 = 49
            v37 = (v37 / (((v37 ^ v37)) | 1))
            _ = v37
        }
    } else {
        var v38 int32 = -1
        _ = v38
    }
    var v39 uint16 = 255
    v39 = 42
    v39 -= v39
    v39 -= 945
    _ = v39
    return 2
}

func main() {
    v1 := []int{9223372036854775807, -1, 32767}
    var v2 uint16 = 1
    v3 := []string{"hello", "foo bar", "\n\t"}
    var v4 float64 = 1e10
    for fi40 := 0; fi40 < 5; fi40++ {
        if true {
            v4 = (v4 - (1e10 + 1.5))
            v4 = float64(3.14159)
            var v41 int16 = -608
            _ = v41
        }
        _ = fi40
    }
    v4 = ((1e10 - 0.1) - 3.14159)
    v42 := GenS4{F0: 14}
    var v43 int16 = 0
    if (!true) {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        v43 += v43
    }
    var v44 uint64 = 42
    var v45 float32 = (((-2.25 / 1.0) + (3.14159 * 0.1)) * -2.25)
    if (1 == 59) {
        v44 = (v44 * 629)
        switch -2147483648 {
        case 702:
            switch "hello" {
            case "foo bar":
                var v46 uint32 = 905
                _ = v46
            case "hello":
                for fi47 := 0; fi47 < 11; fi47++ {
                    v43 = (v43 & (^(v43)))
                    for fi48 := 0; fi48 < 5; fi48++ {
                        switch "" {
                        case "foo bar":
                            v44 -= (v44 << ((0) & 63))
                        default:
                            _ = len(v3)
                        }
                        var v49 uint8 = 128
                        v50 := GenS5{F0: v2, F1: genFn1(), F2: uint16(v49)}
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                    _ = fi47
                }
            case "":
                var v51 float64 = float64(float32(v4))
                _ = v51
            }
        default:
            v2 -= (v2 | v2)
        }
    } else {
        for fi52 := 0; fi52 < 6; fi52++ {
            v43 += 1000
            _ = fi52
        }
        switch -222 {
        case 32768:
            var v53 uint32 = 42
            _ = v53
        case 0:
            var v54 int = 4294967295
            _ = v54
        case 1:
            var v55 uint = genFn1()
            _ = v55
        default:
            v4 = float64(v45)
        }
    }
    v45 -= -2.25
    v1 = append(v1, 2147483647)
    v44 = (v44 - (v44 | 2))
    switch (("foo bar" + ((("hello" + "\n\t") + "foo bar") + "a")) + ("foo bar" + ("foo bar" + ("\n\t" + ("foo bar" + ""))))) {
    case "\n\t":
        v43 = v43
    default:
        v56 := GenS4{F0: 135}
        _ = v56
    }
    if false {
        v57 := GenS2{F0: 42, F1: genFn1(), F2: genFn0()}
        _ = v57
    }
    if true {
        if true {
            v58 := GenS3{F0: 509, F1: -881, F2: (v44 > v44)}
            for fi59 := 0; fi59 < 13; fi59++ {
                _ = len(v3)
                for fi60 := 0; fi60 < 10; fi60++ {
                    var v61 int8 = 34
                    var v62 uint = 0
                    var v63 float64 = 0.1
                    _ = fi60
                    _ = v61
                    _ = v62
                    _ = v63
                }
                for fi64 := 0; fi64 < 6; fi64++ {
                    v65 := GenS6{F0: (fi59 >> ((1) & 63))}
                    _ = fi64
                    _ = v65
                }
                _ = fi59
            }
            _ = v58
        } else {
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        }
    } else {
        switch 761 {
        case 2147483647:
            v3 = append(v3, "a")
        case -628:
            var v66 int = 128
            _ = v66
        default:
            if (true || true) {
                var v67 uint8 = uint8(v44)
                var v68 uint64 = 643
                v69 := GenS7{F0: (((("\n\t" + "\n\t") + "foo bar") + "foo bar") + ("foo bar" + "foo bar")), F1: true}
                _ = v67
                _ = v68
                _ = v69
            } else {
                var v70 int16 = v43
                _ = v70
            }
        }
    }
    v71 := GenS8{F0: 255, F1: 9223372036854775807}
    fmt.Println(v4, v43)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v71
}
