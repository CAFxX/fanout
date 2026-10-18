package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint64; F2 int64 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 bool; F1 int8; F2 int16 }
type GenS5 struct { F0 int64; F1 uint }
func genFn0(p0 string) uint16 {
    v7 := []int{198, 872, 181, 230}
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(fi8, p0)
        _ = fi8
    }
    fmt.Println(p0)
    var v9 uint64 = 42
    _ = p0
    _ = v7
    _ = v9
    return 236
}

func main() {
    var v1 int64 = 397
    var v2 bool = (true && true)
    v3 := make(map[uint32]uint32)
    var v4 uint = 1
    var v5 float32 = (((((3.14159 + float32((-2.25 / (3.14159 - 1e10)))) + 1.5) * 0.1) / -2.25) + float32((0.1 * float64(0.0))))
    v6 := make(map[int]int)
    switch 32767 {
    case 0:
        var v10 uint32 = 410
        _ = v10
    default:
        var v11 uint16 = 34
        _ = v11
    }
    var v12 int8 = 100
    for fi13 := 0; fi13 < 2; fi13++ {
        if v2 {
            v1 += (v1 + v1)
            switch "" {
            case "\n\t":
                v6[-545] = (fi13 >> ((335) & 63))
            default:
                var v14 int = fi13
                _ = v14
            }
            v15 := GenS1{F0: (^(fi13))}
            _ = v15
        } else {
            switch (fi13 << (((v4 - (v4 << (((v4 & (v4 / (((v4 >> ((v4) & 63))) | 1)))) & 63)))) & 63)) {
            case 649:
                if (!v2) {
                    v16 := GenS2{F0: (fi13 << ((v4) & 63)), F1: 1, F2: 1}
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 3; fi17++ {
                        var v18 bool = (v2 || ("hello" == "hello"))
                        v4 = v4
                        v5 += v5
                        var v19 int16 = -1
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
            case 1:
                switch "a" {
                case "a":
                    _ = len(v6)
                case "\n\t":
                    var v20 string = ("foo bar" + "foo bar")
                    _ = v20
                case "":
                    var v21 uint64 = 42
                    _ = v21
                }
            }
        }
        _ = fi13
    }
    v22 := v3[280]
    var v23 bool = (("foo bar" + "\n\t") < ("a" + ""))
    v24 := []int{47, 42}
    fmt.Println(v2, v22)
    var v25 float64 = (float64(v5) / 1.0)
    var v26 uint16 = 2
    v12 -= (v12 % ((v12) | 1))
    var v27 int32 = -2147483648
    if (int8(v26) != (v12 / (((v12 | (v12 & (v12 & int8(v4))))) | 1))) {
        v23 = true
    } else {
        v25 = (3.14159 + (v25 / v25))
    }
    v28 := GenS2{F0: -1, F1: 42, F2: -9223372036854775808}
    switch int(v4) {
    case -1:
        if ((v2 && (!(32767 < -139))) || v23) {
            var v29 uint64 = 183
            for fi30 := 0; fi30 < 1; fi30++ {
                fmt.Println(v27)
                var v31 uint32 = uint32(v4)
                _ = fi30
                _ = v31
            }
            if true {
                var v32 uint64 = (v29 & (^(v29)))
                fmt.Println(v25, v1, v25)
                _ = v32
            }
            _ = v29
        } else {
            for fi33 := 0; fi33 < 1; fi33++ {
                v27 -= (v27 + v27)
                v34 := GenS3{F0: v12, F1: 18446744073709551615}
                if ((("hello" + (("foo bar" + ((("a" + "hello") + "a") + ("a" + (("hello" + ("foo bar" + "hello")) + "foo bar")))) + ("" + ""))) + "\n\t") >= (("a" + ("foo bar" + "")) + "foo bar")) {
                    v25 -= (0.1 - (float64(float32(0.1)) - 1e10))
                }
                _ = fi33
                _ = v34
            }
            for fi35 := 0; fi35 < 11; fi35++ {
                v36 := GenS4{F0: true, F1: v12, F2: 1000}
                v4 += v4
                _ = fi35
                _ = v36
            }
            for fi37 := 0; fi37 < 10; fi37++ {
                if v23 {
                    _ = len(v24)
                } else {
                    v3[v22] = 559
                    v38 := GenS3{F0: v12, F1: (v4 << ((v4) & 63))}
                    _ = v38
                }
                var v39 int64 = (v1 ^ v1)
                _ = fi37
                _ = v39
            }
        }
    case 490:
        v40 := GenS5{F0: -9223372036854775808, F1: v4}
        _ = v40
    case -265:
        var v41 uint32 = v22
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
