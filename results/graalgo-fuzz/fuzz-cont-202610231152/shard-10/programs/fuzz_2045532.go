package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]string)
    var v2 string = ""
    var v3 int = 128
    v4 := []int{327, (v3 >> ((4294967295) & 63))}
    var v5 float64 = 1.5
    var v6 uint = 1
    if (!((((((v5 / (v5 / (v5 * 3.14159))) + (0.0 / v5)) / v5) * 0.1) * (v5 / v5)) <= (float64(0.0) * (v5 / v5)))) {
        var v7 string = v2
        switch v2 {
        case "foo bar":
            var v8 bool = true
            _ = v8
        case "":
            if (2 >= 42) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        case "\n\t":
            switch "foo bar" {
            case "a":
                delete(v1, 722)
            }
        }
        _ = v7
    }
    var v9 int8 = 127
    v2 = ("a" + v2)
    v3 -= (v3 * v3)
    v6 = (v6 & (v6 - (^(v6))))
    if true {
        var v10 int8 = (v9 | v9)
        _ = v10
    }
    v3 = v3
    if true {
        switch 338 {
        case 0:
            for fi11 := 0; fi11 < 13; fi11++ {
                switch v3 {
                case 32768:
                    v2 = v2
                case -1:
                    fmt.Println(v9, fi11, v9)
                }
                _ = fi11
            }
        case -9223372036854775808:
            var v12 uint16 = 1
            _ = v12
        case 127:
            v13 := v1[278]
            _ = v13
        default:
            var v14 int = -2147483648
            _ = v14
        }
        v1[2147483648] = v2
    } else {
        if (v2 < "a") {
            var v15 int64 = 42
            _ = v15
        }
        v5 = v5
    }
    if false {
        v5 = v5
    }
    switch ("foo bar" + "a") {
    case "hello":
        var v16 string = v2
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
