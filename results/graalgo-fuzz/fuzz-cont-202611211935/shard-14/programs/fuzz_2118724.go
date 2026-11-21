package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 int32 }

func main() {
    v1 := []int{}
    var v2 uint16 = 0
    v3 := []uint32{0, 65535, uint32(v2), 65535}
    v4 := []int{-2147483648, -9223372036854775808}
    var v5 string = "\n\t"
    v6 := make(map[uint32]uint32)
    var v7 string = ("" + "hello")
    var v8 uint8 = 232
    var v9 float64 = 0.1
    v10 := GenS1{F0: -9223372036854775808}
    if (!((0 >= 736) && false)) {
        for fi11 := 0; fi11 < 6; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 bool = ((42 < 65535) && (561 >= 0))
                if true {
                    delete(v6, 992)
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v8)
                        switch "a" {
                        case "hello":
                            var v14 string = v5
                            _ = v14
                        default:
                            v15 := GenS2{F0: (v2 - 1), F1: 267, F2: 42}
                            _ = v15
                        }
                    case "\n\t":
                        fmt.Println(v5)
                    case "":
                        switch 670 {
                        case 128:
                            v8 += (v8 - 0)
                        case 127:
                            v16 := GenS2{F0: 65535, F1: 0, F2: 42}
                            v17 := GenS2{F0: 255, F1: -32768, F2: 435}
                            _ = v16
                            _ = v17
                        }
                    default:
                        v18 := GenS2{F0: 0, F1: 550, F2: 895}
                        _ = v18
                    }
                } else {
                    v5 = "a"
                    if len(v1) > 0 {
                        v1[0] = fi12
                    }
                    if v13 {
                        var v19 bool = v13
                        _ = v19
                    } else {
                        if true {
                            var v20 uint16 = (v2 * (v2 * (v2 % ((v2) | 1))))
                            switch (fi11 << ((127) & 63)) {
                            case 680:
                                v21 := GenS1{F0: 42}
                                _ = v21
                            case -62:
                                var v22 uint8 = v8
                                _ = v22
                            case -9223372036854775808:
                                var v23 uint16 = uint16(v8)
                                _ = v23
                            default:
                                v7 = v5
                            }
                            fmt.Println(v9)
                            _ = v20
                        }
                        v24 := GenS2{F0: 1, F1: 1000, F2: 65535}
                        switch "foo bar" {
                        case "":
                            v13 = v13
                        case "\n\t":
                            if ((!v13) || v13) {
                                var v25 int32 = -708
                                var v26 string = "\n\t"
                                fmt.Println(fi12, v25)
                                _ = v25
                                _ = v26
                            } else {
                                delete(v6, 1)
                            }
                        case "hello":
                            _ = len(v6)
                            if v13 {
                                var v27 string = v7
                                _ = v27
                            } else {
                                v28 := GenS3{F0: 42}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v28
                            }
                            _ = len(v1)
                        default:
                            var v29 uint = 682
                            _ = v29
                        }
                        _ = v24
                    }
                }
                if v13 {
                    if (v13 || (-9223372036854775808 < -1)) {
                        v2 = 42
                    }
                } else {
                    if v13 {
                        switch fi11 {
                        case 32767:
                            var v30 int = fi11
                            _ = v30
                        case -247:
                            v7 = (("a" + v5) + "\n\t")
                        default:
                            v1 = append(v1, (^(fi11)))
                        }
                        v8 = (v8 | 255)
                        switch 2147483647 {
                        case 2147483647:
                            fmt.Println(v8, v9, fi11)
                        case 32767:
                            var v31 uint = 1
                            _ = v31
                        }
                    }
                    if v13 {
                        v32 := GenS3{F0: -128}
                        _ = v32
                    }
                }
                _ = fi12
                _ = v13
            }
            if false {
                v33 := GenS3{F0: 1}
                switch 9223372036854775807 {
                case -2147483648:
                    v34 := GenS1{F0: -1}
                    _ = v34
                case 2147483647:
                    for fi35 := 0; fi35 < 8; fi35++ {
                        fmt.Println(fi35, v8, v9)
                        switch v7 {
                        case "":
                            var v36 int64 = -1
                            _ = v36
                        }
                        _ = fi35
                    }
                case 0:
                    var v37 uint16 = (v2 + uint16(v8))
                    _ = v37
                }
                _ = v33
            } else {
                if true {
                    fmt.Println(v7, v8, v5)
                    v2 += (v2 * (v2 << ((18446744073709551615) & 63)))
                } else {
                    v9 = v9
                    v38 := GenS2{F0: 0, F1: 1000, F2: 2147483647}
                    v7 = ("a" + v5)
                    _ = v38
                }
                switch "" {
                case "a":
                    switch (fi11 & 9223372036854775807) {
                    case 42:
                        var v39 float32 = 0.1
                        _ = v39
                    case 127:
                        var v40 int16 = 42
                        _ = v40
                    case 32767:
                        var v41 int16 = 42
                        _ = v41
                    default:
                        v7 = (((v7 + "foo bar") + v7) + v7)
                    }
                case "foo bar":
                    switch (v5 + "hello") {
                    case "":
                        v42 := GenS3{F0: -19}
                        _ = v42
                    default:
                        var v43 string = "hello"
                        _ = v43
                    }
                }
                if false {
                    v44 := GenS4{F0: 1, F1: 723}
                    _ = v44
                } else {
                    var v45 uint16 = (v2 / ((65535) | 1))
                    for fi46 := 0; fi46 < 2; fi46++ {
                        var v47 bool = false
                        v9 -= v9
                        var v48 int32 = -676
                        var v49 int64 = 1
                        switch v5 {
                        case "\n\t":
                            if ((!v47) || false) {
                                if true {
                                    fmt.Println(v7, fi46)
                                    var v50 uint16 = (v45 + v2)
                                    _ = v50
                                } else {
                                    v47 = (v47 || ((0.0 - 3.14159) <= (1e10 / (0.1 + 1e10))))
                                    v51 := v6[42]
                                    var v52 int64 = (v49 & v49)
                                    _ = v51
                                    _ = v52
                                }
                                v2 -= (v45 - (v45 << ((4294967295) & 63)))
                                v53 := GenS5{F0: v48}
                                _ = v53
                            }
                        }
                        _ = fi46
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                    _ = v45
                }
            }
            _ = fi11
        }
    }
    v5 = v7
    v2 -= v2
    switch 127 {
    case 300:
        fmt.Println(v7)
    }
    switch ("a" + ("foo bar" + v5)) {
    case "":
        var v54 int32 = 230
        _ = v54
    case "hello":
        v8 = v8
    default:
        var v55 int8 = 42
        _ = v55
    }
    var v56 int64 = -1
    for fi57 := 0; fi57 < 1; fi57++ {
        var v58 float64 = float64(-2.25)
        var v59 int32 = 235
        _ = fi57
        _ = v58
        _ = v59
    }
    v60 := GenS3{F0: -54}
    v2 = (v2 & v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v60
    _ = v7
    _ = v8
    _ = v9
}
