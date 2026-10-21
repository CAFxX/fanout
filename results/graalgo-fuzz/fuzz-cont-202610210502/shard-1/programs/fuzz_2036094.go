package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 bool }
type GenS2 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint8; F2 string }
type GenS4 struct { F0 uint8; F1 int32 }
func genFn0(p0 bool, p1 int64) uint64 {
    p0 = (!(255 <= 65535))
    if p0 {
        p1 = (p1 >> ((255) & 63))
        var v4 int32 = -2147483648
        for fi5 := 0; fi5 < 7; fi5++ {
            if (!p0) {
                v4 = 1
                if false {
                    var v6 int16 = 0
                    fmt.Println(p0)
                    _ = v6
                }
                v7 := GenS1{F0: 1, F1: p0, F2: (true || ((2 <= 189) || p0))}
                _ = v7
            }
            p1 = p1
            _ = fi5
        }
        _ = v4
    } else {
        v8 := []int{-146, 32768, 2147483647, 32768}
        v8 = append(v8, 4294967295)
        if true {
            if (255 <= 42) {
                v8 = append(v8, 1)
                p1 -= (p1 + (p1 % (((p1 - p1)) | 1)))
            }
            v9 := GenS2{F0: 0, F1: 127, F2: 69}
            var v10 uint32 = 4294967295
            _ = v10
            _ = v9
        } else {
            if len(v8) > 0 {
                v8[0] = 4294967295
            }
            var v11 uint = 42
            _ = v11
        }
        _ = v8
    }
    _ = p0
    _ = p1
    return 454
}

func main() {
    var v1 int8 = 100
    v2 := make(map[string]bool)
    var v3 int = 32767
    v3 = -2147483648
    v12 := GenS2{F0: -32768, F1: 1, F2: 22}
    for fi13 := 0; fi13 < 18; fi13++ {
        v3 -= (fi13 - -1)
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 int64 = 688
            if false {
                v16 := GenS1{F0: 1, F1: ((235 >= 2147483648) || false), F2: ((v1 ^ v1) < -128)}
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    var v17 uint = 42
    for fi18 := 0; fi18 < 12; fi18++ {
        switch fi18 {
        case 226:
            v19 := GenS1{F0: 2147483647, F1: true, F2: true}
            _ = v19
        default:
            v20 := GenS3{F0: 41, F1: 201, F2: (((("" + ("a" + ("" + "a"))) + "") + "hello") + (("hello" + "hello") + "hello"))}
            _ = v20
        }
        delete(v2, "hello")
        if (!(!((("hello" + ("foo bar" + "")) + "\n\t") > ("hello" + "\n\t")))) {
            _ = len(v2)
            switch "hello" {
            case "foo bar":
                v17 += 42
            case "hello":
                for fi21 := 0; fi21 < 7; fi21++ {
                    var v22 int16 = 1
                    _ = fi21
                    _ = v22
                }
            }
            v23 := GenS4{F0: 1, F1: -2147483648}
            _ = v23
        } else {
            v24 := GenS2{F0: 32767, F1: 128, F2: 170}
            var v25 int8 = (v1 ^ v1)
            _ = v24
            _ = v25
        }
        _ = fi18
    }
    v1 = 127
    v3 += v3
    var v26 int32 = -768
    switch (v3 + v3) {
    case -9223372036854775808:
        v27 := []int{(v3 << ((v17) & 63))}
        _ = v27
    case -2147483648:
        fmt.Println(v1, v26)
    case 793:
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 int = (v3 * 46)
            fmt.Println(v17, v26)
            var v30 uint = v17
            _ = fi28
            _ = v29
            _ = v30
        }
    default:
        if true {
            var v31 int8 = -6
            v32 := GenS4{F0: 218, F1: (v26 ^ (v26 | (v26 * (^(v26)))))}
            _ = v31
            _ = v32
        } else {
            v33 := GenS1{F0: 0, F1: true, F2: (!true)}
            var v34 int8 = -16
            _ = v33
            _ = v34
        }
    }
    v17 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v26
    _ = v3
}
