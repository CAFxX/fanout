package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint16) int8 {
    var v6 float32 = (-2.25 - -2.25)
    v7 := GenS1{F0: 208}
    p0 = (p0 - (p0 - 42))
    if true {
        for fi8 := 0; fi8 < 6; fi8++ {
            switch (fi8 << ((255) & 63)) {
            case 1:
                p0 = p0
            case 9223372036854775807:
                if (((((!(5 != 2)) && (2 > 255)) || (2147483648 < 42)) || false) && ((1.5 * 1e10) == 0.1)) {
                    var v9 float32 = v6
                    var v10 bool = false
                    var v11 int64 = 0
                    _ = v10
                    _ = v11
                    _ = v9
                } else {
                    v12 := GenS1{F0: 889}
                    for fi13 := 0; fi13 < 5; fi13++ {
                        v6 -= v6
                        var v14 int = 2147483647
                        _ = fi13
                        _ = v14
                    }
                    _ = v12
                }
            case -698:
                if ((-1 == 202) && true) {
                    v6 = v6
                    var v15 uint32 = 2147483648
                    _ = v15
                }
            }
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := GenS1{F0: 42}
                var v18 uint32 = 2147483647
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 float64 = (((1e10 + (1e10 - 1e10)) + (float64(v6) - -2.25)) * (0.1 / 0.1))
            _ = fi8
            _ = v19
        }
        var v20 bool = true
        if v20 {
            var v21 uint16 = (p0 % ((p0) | 1))
            v22 := []int{32768, 2147483647, 1}
            _ = v21
            _ = v22
        } else {
            fmt.Println(p0, v6)
            v20 = (-128 > -26)
        }
        _ = v20
    } else {
        v23 := GenS2{F0: 42}
        _ = v23
    }
    var v24 string = ("a" + "a")
    _ = p0
    _ = v24
    _ = v6
    _ = v7
    return -125
}

func main() {
    v1 := []int{0, 1, 42}
    var v2 uint8 = 127
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    var v5 uint = 18446744073709551615
    v2 += (v2 ^ (v2 % (((v2 + 2)) | 1)))
    var v25 string = ""
    if (42 >= 1) {
        if ("hello" == "a") {
            v4[(("" + v25) + ("hello" + ("a" + "foo bar")))] = 296
        }
        var v26 uint16 = 32767
        fmt.Println(v2)
        _ = v26
    }
    var v27 float64 = 0.0
    var v28 int = int(v2)
    fmt.Println(v28, v5)
    v28 = (v28 * v28)
    if false {
        var v29 uint = 42
        if (v2 < (v2 & 43)) {
            switch v28 {
            case 9223372036854775807:
                var v30 bool = ((v25 == (((("" + "foo bar") + ((v25 + v25) + (("\n\t" + "a") + (v25 + v25)))) + ("foo bar" + ("" + v25))) + "")) && false)
                _ = v30
            case 0:
                var v31 int16 = -900
                _ = v31
            case 42:
                for fi32 := 0; fi32 < 16; fi32++ {
                    var v33 uint = v5
                    _ = fi32
                    _ = v33
                }
            }
            v34 := GenS2{F0: -1}
            fmt.Println(v5, v27, v28)
            _ = v34
        }
        switch v28 {
        case 982:
            var v35 uint32 = 42
            _ = v35
        case -927:
            v1 = append(v1, (v28 >> ((v29) & 63)))
        default:
            for fi36 := 0; fi36 < 5; fi36++ {
                var v37 uint8 = 99
                if (!true) {
                    v28 += (fi36 - 0)
                }
                var v38 uint8 = (v37 & 255)
                _ = fi36
                _ = v37
                _ = v38
            }
        }
        _ = v29
    } else {
        v39 := GenS2{F0: 0}
        fmt.Println(v28, v5, v28)
        var v40 bool = false
        _ = v39
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
