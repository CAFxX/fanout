package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 string }
func genFn0(p0 int8, p1 int, p2 uint) uint64 {
    p1 += (p1 % (((-(p1))) | 1))
    v6 := GenS1{F0: -1, F1: 2, F2: ("foo bar" + ("a" + "hello"))}
    var v7 string = (("a" + ("a" + ("\n\t" + "foo bar"))) + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 string = "a"
    var v2 uint32 = 2147483648
    var v3 float32 = (float32(1e10) + (1.5 + 0.0))
    v4 := []string{(("hello" + v1) + (v1 + (("foo bar" + "") + v1)))}
    var v5 uint32 = 533
    var v8 float32 = (((0.0 * v3) + ((v3 - v3) / 1e10)) / v3)
    if len(v4) > 0 {
        v4[0] = (v1 + "a")
    }
    v2 -= (v5 | 42)
    for fi9 := 0; fi9 < 20; fi9++ {
        fmt.Println(fi9, v1, v8)
        _ = fi9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 float64 = 3.14159
        v2 += (v2 << ((18446744073709551615) & 63))
        for fi12 := 0; fi12 < 20; fi12++ {
            if (!false) {
                switch (fi10 | (fi12 * (-(fi12)))) {
                case -9223372036854775808:
                    var v13 int = (fi12 & fi12)
                    _ = v13
                case 127:
                    if false {
                        if (!(1 == 0)) {
                            var v14 string = (v1 + "foo bar")
                            v15 := GenS1{F0: 0, F1: uint16(v5), F2: "hello"}
                            v16 := GenS1{F0: -32768, F1: 1, F2: v14}
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                        if true {
                            var v17 float64 = float64((((((v8 / v3) + float32((0.0 - (float64(((3.14159 + (v8 / v8)) - (float32(((3.14159 / 1e10) / 1.5)) / (((float32(v11) - 0.0) - float32(-2.25)) * 0.0)))) + v11)))) + (3.14159 * v3)) + v8) + float32(((0.0 + 1e10) - (v11 + (1e10 / (v11 - v11)))))))
                            var v18 int64 = 1
                            _ = v17
                            _ = v18
                        }
                    }
                }
                switch v1 {
                case "\n\t":
                    fmt.Println(fi10, fi10, v1)
                default:
                    v19 := GenS1{F0: -386, F1: 2, F2: v1}
                    _ = v19
                }
                v20 := GenS1{F0: 1, F1: 633, F2: (v1 + ("hello" + "a"))}
                _ = v20
            } else {
                switch (v1 + "") {
                case "a":
                    fmt.Println(fi10, v3)
                case "hello":
                    var v21 uint16 = 65535
                    _ = v21
                default:
                    var v22 int16 = 225
                    _ = v22
                }
                _ = len(v4)
                var v23 float32 = v3
                _ = v23
            }
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    var v24 uint16 = 2
    switch -9223372036854775808 {
    case -2147483648:
        _ = len(v4)
    }
    v2 = (v2 % (((^(v5))) | 1))
    if len(v4) > 0 {
        v4[0] = v1
    }
    var v25 uint64 = 4294967295
    v26 := GenS1{F0: 1000, F1: v24, F2: (v1 + "a")}
    var v27 float32 = v8
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 float64 = 0.1
        for fi30 := 0; fi30 < 2; fi30++ {
            var v31 float64 = (-2.25 - float64(3.14159))
            if (true && (v24 >= v24)) {
                fmt.Println(v3, fi30)
            }
            var v32 float32 = v3
            _ = fi30
            _ = v31
            _ = v32
        }
        v5 -= v5
        _ = fi28
        _ = v29
    }
    switch 305 {
    case 127:
        _ = len(v4)
    case -2147483648:
        v4 = append(v4, v1)
    case -9223372036854775808:
        var v33 int64 = -1
        _ = v33
    default:
        var v34 uint64 = (v25 >> ((65535) & 63))
        _ = v34
    }
    var v35 uint16 = v24
    var v36 bool = (v5 > (^(v2)))
    var v37 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v8
}
