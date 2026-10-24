package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint8 }
func genFn0(p0 string, p1 uint16, p2 uint64) int16 {
    var v5 int = 290
    for fi6 := 0; fi6 < 9; fi6++ {
        if true {
            p1 = (^(p1))
        }
        v7 := GenS1{F0: 1, F1: 2}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float32 = (1.5 * float32(1e10))
        var v10 uint8 = 42
        var v11 int8 = 100
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println(p2, v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 uint32 = 2
    var v4 float64 = ((1e10 * -2.25) - (1.5 + float64(0.1)))
    var v12 float32 = (((0.1 - 1e10) / ((1e10 + (((0.1 - (3.14159 / (float32(v4) + ((((0.0 / float32(v4)) * 1e10) + 0.0) - 1.5)))) - (((1e10 - 1e10) - ((0.0 - 0.0) + -2.25)) + 1e10)) * 1e10)) + 3.14159)) * ((0.1 - 1.5) - (((0.0 + 3.14159) - 1.5) + -2.25)))
    v13 := GenS2{F0: ("\n\t" + ("" + ""))}
    if false {
        var v14 int16 = genFn0(("" + "foo bar"), 0, 4294967295)
        if (766 <= 1) {
            var v15 bool = true
            v15 = (true || true)
            v12 += (v12 + ((0.0 / 3.14159) / v12))
            _ = v15
        } else {
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 uint32 = (^(v3))
                for fi18 := 0; fi18 < 10; fi18++ {
                    delete(v1, ("\n\t" + ("hello" + "a")))
                    v19 := []int{(fi16 | (fi16 * fi18)), (fi16 - (fi16 << ((770) & 63)))}
                    fmt.Println(v17)
                    _ = fi18
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
            switch "" {
            case "foo bar":
                v20 := GenS1{F0: (v14 / (((-(v14))) | 1)), F1: (v3 & v3)}
                _ = v20
            }
        }
        _ = v14
    } else {
        v21 := GenS3{F0: 18446744073709551615, F1: ""}
        for fi22 := 0; fi22 < 6; fi22++ {
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 uint32 = (v3 * (v3 % ((v3) | 1)))
                if false {
                    var v25 string = "a"
                    _ = v25
                }
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            if false {
                var v27 uint64 = 42
                switch "hello" {
                case "hello":
                    v27 = v27
                case "\n\t":
                    v28 := v1[("" + "hello")]
                    _ = v28
                case "":
                    v12 = 1e10
                }
                _ = v27
            } else {
                fmt.Println(v3, v4, fi26)
                fmt.Println(v12, v12)
                v3 -= v3
            }
            _ = fi26
        }
        _ = v21
    }
    v29 := v2[509]
    var v30 int8 = 42
    v4 += float64(v12)
    for fi31 := 0; fi31 < 13; fi31++ {
        v30 = 42
        v32 := v2[(^(v3))]
        v33 := v2[(^(v3))]
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 int32 = 0
    switch 42 {
    case 848:
        v34 = (v34 + v34)
    case -2147483648:
        if (true && (false && true)) {
            v12 = float32(((3.14159 * v4) + (((0.0 / 3.14159) / v4) - float64(0.0))))
            for fi35 := 0; fi35 < 19; fi35++ {
                v36 := GenS4{F0: 2}
                var v37 int64 = 299
                _ = fi35
                _ = v36
                _ = v37
            }
        }
    case 32767:
        switch "foo bar" {
        case "\n\t":
            v4 = (float64(float32((-2.25 * (v4 + 1e10)))) + v4)
        case "":
            v4 = 1.5
        }
    default:
        v2[v3] = -199
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
}
