package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint8, p1 uint64) int32 {
    var v4 float32 = (-2.25 * -2.25)
    var v5 int16 = 42
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -2147483648
}

func main() {
    var v1 float32 = 3.14159
    v2 := []int{128, 9223372036854775807, -1}
    var v3 uint8 = 127
    v2 = append(v2, 326)
    for fi6 := 0; fi6 < 1; fi6++ {
        v1 = v1
        v3 += (v3 ^ v3)
        _ = fi6
    }
    var v7 uint = 42
    var v8 uint16 = 32767
    var v9 uint8 = v3
    for fi10 := 0; fi10 < 13; fi10++ {
        switch int(v7) {
        case 32768:
            fmt.Println(fi10, v9, v3)
        default:
            if true {
                var v11 int16 = -613
                var v12 int = fi10
                for fi13 := 0; fi13 < 17; fi13++ {
                    var v14 uint = v7
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch (("" + ((("" + ("" + "foo bar")) + "hello") + "a")) + "\n\t") {
                    case "hello":
                        v9 = (^(v3))
                    case "\n\t":
                        fmt.Println(v7, fi13, v14)
                    case "a":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        v7 = (^(v7))
                    }
                    v1 -= v1
                    _ = fi13
                    _ = v14
                }
                _ = v11
                _ = v12
            }
        }
        _ = fi10
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        fmt.Println(v9)
        if (true && false) {
            v8 = v8
        }
        _ = fi15
    }
    switch 9223372036854775807 {
    case 0:
        v3 = v3
    case -2147483648:
        var v16 int8 = 127
        _ = v16
    case 9223372036854775807:
        var v17 string = ("\n\t" + "hello")
        _ = v17
    default:
        var v18 int64 = -185
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        fmt.Println(v9, v3)
        _ = fi19
    }
    switch 32768 {
    case 32768:
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := GenS1{F0: -111}
            _ = fi20
            _ = v21
        }
    case 4294967295:
        if false {
            var v22 uint32 = 2147483647
            if (true || true) {
                for fi23 := 0; fi23 < 11; fi23++ {
                    v24 := GenS1{F0: -128}
                    _ = fi23
                    _ = v24
                }
                fmt.Println(v1)
            }
            if (!true) {
                for fi25 := 0; fi25 < 18; fi25++ {
                    v26 := GenS1{F0: 127}
                    var v27 int = 42
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                v1 -= float32(0.1)
                if len(v2) > 0 {
                    v2[0] = 706
                }
            }
            _ = v22
        }
    case 32767:
        v2 = append(v2, -9223372036854775808)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!true) {
        v9 -= 244
        _ = len(v2)
    }
    if true {
        v8 = v8
        if (!(!true)) {
            for fi28 := 0; fi28 < 19; fi28++ {
                fmt.Println(v8, fi28)
                fmt.Println(v7)
                v9 += (v9 | (v3 & v9))
                _ = fi28
            }
            var v29 float64 = 1.5
            _ = v29
        }
        v1 -= ((v1 / ((float32((-2.25 / 1.0)) - float32(1.5)) * v1)) - v1)
    }
    fmt.Println(v9, v9)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
