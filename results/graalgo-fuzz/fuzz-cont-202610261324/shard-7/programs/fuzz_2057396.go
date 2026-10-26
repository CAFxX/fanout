package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint, p1 int8) int16 {
    v5 := GenS1{F0: 0}
    p0 = p0
    p0 += p0
    var v6 uint = p0
    v7 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 97
}

func main() {
    v1 := []string{("a" + (("hello" + ("hello" + "hello")) + (("foo bar" + "a") + "a")))}
    var v2 bool = (4294967295 != 1)
    var v3 bool = (9223372036854775807 > -804)
    var v4 string = "foo bar"
    var v8 int = -507
    switch v8 {
    case 0:
        v8 -= v8
    case 9223372036854775807:
        var v9 uint8 = 255
        _ = v9
    default:
        v10 := GenS2{F0: -2147483648}
        _ = v10
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v11 := GenS1{F0: 91}
    var v12 float64 = 0.1
    v13 := GenS3{F0: 2147483647, F1: ((float32(v12) * float32(((3.14159 - v12) * v12))) <= 3.14159)}
    var v14 int8 = 0
    switch (((v4 + "") + (v4 + ("\n\t" + "foo bar"))) + (v4 + ("foo bar" + "foo bar"))) {
    case "":
        v15 := GenS1{F0: 41}
        _ = v15
    case "foo bar":
        v14 = (v14 * v14)
    case "hello":
        var v16 int32 = 892
        _ = v16
    }
    fmt.Println(v8, v12, v12)
    v1 = append(v1, v4)
    var v17 int64 = 0
    fmt.Println(v17)
    if (!(v4 > v4)) {
        _ = len(v1)
        _ = len(v1)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (((v12 / -2.25) == -2.25) || ((1 != 1) && v3)) {
        var v18 int16 = genFn0(255, (v14 & (v14 / ((100) | 1))))
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 uint64 = 18446744073709551615
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    var v21 uint16 = 1
    switch "a" {
    case "hello":
        v14 = (v14 * int8(v17))
    case "\n\t":
        if (2147483647 > -2147483648) {
            v1 = append(v1, (v4 + "foo bar"))
            if (127 >= 197) {
                if (!false) {
                    var v22 int16 = 676
                    var v23 int8 = v14
                    _ = v22
                    _ = v23
                } else {
                    v3 = (!false)
                }
                v8 += v8
            } else {
                v2 = true
                fmt.Println(v14, v2)
            }
            v24 := GenS4{F0: 2, F1: (v21 + (v21 | 65535))}
            _ = v24
        }
    case "a":
        if v2 {
            var v25 float64 = (((v12 - v12) / 1.0) + 1.5)
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 uint16 = v21
                for fi28 := 0; fi28 < 18; fi28++ {
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    _ = fi28
                }
                var v29 uint32 = 986
                _ = fi26
                _ = v27
                _ = v29
            }
            if false {
                var v30 float64 = v25
                _ = v30
            } else {
                v31 := GenS5{F0: 4294967295}
                var v32 int8 = (v14 >> ((4294967295) & 63))
                _ = v31
                _ = v32
            }
            _ = v25
        } else {
            var v33 int16 = -28
            for fi34 := 0; fi34 < 16; fi34++ {
                var v35 string = "hello"
                _ = fi34
                _ = v35
            }
            _ = v33
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v8
}
