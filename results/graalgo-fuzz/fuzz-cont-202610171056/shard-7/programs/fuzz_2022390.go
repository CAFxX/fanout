package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int8; F2 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 int; F1 int; F2 bool }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 uint32; F1 string; F2 uint }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64 }
func genFn0(p0 int64, p1 uint, p2 int64) string {
    fmt.Println(p1)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "\n\t"
}
func genFn1() int16 {
    var v8 uint64 = 0
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    fmt.Println(v8)
    _ = v8
    return -32768
}

func main() {
    var v1 uint8 = 234
    var v2 float32 = 3.14159
    var v3 uint8 = 127
    var v4 uint64 = 2
    v5 := []string{}
    v6 := make(map[int]int)
    if false {
        var v11 int = -1
        _ = v11
    } else {
        var v12 bool = (!false)
        _ = len(v6)
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
        _ = v12
    }
    v13 := GenS1{F0: -110, F1: 2}
    if true {
        var v14 uint64 = 910
        _ = v14
    }
    if (!(!(true && true))) {
        if (false || ((!((v1 << ((148) & 63)) <= (v3 % ((v3) | 1)))) || true)) {
            var v15 uint8 = 1
            _ = v15
        } else {
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := GenS2{F0: 0}
                v18 := GenS3{F0: 0, F1: -91, F2: 812}
                if true {
                    v19 := GenS4{F0: -1}
                    fmt.Println(v3, v4)
                    if false {
                        switch ("\n\t" + ("" + "hello")) {
                        case "a":
                            var v20 uint8 = (v1 & v1)
                            _ = v20
                        case "foo bar":
                            fmt.Println(v1, fi16, v4)
                        default:
                            v21 := GenS5{F0: 1, F1: 42}
                            _ = v21
                        }
                        v22 := GenS4{F0: 42}
                        v23 := GenS6{F0: (-(fi16)), F1: (fi16 / ((2147483647) | 1)), F2: (0 < 2)}
                        _ = v22
                        _ = v23
                    } else {
                        v4 -= (v4 >> ((0) & 63))
                        if (uint32(v4) <= 1) {
                            v4 = (v4 ^ (v4 % ((v4) | 1)))
                            v4 = (v4 | v4)
                            v3 -= 0
                        } else {
                            v24 := v6[(fi16 + 1)]
                            switch v24 {
                            case -2147483648:
                                v2 -= ((v2 * 0.1) + v2)
                            case -132:
                                v25 := GenS7{F0: 255, F1: -1}
                                _ = v25
                            }
                            _ = v24
                        }
                    }
                    _ = v19
                }
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        for fi26 := 0; fi26 < 16; fi26++ {
            for fi27 := 0; fi27 < 7; fi27++ {
                v4 -= v4
                _ = fi27
            }
            for fi28 := 0; fi28 < 8; fi28++ {
                var v29 uint64 = 18446744073709551615
                if true {
                    fmt.Println(v4)
                    var v30 uint = 18446744073709551615
                    switch "\n\t" {
                    case "":
                        if ((true && ((((v30 > (v30 - 2)) || (!false)) && ((v4 | 262) > (v4 | v4))) || ((-1 >= -32768) || (42 != 0)))) || true) {
                            var v31 int16 = 0
                            var v32 string = ""
                            _ = v31
                            _ = v32
                        } else {
                            fmt.Println(v2, v3)
                            v33 := GenS8{F0: 1, F1: ("foo bar" + "a"), F2: (v30 / ((463) | 1))}
                            _ = v33
                        }
                    case "a":
                        v34 := GenS2{F0: -1}
                        _ = v34
                    }
                    _ = v30
                }
                var v35 int64 = 0
                _ = fi28
                _ = v29
                _ = v35
            }
            fmt.Println(fi26)
            _ = fi26
        }
        v4 -= v4
    } else {
        v2 += 1.5
        if len(v5) > 0 {
            v5[0] = ("foo bar" + "a")
        }
    }
    if true {
        var v36 int32 = -2147483648
        if false {
            v37 := GenS5{F0: 379, F1: 42}
            _ = v37
        }
        _ = v36
    } else {
        var v38 float32 = v2
        v4 = v4
        for fi39 := 0; fi39 < 4; fi39++ {
            v40 := GenS8{F0: 4294967295, F1: "foo bar", F2: 98}
            v6[1] = 403
            _ = fi39
            _ = v40
        }
        _ = v38
    }
    fmt.Println(v3)
    var v41 float32 = -2.25
    for fi42 := 0; fi42 < 13; fi42++ {
        switch (("foo bar" + "foo bar") + "a") {
        case "":
            var v43 uint16 = 2
            _ = v43
        case "foo bar":
            v44 := v6[127]
            _ = v44
        }
        v45 := GenS8{F0: 621, F1: "\n\t", F2: 35}
        var v46 int = fi42
        _ = fi42
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 2; fi47++ {
        var v48 int64 = 9223372036854775807
        v49 := GenS9{F0: ""}
        _ = fi47
        _ = v48
        _ = v49
    }
    v1 = v3
    v5 = append(v5, "\n\t")
    var v50 int8 = 42
    var v51 int = -388
    if len(v5) > 0 {
        v5[0] = ("" + "a")
    }
    v52 := GenS10{F0: 464}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
