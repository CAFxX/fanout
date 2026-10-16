package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 uint; F1 uint; F2 int64 }
func genFn0(p0 int8) int16 {
    v5 := GenS1{F0: 4294967295}
    var v6 int8 = 127
    var v7 uint = 18446744073709551615
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint8, p1 uint16, p2 uint64) int8 {
    var v8 int16 = genFn0(42)
    var v9 float32 = 1e10
    var v10 uint = 42
    var v11 string = ("foo bar" + "hello")
    for fi12 := 0; fi12 < 12; fi12++ {
        p0 = (p0 | (p0 + 19))
        if (false || (false && false)) {
            v13 := GenS1{F0: uint(p2)}
            _ = v13
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 22
}

func main() {
    var v1 int = 2147483647
    v2 := []string{("" + "foo bar")}
    var v3 int = (v1 << ((882) & 63))
    var v4 int = v3
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS2{F0: 18446744073709551615, F1: -1}
        v3 -= (^(v4))
        _ = fi14
        _ = v15
    }
    fmt.Println(v3)
    v16 := GenS1{F0: 4294967295}
    if false {
        v1 -= (-(v4))
        if (!(((879 != 1) || (true || (86 < 7))) && (!(!false)))) {
            v17 := GenS3{F0: true}
            if len(v2) > 0 {
                v2[0] = ""
            }
            _ = v17
        }
    }
    if true {
        for fi18 := 0; fi18 < 8; fi18++ {
            v3 += v4
            fmt.Println(v4, v3, v3)
            _ = fi18
        }
        if false {
            var v19 int64 = 9223372036854775807
            v1 = (v4 % (((v4 ^ (v3 - (v3 % (((v1 / ((v3) | 1))) | 1))))) | 1))
            fmt.Println(v4, v19)
            _ = v19
        } else {
            v4 = (v3 % (((v1 << ((171) & 63))) | 1))
            if (((-1 >= -1) || ((!false) && (true && true))) && false) {
                v20 := GenS2{F0: 42, F1: (v1 * (v1 % ((v3) | 1)))}
                var v21 uint8 = 127
                for fi22 := 0; fi22 < 10; fi22++ {
                    if true {
                        var v23 uint8 = v21
                        var v24 uint64 = 42
                        v25 := GenS4{F0: "foo bar"}
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    _ = fi22
                }
                _ = v20
                _ = v21
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if false {
                    v26 := GenS5{F0: 0, F1: 0, F2: 42}
                    if false {
                        var v27 string = "a"
                        if (float64(3.14159) <= (3.14159 * 1.5)) {
                            var v28 uint16 = 65535
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v28
                        } else {
                            fmt.Println(v4, v27, v4)
                            switch ("hello" + v27) {
                            case "a":
                                fmt.Println(v3, v3, v3)
                            default:
                                v3 += v4
                            }
                            v4 -= (v1 & v4)
                        }
                        _ = v27
                    } else {
                        for fi29 := 0; fi29 < 18; fi29++ {
                            if true {
                                v4 += v4
                                if true {
                                    var v30 int64 = 0
                                    v31 := GenS4{F0: ""}
                                    _ = v30
                                    _ = v31
                                } else {
                                    for fi32 := 0; fi32 < 6; fi32++ {
                                        var v33 uint64 = 1
                                        _ = fi32
                                        _ = v33
                                    }
                                }
                                var v34 uint8 = 17
                                _ = v34
                            }
                            _ = fi29
                        }
                        if false {
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                            v3 = v4
                        } else {
                            v4 = -550
                        }
                        var v35 string = "a"
                        _ = v35
                    }
                    var v36 int = -118
                    _ = v26
                    _ = v36
                } else {
                    var v37 uint16 = 1
                    if false {
                        v4 = v4
                        var v38 float32 = float32(-2.25)
                        _ = v38
                    }
                    v39 := GenS4{F0: "hello"}
                    _ = v37
                    _ = v39
                }
                v2 = append(v2, "foo bar")
            }
            var v40 float32 = 3.14159
            _ = v40
        }
        var v41 uint8 = 1
        _ = v41
    }
    v1 = v4
    for fi42 := 0; fi42 < 12; fi42++ {
        var v43 uint16 = 2
        v44 := GenS4{F0: ("\n\t" + ("" + (("a" + "a") + "hello")))}
        v45 := GenS6{F0: 32767}
        _ = fi42
        _ = v43
        _ = v44
        _ = v45
    }
    _ = len(v2)
    if true {
        _ = len(v2)
        var v46 uint = 4294967295
        _ = v46
    } else {
        if false {
            var v47 int8 = 66
            if true {
                var v48 uint64 = 401
                _ = v48
            } else {
                v49 := GenS6{F0: genFn0((v47 + v47))}
                var v50 uint16 = 32767
                _ = v49
                _ = v50
            }
            for fi51 := 0; fi51 < 8; fi51++ {
                var v52 uint64 = 18446744073709551615
                _ = fi51
                _ = v52
            }
            _ = v47
        }
        if len(v2) > 0 {
            v2[0] = "a"
        }
        if (false && (false && (!false))) {
            if true {
                v53 := GenS4{F0: ("" + "hello")}
                _ = v53
            }
            var v54 uint32 = 2147483648
            v55 := GenS4{F0: ("\n\t" + ("\n\t" + "foo bar"))}
            _ = v54
            _ = v55
        }
    }
    v56 := GenS7{F0: -1, F1: 345}
    fmt.Println(v3, v4, v3)
    for fi57 := 0; fi57 < 5; fi57++ {
        v58 := GenS1{F0: 42}
        if len(v2) > 0 {
            v2[0] = ("a" + "")
        }
        _ = fi57
        _ = v58
    }
    v59 := GenS8{F0: 18446744073709551615, F1: 42, F2: 506}
    for fi60 := 0; fi60 < 1; fi60++ {
        v61 := GenS7{F0: int64(v1), F1: 0}
        switch (("\n\t" + "\n\t") + "hello") {
        case "a":
            v3 += (v4 | v4)
        case "hello":
            v2 = append(v2, ("foo bar" + "a"))
        case "":
            var v62 uint64 = 563
            _ = v62
        }
        _ = fi60
        _ = v61
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v4, v4, v3)
    case "foo bar":
        var v63 uint8 = 147
        _ = v63
    case "\n\t":
        if (true && false) {
            var v64 int8 = -2
            var v65 int32 = 42
            var v66 uint8 = 0
            _ = v64
            _ = v65
            _ = v66
        } else {
            var v67 int8 = genFn1(0, 0, 18446744073709551615)
            var v68 int16 = -267
            _ = v67
            _ = v68
        }
    }
    fmt.Println(v1, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v59
}
