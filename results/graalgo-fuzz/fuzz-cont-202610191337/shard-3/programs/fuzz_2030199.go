package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint16; F2 bool }
type GenS5 struct { F0 int8; F1 int16 }

func main() {
    var v1 uint64 = 42
    var v2 uint = 336
    var v3 string = ("\n\t" + ("foo bar" + ("a" + ((("\n\t" + "") + "\n\t") + "foo bar"))))
    var v4 int = -56
    var v5 string = (v3 + "\n\t")
    var v6 uint8 = 42
    switch v3 {
    case "a":
        v7 := []int{128, v4, (v4 << ((42) & 63)), 4294967295}
        _ = v7
    }
    v4 = (v4 / (((v4 / ((1) | 1))) | 1))
    switch v5 {
    case "hello":
        switch 32767 {
        case -778:
            v8 := []int{329, (v4 | (v4 >> ((v2) & 63))), (v4 >> (((v2 ^ v2)) & 63))}
            _ = v8
        default:
            switch v3 {
            case "hello":
                v9 := []int{(v4 - v4), v4, 127}
                _ = v9
            case "\n\t":
                switch 127 {
                case 210:
                    v6 += 2
                }
            case "foo bar":
                fmt.Println(v6, v6, v5)
            default:
                switch "" {
                case "a":
                    for fi10 := 0; fi10 < 19; fi10++ {
                        v11 := []int{v4, 32767, 4294967295}
                        v1 = v1
                        _ = fi10
                        _ = v11
                    }
                case "foo bar":
                    v12 := GenS1{F0: 0}
                    _ = v12
                case "hello":
                    switch "foo bar" {
                    case "foo bar":
                        v13 := []int{(v4 >> (((^(v2))) & 63)), (v4 & (v4 - v4)), v4, 0}
                        _ = v13
                    case "a":
                        var v14 int32 = -2147483648
                        _ = v14
                    case "":
                        v2 -= (v2 << (((^(v2))) & 63))
                    default:
                        v15 := []int{v4, (^(v4))}
                        _ = v15
                    }
                default:
                    v16 := []int{-9223372036854775808, v4}
                    _ = v16
                }
            }
        }
    case "":
        v17 := GenS1{F0: 2147483648}
        _ = v17
    case "\n\t":
        var v18 int8 = 0
        _ = v18
    default:
        if (false || true) {
            for fi19 := 0; fi19 < 17; fi19++ {
                v20 := GenS2{F0: v1}
                var v21 int = v4
                for fi22 := 0; fi22 < 20; fi22++ {
                    var v23 bool = false
                    var v24 int8 = -69
                    v25 := GenS3{F0: (^(v6))}
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        } else {
            v26 := GenS3{F0: 0}
            _ = v26
        }
    }
    v27 := GenS2{F0: (^(v1))}
    v28 := GenS4{F0: 1, F1: 0, F2: (!true)}
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 float32 = 1.5
        v31 := GenS3{F0: 128}
        _ = fi29
        _ = v30
        _ = v31
    }
    v32 := []int{4294967295, -462, v4}
    if (v1 >= (v1 & (v1 * (v1 * (^(v1)))))) {
        var v33 uint64 = (v1 | 2)
        switch v4 {
        case 127:
            switch (("foo bar" + (v5 + (v3 + (((v5 + v5) + (v3 + v3)) + v5)))) + (v5 + ("" + v5))) {
            case "\n\t":
                v2 = 65535
            case "hello":
                var v34 int16 = 1000
                _ = v34
            }
        case 1:
            if true {
                v5 = v5
                v4 += -2147483648
            } else {
                v35 := GenS1{F0: 353}
                switch v5 {
                case "a":
                    v36 := GenS5{F0: 0, F1: -378}
                    _ = v36
                case "":
                    var v37 uint64 = (v33 * 133)
                    _ = v37
                case "foo bar":
                    v2 = uint(v6)
                default:
                    v38 := GenS4{F0: 0, F1: 0, F2: true}
                    _ = v38
                }
                if len(v32) >= 2 {
                    v32 = v32[1:len(v32)]
                }
                _ = v35
            }
        case 42:
            var v39 int = 32767
            _ = v39
        default:
            v32 = append(v32, (v4 & v4))
        }
        _ = v33
    }
    if true {
        if false {
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
            switch 9223372036854775807 {
            case -607:
                for fi40 := 0; fi40 < 9; fi40++ {
                    v1 = 4294967295
                    _ = fi40
                }
            }
        } else {
            v32 = append(v32, int(v2))
            if len(v32) > 0 {
                v32[0] = -1
            }
            for fi41 := 0; fi41 < 16; fi41++ {
                var v42 uint8 = 33
                var v43 uint8 = v42
                _ = fi41
                _ = v42
                _ = v43
            }
        }
        v2 = 267
    }
    var v44 float32 = 1e10
    for fi45 := 0; fi45 < 11; fi45++ {
        v32 = append(v32, v4)
        _ = fi45
    }
    for fi46 := 0; fi46 < 3; fi46++ {
        if true {
            v4 -= 127
        } else {
            var v47 int = (^(v4))
            _ = v47
        }
        fmt.Println(v2, v5, v6)
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
