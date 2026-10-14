package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 string; F1 bool; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 int32; F2 int32 }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 bool; F1 uint8 }
func genFn0() uint {
    var v5 int8 = 127
    v6 := []int{}
    if (false || ("" > "\n\t")) {
        if false {
            fmt.Println(v5)
            v5 = 1
            var v7 string = "foo bar"
            _ = v7
        }
    } else {
        var v8 uint64 = 622
        if (42 == 42) {
            v5 += (v5 - (v5 | v5))
            v9 := GenS1{F0: 628, F1: 2147483647}
            for fi10 := 0; fi10 < 3; fi10++ {
                for fi11 := 0; fi11 < 2; fi11++ {
                    switch "\n\t" {
                    case "":
                        if true {
                            fmt.Println(v5, fi10, fi10)
                            fmt.Println(v8)
                            fmt.Println(fi10)
                        }
                    case "foo bar":
                        var v12 uint8 = 2
                        _ = v12
                    default:
                        if ((fi11 > (fi10 | 0)) && false) {
                            v5 = (v5 / (((v5 & (v5 + (v5 | v5)))) | 1))
                            var v13 float32 = ((-2.25 / float32((((3.14159 / 1.5) * 1.5) * ((0.1 * -2.25) + (0.0 * ((float64((3.14159 * (-2.25 * 3.14159))) + -2.25) + 0.1)))))) / float32(0.1))
                            _ = v13
                        }
                    }
                    if false {
                        v5 = 127
                        var v14 float32 = (float32(0.0) - 3.14159)
                        var v15 bool = (0 >= 346)
                        _ = v14
                        _ = v15
                    } else {
                        var v16 int8 = (v5 - (v5 % ((-128) | 1)))
                        v5 += v16
                        _ = v16
                    }
                    v17 := GenS2{F0: "foo bar", F1: false, F2: ((("foo bar" + "hello") + "a") + "hello")}
                    _ = fi11
                    _ = v17
                }
                v18 := GenS3{F0: 2}
                _ = fi10
                _ = v18
            }
            _ = v9
        }
        v19 := GenS1{F0: (v8 ^ v8), F1: 2147483647}
        _ = v19
        _ = v8
    }
    var v20 string = ""
    switch ("" + v20) {
    case "hello":
        v21 := GenS3{F0: 4294967295}
        _ = v21
    default:
        fmt.Println(v5, v5)
    }
    _ = v20
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 bool = (0 <= 42)
    v4 := []int{-660, -9223372036854775808, 32768}
    if (0 > 76) {
        var v22 float64 = 1.5
        _ = v22
    } else {
        v23 := v2[("\n\t" + "a")]
        _ = v23
    }
    v3 = (127 > 2)
    for fi24 := 0; fi24 < 3; fi24++ {
        fmt.Println(v3, v3)
        v25 := GenS4{F0: 1}
        _ = fi24
        _ = v25
    }
    if v3 {
        v26 := GenS2{F0: "hello", F1: true, F2: ("" + "\n\t")}
        if (v3 || (2147483647 >= 1)) {
            var v27 string = "hello"
            _ = v27
        } else {
            fmt.Println(v3)
            fmt.Println(v3)
        }
        _ = v26
    }
    v28 := GenS5{F0: -2147483648, F1: -2147483648, F2: -330}
    var v29 float32 = 0.0
    v3 = true
    _ = len(v4)
    fmt.Println(v3, v3)
    if (!v3) {
        _ = len(v4)
        for fi30 := 0; fi30 < 16; fi30++ {
            v31 := GenS2{F0: "foo bar", F1: v3, F2: ""}
            if v3 {
                v3 = true
                if true {
                    fmt.Println(v3)
                } else {
                    fmt.Println(v29, v3, fi30)
                }
                v32 := GenS4{F0: 42}
                _ = v32
            }
            if true {
                for fi33 := 0; fi33 < 16; fi33++ {
                    v4 = append(v4, (^(fi30)))
                    _ = fi33
                }
            } else {
                v3 = (187 == 1)
                for fi34 := 0; fi34 < 14; fi34++ {
                    v3 = v3
                    var v35 string = ""
                    var v36 uint64 = 42
                    delete(v2, v35)
                    v36 -= v36
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
            _ = fi30
            _ = v31
        }
    }
    switch 42 {
    case 128:
        v37 := GenS6{F0: 42, F1: 165}
        _ = v37
    case 414:
        v29 -= v29
    case 9223372036854775807:
        var v38 uint64 = 18446744073709551615
        _ = v38
    default:
        v29 = v29
    }
    v29 = v29
    for fi39 := 0; fi39 < 11; fi39++ {
        delete(v2, "hello")
        v3 = v3
        switch fi39 {
        case 4294967295:
            for fi40 := 0; fi40 < 1; fi40++ {
                v4 = append(v4, -2147483648)
                _ = fi40
            }
        }
        _ = fi39
    }
    v3 = v3
    fmt.Println(v29, v3)
    for fi41 := 0; fi41 < 4; fi41++ {
        v42 := GenS3{F0: 0}
        switch "hello" {
        case "":
            fmt.Println(v29, fi41)
        case "foo bar":
            switch fi41 {
            case 912:
                for fi43 := 0; fi43 < 20; fi43++ {
                    if (740 != 255) {
                        var v44 int = fi43
                        v44 += fi43
                        v45 := GenS7{F0: "", F1: 127}
                        _ = v44
                        _ = v45
                    } else {
                        if (!(!v3)) {
                            v29 = v29
                            var v46 uint8 = 127
                            _ = v46
                        } else {
                            fmt.Println(v3, v29)
                            var v47 uint16 = 255
                            fmt.Println(fi43, fi41, v47)
                            _ = v47
                        }
                    }
                    _ = fi43
                }
            case -666:
                v48 := v2[""]
                _ = v48
            }
        default:
            v29 = 1e10
        }
        _ = fi41
        _ = v42
    }
    for fi49 := 0; fi49 < 6; fi49++ {
        var v50 int16 = 32767
        v29 += (((float32(3.14159) + (v29 - (v29 * v29))) * 0.0) / v29)
        switch "\n\t" {
        case "foo bar":
            v29 = (1.5 * (v29 * (v29 / v29)))
        case "\n\t":
            v51 := GenS8{F0: v3, F1: 0}
            _ = v51
        }
        _ = fi49
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
