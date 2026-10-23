package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int32; F1 int32; F2 string }
func genFn0(p0 bool) string {
    p0 = true
    var v5 uint8 = 152
    var v6 int16 = 1
    if ((0.0 + (3.14159 + 0.1)) == ((0.0 * 0.1) - -2.25)) {
        switch "\n\t" {
        case "hello":
            for fi7 := 0; fi7 < 8; fi7++ {
                v6 -= 42
                _ = fi7
            }
        case "":
            v8 := GenS1{F0: -1}
            _ = v8
        case "foo bar":
            if p0 {
                if (float32(((((1.5 - (((((1.5 - 0.0) - 1.5) / (((float64((1e10 + -2.25)) + 3.14159) / 3.14159) - 0.1)) / 0.1) * 1e10)) - 0.1) - 3.14159) * (0.0 + (1.5 * 1.5)))) > -2.25) {
                    v6 = (v6 / (((v6 << ((2) & 63))) | 1))
                }
                p0 = (p0 || true)
            }
        }
        v5 -= 127
    }
    v5 += v5
    _ = p0
    _ = v5
    _ = v6
    return "foo bar"
}
func genFn1(p0 bool, p1 uint32, p2 int16) int16 {
    var v9 int8 = -41
    if (-2.25 >= 1e10) {
        p2 = (p2 << ((18446744073709551615) & 63))
        v10 := GenS2{F0: -202, F1: 834}
        v11 := []int{-50}
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := []int{fi12, (-(fi12))}
            _ = fi12
            _ = v13
        }
        v14 := GenS2{F0: 9223372036854775807, F1: 770}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p2
}

func main() {
    var v1 int = 127
    v2 := make(map[string]bool)
    var v3 int = (v1 >> ((18446744073709551615) & 63))
    v4 := make(map[uint32]string)
    v15 := GenS2{F0: (v1 ^ v1), F1: 0}
    v16 := []int{(-(v1)), (v1 * 32767), (v1 % ((v3) | 1))}
    if false {
        var v17 float64 = 0.1
        v3 -= 1
        _ = v17
    }
    v18 := GenS2{F0: (v1 << ((242) & 63)), F1: 255}
    v1 = v1
    v3 = v1
    if false {
        for fi19 := 0; fi19 < 12; fi19++ {
            v3 += (fi19 * (fi19 - (fi19 - (v3 * (fi19 & v1)))))
            _ = fi19
        }
        v3 -= (v1 ^ 4294967295)
        var v20 uint32 = 0
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 bool = (true && false)
        _ = fi21
        _ = v22
    }
    v23 := GenS3{F0: -2147483648, F1: 0, F2: ("\n\t" + "a")}
    var v24 uint8 = 2
    var v25 uint = 4294967295
    v24 = 128
    if (false || (103 >= 1)) {
        v25 += 18446744073709551615
    } else {
        if (-926 <= 101) {
            var v26 int32 = 2147483647
            for fi27 := 0; fi27 < 5; fi27++ {
                for fi28 := 0; fi28 < 8; fi28++ {
                    v2["\n\t"] = ("a" < ("a" + ("" + "\n\t")))
                    _ = fi28
                }
                _ = fi27
            }
            v1 = v3
            _ = v26
        }
    }
    v16 = append(v16, (v1 % ((v1) | 1)))
    var v29 float64 = 1.5
    switch (v1 >> (((v25 & v25)) & 63)) {
    case 32767:
        var v30 uint64 = 0
        _ = v30
    case 167:
        _ = len(v4)
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 uint64 = 534
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
