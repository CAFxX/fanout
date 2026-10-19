package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
func genFn0(p0 uint, p1 int16, p2 int64) bool {
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := []int{fi6, fi6}
        if len(v7) > 0 {
            v7[0] = fi6
        }
        if (fi6 != 973) {
            var v8 uint64 = 2
            if true {
                v9 := GenS1{F0: (p2 >> (((p0 << ((p0) & 63))) & 63)), F1: 0, F2: 1}
                if ((127 <= 161) && true) {
                    v10 := GenS1{F0: (^(p2)), F1: 1, F2: 746}
                    var v11 int16 = p1
                    _ = len(v7)
                    _ = v10
                    _ = v11
                }
                v12 := GenS1{F0: 1, F1: 896, F2: 65535}
                _ = v12
                _ = v9
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 int64 = p2
                if len(v7) > 0 {
                    v7[0] = (-(fi13))
                }
                switch fi6 {
                case 0:
                    switch (("foo bar" + "a") + ("hello" + "a")) {
                    case "\n\t":
                        var v15 int64 = 1
                        _ = v15
                    case "":
                        v16 := GenS2{F0: 32767}
                        _ = v16
                    default:
                        fmt.Println(p0)
                    }
                case 66:
                    var v17 float32 = (0.1 + 3.14159)
                    _ = v17
                default:
                    var v18 int = 2147483647
                    _ = v18
                }
                _ = fi13
                _ = v14
            }
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    if true {
        for fi19 := 0; fi19 < 11; fi19++ {
            v20 := []int{fi19, 2147483647, (fi19 / ((fi19) | 1)), (fi19 * (fi19 & (^(fi19))))}
            _ = fi19
            _ = v20
        }
        var v21 int64 = 42
        switch "\n\t" {
        case "hello":
            for fi22 := 0; fi22 < 15; fi22++ {
                fmt.Println(v21, p1)
                _ = fi22
            }
        case "\n\t":
            v23 := GenS2{F0: 126}
            _ = v23
        }
        _ = v21
    }
    v24 := GenS3{F0: "\n\t"}
    for fi25 := 0; fi25 < 15; fi25++ {
        p0 = (p0 << ((p0) & 63))
        _ = fi25
    }
    var v26 uint8 = 181
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v26
    return (0.0 != -2.25)
}
func genFn1() int64 {
    for fi27 := 0; fi27 < 11; fi27++ {
        fmt.Println(fi27)
        _ = fi27
    }
    var v28 float64 = (0.1 - float64((1e10 + -2.25)))
    _ = v28
    return -1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    var v3 int16 = -344
    var v4 bool = (!(255 > 1))
    var v5 int64 = 9223372036854775807
    v29 := GenS2{F0: 1}
    v5 += (^(v5))
    v3 -= 1
    v3 -= v3
    v30 := GenS3{F0: "\n\t"}
    fmt.Println(v4, v5)
    if (!false) {
        var v31 uint16 = 65535
        _ = v31
    }
    if (284 != 154) {
        _ = len(v2)
        var v32 float32 = 1.5
        delete(v2, 32768)
        _ = v32
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        var v34 float32 = 3.14159
        delete(v1, "\n\t")
        switch (fi33 * (^(fi33))) {
        case 1:
            var v35 float64 = 1e10
            _ = v35
        }
        _ = fi33
        _ = v34
    }
    v36 := []int{127, -9223372036854775808, -462, 127}
    v2[776] = (true || v4)
    v5 -= v5
    if (("hello" != (("foo bar" + "a") + "foo bar")) || (1 >= 42)) {
        var v37 uint32 = 2147483647
        fmt.Println(v3, v5, v5)
        v38 := GenS3{F0: "foo bar"}
        _ = v37
        _ = v38
    } else {
        var v39 uint16 = 1
        v4 = v4
        _ = v39
    }
    v36 = append(v36, int(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
}
