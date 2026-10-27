package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint16 = 560
    var v2 uint = 2
    var v3 uint64 = uint64(v1)
    var v4 uint16 = (v1 ^ 255)
    v5 := []string{"foo bar", ((("a" + "\n\t") + "foo bar") + "\n\t"), "\n\t"}
    v6 := GenS1{F0: 18446744073709551615, F1: (^(v4)), F2: v1}
    var v7 string = "foo bar"
    v7 = v7
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            switch fi8 {
            case 42:
                v10 := GenS1{F0: (v3 - v3), F1: (v1 - (v1 | (v1 | v4))), F2: (v4 >> ((310) & 63))}
                _ = v10
            case -139:
                var v11 uint = (^(v2))
                _ = v11
            case -9223372036854775808:
                if (-2147483648 > (fi9 ^ fi9)) {
                    _ = len(v5)
                    v12 := GenS2{F0: 2147483647}
                    _ = v12
                } else {
                    switch fi9 {
                    case -578:
                        if true {
                            if false {
                                v1 -= 867
                                var v13 uint64 = (v3 << ((1) & 63))
                                _ = v13
                            } else {
                                fmt.Println(v7)
                            }
                            v14 := GenS2{F0: 65535}
                            _ = v14
                        }
                    }
                    v7 = "hello"
                    if false {
                        v15 := GenS3{F0: fi8, F1: v3}
                        _ = v15
                    } else {
                        v16 := GenS2{F0: 2147483647}
                        if true {
                            var v17 uint64 = 18446744073709551615
                            v17 -= (v3 ^ v3)
                            var v18 uint32 = 65535
                            v1 = v1
                            _ = v17
                            _ = v18
                        }
                        _ = v16
                    }
                    fmt.Println(v4)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
            var v19 float32 = (0.1 + 3.14159)
            _ = fi9
            _ = v19
        }
        var v20 int32 = -352
        var v21 int32 = v20
        _ = fi8
        _ = v20
        _ = v21
    }
    if false {
        var v22 int64 = int64(v3)
        for fi23 := 0; fi23 < 2; fi23++ {
            for fi24 := 0; fi24 < 3; fi24++ {
                fmt.Println(fi23, fi24)
                v3 -= v3
                if false {
                    var v25 int8 = 42
                    v4 += (v4 % ((v1) | 1))
                    var v26 uint16 = 32767
                    if false {
                        if ((v4 - 2) >= 32767) {
                            if ((v22 ^ (v22 >> ((v2) & 63))) == v22) {
                                v3 = v3
                                if (((3.14159 * 0.0) + 0.0) != 3.14159) {
                                    v3 += v3
                                    fmt.Println(v26)
                                } else {
                                    var v27 bool = (v25 > v25)
                                    if true {
                                        v7 = v7
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v28 uint64 = v3
                                        fmt.Println(v27, v22)
                                        var v29 int = (fi23 * fi23)
                                        _ = v28
                                        _ = v29
                                    } else {
                                        if (true || v27) {
                                            if (false && false) {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                if false {
                                                    var v30 int8 = (-(v25))
                                                    _ = v30
                                                }
                                            }
                                            var v31 bool = (v2 <= 4294967295)
                                            _ = v31
                                        }
                                        var v32 bool = false
                                        v26 -= (v1 << (((^(v2))) & 63))
                                        _ = v32
                                    }
                                    _ = v27
                                }
                                v2 += v2
                            } else {
                                var v33 float64 = 1.5
                                _ = v33
                            }
                        } else {
                            var v34 float64 = 1e10
                            _ = v34
                        }
                        v7 = (v7 + "hello")
                        switch v7 {
                        case "a":
                            fmt.Println(v7)
                        case "\n\t":
                            var v35 int = 127
                            _ = v35
                        case "":
                            switch v7 {
                            case "\n\t":
                                var v36 int32 = -1
                                _ = v36
                            case "foo bar":
                                v3 += (v3 & v3)
                            default:
                                v1 = (v26 % (((^(v4))) | 1))
                            }
                        }
                    } else {
                        var v37 uint64 = (v3 ^ (v3 + (v3 << ((v2) & 63))))
                        var v38 int8 = (v25 + (v25 - 42))
                        _ = v37
                        _ = v38
                    }
                    _ = v25
                    _ = v26
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v39 := GenS2{F0: 0}
                    _ = v39
                }
                _ = fi24
            }
            var v40 float32 = 1.5
            _ = fi23
            _ = v40
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v22
    }
    var v41 int32 = 0
    fmt.Println(v2, v1)
    var v42 int16 = 0
    if len(v5) > 0 {
        v5[0] = v7
    }
    if true {
        var v43 uint8 = 127
        switch -1 {
        case -358:
            _ = len(v5)
        default:
            if (((v7 + "\n\t") + v7) != v7) {
                v44 := GenS4{F0: (v2 + v2)}
                _ = v44
            }
        }
        _ = v43
    }
    v7 = (v7 + ("a" + "\n\t"))
    v45 := GenS3{F0: 0, F1: (^(v3))}
    switch v7 {
    case "\n\t":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "hello":
        var v46 uint = (v2 & 255)
        _ = v46
    case "foo bar":
        var v47 uint16 = v4
        _ = v47
    }
    v2 = (v2 / (((^(v2))) | 1))
    v42 -= (v42 % ((v42) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
