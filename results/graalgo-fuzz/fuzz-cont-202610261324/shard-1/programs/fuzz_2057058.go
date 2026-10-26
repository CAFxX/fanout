package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 string; F1 uint; F2 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32; F1 int8 }
type GenS9 struct { F0 int8 }

func main() {
    v1 := []string{"", (("\n\t" + ("a" + "\n\t")) + "a"), "hello", ("foo bar" + "foo bar")}
    var v2 uint8 = 92
    var v3 uint64 = 18446744073709551615
    var v4 bool = true
    v2 += (v2 & (v2 << ((18446744073709551615) & 63)))
    v5 := GenS1{F0: -573, F1: 0}
    fmt.Println(v4, v3, v3)
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int = (fi6 + 127)
        if (!false) {
            v8 := GenS1{F0: fi6, F1: 2}
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS2{F0: 42, F1: 695}
                switch 4294967295 {
                case 32767:
                    v11 := GenS3{F0: "", F1: 0, F2: fi6}
                    _ = v11
                case 1:
                    if (!false) {
                        if v4 {
                            _ = len(v1)
                            v12 := GenS4{F0: 0}
                            v13 := GenS2{F0: 2, F1: 42}
                            _ = v12
                            _ = v13
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if (("hello" >= ("" + "hello")) && v4) {
                                v2 -= (^(v2))
                            }
                            if (v4 || v4) {
                                var v14 float32 = (float32(-2.25) * 0.0)
                                if ((!(128 < fi6)) && v4) {
                                    v2 -= v2
                                }
                                _ = v14
                            }
                        }
                        fmt.Println(fi9, fi9)
                        fmt.Println(v3, v3)
                    }
                    var v15 float32 = 0.0
                    if v4 {
                        var v16 uint32 = 563
                        switch (fi9 & v7) {
                        case 683:
                            if true {
                                v17 := GenS5{F0: (2 <= 0)}
                                v18 := GenS6{F0: 42, F1: 42}
                                var v19 uint = 42
                                switch (v7 - -300) {
                                case -133:
                                    v16 = (^(v16))
                                case 0:
                                    v20 := GenS4{F0: 0}
                                    _ = v20
                                default:
                                    var v21 int32 = 0
                                    _ = v21
                                }
                                switch v7 {
                                case 128:
                                    v1 = append(v1, ("\n\t" + "hello"))
                                }
                                v1 = append(v1, ((("a" + "foo bar") + "hello") + "foo bar"))
                                v22 := GenS7{F0: (fi6 >> (((v19 * 65535)) & 63))}
                                switch int(v16) {
                                case -204:
                                    v23 := GenS8{F0: (v16 + v16), F1: 127}
                                    _ = v23
                                default:
                                    v19 -= 0
                                }
                                switch fi6 {
                                case 474:
                                    fmt.Println(v2, v4)
                                }
                                switch "\n\t" {
                                case "":
                                    var v24 uint = (v19 | (^(v19)))
                                    _ = v24
                                }
                                var v25 int = fi6
                                _ = v17
                                _ = v18
                                _ = v19
                                _ = v22
                                _ = v25
                            }
                        case 116:
                            if false {
                                v15 -= v15
                                v26 := GenS9{F0: 0}
                                if (v3 > uint64(v2)) {
                                    var v27 float32 = -2.25
                                    switch "" {
                                    case "hello":
                                        v27 -= 1.5
                                    }
                                    var v28 uint32 = 435
                                    _ = v27
                                    _ = v28
                                } else {
                                    v1 = append(v1, "hello")
                                    if ((1e10 >= ((float64((v15 + (v15 * (v15 * v15)))) * -2.25) - 0.1)) && v4) {
                                        fmt.Println(v4)
                                    }
                                    var v29 uint16 = 1
                                    _ = v29
                                }
                                var v30 uint = 150
                                v3 = (v3 % (((v3 % ((1) | 1))) | 1))
                                _ = v26
                                _ = v30
                            }
                        }
                        _ = v16
                    }
                    _ = v15
                case 32768:
                    if (-37 != 0) {
                        if (931 >= 1) {
                            var v31 uint64 = (^(v3))
                            _ = v31
                        }
                    }
                }
                _ = fi9
                _ = v10
            }
            _ = v8
        }
        v4 = v4
        _ = fi6
        _ = v7
    }
    switch 693 {
    case 534:
        fmt.Println(v4, v4)
    }
    fmt.Println(v3)
    if false {
        switch 746 {
        case 4294967295:
            v1 = append(v1, "foo bar")
        }
    } else {
        v32 := GenS3{F0: "", F1: 18446744073709551615, F2: int(v3)}
        _ = v32
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 int = fi33
        var v35 int = fi33
        if len(v1) > 0 {
            v1[0] = ""
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    v3 = (v3 << ((4294967295) & 63))
    fmt.Println(v4, v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
