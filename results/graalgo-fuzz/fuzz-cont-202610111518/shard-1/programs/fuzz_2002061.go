package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int; F1 uint16 }
func genFn0() uint8 {
    v5 := []int{}
    var v6 string = "\n\t"
    v6 = (v6 + ((("\n\t" + v6) + "") + v6))
    for fi7 := 0; fi7 < 2; fi7++ {
        _ = len(v5)
        if false {
            for fi8 := 0; fi8 < 19; fi8++ {
                fmt.Println(v6)
                _ = fi8
            }
            if true {
                v6 = v6
                v6 = v6
                if ((127 == -1) && (1e10 > 1e10)) {
                    for fi9 := 0; fi9 < 11; fi9++ {
                        v6 = v6
                        v10 := GenS1{F0: 177, F1: 42, F2: 0}
                        _ = fi9
                        _ = v10
                    }
                    v6 = ((v6 + v6) + (v6 + "hello"))
                    switch fi7 {
                    case -697:
                        if ((!true) || (134 == 0)) {
                            var v11 float64 = -2.25
                            v12 := GenS1{F0: 42, F1: -1, F2: 404}
                            _ = v11
                            _ = v12
                        }
                    case 2147483647:
                        var v13 uint64 = 4294967295
                        _ = v13
                    default:
                        for fi14 := 0; fi14 < 8; fi14++ {
                            if false {
                                v6 = "a"
                            }
                            _ = fi14
                        }
                    }
                }
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(fi7)
            }
        } else {
            if false {
                v15 := GenS2{F0: false}
                if ((0.1 != 1e10) || ((127 < 128) || true)) {
                    v6 = (v6 + "foo bar")
                }
                _ = v15
            } else {
                var v16 int8 = 42
                _ = v16
            }
            fmt.Println(v6)
        }
        _ = fi7
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    _ = v5
    _ = v6
    return 3
}
func genFn1() int64 {
    var v17 bool = true
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 uint = 18446744073709551615
        v17 = false
        var v20 float32 = ((0.0 + 0.0) / 1.0)
        _ = fi18
        _ = v19
        _ = v20
    }
    v17 = (false && v17)
    _ = v17
    return 9223372036854775807
}

func main() {
    var v1 int8 = 25
    var v2 int8 = v1
    var v3 int16 = -32768
    v4 := make(map[uint32]uint32)
    var v21 int32 = int32(v3)
    if (135 > 0) {
        v22 := GenS2{F0: false}
        _ = v22
    }
    if (!true) {
        var v23 uint32 = 2147483647
        for fi24 := 0; fi24 < 3; fi24++ {
            v23 -= v23
            _ = fi24
        }
        _ = v23
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        if false {
            var v26 int64 = 0
            v26 = (v26 ^ v26)
            var v27 uint = 42
            _ = v26
            _ = v27
        }
        _ = fi25
    }
    switch -667 {
    case 906:
        delete(v4, 42)
    }
    v28 := GenS1{F0: 175, F1: 1, F2: 32767}
    var v29 bool = true
    if (255 != 2) {
        if false {
            switch 32767 {
            case -2147483648:
                switch ("a" + "foo bar") {
                case "":
                    for fi30 := 0; fi30 < 11; fi30++ {
                        var v31 int32 = (v21 + -2147483648)
                        v32 := GenS2{F0: (v1 >= 100)}
                        if v29 {
                            switch "" {
                            case "a":
                                v33 := GenS3{F0: 565, F1: 255}
                                _ = v33
                            case "\n\t":
                                var v34 string = "a"
                                _ = v34
                            default:
                                v2 -= (-(v2))
                            }
                            v35 := []int{}
                            _ = v35
                        }
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                case "\n\t":
                    switch 9223372036854775807 {
                    case 128:
                        v36 := GenS4{F0: 100, F1: 2147483647}
                        _ = v36
                    case 360:
                        var v37 int16 = (v3 << ((4294967295) & 63))
                        _ = v37
                    case 4294967295:
                        fmt.Println(v21, v21, v2)
                    default:
                        if (2 < 4294967295) {
                            v38 := GenS3{F0: 42, F1: 255}
                            switch -601 {
                            case 4294967295:
                                var v39 uint = 4294967295
                                _ = v39
                            case -448:
                                var v40 int16 = (v3 << ((255) & 63))
                                _ = v40
                            case 32767:
                                var v41 uint = 335
                                _ = v41
                            }
                            _ = v38
                        }
                    }
                case "foo bar":
                    var v42 uint64 = 0
                    _ = v42
                }
            }
            var v43 uint32 = 4294967295
            for fi44 := 0; fi44 < 18; fi44++ {
                _ = len(v4)
                _ = fi44
            }
            _ = v43
        } else {
            var v45 float64 = 0.1
            v29 = v29
            _ = v45
        }
    } else {
        var v46 uint32 = 732
        v1 = v2
        _ = v46
    }
    delete(v4, 2147483648)
    var v47 int = 704
    v29 = ((42 > 65535) && (v3 <= int16(v21)))
    v48 := GenS5{F0: (v47 - (v47 % (((v47 + v47)) | 1))), F1: 0}
    if (661 != 255) {
        v49 := v4[668]
        var v50 int32 = v21
        for fi51 := 0; fi51 < 20; fi51++ {
            v3 = 42
            _ = fi51
        }
        _ = v49
        _ = v50
    }
    var v52 bool = v29
    v53 := []int{(-(v47))}
    v47 += v47
    var v54 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v52
    _ = v53
    _ = v54
}
