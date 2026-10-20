package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int32 }
func genFn0(p0 int8, p1 uint64, p2 int) uint8 {
    v6 := GenS1{F0: true}
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{fi7, 788, (^(fi7))}
        _ = fi7
        _ = v8
    }
    p1 = (p1 >> ((42) & 63))
    switch p2 {
    case 581:
        p1 = p1
    case 1:
        fmt.Println(p2, p1, p0)
    default:
        p0 = p0
    }
    v9 := GenS1{F0: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 1
}
func genFn1(p0 bool, p1 int64, p2 string) uint {
    fmt.Println(p2, p1)
    v10 := GenS1{F0: (p0 || ((((-128 != -128) && p0) && true) || p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 255
}

func main() {
    var v1 float64 = ((3.14159 * 1e10) / 1.0)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    v4 := []string{""}
    var v5 uint16 = 1
    if (9223372036854775807 != 306) {
        switch "\n\t" {
        case "a":
            fmt.Println(v5, v1)
        case "":
            v11 := v2[73]
            _ = v11
        case "foo bar":
            if (false && (0 < 1)) {
                if (v5 < v5) {
                    fmt.Println(v5, v5)
                    v12 := GenS1{F0: false}
                    _ = v12
                } else {
                    fmt.Println(v5, v5)
                    var v13 int8 = 1
                    _ = v13
                }
            } else {
                for fi14 := 0; fi14 < 10; fi14++ {
                    fmt.Println(fi14, fi14, fi14)
                    if (!(!(true || (600 != 2)))) {
                        v15 := GenS1{F0: true}
                        switch (fi14 ^ -1) {
                        case 127:
                            var v16 bool = (39 == 0)
                            _ = v16
                        default:
                            for fi17 := 0; fi17 < 1; fi17++ {
                                v3[65535] = true
                                _ = fi17
                            }
                        }
                        _ = v15
                    } else {
                        var v18 bool = (-1 <= -12)
                        _ = len(v4)
                        for fi19 := 0; fi19 < 14; fi19++ {
                            v1 -= v1
                            fmt.Println(v1, v1)
                            fmt.Println(v5, fi19)
                            _ = fi19
                        }
                        _ = v18
                    }
                    var v20 uint16 = v5
                    _ = fi14
                    _ = v20
                }
            }
        }
        v21 := GenS1{F0: true}
        _ = v21
    } else {
        if true {
            switch 127 {
            case -2147483648:
                var v22 float64 = (v1 + v1)
                _ = v22
            }
        }
        fmt.Println(v1, v5)
    }
    fmt.Println(v1, v1)
    if true {
        var v23 int16 = 0
        v24 := GenS2{F0: 4294967295, F1: 2147483647}
        _ = v23
        _ = v24
    } else {
        _ = len(v4)
    }
    v1 = v1
    switch ("hello" + "a") {
    case "\n\t":
        v1 = 1e10
    case "hello":
        v5 = (^(v5))
    case "":
        for fi25 := 0; fi25 < 19; fi25++ {
            v5 += (v5 | (v5 ^ v5))
            _ = fi25
        }
    }
    if len(v4) > 0 {
        v4[0] = (("hello" + (("foo bar" + "foo bar") + "")) + "\n\t")
    }
    v1 -= v1
    var v26 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
