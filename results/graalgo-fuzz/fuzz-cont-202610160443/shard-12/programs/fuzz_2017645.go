package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint32; F1 int; F2 int32 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{}
    v3 := []string{("hello" + (("" + ("" + ("\n\t" + "foo bar"))) + "")), "\n\t", (("" + ("foo bar" + (("a" + ((("\n\t" + "") + "\n\t") + ("a" + ((((("\n\t" + "a") + ("foo bar" + "foo bar")) + "hello") + "hello") + "a")))) + "foo bar"))) + "\n\t"), ""}
    var v4 int64 = -1
    if ((163 >= 255) || true) {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        v5 := GenS1{F0: 1000, F1: "a"}
        _ = v5
    }
    v1 = v1
    if (9223372036854775807 > 127) {
        if false {
            for fi6 := 0; fi6 < 17; fi6++ {
                fmt.Println(v4, v1, v4)
                v1 = v1
                v1 += (v1 / ((v1) | 1))
                _ = fi6
            }
        } else {
            v7 := GenS1{F0: 1, F1: "hello"}
            var v8 int32 = 42
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 uint64 = 1
            v2 = append(v2, 65535)
            _ = fi9
            _ = v10
        }
    } else {
        v3 = append(v3, "foo bar")
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v11 int = -1
    v12 := GenS1{F0: -32768, F1: ""}
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v13 := GenS1{F0: 1, F1: ("a" + "hello")}
    _ = len(v3)
    v4 += (v4 - (v4 - (v4 + 42)))
    if (100 < -83) {
        _ = len(v2)
        var v14 uint16 = 255
        _ = v14
    } else {
        _ = len(v2)
        for fi15 := 0; fi15 < 3; fi15++ {
            if true {
                fmt.Println(v1, v1)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                if true {
                    for fi16 := 0; fi16 < 10; fi16++ {
                        if false {
                            var v17 uint32 = 2147483648
                            var v18 uint = 42
                            _ = v17
                            _ = v18
                        }
                        _ = fi16
                    }
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                } else {
                    var v19 bool = true
                    if (-32768 < 315) {
                        for fi20 := 0; fi20 < 4; fi20++ {
                            if v19 {
                                if (2147483647 == 1) {
                                    if (v4 < 9223372036854775807) {
                                        if (("hello" + "hello") > ("hello" + "a")) {
                                            fmt.Println(fi15)
                                            var v21 float64 = (0.1 - 3.14159)
                                            _ = v21
                                        } else {
                                            v19 = (42 <= 1)
                                            v1 -= (v1 / (((v1 & v1)) | 1))
                                            v22 := GenS1{F0: 32767, F1: "hello"}
                                            _ = v22
                                        }
                                    }
                                }
                                v1 = 18446744073709551615
                            } else {
                                if (!false) {
                                    switch (^(v11)) {
                                    case 544:
                                        var v23 uint8 = 157
                                        _ = v23
                                    }
                                    fmt.Println(v4, v19)
                                } else {
                                    if (-1 < (v4 ^ 9223372036854775807)) {
                                        v11 += fi15
                                        if (-32768 == 0) {
                                            fmt.Println(v19, v19, v4)
                                        } else {
                                            if (!v19) {
                                                var v24 bool = (32767 < 1000)
                                                v19 = v19
                                                _ = v24
                                            } else {
                                                if v19 {
                                                    v25 := GenS1{F0: -440, F1: "a"}
                                                    _ = v25
                                                }
                                                fmt.Println(v19, fi20, v11)
                                            }
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                        if v19 {
                                            v11 += fi20
                                            fmt.Println(fi20)
                                        } else {
                                            v26 := GenS1{F0: 934, F1: ""}
                                            v19 = (3.14159 > 0.1)
                                            var v27 int64 = v4
                                            _ = v26
                                            _ = v27
                                        }
                                        var v28 int = (fi15 + 889)
                                        v19 = v19
                                        v11 = fi20
                                        _ = v28
                                    } else {
                                        var v29 uint = 0
                                        fmt.Println(fi20)
                                        switch ("\n\t" + "hello") {
                                        case "hello":
                                            var v30 int16 = 1000
                                            _ = v30
                                        }
                                        var v31 float32 = ((0.0 * 0.1) - 0.0)
                                        var v32 int = -17
                                        v29 = (v29 + v29)
                                        _ = v29
                                        _ = v31
                                        _ = v32
                                    }
                                    var v33 int8 = 48
                                    v34 := GenS2{F0: 65535, F1: (v33 >> ((1) & 63))}
                                    _ = v33
                                    _ = v34
                                }
                                v1 = v1
                                if ((v1 != v1) && v19) {
                                    var v35 uint32 = 571
                                    v36 := GenS3{F0: v35, F1: v11, F2: 42}
                                    var v37 string = "foo bar"
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                            }
                            fmt.Println(v19)
                            _ = fi20
                        }
                    } else {
                        switch fi15 {
                        case 460:
                            switch ("hello" + "a") {
                            case "hello":
                                fmt.Println(v11, v4)
                            case "a":
                                var v38 uint = 684
                                _ = v38
                            case "":
                                var v39 uint32 = 42
                                _ = v39
                            }
                        case 0:
                            switch "foo bar" {
                            case "foo bar":
                                if v19 {
                                    v40 := GenS2{F0: 2, F1: -128}
                                    var v41 string = "\n\t"
                                    for fi42 := 0; fi42 < 18; fi42++ {
                                        if (!(1 == -2147483648)) {
                                            var v43 uint64 = (v1 / (((v1 << ((18446744073709551615) & 63))) | 1))
                                            fmt.Println(v41)
                                            fmt.Println(v11, v1)
                                            _ = v43
                                        }
                                        switch ("hello" + v41) {
                                        case "foo bar":
                                            v11 = (fi42 & fi15)
                                        }
                                        _ = fi42
                                    }
                                    _ = v40
                                    _ = v41
                                }
                            case "\n\t":
                                v19 = (!((-2.25 > -2.25) && v19))
                            }
                        default:
                            v11 = fi15
                        }
                        v19 = false
                    }
                    _ = v19
                }
                var v44 float32 = float32(3.14159)
                _ = v44
            }
            _ = fi15
        }
        for fi45 := 0; fi45 < 13; fi45++ {
            for fi46 := 0; fi46 < 15; fi46++ {
                v47 := GenS1{F0: -154, F1: ("foo bar" + "")}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi46
                _ = v47
            }
            v48 := GenS2{F0: 536, F1: 42}
            _ = len(v3)
            _ = fi45
            _ = v48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
