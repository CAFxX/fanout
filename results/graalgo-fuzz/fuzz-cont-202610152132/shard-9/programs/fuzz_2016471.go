package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 uint64 }
func genFn0() int64 {
    var v5 int64 = 940
    switch ("foo bar" + "") {
    case "hello":
        fmt.Println(v5)
    }
    switch ("hello" + "foo bar") {
    case "":
        v5 = v5
    case "a":
        if (42 >= 756) {
            v6 := GenS1{F0: -1}
            _ = v6
        }
    }
    v7 := GenS1{F0: 1}
    _ = v5
    _ = v7
    return v5
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 1
    var v3 string = ""
    var v4 float64 = (0.1 - 3.14159)
    var v8 int16 = -1
    v8 -= v8
    var v9 int16 = 32767
    if false {
        fmt.Println(v8)
    } else {
        v3 = v3
    }
    fmt.Println(v3)
    if (false || ((v3 != ((v3 + v3) + (v3 + ((v3 + "") + "hello")))) && true)) {
        delete(v1, "")
        fmt.Println(v9)
        v1["foo bar"] = ((!(!true)) && true)
    }
    if true {
        v10 := []int{-1}
        if (float64(0.0) <= ((v4 + (v4 * v4)) + 3.14159)) {
            for fi11 := 0; fi11 < 2; fi11++ {
                v3 = ("hello" + v3)
                _ = fi11
            }
            if (!((v2 & -9223372036854775808) >= v2)) {
                var v12 uint64 = 18446744073709551615
                _ = v12
            } else {
                v8 = (v9 ^ v9)
            }
            for fi13 := 0; fi13 < 12; fi13++ {
                v14 := GenS2{F0: false}
                v4 = ((0.0 - 0.0) / (float64((1e10 / 1e10)) + 0.1))
                _ = fi13
                _ = v14
            }
        } else {
            v15 := GenS3{F0: 42}
            v2 += (v2 - genFn0())
            v10 = append(v10, 539)
            _ = v15
        }
        fmt.Println(v3)
        _ = v10
    }
    var v16 int = 42
    var v17 uint16 = 245
    switch 32768 {
    case 4294967295:
        v9 -= (v8 << ((255) & 63))
    case 2147483647:
        var v18 uint = 4294967295
        _ = v18
    default:
        var v19 uint16 = v17
        _ = v19
    }
    _ = len(v1)
    if false {
        var v20 float32 = (-2.25 / ((-2.25 / 1.5) + -2.25))
        v17 -= (v17 & (v17 | 2))
        _ = v20
    }
    v21 := GenS4{F0: 1}
    var v22 float32 = 0.0
    if true {
        switch v3 {
        case "hello":
            for fi23 := 0; fi23 < 10; fi23++ {
                v9 = (v9 >> ((255) & 63))
                switch (v3 + v3) {
                case "a":
                    var v24 string = v3
                    _ = v24
                }
                _ = fi23
            }
        case "\n\t":
            switch v16 {
            case 128:
                for fi25 := 0; fi25 < 10; fi25++ {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        v27 := GenS2{F0: (127 == 127)}
                        var v28 bool = false
                        v8 = v8
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    _ = fi25
                }
            default:
                var v29 bool = false
                _ = v29
            }
        default:
            v30 := GenS5{F0: 4294967295, F1: 1}
            _ = v30
        }
        v16 += v16
    } else {
        v31 := GenS1{F0: 2147483647}
        var v32 int32 = -1
        switch (^(v16)) {
        case 2147483647:
            v33 := v1[(((v3 + ((("" + (("hello" + "\n\t") + "\n\t")) + (((((("" + (("foo bar" + v3) + (("a" + "\n\t") + ((v3 + (v3 + (v3 + ""))) + v3)))) + v3) + ((v3 + v3) + v3)) + v3) + v3) + "a")) + "a")) + v3) + (v3 + ""))]
            _ = v33
        case 9223372036854775807:
            var v34 uint8 = 2
            _ = v34
        default:
            v32 = 888
        }
        _ = v31
        _ = v32
    }
    v8 += v8
    v35 := GenS1{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v35
    _ = v4
    _ = v8
    _ = v9
}
