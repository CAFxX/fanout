package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint8 = 255
    var v2 int16 = 855
    var v3 uint64 = uint64(v1)
    v4 := []int{1, 43, 80}
    var v5 int64 = -455
    var v6 int8 = 127
    var v7 float64 = 0.1
    fmt.Println(v1, v6, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch -797 {
    case 9223372036854775807:
        v6 = v6
    case 187:
        var v8 uint8 = 141
        _ = v8
    default:
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 bool = (true || (false && (-2147483648 >= 0)))
            var v11 uint = 692
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    switch -1 {
    case -1:
        if (!true) {
            var v12 uint32 = 4294967295
            if false {
                _ = len(v4)
            } else {
                v3 = (v3 & 2)
                v3 += 2
            }
            _ = v12
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v1, v6)
            v1 -= v1
        }
    default:
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            if true {
                var v13 uint64 = v3
                fmt.Println(v13, v1, v1)
                _ = v13
            } else {
                _ = len(v4)
                if true {
                    v2 = v2
                    for fi14 := 0; fi14 < 2; fi14++ {
                        v7 -= ((0.0 / v7) + v7)
                        _ = fi14
                    }
                    _ = len(v4)
                } else {
                    for fi15 := 0; fi15 < 16; fi15++ {
                        v4 = append(v4, (fi15 * (fi15 - 32767)))
                        v2 -= 42
                        var v16 int16 = (v2 * 32767)
                        _ = fi15
                        _ = v16
                    }
                    fmt.Println(v5, v1, v2)
                    v7 = (((float64(-2.25) * v7) * (v7 + v7)) / v7)
                }
            }
        default:
            var v17 uint32 = 0
            _ = v17
        }
    }
    var v18 float64 = float64((1.5 - 1.5))
    if (99 <= 0) {
        if (-2147483648 <= 273) {
            v6 = 42
            v4 = append(v4, 4294967295)
        }
    } else {
        if (false || false) {
            for fi19 := 0; fi19 < 14; fi19++ {
                for fi20 := 0; fi20 < 8; fi20++ {
                    v3 -= (v3 ^ 2)
                    v21 := GenS1{F0: true}
                    v2 -= 940
                    v4 = append(v4, 127)
                    _ = fi20
                    _ = v21
                }
                v18 += 1.5
                var v22 int = fi19
                _ = fi19
                _ = v22
            }
            var v23 uint = 929
            switch ("\n\t" + "a") {
            case "":
                v24 := GenS1{F0: false}
                _ = v24
            default:
                var v25 uint64 = v3
                _ = v25
            }
            _ = v23
        } else {
            var v26 uint64 = v3
            for fi27 := 0; fi27 < 9; fi27++ {
                if false {
                    v1 = (v1 & (v1 >> ((254) & 63)))
                    _ = len(v4)
                } else {
                    v5 = (v5 * v5)
                    _ = len(v4)
                }
                _ = fi27
            }
            _ = v26
        }
        if (true && true) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            var v28 float64 = v7
            v5 -= v5
            var v29 uint = 42
            _ = v28
            _ = v29
        }
    }
    if false {
        if (-1 == -244) {
            if false {
                var v30 uint8 = (^(v1))
                v7 = (v7 - (v18 * -2.25))
                v2 += v2
                _ = v30
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v5, v5, v3)
                v31 := GenS1{F0: false}
                _ = v31
            }
        }
    }
    fmt.Println(v1)
    v1 += v1
    var v32 int64 = -1
    var v33 uint16 = 65535
    var v34 float32 = float32(v18)
    var v35 uint32 = uint32(v1)
    for fi36 := 0; fi36 < 5; fi36++ {
        switch "a" {
        case "hello":
            _ = len(v4)
        }
        if len(v4) > 0 {
            v4[0] = fi36
        }
        if false {
            switch 9223372036854775807 {
            case 1:
                fmt.Println(v1, v2)
            case -2147483648:
                v33 = v33
            }
        }
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
