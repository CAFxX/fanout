package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint16 }
func genFn0() int16 {
    if (float32((0.0 + (0.1 - (0.0 / 1.5)))) <= (0.0 - (-2.25 / 0.1))) {
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 int = fi6
                fmt.Println(fi7)
                v9 := []int{fi6, 1}
                v8 -= (fi7 >> ((18) & 63))
                _ = fi7
                _ = v8
                _ = v9
            }
            if (("" + "foo bar") <= "") {
                var v10 bool = (!(-2.25 != -2.25))
                v10 = v10
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 uint32 = 42
                    _ = fi11
                    _ = v12
                }
                _ = v10
            } else {
                var v13 uint = 950
                var v14 uint16 = 919
                _ = v13
                _ = v14
            }
            var v15 int8 = -128
            _ = fi6
            _ = v15
        }
    } else {
        var v16 uint8 = 128
        _ = v16
    }
    v17 := []int{848}
    var v18 int64 = 9223372036854775807
    if false {
        v18 += v18
    }
    _ = v17
    _ = v18
    return 42
}
func genFn1(p0 int32, p1 uint64, p2 uint8) uint32 {
    fmt.Println(p2)
    p0 = p0
    p2 = 0
    _ = p0
    _ = p1
    _ = p2
    return 2147483648
}

func main() {
    var v1 uint = 4294967295
    v2 := []int{-9223372036854775808, 42, -75}
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]uint32)
    var v5 int64 = 0
    v19 := GenS1{F0: int8(v5), F1: 42, F2: v3}
    var v20 float32 = 0.1
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS1{F0: 1, F1: -1, F2: (v3 * v3)}
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 uint64 = v3
            var v25 uint64 = 959
            switch fi23 {
            case 2147483647:
                v26 := GenS1{F0: -128, F1: -1, F2: v25}
                _ = v26
            case 0:
                var v27 float64 = (((3.14159 * 1.5) - 0.0) * (float64(0.0) / 1e10))
                _ = v27
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    v28 := GenS1{F0: -1, F1: -114, F2: v3}
    var v29 int = 32767
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 uint8 = uint8(v3)
        _ = fi30
        _ = v31
    }
    v20 = v20
    switch v29 {
    case 0:
        v2 = append(v2, v29)
    case 407:
        if (!(65535 >= 2)) {
            v32 := GenS2{F0: 255, F1: false, F2: 629}
            switch "hello" {
            case "\n\t":
                var v33 int8 = -128
                _ = v33
            }
            var v34 int = v29
            _ = v32
            _ = v34
        }
    case 128:
        v35 := GenS1{F0: 51, F1: 56, F2: (v3 % (((v3 << ((4294967295) & 63))) | 1))}
        _ = v35
    default:
        _ = len(v4)
    }
    if (!(false && true)) {
        for fi36 := 0; fi36 < 3; fi36++ {
            if true {
                v29 = -9223372036854775808
            } else {
                for fi37 := 0; fi37 < 1; fi37++ {
                    var v38 int = -1
                    _ = fi37
                    _ = v38
                }
            }
            switch "hello" {
            case "hello":
                var v39 int64 = (v5 - (v5 >> (((v1 | (v1 + (v1 * v1)))) & 63)))
                _ = v39
            case "foo bar":
                if false {
                    v5 = v5
                    var v40 int32 = -2147483648
                    fmt.Println(v29, v40)
                    _ = v40
                }
            }
            v20 += v20
            _ = fi36
        }
        v41 := GenS3{F0: 65535}
        _ = v41
    } else {
        _ = len(v2)
        for fi42 := 0; fi42 < 3; fi42++ {
            for fi43 := 0; fi43 < 7; fi43++ {
                v29 = fi43
                delete(v4, 42)
                _ = fi43
            }
            _ = fi42
        }
        var v44 string = ("\n\t" + "a")
        _ = v44
    }
    var v45 int32 = -784
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
