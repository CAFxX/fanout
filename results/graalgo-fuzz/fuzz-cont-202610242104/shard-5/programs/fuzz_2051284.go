package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint8, p1 int) int8 {
    v4 := GenS1{F0: "foo bar"}
    v5 := []int{(p1 ^ (p1 | (p1 ^ int(p0)))), (p1 & p1), -45}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 bool) int16 {
    var v6 uint16 = 313
    fmt.Println(p0)
    if (!(-9223372036854775808 != 42)) {
        var v7 int64 = 42
        _ = v7
    } else {
        var v8 float32 = float32((0.0 + (0.0 / 1.0)))
        fmt.Println(p0, v8, p0)
        _ = v8
    }
    _ = p0
    _ = v6
    return -14
}

func main() {
    var v1 uint16 = 42
    var v2 float64 = (float64(1.5) - (((0.1 + float64(0.0)) / 1e10) / 1e10))
    v3 := []uint32{2147483648, 65535, 0, 950}
    switch (("" + "hello") + "foo bar") {
    case "\n\t":
        v2 -= (0.0 + (v2 / float64((0.0 - -2.25))))
    default:
        if true {
            var v9 float32 = ((float32((v2 * float64(0.1))) + (0.0 - -2.25)) * (((1.5 - (0.1 - (((0.0 + 1.5) / ((1.5 * 0.0) - 0.1)) * (float32(v2) + -2.25)))) / ((-2.25 / (1.5 * 0.1)) + 1.5)) - 1e10))
            v10 := GenS1{F0: "hello"}
            v9 += ((v9 * float32(v2)) / v9)
            _ = v10
            _ = v9
        } else {
            v2 = v2
            v2 = v2
            v11 := GenS1{F0: "hello"}
            _ = v11
        }
    }
    var v12 int32 = -918
    fmt.Println(v12, v2, v2)
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := GenS1{F0: "foo bar"}
        _ = fi13
        _ = v14
    }
    fmt.Println(v12, v1, v1)
    if ((^(v12)) == (v12 - (v12 - (v12 << ((65535) & 63))))) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v15 float32 = 3.14159
        _ = v15
    }
    switch (((("hello" + "\n\t") + "hello") + "\n\t") + (("" + "foo bar") + "")) {
    case "a":
        switch "a" {
        case "\n\t":
            var v16 uint16 = v1
            _ = v16
        case "":
            v17 := GenS1{F0: ("a" + ((("" + ("\n\t" + "a")) + "foo bar") + "hello"))}
            _ = v17
        }
    case "":
        for fi18 := 0; fi18 < 16; fi18++ {
            for fi19 := 0; fi19 < 8; fi19++ {
                if false {
                    v1 -= v1
                    var v20 uint16 = v1
                    v20 -= 307
                    _ = v20
                }
                _ = fi19
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                switch "hello" {
                case "a":
                    fmt.Println(fi18)
                case "foo bar":
                    v22 := GenS1{F0: "foo bar"}
                    _ = v22
                default:
                    switch "hello" {
                    case "a":
                        var v23 bool = true
                        _ = v23
                    }
                }
                _ = fi21
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi18
        }
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 uint16 = 291
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
