package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int8; F2 string }
type GenS3 struct { F0 uint64; F1 string; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 int64) uint32 {
    p0 -= p0
    for fi7 := 0; fi7 < 17; fi7++ {
        p0 = (p0 * p0)
        _ = fi7
    }
    _ = p0
    return 121
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]string)
    var v3 uint8 = 147
    var v4 float32 = 1e10
    var v5 bool = (!(702 != -844))
    var v6 string = "\n\t"
    v8 := GenS1{F0: (v3 | v3)}
    var v9 uint16 = (v1 * v1)
    switch (v6 + v6) {
    case "hello":
        switch -423 {
        case -2147483648:
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := []int{(fi10 << ((327) & 63)), (fi10 & (fi10 % ((4294967295) | 1))), (fi10 | 956)}
                if ((((0.0 + 3.14159) - v4) / v4) > 0.1) {
                    for fi12 := 0; fi12 < 6; fi12++ {
                        if (0 != 2) {
                            v13 := GenS2{F0: 986, F1: 42, F2: (v6 + "a")}
                            fmt.Println(fi12)
                            var v14 int16 = 0
                            if (!(!(!(!(!v5))))) {
                                switch fi10 {
                                case -2147483648:
                                    if len(v11) > 0 {
                                        v11[0] = fi10
                                    }
                                }
                            } else {
                                v2[("hello" + v6)] = v6
                                fmt.Println(v5, v6)
                                v3 = v3
                                fmt.Println(v1)
                                var v15 int8 = 1
                                _ = v15
                            }
                            _ = v13
                            _ = v14
                        } else {
                            if (v9 >= 1) {
                                fmt.Println(fi10)
                            } else {
                                v16 := v2[v6]
                                if ((v4 < (v4 / float32((1.5 / (3.14159 - ((1.5 / (1.5 - 3.14159)) + (0.0 * 0.0))))))) || v5) {
                                    var v17 int8 = -128
                                    _ = v17
                                } else {
                                    switch ("foo bar" + v16) {
                                    case "":
                                        v3 -= (v3 / (((v3 - v3)) | 1))
                                    case "hello":
                                        fmt.Println(v16, v5)
                                    case "\n\t":
                                        var v18 uint16 = (v9 << ((uint(v3)) & 63))
                                        _ = v18
                                    }
                                }
                                _ = v16
                            }
                            v4 -= -2.25
                            switch (fi10 % (((fi12 % ((fi10) | 1))) | 1)) {
                            case 257:
                                fmt.Println(v3, v3)
                            default:
                                var v19 int8 = 38
                                _ = v19
                            }
                        }
                        _ = fi12
                    }
                    var v20 uint16 = (v1 & (v1 >> ((430) & 63)))
                    if v5 {
                        for fi21 := 0; fi21 < 11; fi21++ {
                            var v22 bool = v5
                            _ = fi21
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 8; fi23++ {
                            v24 := GenS1{F0: v3}
                            switch v6 {
                            case "\n\t":
                                var v25 int8 = 102
                                _ = v25
                            default:
                                v26 := GenS3{F0: 42, F1: "", F2: 1}
                                _ = v26
                            }
                            _ = fi23
                            _ = v24
                        }
                        for fi27 := 0; fi27 < 10; fi27++ {
                            var v28 string = v6
                            fmt.Println(v3, v4)
                            _ = fi27
                            _ = v28
                        }
                    } else {
                        v6 = ""
                        var v29 float32 = v4
                        v3 = v3
                        _ = v29
                    }
                    _ = v20
                }
                _ = fi10
                _ = v11
            }
        }
    case "":
        if (-324 > 399) {
            v30 := GenS4{F0: 283}
            _ = v30
        } else {
            var v31 uint32 = 2147483647
            v1 += (v1 >> ((0) & 63))
            _ = v31
        }
    default:
        var v32 int32 = 1
        _ = v32
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 int16 = 1000
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        v36 := GenS5{F0: -1}
        _ = fi35
        _ = v36
    }
    v3 -= 42
    if (((v4 >= (v4 + float32((0.1 * 1e10)))) || (true || (v5 && (!false)))) || ("a" != "")) {
        switch v6 {
        case "":
            delete(v2, v6)
        case "hello":
            v37 := GenS6{F0: -1, F1: 2}
            _ = v37
        case "a":
            v4 -= 3.14159
        default:
            v38 := GenS4{F0: (v9 * 65535)}
            _ = v38
        }
        v39 := GenS7{F0: 4294967295}
        var v40 int32 = -1
        _ = v39
        _ = v40
    } else {
        v4 += 1.5
        fmt.Println(v1, v4)
        var v41 float32 = ((v4 * v4) / 3.14159)
        _ = v41
    }
    v3 += 127
    switch "a" {
    case "":
        fmt.Println(v3, v9, v4)
    default:
        if true {
            switch 42 {
            case 9223372036854775807:
                v42 := []int{-461, -9223372036854775808}
                _ = v42
            case 4294967295:
                if (v5 && v5) {
                    v1 = v9
                    v43 := GenS8{F0: (v1 % (((v1 | (^(v1)))) | 1))}
                    v44 := GenS1{F0: v3}
                    _ = v43
                    _ = v44
                }
            case -1:
                fmt.Println(v3, v9)
            default:
                var v45 string = v6
                _ = v45
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
