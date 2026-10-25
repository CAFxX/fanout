package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 bool }
type GenS2 struct { F0 int16; F1 uint8; F2 int }
type GenS3 struct { F0 int }
func genFn0(p0 uint64) uint16 {
    p0 = p0
    v5 := []int{-2147483648, 0, -9223372036854775808}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = v5
    return 0
}
func genFn1() int8 {
    if true {
        var v6 uint32 = 246
        var v7 uint64 = 2
        var v8 uint = 0
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 string = "a"
    var v10 float32 = (1e10 * ((3.14159 * 1.5) - ((0.0 + float32(1.5)) - (((1e10 + 1e10) - ((((-2.25 + 1.5) * (3.14159 / 1.0)) / -2.25) + 3.14159)) + ((0.0 / (0.1 / float32(float64(float32(1e10))))) / (1e10 - 0.0))))))
    if (1e10 > 0.0) {
        switch 1 {
        case 296:
            v10 = (v10 * 1.5)
        }
        v9 = v9
    }
    _ = v10
    _ = v9
    return 100
}

func main() {
    var v1 uint32 = 65535
    var v2 int16 = -188
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint64 = 42
        if (!false) {
            if true {
                if (53 >= 255) {
                    switch (("hello" + "hello") + "foo bar") {
                    case "hello":
                        v3[fi11] = 42
                    case "foo bar":
                        fmt.Println(v2, v1)
                    case "\n\t":
                        var v13 uint16 = 2
                        _ = v13
                    }
                } else {
                    delete(v3, fi11)
                }
                v14 := GenS1{F0: 53, F1: "\n\t", F2: ((!false) || false)}
                var v15 string = ""
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 14; fi16++ {
                    v17 := []int{fi16, fi16, -9223372036854775808, 32768}
                    _ = fi16
                    _ = v17
                }
                v12 = v12
            }
            var v18 uint32 = v1
            v19 := GenS1{F0: -932, F1: "hello", F2: false}
            _ = v18
            _ = v19
        }
        _ = fi11
        _ = v12
    }
    switch -9223372036854775808 {
    case 0:
        fmt.Println(v1)
    default:
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS1{F0: -442, F1: "\n\t", F2: true}
            delete(v3, -2147483648)
            _ = fi20
            _ = v21
        }
    }
    v22 := GenS1{F0: 1, F1: ("a" + "hello"), F2: (0.0 == ((float64((3.14159 + ((0.1 / (-2.25 - (float32(0.1) - 1.5))) * (((((float32(1e10) - 3.14159) + 1.5) / float32(3.14159)) * (-2.25 + 1.5)) - 0.0)))) - float64(float32(0.1))) + 3.14159))}
    var v23 uint32 = v1
    fmt.Println(v1)
    v24 := GenS2{F0: v2, F1: 2, F2: 32767}
    fmt.Println(v1, v1, v1)
    fmt.Println(v2)
    v2 = (v2 >> ((2) & 63))
    if (!true) {
        for fi25 := 0; fi25 < 6; fi25++ {
            fmt.Println(fi25, fi25, v23)
            v26 := []int{fi25, fi25, -1}
            var v27 int = fi25
            _ = fi25
            _ = v26
            _ = v27
        }
        var v28 float64 = 3.14159
        var v29 uint8 = 237
        _ = v28
        _ = v29
    }
    v2 += 1
    if (-1 != v2) {
        if (0 != -1) {
            v30 := []int{}
            _ = v30
        } else {
            for fi31 := 0; fi31 < 9; fi31++ {
                var v32 int8 = -36
                _ = fi31
                _ = v32
            }
            v23 = v23
        }
        v33 := []int{int(v23), int(v2)}
        _ = v33
    }
    v2 += v2
    v1 = v1
    for fi34 := 0; fi34 < 5; fi34++ {
        if (2147483647 >= -2147483648) {
            v35 := GenS1{F0: 1, F1: ("\n\t" + "\n\t"), F2: (!(4294967295 > 18446744073709551615))}
            _ = v35
        } else {
            var v36 int8 = -128
            if (900 < 18446744073709551615) {
                v37 := []int{-54, fi34, fi34, (fi34 ^ 42)}
                _ = v37
            }
            _ = v36
        }
        var v38 string = "a"
        var v39 int = fi34
        _ = fi34
        _ = v38
        _ = v39
    }
    v23 = v1
    v40 := GenS1{F0: 9223372036854775807, F1: "hello", F2: false}
    v41 := GenS3{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
