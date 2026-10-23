package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 int32 }
func genFn0(p0 int) int32 {
    var v6 int = 32767
    v7 := GenS1{F0: 836, F1: p0, F2: (-616 < 42)}
    var v8 string = "foo bar"
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint16 = 1
    var v2 int8 = 0
    v3 := make(map[int]bool)
    var v4 uint = 42
    var v5 string = ""
    v9 := []int{-1}
    v2 = v2
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint64 = 0
        v3[fi10] = true
        if false {
            fmt.Println(fi10)
        } else {
            for fi12 := 0; fi12 < 14; fi12++ {
                v2 = (v2 | int8(fi12))
                v13 := v3[1]
                fmt.Println(v11, v1, fi12)
                _ = fi12
                _ = v13
            }
            var v14 uint8 = 127
            for fi15 := 0; fi15 < 12; fi15++ {
                v3[208] = (-1 >= 9223372036854775807)
                v1 = v1
                v16 := GenS2{F0: v4}
                v5 = v5
                var v17 string = ""
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    if true {
        for fi18 := 0; fi18 < 4; fi18++ {
            switch v5 {
            case "hello":
                var v19 uint16 = (v1 * 255)
                _ = v19
            case "foo bar":
                fmt.Println(fi18, v2)
            case "":
                fmt.Println(v2, v2)
            default:
                v20 := GenS2{F0: (v4 / ((v4) | 1))}
                _ = v20
            }
            v5 = v5
            _ = fi18
        }
    } else {
        if ((true || true) && true) {
            var v21 float64 = -2.25
            v22 := GenS2{F0: v4}
            _ = v21
            _ = v22
        } else {
            for fi23 := 0; fi23 < 13; fi23++ {
                for fi24 := 0; fi24 < 16; fi24++ {
                    switch (fi23 * (fi24 - fi24)) {
                    case -86:
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    case 4294967295:
                        v2 = v2
                    default:
                        _ = len(v3)
                        v2 = v2
                        switch v5 {
                        case "":
                            v25 := GenS3{F0: (^(v1)), F1: (fi23 << ((65535) & 63))}
                            _ = v25
                        case "a":
                            var v26 bool = (42 > 42)
                            _ = v26
                        }
                        switch v5 {
                        case "\n\t":
                            v1 = (v1 * v1)
                            switch v5 {
                            case "a":
                                v1 -= 32767
                            }
                        case "foo bar":
                            v2 -= v2
                        default:
                            var v27 int = fi23
                            v28 := GenS3{F0: v1, F1: 2147483647}
                            switch (^(fi24)) {
                            case 32767:
                                fmt.Println(v5, v27, v4)
                            case -1:
                                fmt.Println(v1, fi24)
                            default:
                                var v29 int32 = -123
                                _ = v29
                            }
                            v2 = 87
                            var v30 int8 = (v2 ^ (v2 % (((^(v2))) | 1)))
                            _ = v27
                            _ = v28
                            _ = v30
                        }
                        switch fi23 {
                        case 929:
                            v31 := GenS4{F0: 127}
                            _ = v31
                        case 32768:
                            var v32 int64 = 42
                            _ = v32
                        default:
                            v2 += (^(v2))
                        }
                    }
                    _ = fi24
                }
                v33 := GenS2{F0: (v4 - (v4 << (((^(v4))) & 63)))}
                _ = fi23
                _ = v33
            }
            var v34 int64 = 1
            _ = v34
        }
    }
    fmt.Println(v1, v5)
    var v35 uint = v4
    var v36 int16 = 99
    var v37 int = 128
    if ((!true) && false) {
        for fi38 := 0; fi38 < 5; fi38++ {
            var v39 int8 = (v2 / ((v2) | 1))
            v39 -= v39
            _ = fi38
            _ = v39
        }
    } else {
        var v40 float32 = float32((0.0 / 1e10))
        v37 -= (-(v37))
        var v41 int32 = 42
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        for fi43 := 0; fi43 < 16; fi43++ {
            if len(v9) > 0 {
                v9[0] = -9223372036854775808
            }
            v36 += v36
            v44 := GenS4{F0: 1}
            _ = fi43
            _ = v44
        }
        v45 := GenS2{F0: v35}
        switch 687 {
        case 128:
            fmt.Println(v2)
        case -46:
            if (1 > 1) {
                v46 := GenS2{F0: v4}
                var v47 int = fi42
                _ = v46
                _ = v47
            }
        case 32767:
            v48 := GenS5{F0: 1, F1: (v35 ^ 295), F2: 42}
            _ = v48
        default:
            switch "\n\t" {
            case "foo bar":
                for fi49 := 0; fi49 < 1; fi49++ {
                    delete(v3, (fi42 | fi42))
                    _ = fi49
                }
            case "\n\t":
                v35 = (v4 + v4)
            case "hello":
                var v50 uint32 = 858
                _ = v50
            default:
                if (("" + ("hello" + v5)) <= v5) {
                    v4 -= (^(v35))
                    if ((v36 >> (((^(v4))) & 63)) <= v36) {
                        v9 = append(v9, (fi42 | (v37 << ((197) & 63))))
                        if false {
                            v51 := GenS4{F0: 2147483647}
                            v52 := GenS6{F0: -1}
                            _ = v51
                            _ = v52
                        } else {
                            var v53 uint64 = 431
                            switch v5 {
                            case "a":
                                v37 = fi42
                            case "hello":
                                for fi54 := 0; fi54 < 15; fi54++ {
                                    if (true && false) {
                                        if (-966 < -1) {
                                            var v55 int32 = 1
                                            if len(v9) > 0 {
                                                v9[0] = (-(fi54))
                                            }
                                            _ = v55
                                        }
                                    }
                                    if ((1.5 * -2.25) < 0.1) {
                                        var v56 uint = (v35 + (v4 - (^(v35))))
                                        v57 := v3[(fi54 | v37)]
                                        _ = v56
                                        _ = v57
                                    }
                                    _ = fi54
                                }
                            default:
                                v58 := GenS7{F0: 1000, F1: -2147483648}
                                _ = v58
                            }
                            var v59 bool = (2147483647 != 7)
                            _ = v53
                            _ = v59
                        }
                    }
                }
            }
        }
        _ = fi42
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v9
}
