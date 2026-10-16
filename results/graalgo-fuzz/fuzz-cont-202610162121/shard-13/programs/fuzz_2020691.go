package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 2
    var v3 uint = 4294967295
    var v4 uint32 = 644
    if true {
        v4 += (v4 % (((v4 - 2)) | 1))
        fmt.Println(v2)
        v3 -= v3
    } else {
        v4 = 0
    }
    v2 -= v2
    if (!false) {
        switch "hello" {
        case "":
            v2 -= v2
        case "foo bar":
            var v5 string = "foo bar"
            _ = v5
        case "a":
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 int8 = 42
                v1[128] = (("foo bar" + "a") + "hello")
                v1[fi6] = "\n\t"
                _ = fi6
                _ = v7
            }
        default:
            var v8 int16 = 32767
            _ = v8
        }
    }
    var v9 bool = true
    if v9 {
        if (-2.25 < ((1e10 - 1.5) * (0.1 + 0.1))) {
            for fi10 := 0; fi10 < 7; fi10++ {
                if (v3 > (v3 + 766)) {
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v3 = v3
                        if v9 {
                            var v12 bool = v9
                            v13 := GenS1{F0: v3, F1: 647}
                            switch "\n\t" {
                            case "a":
                                switch 2147483647 {
                                case 42:
                                    fmt.Println(fi11, v12, v12)
                                case 9223372036854775807:
                                    v14 := GenS1{F0: v3, F1: 2147483647}
                                    _ = v14
                                case -623:
                                    var v15 bool = v12
                                    _ = v15
                                default:
                                    if v9 {
                                        switch -9223372036854775808 {
                                        case -366:
                                            if ((v12 && v12) && v12) {
                                                fmt.Println(v9)
                                                var v16 uint = v3
                                                v17 := GenS1{F0: (v3 / ((v3) | 1)), F1: 1}
                                                switch fi10 {
                                                case 32767:
                                                    v12 = v12
                                                }
                                                v3 = 42
                                                _ = v16
                                                _ = v17
                                            } else {
                                                if false {
                                                    v12 = v12
                                                    if true {
                                                        var v18 string = "foo bar"
                                                        v4 += uint32(v3)
                                                        _ = v18
                                                    } else {
                                                        v12 = v12
                                                    }
                                                    v19 := GenS2{F0: -128}
                                                    _ = v19
                                                }
                                            }
                                        case -1:
                                            switch ((((("" + ("hello" + ("a" + "hello"))) + ("a" + "hello")) + ("hello" + "")) + ("" + ("foo bar" + "foo bar"))) + "") {
                                            case "hello":
                                                if (v9 && (245 < 1)) {
                                                    v20 := []int{}
                                                    _ = v20
                                                }
                                            case "foo bar":
                                                var v21 int64 = -9223372036854775808
                                                _ = v21
                                            case "\n\t":
                                                v12 = (true && v12)
                                                fmt.Println(fi10, v9)
                                            default:
                                                switch "a" {
                                                case "hello":
                                                    var v22 int16 = -3
                                                    _ = v22
                                                }
                                            }
                                        default:
                                            switch "\n\t" {
                                            case "":
                                                v23 := GenS3{F0: false, F1: -2147483648}
                                                _ = v23
                                            case "foo bar":
                                                v24 := []int{fi10, 1}
                                                _ = v24
                                            }
                                        }
                                        if (1.5 >= float32(3.14159)) {
                                            _ = len(v1)
                                            var v25 uint = (v3 - (v3 | v3))
                                            if ((128 == 0) || true) {
                                                if v9 {
                                                    switch "hello" {
                                                    case "a":
                                                        v26 := GenS4{F0: 42}
                                                        _ = v26
                                                    case "foo bar":
                                                        v2 = 89
                                                    }
                                                    v27 := GenS4{F0: 42}
                                                    fmt.Println(v2, v3, fi10)
                                                    _ = v27
                                                } else {
                                                    switch fi11 {
                                                    case 829:
                                                        var v28 string = ((("" + ("\n\t" + "foo bar")) + "foo bar") + "a")
                                                        _ = v28
                                                    case -289:
                                                        var v29 int32 = 401
                                                        _ = v29
                                                    default:
                                                        v30 := v1[(-(fi11))]
                                                        _ = v30
                                                    }
                                                    switch 886 {
                                                    case -9223372036854775808:
                                                        fmt.Println(v2, fi11, v12)
                                                    case 1:
                                                        v31 := GenS2{F0: 0}
                                                        _ = v31
                                                    case -851:
                                                        v12 = (!v9)
                                                    }
                                                }
                                            } else {
                                                var v32 uint8 = 0
                                                if (2 >= 60) {
                                                    v25 -= (^(v25))
                                                }
                                                v3 = (v25 * 42)
                                                v9 = v12
                                                switch (("\n\t" + "hello") + "") {
                                                case "\n\t":
                                                    var v33 uint8 = v32
                                                    _ = v33
                                                case "a":
                                                    v32 = (v32 + v32)
                                                }
                                                _ = v32
                                            }
                                            _ = v25
                                        }
                                    }
                                }
                            default:
                                v34 := GenS3{F0: (fi10 >= (fi11 / ((fi10) | 1))), F1: 0}
                                _ = v34
                            }
                            _ = v12
                            _ = v13
                        }
                        _ = fi11
                    }
                    for fi35 := 0; fi35 < 19; fi35++ {
                        switch fi10 {
                        case 106:
                            v36 := []int{fi10, 268}
                            _ = v36
                        }
                        _ = fi35
                    }
                    fmt.Println(v3, v4, v2)
                }
                var v37 uint16 = 32767
                v38 := GenS1{F0: (v3 % (((v3 << ((v3) & 63))) | 1)), F1: 13}
                _ = fi10
                _ = v37
                _ = v38
            }
        }
        v39 := GenS2{F0: -1}
        _ = v39
    } else {
        v3 = (^(v3))
        fmt.Println(v2, v9)
    }
    v4 = v4
    v3 = 2
    v40 := []int{0, 32767}
    v9 = true
    if len(v40) > 0 {
        v40[0] = 9223372036854775807
    }
    switch "foo bar" {
    case "":
        var v41 int16 = -32768
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v9
}
