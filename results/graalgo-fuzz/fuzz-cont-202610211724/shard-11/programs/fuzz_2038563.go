package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int32; F2 int16 }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 uint32; F1 uint8; F2 string }
type GenS7 struct { F0 bool; F1 uint64; F2 int64 }
func genFn0(p0 int16) uint64 {
    p0 += p0
    v7 := GenS1{F0: (-70 != 127), F1: 95}
    var v8 float64 = 3.14159
    _ = p0
    _ = v7
    _ = v8
    return 2
}
func genFn1() int16 {
    switch "a" {
    case "\n\t":
        v9 := GenS2{F0: 42, F1: 548}
        _ = v9
    case "":
        var v10 uint64 = 142
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 10; fi11++ {
            switch fi11 {
            case 127:
                if true {
                    var v12 uint32 = 2
                    v13 := []int{}
                    _ = v12
                    _ = v13
                } else {
                    if ((1 == -51) || ("a" > "a")) {
                        var v14 uint = 18446744073709551615
                        _ = v14
                    } else {
                        var v15 uint = 4294967295
                        if true {
                            fmt.Println(v15, v15)
                        } else {
                            v15 = (v15 - (^(v15)))
                        }
                        v16 := GenS1{F0: ((fi11 & fi11) > fi11), F1: 255}
                        _ = v15
                        _ = v16
                    }
                }
            }
            v17 := GenS1{F0: (1 < 42), F1: 197}
            _ = fi11
            _ = v17
        }
    }
    var v18 uint = 1
    if (!(32767 > 2)) {
        fmt.Println(v18)
        v19 := GenS3{F0: ("\n\t" + "a")}
        _ = v19
    } else {
        if true {
            v18 -= 65535
        }
        fmt.Println(v18)
    }
    _ = v18
    return -32768
}

func main() {
    v1 := []int{-160, -865, -773, -1}
    v2 := []uint32{}
    var v3 int16 = -1
    var v4 float64 = 0.0
    v5 := make(map[string]bool)
    var v6 uint32 = 4294967295
    var v20 uint8 = 203
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := GenS4{F0: true, F1: 2147483647, F2: (v3 * (^(v3)))}
        if (0.1 == 3.14159) {
            switch (((("foo bar" + "\n\t") + (("hello" + ("a" + "\n\t")) + ("" + "hello"))) + ("foo bar" + ("\n\t" + "\n\t"))) + ("\n\t" + "hello")) {
            case "hello":
                for fi23 := 0; fi23 < 14; fi23++ {
                    var v24 int = fi23
                    v4 = (float64(((1e10 * 3.14159) - -2.25)) - v4)
                    _ = len(v5)
                    fmt.Println(fi23, v20, fi21)
                    _ = fi23
                    _ = v24
                }
            case "a":
                switch (fi21 << ((281) & 63)) {
                case 4294967295:
                    for fi25 := 0; fi25 < 19; fi25++ {
                        var v26 int16 = (v3 | 42)
                        v6 -= 0
                        if len(v1) > 0 {
                            v1[0] = -692
                        }
                        _ = fi25
                        _ = v26
                    }
                default:
                    if false {
                        v3 = v3
                        var v27 uint16 = 1
                        _ = v27
                    } else {
                        v28 := GenS1{F0: false, F1: 127}
                        v29 := GenS5{F0: 2, F1: -2147483648}
                        fmt.Println(v3)
                        _ = v28
                        _ = v29
                    }
                }
            default:
                v3 = v3
            }
            switch "a" {
            case "foo bar":
                _ = len(v2)
            case "":
                v30 := v5[("hello" + "\n\t")]
                _ = v30
            case "\n\t":
                fmt.Println(v20)
            default:
                switch "\n\t" {
                case "a":
                    for fi31 := 0; fi31 < 1; fi31++ {
                        var v32 float32 = float32(v4)
                        _ = fi31
                        _ = v32
                    }
                case "foo bar":
                    for fi33 := 0; fi33 < 3; fi33++ {
                        fmt.Println(fi33, fi33, v6)
                        v34 := GenS1{F0: true, F1: 222}
                        v3 -= (v3 - (-(v3)))
                        if false {
                            v35 := GenS6{F0: (^(v6)), F1: v20, F2: "a"}
                            v5[((("foo bar" + "") + ("\n\t" + "foo bar")) + "hello")] = (146 != 42)
                            v20 = v20
                            fmt.Println(v3, fi21, fi21)
                            _ = v35
                        } else {
                            if len(v2) > 0 {
                                v2[0] = (v6 % (((v6 >> ((2) & 63))) | 1))
                            }
                        }
                        if true {
                            v36 := GenS6{F0: (v6 >> ((65535) & 63)), F1: v20, F2: (((("\n\t" + ("" + "")) + "\n\t") + "") + ("foo bar" + "hello"))}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch "hello" {
                            case "foo bar":
                                var v37 string = "foo bar"
                                _ = v37
                            case "":
                                delete(v5, "foo bar")
                            case "a":
                                v4 = v4
                            default:
                                var v38 uint64 = 320
                                var v39 int8 = -33
                                _ = v38
                                _ = v39
                            }
                            var v40 uint64 = 20
                            _ = len(v2)
                            _ = v36
                            _ = v40
                        }
                        if (!((^(v20)) <= v20)) {
                            v41 := GenS3{F0: "foo bar"}
                            _ = v41
                        } else {
                            fmt.Println(v20, fi33)
                            _ = len(v5)
                            if false {
                                var v42 uint64 = 322
                                v43 := GenS3{F0: ""}
                                _ = v42
                                _ = v43
                            }
                        }
                        var v44 uint8 = (v20 / ((v20) | 1))
                        switch 9223372036854775807 {
                        case 1:
                            var v45 int32 = int32(v44)
                            _ = v45
                        }
                        _ = fi33
                        _ = v34
                        _ = v44
                    }
                case "\n\t":
                    v3 -= (v3 ^ (v3 + (v3 << ((1) & 63))))
                default:
                    for fi46 := 0; fi46 < 1; fi46++ {
                        v47 := GenS3{F0: ((("foo bar" + "\n\t") + "a") + "\n\t")}
                        v6 += (v6 * v6)
                        _ = len(v2)
                        switch 1 {
                        case 0:
                            switch "\n\t" {
                            case "":
                                var v48 int32 = 2147483647
                                _ = v48
                            }
                        case 32768:
                            delete(v5, "hello")
                        case -2147483648:
                            delete(v5, "a")
                        }
                        _ = fi46
                        _ = v47
                    }
                }
            }
            if ((false || (!(genFn0((v3 % (((v3 * (v3 / (((v3 & int16(v20))) | 1)))) | 1))) > 18446744073709551615))) || false) {
                fmt.Println(v20, v3)
                if (127 >= fi21) {
                    v49 := GenS7{F0: (v4 < v4), F1: 18446744073709551615, F2: -9223372036854775808}
                    v50 := v5[("hello" + "a")]
                    if (0 >= 1) {
                        v6 -= 149
                    } else {
                        if v50 {
                            var v51 string = ""
                            for fi52 := 0; fi52 < 16; fi52++ {
                                v3 = v3
                                v53 := GenS5{F0: 146, F1: 124}
                                var v54 uint = 26
                                _ = fi52
                                _ = v53
                                _ = v54
                            }
                            _ = v51
                        }
                        v55 := v5[("hello" + "\n\t")]
                        _ = v55
                    }
                    _ = v49
                    _ = v50
                } else {
                    v6 -= (v6 % ((v6) | 1))
                    var v56 int16 = genFn1()
                    _ = v56
                }
            }
        }
        fmt.Println(v3)
        _ = fi21
        _ = v22
    }
    v20 = v20
    v3 = v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "a" {
    case "a":
        v5[("hello" + "")] = (1 <= 4294967295)
    case "hello":
        fmt.Println(v4, v3)
    }
    switch -2147483648 {
    case 128:
        v4 -= v4
    case -774:
        for fi57 := 0; fi57 < 14; fi57++ {
            var v58 int = fi57
            delete(v5, "\n\t")
            for fi59 := 0; fi59 < 11; fi59++ {
                v20 -= 42
                _ = fi59
            }
            _ = fi57
            _ = v58
        }
    }
    switch 127 {
    case 32768:
        v3 -= v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
