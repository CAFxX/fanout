package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint }
type GenS5 struct { F0 uint8; F1 int16; F2 int16 }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 uint8 }
func genFn0() uint8 {
    if true {
        v4 := GenS1{F0: -2147483648}
        _ = v4
    } else {
        var v5 uint = 0
        v6 := []int{32768, 0, 32768, -9223372036854775808}
        v7 := GenS1{F0: 0}
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := []int{-442, 928, -497, 42}
    var v9 int = -631
    v10 := GenS1{F0: -625}
    v11 := GenS2{F0: 819}
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{65535, 341, 2147483647, 692}
    var v3 float64 = 3.14159
    v3 -= float64(1e10)
    if ("\n\t" < "") {
        v2 = append(v2, 42)
        fmt.Println(v3)
    } else {
        if false {
            v3 = (v3 - v3)
            for fi12 := 0; fi12 < 16; fi12++ {
                _ = len(v1)
                _ = fi12
            }
            var v13 float32 = (0.0 * 1e10)
            _ = v13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if len(v2) > 0 {
                v2[0] = 91
            }
            _ = fi14
        }
    }
    if false {
        if true {
            if true {
                v3 -= 1e10
            }
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 uint64 = 954
                v17 := GenS3{F0: 127}
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    }
    switch -642 {
    case 42:
        v3 = 3.14159
    case 127:
        v2 = append(v2, 65535)
    case 1:
        if (!((true && (198 <= 65535)) || true)) {
            v1[((("foo bar" + "hello") + ("\n\t" + "a")) + "hello")] = 42
            for fi18 := 0; fi18 < 17; fi18++ {
                fmt.Println(fi18)
                var v19 float64 = v3
                var v20 uint16 = 730
                _ = fi18
                _ = v19
                _ = v20
            }
            if false {
                v3 += v3
                for fi21 := 0; fi21 < 6; fi21++ {
                    v1[""] = fi21
                    if (3.14159 > 3.14159) {
                        fmt.Println(fi21, fi21)
                        fmt.Println(v3, v3)
                    } else {
                        fmt.Println(fi21, v3)
                    }
                    switch (fi21 << ((4294967295) & 63)) {
                    case 42:
                        var v22 float32 = -2.25
                        _ = v22
                    default:
                        v23 := GenS4{F0: -473, F1: 983, F2: 2}
                        _ = v23
                    }
                    _ = fi21
                }
            } else {
                v3 += (v3 * v3)
                v24 := GenS5{F0: 81, F1: -485, F2: 42}
                _ = v24
            }
        }
    }
    var v25 uint16 = 255
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v25 -= (v25 * v25)
    v3 -= ((0.0 + 1e10) + v3)
    if false {
        fmt.Println(v3, v3)
        for fi26 := 0; fi26 < 14; fi26++ {
            if (1 >= 2) {
                fmt.Println(v3)
                v27 := v1[("foo bar" + "")]
                var v28 int = fi26
                _ = v27
                _ = v28
            } else {
                if false {
                    v25 = (v25 ^ v25)
                    v29 := GenS4{F0: 9223372036854775807, F1: 145, F2: 708}
                    _ = v29
                } else {
                    fmt.Println(v3)
                    v3 = 0.0
                }
                var v30 int32 = 2147483647
                v31 := GenS4{F0: int64(fi26), F1: 1, F2: 1}
                _ = v30
                _ = v31
            }
            var v32 string = "foo bar"
            var v33 uint8 = 2
            _ = fi26
            _ = v32
            _ = v33
        }
    }
    v3 -= ((v3 * v3) * float64(0.1))
    _ = len(v1)
    if false {
        v2 = append(v2, 877)
        if ((42 != 108) && (0 >= -2147483648)) {
            fmt.Println(v25, v25)
            v34 := GenS6{F0: 42, F1: 523}
            _ = v34
        } else {
            v3 = (1e10 + float64(1.5))
            switch (("\n\t" + "foo bar") + "a") {
            case "\n\t":
                fmt.Println(v3, v25)
            case "foo bar":
                var v35 int16 = -32768
                _ = v35
            }
        }
        if true {
            var v36 float32 = -2.25
            v37 := GenS7{F0: 127}
            _ = v36
            _ = v37
        } else {
            switch 42 {
            case 2147483647:
                fmt.Println(v3, v25)
            case 32767:
                _ = len(v2)
            case 32768:
                delete(v1, "a")
            default:
                if true {
                    fmt.Println(v3, v3)
                } else {
                    var v38 float32 = -2.25
                    v38 += v38
                    _ = v38
                }
            }
            var v39 float32 = 1e10
            v25 = (v25 - 1)
            _ = v39
        }
    }
    var v40 int = 127
    fmt.Println(v25)
    v3 = 0.0
    for fi41 := 0; fi41 < 16; fi41++ {
        var v42 int64 = -1
        v3 += -2.25
        v3 = v3
        _ = fi41
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v40
}
