package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 string; F2 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16 }
func genFn0(p0 bool, p1 int, p2 uint64) bool {
    var v7 uint = 42
    v8 := []int{}
    v9 := GenS1{F0: 126, F1: "a", F2: 42}
    v7 = (^(v7))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (0 < 1)
}
func genFn1(p0 int) int32 {
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(p0)
        for fi11 := 0; fi11 < 1; fi11++ {
            switch p0 {
            case 127:
                v12 := GenS1{F0: 224, F1: "hello", F2: 32767}
                _ = v12
            default:
                if false {
                    fmt.Println(p0, p0)
                    var v13 uint64 = 42
                    _ = v13
                } else {
                    if true {
                        if false {
                            v14 := []int{fi10, p0, (fi10 - (fi10 * 2147483647)), p0}
                            switch fi11 {
                            case -9223372036854775808:
                                if len(v14) > 0 {
                                    v14[0] = (fi11 ^ -9223372036854775808)
                                }
                            case 713:
                                var v15 uint32 = 4294967295
                                _ = v15
                            default:
                                switch -771 {
                                case -864:
                                    if ("\n\t" != "a") {
                                        if false {
                                            _ = len(v14)
                                            _ = len(v14)
                                            var v16 int8 = -128
                                            fmt.Println(fi11, fi11, fi10)
                                            _ = v16
                                        }
                                        switch (-(fi11)) {
                                        case 9223372036854775807:
                                            v17 := GenS1{F0: 255, F1: "", F2: -1000}
                                            _ = v17
                                        case 628:
                                            if (202 == 18446744073709551615) {
                                                if true {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v18 uint8 = 42
                                                        _ = v18
                                                    case "":
                                                        if len(v14) >= 2 {
                                                            v14 = v14[1:len(v14)]
                                                        }
                                                    case "\n\t":
                                                        var v19 int64 = 42
                                                        _ = v19
                                                    default:
                                                        var v20 float64 = (1e10 - 3.14159)
                                                        _ = v20
                                                    }
                                                    var v21 uint32 = 153
                                                    v21 = (^(v21))
                                                    _ = len(v14)
                                                    v21 = (v21 | v21)
                                                    _ = v21
                                                } else {
                                                    _ = len(v14)
                                                    p0 = 42
                                                    var v22 uint16 = 42
                                                    v23 := GenS1{F0: 56, F1: "hello", F2: -80}
                                                    _ = v22
                                                    _ = v23
                                                }
                                                var v24 int8 = 127
                                                _ = v24
                                            }
                                        }
                                        p0 = (^(fi10))
                                    } else {
                                        p0 += (fi11 + fi11)
                                        p0 -= p0
                                        if true {
                                            switch p0 {
                                            case 127:
                                                v25 := GenS2{F0: -675, F1: (false || ((1.5 + 1.5) > -2.25))}
                                                _ = v25
                                            case 655:
                                                if len(v14) > 0 {
                                                    v14[0] = fi11
                                                }
                                            }
                                            var v26 string = "hello"
                                            p0 = fi11
                                            _ = v26
                                        }
                                    }
                                case -1:
                                    p0 = (-(fi10))
                                case -499:
                                    var v27 bool = false
                                    _ = v27
                                }
                            }
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            var v28 float64 = (float64(((1.5 / float32(3.14159)) * 0.1)) - (float64((-2.25 - (0.1 / 1e10))) / 1e10))
                            var v29 uint16 = 238
                            switch "hello" {
                            case "":
                                v29 += (v29 / (((^(v29))) | 1))
                            case "hello":
                                _ = len(v14)
                            case "\n\t":
                                v28 = ((((-2.25 * (-2.25 + v28)) - -2.25) + ((v28 / -2.25) / 1e10)) - (v28 / v28))
                            default:
                                var v30 float32 = 1.5
                                _ = v30
                            }
                            var v31 uint16 = (v29 - 335)
                            v32 := GenS3{F0: 54}
                            v33 := GenS1{F0: 127, F1: "\n\t", F2: 32767}
                            v34 := GenS4{F0: 460, F1: "a", F2: ("hello" >= "a")}
                            var v35 int32 = -1
                            _ = v14
                            _ = v28
                            _ = v29
                            _ = v31
                            _ = v32
                            _ = v33
                            _ = v34
                            _ = v35
                        } else {
                            v36 := GenS1{F0: 153, F1: "\n\t", F2: -221}
                            var v37 uint64 = 435
                            var v38 uint64 = v37
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                    }
                    p0 = fi11
                    fmt.Println(p0, fi10, fi11)
                }
            }
            var v39 uint16 = 777
            if (false && false) {
                v40 := []int{p0, int(v39), fi11}
                _ = v40
            }
            _ = fi11
            _ = v39
        }
        _ = fi10
    }
    if (!false) {
        switch (-(p0)) {
        case -36:
            if false {
                p0 = 549
                fmt.Println(p0)
            }
        default:
            for fi41 := 0; fi41 < 10; fi41++ {
                var v42 int = 127
                if true {
                    for fi43 := 0; fi43 < 7; fi43++ {
                        var v44 uint32 = 223
                        v45 := []int{9223372036854775807, -936, 729}
                        switch -457 {
                        case 127:
                            fmt.Println(fi43, v44)
                            if false {
                                _ = len(v45)
                            }
                            v46 := GenS1{F0: 96, F1: (("hello" + "\n\t") + ((((("\n\t" + "foo bar") + "hello") + "a") + "\n\t") + (("hello" + (("" + "") + (("foo bar" + (("hello" + ("\n\t" + (("a" + "a") + ""))) + "foo bar")) + ((("hello" + "a") + (("" + "hello") + "a")) + (("foo bar" + "foo bar") + ""))))) + "\n\t"))), F2: -32768}
                            _ = v46
                        }
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                } else {
                    v47 := GenS5{F0: p0}
                    var v48 int = (^(v42))
                    _ = v47
                    _ = v48
                }
                _ = fi41
                _ = v42
            }
        }
    } else {
        fmt.Println(p0)
        p0 = p0
        var v49 int32 = 0
        _ = v49
    }
    var v50 uint64 = 39
    _ = p0
    _ = v50
    return 679
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 128
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    var v5 uint64 = 18446744073709551615
    v6 := make(map[int]string)
    v5 = v5
    var v51 int16 = -1
    var v52 uint8 = 128
    for fi53 := 0; fi53 < 16; fi53++ {
        v54 := GenS3{F0: 113}
        for fi55 := 0; fi55 < 19; fi55++ {
            var v56 int8 = 1
            v51 -= (v51 ^ (v51 * int16(v52)))
            _ = fi55
            _ = v56
        }
        _ = fi53
        _ = v54
    }
    for fi57 := 0; fi57 < 11; fi57++ {
        fmt.Println(fi57, v5, v2)
        _ = len(v1)
        var v58 string = "hello"
        _ = fi57
        _ = v58
    }
    switch "a" {
    case "hello":
        for fi59 := 0; fi59 < 7; fi59++ {
            var v60 uint64 = 2
            _ = fi59
            _ = v60
        }
    }
    fmt.Println(v5)
    if (-1 > -561) {
        v61 := GenS6{F0: 42}
        var v62 int = 543
        _ = v61
        _ = v62
    }
    if false {
        fmt.Println(v52, v5, v2)
        v63 := GenS7{F0: 0}
        _ = v63
    } else {
        fmt.Println(v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
