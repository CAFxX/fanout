package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int64 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint16; F1 int64; F2 string }
type GenS5 struct { F0 bool; F1 uint; F2 int32 }
type GenS6 struct { F0 uint32; F1 int16; F2 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 0
    var v3 uint8 = 42
    var v4 float32 = 0.0
    var v5 float64 = 1e10
    var v6 uint8 = v3
    fmt.Println(v4, v2)
    var v7 uint8 = (v3 + (v6 & (v6 << ((18446744073709551615) & 63))))
    var v8 int64 = -9223372036854775808
    if ((!(false && (-1 > 0))) || (v2 <= 773)) {
        if (!true) {
            switch 2147483647 {
            case 1:
                v5 = v5
            default:
                fmt.Println(v3)
            }
        }
    } else {
        if true {
            v4 -= v4
            v9 := []int{921, -554, 0}
            _ = v9
        }
        if (2147483647 > 42) {
            v4 -= (v4 + ((1.5 * v4) - (v4 - (float32(1.5) / v4))))
            if false {
                v2 -= 2
                var v10 int16 = 1000
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            if true {
                v5 = 0.1
                var v12 int32 = -215
                _ = v12
            }
            v13 := GenS1{F0: 0, F1: (v8 << ((452) & 63)), F2: (v8 >> ((255) & 63))}
            v14 := []int{(fi11 & (fi11 + 4294967295)), fi11, (-(fi11))}
            _ = fi11
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v6 = 42
        if false {
            v16 := GenS1{F0: 396, F1: v8, F2: (v8 << ((4294967295) & 63))}
            _ = v16
        }
        fmt.Println(fi15, v3)
        _ = fi15
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v18 := v1[("hello" + "")]
        var v19 uint8 = 127
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := []int{fi20, fi20, -103}
            if len(v21) > 0 {
                v21[0] = 1
            }
            _ = fi20
            _ = v21
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v7, v6)
    var v22 int16 = -32768
    for fi23 := 0; fi23 < 13; fi23++ {
        for fi24 := 0; fi24 < 10; fi24++ {
            v25 := GenS2{F0: (!false), F1: (^(v2))}
            _ = fi24
            _ = v25
        }
        var v26 float32 = v4
        if (4294967295 != (fi23 * int(v2))) {
            v6 = 127
            var v27 int16 = v22
            _ = v27
        } else {
            for fi28 := 0; fi28 < 13; fi28++ {
                if false {
                    v29 := GenS1{F0: 2, F1: v8, F2: v8}
                    _ = v29
                } else {
                    var v30 uint8 = 255
                    switch "a" {
                    case "":
                        if true {
                            v2 -= v2
                        }
                    default:
                        var v31 int = (^(fi23))
                        _ = v31
                    }
                    switch fi23 {
                    case 2147483647:
                        v2 = v2
                    case 42:
                        var v32 int64 = -819
                        _ = v32
                    case 32768:
                        v22 -= v22
                    }
                    _ = v30
                }
                v2 = 534
                var v33 uint = 880
                v34 := GenS3{F0: v6, F1: (fi23 ^ (-(fi28)))}
                v35 := GenS4{F0: 65535, F1: 9223372036854775807, F2: "\n\t"}
                _ = fi28
                _ = v33
                _ = v34
                _ = v35
            }
            v36 := GenS5{F0: true, F1: 935, F2: 561}
            v4 -= 3.14159
            _ = v36
        }
        _ = fi23
        _ = v26
    }
    v37 := GenS4{F0: 32767, F1: int64(v2), F2: ("a" + ("hello" + (("\n\t" + ((("hello" + "") + ("hello" + ("hello" + "\n\t"))) + ("\n\t" + "hello"))) + ("\n\t" + ("hello" + "a")))))}
    v38 := GenS6{F0: v2, F1: -1, F2: v22}
    v39 := GenS5{F0: (int(v6) < 42), F1: 1, F2: 0}
    for fi40 := 0; fi40 < 14; fi40++ {
        v3 -= (v3 * 127)
        v7 -= 128
        v3 -= v6
        _ = fi40
    }
    v41 := v1["a"]
    v42 := GenS5{F0: true, F1: 0, F2: -1}
    if (!false) {
        fmt.Println(v22)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
