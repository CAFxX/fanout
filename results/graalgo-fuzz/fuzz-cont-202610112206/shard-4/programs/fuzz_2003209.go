package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int; F1 uint32; F2 int8 }
func genFn0(p0 uint64, p1 int32) int64 {
    for fi4 := 0; fi4 < 19; fi4++ {
        p0 -= p0
        _ = fi4
    }
    var v5 uint = 197
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{(fi6 ^ (fi6 % (((fi6 ^ int(p1))) | 1))), 128, fi6, fi6}
        var v8 float64 = (1e10 - (((1.5 * (3.14159 - (3.14159 - -2.25))) * ((float64((1.5 - (3.14159 * (0.0 * 0.0)))) - (1.5 - 1e10)) - (3.14159 - 0.1))) + (1.5 - 3.14159)))
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float32 = -2.25
    p1 = -2147483648
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 899
}
func genFn1(p0 uint32, p1 uint) int16 {
    v10 := []int{0, 32767, 9223372036854775807, -317}
    switch ("foo bar" + "") {
    case "hello":
        switch 0 {
        case 4294967295:
            if (true && true) {
                v11 := GenS1{F0: genFn0(699, -2147483648)}
                if (128 <= -625) {
                    v12 := GenS1{F0: -9223372036854775808}
                    _ = v12
                }
                _ = v11
            } else {
                v13 := GenS1{F0: -1}
                _ = v13
            }
        case 2147483647:
            if true {
                p1 += 42
            }
        }
    case "":
        var v14 uint16 = 65535
        _ = v14
    case "foo bar":
        var v15 string = ""
        _ = v15
    default:
        p0 = 1
    }
    var v16 uint = (^(p1))
    var v17 int64 = -30
    _ = p0
    _ = p1
    _ = v10
    _ = v16
    _ = v17
    return -32768
}

func main() {
    v1 := []string{"hello", "", (("hello" + "") + ""), ""}
    v2 := make(map[uint32]string)
    var v3 bool = (!true)
    v1 = append(v1, "")
    v3 = (float64(0.0) >= (1e10 * 1e10))
    var v18 int8 = 13
    v3 = ((-2147483648 < 42) || false)
    v19 := GenS1{F0: 0}
    for fi20 := 0; fi20 < 7; fi20++ {
        switch "\n\t" {
        case "hello":
            v3 = (42 <= 2147483648)
        case "a":
            for fi21 := 0; fi21 < 6; fi21++ {
                if len(v1) > 0 {
                    v1[0] = ("\n\t" + ((((("" + (((("\n\t" + "hello") + ("foo bar" + "foo bar")) + (("\n\t" + "a") + "hello")) + (("\n\t" + "\n\t") + "foo bar"))) + "\n\t") + "hello") + ("\n\t" + ("hello" + ("\n\t" + "foo bar")))) + ((("" + ("a" + "hello")) + ("" + "foo bar")) + (("foo bar" + ("foo bar" + ("hello" + "hello"))) + "\n\t"))))
                }
                _ = fi21
            }
        case "foo bar":
            var v22 int8 = (v18 & (v18 % ((0) | 1)))
            _ = v22
        default:
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        }
        fmt.Println(v18, fi20)
        if true {
            _ = len(v2)
            switch "" {
            case "foo bar":
                if v3 {
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v18 = int8(fi23)
                        v3 = (v3 && (v3 || (!true)))
                        _ = len(v2)
                        _ = fi23
                    }
                    var v24 int64 = -9223372036854775808
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        v3 = (!v3)
                        v18 = v18
                        _ = fi25
                    }
                    v18 = (v18 - v18)
                }
            }
        } else {
            switch (fi20 / (((fi20 / ((fi20) | 1))) | 1)) {
            case -326:
                switch 42 {
                case -305:
                    if (42 > 42) {
                        var v26 uint = 2
                        var v27 string = "a"
                        switch (fi20 << (((v26 + v26)) & 63)) {
                        case 32768:
                            if v3 {
                                v28 := v2[42]
                                _ = v28
                            }
                        default:
                            var v29 uint32 = 2
                            _ = v29
                        }
                        _ = v26
                        _ = v27
                    } else {
                        v2[2147483647] = "hello"
                        v30 := GenS1{F0: 9223372036854775807}
                        _ = v30
                    }
                case -9223372036854775808:
                    var v31 uint = 0
                    _ = v31
                }
            case 32768:
                v18 = v18
            }
        }
        _ = fi20
    }
    var v32 int64 = -1
    v33 := GenS1{F0: v32}
    if v3 {
        for fi34 := 0; fi34 < 2; fi34++ {
            if (!v3) {
                if false {
                    if (!("hello" >= "hello")) {
                        switch (fi34 >> ((255) & 63)) {
                        case -1:
                            fmt.Println(v3)
                        case 873:
                            v3 = (!v3)
                        case 32767:
                            var v35 uint16 = 65535
                            _ = v35
                        }
                        if false {
                            if v3 {
                                if ((0.0 == -2.25) || (1 <= -583)) {
                                    fmt.Println(v3, v3, fi34)
                                    v36 := GenS1{F0: v32}
                                    fmt.Println(v32, v18)
                                    _ = v36
                                }
                                v32 = (v32 + (v32 * genFn0(0, 42)))
                            } else {
                                if (!(552 <= 679)) {
                                    v3 = ((1 > 42) && v3)
                                }
                                var v37 bool = ((-9223372036854775808 <= fi34) || v3)
                                _ = v37
                            }
                            v18 = v18
                            v38 := GenS2{F0: 255}
                            _ = v38
                        } else {
                            v3 = v3
                        }
                        var v39 int16 = 781
                        _ = v39
                    } else {
                        v40 := GenS1{F0: (v32 >> ((1) & 63))}
                        _ = v40
                    }
                }
                fmt.Println(fi34)
                v41 := GenS3{F0: "foo bar", F1: 18446744073709551615}
                _ = v41
            } else {
                for fi42 := 0; fi42 < 2; fi42++ {
                    switch 0 {
                    case -277:
                        if v3 {
                            v43 := GenS4{F0: fi34, F1: 0, F2: (v18 << ((1) & 63))}
                            fmt.Println(v18, v32)
                            _ = v43
                        } else {
                            var v44 float32 = 1e10
                            _ = v44
                        }
                    case 2147483647:
                        var v45 int64 = 0
                        _ = v45
                    case 126:
                        var v46 uint = 18446744073709551615
                        _ = v46
                    }
                    _ = fi42
                }
                var v47 int32 = -511
                _ = v47
            }
            _ = fi34
        }
    } else {
        var v48 uint8 = 42
        _ = v48
    }
    fmt.Println(v18)
    switch (("a" + "foo bar") + "a") {
    case "foo bar":
        v18 = v18
    default:
        var v49 uint32 = 65535
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
}
