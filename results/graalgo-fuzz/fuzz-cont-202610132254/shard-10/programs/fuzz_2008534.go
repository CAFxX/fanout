package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() string {
    var v5 float64 = -2.25
    v5 += (v5 + v5)
    _ = v5
    return ""
}
func genFn1() int {
    for fi6 := 0; fi6 < 3; fi6++ {
        if (!false) {
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 float32 = (-2.25 * 3.14159)
                var v9 string = ""
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        var v10 uint = 18446744073709551615
        _ = fi6
        _ = v10
    }
    var v11 int = -164
    v11 = 2147483647
    _ = v11
    return (v11 << ((42) & 63))
}

func main() {
    var v1 float64 = float64(0.0)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    var v4 uint = 65535
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(fi12, v4, v4)
        v13 := GenS1{F0: true}
        switch (fi12 & 32767) {
        case -414:
            v14 := GenS1{F0: (0 != 497)}
            _ = v14
        case 1:
            var v15 int16 = 32767
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v4 += (v4 * 0)
    v1 = (v1 - v1)
    v4 = v4
    v4 = (v4 - v4)
    v16 := v2[90]
    var v17 uint = (v4 & (v4 + v4))
    v18 := v3[(v16 % (((^(v16))) | 1))]
    if (!(true || false)) {
        if true {
            for fi19 := 0; fi19 < 18; fi19++ {
                fmt.Println(v4, v17)
                v16 = 42
                _ = fi19
            }
        }
        v17 += (v17 + v4)
        v4 = (v17 - (v17 / ((v4) | 1)))
    } else {
        v20 := GenS1{F0: true}
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        fmt.Println(v18)
        fmt.Println(v17, v4)
        switch (fi21 | -2147483648) {
        case 32767:
            switch v18 {
            case "a":
                if (true || (!((0.0 / 0.1) >= 1e10))) {
                    fmt.Println(v17, v17)
                }
            case "hello":
                if (-1 <= 2147483647) {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v18 = (v18 + v18)
                        _ = fi22
                    }
                }
            }
        }
        _ = fi21
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 int64 = -9223372036854775808
            v17 += (v4 << (((v4 * v17)) & 63))
            v4 -= 255
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
