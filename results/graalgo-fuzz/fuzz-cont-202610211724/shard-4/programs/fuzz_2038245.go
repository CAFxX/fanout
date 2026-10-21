package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
func genFn0(p0 uint16) uint64 {
    var v7 int = 128
    v8 := GenS1{F0: v7}
    fmt.Println(v7, v7)
    var v9 uint = 18446744073709551615
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := []int{(-(v7)), v7, fi10}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 uint64, p1 bool) int8 {
    v12 := []int{-1, 9223372036854775807}
    var v13 uint64 = p0
    var v14 uint64 = (p0 << ((42) & 63))
    var v15 int = 2147483647
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 49
}

func main() {
    var v1 uint64 = 697
    var v2 float64 = 3.14159
    var v3 float64 = -2.25
    v4 := []int{1}
    v5 := []uint32{}
    v6 := []int{-9223372036854775808, 32768, 96, -601}
    var v16 uint = 255
    if true {
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 float64 = v3
            if (!((-128 > 19) || (2147483647 < -1))) {
                v1 -= 4294967295
                fmt.Println(v2, fi17, v1)
                if ((v18 / -2.25) >= 1e10) {
                    var v19 int64 = 1
                    v20 := GenS2{F0: ("foo bar" + (("a" + "a") + ("foo bar" + ("a" + "foo bar"))))}
                    _ = v19
                    _ = v20
                } else {
                    switch "\n\t" {
                    case "":
                        if len(v5) > 0 {
                            v5[0] = 42
                        }
                    case "hello":
                        v3 -= v2
                    }
                }
            }
            _ = fi17
            _ = v18
        }
        if (true || true) {
            var v21 int8 = 0
            var v22 int32 = 1
            _ = v21
            _ = v22
        }
        v4 = append(v4, 227)
    }
    var v23 string = "hello"
    v5 = append(v5, 4294967295)
    v24 := GenS1{F0: 1}
    v25 := GenS1{F0: -2147483648}
    if true {
        var v26 float32 = 0.0
        v23 = v23
        v27 := GenS2{F0: "a"}
        _ = v26
        _ = v27
    }
    switch "foo bar" {
    case "foo bar":
        v1 = v1
    case "\n\t":
        var v28 uint16 = 536
        _ = v28
    }
    var v29 int = 32768
    if false {
        if ((1 > 100) && (42 < (v1 + uint64(v16)))) {
            v2 = (((3.14159 / (1.5 + (v3 - v3))) * 1.5) + v2)
        }
    }
    v30 := GenS2{F0: (v23 + v23)}
    if true {
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 int = -365
            var v33 uint8 = 175
            var v34 float64 = 3.14159
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        v3 -= v2
        fmt.Println(v16, v29, v1)
    } else {
        _ = len(v5)
    }
    var v35 float64 = float64(3.14159)
    if false {
        v3 += ((-2.25 - v3) * (float64(0.0) * ((3.14159 + v35) + (0.0 - ((1.5 / (v2 / v35)) - float64((((-2.25 * (float32(v3) * -2.25)) / 1.0) * 3.14159)))))))
        var v36 uint8 = 42
        _ = v36
    } else {
        var v37 int64 = 1
        _ = v37
    }
    v4 = append(v4, 10)
    fmt.Println(v3, v35, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
