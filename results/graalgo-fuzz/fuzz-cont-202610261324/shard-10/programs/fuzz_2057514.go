package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint; F2 int8 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint8; F1 uint32; F2 int }
type GenS6 struct { F0 bool; F1 uint; F2 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8; F1 uint32; F2 uint8 }
func genFn0(p0 int) int32 {
    var v4 uint64 = 42
    p0 = -755
    fmt.Println(p0, v4)
    fmt.Println(p0, p0)
    var v5 uint32 = 2147483647
    _ = p0
    _ = v4
    _ = v5
    return -2147483648
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"a"}
    var v3 int64 = 9223372036854775807
    var v6 bool = false
    if len(v2) > 0 {
        v2[0] = ("foo bar" + "a")
    }
    if v6 {
        _ = len(v2)
    }
    if (4294967295 == -2147483648) {
        fmt.Println(v3, v3)
        switch 1 {
        case -530:
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 uint32 = 4294967295
                switch fi7 {
                case 42:
                    var v9 uint64 = 313
                    _ = v9
                case 32768:
                    v10 := GenS1{F0: 560, F1: v8}
                    _ = v10
                default:
                    _ = len(v1)
                }
                _ = fi7
                _ = v8
            }
        }
    }
    if true {
        v6 = (42 == -117)
        for fi11 := 0; fi11 < 11; fi11++ {
            v2 = append(v2, "\n\t")
            switch 128 {
            case -9223372036854775808:
                var v12 int32 = -485
                _ = v12
            case 127:
                _ = len(v2)
            case 42:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            for fi13 := 0; fi13 < 12; fi13++ {
                v3 = (v3 * v3)
                switch "foo bar" {
                case "\n\t":
                    v14 := GenS2{F0: 9223372036854775807, F1: 823, F2: 0}
                    _ = v14
                case "":
                    fmt.Println(fi11)
                }
                _ = fi13
            }
            _ = fi11
        }
        v3 = (v3 % ((v3) | 1))
    } else {
        fmt.Println(v3, v6)
        var v15 uint32 = 42
        _ = v15
    }
    var v16 int32 = 1
    var v17 bool = v6
    _ = len(v1)
    var v18 int32 = (v16 * -265)
    if (v6 || (("\n\t" + "a") != "\n\t")) {
        v19 := GenS3{F0: 1, F1: 2}
        for fi20 := 0; fi20 < 17; fi20++ {
            if (334 >= v3) {
                var v21 uint64 = 42
                for fi22 := 0; fi22 < 20; fi22++ {
                    _ = len(v1)
                    switch "a" {
                    case "hello":
                        fmt.Println(v16, v3, v17)
                    }
                    var v23 bool = v6
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi22
                    _ = v23
                }
                switch 32767 {
                case 454:
                    if v17 {
                        var v24 int64 = v3
                        v25 := GenS4{F0: 1000, F1: 42}
                        v3 += (v3 % ((v24) | 1))
                        _ = v24
                        _ = v25
                    }
                case 1:
                    switch ("a" + ("" + "hello")) {
                    case "a":
                        v16 += v18
                    case "":
                        v26 := GenS2{F0: v3, F1: 2, F2: 1}
                        _ = v26
                    }
                case 9223372036854775807:
                    for fi27 := 0; fi27 < 20; fi27++ {
                        if (true || v17) {
                            v2 = append(v2, ("hello" + "foo bar"))
                            fmt.Println(v17)
                        }
                        switch (fi27 + fi27) {
                        case -9223372036854775808:
                            v28 := GenS2{F0: -704, F1: 42, F2: -1}
                            _ = v28
                        case 32768:
                            var v29 int8 = -128
                            _ = v29
                        case -2147483648:
                            v30 := GenS5{F0: 117, F1: 42, F2: fi27}
                            _ = v30
                        }
                        _ = fi27
                    }
                default:
                    v21 -= v21
                }
                _ = v21
            } else {
                fmt.Println(v18)
            }
            _ = fi20
        }
        if v6 {
            var v31 int32 = v16
            _ = v31
        } else {
            fmt.Println(v16, v17, v17)
        }
        _ = v19
    }
    v17 = v17
    v32 := GenS2{F0: v3, F1: 0, F2: -1}
    switch 1 {
    case -9223372036854775808:
        var v33 int = -922
        _ = v33
    case 1:
        fmt.Println(v3, v6)
    default:
        var v34 uint32 = 441
        _ = v34
    }
    switch "foo bar" {
    case "hello":
        if v6 {
            if v17 {
                var v35 uint16 = 42
                _ = v35
            }
            switch -121 {
            case -2147483648:
                var v36 float64 = float64(float32(-2.25))
                _ = v36
            }
            v16 = (v16 ^ (-(v16)))
        }
    case "":
        if (1 == 0) {
            v37 := GenS6{F0: (127 < 84), F1: 65535, F2: "\n\t"}
            _ = v37
        } else {
            var v38 int32 = (^(v16))
            delete(v1, 1)
            for fi39 := 0; fi39 < 8; fi39++ {
                _ = len(v2)
                _ = fi39
            }
            _ = v38
        }
    case "a":
        _ = len(v2)
    default:
        v16 = 1
    }
    v40 := v1[0]
    if v17 {
        v41 := GenS7{F0: "a"}
        fmt.Println(v16, v6, v6)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v41
    } else {
        var v42 int16 = 1000
        switch 128 {
        case -108:
            switch "hello" {
            case "a":
                switch 128 {
                case 32767:
                    v16 -= v18
                }
            case "\n\t":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "foo bar":
                v43 := GenS5{F0: 127, F1: 4294967295, F2: 32767}
                _ = v43
            }
        default:
            v44 := GenS8{F0: 127, F1: 698, F2: 2}
            _ = v44
        }
        v45 := v1[42]
        _ = v42
        _ = v45
    }
    v46 := GenS2{F0: (v3 & (v3 | (v3 % ((876) | 1)))), F1: 42, F2: -113}
    var v47 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v40
    _ = v46
    _ = v47
    _ = v6
}
