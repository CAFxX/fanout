package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 string }

func main() {
    var v1 int = 1
    var v2 uint32 = 566
    var v3 int16 = -32768
    v4 := make(map[int]uint32)
    var v5 int32 = -2147483648
    switch (("hello" + "a") + "\n\t") {
    case "a":
        v6 := []int{(v1 >> ((18446744073709551615) & 63)), (-(v1)), v1}
        _ = v6
    case "foo bar":
        v7 := GenS1{F0: uint8(v2)}
        _ = v7
    default:
        if (!(448 != 18446744073709551615)) {
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 uint16 = 42
                _ = fi8
                _ = v9
            }
            var v10 int32 = v5
            v11 := GenS2{F0: v3}
            _ = v10
            _ = v11
        }
    }
    var v12 uint8 = 255
    v13 := GenS2{F0: -32768}
    v5 -= 1
    v1 = v1
    fmt.Println(v12, v12, v3)
    if (!false) {
        fmt.Println(v12, v12)
        if false {
            if (32767 == v3) {
                var v14 float64 = (float64(-2.25) + float64(float32((1.5 - (1.5 - (3.14159 - ((-2.25 + 0.1) + 1e10)))))))
                switch (("\n\t" + "hello") + ("foo bar" + ("foo bar" + "a"))) {
                case "foo bar":
                    if false {
                        var v15 bool = true
                        if v15 {
                            v12 = 127
                        } else {
                            var v16 bool = false
                            _ = v16
                        }
                        v1 = v1
                        _ = v15
                    }
                case "\n\t":
                    if true {
                        if (-128 > int8(v1)) {
                            for fi17 := 0; fi17 < 6; fi17++ {
                                switch fi17 {
                                case 176:
                                    if (((v12 < v12) && false) && false) {
                                        var v18 string = "\n\t"
                                        _ = v18
                                    }
                                }
                                var v19 int = (-(fi17))
                                _ = fi17
                                _ = v19
                            }
                            fmt.Println(v5, v2)
                        }
                        v20 := []int{v1}
                        var v21 int8 = 67
                        _ = v20
                        _ = v21
                    }
                case "":
                    if (false && true) {
                        var v22 uint = 149
                        v5 = (v5 % ((v5) | 1))
                        if (false || (float32(0.1) > 0.1)) {
                            var v23 float32 = 1e10
                            v2 += (v2 >> ((0) & 63))
                            for fi24 := 0; fi24 < 15; fi24++ {
                                for fi25 := 0; fi25 < 9; fi25++ {
                                    if (v3 < (v3 >> ((65535) & 63))) {
                                        v26 := []int{}
                                        v27 := GenS3{F0: "a", F1: (("foo bar" + ("" + ("hello" + "a"))) + ("foo bar" + "\n\t"))}
                                        var v28 float32 = (float32(float64(0.1)) + 1.5)
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                    v29 := []int{(fi24 ^ fi25), (fi25 - (fi24 * v1)), fi25}
                                    _ = fi25
                                    _ = v29
                                }
                                fmt.Println(fi24, v22, v1)
                                v14 += 1e10
                                _ = fi24
                            }
                            _ = v23
                        }
                        _ = v22
                    }
                }
                _ = v14
            }
            var v30 float64 = ((-2.25 / 1.0) * float64(1e10))
            switch (v1 | 32768) {
            case 221:
                fmt.Println(v30, v5)
            case 0:
                var v31 int16 = v3
                _ = v31
            case -191:
                var v32 int16 = (v3 ^ (^(v3)))
                _ = v32
            }
            _ = v30
        }
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        var v34 int16 = 42
        v35 := GenS3{F0: "hello", F1: "a"}
        v36 := v4[v1]
        _ = fi33
        _ = v34
        _ = v35
        _ = v36
    }
    v3 = (v3 - (v3 + (v3 / ((32767) | 1))))
    var v37 int16 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
