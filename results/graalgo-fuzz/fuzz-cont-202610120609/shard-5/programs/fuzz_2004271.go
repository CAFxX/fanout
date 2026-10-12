package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 string; F1 string; F2 uint16 }
type GenS4 struct { F0 uint32; F1 string; F2 string }
type GenS5 struct { F0 int; F1 uint8; F2 string }
type GenS6 struct { F0 uint8; F1 int }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint16 = 1
    var v3 float32 = (((1.5 / float32(1.5)) * 1.5) - 0.1)
    var v4 bool = false
    v5 := []uint32{65535, 2, 4294967295, 222}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v4, v4, v4)
    if false {
        var v6 float32 = float32(float64(0.0))
        v4 = false
        v3 -= (v6 * (v6 - v3))
        _ = v6
    } else {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint32 = 601
            var v9 int = (fi7 + -177)
            _ = len(v5)
            _ = fi7
            _ = v8
            _ = v9
        }
        v5 = append(v5, uint32(v1))
        if len(v5) > 0 {
            v5[0] = 584
        }
    }
    fmt.Println(v2, v1)
    v3 -= 1.5
    v4 = (v4 && true)
    fmt.Println(v4)
    switch "hello" {
    case "foo bar":
        var v10 bool = ((0.1 / v3) <= v3)
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: 1}
            if (v2 >= (v2 / ((v2) | 1))) {
                v5 = append(v5, 513)
            } else {
                var v13 uint = v1
                if len(v5) > 0 {
                    v5[0] = 42
                }
                _ = len(v5)
                _ = v13
            }
            _ = len(v5)
            _ = fi11
            _ = v12
        }
    }
    switch 128 {
    case 128:
        if len(v5) > 0 {
            v5[0] = 65535
        }
    case -1:
        if (4294967295 > 1) {
            fmt.Println(v3, v2)
        } else {
            var v14 string = ("" + "hello")
            var v15 uint16 = 32767
            _ = v14
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 14; fi16++ {
            v4 = true
            var v17 string = "foo bar"
            _ = fi16
            _ = v17
        }
    }
    v4 = (0 > int32(v1))
    var v18 bool = (-1 < 1)
    if v18 {
        v1 += (v1 - 255)
    } else {
        fmt.Println(v4, v4)
        var v19 uint8 = 67
        switch 32768 {
        case 42:
            if (true || v4) {
                var v20 float32 = v3
                for fi21 := 0; fi21 < 1; fi21++ {
                    v5 = append(v5, 0)
                    _ = fi21
                }
                _ = v20
            } else {
                if (3.14159 != 3.14159) {
                    if v18 {
                        v22 := GenS2{F0: -1, F1: true}
                        switch 199 {
                        case -428:
                            if len(v5) > 0 {
                                v5[0] = 804
                            }
                        case -243:
                            v23 := GenS3{F0: "a", F1: ("foo bar" + ("hello" + "\n\t")), F2: (v2 - v2)}
                            _ = v23
                        case 127:
                            v24 := GenS4{F0: 2, F1: "hello", F2: "hello"}
                            _ = v24
                        default:
                            fmt.Println(v1)
                        }
                        v18 = true
                        _ = v22
                    } else {
                        fmt.Println(v18, v4)
                        for fi25 := 0; fi25 < 14; fi25++ {
                            for fi26 := 0; fi26 < 15; fi26++ {
                                v4 = v4
                                if len(v5) > 0 {
                                    v5[0] = 4294967295
                                }
                                if v4 {
                                    v5 = append(v5, 42)
                                    var v27 uint32 = 65535
                                    var v28 int8 = 42
                                    v4 = (!true)
                                    _ = v27
                                    _ = v28
                                } else {
                                    if v4 {
                                        var v29 int32 = int32(v19)
                                        if false {
                                            var v30 uint8 = v19
                                            _ = v30
                                        } else {
                                            switch "hello" {
                                            case "":
                                                var v31 int16 = -1
                                                _ = v31
                                            }
                                            v32 := GenS3{F0: "a", F1: ("foo bar" + "hello"), F2: v2}
                                            v33 := GenS1{F0: (v19 | (v19 * v19))}
                                            _ = v32
                                            _ = v33
                                        }
                                        _ = v29
                                    }
                                    v4 = (int64(v19) == -1)
                                }
                                switch ("a" + "\n\t") {
                                case "foo bar":
                                    switch (fi25 | (fi25 + (fi25 >> ((255) & 63)))) {
                                    case 42:
                                        fmt.Println(fi25)
                                    case 9223372036854775807:
                                        _ = len(v5)
                                    case 64:
                                        switch fi26 {
                                        case 680:
                                            var v34 uint = v1
                                            _ = v34
                                        }
                                    default:
                                        if (0 > -1) {
                                            v19 = 0
                                            var v35 uint64 = 0
                                            _ = v35
                                        } else {
                                            var v36 uint = (^(v1))
                                            fmt.Println(v1, v3, fi25)
                                            switch -250 {
                                            case 1:
                                                _ = len(v5)
                                            case 4294967295:
                                                v36 += v36
                                            default:
                                                v37 := GenS4{F0: 2147483648, F1: "hello", F2: ""}
                                                _ = v37
                                            }
                                            v38 := GenS5{F0: 1, F1: v19, F2: "a"}
                                            _ = v36
                                            _ = v38
                                        }
                                    }
                                default:
                                    if v18 {
                                        switch fi25 {
                                        case 0:
                                            var v39 string = "foo bar"
                                            _ = v39
                                        case 42:
                                            var v40 uint = 0
                                            _ = v40
                                        case -744:
                                            if v18 {
                                                var v41 uint = 250
                                                v1 -= (v1 - 320)
                                                _ = v41
                                            }
                                        }
                                    } else {
                                        var v42 uint64 = uint64(v2)
                                        _ = v42
                                    }
                                    if v18 {
                                        if len(v5) > 0 {
                                            v5[0] = 2147483648
                                        }
                                    }
                                }
                                if len(v5) > 0 {
                                    v5[0] = 1
                                }
                                _ = fi26
                            }
                            _ = fi25
                        }
                    }
                    v43 := GenS5{F0: 559, F1: 255, F2: "a"}
                    switch int(v1) {
                    case 2147483647:
                        var v44 float64 = 1e10
                        _ = v44
                    case 128:
                        v19 -= (^(v19))
                    default:
                        v45 := GenS6{F0: (v19 | 2), F1: 9223372036854775807}
                        _ = v45
                    }
                    _ = v43
                }
                for fi46 := 0; fi46 < 19; fi46++ {
                    v4 = false
                    _ = fi46
                }
            }
        case 1:
            for fi47 := 0; fi47 < 15; fi47++ {
                v1 = (v1 & 533)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi47
            }
        case -585:
            fmt.Println(v3)
        }
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
