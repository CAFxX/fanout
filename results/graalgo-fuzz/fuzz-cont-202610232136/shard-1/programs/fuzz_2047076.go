package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int, p1 int64, p2 uint32) uint8 {
    if true {
        switch 42 {
        case 32768:
            switch p0 {
            case 32768:
                fmt.Println(p2, p0)
            case 607:
                var v7 int32 = 2147483647
                _ = v7
            }
        case 0:
            var v8 uint = 2
            _ = v8
        case 9223372036854775807:
            switch ((("a" + "hello") + "\n\t") + "\n\t") {
            case "foo bar":
                v9 := []int{p0, (p0 >> ((647) & 63)), (p0 % ((p0) | 1)), (p0 / ((p0) | 1))}
                _ = v9
            default:
                for fi10 := 0; fi10 < 3; fi10++ {
                    v11 := GenS1{F0: 2, F1: 0, F2: ""}
                    v12 := GenS1{F0: 42, F1: -86, F2: ""}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        }
        p0 = p0
        var v13 bool = false
        _ = v13
    }
    var v14 string = (("a" + "a") + ("foo bar" + "\n\t"))
    v15 := GenS1{F0: 1, F1: -1, F2: ""}
    p0 -= (p0 & 4294967295)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 255
}
func genFn1() int64 {
    var v16 uint = 42
    for fi17 := 0; fi17 < 9; fi17++ {
        switch "foo bar" {
        case "":
            var v18 uint32 = 2147483648
            _ = v18
        case "a":
            v16 = v16
        default:
            if true {
                if (false && false) {
                    v19 := GenS1{F0: 1, F1: 42, F2: "a"}
                    var v20 float64 = float64(-2.25)
                    _ = v19
                    _ = v20
                }
                for fi21 := 0; fi21 < 14; fi21++ {
                    v16 -= (v16 * v16)
                    _ = fi21
                }
                if (1 == -314) {
                    v22 := GenS1{F0: 2, F1: -128, F2: "foo bar"}
                    var v23 uint16 = 0
                    v24 := []int{}
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 15; fi25++ {
                        v26 := []int{}
                        var v27 uint = (^(v16))
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    v16 = 4294967295
                    v16 = (v16 * v16)
                }
            }
        }
        _ = fi17
    }
    _ = v16
    return int64(v16)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 4294967295
    v3 := make(map[string]string)
    var v4 int64 = 42
    var v5 int16 = -65
    v6 := make(map[int]bool)
    switch 0 {
    case 32767:
        var v28 bool = (0 >= 42)
        _ = v28
    }
    var v29 int8 = 127
    if (!(!(!false))) {
        v5 -= -1
        var v30 int = 32767
        if (!false) {
            fmt.Println(v30)
        }
        _ = v30
    } else {
        var v31 int8 = (v29 - (v29 >> ((2) & 63)))
        _ = v31
    }
    var v32 bool = ((^(v4)) <= genFn1())
    v33 := []int{32768, 1}
    var v34 string = (("\n\t" + "\n\t") + "")
    var v35 int = 2147483647
    _ = len(v33)
    if v32 {
        var v36 int = v35
        _ = v36
    } else {
        v37 := GenS2{F0: 127}
        if false {
            fmt.Println(v32)
            var v38 float32 = 1.5
            v39 := GenS2{F0: genFn0((-(v35)), -729, (v2 - (v2 << ((2) & 63))))}
            _ = v38
            _ = v39
        }
        _ = v37
    }
    if len(v33) >= 2 {
        v33 = v33[1:len(v33)]
    }
    var v40 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
