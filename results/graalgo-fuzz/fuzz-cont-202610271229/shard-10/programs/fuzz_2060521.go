package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 string; F1 bool; F2 int16 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := []int{1, (fi5 / ((127) | 1))}
            var v8 string = "\n\t"
            _ = fi6
            _ = v7
            _ = v8
        }
        if (!(true && true)) {
            var v9 uint64 = 1
            _ = v9
        } else {
            var v10 string = "\n\t"
            v10 = ("\n\t" + ("\n\t" + v10))
            for fi11 := 0; fi11 < 19; fi11++ {
                if true {
                    var v12 int16 = -461
                    fmt.Println(fi11, v10, v12)
                    _ = v12
                }
                _ = fi11
            }
            _ = v10
        }
        var v13 uint64 = 0
        _ = fi5
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 uint8 = 128
        v16 := GenS1{F0: uint64(v15)}
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := GenS1{F0: 2}
        _ = fi17
        _ = v18
    }
    return 0
}

func main() {
    var v1 uint32 = 65535
    var v2 uint = 0
    v3 := make(map[int]uint32)
    var v4 bool = (false || false)
    v1 += (v1 + 42)
    var v19 float64 = ((1e10 / (1e10 + 3.14159)) + (1e10 - float64(3.14159)))
    var v20 uint8 = 0
    fmt.Println(v1)
    for fi21 := 0; fi21 < 6; fi21++ {
        if (-128 > -82) {
            v19 = (0.1 - (v19 / 1.0))
            for fi22 := 0; fi22 < 15; fi22++ {
                if v4 {
                    var v23 int16 = 1
                    _ = v23
                } else {
                    v24 := GenS2{F0: "", F1: 161}
                    v25 := []int{}
                    fmt.Println(v20, v20)
                    _ = v24
                    _ = v25
                }
                var v26 uint = (v2 << ((v2) & 63))
                v27 := GenS1{F0: 18446744073709551615}
                _ = fi22
                _ = v26
                _ = v27
            }
        }
        var v28 string = "hello"
        _ = fi21
        _ = v28
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        if (!v4) {
            v30 := []int{233}
            v2 = v2
            for fi31 := 0; fi31 < 10; fi31++ {
                fmt.Println(v4)
                if (int64(fi29) == 735) {
                    v32 := v3[fi29]
                    _ = v32
                }
                var v33 int = fi29
                _ = fi31
                _ = v33
            }
            _ = v30
        }
        _ = fi29
    }
    fmt.Println(v4)
    switch ("a" + "a") {
    case "foo bar":
        if (205 != 2) {
            v34 := GenS3{F0: ("" + "a"), F1: false, F2: genFn0()}
            _ = v34
        }
    case "hello":
        switch "" {
        case "":
            var v35 float64 = 3.14159
            _ = v35
        case "a":
            var v36 float32 = (0.0 / 0.1)
            _ = v36
        case "\n\t":
            for fi37 := 0; fi37 < 19; fi37++ {
                v4 = (!v4)
                if v4 {
                    var v38 int16 = -368
                    _ = v38
                }
                _ = fi37
            }
        default:
            var v39 bool = v4
            _ = v39
        }
    case "a":
        var v40 int16 = -1
        _ = v40
    }
    for fi41 := 0; fi41 < 8; fi41++ {
        v19 += v19
        if v4 {
            v19 -= 0.0
        } else {
            for fi42 := 0; fi42 < 4; fi42++ {
                fmt.Println(v20)
                _ = fi42
            }
            var v43 uint32 = v1
            if v4 {
                fmt.Println(v1, v43, v20)
                v44 := []int{-834, (fi41 >> ((v2) & 63))}
                _ = v44
            }
            _ = v43
        }
        var v45 uint8 = (v20 << (((v2 / ((483) | 1))) & 63))
        _ = fi41
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
