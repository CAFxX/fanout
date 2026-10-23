package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 bool; F2 uint }
func genFn0(p0 int16, p1 int64) string {
    v5 := []int{-2147483648, -2147483648, 2147483647}
    _ = len(v5)
    _ = p0
    _ = p1
    _ = v5
    return "foo bar"
}
func genFn1(p0 uint32, p1 string) int8 {
    fmt.Println(p0, p0)
    if (1.5 <= (((0.1 - -2.25) * ((((0.0 + ((float32(float64(0.1)) - float32(3.14159)) + 0.0)) - 1e10) / (1e10 - 3.14159)) * (0.0 * 1e10))) / 1.0)) {
        var v6 int64 = -9223372036854775808
        v7 := []int{}
        _ = v6
        _ = v7
    } else {
        v8 := []int{-9223372036854775808, -379, 32767}
        switch "foo bar" {
        case "\n\t":
            fmt.Println(p0)
        case "foo bar":
            var v9 uint32 = p0
            _ = v9
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = v8
    }
    v10 := GenS1{F0: 2, F1: true, F2: 276}
    for fi11 := 0; fi11 < 8; fi11++ {
        if (-115 >= 107) {
            switch p1 {
            case "hello":
                v12 := []int{fi11}
                _ = v12
            }
            var v13 uint32 = 2147483647
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return 51
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{}
    var v3 uint16 = 65535
    var v4 uint = 0
    for fi14 := 0; fi14 < 1; fi14++ {
        if true {
            if true {
                var v15 string = "hello"
                _ = v15
            }
            var v16 int8 = -1
            v17 := GenS1{F0: (v4 - (v4 ^ 18446744073709551615)), F1: (true || (("foo bar" + "a") <= "")), F2: int16(v3)}
            _ = v16
            _ = v17
        } else {
            var v18 bool = false
            v3 -= (^(v3))
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 uint = (v4 - 1)
                v1[((("\n\t" + "foo bar") + (("\n\t" + ((("hello" + (("foo bar" + "") + "")) + "foo bar") + "foo bar")) + "")) + "")] = fi19
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        v4 = 0
        _ = fi14
    }
    if (0.0 < 1e10) {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        _ = len(v1)
    } else {
        if true {
            v21 := GenS1{F0: 654, F1: (!(42 < -1)), F2: 32767}
            for fi22 := 0; fi22 < 8; fi22++ {
                fmt.Println(v3)
                _ = fi22
            }
            v3 -= (v3 | v3)
            _ = v21
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 uint16 = 42
            if ((134 < 509) || false) {
                var v25 uint16 = 796
                v26 := GenS1{F0: uint(v3), F1: true, F2: 32767}
                v3 += v3
                _ = v25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
    }
    _ = len(v2)
    var v27 float32 = 0.0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v28 int8 = 1
    v27 += v27
    for fi29 := 0; fi29 < 19; fi29++ {
        v3 = (v3 / ((42) | 1))
        _ = fi29
    }
    v1["a"] = -1
    var v30 float64 = 0.1
    for fi31 := 0; fi31 < 16; fi31++ {
        for fi32 := 0; fi32 < 16; fi32++ {
            v30 -= -2.25
            _ = fi32
        }
        _ = fi31
    }
    v33 := GenS2{F0: ("\n\t" + "hello")}
    fmt.Println(v28, v4, v28)
    v34 := GenS3{F0: (v4 | 475), F1: (!true), F2: 1}
    var v35 string = ("" + ("foo bar" + ((("a" + "foo bar") + "foo bar") + "")))
    v4 -= v4
    var v36 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
