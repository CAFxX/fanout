package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int64; F1 bool; F2 uint32 }
func genFn0(p0 int) uint8 {
    if false {
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := GenS1{F0: (!(true && false)), F1: 802, F2: 127}
                p0 += (fi7 / (((-(fi8))) | 1))
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        v10 := GenS1{F0: (!false), F1: 0, F2: 127}
        p0 -= (-(p0))
        _ = v10
    }
    v11 := GenS1{F0: false, F1: 495, F2: 1}
    v12 := GenS2{F0: 2147483647}
    if true {
        v13 := []int{4294967295, p0}
        if (8 != 2) {
            p0 = p0
            fmt.Println(p0)
            fmt.Println(p0)
        } else {
            fmt.Println(p0)
            if false {
                switch p0 {
                case 32768:
                    for fi14 := 0; fi14 < 11; fi14++ {
                        p0 -= fi14
                        p0 -= 127
                        _ = fi14
                    }
                }
                v15 := GenS2{F0: 2}
                _ = v15
            }
            var v16 float64 = 3.14159
            _ = v16
        }
        var v17 int16 = 1
        _ = v13
        _ = v17
    }
    _ = p0
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 float64 = (1e10 * float64(3.14159))
    var v2 uint64 = 0
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    v5 := make(map[int]uint32)
    var v6 int8 = 100
    var v18 float32 = (0.1 * -2.25)
    for fi19 := 0; fi19 < 18; fi19++ {
        delete(v3, fi19)
        v20 := []int{fi19, fi19}
        _ = fi19
        _ = v20
    }
    v6 = int8(v2)
    if (false || (0 <= 42)) {
        v21 := GenS1{F0: false, F1: (v2 / ((42) | 1)), F2: 255}
        v22 := []int{}
        fmt.Println(v2, v1, v6)
        _ = v21
        _ = v22
    } else {
        fmt.Println(v1, v6)
        v6 -= v6
    }
    if (0 <= genFn0(9223372036854775807)) {
        v23 := GenS3{F0: (v2 >> ((0) & 63)), F1: "a"}
        delete(v3, 32767)
        _ = v23
    } else {
        v2 = 2
        v1 = ((0.1 / v1) * (v1 - v1))
    }
    switch -2147483648 {
    case 128:
        v6 -= (^(v6))
    case 127:
        for fi24 := 0; fi24 < 18; fi24++ {
            switch (fi24 - (fi24 % ((fi24) | 1))) {
            case 4294967295:
                v2 = 1
            case 42:
                v2 = v2
            }
            v1 = (v1 * (1.5 - v1))
            _ = fi24
        }
    case 0:
        _ = len(v4)
    default:
        v1 = 3.14159
    }
    var v25 uint16 = 32767
    fmt.Println(v2)
    switch 42 {
    case -2147483648:
        v18 = 0.1
    case 0:
        v3[-9223372036854775808] = (!((0 < 1) && (!(421 <= 802))))
    case 32767:
        for fi26 := 0; fi26 < 16; fi26++ {
            for fi27 := 0; fi27 < 9; fi27++ {
                switch (fi26 ^ fi27) {
                case 1:
                    fmt.Println(v18)
                case 4294967295:
                    v6 = 42
                default:
                    switch "a" {
                    case "a":
                        delete(v4, 32768)
                    case "\n\t":
                        fmt.Println(v6, fi27)
                    }
                }
                fmt.Println(fi26, v25, v18)
                var v28 uint16 = 32767
                fmt.Println(v18, fi26)
                _ = fi27
                _ = v28
            }
            fmt.Println(fi26, fi26, v18)
            _ = fi26
        }
    default:
        switch 0 {
        case 32767:
            var v29 uint32 = 42
            _ = v29
        case 42:
            v25 = (v25 + (^(v25)))
        case 2147483647:
            var v30 int = 625
            _ = v30
        }
    }
    v31 := []int{-752, 32768, -1, 0}
    var v32 float32 = v18
    var v33 uint16 = 0
    if (1 <= 415) {
        for fi34 := 0; fi34 < 18; fi34++ {
            v32 = 0.0
            _ = fi34
        }
    }
    fmt.Println(v25)
    if true {
        v1 = ((((v1 + (v1 - (v1 * (0.0 / v1)))) + v1) + (3.14159 - 3.14159)) / 1.0)
        switch ("foo bar" + "\n\t") {
        case "hello":
            var v35 float64 = float64((v18 - float32(3.14159)))
            _ = v35
        case "":
            if false {
                if (861 < 2147483647) {
                    v36 := GenS4{F0: 42, F1: ((v6 < v6) && (true || true)), F2: 0}
                    if len(v31) >= 2 {
                        v31 = v31[1:len(v31)]
                    }
                    switch -698 {
                    case -809:
                        v18 -= v18
                    case 32767:
                        fmt.Println(v33)
                    }
                    _ = v36
                }
                var v37 int = 1
                _ = v37
            }
        }
    } else {
        var v38 uint8 = 236
        _ = v38
    }
    v25 = 42
    v25 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
