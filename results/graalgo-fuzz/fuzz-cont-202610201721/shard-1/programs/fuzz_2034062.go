package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint64; F2 uint16 }
func genFn0(p0 uint32, p1 uint16) int16 {
    p1 = 65535
    p0 -= p0
    for fi6 := 0; fi6 < 13; fi6++ {
        p1 -= p1
        _ = fi6
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint32 = 375
    var v2 int16 = 454
    var v3 bool = (!true)
    var v4 int8 = -1
    var v5 float64 = 1e10
    var v7 int8 = v4
    v8 := GenS1{F0: v1}
    var v9 string = ((("hello" + (("hello" + ("a" + "")) + "a")) + ("" + "\n\t")) + (("foo bar" + "hello") + "\n\t"))
    v2 -= 590
    if (-2147483648 > 679) {
        var v10 int16 = v2
        v11 := GenS1{F0: v1}
        _ = v10
        _ = v11
    }
    var v12 uint = 1
    fmt.Println(v2, v1)
    v3 = true
    var v13 int8 = -1
    switch "hello" {
    case "foo bar":
        var v14 int = 4294967295
        _ = v14
    case "":
        if (!(v3 || (!true))) {
            switch 32767 {
            case 32767:
                if true {
                    var v15 string = v9
                    v16 := []int{-1, 4294967295, 9223372036854775807}
                    switch 467 {
                    case 0:
                        var v17 int16 = -818
                        _ = v17
                    case -436:
                        _ = len(v16)
                    case 9223372036854775807:
                        var v18 int16 = genFn0((v1 << ((v12) & 63)), 682)
                        _ = v18
                    default:
                        var v19 uint8 = 42
                        _ = v19
                    }
                    _ = v15
                    _ = v16
                }
            }
            var v20 int64 = -9223372036854775808
            _ = v20
        }
    case "hello":
        if (18446744073709551615 != 2) {
            fmt.Println(v9, v1)
            v21 := GenS2{F0: 652}
            for fi22 := 0; fi22 < 17; fi22++ {
                v23 := GenS3{F0: genFn0((v1 / (((v1 % ((v1) | 1))) | 1)), uint16(v1)), F1: 299, F2: 314}
                var v24 uint64 = 1
                var v25 int8 = -1
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = v21
        }
    default:
        if v3 {
            var v26 bool = ((v13 | (v13 % ((-1) | 1))) <= (v4 & v4))
            fmt.Println(v2, v2, v13)
            _ = v26
        } else {
            if v3 {
                v27 := []int{-9223372036854775808, -938, 2147483647, 127}
                var v28 int = -753
                _ = v27
                _ = v28
            } else {
                fmt.Println(v2, v2, v7)
                var v29 bool = v3
                v30 := []int{-145, 42, 4294967295, 1}
                _ = v29
                _ = v30
            }
            var v31 int8 = v4
            _ = v31
        }
    }
    switch ((v9 + (("\n\t" + v9) + (((v9 + v9) + ("" + ("a" + "hello"))) + v9))) + v9) {
    case "hello":
        var v32 uint8 = 2
        _ = v32
    }
    if (42 < 791) {
        if true {
            if (v2 >= v2) {
                var v33 string = v9
                v34 := GenS3{F0: -696, F1: 4294967295, F2: 306}
                _ = v33
                _ = v34
            } else {
                fmt.Println(v12)
                var v35 float64 = (v5 * v5)
                _ = v35
            }
            var v36 float32 = 0.0
            _ = v36
        }
        v37 := []int{115, 32767}
        _ = v37
    } else {
        fmt.Println(v4)
        v38 := GenS2{F0: -9223372036854775808}
        _ = v38
    }
    switch (v9 + v9) {
    case "foo bar":
        v7 = 100
    case "":
        v12 = v12
    default:
        v39 := GenS1{F0: 598}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
