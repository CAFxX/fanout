package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64) uint {
    if ((true && true) && ((278 <= 2147483647) && false)) {
        var v7 uint64 = (p0 + 42)
        _ = v7
    }
    var v8 int8 = -1
    _ = p0
    _ = v8
    return 706
}
func genFn1(p0 bool, p1 int8) uint16 {
    v9 := GenS1{F0: -9223372036854775808}
    var v10 uint = 65535
    var v11 int64 = -1
    if (3.14159 != 0.0) {
        if p0 {
            if (!(v10 >= 18446744073709551615)) {
                v12 := []int{-444, 4294967295, 127, 896}
                _ = v12
            } else {
                for fi13 := 0; fi13 < 16; fi13++ {
                    v11 = v11
                    switch (fi13 | fi13) {
                    case 337:
                        v14 := GenS2{F0: 255}
                        _ = v14
                    case -9223372036854775808:
                        var v15 uint64 = 952
                        _ = v15
                    case 1:
                        var v16 float64 = ((((0.0 + 1.5) * ((1e10 + float64(0.1)) / 3.14159)) / 1e10) * 1.5)
                        _ = v16
                    default:
                        v11 += 42
                    }
                    _ = fi13
                }
                fmt.Println(p1)
            }
            v17 := []int{int(v11), -793, 32768, 127}
            _ = v17
        } else {
            var v18 float64 = 3.14159
            _ = v18
        }
        var v19 uint32 = 2
        _ = v19
    } else {
        if (!p0) {
            p1 = p1
            var v20 int16 = -409
            _ = v20
        } else {
            v21 := []int{32768, 32767}
            p1 = p1
            var v22 int32 = -57
            _ = v21
            _ = v22
        }
        v23 := GenS1{F0: -9223372036854775808}
        _ = v23
    }
    v24 := GenS2{F0: 255}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v24
    _ = v9
    return 4
}

func main() {
    v1 := []string{""}
    var v2 bool = (-9223372036854775808 <= 2147483647)
    v3 := []int{9223372036854775807, 590, 128}
    v4 := []int{-1, -2147483648, 32768}
    var v5 uint8 = 188
    var v6 uint8 = (^(v5))
    fmt.Println(v2, v6)
    for fi25 := 0; fi25 < 5; fi25++ {
        for fi26 := 0; fi26 < 10; fi26++ {
            v1 = append(v1, "foo bar")
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (v2 || (!v2)) {
                v2 = false
            } else {
                v5 = (v5 | (v5 ^ (^(v6))))
                switch "" {
                case "":
                    fmt.Println(v2, fi26)
                case "foo bar":
                    v2 = (-707 <= fi26)
                default:
                    if (!(2 == 42)) {
                        v6 -= (v6 / ((127) | 1))
                        v27 := GenS2{F0: v6}
                        var v28 int32 = 1
                        _ = v27
                        _ = v28
                    }
                }
            }
            _ = fi26
        }
        _ = fi25
    }
    v29 := GenS2{F0: v5}
    for fi30 := 0; fi30 < 10; fi30++ {
        if (((v5 / ((v6) | 1)) > (v5 / (((^(v5))) | 1))) || v2) {
            v31 := GenS3{F0: v2}
            _ = v31
        } else {
            if len(v1) > 0 {
                v1[0] = ("" + "\n\t")
            }
            fmt.Println(v5, v2)
        }
        var v32 float64 = (float64(float32(0.1)) - (3.14159 - 1.5))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi30
        _ = v32
    }
    v33 := GenS3{F0: (0 == 0)}
    fmt.Println(v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 -= (v5 | (v6 + 2))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v34 := GenS4{F0: true}
    var v35 int32 = 2147483647
    if (-389 == 64) {
        v35 -= (v35 ^ 1)
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        fmt.Println(v2, v5, v35)
        _ = fi36
    }
    if ((v6 & 42) > (v6 & (v6 & v6))) {
        v6 = 1
    }
    var v37 uint = uint(v5)
    var v38 uint8 = (^(v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
