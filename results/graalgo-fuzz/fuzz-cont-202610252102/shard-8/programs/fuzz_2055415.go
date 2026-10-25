package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8 }
func genFn0() int64 {
    switch "foo bar" {
    case "":
        v5 := []int{-9223372036854775808, 80}
        _ = v5
    case "foo bar":
        var v6 int64 = 1
        _ = v6
    }
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 17; fi8++ {
        if (1 <= 1) {
            var v9 int8 = -128
            v10 := []int{(fi8 + fi8)}
            _ = v10
            _ = v9
        }
        var v11 string = "\n\t"
        for fi12 := 0; fi12 < 19; fi12++ {
            v11 = (v11 + v11)
            v13 := GenS1{F0: 6}
            _ = fi12
            _ = v13
        }
        _ = fi8
        _ = v11
    }
    _ = v7
    return 0
}
func genFn1() string {
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 uint = 42
        _ = fi14
        _ = v15
    }
    var v16 bool = (1 <= 42)
    v17 := []int{-9223372036854775808, -1}
    for fi18 := 0; fi18 < 19; fi18++ {
        switch "hello" {
        case "hello":
            v16 = true
        case "":
            var v19 float32 = (1.5 + -2.25)
            _ = v19
        case "\n\t":
            v16 = v16
        }
        _ = fi18
    }
    var v20 int64 = -237
    _ = v16
    _ = v17
    _ = v20
    return "hello"
}

func main() {
    var v1 uint8 = 231
    var v2 uint64 = 2
    v3 := make(map[string]string)
    var v4 float64 = 1e10
    v21 := GenS1{F0: (v1 ^ v1)}
    var v22 int = 42
    v3[""] = "foo bar"
    v23 := []int{(-(v22)), 1, 2147483647, v22}
    v24 := GenS1{F0: 2}
    v22 += v22
    var v25 uint8 = v1
    for fi26 := 0; fi26 < 2; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            switch "foo bar" {
            case "\n\t":
                v28 := GenS2{F0: 801}
                _ = v28
            default:
                var v29 int8 = 127
                _ = v29
            }
            _ = fi27
        }
        if ((((v4 * (float64((-2.25 / (1e10 * -2.25))) / 3.14159)) - 0.0) >= v4) && false) {
            fmt.Println(v22, v2, v25)
        }
        _ = fi26
    }
    v22 -= v22
    v25 -= (v1 % (((v25 >> ((4294967295) & 63))) | 1))
    switch (v22 + (-(v22))) {
    case 127:
        v25 = v1
    }
    v30 := GenS3{F0: "a"}
    if false {
        for fi31 := 0; fi31 < 6; fi31++ {
            var v32 uint64 = (v2 << ((42) & 63))
            v4 = v4
            _ = fi31
            _ = v32
        }
        v2 += (v2 & (v2 % (((v2 / (((v2 & (^(v2)))) | 1))) | 1)))
        for fi33 := 0; fi33 < 16; fi33++ {
            var v34 float64 = ((float64(((1.5 - -2.25) * 0.0)) / 1.5) - (1.5 - 3.14159))
            for fi35 := 0; fi35 < 13; fi35++ {
                fmt.Println(v2)
                _ = fi35
            }
            fmt.Println(v2, v4, v1)
            _ = fi33
            _ = v34
        }
    } else {
        var v36 bool = (v2 == (^(v2)))
        _ = v36
    }
    var v37 int = (v22 >> ((42) & 63))
    fmt.Println(v4, v25)
    if (!(v2 >= (v2 << ((690) & 63)))) {
        for fi38 := 0; fi38 < 3; fi38++ {
            if ("" == "foo bar") {
                v4 += (float64((0.1 / 1.5)) / 1.0)
            }
            var v39 string = (("foo bar" + "hello") + "foo bar")
            v3[v39] = v39
            _ = fi38
            _ = v39
        }
        fmt.Println(v25, v2, v1)
        var v40 uint32 = 1
        _ = v40
    } else {
        v41 := GenS4{F0: 1}
        var v42 uint = 748
        _ = v41
        _ = v42
    }
    delete(v3, "hello")
    for fi43 := 0; fi43 < 9; fi43++ {
        fmt.Println(v37, v1)
        var v44 float64 = (v4 * v4)
        if (42 != -128) {
            fmt.Println(v44)
        }
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v37
    _ = v4
}
