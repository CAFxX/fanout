package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int16; F1 int64 }
func genFn0(p0 uint32) uint16 {
    v7 := GenS1{F0: -1}
    if (0.1 <= (float32(0.1) / float32(1e10))) {
        var v8 int8 = 0
        _ = v8
    } else {
        if false {
            var v9 float64 = 0.1
            if false {
                for fi10 := 0; fi10 < 6; fi10++ {
                    switch (-(fi10)) {
                    case -2147483648:
                        v11 := GenS2{F0: 65535, F1: 0, F2: (p0 ^ p0)}
                        _ = v11
                    case 42:
                        fmt.Println(fi10, v9)
                    case -1:
                        v12 := GenS1{F0: -249}
                        _ = v12
                    default:
                        var v13 bool = (0.0 < (1.5 / v9))
                        _ = v13
                    }
                    _ = fi10
                }
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                v9 = v9
                _ = fi14
            }
            _ = v9
        } else {
            if false {
                switch -2147483648 {
                case 32767:
                    switch "hello" {
                    case "":
                        var v15 uint8 = 42
                        _ = v15
                    case "hello":
                        p0 = p0
                    }
                default:
                    for fi16 := 0; fi16 < 17; fi16++ {
                        p0 = p0
                        var v17 uint16 = 2
                        var v18 bool = ("a" != ("hello" + ""))
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
                if (-9223372036854775808 != 1000) {
                    p0 = 78
                    var v19 uint = 50
                    _ = v19
                }
                v20 := []int{1, -733, 32767}
                _ = v20
            }
        }
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        p0 += p0
        if ((false && true) || (("hello" != (("\n\t" + (("a" + "hello") + "hello")) + "")) && ((true || true) && (-2.25 >= 3.14159)))) {
            p0 = (p0 << ((2) & 63))
            v22 := []int{}
            v23 := GenS3{F0: false}
            _ = v22
            _ = v23
        }
        var v24 bool = (!true)
        _ = fi21
        _ = v24
    }
    v25 := GenS3{F0: true}
    var v26 int = 711
    _ = p0
    _ = v25
    _ = v26
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = -2.25
    var v3 uint16 = 255
    var v4 float32 = float32(-2.25)
    v5 := []uint32{2, 0, 65535, 4294967295}
    var v6 int = 551
    if true {
        if false {
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 uint32 = 2147483648
                fmt.Println(v2, v4, v3)
                v4 = v2
                _ = fi27
                _ = v28
            }
        } else {
            var v29 int32 = 2147483647
            var v30 float64 = (0.1 / 1e10)
            _ = v29
            _ = v30
        }
        fmt.Println(v4, v3)
        switch (("\n\t" + "") + (("\n\t" + "a") + "foo bar")) {
        case "hello":
            var v31 int32 = 2147483647
            _ = v31
        default:
            if true {
                if len(v5) > 0 {
                    v5[0] = 426
                }
                for fi32 := 0; fi32 < 12; fi32++ {
                    if false {
                        v33 := GenS4{F0: "hello", F1: 128}
                        if (!(false || false)) {
                            v4 -= float32(0.1)
                        }
                        _ = v33
                    }
                    if false {
                        if true {
                            for fi34 := 0; fi34 < 10; fi34++ {
                                v2 = -2.25
                                v3 += 1
                                var v35 uint = 18446744073709551615
                                _ = fi34
                                _ = v35
                            }
                            if true {
                                v4 = v4
                            }
                        }
                    } else {
                        var v36 float64 = (1e10 - ((0.0 + 0.0) - 1e10))
                        switch "a" {
                        case "foo bar":
                            switch ("a" + "\n\t") {
                            case "foo bar":
                                v37 := GenS1{F0: -392}
                                _ = v37
                            case "a":
                                v38 := GenS4{F0: ("" + "a"), F1: 0}
                                _ = v38
                            case "hello":
                                fmt.Println(v3, v3)
                            }
                        case "a":
                            for fi39 := 0; fi39 < 18; fi39++ {
                                _ = len(v1)
                                var v40 int16 = 905
                                v41 := GenS3{F0: (-911 != 398)}
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                        }
                        _ = v36
                    }
                    var v42 int64 = 42
                    _ = fi32
                    _ = v42
                }
                for fi43 := 0; fi43 < 10; fi43++ {
                    var v44 bool = ((^(v6)) >= fi43)
                    if v44 {
                        v45 := GenS5{F0: 1000, F1: 42}
                        if v44 {
                            v46 := GenS2{F0: v3, F1: 0, F2: 0}
                            v47 := GenS2{F0: v3, F1: 1, F2: 529}
                            var v48 bool = (32767 != int16(fi43))
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                        _ = v45
                    } else {
                        var v49 uint = 4294967295
                        _ = v49
                    }
                    var v50 bool = (-32768 == 1)
                    _ = fi43
                    _ = v44
                    _ = v50
                }
            }
        }
    } else {
        v4 += v2
    }
    if false {
        for fi51 := 0; fi51 < 12; fi51++ {
            v2 -= (v2 + v4)
            _ = fi51
        }
    } else {
        var v52 uint64 = 18446744073709551615
        _ = v52
    }
    fmt.Println(v3)
    fmt.Println(v2, v6, v4)
    switch "hello" {
    case "\n\t":
        v1[("foo bar" + "")] = 4294967295
    case "":
        var v53 uint32 = 543
        _ = v53
    }
    v3 = (^(v3))
    for fi54 := 0; fi54 < 20; fi54++ {
        v55 := GenS2{F0: (v3 & (v3 / ((v3) | 1))), F1: 1, F2: 42}
        if (((4294967295 <= 642) && false) && true) {
            var v56 uint32 = 1
            v57 := GenS2{F0: v3, F1: 65535, F2: (v56 >> ((uint(v56)) & 63))}
            fmt.Println(fi54)
            _ = v56
            _ = v57
        } else {
            v3 -= (v3 << ((255) & 63))
        }
        _ = fi54
        _ = v55
    }
    v6 += v6
    v1[""] = v6
    v58 := GenS3{F0: (v4 != 1.5)}
    if false {
        v59 := GenS4{F0: "hello", F1: 62}
        v3 = (v3 * genFn0(297))
        v4 += ((v4 - v4) * float32(0.0))
        _ = v59
    }
    v4 += float32(0.1)
    v5 = append(v5, 4294967295)
    if false {
        var v60 string = (("" + "a") + ("a" + "hello"))
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
}
