package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 string, p1 int32, p2 bool) uint {
    v5 := GenS1{F0: 0, F1: -34, F2: 1}
    if (-9223372036854775808 == 2147483647) {
        p2 = (true || (42 > -19))
        fmt.Println(p2, p1, p2)
    }
    if true {
        v6 := []int{42, 4294967295}
        _ = v6
    }
    v7 := GenS2{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 255
}

func main() {
    var v1 int8 = 42
    v2 := []uint32{0, 1, 213}
    v3 := []string{"a", ("\n\t" + "foo bar"), ("foo bar" + ("" + "foo bar")), "foo bar"}
    var v4 int64 = 1
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 float32 = 0.0
        _ = fi8
        _ = v9
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v10 int = 42
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 bool = false
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := GenS3{F0: 192, F1: 465, F2: 2}
            fmt.Println(v4, v10, fi13)
            fmt.Println(v12)
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    v10 = (-(v10))
    if (false && (331 <= 42)) {
        switch (("hello" + "foo bar") + "a") {
        case "a":
            fmt.Println(v4, v1, v1)
        default:
            v15 := GenS4{F0: 1}
            _ = v15
        }
    }
    v10 = (v10 << ((676) & 63))
    fmt.Println(v1, v1, v10)
    fmt.Println(v4, v4, v10)
    if (1e10 != ((float64((1.5 + 3.14159)) * float64(0.0)) * 3.14159)) {
        var v16 int8 = -19
        var v17 uint16 = 42
        _ = v16
        _ = v17
    }
    if ((1 < 4294967295) || false) {
        var v18 uint64 = 358
        if len(v2) > 0 {
            v2[0] = 1
        }
        _ = v18
    } else {
        v2 = append(v2, 65535)
        v10 -= (v10 | v10)
        for fi19 := 0; fi19 < 9; fi19++ {
            if ((fi19 % (((v10 & (v10 << ((418) & 63)))) | 1)) <= fi19) {
                var v20 int8 = (v1 % (((v1 % ((v1) | 1))) | 1))
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                v21 := GenS3{F0: 991, F1: 1, F2: 42}
                _ = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 7; fi22++ {
                    var v23 uint32 = 0
                    if (!(v1 >= 1)) {
                        _ = len(v2)
                        if true {
                            var v24 uint64 = 42
                            switch 42 {
                            case 128:
                                v24 -= v24
                            case 32768:
                                v4 -= v4
                            }
                            var v25 int = 0
                            _ = v24
                            _ = v25
                        }
                    }
                    _ = fi22
                    _ = v23
                }
            }
            _ = fi19
        }
    }
    switch (v10 & (v10 >> ((42) & 63))) {
    case 0:
        fmt.Println(v1)
    case 128:
        fmt.Println(v4, v10)
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v10, v1, v4)
    case "hello":
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 int32 = 2147483647
            var v28 bool = true
            switch 2147483647 {
            case 589:
                switch (^(v10)) {
                case 0:
                    _ = len(v2)
                case 4294967295:
                    v29 := GenS3{F0: 2147483647, F1: 42, F2: 4294967295}
                    _ = v29
                default:
                    if v28 {
                        v28 = v28
                    } else {
                        v2 = append(v2, 2147483647)
                        v30 := GenS5{F0: 65535}
                        _ = v30
                    }
                }
            case 9223372036854775807:
                switch fi26 {
                case -9223372036854775808:
                    v10 -= -449
                case 42:
                    v1 -= (v1 << ((65535) & 63))
                }
            default:
                v27 += (v27 << ((65535) & 63))
            }
            _ = fi26
            _ = v27
            _ = v28
        }
    default:
        fmt.Println(v4, v4, v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
