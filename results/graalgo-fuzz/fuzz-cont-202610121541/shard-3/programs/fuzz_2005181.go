package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint; F1 int8; F2 int8 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 int16; F1 uint32; F2 int16 }
func genFn0() string {
    v7 := []int{-787}
    var v8 float64 = 3.14159
    _ = v7
    _ = v8
    return "a"
}

func main() {
    v1 := []string{"hello", ("a" + ""), ("\n\t" + ("\n\t" + "a"))}
    var v2 int = 440
    var v3 uint64 = 1
    var v4 int8 = -128
    v5 := []string{("\n\t" + "a"), "\n\t", "", "a"}
    var v6 uint64 = (v3 ^ v3)
    if false {
        if true {
            fmt.Println(v4, v4, v3)
        } else {
            var v9 int64 = 299
            if (0 == 725) {
                if true {
                    v4 -= (v4 % ((-36) | 1))
                    if (!true) {
                        for fi10 := 0; fi10 < 19; fi10++ {
                            var v11 bool = false
                            v4 += (v4 / ((v4) | 1))
                            v11 = ((0.1 >= (0.1 * 3.14159)) || ((-2.25 > float32(0.0)) || (!v11)))
                            _ = fi10
                            _ = v11
                        }
                        if (false || (529 >= 255)) {
                            var v12 float64 = (0.1 + (((-2.25 - 1.5) - 1e10) / 0.1))
                            if ((3.14159 - 0.1) == 0.1) {
                                v6 += (^(v6))
                                v13 := GenS1{F0: (2147483648 >= 2147483647)}
                                switch (("foo bar" + "foo bar") + "foo bar") {
                                case "a":
                                    fmt.Println(v2)
                                }
                                _ = v13
                            } else {
                                if true {
                                    if true {
                                        v14 := GenS1{F0: (v3 < v3)}
                                        fmt.Println(v9)
                                        _ = v14
                                    }
                                } else {
                                    var v15 string = "hello"
                                    _ = v15
                                }
                                var v16 string = "a"
                                if (!((v9 + 9223372036854775807) != 1)) {
                                    var v17 float32 = 0.1
                                    _ = v17
                                }
                                _ = v16
                            }
                            if len(v5) > 0 {
                                v5[0] = "\n\t"
                            }
                            _ = v12
                        } else {
                            if len(v1) > 0 {
                                v1[0] = "a"
                            }
                        }
                    }
                    v5 = append(v5, "a")
                }
            }
            _ = v9
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (1000 > 32767) {
        fmt.Println(v3, v6, v2)
        if true {
            if true {
                v2 = (v2 ^ int(v4))
                v6 = (v6 + (v3 / ((683) | 1)))
                for fi18 := 0; fi18 < 20; fi18++ {
                    var v19 uint = 2
                    _ = fi18
                    _ = v19
                }
            }
        }
        fmt.Println(v4)
    } else {
        fmt.Println(v2, v3)
        if (int(v3) < (v2 * (v2 << ((42) & 63)))) {
            v20 := GenS1{F0: false}
            v2 += (v2 * v2)
            for fi21 := 0; fi21 < 13; fi21++ {
                v4 += v4
                _ = fi21
            }
            _ = v20
        } else {
            var v22 uint16 = 1
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS1{F0: true}
                _ = fi23
                _ = v24
            }
            v25 := GenS1{F0: (127 != int8(v22))}
            _ = v22
            _ = v25
        }
    }
    var v26 string = "hello"
    if true {
        switch v26 {
        case "a":
            fmt.Println(v3, v3, v4)
        case "":
            switch v26 {
            case "hello":
                v4 = (v4 + (-(v4)))
            default:
                v6 -= 1
            }
        }
    } else {
        if (9223372036854775807 == 1) {
            v6 = (^(v6))
        }
        if (1.5 < 3.14159) {
            v2 += (^(v2))
            v27 := GenS1{F0: false}
            v28 := GenS2{F0: "", F1: 9223372036854775807}
            _ = v27
            _ = v28
        }
        v29 := GenS3{F0: 0, F1: v4, F2: -77}
        _ = v29
    }
    v30 := GenS2{F0: (v26 + (v26 + "foo bar")), F1: 0}
    v31 := GenS4{F0: (!((v2 + v2) == 176)), F1: -9223372036854775808}
    var v32 uint32 = 0
    var v33 string = (v26 + (v26 + (v26 + ("" + v26))))
    for fi34 := 0; fi34 < 19; fi34++ {
        v3 = (v6 * 1)
        if ((v33 + v33) <= "hello") {
            var v35 int64 = int64(v32)
            _ = v35
        } else {
            v3 = v3
        }
        _ = fi34
    }
    if (65535 == 4294967295) {
        v36 := GenS5{F0: int64(v4), F1: v2}
        _ = v36
    } else {
        _ = len(v1)
        var v37 int8 = -1
        _ = v37
    }
    var v38 uint16 = 0
    var v39 int32 = 97
    v40 := GenS6{F0: 877, F1: (v32 * uint32(v6)), F2: 680}
    switch v33 {
    case "hello":
        if (42 > 42) {
            v4 += 1
            for fi41 := 0; fi41 < 1; fi41++ {
                v3 = 18446744073709551615
                var v42 uint32 = 0
                _ = fi41
                _ = v42
            }
        }
    case "\n\t":
        v26 = ("hello" + v33)
    case "foo bar":
        _ = len(v5)
    }
    var v43 string = v26
    fmt.Println(v4)
    v44 := GenS5{F0: 601, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
