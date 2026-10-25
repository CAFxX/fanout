package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]string)
    var v2 int8 = 1
    v3 := make(map[uint32]int)
    var v4 int64 = 213
    v5 := make(map[string]uint32)
    for fi6 := 0; fi6 < 17; fi6++ {
        fmt.Println(fi6, v4)
        _ = fi6
    }
    if (float64(-2.25) > 1.5) {
        _ = len(v1)
    } else {
        v7 := v3[2147483647]
        _ = v7
    }
    var v8 int16 = 32767
    var v9 string = "a"
    delete(v3, 949)
    for fi10 := 0; fi10 < 7; fi10++ {
        switch (-(fi10)) {
        case 1:
            v4 -= (-(v4))
        case -2147483648:
            fmt.Println(v9, v4, v9)
        case 602:
            var v11 int32 = 2147483647
            _ = v11
        default:
            var v12 uint64 = 2
            _ = v12
        }
        _ = fi10
    }
    v9 = v9
    var v13 int16 = 32767
    fmt.Println(v13)
    v13 += -32768
    v4 = (v4 | (v4 >> ((18446744073709551615) & 63)))
    if true {
        var v14 int = 32767
        switch ("\n\t" + "foo bar") {
        case "a":
            delete(v5, v9)
        case "\n\t":
            v2 += 0
        default:
            v15 := []int{127, v14}
            _ = v15
        }
        _ = v14
    }
    var v16 int8 = v2
    if ((4294967295 != 735) && true) {
        v4 = v4
    } else {
        if (true || false) {
            var v17 float64 = 1e10
            _ = v17
        }
    }
    var v18 int64 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
