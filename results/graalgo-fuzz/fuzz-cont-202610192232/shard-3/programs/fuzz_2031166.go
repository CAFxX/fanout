package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 float32 = float32(float64(-2.25))
    var v2 uint16 = 65535
    v3 := make(map[int]uint32)
    var v4 uint64 = 2
    var v5 uint32 = 1
    switch "foo bar" {
    case "foo bar":
        var v6 int = 9223372036854775807
        _ = v6
    case "a":
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: ((float32((-2.25 / (float64(1e10) + (float64(v1) - 3.14159)))) * v1) != v1), F1: (false || false)}
            _ = fi7
            _ = v8
        }
    case "":
        if true {
            var v9 bool = (v5 < (v5 - 2))
            var v10 float64 = 1.5
            if (((v1 * float32((0.0 + (1.5 + 1.5)))) + (3.14159 * -2.25)) != v1) {
                for fi11 := 0; fi11 < 13; fi11++ {
                    if v9 {
                        v12 := GenS1{F0: v9, F1: (4294967295 != (-(fi11)))}
                        v13 := GenS2{F0: 0}
                        _ = v12
                        _ = v13
                    } else {
                        delete(v3, fi11)
                        v1 += 0.1
                    }
                    var v14 int16 = 651
                    v15 := v3[(fi11 & fi11)]
                    _ = fi11
                    _ = v14
                    _ = v15
                }
                v9 = false
            }
            _ = v10
            _ = v9
        }
    }
    v4 = (v4 << ((104) & 63))
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS2{F0: 69}
        switch fi16 {
        case 32768:
            v18 := GenS2{F0: 106}
            _ = v18
        case -454:
            var v19 int8 = -14
            _ = v19
        case 42:
            var v20 uint16 = 423
            _ = v20
        default:
            if ((!true) || (65535 >= 461)) {
                var v21 int = int(v5)
                _ = v21
            }
        }
        _ = fi16
        _ = v17
    }
    var v22 int16 = -1
    var v23 int16 = v22
    v24 := v3[int(v23)]
    for fi25 := 0; fi25 < 14; fi25++ {
        v26 := GenS2{F0: -56}
        v27 := []int{-81, (fi25 % (((fi25 + fi25)) | 1))}
        _ = fi25
        _ = v26
        _ = v27
    }
    switch "hello" {
    case "foo bar":
        if ((0 < 42) || true) {
            var v28 uint8 = 1
            for fi29 := 0; fi29 < 3; fi29++ {
                v22 = (v23 - 32767)
                for fi30 := 0; fi30 < 14; fi30++ {
                    var v31 uint32 = v24
                    _ = fi30
                    _ = v31
                }
                _ = fi29
            }
            for fi32 := 0; fi32 < 4; fi32++ {
                fmt.Println(v5, v5)
                var v33 string = ("" + ("" + ("a" + "")))
                _ = fi32
                _ = v33
            }
            _ = v28
        } else {
            delete(v3, 4294967295)
            fmt.Println(v23, v2)
        }
    }
    v1 += v1
    v3[32768] = v24
    if (42 <= (^(v2))) {
        v5 = v24
    }
    v34 := GenS3{F0: uint16(v24)}
    var v35 uint = 848
    if ((v35 - (^(v35))) <= v35) {
        for fi36 := 0; fi36 < 7; fi36++ {
            v37 := v3[2147483647]
            v38 := []int{}
            _ = fi36
            _ = v37
            _ = v38
        }
        var v39 string = ("foo bar" + "a")
        _ = v39
    }
    v40 := []int{2147483647, 4294967295, 2147483647, 1}
    v35 = v35
    var v41 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
