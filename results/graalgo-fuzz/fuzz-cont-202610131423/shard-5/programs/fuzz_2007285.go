package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 uint64; F1 uint8; F2 int }
type GenS7 struct { F0 uint; F1 int8; F2 uint32 }
type GenS8 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS9 struct { F0 uint16; F1 uint }

func main() {
    var v1 int16 = 1000
    var v2 uint16 = 0
    var v3 string = ""
    v3 = (v3 + (v3 + "\n\t"))
    var v4 uint64 = 18446744073709551615
    if true {
        var v5 uint8 = 0
        _ = v5
    } else {
        for fi6 := 0; fi6 < 18; fi6++ {
            if false {
                v7 := []int{fi6, fi6, (fi6 | 127), fi6}
                v3 = v3
                _ = v7
            }
            _ = fi6
        }
        var v8 string = v3
        var v9 int64 = 42
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: int64(v2), F1: (-(v1))}
    v3 = "a"
    v11 := []int{32767, 32768, 9223372036854775807}
    v2 += 630
    v12 := GenS1{F0: 9223372036854775807, F1: v1}
    v13 := GenS2{F0: 127, F1: 42}
    fmt.Println(v2)
    var v14 string = v3
    if (255 == 748) {
        switch 2147483647 {
        case 4294967295:
            v2 -= 0
        case 42:
            var v15 int64 = 42
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            switch ((v14 + v3) + v14) {
            case "hello":
                v11 = append(v11, 32768)
            }
            v17 := GenS3{F0: 2147483647}
            _ = fi16
            _ = v17
        }
        var v18 float32 = 3.14159
        _ = v18
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            switch ((("hello" + ("foo bar" + v14)) + (v14 + "")) + v3) {
            case "\n\t":
                var v20 string = v3
                _ = v20
            default:
                v14 = v3
            }
            v21 := GenS4{F0: v2, F1: 106, F2: (fi19 / (((fi19 & 388)) | 1))}
            _ = fi19
            _ = v21
        }
    }
    if false {
        if len(v11) > 0 {
            v11[0] = 42
        }
        v2 = 1
        if (1 < -395) {
            v1 -= v1
            for fi22 := 0; fi22 < 2; fi22++ {
                if true {
                    v23 := GenS2{F0: 164, F1: 18446744073709551615}
                    if (true || ((true && true) && false)) {
                        if true {
                            fmt.Println(fi22)
                        }
                        v3 = "foo bar"
                        var v24 string = (v3 + "hello")
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 19; fi25++ {
                        v2 += v2
                        fmt.Println(v1)
                        _ = fi25
                    }
                    _ = v23
                } else {
                    if len(v11) > 0 {
                        v11[0] = (fi22 | (fi22 - fi22))
                    }
                    if false {
                        var v26 int8 = 42
                        v3 = ((v3 + v14) + v14)
                        _ = v26
                    } else {
                        v2 += (v2 % ((v2) | 1))
                    }
                    var v27 uint = 0
                    _ = v27
                }
                switch (fi22 + (fi22 >> ((696) & 63))) {
                case -2147483648:
                    switch v14 {
                    case "foo bar":
                        switch "hello" {
                        case "a":
                            if true {
                                var v28 int16 = (v1 + 32767)
                                if true {
                                    var v29 uint16 = v2
                                    v28 -= v1
                                    v28 += v28
                                    _ = v29
                                } else {
                                    for fi30 := 0; fi30 < 2; fi30++ {
                                        v28 += v28
                                        var v31 int64 = 42
                                        _ = fi30
                                        _ = v31
                                    }
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                }
                                _ = v28
                            }
                        case "\n\t":
                            switch (("a" + v14) + v14) {
                            case "foo bar":
                                v32 := GenS5{F0: 204, F1: 9223372036854775807}
                                _ = v32
                            default:
                                var v33 uint = 332
                                _ = v33
                            }
                        case "hello":
                            fmt.Println(v1)
                        default:
                            switch v3 {
                            case "hello":
                                var v34 int64 = int64(v4)
                                _ = v34
                            case "":
                                v35 := GenS3{F0: 531}
                                _ = v35
                            }
                        }
                    }
                case 4294967295:
                    for fi36 := 0; fi36 < 20; fi36++ {
                        _ = len(v11)
                        if false {
                            v11 = append(v11, (fi22 / (((fi36 >> ((0) & 63))) | 1)))
                        } else {
                            v37 := GenS6{F0: 342, F1: 13, F2: (fi22 / ((fi36) | 1))}
                            switch v14 {
                            case "a":
                                var v38 int16 = 0
                                _ = v38
                            default:
                                if (515 == 1) {
                                    v39 := GenS7{F0: 255, F1: 100, F2: 1}
                                    _ = v39
                                }
                            }
                            _ = v37
                        }
                        switch v14 {
                        case "foo bar":
                            fmt.Println(v2, v14, fi36)
                        case "hello":
                            switch v14 {
                            case "foo bar":
                                v1 += v1
                            default:
                                v1 -= 745
                                var v40 uint = 2
                                _ = v40
                            }
                            switch v3 {
                            case "a":
                                v2 -= v2
                                fmt.Println(v4)
                                switch (fi36 | (fi22 % ((int(v4)) | 1))) {
                                case 32767:
                                    v1 = (v1 * -1)
                                case 12:
                                    fmt.Println(fi36)
                                    v14 = (v14 + (v14 + v14))
                                    var v41 uint16 = v2
                                    _ = v41
                                case 9223372036854775807:
                                    v42 := GenS5{F0: -32768, F1: 9223372036854775807}
                                    _ = v42
                                default:
                                    switch v14 {
                                    case "foo bar":
                                        var v43 uint16 = v2
                                        _ = v43
                                    }
                                }
                            default:
                                var v44 uint32 = 1
                                _ = v44
                            }
                        }
                        _ = fi36
                    }
                case 265:
                    v45 := GenS8{F0: uint32(v4), F1: 127, F2: v1}
                    _ = v45
                default:
                    v46 := GenS9{F0: uint16(v4), F1: 18446744073709551615}
                    _ = v46
                }
                _ = fi22
            }
        }
    }
    v1 = (v1 >> ((2) & 63))
    v3 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
