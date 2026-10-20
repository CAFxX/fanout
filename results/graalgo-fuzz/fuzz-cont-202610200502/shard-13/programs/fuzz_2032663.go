package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int; F1 int; F2 uint }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 string }
func genFn0() int16 {
    var v6 float64 = (1.5 + (0.1 * -2.25))
    if false {
        var v7 int8 = 127
        switch 128 {
        case 2147483647:
            var v8 int16 = -32768
            _ = v8
        }
        _ = v7
    } else {
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := []int{(fi9 + (fi9 & fi9))}
            v11 := GenS1{F0: fi9, F1: false}
            var v12 float64 = float64(0.1)
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    v13 := []int{-9223372036854775808}
    _ = v13
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 42
    var v3 uint = 2
    var v4 uint8 = 255
    v5 := []uint32{872, 869}
    if len(v5) > 0 {
        v5[0] = uint32(v3)
    }
    var v14 uint8 = v4
    v14 = v14
    v14 = v4
    if true {
        fmt.Println(v4, v2)
        var v15 uint32 = 1
        v16 := GenS2{F0: 32768, F1: 0, F2: 901}
        _ = v15
        _ = v16
    }
    var v17 float64 = float64(float32(0.1))
    v4 = v4
    var v18 float32 = (((float32((float64((3.14159 - 1.5)) / v17)) - 3.14159) / 1e10) / (0.1 + 1.5))
    if false {
        if ((0.1 != float32(-2.25)) && true) {
            var v19 uint32 = 227
            switch "a" {
            case "a":
                var v20 int8 = 127
                _ = v20
            }
            _ = v19
        }
        v21 := GenS2{F0: -2147483648, F1: 0, F2: (^(v3))}
        _ = v21
    } else {
        var v22 int16 = 476
        for fi23 := 0; fi23 < 15; fi23++ {
            fmt.Println(v3, v3, v18)
            switch (fi23 + fi23) {
            case 9223372036854775807:
                switch (fi23 + fi23) {
                case 127:
                    _ = len(v5)
                case -754:
                    var v24 uint16 = 65535
                    _ = v24
                default:
                    for fi25 := 0; fi25 < 14; fi25++ {
                        v26 := GenS3{F0: 364, F1: 1}
                        _ = fi25
                        _ = v26
                    }
                }
            case 32767:
                var v27 string = (("a" + "") + "a")
                _ = v27
            default:
                delete(v1, fi23)
            }
            delete(v1, (fi23 | -630))
            _ = fi23
        }
        _ = v22
    }
    switch "\n\t" {
    case "\n\t":
        v2 = 664
    default:
        for fi28 := 0; fi28 < 13; fi28++ {
            if (!false) {
                for fi29 := 0; fi29 < 17; fi29++ {
                    switch int(v14) {
                    case 4294967295:
                        if true {
                            v30 := GenS1{F0: (^(fi29)), F1: ((v17 - 1.5) < v17)}
                            var v31 uint8 = 1
                            v32 := GenS1{F0: fi28, F1: (!((fi28 << ((255) & 63)) == (fi28 & (fi28 ^ (fi28 & (fi28 + (-(fi29))))))))}
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    default:
                        v33 := GenS4{F0: 32767}
                        _ = v33
                    }
                    _ = fi29
                }
            }
            v34 := GenS2{F0: (fi28 >> ((v3) & 63)), F1: fi28, F2: (v3 ^ (^(v3)))}
            var v35 float64 = 3.14159
            _ = fi28
            _ = v34
            _ = v35
        }
    }
    if true {
        v17 += 1e10
        var v36 string = (("" + ("foo bar" + ("\n\t" + "foo bar"))) + "")
        _ = v36
    } else {
        delete(v1, -2147483648)
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 float64 = v17
            for fi39 := 0; fi39 < 10; fi39++ {
                v40 := GenS5{F0: (-(fi39)), F1: "foo bar"}
                _ = fi39
                _ = v40
            }
            v4 += v14
            _ = fi37
            _ = v38
        }
    }
    if true {
        v14 += v14
        for fi41 := 0; fi41 < 7; fi41++ {
            var v42 uint8 = (^(v4))
            _ = len(v1)
            _ = fi41
            _ = v42
        }
        switch 9 {
        case 9223372036854775807:
            switch 0 {
            case 2147483647:
                if true {
                    var v43 float32 = -2.25
                    _ = v43
                } else {
                    var v44 float32 = ((((0.1 + v18) + v18) - (v18 - (float32(v17) * (3.14159 * (v18 / 1.0))))) / float32((((0.0 / v17) * 0.0) * -2.25)))
                    _ = v44
                }
            case 32767:
                if len(v5) > 0 {
                    v5[0] = 370
                }
            default:
                var v45 uint32 = 2
                _ = v45
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
