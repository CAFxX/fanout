package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int }
func genFn0() uint8 {
    switch "hello" {
    case "foo bar":
        switch "" {
        case "\n\t":
            var v6 int16 = -32768
            _ = v6
        }
    case "\n\t":
        var v7 uint = 1
        _ = v7
    default:
        v8 := []int{32767, -679, 0}
        _ = v8
    }
    switch -71 {
    case -1:
        var v9 uint64 = 4294967295
        _ = v9
    case 128:
        var v10 int8 = 53
        _ = v10
    case 761:
        var v11 uint = 42
        _ = v11
    }
    var v12 bool = (1 != 65535)
    v13 := GenS1{F0: 42, F1: 1, F2: -359}
    _ = v12
    _ = v13
    return 100
}
func genFn1(p0 bool) int32 {
    p0 = (!p0)
    fmt.Println(p0)
    var v14 uint32 = 2147483648
    if (p0 || true) {
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := []int{}
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: -1}
        _ = v17
    } else {
        v18 := GenS1{F0: 2, F1: 85, F2: 42}
        v14 = 4294967295
        _ = v18
    }
    _ = p0
    _ = v14
    return 0
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]int)
    var v3 int = 4294967295
    var v4 int = 42
    var v5 float64 = ((3.14159 * ((1e10 - 3.14159) + (((0.0 + 1.5) * 0.0) + 0.0))) * float64(3.14159))
    fmt.Println(v4, v1, v4)
    v19 := []int{v4, (v3 & 4294967295), v3, v3}
    switch "" {
    case "hello":
        var v20 uint16 = 255
        _ = v20
    default:
        v1 -= v1
    }
    fmt.Println(v4, v3)
    if true {
        var v21 int64 = 0
        for fi22 := 0; fi22 < 5; fi22++ {
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 int32 = genFn1(false)
                if true {
                    switch fi23 {
                    case -674:
                        v25 := v2["a"]
                        _ = v25
                    case 32768:
                        switch -1 {
                        case -607:
                            if (1000 >= 42) {
                                switch 2147483647 {
                                case 1:
                                    if (!(true || true)) {
                                        switch v3 {
                                        case 32768:
                                            if (!true) {
                                                if (((1e10 / v5) / ((0.0 * v5) * v5)) > v5) {
                                                    if ((false || (((-2.25 + (v1 / v1)) * float32(1.5)) <= 0.0)) || (((1.5 * v5) + v5) > (((3.14159 + v5) / v5) - 1e10))) {
                                                        v26 := GenS3{F0: -363, F1: (v3 + 418)}
                                                        _ = v26
                                                    } else {
                                                        v4 += v3
                                                    }
                                                    var v27 uint32 = 942
                                                    v3 = 42
                                                    fmt.Println(v1, v24, v21)
                                                    _ = v27
                                                }
                                                v1 -= v1
                                            } else {
                                                var v28 int16 = 583
                                                var v29 uint = 36
                                                var v30 int64 = -1
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            }
                                        case 0:
                                            switch ("a" + "foo bar") {
                                            case "a":
                                                fmt.Println(v1, v24)
                                            case "\n\t":
                                                var v31 uint32 = 306
                                                _ = v31
                                            case "hello":
                                                v3 = (fi22 - fi22)
                                            default:
                                                _ = len(v2)
                                            }
                                        }
                                    }
                                case 0:
                                    var v32 uint64 = 1
                                    _ = v32
                                }
                                var v33 int64 = v21
                                var v34 string = "foo bar"
                                switch (((v34 + ("\n\t" + "foo bar")) + v34) + ((v34 + v34) + (v34 + "hello"))) {
                                case "hello":
                                    fmt.Println(v33, v33)
                                case "":
                                    fmt.Println(v4)
                                default:
                                    var v35 int64 = v21
                                    _ = v35
                                }
                                _ = v33
                                _ = v34
                            } else {
                                if ((v21 << ((4294967295) & 63)) == v21) {
                                    fmt.Println(fi23, v21, v4)
                                }
                                v1 = float32(v5)
                                fmt.Println(v5, v4)
                                var v36 int = (fi22 | v4)
                                v2["hello"] = -188
                                _ = v36
                            }
                        case 42:
                            v5 = v5
                            fmt.Println(v24)
                            var v37 uint = 1
                            _ = v37
                        case 1:
                            v5 = v5
                        default:
                            v24 = v24
                        }
                    case 127:
                        if false {
                            fmt.Println(v5, v24, v4)
                        } else {
                            var v38 int8 = 127
                            _ = v38
                        }
                    default:
                        v3 = (fi23 - 2147483647)
                    }
                } else {
                    var v39 uint32 = 829
                    _ = v39
                }
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        _ = v21
    }
    fmt.Println(v3)
    v5 = float64(float32((1.5 + (v5 - v5))))
    for fi40 := 0; fi40 < 13; fi40++ {
        if false {
            var v41 int = fi40
            _ = v41
        }
        v3 = v3
        _ = fi40
    }
    v5 -= v5
    var v42 uint16 = 42
    switch "a" {
    case "":
        _ = len(v2)
    case "a":
        for fi43 := 0; fi43 < 13; fi43++ {
            fmt.Println(v4, v5)
            for fi44 := 0; fi44 < 6; fi44++ {
                var v45 float64 = (v5 / v5)
                _ = fi44
                _ = v45
            }
            _ = fi43
        }
    default:
        var v46 uint8 = 2
        _ = v46
    }
    for fi47 := 0; fi47 < 4; fi47++ {
        v48 := GenS3{F0: 0, F1: v4}
        if (float64(v1) != v5) {
            v5 += 0.0
            switch (v3 << ((4294967295) & 63)) {
            case 4294967295:
                if true {
                    if (!true) {
                        var v49 float32 = 0.1
                        _ = v49
                    }
                }
            case 128:
                v19 = append(v19, v3)
            case -252:
                var v50 uint32 = 425
                _ = v50
            default:
                var v51 string = "a"
                _ = v51
            }
        } else {
            v5 = float64(3.14159)
        }
        _ = fi47
        _ = v48
    }
    var v52 uint8 = 255
    for fi53 := 0; fi53 < 19; fi53++ {
        for fi54 := 0; fi54 < 16; fi54++ {
            _ = len(v19)
            switch "a" {
            case "\n\t":
                v42 = 32767
            default:
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v1)
                }
                var v55 int8 = int8(fi53)
                var v56 string = "hello"
                _ = v55
                _ = v56
            }
            v57 := GenS1{F0: (^(v52)), F1: (v52 | v52), F2: 1}
            _ = fi54
            _ = v57
        }
        switch (fi53 & 0) {
        case -299:
            fmt.Println(v3, v1, v4)
        }
        _ = fi53
    }
    for fi58 := 0; fi58 < 2; fi58++ {
        var v59 float64 = 0.0
        _ = fi58
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v52
}
