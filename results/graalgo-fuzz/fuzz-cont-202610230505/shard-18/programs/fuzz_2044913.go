package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint; F1 int }
func genFn0() uint16 {
    var v5 uint = 18446744073709551615
    var v6 uint64 = 0
    var v7 uint16 = 65535
    if false {
        var v8 uint = (^(v5))
        v9 := GenS1{F0: v6}
        v10 := []int{4294967295, 670, -881, 2147483647}
        _ = v10
        _ = v8
        _ = v9
    } else {
        v11 := []int{32768, 9223372036854775807}
        _ = v11
    }
    _ = v5
    _ = v6
    _ = v7
    return 513
}

func main() {
    var v1 string = ""
    var v2 int32 = 289
    var v3 uint16 = 536
    var v4 uint32 = 2
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := []int{}
        if false {
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 int64 = 42
                switch ((v1 + "a") + "hello") {
                case "foo bar":
                    v16 := GenS1{F0: 777}
                    _ = v16
                case "hello":
                    switch 32768 {
                    case 9223372036854775807:
                        fmt.Println(fi12, v3, fi12)
                        v3 = 491
                    default:
                        _ = len(v13)
                    }
                }
                fmt.Println(v1, fi14)
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                _ = fi14
                _ = v15
            }
            v1 = v1
        }
        var v17 float64 = ((float64(0.1) + (3.14159 - float64(1.5))) - -2.25)
        _ = fi12
        _ = v13
        _ = v17
    }
    v18 := []int{32768, 4294967295}
    switch 32768 {
    case 4294967295:
        fmt.Println(v3, v1)
    case 127:
        switch 0 {
        case -476:
            if len(v18) > 0 {
                v18[0] = 9223372036854775807
            }
        case 127:
            v1 = ""
        default:
            if false {
                fmt.Println(v3, v3)
                var v19 int64 = 9223372036854775807
                _ = v19
            }
        }
    case 308:
        switch ((("a" + (v1 + (v1 + (((v1 + ("hello" + v1)) + "") + v1)))) + ((((v1 + v1) + (v1 + v1)) + "") + (v1 + v1))) + v1) {
        case "foo bar":
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 uint64 = 42
                _ = fi20
                _ = v21
            }
        case "":
            for fi22 := 0; fi22 < 7; fi22++ {
                switch "" {
                case "hello":
                    for fi23 := 0; fi23 < 10; fi23++ {
                        var v24 int32 = 269
                        v25 := GenS2{F0: -1, F1: v1}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
                var v26 int64 = 9223372036854775807
                _ = fi22
                _ = v26
            }
        case "\n\t":
            v18 = append(v18, -229)
        }
    }
    fmt.Println(v1, v2)
    for fi27 := 0; fi27 < 13; fi27++ {
        v3 = genFn0()
        var v28 int16 = 32767
        _ = fi27
        _ = v28
    }
    var v29 int32 = 0
    switch 2147483647 {
    case 9223372036854775807:
        var v30 int = 128
        _ = v30
    case -824:
        var v31 int8 = -128
        _ = v31
    case 32767:
        if (int16(v3) == 32767) {
            var v32 float32 = 3.14159
            _ = v32
        } else {
            v1 = "foo bar"
        }
    default:
        v4 -= (v4 & v4)
    }
    if len(v18) > 0 {
        v18[0] = -1
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        var v34 float32 = (0.0 / 1.5)
        v4 += 2147483647
        var v35 uint = 4294967295
        _ = fi33
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        fmt.Println(v4, fi36, v29)
        _ = fi36
    }
    var v37 int = 32767
    switch 32767 {
    case 1:
        v38 := GenS2{F0: 100, F1: (v1 + (v1 + v1))}
        _ = v38
    default:
        var v39 uint64 = 18446744073709551615
        _ = v39
    }
    v40 := GenS3{F0: 608, F1: -9223372036854775808}
    var v41 uint64 = 42
    var v42 int8 = -56
    v29 += (v2 & v2)
    for fi43 := 0; fi43 < 16; fi43++ {
        if (9223372036854775807 > 1) {
            v44 := GenS1{F0: v41}
            _ = v44
        } else {
            v42 = (v42 + 1)
        }
        _ = fi43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
