package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int16; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int64; F2 string }
type GenS5 struct { F0 uint16; F1 int8; F2 int64 }
type GenS6 struct { F0 uint16; F1 int }
func genFn0() uint8 {
    var v6 uint16 = 344
    if (466 != 42) {
        v6 += v6
        v7 := GenS1{F0: 1000}
        var v8 float32 = float32(1.5)
        _ = v7
        _ = v8
    }
    switch ("foo bar" + "") {
    case "a":
        v9 := []int{4294967295, 797, 1, 792}
        _ = v9
    default:
        fmt.Println(v6)
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v6 = (v6 - v6)
        _ = fi10
    }
    var v11 uint16 = v6
    _ = v11
    _ = v6
    return 42
}
func genFn1() uint64 {
    var v12 int64 = 1
    v13 := GenS2{F0: 345, F1: 0, F2: 32767}
    _ = v12
    _ = v13
    return 303
}

func main() {
    v1 := []int{-414, 32768}
    var v2 int8 = -37
    var v3 uint8 = 0
    var v4 int16 = 468
    v5 := make(map[int]bool)
    v2 = 42
    var v14 uint64 = 581
    v5[128] = true
    var v15 string = "foo bar"
    if (406 <= 42) {
        var v16 int64 = -1
        _ = v16
    } else {
        var v17 float32 = 0.0
        switch -1 {
        case 407:
            if (true || false) {
                v1 = append(v1, 0)
            } else {
                v5[0] = false
                var v18 int8 = v2
                switch (v15 + v15) {
                case "":
                    v19 := GenS1{F0: (v4 & v4)}
                    _ = v19
                }
                _ = v18
            }
        default:
            fmt.Println(v17)
        }
        if (!(false || false)) {
            if true {
                for fi20 := 0; fi20 < 7; fi20++ {
                    fmt.Println(v17, fi20)
                    v21 := GenS2{F0: (v14 | v14), F1: (v4 - v4), F2: 4294967295}
                    _ = fi20
                    _ = v21
                }
                var v22 int8 = v2
                fmt.Println(v15, v4)
                _ = v22
            } else {
                fmt.Println(v14, v4, v17)
                if true {
                    delete(v5, -785)
                    if (true || (!(v14 < genFn1()))) {
                        if true {
                            v23 := GenS1{F0: (v4 << ((2) & 63))}
                            v24 := GenS3{F0: v3}
                            v4 += -861
                            _ = v23
                            _ = v24
                        } else {
                            v3 = genFn0()
                            for fi25 := 0; fi25 < 1; fi25++ {
                                v26 := GenS4{F0: (fi25 & fi25), F1: 42, F2: ("" + "\n\t")}
                                v2 = (^(v2))
                                var v27 uint = 42
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        }
                    } else {
                        v2 -= (v2 << ((2) & 63))
                    }
                } else {
                    v28 := v5[-872]
                    var v29 uint32 = 42
                    _ = v28
                    _ = v29
                }
            }
            switch "foo bar" {
            case "foo bar":
                switch 407 {
                case 128:
                    switch 127 {
                    case 432:
                        if true {
                            v14 += v14
                        }
                    }
                }
            case "\n\t":
                var v30 uint64 = (^(v14))
                _ = v30
            default:
                v31 := GenS3{F0: v3}
                _ = v31
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            var v32 string = (((v15 + v15) + v15) + "a")
            for fi33 := 0; fi33 < 11; fi33++ {
                if true {
                    v4 -= v4
                    v34 := GenS3{F0: (v3 | v3)}
                    _ = v34
                } else {
                    switch "\n\t" {
                    case "hello":
                        v35 := GenS2{F0: (v14 >> ((24) & 63)), F1: (-(v4)), F2: fi33}
                        _ = v35
                    default:
                        switch 1 {
                        case 32768:
                            v36 := GenS5{F0: 390, F1: v2, F2: 9223372036854775807}
                            _ = v36
                        default:
                            v37 := GenS4{F0: (fi33 << ((566) & 63)), F1: 1, F2: ("foo bar" + v15)}
                            _ = v37
                        }
                    }
                    var v38 float32 = (v17 * (((v17 - v17) + v17) + ((v17 / v17) + (-2.25 - float32(3.14159)))))
                    v3 += (v3 / ((v3) | 1))
                    _ = v38
                }
                v39 := GenS6{F0: 11, F1: (fi33 / ((fi33) | 1))}
                _ = fi33
                _ = v39
            }
            v40 := GenS1{F0: (v4 / (((^(v4))) | 1))}
            _ = v32
            _ = v40
        }
        _ = v17
    }
    fmt.Println(v4, v2)
    _ = len(v1)
    if true {
        v41 := v5[127]
        if (("a" + (("foo bar" + "") + v15)) >= "foo bar") {
            fmt.Println(v41)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v14, v3)
        } else {
            v41 = ((65535 == 32767) || (!(42 == 1)))
            var v42 uint16 = 42
            _ = v42
        }
        _ = v41
    } else {
        fmt.Println(v14, v3, v2)
        v4 += (-(v4))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
