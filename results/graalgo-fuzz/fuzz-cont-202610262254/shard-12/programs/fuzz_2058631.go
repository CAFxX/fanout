package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int16; F1 uint32; F2 int16 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int8 }

func main() {
    var v1 string = ("a" + (("" + "a") + ("" + "a")))
    var v2 int16 = -313
    var v3 uint8 = 221
    var v4 int = 0
    var v5 int = v4
    if (!false) {
        v5 = v5
        switch v1 {
        case "foo bar":
            v5 = v4
        case "a":
            switch v4 {
            case -318:
                fmt.Println(v4)
            default:
                v6 := []int{(v4 | (v4 + (v4 & (v5 << ((2) & 63))))), (v4 | (v5 + (v5 / ((v5) | 1)))), (v4 ^ v4)}
                _ = v6
            }
        case "hello":
            if false {
                for fi7 := 0; fi7 < 16; fi7++ {
                    if false {
                        v2 = (v2 + v2)
                    }
                    switch (((v1 + v1) + v1) + v1) {
                    case "foo bar":
                        switch 158 {
                        case -9223372036854775808:
                            v3 -= (v3 ^ v3)
                        case 32768:
                            v8 := GenS1{F0: int(v3), F1: v3, F2: 42}
                            _ = v8
                        case 2147483647:
                            v2 += (v2 % ((v2) | 1))
                        default:
                            v9 := []int{v4, 2147483647}
                            _ = v9
                        }
                    case "a":
                        v10 := []int{}
                        _ = v10
                    default:
                        for fi11 := 0; fi11 < 4; fi11++ {
                            v1 = "\n\t"
                            fmt.Println(fi7, fi11, v5)
                            if true {
                                if false {
                                    var v12 int = (v5 / ((4294967295) | 1))
                                    var v13 uint16 = 0
                                    v14 := []int{fi7}
                                    _ = v12
                                    _ = v13
                                    _ = v14
                                }
                                v15 := GenS2{F0: 255}
                                if (!false) {
                                    v16 := []int{-1}
                                    fmt.Println(v3)
                                    switch "" {
                                    case "\n\t":
                                        v17 := GenS1{F0: (fi7 | 2147483647), F1: v3, F2: 650}
                                        _ = v17
                                    case "hello":
                                        switch v1 {
                                        case "":
                                            _ = len(v16)
                                        }
                                    case "":
                                        if true {
                                            var v18 int8 = 42
                                            _ = v18
                                        }
                                    default:
                                        v19 := GenS3{F0: 625}
                                        _ = v19
                                    }
                                    _ = v16
                                } else {
                                    var v20 uint32 = 42
                                    fmt.Println(fi7)
                                    v21 := GenS1{F0: 32767, F1: (v3 % ((1) | 1)), F2: 1}
                                    _ = v20
                                    _ = v21
                                }
                                _ = v15
                            } else {
                                if true {
                                    var v22 int = fi11
                                    _ = v22
                                }
                                v23 := GenS4{F0: fi11, F1: (v3 / ((v3) | 1))}
                                var v24 uint32 = 42
                                _ = v23
                                _ = v24
                            }
                            _ = fi11
                        }
                    }
                    fmt.Println(v1)
                    _ = fi7
                }
            } else {
                v25 := []int{42, v4, 2147483647}
                _ = v25
            }
        }
    } else {
        for fi26 := 0; fi26 < 3; fi26++ {
            v5 += 32767
            fmt.Println(fi26, fi26)
            switch ("foo bar" + v1) {
            case "hello":
                v1 = v1
            case "a":
                v3 -= v3
            default:
                var v27 string = v1
                _ = v27
            }
            _ = fi26
        }
        var v28 string = ("a" + v1)
        _ = v28
    }
    if true {
        for fi29 := 0; fi29 < 18; fi29++ {
            v30 := GenS1{F0: 4294967295, F1: (v3 - v3), F2: 0}
            v4 -= v4
            if (!(false || (!(306 != 1)))) {
                v31 := GenS5{F0: v2, F1: 407, F2: v2}
                _ = v31
            }
            _ = fi29
            _ = v30
        }
        var v32 uint16 = 255
        switch "" {
        case "":
            fmt.Println(v2)
        case "hello":
            var v33 uint64 = uint64(v32)
            _ = v33
        default:
            var v34 uint32 = 4294967295
            _ = v34
        }
        _ = v32
    } else {
        var v35 int16 = 32767
        var v36 uint = uint(v3)
        for fi37 := 0; fi37 < 7; fi37++ {
            fmt.Println(v35, v36, v1)
            v35 -= (v2 >> ((v36) & 63))
            _ = fi37
        }
        _ = v35
        _ = v36
    }
    v3 += v3
    var v38 string = ((v1 + v1) + v1)
    v39 := []int{v5}
    var v40 string = v1
    var v41 int32 = -530
    if (v3 <= (v3 - v3)) {
        v42 := GenS3{F0: (v41 << ((658) & 63))}
        _ = v42
    } else {
        v43 := GenS2{F0: 808}
        _ = v43
    }
    v44 := GenS6{F0: 177, F1: ""}
    v45 := GenS7{F0: 8}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v45
    _ = v5
}
