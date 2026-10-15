package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    var v3 string = "foo bar"
    var v4 int8 = 1
    var v5 int64 = 432
    v6 := GenS1{F0: 0}
    switch "foo bar" {
    case "foo bar":
        v4 -= v4
    case "a":
        v4 += v4
    default:
        switch 2147483647 {
        case 127:
            v7 := GenS1{F0: -61}
            _ = v7
        }
    }
    switch 32767 {
    case 2147483647:
        for fi8 := 0; fi8 < 16; fi8++ {
            if (v4 < (v4 & v4)) {
                var v9 int32 = -2147483648
                _ = v9
            }
            _ = fi8
        }
    case -1:
        v10 := GenS1{F0: 513}
        _ = v10
    default:
        v5 = v5
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        if (2 > 0) {
            fmt.Println(v5, v3, v4)
            v12 := GenS1{F0: 696}
            v13 := []int{fi11, -4, (fi11 - (fi11 | fi11))}
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 float64 = ((0.1 - 3.14159) + -2.25)
                _ = fi14
                _ = v15
            }
            v3 = (v3 + v3)
            v3 = "hello"
        }
        if false {
            v4 = v4
            fmt.Println(v5, v5, fi11)
        } else {
            v5 -= -9223372036854775808
            var v16 int64 = v5
            _ = v16
        }
        _ = fi11
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        delete(v1, "hello")
    default:
        v4 = v4
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        v5 += (v5 ^ 42)
        var v18 uint32 = 2
        _ = fi17
        _ = v18
    }
    v3 = "\n\t"
    fmt.Println(v5, v4)
    fmt.Println(v4, v4, v5)
    v19 := GenS2{F0: 255, F1: 453, F2: 555}
    v3 = (((v3 + v3) + ("hello" + v3)) + "\n\t")
    fmt.Println(v4, v3, v4)
    for fi20 := 0; fi20 < 20; fi20++ {
        v5 -= v5
        _ = fi20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 uint8 = 255
        var v23 uint = 65535
        v24 := GenS3{F0: 18446744073709551615, F1: 4294967295, F2: -109}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
