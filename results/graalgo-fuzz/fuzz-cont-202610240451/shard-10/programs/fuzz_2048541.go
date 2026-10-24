package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int64; F1 bool; F2 int8 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int64 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]bool)
    var v3 uint16 = 42
    var v4 int8 = -14
    if false {
        for fi5 := 0; fi5 < 4; fi5++ {
            switch ("a" + "") {
            case "hello":
                v3 -= (^(v3))
            case "\n\t":
                var v6 uint32 = 928
                _ = v6
            case "":
                v7 := GenS1{F0: 42, F1: v3}
                _ = v7
            default:
                v1 -= v1
            }
            _ = fi5
        }
        switch 510 {
        case 789:
            var v8 uint32 = 2147483647
            _ = v8
        case 32768:
            v9 := GenS2{F0: 435, F1: (true || ((4294967295 != 2147483647) && true)), F2: v4}
            _ = v9
        case 128:
            if (v1 < (v1 * 42)) {
                v4 -= v4
                v4 -= (v4 * v4)
                v3 += 42
            } else {
                var v10 uint64 = 979
                v4 -= (v4 | -8)
                v11 := []int{4294967295, -9223372036854775808, 523, 32767}
                _ = v10
                _ = v11
            }
        default:
            for fi12 := 0; fi12 < 2; fi12++ {
                fmt.Println(v1)
                fmt.Println(v1, v1)
                v4 = (v4 ^ v4)
                _ = fi12
            }
        }
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        if false {
            switch ("foo bar" + "") {
            case "a":
                v3 -= v3
            }
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 int32 = 42
            if (false && (2 >= 128)) {
                if (!false) {
                    var v16 uint8 = 255
                    switch (fi13 | (-(fi13))) {
                    case 32768:
                        if (v3 == (^(v3))) {
                            var v17 bool = false
                            v18 := GenS1{F0: 0, F1: v3}
                            v16 -= v16
                            _ = v17
                            _ = v18
                        } else {
                            var v19 int16 = -32768
                            _ = v19
                        }
                    default:
                        v20 := GenS3{F0: "a", F1: true}
                        _ = v20
                    }
                    _ = v16
                }
            }
            switch fi14 {
            case 436:
                v4 = v4
                var v21 uint = 18446744073709551615
                v21 = v21
                _ = v21
            case 42:
                switch fi14 {
                case 128:
                    v4 = v4
                }
            default:
                v2[0] = true
            }
            _ = fi14
            _ = v15
        }
        v22 := GenS4{F0: 0}
        _ = fi13
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        switch (fi23 % ((fi23) | 1)) {
        case 23:
            v24 := []int{(-(fi23)), fi23, fi23}
            _ = v24
        case -232:
            switch fi23 {
            case -135:
                v25 := GenS1{F0: 101, F1: v3}
                _ = v25
            }
        default:
            fmt.Println(v4, fi23, fi23)
        }
        _ = fi23
    }
    var v26 uint8 = 255
    v4 += v4
    if true {
        var v27 uint64 = 42
        var v28 uint8 = (v26 % ((v26) | 1))
        v29 := []int{}
        _ = v27
        _ = v28
        _ = v29
    }
    v30 := GenS2{F0: 1, F1: true, F2: 62}
    v1 += v1
    switch 223 {
    case -1:
        fmt.Println(v26)
    default:
        for fi31 := 0; fi31 < 11; fi31++ {
            fmt.Println(v26, v4)
            _ = fi31
        }
    }
    v3 = 2
    if (("" + (("foo bar" + "") + "foo bar")) > (("foo bar" + "foo bar") + "hello")) {
        var v32 float32 = 1e10
        _ = v32
    }
    var v33 uint = 2
    var v34 string = ("hello" + ("a" + "a"))
    switch int(v4) {
    case 114:
        for fi35 := 0; fi35 < 14; fi35++ {
            switch (fi35 * fi35) {
            case 9223372036854775807:
                v36 := []int{fi35, fi35, (fi35 % (((-(fi35))) | 1))}
                _ = v36
            default:
                v1 -= v1
            }
            v3 = 1
            _ = fi35
        }
    case 32768:
        if true {
            v37 := GenS4{F0: int64(v4)}
            v3 = 68
            _ = v37
        }
    case 42:
        var v38 int16 = 22
        _ = v38
    default:
        var v39 uint64 = (^(v1))
        _ = v39
    }
    v4 += (v4 | 0)
    switch 0 {
    case 9223372036854775807:
        var v40 int = 671
        _ = v40
    case 779:
        for fi41 := 0; fi41 < 2; fi41++ {
            v42 := GenS4{F0: -9223372036854775808}
            _ = fi41
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
}
