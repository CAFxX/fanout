package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint32, p1 string, p2 int) int8 {
    switch (p1 + "\n\t") {
    case "foo bar":
        for fi7 := 0; fi7 < 10; fi7++ {
            if (!true) {
                switch (^(p2)) {
                case 128:
                    p0 -= p0
                }
                if true {
                    p0 += (p0 - p0)
                    v8 := []int{(fi7 - (p2 % ((fi7) | 1))), (fi7 - p2), (p2 * (fi7 & (p2 - 42))), 32767}
                    _ = v8
                }
            } else {
                if true {
                    for fi9 := 0; fi9 < 8; fi9++ {
                        var v10 uint16 = 1
                        fmt.Println(p0, p2)
                        _ = fi9
                        _ = v10
                    }
                    v11 := GenS1{F0: (1.5 <= 0.1)}
                    var v12 uint = 0
                    _ = v11
                    _ = v12
                }
            }
            fmt.Println(p1, p2)
            _ = fi7
        }
    default:
        var v13 string = ("a" + p1)
        _ = v13
    }
    p1 = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    return 10
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 2
    var v3 uint16 = 65535
    v4 := []string{("foo bar" + "a"), ("foo bar" + ""), "a", "hello"}
    v5 := make(map[int]int)
    var v6 uint8 = 147
    if (("hello" + ("a" + "\n\t")) < ("a" + "a")) {
        v2 = v2
        for fi14 := 0; fi14 < 13; fi14++ {
            v3 = (^(v3))
            _ = fi14
        }
    }
    v6 += v6
    var v15 string = "hello"
    var v16 float64 = 1e10
    var v17 float64 = v16
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 float32 = (-2.25 * 3.14159)
        _ = fi18
        _ = v19
    }
    if true {
        v3 = (v3 | 1)
    } else {
        if false {
            if false {
                switch -9223372036854775808 {
                case -9223372036854775808:
                    var v20 int32 = -125
                    _ = v20
                case 32767:
                    for fi21 := 0; fi21 < 16; fi21++ {
                        v22 := GenS1{F0: false}
                        var v23 uint8 = v6
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                default:
                    v15 = ""
                }
            }
        } else {
            if false {
                fmt.Println(v17)
                for fi24 := 0; fi24 < 5; fi24++ {
                    v17 -= v16
                    v25 := v1[141]
                    _ = fi24
                    _ = v25
                }
                for fi26 := 0; fi26 < 18; fi26++ {
                    fmt.Println(v17, v15)
                    if (2147483647 == 2147483647) {
                        v15 = (v15 + "\n\t")
                        switch v15 {
                        case "a":
                            v27 := GenS1{F0: ((fi26 ^ 127) != fi26)}
                            _ = v27
                        default:
                            var v28 int32 = 42
                            _ = v28
                        }
                        var v29 int32 = -2147483648
                        _ = v29
                    }
                    fmt.Println(v17, v3, v3)
                    _ = fi26
                }
            } else {
                v30 := GenS1{F0: false}
                v4 = append(v4, "a")
                _ = v30
            }
            fmt.Println(v6, v15)
        }
    }
    switch -550 {
    case 128:
        var v31 uint = v2
        _ = v31
    case 406:
        v16 -= v17
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        if (0.1 != v17) {
            fmt.Println(v2, v17)
            fmt.Println(v17, v2, v6)
        }
        switch v15 {
        case "":
            v2 -= v2
        case "hello":
            v4 = append(v4, (v15 + ""))
        case "foo bar":
            v16 = v17
        default:
            for fi33 := 0; fi33 < 13; fi33++ {
                if len(v4) > 0 {
                    v4[0] = v15
                }
                switch -1 {
                case 32767:
                    v3 = (v3 | (^(v3)))
                case -85:
                    fmt.Println(v6)
                case 32768:
                    fmt.Println(v17, v6)
                }
                _ = fi33
            }
        }
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
