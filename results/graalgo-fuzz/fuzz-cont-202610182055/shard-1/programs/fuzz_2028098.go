package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    var v3 uint64 = 0
    var v4 uint16 = 530
    var v5 bool = ((!(1e10 <= float64(3.14159))) && (32767 != -705))
    var v6 int8 = 14
    v4 = v4
    for fi7 := 0; fi7 < 1; fi7++ {
        v6 = 1
        var v8 int = (fi7 - fi7)
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 0
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int16 = int16(fi10)
        v12 := GenS1{F0: 525}
        _ = fi10
        _ = v11
        _ = v12
    }
    delete(v2, -23)
    var v13 float32 = float32((1.5 * 0.0))
    v13 -= v13
    if (2 == 255) {
        var v14 int = -9223372036854775808
        if (42 == 42) {
            v15 := GenS1{F0: 489}
            _ = v15
        } else {
            v14 += 42
            var v16 float32 = ((0.0 + (float32((0.1 * (3.14159 / 1.5))) + v13)) * (v13 / (v13 / v13)))
            _ = v16
        }
        _ = v14
    }
    v6 += 1
    v17 := GenS1{F0: 65535}
    var v18 uint32 = 2
    fmt.Println(v6, v5, v13)
    var v19 float32 = v13
    v20 := GenS1{F0: 42}
    if (!v5) {
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(v4, v3, v9)
            v6 -= 127
            v22 := GenS1{F0: 42}
            _ = fi21
            _ = v22
        }
    }
    switch "\n\t" {
    case "hello":
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := v1[(("" + "foo bar") + "foo bar")]
            _ = fi23
            _ = v24
        }
    case "":
        if (3.14159 <= (((3.14159 - (0.1 - 1.5)) + (3.14159 * 0.0)) + 0.1)) {
            v25 := []int{2147483647, -1, 0}
            _ = v25
        } else {
            v6 -= -90
        }
    }
    v13 = ((v13 + v19) / 1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
