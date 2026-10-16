package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
func genFn0() int {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int64 = 1
        _ = fi7
        _ = v8
    }
    var v9 int8 = 42
    _ = v9
    return 32768
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 1
    var v3 int = -1
    var v4 uint64 = 4294967295
    v5 := make(map[int]uint32)
    var v6 uint32 = uint32(v2)
    if false {
        if false {
            v10 := []int{genFn0()}
            _ = v10
        } else {
            v2 = 65535
            for fi11 := 0; fi11 < 12; fi11++ {
                if true {
                    if false {
                        v4 = (v4 & (v4 ^ (^(v4))))
                        v6 = 1
                    }
                    fmt.Println(fi11, fi11)
                }
                _ = fi11
            }
            v12 := []int{}
            _ = v12
        }
        var v13 float64 = 0.0
        v14 := GenS1{F0: 0}
        _ = v13
        _ = v14
    } else {
        v3 = -2147483648
    }
    if true {
        v6 = 42
        var v15 uint = 18446744073709551615
        v16 := []int{}
        _ = v15
        _ = v16
    }
    v5[(v3 >> ((0) & 63))] = (v6 | (v6 % ((v6) | 1)))
    for fi17 := 0; fi17 < 20; fi17++ {
        v18 := GenS2{F0: 2147483647}
        v4 -= 18446744073709551615
        v4 = v4
        _ = fi17
        _ = v18
    }
    var v19 uint8 = 244
    fmt.Println(v6, v4)
    if (!(1.5 != (-2.25 * (float32(0.1) * -2.25)))) {
        v20 := GenS1{F0: (v19 + (v19 ^ (v19 >> ((42) & 63))))}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 16; fi21++ {
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 uint64 = 1
                var v24 uint32 = (v6 | v6)
                switch fi22 {
                case 32767:
                    v25 := []int{}
                    _ = v25
                case 4294967295:
                    v5[fi21] = (v24 / ((876) | 1))
                default:
                    if true {
                        v26 := GenS1{F0: v19}
                        if false {
                            v23 = uint64(v19)
                        }
                        _ = v26
                    } else {
                        if false {
                            v19 = v19
                            var v27 int8 = 100
                            _ = v27
                        }
                    }
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 int = v3
                var v30 uint = 516
                var v31 uint64 = v4
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi21
        }
        switch (v3 & (v3 % ((v3) | 1))) {
        case 503:
            var v32 int32 = -57
            _ = v32
        case -9223372036854775808:
            v33 := GenS2{F0: -832}
            _ = v33
        case 9223372036854775807:
            var v34 bool = false
            _ = v34
        default:
            if (65535 >= 2147483648) {
                fmt.Println(v19, v19, v2)
            } else {
                fmt.Println(v6)
                switch genFn0() {
                case -698:
                    var v35 bool = (215 > -2147483648)
                    _ = v35
                case -1:
                    v3 = v3
                default:
                    switch (v3 + 4294967295) {
                    case 127:
                        var v36 uint = 1
                        _ = v36
                    case -664:
                        var v37 int8 = 7
                        _ = v37
                    default:
                        var v38 uint32 = v6
                        _ = v38
                    }
                }
            }
        }
    }
    var v39 int = v3
    for fi40 := 0; fi40 < 4; fi40++ {
        var v41 string = ""
        fmt.Println(v4, v4, v6)
        _ = fi40
        _ = v41
    }
    v42 := []int{791, v3}
    switch 32767 {
    case -9223372036854775808:
        delete(v1, ((("a" + "foo bar") + ("foo bar" + "\n\t")) + "a"))
    case 127:
        var v43 int8 = -64
        _ = v43
    case 770:
        var v44 string = "\n\t"
        _ = v44
    default:
        var v45 int = (v3 ^ (v39 ^ v39))
        _ = v45
    }
    var v46 uint64 = v4
    delete(v1, "\n\t")
    v46 += (v46 % (((v46 ^ (v46 / (((v46 >> ((0) & 63))) | 1)))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v46
    _ = v5
    _ = v6
}
