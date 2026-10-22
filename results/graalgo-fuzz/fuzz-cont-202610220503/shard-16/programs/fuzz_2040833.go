package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 uint8 = 243
    if ((0 <= 42) || true) {
        for fi4 := 0; fi4 < 17; fi4++ {
            if (1e10 != (3.14159 - 1.5)) {
                var v5 uint64 = 700
                if false {
                    v6 := []int{}
                    switch "" {
                    case "":
                        var v7 int32 = -2147483648
                        _ = v7
                    case "a":
                        var v8 uint8 = 203
                        _ = v8
                    default:
                        delete(v1, "\n\t")
                    }
                    fmt.Println(v5)
                    _ = v6
                }
                fmt.Println(fi4, fi4, v3)
                _ = v5
            }
            v3 -= 231
            fmt.Println(fi4, fi4)
            _ = fi4
        }
        for fi9 := 0; fi9 < 20; fi9++ {
            if true {
                for fi10 := 0; fi10 < 5; fi10++ {
                    switch "a" {
                    case "a":
                        var v11 int16 = 42
                        _ = v11
                    case "\n\t":
                        v3 = (v3 | (^(v3)))
                    case "hello":
                        var v12 int8 = 0
                        _ = v12
                    default:
                        v3 = 172
                    }
                    if (0.0 != 1.5) {
                        fmt.Println(v3)
                        if (false && (2147483647 < 1)) {
                            v13 := GenS1{F0: 151, F1: 2, F2: 562}
                            _ = v13
                        } else {
                            var v14 int8 = -15
                            _ = v14
                        }
                        fmt.Println(v3, v3, fi9)
                    }
                    switch ("\n\t" + (("hello" + "") + "")) {
                    case "":
                        v2[""] = ("foo bar" + "\n\t")
                    case "a":
                        switch "a" {
                        case "\n\t":
                            fmt.Println(fi10, v3)
                        case "":
                            v3 = 247
                        }
                    }
                    v3 = (v3 % ((v3) | 1))
                    _ = fi10
                }
                v15 := []int{fi9, (fi9 / ((9223372036854775807) | 1))}
                if false {
                    v16 := GenS1{F0: 32767, F1: 1, F2: 217}
                    var v17 uint8 = (v3 % (((^(v3))) | 1))
                    _ = v16
                    _ = v17
                }
                _ = v15
            } else {
                v18 := []int{(fi9 | 2147483647), fi9}
                _ = v18
            }
            _ = fi9
        }
        var v19 bool = (true || (!((1e10 != 0.0) || (!(72 >= (v3 / ((42) | 1)))))))
        _ = v19
    } else {
        var v20 uint64 = 590
        _ = v20
    }
    v3 = 127
    v21 := GenS1{F0: 821, F1: 693, F2: 2147483648}
    v3 = 127
    for fi22 := 0; fi22 < 4; fi22++ {
        v3 = (^(v3))
        var v23 int16 = 0
        v23 -= (v23 / ((697) | 1))
        _ = fi22
        _ = v23
    }
    var v24 uint = 255
    v24 = v24
    var v25 int = 1
    if false {
        v3 -= 118
    }
    var v26 float64 = 0.1
    fmt.Println(v24)
    switch ("foo bar" + "\n\t") {
    case "hello":
        var v27 int = 681
        _ = v27
    }
    var v28 uint32 = 2147483648
    var v29 float64 = v26
    for fi30 := 0; fi30 < 1; fi30++ {
        if false {
            for fi31 := 0; fi31 < 14; fi31++ {
                fmt.Println(v29, v25)
                v32 := GenS1{F0: 2, F1: v24, F2: 713}
                var v33 uint64 = 18446744073709551615
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        v29 = v29
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
}
