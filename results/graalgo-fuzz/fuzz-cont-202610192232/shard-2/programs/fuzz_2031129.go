package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 bool; F1 bool; F2 bool }
type GenS4 struct { F0 string; F1 uint8 }
func genFn0(p0 int) int32 {
    fmt.Println(p0)
    var v4 int8 = -1
    var v5 uint8 = 255
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            if (0 >= 353) {
                v5 = (v5 << ((791) & 63))
                var v8 string = "\n\t"
                var v9 uint8 = (v5 % ((v5) | 1))
                _ = v8
                _ = v9
            } else {
                var v10 int = (fi6 + fi7)
                _ = v10
            }
            switch "hello" {
            case "\n\t":
                var v11 int64 = -9223372036854775808
                if (2147483647 <= -2147483648) {
                    v12 := []int{(^(fi7))}
                    _ = v12
                }
                _ = v11
            default:
                if (!((42 >= 0) && false)) {
                    v13 := []int{(p0 << ((285) & 63)), 4294967295}
                    _ = v13
                }
            }
            _ = fi7
        }
        _ = fi6
    }
    _ = p0
    _ = v4
    _ = v5
    return -947
}
func genFn1(p0 uint64) uint8 {
    p0 -= (^(p0))
    for fi14 := 0; fi14 < 18; fi14++ {
        p0 = 1
        _ = fi14
    }
    var v15 string = ""
    p0 += p0
    var v16 float64 = 0.0
    _ = p0
    _ = v15
    _ = v16
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 4294967295
    var v3 int = v2
    v2 = v2
    v3 += (^(v2))
    if (1 >= 4294967295) {
        var v17 uint16 = 343
        switch (v3 % ((v3) | 1)) {
        case 128:
            for fi18 := 0; fi18 < 7; fi18++ {
                switch "foo bar" {
                case "\n\t":
                    var v19 float32 = (1.5 * (0.0 - 3.14159))
                    _ = v19
                case "foo bar":
                    v20 := GenS1{F0: 262, F1: 2, F2: -1}
                    _ = v20
                case "":
                    fmt.Println(v3, v2)
                }
                _ = fi18
            }
        case 32768:
            v21 := GenS2{F0: "foo bar", F1: -880}
            _ = v21
        case 4294967295:
            var v22 uint16 = v17
            _ = v22
        default:
            for fi23 := 0; fi23 < 4; fi23++ {
                for fi24 := 0; fi24 < 18; fi24++ {
                    fmt.Println(v2)
                    var v25 uint16 = (v17 % ((v17) | 1))
                    var v26 uint64 = 471
                    var v27 uint8 = 72
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                var v28 bool = ((true && true) || false)
                _ = fi23
                _ = v28
            }
        }
        _ = v17
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 float64 = 1e10
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 uint16 = 89
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    v33 := GenS3{F0: false, F1: (3.14159 != 0.1), F2: false}
    v3 = (v3 | (v3 >> ((4294967295) & 63)))
    v34 := GenS4{F0: ("a" + "foo bar"), F1: 128}
    fmt.Println(v2, v3)
    var v35 float32 = 1.5
    var v36 uint8 = 88
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
