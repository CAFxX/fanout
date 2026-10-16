package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint8; F1 int; F2 uint32 }
type GenS6 struct { F0 string; F1 int64; F2 uint16 }
type GenS7 struct { F0 uint; F1 uint16; F2 int }
func genFn0(p0 string, p1 uint, p2 uint16) int {
    v5 := GenS1{F0: ("foo bar" + (p0 + (("hello" + ("a" + ("foo bar" + ("a" + "a")))) + p0))), F1: 42}
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 271
}

func main() {
    v1 := []string{}
    var v2 bool = (1 >= 9223372036854775807)
    v3 := []uint32{4294967295, 702}
    var v4 string = "a"
    var v6 float32 = 0.0
    v3 = append(v3, 671)
    switch "" {
    case "foo bar":
        v4 = v4
    default:
        var v7 float64 = 1.5
        _ = v7
    }
    switch 206 {
    case 128:
        v1 = append(v1, ("hello" + v4))
    case 2147483647:
        v2 = v2
    case 0:
        var v8 int = 127
        _ = v8
    }
    switch ((v4 + "foo bar") + "foo bar") {
    case "a":
        if v2 {
            if false {
                for fi9 := 0; fi9 < 9; fi9++ {
                    fmt.Println(v4)
                    _ = fi9
                }
                for fi10 := 0; fi10 < 16; fi10++ {
                    if (true || (v4 == v4)) {
                        v6 = 1.5
                    } else {
                        if false {
                            v11 := GenS2{F0: fi10}
                            _ = len(v1)
                            _ = v11
                        } else {
                            v4 = "hello"
                            v12 := GenS1{F0: v4, F1: 681}
                            v13 := GenS2{F0: genFn0(v4, 2, 947)}
                            _ = v12
                            _ = v13
                        }
                    }
                    for fi14 := 0; fi14 < 4; fi14++ {
                        if true {
                            var v15 uint64 = 248
                            fmt.Println(v4, fi10, v6)
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            _ = v15
                        } else {
                            if true {
                                var v16 uint16 = 32767
                                if v2 {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v16
                            }
                            if false {
                                switch (v4 + v4) {
                                case "a":
                                    v17 := GenS1{F0: ("" + "a"), F1: -211}
                                    _ = v17
                                default:
                                    if (!v2) {
                                        v18 := GenS3{F0: ((!(false && v2)) || (!(-824 < 1)))}
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v19 string = (v4 + "")
                                        v20 := GenS1{F0: "", F1: 9223372036854775807}
                                        v2 = ((v2 || v2) || v2)
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    } else {
                                        var v21 int64 = -1
                                        fmt.Println(fi14, fi14)
                                        v22 := GenS4{F0: 983, F1: 228, F2: 255}
                                        _ = v21
                                        _ = v22
                                    }
                                }
                                _ = len(v1)
                                var v23 uint64 = 1
                                _ = v23
                            }
                            fmt.Println(v6)
                        }
                        _ = fi14
                    }
                    v24 := GenS1{F0: v4, F1: 894}
                    _ = fi10
                    _ = v24
                }
            }
            fmt.Println(v6, v4, v4)
            var v25 bool = false
            _ = v25
        } else {
            if (v2 || (v2 && v2)) {
                v4 = v4
            } else {
                for fi26 := 0; fi26 < 19; fi26++ {
                    var v27 int64 = -1
                    v28 := GenS5{F0: 146, F1: (fi26 - fi26), F2: 493}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            switch ((v4 + v4) + (v4 + v4)) {
            case "hello":
                for fi29 := 0; fi29 < 17; fi29++ {
                    v30 := GenS2{F0: 2147483647}
                    for fi31 := 0; fi31 < 17; fi31++ {
                        var v32 int32 = -462
                        fmt.Println(v32, v4)
                        _ = fi31
                        _ = v32
                    }
                    var v33 int32 = 605
                    _ = fi29
                    _ = v30
                    _ = v33
                }
            default:
                var v34 int64 = 9223372036854775807
                _ = v34
            }
        }
    case "foo bar":
        v6 -= (v6 * v6)
    default:
        var v35 uint = 2
        _ = v35
    }
    fmt.Println(v2, v4, v6)
    var v36 uint8 = 255
    switch v4 {
    case "":
        if ((v36 & v36) != (v36 - v36)) {
            v3 = append(v3, 119)
        } else {
            var v37 uint64 = 306
            v38 := GenS1{F0: (v4 + "a"), F1: 263}
            _ = v37
            _ = v38
        }
    case "foo bar":
        _ = len(v3)
    case "\n\t":
        for fi39 := 0; fi39 < 10; fi39++ {
            v40 := GenS6{F0: v4, F1: 9223372036854775807, F2: 1}
            v6 -= (v6 / v6)
            var v41 int = genFn0(("\n\t" + v4), 65535, 255)
            _ = fi39
            _ = v40
            _ = v41
        }
    default:
        v42 := GenS7{F0: 255, F1: 1, F2: 1}
        _ = v42
    }
    v4 = (v4 + v4)
    var v43 float32 = float32(0.0)
    var v44 uint16 = 255
    if len(v1) > 0 {
        v1[0] = "a"
    }
    var v45 int64 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v6
}
