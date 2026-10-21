package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int64; F2 string }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint8 = 2
    var v2 string = "hello"
    var v3 uint64 = 42
    var v4 int32 = 1
    switch v2 {
    case "hello":
        if (255 >= 4294967295) {
            var v5 int = 42
            var v6 uint64 = 4294967295
            v4 = v4
            _ = v5
            _ = v6
        } else {
            var v7 int8 = 127
            _ = v7
        }
    default:
        if true {
            fmt.Println(v3, v1)
        } else {
            v8 := GenS1{F0: -127, F1: -9223372036854775808, F2: 227}
            _ = v8
        }
    }
    switch -94 {
    case -820:
        if false {
            v9 := []int{-1}
            var v10 int32 = (v4 - 42)
            if len(v9) > 0 {
                v9[0] = -893
            }
            _ = v10
            _ = v9
        } else {
            var v11 uint32 = 1
            var v12 int8 = 100
            v1 += (v1 << ((4294967295) & 63))
            _ = v11
            _ = v12
        }
    case -916:
        var v13 uint8 = (v1 | (v1 << ((0) & 63)))
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(v3, v1, v1)
        v15 := GenS2{F0: 65535, F1: -1, F2: ("\n\t" + v2)}
        _ = fi14
        _ = v15
    }
    if (3.14159 < (1e10 + 0.0)) {
        fmt.Println(v4)
        if (!false) {
            if ((831 == v4) || true) {
                v16 := []int{-2147483648, 42}
                _ = v16
            } else {
                if (true && false) {
                    v1 = v1
                    for fi17 := 0; fi17 < 9; fi17++ {
                        var v18 uint64 = (v3 ^ v3)
                        _ = fi17
                        _ = v18
                    }
                }
            }
            v19 := []int{-2147483648, 127, 2147483647, 32768}
            _ = v19
        }
        var v20 bool = (-346 == 1)
        _ = v20
    } else {
        switch 127 {
        case 0:
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 float64 = float64(0.0)
                _ = fi21
                _ = v22
            }
        case 603:
            v23 := GenS3{F0: v3}
            _ = v23
        case 42:
            for fi24 := 0; fi24 < 12; fi24++ {
                v1 = v1
                v1 += v1
                if (2 > (v3 / (((^(v3))) | 1))) {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        v26 := []int{fi24}
                        if len(v26) > 0 {
                            v26[0] = fi24
                        }
                        _ = fi25
                        _ = v26
                    }
                    fmt.Println(v1, v2)
                    v1 -= v1
                } else {
                    if true {
                        v27 := []int{fi24, fi24, (fi24 / ((fi24) | 1))}
                        _ = v27
                    } else {
                        var v28 uint64 = v3
                        var v29 int16 = 42
                        var v30 uint8 = (v1 | 41)
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    v31 := []int{fi24}
                    var v32 int32 = v4
                    _ = v31
                    _ = v32
                }
                _ = fi24
            }
        default:
            if false {
                v33 := []int{127, 128, 32768, 9223372036854775807}
                v34 := GenS2{F0: 65535, F1: int64(v3), F2: ((("foo bar" + v2) + ("hello" + ("\n\t" + v2))) + "foo bar")}
                _ = v33
                _ = v34
            }
        }
    }
    var v35 bool = true
    v2 = v2
    v36 := []int{885, 2147483647, -933}
    var v37 float32 = float32(0.1)
    var v38 string = ("" + ("foo bar" + "\n\t"))
    var v39 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
