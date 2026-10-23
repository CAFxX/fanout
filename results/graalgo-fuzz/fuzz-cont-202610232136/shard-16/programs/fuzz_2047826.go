package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint8 }

func main() {
    var v1 uint8 = 1
    v2 := []uint32{1, 65535}
    var v3 int64 = 9223372036854775807
    var v4 int64 = v3
    var v5 uint64 = 595
    v2 = append(v2, 643)
    v6 := GenS1{F0: (v4 << ((42) & 63)), F1: 522, F2: (v1 + (v1 >> ((0) & 63)))}
    for fi7 := 0; fi7 < 13; fi7++ {
        switch ("\n\t" + "hello") {
        case "a":
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 float64 = (float64(float32(3.14159)) - ((0.0 * float64(0.1)) - 1.5))
                fmt.Println(fi8, fi8)
                var v10 uint32 = 2
                _ = fi8
                _ = v10
                _ = v9
            }
        case "\n\t":
            var v11 int32 = -2147483648
            _ = v11
        case "hello":
            for fi12 := 0; fi12 < 17; fi12++ {
                v13 := GenS1{F0: v3, F1: 2, F2: v1}
                _ = fi12
                _ = v13
            }
        }
        if false {
            var v14 uint32 = 42
            _ = v14
        } else {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
                if true {
                    fmt.Println(v3)
                } else {
                    for fi15 := 0; fi15 < 8; fi15++ {
                        switch (fi15 | (-(fi15))) {
                        case 900:
                            fmt.Println(fi7, fi15)
                        default:
                            switch "hello" {
                            case "foo bar":
                                _ = len(v2)
                            case "":
                                v2 = append(v2, 0)
                            }
                        }
                        _ = fi15
                    }
                }
            }
            switch -352 {
            case 9223372036854775807:
                v5 = (v5 >> ((0) & 63))
            }
        }
        for fi16 := 0; fi16 < 2; fi16++ {
            fmt.Println(v4, v4, v3)
            fmt.Println(v3)
            fmt.Println(fi16, fi16, v5)
            _ = fi16
        }
        _ = fi7
    }
    v1 = v1
    v3 = -981
    if false {
        v2 = append(v2, 2)
    }
    var v17 int16 = -1
    switch 569 {
    case 876:
        if ((v5 * 18446744073709551615) >= (v5 >> ((42) & 63))) {
            var v18 uint = 65535
            _ = v18
        }
    case 32768:
        v4 = v3
    case 0:
        _ = len(v2)
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v19 uint16 = 1
    v19 = (v19 - (v19 + v19))
    if false {
        var v20 float32 = ((((1e10 / (-2.25 * ((0.1 - float32(float64(-2.25))) / float32(0.1)))) / 1.0) + 3.14159) + 1.5)
        if len(v2) > 0 {
            v2[0] = 30
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v20
    }
    fmt.Println(v3, v4)
    v1 += v1
    v17 -= v17
    fmt.Println(v17, v19)
    v19 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
