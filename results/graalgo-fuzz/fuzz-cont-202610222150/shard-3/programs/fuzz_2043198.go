package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint64, p1 string, p2 int32) uint16 {
    var v4 float32 = 0.1
    v5 := GenS1{F0: 640}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return uint16(p0)
}
func genFn1(p0 int32, p1 int16, p2 int) uint {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 bool = ((!false) && false)
        v8 := GenS1{F0: (-(fi6))}
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 -= (p0 | p0)
    p0 = (p0 % ((p0) | 1))
    v9 := GenS1{F0: 4294967295}
    switch p2 {
    case 4294967295:
        p2 += (p2 >> ((65535) & 63))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    var v1 float64 = (3.14159 + ((1.5 * (((1e10 + float64((((0.0 + 1e10) / (1.5 + 0.1)) + 0.1))) - (((((0.0 / -2.25) * 1e10) + 0.0) + 1.5) / 3.14159)) / ((1e10 * 0.0) - (1.5 + (float64((3.14159 - (1.5 + 1.5))) + (float64((-2.25 * ((((-2.25 + ((float32(1e10) - 0.0) + (3.14159 + 3.14159))) * (3.14159 + -2.25)) - float32(-2.25)) + 1.5))) / (1e10 - -2.25))))))) * (0.1 - float64(-2.25))))
    var v2 uint32 = 42
    v3 := make(map[string]uint32)
    v2 = (v2 >> ((2) & 63))
    switch 2147483647 {
    case 32768:
        for fi10 := 0; fi10 < 12; fi10++ {
            _ = len(v3)
            for fi11 := 0; fi11 < 19; fi11++ {
                v2 -= v2
                delete(v3, ("" + ("foo bar" + "\n\t")))
                v12 := GenS1{F0: (fi10 * (fi11 * (fi10 / (((fi11 & fi10)) | 1))))}
                v1 += ((((1e10 / v1) + v1) / ((v1 * (float64(3.14159) - 1e10)) - -2.25)) - ((v1 / v1) / ((v1 - 1.5) * 1.5)))
                v13 := []int{fi10, fi11, (-(fi10))}
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    default:
        var v14 bool = true
        _ = v14
    }
    fmt.Println(v1, v1)
    switch -1 {
    case -1:
        var v15 uint16 = 1
        _ = v15
    }
    var v16 int64 = 0
    v2 += v2
    if (((127 < 1) || (!true)) && (true || false)) {
        v17 := GenS1{F0: -590}
        var v18 int16 = -32768
        fmt.Println(v2, v2)
        _ = v17
        _ = v18
    } else {
        var v19 int16 = 32767
        var v20 bool = true
        _ = v19
        _ = v20
    }
    if false {
        _ = len(v3)
    }
    v16 = (v16 + -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
}
