package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64 }

func main() {
    var v1 uint = 1
    var v2 float32 = -2.25
    v3 := []string{"hello", "", (((((("foo bar" + "foo bar") + ("\n\t" + "\n\t")) + "") + "\n\t") + "a") + ("a" + ("foo bar" + "foo bar")))}
    var v4 int8 = 82
    if (42 < -2147483648) {
        if len(v3) > 0 {
            v3[0] = ""
        }
        v5 := GenS1{F0: -1, F1: 683}
        _ = v5
    }
    if true {
        switch "" {
        case "hello":
            fmt.Println(v2)
        case "foo bar":
            switch 32767 {
            case -9223372036854775808:
                var v6 int16 = 1000
                _ = v6
            case -262:
                v4 -= v4
            }
        case "":
            v7 := GenS1{F0: -32768, F1: -19}
            _ = v7
        default:
            if (0 <= -352) {
                if ("a" == "a") {
                    for fi8 := 0; fi8 < 3; fi8++ {
                        v9 := GenS1{F0: -309, F1: 42}
                        switch ("" + (("a" + "hello") + ("\n\t" + "a"))) {
                        case "a":
                            var v10 float32 = v2
                            _ = v10
                        }
                        var v11 uint32 = 367
                        _ = fi8
                        _ = v11
                        _ = v9
                    }
                    var v12 uint32 = 65535
                    v2 = (v2 + v2)
                    _ = v12
                }
                v13 := GenS1{F0: 422, F1: 0}
                v14 := GenS1{F0: 0, F1: -241}
                _ = v13
                _ = v14
            } else {
                v15 := GenS2{F0: 9223372036854775807, F1: -331}
                var v16 uint32 = 42
                for fi17 := 0; fi17 < 2; fi17++ {
                    for fi18 := 0; fi18 < 8; fi18++ {
                        if (!(0 != 222)) {
                            v16 = (^(v16))
                            v16 += (^(v16))
                        } else {
                            if (!(-1 != 0)) {
                                v19 := GenS3{F0: fi18, F1: ("foo bar" + "")}
                                _ = v19
                            } else {
                                var v20 uint64 = 0
                                var v21 uint = (v1 | v1)
                                _ = v20
                                _ = v21
                            }
                        }
                        v22 := GenS3{F0: fi17, F1: ""}
                        v1 = (v1 >> ((v1) & 63))
                        v16 = v16
                        if (969 < 42) {
                            v23 := GenS3{F0: (fi17 >> (((^(v1))) & 63)), F1: ""}
                            v4 = v4
                            v2 = 3.14159
                            _ = v23
                        } else {
                            _ = len(v3)
                            if false {
                                if false {
                                    v24 := GenS3{F0: fi17, F1: ""}
                                    _ = v24
                                } else {
                                    var v25 uint32 = (v16 >> ((uint(v16)) & 63))
                                    _ = v25
                                }
                                var v26 int = fi18
                                v16 = v16
                                _ = v26
                            }
                        }
                        var v27 bool = ((v4 / (((^(v4))) | 1)) <= v4)
                        if v27 {
                            if false {
                                switch fi17 {
                                case 32768:
                                    if true {
                                        var v28 int64 = 9223372036854775807
                                        _ = v28
                                    }
                                }
                                v3 = append(v3, "\n\t")
                            } else {
                                v29 := GenS4{F0: 265}
                                if (42 > -855) {
                                    if true {
                                        v1 = (v1 + (^(v1)))
                                        fmt.Println(v16)
                                        v2 -= v2
                                        if v27 {
                                            if v27 {
                                                v30 := GenS3{F0: 127, F1: ""}
                                                fmt.Println(v27)
                                                if false {
                                                    v2 += 0.0
                                                    v27 = v27
                                                    v1 += (v1 * (v1 - v1))
                                                    v16 -= (v16 & (v16 | v16))
                                                    switch "hello" {
                                                    case "hello":
                                                        if len(v3) > 0 {
                                                            v3[0] = "hello"
                                                        }
                                                    }
                                                }
                                                _ = v30
                                            } else {
                                                var v31 int = 128
                                                v32 := GenS4{F0: 42}
                                                _ = v31
                                                _ = v32
                                            }
                                        } else {
                                            v33 := GenS3{F0: (fi17 ^ (-(fi18))), F1: "hello"}
                                            _ = v33
                                        }
                                        var v34 string = "hello"
                                        _ = v34
                                    }
                                }
                                _ = v29
                            }
                            if ((fi17 == (fi18 << (((v1 * (v1 | 255))) & 63))) && v27) {
                                v35 := GenS3{F0: (fi17 << ((4294967295) & 63)), F1: ("a" + "foo bar")}
                                _ = v35
                            } else {
                                var v36 int = fi17
                                fmt.Println(fi17, fi17)
                                _ = v36
                            }
                        }
                        fmt.Println(v2, v16, v1)
                        fmt.Println(v4)
                        _ = fi18
                        _ = v22
                        _ = v27
                    }
                    _ = fi17
                }
                _ = v15
                _ = v16
            }
        }
        var v37 uint16 = 255
        var v38 int8 = v4
        _ = v37
        _ = v38
    }
    var v39 int = -228
    v2 = (v2 / v2)
    if true {
        switch (("hello" + (("a" + (("hello" + "foo bar") + ("\n\t" + (("" + "foo bar") + "hello")))) + ("\n\t" + "foo bar"))) + ("a" + "")) {
        case "\n\t":
            v1 = v1
        default:
            for fi40 := 0; fi40 < 20; fi40++ {
                v4 -= v4
                for fi41 := 0; fi41 < 1; fi41++ {
                    if (0 > 4294967295) {
                        fmt.Println(fi40)
                        v42 := GenS2{F0: v39, F1: 1000}
                        _ = v42
                    }
                    fmt.Println(fi41, v1)
                    v43 := GenS5{F0: 120, F1: 1000, F2: 255}
                    var v44 string = "foo bar"
                    var v45 int32 = 853
                    var v46 uint16 = 42
                    _ = fi41
                    _ = v43
                    _ = v44
                    _ = v45
                    _ = v46
                }
                _ = fi40
            }
        }
    } else {
        v47 := GenS6{F0: ("hello" + "\n\t")}
        _ = v47
    }
    v48 := GenS2{F0: v39, F1: 1000}
    var v49 bool = true
    v2 = (v2 * (3.14159 / (1e10 - v2)))
    if v49 {
        var v50 uint16 = 65535
        v51 := GenS7{F0: -1}
        fmt.Println(v49, v49)
        _ = v50
        _ = v51
    }
    v52 := GenS4{F0: 42}
    v4 = (v4 % ((v4) | 1))
    var v53 uint = v1
    switch ("hello" + "hello") {
    case "a":
        v54 := GenS4{F0: 42}
        _ = v54
    default:
        var v55 uint32 = 756
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v48
    _ = v49
    _ = v52
    _ = v53
}
