package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint32; F2 int32 }
func genFn0() int64 {
    var v7 uint64 = 2
    var v8 int16 = 217
    v7 += (v7 << ((0) & 63))
    var v9 int = 4294967295
    _ = v7
    _ = v8
    _ = v9
    return -804
}
func genFn1() int8 {
    if true {
        var v10 int8 = 63
        var v11 int32 = 42
        v11 = (v11 | 382)
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int8 = int8(fi12)
            var v14 uint32 = 276
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 int = 32768
        _ = v15
    }
    var v16 string = "\n\t"
    var v17 int64 = 0
    v16 = (v16 + v16)
    v18 := GenS1{F0: 255, F1: v16, F2: (v17 ^ v17)}
    _ = v16
    _ = v17
    _ = v18
    return -1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    v3 := []string{}
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]bool)
    var v19 uint = 255
    var v20 int8 = 0
    v20 = (v20 & genFn1())
    fmt.Println(v19, v20)
    if (float64((1e10 / 1.5)) <= (0.0 * 1e10)) {
        var v21 float32 = 0.0
        switch int(v19) {
        case 127:
            var v22 float64 = 0.0
            _ = v22
        default:
            var v23 uint8 = 127
            _ = v23
        }
        var v24 uint = (v19 | (^(v19)))
        _ = v21
        _ = v24
    } else {
        if len(v3) > 0 {
            v3[0] = ("hello" + "a")
        }
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 uint8 = 2
            if false {
                for fi27 := 0; fi27 < 13; fi27++ {
                    _ = len(v3)
                    v26 -= 127
                    if ("\n\t" > "\n\t") {
                        v28 := GenS2{F0: ("a" + "")}
                        _ = v28
                    }
                    _ = fi27
                }
            }
            v29 := GenS1{F0: (v26 | 2), F1: ("hello" + "foo bar"), F2: 526}
            _ = fi25
            _ = v26
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        for fi31 := 0; fi31 < 2; fi31++ {
            v32 := GenS3{F0: 127, F1: 2, F2: 225}
            fmt.Println(fi30, v20, fi30)
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    v2[-9223372036854775808] = 957
    var v33 int = -9223372036854775808
    v34 := v5[2147483648]
    var v35 uint64 = 0
    for fi36 := 0; fi36 < 6; fi36++ {
        v33 -= v33
        var v37 int64 = 42
        for fi38 := 0; fi38 < 8; fi38++ {
            v20 -= (v20 >> ((v19) & 63))
            _ = fi38
        }
        _ = fi36
        _ = v37
    }
    if (!(!(255 != 65535))) {
        if true {
            if false {
                v33 -= int(v20)
            }
            v19 += v19
            var v39 int16 = 42
            _ = v39
        } else {
            v34 = v34
            switch int(v35) {
            case 128:
                for fi40 := 0; fi40 < 3; fi40++ {
                    v41 := v1["\n\t"]
                    v35 += (v35 + 155)
                    _ = fi40
                    _ = v41
                }
            case 2147483647:
                _ = len(v4)
            }
            v34 = v34
        }
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        var v43 float32 = (((0.1 - 0.0) / ((((float32(-2.25) * 0.1) * 0.1) + 1.5) + (float32((1e10 * ((((3.14159 - 3.14159) + 1.5) * 1e10) * 0.1))) - 1e10))) + (3.14159 - ((float32(1e10) / ((0.1 - 3.14159) / 0.1)) / 1.5)))
        _ = fi42
        _ = v43
    }
    v34 = v34
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
