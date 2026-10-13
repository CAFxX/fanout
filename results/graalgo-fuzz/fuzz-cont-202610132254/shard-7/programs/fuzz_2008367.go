package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int16; F2 int64 }
func genFn0() string {
    var v4 uint64 = 4294967295
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int64 = 1
        var v7 uint64 = v4
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v4)
    v4 -= 1
    _ = v4
    return "foo bar"
}
func genFn1(p0 uint, p1 string) int64 {
    v8 := []int{32768, 42}
    var v9 int = 617
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 183
}

func main() {
    v1 := []string{"a", "hello", (("\n\t" + "") + "foo bar"), "hello"}
    var v2 float64 = float64((((float32((3.14159 - (3.14159 * 0.1))) / 1.0) - (3.14159 - float32(float64(-2.25)))) + 1e10))
    var v3 uint = 1
    var v10 string = "foo bar"
    if (361 <= 1) {
        fmt.Println(v3, v3)
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS1{F0: 0, F1: uint8(v3)}
        _ = fi11
        _ = v12
    }
    if (1 >= 9223372036854775807) {
        fmt.Println(v3, v2)
    } else {
        var v13 uint16 = 65535
        if (255 != 138) {
            switch v10 {
            case "a":
                v10 = v10
            case "foo bar":
                if (!(!true)) {
                    var v14 int8 = 1
                    _ = v14
                }
            case "\n\t":
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := GenS1{F0: -80, F1: uint8(v3)}
                    var v17 int = (fi15 + 9223372036854775807)
                    if ((!true) || true) {
                        v2 = float64(((1.5 / 1.0) - 3.14159))
                        switch "hello" {
                        case "\n\t":
                            var v18 uint8 = 128
                            _ = v18
                        case "foo bar":
                            _ = len(v1)
                        }
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            if (v3 < v3) {
                v19 := GenS2{F0: 668, F1: -32768, F2: genFn1(uint(v13), ((v10 + ("a" + "\n\t")) + (v10 + (v10 + v10))))}
                _ = v19
            }
        } else {
            v20 := GenS1{F0: 319, F1: 128}
            var v21 string = "a"
            v22 := GenS1{F0: 1000, F1: 42}
            _ = v20
            _ = v21
            _ = v22
        }
        var v23 string = ""
        _ = v13
        _ = v23
    }
    v24 := GenS1{F0: 836, F1: 0}
    v10 = ((v10 + "\n\t") + "")
    if (248 < 229) {
        fmt.Println(v10, v2, v3)
        if true {
            v25 := GenS1{F0: -926, F1: 127}
            _ = v25
        } else {
            var v26 int = 128
            _ = v26
        }
    }
    v10 = v10
    v1 = append(v1, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
}
