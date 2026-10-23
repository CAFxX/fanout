package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 42
    var v3 int16 = 1000
    v4 := make(map[string]string)
    v5 := GenS1{F0: 312, F1: -1, F2: -8}
    v6 := GenS1{F0: -891, F1: 1, F2: (v3 << ((2) & 63))}
    var v7 int64 = -9223372036854775808
    if (!("foo bar" == ("" + "foo bar"))) {
        var v8 int32 = 42
        var v9 float32 = (((3.14159 + ((1.5 * 1.5) + (float32(1e10) / -2.25))) + float32((-2.25 - 3.14159))) + 0.1)
        for fi10 := 0; fi10 < 14; fi10++ {
            fmt.Println(v8, v9, v3)
            for fi11 := 0; fi11 < 7; fi11++ {
                _ = len(v4)
                v9 = v9
                fmt.Println(fi11, fi10)
                _ = fi11
            }
            v9 -= v9
            _ = fi10
        }
        _ = v8
        _ = v9
    }
    fmt.Println(v2)
    fmt.Println(v2, v7)
    if false {
        switch "\n\t" {
        case "":
            if false {
                if false {
                    if (true || (!true)) {
                        v12 := GenS2{F0: 2, F1: -1}
                        v13 := GenS3{F0: v2, F1: 127}
                        _ = v12
                        _ = v13
                    }
                    v14 := []int{0, 128, -118, int(v7)}
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        var v16 int = -2147483648
                        for fi17 := 0; fi17 < 2; fi17++ {
                            v7 = (v7 & v7)
                            var v18 int = 127
                            _ = fi17
                            _ = v18
                        }
                        if false {
                            v19 := []int{(v16 * 32767), (v16 + (v16 * v16)), v16, (v16 / ((127) | 1))}
                            if true {
                                for fi20 := 0; fi20 < 15; fi20++ {
                                    v7 -= -1
                                    v21 := GenS3{F0: 2, F1: 110}
                                    v2 = v2
                                    if (!false) {
                                        var v22 uint8 = (v2 * (v2 << ((902) & 63)))
                                        if false {
                                            var v23 bool = (false || (4294967295 < 970))
                                            _ = v23
                                        } else {
                                            v24 := GenS3{F0: v22, F1: -128}
                                            _ = v24
                                        }
                                        switch ("\n\t" + ("" + (("hello" + ("foo bar" + "hello")) + "foo bar"))) {
                                        case "\n\t":
                                            switch "\n\t" {
                                            case "a":
                                                var v25 uint16 = 85
                                                _ = v25
                                            case "":
                                                if (!false) {
                                                    var v26 int16 = v3
                                                    v3 += v26
                                                    switch ("\n\t" + "") {
                                                    case "hello":
                                                        var v27 float64 = ((((1e10 - float64((1.5 * 1.5))) - 0.0) * 0.0) * 1e10)
                                                        v27 = v27
                                                        _ = v27
                                                    default:
                                                        v3 -= v3
                                                    }
                                                    _ = v26
                                                }
                                            case "hello":
                                                var v28 uint16 = 255
                                                _ = v28
                                            }
                                        case "":
                                            var v29 uint16 = 107
                                            _ = v29
                                        case "foo bar":
                                            v22 -= 2
                                            var v30 uint64 = 2
                                            if len(v19) > 0 {
                                                v19[0] = 808
                                            }
                                            v7 = v7
                                            v2 = 162
                                            _ = v30
                                        default:
                                            v3 = v3
                                        }
                                        delete(v1, 0)
                                        v16 += fi20
                                        switch (("a" + (("hello" + "foo bar") + "\n\t")) + ("\n\t" + (("\n\t" + "") + "foo bar"))) {
                                        case "":
                                            var v31 float32 = (0.0 / (-2.25 / (1e10 / 0.1)))
                                            _ = v31
                                        default:
                                            v3 += (v3 / (((-(v3))) | 1))
                                        }
                                        _ = v22
                                    }
                                    _ = fi20
                                    _ = v21
                                }
                            } else {
                                v3 += 1
                                var v32 uint64 = 0
                                _ = v32
                            }
                            v3 += (^(v3))
                            _ = v19
                        }
                        _ = fi15
                        _ = v16
                    }
                }
                var v33 float32 = (0.0 - 3.14159)
                delete(v4, "a")
                _ = v33
            } else {
                var v34 int16 = v3
                if (0.0 != ((float32(0.0) - 3.14159) + (1.5 / 1.0))) {
                    var v35 int = -92
                    _ = v35
                } else {
                    if false {
                        var v36 uint32 = 1000
                        _ = v36
                    } else {
                        v37 := []int{2147483647}
                        var v38 float32 = 0.1
                        _ = v37
                        _ = v38
                    }
                    if true {
                        if false {
                            fmt.Println(v3, v7, v2)
                        } else {
                            var v39 int8 = -117
                            var v40 uint = 199
                            _ = v39
                            _ = v40
                        }
                        for fi41 := 0; fi41 < 1; fi41++ {
                            var v42 float64 = 3.14159
                            _ = fi41
                            _ = v42
                        }
                    } else {
                        if (-67 >= -128) {
                            v43 := GenS2{F0: uint16(v2), F1: 108}
                            _ = v43
                        }
                    }
                }
                v7 = (v7 << ((1) & 63))
                _ = v34
            }
        case "hello":
            for fi44 := 0; fi44 < 1; fi44++ {
                if ((v7 | (v7 ^ (^(v7)))) <= v7) {
                    v3 = v3
                    v3 -= (v3 << ((4294967295) & 63))
                } else {
                    fmt.Println(fi44, fi44, v2)
                }
                _ = fi44
            }
        default:
            v45 := GenS4{F0: (v7 % (((v7 - v7)) | 1))}
            _ = v45
        }
    } else {
        var v46 int32 = 624
        switch "\n\t" {
        case "\n\t":
            for fi47 := 0; fi47 < 16; fi47++ {
                fmt.Println(v46, fi47)
                v48 := GenS5{F0: 365}
                v2 = (v2 | v2)
                _ = fi47
                _ = v48
            }
        case "foo bar":
            v46 += (v46 - v46)
        case "hello":
            v49 := GenS3{F0: v2, F1: 0}
            _ = v49
        }
        v7 += int64(v2)
        _ = v46
    }
    for fi50 := 0; fi50 < 3; fi50++ {
        v51 := v1[545]
        var v52 uint32 = 563
        var v53 uint8 = (v2 - v2)
        _ = fi50
        _ = v51
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 19; fi54++ {
        var v55 string = ("" + "\n\t")
        switch fi54 {
        case 789:
            fmt.Println(v2)
        case 32767:
            v7 = (v7 / ((v7) | 1))
        case 32768:
            var v56 uint32 = 4294967295
            _ = v56
        }
        _ = fi54
        _ = v55
    }
    v7 -= v7
    var v57 bool = (42 == 42)
    v2 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v57:", v57)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v7
}
