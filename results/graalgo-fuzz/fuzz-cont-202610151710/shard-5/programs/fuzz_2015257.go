package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint16, p1 uint8) uint8 {
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 uint = 65535
        switch "" {
        case "a":
            var v6 uint64 = 2
            _ = v6
        case "hello":
            fmt.Println(p0, p1, p0)
        case "foo bar":
            var v7 uint = v5
            _ = v7
        }
        v5 += (^(v5))
        _ = fi4
        _ = v5
    }
    switch 528 {
    case 9223372036854775807:
        var v8 float64 = 1e10
        _ = v8
    case 0:
        if (true && true) {
            var v9 string = ("" + ("" + "\n\t"))
            _ = v9
        } else {
            var v10 uint8 = p1
            _ = v10
        }
    }
    _ = p0
    _ = p1
    return uint8(p0)
}
func genFn1(p0 int, p1 uint, p2 int64) string {
    v11 := []int{(p0 * (p0 << ((p1) & 63))), p0, p0}
    fmt.Println(p0)
    switch "foo bar" {
    case "\n\t":
        var v12 int32 = 1
        _ = v12
    default:
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return "hello"
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 3.14159
    var v3 uint16 = 0
    v1["foo bar"] = ("foo bar" + ("a" + ("\n\t" + "hello")))
    var v13 int8 = -1
    fmt.Println(v13, v3, v2)
    v1[""] = (("\n\t" + ((("" + ("\n\t" + "")) + "hello") + "")) + "a")
    var v14 float32 = ((0.0 * 0.1) - (((0.0 + 0.0) / 1.5) / -2.25))
    if false {
        v15 := []int{-149, -230, 691, 701}
        if (-534 == -1) {
            if ((("" + "foo bar") + ((("foo bar" + (("\n\t" + "\n\t") + "foo bar")) + "hello") + "hello")) == "foo bar") {
                for fi16 := 0; fi16 < 12; fi16++ {
                    v17 := GenS1{F0: 127}
                    for fi18 := 0; fi18 < 16; fi18++ {
                        var v19 uint32 = 2147483648
                        var v20 uint64 = 4294967295
                        var v21 int = int(v19)
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    _ = fi16
                    _ = v17
                }
            } else {
                var v22 int64 = 9223372036854775807
                _ = v22
            }
            for fi23 := 0; fi23 < 13; fi23++ {
                v2 = v2
                _ = fi23
            }
        } else {
            if true {
                for fi24 := 0; fi24 < 2; fi24++ {
                    v13 = int8(v3)
                    fmt.Println(v14, v14, v13)
                    _ = fi24
                }
                if false {
                    fmt.Println(v13, v2, v14)
                } else {
                    fmt.Println(v14, v13)
                    fmt.Println(v3, v13, v2)
                    var v25 uint32 = 65535
                    _ = v25
                }
            } else {
                fmt.Println(v14, v2, v3)
            }
        }
        _ = v15
    }
    delete(v1, ("\n\t" + "\n\t"))
    fmt.Println(v2, v2, v2)
    v26 := v1[(("a" + "foo bar") + "a")]
    v27 := GenS1{F0: 163}
    var v28 uint32 = 65535
    for fi29 := 0; fi29 < 10; fi29++ {
        delete(v1, ((v26 + (v26 + ("hello" + v26))) + (v26 + "foo bar")))
        v28 = (v28 + v28)
        for fi30 := 0; fi30 < 10; fi30++ {
            var v31 uint = 0
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    v2 -= (v2 / -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
