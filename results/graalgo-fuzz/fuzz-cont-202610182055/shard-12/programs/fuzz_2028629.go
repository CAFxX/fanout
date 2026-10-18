package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint32; F1 string }

func main() {
    var v1 uint16 = 255
    var v2 uint32 = 1
    v3 := []string{("a" + "a"), "a", ("hello" + "foo bar"), (("foo bar" + ("hello" + "hello")) + "a")}
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v4 := GenS1{F0: 216, F1: 2}
    v1 = v1
    fmt.Println(v2)
    fmt.Println(v1)
    if (-2.25 >= (0.1 + 0.1)) {
        v2 = v2
        var v5 uint64 = 416
        _ = v5
    }
    var v6 string = ""
    switch ("a" + "foo bar") {
    case "":
        if true {
            var v7 float32 = (1.5 - 3.14159)
            for fi8 := 0; fi8 < 19; fi8++ {
                if false {
                    if false {
                        v9 := GenS2{F0: (^(v2))}
                        _ = v9
                    } else {
                        var v10 int = fi8
                        switch -636 {
                        case 2147483647:
                            v6 = v6
                        case 652:
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                        case 9223372036854775807:
                            switch fi8 {
                            case 127:
                                for fi11 := 0; fi11 < 14; fi11++ {
                                    v12 := GenS2{F0: 226}
                                    fmt.Println(fi8)
                                    v13 := GenS1{F0: 0, F1: 128}
                                    _ = fi11
                                    _ = v12
                                    _ = v13
                                }
                            case 42:
                                v2 = (v2 - (^(v2)))
                            case -159:
                                if true {
                                    var v14 int16 = 1
                                    switch "" {
                                    case "\n\t":
                                        var v15 int8 = 1
                                        _ = v15
                                    case "a":
                                        v16 := GenS2{F0: (^(v2))}
                                        _ = v16
                                    case "":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    v17 := GenS2{F0: (v2 ^ (v2 + v2))}
                                    _ = v14
                                    _ = v17
                                }
                            }
                        }
                        for fi18 := 0; fi18 < 1; fi18++ {
                            var v19 int64 = -274
                            _ = fi18
                            _ = v19
                        }
                        _ = v10
                    }
                    v1 -= v1
                    fmt.Println(v2)
                }
                _ = fi8
            }
            for fi20 := 0; fi20 < 13; fi20++ {
                v1 += (v1 >> ((928) & 63))
                if len(v3) > 0 {
                    v3[0] = ("hello" + v6)
                }
                var v21 int64 = 0
                _ = fi20
                _ = v21
            }
            _ = v7
        } else {
            for fi22 := 0; fi22 < 12; fi22++ {
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 uint = 42
                    v25 := GenS1{F0: 206, F1: 128}
                    v26 := GenS2{F0: uint32(v24)}
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                v1 = v1
                var v27 string = v6
                _ = fi22
                _ = v27
            }
        }
    default:
        if true {
            switch ("\n\t" + (v6 + ("foo bar" + "\n\t"))) {
            case "":
                for fi28 := 0; fi28 < 13; fi28++ {
                    v29 := GenS1{F0: 215, F1: 121}
                    _ = fi28
                    _ = v29
                }
            case "\n\t":
                var v30 uint16 = 2
                _ = v30
            case "foo bar":
                fmt.Println(v2, v6)
            }
            fmt.Println(v1)
            if true {
                switch v6 {
                case "\n\t":
                    fmt.Println(v1)
                case "foo bar":
                    v31 := GenS3{F0: 9223372036854775807, F1: 54, F2: 9223372036854775807}
                    _ = v31
                case "a":
                    var v32 uint16 = v1
                    _ = v32
                default:
                    v33 := GenS4{F0: 42, F1: 18446744073709551615}
                    _ = v33
                }
                var v34 uint64 = 2
                _ = v34
            } else {
                if true {
                    var v35 uint = 4294967295
                    v2 -= v2
                    _ = v35
                } else {
                    v2 -= (v2 - v2)
                    v6 = "\n\t"
                }
                v36 := GenS3{F0: -2147483648, F1: 42, F2: -345}
                _ = v36
            }
        }
    }
    if true {
        var v37 float32 = ((float32(0.1) * (3.14159 * ((0.1 + float32(0.1)) + ((((1.5 + 1e10) + 1.5) * 0.0) / (1e10 / 3.14159))))) / ((float32(3.14159) - -2.25) * -2.25))
        _ = v37
    } else {
        for fi38 := 0; fi38 < 8; fi38++ {
            switch fi38 {
            case 932:
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
            case 127:
                v6 = (v6 + v6)
            default:
                var v39 int16 = -1
                _ = v39
            }
            v2 -= (v2 + 1)
            var v40 float64 = 1e10
            _ = fi38
            _ = v40
        }
    }
    switch -1 {
    case 42:
        var v41 bool = (0 >= -315)
        _ = v41
    case 0:
        v42 := GenS3{F0: 2147483647, F1: 2, F2: -9223372036854775808}
        _ = v42
    case 280:
        v2 = uint32(v1)
    default:
        for fi43 := 0; fi43 < 4; fi43++ {
            v1 -= v1
            if (((true && false) && (true || (1 < 127))) && false) {
                v44 := GenS5{F0: v2, F1: "a"}
                v3 = append(v3, "\n\t")
                for fi45 := 0; fi45 < 16; fi45++ {
                    if (true && (65535 == 4294967295)) {
                        if (-336 > 1) {
                            var v46 bool = false
                            _ = v46
                        }
                        var v47 uint32 = v2
                        _ = v47
                    } else {
                        var v48 int16 = 0
                        var v49 bool = true
                        v3 = append(v3, "\n\t")
                        _ = v48
                        _ = v49
                    }
                    var v50 int16 = -501
                    var v51 uint = 671
                    _ = fi45
                    _ = v50
                    _ = v51
                }
                _ = v44
            }
            _ = fi43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
