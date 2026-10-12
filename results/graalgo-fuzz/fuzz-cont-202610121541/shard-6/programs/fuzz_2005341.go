package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int64 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint; F1 bool; F2 int }
func genFn0(p0 string, p1 uint) string {
    v7 := GenS1{F0: 38, F1: 0, F2: 0}
    p0 = "foo bar"
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            switch fi9 {
            case 42:
                var v10 int32 = 1
                _ = v10
            default:
                if true {
                    var v11 uint = (p1 - (p1 - (p1 >> ((158) & 63))))
                    _ = v11
                }
            }
            _ = fi9
        }
        var v12 float64 = (-2.25 - 3.14159)
        p1 += 42
        _ = fi8
        _ = v12
    }
    v13 := []int{-9223372036854775808, -644}
    var v14 uint8 = 2
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v7
    return p0
}

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    var v4 float64 = (float64(1e10) / 1.5)
    v5 := []int{-1, 42, 467, 986}
    var v6 string = ("\n\t" + "a")
    var v15 uint = 65535
    v16 := GenS1{F0: 2, F1: 42, F2: -9223372036854775808}
    v17 := GenS1{F0: 42, F1: 307, F2: -1}
    var v18 bool = (false || (2 != 4294967295))
    v19 := GenS1{F0: 18446744073709551615, F1: 42, F2: 282}
    for fi20 := 0; fi20 < 8; fi20++ {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v18, fi20, v1)
        }
        var v21 int = 127
        _ = fi20
        _ = v21
    }
    var v22 int8 = 80
    var v23 uint8 = 255
    var v24 int32 = -2147483648
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v25 uint8 = v23
    if (v23 <= 204) {
        switch v6 {
        case "\n\t":
            v5 = append(v5, 927)
        case "":
            v26 := GenS2{F0: 9223372036854775807, F1: 976}
            _ = v26
        case "hello":
            switch ("a" + ((v6 + "a") + v6)) {
            case "foo bar":
                if len(v5) > 0 {
                    v5[0] = 32768
                }
            case "hello":
                var v27 uint8 = (v23 << (((v15 + v15)) & 63))
                _ = v27
            }
        }
        fmt.Println(v6)
        if ((334 == 0) && (!v18)) {
            var v28 uint8 = 42
            v25 -= (v25 >> (((v15 | (v15 & v15))) & 63))
            _ = v28
        } else {
            v25 = v25
            switch ("" + "a") {
            case "foo bar":
                if true {
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        var v29 int32 = (v24 - 42)
                        _ = v29
                    default:
                        var v30 uint16 = 5
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 20; fi31++ {
                        var v32 int64 = 9223372036854775807
                        _ = fi31
                        _ = v32
                    }
                    var v33 float64 = -2.25
                    _ = v33
                } else {
                    var v34 uint64 = 330
                    switch "" {
                    case "\n\t":
                        v35 := GenS1{F0: (v34 << ((v15) & 63)), F1: 31, F2: -623}
                        _ = v35
                    }
                    _ = v34
                }
            case "a":
                v36 := GenS3{F0: (4294967295 < 501), F1: 761}
                _ = v36
            case "hello":
                v37 := v2[249]
                _ = v37
            }
            v38 := GenS4{F0: v15, F1: false, F2: -410}
            _ = v38
        }
    } else {
        _ = len(v5)
        fmt.Println(v18, v22)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
