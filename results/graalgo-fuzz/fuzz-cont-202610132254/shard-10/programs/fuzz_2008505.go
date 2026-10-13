package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int64; F1 uint32; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 int64; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int }
func genFn0(p0 uint16, p1 string, p2 int) bool {
    v5 := []int{286, p2, (p2 & (p2 % ((p2) | 1))), p2}
    var v6 int16 = 433
    fmt.Println(p1, p1, p1)
    var v7 uint16 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 string) int {
    for fi8 := 0; fi8 < 18; fi8++ {
        p0 = (p0 + "")
        v9 := []int{(fi8 + (fi8 << ((835) & 63))), fi8}
        switch fi8 {
        case -520:
            fmt.Println(p0)
        }
        _ = fi8
        _ = v9
    }
    p0 = (p0 + p0)
    v10 := []int{127}
    v11 := GenS1{F0: 2147483648, F1: 165}
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 int64 = int64(fi12)
            var v15 int64 = -9223372036854775808
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    _ = p0
    _ = v10
    _ = v11
    return 32767
}

func main() {
    v1 := []string{"foo bar", ("hello" + ("foo bar" + "hello"))}
    var v2 uint16 = 42
    v3 := make(map[string]int)
    var v4 uint32 = 137
    v3[((("a" + "hello") + "hello") + "foo bar")] = -244
    if (false || false) {
        v4 = v4
    }
    if ((4294967295 < 4294967295) || (!false)) {
        var v16 int64 = 9223372036854775807
        switch 127 {
        case 9223372036854775807:
            v2 += 255
        case 0:
            switch "" {
            case "foo bar":
                var v17 int64 = (-(v16))
                _ = v17
            case "hello":
                if (100 > -59) {
                    var v18 bool = false
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v20 := GenS2{F0: v16, F1: v4, F2: "hello"}
                        delete(v3, "")
                        _ = fi19
                        _ = v20
                    }
                    var v21 float64 = ((1e10 / 0.1) * 1.5)
                    _ = v18
                    _ = v21
                } else {
                    v4 = v4
                    if (!((3.14159 - (1.5 + 0.0)) > (1e10 * (float32(((3.14159 + -2.25) * 0.1)) * (((-2.25 - (((((((1.5 + (1e10 - 0.1)) * 0.1) / 1e10) - 1e10) * ((((1e10 / float32(1e10)) / (((1.5 * 1.5) + 3.14159) - 1.5)) / 1.0) + 1.5)) / 3.14159) - -2.25)) / float32((float64(1e10) + 0.1))) - 1.5))))) {
                        var v22 uint16 = v2
                        for fi23 := 0; fi23 < 16; fi23++ {
                            var v24 bool = ((true && true) && true)
                            v25 := GenS2{F0: 1, F1: v4, F2: "a"}
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        _ = v22
                    } else {
                        switch 544 {
                        case 356:
                            v26 := GenS1{F0: 2147483648, F1: 1}
                            _ = v26
                        default:
                            v27 := GenS1{F0: (v4 - 271), F1: -452}
                            _ = v27
                        }
                        if ((-2.25 >= 0.0) || (!true)) {
                            fmt.Println(v4, v4)
                            fmt.Println(v16, v2)
                        } else {
                            switch "a" {
                            case "a":
                                var v28 int = int(v4)
                                _ = v28
                            case "":
                                for fi29 := 0; fi29 < 9; fi29++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi29
                                }
                            }
                        }
                        for fi30 := 0; fi30 < 3; fi30++ {
                            var v31 uint = 18446744073709551615
                            _ = fi30
                            _ = v31
                        }
                    }
                }
            default:
                if true {
                    v32 := GenS3{F0: 75}
                    _ = v32
                }
            }
        case 1:
            switch ("\n\t" + "a") {
            case "hello":
                v33 := GenS2{F0: (v16 / ((v16) | 1)), F1: 42, F2: "\n\t"}
                _ = v33
            default:
                delete(v3, (((("foo bar" + "\n\t") + "") + "foo bar") + "\n\t"))
            }
        default:
            _ = len(v1)
        }
        v4 += 1
        _ = v16
    } else {
        var v34 int = 0
        _ = v34
    }
    if (false || (!false)) {
        v2 = v2
        var v35 int32 = 216
        _ = v35
    } else {
        var v36 uint64 = 4294967295
        fmt.Println(v2, v2)
        _ = v36
    }
    v4 += (v4 << ((0) & 63))
    v3[""] = genFn1("foo bar")
    for fi37 := 0; fi37 < 16; fi37++ {
        var v38 float64 = 0.0
        _ = fi37
        _ = v38
    }
    var v39 uint64 = 4294967295
    var v40 float64 = 0.0
    if (18446744073709551615 == 27) {
        v41 := GenS4{F0: 0, F1: -1, F2: "foo bar"}
        _ = v41
    }
    switch 80 {
    case 4294967295:
        if (((42 >= 0) && (((true && true) && true) && false)) && ("hello" > "")) {
            if (false || ((((3.14159 / 1e10) + 1.5) < (float32(float64(0.1)) + 1.5)) || ((((1e10 + -2.25) * 0.0) - 3.14159) > float32(v40)))) {
                var v42 int64 = -1
                if (true && true) {
                    if true {
                        for fi43 := 0; fi43 < 20; fi43++ {
                            v40 += v40
                            v44 := GenS2{F0: v42, F1: v4, F2: "hello"}
                            _ = fi43
                            _ = v44
                        }
                        _ = len(v1)
                    }
                    switch 423 {
                    case 9223372036854775807:
                        delete(v3, ("hello" + "hello"))
                    case 127:
                        for fi45 := 0; fi45 < 9; fi45++ {
                            v46 := v3["hello"]
                            _ = fi45
                            _ = v46
                        }
                    case 4294967295:
                        var v47 uint64 = v39
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 16; fi48++ {
                        v2 = 65535
                        if false {
                            if ("" != "\n\t") {
                                v3[("a" + (("a" + ("a" + "\n\t")) + "foo bar"))] = genFn1("a")
                                for fi49 := 0; fi49 < 2; fi49++ {
                                    var v50 int = (fi48 >> ((850) & 63))
                                    _ = fi49
                                    _ = v50
                                }
                                switch genFn1(("\n\t" + ("foo bar" + ("\n\t" + "hello")))) {
                                case 127:
                                    v3[("a" + "foo bar")] = fi48
                                case 128:
                                    var v51 int = genFn1(("foo bar" + "a"))
                                    _ = v51
                                default:
                                    var v52 uint64 = (^(v39))
                                    _ = v52
                                }
                            }
                        }
                        _ = fi48
                    }
                } else {
                    var v53 int32 = 0
                    fmt.Println(v2, v53)
                    switch 32767 {
                    case 127:
                        switch ("\n\t" + ("\n\t" + ("a" + (("\n\t" + "\n\t") + (("foo bar" + "foo bar") + ("" + "hello")))))) {
                        case "\n\t":
                            switch (("a" + ("" + "a")) + "\n\t") {
                            case "":
                                if true {
                                    v1 = append(v1, "foo bar")
                                    var v54 uint32 = (v4 << ((65535) & 63))
                                    _ = v54
                                }
                            }
                        case "hello":
                            v55 := GenS4{F0: 9223372036854775807, F1: (v42 >> ((666) & 63)), F2: "\n\t"}
                            _ = v55
                        case "foo bar":
                            fmt.Println(v40)
                        default:
                            if true {
                                v40 += v40
                            } else {
                                v56 := GenS5{F0: (^(v2))}
                                var v57 uint8 = 122
                                _ = len(v1)
                                _ = v56
                                _ = v57
                            }
                        }
                    default:
                        var v58 float32 = 3.14159
                        _ = v58
                    }
                    _ = v53
                }
                _ = v42
            } else {
                v3[""] = -1
                switch "hello" {
                case "":
                    v39 -= (v39 | (^(v39)))
                case "foo bar":
                    v39 -= 329
                case "\n\t":
                    var v59 uint32 = v4
                    _ = v59
                }
                if (42 != -905) {
                    fmt.Println(v4, v39, v4)
                } else {
                    var v60 bool = (42 == -74)
                    _ = v60
                }
            }
            var v61 int8 = 0
            _ = v61
        }
    }
    var v62 uint64 = v39
    delete(v3, "a")
    v63 := GenS6{F0: -178}
    if false {
        fmt.Println(v62)
        for fi64 := 0; fi64 < 7; fi64++ {
            v4 -= 1
            _ = fi64
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v62
    _ = v63
}
