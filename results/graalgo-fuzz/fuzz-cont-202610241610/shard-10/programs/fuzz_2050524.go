package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int16; F1 int; F2 uint }
type GenS3 struct { F0 uint8; F1 int64; F2 int32 }
type GenS4 struct { F0 int32; F1 bool; F2 bool }
type GenS5 struct { F0 int16; F1 uint16; F2 uint64 }
func genFn0() uint8 {
    switch -507 {
    case 32767:
        var v7 float32 = ((1.5 / (1e10 / 1.5)) * ((0.1 / (0.1 + ((0.1 - float32((float64(0.0) * 1e10))) / float32(float64(-2.25))))) - (1e10 + 1e10)))
        _ = v7
    case -2147483648:
        v8 := GenS1{F0: 792, F1: "foo bar"}
        _ = v8
    }
    var v9 int64 = 120
    var v10 int = 32767
    _ = v10
    _ = v9
    return 127
}

func main() {
    var v1 uint = 42
    var v2 uint8 = 251
    var v3 int16 = 42
    var v4 uint16 = 255
    var v5 uint16 = (v4 << ((v1) & 63))
    var v6 bool = true
    if ((v6 || (!true)) && v6) {
        v11 := []int{}
        v3 += v3
        if (false || (false && false)) {
            var v12 int32 = 1
            v13 := GenS1{F0: 344, F1: "\n\t"}
            _ = v12
            _ = v13
        }
        _ = v11
    }
    if (0 > 503) {
        var v14 uint32 = 1
        _ = v14
    } else {
        v2 = (v2 << ((v1) & 63))
    }
    v15 := GenS2{F0: (v3 - v3), F1: 0, F2: (v1 | uint(v4))}
    v3 -= 972
    var v16 int32 = 42
    var v17 float64 = 1.5
    v18 := GenS3{F0: (v2 >> ((511) & 63)), F1: -1, F2: int32(v1)}
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 15; fi20++ {
            v4 += uint16(v1)
            v21 := GenS4{F0: v16, F1: v6, F2: (v17 == v17)}
            switch "hello" {
            case "a":
                var v22 float32 = 3.14159
                _ = v22
            default:
                v6 = v6
            }
            _ = fi20
            _ = v21
        }
        switch (("" + "\n\t") + ("" + "\n\t")) {
        case "hello":
            if (!v6) {
                for fi23 := 0; fi23 < 6; fi23++ {
                    v24 := GenS5{F0: v3, F1: v5, F2: 424}
                    _ = fi23
                    _ = v24
                }
                if v6 {
                    v25 := []int{(-(fi19)), -602, (-(fi19))}
                    v2 = (v2 >> ((v1) & 63))
                    if v6 {
                        if false {
                            v26 := GenS4{F0: v16, F1: ((!(v3 > v3)) && (v16 > (v16 + 2147483647))), F2: v6}
                            fmt.Println(v6)
                            _ = v26
                        }
                        var v27 float64 = (v17 - v17)
                        for fi28 := 0; fi28 < 5; fi28++ {
                            var v29 string = "hello"
                            _ = fi28
                            _ = v29
                        }
                        _ = v27
                    } else {
                        var v30 float32 = 0.1
                        v30 = 3.14159
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                        _ = v30
                    }
                    _ = v25
                }
            }
        case "\n\t":
            var v31 int8 = -1
            _ = v31
        }
        _ = fi19
    }
    v2 = 42
    switch -800 {
    case -9223372036854775808:
        if v6 {
            v32 := GenS1{F0: 4294967295, F1: "a"}
            _ = v32
        }
    default:
        v33 := []int{0, int(v2), 1}
        _ = v33
    }
    v34 := []int{127}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
