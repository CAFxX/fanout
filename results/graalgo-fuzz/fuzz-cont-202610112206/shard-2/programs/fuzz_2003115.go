package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 uint; F2 int16 }
type GenS6 struct { F0 int32; F1 uint32; F2 uint }
type GenS7 struct { F0 uint16; F1 string; F2 bool }
func genFn0() bool {
    if false {
        var v6 uint64 = 42
        var v7 string = "\n\t"
        var v8 int8 = 42
        _ = v6
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 2147483647, F1: 42, F2: 32767}
        v10 := []int{4294967295}
        if true {
            var v11 uint16 = 2
            var v12 uint8 = 56
            _ = v11
            _ = v12
        } else {
            v13 := GenS1{F0: -1, F1: 4294967295, F2: 0}
            switch "\n\t" {
            case "a":
                v14 := GenS1{F0: 0, F1: 565, F2: 677}
                _ = v14
            case "foo bar":
                v15 := GenS1{F0: 1, F1: 0, F2: -9223372036854775808}
                _ = v15
            default:
                for fi16 := 0; fi16 < 6; fi16++ {
                    switch (fi16 >> ((42) & 63)) {
                    case 1:
                        var v17 int64 = 804
                        _ = v17
                    default:
                        var v18 string = "foo bar"
                        _ = v18
                    }
                    _ = fi16
                }
            }
            _ = v13
        }
        _ = v10
        _ = v9
    }
    v19 := []int{}
    if (true && (!false)) {
        if false {
            if false {
                for fi20 := 0; fi20 < 13; fi20++ {
                    switch (fi20 - (fi20 | (^(fi20)))) {
                    case 9223372036854775807:
                        if ((1 <= -1) && false) {
                            for fi21 := 0; fi21 < 15; fi21++ {
                                var v22 int8 = -128
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                                switch "a" {
                                case "a":
                                    if false {
                                        var v23 uint = 42
                                        fmt.Println(fi21)
                                        if false {
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                            if false {
                                                fmt.Println(fi21, v23)
                                                v23 -= v23
                                            } else {
                                                if (int32(fi21) != 121) {
                                                    v19 = append(v19, -2147483648)
                                                } else {
                                                    switch "" {
                                                    case "\n\t":
                                                        v22 = (v22 ^ -79)
                                                    case "foo bar":
                                                        var v24 bool = ("hello" <= "hello")
                                                        _ = v24
                                                    default:
                                                        v25 := GenS1{F0: -2147483648, F1: 1, F2: fi20}
                                                        _ = v25
                                                    }
                                                    switch fi20 {
                                                    case 1:
                                                        var v26 uint16 = 65535
                                                        _ = v26
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if len(v19) > 0 {
                                                            v19[0] = fi21
                                                        }
                                                    case "a":
                                                        fmt.Println(fi21, fi20, v22)
                                                    }
                                                }
                                            }
                                            v22 -= (v22 | (v22 - (v22 + (v22 << (((v23 ^ (v23 - 648))) & 63)))))
                                            v22 = v22
                                            v27 := GenS2{F0: 173, F1: 925}
                                            switch fi21 {
                                            case 9223372036854775807:
                                                fmt.Println(fi20)
                                                var v28 uint8 = 255
                                                _ = v28
                                            default:
                                                if len(v19) >= 2 {
                                                    v19 = v19[1:len(v19)]
                                                }
                                            }
                                            _ = v27
                                        } else {
                                            v23 = v23
                                            v22 += 127
                                            if (!false) {
                                                if len(v19) >= 2 {
                                                    v19 = v19[1:len(v19)]
                                                }
                                                if true {
                                                    _ = len(v19)
                                                    v29 := GenS2{F0: 2, F1: 446}
                                                    _ = v29
                                                } else {
                                                    if ((-548 != 358) && false) {
                                                        var v30 int16 = 0
                                                        var v31 int8 = (v22 / ((v22) | 1))
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                }
                                            }
                                        }
                                        _ = v23
                                    }
                                case "foo bar":
                                    v22 += v22
                                case "":
                                    v32 := GenS3{F0: 708, F1: 2147483647, F2: 0}
                                    _ = v32
                                }
                                _ = fi21
                                _ = v22
                            }
                            fmt.Println(fi20)
                        } else {
                            switch (("a" + "\n\t") + "a") {
                            case "hello":
                                fmt.Println(fi20)
                            case "":
                                if (true || (-9223372036854775808 < 0)) {
                                    var v33 int32 = -1
                                    v33 = (v33 * 136)
                                    _ = v33
                                }
                            default:
                                for fi34 := 0; fi34 < 2; fi34++ {
                                    if true {
                                        var v35 uint8 = 1
                                        v35 = 1
                                        _ = v35
                                    }
                                    var v36 string = (("" + (("\n\t" + "") + "foo bar")) + "\n\t")
                                    _ = fi34
                                    _ = v36
                                }
                            }
                        }
                    case 2147483647:
                        _ = len(v19)
                    default:
                        var v37 int8 = -101
                        _ = v37
                    }
                    _ = fi20
                }
                v38 := GenS4{F0: 1000}
                _ = v38
            } else {
                for fi39 := 0; fi39 < 15; fi39++ {
                    fmt.Println(fi39)
                    _ = fi39
                }
                var v40 float64 = float64(((1e10 / 1.0) - float32(1.5)))
                if (!(!true)) {
                    v40 += (v40 * v40)
                    v40 = (0.1 + (0.1 - 0.0))
                } else {
                    for fi41 := 0; fi41 < 10; fi41++ {
                        if ((fi41 % ((32768) | 1)) <= (fi41 - (fi41 * fi41))) {
                            var v42 int16 = -32768
                            for fi43 := 0; fi43 < 20; fi43++ {
                                v44 := GenS3{F0: 72, F1: 2147483648, F2: 0}
                                var v45 int32 = -578
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                            _ = v42
                        } else {
                            var v46 uint = 65535
                            v47 := GenS5{F0: 773, F1: 614, F2: -1}
                            _ = v46
                            _ = v47
                        }
                        _ = fi41
                    }
                    switch 127 {
                    case 975:
                        if false {
                            for fi48 := 0; fi48 < 4; fi48++ {
                                var v49 string = ("\n\t" + "foo bar")
                                _ = fi48
                                _ = v49
                            }
                        } else {
                            v50 := GenS3{F0: 65535, F1: 65535, F2: 1}
                            fmt.Println(v40)
                            v40 = 1e10
                            _ = v50
                        }
                    default:
                        v40 = float64(3.14159)
                    }
                    for fi51 := 0; fi51 < 19; fi51++ {
                        v19 = append(v19, 0)
                        _ = fi51
                    }
                }
                _ = v40
            }
        }
        var v52 float64 = 1.5
        _ = v52
    } else {
        v53 := GenS6{F0: 609, F1: 749, F2: 388}
        if len(v19) > 0 {
            v19[0] = 128
        }
        _ = v53
    }
    _ = v19
    return (42 >= 2147483648)
}
func genFn1(p0 int16) int16 {
    if (!(-504 <= -696)) {
        v54 := GenS4{F0: p0}
        _ = v54
    }
    var v55 uint = 4294967295
    var v56 int64 = 42
    v56 = (v56 << ((592) & 63))
    v57 := GenS7{F0: 825, F1: ("hello" + "hello"), F2: true}
    _ = p0
    _ = v55
    _ = v56
    _ = v57
    return 0
}

func main() {
    var v1 int = 1
    v2 := make(map[string]string)
    var v3 int = (v1 % ((v1) | 1))
    v4 := make(map[uint32]string)
    var v5 int = 42
    var v58 uint64 = 726
    var v59 uint8 = 0
    _ = len(v4)
    var v60 int16 = 121
    var v61 bool = (v60 >= 1)
    fmt.Println(v5, v1)
    if (("a" <= "") && v61) {
        if (v61 || v61) {
            var v62 int16 = genFn1(genFn1((v60 / (((v60 ^ v60)) | 1))))
            _ = v62
        }
    }
    v61 = true
    v63 := []int{(v3 & int(v60)), (v5 * v5), 857, (v1 >> ((50) & 63))}
    for fi64 := 0; fi64 < 12; fi64++ {
        switch "\n\t" {
        case "\n\t":
            for fi65 := 0; fi65 < 2; fi65++ {
                fmt.Println(v1, v61)
                if len(v63) > 0 {
                    v63[0] = (v1 / (((v5 / (((v1 + v1)) | 1))) | 1))
                }
                _ = fi65
            }
        }
        if (!(((((0.1 * (float32(1e10) / 0.1)) / (3.14159 - 1e10)) / 1.0) - 0.1) != float32(0.1))) {
            v66 := GenS4{F0: int16(v58)}
            v60 = (v60 % ((-1) | 1))
            _ = v66
        } else {
            v67 := GenS6{F0: 42, F1: 1, F2: 374}
            v58 += v58
            _ = v67
        }
        if (v61 && true) {
            if (true && false) {
                v59 -= v59
            }
        }
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v63:", len(v63), v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v60
    _ = v61
    _ = v63
}
