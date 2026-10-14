package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int8; F2 int16 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint64; F1 uint; F2 int }
type GenS11 struct { F0 bool }
func genFn0(p0 uint64, p1 int8) int32 {
    for fi5 := 0; fi5 < 6; fi5++ {
        p0 = (p0 >> ((0) & 63))
        v6 := []int{}
        _ = fi5
        _ = v6
    }
    fmt.Println(p1, p0)
    v7 := GenS1{F0: 9223372036854775807, F1: 0}
    var v8 int64 = 728
    v9 := []int{-248, 796, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 bool, p1 int16, p2 uint32) int32 {
    var v10 string = "hello"
    var v11 uint8 = 224
    switch 9223372036854775807 {
    case 127:
        v12 := []int{-106, 32767, 4294967295}
        _ = v12
    case 42:
        v13 := GenS1{F0: -1, F1: (v11 << ((734) & 63))}
        _ = v13
    case 1:
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := []int{}
            v11 = 42
            _ = fi14
            _ = v15
        }
    }
    if p0 {
        if p0 {
            v16 := GenS1{F0: 0, F1: v11}
            _ = v16
        }
        v10 = (((v10 + v10) + ("foo bar" + v10)) + ((v10 + "hello") + v10))
        var v17 string = ("\n\t" + v10)
        _ = v17
    } else {
        v18 := GenS2{F0: 869, F1: 17, F2: p1}
        if ((v10 + "hello") != v10) {
            if p0 {
                v11 -= v11
            }
            v19 := []int{-964, 2147483647, 4294967295}
            _ = v19
        }
        if (-1 == 1) {
            switch int(p2) {
            case 2147483647:
                v20 := []int{42, -9223372036854775808, -2147483648, 128}
                _ = v20
            case -9223372036854775808:
                v21 := []int{-2147483648}
                _ = v21
            }
            if (p0 && p0) {
                if p0 {
                    v22 := GenS2{F0: 2, F1: 80, F2: 949}
                    v23 := []int{9223372036854775807, 42, -2147483648, 128}
                    var v24 int64 = 42
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    fmt.Println(p2)
                    for fi25 := 0; fi25 < 11; fi25++ {
                        v26 := GenS3{F0: (fi25 + fi25), F1: genFn0(108, 106)}
                        _ = fi25
                        _ = v26
                    }
                }
                for fi27 := 0; fi27 < 8; fi27++ {
                    p1 = 32767
                    _ = fi27
                }
                v28 := []int{1, -2147483648, 32767}
                _ = v28
            } else {
                switch "hello" {
                case "":
                    for fi29 := 0; fi29 < 5; fi29++ {
                        if (65535 != 673) {
                            v30 := GenS4{F0: (v10 + v10)}
                            v31 := []int{fi29, fi29, (-(fi29)), -2147483648}
                            _ = v30
                            _ = v31
                        }
                        if (!p0) {
                            if p0 {
                                switch (((("a" + "\n\t") + v10) + v10) + (v10 + (("" + ("\n\t" + "")) + ((("hello" + v10) + "hello") + "hello")))) {
                                case "hello":
                                    var v32 int8 = -128
                                    _ = v32
                                }
                                if ((!false) && p0) {
                                    p1 += p1
                                    fmt.Println(p1)
                                    v10 = (("foo bar" + "\n\t") + ((("foo bar" + v10) + v10) + (v10 + v10)))
                                }
                            } else {
                                if ((^(p2)) <= p2) {
                                    if false {
                                        fmt.Println(fi29, v11, p0)
                                        if ((!true) || true) {
                                            p2 = (^(p2))
                                        }
                                    }
                                    if p0 {
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(p0, p1)
                                        case "foo bar":
                                            v33 := GenS5{F0: (int16(fi29) == p1)}
                                            _ = v33
                                        case "":
                                            fmt.Println(p2)
                                        default:
                                            var v34 bool = (1e10 < -2.25)
                                            _ = v34
                                        }
                                        var v35 int64 = 42
                                        p1 = (p1 + (^(p1)))
                                        _ = v35
                                    }
                                } else {
                                    if false {
                                        var v36 string = v10
                                        _ = v36
                                    } else {
                                        for fi37 := 0; fi37 < 7; fi37++ {
                                            var v38 int16 = (-(p1))
                                            var v39 int32 = 1
                                            _ = fi37
                                            _ = v38
                                            _ = v39
                                        }
                                    }
                                    v40 := GenS6{F0: 255}
                                    fmt.Println(p0, v11, p0)
                                    _ = v40
                                }
                                for fi41 := 0; fi41 < 19; fi41++ {
                                    v42 := []int{fi29}
                                    if (1 <= 42) {
                                        fmt.Println(p0, fi41)
                                    }
                                    v11 -= 0
                                    _ = fi41
                                    _ = v42
                                }
                            }
                            var v43 int64 = -1
                            _ = v43
                        } else {
                            v44 := GenS4{F0: (("foo bar" + v10) + v10)}
                            _ = v44
                        }
                        v45 := GenS7{F0: 440}
                        _ = fi29
                        _ = v45
                    }
                default:
                    v46 := GenS6{F0: 0}
                    _ = v46
                }
            }
            var v47 int32 = -720
            _ = v47
        }
        _ = v18
    }
    v48 := []int{4294967295, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v48
    return 0
}

func main() {
    v1 := []string{}
    var v2 int32 = 809
    v3 := []string{"a", "foo bar"}
    var v4 string = ("foo bar" + "")
    v4 = "hello"
    for fi49 := 0; fi49 < 17; fi49++ {
        var v50 int16 = 42
        var v51 uint16 = 32767
        switch v4 {
        case "foo bar":
            _ = len(v3)
        case "a":
            if true {
                v52 := GenS1{F0: 0, F1: 2}
                _ = v52
            }
        case "hello":
            var v53 int8 = 42
            _ = v53
        }
        _ = fi49
        _ = v50
        _ = v51
    }
    fmt.Println(v2)
    if false {
        for fi54 := 0; fi54 < 3; fi54++ {
            switch (("a" + v4) + v4) {
            case "hello":
                if (true && (true || true)) {
                    v55 := GenS3{F0: (-(fi54)), F1: v2}
                    for fi56 := 0; fi56 < 5; fi56++ {
                        if (!true) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v57 := GenS8{F0: 9223372036854775807}
                            _ = v57
                        } else {
                            if len(v1) > 0 {
                                v1[0] = ((("a" + v4) + (v4 + v4)) + (v4 + "\n\t"))
                            }
                            v58 := GenS7{F0: 128}
                            var v59 bool = (0.0 > (0.1 + ((1.5 * 3.14159) / 1.5)))
                            _ = v58
                            _ = v59
                        }
                        var v60 uint16 = 0
                        v61 := GenS6{F0: 42}
                        _ = fi56
                        _ = v60
                        _ = v61
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v55
                }
            }
            v62 := GenS5{F0: (false || (255 == 1))}
            _ = len(v1)
            _ = fi54
            _ = v62
        }
    } else {
        v63 := GenS9{F0: 1000}
        _ = v63
    }
    if len(v1) > 0 {
        v1[0] = ("hello" + "a")
    }
    for fi64 := 0; fi64 < 14; fi64++ {
        var v65 int = fi64
        for fi66 := 0; fi66 < 18; fi66++ {
            var v67 int16 = 32767
            switch ((("hello" + "hello") + v4) + v4) {
            case "a":
                v68 := GenS3{F0: (^(fi64)), F1: (v2 / ((v2) | 1))}
                _ = v68
            }
            var v69 float64 = 1e10
            _ = fi66
            _ = v67
            _ = v69
        }
        _ = fi64
        _ = v65
    }
    v70 := GenS10{F0: 661, F1: 42, F2: -892}
    var v71 float64 = 3.14159
    v4 = "foo bar"
    fmt.Println(v4, v71, v71)
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (false && (v4 <= "\n\t")) {
            v71 -= 3.14159
            var v72 float64 = v71
            _ = v72
        } else {
            v4 = v4
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        v73 := GenS11{F0: ((!false) || (1.5 <= 3.14159))}
        v71 += v71
        _ = v73
    }
    v71 += v71
    v74 := GenS6{F0: 847}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v70
    _ = v71
    _ = v74
}
