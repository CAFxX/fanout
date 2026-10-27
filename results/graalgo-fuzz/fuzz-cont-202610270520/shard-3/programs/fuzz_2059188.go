package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint) bool {
    v6 := []int{42}
    fmt.Println(p0)
    _ = len(v6)
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch 555 {
        case -321:
            var v7 float32 = ((0.0 + (-2.25 * (1.5 - (3.14159 * -2.25)))) * 1e10)
            _ = v7
        default:
            v8 := GenS1{F0: 0}
            _ = v8
        }
        switch ("a" + (("a" + "hello") + "")) {
        case "":
            if (32768 <= -2147483648) {
                fmt.Println(p0)
                v9 := GenS1{F0: 85}
                _ = v9
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v10 float64 = -2.25
                _ = v10
            }
        }
    } else {
        for fi11 := 0; fi11 < 12; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                v6 = append(v6, (fi12 % ((fi11) | 1)))
                v13 := GenS1{F0: 37}
                if false {
                    v6 = append(v6, fi12)
                    if (0 != 658) {
                        var v14 string = "a"
                        var v15 int8 = -54
                        _ = v14
                        _ = v15
                    }
                }
                p0 = (^(p0))
                _ = len(v6)
                p0 = p0
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    }
    _ = p0
    _ = v6
    return (!false)
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    v3 := make(map[uint32]int)
    var v4 int8 = 1
    var v5 int8 = v4
    v16 := GenS1{F0: 127}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v5)
    case "hello":
        if len(v2) > 0 {
            v2[0] = ""
        }
    case "\n\t":
        var v17 bool = true
        _ = v17
    default:
        v5 = (v4 + -1)
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        v19 := v1[(fi18 >> ((42) & 63))]
        v19 = v19
        _ = fi18
        _ = v19
    }
    var v20 uint32 = 2147483648
    v4 -= v5
    switch "hello" {
    case "hello":
        var v21 uint8 = 78
        _ = v21
    default:
        if (!(18446744073709551615 != 65535)) {
            v5 -= (v5 + (v5 << ((18446744073709551615) & 63)))
        } else {
            var v22 float32 = 0.1
            v23 := GenS1{F0: 255}
            _ = v22
            _ = v23
        }
    }
    v24 := GenS2{F0: (((!(!true)) && false) || true)}
    v5 = v4
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v5 = 100
    for fi25 := 0; fi25 < 8; fi25++ {
        if (4294967295 < (v20 / ((42) | 1))) {
            switch "" {
            case "":
                v5 -= (v5 * v5)
            case "\n\t":
                v26 := GenS2{F0: true}
                _ = v26
            case "hello":
                v27 := GenS1{F0: 127}
                _ = v27
            default:
                v28 := GenS2{F0: (!false)}
                _ = v28
            }
        } else {
            v29 := GenS2{F0: (false && false)}
            _ = v29
        }
        _ = fi25
    }
    switch ("hello" + "a") {
    case "a":
        v5 = v4
    case "foo bar":
        var v30 int8 = (v4 - 0)
        _ = v30
    case "":
        for fi31 := 0; fi31 < 7; fi31++ {
            var v32 uint8 = 2
            _ = fi31
            _ = v32
        }
    }
    fmt.Println(v4, v20, v5)
    var v33 int32 = -2147483648
    v34 := GenS2{F0: true}
    if false {
        var v35 bool = ((!((uint16(v20) < 421) && (-213 <= -9223372036854775808))) || true)
        _ = len(v2)
        var v36 int16 = -955
        _ = v35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
