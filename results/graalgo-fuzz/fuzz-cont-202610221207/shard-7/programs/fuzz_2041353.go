package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int64 }

func main() {
    var v1 uint8 = 185
    v2 := make(map[int]bool)
    var v3 int = -2147483648
    var v4 int8 = -100
    var v5 uint32 = 2147483648
    var v6 uint16 = 32767
    for fi7 := 0; fi7 < 15; fi7++ {
        v2[(v3 & 564)] = (-2.25 < 3.14159)
        if true {
            var v8 float32 = (1.5 + 1e10)
            var v9 int = -9223372036854775808
            _ = v8
            _ = v9
        }
        var v10 float32 = 1e10
        _ = fi7
        _ = v10
    }
    switch v3 {
    case -194:
        if true {
            v3 -= v3
            var v11 string = ("" + "\n\t")
            _ = v11
        } else {
            v5 += v5
            if false {
                var v12 int = (v3 >> ((674) & 63))
                _ = v12
            } else {
                v1 = v1
            }
            fmt.Println(v6)
        }
    }
    switch "foo bar" {
    case "\n\t":
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 bool = true
            v2[-1] = true
            fmt.Println(v5)
            _ = fi13
            _ = v14
        }
    case "":
        for fi15 := 0; fi15 < 20; fi15++ {
            v1 = 2
            switch "\n\t" {
            case "":
                v16 := []int{(fi15 - fi15), (fi15 * 4294967295), (v3 >> ((4) & 63)), v3}
                _ = v16
            case "a":
                v17 := []int{fi15}
                _ = v17
            case "hello":
                delete(v2, (v3 << ((4294967295) & 63)))
            }
            v5 -= 266
            _ = fi15
        }
    case "a":
        v4 += v4
    }
    v18 := []int{(^(v3)), (v3 >> ((939) & 63)), v3, (-(v3))}
    v3 -= v3
    for fi19 := 0; fi19 < 15; fi19++ {
        v20 := GenS1{F0: fi19}
        for fi21 := 0; fi21 < 18; fi21++ {
            v1 = (^(v1))
            if (4294967295 < 42) {
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                v18 = append(v18, 2147483647)
            } else {
                if (-1 != 639) {
                    fmt.Println(v6, fi19, v1)
                    if true {
                        fmt.Println(v1, v5)
                        v5 += v5
                    } else {
                        v22 := GenS2{F0: 929, F1: 9223372036854775807}
                        var v23 float32 = (0.0 * 1e10)
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS2{F0: 1, F1: -463}
                    _ = v24
                } else {
                    v6 += v6
                }
            }
            _ = fi21
        }
        for fi25 := 0; fi25 < 18; fi25++ {
            v5 += uint32(v6)
            var v26 uint8 = v1
            _ = fi25
            _ = v26
        }
        _ = fi19
        _ = v20
    }
    v3 = 384
    v6 = v6
    v27 := v2[v3]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
