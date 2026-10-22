package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int16, p1 int16) uint64 {
    if (!((!true) || false)) {
        p0 -= (^(p0))
        p0 = p1
    }
    v4 := []int{42}
    var v5 string = "hello"
    switch (v5 + "") {
    case "a":
        if (0 != 19) {
            var v6 int64 = -1
            switch 0 {
            case 0:
                p1 += p0
            case 95:
                var v7 float32 = (((float32(1.5) * 1.5) - 0.1) - float32(-2.25))
                _ = v7
            }
            _ = v6
        } else {
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 uint16 = 32767
                _ = fi8
                _ = v9
            }
            v4 = append(v4, -71)
        }
    case "\n\t":
        var v10 int16 = -1
        _ = v10
    case "foo bar":
        if (false && true) {
            v11 := GenS1{F0: 127, F1: 926, F2: 0}
            for fi12 := 0; fi12 < 18; fi12++ {
                p0 -= (p1 / ((42) | 1))
                v13 := GenS1{F0: 128, F1: 65535, F2: int64(p0)}
                fmt.Println(p0, p1, p1)
                _ = fi12
                _ = v13
            }
            _ = v11
        } else {
            fmt.Println(p1)
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint64 = 463
    v2 := []string{}
    v3 := make(map[string]string)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v14 bool = (!(282 != 0))
    v14 = (!v14)
    v15 := GenS1{F0: 1, F1: 63, F2: -9223372036854775808}
    v16 := GenS2{F0: (v1 | 1)}
    v1 -= (^(v1))
    var v17 string = ("a" + ("" + ("" + "hello")))
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 float64 = (float64(1.5) / 1e10)
        _ = fi18
        _ = v19
    }
    var v20 uint64 = 4294967295
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 int16 = 351
            v24 := GenS2{F0: 398}
            v17 = "\n\t"
            var v25 uint64 = 33
            v2 = append(v2, (v17 + v17))
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        if false {
            fmt.Println(v20)
        } else {
            for fi26 := 0; fi26 < 9; fi26++ {
                v14 = v14
                _ = fi26
            }
        }
        _ = fi21
    }
    v27 := v3[""]
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v27
    _ = v3
}
