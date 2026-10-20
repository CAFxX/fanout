package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 string, p1 int32) int32 {
    var v4 int32 = p1
    var v5 float64 = ((((1.5 + ((1e10 / 3.14159) - (3.14159 / 0.1))) - 3.14159) * 1e10) + 0.0)
    v5 = (float64(float32(v5)) - v5)
    switch ("hello" + (("\n\t" + (p0 + p0)) + p0)) {
    case "a":
        if (!((819 == 32767) && false)) {
            var v6 int64 = -1
            _ = v6
        }
    case "hello":
        var v7 int = 823
        _ = v7
    case "":
        fmt.Println(v4)
    }
    v8 := []int{-9223372036854775808, 32767, -9223372036854775808, 32768}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v8
    return -293
}

func main() {
    v1 := []string{"hello", "\n\t", ("hello" + "a"), "a"}
    var v2 int32 = 265
    var v3 uint32 = 42
    var v9 float64 = 3.14159
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    v3 = (v3 - 420)
    fmt.Println(v9, v3, v9)
    v3 += (v3 ^ 2147483647)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 -= (^(v3))
    v1 = append(v1, "a")
    if true {
        v9 = 0.1
        if true {
            fmt.Println(v2)
            if ((float64(-2.25) * v9) <= (v9 - (v9 / (v9 + 0.0)))) {
                v1 = append(v1, ("a" + ((("hello" + "foo bar") + ("\n\t" + ((("foo bar" + "\n\t") + "hello") + ("foo bar" + "hello")))) + "")))
                v10 := GenS1{F0: 0}
                _ = v10
            } else {
                var v11 uint = 2
                switch -2147483648 {
                case -1:
                    v9 = v9
                case 4294967295:
                    for fi12 := 0; fi12 < 16; fi12++ {
                        var v13 int8 = 42
                        _ = len(v1)
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        _ = fi12
                        _ = v13
                    }
                }
                _ = v11
            }
        } else {
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint64 = 0
                v15 = 4294967295
                _ = fi14
                _ = v15
            }
        }
        if (85 < 255) {
            if ((!false) || true) {
                if true {
                    var v16 int = -210
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v16
                } else {
                    v1 = append(v1, "foo bar")
                }
            }
        }
    }
    var v17 int16 = 42
    v17 -= (v17 & -1)
    if true {
        if false {
            v3 -= v3
        } else {
            switch ("\n\t" + (("hello" + "") + "\n\t")) {
            case "a":
                v18 := GenS1{F0: -9223372036854775808}
                _ = v18
            case "\n\t":
                if false {
                    var v19 bool = (true || true)
                    if false {
                        v2 = 42
                        switch "" {
                        case "foo bar":
                            v20 := GenS1{F0: 769}
                            _ = v20
                        }
                        v19 = v19
                    }
                    v17 += (v17 & (v17 | v17))
                    _ = v19
                } else {
                    for fi21 := 0; fi21 < 5; fi21++ {
                        var v22 int = (fi21 << ((255) & 63))
                        fmt.Println(v2, v9, v3)
                        _ = fi21
                        _ = v22
                    }
                    if (9223372036854775807 != 42) {
                        v17 += (v17 / (((^(v17))) | 1))
                    }
                    switch 97 {
                    case -2147483648:
                        v17 += 332
                    case -9223372036854775808:
                        switch (("a" + "") + "\n\t") {
                        case "\n\t":
                            v17 += (v17 + v17)
                        default:
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                        }
                    }
                }
            default:
                v23 := GenS1{F0: -9223372036854775808}
                _ = v23
            }
        }
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
