package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 bool }
type GenS2 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS3 struct { F0 int; F1 string; F2 int16 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 int64; F1 int64 }
func genFn0(p0 uint32) uint32 {
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 int = (fi6 << ((938) & 63))
            if ((!true) && (true && true)) {
                v9 := GenS1{F0: -32768, F1: 172, F2: false}
                p0 = 1
                if false {
                    v8 = fi6
                }
                _ = v9
            } else {
                if false {
                    if (!true) {
                        switch fi6 {
                        case 32767:
                            v10 := []int{32767}
                            _ = v10
                        default:
                            p0 = 4294967295
                        }
                    } else {
                        p0 = (p0 ^ p0)
                        p0 += 1
                        v11 := GenS1{F0: -1, F1: 4294967295, F2: (398 == -657)}
                        _ = v11
                    }
                }
                if (true && true) {
                    v12 := GenS1{F0: 0, F1: 4294967295, F2: (18446744073709551615 >= 42)}
                    _ = v12
                }
            }
            _ = fi7
            _ = v8
        }
        var v13 uint = 65535
        v14 := GenS1{F0: 32767, F1: 0, F2: false}
        _ = fi6
        _ = v13
        _ = v14
    }
    var v15 string = "a"
    switch 32768 {
    case 1:
        if (!false) {
            var v16 int64 = -1
            var v17 string = "a"
            fmt.Println(p0, v16, v17)
            _ = v16
            _ = v17
        } else {
            var v18 int = 680
            _ = v18
        }
    case 550:
        v19 := GenS1{F0: 538, F1: 1, F2: (false && true)}
        _ = v19
    default:
        v15 = v15
    }
    v15 = ""
    _ = p0
    _ = v15
    return (p0 | p0)
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 32767
    var v3 float64 = 0.0
    var v4 float32 = (1e10 - (-2.25 + ((1.5 / float32(v3)) / (0.0 + 1e10))))
    var v5 uint32 = 65535
    v2 = (v2 >> ((65535) & 63))
    var v20 int32 = -630
    v21 := []int{127, 109, -159}
    v22 := GenS2{F0: 2, F1: 42, F2: (v2 * v2)}
    var v23 uint = 65535
    for fi24 := 0; fi24 < 18; fi24++ {
        var v25 uint8 = 255
        _ = fi24
        _ = v25
    }
    v26 := GenS3{F0: 80, F1: "a", F2: -32768}
    var v27 int64 = 1
    _ = len(v21)
    var v28 float32 = v4
    v21 = append(v21, 921)
    switch "\n\t" {
    case "hello":
        fmt.Println(v23, v23)
    case "\n\t":
        switch 42 {
        case 9223372036854775807:
            if (119 > 2) {
                if ((^(v2)) > v2) {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    v27 = v27
                    if (-537 < 0) {
                        switch -2147483648 {
                        case 996:
                            fmt.Println(v4, v20)
                        case 128:
                            fmt.Println(v2, v5, v2)
                        default:
                            v29 := GenS2{F0: 2, F1: 100, F2: (^(v2))}
                            _ = v29
                        }
                        var v30 int8 = 127
                        _ = v30
                    }
                }
                v31 := GenS4{F0: -32768, F1: 44}
                _ = v31
            }
        case 764:
            switch "\n\t" {
            case "":
                v20 = int32(v27)
            case "\n\t":
                if (true && false) {
                    var v32 uint32 = (v5 + (v5 << ((18446744073709551615) & 63)))
                    v33 := GenS2{F0: 1, F1: -128, F2: 173}
                    fmt.Println(v2, v5)
                    _ = v32
                    _ = v33
                } else {
                    var v34 bool = false
                    _ = v34
                }
            case "hello":
                v35 := GenS5{F0: (v27 << ((v23) & 63)), F1: (v27 << ((922) & 63))}
                _ = v35
            default:
                var v36 int32 = v20
                _ = v36
            }
        case 632:
            v20 = (v20 * 1)
        default:
            v27 -= v27
        }
    }
    fmt.Println(v28, v5)
    var v37 float64 = v3
    v38 := GenS4{F0: 908, F1: 255}
    if (false || (0 == 32767)) {
        if (((v28 - v4) < v4) || ((true || (!((("foo bar" + ("hello" + "a")) + "hello") <= "hello"))) && (false && (!false)))) {
            var v39 uint16 = (^(v2))
            for fi40 := 0; fi40 < 1; fi40++ {
                fmt.Println(v39, v20, v27)
                _ = fi40
            }
            _ = v39
        }
        if true {
            for fi41 := 0; fi41 < 12; fi41++ {
                var v42 uint = v23
                _ = fi41
                _ = v42
            }
            fmt.Println(v2)
        }
        fmt.Println(v23)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
