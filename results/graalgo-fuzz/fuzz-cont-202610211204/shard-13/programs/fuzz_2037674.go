package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 int64 }
func genFn0(p0 int8) int32 {
    var v5 uint16 = 32767
    switch 42 {
    case -2147483648:
        var v6 uint8 = 42
        _ = v6
    default:
        switch ("foo bar" + "foo bar") {
        case "":
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 bool = (0 == 417)
                _ = fi7
                _ = v8
            }
        case "foo bar":
            if (42 <= 1) {
                v9 := []int{663, -1}
                var v10 uint16 = 65535
                _ = v10
                _ = v9
            }
        case "\n\t":
            for fi11 := 0; fi11 < 1; fi11++ {
                for fi12 := 0; fi12 < 17; fi12++ {
                    v13 := GenS1{F0: 65535, F1: -985}
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
        default:
            v14 := []int{42}
            _ = v14
        }
    }
    var v15 uint64 = 18446744073709551615
    var v16 uint32 = 1
    _ = p0
    _ = v15
    _ = v16
    _ = v5
    return 0
}
func genFn1(p0 uint32, p1 string) uint64 {
    for fi17 := 0; fi17 < 10; fi17++ {
        if (!(p1 <= (p1 + (("\n\t" + "hello") + "")))) {
            if true {
                if true {
                    v18 := GenS1{F0: 699, F1: -1}
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        v20 := []int{fi17, (fi17 & (fi17 | 32767))}
                        var v21 int = (fi17 + (-(fi19)))
                        v22 := GenS2{F0: 209, F1: true}
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                }
                if ((328 < 42) || ((1 > int16(p0)) && (p0 != p0))) {
                    v23 := GenS1{F0: 755, F1: -32768}
                    fmt.Println(fi17)
                    _ = v23
                } else {
                    v24 := GenS3{F0: 1, F1: fi17}
                    for fi25 := 0; fi25 < 2; fi25++ {
                        if (9223372036854775807 != -9223372036854775808) {
                            if (255 <= 2) {
                                var v26 string = ((p1 + "foo bar") + ((("\n\t" + "a") + p1) + "hello"))
                                _ = v26
                            } else {
                                p0 = (p0 & (p0 ^ p0))
                                v27 := []int{1, fi25, fi25}
                                _ = v27
                            }
                        } else {
                            switch (p1 + "\n\t") {
                            case "a":
                                var v28 uint16 = 32767
                                _ = v28
                            case "":
                                var v29 int8 = -1
                                _ = v29
                            case "hello":
                                p0 += p0
                            default:
                                v30 := GenS3{F0: -1, F1: (fi25 / ((fi17) | 1))}
                                _ = v30
                            }
                            var v31 int8 = -128
                            if ((v31 << ((0) & 63)) < -12) {
                                fmt.Println(fi25, p1)
                                v32 := []int{fi25, (fi25 * (fi17 >> ((4294967295) & 63))), (fi25 * (fi25 ^ (fi25 % (((fi17 | (fi25 ^ -931))) | 1))))}
                                p1 = "a"
                                _ = v32
                            }
                            _ = v31
                        }
                        switch fi17 {
                        case 4294967295:
                            var v33 int16 = -893
                            _ = v33
                        default:
                            switch 1 {
                            case 127:
                                if true {
                                    v34 := GenS2{F0: 1, F1: false}
                                    v35 := []int{}
                                    fmt.Println(fi25, fi25, p0)
                                    fmt.Println(p1, fi17)
                                    var v36 float64 = float64(-2.25)
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                            default:
                                var v37 bool = true
                                _ = v37
                            }
                        }
                        if false {
                            v38 := GenS3{F0: genFn0(-1), F1: -653}
                            v39 := []int{(fi17 + 128), 378, fi17}
                            _ = v38
                            _ = v39
                        }
                        _ = fi25
                    }
                    fmt.Println(p0, p1, fi17)
                    _ = v24
                }
            }
            if (genFn0(-128) == -733) {
                fmt.Println(p0, fi17)
                v40 := []int{fi17, -613, (fi17 | (fi17 * (fi17 - fi17))), (fi17 ^ fi17)}
                _ = v40
            } else {
                v41 := GenS1{F0: 65535, F1: -1}
                _ = v41
            }
            for fi42 := 0; fi42 < 16; fi42++ {
                v43 := GenS4{F0: 18446744073709551615, F1: 1, F2: 1}
                switch (fi17 + (fi42 & fi42)) {
                case -1:
                    v44 := GenS3{F0: 2147483647, F1: fi42}
                    _ = v44
                case 89:
                    p0 += (^(p0))
                default:
                    var v45 uint = 2
                    _ = v45
                }
                _ = fi42
                _ = v43
            }
        } else {
            v46 := GenS5{F0: 2, F1: 2}
            v47 := []int{(fi17 % ((-2147483648) | 1)), (fi17 >> ((0) & 63)), (fi17 ^ 128), (fi17 - (fi17 * fi17))}
            _ = v46
            _ = v47
        }
        _ = fi17
    }
    v48 := []int{32768}
    _ = p0
    _ = p1
    _ = v48
    return 581
}

func main() {
    var v1 uint8 = 128
    var v2 uint32 = 2
    v3 := make(map[string]uint32)
    var v4 int8 = 0
    var v49 uint32 = 2147483647
    v50 := v3["a"]
    v50 = v49
    if ((v50 & v49) < 65535) {
        v50 = (v49 + v2)
        v2 += v50
        v51 := GenS5{F0: 2, F1: 2}
        _ = v51
    }
    for fi52 := 0; fi52 < 2; fi52++ {
        v53 := []int{}
        _ = fi52
        _ = v53
    }
    var v54 int16 = 42
    fmt.Println(v2, v49)
    var v55 uint64 = 18446744073709551615
    v56 := GenS6{F0: 167, F1: false}
    v54 -= v54
    fmt.Println(v49, v55, v49)
    fmt.Println(v2)
    fmt.Println(v2, v4)
    if true {
        for fi57 := 0; fi57 < 18; fi57++ {
            var v58 string = "hello"
            _ = fi57
            _ = v58
        }
    }
    fmt.Println(v54)
    if (false && false) {
        if (!true) {
            v59 := []int{32768}
            v60 := GenS5{F0: 1, F1: 553}
            v59 = append(v59, -9223372036854775808)
            _ = v59
            _ = v60
        } else {
            v4 -= (v4 - (-(v4)))
        }
        for fi61 := 0; fi61 < 2; fi61++ {
            for fi62 := 0; fi62 < 15; fi62++ {
                if true {
                    v4 = (v4 / ((105) | 1))
                } else {
                    v63 := []int{(-(fi61))}
                    delete(v3, ("" + "\n\t"))
                    if len(v63) >= 2 {
                        v63 = v63[1:len(v63)]
                    }
                    _ = v63
                }
                _ = fi62
            }
            v64 := GenS5{F0: 643, F1: 0}
            _ = fi61
            _ = v64
        }
        v49 += 2147483648
    } else {
        if (0.1 != 1.5) {
            v54 += v54
            if false {
                for fi65 := 0; fi65 < 19; fi65++ {
                    for fi66 := 0; fi66 < 3; fi66++ {
                        if false {
                            v67 := GenS7{F0: -9223372036854775808}
                            _ = v67
                        } else {
                            switch "foo bar" {
                            case "a":
                                switch ("\n\t" + "") {
                                case "\n\t":
                                    var v68 bool = ((0 < -636) && false)
                                    _ = v68
                                }
                            case "":
                                var v69 uint16 = 513
                                _ = v69
                            case "\n\t":
                                v49 -= v49
                            }
                            fmt.Println(v49)
                            var v70 int64 = 138
                            v50 += 2147483648
                            _ = v70
                        }
                        _ = fi66
                    }
                    if (42 != 9223372036854775807) {
                        for fi71 := 0; fi71 < 4; fi71++ {
                            var v72 float64 = (1.5 + 3.14159)
                            var v73 uint64 = (v55 | (v55 - uint64(v50)))
                            _ = fi71
                            _ = v72
                            _ = v73
                        }
                    } else {
                        var v74 float64 = 3.14159
                        _ = v74
                    }
                    _ = fi65
                }
                for fi75 := 0; fi75 < 2; fi75++ {
                    fmt.Println(v50)
                    v76 := v3["foo bar"]
                    _ = fi75
                    _ = v76
                }
            }
        }
        var v77 int32 = 212
        _ = v77
    }
    for fi78 := 0; fi78 < 14; fi78++ {
        v49 = (v2 * 549)
        var v79 int16 = 1
        if false {
            var v80 uint64 = genFn1((v2 & 16), (("a" + "hello") + "foo bar"))
            _ = v80
        } else {
            switch fi78 {
            case -927:
                v55 += v55
            case 32767:
                v81 := []int{-1}
                _ = v81
            }
        }
        _ = fi78
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v54
    _ = v55
    _ = v56
}
