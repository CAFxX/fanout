package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 int; F1 string }

func main() {
    var v1 uint16 = 0
    var v2 int32 = -2147483648
    var v3 bool = true
    var v4 uint = 472
    v5 := []int{127, 128}
    _ = len(v5)
    v4 = v4
    var v6 uint64 = 42
    v7 := GenS1{F0: 9223372036854775807, F1: 0}
    switch "foo bar" {
    case "a":
        if true {
            fmt.Println(v2)
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v5 = append(v5, 127)
        }
    case "foo bar":
        if (2 != 967) {
            if (!(((3.14159 == 1e10) || v3) || v3)) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                switch -1 {
                case 137:
                    v3 = v3
                case 638:
                    v8 := GenS2{F0: 1, F1: v2, F2: 376}
                    _ = v8
                case 32768:
                    for fi9 := 0; fi9 < 4; fi9++ {
                        for fi10 := 0; fi10 < 8; fi10++ {
                            var v11 int32 = v2
                            v2 = (v11 + int32(v4))
                            fmt.Println(v11)
                            _ = fi10
                            _ = v11
                        }
                        var v12 float32 = 0.0
                        if (18446744073709551615 != v6) {
                            v4 += (v4 & v4)
                        }
                        _ = fi9
                        _ = v12
                    }
                default:
                    v3 = (((42 != uint8(v4)) && true) || (!v3))
                }
            } else {
                v3 = (!(!((^(v1)) > 370)))
                v6 += (v6 & (v6 / ((2) | 1)))
            }
        } else {
            v2 = v2
        }
    default:
        v3 = true
    }
    v4 -= v4
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 int16 = 0
        var v15 int8 = int8(v1)
        v4 = (v4 % ((v4) | 1))
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v1, v3, v3)
    var v16 uint64 = (v6 + (^(v6)))
    for fi17 := 0; fi17 < 10; fi17++ {
        v4 = (v4 % ((v4) | 1))
        for fi18 := 0; fi18 < 18; fi18++ {
            if v3 {
                switch 42 {
                case -164:
                    if false {
                        if len(v5) > 0 {
                            v5[0] = (fi18 | 793)
                        }
                    }
                case 2147483647:
                    _ = len(v5)
                }
                v3 = v3
                v19 := GenS3{F0: 1, F1: -109, F2: (v4 % ((2) | 1))}
                switch fi18 {
                case -637:
                    _ = len(v5)
                default:
                    if true {
                        v5 = append(v5, (-(fi18)))
                    } else {
                        var v20 int8 = 1
                        v21 := GenS4{F0: 4294967295}
                        v4 -= v4
                        _ = v20
                        _ = v21
                    }
                }
                if (-1 != int16(v1)) {
                    v22 := GenS5{F0: 255, F1: -260}
                    switch fi18 {
                    case 32768:
                        v23 := GenS2{F0: 1, F1: v2, F2: v1}
                        _ = v23
                    }
                    var v24 bool = true
                    _ = v22
                    _ = v24
                }
                v25 := GenS6{F0: (fi18 >> ((v4) & 63)), F1: ""}
                fmt.Println(fi17, v1, fi18)
                var v26 int16 = 0
                _ = v19
                _ = v25
                _ = v26
            } else {
                fmt.Println(v3, v1, v3)
            }
            _ = fi18
        }
        _ = fi17
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
