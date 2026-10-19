package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int32; F2 int16 }
type GenS4 struct { F0 int8; F1 int }
func genFn0(p0 string, p1 bool, p2 int) uint64 {
    if (!(p1 || (!((0.1 - (0.0 * (0.0 + (0.1 - 3.14159)))) >= (-2.25 - float32((-2.25 - -2.25))))))) {
        v4 := []int{(p2 ^ (p2 - 762)), (^(p2)), p2}
        _ = v4
    }
    var v5 int16 = 96
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}

func main() {
    var v1 float32 = 0.0
    var v2 bool = (65535 >= 2)
    var v3 uint = 0
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 989, F1: -9223372036854775808}
        v8 := GenS1{F0: 65535, F1: -1}
        var v9 int32 = 42
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 float32 = 1.5
    var v11 int16 = 32767
    v12 := GenS1{F0: 65535, F1: 1}
    var v13 int32 = 42
    v14 := GenS2{F0: 4294967295}
    var v15 float64 = 0.1
    for fi16 := 0; fi16 < 9; fi16++ {
        if (v13 > -1) {
            v17 := []int{1, int(v13), fi16}
            v10 = v1
            v1 += v1
            _ = v17
        }
        v10 = float32(((v15 - (v15 * 1e10)) * v15))
        _ = fi16
    }
    v1 += (v1 + v1)
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 bool = (!v2)
            var v21 float32 = float32(1.5)
            var v22 bool = false
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        switch -817 {
        case -2147483648:
            fmt.Println(v1, v11)
        case 2147483647:
            v23 := []int{}
            _ = v23
        default:
            v24 := []int{fi18, (fi18 >> ((v3) & 63)), fi18, 42}
            _ = v24
        }
        fmt.Println(v3)
        _ = fi18
    }
    v25 := GenS3{F0: 2, F1: (v13 << ((v3) & 63)), F2: v11}
    if (!v2) {
        if (v2 || (true || (true || ((^(v13)) != v13)))) {
            switch "hello" {
            case "hello":
                var v26 uint8 = 200
                _ = v26
            case "foo bar":
                var v27 uint8 = 128
                _ = v27
            case "a":
                var v28 int64 = 0
                _ = v28
            }
        } else {
            v11 -= (v11 * -4)
        }
    }
    var v29 bool = (-1 > 42)
    v30 := GenS4{F0: -1, F1: -1}
    var v31 int8 = 5
    if true {
        v2 = (v11 <= (v11 << ((v3) & 63)))
        var v32 uint32 = 2147483648
        _ = v32
    } else {
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
