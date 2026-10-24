package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 1
    v3 := make(map[string]bool)
    if true {
        v4 := []int{}
        var v5 bool = (0 != -30)
        _ = v4
        _ = v5
    }
    v2 += (v2 ^ (^(v2)))
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := GenS1{F0: "a"}
            var v9 uint32 = 352
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v2 += (v2 * (v2 >> ((65535) & 63)))
    v10 := []int{127}
    v1[-2147483648] = 0
    v11 := GenS1{F0: ("foo bar" + ("" + ("hello" + "foo bar")))}
    switch -1 {
    case -2147483648:
        delete(v1, 9223372036854775807)
    case 133:
        if true {
            _ = len(v1)
            v10 = append(v10, -219)
            if true {
                var v12 float32 = (3.14159 / (0.0 + (-2.25 / 1.0)))
                _ = v12
            }
        }
    default:
        v2 += (v2 * (v2 & (v2 % (((^(v2))) | 1))))
    }
    _ = len(v1)
    v2 = v2
    v2 = (v2 / ((v2) | 1))
    v1[911] = 2
    var v13 bool = true
    for fi14 := 0; fi14 < 12; fi14++ {
        if ((((((fi14 * (fi14 + fi14)) > fi14) || ("foo bar" > "")) || v13) && v13) || v13) {
            var v15 bool = v13
            _ = v15
        }
        for fi16 := 0; fi16 < 14; fi16++ {
            v2 = (v2 << ((486) & 63))
            var v17 int16 = 1000
            _ = fi16
            _ = v17
        }
        _ = fi14
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        if true {
            if v13 {
                for fi19 := 0; fi19 < 11; fi19++ {
                    if v13 {
                        fmt.Println(fi18, v2, fi19)
                    } else {
                        var v20 int64 = 510
                        _ = v20
                    }
                    _ = fi19
                }
            } else {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v13 = false
                for fi21 := 0; fi21 < 17; fi21++ {
                    v13 = ("hello" <= "")
                    _ = fi21
                }
            }
            delete(v3, "a")
            v22 := GenS1{F0: "a"}
            _ = v22
        } else {
            if false {
                var v23 int16 = -32768
                _ = v23
            } else {
                var v24 float64 = float64((0.0 / 1.5))
                for fi25 := 0; fi25 < 11; fi25++ {
                    fmt.Println(v24)
                    v26 := GenS2{F0: 42, F1: 376, F2: 1000}
                    _ = fi25
                    _ = v26
                }
                if v13 {
                    v24 -= v24
                    v2 = v2
                    if len(v10) > 0 {
                        v10[0] = 32767
                    }
                }
                _ = v24
            }
        }
        _ = fi18
    }
    switch -424 {
    case 9223372036854775807:
        v2 = v2
    case 3:
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 uint32 = uint32(v2)
            for fi29 := 0; fi29 < 20; fi29++ {
                fmt.Println(v2, v13, v28)
                if (!v13) {
                    var v30 uint = 655
                    _ = v30
                } else {
                    var v31 float64 = (1e10 + 1.5)
                    _ = v31
                }
                v32 := GenS1{F0: (("hello" + ("" + "a")) + (("hello" + ("foo bar" + "hello")) + "\n\t"))}
                _ = fi29
                _ = v32
            }
            fmt.Println(v13, v28)
            _ = fi27
            _ = v28
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
}
