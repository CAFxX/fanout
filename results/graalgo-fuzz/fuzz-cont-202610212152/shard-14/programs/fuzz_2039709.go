package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 2
    var v2 int16 = 31
    var v3 float32 = (0.0 + 1.5)
    v4 := make(map[uint32]int)
    v5 := make(map[string]uint32)
    var v6 int64 = 0
    v1 += 255
    var v7 uint64 = 91
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint32 = 2147483648
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int16 = (v2 % (((^(v2))) | 1))
            var v12 uint32 = v9
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    v3 += (float32((((0.0 / ((0.1 + ((1e10 - (0.1 - -2.25)) * 3.14159)) - (float64(v3) + float64((v3 / float32(1.5)))))) / -2.25) * 0.0)) * (v3 + ((float32(1.5) - float32(3.14159)) - (0.0 * v3))))
    var v13 int8 = 125
    if false {
        var v14 bool = false
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v4[2147483647] = (^(fi15))
        _ = fi15
    }
    if (32767 > 0) {
        v3 = v3
    } else {
        if (v7 < v7) {
            v16 := []int{128, 363}
            _ = v16
        } else {
            switch ("\n\t" + "") {
            case "foo bar":
                var v17 string = "a"
                _ = v17
            default:
                var v18 int32 = 42
                _ = v18
            }
            var v19 float64 = (3.14159 * 1e10)
            _ = v19
        }
    }
    var v20 uint = 105
    v3 = v3
    v4[2147483647] = int(v1)
    v2 += 0
    var v21 float64 = 0.0
    delete(v4, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
