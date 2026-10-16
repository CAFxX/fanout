package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 uint8; F1 int; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0() int32 {
    var v5 string = "\n\t"
    var v6 uint8 = 127
    var v7 uint8 = v6
    _ = v5
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []string{}
    var v3 uint8 = 127
    var v4 int8 = 127
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    var v8 int = -2147483648
    fmt.Println(v1, v1, v3)
    v8 = v8
    var v9 float64 = (float64(float32(1e10)) / 0.1)
    var v10 int16 = 0
    var v11 int64 = int64(v3)
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: v4, F1: 587, F2: (("hello" + ("" + (("\n\t" + "") + "a"))) + ("hello" + ("\n\t" + "")))}
        if (!false) {
            v14 := GenS2{F0: 87, F1: -2147483648, F2: 486}
            _ = v14
        } else {
            if (623 >= 559) {
                v8 = fi12
            }
        }
        _ = fi12
        _ = v13
    }
    v15 := GenS2{F0: v3, F1: 1, F2: 4294967295}
    v10 += v10
    var v16 int16 = (v10 ^ v10)
    switch "a" {
    case "\n\t":
        v17 := GenS1{F0: v4, F1: 42, F2: "hello"}
        _ = v17
    case "foo bar":
        if true {
            fmt.Println(v9, v16, v10)
        } else {
            fmt.Println(v1, v16)
        }
    case "a":
        switch "foo bar" {
        case "a":
            switch v8 {
            case -198:
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 int8 = (v4 - -75)
                    if false {
                        v2 = append(v2, "a")
                        v20 := GenS1{F0: v4, F1: 1, F2: "\n\t"}
                        for fi21 := 0; fi21 < 7; fi21++ {
                            v22 := GenS3{F0: 4294967295, F1: v11}
                            v3 -= 2
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi21
                            _ = v22
                        }
                        _ = v20
                    } else {
                        v2 = append(v2, ("foo bar" + (("\n\t" + ("\n\t" + ((("\n\t" + ("a" + ("hello" + ("hello" + "\n\t")))) + "hello") + ("" + "hello")))) + (("hello" + ("\n\t" + "")) + ("hello" + "")))))
                        v4 -= (v4 << ((453) & 63))
                        var v23 float32 = (((-2.25 - -2.25) / -2.25) * 3.14159)
                        _ = v23
                    }
                    _ = fi18
                    _ = v19
                }
            }
        case "\n\t":
            var v24 int = (-(v8))
            _ = v24
        case "":
            fmt.Println(v16, v16, v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
