package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint16; F2 bool }
func genFn0() int8 {
    if (644 < 32767) {
        var v4 uint = 1
        _ = v4
    }
    var v5 uint8 = 2
    if false {
        if (v5 >= (v5 ^ (v5 + 0))) {
            switch 1 {
            case 357:
                for fi6 := 0; fi6 < 8; fi6++ {
                    v5 -= v5
                    v5 = 1
                    _ = fi6
                }
            case -1:
                v7 := GenS1{F0: ((0.1 != 1.5) && (-32768 < 32767)), F1: 940, F2: -619}
                _ = v7
            case 128:
                if true {
                    fmt.Println(v5)
                }
            default:
                fmt.Println(v5)
            }
        }
        fmt.Println(v5)
        v8 := []int{-670, 32768, 0, -566}
        _ = v8
    } else {
        var v9 float64 = 0.0
        v9 = (v9 + 0.1)
        _ = v9
    }
    v10 := []int{-413}
    var v11 uint32 = 2
    _ = v10
    _ = v11
    _ = v5
    return -12
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    switch "foo bar" {
    case "foo bar":
        switch -2147483648 {
        case 773:
            v1 += 1
        case 9223372036854775807:
            v1 += (-(v1))
        case 42:
            v3[-456] = (!true)
        default:
            fmt.Println(v1)
        }
    case "hello":
        var v12 uint8 = 2
        _ = v12
    case "\n\t":
        delete(v2, -9223372036854775808)
    default:
        if true {
            for fi13 := 0; fi13 < 10; fi13++ {
                if true {
                    var v14 uint32 = 742
                    _ = v14
                }
                fmt.Println(v1)
                _ = fi13
            }
            v15 := GenS1{F0: (42 == 42), F1: 255, F2: 32768}
            _ = v15
        } else {
            v1 += 1
            for fi16 := 0; fi16 < 7; fi16++ {
                if (!(!((false && true) || false))) {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        v1 -= 42
                        _ = fi17
                    }
                    var v18 uint = 735
                    _ = v18
                }
                for fi19 := 0; fi19 < 17; fi19++ {
                    v1 += v1
                    var v20 float32 = ((0.0 + (((1.5 / -2.25) - -2.25) - (0.0 - 1e10))) * float32((float64(0.0) / 3.14159)))
                    _ = fi19
                    _ = v20
                }
                v21 := []int{(fi16 - fi16), (fi16 << ((65535) & 63)), fi16}
                _ = fi16
                _ = v21
            }
            v22 := GenS2{F0: 618, F1: 0, F2: 85}
            _ = v22
        }
    }
    fmt.Println(v1)
    v1 -= (v1 - v1)
    v23 := GenS3{F0: 456, F1: 65535, F2: (0.1 >= -2.25)}
    var v24 uint16 = 42
    switch "\n\t" {
    case "\n\t":
        switch 42 {
        case -2147483648:
            var v25 bool = true
            _ = v25
        case 906:
            v24 = 42
        case 1:
            v26 := []int{1000, -539, -2147483648, -428}
            _ = v26
        }
    case "a":
        var v27 string = ""
        _ = v27
    default:
        v28 := GenS2{F0: 477, F1: 1, F2: 42}
        _ = v28
    }
    var v29 float32 = (3.14159 - -2.25)
    if (!false) {
        switch "foo bar" {
        case "foo bar":
            if (!false) {
                v24 = 1
                fmt.Println(v24, v29, v1)
                v29 = v29
            } else {
                switch 362 {
                case 4294967295:
                    for fi30 := 0; fi30 < 20; fi30++ {
                        for fi31 := 0; fi31 < 10; fi31++ {
                            v32 := GenS3{F0: -620, F1: v24, F2: true}
                            v33 := []int{(fi31 / ((32767) | 1))}
                            var v34 int32 = -2147483648
                            _ = fi31
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        _ = fi30
                    }
                case 42:
                    var v35 uint = 195
                    _ = v35
                case 338:
                    var v36 int = 128
                    _ = v36
                }
                if (!(0 <= 42)) {
                    var v37 uint8 = 127
                    _ = v37
                }
            }
        case "a":
            var v38 int8 = genFn0()
            _ = v38
        }
    }
    fmt.Println(v1, v29, v1)
    v1 -= (v1 / ((v1) | 1))
    v39 := []int{9223372036854775807, 32768, 591}
    if false {
        v29 += v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v39
}
