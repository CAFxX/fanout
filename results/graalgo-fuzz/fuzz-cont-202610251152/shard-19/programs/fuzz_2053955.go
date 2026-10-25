package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int16 }
func genFn0() int64 {
    var v4 uint8 = 254
    var v5 int32 = 2147483647
    _ = v4
    _ = v5
    return -747
}
func genFn1(p0 string) uint16 {
    p0 = ("" + (("" + p0) + ("a" + (p0 + (p0 + ("foo bar" + p0))))))
    if (!true) {
        v6 := GenS1{F0: 2147483647, F1: 42, F2: 42}
        _ = v6
    }
    p0 = ("hello" + p0)
    v7 := GenS2{F0: 593, F1: 18446744073709551615, F2: 42}
    v8 := []int{0, -153}
    _ = p0
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 42
    v9 := GenS3{F0: 42}
    v10 := GenS3{F0: -494}
    v3 = 869
    var v11 int = -157
    switch v11 {
    case 32768:
        if (!(float64((1e10 + 1.5)) == 3.14159)) {
            var v12 int64 = 42
            v13 := []int{(^(v11)), v11, (v11 | (v11 * v11))}
            _ = v12
            _ = v13
        }
    default:
        var v14 uint8 = 42
        _ = v14
    }
    v15 := []int{}
    if (v1 && true) {
        var v16 bool = v1
        _ = v16
    }
    fmt.Println(v11, v1)
    for fi17 := 0; fi17 < 3; fi17++ {
        v1 = ((v3 ^ (v3 / ((genFn1("")) | 1))) > (v3 ^ 190))
        _ = fi17
    }
    if false {
        delete(v2, 256)
    } else {
        v3 = v3
        _ = len(v2)
        v1 = (false && (!(1 >= 2147483647)))
    }
    switch v11 {
    case 0:
        for fi18 := 0; fi18 < 5; fi18++ {
            if len(v15) > 0 {
                v15[0] = v11
            }
            delete(v2, 705)
            _ = fi18
        }
    default:
        switch v11 {
        case -1:
            var v19 float64 = 3.14159
            _ = v19
        case -109:
            v2[0] = 1
        }
    }
    _ = len(v2)
    switch "" {
    case "foo bar":
        v11 = v11
    case "":
        if ((v1 || (42 != -64)) && v1) {
            switch v11 {
            case 4294967295:
                v11 += 96
            case 32767:
                var v20 int16 = 1000
                _ = v20
            case 1:
                switch (v11 | 32767) {
                case -2147483648:
                    fmt.Println(v11, v3, v1)
                default:
                    if (v1 || (42 > -2147483648)) {
                        var v21 float32 = ((float32(0.1) + 1.5) + (-2.25 / 1.0))
                        _ = v21
                    } else {
                        if len(v15) > 0 {
                            v15[0] = 32767
                        }
                    }
                }
            }
            switch 32768 {
            case 747:
                v3 = v3
            case 32768:
                v1 = ((0 > 0) || true)
            }
        } else {
            fmt.Println(v11, v11, v1)
        }
    default:
        if v1 {
            var v22 int = v11
            v22 += 0
            v11 += v11
            _ = v22
        } else {
            for fi23 := 0; fi23 < 1; fi23++ {
                v11 = (fi23 >> ((42) & 63))
                _ = fi23
            }
            switch 855 {
            case 672:
                delete(v2, 10)
            case -9223372036854775808:
                v24 := GenS2{F0: 617, F1: 673, F2: 836}
                _ = v24
            }
        }
    }
    v25 := GenS2{F0: 596, F1: 355, F2: 1}
    var v26 int32 = 1
    v27 := GenS2{F0: 321, F1: 644, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v9
}
