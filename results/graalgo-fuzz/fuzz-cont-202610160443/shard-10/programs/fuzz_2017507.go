package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int; F2 string }
type GenS4 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS5 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int; F1 uint; F2 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64; F1 int64; F2 uint16 }
func genFn0() uint64 {
    if false {
        var v4 float32 = 3.14159
        v5 := GenS1{F0: 1, F1: 196}
        _ = v4
        _ = v5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 bool = true
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: 453}
    var v9 int64 = 1
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 float64 = 1e10
        if false {
            fmt.Println(v9, v11, v11)
            if true {
                for fi12 := 0; fi12 < 10; fi12++ {
                    if (2 <= 255) {
                        if false {
                            v13 := []int{(fi10 + (fi12 ^ fi12)), fi10, (fi12 / ((fi12) | 1)), (fi10 << ((4294967295) & 63))}
                            var v14 int = fi10
                            var v15 string = "\n\t"
                            _ = v13
                            _ = v14
                            _ = v15
                        } else {
                            var v16 uint16 = 42
                            if false {
                                v16 -= (v16 * v16)
                                var v17 float32 = 3.14159
                                var v18 int32 = -1
                                _ = v17
                                _ = v18
                            }
                            if (!(1 == 255)) {
                                v19 := []int{32767, (fi10 / ((9223372036854775807) | 1))}
                                fmt.Println(v9, v11)
                                _ = v19
                            } else {
                                v16 += 963
                                v20 := GenS2{F0: fi12}
                                _ = v20
                            }
                            if true {
                                v21 := GenS1{F0: -2147483648, F1: -2147483648}
                                var v22 int16 = 64
                                _ = v21
                                _ = v22
                            }
                            _ = v16
                        }
                    } else {
                        var v23 uint = 42
                        _ = v23
                    }
                    var v24 string = "a"
                    _ = fi12
                    _ = v24
                }
                var v25 bool = ((("\n\t" + "hello") + "foo bar") < ("a" + ("foo bar" + ("foo bar" + ("a" + "hello")))))
                _ = v25
            } else {
                fmt.Println(v9, v11, v9)
                v26 := GenS1{F0: -848, F1: 2147483647}
                _ = v26
            }
        } else {
            v27 := []int{(fi10 & (fi10 << ((65535) & 63))), 116, (fi10 - 32768)}
            v28 := GenS2{F0: fi10}
            v11 = 1.5
            _ = v27
            _ = v28
        }
        fmt.Println(fi10, fi10, fi10)
        _ = fi10
        _ = v11
    }
    _ = v8
    _ = v9
    return 2
}
func genFn1() string {
    var v29 float64 = -2.25
    v30 := GenS3{F0: 4294967295, F1: -2147483648, F2: ""}
    v31 := GenS1{F0: 29, F1: 0}
    v29 -= ((float64(((1e10 - 0.0) * 0.0)) - v29) * v29)
    _ = v29
    _ = v30
    _ = v31
    return "hello"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 0
    v32 := v1["hello"]
    fmt.Println(v3)
    switch ("" + "") {
    case "":
        if false {
            if true {
                v33 := GenS2{F0: -947}
                v32 -= (^(v32))
                _ = v33
            } else {
                v2 = v2
                if false {
                    switch "hello" {
                    case "hello":
                        v34 := GenS4{F0: 42, F1: 42, F2: v2}
                        _ = v34
                    default:
                        if true {
                            v35 := GenS5{F0: (-(v2)), F1: (^(v3)), F2: v32}
                            v36 := GenS1{F0: -790, F1: -726}
                            _ = v35
                            _ = v36
                        }
                    }
                    if (!(!true)) {
                        v37 := GenS6{F0: -2147483648, F1: 42, F2: v32}
                        var v38 float32 = (1.5 - -2.25)
                        _ = v37
                        _ = v38
                    }
                } else {
                    var v39 uint16 = 42
                    if (!true) {
                        var v40 int64 = (v2 - 0)
                        _ = v40
                    }
                    v39 = (v39 - v39)
                    _ = v39
                }
                if ((65535 == 65535) && true) {
                    v41 := GenS6{F0: 9223372036854775807, F1: 255, F2: (v32 << ((2) & 63))}
                    _ = v41
                }
            }
            if true {
                v3 += (v3 | v3)
                for fi42 := 0; fi42 < 12; fi42++ {
                    switch ("\n\t" + "hello") {
                    case "hello":
                        fmt.Println(v3)
                    case "a":
                        var v43 int = fi42
                        _ = v43
                    case "foo bar":
                        fmt.Println(v2)
                    }
                    _ = fi42
                }
            }
        } else {
            var v44 uint16 = 8
            v45 := GenS7{F0: (true && (-33 < -1))}
            v46 := GenS5{F0: (v2 + (-(v2))), F1: (v3 ^ uint64(v32)), F2: (^(v32))}
            _ = v44
            _ = v45
            _ = v46
        }
    default:
        v47 := v1[(("" + "foo bar") + "")]
        _ = v47
    }
    if false {
        v48 := v1[(("" + ((("foo bar" + "a") + "") + "a")) + "foo bar")]
        _ = v48
    } else {
        var v49 int = 557
        var v50 bool = false
        switch 32767 {
        case 125:
            if (0 != 1) {
                var v51 int = v49
                if (!((v51 + (v51 + v51)) >= (v49 ^ (^(v49))))) {
                    v51 += (v49 >> ((42) & 63))
                    v2 += (v2 | -1)
                    v52 := GenS5{F0: -1, F1: v3, F2: 65535}
                    _ = v52
                }
                _ = v51
            } else {
                v53 := []int{42, v49}
                for fi54 := 0; fi54 < 8; fi54++ {
                    var v55 int8 = 42
                    var v56 uint32 = (v32 * v32)
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                if (42 == 32767) {
                    var v57 int64 = (-(v2))
                    var v58 bool = v50
                    _ = v57
                    _ = v58
                } else {
                    var v59 uint8 = 42
                    _ = v59
                }
                _ = v53
            }
        case 2147483647:
            fmt.Println(v3, v2, v32)
        case 1:
            fmt.Println(v32, v50, v3)
        default:
            var v60 uint8 = 255
            _ = v60
        }
        _ = v49
        _ = v50
    }
    var v61 float32 = float32(1.5)
    for fi62 := 0; fi62 < 2; fi62++ {
        var v63 uint8 = 128
        if (!(int8(v32) == 0)) {
            var v64 uint = 669
            _ = v64
        }
        _ = fi62
        _ = v63
    }
    for fi65 := 0; fi65 < 5; fi65++ {
        v66 := []int{(fi65 - (fi65 + (fi65 ^ 4294967295)))}
        switch (fi65 / ((fi65) | 1)) {
        case 42:
            v32 = 2
        }
        _ = fi65
        _ = v66
    }
    if (false || true) {
        v3 = 508
    } else {
        fmt.Println(v3, v32)
    }
    v32 = 2
    switch "hello" {
    case "foo bar":
        v67 := []int{1}
        _ = v67
    case "\n\t":
        v61 -= v61
    default:
        for fi68 := 0; fi68 < 14; fi68++ {
            if false {
                var v69 uint16 = 664
                _ = v69
            } else {
                var v70 int8 = 1
                var v71 int8 = v70
                _ = v70
                _ = v71
            }
            v61 -= (float32(0.1) / ((v61 / 3.14159) / v61))
            _ = fi68
        }
    }
    for fi72 := 0; fi72 < 4; fi72++ {
        v73 := GenS4{F0: 32767, F1: 32767, F2: (v2 ^ (v2 & (v2 + (v2 * (-(v2))))))}
        v74 := []int{}
        v61 -= 0.0
        _ = fi72
        _ = v73
        _ = v74
    }
    v75 := GenS8{F0: 0}
    var v76 uint16 = 16
    v61 = (v61 + v61)
    var v77 uint32 = 51
    v61 = v61
    v78 := GenS9{F0: (v3 - (v3 + 0)), F1: v2, F2: 65535}
    v79 := GenS5{F0: (v2 & -1), F1: (v3 % ((v3) | 1)), F2: (v77 & uint32(v3))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v61:", v61)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v61
    _ = v75
    _ = v76
    _ = v77
    _ = v78
    _ = v79
}
