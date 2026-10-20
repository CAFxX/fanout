package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 uint64 = 2
    var v2 int16 = -32768
    var v3 string = "a"
    var v4 int8 = 0
    var v5 uint32 = 410
    var v6 int64 = -777
    var v7 uint16 = 2
    v4 -= int8(v5)
    if true {
        var v8 uint16 = (v7 & (v7 << ((373) & 63)))
        var v9 string = ("" + v3)
        if ((-(v4)) > (v4 | 42)) {
            v3 = v9
            v1 = (v1 << ((uint(v1)) & 63))
            var v10 uint = 0
            _ = v10
        } else {
            v7 += (v8 % ((771) | 1))
            if (!true) {
                v11 := GenS1{F0: 4294967295, F1: v6}
                v12 := []int{352, 32767, 9223372036854775807, -239}
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 uint8 = 198
                    var v15 int16 = (v2 >> ((uint(v14)) & 63))
                    v16 := GenS2{F0: v15, F1: (-(fi13))}
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            }
        }
        _ = v8
        _ = v9
    } else {
        v17 := GenS3{F0: 32}
        _ = v17
    }
    v18 := GenS4{F0: -1, F1: (v3 + ("\n\t" + "\n\t"))}
    v3 = ""
    v4 += (v4 & v4)
    fmt.Println(v3, v5)
    if (v1 < 183) {
        if true {
            var v19 float64 = 1e10
            _ = v19
        }
        for fi20 := 0; fi20 < 8; fi20++ {
            if (!(!true)) {
                if false {
                    if (42 >= 42) {
                        var v21 uint64 = (v1 & v1)
                        _ = v21
                    } else {
                        if (true || false) {
                            var v22 uint32 = v5
                            v23 := GenS5{F0: (v7 - (v7 * v7))}
                            var v24 uint8 = 86
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        var v25 string = v3
                        _ = v25
                    }
                    if (254 == 1) {
                        v4 -= v4
                    } else {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            var v27 float64 = 0.1
                            v2 += (v2 * (v2 - v2))
                            _ = fi26
                            _ = v27
                        }
                    }
                } else {
                    for fi28 := 0; fi28 < 10; fi28++ {
                        switch v3 {
                        case "hello":
                            v4 = (v4 | v4)
                            if (v7 < v7) {
                                var v29 int32 = -2147483648
                                var v30 int16 = (v2 % ((v2) | 1))
                                v31 := GenS1{F0: 42, F1: (v6 ^ (v6 << ((65535) & 63)))}
                                _ = v29
                                _ = v30
                                _ = v31
                            } else {
                                var v32 uint32 = 42
                                v33 := []int{-750, (fi28 << ((956) & 63))}
                                if len(v33) >= 2 {
                                    v33 = v33[1:len(v33)]
                                }
                                _ = v32
                                _ = v33
                            }
                        case "":
                            v34 := []int{}
                            _ = v34
                        case "foo bar":
                            fmt.Println(v3)
                        }
                        switch (fi20 ^ (fi28 | 390)) {
                        case 295:
                            fmt.Println(v4, v3)
                        default:
                            v6 = (v6 * 9223372036854775807)
                        }
                        _ = fi28
                    }
                }
                switch ("a" + (v3 + v3)) {
                case "\n\t":
                    v7 = v7
                case "hello":
                    var v35 uint16 = v7
                    _ = v35
                case "foo bar":
                    v36 := []int{fi20, (fi20 / ((-57) | 1)), (fi20 ^ (fi20 / (((fi20 + 564)) | 1))), 127}
                    _ = v36
                }
            }
            if true {
                v6 -= (^(v6))
            } else {
                var v37 uint32 = (v5 + (^(v5)))
                var v38 uint16 = 2
                _ = v37
                _ = v38
            }
            _ = fi20
        }
        v39 := []int{2147483647}
        _ = v39
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        v41 := GenS3{F0: 59}
        for fi42 := 0; fi42 < 3; fi42++ {
            v43 := GenS5{F0: 164}
            _ = fi42
            _ = v43
        }
        _ = fi40
        _ = v41
    }
    fmt.Println(v5, v4)
    v1 = (v1 & 319)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
