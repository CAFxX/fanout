package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int32; F1 int8 }
func genFn0(p0 int64, p1 int64) uint32 {
    if (753 <= 42) {
        var v5 string = "foo bar"
        v6 := []int{-1, 1}
        fmt.Println(p1)
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 string = "hello"
            switch (fi7 << ((468) & 63)) {
            case 128:
                v8 = "a"
            }
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 2, F1: 2}
        if ("hello" <= "") {
            fmt.Println(p0)
        } else {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 string = ("" + "foo bar")
                for fi12 := 0; fi12 < 6; fi12++ {
                    switch v11 {
                    case "a":
                        v13 := []int{}
                        _ = v13
                    }
                    p0 = p1
                    fmt.Println(v11, fi10, p1)
                    _ = fi12
                }
                v14 := []int{42}
                _ = fi10
                _ = v11
                _ = v14
            }
            v15 := GenS1{F0: 768, F1: 1}
            _ = v15
        }
        _ = v9
    }
    if (0 != 42) {
        var v16 uint = 678
        _ = v16
    }
    v17 := GenS1{F0: 65535, F1: 1}
    var v18 int32 = 42
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 4294967295
}
func genFn1(p0 int, p1 int16) int32 {
    var v19 float64 = 0.1
    p0 = p0
    v20 := []int{(p0 & (p0 << ((42) & 63))), (p0 ^ -917), (-(p0)), p0}
    v19 += 0.0
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 127
    var v3 string = ("hello" + ("" + "hello"))
    v4 := make(map[string]uint32)
    v3 = v3
    switch 302 {
    case 551:
        v3 = ""
    case 2147483647:
        v3 = "hello"
    case 1:
        if ((false || (4294967295 >= 18446744073709551615)) || (128 < 4294967295)) {
            var v21 uint32 = 0
            fmt.Println(v21, v2, v21)
            var v22 int16 = 1
            _ = v21
            _ = v22
        } else {
            v2 = v2
            v2 = 42
            fmt.Println(v3, v2)
        }
    default:
        v23 := GenS1{F0: 33, F1: 65535}
        _ = v23
    }
    if false {
        v2 += (v2 ^ 194)
    }
    v3 = ((v3 + v3) + (((v3 + v3) + "hello") + "\n\t"))
    fmt.Println(v3, v2)
    v24 := GenS2{F0: v2, F1: -73}
    switch -1 {
    case 0:
        if (1 != 0) {
            var v25 float64 = float64(float32((3.14159 + (0.1 - (1.5 * (3.14159 * -2.25))))))
            v2 = 2
            _ = v25
        } else {
            if (0 >= 1) {
                var v26 uint32 = 612
                v26 -= 9
                v27 := GenS3{F0: -93}
                _ = v26
                _ = v27
            }
        }
    case 1:
        switch (v3 + "\n\t") {
        case "foo bar":
            v28 := GenS4{F0: 42, F1: 9223372036854775807}
            _ = v28
        default:
            if true {
                delete(v4, (("\n\t" + v3) + (((((("\n\t" + ("hello" + ("" + "foo bar"))) + (v3 + v3)) + (v3 + v3)) + v3) + v3) + v3)))
            } else {
                v4[v3] = 42
                v3 = v3
            }
        }
    case -523:
        fmt.Println(v3, v3)
    default:
        switch (v3 + ((v3 + v3) + "a")) {
        case "\n\t":
            if ("a" > v3) {
                v29 := GenS2{F0: (^(v2)), F1: 128}
                _ = v29
            }
        case "foo bar":
            v3 = v3
        }
    }
    switch ("hello" + "hello") {
    case "a":
        v30 := []int{}
        _ = v30
    case "hello":
        v31 := GenS5{F0: 42, F1: -1}
        _ = v31
    case "":
        for fi32 := 0; fi32 < 18; fi32++ {
            if false {
                var v33 bool = (0.0 >= float64(((((float32((1.5 / 1e10)) - 0.1) * (((3.14159 + 0.1) / 1.0) - 1e10)) / 1.5) / 1e10)))
                var v34 uint64 = 0
                for fi35 := 0; fi35 < 1; fi35++ {
                    var v36 bool = false
                    _ = fi35
                    _ = v36
                }
                _ = v33
                _ = v34
            } else {
                v2 += 1
                var v37 uint8 = (v2 & (^(v2)))
                _ = v37
            }
            _ = fi32
        }
    default:
        _ = len(v4)
    }
    v38 := GenS5{F0: 42, F1: -87}
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v38
    _ = v4
}
