package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int8 = 42
    var v2 int8 = v1
    v3 := []string{("a" + (("hello" + "\n\t") + "\n\t")), (((("hello" + (("\n\t" + "hello") + "a")) + "foo bar") + "hello") + "hello"), "hello", "hello"}
    v4 := []string{("foo bar" + "foo bar"), "\n\t", "\n\t", "a"}
    var v5 float32 = 1.5
    var v6 uint32 = 65535
    v7 := GenS1{F0: 42}
    var v8 int16 = -1
    var v9 string = "hello"
    var v10 float64 = 0.0
    if true {
        if true {
            switch v9 {
            case "a":
                fmt.Println(v9, v5, v1)
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            switch 32767 {
            case -2147483648:
                v11 := GenS1{F0: 18446744073709551615}
                _ = v11
            case 9223372036854775807:
                v9 = "hello"
            case 1:
                v12 := GenS1{F0: 780}
                _ = v12
            default:
                if true {
                    var v13 int64 = 9223372036854775807
                    _ = v13
                }
            }
        } else {
            v5 = v5
            v4 = append(v4, ("" + ((v9 + ("hello" + (v9 + "hello"))) + (v9 + v9))))
            v14 := GenS1{F0: 1}
            _ = v14
        }
    } else {
        var v15 float64 = v10
        fmt.Println(v8, v2)
        _ = v15
    }
    v16 := GenS1{F0: 776}
    v17 := GenS2{F0: 1}
    switch "\n\t" {
    case "\n\t":
        switch v9 {
        case "foo bar":
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = len(v4)
                for fi18 := 0; fi18 < 18; fi18++ {
                    fmt.Println(v10, v8, v1)
                    for fi19 := 0; fi19 < 6; fi19++ {
                        switch "hello" {
                        case "a":
                            v1 = v1
                            v9 = (("" + v9) + ("foo bar" + ""))
                            if false {
                                var v20 float32 = v5
                                v6 = (v6 ^ 149)
                                v21 := GenS3{F0: 217, F1: -749}
                                if false {
                                    v5 = v20
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v22 int8 = (v1 >> ((969) & 63))
                                    v6 = 406
                                    v23 := GenS1{F0: 42}
                                    _ = v22
                                    _ = v23
                                }
                                _ = v20
                                _ = v21
                            }
                        case "\n\t":
                            if true {
                                var v24 string = "hello"
                                v25 := GenS1{F0: 0}
                                _ = v24
                                _ = v25
                            }
                        default:
                            if true {
                                var v26 int64 = -1
                                if false {
                                    switch (("\n\t" + ((v9 + v9) + "hello")) + v9) {
                                    case "\n\t":
                                        v8 = int16(v1)
                                    }
                                    v26 = v26
                                    if false {
                                        if false {
                                            if len(v4) > 0 {
                                                v4[0] = "\n\t"
                                            }
                                            v27 := GenS1{F0: 660}
                                            var v28 bool = true
                                            _ = v27
                                            _ = v28
                                        } else {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        v6 = (v6 << ((1) & 63))
                                    } else {
                                        fmt.Println(v10, v1, v9)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                }
                                _ = v26
                            } else {
                                var v29 int8 = (v1 >> ((0) & 63))
                                _ = v29
                            }
                        }
                        _ = fi19
                    }
                    switch (fi18 | 0) {
                    case 127:
                        var v30 bool = ((v1 - v1) != v1)
                        _ = v30
                    case 136:
                        var v31 float64 = float64((v5 * v5))
                        _ = v31
                    }
                    _ = fi18
                }
            } else {
                v10 = v10
                if false {
                    fmt.Println(v2, v10, v6)
                    switch 42 {
                    case 437:
                        var v32 uint32 = v6
                        _ = v32
                    default:
                        var v33 float64 = (float64(v5) * v10)
                        _ = v33
                    }
                }
                fmt.Println(v9, v6)
            }
        case "a":
            for fi34 := 0; fi34 < 6; fi34++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if false {
                    switch (fi34 | (fi34 / ((-1) | 1))) {
                    case 127:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case -2147483648:
                        v6 = 2147483647
                    case 42:
                        var v35 uint = 65535
                        _ = v35
                    default:
                        if (-2.25 < v5) {
                            var v36 bool = false
                            for fi37 := 0; fi37 < 17; fi37++ {
                                v38 := GenS3{F0: 42, F1: 2147483647}
                                _ = len(v3)
                                _ = fi37
                                _ = v38
                            }
                            v39 := GenS4{F0: (fi34 + 9223372036854775807), F1: (v6 / ((v6) | 1)), F2: 584}
                            _ = v36
                            _ = v39
                        }
                    }
                }
                var v40 uint32 = (^(v6))
                _ = fi34
                _ = v40
            }
        case "hello":
            _ = len(v4)
        default:
            switch -49 {
            case 476:
                v4 = append(v4, "")
            case 32767:
                for fi41 := 0; fi41 < 1; fi41++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v5 += 3.14159
                    v8 -= (^(v8))
                    _ = fi41
                }
            case 1:
                v42 := GenS5{F0: 0}
                _ = v42
            }
        }
    case "foo bar":
        for fi43 := 0; fi43 < 12; fi43++ {
            v9 = "hello"
            if (int64(v6) == -9223372036854775808) {
                for fi44 := 0; fi44 < 7; fi44++ {
                    v6 += 2147483648
                    if len(v4) > 0 {
                        v4[0] = v9
                    }
                    _ = fi44
                }
                switch fi43 {
                case 659:
                    v9 = (v9 + ("\n\t" + v9))
                default:
                    v10 = 0.1
                }
                for fi45 := 0; fi45 < 12; fi45++ {
                    if len(v4) > 0 {
                        v4[0] = v9
                    }
                    v6 = v6
                    _ = fi45
                }
            }
            v1 += 1
            _ = fi43
        }
    case "a":
        var v46 uint16 = 0
        _ = v46
    default:
        if true {
            v5 -= v5
            for fi47 := 0; fi47 < 5; fi47++ {
                fmt.Println(v9)
                var v48 float64 = (v10 + (0.0 + v10))
                _ = fi47
                _ = v48
            }
            var v49 uint8 = 217
            _ = v49
        } else {
            switch -804 {
            case 32767:
                v1 = 100
            case 364:
                var v50 uint32 = 65535
                _ = v50
            default:
                v2 = v1
            }
        }
    }
    for fi51 := 0; fi51 < 12; fi51++ {
        var v52 float64 = 3.14159
        _ = fi51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
