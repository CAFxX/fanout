package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 bool; F1 uint8; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 bool; F1 uint; F2 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32 }
func genFn0() int {
    v4 := []int{32767, -229, 147}
    var v5 string = ("a" + (("a" + "foo bar") + ""))
    v6 := GenS1{F0: 255, F1: -9223372036854775808}
    for fi7 := 0; fi7 < 12; fi7++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v8 := GenS2{F0: ((^(fi7)) <= (^(fi7))), F1: 127, F2: -2147483648}
        _ = fi7
        _ = v8
    }
    _ = v4
    _ = v5
    _ = v6
    return -496
}

func main() {
    v1 := []int{801, 278, 1}
    var v2 int64 = -410
    var v3 bool = (!(true || true))
    _ = len(v1)
    v3 = (1e10 != (-2.25 * -2.25))
    if (!true) {
        if ("\n\t" > "foo bar") {
            var v9 string = ("\n\t" + (("hello" + "\n\t") + "hello"))
            var v10 int16 = 42
            _ = v10
            _ = v9
        }
        v11 := GenS3{F0: (v2 * v2)}
        for fi12 := 0; fi12 < 9; fi12++ {
            for fi13 := 0; fi13 < 2; fi13++ {
                if (true || (!(float32(1.5) != (0.1 / 0.1)))) {
                    v1 = append(v1, (fi12 & (fi12 % ((genFn0()) | 1))))
                    var v14 string = ("\n\t" + "a")
                    fmt.Println(fi12, fi12)
                    var v15 uint = 0
                    _ = v14
                    _ = v15
                }
                var v16 uint64 = 18446744073709551615
                _ = fi13
                _ = v16
            }
            v2 -= (^(v2))
            fmt.Println(fi12, v3)
            _ = fi12
        }
        _ = v11
    } else {
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS4{F0: 217, F1: 32767}
            var v19 uint32 = 42
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    fmt.Println(v3)
    var v20 uint16 = 32767
    switch 4294967295 {
    case -850:
        var v21 float32 = -2.25
        _ = v21
    case -1:
        if true {
            var v22 int32 = -2147483648
            _ = v22
        } else {
            v23 := GenS5{F0: v3, F1: 2, F2: -1}
            _ = v23
        }
    case 1:
        v24 := GenS3{F0: v2}
        _ = v24
    default:
        if v3 {
            v25 := GenS2{F0: (255 > 1), F1: 255, F2: -2147483648}
            if true {
                var v26 uint16 = (v20 ^ v20)
                _ = v26
            } else {
                switch "\n\t" {
                case "\n\t":
                    var v27 int32 = 42
                    _ = v27
                }
            }
            _ = v25
        }
    }
    v20 = v20
    if v3 {
        v2 = (v2 & -1)
        if (v3 && v3) {
            var v28 uint16 = v20
            v28 = (v20 >> ((0) & 63))
            for fi29 := 0; fi29 < 15; fi29++ {
                for fi30 := 0; fi30 < 1; fi30++ {
                    v31 := GenS1{F0: v28, F1: v2}
                    v32 := GenS6{F0: v20}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                if len(v1) > 0 {
                    v1[0] = (fi29 / ((genFn0()) | 1))
                }
                v33 := GenS4{F0: 175, F1: v28}
                _ = fi29
                _ = v33
            }
            _ = v28
        } else {
            var v34 int16 = 1
            v34 = v34
            fmt.Println(v20, v34, v3)
            _ = v34
        }
        v35 := GenS7{F0: 102}
        _ = v35
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
}
