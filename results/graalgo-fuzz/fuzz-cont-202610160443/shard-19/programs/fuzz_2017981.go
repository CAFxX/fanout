package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint32; F2 bool }
func genFn0(p0 int64, p1 uint16, p2 uint64) int32 {
    for fi4 := 0; fi4 < 7; fi4++ {
        if false {
            v5 := []int{}
            p1 = p1
            for fi6 := 0; fi6 < 2; fi6++ {
                fmt.Println(p0)
                _ = fi6
            }
            _ = v5
        }
        _ = fi4
    }
    var v7 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return int32(p1)
}
func genFn1() int {
    v8 := []int{-9223372036854775808, 117}
    if true {
        if true {
            if true {
                var v9 uint32 = 2
                v9 += (v9 % ((v9) | 1))
                _ = v9
            }
        } else {
            var v10 int32 = 2147483647
            for fi11 := 0; fi11 < 20; fi11++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                v10 += (v10 + (v10 ^ (v10 * genFn0(0, 255, 888))))
                _ = fi11
            }
            _ = v10
        }
    } else {
        var v12 uint8 = 127
        for fi13 := 0; fi13 < 4; fi13++ {
            for fi14 := 0; fi14 < 4; fi14++ {
                v15 := GenS1{F0: ("a" + ("hello" + ("a" + "")))}
                var v16 int = fi13
                if true {
                    if (!false) {
                        v12 -= v12
                        var v17 int = 4294967295
                        _ = v17
                    } else {
                        var v18 int64 = 373
                        var v19 uint8 = v12
                        _ = v18
                        _ = v19
                    }
                } else {
                    v20 := GenS2{F0: 2147483647, F1: 421, F2: true}
                    _ = len(v8)
                    _ = v20
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi13
        }
        _ = v12
    }
    var v21 string = (("" + ("\n\t" + ("\n\t" + ""))) + ("a" + "hello"))
    _ = v21
    _ = v8
    return -656
}

func main() {
    v1 := []string{"hello", "a"}
    var v2 int64 = 9223372036854775807
    var v3 uint = 4294967295
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v22 bool = false
    if (!v22) {
        for fi23 := 0; fi23 < 14; fi23++ {
            v24 := GenS2{F0: 1, F1: 2147483647, F2: true}
            if (v22 || (1 > 842)) {
                v2 = v2
                v22 = (float64(-2.25) != (-2.25 - (float64(3.14159) / 3.14159)))
            } else {
                v25 := GenS2{F0: 390, F1: 42, F2: (0 < 32767)}
                v26 := GenS1{F0: (("hello" + ("foo bar" + "\n\t")) + "")}
                var v27 int = genFn1()
                _ = v25
                _ = v26
                _ = v27
            }
            _ = len(v1)
            _ = fi23
            _ = v24
        }
    }
    v22 = (false && v22)
    v28 := GenS1{F0: (("hello" + "") + ("hello" + "a"))}
    for fi29 := 0; fi29 < 8; fi29++ {
        if (255 < 0) {
            v2 -= 0
        } else {
            fmt.Println(fi29, v3, v22)
            fmt.Println(v3)
        }
        var v30 int = int(v3)
        fmt.Println(v2)
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        switch "a" {
        case "foo bar":
            fmt.Println(fi31)
        case "hello":
            v2 += (v2 ^ int64(v3))
        }
        v2 = (v2 - v2)
        v32 := GenS1{F0: "\n\t"}
        _ = fi31
        _ = v32
    }
    var v33 int64 = v2
    v34 := GenS1{F0: (("foo bar" + "a") + ("" + "\n\t"))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v33
    _ = v34
}
