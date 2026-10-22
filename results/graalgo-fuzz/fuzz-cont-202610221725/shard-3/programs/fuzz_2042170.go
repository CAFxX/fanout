package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int32; F1 int8; F2 int8 }
type GenS3 struct { F0 int; F1 uint64; F2 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int; F1 uint8 }
type GenS10 struct { F0 int32; F1 int; F2 uint16 }

func main() {
    v1 := []string{"", "", (("a" + "foo bar") + ("hello" + "hello")), ("a" + (("hello" + "") + "hello"))}
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    var v5 int8 = 127
    v6 := make(map[uint32]int)
    delete(v3, -2147483648)
    v7 := GenS1{F0: 350, F1: 42}
    v5 = v5
    var v8 string = ("\n\t" + "hello")
    v1 = append(v1, "hello")
    var v9 float32 = 1e10
    if true {
        for fi10 := 0; fi10 < 1; fi10++ {
            switch v8 {
            case "\n\t":
                fmt.Println(fi10)
            default:
                v6[422] = (fi10 * fi10)
            }
            v11 := GenS2{F0: int32(v5), F1: (v5 - -128), F2: v5}
            _ = fi10
            _ = v11
        }
        var v12 int = 0
        _ = v12
    } else {
        var v13 uint64 = 42
        fmt.Println(v8, v8, v9)
        for fi14 := 0; fi14 < 13; fi14++ {
            switch 32768 {
            case 0:
                if ("" != (((v8 + ("" + v8)) + "hello") + ((v8 + ((v8 + v8) + (v8 + v8))) + v8))) {
                    v13 -= (v13 * (v13 & v13))
                } else {
                    if (42 < (v13 ^ 4294967295)) {
                        v15 := GenS3{F0: (fi14 & (^(fi14))), F1: 4294967295, F2: 1}
                        var v16 uint64 = 4294967295
                        _ = v15
                        _ = v16
                    } else {
                        v5 += -128
                        var v17 uint64 = 0
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 13; fi18++ {
                        switch fi14 {
                        case -858:
                            v19 := GenS4{F0: 1}
                            _ = v19
                        }
                        _ = fi18
                    }
                }
            case -9223372036854775808:
                for fi20 := 0; fi20 < 15; fi20++ {
                    if true {
                        var v21 int32 = 1
                        _ = v21
                    }
                    if (!false) {
                        v8 = v8
                        v22 := GenS3{F0: (^(fi14)), F1: 4294967295, F2: -1}
                        _ = v22
                    }
                    _ = len(v1)
                    _ = fi20
                }
            case 1:
                v23 := GenS2{F0: int32(v13), F1: -126, F2: (v5 | (v5 << ((18446744073709551615) & 63)))}
                _ = v23
            default:
                v8 = (v8 + "a")
            }
            switch (fi14 / ((fi14) | 1)) {
            case 4294967295:
                switch v8 {
                case "\n\t":
                    for fi24 := 0; fi24 < 14; fi24++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch v8 {
                        case "a":
                            if false {
                                if (139 != 255) {
                                    fmt.Println(v8)
                                }
                            } else {
                                v25 := GenS1{F0: v13, F1: 255}
                                _ = v25
                            }
                        }
                        _ = fi24
                    }
                case "hello":
                    var v26 bool = false
                    _ = v26
                }
            case 873:
                var v27 uint8 = 52
                _ = v27
            }
            switch (^(fi14)) {
            case -110:
                var v28 uint16 = 42
                _ = v28
            case 2147483647:
                var v29 int8 = -128
                _ = v29
            case 626:
                if false {
                    if false {
                        v8 = "hello"
                        for fi30 := 0; fi30 < 10; fi30++ {
                            v31 := GenS2{F0: 0, F1: 1, F2: v5}
                            var v32 int16 = -647
                            var v33 uint32 = 84
                            fmt.Println(v33, fi30)
                            _ = fi30
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                    } else {
                        v34 := GenS5{F0: "a"}
                        _ = v34
                    }
                    switch (^(fi14)) {
                    case -1:
                        v35 := GenS6{F0: v8}
                        _ = v35
                    case 2147483647:
                        var v36 uint32 = 2147483647
                        _ = v36
                    case -549:
                        v8 = ((v8 + ((("" + v8) + ((v8 + (v8 + v8)) + "hello")) + "\n\t")) + v8)
                    }
                    v37 := v3[fi14]
                    _ = v37
                }
            }
            _ = fi14
        }
        _ = v13
    }
    if (!(2 <= 994)) {
        if true {
            var v38 float64 = (1e10 - float64(3.14159))
            _ = v38
        } else {
            switch "a" {
            case "foo bar":
                var v39 uint16 = 255
                _ = v39
            }
            if true {
                v9 = v9
                v40 := GenS7{F0: -487}
                _ = v40
            }
        }
        switch v8 {
        case "a":
            var v41 int8 = v5
            _ = v41
        case "hello":
            switch -9223372036854775808 {
            case 128:
                v42 := GenS2{F0: -829, F1: -124, F2: v5}
                _ = v42
            case 127:
                var v43 int16 = 0
                _ = v43
            case 9223372036854775807:
                var v44 string = "\n\t"
                _ = v44
            }
        }
    }
    var v45 uint16 = 995
    v5 = v5
    if (549 != 4294967295) {
        v3[860] = 294
        var v46 uint8 = 2
        _ = v46
    } else {
        var v47 uint16 = v45
        _ = v47
    }
    v8 = (("\n\t" + v8) + (v8 + "foo bar"))
    _ = len(v1)
    if false {
        for fi48 := 0; fi48 < 4; fi48++ {
            v8 = "foo bar"
            v49 := GenS8{F0: v45}
            if true {
                var v50 float32 = v9
                _ = v50
            } else {
                switch ("\n\t" + "a") {
                case "foo bar":
                    v45 -= v45
                default:
                    var v51 float32 = v9
                    _ = v51
                }
                v52 := GenS9{F0: fi48, F1: 42}
                _ = v52
            }
            _ = fi48
            _ = v49
        }
        for fi53 := 0; fi53 < 5; fi53++ {
            fmt.Println(fi53, v45, fi53)
            _ = len(v4)
            v5 += (v5 >> ((18446744073709551615) & 63))
            _ = fi53
        }
        var v54 int8 = v5
        _ = v54
    }
    v55 := GenS10{F0: 877, F1: 236, F2: v45}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v55
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
