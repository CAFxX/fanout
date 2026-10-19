package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS5 struct { F0 int16; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 42
    var v3 string = ("foo bar" + "\n\t")
    for fi4 := 0; fi4 < 9; fi4++ {
        v2 += v2
        switch fi4 {
        case 557:
            v5 := v1[364]
            _ = v5
        }
        _ = fi4
    }
    v2 -= (-(v2))
    v2 = v2
    v2 += v2
    if (42 != 2) {
        var v6 float32 = 1e10
        v1[65535] = 32768
        _ = v6
    } else {
        if false {
            var v7 bool = (!false)
            switch 127 {
            case 32767:
                switch -9223372036854775808 {
                case 127:
                    v8 := GenS1{F0: 32768, F1: 1}
                    _ = v8
                case -261:
                    fmt.Println(v7, v2, v7)
                }
            case -675:
                v7 = true
            default:
                switch (("hello" + v3) + v3) {
                case "":
                    v1[898] = -9223372036854775808
                default:
                    fmt.Println(v2)
                }
            }
            _ = v7
        }
        var v9 int32 = 0
        _ = v9
    }
    v3 = "foo bar"
    v10 := GenS2{F0: 254}
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{fi11, (^(fi11)), fi11}
        v3 = (v3 + v3)
        v13 := GenS3{F0: 80}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 float32 = (-2.25 * ((1.5 - 1e10) - 0.0))
    v14 = (1e10 / 1.0)
    for fi15 := 0; fi15 < 1; fi15++ {
        v3 = "foo bar"
        _ = fi15
    }
    v16 := GenS1{F0: 128, F1: 0}
    var v17 string = v3
    v3 = v17
    v3 = v3
    v18 := GenS4{F0: 0, F1: 132, F2: -731}
    v14 -= 1e10
    for fi19 := 0; fi19 < 4; fi19++ {
        switch (-(fi19)) {
        case 9223372036854775807:
            if ((fi19 - fi19) >= fi19) {
                v2 = v2
                v20 := GenS5{F0: -421, F1: 131, F2: 0}
                _ = v20
            } else {
                v17 = v3
                var v21 string = (v17 + v3)
                _ = v21
            }
        case -2147483648:
            for fi22 := 0; fi22 < 9; fi22++ {
                v17 = v3
                if ((true && true) || (v17 >= "foo bar")) {
                    fmt.Println(fi22)
                    v2 = (v2 << ((42) & 63))
                } else {
                    var v23 float32 = float32(float64(3.14159))
                    v3 = "foo bar"
                    v24 := GenS4{F0: 835, F1: 128, F2: 42}
                    _ = v23
                    _ = v24
                }
                var v25 int16 = 327
                var v26 float32 = v14
                _ = fi22
                _ = v25
                _ = v26
            }
        default:
            v3 = "\n\t"
        }
        if true {
            v2 -= (v2 - v2)
        } else {
            for fi27 := 0; fi27 < 9; fi27++ {
                v2 += (v2 >> ((18446744073709551615) & 63))
                switch fi27 {
                case 2147483647:
                    if false {
                        v14 -= 1e10
                        v28 := []int{fi27, (fi27 & fi19)}
                        fmt.Println(fi27, fi19)
                        _ = v28
                    }
                case -2147483648:
                    v3 = v17
                case -251:
                    var v29 int8 = -8
                    _ = v29
                default:
                    v30 := []int{fi27, (fi19 >> ((0) & 63)), fi27}
                    if ((0 == -153) || false) {
                        v3 = (("hello" + "\n\t") + v3)
                    } else {
                        var v31 float32 = (0.0 * 0.0)
                        v30 = append(v30, 0)
                        v2 -= -70
                        _ = v31
                    }
                    v32 := GenS6{F0: 255}
                    _ = v30
                    _ = v32
                }
                v33 := GenS2{F0: 779}
                _ = fi27
                _ = v33
            }
        }
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
