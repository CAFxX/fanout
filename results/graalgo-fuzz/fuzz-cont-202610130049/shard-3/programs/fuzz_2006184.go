package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 130
    v3 := []string{"hello", ("a" + "hello"), "hello", "hello"}
    var v4 uint16 = 65535
    v5 := make(map[int]int)
    if true {
        fmt.Println(v2, v4)
        fmt.Println(v4, v2)
        var v6 uint16 = (^(v4))
        _ = v6
    } else {
        v3 = append(v3, ("" + "\n\t"))
        switch ("foo bar" + "a") {
        case "":
            var v7 int64 = 1
            _ = v7
        case "foo bar":
            var v8 int16 = -358
            _ = v8
        }
    }
    var v9 int8 = 42
    if true {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        var v10 uint32 = 1
        for fi11 := 0; fi11 < 14; fi11++ {
            v4 = v4
            _ = fi11
        }
        _ = v10
    }
    var v12 uint32 = 2147483648
    var v13 uint = (v2 + v2)
    switch int(v9) {
    case 0:
        v9 += (v9 | (v9 % ((v9) | 1)))
    case 226:
        var v14 int64 = -1
        _ = v14
    }
    if true {
        var v15 float64 = 3.14159
        if (v9 >= 1) {
            var v16 uint64 = uint64(v2)
            if false {
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 int32 = 826
                    _ = fi17
                    _ = v18
                }
                _ = len(v5)
            } else {
                if true {
                    delete(v1, v12)
                    v15 -= ((v15 - 1e10) + v15)
                    switch (("" + (("" + "") + "\n\t")) + "hello") {
                    case "a":
                        v19 := GenS1{F0: (v12 + (v12 << ((v2) & 63))), F1: 128}
                        _ = v19
                    case "":
                        v20 := GenS1{F0: v12, F1: 132}
                        _ = v20
                    default:
                        for fi21 := 0; fi21 < 7; fi21++ {
                            if true {
                                if (!(!true)) {
                                    fmt.Println(v9, v12, v15)
                                    if false {
                                        for fi22 := 0; fi22 < 4; fi22++ {
                                            v23 := GenS1{F0: 0, F1: 128}
                                            _ = fi22
                                            _ = v23
                                        }
                                    } else {
                                        v4 += 255
                                    }
                                    v12 = 4294967295
                                }
                                var v24 uint64 = v16
                                _ = v24
                            } else {
                                fmt.Println(v2)
                                var v25 uint8 = 78
                                _ = v25
                            }
                            _ = fi21
                        }
                    }
                } else {
                    fmt.Println(v2)
                    if (!((((!true) && false) || (1 < v9)) || false)) {
                        v13 = v2
                        var v26 float32 = 0.1
                        _ = v26
                    }
                    v3 = append(v3, "foo bar")
                }
                _ = len(v1)
            }
            v16 = 0
            _ = v16
        }
        var v27 string = "hello"
        _ = v15
        _ = v27
    }
    if ((("hello" + "a") + "a") == "") {
        v4 += (^(v4))
        var v28 string = "hello"
        for fi29 := 0; fi29 < 6; fi29++ {
            if true {
                var v30 uint32 = 759
                if (!((-706 < 0) && (!(((float64(-2.25) / 1.0) * (1e10 - (1.5 * (0.1 * 1.5)))) >= float64((-2.25 / -2.25)))))) {
                    if (false && true) {
                        for fi31 := 0; fi31 < 12; fi31++ {
                            if (v9 >= (v9 % ((100) | 1))) {
                                if false {
                                    v9 = 100
                                    var v32 uint32 = uint32(v4)
                                    v4 += 42
                                    var v33 float64 = -2.25
                                    _ = v32
                                    _ = v33
                                } else {
                                    fmt.Println(v9, v2)
                                }
                                var v34 uint = v2
                                _ = v34
                            } else {
                                if false {
                                    v9 -= (-(v9))
                                    var v35 int16 = 750
                                    v36 := GenS1{F0: uint32(v4), F1: 1}
                                    _ = v35
                                    _ = v36
                                } else {
                                    v12 = v30
                                    if false {
                                        v37 := GenS2{F0: true}
                                        _ = v37
                                    } else {
                                        switch (-(fi29)) {
                                        case 9223372036854775807:
                                            v3 = append(v3, v28)
                                        case 32767:
                                            v38 := GenS3{F0: (v30 << ((1) & 63))}
                                            _ = v38
                                        case 127:
                                            v9 += (v9 >> ((0) & 63))
                                        default:
                                            v13 += 328
                                        }
                                    }
                                    if (v4 < v4) {
                                        if (!true) {
                                            fmt.Println(v4, v30, v30)
                                            _ = len(v3)
                                            fmt.Println(fi29)
                                        }
                                        v12 = 670
                                        v30 -= (v30 + v12)
                                    } else {
                                        v39 := GenS1{F0: v30, F1: 127}
                                        var v40 float32 = 0.0
                                        _ = v39
                                        _ = v40
                                    }
                                }
                                v41 := GenS3{F0: v30}
                                v4 = v4
                                _ = v41
                            }
                            _ = fi31
                        }
                        v30 -= (v30 - (v30 - 0))
                        v30 = (v12 >> (((v2 | (v2 | 42))) & 63))
                    }
                    switch (fi29 / (((fi29 % (((fi29 | (fi29 & -1))) | 1))) | 1)) {
                    case -1:
                        v30 = (v30 ^ (v30 % (((v12 + v12)) | 1)))
                    case 128:
                        var v42 uint16 = v4
                        _ = v42
                    }
                    v43 := GenS4{F0: (v13 << ((0) & 63)), F1: 2}
                    _ = v43
                } else {
                    if true {
                        v1[v12] = (^(v30))
                        var v44 bool = (!("hello" <= "foo bar"))
                        v30 += 2147483648
                        _ = v44
                    }
                    fmt.Println(fi29, v12)
                    v12 += (v30 - v30)
                }
                for fi45 := 0; fi45 < 2; fi45++ {
                    switch 2147483647 {
                    case 127:
                        v46 := GenS2{F0: (((v30 & v12) == v30) || false)}
                        _ = v46
                    default:
                        v9 = v9
                    }
                    _ = fi45
                }
                _ = v30
            }
            v47 := GenS5{F0: 0}
            var v48 int = (-(fi29))
            _ = fi29
            _ = v47
            _ = v48
        }
        _ = v28
    } else {
        var v49 float32 = 3.14159
        _ = v49
    }
    v2 = (v2 | v13)
    var v50 string = "foo bar"
    v9 = v9
    switch int(v12) {
    case 32767:
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    case 9223372036854775807:
        switch v50 {
        case "foo bar":
            var v51 uint32 = (^(v12))
            _ = v51
        case "hello":
            fmt.Println(v9)
        default:
            v52 := v5[127]
            _ = v52
        }
    }
    _ = len(v5)
    v9 = int8(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v9
}
