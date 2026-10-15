package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 string }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8 }

func main() {
    v1 := []uint32{0, 1, 721, 991}
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    var v4 int8 = 100
    var v5 float64 = float64(3.14159)
    v6 := make(map[int]int)
    v7 := GenS1{F0: 0, F1: 0, F2: ("a" + "foo bar")}
    _ = len(v1)
    v5 = (v5 + -2.25)
    v4 -= v4
    v8 := GenS1{F0: 42, F1: -1, F2: "foo bar"}
    v4 = (v4 - v4)
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            v5 += float64((1e10 * (-2.25 / 0.1)))
            if false {
                if ((-316 <= (fi10 / ((fi10) | 1))) || (true || (-249 != 32767))) {
                    switch -2147483648 {
                    case 127:
                        v11 := GenS1{F0: 32767, F1: 42, F2: "a"}
                        _ = v11
                    case 32767:
                        v5 += v5
                    }
                    if true {
                        if (1 == 77) {
                            fmt.Println(v4)
                            var v12 uint8 = 1
                            if false {
                                var v13 uint32 = 2147483648
                                var v14 uint = 978
                                _ = v13
                                _ = v14
                            } else {
                                fmt.Println(fi10, fi9)
                                if len(v1) > 0 {
                                    v1[0] = 926
                                }
                                if false {
                                    if false {
                                        fmt.Println(v5, v4, fi10)
                                        var v15 bool = (65535 >= 42)
                                        var v16 bool = v15
                                        _ = v15
                                        _ = v16
                                    }
                                } else {
                                    v4 = v4
                                    delete(v2, 366)
                                    var v17 int16 = 32767
                                    _ = v17
                                }
                                if ("\n\t" >= "hello") {
                                    if false {
                                        v18 := GenS1{F0: 1, F1: int32(fi9), F2: "\n\t"}
                                        switch "hello" {
                                        case "\n\t":
                                            if true {
                                                v4 += (v4 - (v4 & int8(fi9)))
                                                v19 := GenS1{F0: 1, F1: -632, F2: ("\n\t" + "")}
                                                _ = v19
                                            }
                                            v20 := GenS1{F0: 214, F1: 2147483647, F2: "foo bar"}
                                            v21 := GenS1{F0: 1, F1: 316, F2: (("hello" + "") + ("foo bar" + (("a" + (("\n\t" + "foo bar") + ("\n\t" + "foo bar"))) + "\n\t")))}
                                            var v22 uint64 = 18446744073709551615
                                            var v23 int = (fi9 + fi10)
                                            v24 := GenS2{F0: (^(v12)), F1: 0}
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                        case "a":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case "":
                                            if (!false) {
                                                v12 -= v12
                                                var v25 int32 = 42
                                                _ = v25
                                            }
                                        default:
                                            if ((v12 ^ v12) <= v12) {
                                                var v26 uint32 = 65
                                                var v27 int64 = 1
                                                v26 += (v26 ^ 402)
                                                _ = v26
                                                _ = v27
                                            } else {
                                                v28 := GenS3{F0: 0}
                                                _ = v28
                                            }
                                        }
                                        if false {
                                            var v29 int = 0
                                            v30 := GenS4{F0: v4}
                                            delete(v3, 4294967295)
                                            _ = v29
                                            _ = v30
                                        } else {
                                            _ = len(v3)
                                            v5 -= v5
                                        }
                                        _ = v18
                                    }
                                    fmt.Println(fi10)
                                }
                                if true {
                                    v31 := GenS2{F0: 128, F1: 511}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v31
                                } else {
                                    fmt.Println(v12, fi10, v5)
                                }
                                switch (fi9 & fi9) {
                                case -847:
                                    var v32 bool = true
                                    _ = v32
                                case 32767:
                                    v5 = (v5 + v5)
                                default:
                                    if (false && false) {
                                        v33 := GenS5{F0: 301}
                                        _ = v33
                                    } else {
                                        var v34 uint64 = 733
                                        var v35 int = fi10
                                        var v36 string = ("hello" + "a")
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                }
                            }
                            if true {
                                v4 = v4
                                v5 -= (v5 / 3.14159)
                                v37 := GenS1{F0: 2, F1: 2147483647, F2: ("foo bar" + "")}
                                switch "" {
                                case "":
                                    fmt.Println(v4)
                                case "foo bar":
                                    v12 = 201
                                }
                                _ = v37
                            } else {
                                v5 = (float64(((1e10 * 1.5) * (-2.25 / ((1e10 - 3.14159) - 1e10)))) + 3.14159)
                                v12 -= 255
                            }
                            var v38 float64 = (3.14159 - v5)
                            v4 -= 42
                            _ = v12
                            _ = v38
                        } else {
                            var v39 uint16 = 42
                            v39 = 65535
                            _ = v39
                        }
                        v40 := GenS6{F0: -2147483648}
                        _ = v40
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v5 = 0.0
            }
            _ = fi10
        }
        var v41 bool = (true || false)
        _ = fi9
        _ = v41
    }
    var v42 uint64 = 4294967295
    switch 127 {
    case -58:
        if true {
            v4 -= v4
            if false {
                fmt.Println(v5, v42, v42)
                v1 = append(v1, 4294967295)
            } else {
                v43 := GenS7{F0: v4}
                _ = v43
            }
        } else {
            fmt.Println(v42, v5)
            if (!(42 <= -9223372036854775808)) {
                if false {
                    var v44 int = -9223372036854775808
                    v4 += v4
                    _ = v44
                }
                v42 = 0
            }
            var v45 bool = ((0.0 + float64((1.5 * (1e10 * 1e10)))) > v5)
            _ = v45
        }
    case 127:
        v4 = (-(v4))
    }
    v46 := GenS5{F0: 4294967295}
    var v47 string = "\n\t"
    switch 42 {
    case 1:
        if false {
            var v48 int32 = -257
            var v49 int16 = 724
            _ = v48
            _ = v49
        }
    case 127:
        _ = len(v1)
    case 280:
        if true {
            delete(v2, 2147483647)
        } else {
            fmt.Println(v42)
        }
    default:
        v50 := v6[-1]
        _ = v50
    }
    fmt.Println(v42, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
