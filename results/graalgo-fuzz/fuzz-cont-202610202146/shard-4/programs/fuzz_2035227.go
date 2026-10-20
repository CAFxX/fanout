package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 int; F2 int32 }
func genFn0(p0 uint8) int16 {
    p0 = p0
    v4 := GenS1{F0: 32767, F1: -32768}
    p0 += (p0 | (p0 + p0))
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 string = ((("foo bar" + ("" + "hello")) + "a") + "foo bar")
        switch fi5 {
        case 128:
            v7 := GenS1{F0: 32767, F1: 835}
            _ = v7
        default:
            v8 := GenS2{F0: 241, F1: 2147483647}
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v9 := GenS3{F0: 0, F1: 42}
    _ = p0
    _ = v4
    _ = v9
    return 0
}
func genFn1(p0 string) bool {
    for fi10 := 0; fi10 < 10; fi10++ {
        if (!false) {
            p0 = p0
        }
        fmt.Println(p0, p0)
        _ = fi10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(fi11)
        var v12 int32 = 809
        _ = fi11
        _ = v12
    }
    switch 1 {
    case 2147483647:
        var v13 int32 = 42
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(p0)
        p0 = (("a" + p0) + (p0 + p0))
        for fi15 := 0; fi15 < 10; fi15++ {
            switch -2147483648 {
            case 32767:
                var v16 string = p0
                _ = v16
            case 4294967295:
                var v17 int32 = 1
                _ = v17
            case -54:
                v18 := GenS1{F0: (fi15 & 32767), F1: 42}
                _ = v18
            default:
                if (0 == 72) {
                    p0 = ("hello" + p0)
                    switch ("hello" + "a") {
                    case "hello":
                        if true {
                            fmt.Println(p0)
                        }
                    case "\n\t":
                        if (45 >= -1) {
                            if (-2147483648 <= 42) {
                                var v19 float64 = 0.0
                                _ = v19
                            }
                            p0 = (p0 + ((((p0 + p0) + "foo bar") + ("\n\t" + (((p0 + p0) + ((p0 + "\n\t") + p0)) + "foo bar"))) + ""))
                        } else {
                            v20 := []int{}
                            var v21 int8 = -128
                            _ = v20
                            _ = v21
                        }
                    }
                } else {
                    p0 = p0
                    if false {
                        fmt.Println(fi14, fi15, fi14)
                        p0 = (p0 + p0)
                    } else {
                        p0 = p0
                    }
                }
            }
            var v22 int = (fi15 - -9223372036854775808)
            _ = fi15
            _ = v22
        }
        _ = fi14
    }
    switch p0 {
    case "foo bar":
        switch ("" + (p0 + "foo bar")) {
        case "":
            if true {
                p0 = "a"
                p0 = p0
                var v23 int32 = 42
                _ = v23
            } else {
                if false {
                    fmt.Println(p0)
                    var v24 float32 = (0.0 * (3.14159 / (0.1 - (0.0 - 0.0))))
                    _ = v24
                }
                p0 = ("hello" + "hello")
            }
        case "foo bar":
            v25 := GenS2{F0: 4294967295, F1: -1}
            _ = v25
        }
    case "hello":
        v26 := GenS4{F0: 2147483648, F1: -2147483648}
        _ = v26
    case "\n\t":
        if ((((42 == 2147483647) || (false && (0 == 42))) || (42 >= 2)) || true) {
            v27 := GenS2{F0: 65535, F1: 1}
            p0 = "foo bar"
            switch 821 {
            case 42:
                var v28 bool = (1 >= 127)
                _ = v28
            default:
                var v29 uint16 = 0
                _ = v29
            }
            _ = v27
        }
    default:
        if (true && false) {
            switch p0 {
            case "hello":
                v30 := GenS5{F0: 73}
                _ = v30
            case "\n\t":
                for fi31 := 0; fi31 < 6; fi31++ {
                    if true {
                        v32 := GenS2{F0: 2, F1: -1}
                        v33 := []int{(fi31 & fi31), (^(fi31))}
                        _ = v32
                        _ = v33
                    } else {
                        switch fi31 {
                        case -9223372036854775808:
                            v34 := []int{fi31, fi31, 762, fi31}
                            _ = v34
                        case 4294967295:
                            v35 := []int{-759, fi31, (fi31 >> ((2) & 63)), (fi31 + fi31)}
                            _ = v35
                        default:
                            fmt.Println(p0)
                        }
                    }
                    fmt.Println(p0, fi31)
                    switch fi31 {
                    case 9223372036854775807:
                        var v36 int32 = -510
                        _ = v36
                    }
                    _ = fi31
                }
            default:
                fmt.Println(p0)
            }
            fmt.Println(p0)
        } else {
            if false {
                p0 = (("a" + p0) + p0)
            } else {
                p0 = p0
            }
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 float64 = -2.25
                var v39 uint = 18446744073709551615
                var v40 int64 = 9223372036854775807
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
            var v41 int32 = -386
            _ = v41
        }
    }
    _ = p0
    return ((!true) || true)
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 2
    v3 := []int{}
    v2 = 18446744073709551615
    v1[0] = (true || false)
    if (!true) {
        if false {
            fmt.Println(v2)
        }
        fmt.Println(v2)
    } else {
        var v42 uint = 0
        _ = v42
    }
    v2 -= (v2 * (v2 | v2))
    v2 = 0
    var v43 uint32 = 23
    switch "" {
    case "":
        var v44 string = "foo bar"
        _ = v44
    }
    switch -1 {
    case 32767:
        var v45 int32 = 42
        _ = v45
    case 2147483647:
        var v46 bool = true
        _ = v46
    case -2147483648:
        if true {
            switch "foo bar" {
            case "\n\t":
                if (!true) {
                    v47 := GenS6{F0: 21, F1: 9223372036854775807, F2: 1}
                    switch "foo bar" {
                    case "":
                        var v48 uint32 = (v43 & (v43 << (((v2 - v2)) & 63)))
                        _ = v48
                    case "a":
                        v49 := GenS5{F0: 42}
                        _ = v49
                    case "foo bar":
                        v2 = 4294967295
                    }
                    v43 = v43
                    _ = v47
                } else {
                    for fi50 := 0; fi50 < 10; fi50++ {
                        for fi51 := 0; fi51 < 6; fi51++ {
                            if true {
                                var v52 uint8 = 0
                                if (!false) {
                                    var v53 uint32 = v43
                                    _ = v53
                                } else {
                                    var v54 uint = v2
                                    v2 += (v2 % ((0) | 1))
                                    _ = v54
                                }
                                var v55 uint64 = 683
                                _ = v52
                                _ = v55
                            }
                            _ = fi51
                        }
                        _ = fi50
                    }
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
}
