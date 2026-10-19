package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint16, p1 int64, p2 uint64) uint8 {
    v6 := GenS1{F0: (!(((!(65535 <= 255)) || (1 < 950)) || false))}
    v7 := []int{234, 4294967295, -2147483648}
    var v8 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 200
}

func main() {
    v1 := []int{391, 42, 32767, 42}
    var v2 float32 = 1.5
    var v3 float64 = float64((v2 / (v2 / (3.14159 - v2))))
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    if false {
        switch 32767 {
        case -207:
            v2 = v2
        case -1:
            v2 = v2
        case 42:
            switch "\n\t" {
            case "a":
                var v9 int64 = -1
                _ = v9
            }
        }
        fmt.Println(v3, v3)
        for fi10 := 0; fi10 < 17; fi10++ {
            fmt.Println(v3, v2)
            var v11 uint32 = 262
            v3 -= 0.1
            _ = fi10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        if (false || ((!(0 >= 18446744073709551615)) || false)) {
            switch (fi12 + 42) {
            case 128:
                var v13 bool = (((v3 + float64(float32((0.1 + v3)))) / 1e10) != (v3 - (v3 / v3)))
                _ = v13
            case 0:
                var v14 int = fi12
                _ = v14
            case 32768:
                delete(v5, (fi12 / ((fi12) | 1)))
            }
        }
        _ = fi12
    }
    var v15 uint8 = 38
    if true {
        v16 := GenS1{F0: false}
        _ = v16
    }
    if (!true) {
        if len(v1) > 0 {
            v1[0] = 0
        }
    }
    v1 = append(v1, 105)
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 uint64 = 42
        var v19 uint8 = v15
        _ = fi17
        _ = v18
        _ = v19
    }
    v3 -= float64(0.1)
    v20 := GenS1{F0: true}
    for fi21 := 0; fi21 < 7; fi21++ {
        v3 += v3
        _ = fi21
    }
    switch -891 {
    case -2147483648:
        fmt.Println(v15)
    case -9223372036854775808:
        v22 := GenS1{F0: (-9223372036854775808 > -322)}
        _ = v22
    }
    var v23 uint16 = 2
    _ = len(v1)
    switch "foo bar" {
    case "hello":
        if true {
            switch 1 {
            case -2147483648:
                switch 32767 {
                case 9223372036854775807:
                    var v24 uint16 = v23
                    _ = v24
                case 4294967295:
                    if false {
                        var v25 uint16 = v23
                        _ = v25
                    }
                case 32768:
                    v23 += v23
                default:
                    switch -151 {
                    case -9223372036854775808:
                        var v26 int32 = -811
                        _ = v26
                    case 404:
                        var v27 int16 = 1
                        _ = v27
                    }
                }
            case 4294967295:
                v28 := GenS1{F0: (((v3 - v3) - float64((0.1 - (v2 + (v2 / v2))))) < v3)}
                _ = v28
            case -59:
                v29 := GenS1{F0: true}
                _ = v29
            }
            fmt.Println(v15, v3)
        }
    default:
        v15 = (^(v15))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
