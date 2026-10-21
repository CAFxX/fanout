package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
func genFn0(p0 uint8) int {
    var v6 uint16 = 32767
    v6 = (v6 >> ((42) & 63))
    for fi7 := 0; fi7 < 8; fi7++ {
        p0 -= 2
        for fi8 := 0; fi8 < 1; fi8++ {
            v6 += v6
            _ = fi8
        }
        _ = fi7
    }
    v6 = 793
    _ = p0
    _ = v6
    return 1
}

func main() {
    v1 := []string{"hello", "\n\t", ((("foo bar" + (("hello" + ((("a" + "") + ("" + "")) + "")) + ("hello" + "\n\t"))) + ("foo bar" + "\n\t")) + "a")}
    v2 := []int{-834, 9223372036854775807, 128}
    var v3 int32 = 42
    var v4 int = 131
    v5 := []uint32{}
    fmt.Println(v4, v3)
    v4 = genFn0(235)
    var v9 int8 = 127
    var v10 bool = true
    if ((^(v9)) <= (v9 >> ((2) & 63))) {
        for fi11 := 0; fi11 < 7; fi11++ {
            switch ("hello" + "hello") {
            case "hello":
                if (!v10) {
                    v12 := GenS1{F0: (!(true && (v9 > v9))), F1: 14}
                    v4 = v4
                    _ = v12
                }
            case "":
                fmt.Println(v9)
            default:
                if (!true) {
                    var v13 int16 = 610
                    _ = v13
                }
            }
            _ = fi11
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v4, v4)
    if len(v2) > 0 {
        v2[0] = v4
    }
    fmt.Println(v4)
    switch "\n\t" {
    case "foo bar":
        switch 9223372036854775807 {
        case -9223372036854775808:
            fmt.Println(v9)
        case 116:
            v3 = -2147483648
        }
    }
    v10 = (v4 >= (v4 * 9223372036854775807))
    fmt.Println(v4)
    if v10 {
        v10 = v10
        if v10 {
            if (255 < 447) {
                fmt.Println(v9, v3, v4)
            }
            var v14 float32 = 0.0
            _ = v14
        } else {
            var v15 uint8 = 255
            var v16 int = genFn0(v15)
            if (float32(float64(0.0)) <= (-2.25 - (0.0 + 0.0))) {
                fmt.Println(v10, v15)
                if (!(true || ((v4 / (((v16 ^ genFn0(242))) | 1)) < v16))) {
                    v15 += v15
                }
                if len(v2) > 0 {
                    v2[0] = genFn0(v15)
                }
            }
            _ = v15
            _ = v16
        }
    }
    var v17 uint = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
