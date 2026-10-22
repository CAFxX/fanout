package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 bool }
func genFn0(p0 int64) uint32 {
    fmt.Println(p0)
    fmt.Println(p0)
    v5 := GenS1{F0: -595}
    _ = p0
    _ = v5
    return 65535
}
func genFn1(p0 int, p1 int32) uint {
    var v6 float64 = 3.14159
    if (!(!false)) {
        fmt.Println(p0, v6)
    }
    var v7 uint32 = 2
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(v7, fi8)
        v9 := []int{(p0 | 32767), (p0 & fi8), fi8}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := []string{("foo bar" + "\n\t"), (("" + (("\n\t" + "hello") + (((((((("\n\t" + "hello") + "hello") + "a") + (((((("\n\t" + "\n\t") + "\n\t") + "a") + ("foo bar" + ("a" + "\n\t"))) + "a") + ("" + "hello"))) + "a") + "hello") + "") + "a"))) + "\n\t")}
    var v2 uint32 = 4294967295
    v3 := make(map[uint32]string)
    var v4 int16 = 927
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v10 := GenS1{F0: -1}
    if true {
        switch ("a" + "") {
        case "a":
            var v11 int64 = 0
            _ = v11
        }
        if false {
            v2 = v2
        } else {
            switch "foo bar" {
            case "a":
                if (!(1 <= -2147483648)) {
                    for fi12 := 0; fi12 < 17; fi12++ {
                        v4 = 42
                        var v13 uint64 = 18446744073709551615
                        v4 = 34
                        _ = fi12
                        _ = v13
                    }
                }
            case "":
                v2 = 4294967295
            case "\n\t":
                switch "a" {
                case "a":
                    var v14 int32 = 2147483647
                    _ = v14
                case "\n\t":
                    switch "foo bar" {
                    case "foo bar":
                        v15 := GenS1{F0: 9223372036854775807}
                        _ = v15
                    case "a":
                        switch "hello" {
                        case "":
                            var v16 uint32 = (v2 * v2)
                            _ = v16
                        }
                    case "":
                        fmt.Println(v2)
                    }
                case "foo bar":
                    var v17 int32 = -1
                    _ = v17
                default:
                    v3[v2] = ""
                }
            default:
                for fi18 := 0; fi18 < 12; fi18++ {
                    if (65535 > 837) {
                        v19 := GenS2{F0: 1, F1: 383, F2: 1}
                        var v20 uint8 = 255
                        _ = v19
                        _ = v20
                    }
                    v21 := GenS2{F0: 356, F1: 4294967295, F2: 18446744073709551615}
                    _ = len(v1)
                    _ = fi18
                    _ = v21
                }
            }
        }
    } else {
        var v22 uint16 = 107
        _ = v22
    }
    switch -1 {
    case -668:
        switch 128 {
        case 438:
            if len(v1) > 0 {
                v1[0] = "a"
            }
        case 127:
            v23 := GenS1{F0: 1}
            _ = v23
        }
    case -646:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -1:
        fmt.Println(v2, v4)
    default:
        for fi24 := 0; fi24 < 7; fi24++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v25 int32 = -704
            _ = fi24
            _ = v25
        }
    }
    if (-1 != -735) {
        v4 = (v4 * v4)
    }
    v4 = v4
    v2 = (v2 & v2)
    var v26 int = -290
    v27 := GenS2{F0: 563, F1: 42, F2: 758}
    var v28 bool = true
    if (v28 && true) {
        fmt.Println(v26, v26, v28)
        v29 := GenS3{F0: 127, F1: 18446744073709551615}
        for fi30 := 0; fi30 < 4; fi30++ {
            v1 = append(v1, "a")
            _ = fi30
        }
        _ = v29
    } else {
        for fi31 := 0; fi31 < 2; fi31++ {
            var v32 uint16 = 32767
            if ("foo bar" < ("\n\t" + ("\n\t" + "a"))) {
                v33 := GenS4{F0: (42 > int16(v2))}
                v32 = 972
                var v34 string = "\n\t"
                _ = v33
                _ = v34
            } else {
                v35 := GenS1{F0: 42}
                _ = v35
            }
            _ = fi31
            _ = v32
        }
        var v36 uint = 1
        _ = v36
    }
    v26 -= v26
    switch v26 {
    case 0:
        var v37 float64 = 0.1
        _ = v37
    case -9223372036854775808:
        var v38 bool = (v2 == v2)
        _ = v38
    }
    v4 += 1000
    v28 = (285 > 42)
    if true {
        v28 = true
    } else {
        switch "foo bar" {
        case "":
            for fi39 := 0; fi39 < 9; fi39++ {
                v28 = v28
                if v28 {
                    v2 = 477
                    v40 := GenS3{F0: 42, F1: 527}
                    v26 = (-(fi39))
                    _ = v40
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                for fi41 := 0; fi41 < 1; fi41++ {
                    var v42 uint32 = (v2 | 65535)
                    if len(v1) > 0 {
                        v1[0] = ("hello" + ("hello" + "hello"))
                    }
                    var v43 uint64 = 42
                    v2 -= genFn0(-9223372036854775808)
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = fi39
            }
        case "\n\t":
            if (true && v28) {
                fmt.Println(v28, v2, v28)
            } else {
                if v28 {
                    v28 = true
                }
                var v44 uint8 = 0
                fmt.Println(v2, v26)
                _ = v44
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
