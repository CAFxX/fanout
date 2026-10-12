package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 bool }
type GenS2 struct { F0 int32; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint64; F1 string; F2 uint }
type GenS4 struct { F0 int64; F1 string }
func genFn0(p0 string, p1 uint32) string {
    var v6 int8 = 100
    v7 := []int{9223372036854775807, 442, 0, -272}
    if ((42 != 2) || (128 >= 2)) {
        fmt.Println(v6)
        switch (p0 + p0) {
        case "\n\t":
            v8 := GenS1{F0: 54, F1: 213, F2: true}
            _ = v8
        case "a":
            v9 := GenS2{F0: 2147483647, F1: -2147483648, F2: 0}
            _ = v9
        }
        if true {
            for fi10 := 0; fi10 < 16; fi10++ {
                v7 = append(v7, -2147483648)
                _ = fi10
            }
            v11 := GenS2{F0: 445, F1: 239, F2: 175}
            _ = v11
        }
    } else {
        fmt.Println(p0, v6, p1)
        v6 -= v6
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        if true {
            switch (p0 + p0) {
            case "hello":
                p1 = 2
            case "foo bar":
                v13 := GenS3{F0: 477, F1: ("\n\t" + ((p0 + (("foo bar" + ("a" + p0)) + p0)) + p0)), F2: 2}
                _ = v13
            case "":
                if len(v7) > 0 {
                    v7[0] = 0
                }
            }
        } else {
            if true {
                v6 -= (v6 - -35)
                var v14 bool = true
                var v15 int8 = (^(v6))
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 5; fi16++ {
                    v6 -= 42
                    _ = fi16
                }
                fmt.Println(fi12, fi12, v6)
                p1 += 1
            }
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]int)
    var v3 uint64 = 1
    var v4 float64 = (((0.0 + (3.14159 * 1e10)) / 1.0) + float64((float32(0.0) - 3.14159)))
    var v5 uint64 = (^(v3))
    v1 -= 0
    fmt.Println(v5, v3)
    var v17 uint8 = 42
    v3 -= v5
    fmt.Println(v5, v3)
    var v18 uint32 = 65535
    var v19 bool = false
    v18 = (^(v18))
    var v20 float32 = 1e10
    fmt.Println(v20, v1)
    switch "hello" {
    case "foo bar":
        var v21 uint = 2
        _ = v21
    case "hello":
        v19 = (100 <= 0)
    }
    if v19 {
        fmt.Println(v20, v1, v1)
    } else {
        if v19 {
            if v19 {
                v19 = ((v18 + v18) <= 1)
                var v22 string = "foo bar"
                v23 := []int{32768, int(v5), 32768}
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 9; fi24++ {
                for fi25 := 0; fi25 < 4; fi25++ {
                    v20 -= 0.1
                    var v26 int8 = 30
                    v27 := GenS1{F0: (v1 & v1), F1: -9223372036854775808, F2: (4294967295 >= (v3 - (v5 - v5)))}
                    var v28 bool = (true && v19)
                    var v29 uint32 = 929
                    fmt.Println(v1, fi24)
                    v30 := GenS1{F0: v1, F1: 69, F2: (false || (v1 < (^(v1))))}
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                fmt.Println(v17)
                if (v1 <= -416) {
                    v31 := GenS4{F0: 42, F1: ("a" + ((("\n\t" + "") + ("\n\t" + "\n\t")) + "hello"))}
                    var v32 uint8 = 127
                    _ = v31
                    _ = v32
                }
                _ = fi24
            }
            switch "foo bar" {
            case "foo bar":
                v18 -= (v18 * 1)
            default:
                v18 -= v18
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
