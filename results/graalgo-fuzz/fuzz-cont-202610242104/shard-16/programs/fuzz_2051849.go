package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 bool) uint8 {
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 uint = 936
            var v7 uint = (v6 + (^(v6)))
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    p0 = ((1e10 + -2.25) == 1.5)
    p0 = (!false)
    p0 = p0
    var v8 float64 = 0.0
    _ = p0
    _ = v8
    return 128
}
func genFn1() int64 {
    var v9 uint16 = 42
    v9 += 0
    switch "a" {
    case "\n\t":
        var v10 float64 = 1.5
        _ = v10
    }
    if ((!true) && ("\n\t" <= "a")) {
        if true {
            for fi11 := 0; fi11 < 6; fi11++ {
                if (255 > (v9 % (((v9 / ((255) | 1))) | 1))) {
                    v9 += 2
                    for fi12 := 0; fi12 < 4; fi12++ {
                        var v13 bool = (false || true)
                        switch "hello" {
                        case "\n\t":
                            var v14 int16 = 32767
                            _ = v14
                        case "":
                            switch (fi11 + fi11) {
                            case 9223372036854775807:
                                var v15 int = 127
                                _ = v15
                            case 32767:
                                v9 += (v9 >> ((2) & 63))
                            case -328:
                                v16 := GenS1{F0: -2147483648}
                                _ = v16
                            }
                        default:
                            v17 := []int{-2147483648, fi12}
                            _ = v17
                        }
                        var v18 bool = v13
                        _ = fi12
                        _ = v13
                        _ = v18
                    }
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v9, v9)
                    case "foo bar":
                        var v19 uint16 = (v9 << ((2) & 63))
                        _ = v19
                    }
                }
                fmt.Println(v9)
                switch "hello" {
                case "":
                    v20 := []int{4294967295, 176, (fi11 + fi11), fi11}
                    _ = v20
                case "foo bar":
                    v21 := []int{-628, fi11, (-(fi11)), (fi11 / ((fi11) | 1))}
                    _ = v21
                }
                _ = fi11
            }
            var v22 uint8 = 255
            _ = v22
        }
    }
    v23 := []int{2147483647}
    _ = v23
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := []string{}
    var v2 int = 128
    var v3 bool = false
    v2 -= (v2 * 419)
    v2 = -9223372036854775808
    var v24 int = -2147483648
    if true {
        v24 -= 12
        if len(v1) > 0 {
            v1[0] = (("\n\t" + "foo bar") + "")
        }
    } else {
        var v25 string = (("\n\t" + "\n\t") + "a")
        var v26 string = "hello"
        var v27 uint16 = 1
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println(v3)
    var v28 int16 = -32768
    v2 = v24
    var v29 int8 = -124
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
}
