package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint8, p1 int32, p2 int16) uint64 {
    v5 := []int{42, 127, int(p0)}
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: true}
        fmt.Println(p1, fi6)
        p0 -= p0
        _ = fi6
        _ = v7
    }
    p0 = p0
    if false {
        v8 := GenS2{F0: true, F1: 65535}
        _ = v8
    }
    p2 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    var v1 int = 750
    var v2 string = ((("foo bar" + "foo bar") + "hello") + "")
    v3 := []string{(v2 + v2), "", (((v2 + ("hello" + v2)) + (v2 + (v2 + "hello"))) + "foo bar")}
    v4 := make(map[string]bool)
    var v9 int8 = -107
    _ = len(v4)
    if (!(!((v9 | (v9 % (((v9 >> ((1) & 63))) | 1))) != -128))) {
        _ = len(v3)
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v2)
            if ("foo bar" >= v2) {
                var v11 uint8 = 128
                if len(v3) > 0 {
                    v3[0] = ""
                }
                var v12 float64 = 1e10
                _ = v11
                _ = v12
            } else {
                switch v2 {
                case "":
                    v1 = 32768
                case "\n\t":
                    var v13 float32 = 3.14159
                    _ = v13
                case "hello":
                    fmt.Println(v2, v2)
                default:
                    v9 = 42
                }
                var v14 bool = true
                _ = v14
            }
            _ = fi10
        }
        fmt.Println(v9, v1, v1)
    } else {
        var v15 uint32 = 1
        _ = v15
    }
    v9 = (^(v9))
    _ = len(v4)
    v16 := GenS1{F0: (true || (4294967295 <= 741))}
    fmt.Println(v9)
    for fi17 := 0; fi17 < 10; fi17++ {
        v1 = (fi17 * 42)
        switch 2147483647 {
        case 491:
            switch v1 {
            case 285:
                v18 := GenS1{F0: true}
                _ = v18
            case -9223372036854775808:
                v1 = fi17
            default:
                var v19 string = (v2 + ((v2 + v2) + ((v2 + (v2 + v2)) + ("\n\t" + v2))))
                _ = v19
            }
        case -399:
            if ((("hello" + (v2 + "")) + (v2 + "a")) <= "foo bar") {
                fmt.Println(fi17, fi17, fi17)
                v1 += (v1 & (^(fi17)))
            }
        case -2147483648:
            switch v2 {
            case "\n\t":
                v9 = v9
            case "hello":
                v2 = "a"
            case "foo bar":
                if (887 >= 782) {
                    _ = len(v3)
                    fmt.Println(v2)
                }
            }
        default:
            var v20 float32 = float32((1.5 / (float64(1.5) - 0.0)))
            _ = v20
        }
        _ = fi17
    }
    switch "hello" {
    case "a":
        var v21 uint8 = 0
        _ = v21
    case "hello":
        for fi22 := 0; fi22 < 15; fi22++ {
            switch v1 {
            case -9223372036854775808:
                if true {
                    switch "" {
                    case "":
                        var v23 uint64 = 2
                        _ = v23
                    default:
                        v24 := GenS3{F0: 1000, F1: v2}
                        _ = v24
                    }
                    v25 := GenS3{F0: int16(v1), F1: v2}
                    _ = v25
                } else {
                    fmt.Println(fi22)
                    v3 = append(v3, ("foo bar" + (((v2 + "\n\t") + "foo bar") + v2)))
                    if true {
                        v26 := GenS1{F0: (722 <= 1)}
                        _ = v26
                    } else {
                        fmt.Println(v9, v2)
                        switch (v1 / (((v1 / (((-(fi22))) | 1))) | 1)) {
                        case -2147483648:
                            var v27 int16 = 42
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 9; fi28++ {
                            v29 := GenS3{F0: 0, F1: "\n\t"}
                            _ = fi28
                            _ = v29
                        }
                    }
                }
            case 819:
                v3 = append(v3, v2)
            case -2147483648:
                v1 = 42
            }
            switch "a" {
            case "a":
                if (((0.0 - (((1.5 / 1.5) - 0.0) * float32(-2.25))) >= 0.1) || (!true)) {
                    v30 := GenS2{F0: true, F1: 4294967295}
                    for fi31 := 0; fi31 < 20; fi31++ {
                        if false {
                            v32 := GenS4{F0: fi31}
                            var v33 string = v2
                            fmt.Println(v1, v2, fi22)
                            fmt.Println(fi22, v33, v2)
                            _ = v32
                            _ = v33
                        } else {
                            v9 -= v9
                            v34 := GenS1{F0: false}
                            _ = v34
                        }
                        v2 = "hello"
                        _ = fi31
                    }
                    _ = v30
                }
            case "\n\t":
                v2 = (v2 + "a")
            }
            if true {
                var v35 bool = false
                _ = v35
            }
            _ = fi22
        }
    case "":
        var v36 uint16 = 624
        _ = v36
    default:
        var v37 uint = 682
        _ = v37
    }
    if ((!(true || true)) && false) {
        switch (-(v1)) {
        case -9223372036854775808:
            var v38 int64 = -9223372036854775808
            _ = v38
        case 854:
            v2 = (((v2 + "foo bar") + ((v2 + (v2 + ((v2 + ("\n\t" + v2)) + "hello"))) + v2)) + "hello")
        case 1:
            v9 = v9
        default:
            if true {
                switch v1 {
                case 576:
                    fmt.Println(v9, v9, v9)
                case 4294967295:
                    v39 := GenS5{F0: 298, F1: 255}
                    _ = v39
                }
                v40 := GenS3{F0: 530, F1: ("hello" + v2)}
                _ = v40
            }
        }
        v2 = v2
        var v41 bool = true
        _ = v41
    } else {
        if (!false) {
            if len(v3) > 0 {
                v3[0] = (v2 + v2)
            }
            _ = len(v3)
        }
    }
    v42 := GenS6{F0: 0}
    var v43 uint = 136
    v2 = (("\n\t" + ("hello" + (v2 + "foo bar"))) + v2)
    var v44 uint64 = 0
    var v45 int8 = v9
    if len(v3) > 0 {
        v3[0] = v2
    }
    var v46 int = v1
    v47 := GenS3{F0: 32767, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v9
}
