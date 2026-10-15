package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 string }

func main() {
    var v1 float32 = (3.14159 / 1.0)
    var v2 uint = 42
    var v3 int8 = 1
    var v4 bool = (-1 > 451)
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: 4, F1: -9223372036854775808, F2: 42}
        _ = fi5
        _ = v6
    }
    var v7 bool = v4
    v8 := GenS1{F0: 128, F1: 9223372036854775807, F2: 18446744073709551615}
    fmt.Println(v1, v2, v1)
    fmt.Println(v2)
    var v9 uint64 = 4294967295
    fmt.Println(v9, v9)
    var v10 int64 = 9223372036854775807
    switch 42 {
    case 128:
        fmt.Println(v3)
    case 0:
        if v7 {
            var v11 uint16 = 0
            var v12 int8 = (-(v3))
            switch ("hello" + (("" + "a") + "foo bar")) {
            case "":
                v13 := GenS2{F0: (v10 / ((0) | 1)), F1: -1}
                _ = v13
            default:
                v1 -= float32(-2.25)
            }
            _ = v11
            _ = v12
        } else {
            if (v7 && ("\n\t" > "a")) {
                for fi14 := 0; fi14 < 17; fi14++ {
                    v15 := GenS3{F0: v9}
                    var v16 uint32 = 26
                    v4 = (v4 && false)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                var v17 float64 = 1e10
                fmt.Println(v17, v9, v3)
                _ = v17
            } else {
                if v4 {
                    v18 := []int{}
                    _ = v18
                }
                v10 += v10
                for fi19 := 0; fi19 < 18; fi19++ {
                    v20 := GenS4{F0: 295, F1: -1, F2: (v9 + 75)}
                    if (9223372036854775807 <= (v10 - -9223372036854775808)) {
                        v21 := GenS5{F0: v9, F1: v9}
                        fmt.Println(v4, fi19)
                        _ = v21
                    }
                    _ = fi19
                    _ = v20
                }
            }
        }
    case -873:
        fmt.Println(v4, v4)
    }
    switch "hello" {
    case "":
        for fi22 := 0; fi22 < 11; fi22++ {
            for fi23 := 0; fi23 < 18; fi23++ {
                v24 := []int{(fi23 << ((4294967295) & 63))}
                _ = fi23
                _ = v24
            }
            if ((!(!v4)) && (!true)) {
                v25 := []int{fi22}
                for fi26 := 0; fi26 < 16; fi26++ {
                    v25 = append(v25, 127)
                    var v27 int = fi26
                    _ = fi26
                    _ = v27
                }
                for fi28 := 0; fi28 < 6; fi28++ {
                    if len(v25) > 0 {
                        v25[0] = fi22
                    }
                    v9 -= v9
                    _ = fi28
                }
                _ = v25
            }
            if false {
                v3 -= v3
                v29 := []int{-1}
                v30 := GenS6{F0: "\n\t"}
                _ = v29
                _ = v30
            }
            _ = fi22
        }
    }
    v3 = v3
    for fi31 := 0; fi31 < 8; fi31++ {
        var v32 string = "a"
        var v33 string = v32
        v32 = ""
        _ = fi31
        _ = v32
        _ = v33
    }
    v1 += float32((0.1 + 0.0))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
