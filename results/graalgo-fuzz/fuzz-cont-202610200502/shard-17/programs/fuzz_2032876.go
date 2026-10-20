package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483647, 2147483647, 2}
    var v3 uint8 = 42
    v4 := []int{9223372036854775807, 0}
    v5 := []int{}
    var v6 uint32 = 66
    v6 = (v6 + 1)
    for fi7 := 0; fi7 < 8; fi7++ {
        v1[(fi7 - (-(fi7)))] = fi7
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int32 = -2147483648
            switch ("" + "\n\t") {
            case "hello":
                fmt.Println(v6, v6)
            }
            _ = len(v2)
            _ = fi8
            _ = v9
        }
        var v10 uint = 0
        _ = fi7
        _ = v10
    }
    var v11 bool = false
    v12 := GenS1{F0: -320}
    var v13 uint32 = (v6 + (v6 << ((18446744073709551615) & 63)))
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 int64 = 271
            if (true || v11) {
                switch (fi15 & fi15) {
                case -9223372036854775808:
                    v6 = v6
                case 0:
                    if false {
                        v11 = (-2.25 < 0.1)
                        fmt.Println(v11)
                        fmt.Println(v11, fi14, fi15)
                    } else {
                        if (!v11) {
                            v16 += (v16 & -9223372036854775808)
                            var v17 int64 = v16
                            if v11 {
                                if true {
                                    if v11 {
                                        fmt.Println(v6, v16, fi15)
                                    }
                                }
                                var v18 float64 = 1e10
                                _ = v18
                            }
                            if v11 {
                                fmt.Println(v16, v3, v3)
                                v11 = (((2147483647 <= v6) || v11) && v11)
                                var v19 uint64 = 32
                                _ = v19
                            }
                            var v20 int16 = -242
                            _ = v17
                            _ = v20
                        }
                        fmt.Println(v11, v6)
                    }
                default:
                    if ((!(-58 <= 2147483647)) && ((v16 >> ((42) & 63)) < (v16 & (v16 & (v16 - -1))))) {
                        var v21 uint16 = 919
                        _ = v21
                    } else {
                        v13 = (v6 + v6)
                        var v22 string = ("\n\t" + (("a" + "\n\t") + (((("hello" + "") + "a") + "") + "\n\t")))
                        _ = v22
                    }
                }
                v3 = v3
                delete(v1, (fi15 & (fi15 & fi15)))
            }
            _ = fi15
            _ = v16
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 bool = (!((0 >= (^(v6))) || false))
            fmt.Println(v3)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi23
            _ = v24
        }
        fmt.Println(v6)
        _ = fi14
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        v11 = (1 < 9223372036854775807)
        switch (fi25 >> ((20) & 63)) {
        case 32767:
            v26 := v1[(fi25 % ((fi25) | 1))]
            _ = v26
        case -9223372036854775808:
            var v27 string = ((("\n\t" + "") + "hello") + (("" + "") + "a"))
            _ = v27
        default:
            var v28 uint = 0
            _ = v28
        }
        v13 -= 70
        _ = fi25
    }
    v1[-1] = 9223372036854775807
    var v29 float32 = float32(0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
