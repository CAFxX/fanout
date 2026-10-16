package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint32) uint8 {
    v4 := []int{9223372036854775807, 9223372036854775807, -9223372036854775808, 127}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v4
    return 2
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]bool)
    var v3 uint32 = 65535
    if false {
        v2[v3] = ("" != "a")
        v5 := GenS1{F0: genFn0(-2147483648, (v3 >> ((1) & 63)))}
        _ = v5
    } else {
        var v6 bool = true
        _ = v6
    }
    var v7 int8 = -1
    var v8 uint16 = 328
    if (((-2.25 * (0.1 * -2.25)) * (0.0 * (3.14159 / 1e10))) <= float32(-2.25)) {
        if (-1 < 1) {
            if (!(128 == int(v7))) {
                var v9 int8 = (v7 << ((65535) & 63))
                v1 += v1
                _ = v9
            } else {
                v10 := []int{}
                fmt.Println(v1, v1)
                var v11 float32 = 1e10
                _ = v10
                _ = v11
            }
            v12 := GenS1{F0: genFn0(-2147483648, 42)}
            _ = v12
        } else {
            v7 -= 100
            switch (("" + "a") + "hello") {
            case "a":
                if true {
                    switch ((("a" + "\n\t") + "") + "") {
                    case "\n\t":
                        v13 := GenS2{F0: 2147483648}
                        _ = v13
                    case "foo bar":
                        var v14 string = "a"
                        _ = v14
                    }
                } else {
                    v15 := GenS1{F0: 127}
                    if (694 >= (v3 - 42)) {
                        var v16 uint64 = 456
                        if true {
                            if true {
                                fmt.Println(v1, v16)
                            } else {
                                v1 = (v1 * -32768)
                            }
                            var v17 int = 9223372036854775807
                            _ = v17
                        } else {
                            switch 128 {
                            case 128:
                                v18 := v2[1]
                                _ = v18
                            case 814:
                                v16 += (v16 * v16)
                            default:
                                for fi19 := 0; fi19 < 14; fi19++ {
                                    v20 := GenS2{F0: uint32(v8)}
                                    _ = fi19
                                    _ = v20
                                }
                            }
                            var v21 int64 = 0
                            _ = v21
                        }
                        if (0.1 == -2.25) {
                            var v22 uint8 = 127
                            _ = v22
                        } else {
                            switch -884 {
                            case 42:
                                if (((false || ((-1 > -9223372036854775808) || false)) || (!true)) && true) {
                                    v23 := GenS2{F0: (v3 / ((v3) | 1))}
                                    for fi24 := 0; fi24 < 2; fi24++ {
                                        switch (-(fi24)) {
                                        case 127:
                                            v25 := GenS3{F0: 1, F1: -622}
                                            _ = v25
                                        case 32768:
                                            var v26 int = fi24
                                            _ = v26
                                        case 2147483647:
                                            fmt.Println(v16, v16, v8)
                                        }
                                        _ = fi24
                                    }
                                    _ = v23
                                } else {
                                    v27 := v2[(^(v3))]
                                    v7 += v7
                                    _ = v27
                                }
                            case -382:
                                var v28 float32 = 3.14159
                                _ = v28
                            default:
                                fmt.Println(v3)
                            }
                        }
                        _ = v16
                    }
                    _ = v15
                }
            }
            v29 := GenS4{F0: v8}
            _ = v29
        }
    }
    var v30 uint64 = 2
    if (!((uint16(v30) != (v8 + v8)) || true)) {
        v30 += 42
    } else {
        v8 = v8
        if (uint64(v8) != v30) {
            switch ("foo bar" + "a") {
            case "foo bar":
                v31 := GenS4{F0: (v8 >> ((108) & 63))}
                _ = v31
            case "a":
                var v32 uint = 42
                _ = v32
            case "hello":
                delete(v2, (v3 * 751))
            default:
                v3 -= v3
            }
        } else {
            v30 = (^(v30))
            v33 := []int{543}
            v7 = v7
            _ = v33
        }
        for fi34 := 0; fi34 < 8; fi34++ {
            var v35 uint32 = uint32(v8)
            v36 := []int{fi34, -9223372036854775808, (fi34 << ((523) & 63)), fi34}
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    var v37 int16 = (-(v1))
    v30 = v30
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v7
    _ = v8
}
