package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16; F1 string }
func genFn0(p0 uint32) int {
    p0 += 1
    p0 = p0
    var v7 uint32 = p0
    v8 := []int{}
    v9 := GenS1{F0: 0, F1: (p0 ^ (^(v7))), F2: 42}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 167
}

func main() {
    var v1 bool = (!(32767 >= 42))
    var v2 int64 = -9223372036854775808
    var v3 int16 = 1
    var v4 int = -368
    var v5 float64 = (1e10 * 0.0)
    var v6 float32 = 1.5
    v4 = genFn0(717)
    v3 -= (v3 - v3)
    switch (("" + "") + "") {
    case "hello":
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 uint = 4294967295
            var v12 uint = v11
            _ = fi10
            _ = v11
            _ = v12
        }
    case "foo bar":
        var v13 uint32 = 2147483648
        _ = v13
    }
    if (165 <= 42) {
        switch ("\n\t" + ("foo bar" + "")) {
        case "foo bar":
            fmt.Println(v2, v2)
        }
        if (0 != 822) {
            var v14 int8 = 115
            v15 := []int{(v4 & (v4 * v4)), genFn0(2), (v4 << ((897) & 63))}
            for fi16 := 0; fi16 < 8; fi16++ {
                v17 := GenS1{F0: 953, F1: 624, F2: -128}
                for fi18 := 0; fi18 < 4; fi18++ {
                    _ = len(v15)
                    v14 = -128
                    v19 := GenS1{F0: 18446744073709551615, F1: 2, F2: (^(v14))}
                    _ = fi18
                    _ = v19
                }
                var v20 float64 = v5
                _ = fi16
                _ = v17
                _ = v20
            }
            _ = v14
            _ = v15
        }
        v21 := []int{v4, (v4 << ((881) & 63)), (-(v4)), genFn0(487)}
        _ = v21
    }
    v5 = v5
    v22 := []int{(-(v4)), (v4 ^ -618), 32767}
    if v1 {
        v2 = (v2 + v2)
        for fi23 := 0; fi23 < 2; fi23++ {
            v5 = 1.5
            _ = fi23
        }
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 int8 = -128
            var v26 uint = 366
            _ = fi24
            _ = v25
            _ = v26
        }
    } else {
        for fi27 := 0; fi27 < 6; fi27++ {
            if v1 {
                for fi28 := 0; fi28 < 19; fi28++ {
                    v22 = append(v22, genFn0(294))
                    _ = fi28
                }
                v2 = v2
                v1 = true
            } else {
                for fi29 := 0; fi29 < 14; fi29++ {
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    fmt.Println(v3)
                    v30 := GenS2{F0: (v3 << ((1) & 63)), F1: ""}
                    _ = fi29
                    _ = v30
                }
            }
            switch "\n\t" {
            case "hello":
                if v1 {
                    _ = len(v22)
                }
            case "\n\t":
                v6 = -2.25
            }
            v31 := GenS2{F0: (^(v3)), F1: (("hello" + ("" + "a")) + ("\n\t" + "a"))}
            _ = fi27
            _ = v31
        }
    }
    v32 := GenS2{F0: 1, F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
