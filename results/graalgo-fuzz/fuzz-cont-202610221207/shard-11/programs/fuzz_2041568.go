package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int }
type GenS2 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int32; F1 string; F2 uint16 }
func genFn0(p0 int8, p1 uint64) int32 {
    var v6 bool = false
    var v7 string = "\n\t"
    v8 := GenS1{F0: 2, F1: -2147483648, F2: 467}
    var v9 int8 = 0
    v6 = (v6 && v6)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int = -1
    var v2 int16 = 1000
    var v3 float64 = 0.1
    var v4 int16 = -1
    v5 := make(map[uint32]bool)
    var v10 bool = (!(2 != 4294967295))
    v11 := GenS1{F0: 255, F1: 1, F2: 2147483647}
    var v12 float32 = (((1e10 * 0.1) + float32(float64((1.5 + 1e10)))) - (1e10 * float32((((((0.1 + (-2.25 - float64((-2.25 + 0.0)))) + v3) + ((((v3 / ((1.5 + v3) * (v3 / (0.1 + ((3.14159 / -2.25) + v3))))) / float64(-2.25)) - (v3 - 1.5)) / -2.25)) / v3) / (float64(3.14159) - v3)))))
    if (v1 < -755) {
        v3 = 1.5
        v10 = true
    } else {
        for fi13 := 0; fi13 < 15; fi13++ {
            switch 1 {
            case 32768:
                var v14 float64 = v3
                _ = v14
            case 19:
                var v15 string = ("hello" + "\n\t")
                _ = v15
            case 9223372036854775807:
                var v16 float32 = (1.5 - float32(v3))
                _ = v16
            default:
                if true {
                    delete(v5, 2)
                    fmt.Println(v1)
                }
            }
            for fi17 := 0; fi17 < 19; fi17++ {
                v4 = (v4 % (((v2 >> ((65535) & 63))) | 1))
                v18 := GenS2{F0: (v4 - v2), F1: 24, F2: 0}
                if v10 {
                    if v10 {
                        if false {
                            v19 := GenS1{F0: 1, F1: genFn0(-1, 0), F2: v1}
                            var v20 uint32 = 2
                            _ = v19
                            _ = v20
                        }
                    }
                }
                _ = fi17
                _ = v18
            }
            _ = fi13
        }
        v12 = 3.14159
    }
    v10 = (!(2 < 1))
    v10 = v10
    switch "a" {
    case "hello":
        v3 = float64(3.14159)
    default:
        v10 = ("a" == "a")
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 uint16 = 255
        delete(v5, 2147483648)
        if v10 {
            var v23 bool = v10
            for fi24 := 0; fi24 < 16; fi24++ {
                v12 = v12
                v4 = v4
                _ = fi24
            }
            _ = v23
        } else {
            fmt.Println(v12, v1, v3)
            if (("" + "a") < ("\n\t" + "a")) {
                if (!(255 < 4294967295)) {
                    var v25 int64 = -9223372036854775808
                    v26 := GenS3{F0: 232, F1: 1}
                    v27 := GenS1{F0: 811, F1: 688, F2: (fi21 << ((65535) & 63))}
                    _ = v25
                    _ = v26
                    _ = v27
                }
                fmt.Println(v12)
            }
        }
        _ = fi21
        _ = v22
    }
    v3 = 0.0
    if ((!(2 >= 581)) && (!(3 == -128))) {
        switch "hello" {
        case "hello":
            for fi28 := 0; fi28 < 9; fi28++ {
                if (0 != 227) {
                    v29 := GenS2{F0: v4, F1: 1, F2: 18446744073709551615}
                    _ = v29
                }
                switch (fi28 / ((906) | 1)) {
                case 32768:
                    var v30 int64 = 42
                    _ = v30
                case 4294967295:
                    switch "\n\t" {
                    case "":
                        var v31 uint = 2
                        _ = v31
                    case "hello":
                        var v32 float64 = (v3 + 0.1)
                        _ = v32
                    case "a":
                        v3 = ((v3 + v3) + v3)
                    default:
                        if ((4294967295 == 0) || false) {
                            var v33 float64 = (1.5 * v3)
                            _ = v33
                        }
                    }
                case -2147483648:
                    for fi34 := 0; fi34 < 5; fi34++ {
                        var v35 uint64 = 617
                        fmt.Println(fi28, v10, fi28)
                        if (!(42 > genFn0(46, v35))) {
                            var v36 float64 = 1.5
                            v1 = 273
                            var v37 bool = false
                            _ = v36
                            _ = v37
                        }
                        _ = fi34
                        _ = v35
                    }
                }
                var v38 int32 = -994
                _ = fi28
                _ = v38
            }
        case "\n\t":
            v12 = v12
        case "":
            v39 := GenS2{F0: (v4 & 32767), F1: 4294967295, F2: 1}
            _ = v39
        }
    } else {
        var v40 uint64 = 478
        _ = v40
    }
    v12 += v12
    for fi41 := 0; fi41 < 1; fi41++ {
        v42 := GenS4{F0: -2147483648, F1: "a", F2: 42}
        var v43 bool = (1.5 <= v3)
        if true {
            v44 := []int{(^(fi41)), (fi41 << ((4294967295) & 63)), fi41, fi41}
            v44 = append(v44, 32767)
            _ = v44
        }
        _ = fi41
        _ = v42
        _ = v43
    }
    v2 = (v2 << ((316) & 63))
    var v45 uint16 = 0
    var v46 int16 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
