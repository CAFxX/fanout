package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int64; F1 int; F2 uint8 }
func genFn0(p0 int16, p1 int64, p2 int64) int64 {
    if false {
        var v5 string = ""
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := GenS1{F0: false, F1: 362, F2: 1}
            switch "foo bar" {
            case "hello":
                v8 := []int{fi6}
                _ = v8
            }
            var v9 int = 4294967295
            _ = fi6
            _ = v7
            _ = v9
        }
        var v10 float64 = 1.5
        _ = v10
        _ = v5
    } else {
        v11 := GenS1{F0: true, F1: 18446744073709551615, F2: -36}
        fmt.Println(p1, p0)
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int = (fi12 >> ((4294967295) & 63))
        if (0 <= 255) {
            var v14 uint = 42
            v15 := GenS1{F0: (p2 <= p2), F1: (v14 % (((v14 - v14)) | 1)), F2: -2147483648}
            v16 := []int{(v13 << (((v14 + (v14 << ((43) & 63)))) & 63)), (-(v13)), (v13 + -993)}
            _ = v14
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 16; fi17++ {
                v13 = (v13 * v13)
                _ = fi17
            }
        }
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int64 = 0
    var v3 bool = false
    var v4 bool = v3
    v1 -= 65535
    v18 := GenS2{F0: 128, F1: -1, F2: (v1 | 885)}
    v19 := GenS2{F0: 32768, F1: -622, F2: v1}
    v20 := GenS3{F0: ("a" + "\n\t"), F1: 42}
    for fi21 := 0; fi21 < 3; fi21++ {
        if (true || (42 <= 18446744073709551615)) {
            for fi22 := 0; fi22 < 9; fi22++ {
                v2 = genFn0(32767, 1, v2)
                _ = fi22
            }
        } else {
            for fi23 := 0; fi23 < 9; fi23++ {
                v24 := GenS4{F0: (v2 ^ v2), F1: int(v2), F2: 127}
                fmt.Println(v1, fi23)
                v1 += (v1 | 1)
                _ = fi23
                _ = v24
            }
        }
        _ = fi21
    }
    v2 += 1
    switch -1 {
    case -2147483648:
        for fi25 := 0; fi25 < 2; fi25++ {
            var v26 string = "\n\t"
            _ = fi25
            _ = v26
        }
    case 330:
        var v27 float32 = 1.5
        _ = v27
    default:
        var v28 int32 = -2147483648
        _ = v28
    }
    v4 = true
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
