package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int16; F2 int32 }
type GenS3 struct { F0 int8; F1 uint64; F2 int64 }

func main() {
    v1 := []string{"hello", "", "foo bar", "hello"}
    var v2 uint64 = 0
    var v3 int8 = 42
    var v4 string = ("" + (("a" + "hello") + "foo bar"))
    v1 = append(v1, "")
    fmt.Println(v2, v3)
    switch 32768 {
    case 32768:
        if false {
            v4 = v4
            var v5 int16 = 1000
            var v6 uint16 = 255
            _ = v5
            _ = v6
        } else {
            if false {
                for fi7 := 0; fi7 < 1; fi7++ {
                    fmt.Println(v3)
                    _ = fi7
                }
                var v8 bool = (!(-1 < -32768))
                _ = v8
            }
            switch 127 {
            case -475:
                for fi9 := 0; fi9 < 17; fi9++ {
                    fmt.Println(v2)
                    _ = len(v1)
                    fmt.Println(fi9)
                    _ = fi9
                }
            default:
                _ = len(v1)
            }
            fmt.Println(v2, v3, v2)
        }
    }
    if (v3 > (^(v3))) {
        var v10 float64 = ((1.5 / ((-2.25 / 1.0) + 1.5)) * 0.1)
        _ = v10
    } else {
        var v11 uint = 65535
        _ = v11
    }
    v12 := GenS1{F0: 42}
    fmt.Println(v3, v2, v3)
    var v13 bool = (!(!false))
    v1 = append(v1, (((v4 + v4) + v4) + "hello"))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        if v13 {
            if ((((65535 != 4294967295) && v13) || (v13 || (!v13))) || (!(-19 > (v3 >> ((42) & 63))))) {
                var v14 float64 = (float64((3.14159 + float32(0.1))) / 3.14159)
                _ = v14
            }
            if (0 <= 2147483647) {
                v2 -= (^(v2))
            }
        } else {
            var v15 uint64 = (v2 ^ (^(v2)))
            v16 := GenS1{F0: 1}
            if (!(-1 == -2147483648)) {
                if v13 {
                    v4 = ""
                    var v17 uint32 = 1
                    _ = v17
                } else {
                    v18 := GenS2{F0: v15, F1: 615, F2: 42}
                    _ = v18
                }
            } else {
                for fi19 := 0; fi19 < 12; fi19++ {
                    v13 = (v13 && v13)
                    v20 := GenS1{F0: 42}
                    var v21 float32 = (1.5 * 1e10)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            _ = v15
            _ = v16
        }
    }
    if ("hello" >= (v4 + "foo bar")) {
        if (255 < 679) {
            v13 = ((v3 - v3) >= 116)
            v22 := GenS2{F0: v2, F1: -104, F2: -387}
            _ = v22
        } else {
            fmt.Println(v4, v3)
            switch 128 {
            case 1:
                v1 = append(v1, v4)
            case -531:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case -9223372036854775808:
                switch 1 {
                case 305:
                    _ = len(v1)
                }
            }
        }
    } else {
        v4 = (((v4 + "a") + v4) + v4)
    }
    switch 115 {
    case 32768:
        switch -9223372036854775808 {
        case -429:
            v23 := GenS2{F0: v2, F1: 160, F2: -392}
            _ = v23
        case 32767:
            fmt.Println(v4, v4, v4)
        case 649:
            v24 := GenS3{F0: 0, F1: (^(v2)), F2: 9223372036854775807}
            _ = v24
        default:
            v3 = v3
        }
    case 127:
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 uint64 = v2
            v27 := GenS1{F0: -828}
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
