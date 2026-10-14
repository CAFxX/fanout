package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint16 }
func genFn0(p0 uint32) bool {
    p0 += (^(p0))
    v4 := []int{0, 42, 1}
    _ = p0
    _ = v4
    return (false || false)
}
func genFn1(p0 uint32, p1 uint16) int32 {
    var v5 int8 = -128
    var v6 uint16 = (^(p1))
    fmt.Println(v6)
    switch 128 {
    case 32767:
        v7 := GenS1{F0: false, F1: (false && (42 != 101))}
        _ = v7
    case 127:
        if false {
            v5 += 127
            var v8 float32 = 1.5
            _ = v8
        }
    default:
        var v9 uint32 = 421
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 314
    var v3 uint32 = 2147483647
    if (!false) {
        v10 := v1[42]
        _ = v10
    } else {
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := []int{978, (fi11 | -733), (fi11 | fi11), fi11}
            _ = fi11
            _ = v12
        }
        if true {
            var v13 float64 = (0.0 * 1.5)
            if true {
                var v14 int64 = int64(v3)
                _ = v14
            } else {
                v15 := GenS1{F0: true, F1: false}
                _ = v15
            }
            _ = v13
        }
        v16 := v1[v3]
        _ = v16
    }
    switch 0 {
    case -797:
        v17 := GenS1{F0: (-9223372036854775808 <= 0), F1: false}
        _ = v17
    }
    v3 = 2
    var v18 uint16 = 192
    var v19 string = (((("a" + ("\n\t" + "\n\t")) + (("" + ("\n\t" + "")) + "foo bar")) + "hello") + "hello")
    switch v19 {
    case "foo bar":
        var v20 string = (("" + "\n\t") + v19)
        _ = v20
    case "hello":
        var v21 uint64 = 42
        _ = v21
    case "":
        if false {
            for fi22 := 0; fi22 < 18; fi22++ {
                v23 := GenS2{F0: 1}
                _ = fi22
                _ = v23
            }
            v3 = v3
        } else {
            v24 := []int{1, 128}
            delete(v1, 2147483648)
            v19 = (("hello" + (v19 + (v19 + ("\n\t" + ("a" + (v19 + ("" + "hello"))))))) + "a")
            _ = v24
        }
    }
    switch 42 {
    case 46:
        v25 := []int{}
        _ = v25
    }
    if (genFn1(4294967295, v18) == genFn1(2, v18)) {
        var v26 string = "foo bar"
        for fi27 := 0; fi27 < 2; fi27++ {
            v19 = v19
            var v28 int16 = -316
            v29 := []int{}
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = v26
    } else {
        v30 := GenS1{F0: ((((-2.25 * -2.25) + 0.0) * 1e10) > -2.25), F1: ((!(3.14159 != float32((1.5 + ((float64(((1.5 * 0.1) * 1.5)) / 1.5) * (((float64(-2.25) + 1e10) - 1e10) + float64(-2.25))))))) && (896 != 118))}
        _ = v30
    }
    v18 = 255
    for fi31 := 0; fi31 < 7; fi31++ {
        fmt.Println(fi31, v2, v3)
        v32 := GenS3{F0: false, F1: v18, F2: (v18 & 252)}
        var v33 uint = v2
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 uint32 = (v3 >> (((v2 / ((1) | 1))) & 63))
    if (false && false) {
        var v35 uint64 = 42
        v36 := []int{32767}
        _ = v35
        _ = v36
    }
    fmt.Println(v19, v2)
    v34 = v3
    for fi37 := 0; fi37 < 12; fi37++ {
        v38 := []int{244}
        v18 += (v18 & (v18 * (^(v18))))
        _ = fi37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
}
