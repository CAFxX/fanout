package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 string; F2 uint16 }
type GenS4 struct { F0 uint; F1 int8; F2 uint }
type GenS5 struct { F0 string; F1 int64 }
func genFn0() int32 {
    var v7 int = 32767
    v7 -= v7
    _ = v7
    return -551
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 uint8 = 127
    var v5 bool = true
    v6 := make(map[int]uint32)
    if (!(0 != genFn0())) {
        var v8 int8 = 100
        _ = v8
    }
    var v9 float32 = (0.0 * (0.0 - float32(0.1)))
    var v10 string = v1
    if true {
        var v11 float32 = 1e10
        _ = v11
    }
    var v12 uint16 = 65535
    fmt.Println(v12, v12)
    v13 := GenS1{F0: 42}
    fmt.Println(v9, v10)
    v5 = true
    v9 += (v9 / (v9 * v9))
    v14 := GenS1{F0: 517}
    for fi15 := 0; fi15 < 17; fi15++ {
        switch "foo bar" {
        case "\n\t":
            v5 = (0 == 0)
        case "":
            if true {
                v16 := []int{(fi15 - (fi15 * fi15)), (fi15 / ((127) | 1)), fi15}
                v17 := GenS2{F0: (^(v4)), F1: 1, F2: 116}
                var v18 uint64 = 42
                _ = v16
                _ = v17
                _ = v18
            } else {
                v19 := GenS1{F0: 0}
                v9 = float32((3.14159 / -2.25))
                _ = v19
            }
        }
        if false {
            v10 = "\n\t"
            switch fi15 {
            case 127:
                v10 = v10
            }
        } else {
            fmt.Println(v1, fi15)
            var v20 uint16 = 1
            v21 := GenS1{F0: 2}
            _ = v20
            _ = v21
        }
        _ = fi15
    }
    var v22 float64 = 3.14159
    var v23 int16 = 1000
    if (((!((1e10 / v22) > 3.14159)) && (0 != 42)) && v5) {
        v12 = v12
    } else {
        v22 = 0.0
        v5 = (v5 && ((v4 ^ uint8(v12)) != v4))
        for fi24 := 0; fi24 < 18; fi24++ {
            switch (fi24 | (-(fi24))) {
            case 1:
                if v5 {
                    v22 = (v22 + 0.0)
                    v6[(fi24 << ((1) & 63))] = 4294967295
                    v25 := GenS3{F0: (v4 & v4), F1: v1, F2: (v12 ^ v12)}
                    _ = v25
                } else {
                    if v5 {
                        v5 = false
                    } else {
                        switch 127 {
                        case -1:
                            if v5 {
                                v3["foo bar"] = (true && v5)
                            }
                        }
                        if (0 == 0) {
                            v4 = (v4 ^ v4)
                            v26 := GenS4{F0: 42, F1: 1, F2: 65535}
                            _ = v26
                        } else {
                            var v27 uint32 = 2
                            fmt.Println(v1, v23)
                            v12 -= (v12 & (^(v12)))
                            _ = v27
                        }
                        delete(v6, 32767)
                    }
                }
            case -252:
                for fi28 := 0; fi28 < 7; fi28++ {
                    v22 -= v22
                    _ = fi28
                }
            default:
                v29 := []int{-2147483648}
                _ = v29
            }
            if (3.14159 <= (v22 * (float64(-2.25) / (1e10 - -2.25)))) {
                switch fi24 {
                case -1:
                    var v30 int = fi24
                    _ = v30
                case 17:
                    if true {
                        if v5 {
                            v4 = v4
                            switch fi24 {
                            case 2147483647:
                                fmt.Println(v22, v10, v23)
                            case -883:
                                for fi31 := 0; fi31 < 2; fi31++ {
                                    v10 = (v1 + v1)
                                    var v32 int64 = -832
                                    _ = fi31
                                    _ = v32
                                }
                            case 32767:
                                for fi33 := 0; fi33 < 13; fi33++ {
                                    var v34 int32 = 0
                                    _ = fi33
                                    _ = v34
                                }
                            default:
                                fmt.Println(v5)
                            }
                            v4 += (v4 + v4)
                        } else {
                            switch fi24 {
                            case -1:
                                v23 += v23
                            case 128:
                                for fi35 := 0; fi35 < 12; fi35++ {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v1 = v10
                                    default:
                                        v36 := GenS5{F0: "", F1: 42}
                                        _ = v36
                                    }
                                    _ = fi35
                                }
                            }
                            fmt.Println(v22, v9)
                        }
                        for fi37 := 0; fi37 < 6; fi37++ {
                            if (!v5) {
                                v38 := GenS1{F0: 0}
                                _ = v38
                            }
                            v39 := []int{fi24, (fi24 & fi37), (fi37 >> ((248) & 63))}
                            _ = fi37
                            _ = v39
                        }
                    }
                }
            }
            v12 = 1
            _ = fi24
        }
    }
    fmt.Println(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
