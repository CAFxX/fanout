package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint; F2 int16 }

func main() {
    v1 := []uint32{2147483647}
    var v2 uint16 = 145
    v3 := make(map[uint32]int)
    var v4 uint8 = 255
    var v5 int8 = 0
    var v6 int = -9223372036854775808
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int8 = v5
        _ = fi7
        _ = v8
    }
    fmt.Println(v2, v5)
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    var v9 string = ("a" + ("a" + "\n\t"))
    var v10 uint64 = 212
    if (!((v5 - (v5 + 1)) > v5)) {
        fmt.Println(v5)
        var v11 uint8 = (^(v4))
        switch 0 {
        case 1:
            v12 := v3[65535]
            _ = v12
        case 979:
            if (-846 > -475) {
                if (32767 == 32767) {
                    for fi13 := 0; fi13 < 5; fi13++ {
                        if len(v1) > 0 {
                            v1[0] = 2147483648
                        }
                        switch (v6 & v6) {
                        case 127:
                            for fi14 := 0; fi14 < 6; fi14++ {
                                var v15 uint32 = 2
                                v9 = v9
                                var v16 int32 = 2147483647
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                        }
                        _ = fi13
                    }
                    if (!(0 == -2147483648)) {
                        v10 -= (v10 * v10)
                        v6 = (v6 + v6)
                        v17 := GenS1{F0: 21}
                        _ = v17
                    } else {
                        _ = len(v1)
                        var v18 string = (("hello" + ("\n\t" + v9)) + v9)
                        _ = v18
                    }
                }
                v19 := GenS2{F0: -2147483648}
                _ = v19
            }
        }
        _ = v11
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 uint8 = (v4 | (v4 + v4))
        if (v10 > v10) {
            v10 -= v10
            var v22 uint16 = v2
            for fi23 := 0; fi23 < 1; fi23++ {
                if false {
                    v22 += (v2 << ((369) & 63))
                    if false {
                        var v24 bool = (2147483647 > int32(v5))
                        v24 = (v24 && (1e10 >= (float64(3.14159) * 3.14159)))
                        _ = v24
                    }
                } else {
                    fmt.Println(v4)
                    v1 = append(v1, 42)
                    switch ("foo bar" + v9) {
                    case "":
                        fmt.Println(fi20, v21, v9)
                    case "foo bar":
                        v25 := v3[2147483648]
                        _ = v25
                    case "\n\t":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
                _ = fi23
            }
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    var v26 uint16 = (^(v2))
    fmt.Println(v4)
    var v27 uint16 = 255
    var v28 int8 = v5
    v10 = v10
    v4 += v4
    for fi29 := 0; fi29 < 19; fi29++ {
        for fi30 := 0; fi30 < 4; fi30++ {
            v31 := GenS1{F0: 1}
            var v32 uint64 = v10
            v33 := GenS3{F0: (v32 + v10), F1: 4294967295, F2: -1}
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
        _ = fi29
    }
    v28 = v5
    for fi34 := 0; fi34 < 19; fi34++ {
        fmt.Println(v2, fi34, v9)
        _ = fi34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
