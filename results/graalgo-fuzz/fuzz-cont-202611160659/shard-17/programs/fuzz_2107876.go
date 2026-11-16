package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int64 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint) int16 {
    for fi6 := 0; fi6 < 9; fi6++ {
        if true {
            var v7 int8 = -128
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 uint32 = 34
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        switch (fi6 % (((fi6 | 2147483647)) | 1)) {
        case 2147483647:
            v10 := []int{(fi6 / (((fi6 >> ((p1) & 63))) | 1)), 42}
            _ = v10
        }
        _ = fi6
    }
    p0 = (p0 % ((1) | 1))
    var v11 int16 = 0
    fmt.Println(p0, p0, v11)
    _ = p0
    _ = p1
    _ = v11
    return v11
}

func main() {
    var v1 float32 = 0.1
    var v2 string = "foo bar"
    var v3 uint8 = 22
    var v4 uint8 = (^(v3))
    v5 := make(map[uint32]int)
    v12 := GenS1{F0: (true && true), F1: 18446744073709551615, F2: 1}
    v13 := []int{1, 0}
    v3 = 42
    var v14 bool = (171 <= 255)
    v15 := GenS1{F0: true, F1: 1, F2: -9223372036854775808}
    var v16 uint32 = 37
    if false {
        v4 -= 127
        v1 += 1.5
        v17 := GenS1{F0: false, F1: 42, F2: -1}
        _ = v17
    }
    v14 = v14
    v1 += (v1 * ((0.0 * (v1 * v1)) + (v1 - v1)))
    v4 += (v4 & 42)
    switch 9223372036854775807 {
    case 2147483647:
        if (v1 >= (float32((((0.1 + (1.5 * float64((v1 * float32((-2.25 / (-2.25 * (-2.25 / ((0.1 / (-2.25 + -2.25)) * (-2.25 * 0.1)))))))))) / 1.0) * float64(0.0))) + ((float32((0.1 * 0.1)) + 0.0) + (0.0 / (((((v1 - v1) / ((v1 / (v1 / (v1 - 1.5))) / (3.14159 + v1))) * v1) / v1) / v1))))) {
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 int = 128
                _ = len(v13)
                _ = fi18
                _ = v19
            }
            v20 := GenS2{F0: -728, F1: (v4 ^ (v3 + v3))}
            v1 += (v1 + v1)
            _ = v20
        }
    case 1:
        fmt.Println(v1)
    }
    if (127 <= 1) {
        if ((^(v4)) < (^(v4))) {
            for fi21 := 0; fi21 < 20; fi21++ {
                for fi22 := 0; fi22 < 17; fi22++ {
                    fmt.Println(v2)
                    _ = fi22
                }
                fmt.Println(v3, v4, v4)
                v2 = v2
                _ = fi21
            }
            if len(v13) > 0 {
                v13[0] = 9223372036854775807
            }
        } else {
            switch 32768 {
            case 935:
                if v14 {
                    var v23 int32 = -1
                    v24 := v5[(v16 * v16)]
                    _ = v23
                    _ = v24
                }
            case 128:
                v25 := GenS3{F0: 128}
                _ = v25
            }
        }
        if (-2147483648 <= int32(v16)) {
            switch "foo bar" {
            case "":
                v26 := GenS4{F0: 65535}
                _ = v26
            case "\n\t":
                v4 += v3
            case "a":
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            }
        }
        var v27 int8 = -4
        _ = v27
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        v3 -= v4
        _ = fi28
    }
    var v29 int = 4294967295
    delete(v5, (v16 & 2))
    var v30 uint16 = 42
    var v31 string = ("\n\t" + ("foo bar" + ((v2 + v2) + v2)))
    v14 = ((1e10 * float32((1e10 + ((float64(1.5) + (0.1 + 1e10)) * (1e10 / 0.1))))) <= float32(1.5))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
