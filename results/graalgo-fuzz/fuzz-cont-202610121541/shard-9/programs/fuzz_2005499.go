package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int }

func main() {
    var v1 float64 = 0.0
    var v2 uint32 = 550
    var v3 int = 4294967295
    var v4 int16 = -203
    var v5 int64 = 9223372036854775807
    v5 += (v5 ^ v5)
    v2 = (v2 - v2)
    var v6 uint64 = 0
    v7 := []int{4294967295, (-(v3))}
    var v8 int = v3
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v5 -= 13
    var v9 float32 = 1.5
    v5 -= -978
    switch 2147483647 {
    case 134:
        switch v8 {
        case 704:
            fmt.Println(v1)
        case -531:
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 uint8 = uint8(v6)
                for fi12 := 0; fi12 < 1; fi12++ {
                    if true {
                        v6 = v6
                    } else {
                        v4 -= v4
                        if (int16(v5) > v4) {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            switch (^(fi10)) {
                            case -388:
                                var v13 string = "\n\t"
                                _ = v13
                            default:
                                _ = len(v7)
                            }
                        } else {
                            var v14 int8 = -25
                            if (1 == (v6 << ((uint(v2)) & 63))) {
                                v7 = append(v7, (^(fi10)))
                            } else {
                                var v15 float32 = (v9 / v9)
                                _ = v15
                            }
                            _ = v14
                        }
                    }
                    switch (("" + "\n\t") + (("hello" + "foo bar") + "")) {
                    case "":
                        var v16 uint32 = 42
                        _ = v16
                    case "a":
                        v8 = fi10
                    }
                    _ = fi12
                }
                _ = fi10
                _ = v11
            }
        }
    case -422:
        v9 = 3.14159
    case 42:
        v6 += (v6 - 4294967295)
    default:
        if true {
            if true {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            v7 = append(v7, v8)
            v17 := GenS1{F0: v5, F1: v4, F2: (v8 / (((v3 - -1)) | 1))}
            _ = v17
        }
    }
    v18 := GenS1{F0: v5, F1: v4, F2: v8}
    var v19 float32 = ((float32(float64((v9 / (v9 * -2.25)))) - v9) - v9)
    v20 := GenS1{F0: (v5 - (v5 & int64(v4))), F1: -580, F2: (v3 % ((-979) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
