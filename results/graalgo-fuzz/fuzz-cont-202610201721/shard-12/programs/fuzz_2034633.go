package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int32; F2 int16 }
type GenS3 struct { F0 int; F1 bool; F2 int8 }
type GenS4 struct { F0 bool; F1 int16; F2 uint }

func main() {
    v1 := []string{"foo bar", "a", (("hello" + ("hello" + "hello")) + "\n\t")}
    var v2 int = 32768
    var v3 uint32 = 0
    v4 := GenS1{F0: 161}
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint8 = 216
        _ = fi5
        _ = v6
    }
    if ((float64(3.14159) >= 3.14159) || true) {
        v3 = (v3 / (((v3 + v3)) | 1))
        if false {
            var v7 int32 = -356
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi8 := 0; fi8 < 9; fi8++ {
                v7 = (v7 / ((v7) | 1))
                var v9 int8 = 127
                for fi10 := 0; fi10 < 17; fi10++ {
                    v11 := GenS2{F0: v7, F1: (v7 / (((-(v7))) | 1)), F2: 1000}
                    v2 = fi8
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    _ = fi10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int16 = 218
            _ = fi12
            _ = v13
        }
    }
    var v14 int32 = -1
    v15 := GenS2{F0: 2147483647, F1: (v14 ^ (v14 << ((42) & 63))), F2: 42}
    var v16 int = v2
    switch v2 {
    case 2147483647:
        v2 += v2
    case 32768:
        v16 = v16
    default:
        switch "foo bar" {
        case "foo bar":
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 int8 = -1
                _ = fi17
                _ = v18
            }
        case "a":
            switch v16 {
            case -1:
                var v19 int = v16
                _ = v19
            case -551:
                switch ("hello" + "\n\t") {
                case "foo bar":
                    fmt.Println(v16)
                default:
                    fmt.Println(v14, v14, v14)
                }
            case 4294967295:
                var v20 int8 = -1
                _ = v20
            default:
                var v21 uint32 = 716
                _ = v21
            }
        }
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        _ = len(v1)
        fmt.Println(v3)
        _ = fi22
    }
    switch "a" {
    case "a":
        fmt.Println(v3)
    }
    var v23 int16 = 42
    var v24 uint32 = (v3 & 662)
    _ = len(v1)
    fmt.Println(v2, v24, v23)
    v14 -= v14
    v1 = append(v1, "hello")
    for fi25 := 0; fi25 < 15; fi25++ {
        switch "a" {
        case "a":
            switch (fi25 ^ (-(v2))) {
            case 4294967295:
                switch "hello" {
                case "foo bar":
                    var v26 uint = 2
                    _ = v26
                case "a":
                    if ((4294967295 >= 152) || ((!true) && true)) {
                        v27 := GenS3{F0: v2, F1: (4294967295 != 45), F2: -88}
                        var v28 int64 = 105
                        var v29 uint = 57
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                }
            default:
                v2 -= (fi25 - (v16 * v16))
            }
        case "hello":
            v16 = fi25
        default:
            v30 := GenS2{F0: v14, F1: v14, F2: 10}
            _ = v30
        }
        _ = fi25
    }
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v16 -= v16
        if false {
            var v31 uint16 = 388
            switch "\n\t" {
            case "foo bar":
                if false {
                    v16 += (-(v2))
                } else {
                    _ = len(v1)
                    v16 = v2
                    switch v2 {
                    case 128:
                        v1 = append(v1, "\n\t")
                    case 9223372036854775807:
                        var v32 float64 = 3.14159
                        _ = v32
                    default:
                        v33 := GenS2{F0: (v14 / (((v14 * v14)) | 1)), F1: (v14 / ((v14) | 1)), F2: (v23 << ((65535) & 63))}
                        _ = v33
                    }
                }
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            default:
                var v34 uint8 = 255
                _ = v34
            }
            _ = v31
        } else {
            _ = len(v1)
            v35 := GenS2{F0: (v14 * (v14 ^ (-(v14)))), F1: (v14 / ((v14) | 1)), F2: -1}
            var v36 uint64 = 2
            _ = v35
            _ = v36
        }
    } else {
        v24 = 2147483648
        v37 := GenS3{F0: 338, F1: (42 <= 65535), F2: 91}
        _ = v37
    }
    for fi38 := 0; fi38 < 6; fi38++ {
        var v39 bool = (v14 > v14)
        v40 := GenS4{F0: v39, F1: v23, F2: 166}
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
