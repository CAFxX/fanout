package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }

func main() {
    var v1 uint8 = 0
    var v2 uint64 = 865
    var v3 uint16 = 255
    v4 := make(map[uint32]uint32)
    var v5 int = -957
    var v6 int32 = -2147483648
    v7 := []int{v5, v5}
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v5, v5)
        _ = fi8
    }
    v9 := v4[400]
    if true {
        switch "\n\t" {
        case "a":
            var v10 bool = false
            _ = v10
        }
        for fi11 := 0; fi11 < 12; fi11++ {
            switch ("\n\t" + "\n\t") {
            case "a":
                if len(v7) > 0 {
                    v7[0] = v5
                }
            case "hello":
                if true {
                    switch 0 {
                    case 2147483647:
                        for fi12 := 0; fi12 < 14; fi12++ {
                            switch (v5 | (^(fi11))) {
                            case -9223372036854775808:
                                var v13 uint64 = (v2 << ((609) & 63))
                                _ = v13
                            case 915:
                                if ((fi12 >> ((862) & 63)) > fi12) {
                                    v4[(v9 / ((930) | 1))] = (v9 % ((376) | 1))
                                    fmt.Println(v2, v2)
                                    v7 = append(v7, (-(fi12)))
                                } else {
                                    v14 := v4[v9]
                                    if true {
                                        v15 := GenS1{F0: v9, F1: v1}
                                        v5 += -559
                                        _ = v15
                                    }
                                    v16 := GenS1{F0: v14, F1: v1}
                                    _ = v14
                                    _ = v16
                                }
                            }
                            v17 := GenS2{F0: v5, F1: 1}
                            delete(v4, v9)
                            _ = fi12
                            _ = v17
                        }
                    case 4294967295:
                        v2 = (v2 * (v2 * (v2 & (v2 + (v2 * (v2 / ((829) | 1)))))))
                    case -113:
                        var v18 uint32 = v9
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 13; fi19++ {
                            v20 := GenS1{F0: 2147483647, F1: uint8(v3)}
                            v21 := GenS3{F0: -84}
                            fmt.Println(v2, v2)
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                    var v22 bool = (-665 > 0)
                    _ = v22
                }
            }
            _ = fi11
        }
    }
    var v23 uint8 = (v1 | v1)
    var v24 string = "foo bar"
    for fi25 := 0; fi25 < 16; fi25++ {
        if (!true) {
            if true {
                var v26 float32 = -2.25
                switch (((v24 + (v24 + v24)) + v24) + "\n\t") {
                case "\n\t":
                    switch (fi25 >> ((183) & 63)) {
                    case 4294967295:
                        _ = len(v4)
                    case 32767:
                        _ = len(v4)
                    default:
                        var v27 uint8 = (v23 >> ((745) & 63))
                        _ = v27
                    }
                default:
                    var v28 uint8 = v1
                    _ = v28
                }
                _ = v26
            }
        }
        fmt.Println(fi25, v9)
        _ = fi25
    }
    v29 := GenS4{F0: v5}
    v30 := GenS4{F0: (-(v5))}
    v31 := v4[uint32(v23)]
    if false {
        v4[(v31 | v9)] = v31
        v32 := GenS2{F0: 754, F1: (v2 + v2)}
        fmt.Println(v24, v3, v3)
        _ = v32
    }
    if (((true && (451 >= 42)) && true) || (-63 >= 42)) {
        var v33 uint8 = v23
        _ = v33
    } else {
        var v34 float64 = 0.0
        var v35 uint16 = 531
        v4[(v9 * (v31 % ((v31) | 1)))] = 42
        _ = v34
        _ = v35
    }
    if (!(1 == 233)) {
        var v36 int = v5
        for fi37 := 0; fi37 < 6; fi37++ {
            v5 -= int(v23)
            _ = fi37
        }
        v24 = (("\n\t" + "\n\t") + v24)
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
