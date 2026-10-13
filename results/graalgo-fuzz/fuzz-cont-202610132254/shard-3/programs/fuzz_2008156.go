package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint64; F1 uint16; F2 bool }
func genFn0(p0 uint16, p1 string) int8 {
    if (255 <= 45) {
        p0 = (p0 ^ (^(p0)))
    } else {
        var v4 uint32 = 42
        if (!(!(100 >= 254))) {
            fmt.Println(p1, p0, v4)
        }
        v4 += (v4 << ((0) & 63))
        _ = v4
    }
    p0 += (p0 / (((^(p0))) | 1))
    if (!((p0 >> ((18446744073709551615) & 63)) <= (p0 << ((18446744073709551615) & 63)))) {
        v5 := []int{}
        if false {
            fmt.Println(p1)
            fmt.Println(p1)
        } else {
            var v6 uint8 = 255
            _ = v6
        }
        _ = v5
    } else {
        if (((!(-118 < 0)) || (0.1 <= -2.25)) || (1 == (p0 % (((p0 - p0)) | 1)))) {
            v7 := GenS1{F0: -9223372036854775808, F1: 65535}
            fmt.Println(p0)
            var v8 uint8 = 0
            _ = v7
            _ = v8
        } else {
            var v9 float32 = (float32(-2.25) - (1e10 * (1.5 + ((float32(0.0) * float32(1e10)) - ((1e10 - float32((0.1 / (1.5 * (1.5 / ((0.0 + 0.1) / 1.0)))))) + 1e10)))))
            _ = v9
        }
        if (-9223372036854775808 > 9223372036854775807) {
            var v10 uint32 = 774
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: fi11, F1: 255}
        _ = fi11
        _ = v12
    }
    var v13 int64 = 0
    _ = p0
    _ = p1
    _ = v13
    return -82
}

func main() {
    var v1 int = 0
    var v2 uint64 = 278
    var v3 int16 = 1
    v14 := GenS2{F0: 255, F1: 1, F2: 495}
    if ((1 <= -1) || false) {
        var v15 uint64 = v2
        switch (v1 | (v1 + v1)) {
        case 128:
            fmt.Println(v15, v3, v2)
        case 32767:
            v2 -= (^(v15))
        }
        _ = v15
    }
    v16 := GenS1{F0: -2147483648, F1: 1}
    switch ("\n\t" + "foo bar") {
    case "a":
        if false {
            v17 := []int{(v1 + 9223372036854775807)}
            _ = v17
        }
    case "foo bar":
        if ((false && false) || true) {
            v18 := GenS1{F0: (v1 / ((v1) | 1)), F1: 2}
            _ = v18
        }
    default:
        var v19 uint64 = (v2 | v2)
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        if false {
            switch "hello" {
            case "a":
                v21 := []int{(v1 << ((101) & 63))}
                _ = v21
            case "hello":
                if (0 < 872) {
                    v2 = 2
                    fmt.Println(v3, fi20, v2)
                } else {
                    var v22 int64 = 1
                    _ = v22
                }
            }
        } else {
            if (false && true) {
                var v23 bool = (false && true)
                v23 = false
                v24 := []int{}
                _ = v23
                _ = v24
            } else {
                v1 += -1
                fmt.Println(v3, v3, v1)
            }
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 uint32 = 2
            _ = fi25
            _ = v26
        }
        _ = fi20
    }
    var v27 float32 = (3.14159 + float32(float64(1.5)))
    v28 := GenS1{F0: v1, F1: 0}
    if (((255 < 162) || false) || (-1 <= 2147483647)) {
        v1 = v1
    }
    if (-128 != 0) {
        switch v1 {
        case 32767:
            var v29 uint32 = 0
            _ = v29
        case -245:
            switch (v1 % ((v1) | 1)) {
            case 0:
                v27 = (-2.25 / v27)
            case -981:
                if false {
                    var v30 uint16 = 32767
                    _ = v30
                }
            }
        case 42:
            if true {
                v31 := GenS3{F0: (v2 - (v2 >> ((255) & 63))), F1: -28}
                _ = v31
            }
        }
        v32 := GenS4{F0: v2, F1: 32767, F2: false}
        v1 += (v1 + v1)
        _ = v32
    } else {
        for fi33 := 0; fi33 < 9; fi33++ {
            v3 = v3
            _ = fi33
        }
        for fi34 := 0; fi34 < 13; fi34++ {
            v3 += v3
            for fi35 := 0; fi35 < 9; fi35++ {
                v2 += (^(v2))
                _ = fi35
            }
            switch v1 {
            case 4294967295:
                fmt.Println(v1, v1, v27)
            }
            _ = fi34
        }
        v36 := []int{v1, v1, (^(v1)), (-(v1))}
        _ = v36
    }
    var v37 int16 = v3
    switch v1 {
    case 843:
        for fi38 := 0; fi38 < 14; fi38++ {
            var v39 int = (fi38 * (v1 % ((fi38) | 1)))
            _ = fi38
            _ = v39
        }
    }
    var v40 string = (("\n\t" + "") + "a")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v40
}
