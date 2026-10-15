package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int8; F1 int16; F2 int32 }
type GenS5 struct { F0 int64; F1 bool; F2 uint8 }
func genFn0(p0 int32, p1 int) uint16 {
    v6 := GenS1{F0: 2, F1: -1}
    var v7 uint16 = 408
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ("a" + ("foo bar" + "foo bar"))
    var v3 int64 = -9223372036854775808
    var v4 bool = (1 == 2147483647)
    v5 := make(map[int]uint32)
    if (0 != -128) {
        var v8 float64 = 1e10
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := []int{}
        if v4 {
            if v4 {
                if ((((65535 >= 842) && v4) && (42 >= 2)) || false) {
                    v3 = (v3 / ((v3) | 1))
                    for fi11 := 0; fi11 < 15; fi11++ {
                        var v12 uint8 = 127
                        _ = fi11
                        _ = v12
                    }
                } else {
                    fmt.Println(v2)
                    v5[fi9] = 2147483648
                }
            } else {
                v13 := GenS2{F0: 2, F1: -241}
                _ = v13
            }
            fmt.Println(v3)
            var v14 string = v2
            _ = v14
        } else {
            v4 = true
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 float64 = 1.5
            if (((!v4) && (!v4)) && v4) {
                fmt.Println(v2, fi9)
                if (true && (-2147483648 <= 950)) {
                    fmt.Println(v2)
                }
            }
            v16 -= 0.0
            _ = fi15
            _ = v16
        }
        _ = fi9
        _ = v10
    }
    v17 := GenS2{F0: 2, F1: -32768}
    var v18 float64 = 0.1
    v18 = v18
    switch "a" {
    case "a":
        v19 := GenS1{F0: 0, F1: -1}
        _ = v19
    case "\n\t":
        v18 = (v18 * v18)
    case "":
        v20 := GenS2{F0: 224, F1: 1}
        _ = v20
    default:
        v18 = 0.1
    }
    var v21 int8 = 127
    for fi22 := 0; fi22 < 5; fi22++ {
        fmt.Println(v4)
        _ = fi22
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        _ = len(v5)
        for fi24 := 0; fi24 < 17; fi24++ {
            delete(v5, fi23)
            switch 128 {
            case -2147483648:
                var v25 float64 = v18
                _ = v25
            case -886:
                var v26 int = 42
                _ = v26
            case 4294967295:
                var v27 int64 = v3
                if (false && v4) {
                    v27 = v3
                    v3 = v3
                    _ = len(v1)
                    var v28 uint = 881
                    if (0 >= 2147483648) {
                        v28 = 18446744073709551615
                        fmt.Println(fi24, v4)
                    } else {
                        v29 := []int{fi23}
                        v4 = true
                        fmt.Println(v21)
                        _ = v29
                    }
                    _ = v28
                } else {
                    v30 := v1[v2]
                    v3 -= v3
                    if false {
                        var v31 int64 = v27
                        var v32 uint32 = 65535
                        _ = v31
                        _ = v32
                    } else {
                        fmt.Println(v4)
                        v33 := GenS2{F0: 1, F1: 27}
                        if (127 > (-(v21))) {
                            v34 := GenS3{F0: 195, F1: 127}
                            v18 = v18
                            _ = v34
                        } else {
                            var v35 int16 = 1
                            var v36 uint64 = 2
                            v37 := []int{(fi23 + (fi24 & int(v36))), fi23, (fi23 * (fi23 << ((uint(v36)) & 63)))}
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        _ = v33
                    }
                    _ = v30
                }
                v38 := []int{}
                _ = v27
                _ = v38
            }
            v39 := GenS3{F0: v3, F1: 255}
            _ = fi24
            _ = v39
        }
        _ = fi23
    }
    var v40 uint32 = 65535
    v41 := v5[-672]
    if (!false) {
        fmt.Println(v21, v4)
        if false {
            v42 := GenS4{F0: (^(v21)), F1: 42, F2: 2147483647}
            var v43 uint8 = uint8(v41)
            var v44 float64 = (v18 * float64(((1e10 / 3.14159) * (1.5 - 1e10))))
            _ = v42
            _ = v43
            _ = v44
        }
    }
    var v45 int32 = -2147483648
    for fi46 := 0; fi46 < 3; fi46++ {
        v47 := GenS5{F0: v3, F1: ((v4 && v4) && ((v40 & v41) <= 433)), F2: 0}
        _ = fi46
        _ = v47
    }
    if (v21 == v21) {
        fmt.Println(v45, v21, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v5
}
