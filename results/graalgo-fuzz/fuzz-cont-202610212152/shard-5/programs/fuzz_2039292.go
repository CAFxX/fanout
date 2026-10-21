package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 string; F1 int; F2 uint16 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 int16; F1 int64; F2 string }
type GenS7 struct { F0 int32 }
func genFn0(p0 uint64, p1 string, p2 bool) int32 {
    p1 = ("hello" + ("" + (("hello" + p1) + p1)))
    switch p1 {
    case "foo bar":
        v6 := GenS1{F0: "a", F1: -688, F2: 0}
        _ = v6
    case "hello":
        fmt.Println(p1)
    default:
        var v7 int16 = 42
        _ = v7
    }
    v8 := []int{}
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS2{F0: 1}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2147483647
}

func main() {
    var v1 uint8 = 0
    var v2 bool = false
    var v3 uint16 = 42
    var v4 int64 = -794
    var v5 uint16 = 255
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(v1)
        for fi12 := 0; fi12 < 19; fi12++ {
            if v2 {
                var v13 string = ""
                v14 := []int{fi12, (-(fi11))}
                _ = v13
                _ = v14
            } else {
                v15 := GenS2{F0: 1}
                v16 := []int{(fi12 / ((42) | 1)), -1, fi11, 156}
                _ = v15
                _ = v16
            }
            _ = fi12
        }
        _ = fi11
    }
    v3 = v5
    switch 4294967295 {
    case 310:
        for fi17 := 0; fi17 < 2; fi17++ {
            if v2 {
                v1 = v1
                var v18 string = "hello"
                v3 = v3
                _ = v18
            }
            v19 := []int{fi17, (fi17 - fi17), fi17, (-(fi17))}
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 float32 = (0.1 / 1.5)
                if ((fi17 * 42) == -312) {
                    var v22 int64 = 9223372036854775807
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
            _ = fi17
            _ = v19
        }
    case 32767:
        var v23 int32 = 2147483647
        _ = v23
    case 127:
        for fi24 := 0; fi24 < 5; fi24++ {
            v3 = 255
            if (!((v1 + v1) <= (v1 - 128))) {
                fmt.Println(v3)
                v25 := GenS1{F0: ("" + "a"), F1: 1000, F2: 42}
                v4 -= 42
                _ = v25
            } else {
                var v26 int8 = -32
                _ = v26
            }
            v3 -= 624
            _ = fi24
        }
    }
    if v2 {
        for fi27 := 0; fi27 < 20; fi27++ {
            v28 := GenS2{F0: -32768}
            _ = fi27
            _ = v28
        }
    } else {
        switch 32768 {
        case 127:
            if v2 {
                for fi29 := 0; fi29 < 14; fi29++ {
                    v30 := GenS2{F0: 32767}
                    _ = fi29
                    _ = v30
                }
                v2 = v2
            } else {
                var v31 uint = 214
                var v32 string = ""
                _ = v31
                _ = v32
            }
        default:
            fmt.Println(v5, v2)
        }
    }
    v33 := []int{42, 2147483647, 32767, -1}
    if v2 {
        v34 := GenS2{F0: 1}
        switch "" {
        case "foo bar":
            if len(v33) > 0 {
                v33[0] = 32767
            }
        }
        if v2 {
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
        } else {
            v4 -= v4
        }
        _ = v34
    } else {
        v35 := GenS3{F0: -1, F1: 339}
        var v36 uint64 = 1
        var v37 int8 = -103
        _ = v35
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 5; fi38++ {
        v39 := GenS2{F0: -32768}
        for fi40 := 0; fi40 < 12; fi40++ {
            v41 := GenS4{F0: (("hello" + "\n\t") + "a"), F1: (fi40 / (((fi40 | fi38)) | 1)), F2: (^(v3))}
            _ = fi40
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    v2 = v2
    for fi42 := 0; fi42 < 4; fi42++ {
        v2 = v2
        var v43 uint32 = 610
        for fi44 := 0; fi44 < 17; fi44++ {
            var v45 int8 = int8(v1)
            if (true || v2) {
                if ((0 == -1) || v2) {
                    fmt.Println(v4, v5, v5)
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                } else {
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                    fmt.Println(v5, fi42, fi44)
                }
                if len(v33) >= 2 {
                    v33 = v33[1:len(v33)]
                }
                fmt.Println(fi44, v43)
            }
            v43 = (v43 | 2147483648)
            _ = fi44
            _ = v45
        }
        _ = fi42
        _ = v43
    }
    if len(v33) >= 2 {
        v33 = v33[1:len(v33)]
    }
    if v2 {
        switch 127 {
        case 42:
            v3 -= v5
        case 32767:
            if len(v33) > 0 {
                v33[0] = 713
            }
        }
    }
    for fi46 := 0; fi46 < 17; fi46++ {
        if v2 {
            for fi47 := 0; fi47 < 17; fi47++ {
                fmt.Println(v2, v3)
                _ = fi47
            }
            var v48 string = ""
            v48 = "a"
            _ = v48
        } else {
            switch (^(fi46)) {
            case -2147483648:
                var v49 string = (("foo bar" + "") + "hello")
                _ = v49
            case -761:
                v50 := GenS3{F0: -11, F1: (v4 & v4)}
                _ = v50
            default:
                var v51 uint16 = v5
                _ = v51
            }
            var v52 int64 = (v4 % ((v4) | 1))
            _ = v52
        }
        v53 := GenS5{F0: "hello", F1: ("foo bar" + "")}
        for fi54 := 0; fi54 < 15; fi54++ {
            var v55 uint = 65535
            v2 = false
            fmt.Println(v2, v2, v4)
            _ = fi54
            _ = v55
        }
        _ = fi46
        _ = v53
    }
    for fi56 := 0; fi56 < 19; fi56++ {
        for fi57 := 0; fi57 < 7; fi57++ {
            fmt.Println(v1, v4, v1)
            if true {
                if (uint16(v1) < 104) {
                    fmt.Println(fi57)
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                    if v2 {
                        if (-374 > -32768) {
                            v58 := GenS6{F0: -32768, F1: (v4 * (v4 % (((-(v4))) | 1))), F2: ("foo bar" + "a")}
                            if v2 {
                                var v59 int = fi56
                                _ = v59
                            } else {
                                v60 := GenS6{F0: 1, F1: (v4 | 9223372036854775807), F2: "foo bar"}
                                _ = v60
                            }
                            _ = v58
                        } else {
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                        }
                        fmt.Println(fi57)
                    } else {
                        fmt.Println(v4, fi56, fi57)
                        v61 := GenS7{F0: 1}
                        _ = v61
                    }
                }
                v5 = v3
            }
            v3 -= (v3 >> ((255) & 63))
            _ = fi57
        }
        var v62 int16 = -686
        _ = fi56
        _ = v62
    }
    v1 += 36
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
