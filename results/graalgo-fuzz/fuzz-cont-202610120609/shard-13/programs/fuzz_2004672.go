package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS4 struct { F0 int; F1 uint32; F2 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16 }

func main() {
    v1 := []string{"a"}
    var v2 float32 = 0.0
    v3 := []string{"a", ("" + "foo bar"), "\n\t", ""}
    var v4 uint32 = 2147483648
    var v5 uint32 = v4
    if (-599 < -881) {
        if true {
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 float32 = (-2.25 + ((v2 / 1.5) + v2))
                var v8 uint = 255
                _ = fi6
                _ = v7
                _ = v8
            }
        } else {
            if ((-518 > 42) && false) {
                v5 -= v4
            } else {
                v5 = v4
            }
        }
        var v9 uint32 = 697
        _ = v9
    } else {
        if true {
            for fi10 := 0; fi10 < 1; fi10++ {
                if (false && true) {
                    if true {
                        switch -2147483648 {
                        case -9223372036854775808:
                            for fi11 := 0; fi11 < 18; fi11++ {
                                v2 = (v2 + 3.14159)
                                _ = fi11
                            }
                        case -599:
                            v12 := GenS1{F0: 567, F1: 42}
                            _ = v12
                        case 4294967295:
                            v4 = v4
                        }
                        var v13 uint = 775
                        _ = v13
                    }
                    if true {
                        if true {
                            var v14 string = (("\n\t" + "a") + (("" + "foo bar") + "\n\t"))
                            v15 := GenS1{F0: int32(v5), F1: 2}
                            _ = v14
                            _ = v15
                        } else {
                            for fi16 := 0; fi16 < 7; fi16++ {
                                v4 += 42
                                if false {
                                    v17 := GenS2{F0: "\n\t", F1: 1}
                                    v5 = v5
                                    _ = v17
                                } else {
                                    v2 = v2
                                }
                                var v18 uint8 = 39
                                var v19 int16 = -32768
                                _ = fi16
                                _ = v18
                                _ = v19
                            }
                        }
                        v20 := GenS3{F0: (v5 >> ((4294967295) & 63)), F1: 252, F2: -128}
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 15; fi21++ {
                            v22 := GenS4{F0: (fi21 | (fi21 ^ -9223372036854775808)), F1: 462, F2: 17}
                            if true {
                                var v23 uint64 = 317
                                var v24 float64 = ((float64((-2.25 / 1.0)) - float64(v2)) * float64(v2))
                                _ = v23
                                _ = v24
                            } else {
                                v2 += 3.14159
                                fmt.Println(v4, fi21, v2)
                                var v25 int = (fi21 ^ (fi10 + fi21))
                                var v26 int = fi21
                                v5 = v5
                                switch ("hello" + "") {
                                case "\n\t":
                                    switch ("a" + ("hello" + "foo bar")) {
                                    case "hello":
                                        v1 = append(v1, "hello")
                                    case "a":
                                        _ = len(v1)
                                    }
                                case "foo bar":
                                    var v27 uint8 = 255
                                    _ = v27
                                case "":
                                    switch fi21 {
                                    case 314:
                                        switch (v25 << ((65535) & 63)) {
                                        case -122:
                                            v25 = fi21
                                        case 32767:
                                            v4 -= 2147483648
                                        case -584:
                                            if len(v1) > 0 {
                                                v1[0] = ""
                                            }
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = "foo bar"
                                            }
                                            switch ("foo bar" + ("hello" + ("\n\t" + ""))) {
                                            case "a":
                                                var v28 int32 = -987
                                                _ = v28
                                            case "foo bar":
                                                if len(v1) > 0 {
                                                    v1[0] = "a"
                                                }
                                            default:
                                                v26 = (fi10 * v26)
                                            }
                                            v3 = append(v3, "\n\t")
                                            var v29 uint16 = 714
                                            _ = v29
                                        }
                                    case -1:
                                        var v30 uint16 = 65535
                                        _ = v30
                                    case 4294967295:
                                        var v31 uint32 = v5
                                        _ = v31
                                    }
                                }
                                if ((!false) && true) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v5 -= (v4 + 42)
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                }
                                _ = v25
                                _ = v26
                            }
                            var v32 float32 = v2
                            _ = fi21
                            _ = v22
                            _ = v32
                        }
                    }
                } else {
                    for fi33 := 0; fi33 < 19; fi33++ {
                        var v34 int8 = 1
                        switch "foo bar" {
                        case "\n\t":
                            v2 = ((-2.25 + v2) * (v2 + float32(3.14159)))
                        case "foo bar":
                            if (false || false) {
                                _ = len(v1)
                                v34 = v34
                                _ = len(v3)
                                v35 := GenS5{F0: ""}
                                switch (fi10 + (fi10 % ((32768) | 1))) {
                                case -1:
                                    fmt.Println(v34)
                                default:
                                    v4 -= (v5 | 2147483648)
                                }
                                var v36 string = ("a" + ("" + "hello"))
                                var v37 uint16 = 42
                                v38 := GenS1{F0: 0, F1: 0}
                                _ = v35
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                        case "a":
                            var v39 uint64 = 1
                            _ = v39
                        }
                        _ = fi33
                        _ = v34
                    }
                    var v40 uint64 = 2
                    _ = v40
                }
                _ = fi10
            }
            fmt.Println(v4, v4, v4)
        }
        var v41 int32 = -2147483648
        v2 = (v2 + v2)
        _ = v41
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        var v43 int16 = 1
        _ = fi42
        _ = v43
    }
    v4 = 42
    _ = len(v1)
    v44 := GenS6{F0: -1}
    var v45 int16 = 32767
    fmt.Println(v2, v45)
    _ = len(v3)
    var v46 int16 = v45
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
