package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 bool; F2 int }
type GenS3 struct { F0 int64; F1 uint8; F2 int8 }
func genFn0(p0 int32) uint8 {
    var v5 float64 = ((3.14159 + 1.5) / 0.1)
    v6 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = v5
    _ = v6
    return 20
}

func main() {
    v1 := []int{}
    var v2 float64 = float64(1.5)
    v3 := make(map[string]int)
    var v4 string = "a"
    if true {
        if true {
            if (-93 != -128) {
                v7 := GenS1{F0: -2147483648}
                fmt.Println(v2, v4)
                _ = v7
            } else {
                v2 += 0.1
            }
            var v8 bool = (18446744073709551615 <= 18446744073709551615)
            v9 := GenS1{F0: 32768}
            _ = v8
            _ = v9
        }
    }
    v2 = v2
    if (1 == 0) {
        for fi10 := 0; fi10 < 5; fi10++ {
            if (0 == 2) {
                if true {
                    _ = len(v1)
                }
            } else {
                var v11 int64 = 1
                delete(v3, (v4 + (v4 + ((v4 + v4) + v4))))
                v1 = append(v1, (-(fi10)))
                _ = v11
            }
            v12 := GenS1{F0: fi10}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi10
            _ = v12
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi13
        }
        v14 := GenS1{F0: 32767}
        _ = v14
    } else {
        var v15 uint8 = 47
        if (805 >= -769) {
            _ = len(v3)
            if (538 <= 32767) {
                v2 = v2
                v2 = -2.25
            }
            if (2147483647 > 9223372036854775807) {
                var v16 int64 = int64(v15)
                v17 := GenS1{F0: 32767}
                _ = v16
                _ = v17
            } else {
                v4 = ((v4 + "") + v4)
            }
        }
        _ = v15
    }
    if (v2 < (-2.25 - 1e10)) {
        var v18 uint32 = 283
        if true {
            var v19 uint = 2
            _ = v19
        }
        if len(v1) > 0 {
            v1[0] = 42
        }
        _ = v18
    }
    v1 = append(v1, 1)
    v20 := GenS2{F0: 42, F1: false, F2: 42}
    var v21 bool = true
    v21 = (!v21)
    if v21 {
        v22 := v3[(v4 + ("foo bar" + "\n\t"))]
        for fi23 := 0; fi23 < 9; fi23++ {
            switch v4 {
            case "\n\t":
                v22 += -433
            }
            var v24 int32 = 0
            var v25 uint32 = 1
            _ = fi23
            _ = v24
            _ = v25
        }
        v26 := GenS3{F0: 9223372036854775807, F1: 42, F2: 0}
        _ = v22
        _ = v26
    } else {
        var v27 float64 = v2
        switch "hello" {
        case "\n\t":
            v28 := GenS1{F0: 42}
            _ = v28
        case "a":
            v29 := v3[("a" + v4)]
            _ = v29
        case "":
            var v30 uint32 = 2147483647
            _ = v30
        }
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
