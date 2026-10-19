package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int16; F2 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS8 struct { F0 uint; F1 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 int32 = 42
    var v3 int = 128
    var v4 int8 = -64
    v5 := []uint32{665}
    for fi6 := 0; fi6 < 15; fi6++ {
        v1 += v1
        for fi7 := 0; fi7 < 2; fi7++ {
            v3 += fi7
            fmt.Println(v4, fi7)
            if (0 <= 42) {
                fmt.Println(fi6, fi6)
                v8 := GenS1{F0: "a", F1: 1}
                _ = v8
            }
            _ = fi7
        }
        v9 := GenS2{F0: (v2 & v2), F1: 2147483647, F2: 0}
        _ = fi6
        _ = v9
    }
    switch v3 {
    case 923:
        v2 -= -2147483648
    }
    v1 = (v1 & v1)
    _ = len(v5)
    var v10 uint8 = 139
    switch v3 {
    case 503:
        switch "\n\t" {
        case "foo bar":
            v4 += v4
        default:
            for fi11 := 0; fi11 < 12; fi11++ {
                for fi12 := 0; fi12 < 5; fi12++ {
                    if (!(42 >= 453)) {
                        v13 := GenS3{F0: 1}
                        if (!false) {
                            var v14 int32 = 0
                            fmt.Println(v14, v14)
                            _ = v14
                        } else {
                            var v15 uint64 = 1
                            _ = v15
                        }
                        _ = v13
                    } else {
                        v2 += v2
                        fmt.Println(v10)
                        v2 = (v2 + 2147483647)
                        if (!(65535 <= 65535)) {
                            if (42 == 42) {
                                v4 = v4
                            } else {
                                v16 := GenS3{F0: 849}
                                v5 = append(v5, 0)
                                var v17 uint16 = 1
                                _ = v16
                                _ = v17
                            }
                            fmt.Println(v1)
                            if (!true) {
                                v3 += (fi12 * (v3 % (((fi12 + fi11)) | 1)))
                            }
                        } else {
                            fmt.Println(v10)
                            if (uint16(v10) != 32767) {
                                if true {
                                    fmt.Println(v4)
                                    if false {
                                        v10 = (v10 * v10)
                                    }
                                } else {
                                    v18 := GenS3{F0: 4294967295}
                                    _ = v18
                                }
                                fmt.Println(v4)
                            } else {
                                var v19 string = "\n\t"
                                v19 = v19
                                var v20 int16 = 1000
                                v21 := GenS4{F0: 1, F1: v20, F2: 18446744073709551615}
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v2 += (v2 ^ (v2 - 42))
                            default:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            v2 += -2147483648
                        }
                        var v22 string = "\n\t"
                        _ = v22
                    }
                    _ = fi12
                }
                v3 += v3
                for fi23 := 0; fi23 < 6; fi23++ {
                    var v24 uint32 = 42
                    _ = fi23
                    _ = v24
                }
                _ = fi11
            }
        }
    case -954:
        var v25 bool = false
        _ = v25
    }
    if false {
        v2 = v2
        if (!(v4 >= v4)) {
            v26 := GenS4{F0: 42, F1: -1, F2: (^(v1))}
            _ = v26
        } else {
            if (v4 <= (v4 - 5)) {
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
                v3 = (v3 / ((2147483647) | 1))
            }
        }
        v27 := GenS1{F0: "foo bar", F1: 822}
        _ = v27
    } else {
        switch ("a" + "\n\t") {
        case "foo bar":
            if false {
                switch 834 {
                case 4294967295:
                    v28 := GenS5{F0: 1}
                    _ = v28
                case 128:
                    if (!false) {
                        var v29 float32 = -2.25
                        var v30 uint8 = v10
                        var v31 bool = true
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                case -2147483648:
                    for fi32 := 0; fi32 < 6; fi32++ {
                        for fi33 := 0; fi33 < 16; fi33++ {
                            if len(v5) > 0 {
                                v5[0] = 159
                            }
                            _ = fi33
                        }
                        v34 := GenS6{F0: 65535}
                        switch v3 {
                        case 128:
                            switch "hello" {
                            case "hello":
                                fmt.Println(v1)
                            case "a":
                                if false {
                                    v35 := GenS7{F0: 1, F1: 65535, F2: 65535}
                                    _ = v35
                                }
                            case "":
                                for fi36 := 0; fi36 < 5; fi36++ {
                                    fmt.Println(v10, fi32, v2)
                                    v2 = v2
                                    if true {
                                        v10 = (v10 >> ((v1) & 63))
                                    }
                                    _ = fi36
                                }
                            }
                        }
                        _ = fi32
                        _ = v34
                    }
                default:
                    v4 -= (v4 - 42)
                }
            } else {
                var v37 float32 = 0.1
                v2 = (v2 * (v2 - 42))
                if (2 <= 4294967295) {
                    fmt.Println(v3, v4)
                    if false {
                        v37 += 1e10
                        v37 -= v37
                    }
                }
                _ = v37
            }
        default:
            fmt.Println(v1, v3)
        }
        v1 = uint(v10)
        switch ("foo bar" + "") {
        case "":
            if (false && true) {
                var v38 float32 = 1e10
                _ = v38
            }
        case "a":
            for fi39 := 0; fi39 < 3; fi39++ {
                var v40 uint32 = 1
                fmt.Println(v10, v10, v4)
                _ = fi39
                _ = v40
            }
        case "hello":
            v10 = 42
        }
    }
    fmt.Println(v3)
    for fi41 := 0; fi41 < 5; fi41++ {
        if (((float64(0.1) <= (-2.25 * (1e10 * (float64(((3.14159 / 0.1) + (1.5 - (1e10 + (1.5 * ((1.5 * (3.14159 - 0.1)) - ((3.14159 + 0.0) / (1e10 - 1.5)))))))) * -2.25)))) && true) || true) {
            v42 := GenS6{F0: 1}
            v43 := GenS8{F0: (v1 * (v1 >> (((v1 >> ((v1) & 63))) & 63))), F1: 22}
            v44 := GenS8{F0: (v1 & v1), F1: 42}
            _ = v42
            _ = v43
            _ = v44
        }
        var v45 uint8 = v10
        _ = fi41
        _ = v45
    }
    var v46 bool = (605 >= 0)
    var v47 uint = v1
    v5 = append(v5, 2147483647)
    fmt.Println(v10)
    v48 := GenS7{F0: 1, F1: (v47 | v47), F2: 341}
    var v49 uint8 = (v10 + v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}
