package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint8; F1 string; F2 string }
func genFn0() uint64 {
    var v5 int = -1
    if false {
        v6 := []int{v5, 0, 0, v5}
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint8 = 255
            var v9 uint8 = (v8 >> ((255) & 63))
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    }
    _ = v5
    return 801
}
func genFn1() int {
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := []int{}
        var v12 uint16 = 1
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{-1}
    var v14 bool = true
    if (true || v14) {
        for fi15 := 0; fi15 < 10; fi15++ {
            v14 = (!v14)
            _ = len(v13)
            _ = fi15
        }
        var v16 uint32 = 148
        _ = v16
    } else {
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 bool = false
            fmt.Println(v18, v14, v18)
            _ = fi17
            _ = v18
        }
        v19 := GenS1{F0: 42, F1: 2147483647}
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 float32 = 0.1
            var v23 int8 = 100
            v24 := GenS2{F0: fi20, F1: 1}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        if (0 <= -9223372036854775808) {
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 string = "hello"
                if ((308 == (fi20 << ((255) & 63))) || false) {
                    v27 := GenS2{F0: fi20, F1: 468}
                    _ = v27
                }
                var v28 int8 = -128
                _ = fi25
                _ = v26
                _ = v28
            }
            if len(v13) > 0 {
                v13[0] = fi20
            }
        }
        _ = fi20
    }
    _ = v13
    _ = v14
    return 208
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 0
    v3 := []string{}
    var v4 float32 = v1
    v1 -= v1
    if ((!(true || (!false))) && ((v2 - (v2 % ((42) | 1))) > 925)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v29 bool = false
        switch -9223372036854775808 {
        case 32767:
            var v30 uint32 = 65535
            _ = v30
        case -967:
            for fi31 := 0; fi31 < 12; fi31++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi31
            }
        }
        _ = v29
    }
    if false {
        var v32 uint64 = 267
        _ = v32
    } else {
        if false {
            if len(v3) > 0 {
                v3[0] = "a"
            }
            v4 = -2.25
        }
    }
    switch -542 {
    case 127:
        var v33 uint = (^(v2))
        _ = v33
    case 443:
        v2 = v2
    case 1:
        v34 := GenS3{F0: 42, F1: ("" + "a"), F2: "hello"}
        _ = v34
    }
    _ = len(v3)
    if (-394 > 1) {
        if false {
            switch -9223372036854775808 {
            case 32768:
                v1 = (v1 + v1)
            default:
                var v35 uint16 = 255
                _ = v35
            }
            var v36 uint = v2
            _ = v36
        }
    }
    var v37 uint8 = 199
    var v38 int16 = 0
    v39 := GenS3{F0: (v37 + (v37 + 2)), F1: "", F2: "a"}
    v4 += 1.5
    v4 += ((v4 * (v1 / v1)) / (v1 / (((v4 * (float32((0.1 / (-2.25 - float64((v4 / ((0.1 * v4) + 3.14159)))))) * float32(((1.5 - float64(v1)) * 3.14159)))) + (3.14159 / v4)) + ((v1 * (v1 * float32((3.14159 - float64(3.14159))))) - v1))))
    v37 = v37
    var v40 uint64 = 854
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
