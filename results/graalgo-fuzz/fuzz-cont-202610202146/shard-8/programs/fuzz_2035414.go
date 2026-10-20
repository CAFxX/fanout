package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint16 = 577
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    var v5 uint16 = v1
    for fi6 := 0; fi6 < 2; fi6++ {
        if ((81 <= -1) || false) {
            fmt.Println(fi6)
        }
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := v4[343]
            v5 += 65535
            v1 -= 42
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v9 := GenS1{F0: -804, F1: 1, F2: (!true)}
    v5 -= 457
    switch ("foo bar" + "") {
    case "foo bar":
        var v10 string = (("hello" + ("foo bar" + "a")) + "\n\t")
        _ = v10
    case "\n\t":
        for fi11 := 0; fi11 < 16; fi11++ {
            v5 += 32767
            var v12 int32 = -1
            _ = fi11
            _ = v12
        }
    case "":
        v1 += v1
    }
    if ((42 <= 2) || (-1 == 1)) {
        _ = len(v3)
        v1 = 171
        var v13 float32 = 1.5
        _ = v13
    } else {
        if (false || true) {
            v3[0] = 473
        } else {
            v14 := []int{-2147483648, 9223372036854775807, 42, 127}
            switch 32768 {
            case 127:
                _ = len(v14)
            case 128:
                var v15 int32 = 42
                _ = v15
            default:
                switch "" {
                case "hello":
                    for fi16 := 0; fi16 < 14; fi16++ {
                        fmt.Println(v5, fi16, fi16)
                        v17 := GenS2{F0: -93}
                        if false {
                            var v18 float32 = (3.14159 + -2.25)
                            v19 := GenS1{F0: -1, F1: (fi16 & 577), F2: ("" <= ("" + ("\n\t" + "")))}
                            v1 = v1
                            _ = v18
                            _ = v19
                        }
                        _ = fi16
                        _ = v17
                    }
                case "\n\t":
                    v1 -= 1
                case "":
                    v1 = (v5 - 65535)
                default:
                    v14 = append(v14, 42)
                }
            }
            _ = v14
        }
    }
    fmt.Println(v1, v5)
    v20 := GenS2{F0: 1}
    var v21 uint32 = 2147483647
    v3[v21] = (v21 + v21)
    switch 127 {
    case -490:
        if true {
            switch "a" {
            case "\n\t":
                var v22 int8 = 100
                _ = v22
            default:
                v21 = (v21 | v21)
            }
            fmt.Println(v1, v1)
            if false {
                v21 -= 798
                if true {
                    var v23 bool = false
                    var v24 uint8 = 42
                    for fi25 := 0; fi25 < 15; fi25++ {
                        for fi26 := 0; fi26 < 7; fi26++ {
                            v27 := v3[53]
                            var v28 int8 = 7
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        v29 := GenS2{F0: 42}
                        if (((1e10 - 1.5) == 1.5) || (fi25 == -2147483648)) {
                            delete(v4, (v21 % ((1) | 1)))
                            v23 = v23
                        } else {
                            var v30 uint = 65535
                            v31 := GenS2{F0: -128}
                            var v32 int = (fi25 * -405)
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        _ = fi25
                        _ = v29
                    }
                    _ = v23
                    _ = v24
                } else {
                    v4[(v21 % ((v21) | 1))] = false
                }
            }
        } else {
            if true {
                var v33 int8 = 0
                if true {
                    v33 -= (v33 | (v33 | int8(v21)))
                    v34 := GenS3{F0: 1}
                    var v35 bool = ((v33 / (((v33 << ((392) & 63))) | 1)) > 1)
                    _ = v34
                    _ = v35
                }
                _ = v33
            }
            var v36 float64 = (0.0 + 1e10)
            for fi37 := 0; fi37 < 3; fi37++ {
                for fi38 := 0; fi38 < 1; fi38++ {
                    var v39 int16 = 920
                    var v40 string = "a"
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                if false {
                    v41 := []int{(fi37 - 32768)}
                    _ = v41
                }
                _ = fi37
            }
            _ = v36
        }
    case 490:
        v1 += v1
    }
    if (!(127 == -1)) {
        v2[(v21 ^ (v21 * v21))] = (("foo bar" + ("" + "foo bar")) + ("\n\t" + ("a" + "foo bar")))
        v42 := GenS3{F0: 127}
        _ = v42
    }
    if (!false) {
        v5 = v5
        v43 := []int{}
        _ = v43
    } else {
        var v44 uint64 = 42
        var v45 int16 = 32767
        _ = v44
        _ = v45
    }
    switch -2147483648 {
    case -108:
        var v46 int16 = 288
        _ = v46
    }
    fmt.Println(v1, v5)
    v21 = (^(v21))
    switch ("hello" + "") {
    case "a":
        v47 := []int{}
        _ = v47
    case "\n\t":
        var v48 int8 = -1
        _ = v48
    case "":
        v21 = v21
    default:
        v5 -= v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
