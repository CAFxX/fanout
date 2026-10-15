package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 int; F1 string; F2 int64 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint16 }

func main() {
    v1 := []string{}
    v2 := make(map[int]bool)
    v3 := []int{2147483647, 128, -835, -9223372036854775808}
    var v4 int8 = -1
    var v5 uint16 = 114
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = 8
    var v6 uint = 0
    v7 := GenS1{F0: 255, F1: -73, F2: 0}
    if false {
        var v8 uint8 = 1
        if false {
            v5 += 844
            var v9 uint32 = 1
            var v10 int32 = -2147483648
            _ = v10
            _ = v9
        }
        switch 4294967295 {
        case 32768:
            if ((((float32((0.0 / -2.25)) - 1.5) / 3.14159) > 1e10) || (!false)) {
                if (-1 < 42) {
                    if true {
                        v5 = 65535
                    } else {
                        v11 := GenS1{F0: (v6 % ((v6) | 1)), F1: 1, F2: uint32(v8)}
                        var v12 uint32 = uint32(v5)
                        _ = v11
                        _ = v12
                    }
                } else {
                    delete(v2, 2147483647)
                    fmt.Println(v8, v8)
                }
            } else {
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    var v13 bool = (42 <= 295)
                    _ = v13
                case "hello":
                    var v14 int32 = -1
                    _ = v14
                }
                v4 = 42
            }
        }
        _ = v8
    }
    if false {
        var v15 string = ""
        v4 += (v4 | v4)
        v16 := GenS2{F0: -2147483648, F1: "", F2: 0}
        _ = v15
        _ = v16
    } else {
        var v17 uint32 = 65535
        for fi18 := 0; fi18 < 1; fi18++ {
            v5 = (v5 + uint16(v17))
            switch ("hello" + ("hello" + "")) {
            case "":
                if ((v5 ^ 255) > (v5 % ((v5) | 1))) {
                    v6 = v6
                    v4 -= v4
                    var v19 float64 = (3.14159 - 1.5)
                    _ = v19
                } else {
                    switch ("foo bar" + "\n\t") {
                    case "a":
                        var v20 uint8 = uint8(v5)
                        _ = v20
                    }
                }
            case "foo bar":
                var v21 float32 = ((3.14159 - 3.14159) + 0.1)
                _ = v21
            default:
                v6 += 255
            }
            _ = fi18
        }
        v2[669] = ((858 == 42) || (("\n\t" + "a") < (("foo bar" + ("foo bar" + "")) + "")))
        _ = v17
    }
    if false {
        var v22 uint32 = 683
        _ = v22
    } else {
        if true {
            v4 = (v4 << ((v6) & 63))
            for fi23 := 0; fi23 < 12; fi23++ {
                fmt.Println(v5)
                var v24 string = "a"
                v24 = "\n\t"
                _ = fi23
                _ = v24
            }
            var v25 float64 = 1.5
            _ = v25
        } else {
            for fi26 := 0; fi26 < 14; fi26++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi26
            }
            fmt.Println(v4, v5, v5)
            var v27 int = -201
            _ = v27
        }
        v6 = v6
    }
    switch -880 {
    case 1:
        v28 := GenS3{F0: 32767, F1: 127, F2: (v5 / (((v5 * 423)) | 1))}
        _ = v28
    case 19:
        var v29 bool = true
        _ = v29
    default:
        if (42 != 529) {
            v6 = 4294967295
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
