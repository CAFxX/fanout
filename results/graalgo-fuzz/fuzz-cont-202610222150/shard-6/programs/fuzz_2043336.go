package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint; F1 uint16; F2 string }
type GenS3 struct { F0 bool; F1 int64; F2 int8 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int32; F1 int8 }

func main() {
    v1 := []uint32{1, 1, 894, 630}
    var v2 uint8 = 255
    v3 := make(map[int]bool)
    v4 := make(map[string]bool)
    var v5 int64 = 0
    v6 := make(map[string]int)
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint8 = 1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 65535, F1: int16(v2)}
    fmt.Println(v2, v2)
    if false {
        if true {
            if true {
                for fi10 := 0; fi10 < 4; fi10++ {
                    switch "" {
                    case "foo bar":
                        v11 := GenS1{F0: 0, F1: 1}
                        _ = v11
                    }
                    fmt.Println(v5)
                    _ = fi10
                }
            }
        } else {
            if ("a" >= "foo bar") {
                for fi12 := 0; fi12 < 8; fi12++ {
                    for fi13 := 0; fi13 < 16; fi13++ {
                        v14 := v3[fi12]
                        switch "a" {
                        case "":
                            switch "\n\t" {
                            case "foo bar":
                                v15 := v4["hello"]
                                _ = v15
                            case "hello":
                                v14 = false
                            }
                        case "hello":
                            var v16 int64 = v5
                            _ = v16
                        default:
                            v2 += (^(v2))
                        }
                        _ = fi13
                        _ = v14
                    }
                    if (1 != 1) {
                        for fi17 := 0; fi17 < 4; fi17++ {
                            v2 += (^(v2))
                            _ = len(v1)
                            _ = fi17
                        }
                        var v18 int64 = (v5 - 1)
                        v4["foo bar"] = (!(false && ((("\n\t" > "") && false) || (false || (643 > v5)))))
                        _ = v18
                    }
                    _ = fi12
                }
            } else {
                var v19 int64 = (^(v5))
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 5; fi20++ {
            fmt.Println(v2, fi20)
            for fi21 := 0; fi21 < 10; fi21++ {
                if false {
                    if (false || ((!true) && true)) {
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                        v22 := GenS2{F0: 4294967295, F1: 537, F2: "\n\t"}
                        v2 -= (v2 | v2)
                        _ = v22
                    } else {
                        switch "a" {
                        case "hello":
                            if (740 <= 339) {
                                v2 += 255
                                switch fi20 {
                                case 128:
                                    switch "\n\t" {
                                    case "a":
                                        if true {
                                            if false {
                                                var v23 uint16 = 159
                                                var v24 int16 = -32768
                                                _ = len(v4)
                                                _ = v23
                                                _ = v24
                                            } else {
                                                fmt.Println(v2, fi20, v5)
                                            }
                                            switch "hello" {
                                            case "":
                                                var v25 int = 665
                                                _ = v25
                                            case "hello":
                                                if (1e10 > (-2.25 * 1.5)) {
                                                    if (-548 == 1000) {
                                                        var v26 uint = uint(v2)
                                                        _ = v26
                                                    } else {
                                                        if (true && ((^(fi20)) == (fi20 & (fi21 & 1)))) {
                                                            var v27 uint32 = 2
                                                            _ = v27
                                                        }
                                                        var v28 bool = false
                                                        var v29 float32 = float32(float64(0.1))
                                                        _ = v28
                                                        _ = v29
                                                    }
                                                }
                                            default:
                                                if (int(v5) != (fi21 << ((0) & 63))) {
                                                    v2 += v2
                                                }
                                            }
                                            var v30 string = "a"
                                            _ = v30
                                        } else {
                                            v2 += v2
                                        }
                                    default:
                                        delete(v4, ("foo bar" + ("a" + ("a" + "foo bar"))))
                                    }
                                case 89:
                                    if true {
                                        var v31 int64 = 1
                                        var v32 int8 = 100
                                        var v33 uint32 = 2147483648
                                        if true {
                                            var v34 string = "\n\t"
                                            fmt.Println(v31)
                                            if false {
                                                if len(v1) > 0 {
                                                    v1[0] = v33
                                                }
                                            } else {
                                                v35 := GenS3{F0: (!((fi20 ^ (fi20 & fi20)) > fi21)), F1: v5, F2: v32}
                                                _ = v35
                                            }
                                            _ = v34
                                        }
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                }
                            }
                        case "a":
                            v36 := GenS4{F0: 1, F1: 32767}
                            _ = v36
                        case "":
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                        }
                        switch "a" {
                        case "a":
                            switch "hello" {
                            case "":
                                v2 += v2
                            case "\n\t":
                                v2 = v2
                            }
                        case "":
                            v37 := GenS5{F0: 0, F1: 127}
                            _ = v37
                        default:
                            if (985 <= 169) {
                                var v38 int = fi20
                                var v39 bool = false
                                _ = v38
                                _ = v39
                            } else {
                                var v40 uint32 = 65535
                                v40 = 65535
                                _ = v40
                            }
                        }
                    }
                    var v41 uint16 = 255
                    _ = v41
                } else {
                    var v42 bool = ((1.5 != (3.14159 + 3.14159)) || (((("hello" + "foo bar") + ("foo bar" + ("a" + "a"))) >= ("hello" + "")) || (false && ((100 >= 127) || (48 != 18446744073709551615)))))
                    _ = v42
                }
                if true {
                    v5 += v5
                }
                v6["a"] = fi20
                _ = fi21
            }
            v3[(fi20 ^ fi20)] = (!(871 < 4294967295))
            _ = fi20
        }
    }
    for fi43 := 0; fi43 < 9; fi43++ {
        var v44 int16 = -736
        _ = fi43
        _ = v44
    }
    fmt.Println(v5, v5)
    v5 = v5
    var v45 uint16 = 42
    var v46 uint64 = 0
    for fi47 := 0; fi47 < 1; fi47++ {
        v2 = v2
        var v48 uint8 = (v2 & v2)
        for fi49 := 0; fi49 < 3; fi49++ {
            v45 = v45
            v46 += 520
            v5 -= (v5 % ((696) | 1))
            _ = fi49
        }
        _ = fi47
        _ = v48
    }
    v45 += (v45 & 32767)
    switch 1 {
    case -1:
        if len(v1) > 0 {
            v1[0] = 1
        }
    case -108:
        switch "a" {
        case "":
            v45 += 32767
        default:
            v45 = uint16(v2)
        }
    case -2147483648:
        var v50 string = ""
        _ = v50
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi51 := 0; fi51 < 13; fi51++ {
        var v52 int = (fi51 * (fi51 / ((fi51) | 1)))
        _ = fi51
        _ = v52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v9
}
