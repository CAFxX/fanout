package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 2147483647
    var v3 int8 = 1
    v4 := make(map[int]string)
    var v5 int32 = 2147483647
    fmt.Println(v3, v5)
    var v6 string = ("\n\t" + (("" + "foo bar") + ((("foo bar" + ("\n\t" + "hello")) + "foo bar") + ("hello" + "hello"))))
    for fi7 := 0; fi7 < 15; fi7++ {
        if false {
            for fi8 := 0; fi8 < 9; fi8++ {
                switch "a" {
                case "":
                    var v9 float32 = float32((float64(0.1) / (1e10 * ((3.14159 + (1.5 + (float64(-2.25) * ((((((3.14159 + 1.5) - (1e10 * 0.0)) + -2.25) - (-2.25 * 1e10)) + ((-2.25 / 1.5) - (-2.25 + (1.5 + 3.14159)))) * (1.5 + ((1e10 - (float64(1.5) + (((3.14159 / 1e10) + 1e10) * 0.0))) - (0.1 / -2.25))))))) * -2.25))))
                    _ = v9
                default:
                    var v10 float64 = (3.14159 * 0.1)
                    _ = v10
                }
                if true {
                    fmt.Println(v2, fi8)
                    v5 = v5
                    var v11 float32 = (float32(0.0) / float32(3.14159))
                    _ = v11
                } else {
                    var v12 int = fi7
                    v5 = v5
                    _ = v12
                }
                _ = fi8
            }
            var v13 int8 = (v3 | -72)
            _ = v13
        }
        _ = fi7
    }
    v5 = (-(v5))
    var v14 uint16 = 339
    for fi15 := 0; fi15 < 20; fi15++ {
        v3 -= int8(v5)
        var v16 float64 = float64(1e10)
        var v17 uint64 = 18446744073709551615
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := []int{32767}
    for fi19 := 0; fi19 < 2; fi19++ {
        v2 -= v2
        _ = fi19
    }
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
