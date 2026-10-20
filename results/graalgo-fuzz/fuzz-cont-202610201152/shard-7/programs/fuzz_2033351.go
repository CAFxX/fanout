package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint; F1 uint; F2 int8 }
func genFn0(p0 int16, p1 int8, p2 int64) uint16 {
    var v4 uint16 = 1
    fmt.Println(p1, p0)
    if true {
        switch (("" + "a") + "hello") {
        case "\n\t":
            var v5 uint64 = 4294967295
            _ = v5
        default:
            if (3.14159 != (1e10 - float32(((-2.25 - -2.25) - 3.14159)))) {
                p1 = p1
                if (322 == 0) {
                    if (2147483648 > 2147483647) {
                        p1 -= -1
                    } else {
                        var v6 bool = false
                        _ = v6
                    }
                } else {
                    v7 := []int{-633, 2147483647, 388}
                    var v8 uint32 = 2147483647
                    var v9 string = "hello"
                    _ = v7
                    _ = v8
                    _ = v9
                }
            } else {
                v10 := []int{32767, 4294967295}
                p0 = 1
                _ = v10
            }
        }
        var v11 uint64 = 861
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint16 = v4
        if true {
            p2 -= 9223372036854775807
        }
        var v14 uint64 = 1
        _ = fi12
        _ = v13
        _ = v14
    }
    p0 = (p0 & (p0 - p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (v4 << ((255) & 63))
}
func genFn1() uint32 {
    var v15 uint = 778
    v15 = v15
    _ = v15
    return 2
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{225, 186, 0}
    var v3 bool = false
    for fi16 := 0; fi16 < 9; fi16++ {
        if v3 {
            if v3 {
                var v17 int32 = 0
                if (2 >= 1) {
                    var v18 uint8 = 2
                    _ = v18
                } else {
                    if true {
                        v3 = (1 == 4294967295)
                        if (!v3) {
                            fmt.Println(v3, v3, fi16)
                        }
                        v19 := GenS1{F0: 2147483648}
                        _ = v19
                    }
                    if (4294967295 >= 255) {
                        var v20 float64 = 0.0
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 11; fi21++ {
                            v22 := GenS2{F0: 1, F1: 42, F2: fi21}
                            v17 = (^(v17))
                            v3 = true
                            _ = fi21
                            _ = v22
                        }
                        if v3 {
                            for fi23 := 0; fi23 < 16; fi23++ {
                                var v24 int32 = 0
                                _ = fi23
                                _ = v24
                            }
                            for fi25 := 0; fi25 < 4; fi25++ {
                                switch (fi16 | (fi25 >> ((255) & 63))) {
                                case 42:
                                    var v26 float64 = 0.1
                                    _ = v26
                                }
                                v2 = append(v2, 65535)
                                _ = fi25
                            }
                            delete(v1, "")
                        }
                    }
                }
                _ = v17
            } else {
                var v27 uint8 = 0
                var v28 float64 = ((3.14159 + 1e10) - 1e10)
                v3 = (v3 || (v3 || (true || false)))
                _ = v27
                _ = v28
            }
        }
        v3 = v3
        _ = fi16
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 int64 = -75
        var v31 float64 = ((1.5 + 1.5) * ((1.5 * 3.14159) * (0.1 + (1e10 + (1.5 + 1e10)))))
        _ = fi29
        _ = v30
        _ = v31
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        v3 = v3
        v33 := GenS3{F0: 42}
        v34 := GenS4{F0: 2, F1: -128, F2: 0}
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        var v36 int64 = 1
        v37 := GenS5{F0: 1, F1: 4294967295, F2: 127}
        v3 = v3
        _ = fi35
        _ = v36
        _ = v37
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v38 float64 = -2.25
    var v39 string = "hello"
    switch 32767 {
    case -1:
        v40 := GenS4{F0: 553, F1: -128, F2: 1}
        _ = v40
    case 32768:
        v38 += 3.14159
    case 0:
        v3 = v3
    }
    fmt.Println(v38, v38, v38)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
}
