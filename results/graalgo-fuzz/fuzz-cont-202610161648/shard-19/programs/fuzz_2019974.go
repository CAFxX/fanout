package main

import (
    "fmt"
)

func genFn0(p0 uint64) int16 {
    var v6 uint64 = 2
    v7 := []int{-1, -203, 487, -683}
    fmt.Println(v6)
    var v8 float64 = 0.0
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 string = ("hello" + "")
    v2 := []string{"", v1, v1, ((v1 + "\n\t") + (v1 + (v1 + "foo bar")))}
    v3 := []uint32{42, 2}
    var v4 uint8 = 127
    var v5 float64 = float64((0.0 / ((1.5 - 3.14159) + (((1.5 / ((3.14159 * -2.25) - (((0.0 + 0.0) + 1e10) - (float32(3.14159) + -2.25)))) * -2.25) * 3.14159))))
    var v9 string = v1
    for fi10 := 0; fi10 < 17; fi10++ {
        v4 = 2
        fmt.Println(v9, v4)
        var v11 int32 = 42
        _ = fi10
        _ = v11
    }
    v2 = append(v2, (v1 + "foo bar"))
    var v12 uint64 = 1
    var v13 int64 = -1
    if len(v2) > 0 {
        v2[0] = (v1 + ((("" + ("hello" + v1)) + v1) + (("foo bar" + (v9 + "")) + ("" + v1))))
    }
    v9 = v9
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    if false {
        fmt.Println(v4)
    }
    v5 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
