package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int32 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS3 struct { F0 bool; F1 int8 }
func genFn0(p0 int8) uint8 {
    v7 := GenS1{F0: 177, F1: 42, F2: 165}
    var v8 int64 = 42
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint32 = 4294967295
        _ = fi9
        _ = v10
    }
    p0 += (p0 * p0)
    if true {
        var v11 string = "hello"
        v12 := []int{127}
        v11 = ((v11 + (v11 + v11)) + (v11 + ((v11 + v11) + (v11 + "hello"))))
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v7
    _ = v8
    return 5
}

func main() {
    var v1 int8 = 127
    v2 := []string{"foo bar", "a", "a"}
    var v3 uint16 = 1
    var v4 int16 = -1
    v5 := []string{"foo bar", ("hello" + ((("\n\t" + "hello") + "foo bar") + "\n\t")), ("" + ("a" + ("hello" + ""))), "\n\t"}
    v6 := make(map[int]string)
    v6[241] = "a"
    for fi13 := 0; fi13 < 13; fi13++ {
        switch ("" + ("foo bar" + "foo bar")) {
        case "a":
            fmt.Println(v4)
        }
        if ((-2.25 / 1e10) >= 0.0) {
            var v14 uint8 = genFn0(-128)
            if false {
                v3 += (^(v3))
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS1{F0: 592, F1: (v4 >> ((271) & 63)), F2: -2147483648}
                    var v17 int32 = -1
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                switch (("foo bar" + "\n\t") + "a") {
                case "hello":
                    if (!(v1 <= v1)) {
                        v4 = v4
                        var v18 uint64 = 679
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            v3 -= (v3 % ((v3) | 1))
                            _ = fi19
                        }
                        switch ("foo bar" + "") {
                        case "\n\t":
                            switch fi13 {
                            case 4294967295:
                                if (!false) {
                                    var v20 int64 = 30
                                    _ = v20
                                } else {
                                    _ = len(v6)
                                    v6[fi13] = "foo bar"
                                    delete(v6, fi13)
                                }
                            case -546:
                                fmt.Println(v4)
                            default:
                                var v21 int = fi13
                                _ = v21
                            }
                        case "foo bar":
                            if true {
                                if false {
                                    v4 = v4
                                    v22 := GenS1{F0: int(v4), F1: 32767, F2: -397}
                                    _ = len(v6)
                                    _ = v22
                                }
                            } else {
                                v3 = (^(v3))
                                v1 -= v1
                                if true {
                                    if (9223372036854775807 <= fi13) {
                                        v23 := v6[fi13]
                                        if (!(!true)) {
                                            switch "\n\t" {
                                            case "a":
                                                v23 = "\n\t"
                                            case "\n\t":
                                                var v24 float32 = -2.25
                                                _ = v24
                                            case "":
                                                fmt.Println(fi13, v3, v1)
                                            default:
                                                fmt.Println(v4, v3)
                                            }
                                        } else {
                                            v23 = (v23 + v23)
                                            switch (v23 + "a") {
                                            case "a":
                                                v25 := GenS1{F0: fi13, F1: (v4 % ((v4) | 1)), F2: 495}
                                                _ = v25
                                            case "":
                                                if (false && true) {
                                                    fmt.Println(v1)
                                                } else {
                                                    switch "" {
                                                    case "":
                                                        for fi26 := 0; fi26 < 17; fi26++ {
                                                            v27 := GenS1{F0: (fi26 % (((fi26 & (fi26 & (fi26 ^ fi13)))) | 1)), F1: v4, F2: -203}
                                                            var v28 uint = 4294967295
                                                            _ = fi26
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                    case "hello":
                                                        var v29 uint = 42
                                                        _ = v29
                                                    default:
                                                        var v30 uint = 864
                                                        _ = v30
                                                    }
                                                    v23 = (("a" + v23) + "a")
                                                    if true {
                                                        var v31 uint32 = 0
                                                        v6[fi13] = (("foo bar" + ("a" + v23)) + ((v23 + "a") + v23))
                                                        if len(v5) > 0 {
                                                            v5[0] = v23
                                                        }
                                                        _ = v31
                                                    }
                                                }
                                            }
                                            v32 := GenS2{F0: 65535, F1: v3, F2: 4294967295}
                                            _ = v32
                                        }
                                        var v33 float32 = ((0.1 * float32(3.14159)) / 1.5)
                                        _ = v23
                                        _ = v33
                                    }
                                }
                            }
                        case "a":
                            fmt.Println(fi13)
                        }
                        v34 := GenS1{F0: (fi13 % ((fi13) | 1)), F1: v4, F2: -534}
                        _ = v34
                    }
                case "a":
                    var v35 int32 = 415
                    _ = v35
                }
            }
            _ = v14
        }
        if true {
            v5 = append(v5, "")
            var v36 float64 = -2.25
            _ = v36
        } else {
            fmt.Println(v4, v1)
        }
        _ = fi13
    }
    v3 = (^(v3))
    v37 := GenS3{F0: (65535 != 18446744073709551615), F1: 1}
    for fi38 := 0; fi38 < 19; fi38++ {
        switch 2147483647 {
        case 215:
            fmt.Println(v4, v3)
        case 659:
            var v39 uint = 0
            _ = v39
        default:
            if true {
                var v40 uint8 = genFn0((v1 << ((210) & 63)))
                var v41 int16 = 32767
                _ = v40
                _ = v41
            }
        }
        var v42 uint16 = (v3 & (v3 % ((v3) | 1)))
        _ = fi38
        _ = v42
    }
    var v43 uint = 1
    if len(v2) > 0 {
        v2[0] = "a"
    }
    if ((v1 - 0) != (v1 % ((-1) | 1))) {
        fmt.Println(v1, v1)
        fmt.Println(v4, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
