package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int16 }
type GenS2 struct { F0 int }
func genFn0() uint32 {
    switch "\n\t" {
    case "foo bar":
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 int8 = -1
            v9 := GenS1{F0: 255, F1: true, F2: -1}
            fmt.Println(v8)
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v10 int64 = 3
    var v11 uint64 = 0
    v12 := []int{127}
    _ = v10
    _ = v11
    _ = v12
    return 641
}

func main() {
    var v1 uint16 = 400
    var v2 uint32 = 424
    var v3 int16 = 1
    var v4 int16 = 1
    var v5 uint16 = (v1 / (((v1 << ((42) & 63))) | 1))
    v6 := []uint32{(v2 * (v2 / ((v2) | 1)))}
    var v13 int64 = -9223372036854775808
    fmt.Println(v1, v4)
    var v14 uint32 = v2
    var v15 int8 = 16
    v16 := GenS2{F0: 127}
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 uint32 = 149
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = len(v6)
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: uint16(v2), F1: false, F2: v3}
    v2 = v14
    if true {
        if (42 < int(v5)) {
            v15 += 22
            v3 += v3
        }
    } else {
        v20 := GenS1{F0: v5, F1: (!((false && (3.14159 != -2.25)) && (3.14159 == (1.5 - -2.25)))), F2: 42}
        for fi21 := 0; fi21 < 18; fi21++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi21
        }
        switch ((("hello" + "a") + "foo bar") + (("a" + "a") + (("" + ("\n\t" + "\n\t")) + "a"))) {
        case "\n\t":
            if (true || (false && false)) {
                switch 42 {
                case 4294967295:
                    v2 = genFn0()
                case 9223372036854775807:
                    switch ("a" + "a") {
                    case "hello":
                        v14 += genFn0()
                    case "\n\t":
                        v1 -= v1
                    case "foo bar":
                        if (true || true) {
                            var v22 bool = false
                            v15 -= (v15 * (v15 | v15))
                            _ = v22
                        }
                    }
                case 128:
                    var v23 int32 = 0
                    _ = v23
                }
                v6 = append(v6, genFn0())
            } else {
                if false {
                    if len(v6) > 0 {
                        v6[0] = genFn0()
                    }
                    v13 = 0
                    fmt.Println(v15, v13)
                }
                v1 -= 2
                var v24 float32 = -2.25
                _ = v24
            }
        }
        _ = v20
    }
    var v25 int32 = -2147483648
    v26 := GenS2{F0: 2147483647}
    var v27 uint8 = 127
    var v28 uint32 = (v14 % ((genFn0()) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
