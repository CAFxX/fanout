package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int64 }

func main() {
    var v1 float64 = (0.1 - (3.14159 - float64(float32(3.14159))))
    var v2 uint8 = 1
    var v3 int16 = 0
    var v4 bool = true
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 0, F1: 42, F2: -1}
        v7 := []int{(fi5 ^ (fi5 / ((4294967295) | 1))), 42, fi5}
        _ = fi5
        _ = v6
        _ = v7
    }
    switch "" {
    case "":
        v8 := []int{2147483647, -1, 128, 2147483647}
        _ = v8
    default:
        v4 = (true || v4)
    }
    v1 += v1
    switch "hello" {
    case "":
        if (v1 != (1e10 - v1)) {
            if (127 < v2) {
                switch -756 {
                case 807:
                    for fi9 := 0; fi9 < 6; fi9++ {
                        v3 -= 465
                        if (uint64(v2) < 163) {
                            v10 := []int{fi9, 128}
                            _ = v10
                        }
                        _ = fi9
                    }
                }
                v4 = true
                for fi11 := 0; fi11 < 1; fi11++ {
                    var v12 string = ("a" + "a")
                    if ((989 >= 1) && true) {
                        v13 := []int{fi11, 42}
                        _ = len(v13)
                        _ = v13
                    }
                    _ = fi11
                    _ = v12
                }
            } else {
                v14 := []int{-9223372036854775808, -199}
                v15 := GenS1{F0: -74, F1: 876, F2: -85}
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 bool = v4
                v18 := GenS1{F0: -1, F1: 1, F2: -191}
                _ = fi16
                _ = v17
                _ = v18
            }
            switch "" {
            case "":
                v19 := GenS1{F0: 42, F1: 381, F2: -1}
                _ = v19
            default:
                switch "a" {
                case "a":
                    if (!v4) {
                        for fi20 := 0; fi20 < 19; fi20++ {
                            if (!false) {
                                v3 = (v3 & 42)
                                switch "\n\t" {
                                case "foo bar":
                                    v1 += ((v1 - (v1 + (v1 - (v1 + float64(-2.25))))) / 0.1)
                                }
                                if v4 {
                                    if v4 {
                                        for fi21 := 0; fi21 < 4; fi21++ {
                                            v22 := []int{42, (fi21 ^ (^(fi21))), -9223372036854775808}
                                            var v23 bool = true
                                            _ = fi21
                                            _ = v22
                                            _ = v23
                                        }
                                        v24 := GenS1{F0: -128, F1: 42, F2: -1}
                                        v25 := []int{}
                                        _ = v24
                                        _ = v25
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        if v4 {
                                            fmt.Println(v1, v4)
                                            switch "hello" {
                                            case "":
                                                v26 := []int{(-(fi20))}
                                                _ = v26
                                            case "foo bar":
                                                v4 = v4
                                            }
                                        }
                                    case "":
                                        v27 := []int{(fi20 * 312)}
                                        _ = v27
                                    case "a":
                                        fmt.Println(v4, v3, fi20)
                                    }
                                }
                            }
                            v1 += float64(0.1)
                            _ = fi20
                        }
                        var v28 uint = 42
                        v28 += 18446744073709551615
                        _ = v28
                    }
                case "\n\t":
                    if false {
                        fmt.Println(v1, v1, v1)
                    }
                default:
                    v3 -= 32767
                }
            }
        }
    default:
        var v29 uint16 = 32767
        _ = v29
    }
    v1 = ((v1 - (v1 - v1)) * (1.5 * (0.1 / 0.1)))
    switch 127 {
    case 128:
        var v30 bool = (false || ("foo bar" == ("" + ((("" + "\n\t") + ((("foo bar" + "foo bar") + ("hello" + "\n\t")) + "\n\t")) + "hello"))))
        _ = v30
    case 32768:
        v31 := GenS1{F0: 111, F1: 18446744073709551615, F2: -9223372036854775808}
        _ = v31
    case 4294967295:
        fmt.Println(v4, v3)
    }
    v2 = (v2 * v2)
    fmt.Println(v3)
    for fi32 := 0; fi32 < 1; fi32++ {
        v3 = -649
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
