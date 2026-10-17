package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    var v3 int = 128
    fmt.Println(v3)
    if (191 >= 42) {
        if (3.14159 == float32((1e10 * ((-2.25 * -2.25) * 3.14159)))) {
            for fi4 := 0; fi4 < 4; fi4++ {
                v5 := GenS1{F0: 2, F1: (1 >= 2147483647)}
                if (("a" + (("" + ("hello" + "")) + ("foo bar" + ((("" + "a") + "foo bar") + ("hello" + "\n\t"))))) >= "hello") {
                    v3 -= v3
                    if true {
                        for fi6 := 0; fi6 < 2; fi6++ {
                            v7 := GenS1{F0: 0, F1: (false && (1e10 != -2.25))}
                            v8 := GenS2{F0: 787, F1: 0}
                            var v9 int8 = 0
                            _ = fi6
                            _ = v7
                            _ = v8
                            _ = v9
                        }
                        v3 -= -2147483648
                        v10 := []int{v3, fi4, -802, 32768}
                        _ = v10
                    }
                }
                _ = fi4
                _ = v5
            }
        }
    } else {
        if (true && (185 > -1)) {
            switch "\n\t" {
            case "a":
                switch v3 {
                case -9223372036854775808:
                    var v11 bool = false
                    _ = v11
                case 32768:
                    var v12 int32 = -406
                    _ = v12
                default:
                    v3 -= v3
                }
            case "foo bar":
                for fi13 := 0; fi13 < 14; fi13++ {
                    fmt.Println(fi13)
                    _ = fi13
                }
            case "\n\t":
                for fi14 := 0; fi14 < 6; fi14++ {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        var v16 uint = 42
                        v17 := GenS3{F0: -1}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    v18 := []int{9223372036854775807, (fi14 - fi14)}
                    _ = fi14
                    _ = v18
                }
            }
            v19 := []int{(v3 + v3), (v3 * v3), -2147483648, (v3 | v3)}
            var v20 int64 = 1
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 16; fi21++ {
                var v22 string = ((("" + ("foo bar" + ("a" + "foo bar"))) + ("\n\t" + "")) + "hello")
                _ = fi21
                _ = v22
            }
        }
        v23 := GenS1{F0: 800, F1: false}
        for fi24 := 0; fi24 < 5; fi24++ {
            fmt.Println(fi24)
            var v25 uint8 = 1
            if (1 > 0) {
                fmt.Println(fi24, fi24, v25)
            }
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    v3 -= (-(v3))
    if (!(!(42 > (^(v3))))) {
        for fi26 := 0; fi26 < 20; fi26++ {
            v3 = fi26
            fmt.Println(v3)
            v27 := GenS4{F0: 42, F1: 1}
            _ = fi26
            _ = v27
        }
        var v28 int = (v3 * v3)
        for fi29 := 0; fi29 < 2; fi29++ {
            if (1 == 42) {
                v30 := GenS2{F0: 0, F1: 42}
                v31 := []int{}
                for fi32 := 0; fi32 < 11; fi32++ {
                    v3 = (fi29 / (((v28 ^ (-(v3)))) | 1))
                    v28 = v3
                    v28 = 9223372036854775807
                    _ = fi32
                }
                _ = v30
                _ = v31
            } else {
                var v33 int = -434
                _ = v33
            }
            v28 += v3
            v34 := GenS5{F0: 961}
            _ = fi29
            _ = v34
        }
        _ = v28
    }
    var v35 int = (v3 ^ v3)
    var v36 uint8 = 2
    var v37 uint8 = (v36 * 0)
    var v38 int16 = 32767
    fmt.Println(v36, v37)
    var v39 uint16 = 419
    for fi40 := 0; fi40 < 20; fi40++ {
        v41 := GenS3{F0: (-(v38))}
        _ = fi40
        _ = v41
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
