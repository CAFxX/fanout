package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint64, p1 uint32, p2 uint16) int8 {
    var v6 uint = 42
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(v6, v6, p1)
        for fi8 := 0; fi8 < 1; fi8++ {
            switch fi8 {
            case 268:
                if (((v6 / ((v6) | 1)) != (v6 ^ (v6 & (v6 + v6)))) && false) {
                    v9 := []int{fi8, fi7, (^(fi7))}
                    p2 = 32767
                    _ = v9
                }
            case 235:
                v6 -= v6
            case 679:
                if (((false || (!("" <= ("a" + "")))) && ((0.0 == (-2.25 - (1e10 - 0.0))) && false)) || ((^(p1)) >= (p1 + p1))) {
                    var v10 uint32 = p1
                    _ = v10
                }
            }
            if false {
                v11 := GenS1{F0: p2}
                _ = v11
            }
            _ = fi8
        }
        for fi12 := 0; fi12 < 16; fi12++ {
            switch "foo bar" {
            case "hello":
                v13 := GenS2{F0: fi7}
                _ = v13
            case "":
                fmt.Println(fi12, fi12)
            case "\n\t":
                if false {
                    var v14 float32 = 0.1
                    _ = v14
                }
            default:
                switch (fi12 >> ((867) & 63)) {
                case 450:
                    var v15 uint64 = 42
                    _ = v15
                case 4294967295:
                    v16 := GenS2{F0: 1}
                    _ = v16
                }
            }
            var v17 float64 = 0.0
            if (363 != p2) {
                fmt.Println(v6, p2, p1)
                v18 := GenS3{F0: v6, F1: -1}
                _ = v18
            }
            v19 := GenS4{F0: 127, F1: 128}
            var v20 int16 = -761
            _ = fi12
            _ = v17
            _ = v19
            _ = v20
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 81
}

func main() {
    v1 := []uint32{798}
    v2 := []uint32{4294967295, 2}
    v3 := make(map[int]int)
    var v4 string = ""
    var v5 int64 = 0
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 20; fi22++ {
            switch 32767 {
            case -530:
                if (false || true) {
                    _ = len(v3)
                    v5 += (v5 >> ((4294967295) & 63))
                    switch fi22 {
                    case -9223372036854775808:
                        switch fi22 {
                        case 42:
                            v23 := GenS4{F0: 4, F1: 153}
                            _ = v23
                        }
                        v4 = (("\n\t" + v4) + ((("\n\t" + "\n\t") + (v4 + v4)) + "a"))
                    case -2147483648:
                        v5 = v5
                    default:
                        v24 := v3[2147483647]
                        _ = v24
                    }
                    v25 := GenS5{F0: 824}
                    _ = v25
                }
            default:
                v1 = append(v1, 2147483648)
            }
            var v26 uint32 = 0
            delete(v3, 32767)
            _ = fi22
            _ = v26
        }
        switch v4 {
        case "hello":
            fmt.Println(fi21, fi21, fi21)
        case "a":
            fmt.Println(v4, v4, fi21)
        case "foo bar":
            for fi27 := 0; fi27 < 10; fi27++ {
                switch v4 {
                case "foo bar":
                    v4 = ("hello" + "hello")
                default:
                    fmt.Println(fi27, v5)
                }
                _ = fi27
            }
        default:
            if ((!true) && (2 != 0)) {
                for fi28 := 0; fi28 < 7; fi28++ {
                    v4 = ("foo bar" + ("a" + (v4 + ("\n\t" + v4))))
                    _ = fi28
                }
                fmt.Println(fi21, v5)
                if (!true) {
                    var v29 int16 = 0
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (true || false) {
                        v4 = (v4 + v4)
                    } else {
                        v30 := GenS5{F0: 434}
                        _ = v30
                    }
                    _ = v29
                }
            } else {
                _ = len(v2)
            }
        }
        _ = fi21
    }
    v5 = 0
    var v31 uint64 = 331
    v32 := GenS2{F0: -9223372036854775808}
    var v33 uint8 = 42
    var v34 int8 = 127
    var v35 int64 = v5
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch 32768 {
        case -1:
            var v36 int16 = 1000
            _ = v36
        case -342:
            v37 := v3[52]
            _ = v37
        }
        if (232 <= 9223372036854775807) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            delete(v3, -310)
            v38 := GenS6{F0: 1}
            _ = v38
        }
    }
    v3[620] = 591
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
