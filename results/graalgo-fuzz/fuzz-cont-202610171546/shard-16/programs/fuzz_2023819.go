package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint8; F1 int }

func main() {
    var v1 uint32 = 65535
    var v2 string = ("a" + ("foo bar" + ("\n\t" + "hello")))
    var v3 bool = true
    v4 := []string{v2, (("hello" + "") + ("" + v2)), (v2 + "\n\t"), (("foo bar" + v2) + v2)}
    v5 := make(map[string]string)
    var v6 int32 = 1
    if v3 {
        switch v2 {
        case "foo bar":
            v7 := GenS1{F0: 1}
            _ = v7
        case "":
            var v8 uint = 255
            _ = v8
        }
        var v9 uint = 1
        v10 := GenS2{F0: 1, F1: v3, F2: 1}
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 0}
    var v12 float64 = 0.1
    switch 128 {
    case 32767:
        if (v3 || (true && v3)) {
            var v13 uint = 2
            _ = v13
        } else {
            v14 := GenS2{F0: 65535, F1: ((!((301 <= 18446744073709551615) && ((v3 || (0.0 >= 0.0)) || (1 > 255)))) && v3), F2: 255}
            var v15 float64 = (1.5 - (((((v12 + 1e10) + (v12 * 0.1)) / v12) * float64((((0.1 * 1e10) - 3.14159) * (float32((0.1 + (float64(-2.25) * (v12 + v12)))) / (float32((float64(0.1) / 1.0)) * (1.5 * ((1.5 * 0.1) - float32(((v12 + v12) - 1e10))))))))) * (v12 / (v12 + 1e10))))
            _ = v14
            _ = v15
        }
    }
    if (!false) {
        for fi16 := 0; fi16 < 13; fi16++ {
            _ = len(v5)
            _ = fi16
        }
        v6 += 2147483647
        var v17 int8 = 1
        _ = v17
    } else {
        var v18 uint8 = 128
        for fi19 := 0; fi19 < 20; fi19++ {
            v20 := GenS2{F0: 516, F1: false, F2: 388}
            v21 := GenS3{F0: 42, F1: -877}
            _ = fi19
            _ = v20
            _ = v21
        }
        v22 := GenS1{F0: v18}
        _ = v18
        _ = v22
    }
    if v3 {
        v3 = (!(v6 == (v6 * v6)))
        v12 = v12
    } else {
        var v23 int8 = int8(v1)
        _ = v23
    }
    v24 := GenS2{F0: 32767, F1: v3, F2: 1}
    var v25 float64 = v12
    switch "a" {
    case "hello":
        if (1e10 == 1.5) {
            switch -54 {
            case 978:
                var v26 int64 = 42
                _ = v26
            }
            if len(v4) > 0 {
                v4[0] = v2
            }
            for fi27 := 0; fi27 < 10; fi27++ {
                v2 = ((v2 + ((((v2 + ("foo bar" + (("" + v2) + (v2 + ((v2 + "") + (v2 + (v2 + v2))))))) + (("" + v2) + v2)) + (v2 + v2)) + "foo bar")) + v2)
                _ = fi27
            }
        } else {
            v1 -= v1
            for fi28 := 0; fi28 < 17; fi28++ {
                v2 = v2
                _ = fi28
            }
            var v29 bool = (-229 >= 42)
            _ = v29
        }
    case "\n\t":
        for fi30 := 0; fi30 < 9; fi30++ {
            for fi31 := 0; fi31 < 3; fi31++ {
                delete(v5, v2)
                if ((v2 + v2) <= v2) {
                    _ = len(v5)
                }
                _ = fi31
            }
            _ = fi30
        }
    default:
        switch 0 {
        case -949:
            for fi32 := 0; fi32 < 14; fi32++ {
                for fi33 := 0; fi33 < 16; fi33++ {
                    var v34 bool = v3
                    _ = fi33
                    _ = v34
                }
                if true {
                    v35 := GenS1{F0: 255}
                    var v36 int8 = int8(v6)
                    v37 := GenS3{F0: (v36 ^ 100), F1: -9223372036854775808}
                    _ = v35
                    _ = v36
                    _ = v37
                } else {
                    var v38 uint8 = 199
                    _ = v38
                }
                if ((!v3) && false) {
                    var v39 float64 = v25
                    for fi40 := 0; fi40 < 19; fi40++ {
                        v1 += (^(v1))
                        _ = fi40
                    }
                    _ = v39
                }
                _ = fi32
            }
        }
    }
    var v41 int = 1
    var v42 uint16 = 3
    var v43 uint16 = 1
    var v44 bool = (!(4294967295 <= 2))
    v45 := GenS4{F0: 173, F1: v41}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
