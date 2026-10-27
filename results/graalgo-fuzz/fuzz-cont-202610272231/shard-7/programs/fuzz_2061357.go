package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int; F1 int64; F2 int }
type GenS3 struct { F0 uint32; F1 int16; F2 int64 }
func genFn0(p0 bool, p1 uint) int8 {
    p0 = ((42 == 4294967295) && (4294967295 <= 1))
    for fi6 := 0; fi6 < 4; fi6++ {
        if p0 {
            var v7 uint64 = 4294967295
            _ = v7
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 uint8, p1 uint8) int8 {
    v8 := GenS1{F0: 4294967295, F1: 2}
    var v9 bool = (true && false)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -715
    var v3 int = -2147483648
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    v10 := GenS1{F0: 1, F1: 2}
    v3 += v3
    v11 := GenS1{F0: 49, F1: 128}
    if ((0 > 9223372036854775807) && (float64(float32(0.1)) < (0.1 + 1.5))) {
        switch 621 {
        case 32767:
            _ = len(v4)
        }
    }
    var v12 int32 = (v2 - (v2 | v2))
    if ((32767 != 0) && true) {
        fmt.Println(v2, v2, v2)
        switch v3 {
        case 4294967295:
            v2 -= (v2 >> ((0) & 63))
        }
    }
    fmt.Println(v12, v3)
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            if false {
                v12 += 1
            } else {
                fmt.Println(v2, v2)
                v3 -= fi14
            }
            _ = fi14
        }
        var v15 int16 = 1000
        _ = fi13
        _ = v15
    }
    v12 -= (-(v2))
    fmt.Println(v2, v2, v2)
    _ = len(v1)
    switch (v3 << ((720) & 63)) {
    case 1:
        if false {
            v1["hello"] = ("" <= "a")
        } else {
            v12 += (v2 + (-(v12)))
        }
    case 9223372036854775807:
        if (false && (94 != v3)) {
            v12 = v2
        } else {
            var v16 float32 = -2.25
            var v17 bool = (0 == 96)
            _ = v16
            _ = v17
        }
    case 127:
        if ((255 == 0) && (!true)) {
            for fi18 := 0; fi18 < 17; fi18++ {
                for fi19 := 0; fi19 < 17; fi19++ {
                    v20 := GenS1{F0: 4294967295, F1: 42}
                    var v21 string = ("foo bar" + "")
                    if (!(767 > 255)) {
                        switch "foo bar" {
                        case "a":
                            v1["hello"] = (!(-655 < -9223372036854775808))
                        case "\n\t":
                            var v22 string = v21
                            _ = v22
                        }
                        v3 -= fi18
                        switch 9223372036854775807 {
                        case 42:
                            fmt.Println(v12, v3)
                        case 128:
                            v23 := []int{-2147483648, (fi19 ^ 34), (v3 % ((fi19) | 1)), -749}
                            _ = v23
                        case -2147483648:
                            var v24 int16 = -1
                            _ = v24
                        default:
                            var v25 uint64 = 2
                            v26 := GenS2{F0: 127, F1: 529, F2: (-(fi18))}
                            var v27 uint8 = 127
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    }
                    var v28 string = v21
                    v21 = v21
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v28
                }
                v3 -= v3
                _ = fi18
            }
            v29 := []int{(v3 % ((-9223372036854775808) | 1)), 835, -427, -9223372036854775808}
            if (v2 > (v2 << ((255) & 63))) {
                if (42 < -9223372036854775808) {
                    delete(v1, ("a" + "a"))
                    if false {
                        if (888 < 59) {
                            v2 -= 0
                            v30 := GenS1{F0: 4294967295, F1: 237}
                            _ = v30
                        }
                        var v31 uint8 = 128
                        fmt.Println(v2, v31)
                        _ = v31
                    }
                    var v32 uint = 4294967295
                    _ = v32
                }
                var v33 uint = 42
                _ = v33
            }
            _ = v29
        }
    }
    var v34 uint16 = 65535
    v35 := GenS3{F0: 418, F1: -1, F2: -279}
    v36 := []int{(v3 - (^(v3))), v3, v3, v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
