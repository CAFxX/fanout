package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int32; F1 uint8; F2 int }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 int16 }

func main() {
    var v1 float32 = 0.1
    var v2 uint32 = 109
    var v3 int32 = 0
    if (true || false) {
        for fi4 := 0; fi4 < 3; fi4++ {
            v5 := GenS1{F0: v2}
            var v6 uint16 = 2
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 int32 = (v3 ^ v3)
                var v9 float64 = ((1.5 - (3.14159 + -2.25)) / -2.25)
                v10 := GenS1{F0: (v2 ^ v2)}
                v11 := GenS2{F0: 2}
                if (!(!true)) {
                    switch -230 {
                    case -374:
                        fmt.Println(v2, v3)
                    case 342:
                        var v12 float32 = v1
                        _ = v12
                    }
                    v13 := []int{(fi7 - fi4), fi4, (fi4 & 249), fi4}
                    var v14 int64 = -247
                    v15 := GenS2{F0: 65535}
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = fi7
                _ = v10
                _ = v11
                _ = v8
                _ = v9
            }
            _ = fi4
            _ = v5
            _ = v6
        }
        var v16 uint = 4294967295
        var v17 bool = false
        _ = v16
        _ = v17
    }
    if true {
        fmt.Println(v2, v2)
        if (!(0.1 > 1e10)) {
            var v18 int32 = v3
            v18 += (-(v3))
            v19 := GenS3{F0: -15, F1: ("foo bar" + ("hello" + "hello"))}
            _ = v18
            _ = v19
        } else {
            switch "foo bar" {
            case "foo bar":
                v1 = 3.14159
            default:
                var v20 int16 = 42
                _ = v20
            }
        }
    } else {
        switch 4294967295 {
        case 1:
            v21 := []int{4294967295, -626, 9223372036854775807, 9223372036854775807}
            _ = v21
        case 277:
            v1 = float32((((1.5 - ((1.5 - -2.25) * (1e10 + ((0.1 / 3.14159) * (1.5 + 1.5))))) - 1.5) + 1.5))
        default:
            v22 := GenS1{F0: 470}
            _ = v22
        }
        var v23 uint8 = 128
        _ = v23
    }
    v3 += (v3 / ((v3) | 1))
    switch -9223372036854775808 {
    case -132:
        fmt.Println(v2)
    default:
        if true {
            var v24 float32 = float32((1e10 / 1.5))
            for fi25 := 0; fi25 < 6; fi25++ {
                v26 := GenS2{F0: 2}
                v27 := []int{(fi25 * fi25)}
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = v24
        } else {
            switch -932 {
            case -9223372036854775808:
                switch 4294967295 {
                case 0:
                    var v28 int16 = -32768
                    _ = v28
                default:
                    v29 := GenS4{F0: 42, F1: 0, F2: 127}
                    _ = v29
                }
            case -516:
                v30 := []int{}
                _ = v30
            case 128:
                var v31 int16 = 0
                _ = v31
            }
        }
    }
    v3 -= (^(v3))
    switch "" {
    case "hello":
        v32 := []int{9223372036854775807, 4294967295, 32768}
        _ = v32
    }
    if false {
        v33 := GenS2{F0: 255}
        _ = v33
    }
    if (-128 != 1) {
        v1 -= 0.1
        v3 -= (v3 - v3)
    }
    switch "foo bar" {
    case "foo bar":
        switch 138 {
        case 500:
            if (1 == 42) {
                v3 += (v3 ^ v3)
                fmt.Println(v3, v2, v3)
                for fi34 := 0; fi34 < 11; fi34++ {
                    v35 := GenS5{F0: "", F1: (v2 ^ v2)}
                    if (-32768 >= 42) {
                        v36 := []int{}
                        v3 += (v3 + v3)
                        v1 = v1
                        _ = v36
                    } else {
                        fmt.Println(v1, v1)
                        v1 += v1
                    }
                    v37 := []int{fi34, (fi34 / ((-77) | 1))}
                    _ = fi34
                    _ = v35
                    _ = v37
                }
            }
        case -312:
            v2 -= (v2 ^ v2)
        case 42:
            if false {
                v38 := GenS6{F0: 468}
                v3 -= 0
                _ = v38
            } else {
                switch ("foo bar" + "hello") {
                case "hello":
                    v2 -= v2
                }
                fmt.Println(v1, v2)
            }
        default:
            v2 += v2
        }
    default:
        v1 += v1
    }
    var v39 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
}
