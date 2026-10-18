package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 bool; F2 uint64 }
type GenS3 struct { F0 int8; F1 int16; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 uint64; F1 int; F2 uint64 }

func main() {
    var v1 uint8 = 2
    var v2 int32 = -1
    var v3 float64 = (0.0 - (((1e10 / 1.5) - 1.5) / (1.5 / float64((-2.25 + 1.5)))))
    var v4 int8 = 94
    v5 := make(map[string]uint32)
    v6 := GenS1{F0: -950}
    v7 := GenS1{F0: 1000}
    v4 = (v4 >> ((232) & 63))
    switch 1 {
    case 128:
        _ = len(v5)
    case 9223372036854775807:
        if (0.1 <= (3.14159 + v3)) {
            fmt.Println(v4, v2, v1)
        }
    default:
        v1 = (v1 / ((v1) | 1))
    }
    fmt.Println(v2)
    switch "\n\t" {
    case "foo bar":
        switch -881 {
        case -144:
            v2 = v2
        case 128:
            fmt.Println(v3, v1)
        case -2147483648:
            if (2147483647 < 2147483647) {
                for fi8 := 0; fi8 < 8; fi8++ {
                    v2 = v2
                    _ = fi8
                }
                v9 := v5["\n\t"]
                _ = v9
            }
        }
    case "\n\t":
        for fi10 := 0; fi10 < 15; fi10++ {
            switch fi10 {
            case 1:
                var v11 int8 = -65
                _ = v11
            default:
                fmt.Println(fi10)
            }
            var v12 float32 = (3.14159 / (0.1 / (0.0 + ((1.5 + 0.1) - -2.25))))
            v13 := GenS2{F0: (^(v4)), F1: false, F2: 2}
            _ = fi10
            _ = v12
            _ = v13
        }
    }
    switch -369 {
    case 410:
        if (353 >= 1) {
            var v14 string = ("foo bar" + ("hello" + "\n\t"))
            var v15 int32 = -2147483648
            _ = v14
            _ = v15
        } else {
            v3 = 0.0
            _ = len(v5)
            fmt.Println(v2, v2, v3)
        }
    }
    if (true || (-1 != -32768)) {
        fmt.Println(v4, v1, v4)
        v16 := []int{42}
        v3 = 3.14159
        _ = v16
    } else {
        v1 += 255
        for fi17 := 0; fi17 < 17; fi17++ {
            if false {
                switch fi17 {
                case 1:
                    fmt.Println(v2)
                case 42:
                    fmt.Println(v2, v3)
                }
                v18 := GenS3{F0: 101, F1: 1000, F2: 580}
                _ = v18
            } else {
                _ = len(v5)
            }
            _ = fi17
        }
    }
    var v19 float32 = (-2.25 - 0.1)
    v20 := GenS1{F0: -984}
    var v21 uint8 = v1
    v22 := GenS4{F0: 32767}
    for fi23 := 0; fi23 < 14; fi23++ {
        v19 = (v19 * v19)
        _ = fi23
    }
    _ = len(v5)
    for fi24 := 0; fi24 < 2; fi24++ {
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 bool = (1e10 > v19)
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    switch 0 {
    case 32767:
        var v27 uint8 = (v21 ^ 193)
        _ = v27
    case 275:
        for fi28 := 0; fi28 < 5; fi28++ {
            if false {
                v29 := v5[("a" + ("foo bar" + ("" + "")))]
                _ = v29
            }
            delete(v5, "hello")
            _ = fi28
        }
    case 9223372036854775807:
        for fi30 := 0; fi30 < 4; fi30++ {
            for fi31 := 0; fi31 < 7; fi31++ {
                if false {
                    if (true && ((false && false) || false)) {
                        v32 := []int{fi30, -2147483648, int(v4), fi30}
                        var v33 int32 = 0
                        v19 = float32((0.1 + ((((v3 * float64(v19)) - 1.5) + v3) + (v3 / (v3 - v3)))))
                        _ = v32
                        _ = v33
                    } else {
                        fmt.Println(fi30, v2)
                        v4 -= 72
                        var v34 int64 = 321
                        v35 := GenS5{F0: fi30, F1: 2}
                        _ = v34
                        _ = v35
                    }
                    v21 = (v1 >> ((4294967295) & 63))
                }
                v2 = 2147483647
                v36 := GenS6{F0: 0, F1: 128, F2: 464}
                var v37 uint = 608
                var v38 uint32 = uint32(v37)
                _ = fi31
                _ = v36
                _ = v37
                _ = v38
            }
            v19 += float32(0.1)
            _ = fi30
        }
    }
    v3 = v3
    if true {
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
