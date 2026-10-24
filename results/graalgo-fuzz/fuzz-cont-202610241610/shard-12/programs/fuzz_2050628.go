package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 string, p1 int16) string {
    p1 += p1
    if (32767 == 804) {
        p1 = (-(p1))
        p1 = p1
    }
    _ = p0
    _ = p1
    return "a"
}
func genFn1(p0 uint, p1 uint64) uint16 {
    fmt.Println(p1, p1)
    if false {
        if (!true) {
            v5 := GenS1{F0: 1}
            v6 := GenS1{F0: 938}
            v7 := []int{128, -937}
            _ = v5
            _ = v6
            _ = v7
        } else {
            v8 := GenS1{F0: 0}
            v9 := GenS1{F0: 65535}
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 42}
        _ = v10
    }
    _ = p0
    _ = p1
    return 851
}

func main() {
    var v1 int32 = -2147483648
    v2 := []int{}
    var v3 int32 = v1
    v4 := make(map[string]bool)
    var v11 uint64 = 42
    var v12 int32 = (^(v3))
    v11 = (v11 / ((42) | 1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        var v13 int32 = 2147483647
        _ = v13
    } else {
        v2 = append(v2, 4294967295)
        v14 := GenS1{F0: 0}
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 uint32 = 1
        if true {
            v17 := GenS1{F0: v16}
            if ((!false) && (0 >= 18446744073709551615)) {
                v16 = 23
            } else {
                var v18 uint8 = 0
                _ = v18
            }
            if false {
                v11 = (v11 - v11)
            } else {
                v19 := v4[("foo bar" + "\n\t")]
                _ = v19
            }
            _ = v17
        } else {
            v1 = v12
            var v20 float64 = 0.1
            _ = v20
        }
        delete(v4, "a")
        _ = fi15
        _ = v16
    }
    if (127 == 89) {
        v21 := GenS2{F0: -192, F1: 1}
        v12 = 42
        for fi22 := 0; fi22 < 14; fi22++ {
            switch "foo bar" {
            case "":
                var v23 uint16 = 358
                _ = v23
            }
            var v24 float32 = ((1e10 - 3.14159) + ((0.1 + 0.1) * (0.0 * ((0.1 - 1e10) * 1.5))))
            fmt.Println(v3)
            _ = fi22
            _ = v24
        }
        _ = v21
    }
    switch -579 {
    case 128:
        var v25 string = "foo bar"
        _ = v25
    case 127:
        for fi26 := 0; fi26 < 13; fi26++ {
            var v27 string = ""
            switch v27 {
            case "\n\t":
                v11 += (v11 - v11)
            }
            _ = fi26
            _ = v27
        }
    }
    v28 := GenS1{F0: 4294967295}
    v4["a"] = false
    v1 -= (v3 % ((2147483647) | 1))
    v29 := GenS3{F0: 65535}
    delete(v4, "")
    for fi30 := 0; fi30 < 19; fi30++ {
        switch "" {
        case "foo bar":
            var v31 float32 = 0.1
            _ = v31
        case "hello":
            if true {
                if true {
                    var v32 uint16 = 0
                    v1 = v3
                    var v33 float32 = 0.1
                    _ = v32
                    _ = v33
                } else {
                    fmt.Println(v1, v1)
                }
            } else {
                v12 = 1
                if (false && (!true)) {
                    var v34 uint16 = genFn1(2, (v11 | (v11 + 545)))
                    _ = v34
                } else {
                    v11 -= 45
                }
            }
        }
        v3 = -229
        var v35 float64 = float64(1.5)
        _ = fi30
        _ = v35
    }
    var v36 bool = (true || (138 != 14))
    v11 = v11
    v4["a"] = (int64(v11) < 52)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
}
