package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 uint64, p1 uint16, p2 bool) int {
    var v7 uint32 = 65535
    for fi8 := 0; fi8 < 1; fi8++ {
        p0 += (p0 << ((255) & 63))
        var v9 string = "\n\t"
        if true {
            if p2 {
                if (!(1.5 != 3.14159)) {
                    var v10 int = (fi8 - (fi8 << ((uint(p0)) & 63)))
                    if false {
                        if ((0.1 < (1e10 * (float32(1.5) + (-2.25 + 1e10)))) || p2) {
                            v9 = "foo bar"
                        }
                    } else {
                        fmt.Println(fi8)
                        switch (v10 << ((647) & 63)) {
                        case -9223372036854775808:
                            p0 = (p0 + (p0 - p0))
                        case 2147483647:
                            if (v9 > v9) {
                                var v11 float64 = 3.14159
                                _ = v11
                            } else {
                                fmt.Println(v10, v9)
                            }
                        }
                    }
                    _ = v10
                }
                p2 = (uint32(p1) > (v7 / ((v7) | 1)))
            } else {
                fmt.Println(p1, p0)
                v7 += v7
                fmt.Println(fi8, v7)
            }
            fmt.Println(p0, p0)
            if p2 {
                v12 := []int{(fi8 + fi8), fi8, fi8}
                p1 += 255
                _ = v12
            } else {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v9 = v9
                    var v14 uint8 = uint8(p1)
                    _ = fi13
                    _ = v14
                }
            }
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2147483647
}
func genFn1() uint {
    var v15 string = "\n\t"
    var v16 uint16 = 439
    _ = v15
    _ = v16
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{1, 2, 131}
    v3 := []uint32{65535, 65535, 42}
    v4 := []int{2147483647, 32768}
    var v5 float64 = -2.25
    v6 := make(map[string]uint32)
    v17 := GenS1{F0: 527}
    fmt.Println(v5)
    v18 := GenS1{F0: -41}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v19 uint64 = 4294967295
    switch "foo bar" {
    case "foo bar":
        v5 -= 3.14159
    case "a":
        fmt.Println(v19)
    default:
        for fi20 := 0; fi20 < 19; fi20++ {
            v19 = (v19 ^ (v19 >> ((255) & 63)))
            _ = fi20
        }
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        v19 = (v19 ^ (v19 / ((111) | 1)))
        v1[227] = true
        v22 := v6["\n\t"]
        _ = fi21
        _ = v22
    }
    switch "\n\t" {
    case "foo bar":
        if false {
            var v23 uint64 = (v19 - (v19 + v19))
            _ = v23
        } else {
            var v24 int32 = 0
            v25 := v1[42]
            _ = v24
            _ = v25
        }
    case "a":
        v6[""] = 933
    case "hello":
        switch (("a" + ("" + ("\n\t" + ""))) + "hello") {
        case "foo bar":
            var v26 int32 = 1
            _ = v26
        }
    }
    _ = len(v6)
    var v27 uint32 = 2147483648
    for fi28 := 0; fi28 < 18; fi28++ {
        if true {
            if (((v19 - v19) < (^(v19))) || true) {
                fmt.Println(v19)
            } else {
                if false {
                    fmt.Println(v27, v5)
                }
            }
            v27 -= 65535
        } else {
            var v29 bool = false
            v30 := GenS1{F0: 32767}
            _ = v29
            _ = v30
        }
        _ = fi28
    }
    var v31 float64 = v5
    for fi32 := 0; fi32 < 14; fi32++ {
        for fi33 := 0; fi33 < 9; fi33++ {
            v34 := GenS2{F0: 0, F1: 707}
            _ = len(v2)
            if false {
                v35 := GenS2{F0: -1, F1: -1}
                var v36 float64 = float64(3.14159)
                _ = v35
                _ = v36
            }
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
