package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64; F1 int8; F2 uint }
func genFn0(p0 int, p1 int16) int8 {
    var v6 int16 = p1
    var v7 int16 = 32767
    var v8 int64 = 42
    var v9 float64 = 3.14159
    v10 := GenS1{F0: 487, F1: 924}
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{0}
    v2 := []uint32{2}
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]string)
    switch -2147483648 {
    case -1:
        var v11 float32 = (-2.25 + float32(((0.0 * 3.14159) - (0.0 + 3.14159))))
        _ = v11
    case 32767:
        var v12 float64 = 1e10
        _ = v12
    case -9223372036854775808:
        var v13 uint8 = 255
        _ = v13
    default:
        var v14 uint16 = 255
        _ = v14
    }
    switch -993 {
    case 1:
        var v15 uint = 868
        _ = v15
    }
    var v16 int8 = 1
    var v17 int64 = 278
    v16 = -100
    for fi18 := 0; fi18 < 3; fi18++ {
        v16 += (v16 ^ genFn0((fi18 - (^(fi18))), 643))
        if (!(!true)) {
            v17 = v17
        }
        var v19 int32 = -637
        _ = fi18
        _ = v19
    }
    v16 = v16
    v2 = append(v2, 691)
    switch ("\n\t" + "hello") {
    case "hello":
        if (((2147483648 < 65535) && false) && (32767 < 42)) {
            var v20 float32 = (0.1 / 1.0)
            _ = v20
        }
    case "foo bar":
        if false {
            v17 += (v17 & v17)
            var v21 uint = 856
            switch 9223372036854775807 {
            case -9223372036854775808:
                v22 := GenS2{F0: (v17 * v17), F1: genFn0(int(v17), 1000), F2: v21}
                _ = v22
            case 9223372036854775807:
                var v23 int8 = 127
                _ = v23
            case 0:
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 uint16 = 255
                    var v26 float32 = 1e10
                    var v27 uint = v21
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
            default:
                if (false || false) {
                    if (true || true) {
                        if ((0.0 > float64(1.5)) && (-266 >= 760)) {
                            switch (("a" + ("" + ("hello" + ""))) + (("a" + (((("\n\t" + "\n\t") + "a") + ("foo bar" + "")) + ("foo bar" + "a"))) + "")) {
                            case "hello":
                                if (!(true || false)) {
                                    v28 := GenS1{F0: 462, F1: 1}
                                    _ = v28
                                }
                            default:
                                if true {
                                    var v29 int = -9223372036854775808
                                    var v30 uint64 = 4294967295
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v31 bool = ((32767 <= 65535) || (!false))
                                    _ = v31
                                }
                            }
                        }
                    }
                }
            }
            _ = v21
        }
    default:
        v17 += (^(v17))
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        var v33 float64 = 3.14159
        var v34 int8 = (^(v16))
        switch "\n\t" {
        case "hello":
            for fi35 := 0; fi35 < 17; fi35++ {
                switch -2147483648 {
                case -1:
                    var v36 float64 = float64(-2.25)
                    _ = v36
                }
                _ = fi35
            }
        case "foo bar":
            var v37 int8 = v16
            _ = v37
        case "\n\t":
            _ = len(v2)
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    var v38 uint32 = 609
    fmt.Println(v16, v17, v17)
    v39 := GenS2{F0: (-(v17)), F1: 127, F2: 665}
    v17 -= v17
    for fi40 := 0; fi40 < 11; fi40++ {
        var v41 bool = false
        v16 -= (v16 | (v16 + (^(v16))))
        var v42 uint64 = 2
        _ = fi40
        _ = v41
        _ = v42
    }
    v17 -= v17
    for fi43 := 0; fi43 < 15; fi43++ {
        if false {
            v16 = (-(v16))
        }
        _ = fi43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
