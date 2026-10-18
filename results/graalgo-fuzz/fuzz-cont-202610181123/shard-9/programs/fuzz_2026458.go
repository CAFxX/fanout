package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int32 }
func genFn0(p0 int32, p1 uint64) bool {
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 uint32 = 2
        p0 = p0
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: "hello", F1: ("foo bar" + ((("\n\t" + "hello") + "hello") + "\n\t")), F2: p0}
    _ = p0
    _ = p1
    _ = v8
    return (false || true)
}

func main() {
    var v1 string = "hello"
    v2 := make(map[int]int)
    v3 := []string{v1, v1}
    var v4 int16 = 0
    var v5 int64 = -9223372036854775808
    fmt.Println(v1, v1)
    var v9 uint64 = 0
    v5 -= v5
    var v10 float64 = (1.5 + 3.14159)
    v4 += (v4 << ((4294967295) & 63))
    if false {
        v3 = append(v3, v1)
    }
    var v11 bool = (42 >= 93)
    switch ("a" + v1) {
    case "foo bar":
        v9 += (v9 * 0)
    case "hello":
        if v11 {
            switch "a" {
            case "a":
                v1 = ""
            case "hello":
                switch "hello" {
                case "hello":
                    var v12 float64 = v10
                    _ = v12
                case "foo bar":
                    var v13 uint64 = (^(v9))
                    _ = v13
                case "a":
                    var v14 uint32 = 1
                    _ = v14
                }
            case "\n\t":
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 string = ("hello" + v1)
                    _ = fi15
                    _ = v16
                }
            default:
                var v17 int32 = 1
                _ = v17
            }
        } else {
            if (-2.25 <= -2.25) {
                switch 32767 {
                case 495:
                    fmt.Println(v10, v4, v1)
                case 0:
                    v5 += v5
                case -852:
                    var v18 bool = ((v10 == v10) && (!(!v11)))
                    _ = v18
                default:
                    var v19 uint8 = 0
                    _ = v19
                }
                var v20 int = 761
                fmt.Println(v4, v4)
                _ = v20
            }
            if len(v3) > 0 {
                v3[0] = ((v1 + v1) + v1)
            }
            if (307 >= 18) {
                var v21 float64 = v10
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                for fi22 := 0; fi22 < 11; fi22++ {
                    switch (("hello" + v1) + "\n\t") {
                    case "a":
                        v10 += (1.5 * v10)
                    case "foo bar":
                        v4 -= v4
                    case "hello":
                        fmt.Println(v4, v10)
                    }
                    switch ("\n\t" + (v1 + v1)) {
                    case "a":
                        v23 := v2[(fi22 - (fi22 >> ((120) & 63)))]
                        _ = v23
                    case "\n\t":
                        var v24 uint16 = 342
                        _ = v24
                    case "hello":
                        for fi25 := 0; fi25 < 18; fi25++ {
                            v3 = append(v3, "foo bar")
                            var v26 uint64 = (^(v9))
                            _ = fi25
                            _ = v26
                        }
                    }
                    _ = fi22
                }
                _ = v21
            }
        }
    case "\n\t":
        var v27 int = 42
        _ = v27
    default:
        fmt.Println(v11, v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
