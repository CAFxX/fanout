package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int8; F1 int16 }
func genFn0(p0 int16) int64 {
    var v6 uint32 = 0
    v6 = (v6 & (v6 | v6))
    if (true && true) {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int = 128
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := []int{fi9}
            v11 := GenS1{F0: (p0 % ((32767) | 1)), F1: false}
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        p0 = p0
    } else {
        if (!false) {
            var v12 bool = (!true)
            var v13 uint = 0
            _ = v12
            _ = v13
        }
        var v14 int16 = p0
        for fi15 := 0; fi15 < 14; fi15++ {
            v16 := GenS2{F0: (p0 > v14), F1: (p0 + (p0 / ((p0) | 1)))}
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    v17 := GenS2{F0: (6 > 255), F1: (-(p0))}
    var v18 int = 4294967295
    _ = p0
    _ = v17
    _ = v18
    _ = v6
    return 418
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 uint8 = 243
    v4 := make(map[uint32]int)
    var v5 bool = (-1 > -90)
    v19 := GenS1{F0: -974, F1: v5}
    if v5 {
        v20 := GenS3{F0: 0, F1: 1}
        _ = v20
    }
    if (0 != 219) {
        v5 = v5
        if (true && false) {
            v5 = false
        } else {
            for fi21 := 0; fi21 < 10; fi21++ {
                v5 = (0 == 2)
                _ = fi21
            }
            var v22 int16 = -1
            v23 := GenS2{F0: v5, F1: v22}
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            fmt.Println(v5)
            _ = fi24
        }
    }
    v25 := GenS4{F0: 100, F1: -1}
    for fi26 := 0; fi26 < 7; fi26++ {
        v5 = v5
        _ = fi26
    }
    v27 := GenS4{F0: 0, F1: -291}
    if false {
        var v28 uint = 522
        if (1 <= 725) {
            if (!(497 <= 65535)) {
                v28 -= 0
            } else {
                v5 = (v5 || (v5 && (!(!v5))))
                for fi29 := 0; fi29 < 11; fi29++ {
                    var v30 float64 = (3.14159 * 1.5)
                    for fi31 := 0; fi31 < 9; fi31++ {
                        var v32 uint8 = 127
                        _ = fi31
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 11; fi33++ {
                        v34 := []int{fi29, (fi33 - (fi33 - fi29)), (fi33 + 128)}
                        v35 := GenS4{F0: -9, F1: int16(v28)}
                        fmt.Println(v3)
                        fmt.Println(v5, v30)
                        var v36 int16 = 32767
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    _ = fi29
                    _ = v30
                }
            }
        } else {
            if (18446744073709551615 >= 1) {
                if v5 {
                    var v37 uint32 = 42
                    for fi38 := 0; fi38 < 17; fi38++ {
                        for fi39 := 0; fi39 < 12; fi39++ {
                            var v40 int8 = 1
                            v40 -= v40
                            if ((^(fi38)) == (fi39 + (fi39 << ((2) & 63)))) {
                                v41 := []int{fi39}
                                var v42 int64 = 9223372036854775807
                                var v43 string = "\n\t"
                                var v44 uint8 = v3
                                _ = v41
                                _ = v42
                                _ = v43
                                _ = v44
                            } else {
                                fmt.Println(v3)
                                fmt.Println(fi38, v37, v28)
                                var v45 bool = (v5 && v5)
                                _ = v45
                            }
                            _ = fi39
                            _ = v40
                        }
                        _ = fi38
                    }
                    _ = v37
                } else {
                    var v46 float32 = (1e10 - 3.14159)
                    for fi47 := 0; fi47 < 14; fi47++ {
                        v48 := GenS1{F0: 42, F1: true}
                        _ = fi47
                        _ = v48
                    }
                    _ = v46
                }
                v28 -= (v28 | (v28 << (((v28 % (((v28 / ((4294967295) | 1))) | 1))) & 63)))
            }
            switch ("foo bar" + "foo bar") {
            case "a":
                for fi49 := 0; fi49 < 9; fi49++ {
                    v3 = v3
                    _ = fi49
                }
            case "hello":
                for fi50 := 0; fi50 < 19; fi50++ {
                    v51 := []int{32767}
                    fmt.Println(v3, fi50)
                    for fi52 := 0; fi52 < 18; fi52++ {
                        var v53 float32 = -2.25
                        switch (fi50 % (((-(fi50))) | 1)) {
                        case 32767:
                            var v54 int = (-(fi52))
                            _ = v54
                        case -2147483648:
                            v28 = 18446744073709551615
                        default:
                            v55 := GenS1{F0: -928, F1: (false && (!((false || false) && (!((v3 ^ v3) > (v3 >> ((189) & 63)))))))}
                            _ = v55
                        }
                        var v56 uint64 = 4294967295
                        _ = fi52
                        _ = v53
                        _ = v56
                    }
                    _ = fi50
                    _ = v51
                }
            default:
                for fi57 := 0; fi57 < 20; fi57++ {
                    v5 = v5
                    _ = fi57
                }
            }
            v5 = false
        }
        _ = v28
    } else {
        var v58 float32 = ((-2.25 + -2.25) / (float32((0.1 + -2.25)) + 0.0))
        var v59 uint32 = 0
        for fi60 := 0; fi60 < 12; fi60++ {
            if (int32(v3) < -1) {
                switch "hello" {
                case "":
                    for fi61 := 0; fi61 < 1; fi61++ {
                        delete(v4, v59)
                        v3 = v3
                        v58 -= v58
                        _ = fi61
                    }
                case "a":
                    for fi62 := 0; fi62 < 6; fi62++ {
                        v63 := GenS4{F0: -5, F1: 1000}
                        _ = fi62
                        _ = v63
                    }
                case "\n\t":
                    for fi64 := 0; fi64 < 14; fi64++ {
                        if false {
                            var v65 uint64 = 2
                            _ = v65
                        }
                        _ = fi64
                    }
                }
                v66 := []int{(fi60 | 32768), 1, fi60, (fi60 << ((1) & 63))}
                v3 -= v3
                _ = v66
            } else {
                var v67 int64 = 0
                v59 -= v59
                _ = v67
            }
            _ = fi60
        }
        _ = v58
        _ = v59
    }
    for fi68 := 0; fi68 < 11; fi68++ {
        var v69 int64 = 0
        for fi70 := 0; fi70 < 6; fi70++ {
            var v71 uint32 = 42
            _ = fi70
            _ = v71
        }
        v69 = (v69 + (v69 >> ((18446744073709551615) & 63)))
        _ = fi68
        _ = v69
    }
    v5 = (!v5)
    v72 := GenS1{F0: 42, F1: v5}
    v5 = ((!(!true)) && true)
    var v73 string = "a"
    fmt.Println(v5, v5, v73)
    fmt.Println(v73, v3)
    switch ((v73 + v73) + v73) {
    case "hello":
        if (-9223372036854775808 >= -9223372036854775808) {
            if v5 {
                v5 = (!false)
                v3 -= 2
            }
            var v74 float64 = (0.1 * 3.14159)
            var v75 int8 = 127
            _ = v74
            _ = v75
        }
    }
    for fi76 := 0; fi76 < 8; fi76++ {
        var v77 uint64 = 24
        _ = fi76
        _ = v77
    }
    v5 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v73
}
