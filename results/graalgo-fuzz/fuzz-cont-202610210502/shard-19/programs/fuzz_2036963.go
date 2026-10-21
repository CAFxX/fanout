package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 string; F2 uint64 }

func main() {
    var v1 uint64 = 1
    var v2 int64 = -9223372036854775808
    var v3 uint = 174
    var v4 int64 = (v2 - v2)
    var v5 float32 = float32((3.14159 + (0.1 / 0.1)))
    for fi6 := 0; fi6 < 15; fi6++ {
        if true {
            var v7 uint = (v3 % ((65535) | 1))
            v8 := []int{fi6}
            _ = v7
            _ = v8
        } else {
            v5 = float32(0.0)
            switch (fi6 >> ((4294967295) & 63)) {
            case 1:
                var v9 int16 = 1
                _ = v9
            }
            for fi10 := 0; fi10 < 6; fi10++ {
                if (false || false) {
                    var v11 uint16 = 42
                    v12 := []int{fi10, fi10, (^(fi6))}
                    switch (fi10 / (((fi6 - (^(fi10)))) | 1)) {
                    case 9223372036854775807:
                        v13 := GenS1{F0: 127, F1: "\n\t", F2: true}
                        _ = v13
                    }
                    fmt.Println(fi10, fi10)
                    _ = v11
                    _ = v12
                }
                _ = fi10
            }
        }
        var v14 bool = false
        _ = fi6
        _ = v14
    }
    v2 -= v2
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int32 = 2147483647
        if true {
            for fi17 := 0; fi17 < 16; fi17++ {
                switch (^(fi15)) {
                case 9223372036854775807:
                    var v18 string = (("a" + ("foo bar" + (("" + "\n\t") + "a"))) + ("hello" + ("hello" + ((("" + (("a" + "\n\t") + "")) + "\n\t") + "\n\t"))))
                    var v19 uint64 = v1
                    v19 = v1
                    v20 := GenS2{F0: (^(v3))}
                    var v21 float32 = float32(0.0)
                    _ = v18
                    _ = v19
                    _ = v20
                    _ = v21
                case 212:
                    fmt.Println(v16)
                case -998:
                    var v22 uint8 = 42
                    fmt.Println(v1)
                    if ((fi17 + -9223372036854775808) == -9223372036854775808) {
                        switch (fi15 % ((fi17) | 1)) {
                        case 483:
                            v22 = (v22 - v22)
                        }
                        var v23 int32 = 2147483647
                        v5 -= (((v5 - v5) - 3.14159) * (3.14159 / v5))
                        _ = v23
                    }
                    _ = v22
                default:
                    var v24 int16 = 32767
                    _ = v24
                }
                if (true || false) {
                    var v25 int32 = (^(v16))
                    v3 -= 18446744073709551615
                    if true {
                        var v26 int32 = v16
                        _ = v26
                    }
                    _ = v25
                } else {
                    var v27 int32 = (v16 | -722)
                    var v28 uint16 = 2
                    _ = v27
                    _ = v28
                }
                _ = fi17
            }
        }
        _ = fi15
        _ = v16
    }
    v3 += 42
    v29 := []int{1}
    if (-2147483648 != -2147483648) {
        v5 = (v5 - 0.1)
    }
    fmt.Println(v4)
    var v30 int64 = (-(v2))
    if true {
        var v31 int64 = v30
        var v32 uint16 = 65535
        _ = v31
        _ = v32
    } else {
        if len(v29) > 0 {
            v29[0] = 214
        }
        var v33 bool = true
        v34 := GenS3{F0: 548, F1: "\n\t", F2: v1}
        _ = v33
        _ = v34
    }
    if true {
        var v35 float64 = (1.5 * 1e10)
        _ = v35
    } else {
        v2 += (v2 >> ((838) & 63))
        fmt.Println(v4, v4)
        var v36 uint16 = 602
        _ = v36
    }
    fmt.Println(v4, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
