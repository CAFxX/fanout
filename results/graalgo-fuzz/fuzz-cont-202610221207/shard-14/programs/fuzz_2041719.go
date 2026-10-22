package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int16 = 42
    v2 := []int{0, 0, -2147483648}
    var v3 int64 = -9223372036854775808
    v3 = v3
    var v4 int8 = 127
    var v5 string = ("foo bar" + ("foo bar" + "hello"))
    v2 = append(v2, 736)
    v3 -= v3
    if ((1e10 == 1e10) && false) {
        _ = len(v2)
    }
    fmt.Println(v3, v1, v1)
    if true {
        if (-1 < 2147483647) {
            v2 = append(v2, 907)
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v1 += (v1 | v1)
        v2 = append(v2, int(v3))
    } else {
        _ = len(v2)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = len(v2)
    }
    if (223 > -2147483648) {
        if (!true) {
            switch 32767 {
            case 32768:
                v3 = (v3 * v3)
            case 9223372036854775807:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            var v6 uint64 = 0
            switch v5 {
            case "\n\t":
                for fi7 := 0; fi7 < 5; fi7++ {
                    fmt.Println(v1, v1)
                    fmt.Println(v4, v3, v6)
                    if ((fi7 >> ((517) & 63)) <= 32767) {
                        v1 = (v1 | v1)
                        var v8 int64 = v3
                        var v9 int8 = v4
                        _ = v8
                        _ = v9
                    } else {
                        if true {
                            _ = len(v2)
                        }
                        v10 := GenS1{F0: -2147483648}
                        _ = v10
                    }
                    _ = fi7
                }
            case "a":
                var v11 int32 = 0
                _ = v11
            }
            _ = v6
        } else {
            var v12 uint64 = 42
            _ = v12
        }
        if (v1 == (v1 & v1)) {
            v4 += v4
            v1 = v1
            v1 -= (v1 - (v1 & (v1 / ((int16(v3)) | 1))))
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 float32 = (float32((3.14159 / 0.1)) + 3.14159)
            _ = len(v2)
            _ = fi13
            _ = v14
        }
    } else {
        if true {
            var v15 uint32 = 324
            if false {
                var v16 uint16 = 2
                v16 = v16
                _ = v16
            }
            _ = v15
        } else {
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 uint8 = 127
                var v19 int64 = -1
                v20 := GenS1{F0: fi17}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            if (!(42 < 42)) {
                var v21 float64 = 1e10
                switch -9223372036854775808 {
                case -1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 32767:
                    v22 := GenS1{F0: 1}
                    _ = v22
                }
                var v23 uint32 = 500
                _ = v21
                _ = v23
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            if (false && true) {
                fmt.Println(v1)
                v25 := GenS1{F0: fi24}
                v1 += (v1 + 32767)
                _ = v25
            } else {
                fmt.Println(v3, v1, v4)
            }
            if len(v2) > 0 {
                v2[0] = (fi24 * (fi24 ^ 32768))
            }
            if true {
                var v26 int32 = 42
                _ = v26
            }
            _ = fi24
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v27 int32 = 2147483647
    var v28 string = (v5 + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
