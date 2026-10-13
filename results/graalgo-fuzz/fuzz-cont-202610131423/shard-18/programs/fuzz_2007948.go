package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int8 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS4 struct { F0 int; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int32; F1 int64; F2 int32 }
func genFn0() uint32 {
    switch 898 {
    case 1:
        v4 := GenS1{F0: -1, F1: 472, F2: 70}
        _ = v4
    }
    v5 := GenS1{F0: -834, F1: 32767, F2: 0}
    var v6 uint = 4294967295
    switch 9223372036854775807 {
    case 128:
        v7 := GenS2{F0: "", F1: uint64(v6)}
        _ = v7
    case 32767:
        v6 = (v6 * v6)
    default:
        var v8 uint = v6
        _ = v8
    }
    if (!(0 < 1000)) {
        v9 := []int{-626, 32767, 581, -485}
        v6 += v6
        if len(v9) > 0 {
            v9[0] = -9223372036854775808
        }
        _ = v9
    } else {
        if ((32767 >= 714) && (false || false)) {
            var v10 int16 = 793
            var v11 int16 = (v10 + 0)
            _ = v10
            _ = v11
        }
        v12 := []int{int(v6), 2147483647}
        _ = v12
    }
    _ = v5
    _ = v6
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -1
    var v3 uint16 = 461
    v2 += (v2 << ((2) & 63))
    v13 := GenS3{F0: 2147483647, F1: ((65535 > 638) || false), F2: v3}
    for fi14 := 0; fi14 < 7; fi14++ {
        switch 42 {
        case 128:
            var v15 string = "\n\t"
            _ = v15
        case 176:
            v16 := []int{-345, fi14, 388}
            _ = v16
        case 2147483647:
            v17 := GenS3{F0: 2147483647, F1: false, F2: (v3 << ((194) & 63))}
            _ = v17
        default:
            if false {
                for fi18 := 0; fi18 < 10; fi18++ {
                    fmt.Println(v2)
                    _ = fi18
                }
            } else {
                if false {
                    v3 += 65535
                    var v19 uint16 = 32767
                    _ = v19
                }
            }
        }
        _ = fi14
    }
    v2 = 0
    if false {
        var v20 uint8 = 42
        v21 := GenS4{F0: 131, F1: (^(v20)), F2: 42}
        _ = v20
        _ = v21
    }
    var v22 uint32 = 2
    if true {
        v23 := GenS5{F0: 2147483647, F1: 42, F2: 2147483647}
        _ = len(v1)
        _ = v23
    }
    if (!false) {
        if false {
            var v24 string = (("hello" + "foo bar") + ("foo bar" + ""))
            v2 -= v2
            _ = v24
        } else {
            v2 = v2
            v25 := GenS2{F0: (("" + "\n\t") + "hello"), F1: 1}
            fmt.Println(v22)
            _ = v25
        }
        if (((true && true) || true) || false) {
            var v26 uint8 = 194
            _ = v26
        } else {
            v27 := GenS1{F0: (v2 / ((-1) | 1)), F1: 1000, F2: 42}
            _ = v27
        }
        var v28 string = "foo bar"
        _ = v28
    } else {
        v29 := GenS1{F0: v2, F1: 1, F2: 127}
        v2 = (v2 * int64(v3))
        _ = v29
    }
    if true {
        if (4294967295 <= 42) {
            if false {
                v30 := v1[32767]
                v31 := GenS3{F0: v30, F1: true, F2: v3}
                _ = v30
                _ = v31
            }
            v2 += v2
        }
        for fi32 := 0; fi32 < 3; fi32++ {
            v33 := v1[943]
            _ = fi32
            _ = v33
        }
        switch (("foo bar" + "foo bar") + "foo bar") {
        case "":
            fmt.Println(v3, v22)
        case "foo bar":
            v34 := []int{127, 4294967295}
            _ = v34
        case "a":
            var v35 uint64 = 526
            _ = v35
        default:
            for fi36 := 0; fi36 < 3; fi36++ {
                switch (fi36 / ((fi36) | 1)) {
                case -1:
                    if (false || ((1.5 / -2.25) == ((((float32(1e10) * 1e10) + 0.0) - ((3.14159 - 0.1) * 1e10)) * (1e10 * (((((0.0 * 1.5) / 1.5) - (3.14159 - (3.14159 + 3.14159))) * (1e10 / ((((0.1 - 1.5) - (0.1 * float32(0.0))) / -2.25) * 0.1))) + ((-2.25 - 0.1) / 1.5)))))) {
                        var v37 int = fi36
                        _ = v37
                    }
                case 42:
                    _ = len(v1)
                default:
                    var v38 uint16 = (^(v3))
                    _ = v38
                }
                v39 := []int{fi36, (fi36 - -336), 9223372036854775807}
                fmt.Println(fi36)
                _ = fi36
                _ = v39
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
}
