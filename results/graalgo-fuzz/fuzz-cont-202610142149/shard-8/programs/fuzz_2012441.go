package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    var v1 uint = 737
    var v2 uint = (v1 - 65535)
    var v3 uint16 = 2
    var v4 string = (("hello" + "\n\t") + ("foo bar" + "foo bar"))
    v5 := make(map[uint32]string)
    v6 := []uint32{}
    v4 = "\n\t"
    delete(v5, 42)
    if ((v2 & (v2 / (((v1 * v2)) | 1))) <= v2) {
        var v7 uint32 = 4294967295
        _ = v7
    }
    switch 32768 {
    case -1:
        v8 := GenS1{F0: 255, F1: -285}
        _ = v8
    default:
        switch 9223372036854775807 {
        case 4294967295:
            v4 = v4
        }
    }
    var v9 string = "\n\t"
    v3 -= (v3 + (^(v3)))
    switch 992 {
    case -126:
        switch 127 {
        case -466:
            if false {
                for fi10 := 0; fi10 < 11; fi10++ {
                    v6 = append(v6, 2147483648)
                    fmt.Println(v3, fi10, v1)
                    fmt.Println(v3, v1)
                    _ = fi10
                }
                v11 := GenS1{F0: 1, F1: -1}
                _ = v11
            }
        default:
            if true {
                var v12 uint16 = 65535
                if len(v6) > 0 {
                    v6[0] = 2147483647
                }
                _ = v12
            }
        }
    case 32767:
        v6 = append(v6, 403)
    default:
        var v13 string = (("hello" + v9) + "hello")
        _ = v13
    }
    if false {
        switch -2147483648 {
        case 598:
            if false {
                delete(v5, 549)
                fmt.Println(v1, v2)
            }
        default:
            var v14 uint16 = 50
            _ = v14
        }
        if true {
            v15 := GenS1{F0: 1, F1: -2147483648}
            _ = v15
        }
    }
    v5[42] = v9
    fmt.Println(v3, v4)
    switch 2147483647 {
    case -9223372036854775808:
        var v16 int16 = -129
        _ = v16
    }
    v2 = v1
    v9 = v9
    var v17 uint = 727
    if (!true) {
        var v18 uint8 = 241
        v19 := GenS1{F0: (v18 << ((v17) & 63)), F1: 1}
        switch (("foo bar" + (v9 + "")) + (v9 + "foo bar")) {
        case "\n\t":
            switch "\n\t" {
            case "foo bar":
                if false {
                    v20 := GenS1{F0: (v18 * (v18 % (((v18 * (v18 - (v18 >> ((v2) & 63))))) | 1))), F1: 2147483647}
                    var v21 int16 = -32768
                    _ = v20
                    _ = v21
                } else {
                    v4 = (v9 + v9)
                    v1 = (^(v17))
                    v1 = (v17 << ((v2) & 63))
                }
            default:
                if (2147483647 < 2147483647) {
                    if (!true) {
                        v4 = v4
                        var v22 int64 = -505
                        _ = v22
                    }
                }
            }
        default:
            v17 -= v17
        }
        _ = v18
        _ = v19
    } else {
        v4 = v9
        v5[2] = v4
    }
    var v23 float64 = 1.5
    var v24 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
