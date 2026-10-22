package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 string; F1 int16 }
func genFn0(p0 uint32, p1 bool, p2 int64) string {
    fmt.Println(p2)
    v4 := []int{127, 2147483647}
    p1 = (2 > 1)
    var v5 int32 = -213
    p1 = (-28 != -2147483648)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (("a" + (("hello" + "\n\t") + (("a" + "\n\t") + "foo bar"))) + "hello")
}
func genFn1() uint64 {
    var v6 uint16 = 989
    fmt.Println(v6)
    var v7 uint8 = 42
    v8 := GenS1{F0: uint8(v6), F1: "hello", F2: ((!(true && true)) && false)}
    var v9 uint16 = (v6 >> ((506) & 63))
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int8 = -104
    v2 := make(map[string]string)
    var v3 int64 = -200
    v3 = v3
    fmt.Println(v1)
    if false {
        v10 := []int{42, 128}
        if true {
            v3 -= (-(v3))
            if (v3 == v3) {
                if true {
                    v1 = (v1 / ((v1) | 1))
                }
                var v11 string = (("a" + (("" + ("a" + "hello")) + "\n\t")) + "\n\t")
                _ = v11
            } else {
                v1 = (v1 ^ (v1 | -1))
            }
        } else {
            fmt.Println(v1, v1)
        }
        switch ("" + ("\n\t" + "hello")) {
        case "":
            v1 += v1
        case "a":
            var v12 float32 = float32(((3.14159 + 0.1) + 3.14159))
            _ = v12
        default:
            var v13 string = "a"
            _ = v13
        }
        _ = v10
    }
    fmt.Println(v1, v1)
    if true {
        v14 := GenS2{F0: 18446744073709551615}
        v1 -= (v1 | v1)
        _ = v14
    } else {
        switch 1 {
        case 127:
            var v15 uint = 575
            _ = v15
        case -9223372036854775808:
            v16 := GenS3{F0: 954, F1: v3}
            _ = v16
        }
        var v17 float32 = -2.25
        _ = v17
    }
    v1 -= (-(v1))
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 5; fi19++ {
            if false {
                v20 := []int{fi19, (fi19 >> ((745) & 63))}
                if (508 == 164) {
                    var v21 uint8 = 255
                    var v22 float64 = ((0.0 * -2.25) + 1e10)
                    _ = len(v2)
                    if (0.0 != 0.1) {
                        if false {
                            v23 := GenS4{F0: ("\n\t" + ("\n\t" + "foo bar")), F1: 42}
                            _ = v23
                        }
                    }
                    _ = v21
                    _ = v22
                }
                var v24 float32 = (-2.25 * 1.5)
                _ = v20
                _ = v24
            }
            _ = fi19
        }
        _ = fi18
    }
    v25 := GenS1{F0: 2, F1: "a", F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
}
