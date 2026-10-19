package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = ((1e10 * 3.14159) + 0.0)
    var v3 float32 = (((((0.1 + ((1e10 + (0.0 / (0.1 + (3.14159 + (float32(0.1) / 1e10))))) * 1e10)) + (float32((v2 / 1e10)) * -2.25)) / float32(float64((3.14159 + float32(0.0))))) / 1e10) * 1.5)
    var v4 bool = false
    if ("foo bar" > "") {
        var v5 uint8 = 0
        v3 += -2.25
        _ = v5
    }
    fmt.Println(v2, v2)
    var v6 int32 = -1
    v1[2147483647] = 2147483648
    if false {
        fmt.Println(v3, v4, v2)
        for fi7 := 0; fi7 < 4; fi7++ {
            v1[fi7] = 961
            switch (("a" + "a") + "hello") {
            case "":
                v2 += 1.5
            }
            for fi8 := 0; fi8 < 17; fi8++ {
                fmt.Println(v6, fi7)
                switch "hello" {
                case "\n\t":
                    var v9 float32 = (1e10 - ((v3 + float32(v2)) - v3))
                    _ = v9
                }
                fmt.Println(v2, v2, v4)
                _ = fi8
            }
            _ = fi7
        }
    } else {
        switch "" {
        case "":
            v4 = true
        case "\n\t":
            fmt.Println(v3)
        case "foo bar":
            var v10 int32 = 1
            _ = v10
        }
    }
    var v11 int32 = v6
    v12 := GenS1{F0: -1}
    var v13 int8 = 127
    switch ("a" + "a") {
    case "hello":
        var v14 string = (("a" + "") + "a")
        _ = v14
    case "a":
        switch "\n\t" {
        case "hello":
            if (v4 && (30 <= 582)) {
                for fi15 := 0; fi15 < 20; fi15++ {
                    v16 := []int{(fi15 | (^(fi15))), fi15, int(v13)}
                    if v4 {
                        fmt.Println(v4)
                        v17 := v1[(fi15 & -2147483648)]
                        var v18 int64 = 0
                        _ = v17
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 18; fi19++ {
                            fmt.Println(fi15, v2, v2)
                            _ = fi19
                        }
                        switch "foo bar" {
                        case "\n\t":
                            v20 := GenS1{F0: 767}
                            _ = v20
                        default:
                            if v4 {
                                v13 += (v13 & (v13 * (v13 << ((18446744073709551615) & 63))))
                                fmt.Println(v4, v3)
                            } else {
                                v21 := v1[(fi15 - 32768)]
                                _ = v21
                            }
                        }
                        fmt.Println(v3)
                    }
                    _ = fi15
                    _ = v16
                }
            }
        }
    case "\n\t":
        var v22 uint8 = 128
        _ = v22
    }
    v11 = -71
    v11 -= (v11 & 169)
    var v23 string = "foo bar"
    fmt.Println(v11)
    v24 := GenS1{F0: -173}
    for fi25 := 0; fi25 < 3; fi25++ {
        if ((290 != 2) || v4) {
            v26 := GenS2{F0: 152}
            v27 := GenS1{F0: 1}
            fmt.Println(v6)
            _ = v26
            _ = v27
        } else {
            fmt.Println(v6)
        }
        _ = fi25
    }
    var v28 int16 = 575
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v6
}
