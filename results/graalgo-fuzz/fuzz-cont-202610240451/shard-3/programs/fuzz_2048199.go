package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint8; F1 int64; F2 int64 }
type GenS4 struct { F0 uint16; F1 string }
func genFn0(p0 uint8) int32 {
    v4 := GenS1{F0: 0, F1: 32767}
    v5 := []int{4294967295}
    v5 = append(v5, 128)
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            p0 = p0
            switch (fi6 | -915) {
            case -9223372036854775808:
                _ = len(v5)
            }
            p0 += p0
            _ = fi7
        }
        if true {
            p0 -= (^(p0))
        }
        _ = fi6
    }
    v5 = append(v5, 119)
    _ = p0
    _ = v4
    _ = v5
    return 585
}

func main() {
    var v1 uint32 = 308
    v2 := []int{-9223372036854775808, 808, -9223372036854775808}
    var v3 float64 = 0.1
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 uint = 255
        if (0 != 32767) {
            if ((fi8 == 0) && true) {
                v10 := GenS1{F0: -2147483648, F1: fi8}
                if false {
                    for fi11 := 0; fi11 < 1; fi11++ {
                        var v12 uint32 = v1
                        if false {
                            var v13 string = (("foo bar" + ("" + "\n\t")) + ("a" + ("a" + ("foo bar" + "hello"))))
                            var v14 bool = (!false)
                            v3 -= v3
                            _ = v13
                            _ = v14
                        }
                        _ = fi11
                        _ = v12
                    }
                }
                var v15 string = ("foo bar" + "hello")
                _ = v10
                _ = v15
            } else {
                switch "\n\t" {
                case "hello":
                    v1 -= 850
                case "foo bar":
                    switch -1 {
                    case 0:
                        var v16 int32 = -1
                        _ = v16
                    case 127:
                        v9 -= 65535
                    }
                }
            }
            var v17 uint32 = v1
            _ = v17
        } else {
            switch (((("a" + "\n\t") + "") + ("a" + "")) + "a") {
            case "a":
                for fi18 := 0; fi18 < 17; fi18++ {
                    v9 += (v9 / (((^(v9))) | 1))
                    switch (fi8 * 74) {
                    case 32767:
                        v2 = append(v2, fi18)
                    }
                    _ = fi18
                }
            case "hello":
                switch "\n\t" {
                case "":
                    if (3.14159 != 1.5) {
                        if (1.5 > 3.14159) {
                            v19 := GenS2{F0: 42, F1: int32(v9)}
                            fmt.Println(v1, v1, v9)
                            _ = v19
                        } else {
                            v3 = v3
                            v9 -= (^(v9))
                        }
                    } else {
                        if true {
                            var v20 uint = (^(v9))
                            for fi21 := 0; fi21 < 14; fi21++ {
                                var v22 int64 = 442
                                switch "a" {
                                case "":
                                    v22 = v22
                                case "hello":
                                    var v23 uint8 = 128
                                    _ = v23
                                case "\n\t":
                                    v20 = 65535
                                }
                                var v24 uint64 = 42
                                var v25 uint = (v20 >> (((^(v9))) & 63))
                                var v26 int = (fi21 / ((int(v22)) | 1))
                                _ = fi21
                                _ = v22
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                            if false {
                                switch ("" + "\n\t") {
                                case "hello":
                                    switch ("" + "\n\t") {
                                    case "\n\t":
                                        var v27 int8 = 100
                                        _ = v27
                                    case "a":
                                        switch ("hello" + "hello") {
                                        case "hello":
                                            if (!true) {
                                                for fi28 := 0; fi28 < 19; fi28++ {
                                                    var v29 int32 = genFn0(127)
                                                    if false {
                                                        v30 := GenS1{F0: 42, F1: (-(fi8))}
                                                        _ = v30
                                                    }
                                                    v3 += 0.1
                                                    var v31 uint32 = v1
                                                    _ = fi28
                                                    _ = v29
                                                    _ = v31
                                                }
                                                v9 = (^(v9))
                                                _ = len(v2)
                                            }
                                        case "foo bar":
                                            var v32 uint16 = 255
                                            _ = v32
                                        }
                                    }
                                case "":
                                    v33 := GenS3{F0: 42, F1: -1, F2: -1}
                                    _ = v33
                                }
                            }
                            _ = v20
                        } else {
                            if true {
                                v2 = append(v2, (fi8 * (fi8 / ((fi8) | 1))))
                            }
                            switch 1 {
                            case 682:
                                var v34 int64 = 1
                                _ = v34
                            case -134:
                                var v35 uint8 = 52
                                _ = v35
                            case 357:
                                v36 := GenS4{F0: 255, F1: ""}
                                _ = v36
                            default:
                                v37 := GenS1{F0: genFn0(42), F1: -172}
                                _ = v37
                            }
                            var v38 float32 = 0.1
                            _ = v38
                        }
                        var v39 int = 2147483647
                        var v40 uint16 = 32767
                        _ = v39
                        _ = v40
                    }
                case "foo bar":
                    var v41 uint = (v9 % ((314) | 1))
                    _ = v41
                }
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                var v42 string = ""
                _ = v42
            }
        }
        _ = fi8
        _ = v9
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        fmt.Println(v3, v3, v1)
        if len(v2) > 0 {
            v2[0] = 32768
        }
        _ = fi43
    }
    if (!(79 > -9223372036854775808)) {
        if (!true) {
            v1 = v1
            var v44 int64 = 9223372036854775807
            v1 -= v1
            _ = v44
        }
        v45 := GenS2{F0: -1, F1: genFn0(1)}
        switch 32768 {
        case 128:
            switch "a" {
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = 128
                }
            case "a":
                _ = len(v2)
            default:
                v46 := GenS1{F0: 0, F1: 4294967295}
                _ = v46
            }
        case 103:
            if false {
                var v47 bool = true
                var v48 uint32 = v1
                v2 = append(v2, -656)
                _ = v47
                _ = v48
            } else {
                fmt.Println(v3, v1)
                if (!false) {
                    switch "a" {
                    case "":
                        v49 := GenS3{F0: 170, F1: -184, F2: 718}
                        _ = v49
                    case "foo bar":
                        for fi50 := 0; fi50 < 3; fi50++ {
                            var v51 bool = false
                            fmt.Println(v1)
                            _ = fi50
                            _ = v51
                        }
                    case "hello":
                        v3 = float64((1.5 * 0.0))
                    }
                    v2 = append(v2, 32768)
                    v1 = (v1 % ((v1) | 1))
                }
                var v52 float64 = 1e10
                _ = v52
            }
        }
        _ = v45
    } else {
        var v53 float32 = (1e10 - 3.14159)
        switch 32767 {
        case 42:
            var v54 int32 = 42
            _ = v54
        case 127:
            var v55 bool = false
            _ = v55
        case -1:
            switch -652 {
            case 32768:
                var v56 uint64 = 42
                _ = v56
            case 9223372036854775807:
                var v57 uint64 = 707
                _ = v57
            default:
                v58 := GenS4{F0: 935, F1: ("foo bar" + "")}
                _ = v58
            }
        }
        var v59 int64 = 152
        _ = v53
        _ = v59
    }
    v1 += (v1 + (v1 << ((620) & 63)))
    v60 := GenS3{F0: 78, F1: -9223372036854775808, F2: 42}
    v3 += v3
    if (127 != 32768) {
        for fi61 := 0; fi61 < 18; fi61++ {
            var v62 uint8 = 128
            _ = fi61
            _ = v62
        }
    }
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v60
}
