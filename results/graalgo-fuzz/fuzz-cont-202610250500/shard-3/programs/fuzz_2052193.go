package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint }
type GenS2 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint8; F1 int }
func genFn0() int16 {
    var v4 int = -1
    var v5 int32 = -1
    if true {
        fmt.Println(v4, v5)
        if (65535 != 867) {
            v6 := []int{v4, (v4 / ((9223372036854775807) | 1)), v4}
            v5 = v5
            _ = v6
        }
        fmt.Println(v5, v4)
    } else {
        var v7 int32 = (^(v5))
        var v8 uint = 4294967295
        if ((v5 % ((v5) | 1)) <= (-(v5))) {
            v5 = v7
            fmt.Println(v5, v5)
            switch ("hello" + "hello") {
            case "hello":
                v8 = v8
            case "a":
                for fi9 := 0; fi9 < 4; fi9++ {
                    var v10 uint64 = 924
                    _ = fi9
                    _ = v10
                }
            case "foo bar":
                var v11 int = 32768
                _ = v11
            default:
                v8 += (^(v8))
            }
        } else {
            v12 := GenS1{F0: 532, F1: (^(v8)), F2: (v8 << ((v8) & 63))}
            _ = v12
        }
        _ = v7
        _ = v8
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        fmt.Println(v5)
        v14 := GenS2{F0: 826, F1: 255, F2: (v5 - v5)}
        v5 = v5
        _ = fi13
        _ = v14
    }
    var v15 int32 = 0
    _ = v15
    _ = v4
    _ = v5
    return 1000
}
func genFn1() int32 {
    v16 := []int{0, 0}
    if false {
        var v17 float32 = 1.5
        _ = v17
    } else {
        var v18 uint32 = 2147483647
        fmt.Println(v18)
        v18 = (v18 ^ v18)
        _ = v18
    }
    if (-128 < 100) {
        var v19 uint8 = 42
        var v20 int64 = 42
        _ = v19
        _ = v20
    } else {
        var v21 bool = false
        switch ("hello" + "\n\t") {
        case "":
            v21 = false
        case "\n\t":
            var v22 bool = v21
            _ = v22
        }
        v21 = (!v21)
        _ = v21
    }
    _ = v16
    return 687
}

func main() {
    var v1 float64 = float64(-2.25)
    v2 := make(map[uint32]bool)
    v3 := []uint32{2, 0, 42}
    var v23 int = 32767
    var v24 float32 = 0.0
    v2[502] = (255 < 593)
    v25 := GenS1{F0: 4294967295, F1: 0, F2: 42}
    for fi26 := 0; fi26 < 6; fi26++ {
        if true {
            fmt.Println(v23)
            if false {
                v27 := GenS3{F0: -9223372036854775808}
                var v28 uint64 = 334
                switch fi26 {
                case -256:
                    v29 := GenS4{F0: -2147483648, F1: (v23 ^ v23)}
                    _ = v29
                }
                _ = v27
                _ = v28
            } else {
                switch v23 {
                case 128:
                    var v30 uint8 = 211
                    _ = v30
                case -1:
                    if false {
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                        v31 := GenS5{F0: -4}
                        _ = v31
                    } else {
                        var v32 int32 = genFn1()
                        switch "foo bar" {
                        case "":
                            _ = len(v3)
                        default:
                            v33 := GenS6{F0: 2}
                            _ = v33
                        }
                        _ = v32
                    }
                case 32767:
                    var v34 uint = 18446744073709551615
                    _ = v34
                }
                v35 := v2[633]
                v36 := GenS6{F0: 588}
                _ = v35
                _ = v36
            }
            var v37 int32 = -1
            _ = v37
        } else {
            v3 = append(v3, 67)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v38 := GenS5{F0: -122}
            _ = v38
        }
        fmt.Println(v23, fi26)
        _ = fi26
    }
    v39 := GenS5{F0: -128}
    switch 42 {
    case 2147483647:
        v23 = v23
    case 32768:
        for fi40 := 0; fi40 < 13; fi40++ {
            if (9223372036854775807 < -901) {
                v23 += fi40
                v23 = v23
                for fi41 := 0; fi41 < 8; fi41++ {
                    if (true && true) {
                        switch (fi40 | (fi41 & 32767)) {
                        case 2147483647:
                            var v42 uint = 639
                            _ = v42
                        }
                    }
                    fmt.Println(v23, v23)
                    v23 -= (^(v23))
                    _ = fi41
                }
            } else {
                v23 = (v23 % (((v23 & v23)) | 1))
                for fi43 := 0; fi43 < 20; fi43++ {
                    if false {
                        fmt.Println(v23, v23, v1)
                    }
                    v44 := GenS3{F0: 1}
                    var v45 uint8 = 0
                    if (490 < 1) {
                        v1 -= v1
                        v46 := v2[401]
                        v46 = (1 != v23)
                        _ = v46
                    } else {
                        var v47 bool = true
                        var v48 float64 = 0.1
                        _ = v47
                        _ = v48
                    }
                    v49 := GenS5{F0: -115}
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v49
                }
                for fi50 := 0; fi50 < 18; fi50++ {
                    v24 -= (float32((v1 - float64(float32(v1)))) * 1e10)
                    v51 := v2[2147483648]
                    _ = fi50
                    _ = v51
                }
            }
            _ = fi40
        }
    default:
        if len(v3) > 0 {
            v3[0] = 65535
        }
    }
    v24 = float32((float64((((v24 * v24) + float32((v1 - (v1 / 1.0)))) + v24)) + 3.14159))
    switch v23 {
    case -2147483648:
        var v52 uint8 = 2
        _ = v52
    default:
        v53 := GenS7{F0: 65535, F1: 0, F2: 0}
        _ = v53
    }
    switch v23 {
    case -800:
        v54 := GenS3{F0: 1}
        _ = v54
    case 127:
        fmt.Println(v1, v24, v23)
    default:
        for fi55 := 0; fi55 < 16; fi55++ {
            v56 := GenS8{F0: (-(fi55))}
            _ = fi55
            _ = v56
        }
    }
    v57 := GenS5{F0: 0}
    var v58 float32 = float32(((0.0 * -2.25) + v1))
    v59 := GenS9{F0: 0, F1: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v39
    _ = v57
    _ = v58
    _ = v59
}
