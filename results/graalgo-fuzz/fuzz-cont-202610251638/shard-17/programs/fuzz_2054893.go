package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 string; F1 int; F2 int }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int32; F1 bool }
func genFn0(p0 int8, p1 int64, p2 uint8) uint16 {
    p0 = 127
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint64 = 42
    v2 := []int{0}
    v3 := []int{-9223372036854775808, -2147483648, 66}
    v1 = v1
    var v4 string = "foo bar"
    switch ((v4 + v4) + ("a" + (("" + v4) + "\n\t"))) {
    case "foo bar":
        v5 := GenS1{F0: ("\n\t" + (((v4 + "hello") + v4) + v4)), F1: ("hello" + "\n\t"), F2: 5}
        _ = v5
    }
    if (!true) {
        v1 += 86
        if true {
            v4 = (v4 + "foo bar")
        } else {
            if true {
                if len(v3) > 0 {
                    v3[0] = -954
                }
                v4 = ((v4 + v4) + v4)
                fmt.Println(v4, v1)
            } else {
                if true {
                    v6 := GenS2{F0: 255}
                    fmt.Println(v4)
                    fmt.Println(v4)
                    _ = v6
                }
                var v7 float32 = float32((0.0 * -2.25))
                _ = v7
            }
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v8 int8 = 1
        _ = v8
    }
    if true {
        if ((!(v1 >= v1)) || true) {
            var v9 uint32 = 2147483647
            _ = v9
        } else {
            v4 = v4
        }
        switch -1 {
        case 824:
            v1 += 562
        case 0:
            switch v4 {
            case "a":
                v4 = v4
            case "hello":
                v10 := GenS3{F0: -1, F1: -686}
                _ = v10
            case "\n\t":
                v1 -= v1
            default:
                v4 = v4
            }
        default:
            if len(v2) > 0 {
                v2[0] = -869
            }
        }
    } else {
        fmt.Println(v4)
        v1 = 0
    }
    var v11 float32 = 0.0
    v12 := GenS3{F0: -9223372036854775808, F1: 418}
    var v13 uint32 = 278
    if true {
        var v14 string = (v4 + "a")
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 int32 = int32(v13)
            v11 = float32(3.14159)
            var v17 float64 = 3.14159
            _ = fi15
            _ = v16
            _ = v17
        }
        if false {
            for fi18 := 0; fi18 < 11; fi18++ {
                v19 := GenS1{F0: v4, F1: (v4 + v14), F2: 1}
                v11 = v11
                _ = fi18
                _ = v19
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v1 = 2
            switch v14 {
            case "foo bar":
                var v20 int32 = -1
                _ = v20
            default:
                var v21 float64 = (-2.25 / ((float64((v11 + float32(1e10))) - (0.1 / 3.14159)) / 1e10))
                _ = v21
            }
            if true {
                v14 = v14
            } else {
                if false {
                    switch v4 {
                    case "foo bar":
                        v13 = v13
                    }
                    var v22 uint32 = uint32(v1)
                    var v23 uint16 = 957
                    _ = v22
                    _ = v23
                }
                var v24 int16 = int16(v1)
                v1 = (v1 / ((v1) | 1))
                _ = v24
            }
        }
        _ = v14
    } else {
        switch 42 {
        case 128:
            var v25 float64 = ((3.14159 / ((-2.25 + (((1.5 + -2.25) / 1.5) / 1.5)) + (1.5 * 0.0))) + 3.14159)
            _ = v25
        case 32767:
            switch 1 {
            case 623:
                v13 += 265
            case 395:
                v26 := GenS4{F0: v4, F1: 32768, F2: 32767}
                _ = v26
            case -471:
                v27 := GenS5{F0: int64(v1), F1: 2}
                _ = v27
            default:
                v28 := GenS4{F0: ("" + v4), F1: -1, F2: 42}
                _ = v28
            }
        case 9223372036854775807:
            for fi29 := 0; fi29 < 19; fi29++ {
                _ = len(v3)
                if len(v3) > 0 {
                    v3[0] = 583
                }
                for fi30 := 0; fi30 < 19; fi30++ {
                    var v31 int32 = 0
                    var v32 int8 = 42
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v33 uint = 886
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = fi29
            }
        default:
            v13 -= (^(v13))
        }
        if (-2.25 != float64(v11)) {
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            var v34 uint = uint(v1)
            if len(v2) > 0 {
                v2[0] = 42
            }
            _ = v34
        } else {
            if true {
                var v35 uint64 = uint64(v13)
                fmt.Println(v11, v35, v11)
                if false {
                    switch -2147483648 {
                    case 1:
                        var v36 string = v4
                        _ = v36
                    case 128:
                        v35 = (v1 >> ((65535) & 63))
                    default:
                        v37 := GenS6{F0: int32(v35), F1: false}
                        _ = v37
                    }
                } else {
                    var v38 uint64 = (v35 >> ((0) & 63))
                    _ = v38
                }
                _ = v35
            } else {
                switch 4294967295 {
                case -487:
                    var v39 float32 = (v11 - v11)
                    _ = v39
                case -2147483648:
                    if true {
                        v40 := GenS2{F0: 65535}
                        for fi41 := 0; fi41 < 6; fi41++ {
                            fmt.Println(v11, fi41)
                            v13 = 65535
                            if len(v2) > 0 {
                                v2[0] = (fi41 >> ((244) & 63))
                            }
                            _ = fi41
                        }
                        switch -2147483648 {
                        case -9223372036854775808:
                            var v42 uint16 = genFn0(int8(v13), 172, 128)
                            _ = v42
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = v40
                    }
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = len(v3)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
