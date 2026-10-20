package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 uint32 }

func main() {
    var v1 float32 = (1.5 + 1e10)
    var v2 uint16 = 75
    var v3 int = 9223372036854775807
    var v4 uint = 65535
    for fi5 := 0; fi5 < 15; fi5++ {
        switch v3 {
        case 42:
            v4 = v4
        case 9223372036854775807:
            fmt.Println(v1)
        }
        fmt.Println(v2, fi5, v2)
        _ = fi5
    }
    for fi6 := 0; fi6 < 12; fi6++ {
        switch ("hello" + ("foo bar" + "foo bar")) {
        case "a":
            fmt.Println(v1)
        }
        v3 += v3
        if (true || (v3 == -1)) {
            v7 := GenS1{F0: true, F1: 2, F2: 42}
            switch "hello" {
            case "foo bar":
                if ((1e10 - (v1 + 3.14159)) != 1.5) {
                    v1 = v1
                    if false {
                        switch "foo bar" {
                        case "a":
                            v8 := []int{}
                            _ = v8
                        }
                        var v9 int8 = 0
                        _ = v9
                    } else {
                        v10 := GenS1{F0: false, F1: 255, F2: 4294967295}
                        v2 -= (v2 % (((v2 % (((v2 | (^(v2)))) | 1))) | 1))
                        _ = v10
                    }
                    v11 := []int{}
                    _ = v11
                } else {
                    switch "foo bar" {
                    case "hello":
                        var v12 uint64 = 2
                        _ = v12
                    }
                    v2 += (^(v2))
                    var v13 int32 = 0
                    _ = v13
                }
            }
            _ = v7
        } else {
            v1 = (-2.25 - v1)
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 string = "foo bar"
                if ((-1 >= 0) || (1.5 < float64((-2.25 - -2.25)))) {
                    if (1 >= 42) {
                        v4 -= (v4 / ((18446744073709551615) | 1))
                        v1 -= ((float32(1e10) * v1) - (((v1 / float32((0.0 + (3.14159 + -2.25)))) + (3.14159 / v1)) + ((0.0 * v1) / ((v1 * -2.25) + -2.25))))
                    }
                    v16 := GenS2{F0: 42}
                    switch (fi6 - fi6) {
                    case -1:
                        if false {
                            var v17 uint32 = 42
                            v3 -= 1
                            _ = v17
                        }
                    case 1:
                        var v18 uint32 = 652
                        _ = v18
                    }
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
            v19 := GenS3{F0: true}
            _ = v19
        }
        _ = fi6
    }
    if (833 > 9223372036854775807) {
        if (0 > -110) {
            var v20 uint8 = 20
            v21 := []int{}
            _ = v20
            _ = v21
        }
        v22 := GenS3{F0: ((v3 * v3) < v3)}
        _ = v22
    }
    var v23 uint8 = 0
    var v24 uint16 = v2
    v25 := GenS4{F0: v3, F1: 448}
    v24 -= v2
    if false {
        for fi26 := 0; fi26 < 10; fi26++ {
            v27 := []int{(fi26 % (((v3 ^ 42)) | 1)), (v3 * (v3 % (((fi26 >> (((v4 + (^(v4)))) & 63))) | 1)))}
            _ = fi26
            _ = v27
        }
        v28 := []int{v3, v3}
        switch ("" + ("a" + "a")) {
        case "foo bar":
            _ = len(v28)
        default:
            var v29 uint32 = 42
            _ = v29
        }
        _ = v28
    }
    fmt.Println(v1, v24, v4)
    var v30 int8 = 0
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 uint16 = v2
        if (!(4294967295 <= 34)) {
            fmt.Println(v32, v30, v1)
            fmt.Println(v24)
            var v33 uint32 = 2
            _ = v33
        } else {
            for fi34 := 0; fi34 < 5; fi34++ {
                var v35 uint16 = v24
                _ = fi34
                _ = v35
            }
            v1 = v1
        }
        v24 = (v32 - (^(v32)))
        _ = fi31
        _ = v32
    }
    var v36 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v36
    _ = v4
}
