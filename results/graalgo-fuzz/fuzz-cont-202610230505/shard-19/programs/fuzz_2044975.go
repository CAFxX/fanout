package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint16) string {
    switch ("hello" + (("a" + "") + "")) {
    case "hello":
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 int32 = 88
            switch fi6 {
            case -266:
                var v8 int64 = 1
                _ = v8
            }
            _ = fi6
            _ = v7
        }
    case "foo bar":
        var v9 int32 = -2147483648
        _ = v9
    default:
        if true {
            if false {
                p0 += 2
            }
            v10 := GenS1{F0: 2, F1: -2147483648}
            _ = v10
        }
    }
    v11 := []int{32768}
    fmt.Println(p0)
    _ = p0
    _ = v11
    return (("a" + (("hello" + "\n\t") + "a")) + "")
}

func main() {
    v1 := []int{-9223372036854775808}
    v2 := make(map[int]bool)
    var v3 bool = (-1 > -9223372036854775808)
    v4 := []string{(("" + ("a" + "")) + "a"), ("a" + ((("" + "foo bar") + "foo bar") + "a")), "", "a"}
    v5 := make(map[string]uint32)
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(fi12, v3)
            for fi13 := 0; fi13 < 20; fi13++ {
                fmt.Println(fi13, fi13, fi13)
                v3 = v3
                _ = len(v5)
                _ = fi13
            }
            var v14 int64 = -9223372036854775808
            _ = fi12
            _ = v14
        }
    case "hello":
        var v15 float32 = (float32(((float64(1e10) / 1.0) + -2.25)) + float32(((1.5 / 0.1) + (((3.14159 - float64((1e10 + (float32(-2.25) - 0.1)))) * (1e10 + ((float64((0.1 - 3.14159)) * 1e10) - 1e10))) - float64(3.14159)))))
        _ = v15
    }
    _ = len(v2)
    v16 := GenS1{F0: 2147483647, F1: 20}
    v17 := GenS1{F0: 805, F1: 42}
    switch 624 {
    case 9223372036854775807:
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 int8 = -12
            _ = fi18
            _ = v19
        }
    }
    var v20 int8 = 100
    v20 -= v20
    fmt.Println(v20, v3)
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
