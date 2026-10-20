package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint }
func genFn0(p0 uint16) int16 {
    p0 = (p0 & p0)
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint16 = 464
        switch "a" {
        case "a":
            p0 -= 32767
        case "foo bar":
            var v8 float32 = 3.14159
            _ = v8
        case "":
            v9 := GenS1{F0: 77, F1: 1}
            _ = v9
        default:
            v10 := []int{9223372036854775807, fi6, (fi6 >> ((18446744073709551615) & 63))}
            _ = v10
        }
        var v11 int16 = int16(fi6)
        _ = fi6
        _ = v11
        _ = v7
    }
    if (!("\n\t" <= "foo bar")) {
        fmt.Println(p0)
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(fi12, p0)
            _ = fi12
        }
    } else {
        fmt.Println(p0)
        switch -749 {
        case 347:
            if true {
                p0 = p0
                v13 := []int{9223372036854775807, 646, 2147483647}
                _ = v13
            } else {
                if false {
                    var v14 uint16 = p0
                    if (!true) {
                        var v15 float32 = 0.0
                        if (!(-594 != int(p0))) {
                            fmt.Println(v14, v15)
                        } else {
                            v16 := []int{32767}
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            _ = v16
                        }
                        v17 := GenS2{F0: "foo bar", F1: 2}
                        _ = v15
                        _ = v17
                    }
                    _ = v14
                }
                for fi18 := 0; fi18 < 6; fi18++ {
                    fmt.Println(fi18, p0)
                    p0 -= (p0 & p0)
                    var v19 uint8 = 114
                    _ = fi18
                    _ = v19
                }
            }
        }
    }
    var v20 int64 = 770
    if true {
        var v21 int8 = 0
        _ = v21
    } else {
        var v22 int32 = 42
        fmt.Println(v20, v22)
        var v23 float32 = 1e10
        _ = v22
        _ = v23
    }
    _ = p0
    _ = v20
    return -1
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"a"}
    var v3 int64 = 0
    v4 := []uint32{}
    var v5 string = (("\n\t" + "a") + "foo bar")
    if len(v4) > 0 {
        v4[0] = 42
    }
    v3 = (v3 >> ((0) & 63))
    if (true || (((1e10 <= 0.1) && (255 <= 4294967295)) && false)) {
        var v24 uint = 18446744073709551615
        switch ((((v5 + "foo bar") + ("a" + "hello")) + v5) + ((v5 + v5) + v5)) {
        case "hello":
            switch (v5 + ("" + "")) {
            case "a":
                v2 = append(v2, (v5 + ((v5 + v5) + "")))
            case "":
                v25 := GenS1{F0: v24, F1: v24}
                _ = v25
            default:
                if ((-816 > 42) && ((981 > genFn0(993)) || false)) {
                    v5 = v5
                    v5 = v5
                } else {
                    v26 := GenS2{F0: v5, F1: 42}
                    _ = v26
                }
            }
        case "":
            var v27 uint32 = 2
            _ = v27
        case "\n\t":
            _ = len(v1)
        default:
            var v28 uint16 = 387
            _ = v28
        }
        _ = v24
    } else {
        v5 = (v5 + v5)
        _ = len(v1)
        var v29 uint8 = 42
        _ = v29
    }
    var v30 uint32 = 2
    switch -915 {
    case 9223372036854775807:
        fmt.Println(v5, v3, v30)
    default:
        v30 = (^(v30))
    }
    var v31 int32 = 920
    if (42 < 42) {
        if (2 >= 391) {
            var v32 string = v5
            var v33 int64 = (v3 | v3)
            _ = v32
            _ = v33
        } else {
            v34 := GenS3{F0: 1, F1: (v30 << ((4294967295) & 63)), F2: 18446744073709551615}
            _ = v34
        }
        fmt.Println(v30, v31, v3)
    }
    var v35 int32 = (^(v31))
    if false {
        if (!true) {
            var v36 uint32 = 1
            switch 2147483647 {
            case 2147483647:
                _ = len(v1)
            default:
                if (((4294967295 < 42) || true) && false) {
                    v31 += (^(v35))
                    fmt.Println(v30, v35, v36)
                } else {
                    var v37 uint16 = 0
                    _ = v37
                }
            }
            _ = v36
        }
        var v38 float64 = 0.1
        for fi39 := 0; fi39 < 15; fi39++ {
            var v40 uint8 = 255
            var v41 float32 = (0.1 - 3.14159)
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v5
}
