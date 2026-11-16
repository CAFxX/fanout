package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int) uint16 {
    fmt.Println(p0)
    p0 -= p0
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 float64 = 1.5
        _ = fi7
        _ = v8
    }
    _ = p0
    return 1
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[string]int)
    var v3 uint64 = 625
    v4 := make(map[string]uint32)
    v5 := []int{(v1 + v1), v1}
    v6 := []string{"", ("a" + "")}
    v1 = (v1 & int(v3))
    v1 += (v1 & (v1 + 32768))
    fmt.Println(v1)
    var v9 uint64 = v3
    if ((65535 != 2147483647) || (-1 == -634)) {
        var v10 int8 = 62
        v6 = append(v6, ("a" + ("" + "\n\t")))
        var v11 int32 = 1
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        switch v1 {
        case 4294967295:
            var v13 bool = false
            _ = v13
        default:
            v9 = v3
        }
        if false {
            v9 = v3
            v3 = v3
        }
        _ = fi12
    }
    var v14 int = v1
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 uint64 = v9
        fmt.Println(v3)
        v16 = v9
        _ = fi15
        _ = v16
    }
    var v17 uint = uint(v3)
    fmt.Println(v14, v3)
    var v18 int32 = 42
    for fi19 := 0; fi19 < 12; fi19++ {
        fmt.Println(v17, v3, v14)
        _ = fi19
    }
    v18 = v18
    switch "foo bar" {
    case "hello":
        v20 := GenS1{F0: v18, F1: v17}
        _ = v20
    case "a":
        if ((v14 | int(v17)) <= v1) {
            var v21 float32 = float32(1.5)
            for fi22 := 0; fi22 < 20; fi22++ {
                delete(v4, "foo bar")
                switch "a" {
                case "hello":
                    for fi23 := 0; fi23 < 16; fi23++ {
                        if true {
                            v21 = v21
                            var v24 int = (v14 / ((324) | 1))
                            _ = v24
                        }
                        var v25 int8 = -1
                        if (42 > 34) {
                            v21 = v21
                            fmt.Println(v17)
                        }
                        _ = fi23
                        _ = v25
                    }
                case "a":
                    for fi26 := 0; fi26 < 13; fi26++ {
                        if true {
                            v21 += 0.0
                            v27 := GenS2{F0: 32767}
                            _ = v27
                        } else {
                            var v28 bool = (!(142 > 1))
                            var v29 uint = (v17 / (((v17 ^ v17)) | 1))
                            if (v9 > (v3 / (((v9 * v3)) | 1))) {
                                if (!true) {
                                    if v28 {
                                        switch "foo bar" {
                                        case "hello":
                                            if true {
                                                var v30 uint = v29
                                                _ = v30
                                            }
                                            fmt.Println(v28, v28, v3)
                                        case "a":
                                            if len(v5) > 0 {
                                                v5[0] = v14
                                            }
                                        case "\n\t":
                                            v3 -= 42
                                        }
                                    } else {
                                        v1 += 233
                                        v28 = v28
                                        fmt.Println(fi22)
                                        var v31 uint32 = 264
                                        _ = v31
                                    }
                                    var v32 uint64 = v3
                                    v29 = 0
                                    _ = v32
                                }
                                if v28 {
                                    v33 := GenS1{F0: v18, F1: 0}
                                    _ = v33
                                }
                            } else {
                                fmt.Println(v9, fi22, v29)
                            }
                            if (!(0 <= 885)) {
                                if v28 {
                                    v34 := GenS3{F0: v9}
                                    _ = v34
                                }
                                if len(v6) > 0 {
                                    v6[0] = "\n\t"
                                }
                                fmt.Println(v3)
                            } else {
                                v21 = ((v21 / (((v21 / -2.25) * float32(((((3.14159 * (0.0 - float64((v21 - (v21 - 0.1))))) * 0.0) * 0.0) * 3.14159))) / v21)) + ((v21 / (1e10 + v21)) + v21))
                                fmt.Println(v14, v28, v17)
                            }
                            var v35 int16 = -32768
                            _ = v28
                            _ = v29
                            _ = v35
                        }
                        var v36 int8 = 42
                        var v37 uint8 = 250
                        _ = fi26
                        _ = v36
                        _ = v37
                    }
                }
                _ = fi22
            }
            for fi38 := 0; fi38 < 6; fi38++ {
                delete(v2, "a")
                _ = fi38
            }
            _ = v21
        }
    }
    if (true || true) {
        var v39 int = (v1 | 128)
        v9 = 303
        _ = v39
    } else {
        var v40 uint = v17
        if true {
            var v41 int16 = 1000
            v42 := GenS3{F0: (v3 ^ (v9 ^ v9))}
            _ = v41
            _ = v42
        } else {
            switch 127 {
            case 32768:
                v3 -= v3
            case 1:
                var v43 uint = (^(v40))
                _ = v43
            }
            v9 -= v9
        }
        _ = len(v5)
        _ = v40
    }
    for fi44 := 0; fi44 < 15; fi44++ {
        _ = len(v4)
        v45 := GenS4{F0: 0}
        _ = len(v2)
        _ = fi44
        _ = v45
    }
    v17 -= (^(v17))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
