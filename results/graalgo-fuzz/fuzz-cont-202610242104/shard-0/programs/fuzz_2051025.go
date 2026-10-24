package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 18446744073709551615
    var v3 int64 = 1
    v3 = v3
    fmt.Println(v2, v3)
    v4 := v1[""]
    var v5 int = 42
    if v4 {
        for fi6 := 0; fi6 < 19; fi6++ {
            if v4 {
                fmt.Println(fi6)
            }
            _ = fi6
        }
        if v4 {
            v7 := GenS1{F0: (v3 << (((v2 % ((4294967295) | 1))) & 63))}
            _ = v7
        }
        v8 := GenS1{F0: (^(v3))}
        _ = v8
    } else {
        switch v5 {
        case -228:
            if false {
                v3 += -669
            } else {
                v3 -= v3
                v9 := GenS1{F0: (v3 ^ 0)}
                switch int(v2) {
                case -1:
                    var v10 float64 = (((float64((((0.0 * 1e10) * -2.25) * 1.5)) + 0.1) + 0.1) * -2.25)
                    _ = v10
                default:
                    v11 := GenS1{F0: 42}
                    _ = v11
                }
                _ = v9
            }
        case 127:
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 string = ("\n\t" + ("hello" + ("a" + "\n\t")))
                _ = fi12
                _ = v13
            }
        case 32768:
            var v14 float32 = float32(0.0)
            _ = v14
        }
    }
    var v15 int = -419
    var v16 int16 = -63
    var v17 uint8 = 42
    for fi18 := 0; fi18 < 19; fi18++ {
        if v4 {
            v17 += 0
        } else {
            var v19 float64 = (((float64(0.0) * float64(float32(float64(0.0)))) - -2.25) - 1.5)
            switch (v5 >> ((4294967295) & 63)) {
            case 32767:
                v20 := []int{v5, fi18, (v15 + (fi18 & (^(fi18))))}
                _ = v20
            default:
                if (0 == 42) {
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 float32 = (((0.1 - float32(float64((0.0 + (0.1 - 3.14159))))) / float32(v19)) * (0.1 * 0.0))
                        if (643 > 564) {
                            fmt.Println(v16, v22)
                            var v23 int = (fi18 & (fi21 + 42))
                            _ = v23
                        }
                        v24 := GenS1{F0: 0}
                        _ = fi21
                        _ = v22
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 18; fi25++ {
                        var v26 uint16 = 0
                        var v27 int64 = (v3 / (((v3 >> (((v2 % ((v2) | 1))) & 63))) | 1))
                        if v4 {
                            fmt.Println(v19)
                            var v28 float32 = -2.25
                            _ = v28
                        }
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    v16 = v16
                }
            }
            _ = v19
        }
        if (v4 && (0.0 <= 3.14159)) {
            if (1 < 777) {
                if v4 {
                    var v29 int8 = -128
                    var v30 uint32 = 2
                    _ = v29
                    _ = v30
                }
            } else {
                for fi31 := 0; fi31 < 19; fi31++ {
                    switch "" {
                    case "foo bar":
                        v1[((((("a" + "") + "a") + "hello") + "") + "a")] = ((0.1 * -2.25) < (3.14159 + (3.14159 / (0.1 + ((float64(0.1) / (float64(3.14159) * -2.25)) * (float64(0.1) + ((0.0 * (1.5 / 1.0)) - 3.14159)))))))
                    default:
                        var v32 uint8 = v17
                        _ = v32
                    }
                    _ = fi31
                }
            }
            v33 := []int{-9223372036854775808, (fi18 + int(v17)), v15, v15}
            _ = v33
        }
        if (-1 <= 1) {
            v4 = ((false && v4) || false)
            _ = len(v1)
            var v34 uint8 = v17
            _ = v34
        }
        _ = fi18
    }
    var v35 int = v5
    v3 -= v3
    var v36 int32 = int32(v17)
    _ = len(v1)
    fmt.Println(v16)
    fmt.Println(v36, v2)
    v17 = (v17 / ((v17) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
