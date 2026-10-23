package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS2 struct { F0 int8; F1 int; F2 uint }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint8; F1 int64; F2 int8 }

func main() {
    var v1 int16 = 1000
    var v2 uint16 = 65535
    var v3 uint8 = 127
    switch 9223372036854775807 {
    case -961:
        fmt.Println(v3)
    case 691:
        switch 32768 {
        case 1:
            v2 += (v2 | (^(v2)))
        }
    }
    if ((float64(1.5) * float64(-2.25)) <= 0.0) {
        v4 := []int{}
        _ = len(v4)
        if (!("hello" != ("a" + "\n\t"))) {
            v3 += 42
            v5 := GenS1{F0: (v2 + 32767), F1: 42, F2: 2}
            var v6 int64 = 997
            _ = v5
            _ = v6
        }
        _ = v4
    } else {
        var v7 int32 = -94
        var v8 uint64 = 1
        _ = v7
        _ = v8
    }
    fmt.Println(v2, v1, v1)
    if ((v1 + (v1 ^ 1)) <= 1) {
        var v9 uint8 = 42
        _ = v9
    } else {
        fmt.Println(v2, v1, v1)
        for fi10 := 0; fi10 < 10; fi10++ {
            v3 += v3
            _ = fi10
        }
        v3 += v3
    }
    switch "" {
    case "foo bar":
        var v11 float32 = ((0.0 * 1e10) - 3.14159)
        _ = v11
    case "\n\t":
        fmt.Println(v3)
    case "a":
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS2{F0: -128, F1: 528, F2: 178}
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS3{F0: 1, F1: 42}
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    default:
        for fi16 := 0; fi16 < 14; fi16++ {
            if false {
                v17 := GenS4{F0: 997, F1: true}
                v18 := []int{(fi16 * 31), fi16, (-(fi16)), -9223372036854775808}
                _ = v17
                _ = v18
            }
            var v19 uint = 2
            fmt.Println(v3)
            _ = fi16
            _ = v19
        }
    }
    if true {
        v1 -= -1
    } else {
        if (!(!(1 <= 65535))) {
            var v20 uint8 = (v3 & v3)
            v21 := []int{1, 32767, 9223372036854775807}
            _ = v20
            _ = v21
        } else {
            v22 := GenS5{F0: 19, F1: 1, F2: 127}
            var v23 int8 = 88
            v24 := GenS4{F0: v2, F1: (0 == 65535)}
            _ = v22
            _ = v23
            _ = v24
        }
        var v25 uint8 = v3
        v1 = 1
        _ = v25
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        for fi27 := 0; fi27 < 15; fi27++ {
            if (63 == 851) {
                v3 += 128
                var v28 uint = 295
                _ = v28
            }
            _ = fi27
        }
        var v29 int32 = 1
        var v30 uint64 = 4294967295
        _ = fi26
        _ = v29
        _ = v30
    }
    switch "hello" {
    case "hello":
        if (!(1 < -258)) {
            var v31 int32 = 2147483647
            v32 := []int{32768, 32767}
            v33 := GenS2{F0: int8(v31), F1: -58, F2: 801}
            _ = v31
            _ = v32
            _ = v33
        }
    default:
        if true {
            fmt.Println(v2, v3, v2)
        } else {
            switch "foo bar" {
            case "":
                var v34 float64 = 0.1
                _ = v34
            default:
                v35 := []int{32767, 128, 0, -1}
                _ = v35
            }
            v1 = v1
        }
    }
    v2 += 212
    var v36 uint8 = (v3 * 164)
    for fi37 := 0; fi37 < 4; fi37++ {
        var v38 uint16 = v2
        _ = fi37
        _ = v38
    }
    var v39 int16 = (v1 % (((v1 | (v1 + -204))) | 1))
    for fi40 := 0; fi40 < 15; fi40++ {
        var v41 uint8 = (v36 * v36)
        if false {
            var v42 string = "a"
            var v43 uint = 42
            _ = v42
            _ = v43
        } else {
            var v44 int64 = -9223372036854775808
            v45 := GenS2{F0: 42, F1: (^(fi40)), F2: 2}
            v1 = v39
            _ = v44
            _ = v45
        }
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v39
}
