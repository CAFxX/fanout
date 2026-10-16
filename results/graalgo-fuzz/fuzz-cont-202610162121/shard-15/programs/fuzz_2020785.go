package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint64) string {
    v5 := GenS1{F0: 509, F1: ("a" + "hello")}
    if (true || false) {
        v6 := GenS1{F0: 32767, F1: "\n\t"}
        p0 += (p0 / ((p0) | 1))
        p0 = (p0 - p0)
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int64 = 1
        var v9 uint64 = (^(p0))
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 = (p0 << ((1) & 63))
    _ = p0
    _ = v5
    return ("foo bar" + "")
}
func genFn1() uint32 {
    v10 := GenS1{F0: 211, F1: ("foo bar" + "\n\t")}
    v11 := GenS1{F0: 2, F1: "a"}
    if (273 < 1) {
        var v12 int8 = -1
        _ = v12
    }
    if ((1.5 >= 1e10) && true) {
        var v13 string = "\n\t"
        var v14 uint64 = 1
        v14 = (v14 | (v14 ^ v14))
        _ = v13
        _ = v14
    }
    _ = v10
    _ = v11
    return 4294967295
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 3.14159
    var v3 string = "hello"
    var v4 uint64 = 458
    switch -80 {
    case 42:
        var v15 uint = 42
        _ = v15
    case -1:
        if (!(v4 < v4)) {
            var v16 int64 = 0
            delete(v1, 32767)
            _ = v16
        } else {
            var v17 float32 = (0.1 - 0.1)
            if (true || (634 <= 880)) {
                v18 := GenS2{F0: 0, F1: 494}
                v2 -= v2
                v17 = 3.14159
                _ = v18
            } else {
                for fi19 := 0; fi19 < 11; fi19++ {
                    var v20 uint8 = 127
                    fmt.Println(v20, v17)
                    var v21 uint32 = 0
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            v22 := []int{}
            _ = v17
            _ = v22
        }
    case 128:
        fmt.Println(v3)
    }
    var v23 uint64 = (v4 - v4)
    var v24 float32 = (float32(float64((0.1 / 1e10))) - 1e10)
    if (-2.25 != (v24 / v24)) {
        v25 := GenS1{F0: 32767, F1: (("" + v3) + ("hello" + (v3 + "\n\t")))}
        _ = v25
    }
    var v26 uint32 = 0
    switch 332 {
    case -1:
        var v27 uint8 = 236
        _ = v27
    }
    v28 := []int{9223372036854775807, -263, 127}
    for fi29 := 0; fi29 < 2; fi29++ {
        delete(v1, (fi29 % ((fi29) | 1)))
        var v30 int64 = 986
        var v31 uint32 = (v26 ^ v26)
        _ = fi29
        _ = v30
        _ = v31
    }
    if false {
        switch v3 {
        case "":
            switch 32767 {
            case -798:
                v32 := v1[0]
                _ = v32
            case 0:
                v33 := GenS1{F0: 890, F1: v3}
                _ = v33
            case 2147483647:
                v1[127] = (v26 * 20)
            default:
                v34 := GenS3{F0: 156}
                _ = v34
            }
        case "foo bar":
            for fi35 := 0; fi35 < 8; fi35++ {
                var v36 uint16 = 32767
                if (!(86 > 24)) {
                    v3 = ("\n\t" + v3)
                } else {
                    v4 -= v23
                    delete(v1, (fi35 | (^(fi35))))
                    v26 += genFn1()
                }
                var v37 uint16 = (v36 / (((v36 - (v36 % ((1) | 1)))) | 1))
                _ = fi35
                _ = v36
                _ = v37
            }
        }
    }
    if ((!true) && true) {
        switch 32768 {
        case 32768:
            var v38 bool = ((1 <= 127) || (v2 != v2))
            _ = v38
        case 42:
            var v39 int = 652
            _ = v39
        default:
            v24 = v24
        }
    } else {
        v23 += v23
    }
    var v40 int = 32767
    v41 := v1[(-(v40))]
    v4 = (v4 - (^(v4)))
    v28 = append(v28, -2147483648)
    switch "hello" {
    case "":
        var v42 uint = 18446744073709551615
        _ = v42
    }
    var v43 int32 = -2147483648
    var v44 uint8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v44
}
