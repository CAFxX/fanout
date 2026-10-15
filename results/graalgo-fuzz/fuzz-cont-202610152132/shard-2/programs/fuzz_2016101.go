package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int16; F1 uint64; F2 int }
type GenS3 struct { F0 uint32; F1 uint16 }
func genFn0() uint8 {
    var v7 uint64 = 0
    var v8 float32 = (1e10 * 0.0)
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 int8) uint8 {
    var v9 uint64 = 236
    var v10 float64 = 3.14159
    var v11 float64 = 1.5
    if (127 >= 0) {
        var v12 uint32 = 2147483647
        for fi13 := 0; fi13 < 10; fi13++ {
            v11 += v10
            v9 += (v9 * 18446744073709551615)
            _ = fi13
        }
        _ = v12
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 int32 = 0
                _ = fi15
                _ = v16
            }
            if false {
                fmt.Println(fi14, v10, v9)
            } else {
                if false {
                    var v17 float32 = float32(v10)
                    v18 := GenS1{F0: (v9 ^ (^(v9))), F1: 1}
                    v19 := []int{0, -80, (-(fi14))}
                    _ = v17
                    _ = v18
                    _ = v19
                }
                if ((fi14 % (((fi14 >> ((255) & 63))) | 1)) > 0) {
                    v20 := GenS1{F0: v9, F1: 9223372036854775807}
                    v21 := []int{fi14, fi14}
                    v11 = v10
                    _ = v20
                    _ = v21
                }
                v22 := GenS1{F0: (v9 % (((^(v9))) | 1)), F1: -250}
                _ = v22
            }
            _ = fi14
        }
        fmt.Println(v9, v9)
        v23 := []int{42, 0}
        _ = v23
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 171
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]uint32)
    var v3 int8 = -85
    var v4 uint32 = 42
    var v5 float32 = 0.1
    v6 := []uint32{v4, (v4 >> ((280) & 63))}
    v24 := GenS2{F0: -1, F1: 4294967295, F2: 4294967295}
    for fi25 := 0; fi25 < 14; fi25++ {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v1, v4, v4)
        case "\n\t":
            fmt.Println(v3, v4)
        default:
            v26 := GenS1{F0: 42, F1: -1}
            _ = v26
        }
        fmt.Println(fi25, v5)
        _ = fi25
    }
    var v27 string = "a"
    if (1 != 2) {
        v28 := v2[v4]
        fmt.Println(v27)
        for fi29 := 0; fi29 < 13; fi29++ {
            for fi30 := 0; fi30 < 8; fi30++ {
                if (!(!(!(0 <= 32767)))) {
                    var v31 int32 = 42
                    v32 := v2[v28]
                    v3 -= v3
                    switch "\n\t" {
                    case "hello":
                        var v33 int8 = v3
                        _ = v33
                    case "\n\t":
                        if false {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v34 int32 = (^(v31))
                            _ = v34
                        } else {
                            v28 -= 65535
                        }
                    }
                    switch v27 {
                    case "a":
                        fmt.Println(v31, v1)
                    case "\n\t":
                        var v35 int16 = 32767
                        _ = v35
                    case "foo bar":
                        v31 -= (v31 / ((-289) | 1))
                    }
                    if (v27 > v27) {
                        if (55 >= v3) {
                            v36 := GenS1{F0: 768, F1: -9223372036854775808}
                            v31 = (v31 ^ v31)
                            fmt.Println(v32, v31)
                            fmt.Println(fi29, v1)
                            var v37 bool = (!true)
                            _ = v36
                            _ = v37
                        }
                    } else {
                        v2[(v32 >> ((2) & 63))] = (v32 + (v4 >> ((1) & 63)))
                    }
                    _ = v31
                    _ = v32
                } else {
                    v38 := v2[v28]
                    _ = v38
                }
                var v39 int8 = (v3 ^ (v3 - v3))
                switch (fi29 / ((32767) | 1)) {
                case 4294967295:
                    var v40 uint8 = 128
                    _ = v40
                }
                _ = fi30
                _ = v39
            }
            for fi41 := 0; fi41 < 20; fi41++ {
                if ((fi29 / ((2147483647) | 1)) == 1) {
                    if ((v5 * (float32(v1) - (v5 * 0.0))) > v5) {
                        v3 -= v3
                        v2[v28] = v4
                    }
                    v3 = (v3 >> ((0) & 63))
                } else {
                    var v42 uint32 = 615
                    _ = v42
                }
                var v43 string = "a"
                _ = fi41
                _ = v43
            }
            _ = fi29
        }
        _ = v28
    } else {
        v44 := GenS2{F0: 892, F1: 18446744073709551615, F2: 128}
        if (!false) {
            switch "\n\t" {
            case "a":
                v45 := GenS3{F0: 0, F1: 65535}
                _ = v45
            }
            for fi46 := 0; fi46 < 12; fi46++ {
                v47 := v2[v4]
                var v48 int32 = 930
                var v49 float64 = 1.5
                _ = fi46
                _ = v47
                _ = v48
                _ = v49
            }
        }
        for fi50 := 0; fi50 < 20; fi50++ {
            var v51 uint32 = 1
            _ = fi50
            _ = v51
        }
        _ = v44
    }
    v52 := GenS3{F0: v4, F1: 1}
    if true {
        v1 -= v1
        if (!(!true)) {
            v1 -= v1
        } else {
            v3 -= v3
            v53 := GenS1{F0: 1, F1: 0}
            v1 = (((float64(1.5) + (v1 - v1)) - 3.14159) * v1)
            _ = v53
        }
    }
    fmt.Println(v27, v3)
    v5 = v5
    fmt.Println(v1, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
