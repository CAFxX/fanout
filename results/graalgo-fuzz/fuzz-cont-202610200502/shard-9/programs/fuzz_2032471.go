package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 bool; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 163
    var v3 uint8 = 127
    v4 := make(map[uint32]string)
    var v5 uint = 65535
    var v6 int8 = -66
    v6 -= v6
    _ = len(v1)
    var v7 int16 = 737
    var v8 int = 456
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int = fi9
        v11 := GenS1{F0: -553, F1: v6}
        v10 -= 1
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 float32 = 3.14159
    var v13 uint8 = (v3 >> ((961) & 63))
    var v14 int8 = (v6 - v6)
    var v15 int16 = 0
    v15 -= v7
    v8 = v8
    var v16 int16 = (-(v7))
    for fi17 := 0; fi17 < 16; fi17++ {
        v8 -= (v8 / (((v8 << (((v5 - (v5 & 393))) & 63))) | 1))
        switch 272 {
        case 1:
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 uint32 = (v2 >> ((v5) & 63))
                if false {
                    if (float32(3.14159) < 0.0) {
                        v20 := GenS1{F0: -524, F1: v6}
                        v12 = v12
                        _ = v20
                    }
                    if (-86 < -801) {
                        var v21 int = 127
                        if false {
                            switch "hello" {
                            case "a":
                                var v22 int64 = -9223372036854775808
                                _ = v22
                            case "hello":
                                v23 := GenS1{F0: -338, F1: (v6 - v14)}
                                _ = v23
                            case "foo bar":
                                var v24 string = ("\n\t" + "foo bar")
                                _ = v24
                            }
                        }
                        v2 = 306
                        v25 := GenS1{F0: 501, F1: -128}
                        switch (v21 & v21) {
                        case 813:
                            v26 := v4[v19]
                            _ = v26
                        case -223:
                            var v27 int = (fi18 << ((1) & 63))
                            _ = v27
                        case 2147483647:
                            v21 += fi17
                            switch "\n\t" {
                            case "\n\t":
                                v5 = v5
                            case "foo bar":
                                fmt.Println(v19, v14)
                            case "":
                                var v28 int32 = -1
                                _ = v28
                            default:
                                _ = len(v1)
                                v29 := GenS1{F0: -2147483648, F1: -112}
                                _ = v29
                            }
                            v30 := GenS2{F0: false, F1: (v19 + 42), F2: -1}
                            _ = v30
                        }
                        var v31 uint32 = v2
                        fmt.Println(v14)
                        var v32 float64 = 1e10
                        _ = v21
                        _ = v25
                        _ = v31
                        _ = v32
                    }
                    v16 = v15
                    var v33 int8 = v6
                    _ = v33
                } else {
                    v14 += int8(fi18)
                    delete(v1, (("\n\t" + "\n\t") + "\n\t"))
                    if (true || true) {
                        var v34 int16 = 42
                        v13 -= v3
                        var v35 uint8 = (^(v13))
                        _ = v34
                        _ = v35
                    }
                    var v36 string = ("\n\t" + "foo bar")
                    if false {
                        v37 := GenS1{F0: -1, F1: (v14 % ((100) | 1))}
                        var v38 uint64 = uint64(v19)
                        var v39 float32 = v12
                        fmt.Println(v3)
                        v12 -= (v39 * v12)
                        var v40 float32 = float32(0.1)
                        _ = v37
                        _ = v38
                        _ = v39
                        _ = v40
                    } else {
                        _ = len(v4)
                        v41 := GenS2{F0: false, F1: 2, F2: 42}
                        v3 = v13
                        _ = v41
                    }
                    _ = v36
                }
                v42 := GenS3{F0: 42, F1: 7}
                _ = fi18
                _ = v19
                _ = v42
            }
        case 128:
            v13 -= 128
        }
        fmt.Println(v12, v13, v5)
        _ = fi17
    }
    for fi43 := 0; fi43 < 14; fi43++ {
        for fi44 := 0; fi44 < 20; fi44++ {
            var v45 int16 = v15
            _ = fi44
            _ = v45
        }
        _ = fi43
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        v47 := []int{}
        fmt.Println(v6)
        switch ("foo bar" + "a") {
        case "\n\t":
            if true {
                for fi48 := 0; fi48 < 18; fi48++ {
                    var v49 uint = (v5 % ((v5) | 1))
                    v50 := GenS4{F0: 1}
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            }
        case "a":
            fmt.Println(v5, v3)
        default:
            for fi51 := 0; fi51 < 7; fi51++ {
                v52 := GenS5{F0: ""}
                _ = fi51
                _ = v52
            }
        }
        _ = fi46
        _ = v47
    }
    var v53 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v7
    _ = v8
}
