package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int32; F2 uint32 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    v3 := []string{""}
    var v4 int64 = -168
    if ((true || (true && false)) || (42 < 0)) {
        delete(v1, 32768)
        v5 := v2[-1]
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v5
    } else {
        v4 -= v4
        v4 = (v4 ^ (v4 % ((v4) | 1)))
    }
    v4 = v4
    v4 = -9223372036854775808
    if true {
        v6 := GenS1{F0: 1, F1: v4, F2: (v4 | v4)}
        v4 += 0
        _ = v6
    } else {
        if false {
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch 405 {
                case 0:
                    var v7 uint64 = 1
                    _ = v7
                }
            } else {
                delete(v1, 32767)
                for fi8 := 0; fi8 < 18; fi8++ {
                    v3 = append(v3, "hello")
                    _ = fi8
                }
            }
        } else {
            switch ("\n\t" + ("" + (("a" + "hello") + ((("hello" + "hello") + (("" + ("" + "\n\t")) + ("hello" + ("foo bar" + "hello")))) + "")))) {
            case "hello":
                fmt.Println(v4)
            case "":
                var v9 int32 = 42
                _ = v9
            default:
                var v10 uint8 = 0
                _ = v10
            }
            var v11 uint = 220
            _ = len(v1)
            _ = v11
        }
        fmt.Println(v4)
    }
    var v12 int = 128
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(v4, v4, fi13)
        v14 := GenS2{F0: 41}
        var v15 uint64 = 2
        _ = fi13
        _ = v14
        _ = v15
    }
    v1[(v12 & v12)] = 317
    var v16 uint = 42
    v12 += (v12 - (v12 ^ 128))
    for fi17 := 0; fi17 < 13; fi17++ {
        switch fi17 {
        case 32768:
            switch -101 {
            case -2147483648:
                if (1 <= 0) {
                    v18 := GenS1{F0: -1, F1: v4, F2: 1}
                    v16 = v16
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        v20 := v2[0]
                        v4 += v4
                        _ = fi19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 15; fi21++ {
                        v22 := GenS2{F0: 42}
                        var v23 uint32 = 383
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                }
            case 9223372036854775807:
                for fi24 := 0; fi24 < 6; fi24++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch "" {
                    case "hello":
                        if true {
                            v16 += v16
                        }
                    case "":
                        v12 -= (fi24 << ((v16) & 63))
                    default:
                        delete(v1, v12)
                    }
                    _ = fi24
                }
            case 32767:
                var v25 uint = (^(v16))
                _ = v25
            }
        case 128:
            switch "hello" {
            case "\n\t":
                for fi26 := 0; fi26 < 10; fi26++ {
                    v27 := GenS3{F0: v4, F1: -2147483648, F2: uint32(v16)}
                    v28 := GenS4{F0: 100, F1: -1}
                    var v29 int32 = 0
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
            case "":
                fmt.Println(v16)
            default:
                v30 := GenS5{F0: 2147483647, F1: 640}
                _ = v30
            }
        }
        fmt.Println(v12)
        v31 := GenS2{F0: uint8(v16)}
        _ = fi17
        _ = v31
    }
    v4 = v4
    var v32 bool = true
    fmt.Println(v4, v32, v16)
    if v32 {
        var v33 float32 = 1.5
        _ = v33
    } else {
        v2[9223372036854775807] = 65535
        for fi34 := 0; fi34 < 7; fi34++ {
            var v35 int32 = 1
            switch v12 {
            case 2147483647:
                _ = len(v3)
            case 1:
                v36 := GenS1{F0: -1, F1: (v4 % ((v4) | 1)), F2: (-(v4))}
                _ = v36
            case 735:
                var v37 uint8 = 127
                _ = v37
            }
            _ = fi34
            _ = v35
        }
        v38 := GenS4{F0: 127, F1: 1}
        _ = v38
    }
    var v39 uint = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
}
