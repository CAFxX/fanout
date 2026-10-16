package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int64; F1 bool; F2 uint16 }

func main() {
    var v1 uint = 47
    var v2 uint32 = 42
    var v3 bool = false
    var v4 uint32 = (v2 & 1)
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(v1, v3)
            if ((611 < 0) || ((!(32767 != -749)) && (v3 && (42 < 42)))) {
                var v7 uint16 = 513
                v1 = v1
                _ = v7
            } else {
                var v8 int32 = 0
                if (!(-128 == 111)) {
                    v9 := GenS1{F0: "hello", F1: (-(v8)), F2: v2}
                    _ = v9
                } else {
                    var v10 int16 = -949
                    if (357 < (v8 % (((v8 >> (((v1 ^ v1)) & 63))) | 1))) {
                        var v11 int16 = (v10 * (^(v10)))
                        if (!(v3 && (!v3))) {
                            v4 += (v2 + v4)
                            var v12 uint64 = 214
                            v3 = v3
                            _ = v12
                        } else {
                            v13 := GenS1{F0: (("hello" + ("foo bar" + ("hello" + "a"))) + "hello"), F1: v8, F2: 4294967295}
                            v14 := GenS2{F0: 210, F1: (v4 & v2), F2: 0}
                            _ = v13
                            _ = v14
                        }
                        _ = v11
                    }
                    fmt.Println(v4, fi6, v4)
                    _ = v10
                }
                _ = v8
            }
            _ = fi6
        }
        fmt.Println(fi5, v4, v1)
        v1 = v1
        _ = fi5
    }
    var v15 int32 = 42
    if (v3 || ((v15 <= v15) || ((float32(1.5) != 0.0) && ((v4 << ((v1) & 63)) == (^(v4)))))) {
        fmt.Println(v1, v2)
        v2 += (v2 >> ((837) & 63))
        var v16 uint8 = 128
        _ = v16
    } else {
        var v17 uint = (^(v1))
        if (32767 == 0) {
            v15 = (v15 ^ (v15 * (v15 ^ -1)))
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 uint32 = (v2 ^ 2)
                var v20 bool = v3
                v19 = v2
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        _ = v17
    }
    v15 -= v15
    if v3 {
        v3 = (-2147483648 <= 128)
        v4 = 905
    } else {
        if false {
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 int8 = 42
                v23 := []int{}
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 bool = (false && v3)
            var v25 string = ""
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 int32 = (v15 >> ((v1) & 63))
        v28 := GenS2{F0: 1, F1: (v4 & (v4 - (v4 * 548))), F2: 0}
        var v29 uint32 = (v2 - v2)
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    fmt.Println(v2, v15, v1)
    v30 := GenS3{F0: -296}
    var v31 int32 = (-(v15))
    switch 4294967295 {
    case -9223372036854775808:
        v32 := GenS4{F0: ("" + "hello"), F1: v3}
        _ = v32
    case -1:
        var v33 int = -2147483648
        _ = v33
    default:
        if false {
            v3 = (((((float32((0.0 - ((1.5 * float64(float32(0.0))) - (float64(3.14159) / 1e10)))) / float32(1.5)) + -2.25) - (float32((-2.25 + 0.0)) + 1.5)) + 0.1) >= (0.1 + (0.0 * (0.0 * 3.14159))))
            switch "foo bar" {
            case "a":
                switch (("hello" + "\n\t") + ("" + (("hello" + "\n\t") + ""))) {
                case "":
                    if true {
                        var v34 int8 = 42
                        _ = v34
                    } else {
                        v3 = (!(42 > -9223372036854775808))
                    }
                case "foo bar":
                    switch 1 {
                    case 42:
                        var v35 uint32 = v4
                        _ = v35
                    }
                case "a":
                    v36 := GenS5{F0: 802, F1: (1.5 > float64((float32(float64(0.1)) / (3.14159 / 1.5)))), F2: 65535}
                    _ = v36
                default:
                    for fi37 := 0; fi37 < 18; fi37++ {
                        var v38 uint = v1
                        v38 -= v38
                        v39 := []int{42, int(v4), 2147483647}
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                }
            }
        } else {
            var v40 int64 = 342
            _ = v40
        }
    }
    v31 = v15
    var v41 int32 = 42
    v15 = v31
    v3 = v3
    for fi42 := 0; fi42 < 19; fi42++ {
        var v43 int32 = v15
        var v44 uint8 = 1
        _ = fi42
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v41
}
