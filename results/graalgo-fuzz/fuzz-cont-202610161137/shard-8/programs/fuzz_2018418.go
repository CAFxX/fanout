package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 float32 = (-2.25 * -2.25)
    var v2 uint8 = 2
    var v3 int32 = 1
    var v4 float64 = -2.25
    var v5 string = (("hello" + "\n\t") + "foo bar")
    var v6 bool = (412 == 65535)
    v7 := GenS1{F0: -128}
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := []int{}
        if false {
            for fi10 := 0; fi10 < 11; fi10++ {
                v3 = (v3 * (v3 + v3))
                if len(v9) > 0 {
                    v9[0] = -107
                }
                var v11 uint32 = 0
                _ = fi10
                _ = v11
            }
            var v12 int8 = 1
            switch 32768 {
            case -2147483648:
                for fi13 := 0; fi13 < 6; fi13++ {
                    fmt.Println(v12, v1, v4)
                    v14 := GenS1{F0: v12}
                    fmt.Println(v12)
                    _ = fi13
                    _ = v14
                }
            case 42:
                if (0 <= -976) {
                    v1 = 1e10
                    for fi15 := 0; fi15 < 13; fi15++ {
                        switch "a" {
                        case "foo bar":
                            switch v5 {
                            case "foo bar":
                                var v16 uint32 = 2
                                switch "hello" {
                                case "foo bar":
                                    switch v5 {
                                    case "foo bar":
                                        if v6 {
                                            if (838 <= 927) {
                                                v17 := GenS2{F0: (^(v2)), F1: fi8, F2: v5}
                                                _ = v17
                                            }
                                            v18 := GenS1{F0: v12}
                                            _ = v18
                                        }
                                    case "hello":
                                        v19 := GenS1{F0: (v12 << ((4294967295) & 63))}
                                        _ = v19
                                    default:
                                        _ = len(v9)
                                    }
                                    if (true || (!v6)) {
                                        var v20 uint64 = uint64(v2)
                                        if v6 {
                                            var v21 uint32 = (v16 & (v16 * (v16 % (((v16 % ((uint32(v20)) | 1))) | 1))))
                                            _ = v21
                                        }
                                        _ = v20
                                    } else {
                                        v9 = append(v9, (^(fi15)))
                                        var v22 int64 = 9223372036854775807
                                        _ = v22
                                    }
                                case "\n\t":
                                    var v23 string = v5
                                    _ = v23
                                default:
                                    v24 := GenS2{F0: 2, F1: -930, F2: v5}
                                    _ = v24
                                }
                                var v25 float32 = float32(v4)
                                _ = v16
                                _ = v25
                            case "a":
                                var v26 float64 = v4
                                _ = v26
                            case "\n\t":
                                if (891 <= 0) {
                                    var v27 int32 = v3
                                    if (true && true) {
                                        v5 = "a"
                                    }
                                    _ = v27
                                }
                            }
                        default:
                            switch fi8 {
                            case -602:
                                v5 = v5
                            case 2147483647:
                                v9 = append(v9, fi8)
                            default:
                                v12 += v12
                            }
                        }
                        switch -689 {
                        case -423:
                            v28 := GenS3{F0: -1}
                            _ = v28
                        default:
                            v12 = (v12 << ((188) & 63))
                        }
                        var v29 uint8 = (v2 + 0)
                        _ = fi15
                        _ = v29
                    }
                    v30 := GenS4{F0: -1, F1: 432}
                    _ = v30
                }
            case 32768:
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            default:
                _ = len(v9)
            }
            _ = v12
        }
        v2 += v2
        _ = fi8
        _ = v9
    }
    switch -674 {
    case 32768:
        var v31 uint64 = 2
        _ = v31
    default:
        fmt.Println(v5)
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        fmt.Println(v5, v6, v2)
        v33 := []int{(fi32 << ((18446744073709551615) & 63)), (fi32 + fi32)}
        switch v5 {
        case "a":
            v6 = v6
        case "hello":
            if v6 {
                var v34 int = fi32
                for fi35 := 0; fi35 < 13; fi35++ {
                    v36 := GenS3{F0: -72}
                    v1 -= v1
                    switch (("hello" + "\n\t") + (v5 + (v5 + ((v5 + "foo bar") + (v5 + ((v5 + v5) + (v5 + v5))))))) {
                    case "foo bar":
                        fmt.Println(fi35, fi32, v34)
                    default:
                        v37 := GenS5{F0: 195}
                        _ = v37
                    }
                    v33 = append(v33, -2147483648)
                    v38 := GenS4{F0: -1, F1: 613}
                    _ = fi35
                    _ = v36
                    _ = v38
                }
                _ = v34
            } else {
                switch v5 {
                case "foo bar":
                    v33 = append(v33, (fi32 & fi32))
                default:
                    v2 += (v2 / ((127) | 1))
                }
                var v39 uint = 745
                _ = v39
            }
        case "":
            v6 = (!((v2 / ((v2) | 1)) > 0))
        }
        _ = fi32
        _ = v33
    }
    var v40 string = v5
    var v41 float32 = v1
    v42 := GenS2{F0: 1, F1: 4294967295, F2: v5}
    v4 = float64(3.14159)
    fmt.Println(v40, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
