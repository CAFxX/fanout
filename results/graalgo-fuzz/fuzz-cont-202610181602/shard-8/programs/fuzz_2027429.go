package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 int8 = 100
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    var v4 uint16 = 42
    v5 := []uint32{65535, 1, 2, 337}
    var v6 int8 = (-(v1))
    var v7 string = "a"
    for fi8 := 0; fi8 < 13; fi8++ {
        v5 = append(v5, 2)
        if false {
            if ((fi8 + (^(fi8))) == 4294967295) {
                v6 = (v1 ^ 11)
                _ = len(v5)
            }
            if (true || true) {
                var v9 int32 = 1
                _ = v9
            }
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := GenS1{F0: true, F1: true}
                var v12 uint16 = (v4 * v4)
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            v5 = append(v5, 4294967295)
            var v13 float32 = float32(-2.25)
            _ = v13
        }
        switch 0 {
        case -519:
            var v14 float32 = 0.1
            _ = v14
        case -2147483648:
            fmt.Println(v4)
        }
        _ = fi8
    }
    v2[4294967295] = ("" + v7)
    v4 = 1
    var v15 int32 = -127
    for fi16 := 0; fi16 < 7; fi16++ {
        v7 = "a"
        switch (v7 + v7) {
        case "a":
            switch 128 {
            case 905:
                for fi17 := 0; fi17 < 17; fi17++ {
                    fmt.Println(v15)
                    v6 = v1
                    _ = fi17
                }
            case -17:
                switch fi16 {
                case -530:
                    fmt.Println(fi16)
                case -6:
                    var v18 uint = 570
                    _ = v18
                }
            }
        case "\n\t":
            var v19 string = "a"
            _ = v19
        case "":
            v1 -= (v6 % (((v6 << ((245) & 63))) | 1))
        default:
            v15 = (v15 + v15)
        }
        _ = fi16
    }
    switch 1 {
    case -9223372036854775808:
        if true {
            var v20 int8 = v1
            switch (v7 + ("a" + "hello")) {
            case "":
                for fi21 := 0; fi21 < 2; fi21++ {
                    switch v7 {
                    case "hello":
                        v22 := GenS1{F0: true, F1: true}
                        _ = v22
                    default:
                        v7 = "hello"
                    }
                    if (true || true) {
                        v23 := GenS1{F0: false, F1: false}
                        v7 = "\n\t"
                        _ = v23
                    }
                    _ = fi21
                }
            case "foo bar":
                var v24 uint16 = (^(v4))
                _ = v24
            case "a":
                var v25 uint64 = 0
                _ = v25
            }
            _ = v20
        } else {
            v6 = (v1 % (((v6 % (((v1 + (v1 * (v6 | (v1 + 127))))) | 1))) | 1))
            _ = len(v2)
        }
    case 4294967295:
        var v26 int16 = -32768
        _ = v26
    case 32767:
        switch (("hello" + ("\n\t" + "")) + "\n\t") {
        case "hello":
            for fi27 := 0; fi27 < 2; fi27++ {
                fmt.Println(v1)
                v6 -= v1
                var v28 uint8 = 42
                _ = fi27
                _ = v28
            }
        default:
            v7 = (v7 + ((v7 + "") + v7))
        }
    default:
        var v29 string = v7
        _ = v29
    }
    switch ("a" + (v7 + "hello")) {
    case "foo bar":
        if true {
            v30 := GenS1{F0: (1.5 < (0.1 / 3.14159)), F1: false}
            var v31 uint64 = 42
            _ = v30
            _ = v31
        } else {
            switch v7 {
            case "foo bar":
                fmt.Println(v6)
            case "hello":
                v32 := v3[2147483647]
                _ = v32
            }
        }
    case "":
        fmt.Println(v4)
    case "a":
        v4 += 255
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        var v34 int = (fi33 >> ((255) & 63))
        fmt.Println(v15, v1, v4)
        _ = fi33
        _ = v34
    }
    switch v7 {
    case "":
        v1 += 42
    case "hello":
        v1 -= (v6 ^ v6)
    case "\n\t":
        switch -2147483648 {
        case 32767:
            v35 := GenS2{F0: 280, F1: -870}
            _ = v35
        }
    }
    v36 := GenS3{F0: v15, F1: 2, F2: 1}
    switch v7 {
    case "hello":
        var v37 float64 = 1e10
        _ = v37
    case "foo bar":
        v1 = v6
    case "":
        for fi38 := 0; fi38 < 19; fi38++ {
            switch v7 {
            case "a":
                switch fi38 {
                case 2147483647:
                    var v39 string = v7
                    _ = v39
                case 1:
                    var v40 int32 = -98
                    _ = v40
                case 32768:
                    if false {
                        v41 := GenS4{F0: -328}
                        _ = v41
                    }
                }
            case "hello":
                v1 = (^(v6))
            }
            v7 = v7
            _ = fi38
        }
    }
    v42 := GenS4{F0: 1000}
    var v43 int = -410
    v1 = v1
    for fi44 := 0; fi44 < 4; fi44++ {
        v45 := GenS4{F0: -32768}
        fmt.Println(v1, v15, fi44)
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
