package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint16 = 42
    var v2 uint = 1
    v3 := []uint32{453}
    v4 := []int{-9223372036854775808, 127, 4294967295}
    var v5 uint = 255
    var v6 bool = (true && (1e10 != 0.0))
    var v7 uint16 = v1
    if (521 != -528) {
        v8 := GenS1{F0: v2}
        for fi9 := 0; fi9 < 2; fi9++ {
            v7 -= v7
            v10 := GenS1{F0: (v2 ^ 255)}
            _ = fi9
            _ = v10
        }
        fmt.Println(v5, v5)
        _ = v8
    }
    var v11 int8 = 55
    var v12 uint32 = 2147483648
    if v6 {
        for fi13 := 0; fi13 < 6; fi13++ {
            switch (fi13 - (fi13 - -194)) {
            case 42:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 9223372036854775807:
                switch "" {
                case "hello":
                    if (0 != 9223372036854775807) {
                        fmt.Println(v12, v7, v2)
                        v7 += (v1 << ((756) & 63))
                        if true {
                            v14 := GenS1{F0: v5}
                            if false {
                                fmt.Println(v2, v1, v7)
                                var v15 uint8 = 93
                                _ = v15
                            }
                            fmt.Println(fi13)
                            _ = v14
                        } else {
                            if v6 {
                                var v16 bool = (42 != 2)
                                var v17 int8 = 42
                                v18 := GenS1{F0: v2}
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                        }
                    } else {
                        v1 -= (^(v7))
                    }
                case "a":
                    var v19 bool = true
                    _ = v19
                default:
                    var v20 float32 = float32(3.14159)
                    _ = v20
                }
            case -2147483648:
                var v21 float64 = 0.1
                _ = v21
            default:
                v1 = v1
            }
            v22 := GenS1{F0: 65535}
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 int = (fi13 & (-(fi23)))
                _ = fi23
                _ = v24
            }
            _ = fi13
            _ = v22
        }
    }
    var v25 uint16 = v7
    var v26 int64 = 0
    switch ("foo bar" + ("foo bar" + "")) {
    case "a":
        v12 += 704
    }
    if (1.5 >= float64(3.14159)) {
        v27 := GenS1{F0: (v2 >> ((4294967295) & 63))}
        v4 = append(v4, -676)
        _ = v27
    } else {
        v5 = v2
        _ = len(v3)
    }
    switch 2147483647 {
    case 4294967295:
        v28 := GenS1{F0: (v2 + v5)}
        _ = v28
    }
    var v29 uint32 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
