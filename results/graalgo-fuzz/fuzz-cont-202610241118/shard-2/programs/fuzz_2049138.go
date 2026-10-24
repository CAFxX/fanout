package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int8; F1 int8 }
func genFn0() int {
    if true {
        if ((!true) || (4294967295 != 4294967295)) {
            switch 128 {
            case 766:
                var v7 int32 = 42
                _ = v7
            }
        } else {
            var v8 int = 9223372036854775807
            v9 := []int{}
            _ = v8
            _ = v9
        }
    } else {
        var v10 int16 = 42
        v10 = 466
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS1{F0: 255, F1: ((228 >= 0) || (float64(1e10) >= -2.25))}
        if (!true) {
            var v13 int8 = 42
            var v14 uint32 = 847
            v13 -= (^(v13))
            _ = v13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    v15 := GenS1{F0: 255, F1: false}
    v16 := []int{-827}
    _ = v15
    _ = v16
    return -373
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 42
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    v6 := make(map[int]uint32)
    _ = len(v4)
    switch 32767 {
    case -230:
        var v17 uint16 = 2
        _ = v17
    case 0:
        var v18 bool = true
        _ = v18
    case 42:
        switch "a" {
        case "":
            if (919 <= 1) {
                fmt.Println(v2)
                v19 := GenS2{F0: 972, F1: 2, F2: (v2 ^ v2)}
                if false {
                    v20 := []int{128, -2147483648}
                    var v21 uint8 = v2
                    if (false || true) {
                        v2 -= (v2 ^ (v21 % (((^(v2))) | 1)))
                        var v22 uint16 = 45
                        _ = len(v5)
                        _ = v22
                    }
                    _ = v20
                    _ = v21
                }
                _ = v19
            }
        }
    default:
        v2 += (v2 << ((4294967295) & 63))
    }
    v2 = (^(v2))
    v2 = v2
    v23 := []int{2147483647, genFn0()}
    for fi24 := 0; fi24 < 3; fi24++ {
        switch "a" {
        case "foo bar":
            v2 = 127
        }
        _ = fi24
    }
    v2 += v2
    var v25 uint32 = 2147483648
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 int8 = -128
        switch ("foo bar" + "") {
        case "":
            switch ("hello" + ("" + ("foo bar" + (("" + (("a" + "") + "foo bar")) + "")))) {
            case "a":
                var v28 int32 = 1
                _ = v28
            case "\n\t":
                for fi29 := 0; fi29 < 15; fi29++ {
                    fmt.Println(v2, v2)
                    _ = fi29
                }
            case "foo bar":
                for fi30 := 0; fi30 < 15; fi30++ {
                    v6[9223372036854775807] = v25
                    delete(v4, "\n\t")
                    var v31 int64 = 42
                    var v32 string = "a"
                    v27 = v27
                    v31 -= v31
                    _ = len(v3)
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
        case "hello":
            v33 := GenS3{F0: (v27 + (^(v27))), F1: (v27 << ((2) & 63))}
            _ = v33
        case "foo bar":
            if true {
                fmt.Println(v27, v25)
            }
        }
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
