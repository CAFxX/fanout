package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int32; F2 int16 }
func genFn0() uint16 {
    v7 := []int{4294967295, 128}
    var v8 int64 = 1
    v9 := GenS1{F0: 735, F1: -194, F2: 127}
    var v10 uint32 = 65535
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    var v3 uint16 = 1
    var v4 float64 = 0.0
    var v5 int = 2147483647
    var v6 int = 419
    switch 301 {
    case 1:
        switch v6 {
        case 736:
            v11 := []int{}
            _ = v11
        case -9223372036854775808:
            if ((v6 ^ (-(v6))) >= 127) {
                delete(v2, v5)
                v3 = 32767
                fmt.Println(v5, v6, v3)
            }
        default:
            v12 := []int{-2147483648, (v6 % ((42) | 1)), 42, (v6 | v6)}
            _ = v12
        }
    case -1:
        fmt.Println(v5, v4, v4)
    }
    var v13 uint16 = 32767
    v3 -= 969
    var v14 int = v6
    switch v5 {
    case -17:
        if (!false) {
            v15 := v1[v6]
            var v16 int8 = 1
            switch ("a" + "foo bar") {
            case "\n\t":
                for fi17 := 0; fi17 < 19; fi17++ {
                    v1[(v5 + (^(v14)))] = (-32768 >= 0)
                    if (-1 != -1) {
                        fmt.Println(fi17, v4, v4)
                    }
                    v18 := []int{(^(fi17)), 32768}
                    _ = fi17
                    _ = v18
                }
            }
            _ = v15
            _ = v16
        }
    case 958:
        switch 4294967295 {
        case 128:
            switch (v6 ^ 9223372036854775807) {
            case 128:
                var v19 uint16 = genFn0()
                _ = v19
            case -2147483648:
                switch v5 {
                case 42:
                    v2[v14] = 0
                case 9223372036854775807:
                    for fi20 := 0; fi20 < 20; fi20++ {
                        _ = len(v1)
                        v6 += -580
                        _ = fi20
                    }
                case 4294967295:
                    v13 += (^(v13))
                }
            }
        case -836:
            v3 += (v3 - genFn0())
        case 32767:
            v6 -= (v5 & 1)
        }
    case 827:
        if false {
            v21 := GenS1{F0: 1, F1: v14, F2: 53}
            var v22 float32 = float32(v4)
            var v23 int16 = 791
            _ = v21
            _ = v22
            _ = v23
        } else {
            v24 := v1[-819]
            v5 = (v14 - 128)
            _ = v24
        }
    default:
        if (32767 <= 42) {
            if true {
                v25 := GenS1{F0: 42, F1: -747, F2: 40}
                for fi26 := 0; fi26 < 3; fi26++ {
                    v27 := GenS1{F0: 18446744073709551615, F1: (v6 & (fi26 * (v6 + v6))), F2: 1}
                    var v28 string = ""
                    v29 := v2[v5]
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                var v30 int8 = 23
                _ = v25
                _ = v30
            } else {
                v31 := GenS2{F0: v5}
                fmt.Println(v6, v6)
                fmt.Println(v6, v3)
                _ = v31
            }
            var v32 uint32 = 0
            _ = v32
        }
    }
    if false {
        v3 = v13
        var v33 float64 = float64(0.1)
        for fi34 := 0; fi34 < 8; fi34++ {
            fmt.Println(fi34, v13, v6)
            _ = fi34
        }
        _ = v33
    } else {
        for fi35 := 0; fi35 < 17; fi35++ {
            var v36 string = ""
            for fi37 := 0; fi37 < 19; fi37++ {
                v14 -= v14
                v14 = -914
                _ = fi37
            }
            _ = fi35
            _ = v36
        }
        v5 = v14
    }
    v14 += (v14 * v14)
    var v38 uint16 = v13
    fmt.Println(v38, v6, v6)
    fmt.Println(v3, v4, v14)
    if false {
        v39 := GenS3{F0: genFn0(), F1: 0, F2: -278}
        fmt.Println(v5, v38, v38)
        fmt.Println(v3)
        _ = v39
    } else {
        switch "" {
        case "":
            switch v14 {
            case 127:
                _ = len(v1)
            }
        case "\n\t":
            if ((-(v6)) > (v5 ^ v5)) {
                var v40 int64 = -1
                var v41 string = ""
                _ = v40
                _ = v41
            }
        }
        fmt.Println(v38)
    }
    v42 := []int{}
    if len(v42) > 0 {
        v42[0] = -972
    }
    if len(v42) > 0 {
        v42[0] = (v6 ^ (v14 * (v14 / ((v5) | 1))))
    }
    var v43 uint64 = 856
    var v44 int = v14
    var v45 float64 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
