package main

import (
    "fmt"
)


func main() {
    var v1 uint = 0
    var v2 float32 = ((((((1e10 + -2.25) / (((((1.5 + ((((((1e10 + 1e10) + 0.0) + (((((((3.14159 / 1.0) / float32(-2.25)) / 3.14159) * (((1e10 * 0.1) / 3.14159) + 3.14159)) / -2.25) + 0.0) * ((float32(1.5) - ((3.14159 + (1e10 + 0.1)) - (0.1 * 0.0))) * ((3.14159 / ((-2.25 / 0.1) + -2.25)) / (3.14159 + (0.1 / -2.25)))))) + 0.1) * 3.14159) / 1.5)) * float32((0.0 - 0.1))) * (3.14159 / 1.0)) + ((((-2.25 * 1e10) / 0.1) / float32(0.1)) / 1.5)) - (1.5 / 0.1))) * 0.1) + (1e10 / float32((3.14159 * (3.14159 + (1.5 - (1e10 * 0.0))))))) - 0.1) * -2.25)
    v3 := make(map[string]string)
    v4 := []int{0, 142, 619, 9223372036854775807}
    v5 := v3[""]
    var v6 int16 = 42
    if len(v4) > 0 {
        v4[0] = -1
    }
    if (false || true) {
        fmt.Println(v1, v5, v2)
        v7 := v3["foo bar"]
        if ((3.14159 * 0.1) <= (1.5 * 1e10)) {
            v2 -= 3.14159
            v6 -= v6
        } else {
            v5 = (v5 + (("hello" + v5) + v7))
            v1 -= (v1 << ((v1) & 63))
            v6 = 1
        }
        _ = v7
    } else {
        var v8 uint32 = 985
        var v9 int32 = -838
        _ = v8
        _ = v9
    }
    v5 = v5
    v1 += (v1 >> ((v1) & 63))
    v1 = (v1 + 362)
    fmt.Println(v5)
    v2 = 1e10
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
