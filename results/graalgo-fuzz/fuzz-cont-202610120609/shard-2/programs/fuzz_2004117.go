package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int64 = 349
    v2 := []string{"a", "foo bar"}
    v3 := make(map[string]int)
    v4 := []int{128, 128, -977, 9223372036854775807}
    var v5 uint64 = 1
    v6 := GenS1{F0: 2, F1: 4294967295, F2: ""}
    if (false || true) {
        for fi7 := 0; fi7 < 17; fi7++ {
            delete(v3, "foo bar")
            _ = fi7
        }
        v8 := GenS1{F0: 769, F1: 1, F2: ("hello" + "")}
        _ = v8
    } else {
        _ = len(v3)
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int8 = 100
            v11 := GenS1{F0: 927, F1: 42, F2: ("\n\t" + ("a" + "\n\t"))}
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 uint16 = 591
                var v14 int64 = (v1 / ((v1) | 1))
                v1 = (-(v1))
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    v1 -= -1
    v1 = 350
    fmt.Println(v5)
    for fi15 := 0; fi15 < 13; fi15++ {
        v5 = 42
        switch "\n\t" {
        case "":
            if false {
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 uint32 = 65535
                    var v18 int16 = 42
                    var v19 float64 = 3.14159
                    var v20 int = (fi15 >> ((118) & 63))
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v18, v17, v18)
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                var v21 int16 = 0
                v1 = v1
                _ = v21
            } else {
                if false {
                    v3[""] = (^(fi15))
                    v5 += v5
                    v5 = v5
                }
                v1 -= (v1 - (-(v1)))
                v1 += (v1 % ((49) | 1))
            }
        case "a":
            if ("a" != "") {
                if true {
                    var v22 float32 = 1e10
                    var v23 int64 = 42
                    _ = v22
                    _ = v23
                } else {
                    var v24 bool = (("foo bar" + (("hello" + ((("" + ("hello" + (("" + ("" + "a")) + "a"))) + "a") + "foo bar")) + "")) <= "hello")
                    v5 = 42
                    fmt.Println(v1, fi15)
                    _ = v24
                }
                v1 = (v1 << ((65535) & 63))
            }
        }
        _ = fi15
    }
    switch "a" {
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = -516
        }
    case "foo bar":
        if true {
            if true {
                for fi25 := 0; fi25 < 13; fi25++ {
                    fmt.Println(fi25)
                    _ = fi25
                }
                if (!false) {
                    switch -70 {
                    case -1:
                        var v26 float64 = 1e10
                        _ = v26
                    case 9223372036854775807:
                        fmt.Println(v1, v1)
                    default:
                        _ = len(v2)
                    }
                    for fi27 := 0; fi27 < 10; fi27++ {
                        if false {
                            if (!false) {
                                if (!false) {
                                    var v28 uint = 255
                                    if (true && true) {
                                        switch 127 {
                                        case 128:
                                            var v29 int16 = 1000
                                            _ = v29
                                        case 733:
                                            v30 := GenS1{F0: 979, F1: (v28 & 2), F2: (("\n\t" + "a") + ("hello" + ("" + ("" + "a"))))}
                                            _ = v30
                                        }
                                    } else {
                                        v31 := GenS2{F0: 255}
                                        if false {
                                            switch "a" {
                                            case "\n\t":
                                                if (42 < 128) {
                                                    for fi32 := 0; fi32 < 15; fi32++ {
                                                        v33 := GenS3{F0: (v5 << ((752) & 63)), F1: v28, F2: v28}
                                                        var v34 uint = v28
                                                        _ = fi32
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = "a"
                                                    }
                                                } else {
                                                    v1 += v1
                                                    for fi35 := 0; fi35 < 4; fi35++ {
                                                        v5 = (v5 & v5)
                                                        v3[("hello" + ("" + "hello"))] = (fi35 >> (((v28 | (v28 ^ v28))) & 63))
                                                        _ = fi35
                                                    }
                                                    v36 := GenS4{F0: (11 == 100)}
                                                    _ = v36
                                                }
                                            case "hello":
                                                v3[("hello" + ("\n\t" + ""))] = 42
                                            }
                                        } else {
                                            var v37 bool = true
                                            _ = v37
                                        }
                                        v5 = v5
                                        _ = v31
                                    }
                                    switch fi27 {
                                    case 127:
                                        v38 := GenS5{F0: 149}
                                        _ = v38
                                    case 32768:
                                        var v39 int = -2147483648
                                        _ = v39
                                    }
                                    _ = v28
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                if false {
                                    if (true || (42 == -1)) {
                                        var v40 uint64 = (v5 ^ (^(v5)))
                                        _ = v40
                                    } else {
                                        var v41 uint64 = v5
                                        var v42 uint64 = (^(v5))
                                        _ = v41
                                        _ = v42
                                    }
                                    var v43 int32 = 337
                                    _ = v43
                                }
                                var v44 float64 = 1.5
                                _ = v44
                            }
                            var v45 float32 = float32(1.5)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v45
                        }
                        var v46 int16 = 42
                        _ = fi27
                        _ = v46
                    }
                } else {
                    _ = len(v2)
                    v1 += 9223372036854775807
                    fmt.Println(v1, v5)
                }
                fmt.Println(v5)
            } else {
                if true {
                    v1 = 1
                    var v47 float32 = 0.1
                    _ = v47
                }
            }
            var v48 uint8 = 42
            v3["a"] = 4294967295
            _ = v48
        } else {
            if true {
                v1 += (v1 & (v1 + (v1 - v1)))
                v49 := GenS3{F0: 18446744073709551615, F1: 600, F2: 200}
                _ = v49
            } else {
                for fi50 := 0; fi50 < 11; fi50++ {
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    _ = fi50
                }
                v1 -= 9223372036854775807
                switch 128 {
                case 32767:
                    if (false && ((!true) && (!true))) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        switch -2147483648 {
                        case 309:
                            v5 = v5
                        default:
                            var v51 uint = 18446744073709551615
                            _ = v51
                        }
                        for fi52 := 0; fi52 < 17; fi52++ {
                            _ = len(v4)
                            _ = fi52
                        }
                    }
                }
            }
        }
    case "hello":
        if (v5 > v5) {
            v3["foo bar"] = 0
            v5 += 4294967295
            if true {
                switch -9223372036854775808 {
                case -300:
                    v53 := GenS1{F0: 826, F1: 2, F2: (("" + "") + ("foo bar" + "\n\t"))}
                    _ = v53
                case 127:
                    var v54 int64 = v1
                    _ = v54
                default:
                    var v55 float64 = ((3.14159 / 1e10) + 0.0)
                    _ = v55
                }
                _ = len(v3)
                var v56 int = 127
                _ = v56
            }
        }
    }
    var v57 uint = 255
    _ = len(v2)
    v57 += (v57 & v57)
    v3["a"] = -9223372036854775808
    v58 := GenS4{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
