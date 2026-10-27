package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint16; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 string; F2 string }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 uint8; F1 int16; F2 uint }
type GenS7 struct { F0 uint; F1 uint8; F2 uint }

func main() {
    v1 := []string{}
    var v2 int8 = 81
    v3 := make(map[string]bool)
    v4 := []uint32{2147483648}
    v3[(("foo bar" + "") + "hello")] = false
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int32 = -2147483648
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 747}
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 bool = (1 != 0)
        fmt.Println(fi8, v2, fi8)
        switch -271 {
        case 611:
            v9 = (2 != 4294967295)
        case 32768:
            v2 = (v2 & (v2 ^ v2))
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(v2)
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS1{F0: 1}
            _ = fi11
            _ = v12
        }
        fmt.Println(fi10)
        _ = fi10
    }
    v2 += -1
    v2 -= (-(v2))
    v13 := GenS2{F0: 18446744073709551615, F1: 2, F2: (("foo bar" + (("a" + "a") + "a")) + "foo bar")}
    var v14 uint64 = 938
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS2{F0: 379, F1: 42, F2: "a"}
        v14 -= v14
        if (-1 == -9223372036854775808) {
            if (!(false && true)) {
                v14 = v14
            }
            v17 := v3["hello"]
            fmt.Println(v2)
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    var v18 uint8 = 214
    switch -662 {
    case 0:
        if false {
            fmt.Println(v18)
            var v19 uint64 = (v14 ^ 1)
            _ = v19
        }
    case 42:
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 int64 = 42
            if false {
                v22 := GenS3{F0: 0}
                if true {
                    for fi23 := 0; fi23 < 14; fi23++ {
                        var v24 uint64 = v14
                        if (32 < (v2 << ((uint(v24)) & 63))) {
                            fmt.Println(fi23, v14, v21)
                        } else {
                            v25 := GenS4{F0: (v21 >> ((2) & 63)), F1: (("a" + "hello") + ("" + ("" + ""))), F2: ("foo bar" + "\n\t")}
                            _ = v25
                        }
                        if ((("" + "") + "foo bar") <= ("hello" + "\n\t")) {
                            _ = len(v3)
                            if false {
                                v26 := GenS5{F0: ((true && (v24 == v24)) && true), F1: 0}
                                var v27 bool = true
                                v28 := GenS3{F0: -32768}
                                var v29 uint16 = 65535
                                _ = len(v4)
                                var v30 uint64 = v14
                                _ = v26
                                _ = v27
                                _ = v28
                                _ = v29
                                _ = v30
                            } else {
                                if (!false) {
                                    _ = len(v1)
                                    if (false || true) {
                                        var v31 float32 = 0.1
                                        delete(v3, "foo bar")
                                        _ = v31
                                    }
                                } else {
                                    v24 -= v14
                                    switch 9223372036854775807 {
                                    case 32767:
                                        fmt.Println(fi20, v24, fi23)
                                    default:
                                        var v32 bool = false
                                        _ = v32
                                    }
                                    var v33 uint = 18446744073709551615
                                    var v34 int64 = (-(v21))
                                    _ = v33
                                    _ = v34
                                }
                                var v35 uint8 = v18
                                _ = v35
                            }
                            v1 = append(v1, (("foo bar" + ("a" + ("\n\t" + "\n\t"))) + ("" + "\n\t")))
                            v36 := GenS6{F0: v18, F1: 275, F2: 42}
                            if false {
                                v37 := v3["\n\t"]
                                _ = len(v1)
                                _ = v37
                            } else {
                                v24 = 2
                            }
                            _ = v36
                        } else {
                            if true {
                                switch 42 {
                                case 42:
                                    delete(v3, "a")
                                    var v38 bool = false
                                    _ = v38
                                }
                                v2 += v2
                                v14 += v24
                            }
                            v14 = v24
                        }
                        if false {
                            fmt.Println(v2, v14)
                        }
                        _ = fi23
                        _ = v24
                    }
                    for fi39 := 0; fi39 < 7; fi39++ {
                        switch (fi39 & (fi20 / ((9223372036854775807) | 1))) {
                        case 217:
                            if len(v1) > 0 {
                                v1[0] = "a"
                            }
                        case 4294967295:
                            var v40 uint16 = 42
                            _ = v40
                        case -857:
                            _ = len(v1)
                        }
                        if true {
                            v41 := GenS4{F0: 137, F1: "foo bar", F2: ("hello" + ("" + (("a" + "\n\t") + "foo bar")))}
                            if false {
                                v14 += v14
                                v2 = 100
                                var v42 float32 = (0.0 / 3.14159)
                                _ = v42
                            }
                            switch "foo bar" {
                            case "a":
                                switch (fi39 * (fi39 & int(v2))) {
                                case 1:
                                    v43 := v3[(("foo bar" + ("a" + "foo bar")) + ("hello" + "hello"))]
                                    _ = v43
                                default:
                                    var v44 bool = false
                                    v2 = 44
                                    var v45 float32 = ((-2.25 + 3.14159) * 1.5)
                                    v14 = 18446744073709551615
                                    var v46 float32 = v45
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                            default:
                                var v47 bool = (true && (2147483647 == int32(fi39)))
                                _ = v47
                            }
                            _ = v41
                        } else {
                            var v48 uint32 = 1
                            v49 := GenS5{F0: (v14 <= v14), F1: v48}
                            if (!false) {
                                var v50 float32 = (float32(0.1) * float32(1.5))
                                var v51 uint32 = 1
                                if false {
                                    v52 := GenS1{F0: 343}
                                    var v53 float64 = 0.0
                                    fmt.Println(v53, v18)
                                    _ = v52
                                    _ = v53
                                }
                                _ = v50
                                _ = v51
                            }
                            var v54 uint = uint(v48)
                            _ = v48
                            _ = v49
                            _ = v54
                        }
                        v55 := v3[""]
                        _ = fi39
                        _ = v55
                    }
                    var v56 uint16 = 65535
                    _ = v56
                } else {
                    v21 = (v21 - (v21 * (v21 & v21)))
                    fmt.Println(v2, v2, v14)
                }
                _ = v22
            } else {
                v57 := GenS7{F0: 385, F1: (v18 & (v18 << ((42) & 63))), F2: 4294967295}
                _ = v57
            }
            _ = fi20
            _ = v21
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
