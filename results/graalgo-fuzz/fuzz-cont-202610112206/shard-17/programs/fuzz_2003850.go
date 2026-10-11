package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := []uint32{42, 2, 4294967295, 65535}
    var v2 string = "foo bar"
    v3 := []int{800, -9223372036854775808, -591}
    var v4 uint = 42
    v5 := make(map[string]string)
    var v6 float64 = 0.0
    for fi7 := 0; fi7 < 15; fi7++ {
        v4 = (v4 | (^(v4)))
        _ = len(v1)
        var v8 int64 = 74
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        if true {
            switch ("\n\t" + v2) {
            case "":
                v3 = append(v3, fi9)
            }
            _ = len(v3)
        } else {
            fmt.Println(fi9, v2, v4)
            v10 := GenS1{F0: 42, F1: (v4 & (v4 % (((v4 % ((v4) | 1))) | 1)))}
            _ = v10
        }
        v11 := GenS1{F0: 984, F1: (v4 - (v4 | v4))}
        _ = fi9
        _ = v11
    }
    var v12 uint32 = 2
    v6 = (v6 - v6)
    v13 := v5[(v2 + v2)]
    v5[v13] = ("hello" + v2)
    v14 := GenS1{F0: 42, F1: v4}
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 uint16 = 255
        var v17 string = v2
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: (v4 != 2), F1: 198}
    if len(v3) > 0 {
        v3[0] = 964
    }
    fmt.Println(v13, v12)
    if true {
        var v19 int8 = 127
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 bool = ((fi20 / (((fi20 + 42)) | 1)) >= (fi20 + 0))
            fmt.Println(v21, v21, v2)
            _ = fi20
            _ = v21
        }
        var v22 string = v2
        _ = v19
        _ = v22
    } else {
        fmt.Println(v13)
    }
    var v23 bool = (true || (!(-744 >= 9223372036854775807)))
    switch 127 {
    case 42:
        v23 = v23
    default:
        _ = len(v3)
    }
    var v24 uint = v4
    if false {
        var v25 int = -346
        for fi26 := 0; fi26 < 10; fi26++ {
            switch (fi26 * fi26) {
            case -535:
                for fi27 := 0; fi27 < 15; fi27++ {
                    if false {
                        v28 := v5[("" + "\n\t")]
                        _ = v28
                    }
                    if false {
                        v29 := GenS3{F0: -9223372036854775808}
                        v24 = (v4 ^ 983)
                        _ = v29
                    } else {
                        switch 4294967295 {
                        case 128:
                            var v30 uint8 = 1
                            _ = v30
                        default:
                            var v31 uint16 = 1
                            _ = v31
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v32 int = fi27
                        v3 = append(v3, -340)
                        _ = v32
                    }
                    _ = fi27
                }
            case 0:
                if (-741 != -9223372036854775808) {
                    v6 -= v6
                    v6 = float64(0.0)
                    v24 = (^(v4))
                } else {
                    fmt.Println(v4, v4, v12)
                    v13 = (v2 + "foo bar")
                    var v33 bool = v23
                    _ = v33
                }
            case 1:
                var v34 int16 = 1
                _ = v34
            }
            var v35 string = (v13 + "")
            if v23 {
                for fi36 := 0; fi36 < 5; fi36++ {
                    fmt.Println(v24, fi36, v24)
                    _ = fi36
                }
            } else {
                v37 := GenS4{F0: -202}
                var v38 uint = v4
                _ = v37
                _ = v38
            }
            _ = fi26
            _ = v35
        }
        switch v25 {
        case 42:
            v39 := GenS4{F0: -1}
            _ = v39
        case 127:
            var v40 int16 = -184
            _ = v40
        default:
            v12 += v12
        }
        _ = v25
    } else {
        var v41 float64 = float64(float32(v6))
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
