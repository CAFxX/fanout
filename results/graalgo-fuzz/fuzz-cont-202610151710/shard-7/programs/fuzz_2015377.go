package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() uint32 {
    var v6 int64 = 9223372036854775807
    var v7 float32 = ((-2.25 / 1.0) / (1e10 - 1.5))
    _ = v6
    _ = v7
    return 481
}

func main() {
    v1 := []uint32{2147483647}
    var v2 float64 = 0.0
    var v3 uint = 2
    v4 := []int{42, int(v3), 222, 1}
    var v5 int = -9223372036854775808
    v8 := GenS1{F0: "foo bar"}
    v4 = append(v4, 127)
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 float32 = -2.25
        v5 = 9223372036854775807
        v11 := GenS1{F0: ("\n\t" + "")}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int = (-(v5))
    if (577 != 2) {
        switch (v12 | 0) {
        case 128:
            var v13 int16 = 1000
            _ = v13
        }
    } else {
        v12 += v12
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            if true {
                if len(v1) > 0 {
                    v1[0] = 310
                }
            } else {
                v17 := GenS1{F0: ("" + "\n\t")}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v18 uint16 = 2
                var v19 bool = true
                var v20 int64 = -797
                _ = v17
                _ = v18
                _ = v19
                _ = v20
            }
            _ = fi16
        }
        _ = fi14
    }
    if len(v1) > 0 {
        v1[0] = 38
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        if len(v1) > 0 {
            v1[0] = 704
        }
        var v22 float32 = (0.1 / float32(((-2.25 + (v2 * -2.25)) + 3.14159)))
        _ = fi21
        _ = v22
    }
    if true {
        v23 := GenS1{F0: ("a" + ("" + "\n\t"))}
        _ = v23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        _ = len(v1)
        if false {
            for fi25 := 0; fi25 < 20; fi25++ {
                var v26 float64 = 1e10
                if false {
                    if true {
                        fmt.Println(v26, fi24, v2)
                        v27 := GenS1{F0: "hello"}
                        _ = len(v4)
                        fmt.Println(v12, v3)
                        _ = v27
                    }
                    v12 -= v12
                }
                var v28 int16 = 1
                _ = fi25
                _ = v26
                _ = v28
            }
            if false {
                switch (v12 ^ 2147483647) {
                case 4294967295:
                    v2 -= (v2 + 0.1)
                case 32768:
                    for fi29 := 0; fi29 < 13; fi29++ {
                        _ = len(v1)
                        var v30 float64 = ((v2 + 1e10) / 1.5)
                        v2 -= (((((-2.25 + float64(1.5)) - (float64(1.5) + -2.25)) / v30) - v30) * ((1.5 - (float64((((-2.25 * -2.25) * (float32(v30) / 1.0)) - 0.0)) + v30)) / (3.14159 - v30)))
                        _ = fi29
                        _ = v30
                    }
                default:
                    _ = len(v1)
                }
                v2 -= 0.1
            }
            var v31 string = ""
            _ = v31
        }
        v2 = v2
        _ = fi24
    }
    v1 = append(v1, 402)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
