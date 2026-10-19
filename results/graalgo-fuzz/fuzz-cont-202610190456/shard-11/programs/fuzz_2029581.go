package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 string }
func genFn0() bool {
    var v7 int16 = 1000
    v7 = v7
    v8 := []int{-2147483648}
    v9 := GenS1{F0: 0, F1: -9223372036854775808, F2: -882}
    _ = v7
    _ = v8
    _ = v9
    return true
}
func genFn1(p0 string) int32 {
    switch p0 {
    case "foo bar":
        var v10 float64 = (0.1 / 0.1)
        _ = v10
    default:
        if (false && (380 == 0)) {
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int64 = 42
                v12 = (v12 ^ 9223372036854775807)
                var v13 bool = true
                _ = fi11
                _ = v12
                _ = v13
            }
        } else {
            if (true || false) {
                switch 32767 {
                case -2147483648:
                    var v14 uint8 = 0
                    _ = v14
                case 128:
                    v15 := GenS2{F0: 274, F1: 74}
                    _ = v15
                case 9223372036854775807:
                    p0 = p0
                default:
                    v16 := GenS2{F0: 2, F1: 1}
                    _ = v16
                }
                v17 := []int{-9223372036854775808, 128, -2147483648}
                _ = v17
            } else {
                var v18 int64 = 918
                _ = v18
            }
            var v19 uint32 = 65535
            _ = v19
        }
    }
    v20 := []int{664, 42, -1}
    var v21 uint32 = 1
    if ((1 == 127) && false) {
        v22 := GenS3{F0: p0}
        _ = v22
    } else {
        var v23 int16 = -1
        v23 = v23
        if false {
            var v24 int8 = -26
            v25 := GenS3{F0: "a"}
            _ = v24
            _ = v25
        } else {
            p0 = p0
            if false {
                fmt.Println(v23, v21, v23)
                var v26 int32 = -704
                _ = v26
            } else {
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 float32 = 1e10
                    _ = fi27
                    _ = v28
                }
            }
        }
        _ = v23
    }
    _ = p0
    _ = v20
    _ = v21
    return 589
}

func main() {
    var v1 int = 32768
    var v2 int = (v1 * -9223372036854775808)
    var v3 string = "\n\t"
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]string)
    var v6 uint8 = 159
    if false {
        v2 += 32768
        var v29 int16 = 1
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        v31 := GenS1{F0: 990, F1: 1, F2: v2}
        _ = fi30
        _ = v31
    }
    switch v3 {
    case "":
        v3 = ("hello" + (v3 + (v3 + (v3 + v3))))
    default:
        if (-2147483648 != -1) {
            var v32 int64 = -807
            if true {
                for fi33 := 0; fi33 < 2; fi33++ {
                    fmt.Println(v2)
                    var v34 uint8 = v6
                    _ = fi33
                    _ = v34
                }
                v35 := GenS1{F0: 42, F1: int64(v2), F2: v2}
                _ = v35
            }
            _ = v32
        }
    }
    switch (v1 >> ((18446744073709551615) & 63)) {
    case 127:
        var v36 int = (v1 ^ 206)
        _ = v36
    default:
        v2 = v2
    }
    if false {
        var v37 float64 = 1.5
        _ = v37
    }
    v38 := []int{v1, (v1 % ((4294967295) | 1)), 128}
    v1 += (v1 << ((4294967295) & 63))
    if (true && true) {
        fmt.Println(v2, v2, v6)
        var v39 float32 = ((-2.25 + 1.5) - 0.0)
        v3 = v3
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
