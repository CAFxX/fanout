package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool; F1 int8; F2 uint }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[int]string)
    var v2 int = -268
    var v3 uint64 = 1
    v2 = (^(v2))
    fmt.Println(v2, v3)
    fmt.Println(v2, v2)
    var v4 float32 = (1e10 - (float32((0.1 + (1.5 + (-2.25 * -2.25)))) - 0.0))
    v3 -= v3
    var v5 int32 = 1
    if true {
        for fi6 := 0; fi6 < 19; fi6++ {
            v2 -= v2
            fmt.Println(v2)
            _ = fi6
        }
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                v9 := GenS1{F0: 192, F1: -114, F2: 32767}
                var v10 int16 = 42
                v4 = 1e10
                fmt.Println(v10, v3)
                v10 = int16(v5)
                _ = fi8
                _ = v10
                _ = v9
            }
            if true {
                v1[(-(fi7))] = ("\n\t" + "hello")
            }
            var v11 int64 = -1
            _ = fi7
            _ = v11
        }
        v2 += (v2 | (v2 << ((0) & 63)))
    }
    fmt.Println(v3, v4, v4)
    var v12 int16 = 1000
    v13 := GenS1{F0: 332, F1: 54, F2: 42}
    fmt.Println(v5, v5, v5)
    switch "hello" {
    case "a":
        if (713 == 982) {
            v4 = v4
            switch (v2 - v2) {
            case 128:
                for fi14 := 0; fi14 < 2; fi14++ {
                    if (!false) {
                        v12 += 1000
                    } else {
                        var v15 float32 = 3.14159
                        v3 = 1
                        _ = v15
                    }
                    _ = fi14
                }
            case 2147483647:
                switch (v2 << ((1) & 63)) {
                case -971:
                    switch 127 {
                    case 542:
                        var v16 int32 = v5
                        _ = v16
                    default:
                        var v17 float64 = 0.0
                        _ = v17
                    }
                }
            case 42:
                for fi18 := 0; fi18 < 3; fi18++ {
                    switch "foo bar" {
                    case "hello":
                        for fi19 := 0; fi19 < 12; fi19++ {
                            v20 := GenS2{F0: true, F1: 100, F2: 42}
                            v5 = (v5 - int32(v12))
                            _ = fi19
                            _ = v20
                        }
                    }
                    _ = fi18
                }
            }
        } else {
            var v21 uint8 = 0
            _ = len(v1)
            v22 := GenS3{F0: 107}
            _ = v21
            _ = v22
        }
    case "":
        if false {
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 int8 = 60
                _ = fi23
                _ = v24
            }
            var v25 float32 = 1e10
            var v26 uint32 = 464
            _ = v25
            _ = v26
        } else {
            _ = len(v1)
            var v27 uint = 495
            _ = v27
        }
    }
    switch "a" {
    case "hello":
        v3 -= v3
    case "\n\t":
        v28 := GenS1{F0: 944, F1: 0, F2: 723}
        _ = v28
    case "a":
        v29 := []int{v2, (v2 >> ((4294967295) & 63)), v2, (-(v2))}
        _ = v29
    }
    fmt.Println(v2, v2, v4)
    for fi30 := 0; fi30 < 18; fi30++ {
        v1[fi30] = ("" + "foo bar")
        var v31 int = (^(fi30))
        v32 := GenS1{F0: 2, F1: 127, F2: 254}
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 uint = 18446744073709551615
    var v34 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
