package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint8, p1 int16) string {
    var v5 float32 = 1e10
    v6 := []int{-227, 32768}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return ""
}

func main() {
    var v1 int32 = 621
    var v2 uint16 = 65535
    var v3 string = ("foo bar" + (("hello" + ("\n\t" + "\n\t")) + ("a" + ("" + "\n\t"))))
    var v4 uint = 4294967295
    if false {
        var v7 uint64 = 0
        _ = v7
    } else {
        if true {
            v2 -= 520
        } else {
            v8 := []int{-2147483648, -773, 966, 844}
            var v9 uint = v4
            var v10 int32 = 42
            _ = v10
            _ = v8
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        if true {
            v3 = v3
            v4 = 413
            fmt.Println(v4, v2, fi11)
        }
        switch ("hello" + v3) {
        case "a":
            var v12 int = (fi11 ^ -969)
            _ = v12
        }
        v4 = v4
        _ = fi11
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v4 -= 42
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 int = (fi13 * (^(fi14)))
            switch (fi13 - (fi14 * (fi13 % ((fi14) | 1)))) {
            case 32768:
                v16 := GenS1{F0: false}
                _ = v16
            case 9223372036854775807:
                fmt.Println(fi13)
            case 1:
                fmt.Println(v3, fi14, fi13)
            default:
                v17 := []int{v15}
                _ = v17
            }
            fmt.Println(v4, v15)
            _ = fi14
            _ = v15
        }
        if true {
            v18 := GenS1{F0: ((v2 != (v2 >> (((v4 >> ((714) & 63))) & 63))) && (1.5 <= ((0.0 * 3.14159) * ((0.1 / ((-2.25 / (3.14159 - float64(-2.25))) - (float64(((float32((3.14159 + 1e10)) / (-2.25 + ((1e10 * -2.25) + 1e10))) / 3.14159)) - (1e10 / 1e10)))) * ((-2.25 * float64((1e10 / 3.14159))) - -2.25)))))}
            _ = v18
        } else {
            if (!true) {
                if true {
                    v19 := []int{(^(fi13)), (-(fi13))}
                    var v20 int = fi13
                    _ = v19
                    _ = v20
                } else {
                    var v21 string = v3
                    var v22 int16 = 32767
                    _ = v21
                    _ = v22
                }
            }
            var v23 float64 = 0.0
            _ = v23
        }
        _ = fi13
    }
    switch 42 {
    case -429:
        if false {
            var v24 int8 = 42
            var v25 uint16 = v2
            _ = v24
            _ = v25
        } else {
            for fi26 := 0; fi26 < 11; fi26++ {
                if true {
                    switch v3 {
                    case "":
                        v27 := []int{4294967295, -1}
                        _ = v27
                    default:
                        for fi28 := 0; fi28 < 3; fi28++ {
                            v29 := GenS1{F0: false}
                            var v30 uint32 = 363
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                } else {
                    var v31 bool = (1.5 < -2.25)
                    var v32 uint64 = 18446744073709551615
                    v2 = (v2 - uint16(v4))
                    _ = v31
                    _ = v32
                }
                v33 := GenS1{F0: false}
                _ = fi26
                _ = v33
            }
        }
    }
    v2 = (^(v2))
    var v34 float64 = -2.25
    if (v34 > (3.14159 - (v34 + (float64(-2.25) / v34)))) {
        var v35 int64 = 971
        v2 = v2
        switch 2147483647 {
        case 9223372036854775807:
            var v36 bool = true
            _ = v36
        case 0:
            fmt.Println(v3)
        }
        _ = v35
    } else {
        v37 := []int{1}
        v37 = append(v37, -2147483648)
        var v38 int = -508
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        v34 -= ((((v34 / (1e10 * (v34 / ((((v34 - v34) * v34) - (v34 - v34)) - (((1e10 - v34) - v34) / (float64(0.0) - 3.14159)))))) - v34) / v34) + ((v34 / -2.25) + v34))
        var v40 int32 = v1
        for fi41 := 0; fi41 < 11; fi41++ {
            v40 = (v40 & (v40 & (v40 | (v40 & v40))))
            _ = fi41
        }
        _ = fi39
        _ = v40
    }
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
