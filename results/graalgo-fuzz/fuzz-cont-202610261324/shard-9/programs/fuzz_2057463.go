package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint64) string {
    var v5 uint32 = 0
    for fi6 := 0; fi6 < 3; fi6++ {
        if true {
            for fi7 := 0; fi7 < 19; fi7++ {
                switch -9223372036854775808 {
                case -2147483648:
                    p0 = (p0 >> ((2) & 63))
                    var v8 bool = (69 == 4294967295)
                    fmt.Println(p0, p0)
                    _ = v8
                case 2147483647:
                    var v9 uint64 = p0
                    _ = v9
                case 42:
                    v10 := GenS1{F0: -1, F1: 312}
                    _ = v10
                }
                _ = fi7
            }
            var v11 string = "foo bar"
            v12 := GenS1{F0: 42, F1: 18446744073709551615}
            _ = v11
            _ = v12
        } else {
            v13 := GenS2{F0: 1}
            v14 := GenS1{F0: -9223372036854775808, F1: 65535}
            _ = v13
            _ = v14
        }
        switch "hello" {
        case "foo bar":
            var v15 uint32 = v5
            _ = v15
        case "":
            var v16 float64 = ((0.0 * (1.5 * 0.0)) + -2.25)
            _ = v16
        case "\n\t":
            v17 := []int{(fi6 % ((2147483647) | 1)), (fi6 | (^(fi6))), (^(fi6)), 4294967295}
            _ = v17
        default:
            v18 := []int{fi6, (fi6 & -623)}
            _ = v18
        }
        v19 := []int{(^(fi6)), fi6}
        _ = fi6
        _ = v19
    }
    var v20 int = int(p0)
    var v21 uint8 = 127
    fmt.Println(v21, p0, v20)
    _ = p0
    _ = v20
    _ = v21
    _ = v5
    return "hello"
}
func genFn1() string {
    for fi22 := 0; fi22 < 14; fi22++ {
        v23 := GenS1{F0: 42, F1: 255}
        _ = fi22
        _ = v23
    }
    var v24 int64 = 42
    var v25 int = 929
    v26 := []int{}
    _ = v24
    _ = v25
    _ = v26
    return (("a" + "") + "")
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 65535
    var v3 uint32 = 173
    var v4 int64 = 9223372036854775807
    if ((v3 * v3) <= v3) {
        var v27 int8 = 0
        _ = v27
    } else {
        v3 = v3
        v4 += 477
    }
    if ((!(0.0 == 0.1)) && false) {
        switch 2147483647 {
        case -1:
            v1[-165] = 372
        case 1:
            v2 = v2
        }
    }
    switch -2147483648 {
    case -221:
        switch "" {
        case "\n\t":
            var v28 string = ((("hello" + "\n\t") + "") + ("hello" + ""))
            _ = v28
        }
    }
    if (v2 != (v2 + 4294967295)) {
        var v29 bool = (42 >= (v3 & v3))
        var v30 uint = v2
        switch 276 {
        case 9223372036854775807:
            v4 = -711
        case 669:
            var v31 int64 = v4
            _ = v31
        }
        _ = v29
        _ = v30
    } else {
        v32 := GenS1{F0: v4, F1: (v2 >> (((v2 & v2)) & 63))}
        _ = v32
    }
    fmt.Println(v3)
    for fi33 := 0; fi33 < 18; fi33++ {
        v4 = (v4 ^ (-(v4)))
        v34 := GenS2{F0: 1000}
        _ = fi33
        _ = v34
    }
    switch 1 {
    case -2147483648:
        var v35 float64 = 0.0
        _ = v35
    case 9223372036854775807:
        v36 := GenS2{F0: 0}
        _ = v36
    default:
        v37 := GenS2{F0: 0}
        _ = v37
    }
    var v38 float64 = 0.1
    var v39 int16 = -1
    var v40 float64 = v38
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
