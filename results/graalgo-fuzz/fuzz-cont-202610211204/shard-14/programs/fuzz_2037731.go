package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int64; F2 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 int8; F1 int }
func genFn0(p0 uint32, p1 int16) uint64 {
    v4 := []int{933, 709}
    for fi5 := 0; fi5 < 18; fi5++ {
        if len(v4) > 0 {
            v4[0] = 32767
        }
        if (true && ((float32(1e10) < 1e10) || (42 == 0))) {
            p0 = (p0 - (p0 >> ((0) & 63)))
        } else {
            if true {
                v6 := GenS1{F0: 255}
                _ = v6
            } else {
                var v7 float64 = ((0.1 / 3.14159) / (((1e10 - 1e10) * 1e10) - 1e10))
                _ = v7
            }
        }
        if true {
            v4 = append(v4, (-(fi5)))
        } else {
            for fi8 := 0; fi8 < 4; fi8++ {
                if true {
                    p0 = p0
                } else {
                    _ = len(v4)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi8
            }
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 4294967295
}
func genFn1(p0 string, p1 int) string {
    fmt.Println(p0, p0)
    v9 := GenS2{F0: false, F1: 2147483647}
    _ = p0
    _ = p1
    _ = v9
    return p0
}

func main() {
    var v1 string = ("a" + "hello")
    v2 := make(map[string]int)
    v3 := []uint32{597}
    v10 := GenS3{F0: 128, F1: 1, F2: 739}
    for fi11 := 0; fi11 < 1; fi11++ {
        switch ("hello" + "\n\t") {
        case "hello":
            var v12 uint16 = 255
            _ = v12
        case "\n\t":
            switch "\n\t" {
            case "foo bar":
                var v13 uint = 566
                _ = v13
            default:
                if true {
                    switch fi11 {
                    case 1:
                        var v14 float32 = 1.5
                        _ = v14
                    case -10:
                        var v15 int8 = 0
                        _ = v15
                    }
                    switch (v1 + (v1 + (v1 + ((v1 + "foo bar") + v1)))) {
                    case "\n\t":
                        fmt.Println(v1, fi11)
                    case "foo bar":
                        fmt.Println(v1, fi11)
                    case "hello":
                        for fi16 := 0; fi16 < 13; fi16++ {
                            v1 = (("" + "a") + "hello")
                            v17 := GenS3{F0: 127, F1: 255, F2: 809}
                            _ = fi16
                            _ = v17
                        }
                    }
                } else {
                    if true {
                        var v18 uint32 = 2147483647
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                            v1 = v1
                            _ = fi19
                        }
                        if true {
                            switch "" {
                            case "a":
                                v1 = ("\n\t" + v1)
                            case "foo bar":
                                fmt.Println(v1, fi11)
                            case "hello":
                                switch v1 {
                                case "a":
                                    for fi20 := 0; fi20 < 19; fi20++ {
                                        fmt.Println(fi11)
                                        v1 = "hello"
                                        _ = fi20
                                    }
                                case "hello":
                                    for fi21 := 0; fi21 < 3; fi21++ {
                                        v1 = "a"
                                        delete(v2, "\n\t")
                                        if true {
                                            v22 := GenS3{F0: 2, F1: 128, F2: 1}
                                            v1 = v1
                                            fmt.Println(v1, fi11)
                                            fmt.Println(fi11, fi21)
                                            switch ((v1 + ((("\n\t" + ("foo bar" + ("foo bar" + (v1 + v1)))) + ("" + v1)) + ((("foo bar" + "\n\t") + "") + v1))) + (v1 + "\n\t")) {
                                            case "foo bar":
                                                var v23 int16 = 32767
                                                _ = v23
                                            default:
                                                v1 = ""
                                            }
                                            _ = v22
                                        } else {
                                            var v24 int = fi11
                                            var v25 int = fi11
                                            fmt.Println(v1, v25, v24)
                                            _ = v24
                                            _ = v25
                                        }
                                        var v26 int64 = 42
                                        delete(v2, "")
                                        _ = fi21
                                        _ = v26
                                    }
                                default:
                                    v3 = append(v3, 1)
                                }
                            }
                            var v27 uint8 = 255
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 14; fi28++ {
                            v1 = "a"
                            _ = fi28
                        }
                    }
                    v1 = ("\n\t" + v1)
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                }
            }
        case "a":
            var v29 float64 = 0.1
            _ = v29
        }
        if false {
            v30 := GenS1{F0: 4294967295}
            var v31 int = (fi11 * (fi11 & fi11))
            v31 -= -332
            _ = v30
            _ = v31
        } else {
            v32 := GenS4{F0: 65535, F1: 64, F2: -163}
            _ = v32
        }
        _ = fi11
    }
    v1 = (("hello" + v1) + v1)
    v1 = "\n\t"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = ("" + v1)
    var v33 int32 = -1
    switch 128 {
    case 1:
        for fi34 := 0; fi34 < 9; fi34++ {
            v33 = (v33 & -1)
            if len(v3) > 0 {
                v3[0] = 200
            }
            _ = fi34
        }
    case 127:
        v2["foo bar"] = -9223372036854775808
    case -1:
        fmt.Println(v1)
    }
    var v35 string = v1
    v36 := v2[v35]
    v37 := GenS5{F0: v36}
    v35 = ((v35 + "a") + (v1 + v1))
    v38 := GenS6{F0: 127}
    switch v36 {
    case -251:
        var v39 string = "a"
        _ = v39
    case 132:
        v40 := GenS7{F0: ("\n\t" >= v35), F1: ("\n\t" >= v1)}
        _ = v40
    case 4294967295:
        var v41 float32 = float32(0.0)
        _ = v41
    default:
        switch "" {
        case "foo bar":
            v42 := GenS8{F0: 0, F1: 32767}
            _ = v42
        case "a":
            fmt.Println(v35)
        case "hello":
            var v43 uint8 = 1
            _ = v43
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    v3 = append(v3, 4294967295)
    switch "foo bar" {
    case "hello":
        _ = len(v3)
    case "\n\t":
        fmt.Println(v36, v35, v36)
    }
    fmt.Println(v33)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
