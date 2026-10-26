package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int16; F1 uint16; F2 bool }
type GenS6 struct { F0 int16; F1 uint; F2 bool }
type GenS7 struct { F0 int16 }

func main() {
    var v1 uint64 = 0
    v2 := []int{}
    v3 := make(map[uint32]int)
    var v4 uint64 = (v1 + (v1 + 0))
    var v5 uint = 18446744073709551615
    var v6 uint16 = 2
    switch 622 {
    case 127:
        for fi7 := 0; fi7 < 12; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                if ((false || false) && (!("foo bar" == (("a" + ("foo bar" + "")) + "hello")))) {
                    v9 := GenS1{F0: "hello"}
                    delete(v3, 2)
                    _ = v9
                } else {
                    if (("a" < "foo bar") || (1.5 != (-2.25 / (3.14159 - float64(1.5))))) {
                        switch "" {
                        case "hello":
                            _ = len(v3)
                        case "foo bar":
                            v10 := GenS2{F0: fi7}
                            v11 := GenS3{F0: (v6 ^ (v6 * 516)), F1: (fi8 * 0), F2: 42}
                            fmt.Println(v1, fi7)
                            if ((^(v4)) == (^(v1))) {
                                v12 := v3[2147483647]
                                fmt.Println(v6)
                                v12 -= 155
                                switch "foo bar" {
                                case "":
                                    switch fi8 {
                                    case 9223372036854775807:
                                        if false {
                                            v13 := GenS2{F0: 9223372036854775807}
                                            _ = v13
                                        } else {
                                            var v14 int32 = 0
                                            var v15 int16 = -540
                                            _ = v14
                                            _ = v15
                                        }
                                    case 42:
                                        v16 := GenS4{F0: 2, F1: (fi7 << ((uint(v1)) & 63))}
                                        _ = v16
                                    case 197:
                                        v17 := GenS5{F0: -32768, F1: (v6 << ((v5) & 63)), F2: true}
                                        _ = v17
                                    default:
                                        var v18 int64 = -9223372036854775808
                                        _ = v18
                                    }
                                case "\n\t":
                                    delete(v3, uint32(v6))
                                }
                                v6 = (v6 << ((v5) & 63))
                                fmt.Println(fi8)
                                var v19 bool = false
                                _ = v12
                                _ = v19
                            } else {
                                v4 = (v4 / ((v1) | 1))
                                if len(v2) > 0 {
                                    v2[0] = -720
                                }
                                var v20 uint16 = (^(v6))
                                _ = v20
                            }
                            _ = v10
                            _ = v11
                        }
                        v21 := GenS4{F0: 241, F1: (fi7 % (((-(fi8))) | 1))}
                        _ = v21
                    } else {
                        if (false && (false && (-87 >= 127))) {
                            if true {
                                fmt.Println(v1)
                                var v22 bool = true
                                var v23 uint = (^(v5))
                                _ = v22
                                _ = v23
                            } else {
                                if (!true) {
                                    v5 = (v5 << (((v5 & v5)) & 63))
                                }
                            }
                        }
                        if true {
                            delete(v3, 361)
                            v4 += uint64(v5)
                            v24 := GenS6{F0: 42, F1: (v5 >> ((18446744073709551615) & 63)), F2: (!false)}
                            _ = v24
                        }
                        var v25 uint32 = uint32(v1)
                        _ = v25
                    }
                }
                if false {
                    _ = len(v2)
                    if true {
                        v1 = (^(v4))
                        v3[1] = (fi8 - (fi7 + 32768))
                        switch -2147483648 {
                        case 42:
                            v1 += 42
                            v6 = (v6 % (((v6 * (v6 + 32767))) | 1))
                            v4 -= (v4 ^ 1)
                        case -211:
                            v5 += (v5 << (((v5 >> (((^(v5))) & 63))) & 63))
                        }
                    }
                    v2 = append(v2, 32767)
                    v4 -= v1
                    var v26 int64 = 1
                    _ = v26
                }
                if (!(1 == -522)) {
                    var v27 string = (("a" + (("foo bar" + "hello") + "\n\t")) + "\n\t")
                    _ = v27
                } else {
                    var v28 float32 = -2.25
                    if ((32767 != v6) || false) {
                        v6 += (v6 | 32767)
                        switch (("\n\t" + ("hello" + ("a" + ("hello" + ("\n\t" + ("\n\t" + "a")))))) + (("foo bar" + ("" + "foo bar")) + "a")) {
                        case "\n\t":
                            if false {
                                var v29 int16 = -1
                                v30 := GenS1{F0: (("hello" + ("" + "a")) + "hello")}
                                _ = v29
                                _ = v30
                            }
                        default:
                            v31 := GenS5{F0: -1, F1: (^(v6)), F2: false}
                            _ = v31
                        }
                    }
                    _ = v28
                }
                _ = fi8
            }
            if (true || true) {
                fmt.Println(v6)
                v3[933] = fi7
            } else {
                v32 := GenS5{F0: 32767, F1: (^(v6)), F2: (100 != 127)}
                v6 += v6
                _ = v32
            }
            _ = fi7
        }
    case -9223372036854775808:
        if true {
            var v33 int8 = 97
            v4 = v4
            var v34 uint64 = (v4 - v4)
            _ = v33
            _ = v34
        } else {
            fmt.Println(v1, v5, v1)
            delete(v3, 1)
        }
    }
    if true {
        v35 := GenS4{F0: 4294967295, F1: 128}
        var v36 string = "hello"
        _ = v35
        _ = v36
    }
    switch "a" {
    case "hello":
        if len(v2) > 0 {
            v2[0] = 1
        }
    }
    v6 = (v6 ^ v6)
    var v37 int8 = -1
    if (false || (!true)) {
        v1 -= 418
        for fi38 := 0; fi38 < 18; fi38++ {
            var v39 int = (fi38 >> ((2) & 63))
            _ = fi38
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        for fi41 := 0; fi41 < 9; fi41++ {
            v37 = v37
            v2 = append(v2, (fi40 - (fi41 + fi40)))
            _ = fi41
        }
        _ = fi40
    }
    v42 := GenS6{F0: 32767, F1: 303, F2: true}
    v37 -= v37
    var v43 uint = v5
    for fi44 := 0; fi44 < 3; fi44++ {
        if ((-1 >= -2147483648) || true) {
            _ = len(v2)
            v45 := GenS7{F0: -32768}
            _ = v45
        } else {
            for fi46 := 0; fi46 < 10; fi46++ {
                v37 = (v37 ^ v37)
                fmt.Println(v37, v5)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi46
            }
            var v47 int8 = v37
            _ = v47
        }
        fmt.Println(v43, v1)
        _ = fi44
    }
    var v48 int64 = -924
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v42
    _ = v43
    _ = v48
    _ = v5
    _ = v6
}
