package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint8; F1 uint16; F2 bool }
func genFn0() int32 {
    switch 32768 {
    case 32768:
        if (42 < 1) {
            v6 := GenS1{F0: 32767, F1: 1}
            var v7 float64 = (0.0 * 3.14159)
            _ = v6
            _ = v7
        } else {
            v8 := []int{0, 127}
            var v9 uint32 = 2147483647
            _ = v8
            _ = v9
        }
    }
    var v10 string = ""
    if (true && true) {
        v11 := GenS2{F0: 0}
        v10 = "hello"
        var v12 int = 127
        _ = v11
        _ = v12
    }
    _ = v10
    return -2147483648
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint8 = 1
    v3 := []string{("a" + ""), "hello", (("hello" + (("hello" + "") + "foo bar")) + "\n\t")}
    var v4 float64 = -2.25
    var v5 float32 = float32(3.14159)
    v1 = (v1 - (v1 % (((^(v1))) | 1)))
    var v13 bool = (!(!(!(-48 != -1))))
    v3 = append(v3, "")
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v13 = (!(((("hello" + "\n\t") + "\n\t") + ("hello" + "")) > "\n\t"))
    var v14 float32 = (v5 + (v5 - (v5 / 1.0)))
    v13 = v13
    switch (("" + "a") + "") {
    case "foo bar":
        var v15 float32 = (0.0 + 1e10)
        _ = v15
    case "\n\t":
        v3 = append(v3, "hello")
    case "hello":
        for fi16 := 0; fi16 < 12; fi16++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if v13 {
                v17 := GenS1{F0: 65535, F1: v2}
                var v18 int32 = v1
                _ = v17
                _ = v18
            } else {
                switch fi16 {
                case -2147483648:
                    switch -738 {
                    case -1:
                        var v19 uint8 = (v2 << ((18446744073709551615) & 63))
                        _ = v19
                    default:
                        for fi20 := 0; fi20 < 4; fi20++ {
                            var v21 int = fi20
                            _ = fi20
                            _ = v21
                        }
                    }
                case 127:
                    switch ((("\n\t" + ("foo bar" + ("hello" + ""))) + "") + "foo bar") {
                    case "foo bar":
                        if (v4 == v4) {
                            if len(v3) > 0 {
                                v3[0] = (("hello" + "foo bar") + "hello")
                            }
                            var v22 float32 = v14
                            v5 = v14
                            _ = v22
                        } else {
                            var v23 int8 = 79
                            fmt.Println(v2, v4)
                            _ = v23
                        }
                    case "":
                        for fi24 := 0; fi24 < 18; fi24++ {
                            v25 := GenS2{F0: 1}
                            var v26 float32 = ((v14 / v14) / v14)
                            if (!(v1 != (^(v1)))) {
                                v14 += v14
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                                switch (fi24 ^ fi24) {
                                case -783:
                                    var v27 int16 = -114
                                    _ = v27
                                case -626:
                                    v28 := GenS1{F0: 1, F1: (v2 << ((1) & 63))}
                                    _ = v28
                                default:
                                    switch "\n\t" {
                                    case "":
                                        var v29 uint64 = 42
                                        _ = v29
                                    case "a":
                                        v30 := GenS1{F0: 518, F1: (^(v2))}
                                        _ = v30
                                    }
                                }
                            } else {
                                if v13 {
                                    switch ("" + ("hello" + ((("foo bar" + "foo bar") + (("" + "\n\t") + "\n\t")) + "foo bar"))) {
                                    case "\n\t":
                                        v13 = (v4 == v4)
                                    case "a":
                                        var v31 uint = 704
                                        _ = v31
                                    case "":
                                        v32 := GenS3{F0: (((^(v2)) != 143) || (v4 != float64(float32((((1.5 + (v4 * (3.14159 + ((v4 + 1e10) - v4)))) - 3.14159) / (v4 * 0.1)))))), F1: 100, F2: 0}
                                        _ = v32
                                    default:
                                        fmt.Println(v2, v1, v2)
                                    }
                                }
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                            }
                            v1 -= genFn0()
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                }
                fmt.Println(v5, v4)
                var v33 uint16 = 895
                _ = v33
            }
            if (v14 < v5) {
                var v34 bool = false
                if (2 <= 4294967295) {
                    v35 := GenS2{F0: 2}
                    v36 := GenS4{F0: v2, F1: 89, F2: true}
                    _ = v35
                    _ = v36
                } else {
                    var v37 int16 = 1000
                    _ = v37
                }
                v2 += v2
                _ = v34
            } else {
                for fi38 := 0; fi38 < 10; fi38++ {
                    _ = len(v3)
                    _ = fi38
                }
                if len(v3) > 0 {
                    v3[0] = (((("hello" + ("" + "hello")) + "a") + "foo bar") + "hello")
                }
            }
            _ = fi16
        }
    default:
        fmt.Println(v13)
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    if v13 {
        for fi39 := 0; fi39 < 14; fi39++ {
            if (42 < 1) {
                var v40 int16 = -1
                v41 := GenS3{F0: ((v2 / ((v2) | 1)) > (v2 << ((1) & 63))), F1: 127, F2: 4294967295}
                _ = v40
                _ = v41
            } else {
                var v42 int32 = v1
                v14 = (((v14 * (v5 + 0.1)) * v14) - v5)
                fmt.Println(fi39, fi39)
                _ = v42
            }
            v1 += (v1 + (^(v1)))
            _ = fi39
        }
    } else {
        for fi43 := 0; fi43 < 5; fi43++ {
            v44 := GenS1{F0: 65535, F1: v2}
            _ = fi43
            _ = v44
        }
        v14 = v14
    }
    v2 -= (^(v2))
    fmt.Println(v13, v14, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
