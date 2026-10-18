package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint16; F1 int }
func genFn0(p0 uint, p1 uint32) uint64 {
    v6 := []int{32768, 0, 2147483647, -817}
    if true {
        var v7 uint16 = 65535
        p1 += p1
        var v8 string = ""
        _ = v7
        _ = v8
    }
    if false {
        if len(v6) > 0 {
            v6[0] = -513
        }
        p0 -= p0
        var v9 bool = true
        _ = v9
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    if false {
        var v10 string = ""
        switch (v10 + (v10 + (v10 + ((("\n\t" + (((v10 + v10) + (("hello" + v10) + v10)) + (v10 + "\n\t"))) + "") + v10)))) {
        case "\n\t":
            for fi11 := 0; fi11 < 17; fi11++ {
                fmt.Println(p0)
                var v12 int32 = -335
                _ = fi11
                _ = v12
            }
        default:
            v6 = append(v6, 663)
        }
        if (p1 >= p1) {
            for fi13 := 0; fi13 < 18; fi13++ {
                for fi14 := 0; fi14 < 18; fi14++ {
                    if len(v6) > 0 {
                        v6[0] = fi13
                    }
                    v15 := GenS1{F0: 1, F1: 1}
                    p1 = p1
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            p1 -= 53
            var v16 int16 = 526
            _ = v16
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    return 4294967295
}
func genFn1(p0 uint8) int8 {
    v17 := []int{32768}
    var v18 int32 = 0
    var v19 uint32 = 4294967295
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = -2147483648
    v3 := make(map[string]string)
    var v4 bool = ((1.5 + (((0.1 * 3.14159) - (0.1 - float64(3.14159))) + -2.25)) > (float64((1e10 * 3.14159)) - ((1e10 / 1.0) + 0.1)))
    v5 := []uint32{1, 65535, 49}
    var v20 float64 = (((1e10 / 1e10) + (float64((1e10 * 0.0)) / ((float64(1e10) - float64(float32(((1e10 * 0.0) + 1.5)))) * -2.25))) / 1e10)
    v4 = (2147483647 >= -2147483648)
    var v21 uint = 2
    var v22 int64 = 0
    var v23 float64 = v20
    var v24 string = "a"
    var v25 uint = v21
    for fi26 := 0; fi26 < 16; fi26++ {
        delete(v1, 1)
        var v27 int16 = 0
        switch v24 {
        case "hello":
            switch (v24 + ("" + "foo bar")) {
            case "foo bar":
                v28 := GenS2{F0: 42, F1: (^(fi26))}
                _ = v28
            case "a":
                v22 = v22
            case "":
                delete(v3, ("" + (v24 + (v24 + v24))))
            }
        }
        _ = fi26
        _ = v27
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
