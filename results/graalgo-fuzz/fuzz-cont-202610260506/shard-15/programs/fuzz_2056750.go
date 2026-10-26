package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int }

func main() {
    v1 := []string{"a", ("\n\t" + (("\n\t" + "a") + "foo bar")), ("foo bar" + ("\n\t" + ((("foo bar" + (("foo bar" + "hello") + "hello")) + "a") + "\n\t"))), "foo bar"}
    v2 := []uint32{2147483647, 65535, 699}
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 uint32 = 357
        _ = fi4
        _ = v5
    }
    if ("hello" >= "\n\t") {
        switch 2147483647 {
        case 32767:
            var v6 float64 = 0.1
            _ = v6
        case -2147483648:
            var v7 int16 = 498
            _ = v7
        default:
            for fi8 := 0; fi8 < 18; fi8++ {
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 uint8 = 1
                    v11 := GenS1{F0: 2, F1: 0}
                    if ((511 >= 0) && (!(439 >= 32767))) {
                        v3["hello"] = 806
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                var v12 string = "foo bar"
                v3[(v12 + (v12 + ("a" + "")))] = 512
                _ = fi8
                _ = v12
            }
        }
        var v13 int = -46
        if false {
            var v14 bool = false
            v1 = append(v1, "foo bar")
            _ = v14
        }
        _ = v13
    } else {
        for fi15 := 0; fi15 < 5; fi15++ {
            if true {
                var v16 uint64 = 143
                switch "" {
                case "hello":
                    var v17 int8 = 100
                    _ = v17
                case "a":
                    v16 += (v16 / ((216) | 1))
                case "\n\t":
                    if true {
                        v16 = (v16 >> ((uint(v16)) & 63))
                        var v18 string = ""
                        _ = v18
                    }
                }
                _ = v16
            } else {
                if true {
                    _ = len(v2)
                } else {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        if true {
                            v2 = append(v2, 216)
                        }
                        var v20 int = fi19
                        _ = fi19
                        _ = v20
                    }
                    _ = len(v2)
                    var v21 uint32 = 2147483647
                    _ = v21
                }
            }
            if false {
                if false {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        switch (fi15 | fi22) {
                        case 702:
                            var v23 int32 = 2147483647
                            _ = v23
                        case 42:
                            fmt.Println(fi15, fi22)
                        default:
                            var v24 int = 519
                            _ = v24
                        }
                        _ = fi22
                    }
                    v25 := GenS1{F0: 65535, F1: 93}
                    for fi26 := 0; fi26 < 14; fi26++ {
                        switch (fi15 * fi15) {
                        case -2147483648:
                            var v27 string = ("\n\t" + "hello")
                            _ = v27
                        default:
                            if false {
                                var v28 int = (fi15 / ((-2147483648) | 1))
                                fmt.Println(v28, fi26, fi15)
                                var v29 uint64 = 2
                                _ = v28
                                _ = v29
                            } else {
                                v30 := GenS1{F0: 0, F1: 742}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v30
                            }
                        }
                        var v31 int32 = -1
                        _ = fi26
                        _ = v31
                    }
                    _ = v25
                }
                var v32 int = fi15
                v33 := v3[("\n\t" + "\n\t")]
                _ = v32
                _ = v33
            } else {
                if false {
                    if len(v1) > 0 {
                        v1[0] = (((("a" + "foo bar") + "hello") + "hello") + ("foo bar" + "foo bar"))
                    }
                    var v34 int = fi15
                    if false {
                        var v35 uint16 = 126
                        v35 -= 42
                        var v36 float32 = (3.14159 - 3.14159)
                        _ = v35
                        _ = v36
                    } else {
                        if true {
                            var v37 uint8 = 135
                            v37 = 128
                            _ = v37
                        }
                        var v38 uint64 = 650
                        fmt.Println(fi15, fi15)
                        _ = v38
                    }
                    _ = v34
                } else {
                    var v39 int8 = -120
                    if false {
                        if (42 != 127) {
                            v40 := GenS2{F0: (v39 + v39), F1: 1}
                            if (2147483647 == 1) {
                                switch "\n\t" {
                                case "hello":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            } else {
                                if false {
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                    var v41 uint32 = 65535
                                    if ((v39 - v39) < v39) {
                                        v42 := GenS2{F0: v39, F1: 245}
                                        var v43 float64 = ((float64((((1.5 + (0.0 + 1e10)) - 0.1) + 1e10)) + 0.0) * 3.14159)
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = v41
                                } else {
                                    delete(v3, (("\n\t" + "") + "foo bar"))
                                }
                                v44 := GenS3{F0: 0}
                                _ = v44
                            }
                            _ = v40
                        } else {
                            var v45 bool = (1 == 0)
                            v46 := GenS4{F0: 1}
                            _ = v45
                            _ = v46
                        }
                        v47 := GenS2{F0: v39, F1: 42}
                        for fi48 := 0; fi48 < 10; fi48++ {
                            v39 = v39
                            _ = fi48
                        }
                        _ = v47
                    }
                    var v49 uint64 = 0
                    _ = v39
                    _ = v49
                }
            }
            _ = fi15
        }
    }
    v3[("hello" + "\n\t")] = 65535
    var v50 int = 128
    v51 := GenS3{F0: 42}
    var v52 string = ((("a" + "foo bar") + "\n\t") + "hello")
    var v53 int16 = 772
    var v54 uint16 = 0
    v55 := GenS5{F0: (^(v53)), F1: int(v54)}
    var v56 uint64 = 1
    fmt.Println(v52, v53, v52)
    v53 += (v53 % ((v53) | 1))
    v57 := GenS4{F0: 665}
    fmt.Println(v53)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
}
