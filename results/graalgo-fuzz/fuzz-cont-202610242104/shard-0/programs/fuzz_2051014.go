package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint = 2
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    v1 -= v1
    if false {
        delete(v3, "foo bar")
    } else {
        for fi4 := 0; fi4 < 18; fi4++ {
            switch (fi4 >> ((v1) & 63)) {
            case -67:
                var v5 int16 = 639
                _ = v5
            case 32768:
                if (128 == 193) {
                    v6 := GenS1{F0: fi4, F1: 18446744073709551615, F2: 1}
                    _ = v6
                }
            case 2147483647:
                v1 -= v1
            default:
                for fi7 := 0; fi7 < 19; fi7++ {
                    var v8 int16 = -32768
                    var v9 uint16 = 406
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
            v10 := GenS1{F0: fi4, F1: 0, F2: 0}
            v11 := []int{-9223372036854775808, (fi4 << ((v1) & 63)), (fi4 ^ (fi4 + (fi4 & (^(fi4)))))}
            _ = fi4
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            fmt.Println(v1)
            _ = fi12
        }
        if true {
            v3["a"] = 992
        } else {
            if (!false) {
                for fi13 := 0; fi13 < 6; fi13++ {
                    v14 := []int{}
                    _ = fi13
                    _ = v14
                }
                var v15 int16 = 355
                _ = v15
            } else {
                switch 4294967295 {
                case 963:
                    v1 = v1
                default:
                    v16 := GenS1{F0: -298, F1: 491, F2: 0}
                    _ = v16
                }
            }
            v17 := []int{42}
            _ = v17
        }
    }
    var v18 int32 = 770
    fmt.Println(v18, v18)
    var v19 string = "a"
    v19 = ((("" + "") + (v19 + "\n\t")) + ("" + (v19 + (((v19 + v19) + "\n\t") + v19))))
    var v20 int32 = (-(v18))
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 int64 = -9223372036854775808
        _ = fi21
        _ = v22
    }
    v19 = (((v19 + ("foo bar" + v19)) + ("foo bar" + ("hello" + "\n\t"))) + v19)
    v23 := GenS2{F0: (65535 < 255), F1: 2}
    for fi24 := 0; fi24 < 17; fi24++ {
        if false {
            var v25 float32 = (1e10 * float32(3.14159))
            v26 := []int{fi24, 22, -1, fi24}
            var v27 int16 = -32768
            _ = v25
            _ = v26
            _ = v27
        } else {
            delete(v3, ("hello" + (v19 + v19)))
            fmt.Println(v19)
        }
        v18 += v18
        _ = fi24
    }
    var v28 int64 = -1
    fmt.Println(v19, v20)
    var v29 int64 = v28
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v28
    _ = v29
    _ = v3
}
