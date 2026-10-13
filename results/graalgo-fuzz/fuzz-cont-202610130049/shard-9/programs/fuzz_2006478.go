package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint8) int16 {
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint64 = 0
        v8 := GenS1{F0: (fi6 ^ (fi6 / (((fi6 % (((fi6 >> ((4294967295) & 63))) | 1))) | 1)))}
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 = (^(p0))
    var v9 uint8 = p0
    _ = p0
    _ = v9
    return 1000
}
func genFn1(p0 int16) int32 {
    v10 := []int{662, 42}
    fmt.Println(p0)
    _ = p0
    _ = v10
    return 377
}

func main() {
    var v1 int8 = -128
    var v2 uint = 106
    var v3 int64 = 42
    v4 := []string{((("" + "\n\t") + "foo bar") + "hello"), "hello"}
    var v5 uint8 = 127
    v1 = -1
    v11 := GenS2{F0: 0}
    var v12 float64 = 1.5
    v5 += v5
    switch "\n\t" {
    case "\n\t":
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 string = (("hello" + "") + ("" + "a"))
            var v15 float64 = ((float64(-2.25) / (v12 + (v12 / float64(0.1)))) - (v12 / (v12 / ((0.1 + v12) * v12))))
            _ = fi13
            _ = v14
            _ = v15
        }
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v1 += (v1 | -10)
    v4 = append(v4, "")
    v5 = v5
    v4 = append(v4, "a")
    v2 += 4294967295
    v1 -= 127
    v5 = 1
    var v16 uint16 = 65535
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS2{F0: int32(v3)}
        _ = fi17
        _ = v18
    }
    v16 += (v16 % (((v16 * v16)) | 1))
    if (127 < 2147483647) {
        var v19 int64 = 416
        switch "" {
        case "\n\t":
            var v20 float64 = ((v12 / ((v12 / float64((1e10 - ((1e10 / 1e10) + 1e10)))) / v12)) / (v12 - 3.14159))
            _ = v20
        case "foo bar":
            v21 := GenS2{F0: 720}
            _ = v21
        case "hello":
            fmt.Println(v3, v1)
        default:
            if true {
                var v22 bool = false
                _ = v22
            }
        }
        v3 -= v19
        _ = v19
    } else {
        var v23 int64 = (v3 + (v3 | (v3 - -9223372036854775808)))
        switch 1 {
        case 4294967295:
            v4 = append(v4, "hello")
        }
        _ = v23
    }
    if (!((^(v5)) != v5)) {
        fmt.Println(v5, v1, v1)
        for fi24 := 0; fi24 < 16; fi24++ {
            _ = len(v4)
            if (!(0 >= -517)) {
                v25 := GenS2{F0: 2147483647}
                v16 += v16
                for fi26 := 0; fi26 < 17; fi26++ {
                    var v27 float64 = (1.5 + ((float64((float32(v12) / 1.0)) * (0.0 * (v12 * v12))) - 3.14159))
                    if (true || true) {
                        v28 := GenS2{F0: genFn1(42)}
                        var v29 uint8 = 42
                        var v30 float32 = -2.25
                        if (v16 != (v16 << (((v2 % (((v2 ^ 18446744073709551615)) | 1))) & 63))) {
                            v5 = (v29 ^ (v29 & v5))
                            v12 -= v27
                            if len(v4) > 0 {
                                v4[0] = "\n\t"
                            }
                        }
                        _ = v28
                        _ = v29
                        _ = v30
                    } else {
                        var v31 uint16 = v16
                        v27 = v12
                        _ = v31
                    }
                    _ = fi26
                    _ = v27
                }
                _ = v25
            } else {
                var v32 bool = false
                if ((v32 || false) || v32) {
                    var v33 int = (fi24 | -9223372036854775808)
                    v33 -= fi24
                    var v34 string = ("" + "foo bar")
                    _ = v33
                    _ = v34
                } else {
                    _ = len(v4)
                }
                _ = v32
            }
            _ = fi24
        }
    } else {
        var v35 uint8 = (^(v5))
        for fi36 := 0; fi36 < 19; fi36++ {
            for fi37 := 0; fi37 < 14; fi37++ {
                var v38 uint64 = 403
                v38 += 1
                switch int(v3) {
                case 9223372036854775807:
                    if len(v4) > 0 {
                        v4[0] = ("\n\t" + "foo bar")
                    }
                }
                v3 -= (v3 << ((334) & 63))
                if true {
                    v35 = uint8(v38)
                    v4 = append(v4, "hello")
                    v39 := GenS1{F0: fi36}
                    _ = v39
                }
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
        _ = v35
    }
    switch 9223372036854775807 {
    case 1:
        v4 = append(v4, ("foo bar" + "hello"))
    case 636:
        v40 := GenS3{F0: 56}
        _ = v40
    case 127:
        fmt.Println(v2, v12, v16)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
