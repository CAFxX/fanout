package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (!false)
    var v3 float64 = (1.5 * -2.25)
    var v4 bool = (2147483647 <= 0)
    var v5 bool = (2147483647 == 2147483647)
    v6 := make(map[int]string)
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := v6[-1]
        for fi9 := 0; fi9 < 3; fi9++ {
            v1[0] = (((v8 + ("" + v8)) + v8) + (v8 + ""))
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    var v10 int16 = -32768
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 bool = v5
        if true {
            fmt.Println(v4)
            v5 = v2
            v4 = (false && true)
        }
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 42, F1: v5, F2: 468}
    delete(v6, 1)
    if (-9223372036854775808 == 708) {
        v10 -= 1000
        if (65535 >= 42) {
            var v14 uint64 = 42
            v15 := GenS2{F0: uint32(v14), F1: 32768}
            var v16 bool = v2
            _ = v14
            _ = v15
            _ = v16
        } else {
            switch "\n\t" {
            case "a":
                var v17 bool = v2
                _ = v17
            case "":
                if v2 {
                    v10 += (v10 | (v10 * (v10 & (v10 + 63))))
                    for fi18 := 0; fi18 < 20; fi18++ {
                        for fi19 := 0; fi19 < 3; fi19++ {
                            var v20 uint8 = 128
                            v21 := GenS2{F0: 4294967295, F1: (fi19 ^ 2147483647)}
                            v22 := GenS1{F0: fi19, F1: true, F2: 4294967295}
                            v23 := GenS1{F0: 781, F1: ((-14 >= -86) && ((!false) && (v2 || v4))), F2: 0}
                            fmt.Println(v4, v5, v5)
                            _ = fi19
                            _ = v20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        switch (("hello" + "") + ("a" + "")) {
                        case "foo bar":
                            var v24 int32 = 0
                            _ = v24
                        }
                        _ = fi18
                    }
                    fmt.Println(v3, v3)
                } else {
                    fmt.Println(v2, v5, v3)
                    var v25 int32 = 2147483647
                    _ = v25
                }
            }
            v26 := []int{9223372036854775807, 4294967295, -9223372036854775808, -9223372036854775808}
            _ = v26
        }
    }
    if (v4 || (!v2)) {
        v4 = (false || (!(-155 == 1)))
    }
    var v27 uint = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
