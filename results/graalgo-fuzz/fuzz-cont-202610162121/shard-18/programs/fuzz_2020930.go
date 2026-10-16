package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int16 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{}
    var v3 float32 = 3.14159
    v4 := GenS1{F0: 975, F1: 255, F2: 1}
    switch -372 {
    case -9223372036854775808:
        if (true || (false && true)) {
            v1 -= (v1 - (v1 % (((v1 * v1)) | 1)))
            v1 += v1
        } else {
            v5 := GenS1{F0: 128, F1: 65535, F2: 847}
            v6 := GenS2{F0: 188, F1: 0}
            v1 = (v1 & (v1 | (v1 << ((847) & 63))))
            _ = v5
            _ = v6
        }
    case 32767:
        for fi7 := 0; fi7 < 11; fi7++ {
            v3 = (v3 - v3)
            var v8 uint64 = 2
            if (65535 >= 475) {
                v2 = append(v2, fi7)
            }
            _ = fi7
            _ = v8
        }
    case -1:
        var v9 uint16 = 42
        _ = v9
    }
    if true {
        var v10 int64 = -1
        if ((3.14159 + ((0.1 - 1e10) * 3.14159)) <= ((1e10 / (1e10 * -2.25)) / -2.25)) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v11 string = ("a" + "hello")
            _ = v11
        }
        _ = v10
    } else {
        if ((!false) && false) {
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 int8 = 39
                v3 -= (v3 + v3)
                _ = fi12
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 3; fi14++ {
                if true {
                    var v15 float32 = v3
                    _ = v15
                }
                var v16 uint64 = 18446744073709551615
                if true {
                    fmt.Println(v16)
                    var v17 string = "a"
                    _ = v17
                }
                _ = fi14
                _ = v16
            }
        }
    }
    fmt.Println(v1, v1)
    if true {
        var v18 float64 = 0.0
        if true {
            var v19 float64 = 1e10
            fmt.Println(v3)
            fmt.Println(v1)
            _ = v19
        } else {
            _ = len(v2)
        }
        if false {
            var v20 uint = 1
            _ = v20
        }
        _ = v18
    }
    switch 32767 {
    case 4294967295:
        v3 -= v3
    case 32768:
        fmt.Println(v3, v1)
    }
    if (v3 != v3) {
        var v21 uint64 = 0
        for fi22 := 0; fi22 < 3; fi22++ {
            v3 = ((v3 - 0.1) + (((float32((3.14159 + -2.25)) / (-2.25 + 0.1)) - (v3 + float32(0.1))) - -2.25))
            _ = fi22
        }
        _ = v21
    } else {
        fmt.Println(v1, v3)
        v2 = append(v2, 127)
        var v23 int8 = 127
        _ = v23
    }
    switch 2147483647 {
    case -2147483648:
        for fi24 := 0; fi24 < 6; fi24++ {
            v25 := GenS3{F0: 0, F1: 255}
            if (true || true) {
                switch fi24 {
                case 2147483647:
                    v26 := GenS4{F0: 0}
                    _ = v26
                case 595:
                    var v27 int64 = v1
                    _ = v27
                case 128:
                    for fi28 := 0; fi28 < 4; fi28++ {
                        if len(v2) > 0 {
                            v2[0] = 712
                        }
                        v2 = append(v2, fi24)
                        _ = fi28
                    }
                default:
                    for fi29 := 0; fi29 < 6; fi29++ {
                        v1 += 42
                        if (!(true && (18446744073709551615 == 647))) {
                            fmt.Println(v3)
                            switch "\n\t" {
                            case "foo bar":
                                if len(v2) > 0 {
                                    v2[0] = fi24
                                }
                            case "":
                                var v30 bool = (false || (116 > 0))
                                _ = v30
                            }
                        }
                        v31 := GenS3{F0: 42, F1: 255}
                        v1 = v1
                        _ = fi29
                        _ = v31
                    }
                }
                switch (fi24 << ((158) & 63)) {
                case 9223372036854775807:
                    v2 = append(v2, (fi24 & -1))
                }
            }
            _ = fi24
            _ = v25
        }
    default:
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 string = ((("foo bar" + "foo bar") + "a") + (("a" + "hello") + ""))
            _ = fi32
            _ = v33
        }
    }
    switch 128 {
    case 119:
        switch "foo bar" {
        case "\n\t":
            if true {
                v3 = 0.0
                for fi34 := 0; fi34 < 8; fi34++ {
                    fmt.Println(v1)
                    _ = fi34
                }
            } else {
                var v35 uint32 = 2
                v35 -= (v35 + v35)
                _ = v35
            }
        default:
            var v36 int64 = v1
            _ = v36
        }
    case -747:
        v37 := GenS5{F0: 255, F1: true}
        _ = v37
    case -1:
        v1 = v1
    default:
        v38 := GenS1{F0: 0, F1: 1, F2: 42}
        _ = v38
    }
    v3 -= 1e10
    v39 := GenS6{F0: 32767}
    v3 = (3.14159 * v3)
    v3 -= float32(0.1)
    var v40 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}
