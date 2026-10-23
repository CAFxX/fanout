package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int, p1 int) int16 {
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := []int{fi8, fi8, 568, p0}
            var v10 float64 = 0.1
            if false {
                var v11 uint = 42
                if len(v9) > 0 {
                    v9[0] = (p0 * fi7)
                }
                _ = v11
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        var v12 float64 = 1.5
        _ = fi7
        _ = v12
    }
    var v13 float32 = 3.14159
    _ = p0
    _ = p1
    _ = v13
    return -1
}
func genFn1(p0 int16, p1 int8, p2 uint8) int8 {
    var v14 uint16 = 1
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS1{F0: (((("a" + ("" + "")) + "\n\t") + ((("" + "a") + "") + ("a" + "foo bar"))) + "foo bar")}
        _ = fi15
        _ = v16
    }
    v17 := []int{1, 128, 0, 508}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v17
    return (^(p1))
}

func main() {
    v1 := []int{771, -9223372036854775808, 2147483647}
    v2 := []uint32{74, 198, 2, 42}
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    var v5 uint8 = 127
    var v6 int8 = 25
    v18 := GenS1{F0: "\n\t"}
    v5 = 42
    for fi19 := 0; fi19 < 5; fi19++ {
        if (!(-745 >= -393)) {
            if (!((("a" + "foo bar") + "\n\t") != ("a" + ("hello" + "a")))) {
                v20 := v4["\n\t"]
                for fi21 := 0; fi21 < 2; fi21++ {
                    v4[((v20 + v20) + v20)] = v20
                    var v22 uint64 = 4294967295
                    _ = fi21
                    _ = v22
                }
                var v23 string = (v20 + v20)
                _ = v20
                _ = v23
            } else {
                delete(v3, "")
            }
            _ = len(v3)
            v24 := GenS1{F0: (("\n\t" + "hello") + "\n\t")}
            _ = v24
        }
        v6 = v6
        _ = fi19
    }
    v6 = (v6 % ((-99) | 1))
    v5 += v5
    var v25 int = 312
    fmt.Println(v6, v25)
    v26 := GenS1{F0: "foo bar"}
    if (42 != (v5 ^ v5)) {
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 uint16 = 1
            _ = len(v4)
            _ = fi27
            _ = v28
        }
        v6 -= (v6 & v6)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
