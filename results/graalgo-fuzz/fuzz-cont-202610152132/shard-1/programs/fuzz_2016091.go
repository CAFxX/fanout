package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS4 struct { F0 int; F1 uint64; F2 int32 }
func genFn0() uint32 {
    v6 := []int{-1}
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint32 = 2147483648
        var v9 uint16 = 765
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v6
    return 1
}

func main() {
    var v1 uint64 = 752
    v2 := make(map[int]bool)
    var v3 int64 = 0
    v4 := []int{9223372036854775807, -105, 42}
    v5 := make(map[string]int)
    for fi10 := 0; fi10 < 11; fi10++ {
        _ = len(v2)
        v11 := GenS1{F0: true, F1: 42, F2: 192}
        v1 += (v1 >> ((255) & 63))
        _ = fi10
        _ = v11
    }
    v4 = append(v4, -893)
    if (-75 == -1) {
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint32 = 2147483647
            for fi14 := 0; fi14 < 18; fi14++ {
                v1 = (v1 ^ (^(v1)))
                var v15 int8 = 99
                v15 = v15
                var v16 int = 4294967295
                if true {
                    v17 := GenS2{F0: (v15 | v15), F1: 255, F2: (v1 << ((4294967295) & 63))}
                    v18 := GenS2{F0: 1, F1: 179, F2: v1}
                    _ = v17
                    _ = v18
                } else {
                    if ((0 > (v1 | 1)) || (v13 < genFn0())) {
                        v19 := GenS3{F0: 4294967295, F1: (^(v15)), F2: v13}
                        v20 := GenS3{F0: v1, F1: (v15 ^ v15), F2: v13}
                        _ = v19
                        _ = v20
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
        for fi21 := 0; fi21 < 18; fi21++ {
            v3 += (^(v3))
            switch fi21 {
            case -9223372036854775808:
                v22 := GenS2{F0: 47, F1: 0, F2: 42}
                _ = v22
            case 32768:
                if true {
                    var v23 int64 = v3
                    fmt.Println(v23, v23, v23)
                    switch (^(fi21)) {
                    case 32767:
                        for fi24 := 0; fi24 < 1; fi24++ {
                            var v25 string = "\n\t"
                            var v26 int32 = 153
                            v3 += v23
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    case 2147483647:
                        var v27 int16 = 1
                        _ = v27
                    case -880:
                        var v28 uint8 = 128
                        _ = v28
                    }
                    _ = v23
                }
            case -2147483648:
                v29 := v5[""]
                _ = v29
            default:
                _ = len(v4)
            }
            fmt.Println(v1)
            _ = fi21
        }
        fmt.Println(v3)
    }
    v3 -= (v3 >> ((2) & 63))
    switch -2147483648 {
    case 1:
        v1 = 42
    case -846:
        for fi30 := 0; fi30 < 4; fi30++ {
            if false {
                if len(v4) > 0 {
                    v4[0] = 1
                }
                var v31 float64 = 0.0
                var v32 uint64 = (v1 << ((60) & 63))
                _ = v31
                _ = v32
            } else {
                for fi33 := 0; fi33 < 10; fi33++ {
                    v1 -= v1
                    _ = fi33
                }
                var v34 uint64 = v1
                var v35 uint32 = 1
                _ = v34
                _ = v35
            }
            v3 = -1
            switch "" {
            case "\n\t":
                v36 := GenS3{F0: v1, F1: 12, F2: 4294967295}
                _ = v36
            case "hello":
                for fi37 := 0; fi37 < 11; fi37++ {
                    var v38 bool = (42 > -741)
                    var v39 int16 = 0
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
            _ = fi30
        }
    case 4294967295:
        for fi40 := 0; fi40 < 12; fi40++ {
            var v41 uint8 = 128
            var v42 float32 = (1e10 + ((1e10 * ((-2.25 + 3.14159) + -2.25)) - (1.5 - float32(0.0))))
            v42 = float32(1e10)
            _ = fi40
            _ = v41
            _ = v42
        }
    default:
        for fi43 := 0; fi43 < 16; fi43++ {
            v44 := GenS2{F0: -1, F1: 29, F2: (v1 << ((170) & 63))}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi43
            _ = v44
        }
    }
    fmt.Println(v1, v3)
    switch 127 {
    case 128:
        var v45 float64 = 1e10
        _ = v45
    default:
        for fi46 := 0; fi46 < 11; fi46++ {
            v47 := GenS4{F0: (fi46 << ((0) & 63)), F1: (v1 % ((v1) | 1)), F2: -1}
            _ = fi46
            _ = v47
        }
    }
    v3 = (v3 >> ((2) & 63))
    var v48 uint64 = v1
    v49 := GenS3{F0: (v1 - (v48 >> ((255) & 63))), F1: 48, F2: genFn0()}
    v3 += (v3 * (v3 >> ((132) & 63)))
    _ = len(v2)
    if (1 <= (v1 / ((v48) | 1))) {
        if (-1 != 0) {
            v1 = v48
            if (42 >= 373) {
                v48 -= (v1 - (v1 & 1))
            } else {
                var v50 string = ""
                _ = v50
            }
            fmt.Println(v48)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
}
