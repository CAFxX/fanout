package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int }
func genFn0(p0 uint16) int {
    fmt.Println(p0)
    p0 += (p0 >> ((577) & 63))
    if (1.5 <= (0.0 * 1.5)) {
        v5 := GenS1{F0: 2, F1: 6, F2: 509}
        v6 := []int{2147483647, -9223372036854775808}
        if (1 >= -1) {
            p0 += p0
        }
        _ = v5
        _ = v6
    }
    var v7 float64 = 3.14159
    _ = p0
    _ = v7
    return 127
}
func genFn1(p0 uint8, p1 uint64, p2 uint8) int64 {
    var v8 int32 = -294
    for fi9 := 0; fi9 < 7; fi9++ {
        if true {
            p2 += 0
            switch ("" + ((((("foo bar" + "foo bar") + "\n\t") + "") + "hello") + "")) {
            case "a":
                fmt.Println(v8)
            case "\n\t":
                p1 -= (^(p1))
            case "":
                fmt.Println(v8, v8, v8)
            default:
                var v10 uint8 = 42
                _ = v10
            }
        }
        p2 -= 0
        var v11 int = fi9
        _ = fi9
        _ = v11
    }
    v8 += -2147483648
    var v12 int = -2147483648
    p0 = (p0 * 150)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v8
    return -985
}

func main() {
    var v1 string = "hello"
    v2 := []int{2147483647, 1, 957, 824}
    var v3 int8 = 1
    var v4 string = v1
    v13 := GenS1{F0: 0, F1: 127, F2: -9223372036854775808}
    v3 += (-(v3))
    if (true || (!false)) {
        v4 = (("a" + v4) + v4)
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 bool = false
                if len(v2) > 0 {
                    v2[0] = (fi14 % ((4294967295) | 1))
                }
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    var v17 int8 = v3
    var v18 int32 = -2147483648
    for fi19 := 0; fi19 < 1; fi19++ {
        if true {
            var v20 uint32 = 4294967295
            v18 += (v18 / (((v18 + -965)) | 1))
            _ = v20
        }
        _ = fi19
    }
    v17 = (v3 >> ((2) & 63))
    fmt.Println(v17, v1, v18)
    v4 = ("a" + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
