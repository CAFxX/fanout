package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint32; F1 bool; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 string; F2 int }
type GenS5 struct { F0 uint8; F1 uint; F2 uint64 }
func genFn0(p0 uint16) uint64 {
    if (-2147483648 < 1) {
        v7 := []int{2147483647}
        var v8 float32 = (1.5 - 1e10)
        switch ("hello" + "a") {
        case "hello":
            for fi9 := 0; fi9 < 3; fi9++ {
                fmt.Println(v8, v8, p0)
                switch fi9 {
                case 42:
                    var v10 float32 = float32(-2.25)
                    _ = v10
                case 32767:
                    if true {
                        for fi11 := 0; fi11 < 12; fi11++ {
                            switch (fi9 / ((fi9) | 1)) {
                            case 32767:
                                fmt.Println(v8, fi9)
                            default:
                                var v12 uint64 = 18446744073709551615
                                _ = v12
                            }
                            v13 := GenS1{F0: 201, F1: 32767}
                            _ = fi11
                            _ = v13
                        }
                    } else {
                        var v14 int64 = -1
                        _ = v14
                    }
                case 4294967295:
                    for fi15 := 0; fi15 < 14; fi15++ {
                        switch (("" + ("a" + "hello")) + "hello") {
                        case "":
                            p0 += 345
                        case "a":
                            var v16 float32 = v8
                            _ = v16
                        case "\n\t":
                            if len(v7) > 0 {
                                v7[0] = (-(fi9))
                            }
                        default:
                            if ((p0 + (p0 % ((p0) | 1))) != (p0 * p0)) {
                                switch "a" {
                                case "foo bar":
                                    v17 := GenS1{F0: 28, F1: 240}
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    var v18 int32 = 2147483647
                                    _ = v17
                                    _ = v18
                                case "":
                                    var v19 uint8 = 2
                                    _ = v19
                                case "\n\t":
                                    v20 := GenS2{F0: 0, F1: true, F2: ""}
                                    _ = v20
                                }
                                v21 := GenS3{F0: -1}
                                _ = v21
                            } else {
                                v22 := GenS2{F0: 4294967295, F1: true, F2: "foo bar"}
                                v23 := GenS4{F0: 0, F1: ("hello" + "hello"), F2: fi15}
                                fmt.Println(v8, v8, v8)
                                _ = v22
                                _ = v23
                            }
                        }
                        var v24 int = fi15
                        _ = fi15
                        _ = v24
                    }
                default:
                    fmt.Println(v8)
                }
                switch (("" + "a") + "a") {
                case "\n\t":
                    fmt.Println(v8, fi9)
                case "hello":
                    var v25 float64 = 1e10
                    _ = v25
                case "":
                    v26 := GenS4{F0: 0, F1: ("" + "\n\t"), F2: 32767}
                    _ = v26
                default:
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v7 = append(v7, fi27)
                        _ = fi27
                    }
                }
                _ = fi9
            }
        case "\n\t":
            switch "a" {
            case "hello":
                v8 = v8
            default:
                p0 = p0
            }
        case "a":
            switch -416 {
            case 751:
                if len(v7) > 0 {
                    v7[0] = -1
                }
            case 633:
                v28 := GenS4{F0: 1, F1: "", F2: 32768}
                _ = v28
            }
        default:
            var v29 float64 = 3.14159
            _ = v29
        }
        _ = v7
        _ = v8
    }
    var v30 bool = false
    if v30 {
        for fi31 := 0; fi31 < 17; fi31++ {
            v32 := []int{(fi31 + fi31), fi31}
            var v33 int32 = -2147483648
            v33 = v33
            _ = fi31
            _ = v32
            _ = v33
        }
        fmt.Println(v30, v30)
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 uint32 = 0
            v36 := GenS4{F0: -32768, F1: "hello", F2: fi34}
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    v37 := []int{32768, 2147483647, 128, 203}
    _ = p0
    _ = v30
    _ = v37
    return 582
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    var v3 string = "hello"
    v4 := make(map[uint32]bool)
    var v5 uint32 = 166
    v6 := make(map[int]bool)
    var v38 float32 = 1.5
    switch "\n\t" {
    case "a":
        v3 = v3
    default:
        v2[4294967295] = (v5 & (^(v5)))
    }
    v39 := GenS2{F0: v5, F1: (true || (2 > v5)), F2: v3}
    for fi40 := 0; fi40 < 16; fi40++ {
        v41 := GenS5{F0: 2, F1: 18446744073709551615, F2: 2}
        var v42 uint16 = 816
        _ = fi40
        _ = v41
        _ = v42
    }
    var v43 int32 = -293
    for fi44 := 0; fi44 < 19; fi44++ {
        v43 = 958
        v45 := []int{(fi44 ^ fi44)}
        v46 := GenS1{F0: 42, F1: int(v5)}
        _ = fi44
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 6; fi47++ {
        var v48 uint = 18446744073709551615
        _ = fi47
        _ = v48
    }
    var v49 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v49
    _ = v5
    _ = v6
}
