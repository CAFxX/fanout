package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 string; F1 int }

func main() {
    v1 := []uint32{0, 2147483648, 0, 498}
    v2 := []int{4294967295, -9223372036854775808, -250, 128}
    v3 := make(map[string]int)
    var v4 bool = (true && (!(65 != 42)))
    v4 = v4
    switch -628 {
    case 32767:
        for fi5 := 0; fi5 < 13; fi5++ {
            var v6 int = (fi5 % ((-2147483648) | 1))
            delete(v3, "hello")
            fmt.Println(fi5)
            _ = fi5
            _ = v6
        }
    case 4294967295:
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 bool = v4
                v10 := GenS1{F0: (fi7 & 938), F1: 127}
                v9 = true
                v11 := GenS1{F0: fi8, F1: 1}
                if (0 >= 1) {
                    fmt.Println(v4)
                    if false {
                        var v12 uint32 = 65535
                        v13 := GenS2{F0: 548, F1: v9}
                        _ = v12
                        _ = v13
                    } else {
                        v9 = (816 >= 255)
                        v3["foo bar"] = fi8
                        var v14 uint32 = 42
                        fmt.Println(v4, v9)
                        _ = v14
                    }
                    v4 = v4
                    switch fi8 {
                    case 42:
                        v4 = (v9 || v4)
                    default:
                        v9 = (!(!(!false)))
                        v15 := GenS3{F0: (((("a" + (((("" + "a") + ("foo bar" + "\n\t")) + "") + (("hello" + "") + ((("" + "foo bar") + "") + (((("" + "a") + "a") + "a") + "a"))))) + "hello") + "hello") + "\n\t"), F1: true}
                        _ = v15
                    }
                    v9 = v4
                } else {
                    v16 := v3[("\n\t" + "a")]
                    switch -2147483648 {
                    case 911:
                        fmt.Println(v9, fi8)
                    case 466:
                        var v17 uint8 = 140
                        _ = v17
                    default:
                        v18 := GenS3{F0: "", F1: false}
                        _ = v18
                    }
                    _ = v16
                }
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi7
        }
    }
    var v19 uint32 = 20
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS4{F0: (("\n\t" + "foo bar") + "\n\t"), F1: -283}
        fmt.Println(fi20, v19)
        v4 = (0.0 != -2.25)
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        if false {
            var v23 int64 = -9223372036854775808
            v23 -= 42
            _ = v23
        } else {
            var v24 uint64 = 9
            fmt.Println(v24, v24, v19)
            if v4 {
                if (65535 > v19) {
                    for fi25 := 0; fi25 < 7; fi25++ {
                        var v26 int64 = 1
                        v4 = (true && (42 == -62))
                        v4 = v4
                        _ = fi25
                        _ = v26
                    }
                } else {
                    v24 = 48
                    for fi27 := 0; fi27 < 12; fi27++ {
                        if ("a" == "a") {
                            v28 := GenS3{F0: "a", F1: (((1.5 - 1.5) + 1e10) == 1.5)}
                            _ = v28
                        }
                        v4 = false
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi27
                    }
                    v29 := v3["hello"]
                    _ = v29
                }
                _ = len(v1)
            }
            _ = v24
        }
        v30 := GenS1{F0: -9223372036854775808, F1: 99}
        _ = fi22
        _ = v30
    }
    var v31 bool = (42 == 2147483647)
    v32 := GenS2{F0: 920, F1: true}
    fmt.Println(v4, v31, v19)
    var v33 float64 = float64(1.5)
    for fi34 := 0; fi34 < 15; fi34++ {
        v31 = (uint16(v19) >= 359)
        _ = fi34
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        switch "foo bar" {
        case "foo bar":
            for fi36 := 0; fi36 < 12; fi36++ {
                v33 = (1e10 - 1.5)
                v33 = v33
                var v37 int8 = 1
                var v38 uint16 = 839
                _ = fi36
                _ = v37
                _ = v38
            }
        }
        _ = fi35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
