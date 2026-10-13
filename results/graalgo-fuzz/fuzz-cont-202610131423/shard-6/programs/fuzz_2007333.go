package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int8 }
type GenS2 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{321, 205, 2147483648}
    v3 := make(map[string]uint32)
    v4 := []int{1, 127, 1, 32768}
    var v5 int64 = -908
    var v6 int = -9223372036854775808
    fmt.Println(v5, v6)
    switch "a" {
    case "a":
        for fi7 := 0; fi7 < 17; fi7++ {
            v5 = v5
            _ = fi7
        }
    case "foo bar":
        v5 += v5
    case "hello":
        if len(v4) > 0 {
            v4[0] = (v6 + 2147483647)
        }
    }
    var v8 string = "foo bar"
    for fi9 := 0; fi9 < 6; fi9++ {
        v6 = (v6 | (fi9 ^ -9223372036854775808))
        _ = fi9
    }
    v6 -= (v6 ^ v6)
    v10 := GenS1{F0: 255, F1: 65535, F2: 100}
    var v11 int = v6
    if (false || false) {
        if (v8 <= (v8 + "\n\t")) {
            var v12 int = 324
            switch 9223372036854775807 {
            case 839:
                v13 := GenS1{F0: 255, F1: 513, F2: 19}
                _ = v13
            case 313:
                v1[(v6 * v12)] = (v8 + "foo bar")
            }
            _ = v12
        } else {
            v4 = append(v4, 128)
        }
        v6 -= v11
    } else {
        var v14 float32 = 0.0
        var v15 string = v8
        var v16 int32 = 1
        _ = v14
        _ = v15
        _ = v16
    }
    v11 = (v11 % ((-857) | 1))
    for fi17 := 0; fi17 < 4; fi17++ {
        v4 = append(v4, fi17)
        if (1e10 != 0.1) {
            if (436 != -1) {
                v11 = v11
                v11 += 4294967295
                for fi18 := 0; fi18 < 20; fi18++ {
                    if (!(!(42 > -32768))) {
                        var v19 int64 = (v5 / (((v5 << ((65535) & 63))) | 1))
                        _ = v19
                    }
                    _ = fi18
                }
            } else {
                if ((v11 == fi17) && true) {
                    var v20 uint8 = 0
                    _ = v20
                } else {
                    v21 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: -1}
                    if (1e10 <= (((0.0 + 1.5) - 1.5) - 0.0)) {
                        v3[v8] = 884
                        v4 = append(v4, (fi17 * v6))
                    } else {
                        var v22 string = ((("foo bar" + v8) + ((v8 + "a") + v8)) + (((("foo bar" + "hello") + "hello") + (v8 + (v8 + v8))) + v8))
                        var v23 float32 = -2.25
                        _ = v22
                        _ = v23
                    }
                    _ = v21
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    v11 += -1
                    v6 -= 4294967295
                    var v25 float32 = 1e10
                    _ = fi24
                    _ = v25
                }
            }
            for fi26 := 0; fi26 < 9; fi26++ {
                var v27 bool = false
                _ = fi26
                _ = v27
            }
        } else {
            for fi28 := 0; fi28 < 4; fi28++ {
                var v29 bool = true
                var v30 float64 = float64(3.14159)
                switch ("a" + ("" + (((v8 + (v8 + ("hello" + "\n\t"))) + v8) + ""))) {
                case "foo bar":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "a":
                    var v31 uint8 = 255
                    _ = v31
                case "":
                    var v32 int = (v11 * 2147483647)
                    _ = v32
                }
                v11 = (v6 << ((0) & 63))
                _ = fi28
                _ = v29
                _ = v30
            }
            if (!(!true)) {
                v33 := GenS2{F0: false, F1: -59}
                v34 := GenS2{F0: false, F1: -386}
                v11 = v6
                _ = v33
                _ = v34
            }
        }
        var v35 int64 = (v5 % ((-550) | 1))
        _ = fi17
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
