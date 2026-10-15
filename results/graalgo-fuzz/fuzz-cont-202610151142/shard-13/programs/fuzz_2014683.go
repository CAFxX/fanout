package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint16) int8 {
    switch "a" {
    case "":
        for fi7 := 0; fi7 < 8; fi7++ {
            switch (fi7 % ((42) | 1)) {
            case 127:
                for fi8 := 0; fi8 < 5; fi8++ {
                    p1 = 255
                    p0 = 42
                    _ = fi8
                }
            case -9223372036854775808:
                var v9 int16 = -1
                _ = v9
            case 554:
                p1 = p1
            }
            fmt.Println(fi7)
            _ = fi7
        }
    case "hello":
        switch "\n\t" {
        case "":
            p0 = (p0 >> ((4294967295) & 63))
        case "hello":
            v10 := []int{127, 32768, 32768}
            _ = v10
        }
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v11
    return 97
}

func main() {
    var v1 int32 = -695
    v2 := make(map[int]bool)
    v3 := []uint32{4294967295, 4294967295}
    var v4 float64 = float64(float32(float64(3.14159)))
    v5 := make(map[uint32]bool)
    var v6 bool = (!false)
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 float64 = (v4 + float64(3.14159))
        var v14 uint64 = 1
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v6)
    v15 := GenS1{F0: 18446744073709551615}
    var v16 float32 = 3.14159
    _ = len(v3)
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := GenS1{F0: 92}
        _ = fi17
        _ = v18
    }
    if (!(!(true && false))) {
        var v19 float32 = 0.0
        _ = v19
    } else {
        delete(v2, -21)
        switch ("a" + ("" + "")) {
        case "hello":
            switch -1 {
            case 42:
                v1 -= v1
            default:
                var v20 bool = true
                _ = v20
            }
        }
        v16 += ((v16 - ((0.0 / float32(float64((0.1 * 0.1)))) - float32((v4 - (v4 / (v4 * (1e10 - (1.5 - v4)))))))) / v16)
    }
    if true {
        v16 = 1.5
        v1 = v1
    } else {
        for fi21 := 0; fi21 < 15; fi21++ {
            if v6 {
                var v22 string = "hello"
                _ = v22
            }
            for fi23 := 0; fi23 < 1; fi23++ {
                v3 = append(v3, 268)
                _ = fi23
            }
            for fi24 := 0; fi24 < 13; fi24++ {
                fmt.Println(v16, v16)
                _ = fi24
            }
            _ = fi21
        }
        var v25 int16 = 42
        if (!true) {
            for fi26 := 0; fi26 < 16; fi26++ {
                v27 := GenS1{F0: 417}
                _ = len(v5)
                _ = fi26
                _ = v27
            }
        }
        _ = v25
    }
    var v28 int64 = -1
    v28 -= v28
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
