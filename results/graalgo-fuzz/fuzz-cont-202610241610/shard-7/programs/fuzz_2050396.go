package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int8 }

func main() {
    var v1 int8 = 42
    v2 := []uint32{2147483647, 323, 319}
    v3 := make(map[string]string)
    var v4 uint8 = 255
    v5 := v3["foo bar"]
    if true {
        v1 -= (v1 + (^(v1)))
        v1 += v1
    } else {
        _ = len(v2)
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := GenS1{F0: 1000, F1: 18446744073709551615, F2: 42}
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 uint8 = (v4 % ((v4) | 1))
                var v10 bool = false
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    }
    switch -286 {
    case 127:
        var v11 uint32 = 65535
        _ = v11
    case 0:
        if (true || true) {
            if (-128 > (v1 >> ((255) & 63))) {
                var v12 int32 = 1
                _ = v12
            }
        }
    }
    fmt.Println(v4, v1, v4)
    var v13 uint8 = (v4 + v4)
    v13 = 2
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    var v14 float64 = (1e10 / float64((1.5 - 3.14159)))
    if (!(32767 > -1)) {
        fmt.Println(v4)
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 int = 42
            _ = fi15
            _ = v16
        }
        var v17 uint8 = 90
        _ = v17
    }
    fmt.Println(v4)
    if ((821 < 329) || (0 == 0)) {
        for fi18 := 0; fi18 < 13; fi18++ {
            if len(v2) > 0 {
                v2[0] = 42
            }
            _ = fi18
        }
        v4 -= (v13 % ((v13) | 1))
    } else {
        switch 3 {
        case -333:
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 uint = 255
                _ = fi19
                _ = v20
            }
        case -9223372036854775808:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    v21 := v3[v5]
    var v22 int32 = -2147483648
    for fi23 := 0; fi23 < 2; fi23++ {
        v4 += v4
        for fi24 := 0; fi24 < 8; fi24++ {
            switch "\n\t" {
            case "a":
                v25 := GenS2{F0: (^(v13)), F1: 2, F2: v1}
                _ = v25
            default:
                switch (fi24 + int(v1)) {
                case 977:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            _ = fi24
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
