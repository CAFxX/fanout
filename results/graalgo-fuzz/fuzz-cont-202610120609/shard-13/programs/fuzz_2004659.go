package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    var v3 uint = 434
    var v4 uint = v3
    var v5 int8 = -128
    if true {
        v4 -= (v3 / (((v3 | (v4 - 840))) | 1))
        var v6 int32 = 0
        _ = v6
    } else {
        v4 += v4
    }
    v5 -= 100
    fmt.Println(v3, v4)
    if (("foo bar" < ("a" + "\n\t")) && (!true)) {
        var v7 int8 = v5
        if true {
            var v8 float64 = 1e10
            _ = v8
        } else {
            fmt.Println(v4)
            v3 = v4
            var v9 int16 = -32768
            _ = v9
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS1{F0: 1, F1: (0 > 522), F2: 9223372036854775807}
            _ = fi10
            _ = v11
        }
        _ = v7
    }
    var v12 uint64 = 2
    v4 += (v4 % ((v3) | 1))
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 float64 = (-2.25 / 0.1)
        v15 := GenS1{F0: 1, F1: false, F2: 154}
        if (!true) {
            v14 = v14
        } else {
            switch fi13 {
            case 162:
                var v16 int64 = -134
                _ = v16
            case 4294967295:
                v5 -= 1
            default:
                v12 = 2
            }
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        switch (fi17 ^ fi17) {
        case -1:
            var v18 int32 = 2147483647
            _ = v18
        case 221:
            v5 = (^(v5))
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            switch ("foo bar" + ("foo bar" + "\n\t")) {
            case "a":
                var v20 uint16 = uint16(v4)
                _ = v20
            case "hello":
                delete(v2, 613)
            default:
                if (1 < 0) {
                    if false {
                        var v21 float64 = (float64(float32(3.14159)) + 1.5)
                        if (!(0 == -32768)) {
                            fmt.Println(v12)
                        }
                        _ = v21
                    }
                    var v22 float32 = -2.25
                    fmt.Println(fi19, v4)
                    switch (fi17 / (((fi19 | (-(fi19)))) | 1)) {
                    case -2147483648:
                        var v23 uint8 = 191
                        _ = v23
                    case -1:
                        var v24 int32 = 0
                        _ = v24
                    }
                    _ = v22
                }
            }
            _ = fi19
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            v5 = -67
            if (!true) {
                v26 := []int{2147483647, fi25}
                v27 := GenS1{F0: 0, F1: true, F2: -1}
                v28 := GenS1{F0: 162, F1: false, F2: 678}
                _ = v26
                _ = v27
                _ = v28
            } else {
                fmt.Println(v3, v3)
                switch 554 {
                case 844:
                    if false {
                        if (3.14159 == ((0.1 + -2.25) / (1.5 / 1e10))) {
                            var v29 bool = false
                            v12 = v12
                            fmt.Println(v12)
                            v29 = v29
                            v30 := []int{fi25}
                            _ = v29
                            _ = v30
                        } else {
                            v31 := GenS1{F0: 42, F1: (v4 >= (v3 - v4)), F2: 42}
                            v3 -= v3
                            _ = v31
                        }
                    }
                }
                v32 := GenS1{F0: 2147483647, F1: ((v5 - (v5 | (v5 << (((v4 * v4)) & 63)))) != 42), F2: 9223372036854775807}
                _ = v32
            }
            _ = fi25
        }
        _ = fi17
    }
    fmt.Println(v12, v4, v12)
    var v33 int16 = 1
    for fi34 := 0; fi34 < 20; fi34++ {
        var v35 uint = v3
        _ = fi34
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
