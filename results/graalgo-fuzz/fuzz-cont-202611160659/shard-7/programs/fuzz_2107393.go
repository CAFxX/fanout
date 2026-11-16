package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int32; F2 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 int }
func genFn0(p0 int) uint8 {
    p0 -= p0
    v5 := GenS1{F0: 65535}
    _ = p0
    _ = v5
    return 100
}
func genFn1(p0 string) uint16 {
    p0 = ""
    v6 := GenS2{F0: 65535}
    _ = p0
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 1
    v3 := []string{}
    v4 := []int{int(v2), 32767}
    v2 -= 235
    switch -9223372036854775808 {
    case -2147483648:
        for fi7 := 0; fi7 < 10; fi7++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 -= 562
            _ = fi7
        }
    case 128:
        v8 := GenS1{F0: 0}
        _ = v8
    case 42:
        fmt.Println(v2)
    default:
        var v9 int16 = 0
        _ = v9
    }
    if ((v2 + genFn1(("hello" + (("\n\t" + (((("hello" + "foo bar") + ("\n\t" + "")) + "") + "\n\t")) + "a")))) < (v2 * 255)) {
        v10 := GenS2{F0: 1}
        var v11 uint16 = (v2 ^ v2)
        _ = v10
        _ = v11
    } else {
        switch 785 {
        case -355:
            fmt.Println(v2)
        default:
            switch 42 {
            case 590:
                if (!true) {
                    var v12 uint16 = 153
                    v13 := GenS3{F0: 2, F1: 81, F2: "hello"}
                    _ = v12
                    _ = v13
                }
            }
        }
        v14 := GenS1{F0: 1}
        fmt.Println(v2)
        _ = v14
    }
    v2 -= (^(v2))
    v2 = (v2 + v2)
    v2 = 0
    v2 = genFn1("hello")
    v2 = v2
    var v15 string = "hello"
    fmt.Println(v15, v15)
    if (65535 <= 42) {
        if (!true) {
            switch -243 {
            case -9223372036854775808:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                v15 = v15
            }
            var v16 uint8 = 0
            if (v15 <= "") {
                v17 := GenS3{F0: v16, F1: 0, F2: (("foo bar" + (v15 + v15)) + v15)}
                if (false || (false && false)) {
                    v18 := GenS1{F0: 255}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v19 float32 = ((0.0 * (3.14159 * float32((1e10 + 1e10)))) + -2.25)
                    _ = v18
                    _ = v19
                } else {
                    if (2147483647 >= -2147483648) {
                        var v20 int16 = -79
                        if (!(false && false)) {
                            if (((1000 > (-(v20))) && (-2.25 < 1.5)) || true) {
                                if true {
                                    if (546 < 749) {
                                        v16 = (v16 + (v16 % ((42) | 1)))
                                        v16 += v16
                                        if false {
                                            v20 = v20
                                        }
                                    }
                                } else {
                                    var v21 float64 = 1e10
                                    _ = v21
                                }
                                if true {
                                    v22 := GenS4{F0: 42}
                                    if ((false || false) || (true || (v2 > 65535))) {
                                        var v23 int = 42
                                        _ = v23
                                    } else {
                                        fmt.Println(v20)
                                        v24 := GenS3{F0: (v16 + v16), F1: -1, F2: v15}
                                        var v25 uint64 = 18446744073709551615
                                        _ = v24
                                        _ = v25
                                    }
                                    v26 := GenS5{F0: 675}
                                    _ = v22
                                    _ = v26
                                } else {
                                    v15 = "hello"
                                    if (42 < v2) {
                                        var v27 string = (v15 + "a")
                                        for fi28 := 0; fi28 < 20; fi28++ {
                                            v15 = v15
                                            _ = fi28
                                        }
                                        _ = v27
                                    } else {
                                        v1[(v15 + "hello")] = 147
                                        fmt.Println(v15)
                                    }
                                    var v29 uint = 0
                                    _ = v29
                                }
                                v15 = ("a" + "\n\t")
                            }
                            for fi30 := 0; fi30 < 15; fi30++ {
                                if false {
                                    var v31 int32 = 2147483647
                                    v2 += (^(v2))
                                    var v32 bool = (42 == 1)
                                    _ = v31
                                    _ = v32
                                } else {
                                    var v33 float64 = 1e10
                                    _ = v33
                                }
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                                fmt.Println(v16, v2, v20)
                                _ = fi30
                            }
                        } else {
                            if ((!(42 < 0)) || (true || false)) {
                                v2 -= (v2 | genFn1(v15))
                                for fi34 := 0; fi34 < 15; fi34++ {
                                    var v35 uint64 = 282
                                    switch (fi34 << ((0) & 63)) {
                                    case 259:
                                        switch "\n\t" {
                                        case "\n\t":
                                            v35 -= 18446744073709551615
                                        default:
                                            var v36 int = int(v35)
                                            _ = v36
                                        }
                                    }
                                    _ = fi34
                                    _ = v35
                                }
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v16 += (v16 << ((42) & 63))
                        }
                        _ = v20
                    }
                    if ((18446744073709551615 <= 483) && (-1 > -32768)) {
                        if false {
                            switch "hello" {
                            case "foo bar":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "hello":
                                v2 = 0
                            default:
                                if false {
                                    v16 = v16
                                    v15 = v15
                                }
                            }
                        } else {
                            for fi37 := 0; fi37 < 17; fi37++ {
                                v38 := GenS4{F0: 42}
                                _ = fi37
                                _ = v38
                            }
                            var v39 string = ("" + v15)
                            v40 := v1[v39]
                            _ = v39
                            _ = v40
                        }
                        for fi41 := 0; fi41 < 14; fi41++ {
                            fmt.Println(v2, v15, fi41)
                            _ = fi41
                        }
                    }
                }
                fmt.Println(v15)
                _ = v17
            }
            _ = v16
        } else {
            v15 = (v15 + (v15 + ("a" + (v15 + "foo bar"))))
            var v42 uint16 = v2
            _ = v42
        }
        v43 := GenS6{F0: 2, F1: 163}
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
