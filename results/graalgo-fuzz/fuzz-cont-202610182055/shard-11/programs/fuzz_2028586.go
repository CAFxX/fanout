package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8; F1 int64 }
type GenS9 struct { F0 int32; F1 bool }
type GenS10 struct { F0 int32; F1 int }
func genFn0(p0 int16, p1 uint32) int8 {
    var v4 uint32 = 4294967295
    v5 := GenS1{F0: (^(p1)), F1: "foo bar"}
    if true {
        var v6 uint = 26
        v7 := GenS2{F0: "foo bar"}
        if false {
            var v8 float32 = 3.14159
            _ = v8
        }
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 115
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 116
    var v3 int16 = -1
    for fi9 := 0; fi9 < 5; fi9++ {
        v2 += v2
        _ = fi9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        v2 = -9223372036854775808
        if ((2147483647 > 42) && (true || (!true))) {
            v11 := GenS3{F0: 255}
            var v12 int8 = -60
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := GenS4{F0: 282}
                switch (("\n\t" + (("hello" + "") + "a")) + "\n\t") {
                case "":
                    fmt.Println(v3)
                case "a":
                    fmt.Println(fi10, v2, fi10)
                default:
                    if ((^(fi10)) <= (^(v2))) {
                        v15 := GenS2{F0: ""}
                        fmt.Println(v2, fi10, fi10)
                        fmt.Println(v3, v2, fi13)
                        _ = v15
                    }
                }
                fmt.Println(fi10, v3, fi13)
                _ = fi13
                _ = v14
            }
            var v16 int16 = (v3 << ((255) & 63))
            if false {
                v17 := GenS5{F0: -128}
                _ = v17
            } else {
                v16 = v3
                for fi18 := 0; fi18 < 18; fi18++ {
                    v16 -= v3
                    _ = fi18
                }
            }
            _ = v16
        }
        _ = fi10
    }
    var v19 int8 = 0
    for fi20 := 0; fi20 < 4; fi20++ {
        v19 -= v19
        _ = fi20
    }
    if (false || (((v19 > 42) || true) && false)) {
        if false {
            switch int(v3) {
            case 2147483647:
                for fi21 := 0; fi21 < 13; fi21++ {
                    v22 := GenS6{F0: 42}
                    var v23 uint32 = 2147483647
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            case 400:
                for fi24 := 0; fi24 < 16; fi24++ {
                    for fi25 := 0; fi25 < 1; fi25++ {
                        var v26 uint32 = 2147483647
                        v26 += v26
                        var v27 int = (v2 & (v2 + fi24))
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    _ = fi24
                }
            case 0:
                v28 := GenS2{F0: "foo bar"}
                _ = v28
            default:
                v29 := []int{v2, 1, (v2 & (-(v2)))}
                _ = v29
            }
        } else {
            v19 -= (v19 / ((v19) | 1))
            if true {
                var v30 bool = (float64(float32(0.0)) <= (1.5 / 1e10))
                _ = v30
            }
            switch (v2 + v2) {
            case -2147483648:
                v31 := []int{(^(v2)), 1, v2, (-(v2))}
                _ = v31
            case -929:
                if (!true) {
                    for fi32 := 0; fi32 < 1; fi32++ {
                        var v33 int16 = (^(v3))
                        var v34 uint32 = 2
                        fmt.Println(v19)
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    var v35 int32 = 1
                    switch (v2 >> ((2) & 63)) {
                    case 0:
                        var v36 int16 = -1
                        _ = v36
                    }
                    _ = v35
                }
            default:
                v3 = v3
            }
        }
        if ((v3 != 430) || (!(("" + (("hello" + "foo bar") + "a")) == ((("a" + "foo bar") + "foo bar") + "\n\t")))) {
            var v37 int8 = 127
            if true {
                for fi38 := 0; fi38 < 3; fi38++ {
                    v3 = 175
                    switch (fi38 << ((974) & 63)) {
                    case 32768:
                        fmt.Println(v3)
                    default:
                        for fi39 := 0; fi39 < 10; fi39++ {
                            v1[65535] = (v2 ^ fi38)
                            _ = fi39
                        }
                    }
                    _ = fi38
                }
            }
            if true {
                var v40 float32 = ((3.14159 - 1e10) / 1e10)
                v41 := GenS5{F0: 1}
                _ = v40
                _ = v41
            } else {
                for fi42 := 0; fi42 < 1; fi42++ {
                    if false {
                        v3 = (v3 - (-(v3)))
                        v43 := GenS5{F0: v19}
                        var v44 int32 = 2147483647
                        _ = v43
                        _ = v44
                    } else {
                        var v45 float32 = -2.25
                        v3 = (v3 % ((v3) | 1))
                        if false {
                            v46 := v1[102]
                            for fi47 := 0; fi47 < 3; fi47++ {
                                v45 = v45
                                _ = fi47
                            }
                            _ = v46
                        } else {
                            fmt.Println(v3)
                        }
                        _ = v45
                    }
                    if false {
                        v37 = (v37 + v37)
                        var v48 int64 = 1
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 10; fi49++ {
                        var v50 uint64 = 211
                        v51 := GenS7{F0: 702}
                        var v52 int64 = 9223372036854775807
                        v53 := GenS4{F0: (v50 >> ((255) & 63))}
                        var v54 float64 = 1.5
                        v55 := GenS8{F0: (v19 / (((^(v19))) | 1)), F1: (v52 ^ (v52 | 42))}
                        _ = fi49
                        _ = v50
                        _ = v51
                        _ = v52
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                    _ = fi42
                }
                for fi56 := 0; fi56 < 2; fi56++ {
                    fmt.Println(v2)
                    v2 += v2
                    v2 -= (fi56 << ((2) & 63))
                    _ = fi56
                }
                v37 -= v37
            }
            _ = v37
        }
    } else {
        switch (v2 ^ (^(v2))) {
        case -188:
            var v57 int8 = (v19 & v19)
            _ = v57
        default:
            for fi58 := 0; fi58 < 10; fi58++ {
                switch (("" + ("hello" + ("foo bar" + "a"))) + "") {
                case "foo bar":
                    v59 := GenS9{F0: 56, F1: (52 != 0)}
                    _ = v59
                case "hello":
                    v3 -= 1
                }
                v60 := []int{v2}
                _ = fi58
                _ = v60
            }
        }
        var v61 int64 = 907
        var v62 int64 = (v61 + (v61 ^ v61))
        _ = v61
        _ = v62
    }
    v63 := GenS10{F0: 1, F1: (-(v2))}
    v2 -= (v2 - (v2 << ((791) & 63)))
    var v64 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v63
    _ = v64
}
