package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint8) string {
    v7 := GenS1{F0: 847, F1: 2147483648}
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS2{F0: -1}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 string = ("" + "\n\t")
        if (100 != 1) {
            v12 := []int{(fi10 << ((18446744073709551615) & 63)), (fi10 / ((1) | 1)), fi10, fi10}
            var v13 float64 = 1.5
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    p0 = p0
    v14 := GenS2{F0: -821}
    _ = p0
    _ = v14
    _ = v7
    return ""
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    v3 := []uint32{}
    v4 := make(map[uint32]string)
    v5 := []string{}
    v6 := make(map[int]int)
    if v1 {
        for fi15 := 0; fi15 < 13; fi15++ {
            switch "foo bar" {
            case "hello":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "foo bar":
                if true {
                    switch 2147483647 {
                    case 790:
                        for fi16 := 0; fi16 < 20; fi16++ {
                            _ = len(v2)
                            v1 = ((-1 != 0) && v1)
                            fmt.Println(v1)
                            _ = fi16
                        }
                    default:
                        var v17 uint = 229
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 6; fi18++ {
                        if (0 != -75) {
                            if v1 {
                                v1 = (fi18 >= fi18)
                                var v19 int32 = -2147483648
                                v19 -= v19
                                _ = v19
                            }
                            fmt.Println(fi18)
                            if true {
                                _ = len(v3)
                                var v20 bool = (1 == fi18)
                                if false {
                                    v21 := GenS1{F0: 32767, F1: 2147483648}
                                    _ = v21
                                }
                                _ = v20
                            } else {
                                _ = len(v3)
                                if len(v3) > 0 {
                                    v3[0] = 2147483648
                                }
                                v1 = (v1 && v1)
                                v22 := GenS2{F0: -32768}
                                _ = v22
                            }
                        }
                        v1 = (!false)
                        _ = fi18
                    }
                    if (((-718 == -712) || v1) && (0 < 0)) {
                        v1 = false
                        var v23 uint32 = 1
                        _ = v23
                    }
                }
            case "\n\t":
                v1 = (!v1)
            default:
                v1 = ((true && true) || v1)
            }
            v24 := GenS2{F0: 439}
            _ = fi15
            _ = v24
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v1 = (1 == 128)
    if len(v3) > 0 {
        v3[0] = 1
    }
    if ("a" != "foo bar") {
        for fi25 := 0; fi25 < 3; fi25++ {
            v26 := GenS3{F0: -32}
            switch 128 {
            case 127:
                v1 = (100 == 99)
            case 9223372036854775807:
                v1 = (128 >= 1)
            }
            _ = fi25
            _ = v26
        }
        delete(v4, 2147483648)
    } else {
        v27 := v2[32768]
        _ = v27
    }
    v28 := GenS1{F0: 1, F1: 65535}
    var v29 uint8 = 1
    v29 += (v29 - (v29 << ((1) & 63)))
    var v30 int = int(v29)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v31 := GenS1{F0: 954, F1: 534}
    var v32 int8 = 24
    v29 -= v29
    v33 := GenS1{F0: 412, F1: 1}
    var v34 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
