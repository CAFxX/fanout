package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int; F2 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 int; F2 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "hello"
    var v3 uint8 = 175
    var v4 uint8 = (v3 >> ((42) & 63))
    var v5 float32 = (float32(-2.25) / 1.5)
    v6 := make(map[string]uint32)
    switch 42 {
    case -207:
        v3 -= (v4 - v3)
    case 104:
        if false {
            v7 := []int{2147483647, 763}
            _ = v7
        } else {
            switch -9223372036854775808 {
            case 81:
                v5 -= -2.25
            case -1:
                v2 = (v2 + (v2 + ""))
            case 9223372036854775807:
                fmt.Println(v5)
            }
        }
    default:
        if false {
            switch (v2 + (v2 + "\n\t")) {
            case "":
                for fi8 := 0; fi8 < 12; fi8++ {
                    var v9 int = fi8
                    if true {
                        var v10 int = v9
                        switch v2 {
                        case "a":
                            for fi11 := 0; fi11 < 17; fi11++ {
                                var v12 string = "\n\t"
                                _ = fi11
                                _ = v12
                            }
                        }
                        var v13 string = v2
                        _ = v10
                        _ = v13
                    } else {
                        v9 = fi8
                    }
                    _ = fi8
                    _ = v9
                }
            case "foo bar":
                v2 = ("" + "")
            case "a":
                if (-2147483648 != 318) {
                    var v14 int8 = 100
                    for fi15 := 0; fi15 < 6; fi15++ {
                        v4 = 3
                        var v16 float64 = (0.0 + -2.25)
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                } else {
                    var v17 int8 = 55
                    _ = v17
                }
            default:
                if true {
                    if true {
                        var v18 int8 = -37
                        fmt.Println(v18, v5, v2)
                        switch -267 {
                        case -885:
                            switch 32768 {
                            case -591:
                                var v19 float64 = ((0.0 + (((float64(v5) * -2.25) / -2.25) / -2.25)) - 1.5)
                                _ = v19
                            case 32768:
                                for fi20 := 0; fi20 < 5; fi20++ {
                                    var v21 uint32 = 734
                                    _ = fi20
                                    _ = v21
                                }
                            case 858:
                                v4 += v3
                            default:
                                fmt.Println(v3, v4, v4)
                            }
                        default:
                            fmt.Println(v4, v5, v5)
                        }
                        _ = v18
                    } else {
                        if (0 == 50) {
                            _ = len(v1)
                            var v22 int32 = -2147483648
                            fmt.Println(v3, v2)
                            _ = v22
                        } else {
                            var v23 int16 = -630
                            v24 := v1[215]
                            _ = v23
                            _ = v24
                        }
                    }
                }
            }
            v6["foo bar"] = 2147483648
            switch ((v2 + "hello") + (((v2 + v2) + v2) + v2)) {
            case "hello":
                if (!(32767 != 0)) {
                    var v25 int8 = 127
                    v26 := GenS1{F0: 42, F1: 0}
                    fmt.Println(v25, v3)
                    _ = v25
                    _ = v26
                } else {
                    if false {
                        for fi27 := 0; fi27 < 18; fi27++ {
                            v28 := GenS2{F0: 1, F1: fi27, F2: -9223372036854775808}
                            switch (^(fi27)) {
                            case -430:
                                v29 := GenS3{F0: (fi27 + fi27)}
                                _ = v29
                            }
                            _ = fi27
                            _ = v28
                        }
                    }
                    fmt.Println(v5, v4, v4)
                }
            case "a":
                for fi30 := 0; fi30 < 3; fi30++ {
                    switch (-(fi30)) {
                    case 863:
                        if (584 >= 0) {
                            v31 := []int{(fi30 + fi30)}
                            _ = v31
                        } else {
                            var v32 float64 = 0.1
                            v3 += (^(v3))
                            var v33 float64 = v32
                            _ = v32
                            _ = v33
                        }
                    }
                    v5 -= ((0.1 - v5) * v5)
                    _ = fi30
                }
            default:
                switch (v2 + (v2 + v2)) {
                case "foo bar":
                    v1[28] = true
                case "hello":
                    var v34 bool = false
                    _ = v34
                case "":
                    v4 -= v4
                default:
                    v2 = v2
                }
            }
        } else {
            fmt.Println(v4, v3)
            v5 -= v5
            fmt.Println(v2, v3, v4)
        }
    }
    v4 -= (v4 >> ((18446744073709551615) & 63))
    var v35 int8 = 1
    v36 := GenS4{F0: 2, F1: 32767, F2: 257}
    var v37 string = "\n\t"
    fmt.Println(v5, v37)
    fmt.Println(v4)
    fmt.Println(v2, v2, v2)
    v35 += v35
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
