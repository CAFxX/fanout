package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 float32 = 3.14159
    var v2 string = "\n\t"
    v3 := make(map[string]uint32)
    var v4 uint64 = 0
    v5 := make(map[string]int)
    for fi6 := 0; fi6 < 11; fi6++ {
        v2 = (((((v2 + ("\n\t" + v2)) + ((("foo bar" + "") + v2) + (v2 + (v2 + "")))) + ((v2 + ((v2 + (("a" + ("\n\t" + "foo bar")) + v2)) + "a")) + "\n\t")) + v2) + (v2 + "foo bar"))
        _ = fi6
    }
    var v7 int64 = -9223372036854775808
    var v8 int8 = 21
    v8 = (v8 >> ((65535) & 63))
    v4 = (v4 % ((18446744073709551615) | 1))
    for fi9 := 0; fi9 < 6; fi9++ {
        v2 = v2
        var v10 uint8 = 255
        fmt.Println(v2, v2)
        _ = fi9
        _ = v10
    }
    if true {
        v2 = (("a" + v2) + "")
    } else {
        var v11 uint16 = 1
        var v12 uint8 = 0
        _ = v11
        _ = v12
    }
    var v13 uint8 = 255
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := []int{(-(fi14)), (-(fi14)), 32768, fi14}
            if (2147483647 > 2147483647) {
                v4 = 4294967295
            }
            v17 := GenS1{F0: ((3.14159 * (-2.25 / 1.0)) > (((1.5 * (1.5 * 0.0)) * 1.5) / -2.25))}
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 uint64 = v4
        var v19 uint64 = (v4 - 18446744073709551615)
        _ = fi14
        _ = v18
        _ = v19
    }
    if true {
        v20 := []int{0, 127, 9223372036854775807}
        for fi21 := 0; fi21 < 4; fi21++ {
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            delete(v3, v2)
            if false {
                for fi22 := 0; fi22 < 10; fi22++ {
                    v1 = 1e10
                    if (!false) {
                        var v23 float32 = v1
                        v2 = v2
                        v1 = v1
                        var v24 float64 = 0.1
                        _ = v23
                        _ = v24
                    } else {
                        if true {
                            var v25 uint8 = (^(v13))
                            v3["a"] = 2147483648
                            fmt.Println(v2)
                            _ = v25
                        }
                        fmt.Println(v8, v7)
                    }
                    if len(v20) > 0 {
                        v20[0] = fi22
                    }
                    v2 = "foo bar"
                    _ = fi22
                }
                fmt.Println(v2)
                v26 := GenS2{F0: true}
                _ = v26
            }
            _ = fi21
        }
        _ = v20
    } else {
        v27 := GenS3{F0: (v4 % ((v4) | 1))}
        _ = v27
    }
    fmt.Println(v13, v7)
    for fi28 := 0; fi28 < 3; fi28++ {
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 int8 = v8
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    var v31 string = ((("" + v2) + v2) + "a")
    var v32 uint = 112
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
