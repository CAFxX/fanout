package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int8; F2 bool }
type GenS4 struct { F0 int64 }
func genFn0(p0 int64, p1 int, p2 uint) int8 {
    fmt.Println(p1, p2)
    if ((p2 * 722) <= 290) {
        if true {
            var v7 uint8 = 152
            v8 := GenS1{F0: 4294967295}
            _ = v7
            _ = v8
        }
        fmt.Println(p0, p2)
    }
    v9 := []int{(p1 + int(p2)), -132}
    var v10 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -7
}

func main() {
    v1 := []string{}
    var v2 int64 = 44
    var v3 uint16 = 32767
    v4 := make(map[int]uint32)
    v5 := []string{(("foo bar" + "a") + "a")}
    v6 := make(map[string]uint32)
    if (2147483647 == -1) {
        v2 += v2
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v11 int16 = 1
        _ = v11
    }
    _ = len(v6)
    var v12 int64 = v2
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            switch ("hello" + (((("a" + "foo bar") + "a") + "a") + "")) {
            case "\n\t":
                switch "a" {
                case "\n\t":
                    var v15 int = (fi14 - (fi14 - 498))
                    _ = v15
                }
            case "":
                var v16 uint = 42
                _ = v16
            default:
                fmt.Println(v12)
            }
            if (false && false) {
                if (42 != 892) {
                    v17 := GenS2{F0: 235}
                    _ = len(v5)
                    _ = v17
                } else {
                    v3 = (v3 | 698)
                    if false {
                        var v18 bool = true
                        var v19 uint32 = 110
                        var v20 uint64 = 27
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        var v21 uint = 42
                        fmt.Println(v21, v3)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v21
                    }
                }
                v12 = v12
                v22 := GenS2{F0: -287}
                _ = v22
            }
            v23 := GenS1{F0: 2}
            _ = fi14
            _ = v23
        }
        v1 = append(v1, "a")
        _ = fi13
    }
    v2 += (^(v2))
    switch "" {
    case "\n\t":
        v12 = v2
    case "":
        var v24 string = ("a" + "a")
        _ = v24
    default:
        for fi25 := 0; fi25 < 8; fi25++ {
            for fi26 := 0; fi26 < 15; fi26++ {
                fmt.Println(v2)
                v27 := GenS3{F0: 2, F1: 42, F2: (!true)}
                var v28 int = -2147483648
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi25
        }
    }
    switch (("hello" + "hello") + "foo bar") {
    case "":
        var v29 uint = 42
        _ = v29
    case "hello":
        delete(v6, "hello")
    default:
        if (float64((float32(-2.25) + (1e10 - 3.14159))) > 1e10) {
            v3 += (v3 | (v3 - v3))
        } else {
            if true {
                if (true || true) {
                    v3 += v3
                    var v30 float64 = -2.25
                    _ = v30
                }
            }
            v1 = append(v1, "hello")
            for fi31 := 0; fi31 < 2; fi31++ {
                for fi32 := 0; fi32 < 3; fi32++ {
                    v3 = v3
                    _ = fi32
                }
                _ = fi31
            }
        }
    }
    var v33 uint64 = 4294967295
    var v34 int = 1
    fmt.Println(v34)
    v35 := GenS1{F0: (v33 >> ((1) & 63))}
    delete(v4, v34)
    v33 = v33
    _ = len(v5)
    fmt.Println(v3, v34, v33)
    v36 := GenS4{F0: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
