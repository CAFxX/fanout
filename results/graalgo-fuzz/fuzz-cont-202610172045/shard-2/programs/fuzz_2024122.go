package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() uint16 {
    var v5 uint64 = 755
    var v6 uint8 = 128
    if true {
        v7 := []int{-2147483648}
        v5 -= (v5 + v5)
        v8 := GenS1{F0: ((v6 % ((177) | 1)) > (v6 % ((112) | 1)))}
        _ = v7
        _ = v8
    } else {
        switch "\n\t" {
        case "foo bar":
            v6 += (^(v6))
        case "":
            v9 := []int{-1, 1, 42, 32767}
            _ = v9
        }
        v5 = (v5 >> ((255) & 63))
        fmt.Println(v5, v6)
    }
    v6 -= v6
    var v10 uint32 = 258
    _ = v10
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 2
    var v3 int32 = 0
    v4 := make(map[uint32]uint32)
    fmt.Println(v3, v3)
    for fi11 := 0; fi11 < 8; fi11++ {
        v1 += (float32(0.0) * v1)
        switch fi11 {
        case 159:
            switch (fi11 << ((18446744073709551615) & 63)) {
            case 1:
                switch "" {
                case "":
                    v2 += 255
                case "foo bar":
                    v2 += v2
                }
            }
        case -9223372036854775808:
            v3 = (v3 * 42)
        default:
            if (77 != 1) {
                if (v1 > 1e10) {
                    fmt.Println(fi11, fi11, fi11)
                    v12 := GenS1{F0: (33 != -47)}
                    v13 := []int{(-(fi11)), 2147483647, (fi11 << ((v2) & 63))}
                    _ = v12
                    _ = v13
                }
                var v14 int32 = v3
                _ = v14
            }
        }
        _ = fi11
    }
    v2 = 388
    var v15 string = ("foo bar" + "hello")
    fmt.Println(v3, v3, v1)
    var v16 int32 = v3
    v17 := v4[4294967295]
    v1 -= -2.25
    v2 = v2
    v18 := []int{-206, 128, 127}
    var v19 uint8 = 255
    _ = len(v18)
    if true {
        var v20 uint16 = 255
        _ = v20
    } else {
        fmt.Println(v3, v17)
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := v4[(^(v17))]
        v2 += uint(v19)
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: false}
    var v24 int32 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
