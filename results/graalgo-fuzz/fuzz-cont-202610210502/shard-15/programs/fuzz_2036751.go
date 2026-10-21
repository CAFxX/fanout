package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint; F2 string }
type GenS3 struct { F0 int; F1 int64; F2 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 uint }
func genFn0(p0 bool, p1 bool, p2 uint16) int64 {
    p1 = p1
    p1 = (0 != -990)
    var v6 int = 306
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 796
}
func genFn1() uint8 {
    var v7 uint = 255
    var v8 int = 2147483647
    v9 := []int{v8}
    v9 = append(v9, v8)
    v7 = (v7 << ((v7) & 63))
    _ = v7
    _ = v8
    _ = v9
    return 119
}

func main() {
    v1 := []string{("\n\t" + "foo bar"), "", "\n\t"}
    var v2 int64 = 1
    var v3 bool = true
    var v4 int16 = 0
    v5 := make(map[uint32]string)
    var v10 int8 = -115
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    fmt.Println(v2, v10)
    v5[2147483648] = "a"
    if (v10 < (v10 | v10)) {
        for fi11 := 0; fi11 < 7; fi11++ {
            v5[65535] = ("\n\t" + ("\n\t" + (("\n\t" + "foo bar") + "foo bar")))
            v12 := GenS1{F0: 0}
            for fi13 := 0; fi13 < 2; fi13++ {
                v3 = false
                fmt.Println(v2)
                v14 := GenS1{F0: 255}
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                switch "hello" {
                case "hello":
                    if v3 {
                        fmt.Println(fi15, fi15)
                        v2 += v2
                        if v3 {
                            fmt.Println(v2, v2)
                            switch fi16 {
                            case 128:
                                if (0 >= 0) {
                                    var v17 uint8 = 0
                                    fmt.Println(v2)
                                    _ = v17
                                }
                            case -2147483648:
                                if v3 {
                                    if v3 {
                                        if v3 {
                                            v2 += genFn0(true, false, 337)
                                        }
                                        v10 += v10
                                    }
                                    v18 := GenS1{F0: 32767}
                                    if v3 {
                                        v19 := GenS2{F0: fi16, F1: 250, F2: "hello"}
                                        _ = v19
                                    }
                                    _ = v18
                                } else {
                                    if v3 {
                                        v20 := GenS2{F0: 2147483647, F1: 806, F2: "foo bar"}
                                        v4 -= 0
                                        switch "\n\t" {
                                        case "hello":
                                            v4 = (v4 ^ 1000)
                                        }
                                        _ = v20
                                    } else {
                                        v5[2] = ("foo bar" + (("hello" + (((((("foo bar" + "\n\t") + ("foo bar" + "foo bar")) + (("\n\t" + ("" + "hello")) + "")) + ((("foo bar" + "") + "a") + "hello")) + "") + "hello")) + "\n\t"))
                                        if true {
                                            v10 -= 0
                                            if v3 {
                                                switch 944 {
                                                case 0:
                                                    var v21 uint8 = 128
                                                    _ = v21
                                                case 10:
                                                    v4 = v4
                                                default:
                                                    if v3 {
                                                        v1 = append(v1, "\n\t")
                                                        switch (("hello" + ("a" + "hello")) + "a") {
                                                        case "hello":
                                                            if (741 <= 2147483648) {
                                                                var v22 float32 = float32(((3.14159 + 1e10) / 1.0))
                                                                if v3 {
                                                                    v5[2147483647] = ""
                                                                    fmt.Println(v22, v3)
                                                                    var v23 int = fi16
                                                                    _ = v23
                                                                } else {
                                                                    v24 := GenS3{F0: (fi15 % ((fi16) | 1)), F1: v2, F2: fi16}
                                                                    var v25 uint32 = 2147483647
                                                                    v4 = (v4 / ((v4) | 1))
                                                                    _ = v24
                                                                    _ = v25
                                                                }
                                                                _ = v22
                                                            }
                                                        default:
                                                            v26 := GenS2{F0: (fi16 << ((4294967295) & 63)), F1: 1, F2: "hello"}
                                                            _ = v26
                                                        }
                                                        var v27 int64 = (^(v2))
                                                        v3 = (!v3)
                                                        _ = v27
                                                    }
                                                }
                                                v10 += v10
                                            } else {
                                                var v28 string = ("\n\t" + (("hello" + ("" + "foo bar")) + ("hello" + ("" + "\n\t"))))
                                                _ = v28
                                            }
                                            var v29 uint8 = 236
                                            var v30 uint32 = 42
                                            _ = len(v5)
                                            v1 = append(v1, "foo bar")
                                            fmt.Println(fi15, v29)
                                            _ = v29
                                            _ = v30
                                        }
                                        v31 := GenS3{F0: fi15, F1: (v2 / (((-(v2))) | 1)), F2: fi16}
                                        v32 := GenS1{F0: 32767}
                                        _ = v31
                                        _ = v32
                                    }
                                    var v33 float32 = (0.1 - 1.5)
                                    v2 -= (-(v2))
                                    var v34 int64 = (v2 ^ (-(v2)))
                                    _ = v33
                                    _ = v34
                                }
                            }
                            var v35 uint32 = 65535
                            if v3 {
                                v4 = (^(v4))
                            }
                            _ = len(v1)
                            _ = v35
                        } else {
                            var v36 float32 = (3.14159 - 0.0)
                            v37 := GenS1{F0: 137}
                            _ = v36
                            _ = v37
                        }
                        var v38 int32 = -1
                        switch "\n\t" {
                        case "foo bar":
                            var v39 int64 = (-(v2))
                            _ = v39
                        case "\n\t":
                            if v3 {
                                v4 = int16(v38)
                            }
                        }
                        _ = v38
                    }
                }
                _ = len(v1)
                _ = fi16
            }
            v4 = v4
            if (2 < 0) {
                v40 := GenS4{F0: v4}
                if len(v1) > 0 {
                    v1[0] = ("" + ("hello" + "\n\t"))
                }
                v2 = -9223372036854775808
                _ = v40
            } else {
                fmt.Println(v3)
                switch fi15 {
                case -9223372036854775808:
                    for fi41 := 0; fi41 < 15; fi41++ {
                        v10 = v10
                        _ = fi41
                    }
                default:
                    var v42 float32 = 0.1
                    _ = v42
                }
                v3 = false
            }
            _ = fi15
        }
        var v43 string = "hello"
        _ = v43
    }
    switch "" {
    case "foo bar":
        var v44 uint32 = 65535
        _ = v44
    default:
        v45 := GenS1{F0: 177}
        _ = v45
    }
    fmt.Println(v10, v3, v3)
    if v3 {
        v2 -= 42
        for fi46 := 0; fi46 < 20; fi46++ {
            v10 += v10
            if true {
                switch (^(fi46)) {
                case -433:
                    _ = len(v5)
                }
            }
            v5[2] = (("\n\t" + (("hello" + "hello") + ("a" + ("" + ((("" + "hello") + "foo bar") + "\n\t"))))) + ("foo bar" + "foo bar"))
            _ = fi46
        }
    } else {
        if (v3 || v3) {
            v4 = 1
        } else {
            var v47 int8 = (v10 - v10)
            fmt.Println(v47, v3, v4)
            _ = v47
        }
        var v48 uint = 18446744073709551615
        _ = v48
    }
    var v49 uint32 = 1
    var v50 int32 = 2147483647
    v51 := GenS4{F0: v4}
    v52 := GenS5{F0: (^(v50))}
    for fi53 := 0; fi53 < 13; fi53++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch "foo bar" {
        case "a":
            var v54 int32 = (v50 | v50)
            _ = v54
        case "\n\t":
            for fi55 := 0; fi55 < 14; fi55++ {
                v56 := v5[(v49 + (v49 - 2147483647))]
                if len(v1) > 0 {
                    v1[0] = (v56 + "hello")
                }
                _ = fi55
                _ = v56
            }
        case "":
            if (!(!v3)) {
                v57 := GenS6{F0: 0, F1: 0}
                _ = v57
            }
        }
        _ = fi53
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v10 = v10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
