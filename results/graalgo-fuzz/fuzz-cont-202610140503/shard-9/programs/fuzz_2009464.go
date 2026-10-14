package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint16; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 bool; F2 uint32 }

func main() {
    var v1 string = ("foo bar" + "")
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    v1 = ""
    switch (v1 + ((((v1 + (v1 + (v1 + ("" + "hello")))) + (v1 + v1)) + "hello") + (((v1 + "hello") + v1) + (v1 + (v1 + (("" + v1) + "\n\t")))))) {
    case "a":
        switch 2147483647 {
        case -1:
            v4 := GenS1{F0: 2, F1: 2147483647, F2: "a"}
            _ = v4
        case 9223372036854775807:
            v1 = "foo bar"
        default:
            if (-2.25 < (float32(3.14159) + (-2.25 - 0.1))) {
                var v5 uint8 = 199
                for fi6 := 0; fi6 < 4; fi6++ {
                    if false {
                        if (!(!(false || true))) {
                            var v7 int32 = -958
                            fmt.Println(v1, fi6)
                            _ = v7
                        }
                        var v8 int16 = 1000
                        _ = v8
                    } else {
                        var v9 uint = 255
                        for fi10 := 0; fi10 < 7; fi10++ {
                            var v11 int32 = 2147483647
                            v9 += v9
                            _ = fi10
                            _ = v11
                        }
                        if (0.1 < 0.1) {
                            v12 := GenS2{F0: (v9 * v9)}
                            var v13 int32 = 42
                            v5 = v5
                            _ = v12
                            _ = v13
                        }
                        _ = v9
                    }
                    _ = fi6
                }
                _ = v5
            }
        }
    default:
        fmt.Println(v1)
    }
    if false {
        if (!(-772 >= 128)) {
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 bool = (true && (false || true))
                if (!false) {
                    v16 := v2["foo bar"]
                    for fi17 := 0; fi17 < 19; fi17++ {
                        switch fi14 {
                        case 1:
                            switch "a" {
                            case "hello":
                                v18 := []int{-2147483648, (fi17 & fi17), 2147483647, 2147483647}
                                _ = v18
                            }
                        }
                        _ = fi17
                    }
                    v2[("\n\t" + "")] = (v16 & v16)
                    _ = v16
                } else {
                    v19 := []int{(fi14 * (fi14 >> ((810) & 63))), (fi14 - fi14), (fi14 << ((4294967295) & 63))}
                    if false {
                        v1 = v1
                        for fi20 := 0; fi20 < 7; fi20++ {
                            var v21 int64 = -108
                            _ = fi20
                            _ = v21
                        }
                    }
                    _ = len(v2)
                    _ = v19
                }
                _ = fi14
                _ = v15
            }
        } else {
            v22 := v2[v1]
            if (322 != 0) {
                v23 := []int{127, 0, 1, 32767}
                for fi24 := 0; fi24 < 15; fi24++ {
                    var v25 uint16 = 32767
                    _ = fi24
                    _ = v25
                }
                _ = v23
            } else {
                for fi26 := 0; fi26 < 1; fi26++ {
                    switch fi26 {
                    case -9223372036854775808:
                        v22 = v22
                    }
                    fmt.Println(v22)
                    _ = fi26
                }
            }
            _ = v22
        }
    } else {
        fmt.Println(v1)
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        if (((!false) && (false || false)) && (-682 == 32767)) {
            v28 := []int{}
            fmt.Println(fi27, v1)
            _ = v28
        } else {
            v29 := GenS1{F0: 32767, F1: 2147483647, F2: v1}
            _ = v29
        }
        var v30 uint8 = 42
        if (42 < 1) {
            if true {
                fmt.Println(fi27, fi27)
            } else {
                v30 += v30
                if false {
                    if true {
                        fmt.Println(fi27)
                    } else {
                        v31 := GenS3{F0: fi27, F1: 12, F2: 42}
                        _ = v31
                    }
                    v1 = "hello"
                }
            }
            _ = len(v2)
        }
        _ = fi27
        _ = v30
    }
    fmt.Println(v1)
    switch v1 {
    case "hello":
        var v32 int32 = -197
        _ = v32
    case "\n\t":
        v33 := GenS2{F0: 65535}
        _ = v33
    default:
        v1 = v1
    }
    v34 := GenS2{F0: 42}
    v35 := []int{}
    v36 := GenS4{F0: 32767}
    for fi37 := 0; fi37 < 8; fi37++ {
        var v38 uint32 = 744
        v1 = "hello"
        _ = fi37
        _ = v38
    }
    v1 = "a"
    switch "" {
    case "a":
        delete(v3, 0)
    default:
        for fi39 := 0; fi39 < 9; fi39++ {
            switch (((((v1 + v1) + "hello") + "hello") + v1) + (v1 + (v1 + v1))) {
            case "foo bar":
                for fi40 := 0; fi40 < 5; fi40++ {
                    if true {
                        if false {
                            _ = len(v35)
                            v35 = append(v35, (fi40 & 32768))
                        }
                        var v41 uint8 = 0
                        _ = v41
                    } else {
                        var v42 int32 = 0
                        v42 -= (^(v42))
                        _ = v42
                    }
                    var v43 int8 = -9
                    _ = fi40
                    _ = v43
                }
            case "\n\t":
                switch "\n\t" {
                case "\n\t":
                    switch "" {
                    case "\n\t":
                        v44 := GenS5{F0: v1}
                        _ = v44
                    case "hello":
                        v1 = (v1 + v1)
                    case "foo bar":
                        if false {
                            for fi45 := 0; fi45 < 3; fi45++ {
                                fmt.Println(v1)
                                _ = fi45
                            }
                            _ = len(v35)
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                        }
                    }
                case "hello":
                    v1 = v1
                default:
                    if false {
                        if false {
                            var v46 uint8 = 22
                            v47 := GenS3{F0: fi39, F1: 32767, F2: -9223372036854775808}
                            v1 = "foo bar"
                            _ = v46
                            _ = v47
                        }
                    }
                }
            case "hello":
                for fi48 := 0; fi48 < 11; fi48++ {
                    v49 := GenS6{F0: false, F1: true, F2: 2}
                    _ = fi48
                    _ = v49
                }
            }
            _ = fi39
        }
    }
    if false {
        var v50 int = 128
        var v51 string = v1
        v50 += 999
        _ = v50
        _ = v51
    } else {
        v1 = "foo bar"
    }
    for fi52 := 0; fi52 < 8; fi52++ {
        var v53 int16 = 65
        v53 = -32768
        _ = fi52
        _ = v53
    }
    switch (v1 + v1) {
    case "a":
        v3[2] = (v1 + "hello")
    case "foo bar":
        v1 = v1
    }
    fmt.Println(v1)
    if (1 <= 1) {
        v1 = "foo bar"
    }
    delete(v3, 2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
