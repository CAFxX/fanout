package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint; F2 string }
type GenS3 struct { F0 uint16; F1 string; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 bool; F1 uint; F2 uint32 }
func genFn0(p0 string) int32 {
    p0 = ("\n\t" + p0)
    v5 := GenS1{F0: 166}
    for fi6 := 0; fi6 < 13; fi6++ {
        p0 = "foo bar"
        p0 = (((p0 + "\n\t") + p0) + p0)
        _ = fi6
    }
    v7 := GenS2{F0: 0, F1: 18446744073709551615, F2: (p0 + "a")}
    _ = p0
    _ = v5
    _ = v7
    return 42
}
func genFn1() uint16 {
    var v8 uint = 830
    if true {
        fmt.Println(v8)
    }
    _ = v8
    return 138
}

func main() {
    var v1 int64 = 96
    v2 := []string{("\n\t" + "foo bar"), ""}
    var v3 int8 = 42
    v4 := make(map[string]bool)
    v1 += (v1 >> ((255) & 63))
    fmt.Println(v3, v1)
    var v9 int8 = v3
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 float64 = (float64(0.1) / 0.1)
        if (1000 < 0) {
            v4["\n\t"] = (!(false || (!false)))
            v3 -= (v9 - 0)
        }
        v12 := GenS2{F0: (v3 * -1), F1: 4294967295, F2: ("foo bar" + "a")}
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = len(v4)
    _ = len(v2)
    switch 32768 {
    case 1:
        if false {
            for fi13 := 0; fi13 < 20; fi13++ {
                switch -9223372036854775808 {
                case 416:
                    switch "foo bar" {
                    case "a":
                        v9 = v3
                    case "hello":
                        v14 := GenS3{F0: 0, F1: ("hello" + "a"), F2: 11}
                        _ = v14
                    case "":
                        v3 = 66
                    }
                }
                _ = fi13
            }
            switch "\n\t" {
            case "foo bar":
                v2 = append(v2, "a")
            case "a":
                var v15 float32 = 1.5
                _ = v15
            case "\n\t":
                v9 += v3
            }
        }
    case 4294967295:
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint = 65535
            switch "hello" {
            case "hello":
                var v18 float32 = -2.25
                _ = v18
            case "":
                if (1 == 32767) {
                    fmt.Println(v1)
                    if true {
                        switch -203 {
                        case 42:
                            for fi19 := 0; fi19 < 18; fi19++ {
                                var v20 uint32 = 27
                                v9 -= v3
                                _ = fi19
                                _ = v20
                            }
                        case 175:
                            for fi21 := 0; fi21 < 1; fi21++ {
                                delete(v4, "")
                                if true {
                                    _ = len(v4)
                                    switch "hello" {
                                    case "foo bar":
                                        var v22 int32 = 1
                                        _ = v22
                                    case "a":
                                        var v23 uint64 = 301
                                        _ = v23
                                    }
                                    var v24 float64 = 0.0
                                    _ = v24
                                } else {
                                    fmt.Println(v3, v3)
                                }
                                if true {
                                    fmt.Println(fi21)
                                    var v25 uint = (^(v17))
                                    _ = v25
                                }
                                _ = fi21
                            }
                        }
                        delete(v4, "a")
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                    }
                } else {
                    for fi26 := 0; fi26 < 17; fi26++ {
                        fmt.Println(v9)
                        if (!false) {
                            var v27 float64 = (((0.0 + 1.5) - float64(-2.25)) + 0.0)
                            var v28 uint32 = 461
                            v3 = (v9 - v9)
                            _ = v27
                            _ = v28
                        }
                        v9 = (v9 + 127)
                        _ = fi26
                    }
                    v2 = append(v2, "a")
                    switch "\n\t" {
                    case "hello":
                        v9 -= v9
                    default:
                        var v29 float64 = ((1.5 - (float64(3.14159) / float64(1e10))) - 0.1)
                        _ = v29
                    }
                }
            case "\n\t":
                var v30 string = ""
                _ = v30
            }
            for fi31 := 0; fi31 < 11; fi31++ {
                if true {
                    fmt.Println(v17)
                    fmt.Println(v1)
                } else {
                    v17 = v17
                }
                var v32 int8 = v3
                _ = fi31
                _ = v32
            }
            _ = fi16
            _ = v17
        }
    case 42:
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 uint = 1
            fmt.Println(fi33)
            v35 := GenS3{F0: 1, F1: "foo bar", F2: -1}
            _ = fi33
            _ = v34
            _ = v35
        }
    default:
        var v36 float32 = -2.25
        _ = v36
    }
    v37 := GenS4{F0: 42, F1: 50, F2: 2}
    var v38 uint = 0
    v39 := GenS5{F0: genFn0("hello"), F1: "foo bar"}
    v3 += v9
    v38 -= 18446744073709551615
    fmt.Println(v3)
    v9 = (-(v9))
    fmt.Println(v38, v1)
    fmt.Println(v3, v9, v3)
    if true {
        var v40 uint64 = 18446744073709551615
        v41 := GenS6{F0: false, F1: v38, F2: 1}
        if len(v2) > 0 {
            v2[0] = ((("a" + "\n\t") + "") + ("hello" + "hello"))
        }
        _ = v40
        _ = v41
    }
    if false {
        v2 = append(v2, ("\n\t" + ("hello" + "hello")))
        var v42 bool = false
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v9
}
