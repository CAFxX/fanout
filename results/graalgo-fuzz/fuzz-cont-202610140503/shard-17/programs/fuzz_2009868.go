package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int16 }
func genFn0(p0 uint32, p1 uint64) uint16 {
    fmt.Println(p0, p0)
    v7 := []int{-748, -427, 2147483647}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    return 2
}

func main() {
    var v1 uint64 = 0
    var v2 int64 = 800
    v3 := make(map[int]string)
    var v4 int = 0
    var v5 float64 = 1.5
    var v6 int8 = 100
    fmt.Println(v5, v5, v4)
    v8 := v3[(^(v4))]
    var v9 int8 = (-(v6))
    fmt.Println(v2, v1, v4)
    var v10 uint16 = genFn0(42, v1)
    for fi11 := 0; fi11 < 6; fi11++ {
        v10 -= 65535
        if ((v5 - 1e10) != (v5 / v5)) {
            v5 -= -2.25
        } else {
            v2 += (v2 * v2)
            for fi12 := 0; fi12 < 18; fi12++ {
                if false {
                    v13 := GenS1{F0: 2, F1: -2147483648, F2: 228}
                    v2 = (v2 << ((uint(v1)) & 63))
                    _ = v13
                }
                fmt.Println(v10)
                _ = fi12
            }
        }
        _ = fi11
    }
    var v14 int32 = 0
    if (!(v2 >= v2)) {
        if false {
            v15 := GenS1{F0: 127, F1: 42, F2: -1}
            _ = v15
        }
    }
    if false {
        if true {
            var v16 string = v8
            var v17 uint8 = 1
            if (!true) {
                switch v4 {
                case 2147483647:
                    for fi18 := 0; fi18 < 8; fi18++ {
                        var v19 int8 = (v6 << ((639) & 63))
                        switch (((("foo bar" + ((((v16 + (("\n\t" + v16) + (v16 + "hello"))) + v8) + (("" + (v16 + ((v8 + v8) + v8))) + (v8 + v8))) + "\n\t")) + (v8 + v16)) + v8) + "hello") {
                        case "foo bar":
                            fmt.Println(v19)
                        }
                        _ = fi18
                        _ = v19
                    }
                default:
                    var v20 int64 = v2
                    _ = v20
                }
            } else {
                v9 += (^(v9))
                for fi21 := 0; fi21 < 6; fi21++ {
                    var v22 int16 = -302
                    _ = fi21
                    _ = v22
                }
                var v23 int8 = v9
                _ = v23
            }
            _ = v16
            _ = v17
        }
    } else {
        v6 = -1
        var v24 uint32 = 831
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
