package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
func genFn0(p0 uint16) uint16 {
    switch 1 {
    case 128:
        for fi5 := 0; fi5 < 1; fi5++ {
            if true {
                if (!(1 == -128)) {
                    if (false && false) {
                        v6 := GenS1{F0: -1, F1: -318, F2: (p0 & (p0 << ((255) & 63)))}
                        _ = v6
                    }
                }
            } else {
                p0 = 65535
                for fi7 := 0; fi7 < 20; fi7++ {
                    var v8 int = fi7
                    var v9 float64 = 0.1
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
            var v10 float64 = (1.5 * 0.1)
            _ = fi5
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: 32767, F1: 1, F2: (p0 % (((p0 ^ p0)) | 1))}
        _ = fi11
        _ = v12
    }
    _ = p0
    return (^(p0))
}
func genFn1(p0 bool, p1 bool, p2 bool) int {
    var v13 bool = p0
    v13 = p1
    v14 := []int{957, -275, -2147483648}
    if len(v14) > 0 {
        v14[0] = 32768
    }
    p2 = ((0.1 / 1e10) == 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 161
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]string)
    var v3 uint = 42
    var v4 float32 = 0.0
    if false {
        for fi15 := 0; fi15 < 12; fi15++ {
            switch -261 {
            case 32767:
                fmt.Println(v3, fi15)
            case 127:
                if true {
                    fmt.Println(v4, fi15, v3)
                    fmt.Println(v3, v4, v1)
                }
            default:
                v1 = float64(0.1)
            }
            switch "a" {
            case "\n\t":
                v16 := GenS2{F0: -1, F1: 65535, F2: ""}
                _ = v16
            case "foo bar":
                var v17 uint64 = 143
                _ = v17
            }
            switch (fi15 % ((fi15) | 1)) {
            case 32767:
                switch int(v3) {
                case 128:
                    v3 += v3
                default:
                    fmt.Println(v3, v1, v1)
                }
            case 32768:
                v18 := []int{fi15, fi15, 799}
                _ = v18
            }
            _ = fi15
        }
        var v19 float32 = (v4 / float32(((((v1 / (v1 * float64(v4))) - ((v1 * 1e10) + v1)) / 1.0) * (float64((((1.5 - v4) / v4) - -2.25)) / v1))))
        _ = v19
    }
    v1 -= v1
    var v20 uint = v3
    fmt.Println(v20, v1)
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 uint8 = 92
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 uint = (v3 & 1)
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    var v25 uint16 = 255
    switch ("foo bar" + "hello") {
    case "":
        var v26 uint32 = 2147483648
        _ = v26
    case "foo bar":
        var v27 float64 = ((v1 / v1) * 1e10)
        _ = v27
    default:
        fmt.Println(v3, v25, v1)
    }
    v28 := []int{-2, genFn1((3.14159 <= v1), (true || true), false), -9223372036854775808, 575}
    for fi29 := 0; fi29 < 20; fi29++ {
        v2[2147483647] = "\n\t"
        v3 = (v3 * 1)
        _ = fi29
    }
    v1 += (v1 * (v1 * float64(v4)))
    fmt.Println(v20)
    v25 += (v25 / ((genFn0((v25 - 32767))) | 1))
    v20 = v20
    var v30 uint32 = 2
    switch genFn1(true, (false || (false && false)), (-331 < 1)) {
    case 32767:
        fmt.Println(v20, v1, v20)
    case -583:
        for fi31 := 0; fi31 < 11; fi31++ {
            v32 := v2[(^(v30))]
            var v33 int8 = 0
            delete(v2, 0)
            _ = fi31
            _ = v32
            _ = v33
        }
    case -9223372036854775808:
        switch ("a" + "") {
        case "foo bar":
            fmt.Println(v30, v1, v4)
        case "\n\t":
            fmt.Println(v4, v30)
        case "hello":
            _ = len(v28)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v28
    _ = v3
    _ = v30
    _ = v4
}
