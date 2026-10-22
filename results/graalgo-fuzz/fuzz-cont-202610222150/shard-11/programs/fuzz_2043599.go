package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() int {
    var v4 float64 = 3.14159
    var v5 int = -9223372036854775808
    fmt.Println(v4, v4)
    var v6 bool = (42 >= 42)
    _ = v4
    _ = v5
    _ = v6
    return (^(v5))
}
func genFn1() int64 {
    var v7 int64 = -9223372036854775808
    var v8 uint = 18446744073709551615
    _ = v7
    _ = v8
    return (v7 / (((v7 + (v7 & (v7 << ((18446744073709551615) & 63))))) | 1))
}

func main() {
    var v1 int32 = 555
    v2 := make(map[int]int)
    v3 := []string{"a"}
    for fi9 := 0; fi9 < 2; fi9++ {
        if false {
            if ((1 <= -682) || (!(1 >= 2147483647))) {
                for fi10 := 0; fi10 < 7; fi10++ {
                    switch genFn0() {
                    case -372:
                        var v11 uint64 = 4294967295
                        _ = v11
                    case 1:
                        v1 = v1
                    default:
                        var v12 int16 = 0
                        _ = v12
                    }
                    _ = len(v2)
                    _ = fi10
                }
                if false {
                    if ((828 <= (fi9 ^ genFn0())) || ((!true) && true)) {
                        var v13 float32 = -2.25
                        _ = v13
                    } else {
                        if false {
                            delete(v2, fi9)
                            v14 := GenS1{F0: 2147483647}
                            _ = v14
                        }
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                    }
                    v1 = v1
                    if (-9223372036854775808 == 0) {
                        for fi15 := 0; fi15 < 8; fi15++ {
                            var v16 uint16 = 32767
                            if ((v16 / (((v16 & 1)) | 1)) >= (v16 << ((65535) & 63))) {
                                var v17 uint8 = 255
                                v1 -= -2147483648
                                _ = v17
                            }
                            _ = fi15
                            _ = v16
                        }
                        v18 := GenS1{F0: 905}
                        var v19 uint = 65535
                        _ = v18
                        _ = v19
                    }
                } else {
                    v1 = (v1 + (v1 >> ((2) & 63)))
                }
            }
        }
        v20 := GenS1{F0: genFn0()}
        _ = len(v2)
        _ = fi9
        _ = v20
    }
    if len(v3) > 0 {
        v3[0] = (("" + "\n\t") + "")
    }
    v1 = 56
    fmt.Println(v1)
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 int16 = 0
        _ = fi21
        _ = v22
    }
    v1 = v1
    var v23 bool = true
    var v24 uint16 = 970
    var v25 string = "a"
    v1 = v1
    switch v25 {
    case "":
        if v23 {
            var v26 float32 = (1.5 + (1e10 - 1e10))
            var v27 float64 = 1e10
            _ = v26
            _ = v27
        } else {
            switch v25 {
            case "hello":
                switch 4294967295 {
                case 829:
                    fmt.Println(v24)
                }
            case "a":
                v3 = append(v3, ((v25 + v25) + v25))
            }
        }
    }
    v28 := GenS1{F0: genFn0()}
    if (128 >= 255) {
        v24 = v24
    } else {
        fmt.Println(v24, v1, v24)
        var v29 float32 = float32(-2.25)
        v23 = (2 != 965)
        _ = v29
    }
    v2[-244] = 471
    v25 = ("hello" + v25)
    for fi30 := 0; fi30 < 5; fi30++ {
        v25 = ("foo bar" + "foo bar")
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
}
