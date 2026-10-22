package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int }
type GenS2 struct { F0 bool; F1 uint; F2 int }
func genFn0() uint32 {
    switch 4294967295 {
    case -2147483648:
        v7 := GenS1{F0: 742, F1: 0, F2: 128}
        _ = v7
    case 42:
        var v8 float64 = 0.1
        _ = v8
    case 517:
        switch -9223372036854775808 {
        case 456:
            var v9 int64 = 9223372036854775807
            _ = v9
        case 32768:
            var v10 string = "hello"
            _ = v10
        default:
            var v11 float32 = 0.1
            _ = v11
        }
    default:
        var v12 string = ("a" + (("hello" + ("foo bar" + "hello")) + "a"))
        _ = v12
    }
    var v13 string = "foo bar"
    var v14 uint16 = 524
    var v15 uint8 = 127
    var v16 float64 = 0.0
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return 2147483648
}

func main() {
    v1 := []uint32{2147483648, 42}
    v2 := make(map[string]uint32)
    v3 := []uint32{2, 305, 2}
    v4 := []string{"foo bar", ("a" + ("a" + "hello")), "a"}
    var v5 int64 = 0
    v6 := []int{585, -445, 127, 32767}
    delete(v2, "foo bar")
    if ((42 <= 255) && true) {
        v17 := GenS1{F0: (-(v5)), F1: -32768, F2: 32767}
        if true {
            if true {
                var v18 uint8 = 0
                var v19 int16 = -964
                _ = v18
                _ = v19
            }
        } else {
            v5 += v5
            var v20 uint64 = 2
            _ = v20
        }
        _ = v17
    }
    var v21 uint64 = 1
    fmt.Println(v5)
    v22 := GenS2{F0: false, F1: 18446744073709551615, F2: 32767}
    for fi23 := 0; fi23 < 7; fi23++ {
        for fi24 := 0; fi24 < 18; fi24++ {
            v5 = (v5 ^ -9223372036854775808)
            _ = fi24
        }
        v25 := GenS1{F0: (v5 % (((v5 ^ v5)) | 1)), F1: 0, F2: (fi23 << ((1) & 63))}
        if true {
            v26 := GenS1{F0: -1, F1: 0, F2: (fi23 / (((^(fi23))) | 1))}
            _ = v26
        } else {
            v21 = 756
        }
        _ = fi23
        _ = v25
    }
    if (42 != 881) {
        v21 = (^(v21))
    }
    var v27 int16 = -1
    for fi28 := 0; fi28 < 8; fi28++ {
        if (!((v27 << ((870) & 63)) >= (v27 >> ((65535) & 63)))) {
            if (!((-2.25 * 1e10) != float64(1.5))) {
                v27 += 441
            }
        }
        _ = fi28
    }
    v27 += v27
    v5 = v5
    _ = len(v4)
    for fi29 := 0; fi29 < 7; fi29++ {
        v21 = v21
        v21 = 0
        _ = fi29
    }
    delete(v2, "foo bar")
    for fi30 := 0; fi30 < 3; fi30++ {
        if true {
            switch -9223372036854775808 {
            case 119:
                if false {
                    var v31 int8 = 100
                    if ((0 != v5) || (!false)) {
                        v21 += (v21 + (^(v21)))
                        var v32 string = "\n\t"
                        v33 := GenS2{F0: true, F1: 1, F2: 9223372036854775807}
                        _ = v32
                        _ = v33
                    }
                    _ = v31
                }
            default:
                var v34 int8 = 100
                _ = v34
            }
            for fi35 := 0; fi35 < 15; fi35++ {
                var v36 uint = 2
                _ = fi35
                _ = v36
            }
        } else {
            v37 := GenS1{F0: (v5 | (v5 ^ 1)), F1: (v27 * v27), F2: 1}
            _ = v37
        }
        v27 = (^(v27))
        for fi38 := 0; fi38 < 11; fi38++ {
            fmt.Println(v21)
            _ = fi38
        }
        _ = fi30
    }
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    fmt.Println(v21, v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
