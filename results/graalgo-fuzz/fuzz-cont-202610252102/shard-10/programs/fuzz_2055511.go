package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 bool; F1 string; F2 bool }
type GenS6 struct { F0 int; F1 int32; F2 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 bool; F1 uint64; F2 uint32 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 bool = false
        if false {
            var v9 int = (fi7 / ((fi7) | 1))
            for fi10 := 0; fi10 < 3; fi10++ {
                v11 := GenS1{F0: 544}
                switch "" {
                case "hello":
                    v12 := GenS1{F0: 65535}
                    _ = v12
                case "foo bar":
                    var v13 int16 = -32768
                    _ = v13
                default:
                    var v14 uint = 4294967295
                    _ = v14
                }
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        var v15 bool = (!v8)
        _ = fi7
        _ = v15
        _ = v8
    }
    if false {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int8 = -74
            if (true || ("\n\t" > "")) {
                v18 := []int{fi16, (fi16 + (fi16 / (((^(fi16))) | 1)))}
                fmt.Println(fi16)
                _ = v18
            } else {
                for fi19 := 0; fi19 < 17; fi19++ {
                    var v20 int64 = 9223372036854775807
                    if true {
                        switch (fi16 / ((fi19) | 1)) {
                        case 32768:
                            switch (fi19 - (^(fi19))) {
                            case 309:
                                if (!true) {
                                    if true {
                                        switch fi19 {
                                        case 127:
                                            if true {
                                                v20 -= -289
                                                fmt.Println(fi16)
                                            } else {
                                                var v21 uint8 = 120
                                                v22 := GenS2{F0: (v17 & 1)}
                                                var v23 int64 = (v20 ^ (-(v20)))
                                                _ = v21
                                                _ = v22
                                                _ = v23
                                            }
                                        case 32768:
                                            v17 = -1
                                        case 967:
                                            v20 = v20
                                        }
                                        switch "hello" {
                                        case "":
                                            switch (("\n\t" + ("hello" + "")) + "a") {
                                            case "a":
                                                v24 := GenS2{F0: v17}
                                                _ = v24
                                            }
                                        case "a":
                                            v25 := []int{}
                                            _ = v25
                                        case "hello":
                                            v26 := GenS3{F0: 656, F1: 4294967295}
                                            _ = v26
                                        default:
                                            var v27 uint = 806
                                            _ = v27
                                        }
                                    } else {
                                        var v28 uint32 = 0
                                        _ = v28
                                    }
                                    fmt.Println(v20, v20, v17)
                                } else {
                                    v17 = v17
                                    var v29 uint = 4294967295
                                    v30 := GenS1{F0: v29}
                                    v31 := GenS1{F0: (v29 | 442)}
                                    v32 := []int{(fi16 - fi16), fi19, int(v20), (fi16 % ((-9223372036854775808) | 1))}
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                            }
                        }
                        v33 := GenS1{F0: 67}
                        _ = v33
                    } else {
                        if false {
                            v34 := GenS2{F0: v17}
                            _ = v34
                        }
                        fmt.Println(fi19)
                    }
                    v35 := GenS2{F0: (^(v17))}
                    _ = fi19
                    _ = v20
                    _ = v35
                }
                if true {
                    var v36 uint = 0
                    v36 += (v36 * v36)
                    var v37 int32 = 0
                    _ = v36
                    _ = v37
                } else {
                    v38 := []int{(fi16 >> ((697) & 63)), (fi16 * (fi16 - fi16)), (fi16 & 42), -804}
                    v17 += (v17 >> ((18446744073709551615) & 63))
                    for fi39 := 0; fi39 < 13; fi39++ {
                        v38 = append(v38, 127)
                        _ = fi39
                    }
                    _ = v38
                }
            }
            _ = fi16
            _ = v17
        }
        var v40 float64 = ((-2.25 / 1e10) / 1.5)
        v41 := []int{-159, 9223372036854775807}
        _ = v40
        _ = v41
    }
    v42 := []int{-856, 32768}
    v43 := GenS4{F0: 42, F1: -698}
    _ = v42
    _ = v43
    return -1
}
func genFn1(p0 uint16) int64 {
    var v44 bool = false
    if v44 {
        v45 := GenS5{F0: v44, F1: ("foo bar" + ("hello" + "")), F2: (128 >= 1)}
        v46 := []int{}
        _ = v45
        _ = v46
    }
    var v47 uint32 = 2147483648
    v48 := []int{8}
    var v49 bool = v44
    _ = p0
    _ = v44
    _ = v47
    _ = v48
    _ = v49
    return 561
}

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 133
    v3 := make(map[uint32]string)
    v4 := []uint32{93, 42}
    var v5 uint8 = 255
    var v6 bool = true
    v5 = 0
    for fi50 := 0; fi50 < 10; fi50++ {
        for fi51 := 0; fi51 < 2; fi51++ {
            switch "" {
            case "foo bar":
                v6 = (float64(float32(1.5)) >= 3.14159)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v52 uint16 = 1
                _ = v52
            default:
                v2 += 2
            }
            _ = fi51
        }
        for fi53 := 0; fi53 < 11; fi53++ {
            v4 = append(v4, 1)
            _ = fi53
        }
        for fi54 := 0; fi54 < 16; fi54++ {
            var v55 uint32 = 42
            var v56 bool = false
            _ = fi54
            _ = v55
            _ = v56
        }
        _ = fi50
    }
    for fi57 := 0; fi57 < 4; fi57++ {
        for fi58 := 0; fi58 < 17; fi58++ {
            v4 = append(v4, 2147483648)
            v3[106] = ""
            v6 = (true && v6)
            _ = fi58
        }
        v59 := GenS5{F0: v6, F1: ("hello" + "\n\t"), F2: v6}
        _ = fi57
        _ = v59
    }
    v1 -= (v1 * 0)
    switch -9223372036854775808 {
    case -76:
        v60 := GenS5{F0: v6, F1: "hello", F2: (!(-159 >= 32767))}
        _ = v60
    case -134:
        v1 = 65535
    case 32768:
        v61 := GenS3{F0: 2, F1: v2}
        _ = v61
    default:
        var v62 uint32 = 896
        _ = v62
    }
    v4 = append(v4, 1)
    v63 := GenS6{F0: 9223372036854775807, F1: -2147483648, F2: -1}
    v2 -= (v2 % (((v2 & (^(v2)))) | 1))
    v4 = append(v4, uint32(v2))
    for fi64 := 0; fi64 < 11; fi64++ {
        v65 := GenS2{F0: int8(fi64)}
        _ = fi64
        _ = v65
    }
    if (v6 && true) {
        v66 := GenS3{F0: 65535, F1: v2}
        switch "foo bar" {
        case "":
            v5 = 227
        default:
            if v6 {
                var v67 int8 = 42
                _ = v67
            }
        }
        fmt.Println(v1, v6, v2)
        _ = v66
    }
    if v6 {
        for fi68 := 0; fi68 < 8; fi68++ {
            v69 := GenS7{F0: v6}
            _ = fi68
            _ = v69
        }
        var v70 float64 = (1.5 - (1e10 - -2.25))
        _ = v70
    } else {
        fmt.Println(v6, v2)
        for fi71 := 0; fi71 < 5; fi71++ {
            var v72 float32 = 0.0
            fmt.Println(fi71, v1, fi71)
            _ = fi71
            _ = v72
        }
    }
    if true {
        v73 := GenS8{F0: (v6 && v6), F1: (v1 * 1)}
        _ = v73
    }
    switch ("foo bar" + ("foo bar" + "")) {
    case "a":
        var v74 int8 = 127
        _ = v74
    case "":
        var v75 uint32 = 2147483647
        _ = v75
    case "hello":
        for fi76 := 0; fi76 < 18; fi76++ {
            var v77 uint16 = v1
            var v78 bool = true
            _ = fi76
            _ = v77
            _ = v78
        }
    default:
        if v6 {
            v79 := GenS9{F0: (!("a" <= ("\n\t" + ""))), F1: v2, F2: 0}
            _ = v79
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
}
