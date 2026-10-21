package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int16, p1 int64) uint64 {
    fmt.Println(p0)
    switch "foo bar" {
    case "":
        v6 := []int{-86, 32768}
        _ = v6
    case "a":
        p0 -= p0
    }
    v7 := GenS1{F0: 42}
    p1 -= p1
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    var v4 float32 = -2.25
    var v5 int16 = 52
    if (true && true) {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint32 = 42
            var v10 int64 = 1
            if false {
                switch ("foo bar" + "\n\t") {
                case "a":
                    v11 := v3["hello"]
                    _ = v11
                case "foo bar":
                    var v12 int8 = 110
                    _ = v12
                default:
                    v5 -= 282
                }
                var v13 float32 = v4
                switch ("a" + "") {
                case "hello":
                    if ((!(!false)) || (0.1 == 3.14159)) {
                        var v14 int32 = -2147483648
                        _ = v14
                    }
                case "foo bar":
                    fmt.Println(v9, v9)
                }
                _ = v13
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        switch -274 {
        case -881:
            v4 += (v4 - 1.5)
        }
        var v15 int32 = -398
        _ = v15
    }
    v4 -= -2.25
    if false {
        if (65535 >= 501) {
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 uint8 = 49
                v4 = -2.25
                var v18 int32 = 1
                _ = fi16
                _ = v17
                _ = v18
            }
        } else {
            v5 -= (v5 ^ (v5 * v5))
            if (0 >= 367) {
                var v19 float32 = float32(1e10)
                var v20 uint16 = 255
                _ = v19
                _ = v20
            }
            v4 -= ((v4 + float32(0.1)) - v4)
        }
        _ = len(v2)
    }
    if (!(42 <= 0)) {
        v21 := []int{868}
        _ = v21
    } else {
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 uint8 = 127
            v24 := GenS1{F0: 120}
            switch (((("a" + "hello") + "a") + "hello") + "hello") {
            case "\n\t":
                var v25 uint8 = v23
                _ = v25
            case "a":
                var v26 uint16 = 1
                _ = v26
            case "foo bar":
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 float64 = float64(3.14159)
                    _ = fi27
                    _ = v28
                }
            }
            _ = fi22
            _ = v23
            _ = v24
        }
        switch 32768 {
        case 4294967295:
            fmt.Println(v5)
        case -2147483648:
            v5 += -988
        }
        v4 = -2.25
    }
    var v29 int32 = 1
    var v30 int8 = 104
    v31 := v3[""]
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
