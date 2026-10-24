package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
func genFn0() int16 {
    if true {
        var v5 float32 = -2.25
        _ = v5
    } else {
        var v6 int16 = -32768
        if (1 >= 281) {
            fmt.Println(v6)
            if false {
                if false {
                    v7 := GenS1{F0: 977}
                    _ = v7
                } else {
                    v6 += v6
                }
                for fi8 := 0; fi8 < 5; fi8++ {
                    v9 := GenS2{F0: -8}
                    _ = fi8
                    _ = v9
                }
            } else {
                v6 -= v6
            }
        }
        _ = v6
    }
    var v10 float64 = 3.14159
    _ = v10
    return -1
}
func genFn1(p0 int, p1 int64, p2 int32) bool {
    for fi11 := 0; fi11 < 18; fi11++ {
        if ((p2 <= p2) || (!true)) {
            if (100 <= 0) {
                p2 = p2
                var v12 int8 = -17
                _ = v12
            }
            switch -9223372036854775808 {
            case 1:
                v13 := GenS1{F0: -682}
                _ = v13
            case 0:
                v14 := GenS3{F0: 0}
                _ = v14
            default:
                var v15 uint = 65535
                _ = v15
            }
            var v16 float32 = 3.14159
            _ = v16
        } else {
            switch ("hello" + ("foo bar" + "\n\t")) {
            case "":
                if true {
                    v17 := GenS2{F0: 0}
                    _ = v17
                }
            }
        }
        switch fi11 {
        case -2147483648:
            var v18 uint8 = 2
            _ = v18
        case 386:
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 float64 = float64((0.1 * 0.0))
                _ = fi19
                _ = v20
            }
        default:
            var v21 uint8 = 65
            _ = v21
        }
        v22 := GenS2{F0: -22}
        _ = fi11
        _ = v22
    }
    var v23 int64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return true
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]uint32)
    var v3 uint64 = 42
    var v4 uint32 = 0
    v1 -= v1
    var v24 int64 = 42
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 bool = false
        v3 = v3
        if ((v24 >> ((42) & 63)) <= -9223372036854775808) {
            v27 := GenS4{F0: 42}
            _ = v27
        }
        _ = fi25
        _ = v26
    }
    v3 = v3
    v24 = (v24 ^ v24)
    var v28 float32 = -2.25
    v1 = v1
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := []int{fi29, -933, (fi29 * fi29)}
        v24 -= (-(v24))
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
}
