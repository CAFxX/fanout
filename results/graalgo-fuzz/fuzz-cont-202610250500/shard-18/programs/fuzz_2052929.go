package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    var v1 uint16 = 869
    var v2 bool = (false || (!(int64(v1) > 0)))
    var v3 string = (("a" + "") + "hello")
    var v4 int = -998
    var v5 int64 = 285
    v6 := GenS1{F0: 2147483647}
    if false {
        if (1 <= 25) {
            v5 = v5
            fmt.Println(v1, v5, v4)
        }
    } else {
        var v7 uint32 = 4294967295
        switch (("\n\t" + ((("foo bar" + (((v3 + v3) + v3) + v3)) + ((("hello" + (v3 + "a")) + ("foo bar" + (v3 + v3))) + ((v3 + "a") + ((("\n\t" + v3) + ((v3 + "foo bar") + v3)) + "")))) + ("a" + ((((("hello" + ("\n\t" + v3)) + v3) + v3) + "a") + v3)))) + v3) {
        case "a":
            switch v4 {
            case 9223372036854775807:
                v8 := GenS2{F0: 42, F1: -1}
                _ = v8
            }
        case "":
            var v9 uint64 = 237
            _ = v9
        default:
            var v10 uint64 = 112
            _ = v10
        }
        fmt.Println(v4, v4)
        _ = v7
    }
    if (v2 || ((v1 % ((v1) | 1)) > (v1 >> ((1) & 63)))) {
        var v11 bool = (!(258 > 65535))
        var v12 bool = false
        v5 = (v5 ^ (v5 / (((v5 + int64(v1))) | 1)))
        _ = v11
        _ = v12
    } else {
        var v13 uint32 = uint32(v1)
        fmt.Println(v5, v2, v13)
        if (100 < 1) {
            var v14 string = v3
            var v15 float64 = 0.0
            v16 := GenS1{F0: 0}
            _ = v14
            _ = v15
            _ = v16
        }
        _ = v13
    }
    v5 = 9223372036854775807
    if true {
        if v2 {
            if false {
                v17 := GenS1{F0: int32(v1)}
                if false {
                    fmt.Println(v4, v3)
                    var v18 float32 = -2.25
                    _ = v18
                }
                _ = v17
            } else {
                v19 := GenS1{F0: int32(v4)}
                _ = v19
            }
            fmt.Println(v4, v5, v3)
            var v20 uint8 = 255
            _ = v20
        } else {
            if (0 > 221) {
                var v21 uint16 = (v1 >> ((2) & 63))
                _ = v21
            }
            var v22 uint32 = 423
            _ = v22
        }
    } else {
        for fi23 := 0; fi23 < 8; fi23++ {
            if (v1 > 280) {
                if v2 {
                    var v24 int16 = -470
                    _ = v24
                } else {
                    switch v3 {
                    case "\n\t":
                        v25 := []int{fi23}
                        _ = v25
                    default:
                        fmt.Println(v1, fi23)
                    }
                    v2 = ((2 < 2) || true)
                }
                for fi26 := 0; fi26 < 6; fi26++ {
                    v27 := []int{fi23}
                    v28 := GenS1{F0: 1}
                    v3 = v3
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            _ = fi23
        }
        if true {
            v29 := []int{(v4 >> ((255) & 63)), v4, v4}
            fmt.Println(v1, v4, v5)
            var v30 uint8 = 128
            _ = v29
            _ = v30
        } else {
            var v31 uint16 = v1
            _ = v31
        }
        var v32 uint64 = 42
        _ = v32
    }
    var v33 int32 = -1
    fmt.Println(v1)
    for fi34 := 0; fi34 < 6; fi34++ {
        v35 := GenS2{F0: 509, F1: 1}
        if (!(0 < (v33 ^ (v33 % ((v33) | 1))))) {
            if (v5 >= (v5 + -1)) {
                v36 := []int{(-(v4))}
                _ = v36
            }
        }
        if (1e10 <= float64(float32(float64(0.0)))) {
            v37 := []int{}
            _ = v37
        }
        _ = fi34
        _ = v35
    }
    switch (v4 - v4) {
    case 2147483647:
        v3 = v3
    case 9223372036854775807:
        if v2 {
            var v38 uint64 = 4294967295
            _ = v38
        }
    default:
        var v39 uint = 2
        _ = v39
    }
    v40 := []int{(v4 * (^(v4))), v4, v4, (-(v4))}
    v40 = append(v40, -2147483648)
    v33 += v33
    for fi41 := 0; fi41 < 20; fi41++ {
        var v42 int32 = (v33 + (v33 % ((v33) | 1)))
        if (103 < 825) {
            for fi43 := 0; fi43 < 3; fi43++ {
                var v44 int32 = v33
                if len(v40) >= 2 {
                    v40 = v40[1:len(v40)]
                }
                _ = fi43
                _ = v44
            }
        }
        _ = fi41
        _ = v42
    }
    if len(v40) > 0 {
        v40[0] = 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
