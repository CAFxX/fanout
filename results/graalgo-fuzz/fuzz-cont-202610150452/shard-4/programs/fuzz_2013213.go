package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint16; F1 uint }
func genFn0(p0 int16) int64 {
    if (1 != -761) {
        v5 := GenS1{F0: 4294967295}
        p0 = p0
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS2{F0: 0}
            v8 := GenS3{F0: 2, F1: -1}
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    } else {
        v9 := []int{0, -9223372036854775808}
        var v10 int8 = 34
        _ = v10
        _ = v9
    }
    switch (("foo bar" + "a") + "foo bar") {
    case "a":
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS2{F0: 2}
            fmt.Println(p0, fi11)
            v13 := GenS1{F0: 565}
            _ = fi11
            _ = v12
            _ = v13
        }
    case "":
        v14 := []int{}
        _ = v14
    case "\n\t":
        if false {
            var v15 string = "\n\t"
            _ = v15
        } else {
            if (true || true) {
                p0 += 1
            } else {
                var v16 int = -75
                _ = v16
            }
            v17 := []int{752, -2147483648}
            _ = v17
        }
    }
    _ = p0
    return 917
}

func main() {
    var v1 string = ("a" + "hello")
    var v2 string = ("foo bar" + v1)
    var v3 uint = 4294967295
    var v4 float32 = 1e10
    if (255 >= 0) {
        v4 = 1.5
    }
    v2 = v2
    v18 := GenS3{F0: (v3 << ((157) & 63)), F1: 1}
    v19 := []int{9223372036854775807, 42, -2147483648, 127}
    var v20 int32 = 197
    var v21 int = 9223372036854775807
    v22 := GenS1{F0: 802}
    for fi23 := 0; fi23 < 1; fi23++ {
        v21 -= 991
        _ = fi23
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        v21 = (v21 * v21)
        var v25 float64 = (((1e10 + 0.0) / -2.25) * float64((v4 + 3.14159)))
        if false {
            v4 -= (1.5 / v4)
            v26 := GenS1{F0: 4294967295}
            if true {
                if true {
                    var v27 int = (fi24 >> ((v3) & 63))
                    v27 -= fi24
                    _ = v27
                } else {
                    v3 = (v3 | 1)
                    v19 = append(v19, 42)
                }
                fmt.Println(v21, v20)
                v21 = v21
            } else {
                v28 := GenS1{F0: 18446744073709551615}
                if ((((1e10 / (3.14159 - (v4 - v4))) > (v4 * v4)) || ((9223372036854775807 == 0) && (true || (42 < -128)))) && (!(42 != 0))) {
                    fmt.Println(fi24, v4, v2)
                    var v29 int64 = 42
                    v30 := GenS4{F0: 2147483647, F1: ((v1 + "a") + v2)}
                    _ = v29
                    _ = v30
                }
                _ = v28
            }
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        if len(v19) > 0 {
            v19[0] = 0
        }
        _ = fi31
    }
    fmt.Println(v21)
    for fi32 := 0; fi32 < 7; fi32++ {
        v33 := GenS5{F0: 65535, F1: v3}
        _ = fi32
        _ = v33
    }
    var v34 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v4
}
