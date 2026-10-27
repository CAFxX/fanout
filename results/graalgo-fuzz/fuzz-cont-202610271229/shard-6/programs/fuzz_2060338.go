package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint }
func genFn0() string {
    v4 := []int{1, -303, -2147483648, 4294967295}
    v5 := GenS1{F0: 0, F1: -57, F2: 128}
    var v6 uint8 = 29
    _ = v4
    _ = v5
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 1.5
    v3 := make(map[string]int)
    v2 += v2
    v2 = -2.25
    fmt.Println(v2)
    if true {
        if true {
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := v3["hello"]
                _ = fi7
                _ = v8
            }
            _ = len(v1)
        } else {
            var v9 uint32 = 65535
            if (0 < 1) {
                fmt.Println(v9, v2)
                fmt.Println(v9)
            } else {
                for fi10 := 0; fi10 < 16; fi10++ {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        var v12 uint = 606
                        v13 := []int{fi11, (fi11 ^ (fi11 ^ (fi11 + (fi10 ^ (fi11 ^ (fi11 + -2147483648))))))}
                        v14 := v1[fi11]
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = fi10
                }
            }
            v15 := GenS1{F0: 759, F1: 127, F2: 128}
            _ = v15
            _ = v9
        }
    }
    fmt.Println(v2)
    var v16 uint64 = 1
    if (!false) {
        switch -129 {
        case 663:
            fmt.Println(v16)
        case 1:
            v16 = v16
        case -884:
            v2 -= v2
        default:
            var v17 int16 = -659
            _ = v17
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            for fi19 := 0; fi19 < 16; fi19++ {
                v20 := []int{fi19, fi19, fi19}
                fmt.Println(v2)
                _ = fi19
                _ = v20
            }
            switch (fi18 / ((fi18) | 1)) {
            case 42:
                fmt.Println(fi18, v16, v2)
            case 633:
                v2 = v2
            default:
                if ((true && false) && false) {
                    for fi21 := 0; fi21 < 5; fi21++ {
                        v1[42] = ("" + "")
                        v22 := GenS1{F0: 2, F1: int8(v16), F2: 178}
                        _ = fi21
                        _ = v22
                    }
                } else {
                    v23 := GenS1{F0: 206, F1: 0, F2: 0}
                    v16 += (v16 << ((336) & 63))
                    _ = v23
                }
            }
            _ = fi18
        }
        if (0.1 >= 0.1) {
            for fi24 := 0; fi24 < 14; fi24++ {
                var v25 int = (fi24 % ((fi24) | 1))
                v26 := GenS2{F0: 373, F1: 2147483648, F2: 647}
                _ = fi24
                _ = v25
                _ = v26
            }
            if false {
                switch 4294967295 {
                case 0:
                    v16 += (v16 & 2)
                case 1:
                    v2 += float32(1e10)
                case 4294967295:
                    var v27 float64 = 0.0
                    _ = v27
                default:
                    v28 := GenS2{F0: 255, F1: 65535, F2: 668}
                    _ = v28
                }
                v2 -= (v2 - 1e10)
            } else {
                for fi29 := 0; fi29 < 4; fi29++ {
                    var v30 float32 = v2
                    v30 = v2
                    for fi31 := 0; fi31 < 16; fi31++ {
                        v16 -= (v16 >> ((65535) & 63))
                        v32 := GenS2{F0: 113, F1: 4294967295, F2: 42}
                        _ = fi31
                        _ = v32
                    }
                    _ = fi29
                    _ = v30
                }
            }
            v2 = 0.1
        }
    }
    if false {
        v2 = 1e10
    } else {
        v33 := []int{128}
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
}
