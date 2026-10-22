package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    v1 := []string{}
    var v2 float32 = (0.0 / 1.5)
    v3 := []uint32{65535, 42}
    var v4 int = -1
    var v5 float32 = (float32(0.1) * (0.0 * v2))
    _ = len(v1)
    switch "" {
    case "hello":
        var v6 uint32 = 2
        _ = v6
    case "foo bar":
        var v7 uint = 65535
        _ = v7
    default:
        if (!(false && (-128 >= 116))) {
            var v8 uint8 = 101
            _ = v8
        } else {
            var v9 uint64 = 4294967295
            var v10 uint = 1
            _ = v10
            _ = v9
        }
    }
    if true {
        switch (-(v4)) {
        case 817:
            switch (-(v4)) {
            case 32767:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 4294967295:
                fmt.Println(v2)
            }
        case 32767:
            if (!(647 == 14)) {
                _ = len(v1)
            } else {
                v11 := GenS1{F0: 42, F1: false}
                var v12 float64 = 0.0
                var v13 uint = 4294967295
                _ = v11
                _ = v12
                _ = v13
            }
        }
        switch (^(v4)) {
        case -2:
            for fi14 := 0; fi14 < 7; fi14++ {
                v1 = append(v1, ("a" + "hello"))
                var v15 uint = 1
                _ = fi14
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            v2 -= (v5 - v5)
            _ = fi16
        }
    }
    if (147 <= 255) {
        var v17 string = ("" + "foo bar")
        _ = v17
    } else {
        var v18 uint8 = 0
        for fi19 := 0; fi19 < 8; fi19++ {
            switch (fi19 - -790) {
            case -863:
                var v20 uint32 = 286
                _ = v20
            case -1:
                var v21 uint8 = v18
                _ = v21
            }
            _ = fi19
        }
        _ = v18
    }
    if (!((18446744073709551615 != 42) || false)) {
        for fi22 := 0; fi22 < 6; fi22++ {
            if false {
                for fi23 := 0; fi23 < 20; fi23++ {
                    var v24 int64 = -760
                    v24 = v24
                    _ = fi23
                    _ = v24
                }
            }
            fmt.Println(v4)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi22
        }
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        var v25 int16 = 1000
        _ = v25
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "hello" {
        case "hello":
            if (42 > 18446744073709551615) {
                v1 = append(v1, "foo bar")
                switch v4 {
                case 1:
                    var v26 uint = 2
                    _ = v26
                case 32768:
                    var v27 float64 = 0.1
                    _ = v27
                default:
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                }
                fmt.Println(v4, v4, v5)
            }
        case "\n\t":
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        default:
            if true {
                var v28 int32 = 0
                _ = v28
            } else {
                if ((0.0 - v5) < v5) {
                    fmt.Println(v4, v5, v4)
                }
            }
        }
        v1 = append(v1, ("hello" + "hello"))
    }
    v29 := GenS1{F0: 42, F1: (true && (false && true))}
    if (((-2.25 * float32(3.14159)) - v5) == v2) {
        if (3.14159 <= 1e10) {
            fmt.Println(v4)
            var v30 int8 = 0
            fmt.Println(v4, v5)
            _ = v30
        }
    }
    fmt.Println(v4, v2)
    v4 -= 1
    var v31 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
