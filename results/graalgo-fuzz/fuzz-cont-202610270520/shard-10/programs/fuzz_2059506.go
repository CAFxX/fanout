package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int32; F1 string; F2 int16 }
type GenS5 struct { F0 int16; F1 string; F2 uint8 }
type GenS6 struct { F0 string }

func main() {
    v1 := []uint32{2, 2147483647, 4294967295}
    v2 := make(map[string]string)
    var v3 float64 = float64(1.5)
    var v4 bool = ((!(!true)) && true)
    var v5 float64 = v3
    v6 := []uint32{2147483647, 762}
    v7 := GenS1{F0: 809, F1: 128, F2: 330}
    v8 := GenS1{F0: 2, F1: 127, F2: 605}
    if true {
        switch "a" {
        case "":
            var v9 uint64 = 18446744073709551615
            _ = v9
        case "hello":
            delete(v2, ("\n\t" + "hello"))
        default:
            for fi10 := 0; fi10 < 5; fi10++ {
                v4 = (v4 || (-12 < -9223372036854775808))
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 int = 9223372036854775807
                    var v13 uint16 = 1
                    switch (("foo bar" + "foo bar") + ("hello" + "a")) {
                    case "hello":
                        fmt.Println(v3)
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                        v14 := GenS2{F0: 0, F1: 65535}
                        _ = v14
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v3 = ((v3 + float64((1e10 + 0.0))) * ((float64((0.0 + 1.5)) + v3) / v3))
                _ = fi10
            }
        }
    }
    var v15 int16 = 32767
    switch 4294967295 {
    case 32767:
        v16 := GenS1{F0: 707, F1: 128, F2: -1}
        _ = v16
    }
    v17 := v2[("" + ("a" + ""))]
    fmt.Println(v17, v15)
    var v18 float64 = v3
    v3 = 0.1
    var v19 uint16 = 2
    if (!v4) {
        fmt.Println(v4, v4, v3)
        for fi20 := 0; fi20 < 10; fi20++ {
            for fi21 := 0; fi21 < 11; fi21++ {
                var v22 int32 = 124
                if v4 {
                    delete(v2, v17)
                } else {
                    switch (v17 + ("\n\t" + "hello")) {
                    case "":
                        var v23 uint8 = 127
                        _ = v23
                    case "hello":
                        v3 += (v3 - ((v18 / 1.5) - ((float64(3.14159) / v3) / 1.5)))
                    case "a":
                        var v24 int32 = (v22 / ((966) | 1))
                        _ = v24
                    }
                }
                _ = fi21
                _ = v22
            }
            for fi25 := 0; fi25 < 15; fi25++ {
                v17 = ("hello" + "foo bar")
                var v26 uint8 = 25
                var v27 int16 = 1000
                if v4 {
                    if (!v4) {
                        var v28 bool = v4
                        _ = len(v2)
                        _ = v28
                    }
                } else {
                    v29 := GenS1{F0: uint32(v26), F1: fi20, F2: 2147483647}
                    _ = v29
                }
                if ((false && (1 > 18446744073709551615)) || (v17 < v17)) {
                    if true {
                        switch (v17 + v17) {
                        case "a":
                            if false {
                                _ = len(v1)
                            } else {
                                if v4 {
                                    v26 = v26
                                    delete(v2, "foo bar")
                                    var v30 int8 = 0
                                    _ = v30
                                } else {
                                    switch -9223372036854775808 {
                                    case 1:
                                        fmt.Println(v15, v26)
                                    case 9223372036854775807:
                                        v31 := GenS1{F0: 42, F1: (fi20 | -195), F2: 1}
                                        if (v4 || false) {
                                            v3 -= v5
                                            v17 = v17
                                            fmt.Println(v19, v26)
                                        }
                                        _ = v31
                                    default:
                                        if (((true && v4) && v4) && v4) {
                                            v27 = (v27 * v27)
                                            v4 = (true && (!v4))
                                        } else {
                                            v27 = (v15 >> ((945) & 63))
                                            var v32 float64 = v3
                                            _ = v32
                                        }
                                    }
                                }
                                var v33 int8 = -126
                                v34 := v2[("\n\t" + (v17 + v17))]
                                _ = v33
                                _ = v34
                            }
                        }
                        if (2 > 787) {
                            var v35 uint16 = (v19 / (((v19 * v19)) | 1))
                            _ = v35
                        } else {
                            if (2147483647 == 2147483647) {
                                v36 := GenS1{F0: 2147483648, F1: fi20, F2: -391}
                                _ = v36
                            } else {
                                v27 -= v15
                            }
                        }
                    }
                    var v37 uint64 = 571
                    switch "" {
                    case "foo bar":
                        if ("a" < (v17 + "")) {
                            v19 = v19
                            if (false || (64 > v26)) {
                                fmt.Println(fi25)
                                _ = len(v2)
                                if v4 {
                                    switch (v17 + "") {
                                    case "a":
                                        var v38 string = ("\n\t" + "a")
                                        _ = v38
                                    case "":
                                        v3 = 1.5
                                    default:
                                        v39 := GenS2{F0: v26, F1: 42}
                                        _ = v39
                                    }
                                } else {
                                    fmt.Println(v5, fi25, v37)
                                }
                                fmt.Println(v19, v5)
                                if true {
                                    v17 = "foo bar"
                                    v27 = v27
                                    v40 := GenS3{F0: v15, F1: 0}
                                    _ = v40
                                }
                                v41 := GenS4{F0: 2147483647, F1: v17, F2: (^(v27))}
                                _ = v41
                            } else {
                                var v42 int = 9223372036854775807
                                v27 += (^(v27))
                                if (!v4) {
                                    v43 := GenS4{F0: 42, F1: ("foo bar" + "hello"), F2: -292}
                                    _ = v43
                                }
                                _ = v42
                            }
                        }
                    case "a":
                        v5 += float64(0.1)
                    }
                    v44 := GenS4{F0: 0, F1: (v17 + v17), F2: (v27 & (v15 / ((v27) | 1)))}
                    switch "hello" {
                    case "":
                        switch "a" {
                        case "foo bar":
                            v19 += v19
                        case "\n\t":
                            v45 := GenS1{F0: 2147483647, F1: (fi20 / (((fi20 - fi25)) | 1)), F2: 586}
                            _ = v45
                        case "a":
                            v46 := GenS5{F0: v27, F1: v17, F2: 42}
                            _ = v46
                        default:
                            var v47 uint32 = 785
                            _ = v47
                        }
                    default:
                        if v4 {
                            v48 := GenS6{F0: v17}
                            var v49 string = v17
                            switch fi20 {
                            case 801:
                                v49 = "\n\t"
                            case 0:
                                switch ("hello" + v49) {
                                case "foo bar":
                                    fmt.Println(v5)
                                }
                            }
                            _ = v48
                            _ = v49
                        }
                    }
                    _ = v37
                    _ = v44
                }
                _ = fi25
                _ = v26
                _ = v27
            }
            v5 = v18
            _ = fi20
        }
        v3 += v18
    } else {
        for fi50 := 0; fi50 < 11; fi50++ {
            v17 = ((v17 + (v17 + v17)) + (((((v17 + "") + "a") + ((v17 + (v17 + "foo bar")) + (v17 + v17))) + (v17 + "")) + ""))
            v3 = v18
            v15 = (v15 + (v15 << ((124) & 63)))
            _ = fi50
        }
        for fi51 := 0; fi51 < 15; fi51++ {
            v15 = (v15 * v15)
            var v52 int8 = -1
            _ = fi51
            _ = v52
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
