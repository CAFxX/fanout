package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint, p1 uint64) uint {
    var v4 uint = p0
    var v5 bool = (-32768 > -1)
    var v6 int16 = 32767
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return (p0 & p0)
}
func genFn1() uint64 {
    v7 := GenS1{F0: ""}
    var v8 float32 = 3.14159
    var v9 int = 349
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint64 = 662
    v3 := make(map[string]uint32)
    var v10 uint64 = genFn1()
    var v11 uint32 = v1
    var v12 uint64 = (v10 / (((v2 ^ 2)) | 1))
    var v13 uint8 = 72
    if true {
        if (32767 == 582) {
            fmt.Println(v10, v10, v2)
            v3[("a" + "foo bar")] = 1
            v2 -= (v12 * (v10 + (v2 + (v10 / ((v2) | 1)))))
        } else {
            var v14 int64 = 1
            _ = v14
        }
    } else {
        var v15 int32 = 1
        _ = v15
    }
    if ((0 != 255) || true) {
        if true {
            var v16 uint32 = uint32(v12)
            _ = v16
        }
        v1 -= 436
    } else {
        v17 := []int{127, -470, 4294967295, 32767}
        v18 := GenS2{F0: v1}
        var v19 float64 = 3.14159
        _ = v17
        _ = v18
        _ = v19
    }
    if (32767 >= 1000) {
        if false {
            v20 := GenS2{F0: (v1 | (v11 >> ((genFn0(1, (v10 * genFn1()))) & 63)))}
            _ = v20
        }
        switch 736 {
        case 32768:
            v21 := v3[("a" + ("foo bar" + ("\n\t" + "")))]
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 16; fi22++ {
            fmt.Println(v11, v13)
            var v23 int32 = 2147483647
            _ = fi22
            _ = v23
        }
    }
    var v24 string = "hello"
    v10 -= (v12 | v2)
    v10 = genFn1()
    switch v24 {
    case "a":
        if true {
            if (false || (!(false || ((v10 | (v12 ^ 2)) <= (v2 & 2))))) {
                v25 := GenS3{F0: v1, F1: v1}
                _ = v25
            }
        }
    case "\n\t":
        if (!true) {
            v26 := GenS4{F0: (v13 % ((42) | 1)), F1: 42, F2: v11}
            switch 4294967295 {
            case 9223372036854775807:
                v27 := GenS5{F0: 0}
                _ = v27
            case 684:
                v1 = v11
            default:
                var v28 int32 = 2147483647
                _ = v28
            }
            var v29 int32 = -624
            _ = v26
            _ = v29
        } else {
            fmt.Println(v12, v24)
        }
    case "hello":
        v3[""] = v1
    default:
        var v30 uint = 42
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
}
