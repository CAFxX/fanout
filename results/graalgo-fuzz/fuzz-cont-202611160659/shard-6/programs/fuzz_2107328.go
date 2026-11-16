package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]string)
    var v2 string = "a"
    v3 := []int{}
    var v4 float32 = 1e10
    var v5 int16 = 304
    if len(v3) > 0 {
        v3[0] = 42
    }
    fmt.Println(v4, v4)
    var v6 float64 = (((((3.14159 + -2.25) * 0.0) / 0.1) / 1e10) / 1.0)
    v2 = (v2 + v2)
    v6 = 0.1
    var v7 string = v2
    switch v7 {
    case "hello":
        switch v7 {
        case "\n\t":
            var v8 uint = 18446744073709551615
            _ = v8
        case "a":
            switch 128 {
            case 2147483647:
                var v9 int = -856
                _ = v9
            }
        default:
            if true {
                v10 := v1[2147483648]
                _ = v10
            }
        }
    default:
        switch v7 {
        case "hello":
            var v11 uint32 = 65535
            _ = v11
        case "\n\t":
            v3 = append(v3, 32767)
        case "":
            for fi12 := 0; fi12 < 3; fi12++ {
                v6 -= float64(float32(3.14159))
                _ = fi12
            }
        default:
            var v13 int64 = -528
            _ = v13
        }
    }
    fmt.Println(v4, v7, v5)
    var v14 int64 = 646
    v7 = (((("a" + (((v2 + v7) + "foo bar") + (v7 + ("hello" + (((v7 + v7) + ("a" + "a")) + v2))))) + v2) + ("hello" + (v2 + v2))) + ("\n\t" + "hello"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
