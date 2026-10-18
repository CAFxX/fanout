package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint8, p1 bool) int64 {
    v7 := []int{42, 1000}
    p0 = p0
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int64 = -529
        p0 = (p0 & 2)
        switch fi8 {
        case -660:
            v10 := GenS1{F0: (p0 + p0)}
            _ = v10
        case -97:
            v11 := GenS2{F0: ("foo bar" > "foo bar"), F1: 2, F2: 24}
            _ = v11
        default:
            v12 := GenS2{F0: (p1 && (!(102 > 400))), F1: 18446744073709551615, F2: 11}
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v13 int8 = -128
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return -1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    v3 := []string{"\n\t"}
    v4 := make(map[string]string)
    var v5 int32 = -2147483648
    var v6 bool = false
    var v14 int16 = 42
    switch 0 {
    case 2147483647:
        v15 := GenS3{F0: 2147483648}
        _ = v15
    default:
        v16 := GenS3{F0: 42}
        _ = v16
    }
    switch 32768 {
    case 203:
        fmt.Println(v14, v5, v14)
    case -222:
        for fi17 := 0; fi17 < 10; fi17++ {
            v4["foo bar"] = "foo bar"
            _ = fi17
        }
    case 449:
        for fi18 := 0; fi18 < 4; fi18++ {
            if false {
                fmt.Println(v5, fi18, v6)
            } else {
                var v19 uint64 = 0
                v20 := GenS2{F0: v6, F1: 522, F2: 0}
                _ = v19
                _ = v20
            }
            var v21 float32 = 3.14159
            fmt.Println(v21, fi18)
            _ = fi18
            _ = v21
        }
    default:
        var v22 uint = 78
        _ = v22
    }
    fmt.Println(v5, v6)
    _ = len(v3)
    fmt.Println(v6, v5)
    var v23 int16 = (^(v14))
    var v24 string = (("\n\t" + "hello") + "")
    switch 4294967295 {
    case 32767:
        var v25 float64 = 0.1
        _ = v25
    case -921:
        if (-2147483648 == 9223372036854775807) {
            var v26 float64 = (1.5 * (float64((3.14159 / 1e10)) - (3.14159 * float64(0.0))))
            _ = v26
        }
    default:
        var v27 uint64 = 1
        _ = v27
    }
    var v28 string = "a"
    v3 = append(v3, v28)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        var v30 int32 = (^(v5))
        v3 = append(v3, v24)
        _ = fi29
        _ = v30
    }
    fmt.Println(v5, v23)
    fmt.Println(v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
