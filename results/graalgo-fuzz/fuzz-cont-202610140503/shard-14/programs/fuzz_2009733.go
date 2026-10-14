package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int) int32 {
    if false {
        if (!(true || false)) {
            p0 = 2147483647
        }
        var v5 uint64 = 18446744073709551615
        var v6 int16 = 1000
        _ = v5
        _ = v6
    } else {
        p0 = 1
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        if true {
            p0 -= (fi7 - 128)
            var v8 uint32 = 555
            _ = v8
        }
        if (2 >= 1) {
            switch "foo bar" {
            case "foo bar":
                var v9 int32 = 1
                _ = v9
            case "hello":
                var v10 float64 = -2.25
                _ = v10
            case "":
                var v11 float32 = 1.5
                _ = v11
            }
        } else {
            p0 = (p0 - (p0 % (((p0 + (p0 << ((4294967295) & 63)))) | 1)))
            if true {
                fmt.Println(p0, fi7)
                p0 = p0
                var v12 int = p0
                _ = v12
            } else {
                p0 -= (p0 / ((32767) | 1))
            }
            var v13 int = 32768
            _ = v13
        }
        var v14 string = "\n\t"
        _ = fi7
        _ = v14
    }
    _ = p0
    return 0
}

func main() {
    var v1 uint64 = 4294967295
    var v2 string = ("" + ("\n\t" + ("" + ("a" + "foo bar"))))
    var v3 uint64 = 4294967295
    var v4 uint16 = 2
    fmt.Println(v3)
    for fi15 := 0; fi15 < 6; fi15++ {
        v4 -= v4
        _ = fi15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 5; fi17++ {
            switch "" {
            case "hello":
                v18 := []int{(-(fi16)), int(v4), (fi16 >> ((1) & 63))}
                _ = v18
            case "\n\t":
                v2 = v2
            case "":
                fmt.Println(fi17, v3, v4)
            default:
                v3 -= (^(v3))
            }
            v19 := []int{fi17, fi17}
            _ = fi17
            _ = v19
        }
        var v20 string = v2
        _ = fi16
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 uint16 = (v4 ^ 32767)
        _ = fi21
        _ = v22
    }
    var v23 string = v2
    fmt.Println(v2)
    v24 := []int{-9223372036854775808, -9223372036854775808, -2147483648}
    var v25 int16 = -1
    if ((float64(-2.25) != ((float64((1.5 - 0.0)) - 1.5) + 0.0)) && true) {
        v26 := GenS1{F0: -9223372036854775808}
        var v27 uint = 917
        v2 = v23
        _ = v26
        _ = v27
    } else {
        v24 = append(v24, 32767)
        fmt.Println(v3)
    }
    v1 -= v3
    fmt.Println(v4, v3, v23)
    switch 32768 {
    case 147:
        _ = len(v24)
    case 0:
        if len(v24) > 0 {
            v24[0] = 4294967295
        }
    }
    if len(v24) > 0 {
        v24[0] = 42
    }
    switch v2 {
    case "a":
        for fi28 := 0; fi28 < 3; fi28++ {
            v1 -= v3
            v29 := GenS1{F0: -1}
            var v30 uint32 = 0
            _ = fi28
            _ = v29
            _ = v30
        }
    case "":
        v2 = v2
    default:
        v4 = v4
    }
    var v31 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v4
}
