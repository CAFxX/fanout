package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int16 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint16, p1 uint, p2 uint8) int32 {
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int32 = -2147483648
        var v8 string = (("hello" + "\n\t") + ("\n\t" + ""))
        fmt.Println(p1)
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        if (771 <= 42) {
            fmt.Println(p1, fi9, p2)
            p2 = p2
        }
        v10 := GenS1{F0: p1, F1: 4294967295, F2: 1}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    var v1 int = 128
    v2 := make(map[uint32]bool)
    var v3 int32 = 1
    v4 := make(map[string]int)
    var v5 int8 = 0
    var v11 int32 = 159
    v5 = v5
    if true {
        if (1 >= 1) {
            v12 := GenS2{F0: (!(32767 > 1)), F1: 1000}
            _ = v12
        } else {
            if (false || false) {
                switch ("\n\t" + "a") {
                case "foo bar":
                    var v13 int8 = v5
                    _ = v13
                case "\n\t":
                    if (3.14159 > 1.5) {
                        v14 := []int{127}
                        v3 += genFn0(42, 1, 2)
                        _ = v14
                    } else {
                        switch 128 {
                        case 351:
                            v15 := []int{}
                            _ = v15
                        }
                        v16 := GenS3{F0: 402, F1: 1000}
                        var v17 uint8 = 128
                        _ = v16
                        _ = v17
                    }
                default:
                    var v18 float64 = float64(0.0)
                    _ = v18
                }
            } else {
                fmt.Println(v11, v1, v1)
            }
        }
    } else {
        v1 -= v1
        v1 -= (-(v1))
        v19 := []int{v1}
        _ = v19
    }
    var v20 int16 = 32767
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 int32 = (v11 - (v3 + -1))
        v23 := GenS2{F0: true, F1: v20}
        if false {
            var v24 int32 = genFn0(1, 0, 127)
            switch (fi21 & v1) {
            case 9223372036854775807:
                var v25 int = -749
                _ = v25
            }
            _ = v24
        } else {
            for fi26 := 0; fi26 < 20; fi26++ {
                v27 := GenS1{F0: 4294967295, F1: 42, F2: v20}
                if ("" <= ("hello" + ("a" + "a"))) {
                    v3 += (-(v3))
                    if (true && (("foo bar" >= ("\n\t" + ("" + "a"))) || (395 > 0))) {
                        var v28 int = 0
                        v1 += 4294967295
                        _ = v28
                    }
                    if ((((fi26 >= (-(fi21))) && ((0 <= 42) || (91 < 18446744073709551615))) && false) || false) {
                        v29 := v4["foo bar"]
                        var v30 float32 = (3.14159 * float32(float64((1e10 - 0.1))))
                        if true {
                            var v31 uint8 = 255
                            _ = len(v2)
                            _ = v31
                        }
                        _ = v29
                        _ = v30
                    } else {
                        var v32 int = (v1 / (((fi26 / ((fi26) | 1))) | 1))
                        v22 = (v3 | (v22 << ((42) & 63)))
                        _ = v32
                    }
                }
                var v33 int8 = (-(v5))
                switch (v1 << ((2) & 63)) {
                case 2147483647:
                    v2[271] = true
                default:
                    v34 := GenS3{F0: 1, F1: -1}
                    _ = v34
                }
                fmt.Println(v33, fi21, v3)
                fmt.Println(v3, v33, v3)
                var v35 uint = 255
                _ = fi26
                _ = v27
                _ = v33
                _ = v35
            }
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    var v36 float32 = (-2.25 - (-2.25 / 3.14159))
    var v37 int16 = (v20 << ((42) & 63))
    fmt.Println(v1)
    for fi38 := 0; fi38 < 15; fi38++ {
        if (false && false) {
            var v39 float32 = v36
            _ = v39
        } else {
            var v40 uint64 = 901
            if false {
                v41 := []int{166, (v1 ^ (^(v1))), -1, (v1 * fi38)}
                _ = v41
            }
            _ = v40
        }
        _ = fi38
    }
    v3 += (v3 / ((v3) | 1))
    v1 = v1
    var v42 float64 = 0.1
    var v43 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
