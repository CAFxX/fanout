package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 string }
type GenS2 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint64 }
func genFn0() uint64 {
    v5 := GenS1{F0: 9223372036854775807, F1: 1, F2: ""}
    switch 15 {
    case -980:
        for fi6 := 0; fi6 < 8; fi6++ {
            switch (fi6 + fi6) {
            case 9223372036854775807:
                var v7 int32 = -2147483648
                _ = v7
            case 2147483647:
                for fi8 := 0; fi8 < 13; fi8++ {
                    var v9 int64 = 372
                    _ = fi8
                    _ = v9
                }
            case 42:
                v10 := []int{fi6, fi6, (fi6 + (fi6 + (fi6 + 42)))}
                _ = v10
            default:
                var v11 bool = ((568 >= 1) && (!false))
                _ = v11
            }
            var v12 int8 = 22
            if (!(0 >= 9223372036854775807)) {
                var v13 uint64 = 1
                if (true && false) {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        v15 := GenS1{F0: 9223372036854775807, F1: (v13 & 18446744073709551615), F2: "\n\t"}
                        if (!false) {
                            var v16 uint32 = 826
                            v13 = v13
                            _ = v16
                        }
                        _ = fi14
                        _ = v15
                    }
                    v17 := GenS2{F0: (v13 % ((v13) | 1)), F1: v13, F2: -9223372036854775808}
                    _ = v17
                } else {
                    v12 -= 127
                }
                _ = v13
            } else {
                var v18 float64 = 3.14159
                var v19 float32 = -2.25
                var v20 int16 = 42
                _ = v18
                _ = v19
                _ = v20
            }
            _ = fi6
            _ = v12
        }
    case 1:
        var v21 int = 128
        _ = v21
    case 127:
        var v22 uint = 65535
        _ = v22
    }
    var v23 uint64 = 0
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 uint8 = 144
        fmt.Println(fi24)
        _ = fi24
        _ = v25
    }
    if (!(-9 > 127)) {
        v23 = v23
        v26 := GenS3{F0: -103}
        var v27 uint8 = 128
        _ = v26
        _ = v27
    }
    _ = v23
    _ = v5
    return v23
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 0.1
    v3 := []int{787, -198}
    v4 := make(map[int]string)
    var v28 int32 = 259
    v29 := GenS4{F0: 125}
    if false {
        v30 := GenS2{F0: genFn0(), F1: 531, F2: 878}
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 uint8 = 2
            _ = fi31
            _ = v32
        }
        _ = v30
    } else {
        for fi33 := 0; fi33 < 14; fi33++ {
            if len(v3) > 0 {
                v3[0] = 790
            }
            _ = fi33
        }
        v2 = float64((1.5 - 0.0))
        fmt.Println(v28, v28)
    }
    var v34 float64 = (v2 - v2)
    var v35 uint16 = 181
    var v36 int32 = (v28 ^ v28)
    if false {
        v35 += (v35 << ((532) & 63))
        fmt.Println(v35, v36)
    } else {
        for fi37 := 0; fi37 < 12; fi37++ {
            v2 -= -2.25
            _ = fi37
        }
        v38 := GenS3{F0: -1}
        _ = v38
    }
    var v39 int = -279
    var v40 float64 = (v34 * v2)
    if (false || (2147483647 >= 330)) {
        v2 = float64(float32(v40))
        if (42 != 22) {
            v40 -= -2.25
            var v41 int32 = 2147483647
            var v42 uint = 18446744073709551615
            _ = v41
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 9; fi43++ {
        if ((!(2 > (v35 | 65535))) && (!true)) {
            v44 := GenS2{F0: 425, F1: 4294967295, F2: 0}
            v45 := GenS1{F0: 42, F1: genFn0(), F2: ""}
            switch "a" {
            case "\n\t":
                for fi46 := 0; fi46 < 10; fi46++ {
                    fmt.Println(fi46)
                    v34 -= ((v40 * 1e10) - (float64(0.0) * v2))
                    var v47 int16 = int16(v28)
                    if false {
                        if (uint64(v35) != genFn0()) {
                            switch (v39 + fi43) {
                            case -1:
                                var v48 float64 = v34
                                _ = v48
                            }
                            v35 = (v35 | (^(v35)))
                        }
                        v49 := GenS3{F0: int64(v28)}
                        _ = v49
                    } else {
                        var v50 uint8 = 127
                        if true {
                            v36 = -1
                            v51 := v4[(v39 | (v39 % ((fi46) | 1)))]
                            _ = v51
                        } else {
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                            v36 = (v28 & 0)
                            v39 += v39
                        }
                        _ = v50
                    }
                    _ = fi46
                    _ = v47
                }
            case "":
                v52 := GenS5{F0: -1, F1: 517, F2: v35}
                _ = v52
            case "hello":
                if len(v3) > 0 {
                    v3[0] = 1
                }
            }
            _ = v44
            _ = v45
        } else {
            for fi53 := 0; fi53 < 12; fi53++ {
                v54 := GenS6{F0: 42}
                fmt.Println(v39, v40, v34)
                v55 := v4[fi53]
                _ = fi53
                _ = v54
                _ = v55
            }
        }
        fmt.Println(v35, v40, v39)
        for fi56 := 0; fi56 < 2; fi56++ {
            v57 := v1[128]
            _ = fi56
            _ = v57
        }
        _ = fi43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
}
