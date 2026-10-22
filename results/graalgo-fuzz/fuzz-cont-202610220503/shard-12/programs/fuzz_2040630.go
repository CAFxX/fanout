package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint8 }
func genFn0(p0 uint64, p1 uint64, p2 uint8) uint {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float64 = 1e10
        var v9 bool = false
        p2 = (p2 % ((p2) | 1))
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int32 = 291
    if (528 >= (p1 ^ p1)) {
        v10 = v10
        v11 := []int{-1}
        var v12 string = "a"
        _ = v11
        _ = v12
    } else {
        p1 -= 363
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}
func genFn1(p0 uint64, p1 uint16, p2 string) string {
    p0 = p0
    v13 := []int{2147483647, 4294967295, 42, 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return (p2 + ((p2 + p2) + (p2 + (p2 + (p2 + (((("hello" + ("hello" + ((p2 + p2) + p2))) + "a") + (p2 + "\n\t")) + (((("hello" + "foo bar") + "foo bar") + "foo bar") + ("\n\t" + ("a" + ("" + p2))))))))))
}

func main() {
    v1 := []int{127, 4294967295, -9223372036854775808, -477}
    var v2 int = 1
    var v3 uint16 = 239
    var v4 bool = true
    v5 := []int{(v2 - v2)}
    v6 := []uint32{}
    var v14 uint8 = 101
    var v15 uint8 = (^(v14))
    var v16 bool = v4
    var v17 int16 = 0
    if ((v2 - 42) <= -223) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v4, v4)
        v18 := GenS1{F0: 18446744073709551615}
        _ = v18
    } else {
        var v19 bool = ((!(2147483647 < 2147483647)) || (399 > 57))
        _ = v19
    }
    _ = len(v6)
    if ((-74 < 1) && (v16 || ((((3.14159 - (0.0 * -2.25)) * 1e10) * 0.1) >= -2.25))) {
        v2 -= v2
        v20 := GenS1{F0: 680}
        for fi21 := 0; fi21 < 18; fi21++ {
            switch "hello" {
            case "\n\t":
                var v22 uint64 = 2
                _ = v22
            case "foo bar":
                v23 := GenS1{F0: 42}
                _ = v23
            case "":
                var v24 int32 = -2147483648
                _ = v24
            }
            _ = fi21
        }
        _ = v20
    } else {
        switch "" {
        case "foo bar":
            v14 += (v14 | (v15 - v14))
        case "hello":
            v15 = 1
        case "":
            var v25 float32 = 3.14159
            _ = v25
        default:
            v15 = v15
        }
        for fi26 := 0; fi26 < 4; fi26++ {
            fmt.Println(v17)
            v16 = (1 == 4294967295)
            _ = fi26
        }
        v15 -= (v15 - 2)
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v2 = v2
        v15 += (v14 << ((0) & 63))
        _ = fi27
    }
    var v28 int8 = 0
    v16 = v4
    v28 -= -31
    var v29 int64 = 9223372036854775807
    var v30 uint64 = 2
    if v4 {
        if ((v2 != (v2 >> ((1) & 63))) || v16) {
            v31 := GenS2{F0: "\n\t", F1: v15}
            var v32 int32 = int32(v17)
            _ = v31
            _ = v32
        } else {
            if len(v1) > 0 {
                v1[0] = v2
            }
            for fi33 := 0; fi33 < 1; fi33++ {
                fmt.Println(v28, v2)
                fmt.Println(v4)
                v4 = v16
                _ = fi33
            }
        }
        v30 += 1
        var v34 int32 = 1
        _ = v34
    } else {
        if len(v6) > 0 {
            v6[0] = 65535
        }
        var v35 float64 = 1e10
        var v36 int32 = -40
        _ = v35
        _ = v36
    }
    switch (v2 / ((32767) | 1)) {
    case -9223372036854775808:
        for fi37 := 0; fi37 < 16; fi37++ {
            _ = len(v5)
            switch v2 {
            case 2147483647:
                v38 := GenS1{F0: (^(v30))}
                _ = v38
            case 128:
                v29 = v29
            case 9223372036854775807:
                v39 := GenS2{F0: "", F1: v14}
                _ = v39
            default:
                for fi40 := 0; fi40 < 20; fi40++ {
                    v28 += v28
                    v30 += v30
                    _ = fi40
                }
            }
            for fi41 := 0; fi41 < 5; fi41++ {
                v17 -= (v17 / ((int16(v3)) | 1))
                _ = fi41
            }
            _ = fi37
        }
    }
    var v42 int16 = int16(v28)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
