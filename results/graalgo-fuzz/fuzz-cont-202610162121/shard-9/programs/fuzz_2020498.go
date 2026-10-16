package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() string {
    switch (("hello" + "") + ((("a" + ((("a" + "\n\t") + "\n\t") + ("\n\t" + "hello"))) + ("" + "foo bar")) + "foo bar")) {
    case "a":
        var v4 bool = (3.14159 <= (((((-2.25 + -2.25) + float32(float64((float32(-2.25) / -2.25)))) / -2.25) / (0.0 + 1.5)) + (-2.25 + 3.14159)))
        _ = v4
    case "foo bar":
        var v5 float64 = (((float64((-2.25 + 0.1)) * 1e10) + 0.1) * 1.5)
        _ = v5
    }
    var v6 float32 = 3.14159
    switch -2147483648 {
    case 217:
        var v7 int32 = 1
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{fi8, fi8}
        var v10 int32 = 1
        v6 -= (v6 + 1.5)
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{-9223372036854775808, 1}
    _ = v11
    _ = v6
    return "hello"
}
func genFn1(p0 uint8) uint32 {
    if ((!true) || (3.14159 != 3.14159)) {
        p0 -= 128
        switch 4294967295 {
        case 90:
            p0 -= (p0 - p0)
        default:
            v12 := []int{32767}
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS1{F0: -1}
            var v15 float64 = 0.0
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(p0)
        var v16 uint64 = 4294967295
        _ = v16
    }
    var v17 string = ((("foo bar" + ("\n\t" + "a")) + "\n\t") + "\n\t")
    p0 += (^(p0))
    _ = p0
    _ = v17
    return 65535
}

func main() {
    v1 := []int{}
    var v2 int64 = 1
    v3 := make(map[string]bool)
    var v18 float64 = 3.14159
    fmt.Println(v2, v18)
    var v19 int8 = int8(v2)
    switch 32767 {
    case 99:
        var v20 float32 = 1.5
        _ = v20
    case -688:
        var v21 int8 = -1
        _ = v21
    case 32768:
        var v22 uint32 = 4294967295
        _ = v22
    }
    var v23 uint32 = 2
    v2 += v2
    if (true && true) {
        v18 -= (v18 + v18)
        switch 9223372036854775807 {
        case -2147483648:
            fmt.Println(v19, v19, v2)
        }
    } else {
        switch 9223372036854775807 {
        case 2147483647:
            var v24 uint32 = (^(v23))
            _ = v24
        case 127:
            var v25 uint = 300
            _ = v25
        case 0:
            var v26 uint = 255
            _ = v26
        }
        v18 -= v18
        v23 = v23
    }
    v27 := GenS1{F0: (v19 % ((v19) | 1))}
    delete(v3, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v27
    _ = v3
}
