package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() int32 {
    var v7 float32 = 3.14159
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint16 = 255
        var v10 uint8 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(fi11, fi11)
        for fi12 := 0; fi12 < 7; fi12++ {
            v7 = v7
            var v13 int16 = 32767
            v14 := []int{fi12, 4294967295}
            fmt.Println(fi12, fi12, fi11)
            var v15 uint64 = 18446744073709551615
            fmt.Println(fi11, v13)
            v16 := GenS1{F0: 396}
            v17 := GenS1{F0: 1}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi11
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := []int{fi18, 127}
        _ = fi18
        _ = v19
    }
    _ = v7
    return 0
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]string)
    v3 := []string{("" + (("\n\t" + "foo bar") + "hello")), "", "foo bar"}
    v4 := []string{("a" + "hello")}
    v5 := []string{"hello", "hello", "a", "a"}
    var v6 int64 = 42
    v6 -= 37
    delete(v2, "")
    switch -241 {
    case 1:
        var v20 int8 = -45
        _ = v20
    case -478:
        v3 = append(v3, "a")
    case 9223372036854775807:
        switch 1 {
        case -823:
            fmt.Println(v1, v1)
        case -760:
            for fi21 := 0; fi21 < 16; fi21++ {
                if ((v1 >> ((0) & 63)) > v1) {
                    if true {
                        for fi22 := 0; fi22 < 7; fi22++ {
                            v23 := v2["a"]
                            v24 := v2[("" + "a")]
                            var v25 float64 = 3.14159
                            _ = fi22
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    }
                } else {
                    var v26 uint16 = 0
                    for fi27 := 0; fi27 < 11; fi27++ {
                        var v28 int = fi27
                        _ = fi27
                        _ = v28
                    }
                    v29 := GenS1{F0: 137}
                    _ = v26
                    _ = v29
                }
                _ = fi21
            }
        default:
            if true {
                for fi30 := 0; fi30 < 4; fi30++ {
                    var v31 float32 = (0.1 * (((0.1 * 3.14159) * (-2.25 / -2.25)) / 1.0))
                    _ = fi30
                    _ = v31
                }
                v32 := GenS1{F0: 1}
                _ = v32
            }
        }
    default:
        fmt.Println(v6)
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        fmt.Println(fi33, v6)
        _ = fi33
    }
    var v34 uint8 = 255
    var v35 string = "hello"
    v36 := GenS1{F0: 42}
    for fi37 := 0; fi37 < 12; fi37++ {
        v38 := GenS1{F0: 42}
        _ = fi37
        _ = v38
    }
    var v39 int = -246
    for fi40 := 0; fi40 < 15; fi40++ {
        for fi41 := 0; fi41 < 7; fi41++ {
            v2[""] = "\n\t"
            var v42 uint64 = 42
            var v43 float32 = 0.0
            _ = fi41
            _ = v42
            _ = v43
        }
        var v44 int = (^(fi40))
        if (65535 == 2147483648) {
            v39 = v39
            for fi45 := 0; fi45 < 16; fi45++ {
                if (false && (((genFn0() <= 0) && false) && (false || ((!(false || true)) && false)))) {
                    v6 += v6
                } else {
                    if false {
                        v34 -= v34
                        v39 -= (v44 << ((4294967295) & 63))
                    } else {
                        var v46 float64 = ((1e10 + 1e10) / ((3.14159 / ((1e10 - 1.5) + 0.0)) - float64((1e10 + 1.5))))
                        _ = v46
                    }
                    var v47 float32 = 1.5
                    v47 = -2.25
                    var v48 int64 = v6
                    _ = v47
                    _ = v48
                }
                _ = fi45
            }
        } else {
            var v49 bool = (2 >= v34)
            _ = v49
        }
        _ = fi40
        _ = v44
    }
    switch v35 {
    case "a":
        var v50 float64 = 1.5
        _ = v50
    case "foo bar":
        fmt.Println(v35, v1, v35)
    }
    for fi51 := 0; fi51 < 3; fi51++ {
        var v52 int64 = -973
        var v53 int32 = -935
        _ = fi51
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 19; fi54++ {
        switch v35 {
        case "":
            for fi55 := 0; fi55 < 20; fi55++ {
                if len(v5) > 0 {
                    v5[0] = v35
                }
                var v56 uint8 = (v34 * v34)
                var v57 float32 = float32(0.0)
                _ = fi55
                _ = v56
                _ = v57
            }
        }
        var v58 int = -2147483648
        var v59 float64 = -2.25
        _ = fi54
        _ = v58
        _ = v59
    }
    v1 = 155
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
