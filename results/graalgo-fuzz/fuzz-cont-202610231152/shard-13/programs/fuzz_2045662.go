package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint8) uint8 {
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: 100, F1: true}
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(fi5, fi7)
            if (!(1.5 < 3.14159)) {
                fmt.Println(p0)
            } else {
                p0 = p0
            }
            _ = fi7
        }
        v8 := []int{}
        _ = fi5
        _ = v6
        _ = v8
    }
    var v9 int64 = -1
    _ = p0
    _ = v9
    return p0
}
func genFn1(p0 int32) bool {
    var v10 int32 = (-(p0))
    if true {
        var v11 uint64 = 4294967295
        p0 = p0
        var v12 uint64 = v11
        _ = v11
        _ = v12
    } else {
        var v13 string = ("hello" + "\n\t")
        var v14 float64 = float64(-2.25)
        if false {
            if (false && (!false)) {
                v13 = v13
                v15 := GenS1{F0: -80, F1: true}
                v16 := GenS2{F0: false}
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 20; fi17++ {
                    for fi18 := 0; fi18 < 8; fi18++ {
                        v19 := []int{127, (fi18 - (fi18 << ((255) & 63))), (fi18 / ((fi17) | 1)), fi18}
                        _ = fi18
                        _ = v19
                    }
                    _ = fi17
                }
                v20 := GenS2{F0: (2 >= 128)}
                _ = v20
            }
            fmt.Println(v14)
            var v21 float64 = float64(-2.25)
            _ = v21
        } else {
            var v22 uint64 = 884
            _ = v22
        }
        _ = v13
        _ = v14
    }
    v23 := []int{1, -137}
    _ = p0
    _ = v10
    _ = v23
    return (float32(1e10) <= ((-2.25 * 0.0) + (float32(-2.25) * float32((1.5 * 0.0)))))
}

func main() {
    var v1 float64 = 0.1
    var v2 int64 = -336
    v3 := make(map[int]string)
    var v4 uint16 = 1
    v24 := GenS3{F0: (v2 ^ (v2 - (v2 / (((v2 >> ((468) & 63))) | 1)))), F1: 434}
    switch "\n\t" {
    case "hello":
        switch 128 {
        case 127:
            var v25 float64 = 1.5
            _ = v25
        case -9223372036854775808:
            v1 += 1e10
        case 166:
            if (65535 < 1) {
                v26 := []int{127, 128, -462, -1}
                _ = v26
            }
        }
    case "\n\t":
        for fi27 := 0; fi27 < 2; fi27++ {
            for fi28 := 0; fi28 < 7; fi28++ {
                v1 -= (v1 + v1)
                _ = fi28
            }
            for fi29 := 0; fi29 < 8; fi29++ {
                v2 -= (v2 / ((9223372036854775807) | 1))
                v30 := GenS1{F0: -1, F1: false}
                v2 -= v2
                v4 += v4
                _ = fi29
                _ = v30
            }
            _ = fi27
        }
    case "a":
        if true {
            fmt.Println(v2)
        } else {
            fmt.Println(v2, v1, v4)
            v2 -= (v2 / (((v2 * 9223372036854775807)) | 1))
        }
    }
    if (true && true) {
        switch 127 {
        case 127:
            if (-9223372036854775808 != v2) {
                v31 := GenS4{F0: 25, F1: 2}
                var v32 float32 = 1.5
                var v33 int16 = 639
                _ = v31
                _ = v32
                _ = v33
            }
        default:
            switch 659 {
            case 4294967295:
                fmt.Println(v4, v1)
            }
        }
        fmt.Println(v4, v4)
        v34 := []int{847, -1, 1, 32768}
        _ = v34
    }
    v1 = v1
    v35 := []int{}
    var v36 int8 = -128
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 bool = (((fi37 + fi37) >= (fi37 * fi37)) && false)
        _ = fi37
        _ = v38
    }
    var v39 bool = false
    fmt.Println(v36, v1)
    v4 -= 65535
    v40 := GenS2{F0: (!true)}
    _ = len(v35)
    switch "a" {
    case "\n\t":
        if v39 {
            var v41 int16 = 1000
            _ = v41
        }
    case "foo bar":
        v39 = v39
    }
    v42 := GenS1{F0: (v36 % ((v36) | 1)), F1: (851 != 1)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v42
}
