package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }
type GenS2 struct { F0 string; F1 uint64 }
func genFn0(p0 uint32, p1 string, p2 uint32) int {
    if (32767 < -32768) {
        p2 += p0
        var v4 int16 = -32768
        _ = v4
    }
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(fi5, p0, p1)
        p1 = "hello"
        _ = fi5
    }
    var v6 bool = (0 > 0)
    if true {
        v7 := []int{127, -876, 23}
        v8 := GenS1{F0: 2, F1: -495, F2: (true && v6)}
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -282
}

func main() {
    var v1 uint64 = 1
    var v2 int8 = 115
    var v3 float64 = -2.25
    fmt.Println(v1, v2, v2)
    switch ("hello" + "") {
    case "hello":
        var v9 float64 = v3
        _ = v9
    }
    var v10 int32 = 42
    var v11 int = 32768
    var v12 int8 = -128
    if false {
        var v13 uint64 = v1
        v14 := GenS1{F0: 54, F1: 69, F2: (((v3 / v3) >= v3) || false)}
        v15 := GenS2{F0: "", F1: (^(v1))}
        _ = v13
        _ = v14
        _ = v15
    } else {
        if ((!false) || (!false)) {
            for fi16 := 0; fi16 < 18; fi16++ {
                if (true && false) {
                    v10 += v10
                    v17 := GenS2{F0: ("\n\t" + ("hello" + "foo bar")), F1: (v1 | (v1 ^ 0))}
                    _ = v17
                }
                v2 = v12
                v18 := []int{128, fi16}
                _ = fi16
                _ = v18
            }
            v19 := []int{v11}
            _ = v19
        } else {
            for fi20 := 0; fi20 < 18; fi20++ {
                fmt.Println(fi20, v2)
                fmt.Println(fi20, v11, fi20)
                v21 := []int{}
                _ = fi20
                _ = v21
            }
        }
        v22 := []int{v11, (v11 - genFn0(2, "", 842)), v11}
        var v23 float64 = (1e10 + ((0.1 - v3) - ((v3 - float64(0.0)) + 0.0)))
        _ = v22
        _ = v23
    }
    var v24 float64 = v3
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 9; fi26++ {
            v27 := []int{genFn0(486, "\n\t", 646), int(v10), -1, v11}
            if (true || (false || false)) {
                var v28 int32 = (v10 | (-(v10)))
                _ = v28
            } else {
                if true {
                    var v29 int8 = v12
                    fmt.Println(v1, v2, v3)
                    v10 = 154
                    _ = v29
                } else {
                    if (!false) {
                        if (!true) {
                            v30 := GenS2{F0: "hello", F1: (v1 - v1)}
                            _ = v30
                        } else {
                            v27 = append(v27, 1)
                        }
                        var v31 uint64 = 2
                        var v32 uint8 = 60
                        v27 = append(v27, (v11 | (v11 - (^(v11)))))
                        _ = v31
                        _ = v32
                    } else {
                        v27 = append(v27, (-(fi25)))
                        v3 += (v24 / v3)
                        v33 := GenS2{F0: "a", F1: 1}
                        _ = v33
                    }
                    var v34 uint8 = 128
                    _ = v34
                }
                var v35 float32 = (1.5 * 0.0)
                _ = v35
            }
            v3 = v3
            _ = fi26
            _ = v27
        }
        v36 := GenS2{F0: (("a" + "a") + ((("\n\t" + (("hello" + "\n\t") + ("" + ("foo bar" + "hello")))) + "foo bar") + "hello")), F1: v1}
        _ = fi25
        _ = v36
    }
    var v37 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v37
}
