package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS5 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    var v3 float64 = (-2.25 - (3.14159 * 0.0))
    v4 := GenS1{F0: 42, F1: -620}
    if (139 == 32767) {
        v3 = 3.14159
        v5 := []int{}
        fmt.Println(v3)
        _ = v5
    } else {
        if (1 >= 389) {
            switch 32767 {
            case 962:
                v3 += float64(3.14159)
            case -711:
                var v6 string = "foo bar"
                _ = v6
            default:
                v7 := GenS2{F0: 0, F1: 4294967295}
                _ = v7
            }
        }
        var v8 uint = 821
        for fi9 := 0; fi9 < 11; fi9++ {
            v3 = 1.5
            _ = fi9
        }
        _ = v8
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        switch "\n\t" {
        case "hello":
            if (-32768 < 0) {
                if false {
                    var v11 int16 = 1
                    _ = v11
                }
                v12 := GenS2{F0: 32767, F1: 1}
                _ = v12
            } else {
                fmt.Println(fi10)
                v13 := []int{(fi10 ^ -1), (fi10 + 32768)}
                _ = v13
            }
        case "":
            if ((false || false) || (fi10 <= fi10)) {
                var v14 int8 = -52
                if false {
                    switch (fi10 - -9223372036854775808) {
                    case 9223372036854775807:
                        fmt.Println(v3)
                    }
                }
                fmt.Println(fi10, v14)
                _ = v14
            }
        default:
            switch "\n\t" {
            case "\n\t":
                for fi15 := 0; fi15 < 7; fi15++ {
                    switch "foo bar" {
                    case "foo bar":
                        v16 := GenS3{F0: 65535, F1: 128, F2: 1000}
                        _ = v16
                    case "\n\t":
                        var v17 int32 = -648
                        _ = v17
                    default:
                        v18 := GenS3{F0: 65535, F1: 127, F2: 0}
                        _ = v18
                    }
                    _ = fi15
                }
            }
        }
        delete(v2, (fi10 & (fi10 % (((-(fi10))) | 1))))
        for fi19 := 0; fi19 < 16; fi19++ {
            if (127 == -128) {
                v20 := []int{fi19, fi10, fi19, (^(fi19))}
                v21 := GenS3{F0: 2147483647, F1: 255, F2: 611}
                _ = v20
                _ = v21
            }
            v3 = v3
            if (65535 == 318) {
                v22 := v1[2]
                _ = v22
            } else {
                fmt.Println(v3)
            }
            var v23 int32 = -2147483648
            _ = fi19
            _ = v23
        }
        _ = fi10
    }
    var v24 int64 = -9223372036854775808
    for fi25 := 0; fi25 < 15; fi25++ {
        if (255 == 255) {
            var v26 uint = 255
            for fi27 := 0; fi27 < 20; fi27++ {
                v3 -= float64(float32(float64(float32(v3))))
                v28 := GenS4{F0: 100, F1: 37, F2: 0}
                fmt.Println(v24, fi27, v24)
                v26 += v26
                _ = fi27
                _ = v28
            }
            if true {
                v29 := GenS2{F0: 65535, F1: 65535}
                _ = v29
            }
            _ = v26
        } else {
            v24 -= v24
            v30 := GenS4{F0: 0, F1: 127, F2: 0}
            _ = v30
        }
        v31 := GenS1{F0: 42, F1: v24}
        v32 := []int{fi25, fi25, 128, (fi25 | (fi25 + -163))}
        _ = fi25
        _ = v31
        _ = v32
    }
    v24 = (v24 & (v24 + v24))
    switch 1 {
    case -511:
        switch "" {
        case "a":
            v33 := GenS3{F0: 1, F1: 0, F2: -1}
            _ = v33
        case "foo bar":
            for fi34 := 0; fi34 < 19; fi34++ {
                switch int(v24) {
                case -581:
                    switch (fi34 & (fi34 * 128)) {
                    case 2147483647:
                        var v35 uint32 = 65535
                        _ = v35
                    case 0:
                        v24 -= -1
                    default:
                        v36 := GenS4{F0: 55, F1: -128, F2: 42}
                        _ = v36
                    }
                }
                if (0 < 255) {
                    if false {
                        var v37 uint8 = 255
                        _ = v37
                    }
                    v24 -= int64(fi34)
                }
                _ = fi34
            }
        case "":
            if (!(978 <= -866)) {
                switch 666 {
                case 32767:
                    fmt.Println(v3, v3)
                case 227:
                    v38 := GenS5{F0: -329}
                    _ = v38
                default:
                    v24 = (v24 / ((v24) | 1))
                }
                v3 = 0.0
                fmt.Println(v3, v3)
            }
        }
    case 9223372036854775807:
        v24 = -9223372036854775808
    case -2147483648:
        fmt.Println(v24, v24)
    }
    var v39 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v39
    _ = v4
}
