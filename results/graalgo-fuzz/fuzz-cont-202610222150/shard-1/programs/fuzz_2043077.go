package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 string) uint32 {
    fmt.Println(p0)
    var v4 float64 = float64(1e10)
    _ = p0
    _ = v4
    return 2147483647
}
func genFn1(p0 int16) uint32 {
    for fi5 := 0; fi5 < 16; fi5++ {
        switch "foo bar" {
        case "foo bar":
            var v6 uint64 = 456
            _ = v6
        }
        _ = fi5
    }
    if false {
        for fi7 := 0; fi7 < 19; fi7++ {
            p0 += -32768
            _ = fi7
        }
        var v8 uint32 = 542
        _ = v8
    }
    if true {
        p0 -= p0
    } else {
        p0 = 32767
        fmt.Println(p0)
    }
    p0 = (p0 % ((42) | 1))
    _ = p0
    return 2
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 187
    v3 := make(map[string]uint32)
    var v9 string = ""
    v10 := []int{4294967295, -384}
    v9 = v9
    if (true || false) {
        v2 -= v2
    } else {
        v1 = v1
    }
    v11 := GenS1{F0: (false && false)}
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint8 = (^(v2))
        switch "a" {
        case "a":
            switch fi12 {
            case 32768:
                var v14 bool = true
                _ = v14
            }
        default:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        }
        _ = fi12
        _ = v13
    }
    v1 = v1
    v1 += v1
    v1 = v1
    var v15 bool = (2 >= 65535)
    if (960 < 18446744073709551615) {
        v15 = v15
    }
    if (68 == 47) {
        v2 -= v2
        v16 := GenS1{F0: true}
        v9 = v9
        _ = v16
    } else {
        var v17 int32 = 0
        var v18 int = 32767
        var v19 float32 = 1.5
        _ = v17
        _ = v18
        _ = v19
    }
    if (256 != 336) {
        v1 = 1
        var v20 bool = ((32767 > 32767) && ((float64(0.1) + (1.5 / 0.1)) > -2.25))
        _ = v20
    } else {
        v9 = ""
        v3[(v9 + (v9 + (v9 + "\n\t")))] = 2147483648
    }
    if (float32(float64((((0.1 - 3.14159) / 1.5) / 1e10))) > 1.5) {
        v9 = v9
        v21 := GenS1{F0: v15}
        _ = v21
    } else {
        v3["a"] = 944
        var v22 float32 = float32(1.5)
        var v23 int = 0
        _ = v22
        _ = v23
    }
    var v24 uint8 = (v2 ^ v2)
    var v25 float64 = ((float64((1.5 / 1e10)) / float64((1.5 + 3.14159))) * -2.25)
    v9 = v9
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v9
}
