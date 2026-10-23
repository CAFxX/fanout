package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []int{4294967295, 0, 1, -1000}
    var v2 float64 = (-2.25 - float64((((float32(0.1) * (3.14159 + (1.5 + (-2.25 * 0.0)))) / 0.1) - -2.25)))
    var v3 int64 = -269
    v4 := []int{42, 0, -409, 32768}
    var v5 uint16 = 65535
    v6 := make(map[int]bool)
    v1 = append(v1, 9223372036854775807)
    v4 = append(v4, 32767)
    var v7 uint8 = 2
    v7 = v7
    var v8 int8 = -128
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := GenS1{F0: 1}
        _ = fi9
        _ = v10
    }
    v2 += (v2 + float64(3.14159))
    if (65535 == 65535) {
        var v11 uint32 = 0
        _ = v11
    } else {
        for fi12 := 0; fi12 < 3; fi12++ {
            for fi13 := 0; fi13 < 10; fi13++ {
                fmt.Println(v2)
                v3 -= (v3 * (v3 - 42))
                if false {
                    delete(v6, (fi13 % (((fi13 * fi12)) | 1)))
                    var v14 uint32 = 0
                    _ = v14
                }
                v15 := GenS2{F0: 2}
                _ = fi13
                _ = v15
            }
            _ = fi12
        }
    }
    var v16 float64 = v2
    if ((false || true) || (v7 <= v7)) {
        fmt.Println(v7, v7)
    }
    switch 2147483647 {
    case 2147483647:
        var v17 int = 352
        _ = v17
    case 87:
        v16 = v16
    default:
        if false {
            switch "\n\t" {
            case "hello":
                if (!(1.5 >= float64(float32((float64(1.5) * float64((3.14159 - (3.14159 - 3.14159)))))))) {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        var v19 int8 = v8
                        _ = fi18
                        _ = v19
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v20 uint8 = 248
                    _ = v20
                }
            case "a":
                v7 = v7
            case "\n\t":
                v5 = v5
            default:
                var v21 int16 = 42
                _ = v21
            }
            _ = len(v6)
            v2 = v16
        } else {
            v22 := v6[-9223372036854775808]
            var v23 int16 = 1
            v24 := GenS1{F0: v8}
            _ = v22
            _ = v23
            _ = v24
        }
    }
    v8 = int8(v7)
    v5 += v5
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 uint32 = 4294967295
        v5 = v5
        v26 = uint32(v5)
        _ = fi25
        _ = v26
    }
    if (-1 == 794) {
        switch "\n\t" {
        case "hello":
            v7 = (v7 / (((v7 + 2)) | 1))
        case "":
            fmt.Println(v5)
        case "foo bar":
            fmt.Println(v3)
        }
    } else {
        v7 = 42
        for fi27 := 0; fi27 < 11; fi27++ {
            v5 = (v5 * (v5 << ((2) & 63)))
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 float64 = float64(3.14159)
                switch "a" {
                case "":
                    v16 = 3.14159
                default:
                    v3 = v3
                }
                if false {
                    if false {
                        v30 := v6[(fi27 / (((fi27 + fi28)) | 1))]
                        v30 = ((((^(v8)) <= (v8 >> ((65535) & 63))) || v30) && (!v30))
                        if (27 <= v7) {
                            var v31 float32 = (3.14159 - (0.0 + ((-2.25 - 3.14159) * (-2.25 / 1.5))))
                            fmt.Println(v16, fi27, v16)
                            _ = v31
                        }
                        _ = v30
                    }
                    var v32 int16 = -232
                    _ = v32
                } else {
                    var v33 int = (fi27 ^ (^(fi28)))
                    _ = v33
                }
                _ = fi28
                _ = v29
            }
            v16 = float64((((1.5 + (-2.25 + float32(v2))) / -2.25) / (((0.0 - -2.25) + float32(v2)) - (1e10 * 1e10))))
            _ = fi27
        }
    }
    var v34 float64 = 1.5
    var v35 int8 = v8
    v7 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
