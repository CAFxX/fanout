package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int64 = -374
    var v2 int8 = -32
    v3 := make(map[uint32]string)
    var v4 float32 = 3.14159
    v1 = -1
    var v5 int16 = 42
    switch (("foo bar" + "a") + "a") {
    case "\n\t":
        v6 := []int{}
        _ = v6
    case "a":
        v2 += (v2 - (v2 % ((v2) | 1)))
    case "foo bar":
        v7 := []int{4294967295}
        _ = v7
    }
    v1 += v1
    v8 := []int{1}
    switch 1 {
    case 931:
        var v9 uint64 = 0
        _ = v9
    case -9223372036854775808:
        switch "" {
        case "":
            v4 += v4
        }
    case -150:
        var v10 int8 = v2
        _ = v10
    default:
        var v11 uint8 = 1
        _ = v11
    }
    if (!true) {
        var v12 uint16 = 1
        _ = v12
    } else {
        v13 := v3[78]
        var v14 uint = 396
        fmt.Println(v13, v14, v1)
        _ = v13
        _ = v14
    }
    _ = len(v3)
    v2 = -1
    var v15 string = "foo bar"
    v16 := GenS1{F0: false}
    v17 := GenS1{F0: true}
    v18 := GenS1{F0: (v5 <= v5)}
    v5 = v5
    v15 = v15
    if true {
        for fi19 := 0; fi19 < 3; fi19++ {
            v5 -= 952
            for fi20 := 0; fi20 < 13; fi20++ {
                v8 = append(v8, (fi19 << ((0) & 63)))
                var v21 float64 = 3.14159
                fmt.Println(v5, v2)
                switch "" {
                case "hello":
                    var v22 int8 = (-(v2))
                    _ = v22
                case "a":
                    var v23 int = -441
                    v2 -= -1
                    v2 = (v2 ^ v2)
                    switch ("a" + ("" + ("" + "\n\t"))) {
                    case "":
                        _ = len(v3)
                    default:
                        var v24 int = (v23 & v23)
                        _ = v24
                    }
                    _ = v23
                default:
                    var v25 uint16 = 255
                    _ = v25
                }
                v4 += v4
                switch 45 {
                case 2147483647:
                    if len(v8) > 0 {
                        v8[0] = (^(fi20))
                    }
                case -918:
                    v1 -= (v1 ^ 9223372036854775807)
                case 918:
                    v1 += (v1 << ((255) & 63))
                }
                v26 := GenS1{F0: (false && false)}
                var v27 string = "foo bar"
                delete(v3, 285)
                _ = fi20
                _ = v21
                _ = v26
                _ = v27
            }
            _ = fi19
        }
        if true {
            switch 684 {
            case 127:
                v2 = (-(v2))
            default:
                if false {
                    fmt.Println(v5, v15)
                    var v28 uint32 = 42
                    v1 += -9223372036854775808
                    _ = v28
                }
            }
        } else {
            if (v15 != "hello") {
                var v29 float64 = float64(v4)
                var v30 uint64 = 228
                _ = v29
                _ = v30
            } else {
                _ = len(v8)
                switch 260 {
                case 0:
                    var v31 uint64 = 0
                    _ = v31
                case -836:
                    delete(v3, 0)
                default:
                    if false {
                        var v32 float64 = 1e10
                        _ = v32
                    }
                }
                var v33 float32 = v4
                _ = v33
            }
            v15 = ""
            v4 = (v4 - (v4 * v4))
        }
        if true {
            var v34 bool = (!(1 == 430))
            delete(v3, 777)
            _ = v34
        }
    } else {
        var v35 uint = 0
        v36 := GenS1{F0: (867 > 276)}
        v37 := GenS1{F0: true}
        _ = v35
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
