package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 int8; F2 int16 }

func main() {
    v1 := []uint32{1, 65535, 245, 65535}
    var v2 int32 = 1
    var v3 uint = 0
    v3 = 18446744073709551615
    for fi4 := 0; fi4 < 9; fi4++ {
        v2 -= (v2 - (^(v2)))
        var v5 uint64 = 42
        var v6 int32 = (-(v2))
        _ = fi4
        _ = v5
        _ = v6
    }
    v3 = v3
    v2 = -2147483648
    if (!(86 >= -1)) {
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: (("foo bar" + ("\n\t" + "\n\t")) + "foo bar"), F1: (v3 * v3)}
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            if (-1 >= -1) {
                if (("foo bar" + "hello") < "hello") {
                    v11 := GenS1{F0: ("" + "\n\t"), F1: (v3 | v3)}
                    _ = v11
                }
            } else {
                for fi12 := 0; fi12 < 9; fi12++ {
                    var v13 float64 = (0.1 - ((float64(1.5) - 0.1) + -2.25))
                    var v14 bool = (false || false)
                    switch (fi10 & fi10) {
                    case 1:
                        if v14 {
                            var v15 int = fi12
                            var v16 uint8 = 127
                            fmt.Println(v16)
                            _ = v15
                            _ = v16
                        }
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                v17 := GenS2{F0: 42, F1: 42}
                _ = v17
            }
            v3 = (v3 - v3)
            _ = fi10
        }
        fmt.Println(v2, v2)
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        v3 += v3
        switch (((("hello" + "") + "\n\t") + (("\n\t" + ("foo bar" + "foo bar")) + "a")) + (("foo bar" + "hello") + "\n\t")) {
        case "\n\t":
            fmt.Println(v2, v2, v2)
        }
        switch (-(fi18)) {
        case 802:
            v19 := GenS2{F0: 0, F1: 1}
            _ = v19
        case 2147483647:
            if (!(9223372036854775807 > int64(v2))) {
                for fi20 := 0; fi20 < 20; fi20++ {
                    if true {
                        var v21 int64 = 9223372036854775807
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v21
                    }
                    v2 += v2
                    var v22 int32 = (-(v2))
                    _ = fi20
                    _ = v22
                }
            }
        default:
            v3 = 145
        }
        _ = fi18
    }
    fmt.Println(v3, v2)
    fmt.Println(v2, v3)
    var v23 int8 = int8(v3)
    v3 += (^(v3))
    v24 := GenS3{F0: (v23 << ((v3) & 63)), F1: (1 == 649)}
    v3 = v3
    fmt.Println(v2, v23, v2)
    for fi25 := 0; fi25 < 15; fi25++ {
        for fi26 := 0; fi26 < 18; fi26++ {
            if false {
                if false {
                    switch fi25 {
                    case -9223372036854775808:
                        if len(v1) > 0 {
                            v1[0] = 412
                        }
                    case -2147483648:
                        if true {
                            if (true || ((v2 >> ((299) & 63)) == (v2 * (v2 / ((v2) | 1))))) {
                                _ = len(v1)
                                fmt.Println(v3, v23, v3)
                                var v27 string = "hello"
                                switch (-(fi25)) {
                                case 9223372036854775807:
                                    v28 := GenS2{F0: 1, F1: 42}
                                    _ = v28
                                }
                                switch fi26 {
                                case -2147483648:
                                    v29 := GenS4{F0: v23}
                                    _ = v29
                                case 128:
                                    fmt.Println(fi26, fi25)
                                case 32768:
                                    fmt.Println(v3, fi25)
                                }
                                v23 = v23
                                _ = len(v1)
                                v1 = append(v1, 2147483647)
                                _ = v27
                            }
                        }
                    case 1:
                        _ = len(v1)
                    }
                    v2 = 241
                } else {
                    switch -352 {
                    case -9223372036854775808:
                        v23 = v23
                    case 4294967295:
                        v30 := GenS2{F0: 2, F1: 454}
                        _ = v30
                    default:
                        v1 = append(v1, 592)
                    }
                    switch (fi26 - (fi26 % ((fi26) | 1))) {
                    case 43:
                        v23 -= (v23 + (-(v23)))
                    case 1:
                        var v31 bool = true
                        v32 := GenS3{F0: v23, F1: ((42 >= 1) || v31)}
                        _ = v31
                        _ = v32
                    case -1:
                        v3 -= (v3 >> (((^(v3))) & 63))
                    default:
                        var v33 uint32 = 965
                        _ = v33
                    }
                    var v34 int8 = (v23 - (v23 | (v23 % (((v23 & v23)) | 1))))
                    _ = v34
                }
            } else {
                fmt.Println(v2, fi25)
                v23 = 127
            }
            switch ((("foo bar" + "\n\t") + "") + ((("hello" + (("a" + ("hello" + ("\n\t" + (("hello" + "a") + ("" + "\n\t"))))) + "foo bar")) + "") + ("" + ("" + "")))) {
            case "a":
                var v35 int64 = -9223372036854775808
                _ = v35
            case "":
                v36 := GenS5{F0: 444, F1: -2147483648}
                _ = v36
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                v23 = v23
            }
            v37 := GenS6{F0: (v23 >> ((42) & 63))}
            _ = fi26
            _ = v37
        }
        v1 = append(v1, 2147483648)
        _ = fi25
    }
    v23 -= v23
    v2 -= (v2 - v2)
    for fi38 := 0; fi38 < 2; fi38++ {
        switch fi38 {
        case 4294967295:
            v39 := GenS1{F0: ("a" + ("\n\t" + "")), F1: (v3 ^ v3)}
            _ = v39
        case 0:
            var v40 string = "hello"
            _ = v40
        default:
            v41 := GenS7{F0: 2, F1: (v23 * 42), F2: -32768}
            _ = v41
        }
        for fi42 := 0; fi42 < 18; fi42++ {
            v23 = 0
            v1 = append(v1, 65535)
            if (((float32(-2.25) - -2.25) - ((0.0 * 1.5) * 3.14159)) <= 3.14159) {
                fmt.Println(v23)
                v2 += 2147483647
            } else {
                var v43 uint8 = 212
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v43
            }
            v44 := GenS5{F0: (fi42 + (^(fi42))), F1: int32(fi42)}
            _ = fi42
            _ = v44
        }
        _ = fi38
    }
    var v45 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v45
}
