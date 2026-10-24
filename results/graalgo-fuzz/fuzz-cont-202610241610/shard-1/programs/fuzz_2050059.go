package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
func genFn0(p0 uint64) uint32 {
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := GenS1{F0: -32768, F1: 255, F2: fi4}
        v6 := GenS2{F0: 819}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint8 = 175
    var v8 bool = false
    switch -438 {
    case 2147483647:
        var v9 uint = 630
        _ = v9
    case -889:
        var v10 uint64 = p0
        _ = v10
    case -2147483648:
        switch "hello" {
        case "foo bar":
            v7 = 128
        }
    default:
        switch 9223372036854775807 {
        case 32767:
            var v11 uint64 = (p0 ^ 42)
            _ = v11
        case 855:
            switch 161 {
            case 127:
                for fi12 := 0; fi12 < 20; fi12++ {
                    v13 := []int{fi12, fi12}
                    var v14 int = 0
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case 2147483647:
                v8 = (!v8)
            case 1:
                if v8 {
                    v15 := []int{}
                    _ = v15
                }
            }
        case 712:
            var v16 float32 = 3.14159
            _ = v16
        default:
            v7 += (v7 << ((104) & 63))
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 bool = (!false)
    var v2 bool = (((((0.0 + 1.5) / float64((-2.25 - 0.1))) - ((3.14159 - (1.5 * float64(1.5))) / float64((1e10 - ((-2.25 - -2.25) * (0.0 - -2.25)))))) <= float64((0.1 * ((((1.5 - (3.14159 - float32(0.0))) / 1.0) - -2.25) - (3.14159 - (float32(0.1) * (-2.25 / float32(-2.25)))))))) && (-1 > 0))
    var v3 uint8 = 127
    v3 -= v3
    if v1 {
        fmt.Println(v3, v2)
        fmt.Println(v2, v1)
    } else {
        var v17 uint = 606
        var v18 float64 = (1e10 - (1e10 / (((1.5 - 3.14159) + 0.0) / (((1.5 - 1.5) + 1.5) / 3.14159))))
        if (0 != 42) {
            v1 = v2
            if v2 {
                if true {
                    var v19 uint32 = 42
                    v17 -= (^(v17))
                    _ = v19
                } else {
                    v20 := GenS2{F0: 2147483647}
                    var v21 uint = v17
                    _ = v20
                    _ = v21
                }
            }
        }
        _ = v17
        _ = v18
    }
    var v22 int64 = 1
    if v1 {
        var v23 string = "foo bar"
        _ = v23
    }
    var v24 float64 = float64(float32(3.14159))
    if (!(float64(1e10) <= (((v24 - v24) * v24) - v24))) {
        if v2 {
            fmt.Println(v2, v22)
            v25 := []int{-634, -126, -2147483648, 215}
            _ = v25
        }
        var v26 int32 = 1
        _ = v26
    } else {
        v27 := GenS2{F0: 336}
        fmt.Println(v1)
        v3 -= (v3 & 0)
        _ = v27
    }
    v28 := GenS3{F0: 128}
    var v29 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v28
    _ = v29
    _ = v3
}
