package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
func genFn0() uint16 {
    var v5 int = 32767
    v5 = 127
    if (9223372036854775807 <= -9223372036854775808) {
        switch "" {
        case "hello":
            v6 := []int{}
            _ = v6
        case "foo bar":
            var v7 float32 = float32(float64(0.1))
            _ = v7
        case "":
            v5 = (v5 % ((v5) | 1))
        default:
            var v8 int64 = 1
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 10; fi9++ {
            if false {
                v10 := []int{}
                for fi11 := 0; fi11 < 20; fi11++ {
                    v12 := GenS1{F0: (((-1 > 42) && true) || (2147483647 == fi11)), F1: 2}
                    v5 = (v5 >> ((139) & 63))
                    v13 := GenS1{F0: (true && (2147483647 > 4294967295)), F1: 2147483647}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v5 = v5
                _ = v10
            }
            if false {
                fmt.Println(v5)
                for fi14 := 0; fi14 < 20; fi14++ {
                    v15 := []int{(v5 & 42)}
                    _ = fi14
                    _ = v15
                }
            } else {
                switch "hello" {
                case "hello":
                    v5 = (v5 ^ -114)
                case "\n\t":
                    v16 := []int{}
                    _ = v16
                default:
                    var v17 float64 = -2.25
                    _ = v17
                }
                v5 = (fi9 + fi9)
            }
            v5 -= (fi9 + fi9)
            _ = fi9
        }
        var v18 uint16 = 65535
        _ = v18
    }
    _ = v5
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = float32((-2.25 / 1.0))
    v3 := make(map[uint32]string)
    var v4 int32 = 0
    v4 = v4
    if (((v2 * v2) > (((0.0 * (((float32(0.0) - v2) + v2) - v2)) * (0.1 * (3.14159 * v2))) - v2)) || false) {
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 int16 = 1000
            _ = fi19
            _ = v20
        }
        v2 = (((float32(1e10) + v2) / float32(3.14159)) + v2)
        v2 += 0.0
    }
    _ = len(v3)
    v21 := GenS1{F0: false, F1: 2147483648}
    _ = len(v1)
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 float64 = float64(float32(((0.0 * 0.0) + 1.5)))
        v4 = v4
        var v24 uint8 = 251
        _ = fi22
        _ = v23
        _ = v24
    }
    v2 -= 1e10
    fmt.Println(v4, v4)
    fmt.Println(v4, v2)
    var v25 float32 = v2
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 string = ("foo bar" + "a")
        var v28 int8 = 127
        v29 := v3[0]
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    var v30 int32 = v4
    var v31 uint = 42
    var v32 float64 = -2.25
    for fi33 := 0; fi33 < 3; fi33++ {
        if true {
            v31 -= 379
        } else {
            switch -1 {
            case -2147483648:
                fmt.Println(v4, v31, v25)
            case -953:
                v34 := v3[589]
                _ = v34
            case 455:
                v35 := []int{fi33}
                _ = v35
            }
        }
        var v36 uint = (v31 + (^(v31)))
        _ = len(v1)
        _ = fi33
        _ = v36
    }
    var v37 bool = false
    var v38 uint8 = 0
    var v39 uint32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
