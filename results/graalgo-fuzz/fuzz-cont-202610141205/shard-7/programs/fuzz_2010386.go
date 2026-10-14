package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint8 }

func main() {
    var v1 int32 = -623
    v2 := make(map[string]int)
    var v3 uint8 = 198
    var v4 uint32 = 0
    v5 := []string{("" + "\n\t"), "a"}
    var v6 bool = true
    v6 = false
    if (int32(v3) > (v1 - 0)) {
        var v7 uint8 = v3
        v7 -= v3
        v3 -= v7
        _ = v7
    } else {
        var v8 uint32 = v4
        v6 = v6
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            if ((v4 | v4) >= (v4 % (((v4 << ((74) & 63))) | 1))) {
                v3 -= (v3 << ((1) & 63))
                switch fi9 {
                case 42:
                    v11 := GenS1{F0: v6, F1: -121, F2: v3}
                    _ = v11
                case 128:
                    _ = len(v5)
                case 0:
                    v4 = (v4 | 2147483648)
                }
                v12 := GenS1{F0: false, F1: 42, F2: v3}
                v1 = v1
                var v13 int16 = -1
                _ = v12
                _ = v13
            } else {
                var v14 uint8 = (v3 % (((v3 & (v3 + 42))) | 1))
                v1 = v1
                if (1 >= 613) {
                    v1 -= (v1 - v1)
                    switch 2147483647 {
                    case -2147483648:
                        var v15 int32 = v1
                        _ = v15
                    }
                }
                _ = v14
            }
            var v16 int16 = -937
            _ = fi10
            _ = v16
        }
        var v17 uint8 = v3
        for fi18 := 0; fi18 < 4; fi18++ {
            v17 += v17
            _ = fi18
        }
        _ = fi9
        _ = v17
    }
    var v19 int = -2147483648
    v2[("\n\t" + "hello")] = (v19 * -76)
    if ((0.1 + -2.25) <= 1.5) {
        if false {
            delete(v2, ("foo bar" + "\n\t"))
            var v20 int64 = 9223372036854775807
            _ = v20
        }
    } else {
        var v21 int64 = 0
        fmt.Println(v4, v19, v3)
        for fi22 := 0; fi22 < 13; fi22++ {
            for fi23 := 0; fi23 < 13; fi23++ {
                v24 := GenS1{F0: (v6 || (int8(v21) != 62)), F1: -1, F2: v3}
                var v25 string = (("hello" + "a") + "a")
                v2[v25] = v19
                v26 := GenS1{F0: v6, F1: 94, F2: (^(v3))}
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi22
        }
        _ = v21
    }
    _ = len(v5)
    var v27 int64 = 42
    fmt.Println(v4)
    fmt.Println(v3, v4, v4)
    v28 := GenS1{F0: true, F1: 127, F2: v3}
    var v29 uint16 = 1
    switch v19 {
    case 264:
        v30 := v2["foo bar"]
        _ = v30
    case -9223372036854775808:
        for fi31 := 0; fi31 < 17; fi31++ {
            v27 = (^(v27))
            if v6 {
                var v32 float64 = 3.14159
                fmt.Println(v32, v19)
                _ = v32
            }
            _ = fi31
        }
    }
    fmt.Println(v3, v4)
    v6 = v6
    v4 = v4
    var v33 int64 = v27
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
