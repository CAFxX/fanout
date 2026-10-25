package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int; F2 uint64 }
type GenS3 struct { F0 int64; F1 string; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint32; F2 bool }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 bool }
func genFn0() int {
    if false {
        v4 := []int{4294967295, 0, 127}
        for fi5 := 0; fi5 < 13; fi5++ {
            v6 := GenS1{F0: (fi5 | fi5), F1: 966}
            var v7 int64 = 1
            if false {
                var v8 float32 = 1e10
                v8 = float32(-2.25)
                _ = v8
            } else {
                v7 += -120
                var v9 int64 = 1
                v7 += v7
                _ = v9
            }
            _ = fi5
            _ = v6
            _ = v7
        }
        var v10 uint16 = 837
        _ = v10
        _ = v4
    }
    var v11 bool = true
    var v12 int32 = -185
    if v11 {
        v11 = true
    }
    v13 := GenS1{F0: 4294967295, F1: 575}
    _ = v11
    _ = v12
    _ = v13
    return 809
}

func main() {
    var v1 uint8 = 86
    v2 := make(map[string]string)
    var v3 uint = 1
    fmt.Println(v3, v3)
    if false {
        switch "hello" {
        case "\n\t":
            switch 4294967295 {
            case -1:
                fmt.Println(v1, v1)
            default:
                v14 := GenS1{F0: 0, F1: 18446744073709551615}
                _ = v14
            }
        default:
            if true {
                var v15 int64 = 1
                v16 := []int{42, genFn0(), 422, 0}
                _ = v15
                _ = v16
            }
        }
        if (true || (-1 == 327)) {
            var v17 int16 = 32767
            var v18 int8 = -1
            _ = v17
            _ = v18
        } else {
            var v19 uint16 = 32767
            var v20 uint = v3
            _ = v19
            _ = v20
        }
    } else {
        fmt.Println(v3, v3)
    }
    var v21 uint = (v3 & 65535)
    switch 2147483647 {
    case -9223372036854775808:
        var v22 float64 = 0.1
        _ = v22
    case 128:
        if (42 != 42) {
            v1 = (^(v1))
        } else {
            switch ("foo bar" + "a") {
            case "hello":
                v23 := GenS1{F0: 42, F1: 756}
                _ = v23
            }
            var v24 uint16 = 173
            _ = v24
        }
    case 9223372036854775807:
        var v25 uint32 = 676
        _ = v25
    default:
        v3 = 65535
    }
    if ((!true) && (1 >= 65535)) {
        var v26 int16 = 1000
        if true {
            for fi27 := 0; fi27 < 8; fi27++ {
                v28 := []int{fi27, fi27}
                for fi29 := 0; fi29 < 5; fi29++ {
                    v30 := GenS2{F0: 1, F1: fi27, F2: 1}
                    if (!(6 != 1)) {
                        var v31 int16 = 32767
                        var v32 uint16 = 2
                        _ = v31
                        _ = v32
                    }
                    _ = fi29
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
            v33 := GenS3{F0: 42, F1: "foo bar", F2: 32}
            v34 := GenS4{F0: v21, F1: 478, F2: false}
            _ = v33
            _ = v34
        } else {
            if false {
                if false {
                    for fi35 := 0; fi35 < 2; fi35++ {
                        var v36 string = "a"
                        _ = fi35
                        _ = v36
                    }
                } else {
                    if ((v26 % ((32767) | 1)) > (-(v26))) {
                        var v37 int32 = 267
                        fmt.Println(v1, v21)
                        _ = v37
                    }
                    v3 = v21
                }
                var v38 float64 = 0.1
                _ = v38
            }
            _ = len(v2)
        }
        _ = v26
    } else {
        v39 := GenS5{F0: -32768, F1: (-9223372036854775808 >= 9223372036854775807)}
        _ = v39
    }
    var v40 uint64 = 2
    v21 -= v21
    var v41 uint = (v21 % (((^(v21))) | 1))
    for fi42 := 0; fi42 < 8; fi42++ {
        v43 := []int{genFn0(), 0}
        var v44 int32 = 1
        _ = fi42
        _ = v43
        _ = v44
    }
    if ((true && true) && (944 <= 42)) {
        v45 := GenS6{F0: true}
        _ = v45
    }
    fmt.Println(v21, v1)
    for fi46 := 0; fi46 < 8; fi46++ {
        fmt.Println(v21, v41, v1)
        if false {
            switch "" {
            case "\n\t":
                v21 = v41
            case "hello":
                var v47 int16 = -1
                _ = v47
            default:
                fmt.Println(v41)
            }
            var v48 uint = (v3 << (((v41 - v21)) & 63))
            fmt.Println(v21, v48, v40)
            _ = v48
        }
        _ = fi46
    }
    for fi49 := 0; fi49 < 6; fi49++ {
        v40 -= 1
        _ = fi49
    }
    var v50 int64 = -1
    v1 += 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v40
    _ = v41
    _ = v50
}
