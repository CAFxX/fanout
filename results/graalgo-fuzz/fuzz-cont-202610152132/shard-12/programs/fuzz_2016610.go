package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int32 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int = 650
    var v2 string = "\n\t"
    var v3 int16 = 27
    var v4 int = v1
    v5 := []int{v1, (v1 - (v4 & 336))}
    var v6 string = v2
    var v7 int = (v4 & v4)
    var v8 uint8 = 127
    var v9 int64 = 42
    var v10 float32 = (((0.0 - 0.0) - (1.5 - ((1.5 / 3.14159) * 0.0))) - 3.14159)
    for fi11 := 0; fi11 < 12; fi11++ {
        if false {
            if true {
                switch ((("\n\t" + v6) + v2) + v2) {
                case "hello":
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v13 := GenS1{F0: 4294967295, F1: false, F2: 0}
                        _ = fi12
                        _ = v13
                    }
                case "\n\t":
                    for fi14 := 0; fi14 < 16; fi14++ {
                        var v15 uint64 = 4294967295
                        _ = fi14
                        _ = v15
                    }
                }
            }
        }
        var v16 int = v7
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int16 = v3
            _ = fi17
            _ = v18
        }
        _ = fi11
        _ = v16
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 uint64 = 526
        switch v7 {
        case 471:
            v9 += 9223372036854775807
        case 927:
            var v21 bool = (-1 >= -828)
            _ = v21
        default:
            for fi22 := 0; fi22 < 8; fi22++ {
                v23 := GenS1{F0: 33, F1: true, F2: -1}
                var v24 float64 = ((0.1 + 0.0) - -2.25)
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        v25 := GenS1{F0: 266, F1: false, F2: 1}
        _ = fi19
        _ = v20
        _ = v25
    }
    if (!false) {
        var v26 float32 = v10
        v26 -= ((float32(-2.25) + v26) * v10)
        for fi27 := 0; fi27 < 6; fi27++ {
            v28 := GenS1{F0: 255, F1: true, F2: 42}
            for fi29 := 0; fi29 < 17; fi29++ {
                v5 = append(v5, v7)
                if false {
                    v30 := GenS2{F0: false}
                    var v31 uint = 0
                    if false {
                        v10 -= v10
                        fmt.Println(v2, v26, v2)
                    }
                    _ = v30
                    _ = v31
                } else {
                    var v32 int32 = 0
                    var v33 float32 = v10
                    if (0.1 == (3.14159 / 3.14159)) {
                        fmt.Println(v2, v1)
                        var v34 uint8 = 255
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if true {
                            v35 := GenS1{F0: 111, F1: false, F2: (-(v32))}
                            var v36 uint16 = 42
                            _ = v35
                            _ = v36
                        } else {
                            v8 -= 123
                        }
                        var v37 uint32 = 1
                        _ = v34
                        _ = v37
                    }
                    _ = v32
                    _ = v33
                }
                switch fi27 {
                case -974:
                    var v38 bool = (-128 >= 109)
                    _ = v38
                }
                _ = fi29
            }
            if len(v5) > 0 {
                v5[0] = v1
            }
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    v2 = (v6 + v2)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
