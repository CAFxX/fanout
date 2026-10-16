package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int16 }
func genFn0(p0 bool) uint32 {
    v7 := []int{-315, -81, 4294967295, 166}
    var v8 int32 = -1
    var v9 uint32 = 611
    fmt.Println(v8)
    var v10 float64 = (1.5 - 0.0)
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return v9
}
func genFn1(p0 bool, p1 int32, p2 int64) uint16 {
    for fi11 := 0; fi11 < 18; fi11++ {
        switch "hello" {
        case "a":
            var v12 int = fi11
            _ = v12
        case "\n\t":
            p2 = p2
        }
        fmt.Println(p2, p2)
        v13 := GenS1{F0: "\n\t"}
        _ = fi11
        _ = v13
    }
    fmt.Println(p2, p1)
    switch 128 {
    case 32767:
        v14 := GenS1{F0: "foo bar"}
        _ = v14
    case 4294967295:
        p0 = true
    case 2147483647:
        v15 := GenS1{F0: "\n\t"}
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := []int{(-(fi16)), (fi16 >> ((905) & 63)), fi16}
        _ = fi16
        _ = v17
    }
    var v18 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 32767
}

func main() {
    var v1 int64 = -466
    var v2 uint = 655
    var v3 int = 909
    v4 := []string{(("\n\t" + ("a" + "")) + "\n\t"), "foo bar", (("" + "foo bar") + ((("" + ("\n\t" + (("foo bar" + "a") + ("\n\t" + "hello")))) + "\n\t") + ("foo bar" + "")))}
    v5 := make(map[int]uint32)
    var v6 bool = false
    var v19 float32 = 3.14159
    switch v3 {
    case 127:
        switch -1 {
        case -476:
            fmt.Println(v6, v6)
        case 1:
            v5[v3] = 2147483647
        }
    case 128:
        fmt.Println(v1)
    case -122:
        var v20 int32 = 498
        _ = v20
    default:
        var v21 int8 = 121
        _ = v21
    }
    fmt.Println(v19, v3, v6)
    v22 := GenS2{F0: 133, F1: -931}
    v5[4294967295] = 2147483647
    v19 -= ((1.5 + v19) / float32((0.1 / (1e10 - -2.25))))
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        v24 := v5[v3]
        v19 -= v19
        if (0 >= 32767) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v1 -= 1
            _ = len(v5)
        } else {
            var v25 float32 = 0.1
            if ((!(true || ("" < ("hello" + ((("foo bar" + "") + "foo bar") + "foo bar"))))) && (!true)) {
                switch "" {
                case "hello":
                    var v26 uint = (v2 - 276)
                    _ = v26
                default:
                    _ = len(v4)
                }
                for fi27 := 0; fi27 < 10; fi27++ {
                    v28 := GenS2{F0: 1, F1: -32768}
                    switch 128 {
                    case -850:
                        v29 := GenS1{F0: "a"}
                        _ = v29
                    default:
                        fmt.Println(v25)
                    }
                    _ = fi27
                    _ = v28
                }
            } else {
                var v30 string = "hello"
                _ = v30
            }
            var v31 uint = 65535
            _ = v25
            _ = v31
        }
        _ = fi23
        _ = v24
    }
    v6 = ((36 < 60) && (genFn1((255 != (v2 / (((v2 | (v2 ^ (v2 | 255)))) | 1))), 144, (^(v1))) >= 255))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
