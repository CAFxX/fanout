package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 bool }
func genFn0(p0 bool, p1 int32, p2 uint) string {
    p1 += p1
    v5 := GenS1{F0: -1, F1: -1}
    var v6 uint32 = 1
    var v7 uint32 = (v6 % ((v6) | 1))
    if ((^(p2)) == uint(v6)) {
        v8 := GenS1{F0: 1, F1: 254}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return ("hello" + ((("a" + "") + (("" + "a") + "")) + ""))
}

func main() {
    var v1 int64 = 130
    var v2 int16 = 42
    var v3 int8 = -1
    var v4 int = 32768
    var v9 uint = 65535
    v4 = int(v9)
    v10 := []int{v4, (v4 ^ v4), (v4 - int(v9))}
    _ = len(v10)
    v9 = 558
    if len(v10) > 0 {
        v10[0] = v4
    }
    _ = len(v10)
    if (-65 == (^(v3))) {
        if (0 > 727) {
            if true {
                var v11 float64 = 0.0
                v11 = v11
                switch ("" + "hello") {
                case "\n\t":
                    if ((v2 << (((v9 + (v9 % ((580) | 1)))) & 63)) > int16(v4)) {
                        v12 := GenS1{F0: v1, F1: (v1 ^ v1)}
                        _ = v12
                    }
                default:
                    if (v3 >= 103) {
                        for fi13 := 0; fi13 < 14; fi13++ {
                            var v14 int = (v4 | (^(fi13)))
                            if true {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                switch ("\n\t" + "a") {
                                case "foo bar":
                                    v9 += (v9 | v9)
                                case "\n\t":
                                    fmt.Println(fi13)
                                }
                                switch 4294967295 {
                                case 32768:
                                    v11 = v11
                                default:
                                    for fi15 := 0; fi15 < 14; fi15++ {
                                        if (true || true) {
                                            v16 := GenS2{F0: ("a" + (("\n\t" + ("" + (("foo bar" + ("foo bar" + "hello")) + ("a" + "foo bar")))) + "hello"))}
                                            _ = v16
                                        } else {
                                            var v17 uint64 = 42
                                            _ = v17
                                        }
                                        switch "hello" {
                                        case "foo bar":
                                            v1 += (v1 + v1)
                                        }
                                        if (0 >= (v1 | v1)) {
                                            if len(v10) > 0 {
                                                v10[0] = -224
                                            }
                                            v2 = v2
                                        }
                                        _ = fi15
                                    }
                                }
                            } else {
                                var v18 int32 = 42
                                var v19 float64 = 1.5
                                _ = v18
                                _ = v19
                            }
                            _ = fi13
                            _ = v14
                        }
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    } else {
                        v20 := GenS1{F0: v1, F1: -916}
                        var v21 string = ("" + "\n\t")
                        _ = v20
                        _ = v21
                    }
                }
                _ = v11
            } else {
                v10 = append(v10, (v4 - (v4 >> (((v9 ^ (v9 << ((v9) & 63)))) & 63))))
                if true {
                    for fi22 := 0; fi22 < 5; fi22++ {
                        switch v4 {
                        case -9223372036854775808:
                            v4 -= (v4 ^ (v4 / ((v4) | 1)))
                        case 1:
                            var v23 int = (fi22 << ((164) & 63))
                            _ = v23
                        case -492:
                            fmt.Println(v2, v3)
                        }
                        var v24 uint = 0
                        fmt.Println(v24, v2, v24)
                        _ = fi22
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 12; fi25++ {
                        switch -354 {
                        case -1:
                            v9 -= v9
                        case -2147483648:
                            v26 := GenS3{F0: "\n\t", F1: 1, F2: 879}
                            _ = v26
                        default:
                            if ((!(!(!true))) && (v3 >= v3)) {
                                v27 := GenS1{F0: -1, F1: v1}
                                for fi28 := 0; fi28 < 12; fi28++ {
                                    _ = len(v10)
                                    var v29 int32 = -506
                                    if false {
                                        if (!false) {
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                            v30 := GenS3{F0: "\n\t", F1: 211, F2: 65535}
                                            v31 := GenS4{F0: v1}
                                            _ = v30
                                            _ = v31
                                        }
                                        v10 = append(v10, 72)
                                    } else {
                                        if len(v10) > 0 {
                                            v10[0] = v4
                                        }
                                        var v32 string = "foo bar"
                                        if (v3 != -1) {
                                            v33 := GenS2{F0: "hello"}
                                            if (!(false || false)) {
                                                v34 := GenS5{F0: (v29 & (v29 >> ((v9) & 63)))}
                                                fmt.Println(v4)
                                                v32 = "\n\t"
                                                _ = v34
                                            } else {
                                                fmt.Println(v9, v3)
                                                var v35 uint = 4294967295
                                                fmt.Println(v2, v2)
                                                _ = v35
                                            }
                                            _ = v33
                                        }
                                        var v36 int64 = (v1 + v1)
                                        _ = v32
                                        _ = v36
                                    }
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    _ = fi28
                                    _ = v29
                                }
                                _ = v27
                            }
                        }
                        switch v4 {
                        case 2147483647:
                            v4 = v4
                        }
                        var v37 int64 = v1
                        _ = fi25
                        _ = v37
                    }
                }
                var v38 uint64 = 18446744073709551615
                _ = v38
            }
            switch "hello" {
            case "foo bar":
                for fi39 := 0; fi39 < 18; fi39++ {
                    if false {
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                var v40 int16 = (v2 | 42)
                                _ = v40
                            default:
                                v3 += 100
                            }
                        }
                        v1 = (v1 * (v1 * v1))
                        var v41 float64 = (float64(1.5) + -2.25)
                        _ = v41
                    }
                    if len(v10) > 0 {
                        v10[0] = v4
                    }
                    _ = fi39
                }
            default:
                var v42 string = ""
                _ = v42
            }
        }
        v1 = v1
        if (255 <= 956) {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        }
    }
    if false {
        if (-128 <= 127) {
            for fi43 := 0; fi43 < 4; fi43++ {
                for fi44 := 0; fi44 < 1; fi44++ {
                    v4 = fi44
                    var v45 int8 = v3
                    _ = fi44
                    _ = v45
                }
                _ = fi43
            }
            if (v3 <= v3) {
                if len(v10) > 0 {
                    v10[0] = (v4 % ((v4) | 1))
                }
                v3 = (v3 | (v3 + v3))
            } else {
                v46 := GenS4{F0: -1}
                _ = v46
            }
        }
    }
    var v47 int = (v4 << ((v9) & 63))
    v48 := GenS6{F0: v1, F1: true}
    if len(v10) > 0 {
        v10[0] = (-(v4))
    }
    switch 0 {
    case -9223372036854775808:
        for fi49 := 0; fi49 < 18; fi49++ {
            var v50 uint = (^(v9))
            fmt.Println(v3)
            _ = fi49
            _ = v50
        }
    case 482:
        v9 -= (v9 * v9)
    }
    v1 = v1
    v9 += 1
    var v51 int32 = 686
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v51:", v51)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v51
    _ = v9
}
