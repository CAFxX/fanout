package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []uint32{2147483647, 2147483647, 65535}
    var v2 uint = 2
    var v3 uint16 = 980
    v4 := make(map[string]string)
    v5 := GenS1{F0: 1}
    _ = len(v4)
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            switch fi7 {
            case -391:
                v3 += 255
            }
            _ = fi7
        }
        for fi8 := 0; fi8 < 3; fi8++ {
            v2 -= (v2 / ((2) | 1))
            _ = fi8
        }
        for fi9 := 0; fi9 < 8; fi9++ {
            switch (fi9 % (((fi9 - fi9)) | 1)) {
            case 1:
                var v10 uint32 = 2147483647
                _ = v10
            default:
                v3 -= 42
            }
            _ = fi9
        }
        _ = fi6
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        switch "" {
        case "":
            var v12 int16 = 32767
            _ = v12
        case "hello":
            v3 = v3
        case "a":
            switch "\n\t" {
            case "a":
                v13 := GenS2{F0: 2}
                _ = v13
            }
        default:
            v2 += (^(v2))
        }
        v3 = v3
        _ = fi11
    }
    switch -9223372036854775808 {
    case 128:
        switch "\n\t" {
        case "\n\t":
            switch 32767 {
            case -2147483648:
                v14 := GenS2{F0: 612}
                _ = v14
            }
        case "a":
            switch "" {
            case "":
                _ = len(v1)
            case "\n\t":
                if (128 != 42) {
                    var v15 int32 = 709
                    _ = v15
                }
            case "a":
                switch "hello" {
                case "a":
                    switch 127 {
                    case -1:
                        switch 9223372036854775807 {
                        case 4294967295:
                            switch "foo bar" {
                            case "\n\t":
                                var v16 int16 = 0
                                _ = v16
                            case "hello":
                                for fi17 := 0; fi17 < 3; fi17++ {
                                    var v18 uint = (^(v2))
                                    v2 = v18
                                    var v19 uint32 = 591
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                            case "":
                                if (-294 != 1) {
                                    switch "a" {
                                    case "hello":
                                        v4["a"] = ((("" + "hello") + ("a" + ("" + "hello"))) + "hello")
                                    case "":
                                        fmt.Println(v3, v2)
                                    }
                                    v20 := GenS1{F0: 42}
                                    _ = v20
                                }
                            default:
                                v21 := GenS2{F0: 114}
                                _ = v21
                            }
                        case 0:
                            v4[("\n\t" + "foo bar")] = "hello"
                        }
                    case 128:
                        var v22 uint8 = 1
                        _ = v22
                    case -46:
                        var v23 int64 = 42
                        _ = v23
                    default:
                        v3 += 65535
                    }
                }
            default:
                if (4294967295 <= 293) {
                    switch "\n\t" {
                    case "a":
                        v2 = (^(v2))
                    case "\n\t":
                        v3 -= v3
                    case "foo bar":
                        var v24 uint16 = 255
                        _ = v24
                    default:
                        v2 = (v2 + 0)
                    }
                    var v25 int64 = 9223372036854775807
                    _ = v25
                }
            }
        case "foo bar":
            switch -386 {
            case 1:
                var v26 int32 = 2147483647
                _ = v26
            default:
                v2 = v2
            }
        default:
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 string = "foo bar"
                if true {
                    if false {
                        switch "foo bar" {
                        case "a":
                            v3 -= 105
                        }
                        v1 = append(v1, 2147483647)
                    } else {
                        v29 := GenS1{F0: -1}
                        if (false && (!(!(v2 != v2)))) {
                            switch (fi27 / (((fi27 - fi27)) | 1)) {
                            case 373:
                                var v30 float32 = float32(3.14159)
                                _ = v30
                            }
                            v28 = v28
                        }
                        _ = v29
                    }
                    var v31 int8 = int8(v2)
                    var v32 uint8 = 127
                    _ = v31
                    _ = v32
                }
                if (1 > v3) {
                    var v33 string = v28
                    v3 += (^(v3))
                    _ = v33
                }
                _ = fi27
                _ = v28
            }
        }
    case 665:
        for fi34 := 0; fi34 < 20; fi34++ {
            delete(v4, "\n\t")
            _ = fi34
        }
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        var v36 float32 = -2.25
        var v37 float64 = -2.25
        _ = fi35
        _ = v36
        _ = v37
    }
    switch "\n\t" {
    case "hello":
        var v38 uint16 = (v3 % ((v3) | 1))
        _ = v38
    case "\n\t":
        var v39 string = "foo bar"
        _ = v39
    default:
        if len(v1) > 0 {
            v1[0] = 42
        }
    }
    if (!(-128 <= 100)) {
        for fi40 := 0; fi40 < 6; fi40++ {
            v3 = v3
            switch "" {
            case "":
                v4["\n\t"] = ((("foo bar" + ("a" + "\n\t")) + "\n\t") + ("hello" + ("hello" + "")))
            }
            _ = fi40
        }
    } else {
        var v41 string = "a"
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
