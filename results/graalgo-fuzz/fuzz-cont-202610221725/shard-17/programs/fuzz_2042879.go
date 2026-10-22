package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int16; F1 int64 }
func genFn0(p0 int32) uint32 {
    if (!(0 >= 810)) {
        if (!(0 <= -32768)) {
            var v5 int = 228
            _ = v5
        }
    } else {
        var v6 float32 = 0.0
        var v7 float64 = 1.5
        fmt.Println(v6)
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint64 = 0
        _ = fi8
        _ = v9
    }
    p0 = (p0 % ((p0) | 1))
    var v10 uint16 = 32767
    _ = p0
    _ = v10
    return 0
}
func genFn1(p0 int32, p1 int32) int8 {
    v11 := GenS1{F0: 673, F1: 25, F2: 2}
    p1 -= p1
    for fi12 := 0; fi12 < 5; fi12++ {
        if ("a" != "hello") {
            switch ("a" + "") {
            case "":
                v13 := []int{(fi12 | (fi12 << ((1) & 63))), 0, -9223372036854775808, fi12}
                _ = v13
            case "foo bar":
                for fi14 := 0; fi14 < 17; fi14++ {
                    v15 := []int{(fi12 - fi14), (fi14 * -892)}
                    var v16 uint16 = 761
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "a":
                var v17 uint32 = 2
                _ = v17
            default:
                v18 := []int{fi12, (fi12 & fi12)}
                _ = v18
            }
        } else {
            fmt.Println(p0)
            v19 := GenS2{F0: -377, F1: 42}
            var v20 string = "a"
            _ = v19
            _ = v20
        }
        var v21 int = (^(fi12))
        _ = fi12
        _ = v21
    }
    p0 -= p1
    _ = p0
    _ = p1
    _ = v11
    return 42
}

func main() {
    var v1 uint64 = 0
    var v2 int16 = 0
    var v3 uint8 = 108
    v4 := make(map[int]uint32)
    v2 = v2
    if true {
        var v22 float64 = 0.1
        var v23 string = "foo bar"
        _ = v22
        _ = v23
    } else {
        v3 -= v3
        delete(v4, 32768)
    }
    var v24 uint8 = v3
    delete(v4, 673)
    switch -1 {
    case 1:
        for fi25 := 0; fi25 < 5; fi25++ {
            var v26 uint64 = (v1 - (v1 / (((v1 ^ v1)) | 1)))
            v27 := GenS2{F0: v2, F1: 1}
            for fi28 := 0; fi28 < 11; fi28++ {
                v4[2147483647] = 2
                v29 := GenS2{F0: (v2 / ((42) | 1)), F1: int64(v2)}
                _ = fi28
                _ = v29
            }
            _ = fi25
            _ = v26
            _ = v27
        }
    default:
        v2 = -32768
    }
    var v30 uint8 = (v3 % ((v3) | 1))
    switch -1 {
    case -9223372036854775808:
        v2 -= 231
    case 45:
        v30 = (v30 >> ((1) & 63))
    default:
        v31 := GenS2{F0: (v2 / ((32767) | 1)), F1: 0}
        _ = v31
    }
    var v32 int = 32767
    if (42 != int64(v24)) {
        for fi33 := 0; fi33 < 11; fi33++ {
            for fi34 := 0; fi34 < 8; fi34++ {
                v24 -= (^(v24))
                var v35 uint16 = 65535
                switch 42 {
                case 9223372036854775807:
                    var v36 bool = (42 <= -2147483648)
                    _ = v36
                case 128:
                    v3 -= 1
                case 463:
                    v1 = (^(v1))
                default:
                    var v37 float64 = 0.1
                    _ = v37
                }
                _ = fi34
                _ = v35
            }
            v24 = v30
            v4[127] = 2
            _ = fi33
        }
    } else {
        v38 := GenS2{F0: (v2 ^ (v2 << ((42) & 63))), F1: -9223372036854775808}
        v32 = int(v2)
        switch "hello" {
        case "hello":
            fmt.Println(v1, v3)
        }
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v32
    _ = v4
}
