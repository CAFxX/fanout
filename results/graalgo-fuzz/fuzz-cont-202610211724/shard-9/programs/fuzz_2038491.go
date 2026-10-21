package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint8 }
func genFn0(p0 uint32, p1 bool, p2 uint32) int {
    var v7 int = 4294967295
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: -590, F1: false, F2: 128}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return v7
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 0
    v3 := make(map[uint32]int)
    var v4 uint64 = 18446744073709551615
    var v5 int16 = 0
    var v6 int64 = -285
    var v10 uint16 = 42
    for fi11 := 0; fi11 < 5; fi11++ {
        fmt.Println(v4, v2)
        v12 := GenS1{F0: v5, F1: false, F2: uint8(v2)}
        v5 += 0
        _ = fi11
        _ = v12
    }
    if (!(false && (0 >= 956))) {
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 float64 = 1e10
            var v15 int64 = -1
            switch "\n\t" {
            case "foo bar":
                v16 := GenS1{F0: 0, F1: ((1e10 / 1.5) == 3.14159), F2: 36}
                _ = v16
            case "\n\t":
                if true {
                    v2 = v2
                    v6 = -1
                    for fi17 := 0; fi17 < 10; fi17++ {
                        v2 = (v2 << ((v2) & 63))
                        v18 := []int{(fi13 ^ (fi13 + (fi13 << (((^(v2))) & 63))))}
                        _ = fi17
                        _ = v18
                    }
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        v5 = v5
                    case "":
                        switch (fi13 | genFn0(2, true, 2)) {
                        case -9223372036854775808:
                            switch ("\n\t" + "") {
                            case "":
                                v2 -= (^(v2))
                            }
                        }
                    }
                }
            case "hello":
                v10 = 2
            default:
                var v19 string = ("\n\t" + (("a" + "a") + "hello"))
                _ = v19
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v4 -= v4
    fmt.Println(v2, v2, v2)
    if (false || false) {
        if (!(false && true)) {
            switch ("" + "\n\t") {
            case "foo bar":
                _ = len(v3)
            case "\n\t":
                if ((88 < 2) && ((^(v10)) == v10)) {
                    var v20 int16 = 1
                    _ = v20
                }
            case "":
                var v21 bool = true
                _ = v21
            }
            var v22 uint32 = 388
            _ = v22
        }
        v2 = v2
    } else {
        for fi23 := 0; fi23 < 18; fi23++ {
            fmt.Println(v5, v2, v10)
            _ = fi23
        }
    }
    switch 32767 {
    case 42:
        v3[873] = 1
    default:
        delete(v1, genFn0(2147483648, ((0.1 >= 1e10) && ((!(100 < -1)) || (false && true))), 2147483648))
    }
    fmt.Println(v4, v6, v4)
    if true {
        v6 -= 0
    }
    _ = len(v1)
    v24 := []int{128, 32768, genFn0(0, false, 371)}
    if (!true) {
        if (!((!true) && false)) {
            fmt.Println(v6, v6)
        }
    }
    var v25 uint16 = (v10 | (v10 + (v10 * (v10 / ((uint16(v4)) | 1)))))
    var v26 int8 = int8(v5)
    fmt.Println(v6, v26)
    if false {
        v27 := GenS1{F0: v5, F1: (v6 < -9223372036854775808), F2: 42}
        _ = v27
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 int32 = -1
        v25 = v25
        _ = len(v24)
        _ = fi28
        _ = v29
    }
    v30 := GenS1{F0: int16(v4), F1: false, F2: 20}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
