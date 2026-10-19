package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int64 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint32; F1 int8 }
func genFn0(p0 bool, p1 uint16) uint8 {
    if true {
        v7 := GenS1{F0: -1, F1: 1, F2: 1}
        _ = v7
    } else {
        if (!((!("a" <= ("a" + "\n\t"))) || true)) {
            fmt.Println(p0, p1)
            var v8 uint8 = 255
            var v9 uint = 42
            _ = v8
            _ = v9
        }
        v10 := []int{127, 32767}
        switch -474 {
        case -9223372036854775808:
            for fi11 := 0; fi11 < 4; fi11++ {
                p1 -= (p1 ^ p1)
                var v12 int16 = -32768
                _ = fi11
                _ = v12
            }
        case -2147483648:
            var v13 uint16 = p1
            _ = v13
        case 4294967295:
            switch 46 {
            case 2147483647:
                v14 := GenS2{F0: 42, F1: 100}
                _ = v14
            }
        }
        _ = v10
    }
    if p0 {
        var v15 int32 = 42
        switch ("a" + "") {
        case "\n\t":
            p1 -= 966
        case "foo bar":
            fmt.Println(v15, v15)
        case "a":
            v16 := []int{26}
            _ = v16
        }
        v17 := GenS2{F0: (^(v15)), F1: 42}
        _ = v15
        _ = v17
    } else {
        var v18 uint16 = (p1 - (p1 & (p1 + 32767)))
        _ = v18
    }
    var v19 int = 4294967295
    v19 = v19
    var v20 uint = 889
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 0
}

func main() {
    v1 := []int{}
    var v2 float64 = ((3.14159 * (0.0 - 1e10)) - (1e10 * ((0.0 - 3.14159) + 1.5)))
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    var v5 float32 = (((-2.25 + float32(0.0)) / 1e10) * ((float32(v2) * 1e10) * (3.14159 * 0.1)))
    v6 := make(map[string]int)
    var v21 bool = (false && false)
    switch "hello" {
    case "foo bar":
        if ((-386 > -68) && (!(v21 && v21))) {
            v5 = v5
            for fi22 := 0; fi22 < 8; fi22++ {
                v23 := GenS3{F0: 42, F1: 1}
                switch (fi22 ^ 112) {
                case 551:
                    v21 = v21
                }
                v2 -= float64(float32(3.14159))
                _ = fi22
                _ = v23
            }
        }
    }
    v21 = v21
    var v24 uint64 = 876
    switch -55 {
    case 2147483647:
        v2 -= 1.5
    }
    fmt.Println(v2, v2, v24)
    v25 := v3[""]
    for fi26 := 0; fi26 < 11; fi26++ {
        fmt.Println(v25, v25)
        _ = fi26
    }
    delete(v4, v25)
    for fi27 := 0; fi27 < 3; fi27++ {
        var v28 int8 = 100
        for fi29 := 0; fi29 < 17; fi29++ {
            v21 = ((2 <= 0) || false)
            _ = len(v1)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (0.0 <= v5) {
                switch ((v25 + ("a" + v25)) + (("hello" + v25) + "")) {
                case "foo bar":
                    var v30 uint = 1
                    _ = v30
                case "\n\t":
                    v25 = ((v25 + "hello") + v25)
                case "hello":
                    var v31 float64 = (0.1 + v2)
                    _ = v31
                }
                var v32 int8 = 42
                _ = v32
            }
            _ = fi29
        }
        _ = fi27
        _ = v28
    }
    v5 = float32(v2)
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
