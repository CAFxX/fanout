package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int64, p1 uint8) uint16 {
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 896}
    for fi10 := 0; fi10 < 18; fi10++ {
        switch 2147483647 {
        case 9223372036854775807:
            var v11 uint32 = 0
            _ = v11
        case -2147483648:
            p0 -= (p0 / ((42) | 1))
        case 127:
            var v12 uint64 = 18446744073709551615
            _ = v12
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return 1
}

func main() {
    var v1 int32 = -2147483648
    var v2 float32 = (1e10 + 1e10)
    v3 := make(map[int]uint32)
    var v4 float64 = 0.1
    var v5 int32 = 2147483647
    v6 := make(map[uint32]uint32)
    v13 := GenS1{F0: 2147483647}
    v14 := GenS1{F0: (-(v5))}
    v15 := []int{}
    v1 -= v1
    var v16 uint = 2
    if (!(!true)) {
        v5 = v1
        var v17 uint64 = 0
        _ = v17
    } else {
        var v18 float64 = float64(float32(-2.25))
        _ = v18
    }
    fmt.Println(v16, v2, v4)
    switch 42 {
    case 9223372036854775807:
        v19 := GenS1{F0: (-(v1))}
        _ = v19
    default:
        v16 -= v16
    }
    switch "a" {
    case "\n\t":
        fmt.Println(v2, v16, v1)
    case "a":
        var v20 uint = 255
        _ = v20
    }
    switch -449 {
    case 32767:
        if (false && true) {
            v15 = append(v15, 32768)
            var v21 float64 = (-2.25 - v4)
            _ = v21
        }
    case -48:
        switch -124 {
        case 128:
            v22 := GenS1{F0: 0}
            _ = v22
        case -2147483648:
            if true {
                var v23 float64 = (v4 - (float64(0.0) + v4))
                var v24 int32 = v1
                if false {
                    v25 := GenS1{F0: (-(v1))}
                    v26 := GenS1{F0: (v24 << (((^(v16))) & 63))}
                    var v27 uint8 = 127
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = v23
                _ = v24
            }
        case 4294967295:
            v28 := GenS1{F0: (v1 >> (((v16 * (v16 % (((v16 | 153)) | 1)))) & 63))}
            _ = v28
        }
    default:
        var v29 int8 = 127
        _ = v29
    }
    var v30 int32 = v5
    var v31 float64 = (((float64((v2 / ((float32(3.14159) / v2) + v2))) / ((v4 + float64((((float32(v4) + (((float32((v4 / v4)) + v2) + ((((v2 / (1e10 - 0.1)) * (v2 * v2)) + (v2 - float32(v4))) * 0.0)) - v2)) - v2) / v2))) - 3.14159)) - float64(1e10)) / v4)
    if false {
        var v32 int32 = v30
        _ = v32
    } else {
        var v33 float32 = -2.25
        _ = v33
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 uint16 = genFn0(-9223372036854775808, 59)
        v30 += (v5 / ((v30) | 1))
        for fi36 := 0; fi36 < 9; fi36++ {
            v37 := GenS1{F0: (v5 % ((v1) | 1))}
            var v38 uint8 = 42
            v16 = uint(v35)
            v39 := GenS2{F0: 100}
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        _ = fi34
        _ = v35
    }
    switch "hello" {
    case "\n\t":
        var v40 string = ""
        _ = v40
    case "hello":
        if true {
            if true {
                fmt.Println(v5, v4, v4)
            }
        } else {
            v2 = -2.25
            v41 := GenS2{F0: -79}
            _ = v41
        }
    case "":
        v42 := GenS2{F0: 124}
        _ = v42
    }
    for fi43 := 0; fi43 < 3; fi43++ {
        var v44 int16 = -656
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
