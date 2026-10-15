package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint32 }
func genFn0() string {
    v6 := []int{4294967295}
    if (!(2 <= 4294967295)) {
        if len(v6) > 0 {
            v6[0] = -1
        }
        switch -2147483648 {
        case -861:
            var v7 float64 = (float64((float32((1e10 + float64((0.0 + (1e10 - 0.1))))) - (3.14159 / 1.0))) / -2.25)
            _ = v7
        case -895:
            var v8 bool = (!(false && (3.14159 != (0.0 * 1.5))))
            _ = v8
        case 295:
            var v9 float32 = 1.5
            _ = v9
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v10 uint64 = 18446744073709551615
    _ = v10
    _ = v6
    return "a"
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]bool)
    var v3 uint32 = 42
    v4 := make(map[string]string)
    var v5 float32 = 3.14159
    var v11 uint8 = 233
    if (!true) {
        v3 -= v3
        if false {
            delete(v4, "foo bar")
            for fi12 := 0; fi12 < 10; fi12++ {
                fmt.Println(v11, v1, fi12)
                v11 = v11
                var v13 int = (^(fi12))
                _ = fi12
                _ = v13
            }
        }
        if ((true && false) || true) {
            v5 -= (v5 - 1e10)
            if true {
                v3 = v3
            } else {
                v11 = 87
                v5 -= (float32((3.14159 + 0.0)) / 1.5)
                if false {
                    v14 := GenS1{F0: uint(v11), F1: 1, F2: (^(v11))}
                    _ = v14
                } else {
                    v11 = (v11 % ((128) | 1))
                    for fi15 := 0; fi15 < 9; fi15++ {
                        switch -243 {
                        case 127:
                            v16 := GenS1{F0: (v1 & (^(v1))), F1: 801, F2: (v11 + 128)}
                            _ = v16
                        case 1:
                            delete(v4, "hello")
                        }
                        var v17 float64 = (1.5 / -2.25)
                        fmt.Println(v17, v11)
                        _ = fi15
                        _ = v17
                    }
                    switch 0 {
                    case -9223372036854775808:
                        fmt.Println(v5, v5, v5)
                    case 32767:
                        v18 := GenS1{F0: 1, F1: uint64(v11), F2: (v11 ^ v11)}
                        _ = v18
                    }
                }
            }
            v3 += v3
        }
    }
    v1 = (v1 + 18446744073709551615)
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 uint16 = 794
        _ = fi19
        _ = v20
    }
    var v21 int8 = 42
    if true {
        var v22 int32 = -83
        _ = v22
    }
    var v23 float64 = ((-2.25 + 1.5) + (1.5 - float64(v5)))
    var v24 int64 = -9223372036854775808
    fmt.Println(v3, v5)
    v1 = (^(v1))
    v25 := []int{}
    _ = len(v25)
    v3 -= v3
    fmt.Println(v11, v1, v24)
    var v26 int = 4294967295
    v23 = v23
    v27 := GenS2{F0: 65535, F1: 792, F2: (^(v3))}
    v5 += -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
