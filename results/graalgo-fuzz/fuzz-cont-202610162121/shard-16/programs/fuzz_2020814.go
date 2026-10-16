package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int, p1 uint16) uint64 {
    v5 := GenS1{F0: (false && true)}
    for fi6 := 0; fi6 < 8; fi6++ {
        switch "\n\t" {
        case "hello":
            v7 := []int{fi6}
            _ = v7
        default:
            var v8 string = "a"
            _ = v8
        }
        var v9 int8 = 94
        _ = fi6
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    return 4294967295
}
func genFn1(p0 string) uint8 {
    p0 = p0
    switch -649 {
    case 32767:
        fmt.Println(p0)
    }
    v10 := []int{1, -944, -793}
    var v11 int8 = 0
    _ = p0
    _ = v10
    _ = v11
    return 128
}

func main() {
    var v1 float32 = float32(3.14159)
    var v2 int64 = -1
    v3 := make(map[string]string)
    var v4 string = "hello"
    if (498 <= -221) {
        v1 += v1
        switch 0 {
        case 242:
            v4 = ""
        case 4294967295:
            if ((v1 + -2.25) < (((0.0 * -2.25) + ((v1 - (1.5 - 1e10)) + v1)) / 3.14159)) {
                v4 = (v4 + (v4 + (v4 + v4)))
            } else {
                var v12 int32 = 2147483647
                _ = v12
            }
        default:
            v4 = ((v4 + (v4 + ("a" + v4))) + v4)
        }
    }
    var v13 string = "a"
    var v14 uint = 18446744073709551615
    v2 += (v2 << ((v14) & 63))
    v2 = 0
    v13 = v13
    var v15 int16 = 537
    switch -1 {
    case 128:
        var v16 int = int(v14)
        _ = v16
    case -1:
        v1 = (1e10 / v1)
    case 42:
        var v17 uint8 = 25
        _ = v17
    default:
        delete(v3, ("" + ((((v13 + "\n\t") + (((v4 + v13) + v13) + (v13 + v4))) + v13) + v4)))
    }
    v1 += v1
    var v18 int32 = -184
    if false {
        v2 = v2
    } else {
        if (v15 != 0) {
            var v19 uint32 = 0
            if true {
                var v20 uint = (v14 / (((v14 & (v14 + v14))) | 1))
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 float32 = 1e10
                    switch -9223372036854775808 {
                    case 32767:
                        v19 -= (^(v19))
                    default:
                        v14 = v14
                    }
                    v3[("a" + v4)] = (v4 + "a")
                    _ = fi21
                    _ = v22
                }
                _ = v20
            } else {
                v23 := []int{0}
                _ = v23
            }
            _ = v19
        }
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 uint16 = 65535
        for fi26 := 0; fi26 < 9; fi26++ {
            if false {
                var v27 bool = (!(false || true))
                v28 := GenS1{F0: (v4 == "a")}
                if (!((((2 != 0) && v27) || (fi26 >= fi26)) && v27)) {
                    v29 := v3[v4]
                    _ = v29
                }
                switch "" {
                case "":
                    v30 := []int{}
                    _ = v30
                case "foo bar":
                    v31 := v3[""]
                    _ = v31
                case "a":
                    v32 := v3[v13]
                    _ = v32
                }
                fmt.Println(v27)
                if (true || v27) {
                    v25 = uint16(v14)
                    var v33 int8 = -1
                    if false {
                        if (uint8(v14) <= 255) {
                            fmt.Println(v2)
                        }
                        v15 -= v15
                    }
                    _ = v33
                }
                _ = v27
                _ = v28
            } else {
                v34 := GenS1{F0: false}
                fmt.Println(v14, v1)
                _ = v34
            }
            v35 := []int{(fi26 * (fi26 % ((fi24) | 1)))}
            _ = fi26
            _ = v35
        }
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
