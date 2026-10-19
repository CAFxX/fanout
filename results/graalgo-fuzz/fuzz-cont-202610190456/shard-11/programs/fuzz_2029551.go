package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint32 }
func genFn0(p0 int64, p1 uint8, p2 bool) int32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        if (0 < 989) {
            p0 = (^(p0))
        }
        fmt.Println(p2, p1, fi4)
        p0 = (p0 % (((p0 & p0)) | 1))
        _ = fi4
    }
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[int]uint32)
    var v3 uint16 = 2
    v3 = (v3 >> ((255) & 63))
    if (0 >= 0) {
        v5 := []int{42, -355, 127}
        _ = v5
    }
    v6 := GenS1{F0: "\n\t", F1: -1, F2: 1}
    if true {
        switch 32768 {
        case 32768:
            v1 += (v1 << ((65535) & 63))
        case -1:
            for fi7 := 0; fi7 < 14; fi7++ {
                if false {
                    var v8 uint32 = 0
                    var v9 string = "foo bar"
                    _ = v8
                    _ = v9
                }
                v10 := GenS1{F0: "a", F1: 9223372036854775807, F2: 749}
                v3 += v3
                _ = fi7
                _ = v10
            }
        case 9223372036854775807:
            for fi11 := 0; fi11 < 15; fi11++ {
                delete(v2, (fi11 ^ fi11))
                var v12 int32 = v1
                fmt.Println(v12, v12, fi11)
                _ = fi11
                _ = v12
            }
        default:
            v13 := []int{}
            _ = v13
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 float32 = 0.1
            var v16 int16 = int16(fi14)
            var v17 float32 = (v15 - (float32(1e10) + (v15 + v15)))
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        var v18 int8 = -45
        _ = v18
    } else {
        switch "" {
        case "hello":
            _ = len(v2)
        case "\n\t":
            switch -2147483648 {
            case 237:
                for fi19 := 0; fi19 < 12; fi19++ {
                    fmt.Println(v1, fi19, v3)
                    _ = fi19
                }
            case -280:
                switch -9223372036854775808 {
                case 1:
                    var v20 int16 = 1
                    _ = v20
                case -2147483648:
                    var v21 bool = true
                    _ = v21
                case 32767:
                    var v22 int32 = v1
                    _ = v22
                default:
                    if true {
                        fmt.Println(v1)
                        v3 -= v3
                    } else {
                        fmt.Println(v3, v1)
                    }
                }
            case 9223372036854775807:
                if (1 == -648) {
                    var v23 int8 = -124
                    fmt.Println(v3, v23)
                    _ = v23
                }
            default:
                var v24 uint8 = 1
                _ = v24
            }
        case "":
            v1 = (^(v1))
        default:
            var v25 bool = ((v1 | v1) <= v1)
            _ = v25
        }
        var v26 float32 = 3.14159
        _ = v26
    }
    if true {
        var v27 bool = false
        if (0 >= uint8(v3)) {
            for fi28 := 0; fi28 < 3; fi28++ {
                v29 := GenS2{F0: (fi28 << ((255) & 63)), F1: 332}
                v30 := GenS3{F0: 0, F1: -32768, F2: 4294967295}
                _ = fi28
                _ = v29
                _ = v30
            }
        } else {
            var v31 uint16 = v3
            _ = v31
        }
        var v32 int8 = -43
        _ = v27
        _ = v32
    }
    var v33 uint64 = 10
    var v34 uint = 151
    v34 = v34
    var v35 uint64 = v33
    if true {
        v36 := GenS3{F0: uint8(v33), F1: 736, F2: 42}
        v37 := GenS3{F0: 1, F1: -1, F2: 2147483648}
        _ = v36
        _ = v37
    }
    var v38 string = ("foo bar" + (("hello" + "foo bar") + "hello"))
    var v39 float32 = 3.14159
    for fi40 := 0; fi40 < 6; fi40++ {
        v41 := []int{318, fi40, (fi40 | -212)}
        _ = fi40
        _ = v41
    }
    var v42 float32 = v39
    v34 = v34
    var v43 int8 = 54
    delete(v2, -572)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v42
    _ = v43
    _ = v6
}
