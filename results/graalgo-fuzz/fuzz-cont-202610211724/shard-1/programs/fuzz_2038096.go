package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS7 struct { F0 uint32; F1 uint32; F2 uint16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    var v3 float32 = (1e10 / (0.1 * 1e10))
    var v4 uint32 = 2
    var v5 float32 = float32(-2.25)
    if true {
        for fi6 := 0; fi6 < 18; fi6++ {
            switch (("a" + (("hello" + "foo bar") + (("hello" + "hello") + ("\n\t" + "a")))) + (((("hello" + ("\n\t" + "a")) + "") + "a") + ("\n\t" + "\n\t"))) {
            case "":
                v4 = v4
            case "hello":
                fmt.Println(fi6, v3)
            }
            var v7 string = "\n\t"
            var v8 int32 = -2147483648
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: (255 == 2), F1: -374}
        v10 := GenS2{F0: (false || false), F1: true}
        _ = v10
        _ = v9
    } else {
        var v11 float64 = 1.5
        v2[127] = (790 != 909)
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        switch (fi12 << ((255) & 63)) {
        case -2147483648:
            var v13 int32 = -1
            _ = v13
        case 128:
            var v14 float32 = v3
            _ = v14
        case 42:
            var v15 float32 = -2.25
            _ = v15
        default:
            fmt.Println(fi12, v4)
        }
        fmt.Println(v4)
        v16 := GenS3{F0: 4294967295, F1: 127}
        _ = fi12
        _ = v16
    }
    var v17 int8 = 119
    v18 := GenS3{F0: 1, F1: v17}
    if true {
        v4 += v4
    }
    var v19 int = 926
    v19 += (v19 >> ((255) & 63))
    var v20 int32 = -1
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            switch "a" {
            case "\n\t":
                v23 := GenS4{F0: v4, F1: 0, F2: 681}
                v17 -= 100
                var v24 float32 = (v3 + -2.25)
                _ = v23
                _ = v24
            default:
                v3 = v3
            }
            if true {
                v25 := GenS2{F0: (true && (2 < 2)), F1: (2 < 668)}
                v19 += v19
                v4 -= 65535
                _ = v25
            } else {
                if true {
                    v26 := GenS5{F0: -357}
                    v17 -= v17
                    _ = v26
                }
                switch "" {
                case "\n\t":
                    fmt.Println(fi21, v3, v17)
                case "a":
                    delete(v1, 2147483647)
                default:
                    v27 := GenS2{F0: (!false), F1: (3.14159 >= v5)}
                    _ = v27
                }
            }
            _ = fi22
        }
        if ((74 > int64(fi21)) || (true && true)) {
            if (true && ((^(v20)) < -2147483648)) {
                for fi28 := 0; fi28 < 4; fi28++ {
                    if true {
                        var v29 uint8 = 0
                        _ = v29
                    }
                    v30 := []int{(-(fi21)), 1}
                    v4 = (v4 | (v4 | (^(v4))))
                    _ = fi28
                    _ = v30
                }
            }
            if true {
                fmt.Println(v20, v19)
                var v31 uint16 = 1
                v32 := GenS6{F0: 116, F1: 1, F2: (v4 & v4)}
                _ = v31
                _ = v32
            }
        } else {
            var v33 uint = 727
            if (!false) {
                switch "foo bar" {
                case "":
                    v19 += -9223372036854775808
                case "foo bar":
                    fmt.Println(v4, fi21)
                case "a":
                    v20 += (-(v20))
                }
            }
            v5 = 0.0
            _ = v33
        }
        if ((float32((1.5 / 1.0)) + (1.5 - v5)) < 1e10) {
            var v34 uint = 137
            if true {
                v35 := GenS5{F0: 42}
                var v36 int64 = 27
                v34 = (v34 ^ v34)
                _ = v35
                _ = v36
            }
            _ = v34
        } else {
            v37 := GenS5{F0: -572}
            v17 = v17
            _ = v37
        }
        _ = fi21
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        fmt.Println(v17, v5, v5)
        v20 -= -2147483648
        _ = fi38
    }
    for fi39 := 0; fi39 < 19; fi39++ {
        v40 := []int{v19, v19, 42, fi39}
        _ = fi39
        _ = v40
    }
    if true {
        fmt.Println(v5)
        if true {
            v2[-789] = false
            for fi41 := 0; fi41 < 11; fi41++ {
                v42 := []int{}
                fmt.Println(fi41, v17)
                _ = fi41
                _ = v42
            }
            v43 := GenS5{F0: 32767}
            _ = v43
        }
        v5 += (v3 - ((0.0 + v3) + 1.5))
    } else {
        v44 := GenS7{F0: 0, F1: v4, F2: 32767}
        var v45 uint = 88
        _ = v44
        _ = v45
    }
    fmt.Println(v3)
    var v46 uint32 = (v4 + v4)
    var v47 uint8 = 85
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
