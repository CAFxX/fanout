package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int8; F2 uint16 }

func main() {
    v1 := []uint32{2, 437}
    var v2 uint = 522
    var v3 bool = (!(!(-815 == 842)))
    var v4 int32 = -2147483648
    var v5 uint16 = 2
    fmt.Println(v5, v5)
    v3 = (!(0 <= 870))
    v6 := GenS1{F0: "\n\t", F1: v4, F2: 746}
    fmt.Println(v3)
    v5 += 42
    v7 := GenS1{F0: "\n\t", F1: v4, F2: 32767}
    if len(v1) > 0 {
        v1[0] = 802
    }
    v4 -= v4
    v2 = (^(v2))
    if true {
        if v3 {
            var v8 uint8 = 2
            var v9 uint = (v2 % ((v2) | 1))
            _ = v8
            _ = v9
        } else {
            fmt.Println(v4, v5)
        }
        v5 = (^(v5))
        v5 = (v5 & 2)
    } else {
        switch "foo bar" {
        case "":
            if false {
                if len(v1) > 0 {
                    v1[0] = 65535
                }
                var v10 int32 = v4
                var v11 bool = ((^(v2)) >= (v2 << ((v2) & 63)))
                _ = v10
                _ = v11
            }
        case "foo bar":
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 int8 = 1
                fmt.Println(v5)
                if (!("a" == "")) {
                    v14 := GenS1{F0: "a", F1: (v4 + (v4 >> ((63) & 63))), F2: (fi12 | (fi12 << ((uint(v5)) & 63)))}
                    var v15 int8 = -118
                    if v3 {
                        fmt.Println(v15)
                        fmt.Println(v5, v5, v2)
                        var v16 uint16 = (v5 << (((v2 << ((113) & 63))) & 63))
                        _ = v16
                    }
                    _ = v14
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
        default:
            if len(v1) > 0 {
                v1[0] = 1
            }
        }
        fmt.Println(v5)
        for fi17 := 0; fi17 < 2; fi17++ {
            v4 = v4
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 uint8 = 156
                var v20 uint64 = 18446744073709551615
                v5 -= v5
                _ = fi18
                _ = v19
                _ = v20
            }
            v21 := GenS2{F0: "hello", F1: v3}
            _ = fi17
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        fmt.Println(fi22, fi22)
        _ = fi22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 uint64 = 768
            if true {
                var v26 bool = v3
                v26 = v3
                _ = v26
            } else {
                v25 -= (v25 % ((0) | 1))
                fmt.Println(v5, fi24)
            }
            _ = fi24
            _ = v25
        }
        if ((uint8(v2) > 128) || false) {
            if v3 {
                for fi27 := 0; fi27 < 16; fi27++ {
                    var v28 bool = false
                    var v29 float32 = -2.25
                    switch "\n\t" {
                    case "a":
                        var v30 uint16 = v5
                        _ = v30
                    case "foo bar":
                        v31 := GenS3{F0: v5}
                        _ = v31
                    default:
                        switch 1 {
                        case 9223372036854775807:
                            var v32 bool = (v3 && (1000 <= 1))
                            _ = v32
                        case -163:
                            switch 0 {
                            case 127:
                                v4 = (^(v4))
                            case -140:
                                fmt.Println(v28)
                            case 335:
                                v33 := GenS1{F0: ("" + ("hello" + "foo bar")), F1: (v4 ^ v4), F2: fi27}
                                _ = v33
                            default:
                                v1 = append(v1, 4294967295)
                            }
                            var v34 uint32 = 42
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v34
                        }
                    }
                    switch (^(fi23)) {
                    case -1:
                        if len(v1) > 0 {
                            v1[0] = 519
                        }
                    case 128:
                        var v35 int = 32768
                        _ = v35
                    default:
                        var v36 uint32 = uint32(v2)
                        if (2147483647 < (v4 >> (((v2 + (^(v2)))) & 63))) {
                            v37 := GenS4{F0: -1, F1: int8(fi27), F2: (^(v5))}
                            v29 += v29
                            v38 := GenS1{F0: "\n\t", F1: (v4 * v4), F2: (^(fi23))}
                            var v39 string = ""
                            _ = v37
                            _ = v38
                            _ = v39
                        } else {
                            var v40 uint = (v2 & (v2 ^ 0))
                            v2 -= (^(v2))
                            _ = v40
                        }
                        v3 = v28
                        _ = v36
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                if ((("\n\t" + "a") == "a") || v3) {
                    var v41 float32 = ((0.0 - 3.14159) * 1.5)
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    _ = v41
                }
            }
            var v42 uint32 = 1
            _ = v42
        }
        _ = fi23
    }
    var v43 int = -71
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
