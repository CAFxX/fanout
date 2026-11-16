package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint, p1 uint32, p2 uint) int16 {
    if true {
        fmt.Println(p0)
        for fi5 := 0; fi5 < 8; fi5++ {
            fmt.Println(p2)
            _ = fi5
        }
    } else {
        p0 += (^(p2))
        if ("\n\t" >= "a") {
            var v6 int = -56
            _ = v6
        }
    }
    switch 1 {
    case 128:
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 float64 = float64((-2.25 / 1.0))
            _ = fi7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 538
}

func main() {
    v1 := []uint32{42, 2147483648}
    v2 := make(map[string]bool)
    var v3 float32 = 0.1
    v4 := []string{}
    v3 -= v3
    switch -905 {
    case 547:
        fmt.Println(v3)
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: true}
        if true {
            if true {
                var v11 int64 = -965
                if false {
                    var v12 uint16 = 0
                    switch "foo bar" {
                    case "a":
                        v3 -= (v3 / v3)
                    case "foo bar":
                        v12 -= 935
                    }
                    if false {
                        v13 := GenS1{F0: true}
                        _ = v13
                    }
                    _ = v12
                }
                v2["\n\t"] = (2 != 42)
                _ = v11
            } else {
                _ = len(v4)
            }
            if false {
                switch (fi9 + 751) {
                case 913:
                    v3 += v3
                case -2147483648:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 1:
                    v3 -= float32(((0.0 - 3.14159) + 1.5))
                }
                if (-1 == 0) {
                    var v14 uint32 = 1
                    _ = v14
                }
            } else {
                v3 += -2.25
                if (true || (!false)) {
                    v3 -= float32((1e10 / 1.0))
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v3 = v3
        } else {
            v3 += v3
        }
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 uint = 475
            v16 = 406
            v17 := GenS1{F0: false}
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi9
        _ = v10
    }
    v3 -= float32(1e10)
    v3 += (float32((1.5 - -2.25)) - 1e10)
    fmt.Println(v3)
    fmt.Println(v3)
    v18 := v2["a"]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
