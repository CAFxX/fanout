package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 bool; F1 string; F2 int32 }
type GenS3 struct { F0 uint16; F1 int64; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16; F1 uint; F2 int32 }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 int16; F1 int32 }
type GenS11 struct { F0 int8; F1 uint32; F2 int64 }
func genFn0(p0 int32, p1 string) int64 {
    v7 := GenS1{F0: (255 == 294), F1: true}
    v8 := GenS2{F0: (-106 >= 435), F1: p1, F2: (p0 >> ((993) & 63))}
    if (!false) {
        v9 := []int{9223372036854775807, -2147483648}
        _ = v9
    }
    var v10 string = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint32, p1 uint64, p2 int32) string {
    var v11 int8 = -128
    if false {
        p0 += (p0 & 2)
        for fi12 := 0; fi12 < 14; fi12++ {
            for fi13 := 0; fi13 < 18; fi13++ {
                if (32767 == 400) {
                    switch fi12 {
                    case 0:
                        if (v11 >= (^(v11))) {
                            if (((!true) && (((1.5 / 0.1) - 0.0) >= (0.1 + 1e10))) || true) {
                                v14 := []int{fi13, fi13}
                                var v15 int16 = 32767
                                p1 = (p1 & 18446744073709551615)
                                _ = v14
                                _ = v15
                            } else {
                                v16 := GenS1{F0: (18446744073709551615 >= 18446744073709551615), F1: (!false)}
                                var v17 uint8 = 117
                                var v18 string = "foo bar"
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            v11 = v11
                        } else {
                            switch (fi13 / ((fi12) | 1)) {
                            case 0:
                                p1 -= 42
                            case -2147483648:
                                var v19 uint32 = p0
                                _ = v19
                            case 512:
                                v20 := GenS2{F0: false, F1: "hello", F2: p2}
                                _ = v20
                            }
                            v21 := GenS2{F0: false, F1: ("\n\t" + "hello"), F2: (p2 * p2)}
                            _ = v21
                        }
                    case 42:
                        fmt.Println(p1, p2, p0)
                    default:
                        var v22 uint8 = 53
                        _ = v22
                    }
                }
                _ = fi13
            }
            _ = fi12
        }
        p2 = p2
    } else {
        for fi23 := 0; fi23 < 9; fi23++ {
            p0 = 2
            _ = fi23
        }
        switch "foo bar" {
        case "foo bar":
            switch (("\n\t" + ("a" + (((((("hello" + "") + ("a" + ("" + "a"))) + (("a" + "\n\t") + "")) + "a") + "hello") + "\n\t"))) + "a") {
            case "\n\t":
                for fi24 := 0; fi24 < 5; fi24++ {
                    if false {
                        var v25 uint16 = 65535
                        p2 += (p2 << ((4294967295) & 63))
                        v26 := []int{(fi24 >> ((42) & 63)), 419}
                        _ = v25
                        _ = v26
                    } else {
                        var v27 uint16 = 42
                        _ = v27
                    }
                    _ = fi24
                }
            case "":
                var v28 float32 = -2.25
                _ = v28
            case "hello":
                v29 := []int{}
                _ = v29
            default:
                var v30 float32 = float32(0.1)
                _ = v30
            }
        case "\n\t":
            var v31 float64 = -2.25
            _ = v31
        case "hello":
            v32 := GenS3{F0: 482, F1: -938, F2: 0}
            _ = v32
        }
    }
    v33 := []int{1, -2147483648, -2147483648}
    v34 := GenS2{F0: (false || true), F1: "a", F2: (p2 + -2147483648)}
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v33
    _ = v34
    return ""
}

func main() {
    v1 := []string{"", (("foo bar" + "foo bar") + ""), "hello"}
    var v2 int32 = -1
    var v3 uint8 = 9
    var v4 string = "a"
    var v5 int = 0
    v6 := make(map[string]string)
    var v35 uint32 = 1
    v35 = (v35 >> ((18446744073709551615) & 63))
    v5 -= -537
    if (!false) {
        for fi36 := 0; fi36 < 3; fi36++ {
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 uint = 18446744073709551615
                v39 := GenS1{F0: (((!(100 >= -29)) || (!true)) && false), F1: (true && false)}
                fmt.Println(v38)
                _ = fi37
                _ = v38
                _ = v39
            }
            for fi40 := 0; fi40 < 2; fi40++ {
                v3 -= v3
                fmt.Println(fi36, fi40)
                var v41 uint16 = 32767
                _ = fi40
                _ = v41
            }
            for fi42 := 0; fi42 < 6; fi42++ {
                v3 = (v3 & 1)
                if true {
                    _ = len(v6)
                    var v43 int32 = 515
                    if (!false) {
                        v44 := GenS4{F0: 4294967295, F1: (^(v3))}
                        switch v4 {
                        case "hello":
                            var v45 int8 = 42
                            _ = v45
                        case "foo bar":
                            if ((!false) || ((!(v4 <= v4)) && (!true))) {
                                if (0.1 >= 0.0) {
                                    fmt.Println(v3, fi36)
                                    switch ("" + (v4 + v4)) {
                                    case "\n\t":
                                        if true {
                                            var v46 float64 = float64(float32(-2.25))
                                            v47 := GenS5{F0: 100}
                                            _ = v46
                                            _ = v47
                                        }
                                    case "foo bar":
                                        var v48 int16 = 354
                                        _ = v48
                                    case "":
                                        var v49 bool = true
                                        _ = v49
                                    }
                                    delete(v6, ((("foo bar" + (v4 + ((("a" + ("" + v4)) + ("foo bar" + (v4 + "hello"))) + v4))) + v4) + (v4 + (v4 + v4))))
                                } else {
                                    var v50 int = (fi36 * 514)
                                    v51 := GenS6{F0: 42, F1: v35, F2: v3}
                                    _ = v50
                                    _ = v51
                                }
                                if (false && false) {
                                    if false {
                                        v6[""] = v4
                                        var v52 float64 = (1e10 + -2.25)
                                        _ = v52
                                    } else {
                                        fmt.Println(v35, fi42)
                                        var v53 uint64 = 18446744073709551615
                                        _ = v53
                                    }
                                    var v54 int64 = -918
                                    v35 += v35
                                    if (!false) {
                                        var v55 bool = (!true)
                                        var v56 uint = 2
                                        v57 := GenS7{F0: 964}
                                        _ = v55
                                        _ = v56
                                        _ = v57
                                    } else {
                                        var v58 int8 = -52
                                        v43 = v43
                                        _ = v58
                                    }
                                    v59 := GenS7{F0: (v35 << ((65535) & 63))}
                                    _ = v54
                                    _ = v59
                                } else {
                                    v35 -= v35
                                    switch v4 {
                                    case "hello":
                                        fmt.Println(v4, v5, v5)
                                    default:
                                        fmt.Println(v2, fi42)
                                    }
                                }
                            }
                        default:
                            if ((v4 + v4) != ("a" + (v4 + "a"))) {
                                var v60 int32 = v2
                                v61 := GenS8{F0: 0, F1: 2, F2: (v2 % (((^(v43))) | 1))}
                                fmt.Println(v2)
                                _ = v60
                                _ = v61
                            } else {
                                v62 := GenS9{F0: "\n\t", F1: 109}
                                v35 = v35
                                _ = v62
                            }
                        }
                        v63 := GenS1{F0: false, F1: false}
                        _ = v44
                        _ = v63
                    } else {
                        switch "foo bar" {
                        case "":
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            _ = len(v1)
                            v64 := GenS10{F0: 32767, F1: 2147483647}
                            _ = len(v1)
                            var v65 int = 1
                            _ = v64
                            _ = v65
                        case "hello":
                            var v66 string = ("foo bar" + ("" + v4))
                            _ = v66
                        default:
                            v67 := v6[v4]
                            _ = v67
                        }
                        v68 := GenS7{F0: v35}
                        _ = v68
                    }
                    _ = v43
                }
                _ = fi42
            }
            _ = fi36
        }
        for fi69 := 0; fi69 < 2; fi69++ {
            var v70 uint16 = 784
            if false {
                switch fi69 {
                case 465:
                    _ = len(v1)
                case -2147483648:
                    var v71 bool = false
                    _ = v71
                default:
                    var v72 int64 = -1
                    _ = v72
                }
            }
            _ = fi69
            _ = v70
        }
    }
    switch v5 {
    case 127:
        var v73 int = v5
        _ = v73
    case 499:
        if false {
            v74 := GenS11{F0: 22, F1: (^(v35)), F2: 805}
            _ = v74
        } else {
            fmt.Println(v3, v3)
            var v75 float64 = (3.14159 + 0.0)
            v35 += v35
            _ = v75
        }
    }
    v35 = (v35 + v35)
    if (!true) {
        if ((v3 == v3) && true) {
            var v76 float64 = (-2.25 + 0.1)
            v77 := GenS10{F0: 0, F1: v2}
            _ = v76
            _ = v77
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
