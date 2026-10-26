package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int8; F1 int64; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint64; F2 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 1000
    var v3 uint8 = 255
    v4 := make(map[uint32]int)
    var v5 uint64 = 0
    var v6 int = 986
    switch ("hello" + "foo bar") {
    case "\n\t":
        var v7 string = "foo bar"
        _ = v7
    case "":
        var v8 int32 = 0
        _ = v8
    case "a":
        var v9 uint16 = 1
        _ = v9
    default:
        if (!true) {
            for fi10 := 0; fi10 < 12; fi10++ {
                _ = len(v4)
                if true {
                    for fi11 := 0; fi11 < 19; fi11++ {
                        v5 = v5
                        var v12 int16 = -608
                        v12 += 1000
                        switch (fi10 >> ((255) & 63)) {
                        case 32768:
                            if true {
                                fmt.Println(v6, fi10)
                            } else {
                                var v13 uint = 18446744073709551615
                                if (!(1 > -1)) {
                                    fmt.Println(v3, v13, v12)
                                    switch v6 {
                                    case -502:
                                        v14 := GenS1{F0: fi11, F1: int(v3)}
                                        _ = v14
                                    default:
                                        v12 = v2
                                    }
                                } else {
                                    fmt.Println(v13, v13)
                                    v12 = 32767
                                }
                                _ = v13
                            }
                        }
                        _ = fi11
                        _ = v12
                    }
                    v2 -= (v2 - v2)
                } else {
                    v2 = -1
                }
                _ = fi10
            }
            fmt.Println(v3, v2, v5)
        } else {
            if (2 != uint16(v3)) {
                switch 667 {
                case 86:
                    v15 := []int{(v6 / (((v6 / ((v6) | 1))) | 1)), 32768, (v6 + (-(v6))), 4294967295}
                    _ = v15
                default:
                    v3 = v3
                }
            }
            fmt.Println(v3)
        }
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        if (true && (false || true)) {
            for fi17 := 0; fi17 < 7; fi17++ {
                if true {
                    v18 := []int{fi17, (fi16 ^ (v6 | v6)), fi16, (fi17 / ((-667) | 1))}
                    _ = v18
                } else {
                    v6 = fi16
                    var v19 uint16 = uint16(v3)
                    _ = v19
                }
                _ = fi17
            }
        } else {
            v4[464] = (v6 << ((4294967295) & 63))
            if false {
                v20 := GenS1{F0: -237, F1: 42}
                fmt.Println(v5)
                _ = v20
            }
            v3 = (v3 + 2)
        }
        var v21 uint64 = v5
        var v22 string = "foo bar"
        _ = fi16
        _ = v21
        _ = v22
    }
    v23 := []int{(v6 ^ v6), (v6 & v6)}
    var v24 bool = false
    v25 := GenS2{F0: 0, F1: 982, F2: v5}
    if false {
        var v26 uint8 = (v3 + (v3 >> ((42) & 63)))
        _ = v26
    } else {
        v3 = (^(v3))
        v5 -= (v5 | (v5 - uint64(v3)))
        fmt.Println(v24, v3)
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        v6 = v6
        v3 = v3
        _ = fi27
    }
    v24 = v24
    fmt.Println(v5, v24, v5)
    v23 = append(v23, int(v3))
    if (!(((0.1 * 1.5) - (-2.25 / 1.0)) < 1.5)) {
        v24 = ((v6 | v6) <= v6)
    }
    v28 := GenS2{F0: -1, F1: 42, F2: uint64(v3)}
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v3 += (v3 & (v3 / ((1) | 1)))
    default:
        switch "a" {
        case "a":
            fmt.Println(v2, v3, v6)
        case "":
            v29 := GenS3{F0: -2147483648, F1: v5, F2: (!v24)}
            _ = v29
        case "hello":
            var v30 float32 = 3.14159
            _ = v30
        }
    }
    var v31 int16 = 1
    if (9223372036854775807 >= -160) {
        v6 = (v6 % ((v6) | 1))
    } else {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        v31 = v2
    }
    switch v6 {
    case 402:
        for fi32 := 0; fi32 < 7; fi32++ {
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            if len(v23) > 0 {
                v23[0] = (v6 & (v6 ^ (v6 | v6)))
            }
            _ = fi32
        }
    default:
        v24 = true
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
