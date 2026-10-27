package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 string }
type GenS2 struct { F0 bool }
func genFn0() string {
    v4 := []int{958, -1, 127}
    var v5 uint32 = 520
    var v6 int = -1
    _ = v4
    _ = v5
    _ = v6
    return ""
}

func main() {
    var v1 float64 = 1.5
    var v2 float64 = v1
    v3 := []uint32{556, 2147483648}
    var v7 string = ""
    if (-128 != 100) {
        var v8 bool = false
        for fi9 := 0; fi9 < 15; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                fmt.Println(v2, fi9, fi9)
                v11 := GenS1{F0: 1, F1: 528, F2: v7}
                var v12 float64 = v1
                _ = fi10
                _ = v11
                _ = v12
            }
            v13 := GenS2{F0: true}
            fmt.Println(v2)
            _ = fi9
            _ = v13
        }
        _ = v8
    } else {
        switch ((v7 + v7) + v7) {
        case "hello":
            switch 989 {
            case 274:
                var v14 int64 = 42
                _ = v14
            }
        default:
            if len(v3) > 0 {
                v3[0] = 1
            }
        }
        fmt.Println(v2)
        for fi15 := 0; fi15 < 9; fi15++ {
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            if true {
                if (!((!true) && false)) {
                    var v16 string = v7
                    if false {
                        v16 = v7
                        v17 := GenS2{F0: ((true && false) && false)}
                        _ = v17
                    } else {
                        switch "hello" {
                        case "a":
                            var v18 uint16 = 32767
                            _ = v18
                        case "foo bar":
                            fmt.Println(v2, v1)
                        }
                    }
                    _ = v16
                } else {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        v7 = ""
                        _ = fi19
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v20 uint32 = 4294967295
                    _ = v20
                }
                if true {
                    if len(v3) > 0 {
                        v3[0] = 65535
                    }
                } else {
                    v21 := GenS1{F0: -9223372036854775808, F1: 1, F2: (v7 + v7)}
                    v1 = float64(1e10)
                    _ = v21
                }
                if true {
                    for fi22 := 0; fi22 < 7; fi22++ {
                        v23 := GenS2{F0: false}
                        var v24 string = "\n\t"
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    var v25 bool = false
                    for fi26 := 0; fi26 < 7; fi26++ {
                        var v27 int8 = 0
                        v1 -= float64((0.1 * ((1.5 - (0.0 - -2.25)) - 1e10)))
                        switch fi15 {
                        case 42:
                            switch "hello" {
                            case "hello":
                                v7 = v7
                            }
                        default:
                            v2 = v1
                        }
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                }
            }
            _ = fi15
        }
    }
    v28 := GenS2{F0: ((true || true) && (-1 == 1000))}
    var v29 int = 187
    v30 := GenS2{F0: (0 > 2147483647)}
    fmt.Println(v2, v29, v29)
    fmt.Println(v29, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 += v1
    v7 = ("foo bar" + "a")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v7
}
