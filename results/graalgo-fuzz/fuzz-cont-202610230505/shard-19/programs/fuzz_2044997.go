package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 107
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    var v4 uint = 4294967295
    var v5 int = 128
    v6 := []int{745, v5}
    if false {
        switch v5 {
        case 665:
            v5 -= 9223372036854775807
        case 524:
            for fi7 := 0; fi7 < 17; fi7++ {
                _ = len(v3)
                _ = fi7
            }
        default:
            var v8 float32 = ((1.5 + 1e10) - 0.1)
            _ = v8
        }
    }
    if false {
        v1 += v1
    } else {
        var v9 string = (("foo bar" + "hello") + ("\n\t" + ((("foo bar" + "a") + "\n\t") + "foo bar")))
        var v10 int32 = -182
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := v3[297]
            var v13 float64 = float64(0.1)
            if false {
                v9 = "a"
                v14 := v3[0]
                _ = v14
            } else {
                v10 -= 2147483647
                v15 := v3[256]
                _ = v15
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
        _ = v9
    }
    var v16 uint16 = 857
    var v17 uint = 470
    var v18 int = 32768
    for fi19 := 0; fi19 < 6; fi19++ {
        v4 -= v4
        _ = fi19
    }
    fmt.Println(v16, v17, v4)
    for fi20 := 0; fi20 < 20; fi20++ {
        if (!(1 == 42)) {
            if (v4 == v4) {
                v18 -= v18
            }
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            v4 = v4
            _ = fi21
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
