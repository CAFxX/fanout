package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int16; F1 int }
func genFn0() bool {
    if true {
        var v5 float32 = float32(0.0)
        v6 := GenS1{F0: 0, F1: 836}
        v5 -= (v5 / v5)
        _ = v5
        _ = v6
    }
    var v7 int = 128
    _ = v7
    return true
}
func genFn1(p0 uint) uint16 {
    var v8 float64 = -2.25
    fmt.Println(v8)
    v9 := []int{}
    v10 := GenS1{F0: 626, F1: 1}
    if false {
        fmt.Println(v8, v8)
        p0 -= (^(p0))
    } else {
        var v11 uint32 = 4294967295
        _ = v11
    }
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 949
}

func main() {
    var v1 uint16 = 32767
    var v2 float64 = (0.1 - 0.0)
    var v3 int32 = 0
    var v4 uint = 426
    switch 305 {
    case 9223372036854775807:
        if (false && false) {
            if (!(128 <= 114)) {
                var v12 int64 = 42
                fmt.Println(v12, v2)
                for fi13 := 0; fi13 < 1; fi13++ {
                    switch (("a" + ("a" + "foo bar")) + ("hello" + "a")) {
                    case "":
                        v12 -= v12
                    default:
                        fmt.Println(v1, v4, v2)
                    }
                    _ = fi13
                }
                _ = v12
            }
            v14 := GenS2{F0: 670, F1: 128}
            _ = v14
        } else {
            var v15 uint = v4
            switch "foo bar" {
            case "":
                v16 := []int{}
                _ = v16
            case "\n\t":
                fmt.Println(v4, v1, v15)
            }
            for fi17 := 0; fi17 < 10; fi17++ {
                if true {
                    for fi18 := 0; fi18 < 10; fi18++ {
                        if (!(!true)) {
                            v2 -= v2
                            v2 = v2
                            fmt.Println(fi18)
                        } else {
                            var v19 uint16 = (v1 + (v1 + (v1 << (((^(v4))) & 63))))
                            if true {
                                v20 := GenS2{F0: -357, F1: -1}
                                v19 -= (^(v19))
                                v2 += 0.1
                                _ = v20
                            } else {
                                var v21 float32 = 0.1
                                v22 := GenS1{F0: 538, F1: v3}
                                var v23 int64 = 0
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            if false {
                                var v24 uint16 = (v1 % ((v19) | 1))
                                if false {
                                    v25 := []int{fi18, (fi17 + fi17)}
                                    _ = v25
                                }
                                _ = v24
                            }
                            _ = v19
                        }
                        _ = fi18
                    }
                } else {
                    v2 = v2
                    var v26 string = "foo bar"
                    var v27 int8 = -1
                    _ = v26
                    _ = v27
                }
                var v28 float64 = (v2 - v2)
                _ = fi17
                _ = v28
            }
            _ = v15
        }
    default:
        fmt.Println(v3, v4)
    }
    var v29 float32 = 0.1
    if false {
        v29 = float32(v2)
        switch 128 {
        case 128:
            var v30 uint32 = 42
            _ = v30
        case 1:
            var v31 int8 = 1
            _ = v31
        case 2147483647:
            v32 := GenS2{F0: 32767, F1: -217}
            _ = v32
        }
        fmt.Println(v29, v3)
    }
    v1 = (^(v1))
    if false {
        if true {
            var v33 string = ("\n\t" + ("a" + "a"))
            v33 = "foo bar"
            _ = v33
        }
    }
    v34 := GenS2{F0: -32768, F1: -1}
    v29 -= 3.14159
    v1 = (v1 * 42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v4
}
