package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint; F2 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 uint32; F1 bool }
func genFn0(p0 bool) uint64 {
    if p0 {
        v7 := GenS1{F0: 96}
        p0 = p0
        _ = v7
    }
    switch "" {
    case "hello":
        switch "\n\t" {
        case "foo bar":
            if (94 != -55) {
                p0 = (1 > 34)
                v8 := []int{-9223372036854775808}
                for fi9 := 0; fi9 < 5; fi9++ {
                    v10 := GenS1{F0: 792}
                    var v11 uint16 = 851
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = v8
            } else {
                v12 := GenS1{F0: -875}
                var v13 uint = 806
                switch "" {
                case "a":
                    var v14 int16 = -32768
                    _ = v14
                default:
                    v13 = v13
                }
                _ = v12
                _ = v13
            }
        default:
            p0 = (true || true)
        }
    }
    if (32767 == 127) {
        p0 = (!((1 < 34) || (!p0)))
    }
    if (42 < 65535) {
        fmt.Println(p0)
    } else {
        if (4294967295 != 2147483647) {
            if false {
                fmt.Println(p0)
                v15 := []int{9223372036854775807, -110, -2147483648, -837}
                p0 = (731 >= 0)
                _ = v15
            } else {
                if false {
                    switch "hello" {
                    case "\n\t":
                        v16 := GenS1{F0: 0}
                        _ = v16
                    case "":
                        var v17 uint16 = 2
                        _ = v17
                    }
                }
            }
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 int8 = 42
                v20 := []int{}
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 float32 = 3.14159
            _ = v21
        } else {
            switch -912 {
            case 48:
                v22 := GenS2{F0: 42, F1: 675, F2: 42}
                _ = v22
            case -39:
                var v23 uint64 = 4294967295
                _ = v23
            case 9223372036854775807:
                v24 := GenS2{F0: 820, F1: 911, F2: 0}
                _ = v24
            default:
                var v25 uint16 = 255
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            if p0 {
                for fi27 := 0; fi27 < 2; fi27++ {
                    var v28 int = 42
                    var v29 int64 = 725
                    v29 = v29
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                fmt.Println(p0, fi30, fi26)
                var v31 uint16 = 32767
                v32 := []int{0}
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi26
        }
    }
    fmt.Println(p0)
    _ = p0
    return 858
}
func genFn1(p0 uint64, p1 uint, p2 int8) uint {
    if false {
        p1 = (p1 ^ p1)
        v33 := GenS1{F0: 1}
        _ = v33
    }
    if false {
        fmt.Println(p1, p0)
        fmt.Println(p0, p1)
    } else {
        var v34 int32 = 42
        var v35 int8 = (^(p2))
        v35 = 127
        _ = v34
        _ = v35
    }
    p0 = (p0 % ((2) | 1))
    _ = p0
    _ = p1
    _ = p2
    return (^(p1))
}

func main() {
    var v1 int = 32767
    var v2 int16 = -370
    var v3 int32 = 0
    var v4 float32 = 0.0
    var v5 uint64 = 211
    var v6 bool = true
    var v36 bool = (v4 != v4)
    if true {
        for fi37 := 0; fi37 < 4; fi37++ {
            if ((!false) || (127 >= 100)) {
                var v38 uint16 = 2
                v39 := GenS3{F0: false}
                _ = v38
                _ = v39
            }
            v40 := GenS1{F0: 0}
            _ = fi37
            _ = v40
        }
    } else {
        v36 = v36
        v41 := GenS1{F0: 474}
        v1 = 128
        _ = v41
    }
    v42 := GenS4{F0: (v1 + v1)}
    var v43 int32 = (v3 / (((v3 - (v3 << ((42) & 63)))) | 1))
    var v44 uint64 = (v5 << ((genFn1(1, 356, 100)) & 63))
    var v45 string = "\n\t"
    for fi46 := 0; fi46 < 2; fi46++ {
        if (4294967295 < 65535) {
            v43 += v3
            switch (fi46 << ((896) & 63)) {
            case -9223372036854775808:
                v47 := GenS5{F0: (^(v43)), F1: 4294967295}
                _ = v47
            default:
                v48 := GenS6{F0: 1, F1: v6}
                _ = v48
            }
            if (((v45 + v45) + "\n\t") != ((("hello" + (("foo bar" + ("a" + ((v45 + v45) + v45))) + v45)) + "foo bar") + v45)) {
                v4 = -2.25
            } else {
                for fi49 := 0; fi49 < 14; fi49++ {
                    v50 := []int{(-(fi46)), fi46, -1, v1}
                    _ = fi49
                    _ = v50
                }
            }
        }
        v43 = (-(v43))
        _ = fi46
    }
    v51 := GenS6{F0: 158, F1: v36}
    var v52 int64 = -9223372036854775808
    var v53 float64 = -2.25
    v6 = v6
    v54 := []int{v1, (v1 + (v1 - 128)), v1, v1}
    for fi55 := 0; fi55 < 20; fi55++ {
        v2 -= (v2 ^ v2)
        v6 = (255 < 434)
        var v56 uint8 = 1
        _ = fi55
        _ = v56
    }
    var v57 uint8 = 2
    switch -2147483648 {
    case 32767:
        for fi58 := 0; fi58 < 13; fi58++ {
            v1 -= int(v2)
            _ = fi58
        }
    case -250:
        v6 = v6
    default:
        v4 = (0.0 * v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v57
    _ = v6
}
