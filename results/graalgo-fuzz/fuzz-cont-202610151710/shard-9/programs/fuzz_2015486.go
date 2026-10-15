package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint16 = 0
    v2 := []string{("" + "\n\t"), "foo bar", ("hello" + ("a" + ("a" + ""))), "\n\t"}
    var v3 uint8 = 255
    v4 := make(map[int]bool)
    var v5 int = 9223372036854775807
    if (false && ((^(v1)) != (v1 ^ v1))) {
        switch v5 {
        case -755:
            v3 = (v3 | v3)
        case 1:
            _ = len(v4)
        case 460:
            if false {
                fmt.Println(v1, v3, v1)
                v6 := GenS1{F0: 459, F1: 83}
                _ = v6
            } else {
                v7 := GenS1{F0: -217, F1: -128}
                var v8 int64 = -156
                _ = v7
                _ = v8
            }
        }
    }
    if (v1 <= v1) {
        var v9 string = (((("hello" + "a") + ("\n\t" + ("foo bar" + "hello"))) + "") + "")
        _ = v9
    } else {
        if false {
            v5 += v5
            switch ("foo bar" + "a") {
            case "":
                for fi10 := 0; fi10 < 4; fi10++ {
                    var v11 bool = true
                    _ = fi10
                    _ = v11
                }
            case "\n\t":
                v12 := GenS1{F0: -565, F1: 0}
                _ = v12
            case "hello":
                delete(v4, v5)
            }
            var v13 uint = 42
            _ = v13
        }
        v1 += v1
    }
    fmt.Println(v3, v5)
    switch ("\n\t" + "\n\t") {
    case "a":
        v14 := GenS2{F0: 100}
        _ = v14
    case "":
        v4[v5] = false
    case "hello":
        var v15 float32 = float32(float64(((1e10 + float32(3.14159)) * 0.1)))
        _ = v15
    default:
        v16 := v4[-1]
        _ = v16
    }
    var v17 bool = ((v1 > v1) || ((!true) && false))
    var v18 bool = (((v3 != 127) || v17) || (-1 != -9223372036854775808))
    var v19 uint = 2
    v20 := GenS1{F0: int32(v1), F1: 45}
    if (v18 && v17) {
        if (483 < uint32(v3)) {
            if ((v1 + v1) > v1) {
                v1 = 250
                var v21 bool = v17
                _ = v21
            }
            for fi22 := 0; fi22 < 4; fi22++ {
                v1 += v1
                _ = fi22
            }
            v5 -= (v5 >> ((779) & 63))
        }
        if v17 {
            switch v5 {
            case 127:
                fmt.Println(v5, v3, v1)
            }
            for fi23 := 0; fi23 < 12; fi23++ {
                if len(v2) > 0 {
                    v2[0] = "a"
                }
                _ = fi23
            }
        }
    } else {
        v3 -= (^(v3))
    }
    v17 = v18
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := GenS1{F0: 42, F1: -128}
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 int32 = -2147483648
        _ = fi26
        _ = v27
    }
    if false {
        if (0 > 0) {
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 uint16 = v1
                v17 = (v18 && (((0 < (^(v3))) && (32767 == v1)) && (0 == 1)))
                if (!((true && true) || (0 != v29))) {
                    v18 = false
                }
                _ = fi28
                _ = v29
            }
            switch (v5 + v5) {
            case 735:
                if false {
                    v30 := GenS2{F0: 42}
                    _ = v30
                }
            }
            v31 := GenS1{F0: 0, F1: -43}
            _ = v31
        }
        var v32 float32 = 3.14159
        _ = v32
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch (v5 * 2147483647) {
    case 9223372036854775807:
        var v33 int8 = 127
        _ = v33
    case 985:
        v18 = false
    default:
        v17 = false
    }
    v34 := GenS3{F0: (!(42 >= -1))}
    var v35 bool = v18
    switch "foo bar" {
    case "foo bar":
        var v36 bool = (false || (v18 || false))
        _ = v36
    case "hello":
        var v37 float64 = 1e10
        _ = v37
    default:
        var v38 int32 = 0
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
