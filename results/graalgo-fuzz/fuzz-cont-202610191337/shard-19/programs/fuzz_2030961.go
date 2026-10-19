package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 uint64 }
func genFn0() int32 {
    if true {
        switch -25 {
        case -583:
            v7 := GenS1{F0: 1, F1: (1 > 2), F2: 660}
            _ = v7
        default:
            var v8 uint32 = 2147483647
            _ = v8
        }
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int = fi9
            var v11 uint16 = 255
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 int64 = 0
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 int = fi12
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    } else {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 float64 = (-2.25 - -2.25)
            _ = fi16
            _ = v17
        }
    }
    var v18 string = ("\n\t" + "")
    v18 = ("" + v18)
    for fi19 := 0; fi19 < 20; fi19++ {
        v18 = (v18 + v18)
        _ = fi19
    }
    _ = v18
    return -585
}

func main() {
    var v1 int64 = 926
    var v2 uint8 = 128
    var v3 int64 = -1
    var v4 float64 = 0.1
    var v5 int = 128
    var v6 uint8 = 1
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 int64 = (v1 % ((-263) | 1))
        var v22 uint = 65535
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println(v4, v5, v2)
    var v23 uint16 = 255
    if (v1 < -1) {
        v24 := GenS2{F0: "a", F1: 4294967295}
        fmt.Println(v5, v6)
        _ = v24
    } else {
        for fi25 := 0; fi25 < 12; fi25++ {
            if (!false) {
                var v26 int = 2147483647
                var v27 uint8 = 222
                _ = v26
                _ = v27
            }
            v28 := GenS3{F0: v23}
            if (v5 >= (-(fi25))) {
                v29 := []int{fi25}
                v30 := GenS4{F0: fi25}
                var v31 int32 = genFn0()
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi25
            _ = v28
        }
        switch v5 {
        case -560:
            v32 := []int{v5, 32768, -2147483648}
            _ = v32
        case -9223372036854775808:
            v33 := GenS3{F0: 65535}
            _ = v33
        default:
            switch v5 {
            case 141:
                var v34 uint8 = (v6 ^ (v2 * v2))
                _ = v34
            default:
                v5 -= (v5 << ((65535) & 63))
            }
        }
        v23 += 2
    }
    var v35 int64 = 1
    v1 = v35
    var v36 bool = true
    v37 := GenS5{F0: v35, F1: 237}
    v38 := []int{(-(v5)), (-(v5))}
    fmt.Println(v23, v6, v4)
    v39 := GenS4{F0: 127}
    var v40 int16 = -806
    for fi41 := 0; fi41 < 3; fi41++ {
        fmt.Println(v23)
        _ = fi41
    }
    for fi42 := 0; fi42 < 6; fi42++ {
        _ = len(v38)
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
