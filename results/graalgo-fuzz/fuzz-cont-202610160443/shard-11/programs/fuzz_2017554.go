package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 int8; F2 string }
type GenS8 struct { F0 uint32; F1 uint8 }
type GenS9 struct { F0 int8 }
func genFn0(p0 uint) int64 {
    v6 := []int{128, -356, 1}
    if false {
        var v7 bool = true
        v6 = append(v6, 0)
        _ = v7
    } else {
        p0 += 319
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint = 255
            p0 -= v9
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 uint8 = 167
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        if true {
            p0 += (p0 >> ((p0) & 63))
        }
        if len(v6) > 0 {
            v6[0] = 128
        }
        _ = fi12
    }
    var v13 int8 = -57
    _ = p0
    _ = v13
    _ = v6
    return int64(v13)
}
func genFn1(p0 uint64) bool {
    if true {
        var v14 int8 = 117
        v15 := []int{}
        _ = v14
        _ = v15
    }
    p0 = (p0 | p0)
    p0 = p0
    _ = p0
    return (("\n\t" + ("a" + "")) < "a")
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]uint32)
    var v3 int16 = 909
    v4 := []string{(("foo bar" + (("" + "foo bar") + "")) + ("a" + "")), "hello", ("foo bar" + "hello"), "foo bar"}
    var v5 bool = (!(3.14159 != 3.14159))
    v16 := GenS1{F0: 263, F1: "\n\t"}
    v5 = (110 != 65535)
    var v17 uint16 = 0
    var v18 uint32 = 2147483648
    v17 = v17
    switch "a" {
    case "a":
        if v5 {
            v19 := v2[""]
            _ = v19
        } else {
            for fi20 := 0; fi20 < 11; fi20++ {
                fmt.Println(v5, v17)
                _ = fi20
            }
            for fi21 := 0; fi21 < 20; fi21++ {
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 uint16 = (v17 ^ v17)
                    var v24 string = "a"
                    if v5 {
                        if (!(0.0 > (-2.25 + 3.14159))) {
                            v2[v24] = 269
                            v25 := GenS2{F0: 2147483647, F1: 804, F2: 42}
                            v26 := GenS1{F0: 545, F1: "foo bar"}
                            v27 := GenS3{F0: 230, F1: 18446744073709551615}
                            v1[(^(v18))] = fi21
                            var v28 int32 = -1
                            v4 = append(v4, (((v24 + ("hello" + v24)) + "hello") + v24))
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            var v29 float64 = -2.25
                            var v30 int16 = 0
                            _ = v29
                            _ = v30
                        }
                        switch (fi21 * -431) {
                        case 714:
                            v5 = ((1 == uint64(v18)) || v5)
                        case 2147483647:
                            v5 = (true || (65535 <= 4294967295))
                        case 219:
                            v31 := GenS4{F0: 42, F1: 2, F2: 1}
                            _ = v31
                        }
                        switch fi21 {
                        case -81:
                            switch ("" + v24) {
                            case "hello":
                                v24 = ""
                                if ("a" <= ("hello" + (("" + v24) + "\n\t"))) {
                                    v32 := GenS2{F0: 0, F1: 4294967295, F2: 127}
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v33 uint8 = 0
                                    if false {
                                        v34 := GenS5{F0: 4294967295, F1: v18, F2: 255}
                                        _ = v34
                                    }
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v35 float64 = -2.25
                                    _ = v35
                                }
                                fmt.Println(v5, fi21)
                            default:
                                v18 = v18
                            }
                        case -1:
                            fmt.Println(v3)
                        case 42:
                            var v36 uint8 = 0
                            _ = v36
                        }
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = fi21
            }
        }
    }
    if ((v5 && v5) || (-1 == 42)) {
        if (42 == 8) {
            v37 := GenS4{F0: uint(v17), F1: 255, F2: 32767}
            if v5 {
                fmt.Println(v17)
            }
            v5 = v5
            _ = v37
        } else {
            for fi38 := 0; fi38 < 13; fi38++ {
                if v5 {
                    v39 := v2[(("foo bar" + ("" + "foo bar")) + ("hello" + "\n\t"))]
                    var v40 uint64 = 301
                    v18 = (v18 << ((18446744073709551615) & 63))
                    _ = v39
                    _ = v40
                }
                _ = fi38
            }
            var v41 int64 = 9223372036854775807
            _ = v41
        }
        fmt.Println(v3, v5)
        var v42 int16 = 42
        _ = v42
    } else {
        for fi43 := 0; fi43 < 12; fi43++ {
            for fi44 := 0; fi44 < 7; fi44++ {
                v45 := GenS6{F0: 42}
                v4 = append(v4, "\n\t")
                _ = fi44
                _ = v45
            }
            var v46 float64 = (-2.25 - -2.25)
            _ = fi43
            _ = v46
        }
        fmt.Println(v18)
    }
    if (v18 > (v18 | 65535)) {
        v47 := GenS7{F0: 42, F1: 0, F2: "hello"}
        if ((905 == 42) || v5) {
            var v48 int8 = -88
            v49 := GenS8{F0: v18, F1: 1}
            for fi50 := 0; fi50 < 9; fi50++ {
                v18 = (^(v18))
                _ = fi50
            }
            _ = v48
            _ = v49
        }
        _ = v47
    } else {
        var v51 uint = 64
        switch 9223372036854775807 {
        case 32768:
            fmt.Println(v3, v51, v3)
        case 862:
            v52 := GenS4{F0: (v51 + v51), F1: 18446744073709551615, F2: 0}
            _ = v52
        default:
            fmt.Println(v18)
        }
        _ = v51
    }
    var v53 int = 4294967295
    delete(v2, ("hello" + "hello"))
    delete(v2, "\n\t")
    if (("" + "hello") <= ("foo bar" + "hello")) {
        var v54 string = "\n\t"
        var v55 int16 = -1
        _ = v54
        _ = v55
    }
    v56 := GenS9{F0: 1}
    v17 = v17
    v57 := GenS4{F0: 1, F1: 917, F2: (v17 % (((^(v17))) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v56
    _ = v57
}
