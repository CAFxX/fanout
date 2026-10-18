package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 string; F1 int64; F2 uint16 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint; F1 uint8; F2 int32 }
type GenS6 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int64, p1 int16, p2 bool) int16 {
    if p2 {
        v7 := []int{164}
        fmt.Println(p1, p0, p1)
        _ = v7
    }
    v8 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 32767
}
func genFn1(p0 uint64, p1 uint, p2 uint16) int32 {
    p0 -= uint64(p2)
    v9 := GenS1{F0: (p1 / ((p1) | 1)), F1: -342}
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := []int{916, fi10, 0}
        p0 += 2
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 703, F1: 2147483647}
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v9
    return 1
}

func main() {
    var v1 float64 = (-2.25 / (float64(((0.0 + (float32(((-2.25 + (-2.25 + (3.14159 - (0.0 - (((0.0 - float64(float32(3.14159))) / 1.5) / 1.0))))) / (-2.25 - (1.5 + 3.14159)))) / 1.0)) + 1e10)) - 1e10))
    v2 := []uint32{1, 1, 4294967295, 2}
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    var v5 int = 959
    var v6 int32 = 2147483647
    for fi13 := 0; fi13 < 9; fi13++ {
        v1 -= 0.1
        _ = fi13
    }
    delete(v4, 0)
    fmt.Println(v5)
    var v14 float32 = (3.14159 / ((1.5 * -2.25) - -2.25))
    if true {
        fmt.Println(v1, v1, v1)
        if false {
            var v15 int64 = -1
            for fi16 := 0; fi16 < 1; fi16++ {
                v2 = append(v2, 7)
                var v17 int16 = 1
                _ = fi16
                _ = v17
            }
            _ = v15
        } else {
            for fi18 := 0; fi18 < 5; fi18++ {
                fmt.Println(v14)
                _ = fi18
            }
        }
    } else {
        if true {
            if true {
                for fi19 := 0; fi19 < 4; fi19++ {
                    if true {
                        v20 := GenS1{F0: 42, F1: fi19}
                        v21 := v3[42]
                        if false {
                            fmt.Println(v21, v21)
                            var v22 uint = 378
                            var v23 string = "a"
                            _ = v22
                            _ = v23
                        }
                        _ = v20
                        _ = v21
                    }
                    var v24 float32 = v14
                    _ = fi19
                    _ = v24
                }
                for fi25 := 0; fi25 < 11; fi25++ {
                    if false {
                        v5 += int(v6)
                        var v26 int32 = 0
                        var v27 uint = 4294967295
                        _ = v26
                        _ = v27
                    }
                    _ = fi25
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                var v28 float64 = v1
                _ = v28
            }
            _ = len(v4)
        }
        for fi29 := 0; fi29 < 4; fi29++ {
            delete(v4, fi29)
            var v30 string = "hello"
            _ = fi29
            _ = v30
        }
    }
    v6 = -770
    v14 = float32(-2.25)
    var v31 float64 = 1.5
    if true {
        var v32 int32 = (v6 + (v6 << ((18446744073709551615) & 63)))
        var v33 bool = ((2 <= 2) || true)
        var v34 int8 = 100
        _ = v32
        _ = v33
        _ = v34
    }
    v35 := GenS1{F0: 4294967295, F1: (v5 - (v5 - -1))}
    for fi36 := 0; fi36 < 18; fi36++ {
        v37 := GenS2{F0: "", F1: 449, F2: 255}
        if (4294967295 == 0) {
            if (1 >= 65535) {
                fmt.Println(v5, v6, v5)
            }
        } else {
            for fi38 := 0; fi38 < 16; fi38++ {
                var v39 uint16 = 32767
                v31 = v1
                _ = len(v2)
                _ = fi38
                _ = v39
            }
            v40 := GenS3{F0: ("foo bar" + "hello"), F1: "a"}
            _ = v40
        }
        _ = fi36
        _ = v37
    }
    var v41 uint32 = 2147483647
    for fi42 := 0; fi42 < 11; fi42++ {
        switch "" {
        case "":
            fmt.Println(v31, v5)
        case "\n\t":
            switch (fi42 | (v5 ^ (^(fi42)))) {
            case 760:
                var v43 int32 = (v6 | 2147483647)
                _ = v43
            case 4294967295:
                switch ("" + "\n\t") {
                case "hello":
                    v44 := GenS4{F0: 1, F1: 0}
                    _ = v44
                case "foo bar":
                    v45 := GenS5{F0: 470, F1: 42, F2: 1}
                    _ = v45
                case "a":
                    v46 := GenS6{F0: 46, F1: 0}
                    _ = v46
                }
            case 780:
                v6 = (v6 | v6)
            default:
                var v47 uint64 = 42
                _ = v47
            }
        }
        fmt.Println(v14, v14, fi42)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi42
    }
    v3[(v41 / (((v41 - v41)) | 1))] = (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
