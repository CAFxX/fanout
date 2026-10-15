package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 string; F2 string }
func genFn0(p0 string) uint64 {
    v5 := []int{32768, 501, 128}
    if len(v5) > 0 {
        v5[0] = -194
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        if len(v5) > 0 {
            v5[0] = fi6
        }
        fmt.Println(p0, fi6)
        _ = fi6
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    v1 := []string{"hello", "hello", "a"}
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    var v4 string = ("a" + "a")
    if true {
        v4 = ("hello" + v4)
        if false {
            for fi7 := 0; fi7 < 17; fi7++ {
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v3[(fi7 | (fi7 ^ fi7))] = false
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    var v8 int64 = 9223372036854775807
                    v4 = ((v4 + "a") + v4)
                    _ = v8
                }
                v4 = v4
                _ = fi7
            }
        }
    } else {
        switch (v4 + (("a" + v4) + ((v4 + v4) + "hello"))) {
        case "\n\t":
            v4 = v4
        case "foo bar":
            var v9 bool = false
            _ = v9
        case "a":
            v2[-168] = false
        }
        switch v4 {
        case "hello":
            for fi10 := 0; fi10 < 9; fi10++ {
                switch fi10 {
                case -2147483648:
                    var v11 string = (v4 + v4)
                    _ = v11
                case -461:
                    var v12 uint8 = 2
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 17; fi13++ {
                        var v14 float64 = 3.14159
                        var v15 int8 = 127
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
                _ = fi10
            }
        case "\n\t":
            var v16 uint32 = 0
            _ = v16
        default:
            var v17 int16 = 1
            _ = v17
        }
    }
    var v18 int16 = 1000
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint64 = 1
        v18 = (-(v18))
        _ = fi19
        _ = v20
    }
    switch v4 {
    case "a":
        fmt.Println(v18, v18)
    default:
        if (1 < 0) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v4 = v4
        }
    }
    if false {
        for fi21 := 0; fi21 < 8; fi21++ {
            if (221 >= 42) {
                var v22 int16 = (^(v18))
                _ = v22
            } else {
                var v23 uint8 = 255
                if (!true) {
                    v24 := GenS1{F0: (v23 + v23), F1: 42, F2: 65535}
                    if ("\n\t" != v4) {
                        for fi25 := 0; fi25 < 19; fi25++ {
                            switch "hello" {
                            case "\n\t":
                                v23 -= 128
                            case "":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            _ = fi25
                        }
                        v23 -= (v23 ^ (^(v23)))
                        fmt.Println(fi21, fi21, fi21)
                    }
                    _ = v24
                }
                _ = v23
            }
            _ = fi21
        }
    } else {
        v4 = v4
        for fi26 := 0; fi26 < 15; fi26++ {
            fmt.Println(fi26, v4, v4)
            fmt.Println(v4, fi26)
            if (!(-9223372036854775808 >= 9223372036854775807)) {
                v27 := GenS2{F0: (^(v18))}
                fmt.Println(v18, v4, fi26)
                _ = v27
            }
            _ = fi26
        }
    }
    v4 = "a"
    switch 0 {
    case 2147483647:
        if true {
            v28 := GenS1{F0: 0, F1: 444, F2: 42}
            for fi29 := 0; fi29 < 12; fi29++ {
                v4 = (v4 + v4)
                if true {
                    switch "" {
                    case "\n\t":
                        v4 = (("a" + ((v4 + v4) + v4)) + v4)
                    default:
                        var v30 int64 = -643
                        _ = v30
                    }
                    v18 = 948
                } else {
                    v4 = v4
                    for fi31 := 0; fi31 < 10; fi31++ {
                        var v32 uint = 18446744073709551615
                        v32 = v32
                        if false {
                            switch v4 {
                            case "hello":
                                var v33 int16 = -1
                                _ = v33
                            case "\n\t":
                                v4 = ((("a" + v4) + (v4 + v4)) + "foo bar")
                            default:
                                v18 = (v18 & (v18 ^ v18))
                            }
                        }
                        _ = fi31
                        _ = v32
                    }
                    var v34 float64 = 0.0
                    _ = v34
                }
                v18 = (v18 | v18)
                _ = fi29
            }
            _ = v28
        }
    case -2147483648:
        var v35 int16 = 989
        _ = v35
    case 275:
        var v36 int8 = 42
        _ = v36
    }
    switch -704 {
    case 0:
        v37 := GenS3{F0: true, F1: (("foo bar" + "\n\t") + ((("foo bar" + v4) + v4) + "foo bar")), F2: ((v4 + (v4 + "a")) + v4)}
        _ = v37
    default:
        var v38 uint8 = 0
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
