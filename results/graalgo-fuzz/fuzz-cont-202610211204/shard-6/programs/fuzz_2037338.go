package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
func genFn0() int64 {
    if ((float32(1.5) <= (3.14159 - 3.14159)) || false) {
        var v5 int8 = 1
        _ = v5
    } else {
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := GenS1{F0: 925, F1: 2}
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := []int{32767, 484}
                var v10 uint = 0
                _ = fi8
                _ = v10
                _ = v9
            }
            for fi11 := 0; fi11 < 6; fi11++ {
                fmt.Println(fi6, fi11)
                fmt.Println(fi11, fi11)
                var v12 int8 = 127
                fmt.Println(fi6, v12)
                _ = fi11
                _ = v12
            }
            _ = fi6
            _ = v7
        }
    }
    var v13 int8 = 127
    var v14 int8 = v13
    _ = v13
    _ = v14
    return 1
}

func main() {
    var v1 uint8 = 127
    var v2 string = ("hello" + "\n\t")
    var v3 uint8 = 42
    var v4 int = 127
    v15 := GenS1{F0: 0, F1: 0}
    switch v2 {
    case "":
        v16 := []int{v4, v4, (v4 - v4)}
        _ = v16
    case "hello":
        v17 := GenS1{F0: 42, F1: 1}
        _ = v17
    case "a":
        v18 := GenS1{F0: 65535, F1: 42}
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        switch (v2 + "") {
        case "":
            var v20 uint = 42
            _ = v20
        }
        _ = fi19
    }
    if (!("a" > (v2 + (v2 + v2)))) {
        fmt.Println(v1, v1)
        fmt.Println(v3, v2, v2)
        var v21 float32 = 3.14159
        _ = v21
    } else {
        fmt.Println(v3)
    }
    var v22 uint64 = 1
    v2 = "a"
    var v23 float32 = 0.1
    v4 = v4
    v22 += (^(v22))
    v23 += v23
    if false {
        v24 := []int{}
        switch v4 {
        case -9223372036854775808:
            fmt.Println(v4, v4, v4)
        case 4294967295:
            if (!(42 == 199)) {
                v25 := GenS2{F0: 0}
                _ = v25
            } else {
                v4 += 745
                var v26 string = "foo bar"
                fmt.Println(v4, v1)
                _ = v26
            }
        case -1:
            v27 := GenS3{F0: v4}
            _ = v27
        default:
            var v28 int16 = 0
            _ = v28
        }
        _ = v24
    } else {
        var v29 uint32 = 979
        for fi30 := 0; fi30 < 15; fi30++ {
            for fi31 := 0; fi31 < 1; fi31++ {
                var v32 uint64 = uint64(v3)
                _ = fi31
                _ = v32
            }
            var v33 float64 = (-2.25 * 0.0)
            v1 += v3
            _ = fi30
            _ = v33
        }
        _ = v29
    }
    switch ((v2 + v2) + v2) {
    case "a":
        v3 -= (v3 & v1)
    case "\n\t":
        v4 = v4
    }
    v1 += v1
    if (v3 == (v3 >> ((837) & 63))) {
        v34 := GenS3{F0: v4}
        var v35 int32 = 0
        _ = v34
        _ = v35
    }
    if true {
        switch (v4 >> ((uint(v22)) & 63)) {
        case 2147483647:
            var v36 uint8 = v3
            _ = v36
        }
        var v37 bool = (-1 > 1)
        fmt.Println(v3, v3, v22)
        _ = v37
    }
    v3 = (v1 | uint8(v22))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
