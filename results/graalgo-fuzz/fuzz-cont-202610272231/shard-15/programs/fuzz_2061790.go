package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint; F1 int16; F2 uint32 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 bool; F1 uint64; F2 uint16 }
func genFn0(p0 uint8) int {
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float32 = -2.25
        switch "a" {
        case "":
            switch "a" {
            case "foo bar":
                if ((((1.5 - (0.0 * -2.25)) / (((1e10 - (-2.25 / 0.1)) + 0.1) - 1e10)) - 0.1) != (-2.25 / 1e10)) {
                    p0 = p0
                    switch "foo bar" {
                    case "foo bar":
                        var v9 int = 845
                        _ = v9
                    default:
                        v10 := []int{fi7}
                        _ = v10
                    }
                }
            case "a":
                var v11 int64 = -9223372036854775808
                _ = v11
            }
        case "\n\t":
            if (42 == -71) {
                v12 := []int{(fi7 - fi7)}
                _ = v12
            }
        default:
            if (32767 != -435) {
                fmt.Println(fi7, p0, fi7)
            } else {
                var v13 int32 = -42
                for fi14 := 0; fi14 < 5; fi14++ {
                    if true {
                        var v15 float32 = v8
                        _ = v15
                    } else {
                        if true {
                            v16 := []int{(-(fi14)), 128}
                            _ = v16
                        }
                    }
                    _ = fi14
                }
                _ = v13
            }
        }
        if true {
            v8 += (v8 / 0.1)
        } else {
            switch ("hello" + "hello") {
            case "hello":
                v8 = v8
            case "\n\t":
                for fi17 := 0; fi17 < 2; fi17++ {
                    v18 := GenS1{F0: ("hello" + ("hello" + "foo bar")), F1: ""}
                    var v19 string = ""
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case "":
                v20 := GenS1{F0: "", F1: "hello"}
                _ = v20
            default:
                var v21 int = (fi7 * 2147483647)
                _ = v21
            }
            for fi22 := 0; fi22 < 19; fi22++ {
                switch "a" {
                case "":
                    v23 := GenS2{F0: 0, F1: 42, F2: 23}
                    _ = v23
                default:
                    var v24 int = 395
                    _ = v24
                }
                _ = fi22
            }
        }
        _ = fi7
        _ = v8
    }
    var v25 uint64 = 2
    v26 := []int{42, 1, 128, -841}
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 int32 = 2147483647
        for fi29 := 0; fi29 < 6; fi29++ {
            v30 := GenS3{F0: fi27, F1: 0}
            _ = fi29
            _ = v30
        }
        v25 -= (v25 & v25)
        _ = fi27
        _ = v28
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        if (9223372036854775807 == -1) {
            var v32 uint32 = 581
            switch (^(fi31)) {
            case -9223372036854775808:
                v32 += 676
            case 2147483647:
                _ = len(v26)
            }
            for fi33 := 0; fi33 < 7; fi33++ {
                v34 := GenS2{F0: 4294967295, F1: -514, F2: v32}
                _ = fi33
                _ = v34
            }
            _ = v32
        }
        for fi35 := 0; fi35 < 18; fi35++ {
            switch 4294967295 {
            case 128:
                p0 = (p0 % (((^(p0))) | 1))
                var v36 int = (^(fi31))
                v36 -= -9223372036854775808
                _ = v36
            case 2147483647:
                p0 += (p0 % (((p0 % (((^(p0))) | 1))) | 1))
            default:
                fmt.Println(v25, fi31, p0)
            }
            _ = fi35
        }
        p0 = (p0 >> ((65535) & 63))
        _ = fi31
    }
    _ = p0
    _ = v25
    _ = v26
    return -9223372036854775808
}

func main() {
    var v1 bool = false
    v2 := make(map[int]bool)
    var v3 int32 = -1
    var v4 uint = 42
    var v5 int64 = 0
    v6 := []uint32{}
    v3 = (v3 | (v3 + v3))
    switch "\n\t" {
    case "a":
        if true {
            fmt.Println(v4, v1)
            if v1 {
                v5 = 9223372036854775807
            }
        }
    case "foo bar":
        var v37 int8 = 0
        _ = v37
    case "\n\t":
        var v38 bool = (428 <= 545)
        _ = v38
    }
    switch ("\n\t" + "a") {
    case "foo bar":
        for fi39 := 0; fi39 < 15; fi39++ {
            v5 = (v5 >> ((821) & 63))
            _ = fi39
        }
    case "":
        var v40 string = ""
        _ = v40
    case "hello":
        var v41 int16 = 1
        _ = v41
    }
    switch "foo bar" {
    case "hello":
        v5 -= (^(v5))
    case "foo bar":
        v5 = (v5 / (((v5 + v5)) | 1))
    case "\n\t":
        fmt.Println(v4, v5)
    default:
        var v42 uint32 = 42
        _ = v42
    }
    var v43 int8 = -128
    fmt.Println(v3, v43)
    v3 = v3
    v44 := GenS4{F0: v1, F1: 2, F2: 42}
    var v45 int8 = (v43 - v43)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
