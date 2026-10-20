package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
func genFn0(p0 int, p1 int32, p2 string) uint8 {
    if false {
        fmt.Println(p0, p0, p1)
        for fi5 := 0; fi5 < 14; fi5++ {
            p0 += (-(fi5))
            _ = fi5
        }
    } else {
        var v6 int = p0
        switch v6 {
        case 835:
            v7 := []int{(p0 % ((v6) | 1))}
            _ = v7
        case -353:
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := []int{(v6 | -9223372036854775808), (p0 >> ((65535) & 63)), p0, (p0 >> ((969) & 63))}
                if false {
                    var v10 string = p2
                    var v11 uint16 = 43
                    _ = v10
                    _ = v11
                }
                for fi12 := 0; fi12 < 13; fi12++ {
                    p0 = -206
                    _ = fi12
                }
                _ = fi8
                _ = v9
            }
        case 0:
            fmt.Println(p1)
        }
        _ = v6
    }
    var v13 uint = 723
    switch (p2 + p2) {
    case "":
        v14 := []int{p0, (p0 & 0), (p0 % ((p0) | 1))}
        _ = v14
    case "a":
        var v15 bool = (1e10 == -2.25)
        _ = v15
    }
    var v16 uint = (v13 >> ((v13) & 63))
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v16
    return 255
}

func main() {
    v1 := []uint32{}
    v2 := make(map[string]uint32)
    var v3 float64 = (float64(-2.25) / 1e10)
    var v4 uint = 1
    v1 = append(v1, 296)
    var v17 int16 = 32767
    if true {
        var v18 int16 = int16(v4)
        v19 := GenS1{F0: 127, F1: (v17 / ((v17) | 1)), F2: (false || false)}
        for fi20 := 0; fi20 < 15; fi20++ {
            fmt.Println(v18)
            switch "hello" {
            case "a":
                v4 = v4
            }
            _ = fi20
        }
        _ = v18
        _ = v19
    } else {
        if false {
            var v21 int16 = (-(v17))
            _ = v21
        }
        var v22 uint32 = 1
        _ = v22
    }
    v1 = append(v1, 4294967295)
    if (!false) {
        var v23 int32 = -2147483648
        _ = v23
    }
    var v24 int64 = 1
    var v25 float64 = v3
    var v26 int64 = (v24 + v24)
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 bool = false
        fmt.Println(v25, fi27, v26)
        if v28 {
            v29 := GenS1{F0: -127, F1: v17, F2: (!v28)}
            _ = v29
        }
        _ = fi27
        _ = v28
    }
    if false {
        if true {
            v4 = 2
            v1 = append(v1, 342)
        }
    }
    v26 += -896
    var v30 int8 = 0
    if false {
        var v31 uint32 = 42
        _ = v31
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        v17 -= -32768
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi33 := 0; fi33 < 5; fi33++ {
            v3 += v3
            if (751 == (^(v17))) {
                v30 = v30
            } else {
                if (0 >= 2147483647) {
                    v25 -= (((((3.14159 * (3.14159 + v3)) / v3) / (0.1 / (v3 * v25))) + v3) / -2.25)
                    fmt.Println(v24)
                }
                var v34 int8 = (v30 >> ((551) & 63))
                _ = len(v1)
                _ = v34
            }
            fmt.Println(v3, fi32, v25)
            _ = fi33
        }
        _ = fi32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
}
