package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int64; F1 int32; F2 int16 }
type GenS3 struct { F0 string; F1 uint; F2 int8 }
type GenS4 struct { F0 int8; F1 uint8; F2 int }
type GenS5 struct { F0 bool; F1 int16; F2 uint32 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 bool }
func genFn0(p0 int64) int64 {
    var v5 int32 = 42
    var v6 uint = 65535
    if false {
        v6 = 18446744073709551615
    }
    v6 = v6
    fmt.Println(v5, v6, v6)
    _ = p0
    _ = v5
    _ = v6
    return (p0 >> (((v6 | v6)) & 63))
}
func genFn1(p0 bool, p1 uint32, p2 bool) uint {
    v7 := []int{-571}
    v8 := GenS1{F0: ((("" + ("" + ("hello" + ("" + "foo bar")))) + ("a" + "\n\t")) + "foo bar"), F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := []string{"hello", "a", "a"}
    var v2 int16 = -212
    var v3 string = (("a" + "hello") + "a")
    var v4 int8 = 127
    v2 = v2
    var v9 float64 = 0.1
    v4 = 100
    for fi10 := 0; fi10 < 8; fi10++ {
        if true {
            switch (fi10 * (fi10 >> ((255) & 63))) {
            case 32768:
                v9 = 1.5
            case 127:
                v9 = (1.5 * v9)
            }
            v11 := GenS2{F0: 1, F1: 1, F2: v2}
            _ = v11
        }
        _ = fi10
    }
    if true {
        if true {
            v3 = (v3 + v3)
            v1 = append(v1, v3)
        } else {
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS3{F0: "foo bar", F1: 738, F2: (v4 / ((-128) | 1))}
                v14 := GenS3{F0: "a", F1: 0, F2: 42}
                v9 -= v9
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    } else {
        switch v3 {
        case "\n\t":
            switch -70 {
            case 614:
                switch 920 {
                case -9223372036854775808:
                    if (!false) {
                        for fi15 := 0; fi15 < 4; fi15++ {
                            fmt.Println(v9, v9)
                            _ = fi15
                        }
                        v16 := GenS4{F0: (v4 | v4), F1: 1, F2: 32768}
                        _ = v16
                    } else {
                        v17 := GenS1{F0: "hello", F1: v2}
                        _ = v17
                    }
                case -296:
                    v9 -= v9
                case -1:
                    if (32767 == -1) {
                        for fi18 := 0; fi18 < 4; fi18++ {
                            if len(v1) > 0 {
                                v1[0] = v3
                            }
                            if false {
                                switch v3 {
                                case "hello":
                                    var v19 int64 = 42
                                    _ = v19
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                }
                                var v20 uint = 18446744073709551615
                                _ = v20
                            }
                            v21 := GenS5{F0: ((v2 + v2) <= (v2 - int16(v4))), F1: v2, F2: 1}
                            _ = fi18
                            _ = v21
                        }
                    } else {
                        fmt.Println(v3, v4, v3)
                        for fi22 := 0; fi22 < 17; fi22++ {
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            var v23 uint64 = 0
                            for fi24 := 0; fi24 < 15; fi24++ {
                                _ = len(v1)
                                var v25 uint16 = 2
                                v26 := GenS4{F0: (v4 / ((v4) | 1)), F1: 127, F2: 0}
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            _ = fi22
                            _ = v23
                        }
                    }
                default:
                    switch "\n\t" {
                    case "":
                        switch 128 {
                        case -875:
                            _ = len(v1)
                        case -267:
                            fmt.Println(v2)
                        case 597:
                            v27 := GenS5{F0: false, F1: (-(v2)), F2: 283}
                            _ = v27
                        default:
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                        }
                    case "foo bar":
                        if (!(216 > -975)) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch v3 {
                            case "a":
                                v28 := GenS6{F0: 2, F1: 32767}
                                _ = v28
                            default:
                                var v29 uint32 = 1
                                _ = v29
                            }
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            for fi30 := 0; fi30 < 3; fi30++ {
                                var v31 uint8 = 0
                                v1 = append(v1, v3)
                                _ = fi30
                                _ = v31
                            }
                        }
                    case "\n\t":
                        switch (v3 + "foo bar") {
                        case "foo bar":
                            switch ("foo bar" + ("" + v3)) {
                            case "foo bar":
                                for fi32 := 0; fi32 < 13; fi32++ {
                                    switch (v3 + "foo bar") {
                                    case "foo bar":
                                        for fi33 := 0; fi33 < 10; fi33++ {
                                            v3 = ((("a" + "") + v3) + v3)
                                            _ = fi33
                                        }
                                    }
                                    _ = fi32
                                }
                            case "\n\t":
                                fmt.Println(v4, v9, v9)
                            }
                        case "":
                            var v34 uint64 = 0
                            _ = v34
                        default:
                            var v35 uint32 = 1
                            _ = v35
                        }
                    }
                }
            }
        case "":
            var v36 uint32 = 0
            _ = v36
        case "foo bar":
            v37 := GenS7{F0: -2147483648}
            _ = v37
        }
        if true {
            var v38 int16 = v2
            _ = v38
        } else {
            v39 := GenS8{F0: 2147483648, F1: 42}
            var v40 uint32 = 1
            _ = v39
            _ = v40
        }
    }
    v1 = append(v1, (("\n\t" + ("" + v3)) + v3))
    if len(v1) > 0 {
        v1[0] = ((v3 + v3) + ("hello" + v3))
    }
    fmt.Println(v9, v3)
    var v41 float64 = 3.14159
    for fi42 := 0; fi42 < 4; fi42++ {
        v43 := GenS9{F0: (-2.25 > ((3.14159 + 0.0) + (0.0 + 0.1)))}
        _ = fi42
        _ = v43
    }
    if (!false) {
        v4 = (v4 & v4)
        v41 -= v9
    } else {
        switch "\n\t" {
        case "hello":
            v9 -= (v41 / v9)
        case "\n\t":
            for fi44 := 0; fi44 < 7; fi44++ {
                if false {
                    v45 := GenS7{F0: 1}
                    var v46 int32 = -1
                    _ = v45
                    _ = v46
                } else {
                    switch (fi44 % ((fi44) | 1)) {
                    case 4294967295:
                        v3 = ""
                    case 0:
                        var v47 uint16 = 65535
                        _ = v47
                    }
                    var v48 int8 = -128
                    _ = v48
                }
                for fi49 := 0; fi49 < 5; fi49++ {
                    _ = len(v1)
                    if false {
                        v2 = v2
                    }
                    _ = fi49
                }
                var v50 int64 = 0
                _ = fi44
                _ = v50
            }
        default:
            for fi51 := 0; fi51 < 15; fi51++ {
                _ = len(v1)
                v9 += v41
                _ = fi51
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v52 float32 = (0.0 / -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v52:", v52)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v52
    _ = v9
}
