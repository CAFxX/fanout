package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 string; F2 int8 }
type GenS3 struct { F0 int; F1 uint8; F2 int32 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    v3 := []int{32768, 0}
    var v4 int16 = 676
    _ = len(v3)
    fmt.Println(v4)
    var v5 int8 = 17
    if (577 >= 0) {
        if false {
            if true {
                if true {
                    v6 := GenS1{F0: (("a" + "hello") + "foo bar")}
                    v5 += v5
                    _ = v6
                } else {
                    for fi7 := 0; fi7 < 19; fi7++ {
                        var v8 uint32 = 441
                        _ = fi7
                        _ = v8
                    }
                }
                if false {
                    var v9 bool = true
                    v5 -= -108
                    _ = v9
                } else {
                    var v10 string = "\n\t"
                    fmt.Println(v10)
                    _ = v10
                }
                if (18446744073709551615 != 2) {
                    var v11 bool = (true && true)
                    if v11 {
                        v5 = (-(v5))
                        v5 = (v5 * v5)
                        var v12 int16 = 541
                        _ = v12
                    }
                    _ = v11
                } else {
                    fmt.Println(v4, v5)
                    var v13 int64 = -1
                    _ = v13
                }
            }
            _ = len(v2)
        }
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v4, v4)
        case "hello":
            for fi14 := 0; fi14 < 16; fi14++ {
                v15 := GenS1{F0: ("foo bar" + ("a" + "hello"))}
                _ = fi14
                _ = v15
            }
        case "foo bar":
            v5 = v5
        }
    }
    v5 -= (v5 << ((255) & 63))
    _ = len(v3)
    for fi16 := 0; fi16 < 20; fi16++ {
        if false {
            if (-1 > 607) {
                v4 = -642
                delete(v1, "a")
                v17 := GenS1{F0: ""}
                _ = v17
            }
            var v18 uint64 = 0
            _ = v18
        } else {
            if true {
                for fi19 := 0; fi19 < 16; fi19++ {
                    v20 := GenS2{F0: 753, F1: "hello", F2: -48}
                    fmt.Println(v4)
                    if false {
                        switch (fi16 / (((^(fi16))) | 1)) {
                        case -911:
                            var v21 string = (("\n\t" + "foo bar") + (("hello" + ((("\n\t" + "foo bar") + "hello") + "\n\t")) + ""))
                            _ = v21
                        case 127:
                            var v22 uint32 = 603
                            _ = v22
                        case 1:
                            v23 := GenS3{F0: (^(fi19)), F1: 99, F2: -1}
                            _ = v23
                        default:
                            v3 = append(v3, (fi16 | -9223372036854775808))
                        }
                        var v24 int = (fi19 & (^(fi19)))
                        _ = v24
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = len(v2)
                    if (42 <= 232) {
                        var v25 int = (fi19 | (fi16 + fi16))
                        fmt.Println(v4, fi16)
                        _ = len(v3)
                        _ = v25
                    }
                    _ = fi19
                    _ = v20
                }
                var v26 float64 = (0.1 * (((0.0 + 1e10) * 3.14159) * 1e10))
                _ = v26
            } else {
                switch -706 {
                case 32767:
                    var v27 int = (fi16 % ((fi16) | 1))
                    _ = v27
                case -2147483648:
                    var v28 float32 = 1.5
                    _ = v28
                case 9223372036854775807:
                    if true {
                        if true {
                            v5 = (-(v5))
                        }
                    }
                }
            }
        }
        v29 := GenS4{F0: 861}
        _ = fi16
        _ = v29
    }
    v4 -= v4
    var v30 uint = 824
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
