package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint64; F1 string; F2 uint }

func main() {
    v1 := []int{42, -1, 1}
    v2 := make(map[string]bool)
    var v3 string = ""
    v4 := []uint32{2147483648}
    v5 := GenS1{F0: 4294967295, F1: "\n\t"}
    var v6 float32 = (0.0 + ((1.5 + 1.5) + 0.0))
    if (127 == 4294967295) {
        switch -2147483648 {
        case 42:
            for fi7 := 0; fi7 < 14; fi7++ {
                v6 += (((v6 - 1.5) / (v6 - v6)) / v6)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if ((false && false) && true) {
                    v1 = append(v1, (fi7 ^ fi7))
                    var v8 float32 = (0.0 + (v6 * (((0.1 + ((v6 * 0.1) / ((3.14159 - (v6 - 1.5)) / v6))) + 0.0) + (v6 + v6))))
                    if true {
                        v3 = v3
                    }
                    _ = v8
                }
                _ = fi7
            }
        case 291:
            if false {
                v2[v3] = false
            } else {
                if (false || false) {
                    var v9 string = v3
                    v10 := GenS1{F0: 2147483648, F1: "hello"}
                    _ = v10
                    _ = v9
                } else {
                    v11 := GenS1{F0: 42, F1: v3}
                    v12 := GenS2{F0: 0, F1: ("hello" + v3), F2: 2}
                    _ = v11
                    _ = v12
                }
            }
        default:
            var v13 uint32 = 2
            _ = v13
        }
        for fi14 := 0; fi14 < 6; fi14++ {
            for fi15 := 0; fi15 < 3; fi15++ {
                v6 += v6
                v3 = (("hello" + ((v3 + "foo bar") + ((v3 + v3) + v3))) + v3)
                var v16 uint16 = 0
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 13; fi17++ {
                if true {
                    switch ((v3 + ((((v3 + "") + "") + ((v3 + v3) + v3)) + (("\n\t" + v3) + "a"))) + v3) {
                    case "":
                        var v18 int = 2147483647
                        _ = v18
                    default:
                        var v19 uint8 = 0
                        _ = v19
                    }
                    if (314 < 65535) {
                        var v20 int = fi17
                        _ = v20
                    } else {
                        v6 = (v6 / v6)
                        v6 -= v6
                        var v21 string = v3
                        _ = v21
                    }
                } else {
                    if ((("\n\t" == v3) && (100 < 0)) && (false && true)) {
                        var v22 int = fi17
                        v6 = v6
                        _ = v22
                    }
                }
                _ = fi17
            }
            v6 -= (float32((3.14159 + (3.14159 - 0.0))) / 1.5)
            _ = fi14
        }
        var v23 float64 = -2.25
        _ = v23
    }
    v3 = v3
    v6 -= 0.1
    fmt.Println(v3)
    var v24 int32 = -2147483648
    for fi25 := 0; fi25 < 18; fi25++ {
        if len(v1) > 0 {
            v1[0] = fi25
        }
        fmt.Println(v3)
        var v26 int64 = -9223372036854775808
        _ = fi25
        _ = v26
    }
    var v27 string = v3
    var v28 int16 = 0
    _ = len(v1)
    fmt.Println(v28, v28)
    for fi29 := 0; fi29 < 14; fi29++ {
        switch ((v3 + "\n\t") + v3) {
        case "":
            for fi30 := 0; fi30 < 11; fi30++ {
                v3 = "a"
                v31 := GenS2{F0: 260, F1: ((((v27 + (((v27 + "\n\t") + "") + v27)) + v3) + (v27 + ("" + v3))) + v3), F2: 2}
                _ = fi30
                _ = v31
            }
        default:
            fmt.Println(v6, v6)
        }
        for fi32 := 0; fi32 < 1; fi32++ {
            var v33 int64 = 1
            v1 = append(v1, fi32)
            _ = fi32
            _ = v33
        }
        _ = fi29
    }
    var v34 int = -2147483648
    v27 = (v3 + "a")
    if ((v34 + -795) >= v34) {
        var v35 int64 = 1
        v24 = (v24 % ((2147483647) | 1))
        for fi36 := 0; fi36 < 19; fi36++ {
            switch (fi36 ^ -29) {
            case 548:
                var v37 uint = 255
                _ = v37
            case 8:
                var v38 int8 = 95
                _ = v38
            default:
                for fi39 := 0; fi39 < 6; fi39++ {
                    delete(v2, (v3 + "\n\t"))
                    var v40 float32 = ((v6 - (1e10 + v6)) / float32(1e10))
                    _ = fi39
                    _ = v40
                }
            }
            var v41 int8 = 0
            v42 := GenS2{F0: 2, F1: v27, F2: 2}
            _ = fi36
            _ = v41
            _ = v42
        }
        _ = v35
    }
    v43 := v2["foo bar"]
    var v44 int64 = 637
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
