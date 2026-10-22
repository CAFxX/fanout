package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0(p0 int32) string {
    var v4 int32 = p0
    v5 := []int{}
    _ = p0
    _ = v4
    _ = v5
    return "hello"
}
func genFn1() uint8 {
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int8 = -54
            if false {
                v9 := []int{986}
                if false {
                    var v10 uint = 476
                    _ = v10
                }
                var v11 float64 = float64(-2.25)
                _ = v11
                _ = v9
            } else {
                switch (fi7 - fi6) {
                case -2147483648:
                    v12 := GenS1{F0: 2, F1: 285}
                    _ = v12
                }
                fmt.Println(fi7, fi7)
            }
            if (false || (32767 <= -4)) {
                v8 = (v8 ^ (v8 ^ 42))
                var v13 string = ((("\n\t" + "foo bar") + "foo bar") + ((("" + "") + (("" + "foo bar") + "\n\t")) + "hello"))
                fmt.Println(v8, v13)
                _ = v13
            }
            _ = fi7
            _ = v8
        }
        var v14 float64 = 0.0
        fmt.Println(fi6, fi6)
        _ = fi6
        _ = v14
    }
    v15 := []int{32767, 32767, -1, 210}
    if false {
        var v16 int = 0
        _ = v16
    } else {
        var v17 string = ((("" + "\n\t") + "a") + (("hello" + "") + ("hello" + "a")))
        v17 = "hello"
        var v18 uint64 = 18446744073709551615
        _ = v17
        _ = v18
    }
    var v19 float64 = -2.25
    _ = v15
    _ = v19
    return 1
}

func main() {
    var v1 int32 = 458
    v2 := make(map[int]string)
    var v3 uint = 0
    var v20 int8 = 0
    var v21 int32 = v1
    var v22 uint64 = 18446744073709551615
    switch "a" {
    case "foo bar":
        var v23 int8 = 100
        _ = v23
    default:
        var v24 int16 = 1000
        _ = v24
    }
    var v25 uint64 = (v22 % (((v22 ^ (v22 << ((v3) & 63)))) | 1))
    switch 0 {
    case 32767:
        var v26 string = (("hello" + "a") + "a")
        _ = v26
    default:
        if false {
            for fi27 := 0; fi27 < 20; fi27++ {
                switch (-(fi27)) {
                case -512:
                    if (true && false) {
                        v3 = v3
                        fmt.Println(v3)
                    } else {
                        v28 := []int{(fi27 & fi27), (fi27 % ((42) | 1)), fi27}
                        var v29 uint16 = uint16(v3)
                        _ = v28
                        _ = v29
                    }
                default:
                    var v30 float64 = 1.5
                    _ = v30
                }
                var v31 float32 = 0.1
                v25 -= (v25 - v22)
                _ = fi27
                _ = v31
            }
        } else {
            _ = len(v2)
            v32 := GenS2{F0: 42, F1: -115}
            if false {
                v33 := []int{}
                for fi34 := 0; fi34 < 2; fi34++ {
                    var v35 float32 = -2.25
                    _ = fi34
                    _ = v35
                }
                _ = v33
            } else {
                var v36 uint64 = (v22 * v25)
                _ = v36
            }
            _ = v32
        }
    }
    if (v3 < v3) {
        var v37 int = 4294967295
        var v38 int32 = 1
        _ = v37
        _ = v38
    } else {
        v1 += (-(v21))
        v21 -= -1
    }
    v21 -= int32(v22)
    switch "\n\t" {
    case "hello":
        var v39 float64 = 0.1
        _ = v39
    case "":
        v40 := []int{808}
        _ = v40
    case "foo bar":
        if (((("" + "hello") + "a") + ("" + "\n\t")) >= "a") {
            var v41 int8 = (v20 - (v20 << ((2) & 63)))
            v21 = 924
            var v42 bool = true
            _ = v41
            _ = v42
        }
    default:
        v43 := v2[32767]
        _ = v43
    }
    switch ("hello" + "foo bar") {
    case "":
        v20 = (v20 - (v20 * v20))
    }
    fmt.Println(v21, v25, v25)
    v44 := GenS2{F0: 2, F1: (v20 + v20)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v44
}
