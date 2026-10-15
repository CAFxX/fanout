package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 string = ""
    var v2 float32 = 3.14159
    v3 := make(map[int]int)
    var v4 int32 = 0
    var v5 float32 = 0.1
    var v6 int32 = (v4 & (-(v4)))
    var v7 int16 = 0
    var v8 uint = 18446744073709551615
    var v9 string = (("" + v1) + "")
    var v10 uint16 = 0
    fmt.Println(v4, v4)
    if false {
        v11 := GenS1{F0: (v8 % ((v8) | 1))}
        _ = v11
    }
    if (v8 > v8) {
        v12 := GenS1{F0: uint(v10)}
        _ = v12
    } else {
        var v13 string = (v9 + "hello")
        var v14 uint32 = 0
        switch v1 {
        case "foo bar":
            v15 := GenS1{F0: 0}
            _ = v15
        default:
            v5 += 1.5
        }
        _ = v13
        _ = v14
    }
    v16 := GenS1{F0: 65535}
    if (v7 > (v7 * -32768)) {
        var v17 uint8 = 42
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS1{F0: v8}
            _ = fi18
            _ = v19
        }
        v20 := GenS1{F0: 0}
        _ = v17
        _ = v20
    } else {
        v21 := []int{-1}
        _ = v21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        if false {
            delete(v3, int(v8))
            var v23 int = (fi22 + -225)
            for fi24 := 0; fi24 < 17; fi24++ {
                v6 = 441
                _ = fi24
            }
            _ = v23
        } else {
            if (v8 >= (v8 % ((v8) | 1))) {
                fmt.Println(v2)
                fmt.Println(v4, v6)
            }
            var v25 string = v1
            if true {
                v7 -= v7
                fmt.Println(v10)
                var v26 uint16 = (v10 + 359)
                _ = v26
            }
            _ = v25
        }
        v7 += v7
        _ = fi22
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        v8 = (v8 & uint(v10))
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
