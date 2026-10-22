package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
func genFn0() bool {
    v5 := []int{32768}
    _ = len(v5)
    switch (("" + ("a" + "foo bar")) + ("foo bar" + ("foo bar" + (((("foo bar" + "\n\t") + "") + "a") + "\n\t")))) {
    case "":
        v6 := GenS1{F0: 2}
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 string = (((("\n\t" + ((("a" + (((("\n\t" + ("a" + ("" + "hello"))) + ("\n\t" + ("" + "\n\t"))) + ("" + "")) + "hello")) + "") + "hello")) + "") + "foo bar") + "a")
            v9 := GenS2{F0: 101}
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    _ = v5
    return (1 > 589)
}
func genFn1() uint {
    if false {
        var v10 int8 = -128
        var v11 uint = 0
        _ = v10
        _ = v11
    } else {
        v12 := []int{}
        _ = v12
    }
    var v13 string = "foo bar"
    _ = v13
    return 713
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]bool)
    var v3 float64 = -2.25
    var v4 uint = 18446744073709551615
    if false {
        var v14 int32 = 42
        var v15 int16 = 1
        fmt.Println(v15)
        _ = v14
        _ = v15
    }
    switch -429 {
    case 32767:
        v3 -= v3
    case -708:
        for fi16 := 0; fi16 < 19; fi16++ {
            _ = len(v2)
            _ = fi16
        }
    default:
        v17 := []int{-1, -2147483648, 32768, -250}
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 float32 = ((1.5 / 1e10) * 1.5)
        var v20 float64 = (v3 + float64(v19))
        v21 := []int{fi18, (^(fi18)), 4294967295}
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    var v22 int16 = 946
    var v23 bool = (4294967295 == 929)
    fmt.Println(v3, v3, v22)
    v24 := GenS2{F0: (v1 << (((v4 << (((v4 >> (((v4 / ((v4) | 1))) & 63))) & 63))) & 63))}
    switch -2147483648 {
    case -9223372036854775808:
        switch 32768 {
        case 4294967295:
            v23 = (0 < 0)
        }
    case 4294967295:
        v25 := GenS2{F0: 42}
        _ = v25
    default:
        fmt.Println(v1, v23, v23)
    }
    v1 -= (v1 + 77)
    v3 = (1.5 - (0.1 + (float64((3.14159 - 0.0)) * v3)))
    fmt.Println(v3, v4)
    v26 := []int{376, -536, 128}
    v3 = v3
    v22 -= v22
    for fi27 := 0; fi27 < 14; fi27++ {
        switch fi27 {
        case 439:
            v28 := GenS2{F0: 1}
            _ = v28
        case 0:
            if v23 {
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                var v29 int32 = 2147483647
                _ = v29
            }
        case 2147483647:
            v30 := GenS1{F0: (v1 >> ((685) & 63))}
            _ = v30
        default:
            switch int(v1) {
            case -900:
                v31 := GenS2{F0: (^(v1))}
                _ = v31
            }
        }
        if len(v26) > 0 {
            v26[0] = (fi27 % (((fi27 >> ((18446744073709551615) & 63))) | 1))
        }
        var v32 int = 9223372036854775807
        _ = fi27
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
}
