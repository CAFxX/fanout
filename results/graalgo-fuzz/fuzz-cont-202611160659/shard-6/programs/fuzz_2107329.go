package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint8, p1 uint32) int64 {
    var v6 uint16 = 255
    if ((p1 % (((^(p1))) | 1)) != 155) {
        var v7 uint32 = p1
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint64 = 997
        fmt.Println(v9)
        v9 -= (v9 * v9)
        _ = fi8
        _ = v9
    }
    var v10 bool = false
    var v11 uint32 = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 bool) uint64 {
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := []int{fi12, 128, -2147483648, (fi12 ^ 436)}
        var v14 uint32 = 42
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int32 = 42
    v16 := GenS1{F0: 4294967295, F1: (v15 % (((v15 % ((1) | 1))) | 1))}
    var v17 uint32 = 2147483647
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return 4294967295
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 229
    v3 := []string{}
    var v4 int8 = 80
    var v5 uint = (v2 * (v2 << (((v2 >> ((v2) & 63))) & 63)))
    var v18 int8 = v4
    switch -663 {
    case -2147483648:
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v5)
            v20 := GenS1{F0: 2147483648, F1: -525}
            _ = fi19
            _ = v20
        }
    case -516:
        v5 = v5
    default:
        if ((1e10 != 1e10) && false) {
            var v21 uint8 = 127
            v5 = 73
            _ = v21
        }
    }
    switch -2147483648 {
    case -414:
        v2 += v2
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    v2 -= 817
    fmt.Println(v18)
    v18 = v4
    if true {
        if (593 == 1) {
            v22 := GenS1{F0: 1, F1: 2147483647}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v22
        }
        switch ("foo bar" + "\n\t") {
        case "a":
            _ = len(v3)
        case "\n\t":
            switch 0 {
            case 4294967295:
                var v23 uint64 = genFn1(false)
                _ = v23
            default:
                for fi24 := 0; fi24 < 1; fi24++ {
                    if true {
                        for fi25 := 0; fi25 < 13; fi25++ {
                            var v26 int8 = v18
                            _ = fi25
                            _ = v26
                        }
                        v2 = v2
                        v18 = (v18 % ((-23) | 1))
                    }
                    var v27 int16 = 1000
                    for fi28 := 0; fi28 < 1; fi28++ {
                        var v29 int = fi24
                        v30 := GenS1{F0: 42, F1: 0}
                        var v31 int16 = (v27 - 42)
                        _ = fi28
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    _ = fi24
                    _ = v27
                }
            }
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = ""
            }
        default:
            var v32 uint16 = 1
            _ = v32
        }
    }
    var v33 int = 127
    if false {
        if false {
            switch "" {
            case "\n\t":
                v34 := GenS1{F0: 42, F1: -2147483648}
                _ = v34
            case "a":
                v1 += genFn1(((v1 / ((v1) | 1)) >= v1))
            default:
                switch ("a" + "") {
                case "":
                    v3 = append(v3, "hello")
                case "foo bar":
                    var v35 string = "hello"
                    _ = v35
                case "hello":
                    var v36 bool = false
                    _ = v36
                }
            }
            v5 = v2
        }
        for fi37 := 0; fi37 < 7; fi37++ {
            _ = len(v3)
            _ = fi37
        }
        v38 := GenS1{F0: 1, F1: 42}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
