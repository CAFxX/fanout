package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 2
    v2 := []uint32{251, 2147483648, 417}
    v3 := []int{9223372036854775807, -9223372036854775808}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = append(v2, 2147483647)
    if true {
        var v4 bool = (1 <= 2147483647)
        var v5 uint64 = 4294967295
        v4 = false
        _ = v4
        _ = v5
    } else {
        var v6 int16 = 32767
        _ = v6
    }
    v1 += (^(v1))
    v1 = v1
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    fmt.Println(v1)
    var v7 bool = false
    if (!(-9223372036854775808 != -9223372036854775808)) {
        if (32767 < -32768) {
            v1 = v1
            var v8 uint8 = 2
            _ = v8
        } else {
            for fi9 := 0; fi9 < 8; fi9++ {
                if (2147483648 != 494) {
                    fmt.Println(v7)
                    var v10 float32 = ((0.1 / (0.0 + 0.1)) * 3.14159)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v10
                }
                _ = fi9
            }
            var v11 uint16 = v1
            v11 = v11
            _ = v11
        }
        switch -1 {
        case 1:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 4294967295:
            var v12 string = "a"
            _ = v12
        case 0:
            if v7 {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        default:
            var v13 uint8 = 128
            _ = v13
        }
    }
    var v14 bool = false
    fmt.Println(v7)
    var v15 float64 = 1e10
    var v16 uint64 = 2
    var v17 uint = 0
    if (0 < -1) {
        switch (("hello" + "a") + "a") {
        case "\n\t":
            var v18 uint64 = v16
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
}
