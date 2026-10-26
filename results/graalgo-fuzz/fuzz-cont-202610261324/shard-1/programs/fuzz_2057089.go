package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint16 }
type GenS2 struct { F0 int64; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]string)
    var v3 int16 = 1
    var v4 uint64 = 807
    v5 := make(map[string]uint32)
    var v6 int16 = (-(v3))
    if (9223372036854775807 != 42) {
        if false {
            var v7 int16 = v6
            switch 4294967295 {
            case -569:
                v3 = 32767
            case -814:
                v6 = 0
            default:
                v8 := []int{128}
                _ = v8
            }
            _ = v7
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 uint16 = 854
            _ = fi9
            _ = v10
        }
        var v11 string = ((((("\n\t" + "a") + "hello") + "") + "") + (("foo bar" + "\n\t") + ("" + "a")))
        _ = v11
    }
    v12 := GenS1{F0: 78, F1: (!false), F2: 211}
    fmt.Println(v6)
    if (!(false || (!true))) {
        v4 = v4
        var v13 uint32 = 2147483648
        _ = v13
    } else {
        switch -2147483648 {
        case 4294967295:
            if (true || true) {
                var v14 uint64 = 1
                var v15 int = 127
                for fi16 := 0; fi16 < 16; fi16++ {
                    v17 := GenS1{F0: 255, F1: (2 <= 892), F2: 662}
                    _ = fi16
                    _ = v17
                }
                _ = v14
                _ = v15
            }
        case 32768:
            _ = len(v5)
        case -829:
            if true {
                var v18 uint8 = 42
                v4 = 1
                var v19 bool = (true || (v4 >= (v4 << ((0) & 63))))
                _ = v18
                _ = v19
            } else {
                v3 = -759
                v20 := GenS2{F0: 1, F1: 0}
                _ = v20
            }
        }
        v6 += (v3 & v6)
    }
    v3 += 795
    v6 = (^(v3))
    v3 -= (v3 / (((v6 & (v6 >> ((uint(v4)) & 63)))) | 1))
    v21 := GenS2{F0: -310, F1: 2147483647}
    fmt.Println(v6)
    var v22 int16 = 538
    fmt.Println(v22, v4)
    v23 := GenS1{F0: 507, F1: false, F2: 254}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
