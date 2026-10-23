package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint16, p1 uint) uint32 {
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int64 = 9223372036854775807
        p0 = (^(p0))
        v7 := []int{(fi5 | -1), fi5, 9223372036854775807}
        _ = fi5
        _ = v6
        _ = v7
    }
    if (3.14159 >= float64(3.14159)) {
        var v8 string = ((("hello" + ("a" + ("a" + ""))) + ((("" + "\n\t") + "hello") + "hello")) + "")
        v9 := []int{42}
        if true {
            v10 := GenS1{F0: (false && false), F1: uint16(p1), F2: 42}
            v11 := GenS2{F0: (^(p1))}
            _ = v10
            _ = v11
        }
        _ = v8
        _ = v9
    }
    v12 := GenS3{F0: 2, F1: 32768}
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 string = "hello"
        var v15 int16 = 984
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v12
    return 201
}

func main() {
    var v1 bool = false
    var v2 float64 = 3.14159
    var v3 int16 = 1
    var v4 uint = 2
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(v1, v3)
        v4 = v4
        _ = fi16
    }
    if v1 {
        fmt.Println(v4, v4)
        switch 323 {
        case -292:
            fmt.Println(v3)
        default:
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS2{F0: 420}
                v19 := GenS4{F0: v1, F1: 2}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    }
    var v20 float64 = v2
    if false {
        for fi21 := 0; fi21 < 8; fi21++ {
            v22 := []int{}
            fmt.Println(v1, v4, v1)
            var v23 float64 = ((v2 - ((((float64(0.1) - (1e10 / v20)) - -2.25) / ((((v2 - (v2 * ((((-2.25 / (v2 / v2)) - (v2 - (1.5 * (1e10 * 1e10)))) / (-2.25 / 1.0)) - v2))) + 1e10) - ((v20 - (v20 + 0.1)) * (((0.0 + ((3.14159 - v20) - float64(0.1))) * (-2.25 - (0.1 + v2))) + v20))) - -2.25)) - float64((0.0 / 1e10)))) + v2)
            _ = fi21
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 8; fi24++ {
            v25 := GenS2{F0: v4}
            _ = fi24
            _ = v25
        }
        v2 -= 1.5
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        if (42 != 619) {
            for fi27 := 0; fi27 < 1; fi27++ {
                v20 = (v2 / 1e10)
                v4 = 18446744073709551615
                var v28 float32 = 0.1
                _ = fi27
                _ = v28
            }
            var v29 int32 = 1
            _ = v29
        } else {
            var v30 uint8 = 34
            _ = v30
        }
        _ = fi26
    }
    v20 = -2.25
    for fi31 := 0; fi31 < 10; fi31++ {
        if (42 == 968) {
            var v32 float32 = float32(3.14159)
            v2 -= v20
            _ = v32
        } else {
            switch fi31 {
            case 817:
                var v33 uint = (v4 >> (((v4 >> ((v4) & 63))) & 63))
                _ = v33
            case 1:
                var v34 uint32 = 0
                _ = v34
            }
            v35 := []int{-2147483648, -1, fi31, (fi31 | 42)}
            _ = v35
        }
        var v36 bool = (!(!false))
        _ = fi31
        _ = v36
    }
    var v37 bool = ((0.0 / ((-2.25 + 0.1) + 0.0)) < 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
}
