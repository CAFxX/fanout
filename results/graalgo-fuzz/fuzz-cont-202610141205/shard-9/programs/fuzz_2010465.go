package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 uint32 = 538
    var v2 float32 = 1e10
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    var v5 int64 = 42
    var v6 uint = 65535
    v4[v1] = 65535
    v7 := GenS1{F0: 65535, F1: 279}
    fmt.Println(v2, v1, v5)
    switch "a" {
    case "a":
        for fi8 := 0; fi8 < 8; fi8++ {
            v6 = uint(v1)
            for fi9 := 0; fi9 < 16; fi9++ {
                if (!true) {
                    v5 -= (^(v5))
                }
                var v10 uint8 = 2
                _ = fi9
                _ = v10
            }
            if (-1 >= 2147483647) {
                var v11 string = "foo bar"
                for fi12 := 0; fi12 < 12; fi12++ {
                    var v13 int = (-(fi8))
                    v2 += -2.25
                    _ = fi12
                    _ = v13
                }
                _ = v11
            } else {
                var v14 uint32 = (v1 << (((v6 * 18446744073709551615)) & 63))
                _ = v14
            }
            _ = fi8
        }
    case "foo bar":
        if true {
            if false {
                if false {
                    fmt.Println(v1, v6, v5)
                    for fi15 := 0; fi15 < 8; fi15++ {
                        var v16 bool = (v6 == v6)
                        if false {
                            var v17 int64 = (v5 ^ v5)
                            fmt.Println(v16, v17)
                            fmt.Println(v5, v6)
                            _ = v17
                        }
                        var v18 int16 = int16(v5)
                        _ = fi15
                        _ = v16
                        _ = v18
                    }
                    switch "" {
                    case "":
                        var v19 bool = (!true)
                        _ = v19
                    default:
                        v2 = float32(0.1)
                    }
                }
                if false {
                    var v20 int16 = 234
                    v4[v1] = (v1 / ((583) | 1))
                    _ = v20
                }
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                v22 := []int{661, (fi21 << ((v6) & 63)), fi21, int(v1)}
                var v23 int8 = 42
                var v24 uint = v6
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            v1 -= v1
        }
    case "\n\t":
        var v25 int16 = 1000
        _ = v25
    default:
        if (-37 != 106) {
            for fi26 := 0; fi26 < 3; fi26++ {
                v1 += (^(v1))
                _ = fi26
            }
            v27 := GenS2{F0: 0, F1: 255}
            _ = v27
        } else {
            for fi28 := 0; fi28 < 3; fi28++ {
                switch (^(fi28)) {
                case 125:
                    v29 := GenS3{F0: "hello"}
                    _ = v29
                }
                var v30 uint = 18446744073709551615
                v6 += v30
                _ = fi28
                _ = v30
            }
            for fi31 := 0; fi31 < 18; fi31++ {
                fmt.Println(v5)
                _ = fi31
            }
            var v32 uint = v6
            _ = v32
        }
    }
    v33 := GenS2{F0: -46, F1: 185}
    v34 := []int{42, 0, -1}
    for fi35 := 0; fi35 < 4; fi35++ {
        if (v6 == v6) {
            v6 += (v6 * 42)
        } else {
            if false {
                var v36 uint64 = 1
                v5 += (^(v5))
                var v37 int64 = v5
                _ = v36
                _ = v37
            } else {
                v6 = v6
                switch -9223372036854775808 {
                case 2147483647:
                    for fi38 := 0; fi38 < 16; fi38++ {
                        v39 := GenS4{F0: v6, F1: 63}
                        _ = fi38
                        _ = v39
                    }
                default:
                    fmt.Println(v5, v2, v1)
                }
            }
        }
        if (227 > 269) {
            v2 = v2
        } else {
            for fi40 := 0; fi40 < 18; fi40++ {
                if len(v34) >= 2 {
                    v34 = v34[1:len(v34)]
                }
                _ = fi40
            }
        }
        _ = fi35
    }
    var v41 string = ""
    v3[("a" + ("a" + ("" + v41)))] = (((v1 - (v1 & 2)) != v1) && (!(9223372036854775807 < (v5 * (^(v5))))))
    var v42 int8 = -1
    v43 := v3[v41]
    if (v43 || ((!v43) || v43)) {
        var v44 bool = (127 <= int(v1))
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
        _ = v44
    } else {
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
        v45 := GenS5{F0: (v1 >> (((^(v6))) & 63)), F1: 4294967295}
        for fi46 := 0; fi46 < 19; fi46++ {
            v2 = 1e10
            v47 := GenS2{F0: 65, F1: uint8(v6)}
            _ = fi46
            _ = v47
        }
        _ = v45
    }
    var v48 uint16 = 624
    v49 := GenS5{F0: 2147483647, F1: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
