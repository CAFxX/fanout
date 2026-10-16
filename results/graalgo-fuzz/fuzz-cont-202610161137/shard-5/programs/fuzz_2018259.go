package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int, p1 uint) int32 {
    fmt.Println(p0, p1)
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            switch "\n\t" {
            case "foo bar":
                var v8 string = "\n\t"
                v9 := []int{p0, (-(fi6)), p0}
                _ = v8
                _ = v9
            case "hello":
                var v10 int = (fi7 ^ 127)
                _ = v10
            default:
                switch "a" {
                case "a":
                    switch (p0 % ((128) | 1)) {
                    case -1:
                        var v11 int64 = 290
                        if (-1 < int32(fi7)) {
                            v11 += v11
                            fmt.Println(p1, fi7, p0)
                            v12 := []int{fi7, 2147483647, fi7}
                            _ = v12
                        }
                        _ = v11
                    case 2147483647:
                        var v13 uint32 = 65535
                        _ = v13
                    default:
                        v14 := []int{(p0 / ((4294967295) | 1)), 4294967295, (p0 + p0), fi6}
                        _ = v14
                    }
                case "":
                    p1 += (p1 / ((4294967295) | 1))
                }
            }
            _ = fi7
        }
        _ = fi6
    }
    var v15 uint16 = 961
    _ = p0
    _ = p1
    _ = v15
    return -67
}

func main() {
    var v1 float32 = ((-2.25 + -2.25) * 1e10)
    var v2 bool = false
    var v3 uint8 = 130
    var v4 int16 = 714
    v5 := make(map[string]string)
    v4 = 42
    v4 += -53
    fmt.Println(v1, v2, v4)
    v2 = v2
    v3 -= v3
    if (v4 <= v4) {
        v16 := GenS1{F0: v2}
        var v17 string = "a"
        _ = v16
        _ = v17
    }
    v1 = v1
    v3 = v3
    var v18 float32 = 0.1
    switch ("a" + "") {
    case "\n\t":
        if (!(true && v2)) {
            v2 = v2
            switch 32767 {
            case 32767:
                delete(v5, "a")
            case 0:
                fmt.Println(v18, v3)
            default:
                delete(v5, "foo bar")
            }
        } else {
            switch "\n\t" {
            case "foo bar":
                v19 := GenS1{F0: (42 < 0)}
                _ = v19
            case "":
                var v20 string = "hello"
                _ = v20
            }
        }
    case "foo bar":
        v18 -= v18
    case "":
        var v21 int64 = -1
        _ = v21
    default:
        v22 := []int{}
        _ = v22
    }
    var v23 int16 = v4
    var v24 uint = 65535
    v24 = v24
    fmt.Println(v3, v18)
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 uint = (v24 >> (((v24 << ((v24) & 63))) & 63))
        if (!true) {
            v27 := GenS1{F0: false}
            _ = v27
        }
        _ = fi25
        _ = v26
    }
    switch "" {
    case "foo bar":
        if false {
            v3 -= v3
        }
    case "hello":
        v2 = (!v2)
    case "a":
        if true {
            var v28 uint = v24
            _ = v28
        } else {
            var v29 int64 = 1
            for fi30 := 0; fi30 < 11; fi30++ {
                v3 = v3
                var v31 uint32 = 0
                switch ("" + (("foo bar" + "a") + ("foo bar" + "foo bar"))) {
                case "hello":
                    v32 := GenS1{F0: (!(genFn0(fi30, v24) >= -1))}
                    _ = v32
                default:
                    v33 := GenS1{F0: v2}
                    _ = v33
                }
                _ = fi30
                _ = v31
            }
            v5[((("a" + ("" + (("" + "a") + "\n\t"))) + ("foo bar" + "\n\t")) + "hello")] = "foo bar"
            _ = v29
        }
    default:
        var v34 float64 = (-2.25 + (float64(v1) * float64(v1)))
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
