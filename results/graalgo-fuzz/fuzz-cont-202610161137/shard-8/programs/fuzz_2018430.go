package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint8; F1 uint; F2 int64 }
func genFn0(p0 uint64, p1 int, p2 int64) uint32 {
    p2 = (p2 & 42)
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{991}
    v3 := []string{"hello", "foo bar", ("\n\t" + "foo bar")}
    var v4 int32 = -1
    var v5 uint = 231
    v7 := GenS1{F0: 42}
    if (false || true) {
        for fi8 := 0; fi8 < 8; fi8++ {
            switch (fi8 + fi8) {
            case 1:
                v9 := GenS2{F0: fi8, F1: (!("" == "\n\t"))}
                _ = v9
            case -1:
                _ = len(v3)
            case 32767:
                v5 += 2
            }
            _ = fi8
        }
    }
    _ = len(v3)
    if (-422 >= (v4 - (v4 * 42))) {
        for fi10 := 0; fi10 < 20; fi10++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi10
        }
        switch 1 {
        case -2147483648:
            fmt.Println(v5, v4)
        case 9223372036854775807:
            var v11 float64 = (1.5 - -2.25)
            _ = v11
        default:
            _ = len(v3)
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            v5 += v5
            if (0.1 != 1.5) {
                var v13 uint32 = 0
                v14 := GenS1{F0: uint16(v5)}
                v13 += genFn0(814, 903, -501)
                _ = v13
                _ = v14
            }
            _ = fi12
        }
    }
    v4 = (v4 | v4)
    v4 -= v4
    v5 = 1
    if false {
        switch 183 {
        case 1:
            var v15 uint8 = 0
            _ = v15
        case 0:
            var v16 uint32 = 42
            _ = v16
        default:
            v1["\n\t"] = (("\n\t" + ("" + ("" + "\n\t"))) + "a")
        }
        v17 := GenS3{F0: 48, F1: 1, F2: 42}
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v4 += v4
        switch "hello" {
        case "\n\t":
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 uint8 = uint8(v5)
                _ = fi19
                _ = v20
            }
        default:
            var v21 uint64 = 522
            _ = v21
        }
        if false {
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := GenS2{F0: 32768, F1: true}
                var v24 float32 = 1e10
                var v25 float64 = 0.0
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            v4 += (v4 >> ((v5) & 63))
        } else {
            if ((255 < 33) && false) {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v5, v5, fi18)
                case "foo bar":
                    v26 := GenS3{F0: 211, F1: v5, F2: -1}
                    _ = v26
                default:
                    for fi27 := 0; fi27 < 7; fi27++ {
                        fmt.Println(v4, v5, v4)
                        if true {
                            _ = len(v1)
                        } else {
                            var v28 bool = true
                            _ = v28
                        }
                        v5 -= v5
                        _ = fi27
                    }
                }
                if true {
                    fmt.Println(fi18, v4, v4)
                    var v29 int16 = -627
                    _ = v29
                } else {
                    v4 = (v4 & 0)
                }
                fmt.Println(v5, v4)
            }
            v30 := GenS2{F0: 128, F1: (1 >= (v5 * 1))}
            for fi31 := 0; fi31 < 5; fi31++ {
                var v32 uint = (v5 + 255)
                _ = fi31
                _ = v32
            }
            _ = v30
        }
        _ = fi18
    }
    var v33 int8 = -128
    v33 += 1
    _ = len(v1)
    v34 := GenS2{F0: 984, F1: (false && true)}
    var v35 int = 1
    v4 += (v4 % ((v4) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v7
}
