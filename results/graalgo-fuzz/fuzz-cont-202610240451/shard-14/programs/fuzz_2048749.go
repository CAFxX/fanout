package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint8; F2 int }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 int8 = 1
    var v4 int = 9223372036854775807
    v5 := make(map[int]string)
    var v6 uint32 = 0
    if false {
        v7 := GenS1{F0: 1, F1: true, F2: v6}
        _ = v7
    } else {
        var v8 float64 = 3.14159
        switch -224 {
        case 32767:
            fmt.Println(v6, v8)
        default:
            switch v4 {
            case 128:
                fmt.Println(v4, v3)
            default:
                fmt.Println(v6, v6, v4)
            }
        }
        _ = v8
    }
    switch "" {
    case "\n\t":
        delete(v2, "")
    case "":
        v3 = v3
    default:
        v9 := GenS1{F0: v3, F1: ((32767 <= 619) || (false || true)), F2: (v6 - v6)}
        _ = v9
    }
    v3 += v3
    var v10 uint16 = 65535
    v10 += 32767
    var v11 bool = false
    if true {
        if (("foo bar" + "\n\t") != ("hello" + ("hello" + "hello"))) {
            var v12 uint16 = 42
            v11 = (!(v11 || ((int32(v10) < -1) || (1 < 2))))
            v3 += (-(v3))
            _ = v12
        } else {
            var v13 int16 = 152
            v14 := v1[v4]
            var v15 bool = v11
            _ = v13
            _ = v14
            _ = v15
        }
        _ = len(v1)
    }
    var v16 int = (v4 + v4)
    v10 = (v10 - v10)
    var v17 int64 = -262
    fmt.Println(v4, v6, v11)
    v18 := GenS2{F0: 700, F1: 1, F2: (v4 / ((42) | 1))}
    v19 := GenS2{F0: 1000, F1: 211, F2: (v4 & (^(v16)))}
    v20 := GenS2{F0: 42, F1: uint8(v10), F2: 732}
    var v21 uint64 = 914
    if true {
        v17 += 1
        v5[v16] = "a"
    }
    var v22 int = v16
    for fi23 := 0; fi23 < 11; fi23++ {
        v21 += v21
        for fi24 := 0; fi24 < 20; fi24++ {
            if v11 {
                v25 := []int{fi24, (-(v22)), (fi23 >> ((255) & 63))}
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                v22 = (v16 % ((v4) | 1))
                _ = v25
            } else {
                if ((v4 - int(v6)) > fi23) {
                    v3 -= (v3 - v3)
                    v26 := GenS1{F0: (v3 & (v3 + 100)), F1: (("" < "foo bar") || (!((v6 == (v6 | v6)) || v11))), F2: v6}
                    var v27 int32 = 279
                    _ = v26
                    _ = v27
                }
            }
            v10 -= (v10 + v10)
            v28 := GenS2{F0: 1000, F1: 128, F2: (fi23 / ((int(v6)) | 1))}
            _ = fi24
            _ = v28
        }
        fmt.Println(v11, v10)
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
