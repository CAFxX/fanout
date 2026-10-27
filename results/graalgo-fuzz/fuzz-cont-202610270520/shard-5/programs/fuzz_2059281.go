package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int16; F1 uint; F2 int }
type GenS5 struct { F0 uint; F1 int16; F2 int16 }
type GenS6 struct { F0 int; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"\n\t", "hello", ("foo bar" + "hello")}
    v3 := make(map[uint32]bool)
    var v4 uint8 = 255
    var v5 uint16 = 255
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: -74, F1: 55, F2: 205}
        var v8 uint32 = 2
        for fi9 := 0; fi9 < 20; fi9++ {
            switch (("hello" + (("" + ("a" + ("foo bar" + (("" + "hello") + ("foo bar" + "\n\t"))))) + "\n\t")) + "a") {
            case "foo bar":
                fmt.Println(v4)
            default:
                fmt.Println(fi9, v8)
            }
            if false {
                v10 := GenS2{F0: ("\n\t" + "\n\t")}
                v5 -= 269
                var v11 uint64 = 252
                _ = v10
                _ = v11
            } else {
                v4 -= v4
                delete(v3, (v8 - (v8 + (v8 / ((v8) | 1)))))
            }
            _ = fi9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v12 int32 = 42
    var v13 float32 = (1e10 + ((1e10 + 0.0) / float32(-2.25)))
    for fi14 := 0; fi14 < 11; fi14++ {
        if true {
            v4 = (v4 + (v4 >> ((4294967295) & 63)))
        } else {
            fmt.Println(v12)
            for fi15 := 0; fi15 < 1; fi15++ {
                switch ("\n\t" + "") {
                case "hello":
                    fmt.Println(fi14, v5, fi14)
                case "foo bar":
                    var v16 string = "\n\t"
                    _ = v16
                }
                _ = fi15
            }
            v5 = 32
        }
        var v17 int8 = 59
        switch (fi14 | (fi14 - 32768)) {
        case 178:
            switch (-(fi14)) {
            case 32768:
                switch "hello" {
                case "hello":
                    if false {
                        switch "" {
                        case "a":
                            switch "\n\t" {
                            case "a":
                                for fi18 := 0; fi18 < 8; fi18++ {
                                    v12 -= v12
                                    _ = fi18
                                }
                            default:
                                v19 := GenS3{F0: fi14, F1: 969}
                                _ = v19
                            }
                        case "":
                            for fi20 := 0; fi20 < 10; fi20++ {
                                v21 := v1[(fi14 % ((int(v12)) | 1))]
                                var v22 string = "foo bar"
                                switch (fi14 >> ((2) & 63)) {
                                case 324:
                                    fmt.Println(v12, v13, fi14)
                                }
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        }
                        var v23 int32 = 335
                        v1[fi14] = 276
                        _ = v23
                    } else {
                        if false {
                            if true {
                                var v24 string = ("\n\t" + "")
                                var v25 int16 = 32767
                                _ = v24
                                _ = v25
                            } else {
                                if true {
                                    switch (fi14 >> ((uint(v5)) & 63)) {
                                    case -2147483648:
                                        v26 := GenS1{F0: 0, F1: 2147483647, F2: 0}
                                        _ = v26
                                    }
                                    if ((((755 < 2) && false) || false) && true) {
                                        fmt.Println(v4)
                                    }
                                    for fi27 := 0; fi27 < 17; fi27++ {
                                        fmt.Println(v5, v12)
                                        v28 := GenS4{F0: -1, F1: 326, F2: int(v12)}
                                        _ = fi27
                                        _ = v28
                                    }
                                }
                                delete(v3, 767)
                            }
                            if true {
                                var v29 string = ""
                                v3[42] = false
                                _ = v29
                            }
                        } else {
                            var v30 float64 = 1e10
                            switch fi14 {
                            case 2147483647:
                                var v31 uint8 = (^(v4))
                                _ = v31
                            case 961:
                                v32 := v1[-787]
                                _ = v32
                            case 42:
                                var v33 uint16 = (v5 | 42)
                                _ = v33
                            default:
                                var v34 int16 = 1000
                                _ = v34
                            }
                            _ = v30
                        }
                    }
                case "foo bar":
                    v12 = (v12 ^ 0)
                }
            case 4294967295:
                v35 := v3[1]
                _ = v35
            default:
                for fi36 := 0; fi36 < 20; fi36++ {
                    v37 := GenS5{F0: 42, F1: -1, F2: -44}
                    _ = fi36
                    _ = v37
                }
            }
        case 32767:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v13 -= float32(1.5)
        }
        _ = fi14
        _ = v17
    }
    v38 := v3[2147483647]
    var v39 uint64 = 2
    var v40 int8 = -67
    var v41 uint16 = v5
    var v42 uint = 42
    _ = len(v1)
    v5 -= 442
    switch -9223372036854775808 {
    case 2147483647:
        v41 -= v5
    case 424:
        v43 := GenS6{F0: 670, F1: 919}
        _ = v43
    case 1:
        v42 = 42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
