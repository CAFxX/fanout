package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int8 }

func main() {
    var v1 int16 = 42
    v2 := []string{("\n\t" + "foo bar")}
    v3 := make(map[uint32]int)
    var v4 float64 = float64(1.5)
    v5 := make(map[string]bool)
    v6 := []uint32{1, 2, 65535, 1}
    v4 = (((1.5 * v4) * float64((1.5 / 0.1))) * (v4 * 1e10))
    fmt.Println(v4, v4)
    fmt.Println(v4, v4)
    var v7 string = ("\n\t" + "hello")
    if len(v2) > 0 {
        v2[0] = v7
    }
    switch ((("a" + v7) + "hello") + "foo bar") {
    case "":
        var v8 bool = ((2147483648 > 0) && (!false))
        _ = v8
    }
    v6 = append(v6, 660)
    v9 := GenS1{F0: v7, F1: v7, F2: 4294967295}
    v6 = append(v6, 2147483647)
    fmt.Println(v1, v4)
    _ = len(v5)
    var v10 bool = true
    if v10 {
        if v10 {
            if v10 {
                switch 9223372036854775807 {
                case -1:
                    if v10 {
                        if v10 {
                            v11 := GenS2{F0: 874}
                            if ((820 == 0) && v10) {
                                var v12 uint8 = 255
                                _ = v12
                            }
                            _ = v11
                        }
                        if v10 {
                            v6 = append(v6, 950)
                        } else {
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                        }
                        v3[42] = 32768
                    } else {
                        for fi13 := 0; fi13 < 10; fi13++ {
                            v2 = append(v2, v7)
                            switch "foo bar" {
                            case "foo bar":
                                var v14 uint16 = 65535
                                _ = v14
                            case "a":
                                var v15 int64 = -539
                                _ = v15
                            case "":
                                v16 := GenS2{F0: 930}
                                _ = v16
                            }
                            _ = fi13
                        }
                        v7 = v7
                        if (42 != 255) {
                            v4 += 3.14159
                            for fi17 := 0; fi17 < 14; fi17++ {
                                var v18 int16 = int16(fi17)
                                _ = fi17
                                _ = v18
                            }
                            if ((!v10) || (!(532 <= 631))) {
                                switch 788 {
                                case -32:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    switch v7 {
                                    case "foo bar":
                                        var v19 uint = 32
                                        _ = v19
                                    case "":
                                        for fi20 := 0; fi20 < 16; fi20++ {
                                            var v21 string = ("\n\t" + (v7 + v7))
                                            v22 := GenS1{F0: "", F1: "hello", F2: 741}
                                            _ = fi20
                                            _ = v21
                                            _ = v22
                                        }
                                    case "hello":
                                        if true {
                                            v7 = v7
                                            v23 := GenS3{F0: int32(v1)}
                                            v24 := GenS1{F0: "a", F1: v7, F2: 42}
                                            _ = v23
                                            _ = v24
                                        }
                                    }
                                }
                                v4 = v4
                                fmt.Println(v10, v4, v7)
                            }
                        } else {
                            if true {
                                v25 := GenS2{F0: 436}
                                v1 = -278
                                _ = v25
                            } else {
                                v1 += v1
                                switch 32767 {
                                case 2147483647:
                                    var v26 bool = v10
                                    _ = v26
                                case 396:
                                    switch ("hello" + v7) {
                                    case "hello":
                                        v7 = (v7 + ((v7 + (v7 + v7)) + ""))
                                    case "\n\t":
                                        v27 := GenS2{F0: 0}
                                        _ = v27
                                    default:
                                        v28 := v5[""]
                                        _ = v28
                                    }
                                }
                                v7 = "foo bar"
                            }
                            switch v7 {
                            case "hello":
                                var v29 int32 = 2147483647
                                _ = v29
                            default:
                                var v30 float32 = 1.5
                                _ = v30
                            }
                            if v10 {
                                var v31 int8 = 1
                                _ = v31
                            }
                        }
                    }
                case 1:
                    for fi32 := 0; fi32 < 18; fi32++ {
                        if (true && v10) {
                            var v33 uint = 1
                            fmt.Println(v4, v33, v4)
                            v4 += v4
                            _ = v33
                        } else {
                            v34 := GenS1{F0: v7, F1: (((v7 + "foo bar") + (((("\n\t" + "") + "foo bar") + v7) + (v7 + v7))) + (v7 + (("" + v7) + v7))), F2: 1}
                            var v35 uint8 = 128
                            if len(v6) > 0 {
                                v6[0] = 453
                            }
                            _ = v34
                            _ = v35
                        }
                        var v36 uint8 = 0
                        _ = fi32
                        _ = v36
                    }
                case 32767:
                    var v37 float32 = ((3.14159 / 1.0) + (0.0 - float32(v4)))
                    _ = v37
                default:
                    for fi38 := 0; fi38 < 4; fi38++ {
                        fmt.Println(v10)
                        switch "\n\t" {
                        case "foo bar":
                            switch ((v7 + (((v7 + v7) + v7) + "")) + ("foo bar" + "\n\t")) {
                            case "\n\t":
                                v10 = true
                            case "":
                                var v39 uint16 = 1
                                _ = v39
                            case "foo bar":
                                v40 := GenS4{F0: -17, F1: 32767, F2: 966}
                                _ = v40
                            }
                        case "":
                            v1 = (v1 << ((4294967295) & 63))
                        case "hello":
                            var v41 int8 = -73
                            _ = v41
                        default:
                            var v42 uint8 = 255
                            _ = v42
                        }
                        fmt.Println(v4, v4, fi38)
                        _ = fi38
                    }
                }
                _ = len(v2)
                v43 := GenS5{F0: 42, F1: 100}
                _ = v43
            }
            fmt.Println(v7)
            var v44 string = "foo bar"
            _ = v44
        } else {
            if (2147483648 > 42) {
                v45 := GenS2{F0: 1}
                _ = v45
            }
            if v10 {
                var v46 int8 = 0
                _ = v46
            }
        }
        for fi47 := 0; fi47 < 1; fi47++ {
            v48 := GenS2{F0: 2}
            _ = fi47
            _ = v48
        }
    } else {
        var v49 int64 = -9223372036854775808
        fmt.Println(v10)
        v4 = v4
        _ = v49
    }
    fmt.Println(v1, v1, v4)
    fmt.Println(v4, v4)
    var v50 uint32 = 0
    var v51 float64 = 0.1
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
    _ = v9
}
