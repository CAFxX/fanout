package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int16; F1 int64; F2 uint16 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 string; F1 int8; F2 int8 }
func genFn0(p0 string) uint32 {
    v4 := []int{-2147483648, -2147483648, 0}
    switch "" {
    case "\n\t":
        v5 := GenS1{F0: -126, F1: 65535, F2: 504}
        _ = v5
    case "":
        switch "\n\t" {
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = 698
            }
        default:
            fmt.Println(p0)
        }
    case "hello":
        switch "foo bar" {
        case "\n\t":
            if true {
                if (((true && true) || (!(1 == -389))) && false) {
                    p0 = "hello"
                } else {
                    if false {
                        _ = len(v4)
                        fmt.Println(p0)
                        if (true || (127 <= 255)) {
                            if false {
                                p0 = ("" + p0)
                                p0 = ((((p0 + ("foo bar" + p0)) + p0) + (p0 + "foo bar")) + p0)
                                if true {
                                    for fi6 := 0; fi6 < 8; fi6++ {
                                        var v7 int32 = -2147483648
                                        _ = len(v4)
                                        v8 := GenS1{F0: 42, F1: 32767, F2: 2}
                                        _ = fi6
                                        _ = v7
                                        _ = v8
                                    }
                                } else {
                                    v4 = append(v4, -1)
                                    p0 = p0
                                }
                            } else {
                                for fi9 := 0; fi9 < 20; fi9++ {
                                    fmt.Println(p0, fi9)
                                    if ((42 == 4294967295) && true) {
                                        v10 := GenS2{F0: -372, F1: 9223372036854775807, F2: 2}
                                        fmt.Println(fi9)
                                        p0 = (p0 + p0)
                                        _ = v10
                                    }
                                    _ = fi9
                                }
                            }
                            p0 = p0
                        }
                    }
                    var v11 string = "foo bar"
                    _ = v11
                }
                switch ((p0 + p0) + (p0 + p0)) {
                case "a":
                    switch -1 {
                    case 127:
                        if false {
                            if len(v4) > 0 {
                                v4[0] = 127
                            }
                            for fi12 := 0; fi12 < 5; fi12++ {
                                if (!(2147483648 > 42)) {
                                    switch fi12 {
                                    case 128:
                                        for fi13 := 0; fi13 < 1; fi13++ {
                                            p0 = p0
                                            v14 := GenS1{F0: 1, F1: 42, F2: 2147483647}
                                            _ = fi13
                                            _ = v14
                                        }
                                    case -158:
                                        for fi15 := 0; fi15 < 12; fi15++ {
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v16 float64 = 0.0
                                                _ = v16
                                            }
                                            v17 := GenS2{F0: 42, F1: 42, F2: 65535}
                                            if false {
                                                var v18 int16 = 368
                                                fmt.Println(fi15, p0, p0)
                                                _ = v18
                                            }
                                            _ = fi15
                                            _ = v17
                                        }
                                    case 1:
                                        p0 = p0
                                    }
                                }
                                _ = fi12
                            }
                            p0 = "hello"
                        } else {
                            for fi19 := 0; fi19 < 19; fi19++ {
                                var v20 uint = 65535
                                _ = fi19
                                _ = v20
                            }
                        }
                    case 2147483647:
                        v21 := GenS1{F0: -1, F1: 740, F2: 123}
                        _ = v21
                    case -746:
                        var v22 int8 = 42
                        _ = v22
                    default:
                        fmt.Println(p0)
                    }
                case "\n\t":
                    var v23 float32 = 3.14159
                    _ = v23
                default:
                    fmt.Println(p0)
                }
                v24 := GenS2{F0: 42, F1: -1, F2: 0}
                _ = v24
            }
        case "":
            v25 := GenS3{F0: 100, F1: -58}
            _ = v25
        default:
            for fi26 := 0; fi26 < 9; fi26++ {
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 uint8 = 236
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
        }
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    _ = p0
    _ = v4
    return 995
}

func main() {
    v1 := []uint32{781, 4294967295, 2147483648}
    v2 := make(map[string]bool)
    v3 := []int{-106}
    var v29 float64 = float64(0.0)
    var v30 string = "foo bar"
    v31 := GenS1{F0: 0, F1: 0, F2: 42}
    for fi32 := 0; fi32 < 4; fi32++ {
        switch (fi32 << ((42) & 63)) {
        case 172:
            switch v30 {
            case "\n\t":
                var v33 uint = 2
                _ = v33
            case "hello":
                var v34 uint = 2
                _ = v34
            }
        }
        var v35 int = (fi32 * 0)
        _ = fi32
        _ = v35
    }
    v30 = "hello"
    v36 := GenS2{F0: 0, F1: -555, F2: 255}
    v30 = v30
    v37 := GenS4{F0: "hello", F1: 42, F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
}
