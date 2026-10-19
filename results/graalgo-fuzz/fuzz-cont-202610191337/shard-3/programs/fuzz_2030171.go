package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint64; F2 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16; F1 bool; F2 int64 }
func genFn0(p0 bool) int16 {
    for fi5 := 0; fi5 < 10; fi5++ {
        fmt.Println(fi5, fi5)
        p0 = p0
        _ = fi5
    }
    fmt.Println(p0)
    _ = p0
    return 1000
}

func main() {
    var v1 uint16 = 518
    var v2 float32 = -2.25
    v3 := []int{}
    var v4 uint32 = 910
    if false {
        v4 = (v4 | v4)
        if (((0.0 - 1e10) - 1.5) >= 3.14159) {
            var v6 int8 = 0
            var v7 float32 = v2
            switch "foo bar" {
            case "hello":
                var v8 int64 = 42
                _ = v8
            case "a":
                var v9 uint32 = (v4 + v4)
                _ = v9
            case "foo bar":
                var v10 int64 = -1
                _ = v10
            }
            _ = v6
            _ = v7
        }
        v11 := GenS1{F0: 67, F1: 1}
        _ = v11
    }
    v4 = (v4 ^ (^(v4)))
    if false {
        v4 += 918
        if (!true) {
            fmt.Println(v4)
        }
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        switch fi12 {
        case -2147483648:
            v2 = (((1e10 * v2) * 1e10) + v2)
        case 32767:
            v3 = append(v3, -2147483648)
        case 766:
            fmt.Println(v2)
        default:
            var v13 int16 = 32767
            _ = v13
        }
        for fi14 := 0; fi14 < 6; fi14++ {
            if true {
                v2 -= 0.1
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
            }
            _ = fi14
        }
        var v15 uint = 2
        _ = fi12
        _ = v15
    }
    if ((("hello" + "\n\t") + (("foo bar" + "") + "foo bar")) <= "") {
        _ = len(v3)
        if false {
            v16 := GenS1{F0: 255, F1: 100}
            _ = v16
        } else {
            switch "foo bar" {
            case "hello":
                switch 960 {
                case 4294967295:
                    for fi17 := 0; fi17 < 9; fi17++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v18 := GenS2{F0: (v1 >> ((2) & 63)), F1: 1, F2: "a"}
                        v19 := GenS3{F0: 1000}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                case 0:
                    v20 := GenS4{F0: 2, F1: 9223372036854775807}
                    _ = v20
                case 9223372036854775807:
                    v21 := GenS4{F0: 2, F1: 1}
                    _ = v21
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            if false {
                v22 := GenS5{F0: ((3.14159 * ((3.14159 / 1.5) * 0.1)) == 3.14159), F1: 973}
                if ((float32(3.14159) < v2) && (32767 > (v1 << ((616) & 63)))) {
                    var v23 float64 = float64(v2)
                    _ = v23
                }
                _ = v22
            }
            _ = len(v3)
        }
    } else {
        var v24 int64 = int64(v4)
        var v25 string = ("a" + "a")
        v25 = "a"
        _ = v24
        _ = v25
    }
    var v26 bool = (!(!(506 >= 854)))
    _ = len(v3)
    var v27 int32 = -308
    var v28 uint = 0
    v4 = 493
    v29 := GenS6{F0: (v27 - (v27 / (((^(v27))) | 1))), F1: 135, F2: "a"}
    v30 := GenS1{F0: 4294967295, F1: -4}
    v31 := GenS7{F0: -9223372036854775808}
    if (v1 < (v1 % ((v1) | 1))) {
        for fi32 := 0; fi32 < 16; fi32++ {
            v33 := GenS8{F0: (v1 / (((v1 ^ (^(v1)))) | 1)), F1: true, F2: -411}
            _ = fi32
            _ = v33
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if v26 {
        var v34 int16 = -1
        _ = v34
    } else {
        v28 = (v28 << ((v28) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
