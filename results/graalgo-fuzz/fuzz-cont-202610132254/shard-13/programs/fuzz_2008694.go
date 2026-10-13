package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 uint32 }
func genFn0() uint32 {
    var v7 uint64 = 2
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint8 = 127
        _ = fi8
        _ = v9
    }
    v10 := []int{42, 128, 9223372036854775807}
    v11 := GenS1{F0: -1}
    var v12 int32 = -624
    _ = v10
    _ = v11
    _ = v12
    _ = v7
    return 2
}

func main() {
    var v1 uint = 1
    v2 := make(map[int]string)
    var v3 int64 = 0
    var v4 float64 = 1.5
    var v5 int8 = -1
    v6 := make(map[string]int)
    if true {
        v4 = float64((1.5 * -2.25))
        if (4294967295 < 323) {
            v13 := GenS2{F0: 127, F1: -1}
            if (0.0 == 0.0) {
                for fi14 := 0; fi14 < 3; fi14++ {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        v16 := v6["foo bar"]
                        _ = fi15
                        _ = v16
                    }
                    var v17 string = "hello"
                    v3 += v3
                    _ = fi14
                    _ = v17
                }
                var v18 float64 = (0.1 * v4)
                _ = v18
            }
            _ = v13
        }
        v19 := []int{-961}
        _ = v19
    }
    var v20 int16 = 699
    v20 += (v20 / ((v20) | 1))
    v6["a"] = 127
    v21 := []int{0}
    if true {
        v4 = v4
        fmt.Println(v3, v3)
    } else {
        var v22 uint64 = 2
        switch "hello" {
        case "a":
            switch "foo bar" {
            case "":
                var v23 uint8 = 2
                _ = v23
            case "a":
                v5 -= v5
            case "hello":
                var v24 int64 = 1
                _ = v24
            }
        case "\n\t":
            for fi25 := 0; fi25 < 12; fi25++ {
                var v26 int8 = v5
                v5 += (v5 / (((v5 * v26)) | 1))
                for fi27 := 0; fi27 < 6; fi27++ {
                    v28 := GenS3{F0: uint64(v1), F1: "hello"}
                    v22 -= 4294967295
                    var v29 uint16 = 2
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = fi25
                _ = v26
            }
        case "hello":
            v30 := GenS1{F0: v3}
            _ = v30
        }
        switch "" {
        case "foo bar":
            var v31 int32 = 0
            _ = v31
        }
        _ = v22
    }
    var v32 string = ("foo bar" + ("a" + ("\n\t" + "hello")))
    fmt.Println(v32, v1, v20)
    v5 = (v5 | (^(v5)))
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    v1 = (v1 + 42)
    v33 := GenS4{F0: 127, F1: -511}
    fmt.Println(v5)
    v34 := GenS5{F0: (v3 - v3), F1: 73}
    if true {
        var v35 uint = v1
        for fi36 := 0; fi36 < 20; fi36++ {
            for fi37 := 0; fi37 < 17; fi37++ {
                v32 = ("a" + v32)
                v38 := GenS6{F0: true, F1: 1}
                var v39 int64 = v3
                _ = fi37
                _ = v38
                _ = v39
            }
            for fi40 := 0; fi40 < 8; fi40++ {
                v41 := GenS5{F0: v3, F1: v1}
                _ = fi40
                _ = v41
            }
            switch -9223372036854775808 {
            case 2147483647:
                v32 = v32
            case 9223372036854775807:
                v32 = ((v32 + (("" + ("foo bar" + "")) + (v32 + "foo bar"))) + "a")
            }
            _ = fi36
        }
        if true {
            var v42 uint32 = 65535
            if len(v21) > 0 {
                v21[0] = 1
            }
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            _ = v42
        } else {
            var v43 uint16 = 300
            v35 = v1
            _ = v43
        }
        _ = v35
    } else {
        var v44 uint = v1
        v1 = v1
        delete(v6, ("a" + ((v32 + v32) + v32)))
        _ = v44
    }
    v45 := GenS7{F0: genFn0()}
    if len(v21) > 0 {
        v21[0] = 42
    }
    v32 = ("a" + (((v32 + v32) + (((("a" + v32) + v32) + v32) + "")) + (v32 + v32)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
