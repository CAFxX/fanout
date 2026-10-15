package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int8 = -128
    var v2 int32 = 505
    var v3 int16 = 42
    var v4 int32 = 1
    var v5 uint64 = 4294967295
    for fi6 := 0; fi6 < 13; fi6++ {
        v3 = (v3 - v3)
        v2 += (-(v4))
        _ = fi6
    }
    fmt.Println(v3)
    v7 := []int{127, 9223372036854775807}
    var v8 uint32 = 65535
    var v9 int32 = v4
    if (65535 == 1) {
        var v10 float32 = 0.0
        _ = v10
    } else {
        if ((^(v8)) < v8) {
            var v11 uint16 = 0
            if true {
                var v12 uint8 = uint8(v11)
                _ = v12
            }
            _ = v11
        }
        var v13 string = (("foo bar" + (("" + "foo bar") + "foo bar")) + "hello")
        _ = v13
    }
    fmt.Println(v9)
    switch 42 {
    case -579:
        if len(v7) > 0 {
            v7[0] = -45
        }
    case 0:
        if true {
            var v14 uint16 = 2
            var v15 float64 = -2.25
            for fi16 := 0; fi16 < 12; fi16++ {
                v9 += (v2 ^ 42)
                fmt.Println(v8, v8)
                var v17 int16 = 1
                _ = fi16
                _ = v17
            }
            _ = v14
            _ = v15
        } else {
            v18 := GenS1{F0: -9223372036854775808}
            var v19 int32 = (v9 * v2)
            v20 := GenS1{F0: -487}
            _ = v18
            _ = v19
            _ = v20
        }
    }
    v7 = append(v7, 1)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(v9, v9, v3)
    if (false || ((v1 + v1) > 127)) {
        if false {
            if (true || (0.0 < 0.0)) {
                var v21 int64 = 0
                for fi22 := 0; fi22 < 5; fi22++ {
                    v3 -= v3
                    v3 = (v3 >> ((4294967295) & 63))
                    var v23 int64 = -9223372036854775808
                    _ = fi22
                    _ = v23
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v21
            } else {
                v4 = v4
                for fi24 := 0; fi24 < 16; fi24++ {
                    fmt.Println(v2, fi24, fi24)
                    if len(v7) > 0 {
                        v7[0] = fi24
                    }
                    for fi25 := 0; fi25 < 2; fi25++ {
                        v3 -= v3
                        v7 = append(v7, fi24)
                        var v26 float32 = -2.25
                        _ = fi25
                        _ = v26
                    }
                    _ = fi24
                }
                var v27 uint8 = 87
                _ = v27
            }
            v4 -= (v2 | 730)
        } else {
            v1 -= v1
            var v28 int = 9223372036854775807
            _ = v28
        }
    }
    switch 32767 {
    case -2147483648:
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 uint8 = uint8(v8)
            _ = fi29
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
