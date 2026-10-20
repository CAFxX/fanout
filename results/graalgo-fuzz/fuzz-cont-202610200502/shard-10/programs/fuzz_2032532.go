package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
func genFn0() int64 {
    var v6 uint64 = 42
    v7 := []int{9, -800, 9223372036854775807, 127}
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 string = (("\n\t" + ("a" + "hello")) + ("\n\t" + "a"))
    var v3 uint = 18446744073709551615
    var v4 int = 4294967295
    var v5 int = (-(v4))
    for fi8 := 0; fi8 < 9; fi8++ {
        v3 -= 42
        delete(v1, fi8)
        v2 = "a"
        _ = fi8
    }
    var v9 int = v5
    switch v2 {
    case "hello":
        v5 -= (-(v9))
    }
    switch v5 {
    case 32768:
        if true {
            var v10 string = v2
            v11 := GenS1{F0: 32768}
            _ = v10
            _ = v11
        } else {
            v1[v9] = (v4 * -9223372036854775808)
            if false {
                var v12 uint16 = 32767
                _ = v12
            }
        }
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := []int{(fi13 ^ v9), 4294967295, fi13}
        for fi15 := 0; fi15 < 18; fi15++ {
            if ((2 != 2) || true) {
                var v16 uint16 = 42
                v4 += fi15
                v17 := v1[v4]
                _ = v16
                _ = v17
            }
            var v18 int16 = 32767
            if (42 != 119) {
                v19 := GenS1{F0: fi15}
                var v20 uint = 18446744073709551615
                _ = v19
                _ = v20
            } else {
                v5 = (fi13 | (v5 >> ((249) & 63)))
                if (true && ((v9 - (fi15 - 127)) < -573)) {
                    fmt.Println(v4, v5, v5)
                    if (!false) {
                        switch v2 {
                        case "":
                            switch (^(v9)) {
                            case -154:
                                var v21 string = ("" + (v2 + "hello"))
                                _ = v21
                            case 32768:
                                v14 = append(v14, fi13)
                            }
                        case "hello":
                            v9 = v5
                        case "a":
                            v14 = append(v14, int(v3))
                        }
                        v9 = v4
                        switch v2 {
                        case "foo bar":
                            v3 = v3
                        case "a":
                            fmt.Println(v9)
                        case "hello":
                            v22 := GenS2{F0: 42}
                            _ = len(v14)
                            _ = v22
                        default:
                            v23 := GenS2{F0: 32767}
                            _ = v23
                        }
                    } else {
                        fmt.Println(fi15)
                        switch 2147483647 {
                        case -873:
                            v3 += v3
                        case -693:
                            v1[127] = (fi13 & fi15)
                        }
                        var v24 uint64 = 18446744073709551615
                        v4 = (fi13 >> ((v3) & 63))
                        if len(v14) > 0 {
                            v14[0] = v4
                        }
                        _ = v24
                    }
                } else {
                    v5 = fi13
                    var v25 string = v2
                    _ = v25
                }
                v26 := GenS1{F0: 2147483647}
                _ = v26
            }
            v3 = v3
            _ = fi15
            _ = v18
        }
        _ = fi13
        _ = v14
    }
    v3 = (v3 << ((v3) & 63))
    if true {
        v27 := GenS3{F0: -71}
        _ = v27
    }
    var v28 float64 = ((0.0 * 1e10) * -2.25)
    var v29 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
