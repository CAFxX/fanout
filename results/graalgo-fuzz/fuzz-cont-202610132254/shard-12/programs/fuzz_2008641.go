package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 bool }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint; F1 int8; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint }

func main() {
    v1 := make(map[int]string)
    v2 := []string{("foo bar" + "")}
    v3 := make(map[string]string)
    var v4 int8 = 1
    v2 = append(v2, ("foo bar" + ""))
    if false {
        if (!(false || true)) {
            for fi5 := 0; fi5 < 11; fi5++ {
                fmt.Println(v4)
                _ = fi5
            }
            for fi6 := 0; fi6 < 15; fi6++ {
                fmt.Println(v4)
                for fi7 := 0; fi7 < 13; fi7++ {
                    v4 += (v4 ^ 100)
                    v4 -= v4
                    v2 = append(v2, (("foo bar" + "\n\t") + "\n\t"))
                    _ = fi7
                }
                fmt.Println(v4)
                _ = fi6
            }
            v4 = (^(v4))
        } else {
            fmt.Println(v4)
        }
        _ = len(v2)
        fmt.Println(v4)
    }
    v8 := v1[127]
    if false {
        if (0 >= 127) {
            if (0 > 1) {
                fmt.Println(v4)
                v4 -= (-(v4))
                v9 := GenS1{F0: 42, F1: 9223372036854775807, F2: (1e10 < float64(0.1))}
                _ = v9
            } else {
                var v10 float32 = (1e10 * -2.25)
                v10 -= (((float32(float64((v10 / v10))) / (float32((-2.25 * 3.14159)) * v10)) + 1e10) / (v10 / (-2.25 * 3.14159)))
                _ = v10
            }
            v4 += 127
            var v11 uint = 1
            _ = v11
        } else {
            var v12 uint64 = 42
            _ = v12
        }
    } else {
        if len(v2) > 0 {
            v2[0] = (v8 + v8)
        }
        var v13 string = v8
        _ = v13
    }
    if (!false) {
        v8 = "a"
        if false {
            var v14 float32 = float32(float64((float32((1.5 * (0.0 + float64((0.1 + 1.5))))) - -2.25)))
            _ = v14
        } else {
            v15 := GenS1{F0: -1, F1: 1, F2: ((0.0 + 0.1) < ((1e10 * 1e10) / -2.25))}
            _ = v15
        }
        v16 := GenS2{F0: 1, F1: 982}
        _ = v16
    }
    v2 = append(v2, (("\n\t" + "hello") + (v8 + "foo bar")))
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int64 = 1
        v19 := GenS1{F0: (fi17 % ((822) | 1)), F1: fi17, F2: true}
        _ = fi17
        _ = v18
        _ = v19
    }
    if true {
        v20 := GenS1{F0: int(v4), F1: 32768, F2: true}
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 uint64 = 18446744073709551615
            _ = fi21
            _ = v22
        }
        switch "" {
        case "hello":
            switch 729 {
            case 0:
                switch v8 {
                case "hello":
                    v23 := GenS3{F0: 42, F1: 1, F2: -1}
                    _ = v23
                case "foo bar":
                    var v24 uint8 = 207
                    _ = v24
                case "":
                    if false {
                        if true {
                            v25 := GenS1{F0: 818, F1: 42, F2: false}
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 10; fi26++ {
                                var v27 bool = (false && true)
                                v1[fi26] = "\n\t"
                                fmt.Println(v4)
                                _ = fi26
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 15; fi28++ {
                                for fi29 := 0; fi29 < 10; fi29++ {
                                    switch ("a" + "\n\t") {
                                    case "\n\t":
                                        if true {
                                            if false {
                                                v30 := v3[v8]
                                                v30 = ((v30 + "a") + "hello")
                                                var v31 bool = true
                                                _ = v30
                                                _ = v31
                                            }
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    }
                                    var v32 int16 = 310
                                    _ = fi29
                                    _ = v32
                                }
                                _ = fi28
                            }
                            v33 := GenS1{F0: 42, F1: -2147483648, F2: ((true && true) || (-798 != 618))}
                            _ = v33
                        }
                        _ = len(v2)
                        fmt.Println(v4, v4)
                    } else {
                        v4 -= v4
                    }
                default:
                    var v34 float32 = 3.14159
                    _ = v34
                }
            case -9223372036854775808:
                v2 = append(v2, (v8 + v8))
            }
        default:
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        _ = v20
    }
    fmt.Println(v4)
    if (3.14159 == 0.0) {
        v2 = append(v2, (((("a" + (v8 + v8)) + v8) + ("" + (v8 + v8))) + (((v8 + v8) + ("" + v8)) + "\n\t")))
        for fi35 := 0; fi35 < 5; fi35++ {
            v36 := GenS3{F0: 2, F1: 101, F2: 1}
            var v37 uint64 = 1
            v38 := GenS4{F0: 32767, F1: 46}
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
