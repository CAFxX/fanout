package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 string }
type GenS3 struct { F0 bool; F1 int8; F2 uint }
type GenS4 struct { F0 int; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint8 = 0
    var v2 string = (("hello" + "foo bar") + ("" + (("" + "a") + ((((("foo bar" + "\n\t") + "a") + ("" + ("hello" + "hello"))) + "hello") + ("foo bar" + "")))))
    v3 := []string{"foo bar", "\n\t"}
    v4 := make(map[string]int)
    var v5 string = (v2 + v2)
    if (-2.25 >= (float64(-2.25) / -2.25)) {
        var v6 uint64 = 4294967295
        _ = v6
    }
    v5 = v5
    var v7 float32 = 3.14159
    var v8 int64 = -1
    for fi9 := 0; fi9 < 17; fi9++ {
        v1 = 133
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v10 := GenS1{F0: -338}
            _ = v10
        } else {
            switch (fi9 ^ fi9) {
            case 9223372036854775807:
                v11 := v4[v2]
                _ = v11
            case -2147483648:
                v7 -= 3.14159
            default:
                switch (v5 + v2) {
                case "hello":
                    v12 := GenS1{F0: 420}
                    _ = v12
                case "":
                    if ((false && false) || false) {
                        var v13 int8 = 89
                        var v14 bool = (false || true)
                        v13 = (v13 + v13)
                        _ = v13
                        _ = v14
                    }
                case "foo bar":
                    v5 = v5
                default:
                    v7 = (1.5 * (((v7 + ((v7 * v7) + (float32(3.14159) * (0.0 + 1e10)))) - (v7 / (v7 / 1e10))) + (v7 / v7)))
                }
            }
            var v15 uint8 = v1
            if (!false) {
                v16 := GenS2{F0: (v8 * 348), F1: 858, F2: v5}
                var v17 bool = ((((v8 * 42) > 1) || true) && false)
                _ = v16
                _ = v17
            } else {
                var v18 uint64 = 2
                var v19 int = (fi9 ^ fi9)
                if false {
                    v7 += (v7 + (v7 / v7))
                } else {
                    if len(v3) > 0 {
                        v3[0] = (v2 + (v2 + v2))
                    }
                }
                _ = v18
                _ = v19
            }
            _ = v15
        }
        for fi20 := 0; fi20 < 10; fi20++ {
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v21 int = fi20
                if (42 == 1) {
                    v22 := GenS2{F0: (v8 / ((42) | 1)), F1: 65535, F2: v5}
                    var v23 float32 = v7
                    _ = v22
                    _ = v23
                }
                _ = v21
            } else {
                var v24 int = (fi20 * 2147483647)
                _ = len(v4)
                var v25 float64 = 0.1
                _ = v24
                _ = v25
            }
            _ = fi20
        }
        _ = fi9
    }
    var v26 string = "\n\t"
    _ = len(v3)
    v5 = "hello"
    v27 := GenS3{F0: true, F1: -89, F2: 834}
    delete(v4, ("a" + ((v26 + v2) + v5)))
    switch -9223372036854775808 {
    case -2147483648:
        if (-1 < 0) {
            v28 := GenS1{F0: 0}
            if true {
                if (0 >= 618) {
                    switch v26 {
                    case "\n\t":
                        v29 := v4["foo bar"]
                        _ = v29
                    default:
                        v26 = v2
                    }
                } else {
                    var v30 uint = 255
                    fmt.Println(v26)
                    _ = v30
                }
                if (32767 == 32767) {
                    v31 := GenS4{F0: 567, F1: 482, F2: 32767}
                    if ((!(65535 != 65535)) || (false || true)) {
                        if true {
                            for fi32 := 0; fi32 < 9; fi32++ {
                                if true {
                                    _ = len(v3)
                                }
                                _ = fi32
                            }
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        v33 := GenS4{F0: 0, F1: 65535, F2: 32767}
                        for fi34 := 0; fi34 < 14; fi34++ {
                            switch v2 {
                            case "foo bar":
                                v2 = "foo bar"
                            case "a":
                                for fi35 := 0; fi35 < 17; fi35++ {
                                    var v36 float64 = 1.5
                                    fmt.Println(v36)
                                    _ = fi35
                                    _ = v36
                                }
                            }
                            if true {
                                v37 := GenS3{F0: false, F1: 100, F2: 65535}
                                _ = v37
                            } else {
                                switch "\n\t" {
                                case "foo bar":
                                    v8 = v8
                                case "":
                                    switch v5 {
                                    case "a":
                                        v26 = (v5 + "foo bar")
                                    }
                                case "a":
                                    v2 = v2
                                default:
                                    var v38 uint16 = uint16(v1)
                                    _ = v38
                                }
                                v39 := GenS5{F0: 95}
                                var v40 bool = true
                                _ = v39
                                _ = v40
                            }
                            v26 = (v26 + (v5 + ("foo bar" + "a")))
                            _ = fi34
                        }
                        _ = v33
                    }
                    fmt.Println(v7, v26, v8)
                    _ = v31
                }
            } else {
                var v41 int = 127
                v3 = append(v3, v26)
                var v42 bool = ((true || (int(v8) <= 169)) && (false && false))
                _ = v41
                _ = v42
            }
            v7 = (v7 + (float32(3.14159) + v7))
            _ = v28
        }
    case 706:
        switch (((v5 + v5) + (((v5 + (v26 + v2)) + v5) + v26)) + v2) {
        case "foo bar":
            _ = len(v4)
        }
    }
    var v43 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v7
    _ = v8
}
