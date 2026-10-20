package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int64; F2 bool }

func main() {
    var v1 uint16 = 1
    var v2 string = ""
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 uint16 = (v1 << ((42) & 63))
    var v6 float64 = 3.14159
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{fi7}
        v4[-333] = 0
        var v9 int16 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    if (false && true) {
        v2 = v2
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 bool = (39 == -117)
            _ = fi10
            _ = v11
        }
    }
    switch 32767 {
    case 127:
        var v12 float64 = v6
        _ = v12
    case 0:
        for fi13 := 0; fi13 < 9; fi13++ {
            fmt.Println(fi13, v1, v1)
            fmt.Println(v2, v1, v2)
            v14 := GenS1{F0: -1, F1: 124}
            _ = fi13
            _ = v14
        }
    case -1:
        if false {
            v15 := []int{42, 42}
            v16 := GenS1{F0: -788, F1: 183}
            _ = v15
            _ = v16
        }
    default:
        fmt.Println(v5)
    }
    if (0 != 363) {
        var v17 int64 = 9223372036854775807
        var v18 uint8 = uint8(v5)
        _ = v17
        _ = v18
    }
    fmt.Println(v5)
    var v19 string = v2
    if (false || false) {
        v20 := GenS1{F0: -1, F1: 128}
        _ = v20
    }
    if (231 <= 127) {
        v6 = v6
        v21 := GenS2{F0: 1000}
        v5 = (^(v1))
        _ = v21
    }
    v22 := []int{2147483647}
    var v23 uint16 = v1
    v23 += (v23 >> ((385) & 63))
    switch 32768 {
    case 4294967295:
        if true {
            var v24 uint8 = 128
            if false {
                var v25 int16 = -32768
                for fi26 := 0; fi26 < 12; fi26++ {
                    v22 = append(v22, (fi26 * 127))
                    for fi27 := 0; fi27 < 6; fi27++ {
                        v23 += v5
                        if true {
                            v25 += -105
                            if true {
                                switch v19 {
                                case "foo bar":
                                    if (32767 != (v25 << ((4294967295) & 63))) {
                                        fmt.Println(v5)
                                        var v28 int = 9223372036854775807
                                        v29 := GenS2{F0: (v25 | 0)}
                                        _ = v28
                                        _ = v29
                                    }
                                case "hello":
                                    v2 = v19
                                }
                                if false {
                                    if true {
                                        fmt.Println(v5, v6)
                                        delete(v3, (-(fi26)))
                                    }
                                    if false {
                                        var v30 int16 = int16(v1)
                                        v19 = "\n\t"
                                        switch ("" + v19) {
                                        case "foo bar":
                                            v31 := GenS1{F0: 661, F1: v24}
                                            _ = v31
                                        case "hello":
                                            v32 := GenS3{F0: uint32(v24), F1: v23}
                                            _ = v32
                                        }
                                        _ = v30
                                    }
                                }
                                v33 := GenS3{F0: 65535, F1: v1}
                                fmt.Println(v23, v23, fi26)
                                _ = v33
                            } else {
                                var v34 int16 = v25
                                var v35 float64 = 3.14159
                                _ = v34
                                _ = v35
                            }
                            v6 -= v6
                        } else {
                            var v36 string = v2
                            _ = v36
                        }
                        _ = fi27
                    }
                    switch fi26 {
                    case 9223372036854775807:
                        v37 := GenS1{F0: (v25 + v25), F1: (v24 >> ((65535) & 63))}
                        _ = v37
                    case 143:
                        var v38 int64 = -1
                        _ = v38
                    }
                    _ = fi26
                }
                var v39 int16 = v25
                _ = v25
                _ = v39
            } else {
                fmt.Println(v24, v1, v24)
                var v40 int8 = 1
                v41 := v3[255]
                _ = v40
                _ = v41
            }
            var v42 uint = 220
            _ = v24
            _ = v42
        } else {
            v43 := GenS4{F0: -32768}
            switch 2147483647 {
            case 4294967295:
                for fi44 := 0; fi44 < 17; fi44++ {
                    v45 := GenS4{F0: -1}
                    var v46 int16 = 1000
                    v5 -= v1
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            case 2147483647:
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            }
            _ = v43
        }
    case 127:
        v47 := GenS5{F0: 42, F1: -1, F2: false}
        _ = v47
    }
    var v48 float32 = -2.25
    switch "foo bar" {
    case "\n\t":
        if false {
            fmt.Println(v19, v2)
        }
    case "":
        var v49 uint32 = 4294967295
        _ = v49
    }
    fmt.Println(v5, v23, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
