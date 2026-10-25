package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 bool; F2 int }
func genFn0(p0 uint16, p1 int64, p2 uint) uint8 {
    v5 := GenS1{F0: (p1 | -9223372036854775808), F1: (!(3.14159 >= -2.25)), F2: 127}
    if (0 >= -1) {
        v6 := []int{}
        _ = v6
    }
    var v7 uint64 = 410
    switch 4294967295 {
    case -165:
        p1 = 42
    case -30:
        var v8 float32 = (0.1 / ((0.0 - ((3.14159 - (1e10 / -2.25)) - float32(0.1))) - -2.25))
        _ = v8
    default:
        if true {
            fmt.Println(v7, p1, p0)
            v7 = uint64(p2)
        } else {
            p2 = 18
            fmt.Println(v7, v7)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 128
}
func genFn1() bool {
    if ((true && (9223372036854775807 < -9223372036854775808)) && true) {
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := GenS2{F0: 0}
            var v11 string = ("foo bar" + "\n\t")
            _ = fi9
            _ = v10
            _ = v11
        }
        if (2147483647 >= 32768) {
            var v12 bool = (32767 < 301)
            for fi13 := 0; fi13 < 19; fi13++ {
                v12 = (fi13 == fi13)
                v12 = false
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 uint32 = 1
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            if (!(0 != -1)) {
                switch (("" + "\n\t") + "hello") {
                case "\n\t":
                    v16 := []int{-1, 4294967295}
                    _ = v16
                case "foo bar":
                    v12 = (!(((0.0 / 1.5) * ((3.14159 + 0.1) - -2.25)) > 3.14159))
                }
            }
            _ = v12
        } else {
            if (284 == 2) {
                if false {
                    var v17 int8 = 0
                    if (1 > 42) {
                        var v18 uint16 = 32767
                        _ = v18
                    } else {
                        var v19 uint = 0
                        v20 := GenS3{F0: 128}
                        var v21 uint32 = 1
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    _ = v17
                }
            } else {
                var v22 float64 = 1.5
                _ = v22
            }
            v23 := []int{32767}
            var v24 string = "hello"
            _ = v23
            _ = v24
        }
        v25 := []int{}
        _ = v25
    } else {
        var v26 string = ("\n\t" + "foo bar")
        v27 := GenS4{F0: 123}
        v28 := GenS5{F0: 42, F1: false, F2: 4294967295}
        _ = v26
        _ = v27
        _ = v28
    }
    switch 9223372036854775807 {
    case 128:
        v29 := []int{1, 127, 1}
        _ = v29
    }
    return ((255 < 0) || true)
}

func main() {
    var v1 bool = (0.0 > 3.14159)
    var v2 int16 = -267
    var v3 int32 = 42
    var v4 uint8 = 1
    v4 += (v4 ^ 2)
    v4 = 255
    fmt.Println(v2)
    var v30 int64 = -230
    v30 += (v30 ^ (v30 % (((v30 % ((v30) | 1))) | 1)))
    var v31 float64 = ((0.0 + (1e10 * (-2.25 * (1e10 - 1.5)))) - 0.1)
    var v32 float32 = (0.1 / ((0.1 / 1e10) * 0.1))
    var v33 float64 = ((v31 + v31) * ((v31 - float64(float32(float64(((v32 / (1e10 / v32)) + v32))))) - float64(float32((v31 / v31)))))
    for fi34 := 0; fi34 < 6; fi34++ {
        v4 = (^(v4))
        var v35 bool = v1
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
