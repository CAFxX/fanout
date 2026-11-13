package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
func genFn0(p0 int32, p1 string, p2 uint16) uint {
    p2 = p2
    fmt.Println(p2)
    switch -559 {
    case -1:
        for fi6 := 0; fi6 < 1; fi6++ {
            p1 = (p1 + ("a" + p1))
            _ = fi6
        }
    case 32768:
        switch (p1 + p1) {
        case "\n\t":
            var v7 float32 = 3.14159
            _ = v7
        case "hello":
            v8 := []int{-982, 128, 32767, -2147483648}
            _ = v8
        default:
            v9 := []int{-614}
            _ = v9
        }
    }
    var v10 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 848
}

func main() {
    var v1 string = ""
    v2 := make(map[int]string)
    var v3 int64 = 42
    var v4 float64 = 0.1
    var v5 float32 = 3.14159
    v11 := GenS1{F0: 944, F1: (false || true)}
    fmt.Println(v3, v3)
    var v12 uint = 969
    fmt.Println(v1, v3, v3)
    fmt.Println(v1, v3)
    var v13 string = v1
    switch 4294967295 {
    case 32767:
        if (!(false && (-93 < 42))) {
            fmt.Println(v4, v13, v3)
        }
    case 127:
        v1 = "hello"
    case -2147483648:
        var v14 uint8 = 42
        _ = v14
    }
    switch 32768 {
    case 0:
        var v15 int32 = int32(v12)
        _ = v15
    case 32768:
        v12 = 897
    case 4294967295:
        if (0 > 1) {
            if false {
                switch -809 {
                case -375:
                    if false {
                        if false {
                            fmt.Println(v13)
                        }
                        v16 := GenS1{F0: 2147483648, F1: (!(!false))}
                        _ = v16
                    }
                }
                var v17 string = "\n\t"
                _ = v17
            }
        } else {
            switch v1 {
            case "foo bar":
                for fi18 := 0; fi18 < 3; fi18++ {
                    _ = len(v2)
                    _ = fi18
                }
            }
        }
    default:
        v12 -= v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
