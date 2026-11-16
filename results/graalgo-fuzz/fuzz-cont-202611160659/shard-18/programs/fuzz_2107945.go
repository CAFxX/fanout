package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 string; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint8 }
func genFn0() int {
    var v6 int16 = 0
    for fi7 := 0; fi7 < 15; fi7++ {
        v6 -= (v6 * (v6 & 1))
        _ = fi7
    }
    var v8 float64 = 0.1
    v8 = ((v8 + v8) / v8)
    _ = v6
    _ = v8
    return 32768
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    v4 := make(map[int]string)
    var v5 int64 = 1
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := []int{fi9, -2147483648, fi9, genFn0()}
        if (127 > (fi9 % ((fi9) | 1))) {
            fmt.Println(fi9, v5)
        } else {
            var v11 string = (("" + "a") + "a")
            _ = v11
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 uint32 = 2147483647
            if true {
                v14 := GenS1{F0: 2}
                var v15 uint32 = (v13 & v13)
                _ = v14
                _ = v15
            } else {
                v16 := GenS2{F0: (fi9 & fi12), F1: (("hello" + "foo bar") + "foo bar"), F2: (v13 - (v13 >> ((15) & 63)))}
                switch fi12 {
                case 32767:
                    if (!true) {
                        fmt.Println(v5, fi12)
                    } else {
                        v17 := GenS3{F0: 42, F1: 173}
                        _ = v17
                    }
                case 42:
                    var v18 int = (-(fi12))
                    _ = v18
                }
                var v19 int16 = -1
                var v20 int = fi12
                v13 += (v13 >> ((2) & 63))
                v13 -= v13
                _ = v16
                _ = v19
                _ = v20
            }
            fmt.Println(v5, v13, v13)
            _ = fi12
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    v21 := GenS2{F0: genFn0(), F1: "foo bar", F2: 0}
    var v22 uint16 = 255
    var v23 bool = false
    v24 := []int{}
    if (v23 || (4294967295 != 0)) {
        switch -1 {
        case 127:
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 uint64 = 953
                if v23 {
                    v26 = v26
                }
                v2[872] = v23
                _ = fi25
                _ = v26
            }
        case 670:
            v23 = true
        case 42:
            var v27 uint8 = 141
            _ = v27
        default:
            v22 += 0
        }
        v2[680] = (255 == 2)
        if (127 != -1) {
            v24 = append(v24, 64)
        } else {
            var v28 uint16 = v22
            v28 = v22
            var v29 uint32 = 2147483647
            _ = v28
            _ = v29
        }
    }
    v22 = (v22 ^ (v22 % (((v22 >> ((1) & 63))) | 1)))
    if (!(2 == 42)) {
        _ = len(v24)
    } else {
        if len(v24) > 0 {
            v24[0] = 282
        }
        _ = len(v1)
    }
    v30 := GenS3{F0: 38, F1: 0}
    v22 -= 1
    var v31 uint64 = 110
    if ("" > ((("" + ("" + "hello")) + "a") + "\n\t")) {
        var v32 uint16 = v22
        if len(v24) > 0 {
            v24[0] = 32768
        }
        switch 0 {
        case -850:
            v24 = append(v24, genFn0())
        case 128:
            fmt.Println(v23, v32, v5)
        case 1:
            if (v23 || true) {
                for fi33 := 0; fi33 < 11; fi33++ {
                    for fi34 := 0; fi34 < 7; fi34++ {
                        var v35 string = ("\n\t" + ("foo bar" + "a"))
                        var v36 string = "foo bar"
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    if (64 == 128) {
                        v31 = v31
                    } else {
                        v37 := GenS3{F0: 42, F1: 255}
                        _ = v37
                    }
                    _ = fi33
                }
            } else {
                v22 += (v22 | v32)
                if (1000 >= 42) {
                    fmt.Println(v23, v23)
                    v5 = v5
                    for fi38 := 0; fi38 < 1; fi38++ {
                        var v39 bool = true
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                        _ = fi38
                        _ = v39
                    }
                } else {
                    v31 += 2
                }
            }
        default:
            if v23 {
                v40 := v3[497]
                _ = v40
            } else {
                v22 -= v22
                _ = len(v1)
                _ = len(v24)
            }
        }
        _ = v32
    } else {
        if (false || v23) {
            if (!v23) {
                switch "" {
                case "\n\t":
                    v31 = 1
                }
                switch genFn0() {
                case 0:
                    v31 = v31
                case 127:
                    var v41 int = -351
                    _ = v41
                }
                var v42 int32 = 0
                _ = v42
            }
            var v43 int = 2147483647
            var v44 uint16 = (v22 - v22)
            _ = v43
            _ = v44
        }
        var v45 int8 = 19
        _ = v45
    }
    if (100 >= -128) {
        v31 += (v31 + (^(v31)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
