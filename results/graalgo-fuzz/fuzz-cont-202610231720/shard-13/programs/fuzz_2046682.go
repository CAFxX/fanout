package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 bool }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint64, p1 int32, p2 int64) int16 {
    if (4294967295 < 427) {
        v4 := GenS1{F0: 42, F1: 42, F2: false}
        _ = v4
    }
    if (false || ((4294967295 > 728) && false)) {
        if false {
            var v5 int = 1
            switch "" {
            case "hello":
                var v6 string = "hello"
                _ = v6
            case "\n\t":
                v7 := GenS1{F0: 4294967295, F1: (p1 + (p1 + p1)), F2: true}
                _ = v7
            case "a":
                var v8 uint8 = 93
                _ = v8
            }
            if true {
                switch ("a" + "\n\t") {
                case "hello":
                    if true {
                        v5 = v5
                        if ((("a" + ("" + "\n\t")) + "") < "") {
                            var v9 string = "a"
                            v10 := GenS1{F0: 65535, F1: (p1 * 463), F2: false}
                            _ = v10
                            _ = v9
                        } else {
                            v11 := GenS2{F0: p1, F1: true}
                            var v12 int8 = 100
                            var v13 float32 = (3.14159 * 1e10)
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        p0 -= p0
                    } else {
                        switch v5 {
                        case -1:
                            switch "a" {
                            case "":
                                p0 = 327
                            case "hello":
                                v14 := []int{(v5 ^ int(p0))}
                                _ = v14
                            default:
                                var v15 int64 = int64(p1)
                                _ = v15
                            }
                        }
                        if (32767 > 0) {
                            p1 = p1
                        } else {
                            var v16 string = "\n\t"
                            _ = v16
                        }
                    }
                case "":
                    p2 += p2
                }
                v17 := GenS2{F0: p1, F1: true}
                if (p2 > (p2 >> ((42) & 63))) {
                    p2 = (-(p2))
                    p0 = (p0 % (((p0 ^ p0)) | 1))
                } else {
                    v18 := []int{}
                    v19 := GenS3{F0: 0, F1: -128}
                    _ = v18
                    _ = v19
                }
                _ = v17
            } else {
                v20 := GenS3{F0: 639, F1: 42}
                var v21 uint8 = 191
                _ = v20
                _ = v21
            }
            _ = v5
        }
    }
    var v22 uint = 552
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 32767
}

func main() {
    var v1 uint8 = 255
    var v2 string = ""
    var v3 int64 = 42
    if (-23 != 100) {
        if (!false) {
            v23 := GenS4{F0: v2, F1: 702}
            if false {
                v1 -= (v1 & 124)
                v24 := GenS3{F0: 571, F1: -35}
                _ = v24
            }
            _ = v23
        }
        for fi25 := 0; fi25 < 6; fi25++ {
            switch fi25 {
            case -9223372036854775808:
                var v26 uint = 4294967295
                _ = v26
            }
            _ = fi25
        }
        switch v2 {
        case "\n\t":
            var v27 uint32 = 843
            _ = v27
        case "a":
            var v28 int = 2147483647
            _ = v28
        }
    } else {
        fmt.Println(v3, v1)
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        if true {
            v30 := GenS5{F0: genFn0(1, -655, v3)}
            fmt.Println(v3, fi29)
            _ = v30
        }
        for fi31 := 0; fi31 < 11; fi31++ {
            v2 = v2
            if (-1 <= -2147483648) {
                var v32 float32 = 1.5
                v33 := []int{fi31}
                var v34 uint = 721
                _ = v32
                _ = v33
                _ = v34
            }
            _ = fi31
        }
        _ = fi29
    }
    v35 := []int{-2147483648, 127, -348}
    for fi36 := 0; fi36 < 6; fi36++ {
        v3 += (v3 - v3)
        var v37 int64 = (v3 + v3)
        _ = fi36
        _ = v37
    }
    if false {
        fmt.Println(v3, v2)
        for fi38 := 0; fi38 < 18; fi38++ {
            v39 := GenS2{F0: -159, F1: false}
            _ = fi38
            _ = v39
        }
        if true {
            v40 := GenS6{F0: 581, F1: 817}
            _ = v40
        }
    }
    if len(v35) >= 2 {
        v35 = v35[1:len(v35)]
    }
    if len(v35) > 0 {
        v35[0] = 9223372036854775807
    }
    switch v2 {
    case "a":
        for fi41 := 0; fi41 < 9; fi41++ {
            for fi42 := 0; fi42 < 13; fi42++ {
                if false {
                    if len(v35) >= 2 {
                        v35 = v35[1:len(v35)]
                    }
                    fmt.Println(v3)
                }
                v2 = "hello"
                if (true || true) {
                    v1 = v1
                } else {
                    _ = len(v35)
                }
                v2 = ((v2 + "foo bar") + (((v2 + ((("foo bar" + v2) + ((v2 + (v2 + ("foo bar" + v2))) + (v2 + "hello"))) + (((v2 + v2) + v2) + (("" + v2) + ("" + "\n\t"))))) + "foo bar") + "hello"))
                _ = fi42
            }
            _ = fi41
        }
    }
    v2 = v2
    if len(v35) > 0 {
        v35[0] = 0
    }
    v2 = (v2 + v2)
    var v43 uint = 401
    fmt.Println(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v43
}
