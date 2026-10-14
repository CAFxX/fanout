package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 string = "foo bar"
    var v2 float32 = ((0.1 - (3.14159 / 1.0)) - 3.14159)
    var v3 uint16 = 644
    v4 := make(map[uint32]int)
    var v5 bool = ((false || (!false)) && false)
    fmt.Println(v1)
    v4[398] = -2147483648
    v6 := GenS1{F0: 65535}
    if true {
        var v7 uint8 = 189
        var v8 uint8 = (v7 << ((183) & 63))
        v7 += 65
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 0}
    if v5 {
        var v10 int16 = -32768
        v10 = v10
        if (!v5) {
            var v11 string = ((v1 + ("foo bar" + v1)) + v1)
            switch 2147483647 {
            case -634:
                switch 42 {
                case 0:
                    v4[4294967295] = 738
                }
            }
            v12 := []int{-9223372036854775808, 9223372036854775807, 32767}
            _ = v11
            _ = v12
        } else {
            if false {
                fmt.Println(v10)
                var v13 uint16 = v3
                if (!v5) {
                    var v14 uint32 = 2147483647
                    fmt.Println(v1)
                    for fi15 := 0; fi15 < 2; fi15++ {
                        v4[1] = (fi15 | fi15)
                        v16 := GenS1{F0: 820}
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                }
                _ = v13
            } else {
                v1 = "foo bar"
                _ = len(v4)
                delete(v4, 4294967295)
            }
            delete(v4, 42)
            var v17 string = v1
            _ = v17
        }
        _ = v10
    } else {
        _ = len(v4)
        v18 := []int{129}
        _ = v18
    }
    v5 = (166 >= 4294967295)
    v2 += ((v2 * ((v2 - (((3.14159 / (v2 - v2)) / v2) / float32((((1.5 - (float64(v2) * (0.1 * 1.5))) * 1.5) / 1e10)))) / float32(1.5))) - v2)
    if true {
        if (v5 || true) {
            var v19 int16 = 32767
            _ = v19
        }
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 string = ""
            v5 = v5
            fmt.Println(v2, fi20, v2)
            _ = fi20
            _ = v21
        }
        var v22 uint = 42
        _ = v22
    }
    var v23 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
