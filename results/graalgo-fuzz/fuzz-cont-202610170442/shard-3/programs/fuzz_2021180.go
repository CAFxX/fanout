package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 bool }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8) int16 {
    var v7 uint8 = 42
    v8 := []int{-1, 2147483647, -2147483648}
    p0 = (p0 * (v7 ^ 42))
    _ = p0
    _ = v7
    _ = v8
    return -969
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 42
    v3 := make(map[string]uint32)
    var v4 float64 = 1e10
    v5 := make(map[uint32]uint32)
    var v6 int16 = -665
    var v9 int = -96
    v9 = v9
    var v10 int32 = 0
    switch "" {
    case "\n\t":
        var v11 int32 = (v2 - v2)
        _ = v11
    case "":
        switch (v9 * v9) {
        case -1:
            for fi12 := 0; fi12 < 12; fi12++ {
                if ((0 > 9223372036854775807) && (!((v4 <= ((((v4 - v4) * v4) / (v4 + (((v4 + v4) + ((v4 - 1e10) - (float64(((0.1 - 1e10) + (float32((3.14159 / (0.0 + -2.25))) / 3.14159))) / (v4 * ((-2.25 - v4) + -2.25))))) + 0.1))) * (v4 / 3.14159))) && true))) {
                    var v13 float64 = v4
                    _ = v13
                }
                if true {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        if true {
                            var v15 bool = true
                            var v16 uint8 = 255
                            _ = v15
                            _ = v16
                        }
                        _ = fi14
                    }
                } else {
                    switch ("hello" + "hello") {
                    case "hello":
                        v17 := v1[""]
                        _ = v17
                    case "\n\t":
                        fmt.Println(v4, v2)
                    }
                    if true {
                        v18 := GenS1{F0: 42, F1: ("a" + "a"), F2: false}
                        switch ("" + (("\n\t" + (("a" + "") + "hello")) + "a")) {
                        case "\n\t":
                            v19 := GenS1{F0: 4294967295, F1: "a", F2: true}
                            _ = v19
                        case "hello":
                            var v20 uint16 = 0
                            _ = v20
                        default:
                            v21 := v3[""]
                            _ = v21
                        }
                        _ = v18
                    }
                    v4 = v4
                }
                v10 -= 42
                _ = fi12
            }
        case 2147483647:
            v4 += ((((1e10 + 3.14159) + v4) - (v4 - v4)) - v4)
        default:
            v22 := GenS2{F0: true}
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 14; fi23++ {
            if false {
                v24 := v3[(("" + "\n\t") + (("\n\t" + "\n\t") + ""))]
                for fi25 := 0; fi25 < 14; fi25++ {
                    var v26 int16 = v6
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
            delete(v3, ("foo bar" + ""))
            _ = fi23
        }
    }
    var v27 uint16 = 0
    switch "hello" {
    case "a":
        for fi28 := 0; fi28 < 5; fi28++ {
            v6 = v6
            _ = fi28
        }
    case "":
        var v29 int32 = v2
        _ = v29
    }
    v2 = int32(v9)
    var v30 string = (("a" + "") + ("" + ""))
    var v31 uint8 = 1
    v32 := []int{32767, 32767, 285}
    v31 = (v31 - (v31 >> ((42) & 63)))
    if ((v4 <= float64((float32(v4) + float32(3.14159)))) || false) {
        v31 += (v31 ^ v31)
    } else {
        if false {
            _ = len(v1)
        } else {
            if ((!false) || (!true)) {
                v31 = (v31 - 128)
            }
            if true {
                fmt.Println(v6, v6, v31)
                _ = len(v32)
            } else {
                v33 := GenS1{F0: 18446744073709551615, F1: v30, F2: true}
                var v34 uint64 = 42
                v2 = -2147483648
                _ = v33
                _ = v34
            }
            var v35 string = (v30 + v30)
            _ = v35
        }
        v10 += (v10 ^ 42)
        v36 := v1["foo bar"]
        _ = v36
    }
    fmt.Println(v9, v27)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
