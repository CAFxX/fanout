package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int) uint8 {
    var v5 int64 = -1
    v6 := []int{p0}
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 float32 = -2.25
    v2 := []int{-9223372036854775808, -2147483648}
    var v3 string = "a"
    var v4 uint32 = 2147483648
    var v7 int64 = 679
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint64 = 18446744073709551615
        switch v3 {
        case "":
            var v10 int16 = -1
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    switch "a" {
    case "":
        if (false && false) {
            v7 = -862
            fmt.Println(v4, v4, v4)
            for fi11 := 0; fi11 < 15; fi11++ {
                for fi12 := 0; fi12 < 15; fi12++ {
                    var v13 uint32 = (v4 >> ((255) & 63))
                    fmt.Println(v13)
                    switch "a" {
                    case "":
                        if (127 == 41) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            v7 -= -1
                            var v14 bool = false
                            _ = v14
                        }
                    }
                    _ = fi12
                    _ = v13
                }
                for fi15 := 0; fi15 < 15; fi15++ {
                    var v16 int16 = 1
                    _ = fi15
                    _ = v16
                }
                var v17 float64 = 1.5
                _ = fi11
                _ = v17
            }
        } else {
            var v18 float64 = 1.5
            _ = v18
        }
    }
    _ = len(v2)
    switch -651 {
    case 0:
        if false {
            v19 := GenS1{F0: 42, F1: 65535}
            fmt.Println(v7, v4, v3)
            var v20 int16 = 1
            _ = v19
            _ = v20
        } else {
            v3 = v3
            fmt.Println(v4, v3)
        }
    case 128:
        v7 = (v7 >> ((0) & 63))
    default:
        if true {
            switch v3 {
            case "hello":
                var v21 uint8 = 255
                _ = v21
            default:
                for fi22 := 0; fi22 < 15; fi22++ {
                    v4 += (v4 + 1)
                    switch v3 {
                    case "hello":
                        var v23 uint32 = 65535
                        _ = v23
                    case "\n\t":
                        var v24 float32 = -2.25
                        _ = v24
                    case "a":
                        v7 -= v7
                    default:
                        var v25 int8 = 0
                        _ = v25
                    }
                    _ = fi22
                }
            }
            for fi26 := 0; fi26 < 5; fi26++ {
                if true {
                    fmt.Println(v1, fi26)
                    v1 += (v1 + v1)
                    v7 = v7
                } else {
                    v3 = "\n\t"
                    fmt.Println(v1, v1, v1)
                    var v27 int16 = 0
                    _ = v27
                }
                switch v3 {
                case "\n\t":
                    if len(v2) > 0 {
                        v2[0] = fi26
                    }
                case "foo bar":
                    var v28 float32 = v1
                    _ = v28
                case "":
                    var v29 float32 = v1
                    _ = v29
                }
                _ = fi26
            }
            v30 := GenS1{F0: -1, F1: 0}
            _ = v30
        } else {
            for fi31 := 0; fi31 < 12; fi31++ {
                if true {
                    if true {
                        fmt.Println(v4)
                    }
                    var v32 int32 = -769
                    _ = len(v2)
                    _ = v32
                }
                if (2 <= 1) {
                    v4 = v4
                } else {
                    for fi33 := 0; fi33 < 19; fi33++ {
                        v34 := GenS1{F0: 127, F1: 1}
                        switch fi31 {
                        case 32767:
                            v7 = -1
                        }
                        v3 = v3
                        _ = fi33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 4; fi35++ {
                        v36 := GenS1{F0: 1, F1: uint16(v4)}
                        fmt.Println(v3)
                        var v37 uint16 = 65535
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    v38 := GenS2{F0: 65535}
                    _ = v38
                }
                v4 += v4
                _ = fi31
            }
            var v39 int8 = 111
            _ = v39
        }
    }
    v1 = -2.25
    v40 := GenS2{F0: 479}
    v4 = (^(v4))
    var v41 float64 = ((-2.25 / (3.14159 + (0.0 / 3.14159))) * 1.5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        v42 := GenS3{F0: 1, F1: (v7 & v7)}
        v43 := GenS4{F0: -435}
        v2 = append(v2, 2147483647)
        _ = v42
        _ = v43
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        v2 = append(v2, 469)
        var v45 int64 = v7
        _ = fi44
        _ = v45
    }
    v46 := GenS5{F0: v4}
    var v47 int16 = int16(v7)
    if true {
        if false {
            switch "hello" {
            case "a":
                v1 -= 1e10
            default:
                var v48 int = -765
                _ = v48
            }
            v1 = -2.25
        }
        switch (v3 + v3) {
        case "a":
            if true {
                var v49 uint32 = (v4 ^ v4)
                _ = v49
            }
        }
        var v50 string = v3
        _ = v50
    } else {
        v7 += v7
        var v51 uint = 4294967295
        _ = v51
    }
    v41 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v46
    _ = v47
    _ = v7
}
