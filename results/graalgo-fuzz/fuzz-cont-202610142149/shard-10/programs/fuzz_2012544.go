package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int64, p1 uint16, p2 int32) bool {
    p2 -= p2
    var v7 bool = (true || ((4294967295 > 42) && true))
    var v8 uint64 = 4294967295
    if (3.14159 == (0.1 / 1.0)) {
        if true {
            var v9 bool = (-62 < int16(p2))
            switch 4294967295 {
            case 4294967295:
                var v10 bool = v7
                _ = v10
            case 42:
                var v11 int8 = 1
                _ = v11
            case 32768:
                var v12 int64 = int64(p1)
                _ = v12
            }
            p1 = (p1 % ((p1) | 1))
            _ = v9
        } else {
            switch 1 {
            case 177:
                if ((p1 & (p1 * (p1 - (p1 | (p1 % ((p1) | 1)))))) <= 255) {
                    for fi13 := 0; fi13 < 5; fi13++ {
                        switch (fi13 - fi13) {
                        case 423:
                            var v14 string = "hello"
                            _ = v14
                        case 32767:
                            var v15 int = (fi13 & (-(fi13)))
                            _ = v15
                        default:
                            var v16 int16 = -106
                            _ = v16
                        }
                        fmt.Println(p2, p1, v8)
                        _ = fi13
                    }
                    p1 += p1
                    if true {
                        var v17 bool = (p2 != p2)
                        v18 := []int{}
                        _ = v17
                        _ = v18
                    }
                } else {
                    switch 32767 {
                    case -2147483648:
                        p0 += p0
                    case -1:
                        var v19 string = "\n\t"
                        _ = v19
                    case 282:
                        v20 := []int{128, 2147483647}
                        _ = v20
                    default:
                        v21 := []int{128, -2147483648, -1}
                        _ = v21
                    }
                }
            case 897:
                p2 += p2
            }
            v7 = false
            v8 = (v8 - v8)
        }
        var v22 bool = true
        var v23 uint16 = 255
        _ = v22
        _ = v23
    } else {
        fmt.Println(p2, p0, p0)
    }
    if (("a" + "\n\t") == "hello") {
        p1 += p1
        fmt.Println(v7)
        fmt.Println(v8)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ((p0 & p0) < p0)
}
func genFn1(p0 int32) string {
    fmt.Println(p0)
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 uint64 = 998
        fmt.Println(fi24, p0)
        _ = fi24
        _ = v25
    }
    var v26 bool = true
    if v26 {
        switch "" {
        case "hello":
            var v27 int16 = -694
            _ = v27
        case "a":
            p0 += p0
        default:
            v28 := GenS1{F0: 0}
            _ = v28
        }
    } else {
        v26 = v26
        var v29 int32 = p0
        var v30 float32 = -2.25
        _ = v29
        _ = v30
    }
    var v31 int32 = p0
    _ = p0
    _ = v26
    _ = v31
    return ""
}

func main() {
    v1 := []string{"\n\t"}
    var v2 uint = 42
    v3 := make(map[string]string)
    var v4 float64 = 3.14159
    v5 := make(map[int]int)
    var v6 uint8 = 42
    v5[32768] = int(v2)
    v3["a"] = ("" + "foo bar")
    var v32 bool = (42 >= 1)
    var v33 uint64 = 594
    v4 -= v4
    var v34 int = -2147483648
    var v35 bool = v32
    for fi36 := 0; fi36 < 15; fi36++ {
        var v37 string = "foo bar"
        if v35 {
            for fi38 := 0; fi38 < 11; fi38++ {
                var v39 uint16 = 255
                v32 = (42 < -89)
                v40 := GenS1{F0: -624}
                _ = fi38
                _ = v39
                _ = v40
            }
            v6 = (v6 - v6)
            fmt.Println(v4, v34)
        }
        _ = fi36
        _ = v37
    }
    v32 = (v33 >= 18446744073709551615)
    v35 = (38 >= 2)
    v33 += (v33 & v33)
    for fi41 := 0; fi41 < 16; fi41++ {
        v35 = ((42 == 65535) && (((fi41 << (((v2 ^ (v2 + v2))) & 63)) == (fi41 ^ -2147483648)) || (v32 || (39 >= 2147483647))))
        _ = len(v1)
        v4 += v4
        _ = fi41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
