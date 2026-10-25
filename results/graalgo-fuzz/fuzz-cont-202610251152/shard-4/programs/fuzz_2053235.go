package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() string {
    if true {
        var v6 int64 = 9223372036854775807
        _ = v6
    } else {
        if false {
            v7 := []int{998}
            if ((1.5 / (float64((3.14159 / 1e10)) * 0.1)) < 0.1) {
                _ = len(v7)
                var v8 int = -9223372036854775808
                var v9 uint16 = 0
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: 18446744073709551615}
            _ = v10
            _ = v7
        } else {
            var v11 float64 = 1.5
            _ = v11
        }
        var v12 uint8 = 161
        _ = v12
    }
    var v13 float32 = (3.14159 / (float32(0.0) + 0.1))
    v14 := GenS1{F0: 2}
    var v15 uint8 = 1
    v13 += (0.1 * ((v13 - (v13 / 1.5)) / v13))
    _ = v13
    _ = v14
    _ = v15
    return ("hello" + "")
}
func genFn1(p0 uint, p1 uint8, p2 int16) int16 {
    switch 1 {
    case 9223372036854775807:
        fmt.Println(p2)
    }
    var v16 int32 = 137
    var v17 uint16 = 32767
    var v18 int = 32768
    switch v18 {
    case 171:
        if (!false) {
            if true {
                var v19 int = v18
                v18 += v18
                var v20 int = 117
                _ = v19
                _ = v20
            } else {
                for fi21 := 0; fi21 < 4; fi21++ {
                    var v22 bool = true
                    if (false || true) {
                        p0 = 0
                    } else {
                        var v23 int8 = 116
                        var v24 int = (-(fi21))
                        _ = v23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 8; fi25++ {
                        var v26 uint16 = 32767
                        fmt.Println(p0)
                        _ = fi25
                        _ = v26
                    }
                    _ = fi21
                    _ = v22
                }
                switch ("\n\t" + "") {
                case "":
                    var v27 uint16 = 42
                    _ = v27
                case "a":
                    v28 := GenS1{F0: 42}
                    _ = v28
                case "hello":
                    switch (v18 ^ 738) {
                    case 4294967295:
                        var v29 uint64 = 4294967295
                        _ = v29
                    case 1:
                        p0 = p0
                    }
                }
                var v30 uint32 = 2147483647
                _ = v30
            }
        }
    case 0:
        v18 += (^(v18))
    case 127:
        var v31 float64 = (0.1 - (((3.14159 * (3.14159 + ((-2.25 - (((0.1 + 1e10) - (1.5 * 1e10)) / 3.14159)) - 1.5))) + -2.25) + ((1e10 - 3.14159) + 0.0)))
        _ = v31
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return (p2 & p2)
}

func main() {
    var v1 float32 = float32(1.5)
    var v2 uint64 = 2
    var v3 uint = 155
    var v4 uint64 = v2
    var v5 float64 = 0.1
    for fi32 := 0; fi32 < 11; fi32++ {
        if (42 <= 42) {
            var v33 uint16 = 255
            var v34 float64 = ((float64((v1 + v1)) - -2.25) + v5)
            _ = v33
            _ = v34
        }
        v5 += (0.1 - (float64(v1) + float64(1.5)))
        _ = fi32
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        v36 := GenS1{F0: v2}
        _ = fi35
        _ = v36
    }
    fmt.Println(v4)
    v3 = 1
    for fi37 := 0; fi37 < 7; fi37++ {
        var v38 int = fi37
        _ = fi37
        _ = v38
    }
    var v39 uint8 = 0
    fmt.Println(v39, v5)
    var v40 uint32 = 42
    fmt.Println(v39, v5, v40)
    v41 := []int{-98, -583, 132, -1}
    v39 += (v39 % ((v39) | 1))
    v41 = append(v41, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
