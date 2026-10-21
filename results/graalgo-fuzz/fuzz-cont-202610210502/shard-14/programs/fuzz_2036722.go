package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 bool, p1 uint64) uint {
    p0 = (true || (42 != 2))
    fmt.Println(p1)
    for fi6 := 0; fi6 < 20; fi6++ {
        p1 = p1
        var v7 bool = (!(1e10 == (0.1 / ((3.14159 * 1.5) / (((0.0 / (((0.1 - 1e10) * 0.1) + 0.0)) + 1.5) / 1e10)))))
        v7 = (2147483647 > -2147483648)
        _ = fi6
        _ = v7
    }
    var v8 float32 = 3.14159
    if ((1.5 + 1.5) != (-2.25 + (0.1 * (-2.25 / 1.0)))) {
        v8 = v8
    } else {
        v9 := []int{}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v8
    return 42
}
func genFn1() uint8 {
    var v10 uint64 = 426
    var v11 string = ((("\n\t" + ("" + "hello")) + "\n\t") + "foo bar")
    var v12 int = 0
    fmt.Println(v11)
    _ = v10
    _ = v11
    _ = v12
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 2
    var v3 int16 = 0
    v4 := make(map[string]int)
    v5 := make(map[string]int)
    v13 := GenS1{F0: 1}
    if true {
        for fi14 := 0; fi14 < 8; fi14++ {
            v3 -= (v3 / ((32767) | 1))
            switch fi14 {
            case 0:
                if false {
                    for fi15 := 0; fi15 < 14; fi15++ {
                        var v16 string = ""
                        _ = fi15
                        _ = v16
                    }
                    v3 -= v3
                    var v17 uint32 = 2147483647
                    _ = v17
                } else {
                    switch (-(fi14)) {
                    case 127:
                        fmt.Println(v3, v2)
                    case -75:
                        for fi18 := 0; fi18 < 1; fi18++ {
                            v19 := GenS1{F0: 552}
                            var v20 uint = genFn0((-276 > -1), 4294967295)
                            fmt.Println(v20)
                            var v21 int32 = 945
                            _ = fi18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    case -1:
                        v22 := []int{(fi14 >> ((4294967295) & 63))}
                        _ = v22
                    }
                    var v23 int32 = -2147483648
                    _ = v23
                }
            }
            _ = fi14
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            if false {
                delete(v4, "")
                v25 := []int{4294967295, fi24, fi24}
                v2 -= 255
                _ = v25
            } else {
                v3 = (^(v3))
            }
            _ = fi24
        }
        _ = len(v1)
    }
    v26 := GenS1{F0: 1}
    v27 := v4[("a" + "")]
    var v28 string = "hello"
    fmt.Println(v3, v27, v28)
    v2 -= (v2 | v2)
    var v29 string = ("foo bar" + (((v28 + (("hello" + "\n\t") + (v28 + v28))) + ("\n\t" + (v28 + v28))) + v28))
    v3 -= v3
    var v30 int64 = -9223372036854775808
    v3 = v3
    v31 := v4["\n\t"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
