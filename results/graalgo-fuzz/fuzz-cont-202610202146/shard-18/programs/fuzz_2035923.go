package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint }
type GenS2 struct { F0 uint; F1 int32; F2 bool }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 uint8; F2 bool }
func genFn0() int64 {
    if true {
        for fi4 := 0; fi4 < 6; fi4++ {
            var v5 int64 = -1
            switch 128 {
            case -695:
                v5 = -1
            default:
                v6 := []int{}
                _ = v6
            }
            _ = fi4
            _ = v5
        }
        var v7 uint64 = 2
        var v8 int8 = 0
        _ = v7
        _ = v8
    } else {
        var v9 string = ((("foo bar" + (("a" + "foo bar") + "foo bar")) + "") + ((("foo bar" + "a") + "hello") + ("\n\t" + "foo bar")))
        v10 := GenS1{F0: 1, F1: 2147483647, F2: 42}
        v11 := []int{}
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS1{F0: 32767, F1: -1, F2: 1}
            _ = fi13
            _ = v14
        }
        v15 := []int{(-(fi12)), -527, fi12}
        _ = fi12
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        if true {
            var v17 int32 = -2147483648
            _ = v17
        } else {
            v18 := GenS1{F0: 42, F1: (-(fi16)), F2: 65535}
            _ = v18
        }
        var v19 int8 = -79
        v20 := []int{(fi16 % ((604) | 1)), (fi16 ^ (fi16 & fi16)), fi16}
        _ = fi16
        _ = v19
        _ = v20
    }
    if false {
        for fi21 := 0; fi21 < 20; fi21++ {
            switch "" {
            case "foo bar":
                v22 := GenS1{F0: 1, F1: (fi21 + (fi21 >> ((776) & 63))), F2: 42}
                _ = v22
            default:
                v23 := []int{(fi21 << ((2) & 63))}
                _ = v23
            }
            var v24 float32 = (float32(-2.25) / -2.25)
            if ("hello" <= ("hello" + "a")) {
                var v25 uint64 = 4294967295
                for fi26 := 0; fi26 < 19; fi26++ {
                    v27 := []int{(fi21 | -613), (fi26 + -2147483648), -128}
                    if true {
                        var v28 string = "foo bar"
                        _ = v28
                    } else {
                        var v29 uint32 = 2147483647
                        _ = v29
                    }
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
            _ = fi21
            _ = v24
        }
        v30 := GenS2{F0: 675, F1: 0, F2: (true && true)}
        _ = v30
    }
    if false {
        if true {
            var v31 float32 = 1e10
            _ = v31
        } else {
            var v32 int16 = 195
            switch "a" {
            case "a":
                if true {
                    v33 := GenS3{F0: -15, F1: 2147483647}
                    if (2147483647 == 903) {
                        if false {
                            var v34 float32 = 0.0
                            v34 = (v34 + ((1.5 + v34) * v34))
                            switch -216 {
                            case 32767:
                                v34 = (v34 * (v34 / 3.14159))
                            default:
                                if (127 <= -28) {
                                    for fi35 := 0; fi35 < 4; fi35++ {
                                        if true {
                                            var v36 int = (^(fi35))
                                            switch "hello" {
                                            case "":
                                                var v37 int8 = -128
                                                _ = v37
                                            case "foo bar":
                                                var v38 uint32 = 610
                                                _ = v38
                                            default:
                                                var v39 uint = 255
                                                _ = v39
                                            }
                                            _ = v36
                                        } else {
                                            var v40 int8 = 0
                                            var v41 uint32 = 65535
                                            _ = v40
                                            _ = v41
                                        }
                                        v34 = v34
                                        _ = fi35
                                    }
                                    fmt.Println(v34, v32)
                                }
                            }
                            _ = v34
                        }
                        var v42 int = 128
                        for fi43 := 0; fi43 < 1; fi43++ {
                            if (!false) {
                                v44 := []int{(v42 / (((^(v42))) | 1)), (-(v42)), fi43}
                                if false {
                                    switch "foo bar" {
                                    case "a":
                                        if len(v44) > 0 {
                                            v44[0] = (^(fi43))
                                        }
                                    case "hello":
                                        v45 := GenS1{F0: 2, F1: (fi43 & v42), F2: 700}
                                        _ = v45
                                    }
                                    v32 += (-(v32))
                                } else {
                                    v42 -= (fi43 % ((523) | 1))
                                    if (0 > 65535) {
                                        var v46 uint64 = 4294967295
                                        _ = v46
                                    } else {
                                        for fi47 := 0; fi47 < 18; fi47++ {
                                            v42 = (^(fi43))
                                            v48 := GenS1{F0: 42, F1: 32767, F2: 909}
                                            switch (fi47 - 42) {
                                            case 845:
                                                fmt.Println(fi47, v32, v42)
                                                var v49 string = "hello"
                                                if (("a" + v49) <= (v49 + v49)) {
                                                    if false {
                                                        var v50 uint16 = 255
                                                        _ = v50
                                                    } else {
                                                        var v51 bool = false
                                                        _ = v51
                                                    }
                                                    v32 -= (v32 & -32768)
                                                }
                                                _ = v49
                                            case -898:
                                                var v52 int32 = 42
                                                _ = v52
                                            default:
                                                if true {
                                                    v32 += 32767
                                                }
                                            }
                                            _ = fi47
                                            _ = v48
                                        }
                                        fmt.Println(v32, fi43, v32)
                                    }
                                }
                                var v53 uint8 = 128
                                _ = v44
                                _ = v53
                            } else {
                                for fi54 := 0; fi54 < 18; fi54++ {
                                    var v55 float32 = 0.0
                                    _ = fi54
                                    _ = v55
                                }
                                var v56 int8 = -1
                                v57 := GenS4{F0: 218}
                                _ = v56
                                _ = v57
                            }
                            var v58 int = 60
                            v59 := []int{fi43, fi43, v42, v58}
                            _ = fi43
                            _ = v58
                            _ = v59
                        }
                        _ = v42
                    }
                    _ = v33
                } else {
                    v32 -= v32
                }
            case "":
                var v60 uint16 = 0
                _ = v60
            case "foo bar":
                var v61 bool = true
                _ = v61
            default:
                v32 -= v32
            }
            _ = v32
        }
        var v62 uint8 = 17
        var v63 uint8 = (v62 >> ((1) & 63))
        _ = v62
        _ = v63
    }
    return 1
}
func genFn1(p0 int) int64 {
    var v64 float64 = (float64(3.14159) - -2.25)
    if true {
        v65 := []int{(p0 | 9223372036854775807)}
        v66 := GenS5{F0: false}
        var v67 uint16 = 15
        _ = v65
        _ = v66
        _ = v67
    }
    v64 = (0.0 - 3.14159)
    _ = p0
    _ = v64
    return 1
}

func main() {
    var v1 int16 = -32768
    var v2 uint = 67
    v3 := make(map[uint32]uint32)
    v1 += v1
    var v68 bool = true
    switch "\n\t" {
    case "\n\t":
        v2 = v2
    }
    var v69 int8 = 127
    var v70 int16 = 773
    switch -275 {
    case -744:
        if (2 == 360) {
            v1 -= (v70 & v70)
        }
    case -2147483648:
        var v71 uint = (v2 * v2)
        _ = v71
    }
    v72 := []int{42}
    v68 = true
    for fi73 := 0; fi73 < 4; fi73++ {
        fmt.Println(v68)
        if len(v72) > 0 {
            v72[0] = (-(fi73))
        }
        _ = len(v3)
        _ = fi73
    }
    v74 := v3[2]
    var v75 uint64 = 18446744073709551615
    if len(v72) >= 2 {
        v72 = v72[1:len(v72)]
    }
    fmt.Println(v2)
    if false {
        if v68 {
            var v76 int8 = (^(v69))
            _ = v76
        }
        switch (("foo bar" + "a") + "hello") {
        case "foo bar":
            var v77 int64 = -797
            _ = v77
        case "":
            if true {
                v78 := GenS3{F0: -73, F1: int32(v2)}
                _ = v78
            }
        case "hello":
            v79 := GenS3{F0: v69, F1: 1}
            _ = v79
        }
    }
    var v80 int64 = -9223372036854775808
    switch 0 {
    case 32767:
        if true {
            switch (((("hello" + "\n\t") + "hello") + "hello") + ((((("" + (("a" + "\n\t") + ("a" + ("foo bar" + "")))) + "a") + ("\n\t" + "\n\t")) + "hello") + "")) {
            case "a":
                v69 = v69
            case "hello":
                var v81 int64 = genFn1(127)
                _ = v81
            case "\n\t":
                for fi82 := 0; fi82 < 9; fi82++ {
                    v83 := GenS6{F0: (^(v80)), F1: 203, F2: ((v80 << ((v2) & 63)) <= (v80 >> (((v2 * (^(v2)))) & 63)))}
                    _ = fi82
                    _ = v83
                }
            }
            fmt.Println(v74)
            v84 := v3[v74]
            _ = v84
        }
    case 42:
        if true {
            v80 = v80
        } else {
            v69 = 127
        }
    case -776:
        for fi85 := 0; fi85 < 12; fi85++ {
            switch fi85 {
            case -653:
                _ = len(v3)
            case 32768:
                var v86 uint32 = 4294967295
                _ = v86
            default:
                var v87 int = fi85
                _ = v87
            }
            v69 -= (-(v69))
            _ = fi85
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v68
    _ = v69
    _ = v70
    _ = v72
    _ = v74
    _ = v75
    _ = v80
}
