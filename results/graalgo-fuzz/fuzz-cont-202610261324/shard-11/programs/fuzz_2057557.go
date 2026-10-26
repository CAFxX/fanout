package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint }
type GenS3 struct { F0 int64; F1 uint8 }
func genFn0() uint64 {
    var v7 int16 = -32768
    var v8 uint16 = 65535
    v9 := GenS1{F0: 0, F1: 2147483648, F2: 326}
    var v10 uint16 = (^(v8))
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 846
}
func genFn1(p0 int64, p1 int32, p2 uint64) int32 {
    p1 = 1
    p1 = (p1 << ((659) & 63))
    var v11 float32 = 3.14159
    v12 := GenS1{F0: 128, F1: 0, F2: 65535}
    var v13 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return p1
}

func main() {
    v1 := []int{2147483647, 9223372036854775807, 4294967295, 128}
    var v2 uint16 = 42
    v3 := make(map[string]bool)
    var v4 float32 = (-2.25 * 0.1)
    var v5 uint = 2
    var v6 uint16 = v2
    var v14 uint64 = uint64(v2)
    fmt.Println(v2, v14, v6)
    var v15 int = 4294967295
    if (128 > 1) {
        if true {
            switch "\n\t" {
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    } else {
        switch (("hello" + ("\n\t" + "\n\t")) + "\n\t") {
        case "":
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 uint8 = 127
                v6 += 2
                _ = fi16
                _ = v17
            }
        }
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 int64 = -9223372036854775808
        fmt.Println(v4)
        fmt.Println(v19, v6, v14)
        _ = fi18
        _ = v19
    }
    if len(v1) > 0 {
        v1[0] = v15
    }
    if (!(true && true)) {
        v20 := GenS1{F0: int(v2), F1: 2147483647, F2: v2}
        for fi21 := 0; fi21 < 4; fi21++ {
            v1 = append(v1, v15)
            var v22 int32 = 0
            _ = fi21
            _ = v22
        }
        v23 := GenS2{F0: (v14 | (v14 + (v14 % ((genFn0()) | 1)))), F1: -32768, F2: (v5 & v5)}
        _ = v20
        _ = v23
    } else {
        v6 = (v6 + v6)
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        v4 = 3.14159
        v25 := GenS1{F0: fi24, F1: 2147483647, F2: (v6 >> ((65535) & 63))}
        _ = fi24
        _ = v25
    }
    _ = len(v3)
    switch v15 {
    case 476:
        for fi26 := 0; fi26 < 15; fi26++ {
            if (((1.5 <= 3.14159) && (true && (0 != v5))) || (42 <= -206)) {
                var v27 bool = true
                v6 += v2
                switch ("foo bar" + "") {
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "":
                    if false {
                        var v28 int8 = -1
                        _ = v28
                    } else {
                        v5 += (v5 & v5)
                        v14 = 912
                        var v29 int8 = 127
                        _ = v29
                    }
                default:
                    for fi30 := 0; fi30 < 18; fi30++ {
                        if len(v1) > 0 {
                            v1[0] = (fi26 + 896)
                        }
                        _ = fi30
                    }
                }
                _ = v27
            } else {
                var v31 float64 = 1.5
                _ = v31
            }
            if (!true) {
                v32 := GenS2{F0: v14, F1: -857, F2: 65535}
                switch "foo bar" {
                case "":
                    var v33 bool = (true && (true || (!false)))
                    _ = v33
                case "hello":
                    if (42 >= 462) {
                        var v34 int16 = 1
                        fmt.Println(v4, v34, v14)
                        fmt.Println(v15, v14, v2)
                        _ = v34
                    }
                case "\n\t":
                    v35 := v3[""]
                    _ = v35
                }
                var v36 bool = (42 == -28)
                _ = v32
                _ = v36
            } else {
                for fi37 := 0; fi37 < 4; fi37++ {
                    v6 = 475
                    _ = fi37
                }
                for fi38 := 0; fi38 < 12; fi38++ {
                    v39 := v3[("hello" + "a")]
                    _ = fi38
                    _ = v39
                }
            }
            _ = fi26
        }
    case 0:
        for fi40 := 0; fi40 < 18; fi40++ {
            for fi41 := 0; fi41 < 14; fi41++ {
                var v42 uint64 = (v14 | v14)
                _ = fi41
                _ = v42
            }
            v2 += (v6 & v6)
            v4 -= v4
            _ = fi40
        }
    }
    v14 += v14
    for fi43 := 0; fi43 < 18; fi43++ {
        v5 -= v5
        var v44 string = ("a" + "")
        var v45 uint64 = genFn0()
        _ = fi43
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 19; fi46++ {
        v14 = (v14 / (((v14 * genFn0())) | 1))
        _ = fi46
    }
    v15 = v15
    if (805 > 10) {
        fmt.Println(v5, v6, v15)
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v47 := GenS3{F0: -9223372036854775808, F1: 42}
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
