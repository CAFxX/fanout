package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
func genFn0() int16 {
    v6 := GenS1{F0: 18446744073709551615}
    var v7 uint16 = 65535
    v8 := []int{42}
    var v9 uint32 = 392
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 895
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 1
    var v3 float64 = 0.1
    var v4 float64 = v3
    var v5 float32 = 0.0
    delete(v1, ("a" + ("hello" + ("hello" + ((("a" + "hello") + ("a" + (("\n\t" + "a") + "hello"))) + "hello")))))
    var v10 int64 = 42
    v11 := GenS1{F0: v2}
    for fi12 := 0; fi12 < 19; fi12++ {
        v3 -= (v4 / 0.1)
        _ = fi12
    }
    v10 -= v10
    v13 := GenS1{F0: 255}
    for fi14 := 0; fi14 < 18; fi14++ {
        switch (fi14 ^ (fi14 * fi14)) {
        case 0:
            _ = len(v1)
        case 32767:
            for fi15 := 0; fi15 < 4; fi15++ {
                v16 := GenS2{F0: 128}
                var v17 int = -1
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        var v18 uint16 = 122
        _ = fi14
        _ = v18
    }
    switch "foo bar" {
    case "":
        v19 := GenS2{F0: 1}
        _ = v19
    case "\n\t":
        v10 += 968
    case "foo bar":
        v1["hello"] = false
    default:
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 uint = 8
            _ = fi20
            _ = v21
        }
    }
    v22 := v1["foo bar"]
    switch "foo bar" {
    case "\n\t":
        var v23 float64 = float64((1e10 + v5))
        _ = v23
    }
    if true {
        for fi24 := 0; fi24 < 11; fi24++ {
            v25 := GenS3{F0: (^(v10))}
            if v22 {
                var v26 int8 = 111
                _ = v26
            }
            _ = fi24
            _ = v25
        }
        switch ((("" + "\n\t") + "foo bar") + (("" + "foo bar") + ("hello" + ("" + ("hello" + "hello"))))) {
        case "\n\t":
            v3 -= (((float64(float32((0.0 * ((-2.25 - (0.1 - float64((v5 / v5)))) / v3)))) + v3) / float64(((-2.25 + (v5 + float32(v3))) * v5))) / float64(v5))
        }
        fmt.Println(v10, v22)
    }
    v27 := GenS1{F0: (^(v2))}
    if v22 {
        fmt.Println(v5, v10, v3)
        if (106 >= 42) {
            v28 := []int{4294967295, 4294967295}
            delete(v1, "\n\t")
            if (!false) {
                var v29 uint64 = 67
                var v30 int64 = v10
                v30 = (v10 ^ v10)
                _ = v29
                _ = v30
            } else {
                v22 = true
                fmt.Println(v4)
            }
            _ = v28
        }
        v31 := GenS2{F0: 1}
        _ = v31
    } else {
        var v32 int = 0
        v2 = (v2 + v2)
        _ = v32
    }
    v2 += (v2 << (((v2 + v2)) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
