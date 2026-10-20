package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int32; F2 uint64 }
func genFn0(p0 bool) uint {
    p0 = (!(-9223372036854775808 == -122))
    for fi4 := 0; fi4 < 7; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 uint8 = 127
            switch ("hello" + "a") {
            case "foo bar":
                p0 = (136 < 60)
            }
            if p0 {
                fmt.Println(fi5, fi4, fi4)
            }
            v7 := GenS1{F0: (-(fi4)), F1: (^(v6))}
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := []int{-9223372036854775808, fi4, (fi4 * (fi4 | (^(fi4)))), fi4}
        fmt.Println(p0)
        _ = fi4
        _ = v8
    }
    var v9 int16 = -32768
    fmt.Println(v9, p0)
    var v10 int = 1
    _ = p0
    _ = v10
    _ = v9
    return 3
}

func main() {
    v1 := []int{2147483647, -9223372036854775808, -2147483648, -719}
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    switch ("foo bar" + ("hello" + ("foo bar" + ("hello" + "")))) {
    case "hello":
        var v11 uint = 136
        _ = v11
    case "a":
        var v12 float64 = 1.5
        _ = v12
    case "":
        if true {
            v13 := GenS1{F0: 127, F1: 0}
            var v14 string = "foo bar"
            _ = v13
            _ = v14
        }
    }
    var v15 int64 = 42
    switch "" {
    case "\n\t":
        v15 = (-(v15))
    case "hello":
        v16 := GenS2{F0: (v15 + v15)}
        _ = v16
    }
    v15 -= (v15 >> ((1) & 63))
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := GenS3{F0: (((!((0 != 4294967295) || (!(true || ("\n\t" <= ""))))) && false) || (2 > 860)), F1: 0, F2: 42}
        v15 = (v15 / ((int64(fi17)) | 1))
        _ = fi17
        _ = v18
    }
    switch 9223372036854775807 {
    case -324:
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 uint16 = 42
            _ = fi19
            _ = v20
        }
    case 9223372036854775807:
        v15 -= (v15 / ((1) | 1))
    }
    v21 := GenS1{F0: -233, F1: 198}
    v15 -= (v15 % (((^(v15))) | 1))
    v15 = (v15 - v15)
    if ((0.0 > (1.5 / 1.0)) || (false && true)) {
        fmt.Println(v15)
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        fmt.Println(v15)
        v15 -= v15
        _ = fi22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
}
