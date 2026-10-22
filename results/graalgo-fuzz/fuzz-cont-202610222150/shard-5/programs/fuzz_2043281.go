package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint64; F1 int16; F2 bool }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64; F1 int; F2 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int32 }

func main() {
    var v1 int8 = 42
    var v2 int16 = -32768
    var v3 uint8 = 42
    var v4 int = -414
    var v5 uint16 = 414
    var v6 uint16 = v5
    v4 = v4
    v7 := []int{-219}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v1 = (v1 >> ((255) & 63))
    v8 := GenS1{F0: 0}
    v9 := GenS1{F0: 1}
    for fi10 := 0; fi10 < 20; fi10++ {
        if (!false) {
            fmt.Println(fi10, v6, v6)
        } else {
            var v11 uint64 = 2
            if (9223372036854775807 != 42) {
                switch ("" + "") {
                case "foo bar":
                    v5 = (^(v6))
                case "hello":
                    v12 := GenS2{F0: v3, F1: (v2 - (v2 ^ 1000))}
                    _ = v12
                default:
                    if ("\n\t" != "\n\t") {
                        fmt.Println(v5, v1, v2)
                        var v13 string = ("\n\t" + "hello")
                        v13 = "\n\t"
                        _ = v13
                    } else {
                        switch (^(fi10)) {
                        case 9223372036854775807:
                            var v14 int16 = v2
                            _ = v14
                        case 32767:
                            for fi15 := 0; fi15 < 17; fi15++ {
                                fmt.Println(v11)
                                var v16 string = ""
                                _ = fi15
                                _ = v16
                            }
                        }
                        switch "hello" {
                        case "hello":
                            v11 += v11
                        case "\n\t":
                            var v17 int8 = 38
                            _ = v17
                        case "a":
                            switch "" {
                            case "hello":
                                if (!(1 < v11)) {
                                    v18 := GenS3{F0: -1, F1: (v2 ^ (v2 | v2))}
                                    for fi19 := 0; fi19 < 4; fi19++ {
                                        var v20 int16 = v2
                                        v21 := GenS4{F0: (v11 >> ((337) & 63)), F1: 32767, F2: (true && (!((fi10 >> ((2) & 63)) <= -761)))}
                                        _ = fi19
                                        _ = v20
                                        _ = v21
                                    }
                                    _ = v18
                                }
                            case "foo bar":
                                switch "" {
                                case "\n\t":
                                    fmt.Println(v5, v3, v4)
                                case "hello":
                                    var v22 int64 = 1
                                    _ = v22
                                default:
                                    var v23 string = ("\n\t" + ("\n\t" + ""))
                                    _ = v23
                                }
                            case "":
                                fmt.Println(v6, v3)
                            }
                        }
                        switch "a" {
                        case "":
                            v11 -= (v11 * v11)
                        case "hello":
                            for fi24 := 0; fi24 < 20; fi24++ {
                                fmt.Println(v5, v6)
                                if (false && (!true)) {
                                    var v25 int = fi24
                                    v4 += fi10
                                    if (v11 != (v11 ^ (v11 | (v11 >> ((1) & 63))))) {
                                        v5 += (^(v5))
                                        if len(v7) > 0 {
                                            v7[0] = 0
                                        }
                                    }
                                    _ = v25
                                } else {
                                    fmt.Println(v1)
                                    fmt.Println(v6, v4, v2)
                                }
                                var v26 bool = true
                                _ = fi24
                                _ = v26
                            }
                        case "a":
                            for fi27 := 0; fi27 < 9; fi27++ {
                                switch "a" {
                                case "hello":
                                    v28 := GenS2{F0: v3, F1: v2}
                                    _ = v28
                                case "foo bar":
                                    var v29 uint16 = (v6 * v5)
                                    _ = v29
                                default:
                                    if true {
                                        v1 += 1
                                        switch 440 {
                                        case 1:
                                            if true {
                                                if len(v7) > 0 {
                                                    v7[0] = -9223372036854775808
                                                }
                                                v7 = append(v7, 0)
                                            } else {
                                                var v30 int = -225
                                                v31 := GenS3{F0: int64(v30), F1: v2}
                                                fmt.Println(fi10)
                                                _ = v30
                                                _ = v31
                                            }
                                            v2 += (-(v2))
                                        default:
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                        }
                                    } else {
                                        if true {
                                            fmt.Println(v3, v1)
                                            v1 = -128
                                        }
                                    }
                                }
                                _ = fi27
                            }
                        }
                    }
                }
                switch (fi10 >> ((0) & 63)) {
                case -243:
                    v32 := GenS1{F0: 9223372036854775807}
                    _ = v32
                default:
                    v7 = append(v7, 32768)
                }
                v7 = append(v7, (fi10 ^ (^(fi10))))
            } else {
                for fi33 := 0; fi33 < 15; fi33++ {
                    if ((v6 - v6) < (v5 / ((v6) | 1))) {
                        var v34 string = ("foo bar" + "a")
                        _ = v34
                    } else {
                        var v35 bool = (false && false)
                        v5 = v6
                        _ = v35
                    }
                    v7 = append(v7, int(v3))
                    switch "a" {
                    case "hello":
                        v4 = fi33
                    case "":
                        var v36 uint = 0
                        _ = v36
                    case "\n\t":
                        fmt.Println(v4, v1, fi10)
                    default:
                        v37 := GenS5{F0: int8(fi10), F1: "hello"}
                        _ = v37
                    }
                    _ = fi33
                }
                for fi38 := 0; fi38 < 16; fi38++ {
                    var v39 string = "foo bar"
                    v40 := GenS3{F0: -636, F1: v2}
                    v41 := GenS6{F0: 2147483647}
                    switch v39 {
                    case "foo bar":
                        if (!true) {
                            if false {
                                v11 += 42
                                v42 := GenS7{F0: 0, F1: fi38, F2: 2147483647}
                                var v43 float64 = 0.0
                                _ = v42
                                _ = v43
                            }
                        } else {
                            switch (v4 ^ v4) {
                            case 1:
                                var v44 int32 = 0
                                _ = v44
                            case 2147483647:
                                switch v4 {
                                case -386:
                                    fmt.Println(fi10, v6)
                                case -9223372036854775808:
                                    v7 = append(v7, (v4 >> ((18446744073709551615) & 63)))
                                }
                            default:
                                v1 = 0
                            }
                        }
                    case "\n\t":
                        if (!(!true)) {
                            v45 := GenS8{F0: ("a" + v39)}
                            var v46 bool = true
                            v47 := GenS9{F0: (fi10 >> ((uint(v3)) & 63))}
                            _ = v45
                            _ = v46
                            _ = v47
                        } else {
                            v2 = (v2 << ((65535) & 63))
                            var v48 uint = 0
                            _ = v48
                        }
                    default:
                        v6 = 156
                    }
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                }
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v11
        }
        var v49 int = (v4 - -368)
        v50 := GenS10{F0: 127}
        _ = fi10
        _ = v49
        _ = v50
    }
    fmt.Println(v6, v2, v2)
    v51 := GenS8{F0: "\n\t"}
    var v52 float32 = 0.0
    v2 = (v2 - v2)
    switch (("a" + ("\n\t" + "foo bar")) + ("hello" + (("" + "\n\t") + ("hello" + "a")))) {
    case "foo bar":
        for fi53 := 0; fi53 < 8; fi53++ {
            if (65535 == uint(v5)) {
                v1 = (v1 >> ((42) & 63))
            }
            switch "\n\t" {
            case "foo bar":
                v5 -= 65535
            case "a":
                v4 += fi53
            case "":
                var v54 float32 = v52
                _ = v54
            default:
                v55 := GenS11{F0: -2147483648}
                _ = v55
            }
            _ = fi53
        }
    case "hello":
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    var v56 uint = 2
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v57 uint32 = 2147483648
    v58 := GenS1{F0: 9223372036854775807}
    if (!(!(!true))) {
        v4 -= int(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v56
    _ = v57
    _ = v58
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
