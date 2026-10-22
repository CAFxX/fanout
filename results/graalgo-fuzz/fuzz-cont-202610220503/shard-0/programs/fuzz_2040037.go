package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint; F1 bool; F2 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 int32; F1 int8; F2 uint8 }
type GenS9 struct { F0 int32; F1 int }
func genFn0(p0 uint32, p1 int64) uint16 {
    var v5 uint32 = p0
    v6 := GenS1{F0: 350, F1: 0, F2: p1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 629
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 42
    var v3 uint16 = 1
    var v4 string = "foo bar"
    if (!false) {
        switch 0 {
        case -9223372036854775808:
            fmt.Println(v4)
        case -206:
            if true {
                var v7 int = 1
                _ = v7
            }
        default:
            v8 := GenS1{F0: -1, F1: 1, F2: 572}
            _ = v8
        }
    }
    var v9 uint16 = 255
    if false {
        var v10 int8 = (v2 & v2)
        _ = v10
    }
    if true {
        var v11 int64 = 134
        _ = v11
    } else {
        for fi12 := 0; fi12 < 4; fi12++ {
            if false {
                for fi13 := 0; fi13 < 13; fi13++ {
                    v2 = v2
                    v14 := GenS1{F0: 422, F1: -1, F2: -257}
                    if ((1 > 4294967295) && ((1 > 2) && (0 == 1))) {
                        v2 = 0
                        v3 += v9
                    }
                    _ = fi13
                    _ = v14
                }
                fmt.Println(fi12)
                var v15 string = ("a" + "hello")
                _ = v15
            } else {
                var v16 float64 = (1.5 / (((3.14159 / 1.0) - (1e10 / 1e10)) + 0.1))
                if (-786 != (fi12 ^ 0)) {
                    if (710 >= 1) {
                        v4 = v4
                        v2 = -63
                    } else {
                        for fi17 := 0; fi17 < 18; fi17++ {
                            if false {
                                v18 := GenS1{F0: -2147483648, F1: -725, F2: -898}
                                _ = v18
                            } else {
                                v19 := GenS1{F0: -340, F1: 42, F2: -1}
                                _ = v19
                            }
                            if (!((!false) || ((fi17 >> ((4294967295) & 63)) <= (fi12 + -1)))) {
                                if true {
                                    v20 := GenS1{F0: -963, F1: 1, F2: 0}
                                    v2 -= (v2 >> ((813) & 63))
                                    _ = v20
                                }
                                if true {
                                    var v21 uint = 42
                                    v3 -= genFn0(1, 0)
                                    var v22 int64 = -352
                                    if ((((((v16 + v16) + (v16 + -2.25)) + float64((0.1 + (float32(float64((1e10 / 1.0))) / 0.1)))) * v16) * 1e10) >= v16) {
                                        v23 := GenS2{F0: (v22 | -856)}
                                        var v24 bool = (!((v21 / ((uint(v3)) | 1)) != v21))
                                        v24 = (((215 <= (^(v21))) && true) && (!(!false)))
                                        v24 = v24
                                        _ = v23
                                        _ = v24
                                    }
                                    _ = v21
                                    _ = v22
                                }
                            } else {
                                v25 := GenS3{F0: 88, F1: 2}
                                _ = v25
                            }
                            _ = fi17
                        }
                        v26 := GenS4{F0: 18446744073709551615, F1: false, F2: 2}
                        v27 := []int{}
                        _ = v26
                        _ = v27
                    }
                    delete(v1, (fi12 % ((128) | 1)))
                    var v28 string = v4
                    _ = v28
                }
                _ = v16
            }
            _ = fi12
        }
        v2 += 84
    }
    v3 += v9
    var v29 uint = 0
    if true {
        var v30 float64 = ((0.0 / ((float64((float32(1.5) + 3.14159)) * 1e10) / ((1.5 / 1.5) + ((0.1 / 1.0) + (0.1 / (1e10 - 0.0)))))) - ((-2.25 + -2.25) + (-2.25 - float64((0.0 - float32((0.0 - -2.25)))))))
        _ = v30
    } else {
        var v31 int64 = int64(v29)
        if false {
            if false {
                var v32 int8 = 0
                if (!true) {
                    v33 := GenS3{F0: 2, F1: 2}
                    v34 := GenS3{F0: 255, F1: 42}
                    if true {
                        v29 = v29
                        v35 := []int{-1, int(v31)}
                        _ = v35
                    }
                    _ = v33
                    _ = v34
                } else {
                    var v36 int64 = (^(v31))
                    v37 := GenS5{F0: false}
                    switch 9223372036854775807 {
                    case 881:
                        fmt.Println(v29, v31, v36)
                    case 128:
                        switch "foo bar" {
                        case "":
                            v38 := GenS6{F0: uint32(v9), F1: uint32(v9), F2: 0}
                            _ = v38
                        case "a":
                            v39 := GenS1{F0: 2147483647, F1: 375, F2: v36}
                            _ = v39
                        }
                    case -452:
                        var v40 int32 = -145
                        _ = v40
                    }
                    _ = v36
                    _ = v37
                }
                v41 := GenS7{F0: -1, F1: 42}
                _ = v32
                _ = v41
            }
        } else {
            fmt.Println(v31, v2, v9)
            switch int(v29) {
            case 459:
                v42 := GenS7{F0: -1, F1: 128}
                _ = v42
            }
        }
        _ = v31
    }
    fmt.Println(v4, v3, v9)
    if (!false) {
        if false {
            for fi43 := 0; fi43 < 1; fi43++ {
                fmt.Println(v9, v2, v29)
                for fi44 := 0; fi44 < 8; fi44++ {
                    v45 := v1[127]
                    var v46 string = "\n\t"
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                _ = fi43
            }
            switch 127 {
            case 4294967295:
                var v47 uint8 = 128
                _ = v47
            case 701:
                v48 := GenS7{F0: 2147483647, F1: 42}
                _ = v48
            }
            if true {
                var v49 float64 = 0.0
                _ = v49
            } else {
                switch (v4 + v4) {
                case "a":
                    v50 := GenS6{F0: 2147483648, F1: 1, F2: 427}
                    _ = v50
                }
                if (!true) {
                    for fi51 := 0; fi51 < 2; fi51++ {
                        if (!false) {
                            var v52 uint32 = 758
                            v29 = v29
                            _ = v52
                        } else {
                            var v53 int8 = 100
                            _ = v53
                        }
                        for fi54 := 0; fi54 < 5; fi54++ {
                            v29 = 255
                            v55 := []int{(fi51 * (fi51 / ((fi54) | 1))), (fi54 >> ((80) & 63)), fi54, (fi54 * (fi54 >> ((908) & 63)))}
                            fmt.Println(v4)
                            _ = fi54
                            _ = v55
                        }
                        var v56 uint16 = 42
                        _ = fi51
                        _ = v56
                    }
                    var v57 uint8 = 1
                    _ = v57
                }
            }
        } else {
            var v58 int64 = 429
            v9 -= (v3 << ((v29) & 63))
            _ = v58
        }
        fmt.Println(v2, v3, v3)
        v59 := []int{}
        _ = v59
    } else {
        v60 := GenS8{F0: int32(v29), F1: 112, F2: 42}
        switch v4 {
        case "hello":
            v61 := GenS9{F0: 647, F1: 9223372036854775807}
            _ = v61
        }
        var v62 uint32 = 1
        _ = v60
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
