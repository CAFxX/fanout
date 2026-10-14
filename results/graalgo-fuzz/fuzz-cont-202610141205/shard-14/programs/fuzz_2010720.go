package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint8 = 42
    var v2 int32 = -800
    var v3 float64 = (0.0 * 1.5)
    var v4 int = 527
    v3 -= 0.0
    for fi5 := 0; fi5 < 9; fi5++ {
        if true {
            var v6 float32 = (1.5 - (1.5 + -2.25))
            v3 -= -2.25
            v6 += v6
            _ = v6
        } else {
            for fi7 := 0; fi7 < 20; fi7++ {
                v2 = (^(v2))
                _ = fi7
            }
        }
        var v8 bool = false
        v9 := GenS1{F0: -631}
        _ = fi5
        _ = v8
        _ = v9
    }
    v10 := []int{(-(v4)), 9223372036854775807, v4}
    var v11 int16 = 0
    var v12 int = 2147483647
    for fi13 := 0; fi13 < 13; fi13++ {
        if (1 == 324) {
            var v14 float32 = 0.0
            v3 -= (v3 + ((v3 / (1.5 / 0.1)) + (v3 - v3)))
            _ = v14
        } else {
            fmt.Println(v3, v11, v1)
        }
        v12 = int(v1)
        _ = fi13
    }
    if false {
        var v15 float32 = -2.25
        v11 += v11
        if (-552 == 1) {
            for fi16 := 0; fi16 < 17; fi16++ {
                fmt.Println(fi16)
                _ = fi16
            }
            v1 = 125
            v4 = (v12 >> ((352) & 63))
        } else {
            var v17 int8 = 100
            _ = v17
        }
        _ = v15
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 bool = false
        v19 = v19
        _ = fi18
        _ = v19
    }
    v2 = (^(v2))
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v20 int16 = (v11 % ((1) | 1))
    v21 := GenS1{F0: -9223372036854775808}
    switch (v4 ^ v4) {
    case 1:
        if ((4294967295 < 18446744073709551615) && (true || false)) {
            v22 := GenS1{F0: -46}
            var v23 int16 = -841
            _ = v22
            _ = v23
        } else {
            if ((v3 - 1.5) <= 3.14159) {
                var v24 uint = 2
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                var v25 float32 = (3.14159 - -2.25)
                _ = v24
                _ = v25
            }
        }
    case -2147483648:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    case 42:
        v3 = v3
    default:
        if true {
            switch (("hello" + "a") + "foo bar") {
            case "a":
                var v26 string = "\n\t"
                _ = v26
            case "foo bar":
                var v27 uint64 = 42
                _ = v27
            case "":
                v20 = v20
            }
        } else {
            var v28 uint32 = 713
            for fi29 := 0; fi29 < 9; fi29++ {
                v4 += v12
                if (247 >= 677) {
                    var v30 uint64 = 1
                    _ = len(v10)
                    _ = v30
                }
                v1 += (v1 + 255)
                _ = fi29
            }
            _ = v28
        }
    }
    var v31 uint32 = 65535
    v12 -= (v4 << ((1) & 63))
    if true {
        v31 += 1
        var v32 uint64 = 1
        v11 = v11
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v4
}
