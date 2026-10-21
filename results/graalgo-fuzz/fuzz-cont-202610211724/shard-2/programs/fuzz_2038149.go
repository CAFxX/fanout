package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint64 }
func genFn0() uint8 {
    var v5 int16 = 756
    if ((!true) || false) {
        v6 := GenS1{F0: 1, F1: true, F2: 180}
        v7 := GenS1{F0: 42, F1: (-55 == 171), F2: 6}
        v5 += (v5 >> ((65535) & 63))
        _ = v6
        _ = v7
    } else {
        var v8 float32 = -2.25
        v5 += v5
        _ = v8
    }
    var v9 uint64 = 393
    if ((v9 & 0) <= 42) {
        v5 = (v5 & v5)
        if false {
            switch -2147483648 {
            case 714:
                v9 += (^(v9))
            default:
                v5 += v5
            }
        } else {
            v10 := GenS1{F0: 1, F1: (-717 != 2147483647), F2: 255}
            _ = v10
        }
        var v11 uint8 = 1
        _ = v11
    } else {
        v5 += (^(v5))
        var v12 uint = 385
        _ = v12
    }
    _ = v5
    _ = v9
    return 56
}
func genFn1(p0 bool, p1 uint32, p2 bool) int64 {
    v13 := []int{0, -2147483648, 560}
    p2 = (!false)
    p1 -= p1
    var v14 float64 = float64(1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 717
}

func main() {
    var v1 uint16 = 42
    var v2 uint16 = (v1 + (v1 >> ((18446744073709551615) & 63)))
    var v3 uint64 = 811
    var v4 int8 = -123
    var v15 uint8 = 127
    var v16 bool = false
    var v17 int16 = 0
    for fi18 := 0; fi18 < 9; fi18++ {
        if v16 {
            switch fi18 {
            case 163:
                for fi19 := 0; fi19 < 14; fi19++ {
                    v20 := GenS1{F0: (v1 - v1), F1: false, F2: (v15 | 97)}
                    var v21 int8 = v4
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 18; fi22++ {
                switch "a" {
                case "hello":
                    fmt.Println(v1)
                default:
                    var v23 uint64 = 1
                    _ = v23
                }
                _ = fi22
            }
        } else {
            for fi24 := 0; fi24 < 19; fi24++ {
                switch fi18 {
                case 2147483647:
                    v25 := GenS2{F0: -854, F1: 65535}
                    _ = v25
                case 0:
                    if true {
                        var v26 string = ""
                        fmt.Println(v17, v15)
                        var v27 int64 = 42
                        _ = v26
                        _ = v27
                    }
                }
                v28 := GenS3{F0: -1, F1: (v17 & v17)}
                _ = fi24
                _ = v28
            }
            var v29 float32 = 1e10
            switch (fi18 | 599) {
            case -595:
                for fi30 := 0; fi30 < 15; fi30++ {
                    var v31 int = (fi30 / ((fi30) | 1))
                    if (!v16) {
                        fmt.Println(v31)
                        if (false && true) {
                            if v16 {
                                if (fi30 == fi18) {
                                    v32 := GenS3{F0: 659, F1: -1}
                                    _ = v32
                                } else {
                                    v33 := []int{128}
                                    v16 = false
                                    v4 = -20
                                    _ = v33
                                }
                                v34 := []int{(fi18 ^ fi18), (fi30 - int(v3)), v31}
                                _ = v34
                            } else {
                                var v35 uint32 = 65535
                                v1 = (v2 & (v1 - 414))
                                v36 := []int{fi30, (fi18 | -1), (v31 / ((fi18) | 1))}
                                _ = v35
                                _ = v36
                            }
                        } else {
                            v3 += (v3 >> ((255) & 63))
                        }
                        var v37 uint16 = v2
                        var v38 uint64 = (v3 % ((42) | 1))
                        _ = v37
                        _ = v38
                    }
                    if v16 {
                        var v39 int64 = -1
                        _ = v39
                    }
                    var v40 int64 = -9223372036854775808
                    v41 := GenS4{F0: 2147483647, F1: 2}
                    _ = fi30
                    _ = v31
                    _ = v40
                    _ = v41
                }
            case -9223372036854775808:
                v42 := GenS4{F0: 0, F1: v3}
                _ = v42
            default:
                var v43 float64 = (1e10 * -2.25)
                _ = v43
            }
            _ = v29
        }
        var v44 int32 = 1
        v2 = 1
        _ = fi18
        _ = v44
    }
    switch int(v3) {
    case 2147483647:
        v45 := GenS2{F0: -1, F1: 255}
        _ = v45
    default:
        for fi46 := 0; fi46 < 19; fi46++ {
            fmt.Println(v1, v15)
            v17 = v17
            _ = fi46
        }
    }
    var v47 float32 = 1e10
    v2 = (v1 | v2)
    var v48 int32 = 240
    fmt.Println(v2, v1)
    fmt.Println(v15, v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
}
