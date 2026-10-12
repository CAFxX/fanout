package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int16 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint }
type GenS4 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS5 struct { F0 uint64; F1 int; F2 uint64 }

func main() {
    var v1 uint = 2
    var v2 uint64 = 18446744073709551615
    var v3 int8 = 90
    var v4 uint64 = (v2 ^ 2)
    v5 := GenS1{F0: (("a" + "\n\t") + ""), F1: 32767, F2: -32768}
    v3 = v3
    v1 -= v1
    v6 := GenS2{F0: ("" + "\n\t"), F1: -1}
    v2 -= 4294967295
    if (560 <= 42) {
        var v7 uint32 = 0
        var v8 uint64 = (v2 + (v4 - (v4 + v4)))
        v2 += v8
        _ = v7
        _ = v8
    }
    switch "foo bar" {
    case "":
        for fi9 := 0; fi9 < 19; fi9++ {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 float32 = 1.5
                var v12 uint64 = 142
                _ = fi10
                _ = v11
                _ = v12
            }
            if (true || (fi9 < fi9)) {
                if false {
                    v4 = v4
                    v13 := GenS1{F0: ("hello" + "hello"), F1: 1000, F2: -32768}
                    _ = v13
                } else {
                    fmt.Println(v4)
                    v14 := []int{fi9, fi9}
                    _ = v14
                }
                var v15 uint64 = (v2 % ((v2) | 1))
                fmt.Println(v2)
                _ = v15
            } else {
                v16 := []int{(fi9 / ((fi9) | 1)), (fi9 | (-(fi9))), (fi9 | (fi9 | int(v4))), 906}
                for fi17 := 0; fi17 < 9; fi17++ {
                    if true {
                        v18 := GenS3{F0: 48, F1: 144, F2: v1}
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        if (!(v1 < v1)) {
                            v19 := GenS2{F0: "\n\t", F1: -189}
                            switch fi9 {
                            case 42:
                                if (float32((3.14159 - 0.1)) != (1e10 / -2.25)) {
                                    fmt.Println(v4, fi9, v2)
                                    if true {
                                        switch ("a" + "foo bar") {
                                        case "\n\t":
                                            v20 := GenS3{F0: 32767, F1: 65535, F2: 255}
                                            _ = v20
                                        }
                                        v21 := GenS1{F0: "", F1: int16(v4), F2: 639}
                                        v22 := GenS3{F0: 42, F1: 255, F2: (v1 * uint(v2))}
                                        _ = v21
                                        _ = v22
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(v1)
                                    case "hello":
                                        switch 0 {
                                        case 4294967295:
                                            v3 = (v3 % ((v3) | 1))
                                        case -2147483648:
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                        }
                                    case "a":
                                        var v23 uint64 = (v2 | (v4 / ((2) | 1)))
                                        _ = v23
                                    }
                                    v1 = (^(v1))
                                    if false {
                                        v4 += v4
                                        switch ("\n\t" + "\n\t") {
                                        case "\n\t":
                                            var v24 string = "a"
                                            _ = v24
                                        case "foo bar":
                                            if ("\n\t" < (("\n\t" + "hello") + ("" + (((("hello" + "") + "\n\t") + ("a" + ("\n\t" + "hello"))) + "hello")))) {
                                                var v25 int32 = 0
                                                _ = v25
                                            } else {
                                                var v26 int64 = 479
                                                v16 = append(v16, -488)
                                                fmt.Println(v3)
                                                _ = v26
                                            }
                                        }
                                        v27 := GenS4{F0: 128, F1: 1, F2: -1}
                                        _ = v27
                                    } else {
                                        switch ("\n\t" + ("" + "\n\t")) {
                                        case "a":
                                            v1 += (v1 >> (((v1 + (v1 & (v1 ^ 65535)))) & 63))
                                        case "hello":
                                            switch 32768 {
                                            case 580:
                                                if (uint32(v1) > 4294967295) {
                                                    var v28 float32 = ((((float32(3.14159) * ((-2.25 / float32(((1.5 / 3.14159) * float64(-2.25)))) * (0.0 * 0.1))) + 1.5) - 3.14159) / 3.14159)
                                                    _ = v28
                                                }
                                            default:
                                                if true {
                                                    var v29 int32 = 2147483647
                                                    _ = v29
                                                } else {
                                                    fmt.Println(fi17, v1)
                                                    switch int(v1) {
                                                    case -1:
                                                        var v30 int = (fi17 & int(v1))
                                                        _ = v30
                                                    case 1:
                                                        v1 = v1
                                                    case 32767:
                                                        var v31 uint16 = 32767
                                                        _ = v31
                                                    }
                                                    var v32 int64 = 1
                                                    _ = v32
                                                }
                                            }
                                            if (!(1 == 2147483647)) {
                                                var v33 uint64 = v2
                                                if (!false) {
                                                    switch "hello" {
                                                    case "":
                                                        if true {
                                                            var v34 uint = 42
                                                            var v35 float64 = 0.0
                                                            _ = v34
                                                            _ = v35
                                                        } else {
                                                            v36 := GenS4{F0: 42, F1: 0, F2: -9223372036854775808}
                                                            _ = v36
                                                        }
                                                    case "hello":
                                                        if false {
                                                            v16 = append(v16, (fi9 / ((int(v33)) | 1)))
                                                        }
                                                    case "a":
                                                        if ((^(v3)) != -62) {
                                                            var v37 uint32 = 2147483648
                                                            _ = v37
                                                        } else {
                                                            v1 += 1
                                                            v4 = v2
                                                            var v38 int16 = 272
                                                            _ = v38
                                                        }
                                                    }
                                                    var v39 int64 = 0
                                                    _ = v39
                                                } else {
                                                    v4 -= (v4 << (((v1 + v1)) & 63))
                                                }
                                                if false {
                                                    fmt.Println(fi17, v4, fi17)
                                                    var v40 uint32 = 65535
                                                    _ = v40
                                                } else {
                                                    fmt.Println(v3, v1)
                                                    v4 = v2
                                                }
                                                _ = v33
                                            }
                                        default:
                                            v41 := GenS1{F0: ("a" + "hello"), F1: 1000, F2: -32768}
                                            _ = v41
                                        }
                                        v42 := GenS5{F0: v2, F1: fi9, F2: v4}
                                        _ = v42
                                    }
                                }
                                switch "hello" {
                                case "a":
                                    fmt.Println(v1, fi17, v3)
                                }
                                fmt.Println(v1, fi17)
                            default:
                                fmt.Println(v3)
                            }
                            _ = v19
                        } else {
                            fmt.Println(v1, fi17)
                            var v43 uint = 888
                            var v44 float64 = float64(1.5)
                            _ = v43
                            _ = v44
                        }
                        _ = v18
                    }
                    _ = fi17
                }
                _ = v16
            }
            var v45 int16 = -1
            _ = fi9
            _ = v45
        }
    case "hello":
        if false {
            v46 := []int{-848}
            _ = v46
        } else {
            if (uint8(v2) == 127) {
                v4 += 26
                v3 += v3
            }
        }
    case "foo bar":
        fmt.Println(v4, v4, v2)
    }
    v2 -= 4294967295
    v47 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
