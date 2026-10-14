package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint; F2 int32 }
type GenS5 struct { F0 int16; F1 string; F2 uint16 }
type GenS6 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int16) uint16 {
    fmt.Println(p0)
    if (!(-2.25 <= 0.1)) {
        if true {
            fmt.Println(p0)
            var v7 uint8 = 2
            _ = v7
        }
    } else {
        p0 = -714
    }
    var v8 uint = 467
    var v9 int = 32768
    _ = p0
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "\n\t"
    var v3 int8 = 83
    var v4 uint32 = 65535
    var v5 float32 = -2.25
    var v6 int32 = -2147483648
    v10 := GenS1{F0: (int(v3) < 9223372036854775807), F1: 0}
    v11 := GenS1{F0: (32767 == 249), F1: 0}
    var v12 uint = 2
    if (!(!(v5 > float32((3.14159 + 3.14159))))) {
        switch "hello" {
        case "hello":
            switch -2147483648 {
            case 128:
                if (761 == 42) {
                    switch "a" {
                    case "hello":
                        for fi13 := 0; fi13 < 5; fi13++ {
                            v12 = v12
                            v2 = v2
                            _ = fi13
                        }
                    default:
                        var v14 int16 = 915
                        _ = v14
                    }
                    var v15 bool = (-2.25 <= float64((v5 * 0.1)))
                    _ = v15
                }
            case -821:
                v6 = (v6 << ((2) & 63))
            }
        case "a":
            v16 := []int{9223372036854775807, 127}
            _ = v16
        case "foo bar":
            if (!true) {
                v17 := GenS1{F0: ((0.1 - v5) > v5), F1: (v6 + (v6 << ((18446744073709551615) & 63)))}
                var v18 int = 128
                for fi19 := 0; fi19 < 18; fi19++ {
                    for fi20 := 0; fi20 < 20; fi20++ {
                        if (32767 == 778) {
                            var v21 int32 = (^(v6))
                            var v22 string = v2
                            _ = v21
                            _ = v22
                        } else {
                            if (true && false) {
                                v23 := GenS1{F0: (0 < (v6 << ((v12) & 63))), F1: (v6 & (v6 >> (((v12 & (v12 % ((v12) | 1)))) & 63)))}
                                v24 := []int{(-(v18))}
                                _ = v23
                                _ = v24
                            }
                        }
                        fmt.Println(v18)
                        fmt.Println(v4)
                        var v25 string = ((v2 + ((("hello" + "") + v2) + "a")) + "a")
                        if false {
                            _ = len(v1)
                        }
                        var v26 int64 = 94
                        _ = fi20
                        _ = v25
                        _ = v26
                    }
                    if false {
                        if (42 == 4294967295) {
                            v27 := GenS2{F0: (v4 | v4)}
                            var v28 bool = true
                            v2 = v2
                            _ = v27
                            _ = v28
                        }
                        switch v2 {
                        case "hello":
                            if false {
                                fmt.Println(v2, v18, v6)
                                v29 := GenS3{F0: int(v4), F1: 42}
                                for fi30 := 0; fi30 < 10; fi30++ {
                                    var v31 float64 = (-2.25 * 0.0)
                                    _ = fi30
                                    _ = v31
                                }
                                _ = v29
                            }
                        }
                        if (-1 > 1) {
                            v32 := GenS1{F0: (float64(0.1) != (1.5 + 0.1)), F1: v6}
                            v33 := []int{(v18 / (((fi19 - (v18 / (((v18 / (((fi19 % ((v18) | 1))) | 1))) | 1)))) | 1)), (v18 & 128)}
                            _ = v32
                            _ = v33
                        } else {
                            if ((true && (true || (false || (!(-128 != (^(v3))))))) && true) {
                                v2 = v2
                            } else {
                                v34 := GenS4{F0: -1, F1: (v12 % ((0) | 1)), F2: int32(v4)}
                                _ = v34
                            }
                            for fi35 := 0; fi35 < 10; fi35++ {
                                v36 := []int{v18, (fi35 / (((fi35 << ((v12) & 63))) | 1))}
                                v3 = v3
                                var v37 uint = v12
                                var v38 int64 = 9223372036854775807
                                _ = fi35
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                        }
                    }
                    v39 := GenS5{F0: 786, F1: v2, F2: 2}
                    _ = fi19
                    _ = v39
                }
                _ = v17
                _ = v18
            }
        default:
            if true {
                if (((1 < 766) && (14 <= 2)) || true) {
                    var v40 uint16 = 0
                    if false {
                        v41 := v1[(v2 + "\n\t")]
                        _ = v41
                    }
                    _ = v40
                }
                var v42 int64 = 9223372036854775807
                v4 = uint32(v12)
                _ = v42
            } else {
                fmt.Println(v6, v2, v12)
                for fi43 := 0; fi43 < 5; fi43++ {
                    v44 := GenS4{F0: (v6 << ((354) & 63)), F1: (v12 >> ((255) & 63)), F2: 2147483647}
                    _ = fi43
                    _ = v44
                }
            }
        }
    }
    var v45 uint64 = 2
    var v46 string = (v2 + v2)
    v47 := GenS4{F0: 1, F1: (v12 % ((v12) | 1)), F2: -881}
    v5 -= 0.0
    _ = len(v1)
    v48 := GenS6{F0: 255, F1: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
