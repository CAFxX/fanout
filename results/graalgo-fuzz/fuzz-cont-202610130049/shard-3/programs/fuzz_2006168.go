package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint }
func genFn0(p0 uint16) int16 {
    p0 = (p0 % ((563) | 1))
    var v5 uint16 = (p0 | p0)
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 uint16, p1 int16) int {
    v6 := []int{0, -9223372036854775808, 449}
    for fi7 := 0; fi7 < 4; fi7++ {
        if len(v6) > 0 {
            v6[0] = (^(fi7))
        }
        switch "\n\t" {
        case "\n\t":
            p1 += p1
        case "a":
            for fi8 := 0; fi8 < 16; fi8++ {
                switch "a" {
                case "foo bar":
                    v6 = append(v6, (-(fi7)))
                default:
                    if false {
                        var v9 int = (^(fi7))
                        _ = v9
                    }
                }
                p0 = (^(p0))
                if (("" + "") != "hello") {
                    if (635 == 4294967295) {
                        fmt.Println(p1, p0, p0)
                    } else {
                        v10 := GenS1{F0: 42}
                        _ = v10
                    }
                }
                p0 -= (^(p0))
                if len(v6) > 0 {
                    v6[0] = (fi7 + (-(fi8)))
                }
                _ = fi8
            }
        }
        _ = fi7
    }
    p0 -= 445
    var v11 uint32 = 47
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 int = 265
    var v4 string = "hello"
    v12 := GenS2{F0: 0}
    delete(v1, "foo bar")
    switch (v3 / ((genFn1(1, 284)) | 1)) {
    case 127:
        var v13 int = (v3 >> ((534) & 63))
        _ = v13
    }
    v2[((v4 + v4) + (("foo bar" + v4) + "a"))] = true
    var v14 int16 = 0
    for fi15 := 0; fi15 < 7; fi15++ {
        if false {
            _ = len(v2)
        } else {
            v14 = (v14 << ((65535) & 63))
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(fi15, v3, fi16)
                _ = fi16
            }
        }
        v17 := []int{}
        if len(v17) > 0 {
            v17[0] = (^(v3))
        }
        _ = fi15
        _ = v17
    }
    v18 := GenS1{F0: 213}
    for fi19 := 0; fi19 < 19; fi19++ {
        if true {
            if true {
                for fi20 := 0; fi20 < 19; fi20++ {
                    fmt.Println(v4)
                    if ((42 == 42) || (!true)) {
                        var v21 uint32 = 0
                        if true {
                            v4 = v4
                            v22 := GenS1{F0: 183}
                            var v23 string = v4
                            _ = v22
                            _ = v23
                        } else {
                            v24 := []int{(fi20 * (v3 | fi20)), genFn1(255, v14), fi19}
                            var v25 uint64 = 2
                            _ = v24
                            _ = v25
                        }
                        _ = v21
                    } else {
                        switch -2147483648 {
                        case 0:
                            var v26 uint = 122
                            _ = v26
                        default:
                            switch v4 {
                            case "a":
                                var v27 float32 = float32(3.14159)
                                _ = v27
                            case "foo bar":
                                var v28 bool = true
                                _ = v28
                            case "hello":
                                v29 := []int{genFn1(255, (-(v14))), (fi20 ^ (fi19 >> ((0) & 63))), (v3 | genFn1(677, v14))}
                                _ = v29
                            }
                        }
                        v14 = (v14 - (-(v14)))
                    }
                    _ = fi20
                }
                if (false && (2147483648 == 1)) {
                    fmt.Println(v14)
                } else {
                    v30 := []int{(v3 << ((905) & 63)), (-(v3)), v3}
                    _ = v30
                }
            } else {
                if true {
                    for fi31 := 0; fi31 < 15; fi31++ {
                        v32 := GenS2{F0: 42}
                        v33 := GenS3{F0: 42, F1: 32767}
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    switch (-(v3)) {
                    case 303:
                        v14 = genFn0(0)
                    default:
                        for fi34 := 0; fi34 < 13; fi34++ {
                            v35 := []int{(fi34 | fi19), genFn1(1, v14)}
                            switch "" {
                            case "foo bar":
                                _ = len(v35)
                            case "a":
                                if true {
                                    var v36 uint64 = 18446744073709551615
                                    v35 = append(v35, (-(fi19)))
                                    if len(v35) > 0 {
                                        v35[0] = 127
                                    }
                                    switch 4294967295 {
                                    case -2147483648:
                                        fmt.Println(v4)
                                    case -128:
                                        v37 := v2["hello"]
                                        _ = v37
                                    case -797:
                                        _ = len(v35)
                                    }
                                    _ = v36
                                } else {
                                    delete(v1, v4)
                                    v14 += (v14 % ((genFn0(32767)) | 1))
                                }
                            case "\n\t":
                                switch (v3 / ((787) | 1)) {
                                case -1:
                                    var v38 float64 = -2.25
                                    _ = v38
                                case 32768:
                                    v39 := GenS4{F0: (-(v14)), F1: (!false)}
                                    _ = v39
                                case 42:
                                    var v40 float32 = 0.0
                                    _ = v40
                                }
                            }
                            v41 := GenS5{F0: (v4 + ""), F1: 1}
                            _ = fi34
                            _ = v35
                            _ = v41
                        }
                    }
                } else {
                    fmt.Println(v14, v14)
                }
                if true {
                    v42 := GenS6{F0: 255}
                    var v43 int16 = (v14 / ((v14) | 1))
                    v43 = genFn0(646)
                    _ = v42
                    _ = v43
                } else {
                    fmt.Println(v14)
                }
                v4 = v4
            }
            switch ((v4 + (("" + (v4 + v4)) + v4)) + "foo bar") {
            case "\n\t":
                switch (-(fi19)) {
                case 9223372036854775807:
                    v4 = ((v4 + "foo bar") + "a")
                case 127:
                    var v44 uint32 = 4294967295
                    _ = v44
                }
            }
            v45 := []int{(v3 | (v3 | (^(v3)))), v3, (fi19 ^ (v3 | v3)), (v3 / (((v3 ^ (v3 % (((-(fi19))) | 1)))) | 1))}
            _ = v45
        }
        v46 := []int{(v3 ^ (v3 & v3)), genFn1(894, v14), (v3 ^ 0)}
        v46 = append(v46, v3)
        _ = fi19
        _ = v46
    }
    if (9223372036854775807 > -1) {
        v47 := []int{(v3 | (v3 / ((genFn1(42, v14)) | 1)))}
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
