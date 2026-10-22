package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int64; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint8 }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 1
    v3 := make(map[int]string)
    v4 := GenS1{F0: 42}
    v5 := GenS1{F0: 647}
    fmt.Println(v2)
    var v6 uint8 = v2
    var v7 uint16 = 240
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(v6)
        if false {
            v6 -= 2
        } else {
            if ((v7 < 542) || false) {
                var v9 uint64 = 2
                fmt.Println(v6, v6, v9)
                _ = v9
            } else {
                if (!(!(0 <= -1))) {
                    var v10 uint64 = 2
                    _ = v10
                }
            }
            v7 += 255
        }
        fmt.Println(v2, fi8)
        _ = fi8
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(v7, fi11)
        _ = fi11
    }
    var v12 uint = 902
    var v13 string = "foo bar"
    switch -29 {
    case 9223372036854775807:
        for fi14 := 0; fi14 < 10; fi14++ {
            switch (fi14 << (((v12 ^ (v12 * (v12 << ((255) & 63))))) & 63)) {
            case 2147483647:
                v6 -= (v6 - (v6 | 55))
            default:
                v7 = v7
            }
            _ = fi14
        }
    case 128:
        var v15 string = v13
        _ = v15
    case 920:
        switch -9223372036854775808 {
        case -18:
            for fi16 := 0; fi16 < 16; fi16++ {
                if false {
                    var v17 string = (("" + v13) + "\n\t")
                    _ = v17
                } else {
                    v13 = ""
                    if (!(!false)) {
                        var v18 uint64 = 18446744073709551615
                        _ = v18
                    }
                    fmt.Println(v12)
                }
                if true {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        if true {
                            if false {
                                v6 = v2
                                if true {
                                    switch fi16 {
                                    case -677:
                                        v13 = "foo bar"
                                    case 42:
                                        var v20 int16 = 0
                                        _ = v20
                                    case 127:
                                        fmt.Println(v13, fi16)
                                    default:
                                        v13 = ((v13 + v13) + v13)
                                    }
                                    switch "a" {
                                    case "a":
                                        v2 = (v6 | 1)
                                    case "":
                                        v21 := GenS2{F0: 0, F1: -207, F2: 2}
                                        _ = v21
                                    }
                                    if false {
                                        fmt.Println(fi16, v13)
                                        var v22 int = fi16
                                        _ = v22
                                    }
                                } else {
                                    _ = len(v1)
                                }
                            } else {
                                _ = len(v3)
                                var v23 float64 = -2.25
                                v2 = (v6 * v6)
                                _ = v23
                            }
                            v24 := GenS2{F0: 42, F1: -9223372036854775808, F2: uint64(v6)}
                            v2 = 2
                            _ = v24
                        }
                        v7 = v7
                        v25 := GenS2{F0: 964, F1: 42, F2: uint64(v2)}
                        _ = fi19
                        _ = v25
                    }
                    if (true || (true || true)) {
                        if true {
                            var v26 int = (fi16 >> ((209) & 63))
                            var v27 uint = 421
                            var v28 int8 = 0
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            var v29 int64 = 9223372036854775807
                            var v30 uint8 = v6
                            fmt.Println(fi16, v30)
                            _ = v29
                            _ = v30
                        }
                    } else {
                        v31 := GenS1{F0: 5}
                        if false {
                            fmt.Println(v13, v7)
                            fmt.Println(v2)
                            fmt.Println(v7, v12)
                        } else {
                            v32 := []int{fi16, (-(fi16)), (fi16 & fi16)}
                            v32 = append(v32, (fi16 / ((fi16) | 1)))
                            _ = v32
                        }
                        _ = v31
                    }
                    if (!((1 <= 142) && true)) {
                        v33 := []int{490, (fi16 & 128), (fi16 - 805)}
                        if ((true && (3.14159 != 0.0)) && false) {
                            v12 += (^(v12))
                            switch (fi16 - fi16) {
                            case 0:
                                for fi34 := 0; fi34 < 16; fi34++ {
                                    fmt.Println(v6, v2)
                                    _ = fi34
                                }
                            case -217:
                                for fi35 := 0; fi35 < 5; fi35++ {
                                    v36 := GenS2{F0: 715, F1: -440, F2: 641}
                                    fmt.Println(fi16, v2)
                                    _ = fi35
                                    _ = v36
                                }
                            }
                        } else {
                            switch v13 {
                            case "hello":
                                v37 := GenS3{F0: 0, F1: v2}
                                _ = v37
                            case "foo bar":
                                if len(v33) > 0 {
                                    v33[0] = fi16
                                }
                            default:
                                v13 = ""
                            }
                            fmt.Println(v2, v6)
                        }
                        fmt.Println(v2)
                        _ = v33
                    }
                } else {
                    fmt.Println(v12, v7, fi16)
                }
                v6 += 2
                _ = fi16
            }
        case 9223372036854775807:
            var v38 string = v13
            _ = v38
        }
    }
    var v39 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
