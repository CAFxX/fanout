package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint }
type GenS2 struct { F0 uint64 }
func genFn0() int8 {
    v6 := GenS1{F0: 127, F1: 0, F2: 42}
    var v7 string = "a"
    switch 0 {
    case 534:
        switch v7 {
        case "":
            if ((!(42 == -61)) || true) {
                var v8 string = v7
                _ = v8
            }
        default:
            switch 0 {
            case 535:
                if false {
                    var v9 int32 = 0
                    var v10 int16 = 0
                    _ = v10
                    _ = v9
                }
            case 0:
                var v11 int = 9223372036854775807
                _ = v11
            case 1:
                var v12 uint16 = 1
                _ = v12
            }
        }
    case 127:
        fmt.Println(v7)
    case -9223372036854775808:
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := []int{}
            switch fi13 {
            case -9223372036854775808:
                var v15 uint = 635
                _ = v15
            case -335:
                if (true && (1 == 123)) {
                    switch fi13 {
                    case 128:
                        v16 := GenS1{F0: -128, F1: fi13, F2: 1}
                        _ = v16
                    }
                    var v17 uint16 = 1
                    _ = v17
                }
            case 32768:
                if len(v14) > 0 {
                    v14[0] = fi13
                }
            default:
                v18 := GenS2{F0: 4294967295}
                _ = v18
            }
            v19 := GenS2{F0: 42}
            _ = fi13
            _ = v14
            _ = v19
        }
    }
    _ = v6
    _ = v7
    return 42
}
func genFn1() string {
    v20 := []int{42}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    var v21 float64 = -2.25
    v21 = v21
    _ = v20
    _ = v21
    return ""
}

func main() {
    var v1 uint = 65535
    var v2 int = -9223372036854775808
    v3 := make(map[uint32]bool)
    var v4 int32 = -583
    var v5 string = "a"
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 uint8 = 255
        if true {
            fmt.Println(v4, v2)
            for fi24 := 0; fi24 < 7; fi24++ {
                v4 -= v4
                _ = fi24
            }
        } else {
            var v25 int = fi22
            var v26 int = (v25 / ((42) | 1))
            delete(v3, 502)
            _ = v25
            _ = v26
        }
        switch ((((v5 + (v5 + v5)) + (v5 + v5)) + ((v5 + "") + v5)) + ((v5 + "hello") + v5)) {
        case "\n\t":
            var v27 int16 = -879
            _ = v27
        case "hello":
            fmt.Println(v1, v5)
        default:
            var v28 uint64 = 4294967295
            _ = v28
        }
        _ = fi22
        _ = v23
    }
    v29 := GenS2{F0: 0}
    v3[0] = (-412 != -261)
    if (1 > -569) {
        v30 := GenS1{F0: 20, F1: (^(v2)), F2: v1}
        v5 = ""
        _ = v30
    }
    if false {
        for fi31 := 0; fi31 < 9; fi31++ {
            fmt.Println(v2)
            _ = fi31
        }
    } else {
        for fi32 := 0; fi32 < 1; fi32++ {
            fmt.Println(v5, v5)
            var v33 int16 = -1
            _ = fi32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        v1 = 486
        _ = fi34
    }
    fmt.Println(v2, v4, v5)
    v35 := []int{(v2 | (v2 | (v2 + (v2 << ((42) & 63)))))}
    var v36 uint64 = 42
    if (3.14159 <= 1.5) {
        var v37 uint = (v1 & (v1 * v1))
        switch v5 {
        case "a":
            var v38 int16 = 32767
            _ = v38
        case "\n\t":
            if true {
                var v39 uint16 = uint16(v36)
                _ = len(v35)
                var v40 int = -2147483648
                _ = v39
                _ = v40
            } else {
                v36 += v36
                v36 = v36
                _ = len(v35)
            }
        }
        v37 += 596
        _ = v37
    } else {
        v3[0] = (((((float32(-2.25) / 1.5) / (0.1 / float32(-2.25))) != 0.1) && true) && (false && true))
        if true {
            if false {
                v4 = v4
            } else {
                v1 = v1
            }
            v35 = append(v35, v2)
        } else {
            switch v2 {
            case 9223372036854775807:
                var v41 uint = v1
                _ = v41
            }
            fmt.Println(v36, v5, v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
