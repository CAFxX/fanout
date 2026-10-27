package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint16 }
type GenS2 struct { F0 string; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int) uint8 {
    var v5 int64 = -1
    v5 -= (v5 >> ((43) & 63))
    _ = p0
    _ = v5
    return 128
}
func genFn1(p0 uint16, p1 int8, p2 int8) int {
    if (1.5 >= 1e10) {
        v6 := []int{877}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 3; fi7++ {
            p1 += (p1 | (p2 >> ((255) & 63)))
            p2 = p1
            _ = fi7
        }
    }
    switch 42 {
    case 4294967295:
        if false {
            switch "foo bar" {
            case "hello":
                p0 = (p0 / (((p0 >> ((65535) & 63))) | 1))
            case "\n\t":
                var v8 uint8 = 1
                _ = v8
            default:
                for fi9 := 0; fi9 < 7; fi9++ {
                    p1 += 127
                    var v10 int = -1
                    var v11 bool = true
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        }
    case 32767:
        if (((!(false && true)) && false) && (!((false || (!(2147483647 >= 698))) && (false && (true && (p2 >= p2)))))) {
            if ((288 > 255) || ((0.0 != (1.5 / (((3.14159 + float64(0.0)) * (0.0 - 1.5)) / (3.14159 - 1e10)))) && true)) {
                var v12 int = -2147483648
                v13 := []int{-564, v12, (v12 - v12)}
                v14 := GenS1{F0: 0, F1: -32768, F2: 135}
                _ = v12
                _ = v13
                _ = v14
            }
            var v15 uint64 = 42
            v16 := GenS2{F0: "", F1: 500, F2: 56}
            _ = v15
            _ = v16
        } else {
            v17 := GenS2{F0: "a", F1: 42, F2: 132}
            var v18 float64 = float64(-2.25)
            v18 += v18
            _ = v17
            _ = v18
        }
    case -9223372036854775808:
        if false {
            if (!(0 >= 2147483647)) {
                if true {
                    p0 = 1
                    v19 := GenS2{F0: "", F1: -1, F2: 42}
                    _ = v19
                }
            } else {
                p2 = 127
            }
            var v20 uint64 = 1
            fmt.Println(p0, p2, v20)
            _ = v20
        } else {
            for fi21 := 0; fi21 < 4; fi21++ {
                p2 += 100
                _ = fi21
            }
        }
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int = 0
    v2 := []int{(v1 + (v1 / (((v1 ^ v1)) | 1))), 128, v1}
    var v3 int64 = -328
    var v4 int16 = 32767
    var v22 bool = (!true)
    if false {
        var v23 int = (v1 >> ((2) & 63))
        var v24 float32 = 3.14159
        v24 = v24
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        if (((("\n\t" + "foo bar") + ("hello" + ("a" + "\n\t"))) + "a") >= "a") {
            var v26 int = (-(fi25))
            _ = v26
        }
        for fi27 := 0; fi27 < 12; fi27++ {
            fmt.Println(v3, v1, v4)
            _ = fi27
        }
        if (!v22) {
            if (genFn0((v1 * 0)) < 1) {
                var v28 float32 = (0.1 / 1.0)
                var v29 int8 = 1
                _ = v28
                _ = v29
            }
            var v30 float64 = -2.25
            _ = v30
        }
        _ = fi25
    }
    v31 := GenS1{F0: (v3 >> ((255) & 63)), F1: (-(v4)), F2: 32767}
    v22 = (!v22)
    fmt.Println(v3)
    var v32 string = "foo bar"
    v2 = append(v2, 32768)
    v33 := GenS1{F0: (v3 / ((397) | 1)), F1: v4, F2: 255}
    for fi34 := 0; fi34 < 15; fi34++ {
        if (152 < 1) {
            v32 = v32
            v35 := GenS3{F0: 65535}
            var v36 uint8 = genFn0((fi34 << ((185) & 63)))
            _ = v35
            _ = v36
        }
        fmt.Println(v22)
        _ = fi34
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
