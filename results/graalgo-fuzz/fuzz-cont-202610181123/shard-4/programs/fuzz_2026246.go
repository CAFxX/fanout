package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
func genFn0() bool {
    var v6 string = (("hello" + "foo bar") + "hello")
    v7 := []int{-2147483648, -836, 4294967295, 905}
    var v8 uint64 = 8
    v8 -= (v8 | v8)
    var v9 float64 = (-2.25 / 1e10)
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return ((v9 != (v9 - 1e10)) && (v8 == (v8 << ((2) & 63))))
}
func genFn1() uint32 {
    var v10 int = -9223372036854775808
    fmt.Println(v10)
    v10 = (v10 | (v10 << ((726) & 63)))
    v10 -= 128
    v10 += v10
    _ = v10
    return 2147483647
}

func main() {
    var v1 uint32 = 2
    var v2 uint = 65535
    v3 := []string{"foo bar"}
    var v4 bool = true
    v5 := make(map[uint32]bool)
    switch 1 {
    case 42:
        delete(v5, v1)
    case 2147483647:
        if false {
            for fi11 := 0; fi11 < 16; fi11++ {
                v5[v1] = (v4 && v4)
                v4 = v4
                if v4 {
                    switch (("\n\t" + "\n\t") + "foo bar") {
                    case "":
                        switch (fi11 >> ((v2) & 63)) {
                        case 42:
                            var v12 int8 = 1
                            _ = v12
                        case 0:
                            for fi13 := 0; fi13 < 12; fi13++ {
                                v2 = 2
                                _ = fi13
                            }
                        case -462:
                            var v14 string = "a"
                            _ = v14
                        }
                    default:
                        var v15 int64 = -796
                        _ = v15
                    }
                    var v16 uint16 = 644
                    _ = v16
                } else {
                    var v17 int = (fi11 / ((fi11) | 1))
                    switch (("a" + "\n\t") + ("hello" + "\n\t")) {
                    case "\n\t":
                        if v4 {
                            v17 -= (v17 ^ 9223372036854775807)
                        } else {
                            var v18 string = "\n\t"
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            for fi19 := 0; fi19 < 10; fi19++ {
                                v20 := GenS1{F0: 20, F1: 1}
                                v21 := GenS1{F0: 1, F1: -32768}
                                fmt.Println(fi11, v18)
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            _ = v18
                        }
                    case "foo bar":
                        v1 = 0
                    case "a":
                        v22 := GenS1{F0: 128, F1: 538}
                        _ = v22
                    default:
                        var v23 uint32 = v1
                        _ = v23
                    }
                    _ = v17
                }
                _ = fi11
            }
            v24 := GenS1{F0: 128, F1: 32767}
            _ = v24
        }
    default:
        v25 := GenS1{F0: 91, F1: 0}
        _ = v25
    }
    v26 := GenS1{F0: 0, F1: 32767}
    for fi27 := 0; fi27 < 20; fi27++ {
        fmt.Println(fi27, fi27, fi27)
        switch fi27 {
        case 42:
            v28 := GenS1{F0: 255, F1: 193}
            _ = v28
        case 410:
            v2 = v2
        default:
            if false {
                v2 += 2
            }
        }
        v4 = v4
        _ = fi27
    }
    v1 -= (^(v1))
    var v29 string = "\n\t"
    v30 := GenS2{F0: 330}
    var v31 int32 = -2147483648
    var v32 int64 = 1
    if (((v32 ^ 9223372036854775807) <= (v32 ^ 1)) || ((!((float64((-2.25 / 0.1)) - 1.5) == 1e10)) || true)) {
        v32 = (v32 >> ((v2) & 63))
        switch -1 {
        case -410:
            var v33 uint8 = 42
            _ = v33
        case 0:
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 uint32 = v1
                fmt.Println(v1, v4)
                _ = fi34
                _ = v35
            }
        default:
            v36 := GenS1{F0: 127, F1: 1}
            _ = v36
        }
        v1 += 0
    } else {
        v2 = uint(v1)
        v37 := GenS3{F0: 618}
        _ = v37
    }
    var v38 uint64 = 170
    var v39 int32 = v31
    var v40 float32 = 1.5
    v2 -= v2
    fmt.Println(v2, v39)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
