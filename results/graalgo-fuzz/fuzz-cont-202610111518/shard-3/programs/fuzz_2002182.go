package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 int32 = 1
    v2 := make(map[int]string)
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]uint32)
    var v5 int = 9223372036854775807
    if (-32768 != 369) {
        var v6 int = 0
        if (false || (756 == 2)) {
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 int32 = 42
                for fi9 := 0; fi9 < 13; fi9++ {
                    v8 = -362
                    _ = fi9
                }
                v1 -= (-(v8))
                _ = fi7
                _ = v8
            }
        }
        delete(v2, -48)
        _ = v6
    } else {
        for fi10 := 0; fi10 < 11; fi10++ {
            if (false && true) {
                if (2 == 229) {
                    var v11 int = -955
                    for fi12 := 0; fi12 < 18; fi12++ {
                        v11 += (fi10 >> ((255) & 63))
                        _ = fi12
                    }
                    _ = v11
                }
            }
            fmt.Println(v1, fi10)
            v13 := GenS1{F0: 1, F1: 968}
            _ = fi10
            _ = v13
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 uint16 = 2
            if true {
                switch ("\n\t" + "hello") {
                case "":
                    v16 := v4[v5]
                    _ = v16
                case "hello":
                    v17 := []int{v5, (fi14 + fi14)}
                    _ = v17
                }
                switch "foo bar" {
                case "":
                    if false {
                        v3 -= (v3 >> ((2) & 63))
                        switch "hello" {
                        case "\n\t":
                            var v18 int8 = 100
                            _ = v18
                        case "":
                            fmt.Println(fi14)
                        }
                    }
                case "a":
                    var v19 int8 = -82
                    _ = v19
                }
                var v20 int16 = 32767
                _ = v20
            } else {
                for fi21 := 0; fi21 < 4; fi21++ {
                    var v22 float64 = -2.25
                    _ = fi21
                    _ = v22
                }
            }
            _ = fi14
            _ = v15
        }
    }
    var v23 uint32 = 0
    switch "a" {
    case "a":
        delete(v2, v5)
    case "\n\t":
        v24 := GenS1{F0: 2, F1: 311}
        _ = v24
    case "foo bar":
        fmt.Println(v1, v3)
    }
    var v25 string = ("" + "foo bar")
    fmt.Println(v23, v1, v25)
    v26 := []int{(v5 ^ (v5 >> ((uint(v23)) & 63)))}
    v1 = (v1 - (v1 & v1))
    fmt.Println(v23)
    for fi27 := 0; fi27 < 18; fi27++ {
        v1 += (v1 | v1)
        switch (v5 * (-(fi27))) {
        case 4294967295:
            var v28 uint64 = 4294967295
            _ = v28
        default:
            for fi29 := 0; fi29 < 7; fi29++ {
                var v30 string = ("foo bar" + ("foo bar" + "a"))
                _ = fi29
                _ = v30
            }
        }
        _ = fi27
    }
    if false {
        v25 = (((v25 + "hello") + v25) + "hello")
        var v31 float32 = 0.1
        _ = v31
    } else {
        fmt.Println(v1, v1)
    }
    var v32 bool = (false && (3.14159 != (-2.25 * float64((1.5 * (1.5 + ((0.0 + (-2.25 + (3.14159 / 1.0))) / ((-2.25 - 0.0) * float32((float64(((0.1 * (3.14159 + (float32(0.0) + 0.0))) - -2.25)) * (float64((0.1 * (float32((-2.25 - 0.0)) - float32((1e10 / 1.0))))) + 0.0)))))))))))
    switch v25 {
    case "a":
        for fi33 := 0; fi33 < 16; fi33++ {
            switch 1 {
            case 1:
                v34 := v2[fi33]
                _ = v34
            case 4294967295:
                for fi35 := 0; fi35 < 14; fi35++ {
                    var v36 string = "\n\t"
                    _ = fi35
                    _ = v36
                }
            }
            v37 := GenS1{F0: 42, F1: 0}
            v3 = (v3 ^ (v3 | -403))
            _ = fi33
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
