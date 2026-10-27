package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int32) string {
    for fi5 := 0; fi5 < 5; fi5++ {
        p0 -= p0
        var v6 int32 = (p0 >> ((4294967295) & 63))
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 0
    if false {
        var v8 string = (("a" + "") + "a")
        _ = v8
    } else {
        v7 = v7
    }
    var v9 uint32 = 4294967295
    fmt.Println(v7, v7, p0)
    _ = p0
    _ = v7
    _ = v9
    return "hello"
}

func main() {
    v1 := []int{9223372036854775807}
    var v2 int16 = -616
    var v3 uint8 = 255
    var v4 int = 9223372036854775807
    for fi10 := 0; fi10 < 13; fi10++ {
        v3 += v3
        var v11 int16 = (v2 * (v2 << ((0) & 63)))
        for fi12 := 0; fi12 < 15; fi12++ {
            fmt.Println(v3, v3, fi12)
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint = 1
        _ = fi13
        _ = v14
    }
    if false {
        var v15 string = (("hello" + "a") + "")
        v16 := GenS1{F0: 0, F1: 496}
        if (((3.14159 - ((float32(0.0) * 3.14159) + (0.0 - 0.1))) + ((3.14159 + 3.14159) * -2.25)) <= (0.0 / 3.14159)) {
            var v17 int16 = v2
            for fi18 := 0; fi18 < 2; fi18++ {
                v17 = v2
                _ = fi18
            }
            switch v15 {
            case "foo bar":
                fmt.Println(v3, v4)
            case "\n\t":
                var v19 int32 = 87
                _ = v19
            default:
                var v20 int32 = -2147483648
                _ = v20
            }
            _ = v17
        } else {
            v3 += v3
            v15 = "foo bar"
            var v21 float64 = 0.0
            _ = v21
        }
        _ = v15
        _ = v16
    } else {
        var v22 int16 = v2
        if false {
            for fi23 := 0; fi23 < 16; fi23++ {
                v3 = v3
                _ = fi23
            }
        }
        _ = v22
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!false) {
        for fi24 := 0; fi24 < 17; fi24++ {
            if (!(-4 <= -1)) {
                if true {
                    v2 = 783
                }
            }
            _ = fi24
        }
        var v25 int16 = 32767
        for fi26 := 0; fi26 < 9; fi26++ {
            switch v4 {
            case 209:
                var v27 int = v4
                _ = v27
            case 32767:
                v4 += 9223372036854775807
            default:
                for fi28 := 0; fi28 < 1; fi28++ {
                    var v29 uint64 = 2
                    if len(v1) > 0 {
                        v1[0] = -101
                    }
                    v30 := GenS1{F0: 9223372036854775807, F1: 42}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
            _ = fi26
        }
        _ = v25
    } else {
        v31 := GenS1{F0: -1, F1: 326}
        _ = v31
    }
    _ = len(v1)
    var v32 uint16 = 2
    var v33 uint8 = (v3 >> ((uint(v32)) & 63))
    for fi34 := 0; fi34 < 3; fi34++ {
        var v35 uint32 = 1
        for fi36 := 0; fi36 < 3; fi36++ {
            var v37 int = (v4 - 128)
            switch (v4 * (v37 / ((1) | 1))) {
            case 128:
                v38 := GenS2{F0: -489}
                _ = v38
            case -333:
                var v39 string = "\n\t"
                _ = v39
            }
            _ = fi36
            _ = v37
        }
        _ = fi34
        _ = v35
    }
    if len(v1) > 0 {
        v1[0] = v4
    }
    v33 += (v3 | 172)
    v32 += (v32 / ((32767) | 1))
    var v40 uint32 = uint32(v33)
    var v41 uint16 = v32
    fmt.Println(v2, v33)
    for fi42 := 0; fi42 < 6; fi42++ {
        _ = len(v1)
        _ = fi42
    }
    v4 = v4
    var v43 int16 = (v2 ^ v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v40
    _ = v41
    _ = v43
}
