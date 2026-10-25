package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint32, p1 int8) int16 {
    switch -2147483648 {
    case 0:
        var v4 uint8 = 255
        _ = v4
    case -4:
        p0 += 2
    }
    if true {
        if false {
            switch 32767 {
            case 4294967295:
                fmt.Println(p0, p0)
            }
            v5 := []int{}
            _ = v5
        }
        for fi6 := 0; fi6 < 8; fi6++ {
            if true {
                p1 = p1
            }
            _ = fi6
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint8 = 2
    var v2 int = 0
    var v3 int16 = -883
    var v7 float64 = 3.14159
    if (((float32(0.1) - 3.14159) * -2.25) <= 0.1) {
        var v8 uint32 = 2
        for fi9 := 0; fi9 < 5; fi9++ {
            if true {
                v1 += (^(v1))
            } else {
                v1 = 255
                v10 := GenS1{F0: 255, F1: (v8 + v8), F2: (v2 != fi9)}
                v11 := GenS1{F0: 0, F1: (v8 >> ((65535) & 63)), F2: (0 == 127)}
                _ = v10
                _ = v11
            }
            v12 := GenS1{F0: 42, F1: 2, F2: true}
            switch "\n\t" {
            case "":
                v13 := []int{v2, 115, 2147483647}
                _ = v13
            case "a":
                fmt.Println(v8, v2)
            case "hello":
                v14 := GenS1{F0: 355, F1: (^(v8)), F2: true}
                _ = v14
            default:
                v15 := GenS2{F0: (42 <= 18446744073709551615)}
                _ = v15
            }
            _ = fi9
            _ = v12
        }
        _ = v8
    } else {
        var v16 int64 = 42
        if (0 < 65535) {
            for fi17 := 0; fi17 < 17; fi17++ {
                v18 := GenS1{F0: 860, F1: 458, F2: true}
                v19 := GenS1{F0: 42, F1: 2147483648, F2: false}
                _ = fi17
                _ = v18
                _ = v19
            }
            switch int(v1) {
            case -597:
                var v20 float32 = float32(v7)
                _ = v20
            }
        }
        _ = v16
    }
    fmt.Println(v2)
    var v21 int64 = 0
    if (!false) {
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := GenS2{F0: false}
            for fi24 := 0; fi24 < 5; fi24++ {
                switch 32768 {
                case 0:
                    v2 += (fi22 - -641)
                }
                v25 := []int{fi22, v2, fi22}
                if true {
                    switch (fi22 >> ((4294967295) & 63)) {
                    case 2147483647:
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                    case -925:
                        v2 = fi24
                        if (float32(v7) > float32(0.0)) {
                            _ = len(v25)
                            switch "\n\t" {
                            case "hello":
                                v1 = 198
                            }
                        }
                    case 127:
                        v26 := GenS2{F0: (0 != -82)}
                        var v27 int32 = 42
                        var v28 int64 = 42
                        v29 := GenS3{F0: v1, F1: 1}
                        v28 = -1
                        _ = v26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                } else {
                    _ = len(v25)
                    _ = len(v25)
                    v1 -= 128
                }
                _ = fi24
                _ = v25
            }
            _ = fi22
            _ = v23
        }
        switch v2 {
        case -1:
            for fi30 := 0; fi30 < 19; fi30++ {
                var v31 float64 = -2.25
                _ = fi30
                _ = v31
            }
        }
    } else {
        v32 := GenS4{F0: 0, F1: true, F2: 1}
        for fi33 := 0; fi33 < 8; fi33++ {
            fmt.Println(v2, v1, v2)
            for fi34 := 0; fi34 < 13; fi34++ {
                fmt.Println(v21, v1)
                _ = fi34
            }
            var v35 uint8 = v1
            _ = fi33
            _ = v35
        }
        switch (^(v2)) {
        case 9223372036854775807:
            var v36 uint32 = 37
            _ = v36
        case 32768:
            if true {
                for fi37 := 0; fi37 < 2; fi37++ {
                    v2 += v2
                    fmt.Println(v7)
                    _ = fi37
                }
            }
        default:
            fmt.Println(v3)
        }
        _ = v32
    }
    var v38 uint16 = 2
    v39 := []int{(v2 / (((v2 & 1)) | 1))}
    for fi40 := 0; fi40 < 15; fi40++ {
        if false {
            var v41 uint16 = (v38 / ((v38) | 1))
            _ = v41
        }
        v42 := GenS5{F0: (v3 + v3)}
        v38 = v38
        _ = fi40
        _ = v42
    }
    switch int(v3) {
    case -1:
        v2 = (v2 / ((128) | 1))
    default:
        v2 = 541
    }
    var v43 int8 = 0
    fmt.Println(v1, v43, v2)
    v44 := GenS6{F0: 327}
    v45 := GenS1{F0: v38, F1: 65535, F2: true}
    v46 := GenS1{F0: uint16(v1), F1: 65535, F2: true}
    v1 -= (v1 - (v1 & v1))
    switch "\n\t" {
    case "a":
        _ = len(v39)
    default:
        if len(v39) > 0 {
            v39[0] = 874
        }
    }
    fmt.Println(v2, v43, v38)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v43:", v43)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v38
    _ = v39
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v7
}
