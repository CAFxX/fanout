package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 string; F2 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (1 < 923)
    var v3 uint16 = 152
    var v4 int8 = -128
    var v5 int8 = -96
    fmt.Println(v2)
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: 19, F1: 72, F2: v5}
        v3 = v3
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 2
    for fi9 := 0; fi9 < 3; fi9++ {
        if (439 <= 12) {
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := GenS2{F0: v2, F1: (^(fi10))}
                var v12 int32 = -723
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 bool = v2
            fmt.Println(v8)
            _ = v13
        } else {
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := []int{9223372036854775807}
                if false {
                    fmt.Println(v5)
                    v16 := GenS1{F0: 14, F1: 0, F2: (v5 - (^(v5)))}
                    v5 += v4
                    v17 := GenS2{F0: true, F1: fi14}
                    _ = v16
                    _ = v17
                } else {
                    fmt.Println(fi14, v8)
                    var v18 uint8 = 193
                    _ = v18
                }
                _ = fi14
                _ = v15
            }
        }
        var v19 uint64 = 720
        if true {
            if ((-975 <= -1) && v2) {
                var v20 uint8 = 128
                _ = v20
            } else {
                if ((("\n\t" + "foo bar") <= ("hello" + "")) && v2) {
                    v3 -= (v3 / ((v3) | 1))
                    fmt.Println(v2, fi9, v2)
                    fmt.Println(v4, v2)
                }
                if ((!v2) || (!((!(!("a" != ""))) && v2))) {
                    var v21 string = "hello"
                    v22 := []int{fi9, fi9, (fi9 + (fi9 - (fi9 << ((uint(v3)) & 63))))}
                    _ = v21
                    _ = v22
                }
                v1["a"] = v8
            }
            v3 = (v3 * v3)
            var v23 uint8 = 2
            _ = v23
        } else {
            switch 127 {
            case 42:
                for fi24 := 0; fi24 < 10; fi24++ {
                    delete(v1, ("a" + (("a" + (("" + ("foo bar" + ("a" + (("" + "foo bar") + ("hello" + ""))))) + ("hello" + "\n\t"))) + ("\n\t" + "foo bar"))))
                    _ = fi24
                }
            default:
                fmt.Println(v2, v8, v4)
            }
            fmt.Println(v2)
            v4 = -81
        }
        _ = fi9
        _ = v19
    }
    v25 := GenS3{F0: 32767, F1: -719}
    var v26 uint16 = v3
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 uint = 255
        var v29 int16 = 42
        fmt.Println(v4, v8, v4)
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int = -9223372036854775808
    for fi31 := 0; fi31 < 14; fi31++ {
        fmt.Println(v2, v5, v30)
        v32 := []int{(v30 ^ fi31)}
        _ = fi31
        _ = v32
    }
    switch (("\n\t" + ("" + "foo bar")) + "\n\t") {
    case "a":
        v30 = v30
    case "foo bar":
        var v33 int = v30
        _ = v33
    default:
        v34 := GenS4{F0: 42, F1: false}
        _ = v34
    }
    var v35 int32 = 0
    fmt.Println(v2)
    if v2 {
        v5 = (v4 + v5)
    } else {
        fmt.Println(v3)
    }
    var v36 float32 = -2.25
    if (0 <= -9223372036854775808) {
        delete(v1, "\n\t")
        if (v30 == v30) {
            for fi37 := 0; fi37 < 9; fi37++ {
                for fi38 := 0; fi38 < 8; fi38++ {
                    switch "\n\t" {
                    case "":
                        var v39 float64 = (1e10 - 3.14159)
                        _ = v39
                    case "a":
                        fmt.Println(v5, v26)
                    case "hello":
                        v36 -= v36
                    }
                    if (true && (!false)) {
                        v4 = 0
                    }
                    _ = fi38
                }
                for fi40 := 0; fi40 < 12; fi40++ {
                    switch "" {
                    case "":
                        if v2 {
                            var v41 int8 = (v4 ^ 100)
                            v8 = 1
                            if v2 {
                                if ((v41 <= v5) && (0 <= 42)) {
                                    if v2 {
                                        v42 := GenS5{F0: 1}
                                        _ = v42
                                    } else {
                                        v1[(("foo bar" + "\n\t") + (((((("\n\t" + ((((("" + "") + "\n\t") + "foo bar") + ("a" + "\n\t")) + "foo bar")) + "foo bar") + "") + "foo bar") + (("hello" + "\n\t") + ("hello" + "\n\t"))) + "\n\t"))] = v8
                                        var v43 string = ("foo bar" + ("" + "\n\t"))
                                        var v44 uint8 = 42
                                        _ = v43
                                        _ = v44
                                    }
                                } else {
                                    var v45 string = "a"
                                    fmt.Println(v2, v3, v4)
                                    _ = v45
                                }
                                var v46 uint16 = 32767
                                _ = v46
                            } else {
                                var v47 int8 = int8(v8)
                                _ = v47
                            }
                            v48 := []int{}
                            _ = v41
                            _ = v48
                        } else {
                            v35 += (v35 >> ((4294967295) & 63))
                            switch 32767 {
                            case 9223372036854775807:
                                v49 := GenS5{F0: v35}
                                _ = v49
                            case -9223372036854775808:
                                v4 -= -128
                            default:
                                v1["\n\t"] = v8
                                v50 := []int{v30, fi40}
                                var v51 bool = false
                                var v52 int = (v30 | 32767)
                                fmt.Println(v35, fi40, v52)
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                            v53 := GenS6{F0: 564, F1: "hello", F2: "a"}
                            _ = v53
                        }
                    case "foo bar":
                        v54 := []int{-898, (v30 ^ (fi40 | -1)), (fi40 * fi40)}
                        _ = v54
                    case "a":
                        v4 += v4
                    default:
                        if false {
                            v4 = (v5 << ((1) & 63))
                            v35 -= 365
                            v3 -= v3
                            v55 := GenS1{F0: 80, F1: 128, F2: (-(v4))}
                            _ = v55
                        }
                    }
                    _ = len(v1)
                    _ = fi40
                }
                _ = fi37
            }
            v56 := GenS3{F0: 0, F1: 9223372036854775807}
            v57 := GenS5{F0: v35}
            _ = v56
            _ = v57
        } else {
            for fi58 := 0; fi58 < 4; fi58++ {
                v26 -= 896
                var v59 int16 = 58
                switch ("\n\t" + ("" + "")) {
                case "\n\t":
                    v60 := []int{fi58, 4294967295, 128, (fi58 | -2147483648)}
                    _ = v60
                }
                _ = fi58
                _ = v59
            }
            var v61 uint32 = v8
            v36 -= (v36 - (v36 + float32((float64(1.5) + -2.25))))
            _ = v61
        }
        if false {
            var v62 uint8 = 0
            _ = v62
        }
    }
    if v2 {
        for fi63 := 0; fi63 < 13; fi63++ {
            for fi64 := 0; fi64 < 18; fi64++ {
                var v65 uint16 = v26
                v4 += v4
                _ = fi64
                _ = v65
            }
            for fi66 := 0; fi66 < 14; fi66++ {
                fmt.Println(v26, v35, v26)
                _ = fi66
            }
            _ = fi63
        }
        v67 := GenS6{F0: -32768, F1: "hello", F2: ("a" + "hello")}
        _ = v67
    }
    fmt.Println(v5, v36)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v8
}
