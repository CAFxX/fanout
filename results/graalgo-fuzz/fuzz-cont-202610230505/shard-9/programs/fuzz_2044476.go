package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint8 }

func main() {
    v1 := []uint32{42}
    v2 := []string{"\n\t"}
    v3 := make(map[int]uint32)
    v4 := make(map[int]bool)
    var v5 int16 = -633
    v5 += (v5 ^ v5)
    v6 := GenS1{F0: -1, F1: (v5 >> ((255) & 63)), F2: 128}
    for fi7 := 0; fi7 < 2; fi7++ {
        if true {
            var v8 int32 = 629
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v8
        } else {
            v5 = v5
        }
        v3[32768] = 42
        var v9 int8 = -13
        _ = fi7
        _ = v9
    }
    v10 := GenS1{F0: -1, F1: -1, F2: 123}
    var v11 bool = true
    v5 += (v5 ^ 1000)
    var v12 uint16 = 312
    for fi13 := 0; fi13 < 16; fi13++ {
        switch "hello" {
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            if true {
                if false {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        fmt.Println(fi14)
                        fmt.Println(v12, v12, fi13)
                        _ = fi14
                    }
                    for fi15 := 0; fi15 < 17; fi15++ {
                        switch fi15 {
                        case 128:
                            v2 = append(v2, "foo bar")
                            var v16 int = fi15
                            if (42 < 0) {
                                v4[524] = (v11 || (42 == 42))
                                v17 := GenS1{F0: 100, F1: v5, F2: 42}
                                _ = v17
                            }
                            v16 -= (fi13 / ((v16) | 1))
                            _ = v16
                        case 4294967295:
                            if v11 {
                                if v11 {
                                    var v18 float64 = 3.14159
                                    var v19 string = "foo bar"
                                    var v20 int16 = (v5 ^ v5)
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                } else {
                                    fmt.Println(fi15)
                                    var v21 bool = ((("foo bar" + "hello") + "\n\t") >= "a")
                                    v21 = v21
                                    switch "" {
                                    case "foo bar":
                                        switch (("foo bar" + "a") + "foo bar") {
                                        case "hello":
                                            var v22 bool = (false || ((0.0 / (1e10 - 1.5)) > (1e10 - ((3.14159 * 3.14159) * (1e10 - 0.1)))))
                                            _ = v22
                                        case "":
                                            v11 = (!(846 <= 42))
                                        case "foo bar":
                                            if (255 != 187) {
                                                var v23 int8 = 119
                                                v21 = (v11 && true)
                                                _ = v23
                                            } else {
                                                var v24 float64 = 0.1
                                                v25 := v3[128]
                                                if ((!((!false) || true)) || v21) {
                                                    var v26 bool = (7 >= 133)
                                                    _ = v26
                                                }
                                                _ = v24
                                                _ = v25
                                            }
                                        default:
                                            v12 = 65535
                                        }
                                    case "a":
                                        v21 = v11
                                    case "\n\t":
                                        var v27 float64 = 0.0
                                        _ = v27
                                    }
                                    v2 = append(v2, "a")
                                    var v28 float32 = 1e10
                                    v29 := GenS1{F0: -1, F1: 84, F2: 255}
                                    _ = v21
                                    _ = v28
                                    _ = v29
                                }
                            }
                        }
                        v11 = v11
                        _ = fi15
                    }
                } else {
                    if (!(127 < 208)) {
                        if ((!true) && (("foo bar" + "") < "hello")) {
                            if (v11 && false) {
                                var v30 uint = 42
                                _ = v30
                            } else {
                                v31 := v4[(fi13 >> ((18446744073709551615) & 63))]
                                var v32 uint64 = 1
                                var v33 int32 = 944
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                        }
                        v34 := v3[(fi13 / ((128) | 1))]
                        _ = v34
                    } else {
                        if true {
                            if ("" != ("hello" + ("a" + ("" + "")))) {
                                var v35 string = ("foo bar" + "hello")
                                _ = v35
                            } else {
                                for fi36 := 0; fi36 < 17; fi36++ {
                                    v37 := GenS1{F0: 127, F1: (v5 << ((0) & 63)), F2: 255}
                                    v2 = append(v2, "hello")
                                    _ = fi36
                                    _ = v37
                                }
                            }
                            var v38 float32 = 0.1
                            _ = v38
                        }
                    }
                    switch 32768 {
                    case 32767:
                        v11 = v11
                    default:
                        for fi39 := 0; fi39 < 19; fi39++ {
                            var v40 string = "foo bar"
                            fmt.Println(v40, v40, v5)
                            _ = fi39
                            _ = v40
                        }
                    }
                }
                v12 -= (v12 | (v12 + 2))
            } else {
                var v41 uint64 = 42
                fmt.Println(v5)
                _ = v41
            }
        }
        _ = fi13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
