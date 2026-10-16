package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int64; F2 bool }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 bool }
func genFn0() uint16 {
    if false {
        switch "" {
        case "a":
            if false {
                switch 1 {
                case 32768:
                    var v4 string = "hello"
                    _ = v4
                default:
                    for fi5 := 0; fi5 < 2; fi5++ {
                        switch -1 {
                        case 32767:
                            switch -1 {
                            case 785:
                                var v6 bool = (("\n\t" + "\n\t") != "hello")
                                _ = v6
                            default:
                                v7 := GenS1{F0: 2147483647, F1: -927}
                                _ = v7
                            }
                        }
                        _ = fi5
                    }
                }
            } else {
                switch ("a" + "foo bar") {
                case "foo bar":
                    var v8 bool = false
                    _ = v8
                case "\n\t":
                    var v9 int64 = 1
                    _ = v9
                default:
                    switch -1 {
                    case -433:
                        v10 := GenS2{F0: 1}
                        _ = v10
                    default:
                        var v11 int32 = -106
                        _ = v11
                    }
                }
                v12 := []int{}
                _ = v12
            }
        case "foo bar":
            var v13 uint64 = 770
            _ = v13
        }
        if ((true || false) || true) {
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 uint64 = 42
                var v16 bool = (false && true)
                v17 := GenS2{F0: 0}
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        } else {
            var v18 float64 = (0.1 * -2.25)
            _ = v18
        }
    }
    v19 := GenS1{F0: -95, F1: 42}
    for fi20 := 0; fi20 < 4; fi20++ {
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 bool = (false || (!false))
            v22 = (1.5 < (0.0 * 1.5))
            var v23 int32 = -2147483648
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    _ = v19
    return 255
}

func main() {
    var v1 int8 = -115
    var v2 float32 = -2.25
    var v3 uint = 286
    fmt.Println(v1, v3, v2)
    for fi24 := 0; fi24 < 17; fi24++ {
        v3 += (v3 - (v3 % ((v3) | 1)))
        var v25 bool = (!true)
        v3 = v3
        _ = fi24
        _ = v25
    }
    fmt.Println(v1, v3)
    if (!true) {
        var v26 int16 = -32768
        fmt.Println(v26)
        for fi27 := 0; fi27 < 9; fi27++ {
            switch -1 {
            case 32767:
                switch (fi27 >> (((v3 * (v3 | 761))) & 63)) {
                case 2147483647:
                    if ((fi27 ^ (fi27 ^ (fi27 + 32768))) < (fi27 & (fi27 ^ fi27))) {
                        v28 := []int{(fi27 % ((9223372036854775807) | 1)), fi27, (fi27 & fi27), -783}
                        _ = v28
                    }
                default:
                    v29 := []int{(fi27 % ((int(v1)) | 1)), fi27, fi27, fi27}
                    _ = v29
                }
            case -1:
                var v30 int = (fi27 & (fi27 ^ fi27))
                _ = v30
            }
            var v31 uint64 = 42
            if true {
                v32 := GenS1{F0: 2147483647, F1: int32(fi27)}
                _ = v32
            }
            _ = fi27
            _ = v31
        }
        _ = v26
    }
    fmt.Println(v2)
    switch -137 {
    case -611:
        v33 := GenS1{F0: -123, F1: 2147483647}
        _ = v33
    case 809:
        v34 := GenS2{F0: 0}
        _ = v34
    }
    var v35 float64 = 0.0
    v36 := GenS3{F0: 267, F1: 2}
    if (-2147483648 < 1) {
        var v37 int8 = 86
        v35 += ((v35 / v35) / 0.1)
        v38 := []int{9223372036854775807, 1}
        _ = v37
        _ = v38
    }
    v3 += v3
    for fi39 := 0; fi39 < 5; fi39++ {
        v40 := GenS4{F0: 128, F1: -9223372036854775808, F2: (2147483647 >= 2147483648)}
        for fi41 := 0; fi41 < 5; fi41++ {
            var v42 int = 781
            _ = fi41
            _ = v42
        }
        _ = fi39
        _ = v40
    }
    if (32767 == 255) {
        for fi43 := 0; fi43 < 1; fi43++ {
            switch int(v1) {
            case -597:
                v44 := []int{fi43, 32767, (fi43 * -9223372036854775808)}
                _ = v44
            }
            var v45 int32 = 0
            v2 -= v2
            _ = fi43
            _ = v45
        }
        for fi46 := 0; fi46 < 13; fi46++ {
            var v47 uint64 = 263
            _ = fi46
            _ = v47
        }
        v35 -= (((-2.25 + float64((float32((v35 * v35)) / v2))) + v35) * float64(float32((0.1 - -2.25))))
    } else {
        var v48 int8 = (^(v1))
        _ = v48
    }
    v1 = (v1 * (v1 >> (((v3 % ((65535) | 1))) & 63)))
    for fi49 := 0; fi49 < 10; fi49++ {
        if true {
            v50 := GenS5{F0: 42, F1: ""}
            _ = v50
        } else {
            v51 := GenS6{F0: (((false && true) && true) && (fi49 > (fi49 % ((fi49) | 1))))}
            fmt.Println(v35, v35)
            _ = v51
        }
        v1 = (v1 << (((v3 | (v3 & v3))) & 63))
        var v52 float64 = v35
        _ = fi49
        _ = v52
    }
    v53 := []int{}
    var v54 uint16 = uint16(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v53
    _ = v54
}
