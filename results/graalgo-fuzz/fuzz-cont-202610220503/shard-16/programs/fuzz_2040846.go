package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int8; F2 uint }
type GenS3 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 string; F2 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint32; F1 uint16; F2 uint16 }
func genFn0(p0 uint16) uint32 {
    v5 := GenS1{F0: -159}
    var v6 uint8 = 2
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS2{F0: p0, F1: -35, F2: 321}
        v9 := []int{}
        v10 := GenS1{F0: (fi7 * (fi7 - 128))}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    fmt.Println(p0, v6)
    _ = p0
    _ = v5
    _ = v6
    return 2147483648
}
func genFn1(p0 uint, p1 int32) int32 {
    if ((p1 >> ((p0) & 63)) <= p1) {
        v11 := GenS3{F0: (p1 % ((p1) | 1)), F1: 1, F2: 2}
        _ = v11
    }
    var v12 uint16 = 231
    _ = p0
    _ = p1
    _ = v12
    return -2147483648
}

func main() {
    var v1 uint = 11
    var v2 uint64 = 314
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    var v13 bool = true
    v14 := GenS4{F0: 42}
    if (!true) {
        v1 -= v1
    }
    var v15 bool = (true || ((v2 < (v2 << ((v1) & 63))) && v13))
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 uint64 = v2
        for fi18 := 0; fi18 < 12; fi18++ {
            if false {
                if (127 != 127) {
                    v15 = v13
                    v19 := GenS5{F0: genFn1(18446744073709551615, 1), F1: "hello", F2: 1}
                    v20 := []int{32767, (fi16 * (-(fi18)))}
                    _ = v19
                    _ = v20
                }
                v21 := GenS6{F0: 0}
                _ = v21
            } else {
                switch 1 {
                case -207:
                    var v22 uint8 = 0
                    _ = v22
                }
            }
            v17 -= 2
            _ = fi18
        }
        var v23 string = ("a" + ("hello" + "hello"))
        _ = fi16
        _ = v17
        _ = v23
    }
    v1 += v1
    var v24 int = 60
    switch (v24 - (v24 / ((int(v2)) | 1))) {
    case 683:
        v25 := []int{}
        _ = v25
    case -864:
        v26 := GenS4{F0: -2147483648}
        _ = v26
    case 32767:
        v27 := []int{2147483647}
        _ = v27
    default:
        switch 108 {
        case 62:
            var v28 float32 = 1e10
            _ = v28
        case 128:
            if false {
                for fi29 := 0; fi29 < 14; fi29++ {
                    var v30 uint8 = 255
                    _ = fi29
                    _ = v30
                }
            }
        default:
            switch (^(v24)) {
            case 2147483647:
                v2 += 629
            case 553:
                v31 := GenS1{F0: v24}
                _ = v31
            case 4294967295:
                for fi32 := 0; fi32 < 18; fi32++ {
                    var v33 int32 = -549
                    _ = fi32
                    _ = v33
                }
            default:
                var v34 uint = v1
                _ = v34
            }
        }
    }
    v24 = (v24 / ((v24) | 1))
    v35 := GenS7{F0: (^(v1))}
    var v36 uint8 = 128
    v37 := GenS8{F0: -945}
    v38 := GenS9{F0: genFn0(42), F1: 1, F2: 0}
    v39 := []int{4294967295, -9223372036854775808, (v24 / ((32768) | 1)), 4294967295}
    var v40 uint64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
