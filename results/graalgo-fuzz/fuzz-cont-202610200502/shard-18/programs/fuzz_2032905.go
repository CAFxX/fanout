package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint8 }
func genFn0(p0 uint8, p1 int16, p2 int8) int8 {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 float32 = 1e10
        fmt.Println(fi6, p1, p2)
        _ = fi6
        _ = v7
    }
    if false {
        var v8 int8 = p2
        _ = v8
    }
    var v9 float64 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p2
}

func main() {
    var v1 uint16 = 1
    var v2 int8 = -128
    v3 := []int{-266, 272, 2147483647}
    var v4 string = "foo bar"
    var v5 string = (v4 + ((v4 + (v4 + v4)) + (v4 + "")))
    var v10 bool = false
    var v11 uint16 = 355
    var v12 int = 128
    for fi13 := 0; fi13 < 11; fi13++ {
        switch v5 {
        case "foo bar":
            for fi14 := 0; fi14 < 6; fi14++ {
                switch 32767 {
                case -2147483648:
                    if len(v3) > 0 {
                        v3[0] = -634
                    }
                case 4294967295:
                    var v15 bool = v10
                    _ = v15
                }
                if v10 {
                    fmt.Println(v4, v5, v12)
                    var v16 string = v4
                    _ = v16
                }
                _ = fi14
            }
        case "":
            v4 = ((((v4 + v4) + "foo bar") + "a") + ((("hello" + v4) + ("hello" + v4)) + (v4 + ("a" + "a"))))
        }
        if false {
            if v10 {
                for fi17 := 0; fi17 < 17; fi17++ {
                    fmt.Println(v2)
                    fmt.Println(v2)
                    v12 += fi13
                    if (!(!(!(v10 && false)))) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v18 := GenS1{F0: v11}
                        _ = len(v3)
                        var v19 int = (v12 << ((196) & 63))
                        _ = v18
                        _ = v19
                    }
                    _ = fi17
                }
                if true {
                    fmt.Println(v10, v2)
                }
            }
            v2 = 104
        }
        _ = fi13
    }
    if (!(true && v10)) {
        var v20 int8 = (v2 + (v2 % (((v2 | v2)) | 1)))
        v4 = v4
        _ = v20
    }
    var v21 uint8 = 149
    if true {
        if (-77 < (v2 % (((v2 << ((336) & 63))) | 1))) {
            for fi22 := 0; fi22 < 12; fi22++ {
                _ = len(v3)
                _ = fi22
            }
        } else {
            if v10 {
                switch (v12 ^ (v12 / ((v12) | 1))) {
                case 0:
                    var v23 string = v5
                    _ = v23
                case -1:
                    for fi24 := 0; fi24 < 4; fi24++ {
                        v21 = v21
                        _ = fi24
                    }
                case 128:
                    v25 := GenS2{F0: 18446744073709551615}
                    _ = v25
                }
                for fi26 := 0; fi26 < 17; fi26++ {
                    fmt.Println(v10, v21)
                    v10 = ((!v10) || v10)
                    _ = fi26
                }
            }
            switch (v12 / ((4294967295) | 1)) {
            case 1:
                if v10 {
                    v1 = (v1 ^ v11)
                } else {
                    var v27 float32 = -2.25
                    _ = len(v3)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v27
                }
            }
            var v28 bool = v10
            _ = v28
        }
    }
    fmt.Println(v5, v12)
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 string = ""
        _ = fi29
        _ = v30
    }
    v12 -= v12
    v10 = (0 <= 2147483647)
    switch ("a" + "\n\t") {
    case "foo bar":
        if true {
            if v10 {
                v21 = (v21 * 42)
                if len(v3) > 0 {
                    v3[0] = 0
                }
            }
        } else {
            fmt.Println(v5)
            v31 := GenS3{F0: 0, F1: 4294967295, F2: 138}
            _ = v31
        }
    case "a":
        if len(v3) > 0 {
            v3[0] = v12
        }
    default:
        v3 = append(v3, 4294967295)
    }
    v12 += (v12 ^ (v12 / (((v12 * (v12 - v12))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
