package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int16; F2 uint32 }

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]string)
    var v3 bool = (0.0 <= 0.0)
    v4 := make(map[int]uint32)
    var v5 int64 = -9223372036854775808
    v6 := make(map[int]uint32)
    var v7 string = ("\n\t" + "\n\t")
    v8 := []int{-9223372036854775808, 270, 184, -1}
    fmt.Println(v1)
    var v9 uint64 = 4294967295
    switch "\n\t" {
    case "a":
        fmt.Println(v7)
    case "hello":
        for fi10 := 0; fi10 < 16; fi10++ {
            v5 += (v1 >> ((633) & 63))
            var v11 int16 = int16(v9)
            fmt.Println(fi10, v9)
            _ = fi10
            _ = v11
        }
    default:
        v12 := GenS1{F0: 65535, F1: v7}
        _ = v12
    }
    var v13 float64 = ((3.14159 * -2.25) + 3.14159)
    var v14 uint = 541
    delete(v4, 32767)
    if (v3 && true) {
        if v3 {
            switch 1 {
            case -1:
                if (float32(0.1) == 0.1) {
                    if (v1 > v1) {
                        v13 = (3.14159 + v13)
                    } else {
                        var v15 int16 = -704
                        v15 = v15
                        _ = v15
                    }
                    v9 = (v9 + v9)
                    switch "\n\t" {
                    case "foo bar":
                        if len(v8) > 0 {
                            v8[0] = -316
                        }
                    case "":
                        var v16 bool = v3
                        _ = v16
                    }
                } else {
                    v17 := GenS1{F0: 226, F1: "hello"}
                    _ = v17
                }
            case 127:
                delete(v2, 1)
            }
        }
        if (v3 || (!(267 <= 316))) {
            fmt.Println(v3, v7)
            v3 = v3
            v13 = (0.0 + (float64((1.5 + ((3.14159 + (float32((v13 + v13)) - (float32((float64((3.14159 / (0.0 - (0.1 / (0.0 - 3.14159))))) - 1.5)) - 1.5))) - -2.25))) - 0.1))
        }
        v18 := GenS2{F0: v1}
        _ = v18
    } else {
        switch v7 {
        case "foo bar":
            for fi19 := 0; fi19 < 7; fi19++ {
                v20 := GenS2{F0: (v5 * (v1 << ((uint(v9)) & 63)))}
                v7 = v7
                v14 += 255
                _ = fi19
                _ = v20
            }
        case "a":
            if v3 {
                v21 := GenS3{F0: 938}
                switch -475 {
                case 565:
                    fmt.Println(v9, v3, v5)
                default:
                    if (32767 == -153) {
                        if len(v8) > 0 {
                            v8[0] = -1
                        }
                        var v22 bool = v3
                        v23 := GenS3{F0: (^(v9))}
                        _ = v22
                        _ = v23
                    }
                }
                var v24 uint32 = 4294967295
                _ = v21
                _ = v24
            }
        case "":
            for fi25 := 0; fi25 < 14; fi25++ {
                if false {
                    var v26 uint16 = 1
                    _ = v26
                } else {
                    v3 = (751 > 0)
                }
                for fi27 := 0; fi27 < 4; fi27++ {
                    if (69 > 1) {
                        var v28 int32 = 42
                        v13 -= v13
                        var v29 bool = ((1 > -1) && ((("" + "foo bar") + v7) != v7))
                        _ = v28
                        _ = v29
                    } else {
                        v30 := GenS4{F0: v1, F1: -32768, F2: 65535}
                        var v31 uint32 = 42
                        _ = v30
                        _ = v31
                    }
                    v9 = 1
                    var v32 uint8 = 2
                    _ = fi27
                    _ = v32
                }
                _ = fi25
            }
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        v5 -= v1
    }
    if (689 > 128) {
        var v33 float64 = -2.25
        _ = v33
    } else {
        var v34 bool = v3
        for fi35 := 0; fi35 < 6; fi35++ {
            if true {
                if v34 {
                    for fi36 := 0; fi36 < 3; fi36++ {
                        v37 := GenS2{F0: (v5 & 1)}
                        v14 = v14
                        _ = fi36
                        _ = v37
                    }
                }
                switch fi35 {
                case 9223372036854775807:
                    fmt.Println(v13)
                }
                var v38 float32 = 3.14159
                _ = v38
            } else {
                v9 = v9
            }
            _ = fi35
        }
        _ = v34
    }
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
