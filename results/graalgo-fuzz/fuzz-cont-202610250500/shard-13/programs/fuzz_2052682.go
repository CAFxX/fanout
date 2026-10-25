package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool; F1 uint8; F2 int16 }
type GenS10 struct { F0 uint; F1 int16 }

func main() {
    var v1 uint64 = 2
    var v2 int16 = -892
    var v3 uint = 255
    for fi4 := 0; fi4 < 12; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    v1 -= v1
    fmt.Println(v2, v2)
    var v5 string = ""
    if (true || (!(!(1 < -751)))) {
        v1 = (v1 & 2)
        v6 := GenS1{F0: 2147483647, F1: 1, F2: false}
        _ = v6
    }
    fmt.Println(v3, v2)
    var v7 bool = (0 <= 42)
    var v8 float32 = 1e10
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := []int{(fi9 | -573)}
        v10 = append(v10, (-(fi9)))
        _ = fi9
        _ = v10
    }
    var v11 float64 = (-2.25 * 1.5)
    var v12 int8 = -128
    if ((255 != 16) && v7) {
        v8 += (v8 - v8)
        switch 128 {
        case 4294967295:
            if false {
                v1 = v1
                v13 := GenS2{F0: 42}
                var v14 int = -1
                _ = v13
                _ = v14
            } else {
                if true {
                    if (uint8(v1) <= 12) {
                        v12 = (v12 + 0)
                        if (1 != 255) {
                            switch v5 {
                            case "hello":
                                fmt.Println(v12, v1, v11)
                            case "foo bar":
                                v15 := GenS1{F0: 278, F1: 38, F2: v7}
                                _ = v15
                            }
                            v16 := []int{32768, -2147483648, int(v2)}
                            _ = len(v16)
                            _ = v16
                        } else {
                            v17 := GenS3{F0: 128, F1: -213}
                            var v18 float64 = (v11 * 1e10)
                            _ = v17
                            _ = v18
                        }
                        if (306 != 0) {
                            fmt.Println(v2, v5)
                        } else {
                            v7 = v7
                            if v7 {
                                v7 = v7
                                v3 = 42
                                fmt.Println(v12)
                            }
                        }
                    }
                    switch 1 {
                    case -364:
                        v19 := GenS3{F0: int(v1), F1: -2147483648}
                        _ = v19
                    case 128:
                        if v7 {
                            fmt.Println(v2)
                        }
                    }
                } else {
                    v20 := GenS4{F0: false, F1: (-(v2))}
                    switch 598 {
                    case 2147483647:
                        v21 := GenS5{F0: v2}
                        _ = v21
                    }
                    v22 := []int{4294967295, -2147483648, 4294967295}
                    _ = v20
                    _ = v22
                }
                fmt.Println(v11)
            }
        }
    }
    if true {
        var v23 uint8 = 255
        switch -269 {
        case 376:
            v24 := []int{4294967295}
            _ = v24
        case 663:
            for fi25 := 0; fi25 < 7; fi25++ {
                for fi26 := 0; fi26 < 17; fi26++ {
                    switch v5 {
                    case "":
                        var v27 int64 = 42
                        _ = v27
                    case "foo bar":
                        v28 := []int{}
                        _ = v28
                    case "hello":
                        v29 := GenS6{F0: 0, F1: int16(fi25)}
                        _ = v29
                    }
                    var v30 float32 = v8
                    _ = fi26
                    _ = v30
                }
                switch fi25 {
                case 785:
                    var v31 uint16 = 11
                    _ = v31
                case 0:
                    v32 := GenS4{F0: v7, F1: -32768}
                    _ = v32
                case 1:
                    for fi33 := 0; fi33 < 16; fi33++ {
                        if (fi33 == (fi33 % ((fi33) | 1))) {
                            v5 = v5
                            if (42 > 0) {
                                v7 = false
                            }
                            fmt.Println(v8, v12, v23)
                        } else {
                            if (v5 > "hello") {
                                switch "hello" {
                                case "hello":
                                    if (1.5 >= -2.25) {
                                        var v34 uint64 = 628
                                        _ = v34
                                    }
                                case "\n\t":
                                    v3 -= (v3 << (((^(v3))) & 63))
                                case "a":
                                    v35 := GenS5{F0: v2}
                                    _ = v35
                                default:
                                    v36 := []int{42, fi25, fi25, fi25}
                                    _ = v36
                                }
                            }
                        }
                        v37 := GenS7{F0: "foo bar", F1: 42}
                        v38 := GenS8{F0: (-(v12))}
                        var v39 float32 = v8
                        _ = fi33
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                default:
                    for fi40 := 0; fi40 < 20; fi40++ {
                        v7 = (!(v11 < (float64((3.14159 / v8)) + v11)))
                        v41 := []int{42, (^(fi40)), 42}
                        _ = fi40
                        _ = v41
                    }
                }
                if (v8 > 0.1) {
                    v3 -= 42
                    v42 := GenS8{F0: v12}
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 8; fi43++ {
                        v2 -= -1
                        v44 := GenS8{F0: (-(v12))}
                        if ((((-59 >= v2) && (v23 > v23)) || (!v7)) || v7) {
                            var v45 uint16 = 32767
                            v46 := []int{(^(fi25)), -9223372036854775808}
                            fmt.Println(v7, fi25)
                            fmt.Println(v2, v5)
                            _ = v45
                            _ = v46
                        } else {
                            v47 := []int{fi43, 32767}
                            var v48 uint16 = 95
                            _ = v47
                            _ = v48
                        }
                        _ = fi43
                        _ = v44
                    }
                    switch -190 {
                    case 754:
                        v49 := GenS9{F0: ((^(v2)) < 1), F1: (v23 & (v23 >> ((v3) & 63))), F2: -32768}
                        _ = v49
                    case -825:
                        for fi50 := 0; fi50 < 6; fi50++ {
                            fmt.Println(fi50, fi25)
                            _ = fi50
                        }
                    }
                    if (!v7) {
                        var v51 int16 = int16(v3)
                        _ = v51
                    } else {
                        v8 = v8
                        fmt.Println(v1, v12, v11)
                        v1 += 4294967295
                    }
                }
                _ = fi25
            }
        case 2147483647:
            var v52 uint = (v3 ^ v3)
            _ = v52
        }
        switch int(v12) {
        case -2147483648:
            if v7 {
                switch 2147483647 {
                case 9223372036854775807:
                    v7 = (1.5 <= float64(v8))
                case -574:
                    v53 := GenS10{F0: 27, F1: (^(v2))}
                    _ = v53
                }
                v7 = true
            }
        case 42:
            for fi54 := 0; fi54 < 20; fi54++ {
                v2 += v2
                v7 = (v11 >= (float64(v8) + v11))
                if v7 {
                    var v55 int32 = 2147483647
                    var v56 int8 = (v12 % ((-21) | 1))
                    fmt.Println(v12)
                    _ = v55
                    _ = v56
                }
                _ = fi54
            }
        default:
            v23 += 2
        }
        _ = v23
    } else {
        switch ("\n\t" + v5) {
        case "a":
            var v57 int64 = -1
            _ = v57
        default:
            v1 -= 4294967295
        }
        v3 = v3
    }
    v58 := []int{}
    v58 = append(v58, 0)
    if len(v58) >= 2 {
        v58 = v58[1:len(v58)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v58
    _ = v7
    _ = v8
}
