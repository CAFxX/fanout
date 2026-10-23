package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint16) uint32 {
    p0 += p0
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint16 = 0
        for fi9 := 0; fi9 < 19; fi9++ {
            p0 = (^(v8))
            v10 := GenS1{F0: 220}
            v8 -= 255
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    return 2147483648
}
func genFn1() string {
    var v11 uint32 = 42
    var v12 uint32 = (^(v11))
    v11 = (v12 * genFn0(42))
    if true {
        v13 := GenS1{F0: 42}
        _ = v13
    } else {
        v14 := []int{127}
        _ = v14
    }
    _ = v11
    _ = v12
    return "hello"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"", "foo bar"}
    v3 := []uint32{1, 2, 590}
    v4 := make(map[int]string)
    var v5 float32 = 0.0
    v6 := []uint32{78, 65535, 380}
    switch "\n\t" {
    case "\n\t":
        v5 = ((v5 / v5) - v5)
    }
    var v15 uint32 = genFn0(32767)
    fmt.Println(v15)
    v16 := v1[(v15 / ((v15) | 1))]
    v15 = (v15 >> ((840) & 63))
    var v17 int16 = 691
    v5 = float32(1.5)
    if len(v3) > 0 {
        v3[0] = (^(v15))
    }
    var v18 float32 = (-2.25 / (((float32(1e10) + v5) * -2.25) - 1.5))
    for fi19 := 0; fi19 < 1; fi19++ {
        v20 := GenS1{F0: 1}
        fmt.Println(v17, v17)
        var v21 int = -2147483648
        _ = fi19
        _ = v20
        _ = v21
    }
    switch -1 {
    case 1:
        for fi22 := 0; fi22 < 11; fi22++ {
            v23 := GenS1{F0: 1}
            _ = fi22
            _ = v23
        }
    case 0:
        switch 4294967295 {
        case 0:
            v1[v15] = v16
        }
    }
    v17 = (-(v17))
    var v24 int64 = 42
    switch -180 {
    case 9223372036854775807:
        switch v16 {
        case "foo bar":
            if false {
                var v25 int64 = (v24 | 0)
                if ((!(4294967295 == 736)) && true) {
                    switch 4294967295 {
                    case -9223372036854775808:
                        fmt.Println(v25)
                    case 4294967295:
                        var v26 uint8 = 195
                        _ = v26
                    case 9223372036854775807:
                        switch 312 {
                        case 127:
                            fmt.Println(v24, v25)
                        }
                    }
                    var v27 uint8 = 128
                    _ = v27
                } else {
                    v16 = ("\n\t" + ((v16 + (v16 + "\n\t")) + v16))
                }
                v24 = (v24 | (v24 ^ (^(v24))))
                _ = v25
            } else {
                v17 -= 32767
                v18 = v18
            }
        case "a":
            var v28 float32 = ((v5 + v18) + 0.0)
            _ = v28
        }
    default:
        fmt.Println(v17, v5, v18)
    }
    var v29 int = 4294967295
    fmt.Println(v29, v5)
    switch "" {
    case "foo bar":
        if (v5 >= (-2.25 * v18)) {
            v30 := GenS2{F0: 1}
            _ = v30
        }
    case "hello":
        fmt.Println(v17, v15)
    case "":
        for fi31 := 0; fi31 < 4; fi31++ {
            v1[v15] = (v16 + v16)
            for fi32 := 0; fi32 < 4; fi32++ {
                var v33 int64 = v24
                v24 = v33
                if true {
                    var v34 uint16 = 32767
                    v16 = ("foo bar" + "a")
                    _ = v34
                } else {
                    fmt.Println(fi31, v24, v18)
                }
                _ = fi32
                _ = v33
            }
            _ = fi31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
