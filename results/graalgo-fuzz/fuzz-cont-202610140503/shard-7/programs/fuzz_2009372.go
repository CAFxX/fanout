package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int32; F2 uint64 }
type GenS4 struct { F0 int; F1 uint8; F2 bool }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint8; F2 int }
func genFn0() int8 {
    var v5 uint16 = 264
    v5 = (v5 | v5)
    if (true && true) {
        v6 := []int{2147483647, 0, 32767, 1}
        _ = v6
    } else {
        v7 := GenS1{F0: 32767, F1: 0, F2: -913}
        v5 = v5
        _ = v7
    }
    v8 := GenS2{F0: 128}
    _ = v5
    _ = v8
    return 42
}
func genFn1() uint32 {
    var v9 float64 = 3.14159
    for fi10 := 0; fi10 < 16; fi10++ {
        v9 += -2.25
        _ = fi10
    }
    switch "a" {
    case "\n\t":
        v11 := GenS1{F0: 32767, F1: 2, F2: -5}
        _ = v11
    case "foo bar":
        switch -2147483648 {
        case 42:
            v12 := []int{32768}
            _ = v12
        }
    }
    if false {
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 int64 = -101
            v14 = v14
            v15 := GenS2{F0: fi13}
            _ = fi13
            _ = v14
            _ = v15
        }
        if false {
            var v16 int8 = 0
            _ = v16
        } else {
            v17 := GenS3{F0: ((-2147483648 > 42) || (!(65535 >= 65535))), F1: 1, F2: 18446744073709551615}
            _ = v17
        }
    } else {
        fmt.Println(v9)
        v9 = v9
        v9 -= float64((-2.25 * 0.0))
    }
    _ = v9
    return 795
}

func main() {
    v1 := make(map[string]int)
    var v2 int = -952
    var v3 int32 = -38
    var v4 uint32 = 458
    if (!(1.5 == (-2.25 - 0.1))) {
        if false {
            if (!false) {
                fmt.Println(v3, v4)
                v18 := GenS4{F0: v2, F1: 128, F2: true}
                _ = v18
            }
        } else {
            v19 := []int{v2, 2147483647, (v2 << ((460) & 63))}
            fmt.Println(v4, v3)
            _ = v19
        }
        v4 = genFn1()
    } else {
        v20 := []int{(-(v2)), (v2 | v2)}
        _ = v20
    }
    v21 := []int{v2}
    for fi22 := 0; fi22 < 17; fi22++ {
        v4 += v4
        v4 = 42
        fmt.Println(v2, v2)
        _ = fi22
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        if false {
            v3 = (v3 / (((v3 & (v3 % ((v3) | 1)))) | 1))
            for fi24 := 0; fi24 < 2; fi24++ {
                fmt.Println(v2, v2)
                _ = fi24
            }
        } else {
            if len(v21) > 0 {
                v21[0] = v2
            }
            v4 -= 26
        }
        switch (v2 >> ((2) & 63)) {
        case 4294967295:
            switch "foo bar" {
            case "\n\t":
                var v25 int32 = v3
                _ = v25
            case "foo bar":
                fmt.Println(v4)
            case "hello":
                v26 := GenS4{F0: 127, F1: 38, F2: (!(true && true))}
                _ = v26
            default:
                if false {
                    var v27 uint8 = 127
                    _ = v27
                } else {
                    v3 = (v3 >> ((255) & 63))
                    if (0 != 0) {
                        if false {
                            switch "" {
                            case "foo bar":
                                fmt.Println(v4, v2, fi23)
                            case "":
                                v2 -= 184
                            case "a":
                                switch fi23 {
                                case -384:
                                    switch (-(v2)) {
                                    case 0:
                                        var v28 uint64 = 1
                                        _ = v28
                                    case 220:
                                        var v29 uint64 = 2
                                        _ = v29
                                    default:
                                        v30 := GenS5{F0: 29, F1: 309, F2: (v4 / ((v4) | 1))}
                                        _ = v30
                                    }
                                }
                            default:
                                v3 += v3
                            }
                        }
                        v1[((("" + "") + ("" + ("foo bar" + "hello"))) + ((("a" + ("hello" + ("" + "foo bar"))) + "foo bar") + ("\n\t" + "foo bar")))] = fi23
                    }
                }
            }
        default:
            if (173 != 1) {
                for fi31 := 0; fi31 < 19; fi31++ {
                    v32 := GenS5{F0: 42, F1: 1, F2: genFn1()}
                    switch 128 {
                    case -9223372036854775808:
                        v33 := GenS6{F0: 255, F1: 42, F2: -697}
                        var v34 uint8 = 1
                        _ = v33
                        _ = v34
                    default:
                        v35 := GenS1{F0: fi23, F1: 375, F2: v3}
                        v36 := v1[""]
                        _ = v35
                        _ = v36
                    }
                    _ = fi31
                    _ = v32
                }
            } else {
                v37 := v1["foo bar"]
                _ = v37
            }
        }
        _ = len(v1)
        _ = fi23
    }
    _ = len(v1)
    _ = len(v21)
    _ = len(v21)
    v38 := GenS4{F0: -9223372036854775808, F1: 2, F2: false}
    var v39 uint = 1
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
