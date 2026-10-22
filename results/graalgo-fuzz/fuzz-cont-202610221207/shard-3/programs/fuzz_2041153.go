package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int8 {
    var v4 uint32 = 2
    var v5 int16 = 990
    var v6 uint64 = 4294967295
    if (true || (!(0 < 1))) {
        var v7 float32 = 0.0
        v8 := []int{42}
        _ = v7
        _ = v8
    }
    if false {
        switch ("foo bar" + "") {
        case "foo bar":
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 int8 = 1
                v6 = v6
                _ = fi9
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            fmt.Println(v5)
            _ = fi11
        }
    } else {
        v5 += (v5 * (v5 - (-(v5))))
    }
    _ = v4
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 uint64 = 221
    v2 := make(map[int]bool)
    var v3 int8 = -45
    fmt.Println(v3)
    v1 = 0
    v3 = 127
    var v12 uint = 65535
    v13 := GenS1{F0: -32768}
    v14 := []int{int(v3), 687}
    switch 109 {
    case 4294967295:
        fmt.Println(v1, v1)
    case 128:
        if ((!(!false)) && false) {
            fmt.Println(v3)
            v3 = 57
        }
    case 9223372036854775807:
        if true {
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            switch 948 {
            case 42:
                for fi15 := 0; fi15 < 12; fi15++ {
                    v3 -= v3
                    _ = fi15
                }
            case 9223372036854775807:
                var v16 uint8 = 1
                _ = v16
            default:
                v3 += -56
            }
            fmt.Println(v1, v3, v3)
        } else {
            var v17 bool = (1 < 42)
            _ = v17
        }
    default:
        fmt.Println(v1, v12, v12)
    }
    fmt.Println(v12, v3, v12)
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := v2[fi18]
        _ = fi18
        _ = v19
    }
    if true {
        v3 = v3
        var v20 float32 = 0.1
        v21 := GenS1{F0: 678}
        _ = v20
        _ = v21
    }
    fmt.Println(v12, v3, v1)
    v22 := GenS1{F0: 1000}
    fmt.Println(v1, v3)
    fmt.Println(v12, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
}
