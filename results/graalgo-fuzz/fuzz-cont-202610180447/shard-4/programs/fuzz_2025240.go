package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int16 }
type GenS2 struct { F0 int32; F1 int64; F2 int64 }
type GenS3 struct { F0 uint16; F1 bool; F2 int64 }
type GenS4 struct { F0 int32 }
func genFn0() uint16 {
    if false {
        var v4 uint32 = 321
        var v5 int16 = 0
        v6 := []int{128, 32768}
        _ = v4
        _ = v5
        _ = v6
    } else {
        if (true || (!(1 != 18446744073709551615))) {
            for fi7 := 0; fi7 < 5; fi7++ {
                for fi8 := 0; fi8 < 17; fi8++ {
                    var v9 uint = 18446744073709551615
                    v9 += v9
                    v10 := []int{fi8, (-(fi8)), 32767, (fi7 % ((4294967295) | 1))}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                switch "" {
                case "\n\t":
                    if (1 > 127) {
                        var v11 float64 = float64(float32((((float64(1e10) - -2.25) / (float64(0.1) - (0.1 + 0.1))) / 1e10)))
                        _ = v11
                    } else {
                        fmt.Println(fi7)
                        fmt.Println(fi7)
                        fmt.Println(fi7)
                    }
                }
                _ = fi7
            }
        }
    }
    var v12 string = (("hello" + ("hello" + (("a" + ("hello" + "hello")) + "hello"))) + ("a" + ""))
    _ = v12
    return 42
}
func genFn1(p0 uint64, p1 uint16, p2 int16) uint8 {
    var v13 int = -2147483648
    v14 := []int{(v13 | -1), 9223372036854775807}
    p2 += -32768
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 255
}

func main() {
    var v1 uint64 = 2
    v2 := []int{128, -343, 9223372036854775807, 32767}
    var v3 uint64 = (v1 ^ v1)
    v3 += v1
    switch ((("" + ((("\n\t" + ("a" + "a")) + "foo bar") + "hello")) + "") + (("foo bar" + "hello") + ("\n\t" + "\n\t"))) {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        var v15 int = 9223372036854775807
        _ = v15
    default:
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint = 255
            v3 -= (v1 - (v3 & v3))
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        if (true || (1 != 282)) {
            if (388 != 0) {
                var v19 float32 = (((((0.1 * float32(0.1)) / 1e10) * 1.5) - 1e10) - 0.1)
                var v20 uint16 = 999
                switch (fi18 ^ 128) {
                case -363:
                    v21 := GenS1{F0: 42, F1: genFn1((v3 / ((v1) | 1)), (v20 * (^(v20))), 0), F2: 42}
                    _ = v21
                case -2147483648:
                    for fi22 := 0; fi22 < 4; fi22++ {
                        if true {
                            v23 := GenS2{F0: -324, F1: -9223372036854775808, F2: 9223372036854775807}
                            _ = v23
                        }
                        switch (^(fi18)) {
                        case 32768:
                            var v24 int64 = -1
                            _ = v24
                        case -347:
                            v19 = 3.14159
                        }
                        _ = fi22
                    }
                }
                _ = v19
                _ = v20
            }
            _ = len(v2)
            v25 := GenS2{F0: 1, F1: 9223372036854775807, F2: 42}
            _ = v25
        }
        _ = fi18
    }
    v2 = append(v2, 32767)
    for fi26 := 0; fi26 < 1; fi26++ {
        v27 := GenS2{F0: 2147483647, F1: 1, F2: 42}
        var v28 int = fi26
        v29 := GenS2{F0: -708, F1: -1, F2: -9223372036854775808}
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    v30 := GenS3{F0: 0, F1: true, F2: 139}
    for fi31 := 0; fi31 < 20; fi31++ {
        v1 = (v1 * (^(v3)))
        switch "" {
        case "hello":
            v2 = append(v2, (-(fi31)))
        }
        v32 := GenS4{F0: -467}
        _ = fi31
        _ = v32
    }
    v33 := GenS3{F0: 2, F1: true, F2: -718}
    var v34 int = 887
    fmt.Println(v1, v1)
    if true {
        fmt.Println(v1, v1, v34)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v3 = (v1 - v3)
    }
    v35 := GenS4{F0: -1}
    fmt.Println(v3, v34, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
}
