package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 uint8; F2 int16 }
func genFn0(p0 bool) bool {
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int8 = 65
        switch fi4 {
        case 102:
            v6 := []int{}
            _ = v6
        case 42:
            v5 = v5
        case -701:
            fmt.Println(fi4, fi4, fi4)
        }
        for fi7 := 0; fi7 < 12; fi7++ {
            if false {
                v8 := GenS1{F0: 255, F1: 0, F2: -343}
                switch "a" {
                case "\n\t":
                    if p0 {
                        p0 = (9223372036854775807 != 0)
                        fmt.Println(fi7, fi7, fi4)
                    }
                case "foo bar":
                    var v9 bool = p0
                    _ = v9
                case "":
                    if (!p0) {
                        if (0 == 9223372036854775807) {
                            v10 := GenS1{F0: 402, F1: 1, F2: 1}
                            var v11 uint = 754
                            v12 := []int{818, 42, (fi4 | fi7), (-(fi4))}
                            _ = v10
                            _ = v11
                            _ = v12
                        } else {
                            var v13 int16 = 372
                            var v14 uint64 = 18446744073709551615
                            _ = v13
                            _ = v14
                        }
                    } else {
                        var v15 uint8 = 1
                        switch fi4 {
                        case 2147483647:
                            v5 += (^(v5))
                        default:
                            var v16 uint64 = 4294967295
                            _ = v16
                        }
                        _ = v15
                    }
                default:
                    if false {
                        fmt.Println(fi4)
                        v17 := []int{fi7, 9223372036854775807, fi4}
                        fmt.Println(fi4)
                        _ = v17
                    } else {
                        var v18 int8 = (v5 / ((-1) | 1))
                        v5 = v5
                        _ = v18
                    }
                }
                _ = v8
            } else {
                v19 := GenS2{F0: -1}
                v20 := []int{-1, fi4, fi4, (fi4 / ((fi4) | 1))}
                var v21 uint = 18446744073709551615
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                _ = v19
                _ = v20
                _ = v21
            }
            var v22 bool = true
            var v23 int16 = 42
            _ = fi7
            _ = v22
            _ = v23
        }
        _ = fi4
        _ = v5
    }
    v24 := []int{20, -9223372036854775808, 9223372036854775807}
    _ = p0
    _ = v24
    return p0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 9223372036854775807
    v3 := []string{}
    v25 := GenS3{F0: 42}
    fmt.Println(v2)
    for fi26 := 0; fi26 < 11; fi26++ {
        if false {
            var v27 uint8 = 67
            delete(v1, ("hello" + "a"))
            _ = v27
        }
        _ = fi26
    }
    v2 -= -140
    var v28 int32 = 1
    v29 := GenS3{F0: 1}
    if ((v2 & (v2 | (v2 - 733))) <= (v2 ^ v2)) {
        v28 = (v28 & (v28 % ((0) | 1)))
        if len(v3) > 0 {
            v3[0] = ("foo bar" + "\n\t")
        }
        for fi30 := 0; fi30 < 11; fi30++ {
            v28 += -1
            _ = fi30
        }
    } else {
        v31 := GenS4{F0: -981}
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        v28 -= (^(v28))
        _ = v31
    }
    if false {
        var v32 uint64 = 18446744073709551615
        for fi33 := 0; fi33 < 16; fi33++ {
            for fi34 := 0; fi34 < 6; fi34++ {
                v2 = fi33
                _ = fi34
            }
            v35 := GenS3{F0: 14}
            var v36 uint16 = 255
            _ = fi33
            _ = v35
            _ = v36
        }
        _ = v32
    } else {
        for fi37 := 0; fi37 < 6; fi37++ {
            v3 = append(v3, (("foo bar" + "foo bar") + "a"))
            _ = fi37
        }
        v38 := GenS5{F0: 1, F1: 255, F2: 1}
        v2 += v2
        _ = v38
    }
    var v39 uint64 = 203
    v28 -= v28
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v39
}
