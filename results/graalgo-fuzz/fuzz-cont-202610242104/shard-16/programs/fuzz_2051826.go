package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 bool, p1 int8) int {
    var v5 string = ((((("\n\t" + ("" + (("foo bar" + "foo bar") + ("foo bar" + ("foo bar" + "a"))))) + (("hello" + "hello") + "\n\t")) + "\n\t") + ("" + "\n\t")) + ("foo bar" + (("hello" + "foo bar") + (("a" + "hello") + "hello"))))
    fmt.Println(p1, v5)
    fmt.Println(p0, v5)
    _ = p0
    _ = p1
    _ = v5
    return 32768
}
func genFn1(p0 int64, p1 uint16) string {
    var v6 uint = 376
    p1 = (p1 << ((2) & 63))
    p1 = p1
    p0 = (p0 + p0)
    v6 = 290
    _ = p0
    _ = p1
    _ = v6
    return "foo bar"
}

func main() {
    var v1 bool = (2147483648 >= 42)
    v2 := []uint32{42}
    var v3 uint16 = 2
    var v4 bool = ((((38 >= 255) || (((1e10 != 1e10) || v1) || v1)) && v1) || false)
    switch 127 {
    case -531:
        v7 := GenS1{F0: 0, F1: 1}
        _ = v7
    case -1:
        v4 = (!true)
    }
    v1 = (!(42 >= 0))
    var v8 uint = 255
    var v9 uint32 = 2
    v8 = v8
    var v10 uint8 = 255
    if false {
        for fi11 := 0; fi11 < 15; fi11++ {
            switch fi11 {
            case -456:
                switch "foo bar" {
                case "foo bar":
                    switch ("hello" + "") {
                    case "\n\t":
                        fmt.Println(v4, v3, v1)
                    case "hello":
                        fmt.Println(v3)
                    }
                case "\n\t":
                    v9 += (v9 * 841)
                case "":
                    v12 := GenS1{F0: (v10 + 255), F1: 42}
                    _ = v12
                default:
                    if false {
                        var v13 float64 = (0.0 / -2.25)
                        v13 = float64(1e10)
                        for fi14 := 0; fi14 < 1; fi14++ {
                            var v15 bool = v4
                            _ = fi14
                            _ = v15
                        }
                        _ = v13
                    } else {
                        fmt.Println(v10, v10)
                        v16 := GenS1{F0: v10, F1: 0}
                        _ = v16
                    }
                }
            case 32768:
                var v17 int16 = 388
                _ = v17
            case -2147483648:
                fmt.Println(v8, v1)
            }
            v18 := GenS2{F0: 438}
            _ = fi11
            _ = v18
        }
        v3 = v3
    } else {
        for fi19 := 0; fi19 < 12; fi19++ {
            switch ("a" + "\n\t") {
            case "a":
                switch ("\n\t" + (("a" + "") + "foo bar")) {
                case "\n\t":
                    v10 = v10
                case "a":
                    if v4 {
                        var v20 uint8 = (^(v10))
                        _ = v20
                    } else {
                        var v21 uint8 = (v10 & v10)
                        fmt.Println(v21)
                        if len(v2) > 0 {
                            v2[0] = v9
                        }
                        _ = v21
                    }
                case "":
                    if ((288 != -2147483648) || v4) {
                        var v22 int8 = 0
                        _ = v22
                    }
                default:
                    for fi23 := 0; fi23 < 12; fi23++ {
                        fmt.Println(v8, fi19)
                        _ = fi23
                    }
                }
            case "":
                _ = len(v2)
            default:
                switch int(v9) {
                case -48:
                    _ = len(v2)
                case 127:
                    v24 := GenS1{F0: v10, F1: -1}
                    _ = v24
                }
            }
            _ = fi19
        }
        fmt.Println(v10)
        v25 := GenS2{F0: (v8 * uint(v10))}
        _ = v25
    }
    v1 = false
    var v26 uint = 251
    v9 += 0
    if false {
        v27 := GenS2{F0: (v8 ^ (v8 ^ (v8 << ((18446744073709551615) & 63))))}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        var v29 int64 = -1
        _ = fi28
        _ = v29
    }
    var v30 int = 32768
    v31 := GenS1{F0: v10, F1: 9223372036854775807}
    v26 += (v26 | (v8 / ((v26) | 1)))
    var v32 int32 = -378
    for fi33 := 0; fi33 < 17; fi33++ {
        v2 = append(v2, (^(v9)))
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v8
    _ = v9
}
