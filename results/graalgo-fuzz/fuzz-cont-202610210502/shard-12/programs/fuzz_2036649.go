package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 bool = false
    v2 := []int{127, 1, -600}
    var v3 bool = (2 <= 593)
    var v4 int = 0
    v3 = (v1 || v1)
    v1 = v1
    var v5 int32 = 1
    var v6 float32 = 1e10
    fmt.Println(v5, v4)
    fmt.Println(v4, v5)
    var v7 int8 = -1
    v8 := GenS1{F0: 0}
    fmt.Println(v1, v6, v3)
    fmt.Println(v5, v1)
    var v9 uint = 255
    var v10 uint64 = 42
    switch (v4 / (((v4 - (v4 * 9223372036854775807))) | 1)) {
    case 127:
        if v1 {
            fmt.Println(v6, v1, v7)
            switch (v4 - -2147483648) {
            case 128:
                for fi11 := 0; fi11 < 10; fi11++ {
                    _ = len(v2)
                    if (!(v1 && (v7 < v7))) {
                        if len(v2) > 0 {
                            v2[0] = (v4 & (v4 << ((2) & 63)))
                        }
                    } else {
                        v2 = append(v2, (^(v4)))
                        for fi12 := 0; fi12 < 13; fi12++ {
                            var v13 int8 = (v7 | 55)
                            var v14 uint64 = v10
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                    }
                    _ = fi11
                }
            case 663:
                fmt.Println(v3)
            case 9223372036854775807:
                for fi15 := 0; fi15 < 6; fi15++ {
                    var v16 uint8 = 127
                    v17 := GenS1{F0: 2}
                    v7 = v7
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            default:
                var v18 uint8 = 255
                _ = v18
            }
        }
    case 32768:
        for fi19 := 0; fi19 < 13; fi19++ {
            if len(v2) > 0 {
                v2[0] = (v4 | 2147483647)
            }
            var v20 int64 = 0
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v5, v6, v20)
            case "":
                var v21 float32 = 0.1
                _ = v21
            case "hello":
                var v22 bool = ("hello" < ("\n\t" + "a"))
                _ = v22
            default:
                for fi23 := 0; fi23 < 5; fi23++ {
                    v6 = 1.5
                    var v24 uint32 = 2147483648
                    _ = fi23
                    _ = v24
                }
            }
            _ = fi19
            _ = v20
        }
    }
    if v1 {
        var v25 int64 = 42
        if (42 != 2) {
            v9 -= (v9 | (v9 * (v9 % ((v9) | 1))))
        }
        _ = v25
    } else {
        var v26 uint32 = 2
        for fi27 := 0; fi27 < 10; fi27++ {
            v7 = v7
            _ = fi27
        }
        v28 := GenS1{F0: 42}
        _ = v26
        _ = v28
    }
    v7 += 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
