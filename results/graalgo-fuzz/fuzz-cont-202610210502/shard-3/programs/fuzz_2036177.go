package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int16; F1 bool; F2 int8 }
func genFn0() uint64 {
    v7 := GenS1{F0: 4294967295, F1: 1, F2: 255}
    if (-9223372036854775808 < -912) {
        v8 := GenS2{F0: -2147483648, F1: 1}
        v9 := []int{2147483647}
        _ = v8
        _ = v9
    } else {
        if false {
            var v10 uint = 255
            v10 = v10
            if false {
                v11 := GenS3{F0: (v10 ^ 255)}
                v12 := []int{876, -312, 32768, 809}
                _ = v11
                _ = v12
            } else {
                var v13 int = 0
                _ = v13
            }
            _ = v10
        } else {
            if true {
                var v14 uint64 = 4294967295
                _ = v14
            }
            v15 := []int{255, -838, 2147483647, 0}
            var v16 int32 = 59
            _ = v15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        if (42 == -1) {
            if false {
                v18 := []int{fi17, fi17, fi17, (fi17 << ((18446744073709551615) & 63))}
                _ = v18
            } else {
                var v19 int32 = 0
                var v20 int16 = 347
                fmt.Println(v20, v20)
                _ = v19
                _ = v20
            }
        }
        _ = fi17
    }
    _ = v7
    return 42
}
func genFn1(p0 uint8) uint8 {
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 int16 = 1
        v23 := GenS4{F0: 100}
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := []int{}
    fmt.Println(p0)
    v25 := GenS4{F0: -128}
    _ = p0
    _ = v24
    _ = v25
    return (p0 << ((uint(p0)) & 63))
}

func main() {
    var v1 int = -234
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    v6 := []string{("hello" + ("hello" + "\n\t")), ""}
    v26 := v5[v1]
    var v27 uint32 = v26
    delete(v4, 1)
    v27 = v27
    fmt.Println(v27, v27, v1)
    switch "a" {
    case "foo bar":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "":
        if false {
            switch v1 {
            case 32768:
                v28 := GenS2{F0: int32(v1), F1: v1}
                _ = v28
            default:
                for fi29 := 0; fi29 < 20; fi29++ {
                    v1 = v1
                    switch (v1 % (((v1 & fi29)) | 1)) {
                    case 128:
                        for fi30 := 0; fi30 < 7; fi30++ {
                            v31 := GenS5{F0: 1, F1: 1, F2: v26}
                            v26 = v27
                            fmt.Println(v27)
                            v3[(v27 - (^(v26)))] = fi30
                            fmt.Println(v27, v26)
                            _ = fi30
                            _ = v31
                        }
                    case 2147483647:
                        for fi32 := 0; fi32 < 5; fi32++ {
                            v5[-2147483648] = v26
                            _ = fi32
                        }
                    default:
                        v33 := GenS3{F0: 1}
                        _ = v33
                    }
                    if (2 >= (^(v27))) {
                        v34 := GenS6{F0: -1, F1: (!false), F2: -128}
                        var v35 string = ""
                        switch ((v35 + "foo bar") + v35) {
                        case "":
                            if ((v35 + (v35 + v35)) != v35) {
                                _ = len(v2)
                                fmt.Println(v1, v26)
                            } else {
                                v26 = (v27 >> ((2) & 63))
                            }
                        case "a":
                            v6 = append(v6, v35)
                        case "foo bar":
                            fmt.Println(v26, v1)
                        default:
                            v26 = (v27 | v27)
                        }
                        _ = v34
                        _ = v35
                    }
                    _ = fi29
                }
            }
            switch v1 {
            case 525:
                var v36 uint64 = 18446744073709551615
                _ = v36
            default:
                var v37 int32 = 394
                _ = v37
            }
        }
    case "a":
        switch "" {
        case "\n\t":
            switch "" {
            case "hello":
                var v38 int64 = -1
                _ = v38
            case "a":
                var v39 int = (^(v1))
                _ = v39
            }
        case "hello":
            v27 += (^(v26))
        case "foo bar":
            v1 = (v1 | (-(v1)))
        }
    }
    if (1 > 213) {
        var v40 uint32 = 869
        switch "foo bar" {
        case "foo bar":
            v5[(v1 ^ v1)] = (v26 & v40)
        default:
            fmt.Println(v27, v40, v1)
        }
        switch (v1 ^ v1) {
        case -2147483648:
            delete(v5, -1)
        default:
            v41 := v5[9223372036854775807]
            _ = v41
        }
        _ = v40
    }
    var v42 uint8 = 127
    v42 = (v42 - 0)
    if (1e10 <= 1.5) {
        if false {
            fmt.Println(v1)
        }
        var v43 bool = ((genFn1(1) < (v42 << ((42) & 63))) || false)
        _ = v43
    }
    var v44 uint64 = 0
    v45 := GenS1{F0: 255, F1: (v27 >> ((1) & 63)), F2: (v42 + 128)}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi46 := 0; fi46 < 5; fi46++ {
        switch "" {
        case "":
            v1 = 4294967295
        default:
            _ = len(v3)
        }
        _ = fi46
    }
    if (1 > 1) {
        var v47 int16 = 1
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
