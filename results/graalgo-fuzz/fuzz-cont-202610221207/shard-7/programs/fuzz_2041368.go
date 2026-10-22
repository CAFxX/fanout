package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint }
type GenS2 struct { F0 int64 }
func genFn0(p0 int64, p1 int16) int64 {
    v6 := []int{525, -165, 681, 32768}
    var v7 uint = 1
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    switch -9223372036854775808 {
    case -1:
        var v8 bool = (-160 <= 920)
        _ = v8
    case 915:
        var v9 int32 = 2147483647
        _ = v9
    }
    fmt.Println(p0, v7)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p0 - p0)
}

func main() {
    var v1 uint32 = 2147483648
    var v2 float64 = (3.14159 / (1.5 / 1e10))
    var v3 int = 42
    v4 := make(map[int]bool)
    var v5 uint64 = 1
    fmt.Println(v5)
    var v10 string = (("hello" + "\n\t") + (("foo bar" + ("hello" + "foo bar")) + "hello"))
    if (77 >= 139) {
        var v11 uint16 = 527
        _ = v11
    }
    fmt.Println(v1)
    fmt.Println(v5)
    v12 := GenS1{F0: (v5 << ((4294967295) & 63)), F1: 1, F2: 237}
    if ((-2.25 / ((3.14159 / 1.5) * float64((((0.1 - -2.25) + (3.14159 - ((1e10 / (0.1 - 0.0)) * (1.5 + 1.5)))) - 0.0)))) >= float64(1.5)) {
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 int16 = -32768
            var v15 int32 = 0
            _ = fi13
            _ = v14
            _ = v15
        }
        if ((-1 < 42) && false) {
            switch "foo bar" {
            case "foo bar":
                var v16 float32 = 1e10
                _ = v16
            case "":
                var v17 string = v10
                _ = v17
            default:
                var v18 uint32 = (v1 >> ((uint(v1)) & 63))
                _ = v18
            }
            if (true && true) {
                v19 := []int{v3}
                _ = v19
            }
        } else {
            switch (v3 + v3) {
            case 0:
                if true {
                    var v20 int = (v3 & (v3 / ((v3) | 1)))
                    _ = v20
                } else {
                    v2 += v2
                    v21 := GenS2{F0: 406}
                    _ = v21
                }
            default:
                v2 += v2
            }
            if true {
                for fi22 := 0; fi22 < 1; fi22++ {
                    fmt.Println(v5, v3, fi22)
                    switch v3 {
                    case -700:
                        if (!((127 == 128) || false)) {
                            v10 = v10
                            var v23 int8 = -77
                            _ = v23
                        }
                    case 128:
                        for fi24 := 0; fi24 < 6; fi24++ {
                            switch 42 {
                            case 565:
                                var v25 uint32 = (v1 >> ((1) & 63))
                                _ = v25
                            case 32768:
                                var v26 int32 = -245
                                _ = v26
                            default:
                                fmt.Println(v10, fi24, v3)
                            }
                            _ = fi24
                        }
                    }
                    v4[(v3 / (((-(v3))) | 1))] = false
                    _ = fi22
                }
                var v27 bool = true
                v28 := []int{v3}
                _ = v27
                _ = v28
            }
            v29 := v4[v3]
            _ = v29
        }
        v30 := v4[9223372036854775807]
        _ = v30
    }
    if false {
        var v31 int = v3
        var v32 int16 = -1
        var v33 int = (v31 - (v3 ^ v31))
        _ = v31
        _ = v32
        _ = v33
    }
    v3 = v3
    var v34 uint32 = v1
    fmt.Println(v10, v3)
    switch -246 {
    case 653:
        switch ("" + "hello") {
        case "foo bar":
            v35 := []int{int(v5), -558, v3}
            _ = v35
        case "a":
            var v36 string = ""
            _ = v36
        case "":
            fmt.Println(v3)
        default:
            v37 := v4[(v3 % ((v3) | 1))]
            _ = v37
        }
    case 32768:
        var v38 float32 = float32(-2.25)
        _ = v38
    case 1:
        v1 += 2147483647
    }
    v39 := GenS2{F0: -1}
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v5
}
