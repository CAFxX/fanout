package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int16; F1 uint8 }
func genFn0() int64 {
    var v7 int32 = 2147483647
    if (-1 >= 127) {
        if ((2 > 2) || false) {
            v7 += v7
            for fi8 := 0; fi8 < 18; fi8++ {
                if (1 < 0) {
                    var v9 uint16 = 532
                    _ = v9
                } else {
                    var v10 uint16 = 1
                    _ = v10
                }
                if (!false) {
                    v11 := []int{-559}
                    _ = v11
                }
                _ = fi8
            }
            var v12 int = 42
            _ = v12
        }
    } else {
        if (114 == 4294967295) {
            var v13 bool = false
            if (v13 && v13) {
                for fi14 := 0; fi14 < 11; fi14++ {
                    fmt.Println(fi14)
                    var v15 uint8 = 128
                    v16 := GenS1{F0: ""}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            fmt.Println(v13)
            _ = v13
        }
    }
    v17 := []int{-1}
    _ = v17
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 uint32, p1 int32) string {
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 uint16 = 255
        p1 -= (^(p1))
        _ = fi18
        _ = v19
    }
    switch 32767 {
    case -290:
        if (0 == 772) {
            v20 := GenS2{F0: 42}
            _ = v20
        } else {
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 uint = 18446744073709551615
                switch ("\n\t" + "") {
                case "a":
                    var v23 uint16 = 32767
                    _ = v23
                default:
                    if true {
                        v24 := []int{fi21}
                        _ = v24
                    }
                }
                p1 -= p1
                _ = fi21
                _ = v22
            }
            v25 := GenS2{F0: 1}
            v26 := GenS3{F0: -1, F1: 9223372036854775807}
            _ = v25
            _ = v26
        }
    default:
        p1 = -2147483648
    }
    _ = p0
    _ = p1
    return (("hello" + ("" + "hello")) + ("" + "hello"))
}

func main() {
    var v1 float64 = (((float64(0.0) * float64((3.14159 - -2.25))) + float64(1.5)) + 3.14159)
    var v2 uint8 = 128
    v3 := []int{0}
    var v4 uint64 = 796
    v5 := make(map[int]string)
    v6 := make(map[int]uint32)
    fmt.Println(v4)
    fmt.Println(v4)
    var v27 uint64 = (v4 >> ((864) & 63))
    delete(v5, 4294967295)
    v2 -= v2
    switch (("foo bar" + "hello") + ("foo bar" + ("hello" + "a"))) {
    case "hello":
        delete(v5, int(v27))
    }
    switch -2147483648 {
    case 127:
        fmt.Println(v27, v4, v2)
    }
    if false {
        if (((-42 <= 107) && (!true)) || false) {
            var v28 float64 = float64(((0.0 / 1.5) / 1.5))
            _ = v28
        } else {
            v1 -= 0.1
        }
    } else {
        v6[-2147483648] = 2147483647
        v29 := GenS2{F0: 1}
        _ = v29
    }
    var v30 bool = (!(1000 == -954))
    var v31 int = -9223372036854775808
    switch (v31 >> ((2) & 63)) {
    case 2147483647:
        v32 := v5[123]
        _ = v32
    case 9223372036854775807:
        v33 := GenS4{F0: 1000, F1: v2}
        _ = v33
    default:
        _ = len(v6)
    }
    if len(v3) > 0 {
        v3[0] = (v31 / (((v31 - (^(v31)))) | 1))
    }
    v31 = v31
    for fi34 := 0; fi34 < 6; fi34++ {
        fmt.Println(v27, v30)
        _ = fi34
    }
    v30 = (0.1 <= v1)
    for fi35 := 0; fi35 < 5; fi35++ {
        if (42 == 65535) {
            for fi36 := 0; fi36 < 7; fi36++ {
                switch "hello" {
                case "a":
                    v37 := GenS2{F0: genFn0()}
                    _ = v37
                case "foo bar":
                    var v38 uint64 = (v4 / ((2) | 1))
                    _ = v38
                }
                switch "" {
                case "hello":
                    v4 = v27
                }
                _ = fi36
            }
            var v39 int8 = 42
            _ = v39
        } else {
            var v40 uint16 = 65535
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v40
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
