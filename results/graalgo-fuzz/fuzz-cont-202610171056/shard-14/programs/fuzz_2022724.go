package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64, p1 uint, p2 int64) int8 {
    switch "hello" {
    case "hello":
        if (!false) {
            var v7 float32 = ((0.0 * 0.1) * -2.25)
            if (42 < 32767) {
                switch "\n\t" {
                case "foo bar":
                    v8 := GenS1{F0: 402}
                    _ = v8
                default:
                    var v9 int = 32768
                    _ = v9
                }
                fmt.Println(v7, p0)
            }
            _ = v7
        } else {
            p2 = p0
        }
    case "a":
        var v10 float32 = ((-2.25 + 3.14159) * (float32((1.5 / float64((((3.14159 / (-2.25 * 1e10)) - (3.14159 * 3.14159)) + 0.0)))) - float32(((-2.25 / 1e10) + -2.25))))
        _ = v10
    default:
        v11 := []int{}
        _ = v11
    }
    var v12 int32 = -835
    fmt.Println(v12, p0)
    v12 -= v12
    fmt.Println(v12)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 127
}

func main() {
    var v1 int16 = -1
    var v2 uint = 18446744073709551615
    var v3 int8 = -124
    var v4 bool = (false && false)
    var v5 uint32 = 766
    v6 := []int{1, 32768}
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            v3 = v3
            switch "" {
            case "\n\t":
                fmt.Println(v5)
            case "hello":
                var v15 int64 = -1
                _ = v15
            }
            v3 += (v3 | -1)
            _ = fi14
        }
        var v16 int64 = -9223372036854775808
        v3 -= 46
        _ = fi13
        _ = v16
    }
    if ((719 == 18446744073709551615) && v4) {
        var v17 float64 = -2.25
        _ = v17
    } else {
        switch -2147483648 {
        case 9223372036854775807:
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 uint32 = (^(v5))
                _ = fi18
                _ = v19
            }
        }
        v20 := GenS1{F0: 1}
        if v4 {
            if v4 {
                var v21 float64 = -2.25
                switch ("foo bar" + (("a" + (("" + "") + "\n\t")) + "")) {
                case "a":
                    v22 := GenS1{F0: 0}
                    _ = v22
                }
                _ = v21
            } else {
                fmt.Println(v4, v1, v4)
                v4 = (v4 && (-644 >= 393))
                v23 := GenS1{F0: 2}
                _ = v23
            }
            v24 := GenS2{F0: 42}
            _ = v24
        } else {
            v4 = v4
            v5 += (v5 ^ 65535)
        }
        _ = v20
    }
    v4 = (v4 && v4)
    fmt.Println(v2, v3, v5)
    v2 = v2
    var v25 int16 = 1
    for fi26 := 0; fi26 < 10; fi26++ {
        v2 = 4294967295
        var v27 string = (("hello" + (("a" + ("" + "a")) + "foo bar")) + "a")
        v28 := GenS3{F0: (v1 - v25)}
        _ = fi26
        _ = v27
        _ = v28
    }
    v3 -= (^(v3))
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 string = ("a" + "")
        _ = fi29
        _ = v30
    }
    fmt.Println(v5, v1)
    var v31 uint16 = 916
    var v32 string = ""
    v1 = -1
    v6 = append(v6, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
