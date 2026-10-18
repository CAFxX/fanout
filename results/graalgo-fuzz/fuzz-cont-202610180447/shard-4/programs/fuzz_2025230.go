package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
func genFn0(p0 int8, p1 uint32, p2 uint64) string {
    p1 -= uint32(p2)
    switch "" {
    case "hello":
        p2 += 0
    case "\n\t":
        if false {
            v6 := []int{}
            switch 9223372036854775807 {
            case -1:
                var v7 bool = false
                _ = v7
            default:
                p0 += 127
            }
            _ = v6
        }
    case "a":
        var v8 uint32 = (p1 & (p1 - p1))
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return ""
}
func genFn1(p0 string) int16 {
    v9 := []int{657, 71, 9223372036854775807, -732}
    v10 := GenS1{F0: 32768}
    if len(v9) > 0 {
        v9[0] = 32768
    }
    if false {
        var v11 uint = 2
        _ = v11
    }
    _ = p0
    _ = v10
    _ = v9
    return -32768
}

func main() {
    var v1 float64 = 1.5
    var v2 int32 = -733
    v3 := make(map[string]string)
    var v4 bool = (2 == 488)
    var v5 int32 = (v2 | (v2 - v2))
    v1 = 0.0
    v2 = v2
    v5 -= 2147483647
    if v4 {
        for fi12 := 0; fi12 < 19; fi12++ {
            v4 = true
            _ = fi12
        }
    }
    v13 := GenS1{F0: 861}
    fmt.Println(v4, v1)
    fmt.Println(v4, v4, v5)
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            v1 = (v1 - v1)
            fmt.Println(v4)
            v4 = false
            _ = fi15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 uint = 728
            if v4 {
                v18 := GenS1{F0: (fi14 | fi14)}
                var v19 string = "a"
                v20 := GenS1{F0: fi14}
                _ = v18
                _ = v19
                _ = v20
            } else {
                v21 := GenS1{F0: 128}
                if (2 == uint16(v17)) {
                    v22 := GenS1{F0: (-(fi16))}
                    v2 += (v5 + (v5 ^ v2))
                    _ = v22
                }
                v23 := v3["a"]
                _ = v21
                _ = v23
            }
            _ = fi16
            _ = v17
        }
        _ = fi14
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := []int{127, fi24}
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 int16 = -350
            v28 := GenS2{F0: "a"}
            if (!((0 < 166) && true)) {
                var v29 int64 = 0
                switch (fi26 % (((fi24 % ((32767) | 1))) | 1)) {
                case 1:
                    v30 := GenS2{F0: (("" + ("hello" + "hello")) + "foo bar")}
                    _ = v30
                case 4294967295:
                    v31 := v3[((((("foo bar" + "") + "\n\t") + "foo bar") + "foo bar") + "a")]
                    _ = v31
                default:
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                }
                fmt.Println(v2, v5, v5)
                _ = v29
            }
            var v32 int8 = 100
            if (fi24 < fi26) {
                v33 := v3["foo bar"]
                _ = v33
            } else {
                if (!((2 <= 1) || false)) {
                    switch "a" {
                    case "hello":
                        var v34 int8 = (v32 + v32)
                        _ = v34
                    }
                }
                v1 -= 0.1
            }
            _ = fi26
            _ = v27
            _ = v28
            _ = v32
        }
        _ = fi24
        _ = v25
    }
    if (-341 <= 127) {
        for fi35 := 0; fi35 < 6; fi35++ {
            if (!v4) {
                v2 -= 838
                var v36 int = (fi35 | (fi35 % ((fi35) | 1)))
                _ = v36
            }
            for fi37 := 0; fi37 < 9; fi37++ {
                var v38 string = "hello"
                delete(v3, ("\n\t" + v38))
                v3[("" + v38)] = (("foo bar" + ((((v38 + (v38 + ("" + v38))) + ("foo bar" + v38)) + "a") + v38)) + (v38 + v38))
                _ = fi37
                _ = v38
            }
            _ = fi35
        }
    }
    v2 -= (v5 & 488)
    v2 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
