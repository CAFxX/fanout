package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 string; F2 bool }
type GenS3 struct { F0 bool; F1 bool; F2 uint }
type GenS4 struct { F0 int; F1 int8; F2 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 int16; F2 uint16 }
func genFn0(p0 int16, p1 uint64, p2 string) int32 {
    p0 -= -193
    v4 := GenS1{F0: -9223372036854775808}
    var v5 int16 = (-(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := []int{128, 9223372036854775807, -9223372036854775808}
    v2 := make(map[string]string)
    v3 := []int{}
    var v6 uint16 = 790
    v2["\n\t"] = "hello"
    var v7 int16 = -1
    var v8 float32 = ((0.1 / 1.0) / 1.5)
    v9 := GenS1{F0: int64(v6)}
    var v10 int = -35
    if false {
        var v11 float32 = ((v8 + 3.14159) * v8)
        if false {
            v2["hello"] = ""
            var v12 bool = (1e10 != 1.5)
            v2["a"] = ((("hello" + "foo bar") + ("a" + "\n\t")) + ("foo bar" + "foo bar"))
            _ = v12
        }
        _ = v11
    } else {
        if (false && (!true)) {
            v2[""] = "hello"
        }
        var v13 int = (-(v10))
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint16 = (^(v6))
        for fi16 := 0; fi16 < 5; fi16++ {
            v10 += fi16
            var v17 uint = 383
            switch v10 {
            case -232:
                v18 := GenS1{F0: int64(v7)}
                _ = v18
            case 128:
                v19 := GenS1{F0: int64(v10)}
                _ = v19
            case 0:
                if true {
                    if ("a" != "foo bar") {
                        var v20 uint32 = 198
                        var v21 int8 = 127
                        switch 32767 {
                        case 128:
                            fmt.Println(v20)
                        case 32768:
                            v22 := GenS2{F0: 4294967295, F1: "", F2: true}
                            _ = v22
                        case 916:
                            v23 := GenS1{F0: 42}
                            _ = v23
                        default:
                            var v24 float32 = v8
                            v7 -= -1
                            _ = v24
                        }
                        _ = v20
                        _ = v21
                    } else {
                        var v25 uint8 = 2
                        v26 := GenS3{F0: (!false), F1: true, F2: v17}
                        var v27 int64 = 788
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    var v28 uint8 = 6
                    _ = v28
                } else {
                    v8 = v8
                    switch "foo bar" {
                    case "foo bar":
                        if false {
                            v8 -= float32(((3.14159 / 0.1) + 3.14159))
                            v17 += (v17 >> ((279) & 63))
                        } else {
                            v7 -= v7
                            v29 := GenS4{F0: 2147483647, F1: 127, F2: -721}
                            var v30 uint8 = 174
                            _ = v29
                            _ = v30
                        }
                    case "\n\t":
                        switch (("a" + "") + "hello") {
                        case "foo bar":
                            var v31 int = 85
                            _ = v31
                        case "\n\t":
                            v3 = append(v3, (^(fi14)))
                            var v32 string = "hello"
                            _ = v32
                        }
                    case "a":
                        v2[""] = ("a" + "foo bar")
                    }
                    fmt.Println(v10)
                }
            default:
                fmt.Println(v7)
            }
            _ = fi16
            _ = v17
        }
        var v33 bool = (0.1 != (1e10 / 1.5))
        _ = fi14
        _ = v15
        _ = v33
    }
    if (((!true) && true) || (!false)) {
        for fi34 := 0; fi34 < 7; fi34++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v35 := GenS5{F0: 671}
            _ = fi34
            _ = v35
        }
    } else {
        if true {
            var v36 int16 = (v7 - v7)
            _ = v36
        } else {
            v37 := GenS6{F0: -872, F1: (v7 - (v7 + v7)), F2: v6}
            _ = v37
        }
        if true {
            var v38 int = 32767
            _ = len(v2)
            _ = v38
        } else {
            v7 += (v7 + v7)
            if true {
                for fi39 := 0; fi39 < 11; fi39++ {
                    var v40 uint = 1
                    v41 := v2[("foo bar" + "\n\t")]
                    switch (("foo bar" + "hello") + ("\n\t" + "hello")) {
                    case "":
                        v7 = 32767
                    case "a":
                        switch "hello" {
                        case "\n\t":
                            for fi42 := 0; fi42 < 16; fi42++ {
                                if false {
                                    var v43 int64 = 144
                                    v7 = (v7 ^ int16(v6))
                                    var v44 uint = v40
                                    _ = v43
                                    _ = v44
                                }
                                v8 = v8
                                var v45 int16 = v7
                                _ = fi42
                                _ = v45
                            }
                        case "hello":
                            if false {
                                fmt.Println(v6, v6)
                                var v46 uint = 2
                                var v47 uint16 = v6
                                _ = v46
                                _ = v47
                            } else {
                                var v48 bool = true
                                _ = v48
                            }
                        }
                    default:
                        var v49 uint16 = 884
                        _ = v49
                    }
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                if (false || true) {
                    v8 = v8
                }
                if (false || (!(v7 != (v7 << ((65535) & 63))))) {
                    v50 := GenS1{F0: -134}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v8 = 1.5
                    _ = v50
                }
            }
            var v51 int16 = 1000
            _ = v51
        }
    }
    v6 = (v6 & (v6 / ((v6) | 1)))
    var v52 uint32 = 42
    var v53 int32 = -877
    var v54 uint = 65535
    if (944 >= (v52 + (v52 / ((v52) | 1)))) {
        var v55 int32 = int32(v7)
        v2["hello"] = ("foo bar" + ("a" + ("" + "a")))
        v6 = (v6 ^ (v6 | (v6 & (v6 << ((uint(v6)) & 63)))))
        _ = v55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v52
    _ = v53
    _ = v54
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
