package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint }
type GenS4 struct { F0 int64; F1 uint32; F2 uint32 }

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    var v4 int = -9223372036854775808
    var v5 int = -1
    var v6 uint16 = 42
    var v7 uint8 = 127
    v8 := GenS1{F0: 876, F1: 1, F2: false}
    switch "a" {
    case "foo bar":
        v1 -= v1
    }
    v6 -= v6
    v6 -= (v6 * v6)
    fmt.Println(v4, v5, v7)
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(v4, fi9, fi9)
        switch 395 {
        case 32767:
            for fi10 := 0; fi10 < 5; fi10++ {
                switch v4 {
                case 42:
                    var v11 uint64 = 1
                    _ = v11
                default:
                    if false {
                        v12 := GenS1{F0: 209, F1: 1, F2: false}
                        switch "" {
                        case "hello":
                            v1 -= (v1 & v1)
                        }
                        switch "a" {
                        case "":
                            if true {
                                var v13 int8 = 1
                                switch "\n\t" {
                                case "\n\t":
                                    if false {
                                        if false {
                                            _ = len(v3)
                                            var v14 bool = true
                                            v1 += v1
                                            _ = v14
                                        } else {
                                            var v15 uint64 = 1
                                            _ = v15
                                        }
                                        if true {
                                            _ = len(v2)
                                            v5 = fi10
                                            v16 := GenS2{F0: v6}
                                            _ = v16
                                        } else {
                                            v17 := GenS3{F0: 4294967295, F1: 42, F2: 4294967295}
                                            _ = v17
                                        }
                                        var v18 int16 = -1
                                        _ = v18
                                    } else {
                                        fmt.Println(v5, fi10, v6)
                                    }
                                case "foo bar":
                                    switch ("a" + "") {
                                    case "hello":
                                        v19 := GenS1{F0: 255, F1: 65535, F2: (false && (62 == -1))}
                                        switch (v4 % ((v4) | 1)) {
                                        case 1:
                                            var v20 uint = 18446744073709551615
                                            v2["hello"] = v1
                                            if true {
                                                v21 := v3[""]
                                                var v22 int8 = (v13 >> ((65535) & 63))
                                                _ = v21
                                                _ = v22
                                            } else {
                                                v6 -= 255
                                                v23 := GenS2{F0: (v6 >> (((v20 & uint(v1))) & 63))}
                                                _ = len(v3)
                                                switch "hello" {
                                                case "foo bar":
                                                    v24 := GenS2{F0: (v6 + uint16(v20))}
                                                    _ = v24
                                                default:
                                                    v25 := GenS2{F0: 255}
                                                    _ = v25
                                                }
                                                _ = v23
                                            }
                                            switch "\n\t" {
                                            case "foo bar":
                                                v4 -= fi10
                                            case "a":
                                                v13 += v13
                                            case "\n\t":
                                                v26 := []int{(v4 & (fi10 | int(v1))), v5, 127}
                                                _ = v26
                                            default:
                                                delete(v3, "\n\t")
                                            }
                                            _ = v20
                                        case 36:
                                            v1 += (v1 + 42)
                                        }
                                        _ = v19
                                    case "\n\t":
                                        v1 += (v1 & 2147483648)
                                    default:
                                        v7 += 1
                                    }
                                }
                                _ = v13
                            }
                            switch "a" {
                            case "hello":
                                v2["foo bar"] = v1
                            }
                            v4 += (v5 & fi10)
                        case "\n\t":
                            v6 = v6
                        default:
                            fmt.Println(v5, fi9)
                            v7 += v7
                            fmt.Println(fi9)
                            _ = len(v3)
                            var v27 uint16 = (v6 ^ 0)
                            _ = v27
                        }
                        _ = v12
                    } else {
                        var v28 uint8 = 127
                        switch 32768 {
                        case -9223372036854775808:
                            var v29 uint16 = 42
                            _ = v29
                        case 0:
                            v7 = (v28 + 2)
                        case 32767:
                            v5 -= (v4 ^ -287)
                        default:
                            var v30 uint = 1
                            _ = v30
                        }
                        fmt.Println(v7, v6)
                        var v31 uint16 = (v6 >> ((637) & 63))
                        var v32 int32 = 1
                        _ = v28
                        _ = v31
                        _ = v32
                    }
                }
                v6 = v6
                if true {
                    var v33 uint64 = 0
                    _ = v33
                } else {
                    var v34 uint8 = (^(v7))
                    v1 = v1
                    v2[""] = v1
                    v34 = v7
                    _ = v34
                }
                v7 = (v7 % (((v7 + 42)) | 1))
                var v35 int = -967
                _ = fi10
                _ = v35
            }
        }
        _ = fi9
    }
    var v36 uint16 = (v6 << ((969) & 63))
    v37 := GenS4{F0: -1, F1: v1, F2: 2147483648}
    var v38 bool = true
    var v39 int32 = 2147483647
    var v40 int32 = v39
    if v38 {
        var v41 int32 = -2147483648
        fmt.Println(v36, v36, v6)
        _ = v41
    } else {
        v6 = 1
        for fi42 := 0; fi42 < 14; fi42++ {
            fmt.Println(v6, v4)
            for fi43 := 0; fi43 < 13; fi43++ {
                if v38 {
                    if ((float32(3.14159) - 0.0) < 0.1) {
                        var v44 uint16 = v36
                        _ = v44
                    }
                    var v45 uint16 = 304
                    _ = v45
                } else {
                    switch "foo bar" {
                    case "hello":
                        switch 32768 {
                        case 944:
                            var v46 string = ""
                            _ = v46
                        }
                    case "a":
                        var v47 uint = 1
                        _ = v47
                    default:
                        var v48 int16 = 186
                        _ = v48
                    }
                    v49 := GenS1{F0: 156, F1: 42, F2: (-574 >= 846)}
                    var v50 float64 = (((0.0 + ((3.14159 - float64(-2.25)) * float64(-2.25))) * 1e10) / 1e10)
                    _ = v49
                    _ = v50
                }
                switch "foo bar" {
                case "hello":
                    fmt.Println(fi43, v1, v39)
                case "\n\t":
                    v4 = -9223372036854775808
                }
                _ = fi43
            }
            _ = fi42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
