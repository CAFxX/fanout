package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int; F1 uint16; F2 uint16 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 bool }

func main() {
    var v1 int8 = -1
    v2 := make(map[int]uint32)
    var v3 uint = 42
    var v4 uint32 = 0
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(v3, v1, v4)
        v1 = (v1 ^ (v1 ^ 26))
        _ = fi5
    }
    if (289 < v4) {
        switch 96 {
        case 128:
            v4 = 584
        case 9223372036854775807:
            v6 := v2[4294967295]
            _ = v6
        case 1:
            var v7 int64 = -452
            _ = v7
        default:
            if (-1 >= 0) {
                var v8 uint16 = 42
                fmt.Println(v4)
                _ = v8
            } else {
                for fi9 := 0; fi9 < 17; fi9++ {
                    for fi10 := 0; fi10 < 19; fi10++ {
                        if false {
                            var v11 uint = (^(v3))
                            fmt.Println(v1, v3, fi10)
                            v12 := GenS1{F0: 32767, F1: -229}
                            _ = v11
                            _ = v12
                        }
                        _ = fi10
                    }
                    v3 -= (^(v3))
                    v3 = v3
                    _ = fi9
                }
                v13 := []int{}
                _ = v13
            }
        }
    } else {
        var v14 uint16 = 777
        v1 = -1
        _ = v14
    }
    v1 += 1
    if true {
        if ((v1 + (^(v1))) >= (v1 / (((v1 % ((-128) | 1))) | 1))) {
            if (!true) {
                for fi15 := 0; fi15 < 3; fi15++ {
                    v16 := GenS1{F0: 714, F1: 522}
                    _ = fi15
                    _ = v16
                }
                if (0.0 <= -2.25) {
                    v17 := GenS2{F0: 0, F1: 42, F2: 216}
                    var v18 float32 = (0.1 * float32(1e10))
                    _ = v17
                    _ = v18
                } else {
                    v19 := GenS3{F0: "foo bar", F1: (^(v3))}
                    _ = v19
                }
            } else {
                var v20 int16 = 0
                delete(v2, 2147483647)
                _ = v20
            }
            for fi21 := 0; fi21 < 14; fi21++ {
                if (893 >= int(v3)) {
                    if false {
                        v22 := []int{}
                        var v23 int32 = -1
                        _ = v22
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 8; fi24++ {
                            if false {
                                var v25 float64 = 0.1
                                _ = v25
                            } else {
                                delete(v2, fi21)
                                var v26 float32 = 1e10
                                v3 = (v3 / (((v3 / ((v3) | 1))) | 1))
                                _ = v26
                            }
                            var v27 uint8 = 1
                            _ = fi24
                            _ = v27
                        }
                    }
                    var v28 uint = v3
                    for fi29 := 0; fi29 < 13; fi29++ {
                        v3 -= (^(v3))
                        v2[fi21] = (^(v4))
                        _ = fi29
                    }
                    _ = v28
                } else {
                    v30 := GenS4{F0: v1}
                    v2[(fi21 & fi21)] = (^(v4))
                    for fi31 := 0; fi31 < 5; fi31++ {
                        var v32 int8 = v1
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                }
                if (!("\n\t" <= (("foo bar" + "\n\t") + "hello"))) {
                    v1 = -93
                    for fi33 := 0; fi33 < 4; fi33++ {
                        v1 += (v1 & (v1 * v1))
                        _ = fi33
                    }
                } else {
                    var v34 int8 = v1
                    switch fi21 {
                    case 426:
                        v1 -= 28
                    case -2147483648:
                        var v35 int64 = -483
                        _ = v35
                    default:
                        v1 = 100
                    }
                    _ = v34
                }
                fmt.Println(v1, v1, v4)
                _ = fi21
            }
        } else {
            if true {
                if false {
                    v36 := []int{}
                    _ = v36
                }
                if (!false) {
                    if ((-85 == 323) && (v1 < (-(v1)))) {
                        v3 = (^(v3))
                        v1 -= v1
                    } else {
                        v4 -= (^(v4))
                        var v37 string = "hello"
                        _ = v37
                    }
                    var v38 uint32 = v4
                    for fi39 := 0; fi39 < 11; fi39++ {
                        v38 += (v4 & 503)
                        v4 -= v4
                        var v40 uint32 = v38
                        _ = fi39
                        _ = v40
                    }
                    _ = v38
                }
                v41 := v2[324]
                _ = v41
            } else {
                var v42 bool = (false && true)
                switch (("a" + ("hello" + ("\n\t" + ("hello" + "foo bar")))) + "foo bar") {
                case "foo bar":
                    v42 = (false && (v42 || v42))
                case "\n\t":
                    v43 := GenS3{F0: ("" + "a"), F1: v3}
                    _ = v43
                default:
                    v42 = (82 > 128)
                }
                _ = v42
            }
            if false {
                if (1e10 <= ((3.14159 / 0.1) * 3.14159)) {
                    fmt.Println(v1)
                }
                v44 := GenS3{F0: ("foo bar" + ""), F1: (^(v3))}
                v2[-817] = (v4 & 4294967295)
                _ = v44
            } else {
                fmt.Println(v1, v3, v3)
                var v45 uint32 = 42
                _ = v45
            }
        }
        if (-2147483648 < 53) {
            v46 := GenS2{F0: -1, F1: 24, F2: 65535}
            for fi47 := 0; fi47 < 17; fi47++ {
                v3 += (v3 & (v3 | v3))
                _ = fi47
            }
            _ = v46
        } else {
            if false {
                fmt.Println(v3, v3, v4)
                for fi48 := 0; fi48 < 13; fi48++ {
                    v49 := GenS5{F0: "foo bar", F1: v4}
                    _ = fi48
                    _ = v49
                }
            } else {
                v50 := GenS6{F0: ((v3 | v3) >= v3)}
                fmt.Println(v4)
                for fi51 := 0; fi51 < 10; fi51++ {
                    v52 := GenS5{F0: ("" + ("hello" + "\n\t")), F1: 2}
                    var v53 int8 = (v1 + (v1 - (^(v1))))
                    v3 = (^(v3))
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                _ = v50
            }
            v4 -= v4
        }
    } else {
        v3 += (v3 + 65535)
    }
    var v54 uint64 = 487
    fmt.Println(v54)
    fmt.Println(v54)
    var v55 string = "foo bar"
    var v56 int8 = 42
    var v57 float64 = 0.1
    var v58 uint32 = (v4 / ((v4) | 1))
    var v59 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
