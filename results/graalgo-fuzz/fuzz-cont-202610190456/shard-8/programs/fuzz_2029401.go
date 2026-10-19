package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int32; F2 int }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int16, p1 string) int {
    var v7 int32 = -212
    if true {
        for fi8 := 0; fi8 < 6; fi8++ {
            switch p1 {
            case "":
                for fi9 := 0; fi9 < 10; fi9++ {
                    var v10 bool = false
                    v11 := GenS1{F0: 2147483648, F1: (fi9 % ((fi9) | 1)), F2: 42}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case "a":
                v12 := GenS1{F0: 709, F1: 9223372036854775807, F2: 1}
                _ = v12
            case "foo bar":
                if (((((p1 + (p1 + "hello")) + (p1 + p1)) + p1) + p1) > "\n\t") {
                    p1 = "hello"
                    v13 := GenS1{F0: 22, F1: (fi8 % ((410) | 1)), F2: 202}
                    var v14 bool = ("" >= p1)
                    _ = v13
                    _ = v14
                }
            }
            _ = fi8
        }
        p0 += 32767
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 float32 = (float32(1.5) - 1e10)
    v2 := make(map[string]int)
    v3 := []string{}
    var v4 int = 966
    v5 := make(map[int]string)
    var v6 uint32 = 2147483648
    if true {
        for fi15 := 0; fi15 < 1; fi15++ {
            if true {
                if true {
                    if false {
                        v16 := GenS2{F0: -9223372036854775808, F1: (v6 * (^(v6)))}
                        v1 = v1
                        var v17 uint8 = 255
                        _ = v16
                        _ = v17
                    }
                    var v18 int = (fi15 * genFn0(0, ""))
                    _ = v18
                } else {
                    var v19 bool = (!(42 >= 489))
                    _ = v19
                }
                v20 := GenS3{F0: 1}
                v21 := GenS4{F0: 132, F1: 806, F2: 32768}
                _ = v20
                _ = v21
            }
            v22 := v5[v4]
            _ = fi15
            _ = v22
        }
    } else {
        fmt.Println(v4, v4)
    }
    if false {
        if true {
            for fi23 := 0; fi23 < 12; fi23++ {
                v6 = v6
                if len(v3) > 0 {
                    v3[0] = ""
                }
                fmt.Println(v1, v4)
                _ = fi23
            }
            v2[(("\n\t" + "foo bar") + "hello")] = 42
            switch 2147483647 {
            case 4294967295:
                delete(v5, v4)
            case 760:
                v6 += 65535
            case 1:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                if (false && (9223372036854775807 <= -9223372036854775808)) {
                    if (!(true && false)) {
                        if true {
                            switch v4 {
                            case 32768:
                                var v24 uint32 = (v6 + v6)
                                _ = v24
                            case 2147483647:
                                var v25 uint = 65535
                                _ = v25
                            case 9223372036854775807:
                                delete(v5, (v4 - genFn0(1, "foo bar")))
                            }
                        } else {
                            for fi26 := 0; fi26 < 18; fi26++ {
                                v1 -= (v1 - 1.5)
                                var v27 int16 = 901
                                v28 := GenS5{F0: 0, F1: ("" + "")}
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        }
                        delete(v2, "a")
                    }
                    v4 -= v4
                }
            }
        }
    } else {
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    }
    switch (^(v4)) {
    case 194:
        v1 += (float32(0.1) / v1)
    }
    if (true && (164 != 0)) {
        fmt.Println(v4)
        fmt.Println(v4)
    }
    v29 := GenS6{F0: 101}
    var v30 uint16 = 1
    if (!false) {
        var v31 int32 = -773
        fmt.Println(v6, v6, v6)
        _ = v31
    } else {
        switch ("foo bar" + ("" + "a")) {
        case "":
            v2["hello"] = 1
        case "\n\t":
            var v32 uint16 = (v30 / ((v30) | 1))
            _ = v32
        }
        var v33 int32 = -1
        if (-1 > 127) {
            var v34 int8 = -22
            fmt.Println(v4, v4, v1)
            v33 = v33
            _ = v34
        }
        _ = v33
    }
    var v35 int64 = -358
    fmt.Println(v30)
    v2["\n\t"] = (v4 % ((v4) | 1))
    switch "" {
    case "":
        for fi36 := 0; fi36 < 15; fi36++ {
            v6 += (v6 / ((0) | 1))
            var v37 int64 = (^(v35))
            if ((v6 | (v6 >> ((0) & 63))) < v6) {
                if (4294967295 < 4294967295) {
                    for fi38 := 0; fi38 < 3; fi38++ {
                        var v39 uint32 = 1
                        if true {
                            fmt.Println(v35, v1, v37)
                            var v40 int32 = 382
                            _ = v40
                        } else {
                            fmt.Println(fi36, v1, v1)
                            if (545 <= -621) {
                                v5[(-(fi36))] = "foo bar"
                                v6 = (v39 >> ((0) & 63))
                                v41 := v5[v4]
                                _ = v41
                            } else {
                                _ = len(v3)
                                var v42 int16 = -1
                                v2[("a" + ("\n\t" + "a"))] = v4
                                var v43 float64 = 0.1
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                                var v44 float64 = 1e10
                                var v45 uint64 = 486
                                _ = v42
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                            v39 = (v39 >> ((2) & 63))
                            if ((-124 < 1) && ("" > "hello")) {
                                fmt.Println(fi38, v1)
                            } else {
                                v46 := GenS4{F0: (v30 ^ uint16(v39)), F1: -737, F2: -328}
                                var v47 int = fi38
                                _ = v46
                                _ = v47
                            }
                        }
                        delete(v5, genFn0(1, "a"))
                        _ = fi38
                        _ = v39
                    }
                } else {
                    switch "a" {
                    case "\n\t":
                        switch (fi36 + (v4 * v4)) {
                        case 128:
                            var v48 float64 = 0.1
                            _ = v48
                        case -1:
                            v49 := v5[genFn0(0, "a")]
                            _ = v49
                        case 9223372036854775807:
                            if (float64((-2.25 / v1)) >= 3.14159) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                var v50 uint64 = 1
                                switch "foo bar" {
                                case "":
                                    v51 := GenS5{F0: 42, F1: "a"}
                                    _ = v51
                                case "a":
                                    v52 := v5[v4]
                                    _ = v52
                                }
                                _ = v50
                            }
                        }
                    }
                    var v53 uint = 2
                    v5[(-(v4))] = "foo bar"
                    _ = v53
                }
            } else {
                v35 -= -9223372036854775808
                if len(v3) > 0 {
                    v3[0] = ("\n\t" + "hello")
                }
            }
            _ = fi36
            _ = v37
        }
    default:
        v5[v4] = "hello"
    }
    fmt.Println(v6, v35)
    if false {
        fmt.Println(v30)
    }
    for fi54 := 0; fi54 < 7; fi54++ {
        var v55 uint = 4294967295
        for fi56 := 0; fi56 < 20; fi56++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi56
        }
        _ = fi54
        _ = v55
    }
    for fi57 := 0; fi57 < 16; fi57++ {
        v58 := GenS1{F0: v6, F1: (fi57 & (v4 ^ (fi57 << ((18446744073709551615) & 63)))), F2: (v30 & 255)}
        if (true && false) {
            var v59 uint = uint(v30)
            switch (("hello" + (("" + (("\n\t" + "\n\t") + "")) + ("foo bar" + ("foo bar" + ("\n\t" + "a"))))) + "") {
            case "":
                v60 := v2["\n\t"]
                _ = v60
            case "foo bar":
                v59 = v59
            }
            _ = v59
        }
        _ = fi57
        _ = v58
    }
    fmt.Println(v1, v6, v35)
    var v61 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
