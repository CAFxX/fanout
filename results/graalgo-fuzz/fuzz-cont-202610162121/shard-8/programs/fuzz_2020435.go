package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int8, p1 int, p2 uint8) bool {
    p2 += p2
    p1 += 1
    v7 := []int{-2147483648}
    v8 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    var v4 float64 = ((((3.14159 + 0.1) - (3.14159 / ((1.5 + 0.0) / 1.5))) / 1.0) + 0.0)
    v5 := []int{0, -1}
    var v6 uint = 65535
    var v9 float64 = ((0.1 / (float64((1e10 / float32(v4))) / 0.1)) - (0.0 * v4))
    _ = len(v5)
    v10 := v3[65535]
    for fi11 := 0; fi11 < 2; fi11++ {
        if (("a" + "foo bar") == "\n\t") {
            switch (v10 - v10) {
            case 9223372036854775807:
                fmt.Println(v1)
            case 614:
                var v12 int64 = 9223372036854775807
                _ = v12
            case 0:
                for fi13 := 0; fi13 < 9; fi13++ {
                    switch fi13 {
                    case -9223372036854775808:
                        var v14 int16 = -32768
                        _ = v14
                    default:
                        if (170 == 0) {
                            fmt.Println(fi13, v10)
                        }
                    }
                    _ = fi13
                }
            }
        } else {
            v9 = (v9 - (v9 / (0.0 - (v4 + v4))))
        }
        _ = fi11
    }
    switch 4294967295 {
    case 0:
        delete(v3, 1)
    case -2147483648:
        for fi15 := 0; fi15 < 14; fi15++ {
            v1 -= v1
            switch "a" {
            case "a":
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := GenS1{F0: 65535}
                    _ = fi16
                    _ = v17
                }
            }
            v18 := GenS1{F0: 65535}
            _ = fi15
            _ = v18
        }
    default:
        switch (v10 / ((v10) | 1)) {
        case -9223372036854775808:
            for fi19 := 0; fi19 < 17; fi19++ {
                fmt.Println(v1, v9)
                v9 = 1e10
                _ = fi19
            }
        case -640:
            fmt.Println(v10, v6)
        case 9223372036854775807:
            fmt.Println(v10, v6)
        }
    }
    if (-2.25 > 1e10) {
        fmt.Println(v9)
        v4 += v4
    } else {
        switch "hello" {
        case "":
            var v20 int = v10
            _ = v20
        default:
            switch ((("hello" + "hello") + ("a" + (("foo bar" + "foo bar") + "a"))) + ((("foo bar" + "foo bar") + ("a" + ("hello" + ((("a" + "hello") + (("foo bar" + "\n\t") + ("hello" + "foo bar"))) + ("" + ("\n\t" + "foo bar")))))) + ("foo bar" + ""))) {
            case "\n\t":
                var v21 int = (v10 << ((42) & 63))
                _ = v21
            case "":
                var v22 int8 = 0
                _ = v22
            default:
                var v23 int32 = v1
                _ = v23
            }
        }
    }
    if ((^(v6)) != (^(v6))) {
        var v24 uint8 = 0
        if true {
            if (3.14159 <= (-2.25 * -2.25)) {
                var v25 int = 32767
                _ = v25
            }
            _ = len(v2)
            v1 = (v1 << ((v6) & 63))
        } else {
            if true {
                var v26 uint16 = 32767
                _ = v26
            } else {
                v1 -= (v1 % ((-163) | 1))
                var v27 int8 = 42
                _ = v27
            }
        }
        _ = v24
    }
    var v28 string = "\n\t"
    v1 = 0
    v28 = v28
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
