package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint64; F1 string; F2 int8 }
type GenS4 struct { F0 int32; F1 uint; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint; F2 uint32 }

func main() {
    var v1 string = ""
    v2 := []int{274, 127, 1, -158}
    var v3 int32 = 2147483647
    var v4 uint16 = 562
    var v5 string = ("" + ("\n\t" + ("foo bar" + ("foo bar" + "a"))))
    var v6 uint8 = 189
    switch (v1 + v5) {
    case "hello":
        v4 = (v4 & v4)
    case "\n\t":
        switch 2147483647 {
        case -1:
            var v7 uint32 = 185
            _ = v7
        case 370:
            v3 = v3
        default:
            for fi8 := 0; fi8 < 12; fi8++ {
                v6 += v6
                _ = fi8
            }
        }
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: 0}
        for fi11 := 0; fi11 < 6; fi11++ {
            if (true || (-906 < -323)) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v3 = (v3 << ((42) & 63))
                fmt.Println(fi9, v1)
            } else {
                var v12 float32 = 0.0
                switch ("a" + "a") {
                case "":
                    fmt.Println(v6, v3, v6)
                }
                if ((v4 * (v4 >> ((4294967295) & 63))) != v4) {
                    v2 = append(v2, fi9)
                    if true {
                        v6 += (v6 - uint8(v4))
                        v5 = v1
                    } else {
                        var v13 bool = (!true)
                        _ = v13
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = fi11
                    }
                    v4 = 255
                    var v14 int32 = (v3 - v3)
                    _ = v14
                }
                if (!false) {
                    v15 := GenS2{F0: fi11, F1: 1}
                    _ = v15
                }
                _ = v12
            }
            var v16 string = v1
            _ = fi11
            _ = v16
        }
        var v17 uint8 = (v6 + (v6 & v6))
        _ = fi9
        _ = v10
        _ = v17
    }
    var v18 int16 = -32768
    var v19 int32 = (v3 << ((0) & 63))
    if true {
        v20 := GenS1{F0: 42}
        fmt.Println(v4, v3)
        var v21 int32 = 31
        _ = v20
        _ = v21
    }
    v4 = v4
    v22 := GenS2{F0: 127, F1: 4294967295}
    fmt.Println(v1, v6)
    v23 := GenS3{F0: 2, F1: v5, F2: -128}
    if false {
        v4 += 255
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        fmt.Println(v6, v4)
        v4 += (^(v4))
        for fi25 := 0; fi25 < 17; fi25++ {
            if len(v2) > 0 {
                v2[0] = fi25
            }
            v26 := GenS4{F0: -2147483648, F1: 0, F2: 18446744073709551615}
            v27 := GenS4{F0: (-(v19)), F1: 437, F2: 4294967295}
            v28 := GenS5{F0: false, F1: 255, F2: 881}
            if true {
                var v29 uint8 = (^(v6))
                _ = v29
            } else {
                _ = len(v2)
            }
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        _ = fi24
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        v2 = append(v2, (fi30 & fi30))
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 int16 = int16(v6)
            if false {
                v33 := GenS4{F0: (v3 << ((916) & 63)), F1: 656, F2: 717}
                var v34 uint32 = 1
                if ((-1 >= -128) && false) {
                    fmt.Println(v1)
                    switch v1 {
                    case "\n\t":
                        switch fi31 {
                        case 127:
                            v5 = "a"
                        case 42:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            if false {
                                _ = len(v2)
                                fmt.Println(v19, fi31, fi30)
                            }
                        }
                    case "a":
                        fmt.Println(v3, v34)
                    case "hello":
                        var v35 int32 = -398
                        _ = v35
                    default:
                        switch "" {
                        case "":
                            var v36 uint64 = 0
                            switch (v1 + (v5 + v1)) {
                            case "foo bar":
                                var v37 int8 = 0
                                _ = v37
                            default:
                                var v38 int16 = (v32 * int16(v34))
                                _ = v38
                            }
                            v5 = ((v5 + ((("\n\t" + (v1 + "foo bar")) + v5) + v5)) + v5)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v36
                        default:
                            switch fi30 {
                            case 32768:
                                if true {
                                    v3 += (-(v3))
                                    var v39 uint16 = v4
                                    v19 -= (v19 / ((int32(v18)) | 1))
                                    _ = v39
                                } else {
                                    var v40 uint64 = 1
                                    _ = v40
                                }
                                if false {
                                    fmt.Println(v5, v3, v32)
                                    v2 = append(v2, (fi31 * fi30))
                                }
                                if (((-(v3)) == v3) || true) {
                                    v3 -= v19
                                } else {
                                    v3 = (v19 >> ((2) & 63))
                                    v41 := GenS1{F0: (^(v34))}
                                    _ = v41
                                }
                            default:
                                if (!((!false) || true)) {
                                    var v42 int = 4294967295
                                    _ = v42
                                }
                            }
                        }
                    }
                    var v43 string = ("foo bar" + "foo bar")
                    fmt.Println(v5)
                    var v44 uint16 = (v4 << ((998) & 63))
                    var v45 string = ("" + v1)
                    _ = v43
                    _ = v44
                    _ = v45
                }
                _ = v33
                _ = v34
            }
            var v46 int8 = -128
            _ = fi31
            _ = v32
            _ = v46
        }
        var v47 uint16 = v4
        _ = fi30
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
