package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint64 }
func genFn0(p0 int16) uint16 {
    p0 = p0
    fmt.Println(p0)
    p0 += p0
    fmt.Println(p0)
    var v5 float32 = (1e10 - (1.5 * 3.14159))
    _ = p0
    _ = v5
    return 827
}

func main() {
    var v1 string = ("foo bar" + ("hello" + ("foo bar" + (("\n\t" + "a") + "foo bar"))))
    var v2 int8 = 127
    var v3 uint32 = 1
    v4 := make(map[uint32]int)
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int8 = (v2 % (((^(v2))) | 1))
        switch v1 {
        case "foo bar":
            var v8 uint16 = 627
            _ = v8
        default:
            var v9 int16 = 389
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint = 255
        _ = fi10
        _ = v11
    }
    var v12 uint8 = 0
    switch 4294967295 {
    case -14:
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 uint = 255
            _ = fi13
            _ = v14
        }
    case 2147483647:
        v3 += (^(v3))
    }
    var v15 uint16 = 1
    fmt.Println(v3, v1)
    var v16 bool = (!true)
    v17 := GenS1{F0: v2, F1: v1, F2: 722}
    var v18 int32 = -1
    switch ("hello" + v1) {
    case "a":
        var v19 uint16 = (^(v15))
        _ = v19
    case "foo bar":
        var v20 float32 = -2.25
        _ = v20
    default:
        v15 -= 65535
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        if v16 {
            if ((v1 + v1) <= (("" + "\n\t") + "foo bar")) {
                var v22 uint16 = v15
                _ = v22
            } else {
                v3 = 2
                var v23 int16 = 1
                _ = v23
            }
            switch (fi21 + fi21) {
            case -2147483648:
                if false {
                    fmt.Println(v18)
                    for fi24 := 0; fi24 < 8; fi24++ {
                        switch (fi21 >> ((25) & 63)) {
                        case -9223372036854775808:
                            v15 += uint16(v12)
                        case 346:
                            switch (-(fi24)) {
                            case 128:
                                switch ((("" + ("" + v1)) + "a") + (v1 + "\n\t")) {
                                case "":
                                    v15 += genFn0(42)
                                default:
                                    if (!v16) {
                                        v12 -= (v12 << ((1) & 63))
                                    } else {
                                        var v25 int8 = -114
                                        var v26 float64 = ((1.5 + 0.0) * (float64(((float32(((1e10 - (3.14159 * 0.1)) + (float64(1e10) * (float64(0.0) + 0.0)))) / -2.25) / (-2.25 / -2.25))) - (0.0 * float64(((-2.25 / -2.25) / 1e10)))))
                                        _ = v25
                                        _ = v26
                                    }
                                }
                                if (2 != (^(v3))) {
                                    var v27 int8 = 0
                                    _ = v27
                                }
                                v16 = (3.14159 < (float64(1.5) / 1.5))
                            case 127:
                                var v28 float64 = 1.5
                                _ = v28
                            default:
                                var v29 float32 = 3.14159
                                _ = v29
                            }
                        }
                        _ = fi24
                    }
                }
            case -9223372036854775808:
                if (!true) {
                    v3 += v3
                }
            case 745:
                switch (((v1 + v1) + "foo bar") + "foo bar") {
                case "hello":
                    v30 := []int{fi21, fi21}
                    _ = v30
                }
            }
            delete(v4, 795)
        }
        _ = fi21
    }
    v12 += (v12 + (v12 << ((1) & 63)))
    if (42 == 794) {
        v31 := v4[(v3 * (^(v3)))]
        if v16 {
            var v32 int64 = 538
            fmt.Println(v16, v12)
            _ = v32
        } else {
            var v33 float64 = 0.1
            _ = v33
        }
        _ = v31
    } else {
        v16 = v16
        v34 := GenS1{F0: v2, F1: v1, F2: 4294967295}
        if true {
            var v35 int16 = -195
            _ = v35
        } else {
            for fi36 := 0; fi36 < 19; fi36++ {
                v16 = (0 > -1)
                for fi37 := 0; fi37 < 16; fi37++ {
                    v3 += uint32(v12)
                    _ = fi37
                }
                _ = fi36
            }
        }
        _ = v34
    }
    v4[v3] = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
