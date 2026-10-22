package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v4 float32 = float32((0.0 - ((3.14159 + (0.1 * -2.25)) + float64(v1))))
    v5 := make(map[string]int)
    var v6 bool = (723 < 105)
    if true {
        v1 = v4
        switch ("a" + "\n\t") {
        case "hello":
            v6 = v6
        case "":
            var v7 bool = true
            _ = v7
        default:
            var v8 uint8 = 127
            _ = v8
        }
        v1 += float32((1.5 / 0.1))
    }
    v4 -= v4
    v6 = v6
    v9 := []int{133}
    var v10 int = 918
    switch "a" {
    case "hello":
        var v11 uint32 = 0
        _ = v11
    }
    if len(v9) > 0 {
        v9[0] = v10
    }
    if v6 {
        if true {
            if len(v9) > 0 {
                v9[0] = -388
            }
        } else {
            var v12 int64 = 42
            _ = v12
        }
        v6 = (!v6)
    }
    fmt.Println(v10)
    var v13 float64 = (0.0 - float64(((v1 + float32((0.1 - (3.14159 - (1.5 + (0.1 * 1.5)))))) + (v1 * (float32((-2.25 * 0.0)) * v4)))))
    v13 = -2.25
    v6 = v6
    switch 569 {
    case 76:
        v1 += (0.0 - 0.1)
    case -847:
        v14 := v5[(("a" + "hello") + "foo bar")]
        _ = v14
    }
    v4 += (v4 / v4)
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS1{F0: 916}
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 42}
    v18 := v5["hello"]
    if false {
        var v19 uint8 = 127
        switch ("foo bar" + "\n\t") {
        case "a":
            fmt.Println(v4, v10)
        case "hello":
            v13 += 3.14159
        case "":
            for fi20 := 0; fi20 < 15; fi20++ {
                v6 = v6
                var v21 uint = 255
                _ = fi20
                _ = v21
            }
        default:
            if (("" + "foo bar") <= ("foo bar" + ("hello" + "hello"))) {
                _ = len(v9)
            }
        }
        _ = v19
    } else {
        var v22 uint32 = 0
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
