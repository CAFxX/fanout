package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() uint64 {
    var v7 uint64 = 4294967295
    var v8 uint = 2
    var v9 uint32 = 983
    v10 := GenS1{F0: 680}
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := []int{(fi11 % (((fi11 + fi11)) | 1)), fi11}
        v8 = (v8 << (((v8 / ((v8) | 1))) & 63))
        _ = fi11
        _ = v12
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (v7 % (((v7 & v7)) | 1))
}

func main() {
    var v1 uint = 65535
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    var v5 uint64 = 1
    var v6 uint8 = 2
    v1 = v1
    delete(v3, 334)
    for fi13 := 0; fi13 < 2; fi13++ {
        if true {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 float32 = (float32((float64(-2.25) - 3.14159)) * -2.25)
                _ = fi14
                _ = v15
            }
        }
        _ = fi13
    }
    switch 2147483647 {
    case 397:
        switch "a" {
        case "\n\t":
            v5 = (v5 & v5)
        case "a":
            var v16 string = "a"
            _ = v16
        default:
            v6 = (v6 + v6)
        }
    default:
        var v17 string = "\n\t"
        _ = v17
    }
    switch 0 {
    case -9223372036854775808:
        var v18 uint16 = 30
        _ = v18
    case -991:
        v19 := []int{-666, 128}
        _ = v19
    }
    v6 = v6
    if true {
        var v20 int64 = 0
        for fi21 := 0; fi21 < 20; fi21++ {
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 int8 = 100
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
        delete(v2, 9223372036854775807)
        _ = v20
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        if false {
            if (!(-2.25 <= (1e10 + 1.5))) {
                v6 = (^(v6))
                switch (^(fi24)) {
                case -2147483648:
                    var v25 int8 = 66
                    _ = v25
                case -750:
                    if false {
                        if ((false || true) && true) {
                            for fi26 := 0; fi26 < 19; fi26++ {
                                var v27 int16 = 821
                                var v28 int32 = 42
                                var v29 uint64 = (^(v5))
                                _ = fi26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            var v30 uint = v1
                            fmt.Println(v1, v6)
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 19; fi31++ {
                            var v32 int32 = 2147483647
                            _ = fi31
                            _ = v32
                        }
                        var v33 int32 = -844
                        _ = v33
                    } else {
                        switch 0 {
                        case 239:
                            fmt.Println(v6)
                        case -338:
                            for fi34 := 0; fi34 < 12; fi34++ {
                                v35 := v2[(fi24 >> ((42) & 63))]
                                switch fi34 {
                                case -50:
                                    v36 := []int{fi34, -9223372036854775808, fi24}
                                    _ = v36
                                case -9:
                                    switch (fi34 - fi24) {
                                    case -905:
                                        v6 += (v6 << (((v1 << ((2) & 63))) & 63))
                                    case 0:
                                        v37 := []int{}
                                        _ = v37
                                    default:
                                        var v38 bool = true
                                        _ = v38
                                    }
                                case -383:
                                    if (354 == 680) {
                                        v39 := GenS1{F0: (fi24 & (-(fi34)))}
                                        v40 := []int{fi24}
                                        var v41 int = (fi24 * fi34)
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    } else {
                                        fmt.Println(v1, fi24)
                                        if false {
                                            v42 := []int{fi34, 128, fi34, (fi24 - fi24)}
                                            _ = v42
                                        }
                                    }
                                default:
                                    _ = len(v3)
                                }
                                _ = fi34
                                _ = v35
                            }
                        case 157:
                            var v43 uint32 = 2147483647
                            _ = v43
                        default:
                            for fi44 := 0; fi44 < 6; fi44++ {
                                fmt.Println(v1, fi24)
                                _ = fi44
                            }
                        }
                        fmt.Println(v6, v1, fi24)
                    }
                }
            }
        }
        _ = fi24
    }
    if true {
        var v45 int8 = 1
        for fi46 := 0; fi46 < 20; fi46++ {
            v47 := v3[42]
            var v48 string = "a"
            v49 := []int{-1, 0, (^(fi46)), (fi46 / ((fi46) | 1))}
            _ = fi46
            _ = v47
            _ = v48
            _ = v49
        }
        var v50 int16 = 1000
        _ = v45
        _ = v50
    }
    fmt.Println(v1, v1)
    fmt.Println(v5, v5, v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
