package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint; F2 int16 }
type GenS3 struct { F0 uint8; F1 uint32; F2 bool }
type GenS4 struct { F0 uint; F1 int8 }
func genFn0(p0 uint16, p1 uint16, p2 bool) int16 {
    p2 = p2
    for fi5 := 0; fi5 < 4; fi5++ {
        switch (fi5 + fi5) {
        case -1:
            var v6 string = "foo bar"
            _ = v6
        case 2147483647:
            p1 = p1
        }
        fmt.Println(p0, p0, p0)
        _ = fi5
    }
    v7 := []int{128, 2147483647, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -687
}

func main() {
    var v1 uint16 = 2
    var v2 uint = 2
    var v3 uint32 = 590
    v4 := make(map[int]bool)
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: (v2 + 4294967295)}
        var v10 int32 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 -= 2
    fmt.Println(v2, v2)
    var v11 bool = false
    var v12 bool = (((v1 / ((v1) | 1)) < v1) && v11)
    fmt.Println(v2)
    v13 := GenS1{F0: 0}
    v1 = (v1 * 255)
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 uint8 = 42
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: -1, F1: v2, F2: 42}
    var v17 string = ("hello" + "foo bar")
    for fi18 := 0; fi18 < 20; fi18++ {
        delete(v4, (fi18 | -520))
        switch (fi18 | fi18) {
        case -358:
            if (!(42 != genFn0(v1, v1, false))) {
                if false {
                    var v19 int32 = 817
                    for fi20 := 0; fi20 < 1; fi20++ {
                        v19 += v19
                        v11 = (!(1 == 127))
                        v1 = v1
                        _ = fi20
                    }
                    _ = v19
                }
                for fi21 := 0; fi21 < 11; fi21++ {
                    v22 := v4[(fi21 * int(v3))]
                    v12 = v12
                    fmt.Println(v17)
                    _ = fi21
                    _ = v22
                }
                v23 := GenS3{F0: 45, F1: (v3 / ((2) | 1)), F2: true}
                _ = v23
            }
        case 127:
            v24 := []int{fi18, (^(fi18))}
            _ = v24
        }
        v2 += v2
        _ = fi18
    }
    v17 = v17
    if v12 {
        for fi25 := 0; fi25 < 15; fi25++ {
            for fi26 := 0; fi26 < 11; fi26++ {
                v27 := []int{}
                v11 = (139 >= 65535)
                _ = fi26
                _ = v27
            }
            fmt.Println(v11, v3)
            _ = fi25
        }
    } else {
        var v28 int8 = -5
        _ = v28
    }
    if (42 > -9223372036854775808) {
        for fi29 := 0; fi29 < 18; fi29++ {
            if (((v1 << (((v2 / (((v2 & 743)) | 1))) & 63)) <= (v1 << (((v2 - (v2 << ((uint(v3)) & 63)))) & 63))) || true) {
                var v30 int16 = 42
                _ = v30
            } else {
                switch "a" {
                case "\n\t":
                    v31 := GenS4{F0: v2, F1: 42}
                    _ = v31
                case "hello":
                    v32 := v4[1]
                    _ = v32
                case "":
                    v11 = false
                default:
                    fmt.Println(v2, fi29, v12)
                }
            }
            _ = fi29
        }
    } else {
        var v33 float64 = (0.1 / float64(((float32(0.0) + (((1.5 + 1e10) / ((-2.25 / (1.5 / 3.14159)) / ((((0.0 - 0.0) * 0.1) - ((0.1 - 3.14159) + ((0.0 / 3.14159) * -2.25))) / (float32(0.1) - (-2.25 * (((float32((1e10 + 0.1)) - 0.0) / -2.25) + 1e10)))))) - (1e10 + 0.1))) + 0.1)))
        _ = v33
    }
    var v34 uint64 = 2
    fmt.Println(v1)
    var v35 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
