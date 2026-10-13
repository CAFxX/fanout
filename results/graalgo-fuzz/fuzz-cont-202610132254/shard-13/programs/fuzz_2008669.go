package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }
type GenS3 struct { F0 int32; F1 int8 }
func genFn0() int8 {
    v7 := GenS1{F0: 0}
    v8 := GenS1{F0: 42}
    var v9 uint16 = 937
    fmt.Println(v9)
    v10 := []int{}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 uint32, p1 uint16, p2 uint64) string {
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(p1)
        p0 -= (p0 >> ((4294967295) & 63))
        if (((1e10 - float32(1e10)) / -2.25) != 1.5) {
            switch (^(fi11)) {
            case 42:
                var v12 uint32 = p0
                _ = v12
            case 32767:
                p0 += p0
            case -1:
                if (!true) {
                    fmt.Println(p0)
                    for fi13 := 0; fi13 < 7; fi13++ {
                        fmt.Println(p0, fi11, fi11)
                        _ = fi13
                    }
                }
            default:
                for fi14 := 0; fi14 < 9; fi14++ {
                    var v15 int = (fi11 ^ int(p2))
                    _ = fi14
                    _ = v15
                }
            }
        }
        _ = fi11
    }
    switch 9223372036854775807 {
    case -419:
        if false {
            var v16 uint = 65535
            p2 = p2
            _ = v16
        }
    case 999:
        v17 := []int{}
        _ = v17
    }
    var v18 uint = 42
    v19 := GenS2{F0: -9223372036854775808, F1: 42, F2: 1}
    if false {
        if true {
            p1 = p1
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 float32 = 0.0
                var v22 int8 = -1
                _ = fi20
                _ = v21
                _ = v22
            }
            var v23 uint = (v18 << ((v18) & 63))
            _ = v23
        } else {
            var v24 int32 = 2147483647
            switch -2147483648 {
            case -221:
                fmt.Println(p1, p0, v24)
            case 128:
                if false {
                    if true {
                        var v25 int16 = 42
                        var v26 int64 = 317
                        _ = v25
                        _ = v26
                    }
                    p2 = (p2 - (p2 - (p2 % ((p2) | 1))))
                } else {
                    var v27 uint = 65535
                    var v28 uint8 = uint8(p1)
                    _ = v27
                    _ = v28
                }
            default:
                for fi29 := 0; fi29 < 16; fi29++ {
                    var v30 float32 = (1e10 * float32((-2.25 / -2.25)))
                    _ = fi29
                    _ = v30
                }
            }
            var v31 bool = (v18 <= v18)
            _ = v24
            _ = v31
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return "a"
}

func main() {
    var v1 uint32 = 42
    var v2 float32 = (1.5 + float32((float64(1.5) / 1.0)))
    v3 := []uint32{42, v1, v1}
    v4 := []string{("a" + "a"), "\n\t", "hello"}
    v5 := []int{-9223372036854775808, -1, 12}
    var v6 float64 = (1e10 / 3.14159)
    v32 := GenS1{F0: 128}
    v33 := GenS2{F0: -990, F1: 911, F2: 9223372036854775807}
    var v34 int = 42
    v34 += v34
    switch v34 {
    case 32767:
        var v35 bool = (!(92 >= 1))
        _ = v35
    default:
        v1 -= (^(v1))
    }
    var v36 uint8 = 2
    var v37 uint = 255
    var v38 string = "foo bar"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v39 uint64 = 984
    v40 := GenS2{F0: 9223372036854775807, F1: (v39 ^ 1), F2: 0}
    if (v38 == v38) {
        fmt.Println(v36, v36)
    }
    if true {
        var v41 string = (v38 + v38)
        if true {
            v41 = ((v38 + v38) + ("\n\t" + (v38 + "foo bar")))
            v36 = 2
            v42 := GenS3{F0: -2147483648, F1: -128}
            _ = v42
        } else {
            var v43 int8 = 1
            _ = v43
        }
        _ = v41
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        var v45 int8 = 0
        if len(v3) > 0 {
            v3[0] = 2
        }
        var v46 uint64 = v39
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
