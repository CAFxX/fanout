package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint; F1 int16; F2 uint16 }
func genFn0(p0 int) uint32 {
    var v4 int64 = 519
    for fi5 := 0; fi5 < 10; fi5++ {
        v4 = v4
        _ = fi5
    }
    if true {
        v4 = v4
    } else {
        var v6 int32 = 1
        switch "foo bar" {
        case "\n\t":
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := GenS1{F0: 2, F1: 2, F2: 255}
                v9 := GenS2{F0: "", F1: 255, F2: 1}
                _ = fi7
                _ = v8
                _ = v9
            }
        case "hello":
            v10 := GenS3{F0: -84, F1: 42}
            _ = v10
        case "":
            if true {
                var v11 int32 = (v6 ^ v6)
                _ = v11
            }
        }
        if ((-2.25 - 1.5) <= (1.5 - ((float64((-2.25 + ((((1e10 - 0.0) - (((1.5 / -2.25) + 0.0) - 0.0)) + float32(1e10)) - 0.1))) * (((3.14159 + 0.1) + 1.5) * float64(1.5))) / 1.0))) {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 float32 = -2.25
                v14 := GenS2{F0: "", F1: 140, F2: 0}
                _ = fi12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                for fi16 := 0; fi16 < 6; fi16++ {
                    v17 := GenS2{F0: "hello", F1: 127, F2: 128}
                    v4 = v4
                    _ = fi16
                    _ = v17
                }
                switch "" {
                case "a":
                    v18 := GenS4{F0: 0, F1: 42}
                    _ = v18
                case "hello":
                    v6 += (-(v6))
                default:
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v20 := []int{-2147483648}
                        _ = fi19
                        _ = v20
                    }
                }
                _ = fi15
            }
            if false {
                v21 := GenS4{F0: 17, F1: 2}
                _ = v21
            } else {
                var v22 string = "hello"
                _ = v22
            }
        } else {
            v23 := GenS2{F0: "", F1: 226, F2: 0}
            for fi24 := 0; fi24 < 18; fi24++ {
                switch "" {
                case "hello":
                    var v25 int64 = (-(v4))
                    _ = v25
                case "\n\t":
                    p0 = fi24
                case "foo bar":
                    var v26 uint32 = 818
                    _ = v26
                }
                v27 := GenS1{F0: 1, F1: 877, F2: 659}
                _ = fi24
                _ = v27
            }
            v28 := []int{p0, (p0 << ((232) & 63)), 422}
            _ = v23
            _ = v28
        }
        _ = v6
    }
    _ = p0
    _ = v4
    return 535
}
func genFn1() uint32 {
    var v29 bool = true
    for fi30 := 0; fi30 < 17; fi30++ {
        v29 = v29
        v31 := GenS1{F0: 1, F1: 163, F2: 1}
        _ = fi30
        _ = v31
    }
    var v32 uint8 = 0
    _ = v29
    _ = v32
    return 239
}

func main() {
    var v1 int64 = -548
    v2 := make(map[string]bool)
    var v3 int64 = v1
    v33 := []int{782, 127}
    for fi34 := 0; fi34 < 5; fi34++ {
        switch -258 {
        case 128:
            for fi35 := 0; fi35 < 2; fi35++ {
                var v36 int8 = -1
                if ("foo bar" != "foo bar") {
                    v1 = 0
                }
                _ = fi35
                _ = v36
            }
        default:
            v3 = v1
        }
        _ = fi34
    }
    var v37 uint32 = 2
    var v38 uint8 = 2
    var v39 uint16 = 1
    for fi40 := 0; fi40 < 20; fi40++ {
        v3 = (v3 >> ((562) & 63))
        _ = fi40
    }
    switch -292 {
    case 2147483647:
        var v41 uint64 = 2
        _ = v41
    default:
        v42 := GenS5{F0: -70, F1: v39}
        _ = v42
    }
    v43 := v2["foo bar"]
    v44 := v2[("foo bar" + "hello")]
    var v45 uint8 = v38
    var v46 bool = (4294967295 != 42)
    fmt.Println(v3, v37, v44)
    var v47 float64 = (((1e10 - 1.5) - 0.1) + (0.1 - 1.5))
    var v48 int16 = 1
    v49 := GenS5{F0: -128, F1: (v39 + 317)}
    for fi50 := 0; fi50 < 15; fi50++ {
        for fi51 := 0; fi51 < 18; fi51++ {
            if (v39 < v39) {
                switch ("" + "a") {
                case "":
                    v39 = 32767
                }
                if v44 {
                    switch ((("\n\t" + "\n\t") + "\n\t") + "a") {
                    case "hello":
                        v3 = (-(v3))
                    case "":
                        var v52 uint64 = 15
                        _ = v52
                    default:
                        var v53 int64 = v3
                        _ = v53
                    }
                }
                var v54 uint16 = (v39 + v39)
                _ = v54
            } else {
                v39 = 0
                v55 := GenS6{F0: 255, F1: (v48 + (v48 & (v48 / (((v48 * (-(v48)))) | 1)))), F2: 42}
                _ = len(v33)
                _ = v55
            }
            _ = fi51
        }
        _ = fi50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
