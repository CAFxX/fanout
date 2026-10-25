package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int8; F1 int64; F2 bool }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint32; F1 string; F2 uint16 }
type GenS6 struct { F0 int16; F1 uint8 }
func genFn0(p0 int64, p1 int64) int16 {
    p0 += (p1 >> ((412) & 63))
    if false {
        var v7 uint32 = 2147483647
        _ = v7
    } else {
        v8 := GenS1{F0: 2147483647, F1: 4294967295}
        _ = v8
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]uint32)
    var v3 bool = true
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]int)
    var v9 uint64 = 358
    var v10 int32 = v1
    if v3 {
        if (((1 < -1) && v3) && v3) {
            fmt.Println(v3, v9)
            var v11 uint64 = (v9 << ((65535) & 63))
            v1 = (v10 / ((42) | 1))
            _ = v11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(v1)
            _ = fi12
        }
    }
    var v13 float32 = float32(3.14159)
    fmt.Println(v9, v9)
    var v14 int32 = (v1 >> ((810) & 63))
    v15 := GenS1{F0: -787, F1: (v9 | v9)}
    v6[(((("" + "a") + "") + "foo bar") + "hello")] = -615
    for fi16 := 0; fi16 < 15; fi16++ {
        if (v9 <= v9) {
            v17 := []int{(fi16 & (fi16 << ((4294967295) & 63))), (-(fi16)), fi16, fi16}
            var v18 int16 = 32767
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 int32 = (v14 + v1)
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        }
        if v3 {
            v13 += 0.0
            v21 := GenS1{F0: (^(fi16)), F1: v9}
            _ = v21
        }
        _ = fi16
    }
    v22 := []int{-9223372036854775808, -681}
    switch "hello" {
    case "\n\t":
        v10 -= (v1 >> ((203) & 63))
    case "hello":
        var v23 int64 = -267
        _ = v23
    case "a":
        for fi24 := 0; fi24 < 18; fi24++ {
            for fi25 := 0; fi25 < 5; fi25++ {
                v26 := GenS1{F0: -1, F1: 1}
                if true {
                    if ("hello" != "foo bar") {
                        _ = len(v6)
                        if true {
                            delete(v6, "\n\t")
                        } else {
                            if v3 {
                                v9 += v9
                                v13 -= -2.25
                            } else {
                                v9 += v9
                                if true {
                                    switch "a" {
                                    case "\n\t":
                                        var v27 uint16 = 0
                                        _ = v27
                                    case "":
                                        var v28 string = "\n\t"
                                        _ = v28
                                    }
                                    var v29 int16 = 59
                                    _ = v29
                                }
                                switch (fi24 / (((fi24 % (((fi24 & fi25)) | 1))) | 1)) {
                                case -1:
                                    v13 = (1e10 - (float32(-2.25) * 1e10))
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    var v30 uint8 = 0
                                    _ = v30
                                case "":
                                    switch -704 {
                                    case 42:
                                        v31 := GenS2{F0: 100, F1: 42, F2: (2 <= 304)}
                                        _ = v31
                                    case 128:
                                        var v32 int = -2147483648
                                        _ = v32
                                    default:
                                        _ = len(v4)
                                    }
                                case "foo bar":
                                    fmt.Println(v14, fi24, fi24)
                                }
                                if v3 {
                                    switch "\n\t" {
                                    case "a":
                                        if true {
                                            fmt.Println(v14, v3, v14)
                                        } else {
                                            var v33 float32 = (v13 - (0.1 / (0.1 + -2.25)))
                                            v22 = append(v22, fi24)
                                            _ = v33
                                        }
                                    case "hello":
                                        var v34 float64 = (1e10 / (0.0 + (1.5 * 3.14159)))
                                        _ = v34
                                    case "\n\t":
                                        fmt.Println(v1, v3)
                                        switch "foo bar" {
                                        case "a":
                                            if true {
                                                fmt.Println(fi25, v14)
                                                var v35 float32 = v13
                                                _ = v35
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v36 := GenS2{F0: 100, F1: -1, F2: ((!true) && v3)}
                                                    var v37 int32 = (v14 | 42)
                                                    v38 := GenS3{F0: 0, F1: "a"}
                                                    _ = v36
                                                    _ = v37
                                                    _ = v38
                                                }
                                                v14 = 0
                                                v39 := GenS3{F0: 42, F1: ("hello" + "")}
                                                v40 := v2[2147483648]
                                                _ = v39
                                                _ = v40
                                            }
                                        case "":
                                            fmt.Println(v10)
                                        case "foo bar":
                                            var v41 uint = 18446744073709551615
                                            _ = v41
                                        }
                                    }
                                } else {
                                    var v42 float32 = ((v13 - 1.5) - 1e10)
                                    var v43 int = (fi24 << ((255) & 63))
                                    _ = v42
                                    _ = v43
                                }
                            }
                        }
                        _ = len(v2)
                        v44 := GenS4{F0: (-2147483648 >= (v1 << ((559) & 63))), F1: 1}
                        var v45 int = fi25
                        v9 = 2
                        v46 := GenS1{F0: (-(fi24)), F1: v9}
                        _ = v44
                        _ = v45
                        _ = v46
                    } else {
                        var v47 uint16 = 65535
                        _ = v47
                    }
                    _ = len(v22)
                    if v3 {
                        var v48 int64 = -411
                        fmt.Println(v13, fi25)
                        var v49 uint8 = 0
                        _ = v48
                        _ = v49
                    } else {
                        fmt.Println(v10)
                        var v50 uint8 = 2
                        _ = v50
                    }
                }
                fmt.Println(v3)
                _ = fi25
                _ = v26
            }
            var v51 uint = 4294967295
            var v52 uint16 = 255
            _ = fi24
            _ = v51
            _ = v52
        }
    default:
        if ((genFn0(-48, 1) > 32767) && false) {
            v53 := GenS4{F0: false, F1: 42}
            for fi54 := 0; fi54 < 11; fi54++ {
                v1 = (^(v10))
                _ = fi54
            }
            for fi55 := 0; fi55 < 6; fi55++ {
                if v3 {
                    var v56 uint64 = v9
                    if (true && true) {
                        v57 := GenS5{F0: 251, F1: "", F2: 0}
                        var v58 int16 = -763
                        for fi59 := 0; fi59 < 17; fi59++ {
                            v60 := GenS6{F0: -1, F1: 42}
                            _ = fi59
                            _ = v60
                        }
                        _ = v57
                        _ = v58
                    }
                    v56 = 0
                    _ = v56
                }
                _ = fi55
            }
            _ = v53
        } else {
            for fi61 := 0; fi61 < 14; fi61++ {
                if v3 {
                    v62 := GenS6{F0: -32768, F1: 108}
                    if ((v10 >> ((530) & 63)) != v14) {
                        var v63 float64 = 1e10
                        v13 += v13
                        fmt.Println(v9)
                        _ = v63
                    }
                    _ = v62
                }
                _ = fi61
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
