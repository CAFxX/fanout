package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 int16 }
func genFn0() uint {
    v6 := []int{42, 1, 127, -954}
    v6 = append(v6, 0)
    v7 := GenS1{F0: 42, F1: "foo bar"}
    var v8 int32 = -611
    _ = v6
    _ = v7
    _ = v8
    return 65535
}
func genFn1() int {
    var v9 float64 = -2.25
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int32 = -2147483648
        fmt.Println(v11)
        _ = fi10
        _ = v11
    }
    _ = v9
    return -383
}

func main() {
    var v1 float32 = 0.0
    var v2 int16 = 0
    v3 := []uint32{1}
    var v4 uint8 = 255
    v5 := []string{"", "\n\t", "a"}
    v2 = v2
    var v12 uint16 = 763
    v2 += (v2 | v2)
    v13 := GenS1{F0: genFn0(), F1: "a"}
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    var v14 float32 = (v1 * 3.14159)
    if (genFn0() <= 65535) {
        if ((!(1 > 42)) || (true && (34 != 100))) {
            v12 = (v12 % ((v12) | 1))
            if (genFn1() > 9223372036854775807) {
                _ = len(v3)
                var v15 uint32 = 0
                var v16 float64 = (-2.25 / (1e10 + ((0.1 / (3.14159 / (3.14159 / 1e10))) - (0.0 / 1.0))))
                _ = v15
                _ = v16
            } else {
                var v17 int16 = (-(v2))
                for fi18 := 0; fi18 < 12; fi18++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v19 := GenS1{F0: genFn0(), F1: "\n\t"}
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                v21 := GenS2{F0: 120}
                v4 += v4
                _ = fi20
                _ = v21
            }
        } else {
            if false {
                _ = len(v5)
                v22 := GenS1{F0: 297, F1: "hello"}
                _ = v22
            } else {
                var v23 float32 = v14
                if (false || (-1 == 2147483647)) {
                    var v24 float64 = 0.1
                    v2 -= (v2 >> ((1) & 63))
                    for fi25 := 0; fi25 < 16; fi25++ {
                        v26 := GenS3{F0: 371, F1: ((((((("\n\t" + "hello") + "\n\t") + "foo bar") + "a") + "a") + "foo bar") + ("a" + ""))}
                        _ = fi25
                        _ = v26
                    }
                    _ = v24
                } else {
                    for fi27 := 0; fi27 < 10; fi27++ {
                        if true {
                            for fi28 := 0; fi28 < 11; fi28++ {
                                v29 := GenS4{F0: -1}
                                v30 := GenS4{F0: 32767}
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                            var v31 string = "hello"
                            _ = v31
                        } else {
                            if false {
                                var v32 float32 = 1e10
                                switch -85 {
                                case -2147483648:
                                    for fi33 := 0; fi33 < 12; fi33++ {
                                        v12 -= v12
                                        _ = fi33
                                    }
                                default:
                                    var v34 string = ""
                                    _ = v34
                                }
                                _ = v32
                            } else {
                                v1 = (v23 + 3.14159)
                                fmt.Println(v23, v14, v12)
                                var v35 int32 = -442
                                _ = v35
                            }
                            if (!((!false) && (float64(0.1) < 1.5))) {
                                var v36 int64 = -835
                                _ = v36
                            }
                            v14 += 1.5
                        }
                        _ = fi27
                    }
                    v37 := GenS5{F0: 42}
                    _ = v37
                }
                v38 := GenS1{F0: 673, F1: ("" + (("foo bar" + ((("" + "a") + ("a" + (("" + "\n\t") + "foo bar"))) + "\n\t")) + (("" + "foo bar") + (("a" + "\n\t") + "a"))))}
                _ = v23
                _ = v38
            }
            v14 -= 0.0
            var v39 uint64 = 1
            _ = v39
        }
    } else {
        var v40 int32 = 999
        v14 -= v1
        _ = v40
    }
    var v41 int16 = v2
    for fi42 := 0; fi42 < 11; fi42++ {
        fmt.Println(v12, v14, v12)
        fmt.Println(v12, fi42)
        _ = fi42
    }
    switch ("a" + "foo bar") {
    case "foo bar":
        for fi43 := 0; fi43 < 1; fi43++ {
            v14 = 1e10
            var v44 uint8 = 42
            for fi45 := 0; fi45 < 5; fi45++ {
                v41 += v41
                var v46 uint8 = (v4 % ((1) | 1))
                v4 += v44
                v5 = append(v5, "")
                _ = fi45
                _ = v46
            }
            _ = fi43
            _ = v44
        }
    case "hello":
        if len(v3) > 0 {
            v3[0] = 1
        }
    case "\n\t":
        fmt.Println(v2, v14, v14)
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        switch (fi47 * (-(fi47))) {
        case -2147483648:
            var v48 uint64 = 245
            _ = v48
        default:
            if false {
                if false {
                    v41 += 840
                    if true {
                        var v49 uint16 = (v12 + 53)
                        _ = v49
                    } else {
                        v41 -= (-(v2))
                        for fi50 := 0; fi50 < 11; fi50++ {
                            if len(v5) > 0 {
                                v5[0] = ("" + "")
                            }
                            _ = fi50
                        }
                        fmt.Println(v12, v1)
                    }
                    if ((3.14159 + ((1.5 * float64(float32((-2.25 * 0.1)))) + 1e10)) >= ((1e10 * 0.0) / (0.1 / -2.25))) {
                        v12 = 65535
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v2 = (-(v41))
                    } else {
                        v14 += float32(1e10)
                        for fi51 := 0; fi51 < 20; fi51++ {
                            v52 := GenS6{F0: v2, F1: (v2 - -201)}
                            _ = fi51
                            _ = v52
                        }
                        for fi53 := 0; fi53 < 13; fi53++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v54 uint32 = 65535
                            _ = fi53
                            _ = v54
                        }
                    }
                } else {
                    var v55 int16 = (v41 * -32768)
                    _ = v55
                }
                fmt.Println(v4)
                var v56 float32 = v14
                _ = v56
            }
        }
        var v57 int16 = v2
        _ = fi47
        _ = v57
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v58 int32 = -422
    switch (("" + ("a" + "a")) + ("hello" + "hello")) {
    case "hello":
        for fi59 := 0; fi59 < 19; fi59++ {
            var v60 int8 = 43
            v4 -= (v4 * (v4 % ((42) | 1)))
            if (((0.1 / (1e10 / -2.25)) + 0.1) >= 1e10) {
                var v61 uint16 = (^(v12))
                var v62 int32 = v58
                v63 := GenS7{F0: (v2 % ((v2) | 1))}
                _ = v61
                _ = v62
                _ = v63
            } else {
                for fi64 := 0; fi64 < 6; fi64++ {
                    var v65 uint8 = (v4 ^ v4)
                    _ = fi64
                    _ = v65
                }
                for fi66 := 0; fi66 < 13; fi66++ {
                    var v67 int8 = (v60 * (v60 + v60))
                    v2 -= v41
                    _ = fi66
                    _ = v67
                }
            }
            _ = fi59
            _ = v60
        }
    }
    switch ("hello" + "a") {
    case "\n\t":
        v2 -= (^(v2))
    case "a":
        v14 = 1.5
    default:
        if len(v5) > 0 {
            v5[0] = ("a" + "hello")
        }
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v58
}
