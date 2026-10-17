package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int64; F2 int8 }
type GenS3 struct { F0 uint32; F1 string; F2 uint }

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 467
    var v3 int8 = 1
    var v4 int = 127
    v5 := []string{"\n\t", ("\n\t" + ("foo bar" + "hello")), (("a" + "foo bar") + "a")}
    var v6 uint = 4294967295
    fmt.Println(v3, v2)
    v2 -= (v2 / ((-2147483648) | 1))
    fmt.Println(v3, v4)
    if ("a" >= (("a" + "a") + "foo bar")) {
        if (0 == 77) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    fmt.Println(v4)
    fmt.Println(v3, v4, v4)
    if true {
        v4 -= -1
        delete(v1, "\n\t")
    }
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 19; fi8++ {
        v7 = (v7 << ((v6) & 63))
        fmt.Println(v7, v3)
        v6 = v6
        _ = fi8
    }
    var v9 uint64 = 2
    switch 297 {
    case -946:
        fmt.Println(v2, v3)
    case 32767:
        v10 := GenS1{F0: uint8(v6)}
        _ = v10
    case 0:
        v2 += v2
    }
    switch "" {
    case "\n\t":
        var v11 uint32 = 358
        _ = v11
    case "foo bar":
        if false {
            var v12 bool = (!("foo bar" > "a"))
            for fi13 := 0; fi13 < 7; fi13++ {
                switch fi13 {
                case 1:
                    if false {
                        switch "a" {
                        case "":
                            v14 := v1["foo bar"]
                            _ = v14
                        }
                        v15 := v1["hello"]
                        _ = v15
                    }
                default:
                    var v16 int = v4
                    _ = v16
                }
                _ = fi13
            }
            _ = v12
        } else {
            fmt.Println(v9, v9)
            switch v4 {
            case 42:
                v17 := GenS1{F0: 101}
                _ = v17
            case 127:
                if true {
                    var v18 float32 = 3.14159
                    if false {
                        v2 = (v2 << ((v6) & 63))
                        v19 := GenS2{F0: 935, F1: 42, F2: v3}
                        _ = v19
                    } else {
                        v7 -= uint16(v9)
                        v3 = v3
                        switch ("foo bar" + "foo bar") {
                        case "hello":
                            if false {
                                var v20 int64 = 1
                                switch (v4 - (v4 / ((215) | 1))) {
                                case 32767:
                                    v7 = 255
                                default:
                                    if false {
                                        v9 += (v9 % ((18446744073709551615) | 1))
                                        fmt.Println(v7, v9)
                                    } else {
                                        v5 = append(v5, "hello")
                                        v2 += (-(v2))
                                    }
                                }
                                for fi21 := 0; fi21 < 19; fi21++ {
                                    for fi22 := 0; fi22 < 2; fi22++ {
                                        _ = len(v5)
                                        delete(v1, ("a" + ""))
                                        _ = fi22
                                    }
                                    var v23 string = ""
                                    _ = fi21
                                    _ = v23
                                }
                                _ = v20
                            } else {
                                var v24 int32 = (^(v2))
                                _ = v24
                            }
                        case "a":
                            v9 = v9
                        case "\n\t":
                            var v25 float32 = v18
                            _ = v25
                        default:
                            v1[("a" + "hello")] = v4
                        }
                    }
                    _ = v18
                }
            case 574:
                switch "foo bar" {
                case "":
                    v26 := GenS1{F0: 1}
                    _ = v26
                case "hello":
                    var v27 int8 = (v3 & (v3 | (v3 | (v3 | int8(v6)))))
                    _ = v27
                }
            default:
                var v28 float32 = 3.14159
                _ = v28
            }
            switch -282 {
            case 32768:
                if true {
                    fmt.Println(v4, v3, v7)
                } else {
                    fmt.Println(v7)
                    fmt.Println(v2, v2, v2)
                }
            default:
                for fi29 := 0; fi29 < 20; fi29++ {
                    for fi30 := 0; fi30 < 4; fi30++ {
                        if false {
                            var v31 uint8 = 42
                            if len(v5) > 0 {
                                v5[0] = ""
                            }
                            _ = v31
                        }
                        v9 -= v9
                        v9 += (v9 / ((363) | 1))
                        _ = fi30
                    }
                    _ = fi29
                }
            }
        }
    case "a":
        for fi32 := 0; fi32 < 3; fi32++ {
            for fi33 := 0; fi33 < 6; fi33++ {
                var v34 uint8 = uint8(v9)
                switch (("hello" + "\n\t") + "") {
                case "\n\t":
                    fmt.Println(v2, v6)
                case "foo bar":
                    var v35 string = ("\n\t" + "hello")
                    _ = v35
                case "hello":
                    if (0 < 528) {
                        v7 = v7
                        var v36 uint64 = 496
                        _ = v36
                    } else {
                        if len(v5) > 0 {
                            v5[0] = ("foo bar" + ("foo bar" + ""))
                        }
                    }
                default:
                    v9 = v9
                }
                v34 = v34
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
    default:
        var v37 float64 = -2.25
        _ = v37
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        v4 = (v4 / ((v4) | 1))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi38
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        var v40 int16 = 794
        _ = fi39
        _ = v40
    }
    v41 := GenS3{F0: 42, F1: (("foo bar" + "hello") + ("" + "a")), F2: (^(v6))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
