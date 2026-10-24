package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 int8; F1 bool; F2 int8 }

func main() {
    v1 := []string{("foo bar" + (("hello" + ((("hello" + ("\n\t" + "hello")) + "foo bar") + ((("a" + "a") + "hello") + "\n\t"))) + ("" + "a")))}
    var v2 int64 = 784
    var v3 bool = true
    var v4 string = ("hello" + ("a" + ("" + "hello")))
    var v5 int16 = 32767
    var v6 int8 = -1
    v1 = append(v1, v4)
    switch v4 {
    case "":
        fmt.Println(v4, v5)
    case "hello":
        fmt.Println(v5, v5)
    case "\n\t":
        v4 = v4
    default:
        v7 := GenS1{F0: 0, F1: 917}
        _ = v7
    }
    v6 = (v6 + v6)
    fmt.Println(v3, v4)
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(v5)
        _ = fi8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int32 = 1
        switch fi9 {
        case 32768:
            v2 = (v2 >> ((1) & 63))
        }
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 18446744073709551615, F1: 765}
    fmt.Println(v3, v5, v6)
    v4 = ("foo bar" + v4)
    v12 := GenS1{F0: 1, F1: 227}
    switch v4 {
    case "":
        v3 = (-17 == -695)
    case "hello":
        v2 = (-(v2))
    case "foo bar":
        v2 += v2
    default:
        v1 = append(v1, ("" + v4))
    }
    switch 997 {
    case -676:
        v13 := GenS2{F0: true, F1: 944}
        _ = v13
    case -2147483648:
        if len(v1) > 0 {
            v1[0] = v4
        }
    default:
        var v14 uint = 934
        _ = v14
    }
    switch (v4 + v4) {
    case "":
        v5 = -350
    case "\n\t":
        switch (v4 + "hello") {
        case "":
            v15 := GenS3{F0: v4}
            _ = v15
        default:
            var v16 string = v4
            _ = v16
        }
    }
    switch v4 {
    case "hello":
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 int = fi17
            _ = fi17
            _ = v18
        }
    case "":
        var v19 bool = v3
        _ = v19
    case "a":
        var v20 int64 = (v2 / ((v2) | 1))
        _ = v20
    default:
        v21 := GenS1{F0: 42, F1: 0}
        _ = v21
    }
    if false {
        var v22 string = v4
        fmt.Println(v4)
        var v23 float64 = 0.1
        _ = v22
        _ = v23
    }
    if (v5 >= v5) {
        if v3 {
            v24 := GenS4{F0: -1}
            var v25 int16 = (v5 * 1)
            switch v4 {
            case "a":
                if v3 {
                    switch v4 {
                    case "foo bar":
                        var v26 uint16 = 2
                        _ = v26
                    case "":
                        v27 := GenS5{F0: 98}
                        _ = v27
                    case "hello":
                        if v3 {
                            for fi28 := 0; fi28 < 4; fi28++ {
                                var v29 uint64 = 937
                                _ = fi28
                                _ = v29
                            }
                            switch 2147483647 {
                            case 803:
                                var v30 int8 = (v6 * (v6 << ((384) & 63)))
                                _ = v30
                            }
                            if v3 {
                                var v31 bool = v3
                                if v31 {
                                    if false {
                                        _ = len(v1)
                                    } else {
                                        if true {
                                            var v32 float32 = (0.0 / ((1.5 - 0.1) / 0.1))
                                            _ = v32
                                        }
                                        v1 = append(v1, v4)
                                        v3 = (v25 <= (v25 << ((4294967295) & 63)))
                                    }
                                    var v33 uint32 = 704
                                    switch 128 {
                                    case 2147483647:
                                        for fi34 := 0; fi34 < 16; fi34++ {
                                            var v35 bool = true
                                            var v36 int16 = v5
                                            _ = fi34
                                            _ = v35
                                            _ = v36
                                        }
                                    case 709:
                                        v6 -= (v6 >> ((42) & 63))
                                    }
                                    _ = v33
                                } else {
                                    var v37 uint32 = 4294967295
                                    v3 = v31
                                    switch -592 {
                                    case -9223372036854775808:
                                        switch v4 {
                                        case "hello":
                                            var v38 uint = 2
                                            _ = v38
                                        case "\n\t":
                                            for fi39 := 0; fi39 < 6; fi39++ {
                                                fmt.Println(v5, v5, v3)
                                                fmt.Println(v3)
                                                if ((-(v2)) == v2) {
                                                    var v40 string = (((v4 + v4) + v4) + v4)
                                                    _ = v40
                                                }
                                                _ = fi39
                                            }
                                        }
                                    case 32767:
                                        if ((1 != (v37 & 42)) || v3) {
                                            if len(v1) > 0 {
                                                v1[0] = v4
                                            }
                                            v5 = 42
                                        } else {
                                            v41 := GenS6{F0: 1000, F1: 42}
                                            _ = v41
                                        }
                                    case 42:
                                        var v42 uint32 = v37
                                        _ = v42
                                    }
                                    _ = v37
                                }
                                _ = v31
                            } else {
                                var v43 uint8 = 166
                                _ = v43
                            }
                        }
                    default:
                        v1 = append(v1, v4)
                    }
                    for fi44 := 0; fi44 < 14; fi44++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (634 <= 356) {
                            var v45 float32 = 0.0
                            _ = v45
                        } else {
                            v2 -= (v2 & 9223372036854775807)
                        }
                        var v46 bool = (v3 && true)
                        _ = fi44
                        _ = v46
                    }
                    if v3 {
                        v47 := GenS5{F0: 128}
                        v25 = (v5 & v5)
                        _ = v47
                    }
                } else {
                    if true {
                        v3 = (v4 >= v4)
                        var v48 int = 2147483647
                        _ = v48
                    } else {
                        fmt.Println(v2, v25, v2)
                        var v49 int64 = (v2 << ((283) & 63))
                        _ = v49
                    }
                }
            case "\n\t":
                v25 = -1
            case "":
                var v50 uint16 = 1
                _ = v50
            default:
                for fi51 := 0; fi51 < 12; fi51++ {
                    v25 -= 42
                    _ = fi51
                }
            }
            _ = v24
            _ = v25
        } else {
            var v52 uint32 = 363
            if (-1 == 674) {
                var v53 int = 855
                for fi54 := 0; fi54 < 7; fi54++ {
                    v1 = append(v1, "")
                    v55 := GenS4{F0: v53}
                    var v56 uint8 = 2
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                v57 := GenS7{F0: int8(v53), F1: false, F2: 0}
                _ = v53
                _ = v57
            } else {
                v58 := GenS7{F0: v6, F1: v3, F2: 42}
                var v59 int = 9223372036854775807
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v58
                _ = v59
            }
            var v60 int = 4294967295
            _ = v52
            _ = v60
        }
    }
    v4 = ((v4 + v4) + ("hello" + "foo bar"))
    v61 := GenS5{F0: 97}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
