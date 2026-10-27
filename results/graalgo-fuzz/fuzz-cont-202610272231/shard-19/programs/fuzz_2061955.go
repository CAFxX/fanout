package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 bool = (-2147483648 > 42)
    var v2 uint8 = 107
    v3 := []string{(("hello" + "a") + "\n\t"), "a", "", ("foo bar" + ((("a" + "a") + (("" + "a") + "a")) + ("\n\t" + "")))}
    v1 = (0.0 <= 1.5)
    var v4 float32 = 3.14159
    var v5 int8 = -73
    var v6 float64 = (((((float64(((v4 - v4) / v4)) - ((1.5 / 1.0) / 1e10)) - ((((float64(0.0) / 3.14159) / -2.25) + (float64(1e10) * 3.14159)) + 1e10)) * 1.5) * ((0.0 - 1.5) - 1e10)) / 3.14159)
    v2 = 128
    v3 = append(v3, "")
    v7 := GenS1{F0: -9223372036854775808}
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int16 = -32768
        v5 += v5
        var v10 uint32 = 456
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        switch (^(fi11)) {
        case 978:
            for fi12 := 0; fi12 < 20; fi12++ {
                v5 -= (v5 / (((v5 - v5)) | 1))
                _ = len(v3)
                _ = fi12
            }
        default:
            fmt.Println(v2, v5, v6)
        }
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        if v1 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 += (^(v2))
        } else {
            fmt.Println(v1, v2, v4)
            v13 := GenS1{F0: -36}
            v14 := GenS1{F0: -170}
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    v4 = ((0.1 * ((1e10 - v4) / (float32(0.1) - (3.14159 / ((1e10 * (((1.5 / (v4 * v4)) + v4) * 3.14159)) * (1.5 * (0.1 * 1e10))))))) / v4)
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int16 = 1
        v17 := GenS1{F0: -1}
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := GenS1{F0: -9223372036854775808}
        fmt.Println(v6, v1)
        v5 -= (v5 << ((255) & 63))
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 9223372036854775807}
    _ = len(v3)
    fmt.Println(v4, v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
