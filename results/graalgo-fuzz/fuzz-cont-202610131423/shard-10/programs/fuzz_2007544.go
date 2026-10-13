package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 string }
type GenS2 struct { F0 int16; F1 int32; F2 bool }
type GenS3 struct { F0 string; F1 string; F2 int }
type GenS4 struct { F0 bool; F1 int8 }
func genFn0(p0 uint) uint {
    p0 += p0
    p0 -= p0
    switch 128 {
    case -951:
        v5 := []int{}
        _ = v5
    default:
        p0 += (p0 % ((p0) | 1))
    }
    _ = p0
    return 0
}
func genFn1(p0 int8, p1 int16, p2 int64) string {
    if true {
        var v6 int16 = p1
        var v7 float32 = 1.5
        _ = v6
        _ = v7
    } else {
        p2 -= p2
    }
    var v8 int16 = (-(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return ""
}

func main() {
    var v1 uint16 = 42
    var v2 uint64 = 626
    v3 := []string{"\n\t", "\n\t"}
    var v4 uint = 2
    switch 0 {
    case -925:
        var v9 int32 = 42
        _ = v9
    case 691:
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    }
    var v10 uint16 = (^(v1))
    v2 += (v2 | v2)
    fmt.Println(v4)
    var v11 int64 = 722
    if false {
        v12 := GenS1{F0: "a", F1: v4, F2: "foo bar"}
        v3 = append(v3, (("\n\t" + "") + ("" + ("" + ("" + ((("hello" + "hello") + "foo bar") + "hello"))))))
        if (false && true) {
            v13 := GenS1{F0: "foo bar", F1: 4294967295, F2: (((("foo bar" + ((((("\n\t" + "foo bar") + ((("" + "") + "hello") + "foo bar")) + "\n\t") + ("foo bar" + "")) + "")) + "") + (("hello" + (("\n\t" + "foo bar") + "hello")) + "hello")) + "\n\t")}
            if (!(false || (!(!(true && false))))) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v4 = genFn0(v4)
            _ = v13
        }
        _ = v12
    } else {
        v14 := GenS2{F0: 1, F1: -1, F2: (-1 >= 100)}
        v4 = (v4 * (v4 ^ (v4 + 4294967295)))
        _ = v14
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        if (!false) {
            for fi16 := 0; fi16 < 16; fi16++ {
                v11 -= v11
                v17 := GenS1{F0: ("\n\t" + (("hello" + "") + ("\n\t" + ("\n\t" + "hello")))), F1: (^(v4)), F2: "hello"}
                var v18 uint32 = 1
                var v19 int64 = (v11 + (v11 * (-(v11))))
                v20 := GenS2{F0: -52, F1: 1, F2: false}
                var v21 int = int(v18)
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
                _ = v20
                _ = v21
            }
        } else {
            var v22 string = ("a" + "a")
            var v23 uint32 = 0
            for fi24 := 0; fi24 < 17; fi24++ {
                fmt.Println(v22, v23, v22)
                var v25 float64 = 0.1
                _ = fi24
                _ = v25
            }
            _ = v22
            _ = v23
        }
        v4 += 496
        _ = fi15
    }
    switch 46 {
    case 127:
        for fi26 := 0; fi26 < 5; fi26++ {
            v10 = v10
            _ = fi26
        }
    case 32768:
        if false {
            v3 = append(v3, "")
            for fi27 := 0; fi27 < 20; fi27++ {
                if (1 < 42) {
                    v1 = (^(v1))
                    v28 := GenS2{F0: 874, F1: 42, F2: (0.0 == 0.0)}
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 6; fi29++ {
                        v30 := GenS3{F0: ("\n\t" + ("a" + "hello")), F1: ("\n\t" + ("\n\t" + ("a" + "foo bar"))), F2: 1}
                        switch (fi27 / (((fi29 - -9223372036854775808)) | 1)) {
                        case -9223372036854775808:
                            var v31 int16 = 754
                            _ = v31
                        case 498:
                            if false {
                                if (141 != 0) {
                                    var v32 uint16 = (^(v10))
                                    switch -2147483648 {
                                    case 127:
                                        var v33 float32 = (1.5 + (float32(-2.25) / 1.0))
                                        _ = v33
                                    }
                                    v10 = (v10 % ((v10) | 1))
                                    _ = v32
                                }
                            }
                        }
                        _ = fi29
                        _ = v30
                    }
                    var v34 bool = (v4 <= (v4 | (v4 >> ((65535) & 63))))
                    var v35 uint64 = v2
                    _ = v34
                    _ = v35
                }
                var v36 float32 = 3.14159
                _ = fi27
                _ = v36
            }
        } else {
            v4 = (v4 | v4)
            var v37 uint16 = (v10 << ((v4) & 63))
            fmt.Println(v10, v1)
            _ = v37
        }
    default:
        switch "\n\t" {
        case "foo bar":
            v38 := GenS3{F0: (("\n\t" + ("\n\t" + "")) + "a"), F1: ((("foo bar" + "") + (((("\n\t" + "") + "") + "a") + "\n\t")) + ((("" + "hello") + "a") + "\n\t")), F2: -728}
            _ = v38
        default:
            v39 := GenS2{F0: 1, F1: 0, F2: false}
            _ = v39
        }
    }
    fmt.Println(v10, v4)
    switch "\n\t" {
    case "a":
        var v40 bool = false
        _ = v40
    case "hello":
        v41 := GenS4{F0: true, F1: 100}
        _ = v41
    case "\n\t":
        var v42 float32 = (1e10 + 3.14159)
        _ = v42
    }
    v4 = genFn0((v4 - v4))
    var v43 uint = genFn0((v4 << ((v4) & 63)))
    v44 := GenS3{F0: ("hello" + ""), F1: ("" + ""), F2: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
}
