package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 string; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 int; F2 int64 }
func genFn0(p0 int32, p1 string, p2 int32) uint16 {
    v7 := GenS1{F0: (42 > 575)}
    p2 = 114
    v8 := []int{192, 127}
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint16 = 255
        if len(v8) > 0 {
            v8[0] = fi9
        }
        if true {
            fmt.Println(p2, p0)
        } else {
            var v11 int64 = 9223372036854775807
            v11 = 9223372036854775807
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]string)
    var v3 bool = true
    v4 := []string{"foo bar", "foo bar", ("" + ("" + "foo bar"))}
    v5 := make(map[int]int)
    var v6 int64 = 42
    for fi12 := 0; fi12 < 16; fi12++ {
        switch -1 {
        case 4294967295:
            v13 := GenS2{F0: (-275 < -956), F1: ("" + ("foo bar" + "\n\t")), F2: (98 == 238)}
            _ = v13
        default:
            for fi14 := 0; fi14 < 5; fi14++ {
                if true {
                    _ = len(v2)
                    if ((fi14 != fi12) || v3) {
                        if v3 {
                            if len(v4) > 0 {
                                v4[0] = "hello"
                            }
                            switch fi14 {
                            case 2147483647:
                                var v15 string = "hello"
                                _ = v15
                            default:
                                v16 := GenS3{F0: 0}
                                _ = v16
                            }
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                        }
                    } else {
                        v3 = v3
                        v1 = (v1 << ((187) & 63))
                    }
                    v6 = (v6 << ((321) & 63))
                    v4 = append(v4, "foo bar")
                }
                var v17 int16 = 1
                _ = fi14
                _ = v17
            }
        }
        v18 := GenS1{F0: v3}
        v6 -= v6
        _ = fi12
        _ = v18
    }
    v4 = append(v4, "foo bar")
    for fi19 := 0; fi19 < 2; fi19++ {
        v1 -= 4294967295
        v20 := GenS1{F0: false}
        _ = fi19
        _ = v20
    }
    v6 -= v6
    var v21 int8 = 0
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 uint64 = 58
        _ = fi22
        _ = v23
    }
    var v24 bool = v3
    var v25 bool = (959 != 0)
    var v26 bool = (!false)
    if v24 {
        var v27 string = ("foo bar" + "\n\t")
        if ((v6 + 807) == v6) {
            var v28 int16 = 133
            _ = v28
        } else {
            var v29 float64 = (0.1 - (-2.25 - 1.5))
            switch v27 {
            case "\n\t":
                if (2 <= 619) {
                    v1 = v1
                } else {
                    fmt.Println(v6)
                    v24 = v24
                }
            case "a":
                var v30 uint64 = v1
                _ = v30
            case "hello":
                for fi31 := 0; fi31 < 17; fi31++ {
                    v32 := GenS1{F0: v25}
                    if (v27 <= "a") {
                        var v33 int64 = -9223372036854775808
                        _ = v33
                    }
                    if (42 >= 2147483647) {
                        for fi34 := 0; fi34 < 5; fi34++ {
                            v27 = v27
                            v5[(fi34 ^ (fi31 | fi34))] = (^(fi34))
                            v4 = append(v4, ("hello" + "hello"))
                            _ = fi34
                        }
                    }
                    _ = fi31
                    _ = v32
                }
            }
            var v35 int16 = -1
            _ = v29
            _ = v35
        }
        _ = v27
    } else {
        switch "a" {
        case "a":
            fmt.Println(v6)
        default:
            var v36 uint = 255
            _ = v36
        }
        for fi37 := 0; fi37 < 2; fi37++ {
            switch "a" {
            case "foo bar":
                if (v21 >= v21) {
                    var v38 int = -2147483648
                    _ = v38
                }
            }
            var v39 int32 = 1
            _ = fi37
            _ = v39
        }
        v40 := GenS4{F0: 42}
        _ = v40
    }
    v41 := GenS5{F0: "foo bar", F1: 32767, F2: v6}
    if (!v3) {
        var v42 int = 0
        delete(v5, (v42 ^ -2147483648))
        _ = v42
    } else {
        v1 = (v1 % ((v1) | 1))
        v43 := GenS2{F0: v25, F1: "foo bar", F2: (((v6 * v6) != v6) && (!(55 <= -1)))}
        for fi44 := 0; fi44 < 5; fi44++ {
            if v25 {
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                if (2147483647 > 0) {
                    if false {
                        if false {
                            if true {
                                var v45 uint8 = 255
                                _ = v45
                            }
                        }
                    }
                    for fi46 := 0; fi46 < 12; fi46++ {
                        v47 := GenS2{F0: v3, F1: "", F2: (!false)}
                        v48 := GenS4{F0: genFn0(-1, (("hello" + "") + "\n\t"), -1)}
                        delete(v2, 4294967295)
                        v21 -= v21
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                }
                if len(v4) > 0 {
                    v4[0] = "a"
                }
            } else {
                v49 := GenS5{F0: "hello", F1: (fi44 & fi44), F2: v6}
                _ = v49
            }
            for fi50 := 0; fi50 < 7; fi50++ {
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                var v51 uint64 = (v1 * 2)
                _ = len(v4)
                _ = len(v2)
                _ = fi50
                _ = v51
            }
            v5[int(v1)] = fi44
            _ = fi44
        }
        _ = v43
    }
    v52 := GenS2{F0: true, F1: "", F2: ((((!(v6 >= v6)) && v3) && ((0.1 / (3.14159 - 0.0)) != (float64(-2.25) + -2.25))) && (-1 > 127))}
    v4 = append(v4, ("" + ""))
    v53 := GenS4{F0: 2}
    for fi54 := 0; fi54 < 14; fi54++ {
        v55 := GenS2{F0: (float64((3.14159 + 1e10)) <= (float64(1.5) + float64((-2.25 + 3.14159)))), F1: "foo bar", F2: v25}
        _ = fi54
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
