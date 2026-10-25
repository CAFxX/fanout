package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int; F1 int16; F2 int64 }
type GenS3 struct { F0 int; F1 uint64 }

func main() {
    v1 := []string{("\n\t" + ("" + "hello")), "foo bar", ("\n\t" + "foo bar"), ("" + ("a" + ("" + "")))}
    var v2 uint8 = 108
    var v3 int = -2147483648
    fmt.Println(v3, v3)
    switch "hello" {
    case "hello":
        var v4 int8 = 0
        _ = v4
    case "a":
        v5 := GenS1{F0: "\n\t", F1: v3}
        _ = v5
    }
    v3 = 912
    var v6 int16 = 647
    for fi7 := 0; fi7 < 6; fi7++ {
        switch "hello" {
        case "":
            v3 -= v3
        case "foo bar":
            switch "foo bar" {
            case "\n\t":
                v1 = append(v1, "")
            default:
                if (-32768 > (v6 << ((4294967295) & 63))) {
                    var v8 int64 = -99
                    v3 = fi7
                    if true {
                        v8 = -195
                        v2 += v2
                        if (!false) {
                            switch v3 {
                            case 206:
                                var v9 uint8 = (v2 * 128)
                                _ = v9
                            case -200:
                                for fi10 := 0; fi10 < 1; fi10++ {
                                    switch 2147483647 {
                                    case -329:
                                        v3 = (fi7 * v3)
                                    }
                                    _ = fi10
                                }
                            case 163:
                                var v11 int8 = -79
                                _ = v11
                            default:
                                for fi12 := 0; fi12 < 20; fi12++ {
                                    switch 1 {
                                    case -2147483648:
                                        switch "\n\t" {
                                        case "a":
                                            v3 -= 128
                                        case "hello":
                                            v13 := GenS1{F0: "\n\t", F1: (fi12 % (((fi12 / ((v3) | 1))) | 1))}
                                            _ = v13
                                        case "\n\t":
                                            switch (fi7 * v3) {
                                            case 658:
                                                v3 += v3
                                            case 127:
                                                v1 = append(v1, "hello")
                                            }
                                        default:
                                            v14 := GenS1{F0: ("" + "\n\t"), F1: fi12}
                                            _ = v14
                                        }
                                    case 4294967295:
                                        if (false || (1 < 1)) {
                                            var v15 int64 = int64(v6)
                                            v1 = append(v1, (("" + (("\n\t" + "a") + "hello")) + "a"))
                                            _ = v15
                                        } else {
                                            _ = len(v1)
                                        }
                                    case 408:
                                        v16 := GenS1{F0: "\n\t", F1: (^(v3))}
                                        _ = v16
                                    }
                                    _ = fi12
                                }
                            }
                            _ = len(v1)
                        } else {
                            v17 := GenS1{F0: "foo bar", F1: (v3 >> ((4294967295) & 63))}
                            fmt.Println(v6)
                            v3 = (v3 - (-(fi7)))
                            _ = v17
                        }
                    }
                    _ = v8
                }
            }
        case "a":
            var v18 uint64 = 1
            _ = v18
        default:
            v2 += v2
        }
        _ = fi7
    }
    var v19 uint32 = 0
    if false {
        for fi20 := 0; fi20 < 9; fi20++ {
            v3 = -745
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 int32 = 2147483647
                switch (fi21 ^ (fi20 << ((18446744073709551615) & 63))) {
                case 32767:
                    if (false || false) {
                        v23 := GenS1{F0: "foo bar", F1: (fi20 % ((int(v19)) | 1))}
                        switch ("\n\t" + "a") {
                        case "hello":
                            switch ("hello" + "a") {
                            case "foo bar":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "\n\t":
                                var v24 int = (fi20 + fi20)
                                var v25 uint = 513
                                switch ("foo bar" + "\n\t") {
                                case "a":
                                    if false {
                                        v19 -= 2147483648
                                    } else {
                                        v1 = append(v1, ("" + "a"))
                                        var v26 string = "foo bar"
                                        var v27 uint = v25
                                        _ = v26
                                        _ = v27
                                    }
                                case "\n\t":
                                    v28 := GenS2{F0: int(v22), F1: v6, F2: 829}
                                    _ = v28
                                case "foo bar":
                                    var v29 uint64 = 18446744073709551615
                                    _ = v29
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v24
                                _ = v25
                            }
                        case "\n\t":
                            if true {
                                var v30 int64 = 9223372036854775807
                                _ = v30
                            }
                        case "foo bar":
                            var v31 int = (^(fi20))
                            _ = v31
                        default:
                            v6 += (v6 >> ((0) & 63))
                        }
                        fmt.Println(fi21)
                        _ = v23
                    } else {
                        var v32 float32 = 1e10
                        fmt.Println(v22, v2)
                        v19 = (v19 + (v19 - 2))
                        _ = v32
                    }
                }
                fmt.Println(v19)
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    } else {
        switch 127 {
        case -138:
            fmt.Println(v19, v3, v6)
        default:
            fmt.Println(v3)
        }
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 bool = (!true)
            _ = fi33
            _ = v34
        }
        var v35 float64 = 1e10
        _ = v35
    }
    v3 = v3
    switch v3 {
    case 1:
        v36 := GenS3{F0: 127, F1: 2}
        _ = v36
    case 9223372036854775807:
        v2 = 2
    case 568:
        var v37 uint32 = (^(v19))
        _ = v37
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v19 = (^(v19))
    v1 = append(v1, "\n\t")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v6
}
