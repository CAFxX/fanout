package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint8; F1 bool; F2 int8 }
func genFn0() uint16 {
    v6 := []int{127}
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 378, F1: 1000}
        _ = fi7
        _ = v8
    }
    if (-81 <= 32767) {
        v9 := GenS1{F0: 301, F1: 42}
        _ = v9
    }
    _ = v6
    return 32767
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int32 = 999
    v3 := []string{}
    var v4 bool = (1 <= 42)
    v5 := []uint32{42, 271}
    v2 = (^(v2))
    for fi10 := 0; fi10 < 9; fi10++ {
        v1 = (v1 & v1)
        switch 138 {
        case 625:
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 int16 = int16(v2)
                _ = fi11
                _ = v12
            }
        case -2147483648:
            if true {
                var v13 int16 = -1
                for fi14 := 0; fi14 < 16; fi14++ {
                    v5 = append(v5, 65535)
                    _ = fi14
                }
                if v4 {
                    var v15 int16 = (v13 + 32767)
                    v15 = 0
                    var v16 uint64 = (^(v1))
                    _ = v15
                    _ = v16
                }
                _ = v13
            }
        case 42:
            var v17 uint64 = (^(v1))
            _ = v17
        }
        var v18 float64 = float64(((float32(1.5) - -2.25) - (1.5 + -2.25)))
        _ = fi10
        _ = v18
    }
    v19 := GenS1{F0: 42, F1: 42}
    var v20 uint8 = 127
    var v21 bool = v4
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := GenS2{F0: (v1 / ((v1) | 1)), F1: 771}
        v4 = true
        v21 = v4
        _ = fi22
        _ = v23
    }
    if v4 {
        for fi24 := 0; fi24 < 20; fi24++ {
            for fi25 := 0; fi25 < 13; fi25++ {
                fmt.Println(v1)
                v4 = (4294967295 != 504)
                _ = fi25
            }
            _ = fi24
        }
    }
    v26 := GenS3{F0: v20, F1: (("foo bar" > "hello") || (-1 >= 0)), F2: 1}
    if len(v5) > 0 {
        v5[0] = 639
    }
    if ((v1 ^ v1) == v1) {
        fmt.Println(v21, v21)
        var v27 int32 = 1
        v28 := GenS2{F0: 592, F1: -32768}
        _ = v27
        _ = v28
    }
    var v29 int8 = 42
    if true {
        v29 = (v29 * -15)
    }
    if len(v5) > 0 {
        v5[0] = 0
    }
    if v21 {
        fmt.Println(v29)
        v29 += (v29 * v29)
        if false {
            for fi30 := 0; fi30 < 12; fi30++ {
                for fi31 := 0; fi31 < 19; fi31++ {
                    v3 = append(v3, "hello")
                    _ = fi31
                }
                _ = fi30
            }
            v32 := GenS3{F0: (v20 * (v20 * (v20 * (v20 + (^(v20)))))), F1: (!((((("foo bar" + (("\n\t" + "\n\t") + ("hello" + ((("" + ("foo bar" + "hello")) + "\n\t") + "\n\t")))) + ("hello" + ("" + (("\n\t" + "foo bar") + "\n\t")))) == "hello") || v4) || (v20 != v20))), F2: v29}
            for fi33 := 0; fi33 < 20; fi33++ {
                v20 += v20
                _ = fi33
            }
            _ = v32
        } else {
            fmt.Println(v2, v4)
            if true {
                var v34 string = "\n\t"
                _ = v34
            }
        }
    }
    fmt.Println(v21, v4)
    if false {
        var v35 bool = (v21 || v21)
        _ = v35
    } else {
        for fi36 := 0; fi36 < 7; fi36++ {
            v5 = append(v5, 65535)
            _ = fi36
        }
        fmt.Println(v1, v4, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
