package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() int8 {
    var v6 uint64 = 18446744073709551615
    if ((65535 >= 2147483647) && false) {
        switch 644 {
        case 1:
            var v7 bool = false
            _ = v7
        }
        if (!(775 == 960)) {
            switch "hello" {
            case "a":
                var v8 string = "foo bar"
                _ = v8
            case "hello":
                if true {
                    switch 4294967295 {
                    case 128:
                        fmt.Println(v6)
                    }
                    var v9 uint64 = (v6 >> ((2) & 63))
                    _ = v9
                } else {
                    if true {
                        v10 := []int{9223372036854775807}
                        _ = v10
                    } else {
                        v6 += v6
                    }
                    v6 = v6
                    switch "a" {
                    case "hello":
                        v6 -= (v6 - (v6 ^ v6))
                    }
                }
            default:
                v6 += v6
            }
        } else {
            v11 := GenS1{F0: 2}
            fmt.Println(v6)
            fmt.Println(v6)
            _ = v11
        }
    } else {
        v6 = (v6 << ((1) & 63))
        fmt.Println(v6)
        v12 := []int{127, -384, -229, 113}
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint32 = 0
        _ = fi13
        _ = v14
    }
    _ = v6
    return 0
}

func main() {
    var v1 int = 284
    v2 := []int{}
    v3 := make(map[string]int)
    var v4 uint64 = 2
    var v5 float32 = (-2.25 + -2.25)
    switch "hello" {
    case "\n\t":
        if true {
            switch ("foo bar" + "\n\t") {
            case "":
                v1 = v1
            case "a":
                var v15 uint8 = 127
                _ = v15
            }
        }
    case "hello":
        fmt.Println(v4, v5)
    }
    if ((v5 >= (v5 * ((float32(((float64(v5) * 1.5) / float64(v5))) / -2.25) - 1e10))) || (!(!false))) {
        v5 = float32(3.14159)
    } else {
        v16 := GenS1{F0: 42}
        fmt.Println(v4, v5)
        _ = v16
    }
    v1 = v1
    v1 = (v1 + (v1 | (v1 ^ -397)))
    fmt.Println(v1)
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 int64 = 42
            var v20 bool = true
            _ = fi18
            _ = v19
            _ = v20
        }
        v21 := GenS1{F0: 65535}
        _ = fi17
        _ = v21
    }
    fmt.Println(v1, v1, v5)
    fmt.Println(v1, v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
