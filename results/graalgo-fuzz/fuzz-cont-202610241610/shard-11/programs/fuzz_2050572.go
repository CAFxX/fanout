package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int32 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 bool; F1 int }
func genFn0(p0 int16, p1 uint8) int8 {
    var v6 uint = 397
    v7 := GenS1{F0: 63, F1: (-(p0)), F2: -2147483648}
    v6 += (^(v6))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 int16) uint8 {
    switch -583 {
    case 32767:
        if (65535 == 65535) {
            var v8 int = -2147483648
            v9 := GenS2{F0: 65535, F1: 255}
            _ = v8
            _ = v9
        }
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(p0)
    default:
        switch ("" + (("a" + "hello") + "a")) {
        case "hello":
            v10 := GenS3{F0: ("a" == ("\n\t" + "hello")), F1: 4294967295}
            _ = v10
        }
    }
    _ = p0
    return 127
}

func main() {
    v1 := []string{"hello", "hello", "\n\t"}
    v2 := make(map[int]string)
    var v3 int8 = -103
    var v4 uint32 = 723
    var v5 int16 = 0
    var v11 int32 = 2147483647
    v5 = v5
    v3 = genFn0((v5 + (^(v5))), genFn1(32767))
    if (0 <= 4294967295) {
        var v12 uint8 = 255
        _ = v12
    } else {
        var v13 uint64 = 42
        _ = v13
    }
    v5 += v5
    v14 := GenS3{F0: true, F1: -164}
    fmt.Println(v4, v4, v4)
    if (v11 > 2147483647) {
        if (true || (!(true || ((-2.25 / 1.0) <= 3.14159)))) {
            for fi15 := 0; fi15 < 6; fi15++ {
                if (true && (1 >= 0)) {
                    v16 := GenS1{F0: 1, F1: (v5 + (^(v5))), F2: 42}
                    _ = v16
                } else {
                    var v17 float64 = 1e10
                    var v18 int64 = -505
                    v5 = (^(v5))
                    _ = v17
                    _ = v18
                }
                v1 = append(v1, "")
                v3 -= (v3 * v3)
                _ = fi15
            }
        } else {
            for fi19 := 0; fi19 < 17; fi19++ {
                for fi20 := 0; fi20 < 2; fi20++ {
                    fmt.Println(v11, fi19, fi19)
                    _ = fi20
                }
                v21 := GenS2{F0: 645, F1: 42}
                var v22 uint16 = 65535
                _ = fi19
                _ = v21
                _ = v22
            }
            v11 -= 2147483647
            v11 = v11
        }
    }
    if true {
        fmt.Println(v5, v5, v4)
        switch 127 {
        case 484:
            if false {
                var v23 float64 = 3.14159
                var v24 float32 = ((3.14159 - 0.0) / 0.1)
                _ = v23
                _ = v24
            }
        }
    } else {
        var v25 string = "foo bar"
        _ = v25
    }
    v2[895] = ("\n\t" + "\n\t")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
