package main

import (
    "fmt"
)


func main() {
    var v1 bool = true
    var v2 float64 = 1e10
    v3 := []int{1}
    var v4 int16 = -32768
    v1 = true
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint16 = 2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi5
        _ = v6
    }
    v1 = v1
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            if (1.5 <= (-2.25 / 3.14159)) {
                switch (("a" + "") + "hello") {
                case "hello":
                    v4 = v4
                }
            }
            if (800 != -9223372036854775808) {
                fmt.Println(v1)
            }
            var v9 int16 = v4
            _ = fi8
            _ = v9
        }
        fmt.Println(fi7)
        var v10 float32 = float32(v2)
        _ = fi7
        _ = v10
    }
    switch "a" {
    case "hello":
        var v11 uint16 = 255
        _ = v11
    }
    v4 = v4
    v4 = v4
    if (v1 || v1) {
        fmt.Println(v1)
        v4 += (-(v4))
    }
    var v12 int16 = v4
    v12 -= v12
    var v13 int8 = -1
    v12 = v12
    var v14 float32 = ((((0.0 * (3.14159 - 0.0)) - 0.0) * float32(1.5)) - (((float32(1.5) * 3.14159) - 1e10) * float32((0.0 * 3.14159))))
    var v15 int = -9223372036854775808
    if (770 >= 0) {
        for fi16 := 0; fi16 < 9; fi16++ {
            fmt.Println(v14)
            _ = fi16
        }
    }
    fmt.Println(v1)
    var v17 float32 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
